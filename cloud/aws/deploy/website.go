// Copyright Nitric Pty Ltd.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package deploy

import (
	"fmt"
	"io/fs"
	"math"
	"mime"
	"path/filepath"
	"runtime"
	"slices"
	"sort"
	"strings"

	"github.com/nitrictech/nitric/cloud/aws/deploy/embeds"
	"github.com/nitrictech/nitric/cloud/common/deploy/resources"
	common "github.com/nitrictech/nitric/cloud/common/deploy/tags"
	deploymentspb "github.com/nitrictech/nitric/core/pkg/proto/deployments/v1"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/cloudfront"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
	"github.com/pulumi/pulumi-command/sdk/go/command/local"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/samber/lo"
)

type website struct {
	bucket   *s3.Bucket
	basePath string
}

// Website - Implements the Website deployment method for the AWS provider
func (a *NitricAwsPulumiProvider) Website(ctx *pulumi.Context, parent pulumi.Resource, name string, config *deploymentspb.Website) error {
	opts := []pulumi.ResourceOption{pulumi.Parent(parent)}

	localDir, ok := config.AssetSource.(*deploymentspb.Website_LocalDirectory)
	if !ok {
		return fmt.Errorf("unsupported asset source type for website: %s", name)
	}

	cleanedPath := filepath.ToSlash(filepath.Clean(localDir.LocalDirectory))

	if config.BasePath == "/" {
		a.websiteIndexDocument = config.IndexDocument
		a.websiteErrorDocument = config.ErrorDocument
	}

	websiteBucketName := fmt.Sprintf("%s-website-bucket", name)
	websiteBucket, err := s3.NewBucket(ctx, websiteBucketName, &s3.BucketArgs{
		Tags: pulumi.ToStringMap(common.Tags(a.StackId, websiteBucketName, resources.Website)),
	})
	if err != nil {
		return err
	}

	// Enumerate the public directory in pwd and upload all files to the public bucket
	// This will be the source for our cloudfront distribution
	err = filepath.WalkDir(cleanedPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// Skip directories
		if d.IsDir() {
			return nil
		}

		// Get file info to check for special types
		info, err := d.Info()
		if err != nil {
			return err
		}

		// Skip non-regular files (e.g., symlinks, sockets, devices)
		if info.Mode()&fs.ModeType != 0 {
			return nil
		}

		// Determine the content type based on the file extension
		contentType := mime.TypeByExtension(filepath.Ext(path))
		if contentType == "" {
			contentType = "application/octet-stream"
		}

		// Generate the object key to include the folder structure
		var objectKey string

		filePath := path[len(cleanedPath):]

		arn := filepath.ToSlash(filepath.Join(name, filePath))

		// If the base path is not the root, include it in the object key
		objectKey = filepath.ToSlash(filePath)

		obj, err := s3.NewBucketObject(ctx, arn, &s3.BucketObjectArgs{
			Bucket:      websiteBucket.Bucket,
			Source:      pulumi.NewFileAsset(path),
			ContentType: pulumi.String(contentType),
			Key:         pulumi.String(objectKey),
		}, opts...)
		if err != nil {
			return err
		}

		a.websiteFileMd5Outputs = append(a.websiteFileMd5Outputs, obj.Etag)

		return nil
	})
	if err != nil {
		return err
	}

	a.Websites[name] = &website{
		bucket:   websiteBucket,
		basePath: config.BasePath,
	}

	return nil
}

