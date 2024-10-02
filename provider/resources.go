// Copyright 2016-2018, Pulumi Corporation.
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

package tailscale

import (
	"fmt"
	"path"

	// Allow embedding metadata
	_ "embed"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tailscale/terraform-provider-tailscale/tailscale"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"

	"github.com/pulumi/pulumi-tailscale/provider/pkg/version"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "tailscale"
	// modules:
	mainMod = "index" // the y module
)

//go:embed cmd/pulumi-resource-tailscale/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(tailscale.Provider(addUserAgent))

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:            p,
		Name:         "tailscale",
		DisplayName:  "Tailscale",
		Description:  "A Pulumi package for creating and managing Tailscale cloud resources.",
		Keywords:     []string{"pulumi", "tailscale"},
		License:      "Apache-2.0",
		Homepage:     "https://pulumi.io",
		GitHubOrg:    "tailscale",
		Repository:   "https://github.com/pulumi/pulumi-tailscale",
		Config:       map[string]*tfbridge.SchemaInfo{},
		MetadataInfo: tfbridge.NewProviderMetadata(metadata),
		Resources: map[string]*tfbridge.ResourceInfo{
			"tailscale_acl": {
				Fields: map[string]*tfbridge.SchemaInfo{
					"acl": {CSharpName: "AclJson"},
				},
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Overridden for back-compat reasons. Defaults to "get4via6".
			"tailscale_4via6": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "get4Via6")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			RespectSchemaVersion: true,
		},
		Python: &tfbridge.PythonInfo{
			RespectSchemaVersion: true,

			PyProject: struct{ Enabled bool }{true},
		},

		Golang: &tfbridge.GolangInfo{
			ImportBasePath: path.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
			RespectSchemaVersion:           true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RespectSchemaVersion: true,
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.MustComputeTokens(tokens.SingleModule("tailscale_", mainMod, tokens.MakeStandard(mainPkg)))
	prov.SetAutonaming(255, "-")
	prov.MustApplyAutoAliases()

	return prov
}

// addUserAgent adds a `user_agent` configuration key to the provider with a
// default value based on provider version. This is adapted from the upstream provider.
// See: https://github.com/tailscale/terraform-provider-tailscale/commit/e2961ac83f24bc2cc279177abcf23e28570815c6
func addUserAgent(p *schema.Provider) {
	p.Schema["user_agent"] = &schema.Schema{
		Type:        schema.TypeString,
		Default:     fmt.Sprintf("Pulumi/3.0 (https://www.pulumi.com) pulumi-tailscale/%s", version.Version),
		Optional:    true,
		Description: "User-Agent header for API requests.",
	}
}
