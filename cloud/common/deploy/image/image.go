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

package image

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"golang.org/x/exp/maps"

	"github.com/nitrictech/nitric/cloud/common/deploy/telemetry"
)

type ImageArgs struct {
	SourceImage   string
	Runtime       []byte
	RepositoryUrl pulumi.StringInput
	Server        pulumi.StringInput
	Username      pulumi.StringInput
	Password      pulumi.StringInput
	Telemetry *telemetry.TelemetryConfigArgs
}

type Image struct {
	pulumi.ResourceState

	Name        string
	DockerImage *docker.Image
}

type WrappedBuildInput struct {
	Args       map[string]string
	Dockerfile string
}

var (
	//go:embed wrapper.dockerfile
	imageWrapper string
	//go:embed wrapper-telemetry.dockerfile
	telemetryImageWrapper string
)

func NewImage(ctx *pulumi.Context, name string, args *ImageArgs, opts ...pulumi.ResourceOption) (*Image, error) {
	res := &Image{Name: name}

	err := ctx.RegisterComponentResource("nitric:Image", name, res, opts...)
	if err != nil {
		return nil, err
	}

	imageWrapper, err := getWrapperDockerfile(args.Telemetry)
	if err != nil {
		return nil, err
	}

	dockerfileContent, sourceImageID, err := wrapDockerImage(imageWrapper.Dockerfile, args.SourceImage)
	if err != nil {
		return nil, err
	}

	buildContext := fmt.Sprintf("%s/build-%s", os.TempDir(), name)
	os.MkdirAll(buildContext, os.ModePerm)

	dockerfile, err := os.Create(path.Join(buildContext, "Dockerfile"))
	if err != nil {
		return nil, err
	}

	dockerfile.Write([]byte(dockerfileContent))
	dockerfile.Close()

	runtimefile, err := os.Create(path.Join(buildContext, "runtime"))
	if err != nil {
		return nil, err
	}

	runtimefile.Write(args.Runtime)
	runtimefile.Close()			
	
	buildArgs := combineBuildArgs(map[string]string{
		"BASE_IMAGE": args.SourceImage,
		"RUNTIME_FILE": "runtime",
		"BASE_IMAGE_ID": sourceImageID,
	}, imageWrapper.Args)

	res.DockerImage, err = docker.NewImage(ctx, name+"-image", &docker.ImageArgs{
		ImageName:       args.RepositoryUrl,
		Build: docker.DockerBuildArgs{
			Context: pulumi.String(buildContext),
			Dockerfile: pulumi.String(path.Join(buildContext, "Dockerfile")),
			Args: buildArgs,
			Platform: pulumi.String("linux/amd64"),
		},
		Registry: docker.RegistryArgs{
			Server:   args.Server,
			Username: args.Username,
			Password: args.Password,
		},
		SkipPush: pulumi.Bool(false),
	}, pulumi.Parent(res))
	if err != nil {
		return nil, err
	}

	return res, ctx.RegisterResourceOutputs(res, pulumi.Map{
		"name":     pulumi.String(res.Name),
		"imageUri": res.DockerImage.ImageName,
	})
}

func (d *Image) URI() pulumi.StringOutput {
	return d.DockerImage.RepoDigest.Elem().ToStringOutput()
}

// Returns the default docker file if telemetry sampling is disabled for this execution unit. Otherwise, will return a wrapped telemetry image.
func getWrapperDockerfile(configArgs *telemetry.TelemetryConfigArgs) (*WrappedBuildInput, error) {
	if configArgs != nil && configArgs.TraceSampling > 0 {
		config, err := telemetry.NewTelemetryConfig(configArgs)
		if err != nil {
			return nil, err
		}

		return &WrappedBuildInput{
			Dockerfile: telemetryImageWrapper,
			Args: map[string]string{
				"OTELCOL_CONFIG":              config.Config,
				"OTELCOL_CONTRIB_URI":         config.Uri,
				"NITRIC_TRACE_SAMPLE_PERCENT": fmt.Sprint(configArgs.TraceSampling),
			},
		}, nil
	}

	return &WrappedBuildInput{
		Dockerfile: imageWrapper,
		Args: map[string]string{},
	}, nil
}

func combineBuildArgs(baseArgs, wrapperArgs map[string]string) (pulumi.StringMap) {
	maps.Copy(wrapperArgs, baseArgs)

	return pulumi.ToStringMap(wrapperArgs)
}


// Wraps the source image with the wrapper image, acknowledging the command from the source image
func wrapDockerImage(wrapper, sourceImage string) (string, string, error) {
	if sourceImage == "" {
		return "", "", fmt.Errorf("blank sourceImage provided")
	}

	client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return "", "", err
	}

	imageInspect, _, err := client.ImageInspectWithRaw(context.Background(), sourceImage)
	if err != nil {
		return "", "", errors.WithMessage(err, fmt.Sprintf("could not inspect image: %s", sourceImage))
	}

	cmdStr, err := commandFromImageInspect(imageInspect)
	if err != nil {
		return "", "", err
	}

	return fmt.Sprintf(wrapper, cmdStr), imageInspect.ID, nil
}

// Gets the command from the source image and returns as a comma separated string
func commandFromImageInspect(imageInspect types.ImageInspect) (string, error) {
	// Get the original command of the source image
	// and inject it into the wrapper image
	cmd := append(imageInspect.Config.Entrypoint, imageInspect.Config.Cmd...)

	// Wrap each command in string quotes
	cmdStr := []string{}
	for _, c := range cmd {
		cmdStr = append(cmdStr, "\""+c+"\"")
	}

	return strings.Join(cmdStr, ","), nil
}