func (a *NitricAwsPulumiProvider) deployCloudfrontDistribution(ctx *pulumi.Context) error {
	origins := cloudfront.DistributionOriginArray{}
	orderedCacheBehaviors := cloudfront.DistributionOrderedCacheBehaviorArray{}
	var defaultCacheBehavior cloudfront.DistributionDefaultCacheBehaviorArgs

	oai, err := cloudfront.NewOriginAccessIdentity(ctx, "oai", &cloudfront.OriginAccessIdentityArgs{
		Comment: pulumi.String("OAI for accessing S3 bucket"),
	})
	if err != nil {
		return err
	}

	for websiteName, website := range a.Websites {
		policy := pulumi.All(website.bucket.Arn, oai.IamArn).ApplyT(func(args []interface{}) (string, error) {
			bucketID, bucketIdOk := args[0].(string)
			oaiPath, oaiPathOk := args[1].(string)

			if !bucketIdOk || !oaiPathOk {
				return "", fmt.Errorf("failed to get bucket ID or OAI path")
			}

			return fmt.Sprintf(`{
				"Version": "2012-10-17",
				"Statement": [
					{
						"Effect": "Allow",
						"Principal": {
							"AWS": "%s"
						},
						"Action": "s3:GetObject",
						"Resource": "%s/*"
					}
				]
			}`, oaiPath, bucketID), nil
		}).(pulumi.StringOutput)

		_, err = s3.NewBucketPolicy(ctx, fmt.Sprintf("bucket-policy-%s", websiteName), &s3.BucketPolicyArgs{
			Bucket: website.bucket.Bucket,
			Policy: policy,
		})
		if err != nil {
			return err
		}

		origins = append(origins, &cloudfront.DistributionOriginArgs{
			DomainName: website.bucket.BucketRegionalDomainName,
			OriginId:   pulumi.String(websiteName),
			S3OriginConfig: cloudfront.DistributionOriginS3OriginConfigArgs{
				OriginAccessIdentity: oai.CloudfrontAccessIdentityPath,
			},
		})

		code, err := embeds.GetUrlRewriteFunction(website.basePath)
		if err != nil {
			return err
		}

		rewriteFun, err := cloudfront.NewFunction(ctx, fmt.Sprintf("url-rewrite-function-%s", websiteName), &cloudfront.FunctionArgs{
			Comment: pulumi.String("Rewrite URLs to default index document"),
			Code:    code,
			Runtime: pulumi.String("cloudfront-js-1.0"),
		})
		if err != nil {
			return err
		}

		// Make cache behaviour for all but the root origin. The root origin uses the default cache behaviour
		if website.basePath != "/" {
			rootCacheBehavior := &cloudfront.DistributionOrderedCacheBehaviorArgs{
				PathPattern:          pulumi.String(strings.TrimPrefix(website.basePath, "/")),
				TargetOriginId:       pulumi.String(websiteName),
				ViewerProtocolPolicy: pulumi.String("redirect-to-https"),
				AllowedMethods: pulumi.StringArray{
					pulumi.String("GET"),
					pulumi.String("HEAD"),
					pulumi.String("OPTIONS"),
				},
				CachedMethods: pulumi.StringArray{
					pulumi.String("GET"),
					pulumi.String("HEAD"),
					pulumi.String("OPTIONS"),
				},
				ForwardedValues: &cloudfront.DistributionOrderedCacheBehaviorForwardedValuesArgs{
					QueryString: pulumi.Bool(false),
					Cookies: &cloudfront.DistributionOrderedCacheBehaviorForwardedValuesCookiesArgs{
						Forward: pulumi.String("none"),
					},
					Headers: pulumi.StringArray{
						pulumi.String("x-nitric-cache-key"),
					},
				},
				FunctionAssociations: cloudfront.DistributionOrderedCacheBehaviorFunctionAssociationArray{
					&cloudfront.DistributionOrderedCacheBehaviorFunctionAssociationArgs{
						EventType:   pulumi.String("viewer-request"),
						FunctionArn: rewriteFun.Arn,
					},
				},
				// could be added to stack config in the future
				MinTtl:     pulumi.Int(0),
				DefaultTtl: pulumi.Int(3600),
				MaxTtl:     pulumi.Int(86400),
			}

			orderedCacheBehaviors = append(orderedCacheBehaviors, rootCacheBehavior)

			// Create a new cache behavior for subpaths rather than modifying the root one
			subpathCacheBehavior := &cloudfront.DistributionOrderedCacheBehaviorArgs{}
			*subpathCacheBehavior = *rootCacheBehavior // Copy all fields
			subpathCacheBehavior.PathPattern = pulumi.Sprintf("%s/*", strings.TrimPrefix(website.basePath, "/"))

			orderedCacheBehaviors = append(orderedCacheBehaviors, subpathCacheBehavior)
		} else {
			defaultCacheBehavior = cloudfront.DistributionDefaultCacheBehaviorArgs{
				TargetOriginId:       pulumi.String(websiteName),
				ViewerProtocolPolicy: pulumi.String("redirect-to-https"),
				AllowedMethods: pulumi.StringArray{
					pulumi.String("GET"),
					pulumi.String("HEAD"),
					pulumi.String("OPTIONS"),
				},
				CachedMethods: pulumi.StringArray{
					pulumi.String("GET"),
					pulumi.String("HEAD"),
					pulumi.String("OPTIONS"),
				},
				ForwardedValues: &cloudfront.DistributionDefaultCacheBehaviorForwardedValuesArgs{
					QueryString: pulumi.Bool(false),
					Cookies: &cloudfront.DistributionDefaultCacheBehaviorForwardedValuesCookiesArgs{
						Forward: pulumi.String("none"),
					},
				},
				FunctionAssociations: cloudfront.DistributionDefaultCacheBehaviorFunctionAssociationArray{
					&cloudfront.DistributionDefaultCacheBehaviorFunctionAssociationArgs{
						EventType:   pulumi.String("viewer-request"),
						FunctionArn: rewriteFun.Arn,
					},
				},
				// could be added to stack config in the future
				MinTtl:     pulumi.Int(0),
				DefaultTtl: pulumi.Int(3600),
				MaxTtl:     pulumi.Int(86400),
			}
		}
	}

	// We conventionally route to nitric resources from this distribution to create a single entry point
	// for the entire stack. e.g. /api/main/* will route to a nitric api named "main"
	apiRewriteFun, err := cloudfront.NewFunction(ctx, "api-url-rewrite-function", &cloudfront.FunctionArgs{
		Comment: pulumi.String("Rewrite API URLs routed to nitric services"),
		Code:    embeds.GetApiUrlRewriteFunction(),
		Runtime: pulumi.String("cloudfront-js-1.0"),
	})
	if err != nil {
		return err
	}

	// Sort the APIs by name
	sortedApiKeys := lo.Keys(a.Apis)
	slices.Sort(sortedApiKeys)

	// For each API forward to the appropriate API gateway
	for _, name := range sortedApiKeys {
		api := a.Apis[name]

		apiDomainName := api.ApiEndpoint.ApplyT(func(endpoint string) string {
			return strings.Replace(endpoint, "https://", "", 1)
		}).(pulumi.StringOutput)

		origins = append(origins, &cloudfront.DistributionOriginArgs{
			DomainName: apiDomainName,
			OriginId:   pulumi.String(name),
			CustomOriginConfig: &cloudfront.DistributionOriginCustomOriginConfigArgs{
				OriginReadTimeout:    pulumi.Int(30),
				OriginProtocolPolicy: pulumi.String("https-only"),
				OriginSslProtocols: pulumi.StringArray{
					pulumi.String("TLSv1.2"),
					pulumi.String("SSLv3"),
				},
				HttpPort:  pulumi.Int(80),
				HttpsPort: pulumi.Int(443),
			},
		})

		orderedCacheBehaviors = append(orderedCacheBehaviors,
			&cloudfront.DistributionOrderedCacheBehaviorArgs{
				PathPattern: pulumi.Sprintf("api/%s/*", name),
				// rewrite the URL to the nitric service
				FunctionAssociations: cloudfront.DistributionOrderedCacheBehaviorFunctionAssociationArray{
					&cloudfront.DistributionOrderedCacheBehaviorFunctionAssociationArgs{
						EventType:   pulumi.String("viewer-request"),
						FunctionArn: apiRewriteFun.Arn,
					},
				},
				AllowedMethods: pulumi.ToStringArray([]string{"GET", "HEAD", "OPTIONS", "PUT", "POST", "PATCH", "DELETE"}),
				CachedMethods:  pulumi.ToStringArray([]string{"GET", "HEAD", "OPTIONS"}),
				TargetOriginId: pulumi.String(name),
				ForwardedValues: &cloudfront.DistributionOrderedCacheBehaviorForwardedValuesArgs{
					QueryString: pulumi.Bool(true),
					Cookies: &cloudfront.DistributionOrderedCacheBehaviorForwardedValuesCookiesArgs{
						Forward: pulumi.String("all"),
					},
				},
				ViewerProtocolPolicy: pulumi.String("https-only"),
			},
		)
	}

	name := fmt.Sprintf("%s-cdn", a.StackId)

	tags := common.Tags(a.StackId, name, resources.Website)

	// Deploy a CloudFront distribution for the S3 bucket
	a.Distribution, err = cloudfront.NewDistribution(ctx, name, &cloudfront.DistributionArgs{
		Origins:               origins,
		Enabled:               pulumi.Bool(true),
		DefaultCacheBehavior:  defaultCacheBehavior,
		DefaultRootObject:     pulumi.String(a.websiteIndexDocument),
		OrderedCacheBehaviors: orderedCacheBehaviors,
		Restrictions: &cloudfront.DistributionRestrictionsArgs{
			GeoRestriction: &cloudfront.DistributionRestrictionsGeoRestrictionArgs{
				RestrictionType: pulumi.String("none"),
			},
		},
		Tags: pulumi.ToStringMap(tags),
		ViewerCertificate: &cloudfront.DistributionViewerCertificateArgs{
			CloudfrontDefaultCertificate: pulumi.Bool(true),
		},
		CustomErrorResponses: cloudfront.DistributionCustomErrorResponseArray{
			&cloudfront.DistributionCustomErrorResponseArgs{
				ErrorCode:        pulumi.Int(404),
				ResponseCode:     pulumi.Int(200),
				ResponsePagePath: pulumi.String(fmt.Sprintf("/%v", a.websiteErrorDocument)),
			},
			// Redirect all 403 errors to the error page, s3 by default will return a 403 for missing files
			&cloudfront.DistributionCustomErrorResponseArgs{
				ErrorCode:        pulumi.Int(403),
				ResponseCode:     pulumi.Int(200),
				ResponsePagePath: pulumi.String(fmt.Sprintf("/%v", a.websiteErrorDocument)),
			},
		},
	})
	if err != nil {
		return err
	}

	ctx.Export("cdn", pulumi.Sprintf("https://%s", a.Distribution.DomainName))

	if a.AwsConfig.Cdn.SkipCacheInvalidation {
		return nil
	}

	// Apply a function to sort the array
	sortedMd5Result := a.websiteFileMd5Outputs.ToArrayOutput().ApplyT(func(arr []interface{}) string {
		// Convert each element to string
		md5Strings := []string{}
		for _, md5 := range arr {
			if md5Str, ok := md5.(string); ok {
				if md5Str != "" {
					md5Strings = append(md5Strings, md5Str)
				}
			}
		}

		sort.Strings(md5Strings)

		return strings.Join(md5Strings, "")
	}).(pulumi.StringOutput)

	var interpreter pulumi.StringArrayInput

	// change the interpreter to PowerShell if running on Windows due to issues regarding double quotes
	// https://github.com/pulumi/pulumi-command/issues/271
	if runtime.GOOS == "windows" {
		interpreter = pulumi.StringArray{
			pulumi.String("powershell"),
			pulumi.String("-Command"),
		}
	}

	// Invalidate the CDN Cache
	_, err = local.NewCommand(ctx, "invalidate-cache", &local.CommandArgs{
		Create: pulumi.Sprintf(`aws cloudfront create-invalidation --distribution-id %s --paths "/*"`,
			a.Distribution.ID().ToStringOutput()),
		Triggers: pulumi.Array{
			sortedMd5Result,
		},
		Logging:     local.LoggingStdoutAndStderr,
		Interpreter: interpreter,
	}, pulumi.DependsOn([]pulumi.Resource{a.Distribution}))
	if err != nil {
		return err
	}

	return nil
}

// SafeInt32 - Safely convert an int to an int32
func SafeInt32(n int) (int32, error) {
	if n > math.MaxInt32 {
		return 0, fmt.Errorf("value exceeds int32 limit: %d", n)
	}

	return int32(n), nil //#nosec G115 -- n is checked to be within int32 range
}
