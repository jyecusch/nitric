// Copyright 2021 Nitric Technologies Pty Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package deploytf

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/nitrictech/nitric/cloud/aws/deploytf/generated/website"
	deploymentspb "github.com/nitrictech/nitric/core/pkg/proto/deployments/v1"
)

// Websites - Deploy a Website
func (a *NitricAwsTerraformProvider) Website(stack cdktf.TerraformStack, name string, config *deploymentspb.Website) error {
	allDependants := []cdktf.ITerraformDependable{a.Stack}

	a.Websites[name] = website.NewWebsite(stack, jsii.String(name), &website.WebsiteConfig{
		Name:           jsii.String(name),
		StackId:        a.Stack.StackIdOutput(),
		LocalDirectory: jsii.String(config.GetLocalDirectory()),
		BasePath:       jsii.String(config.GetBasePath()),
		DependsOn:      &allDependants,
	})

	return nil
}
