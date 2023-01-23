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
	"path/filepath"
	"unicode"

	"github.com/pulumi/pulumi-tailscale/provider/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/tailscale/terraform-provider-tailscale/tailscale"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "tailscale"
	// modules:
	mainMod = "index" // the y module
)

// makeMember manufactures a type token for the package and the given module and type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeDataSource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the data source's
// first character.
func makeDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeMember(mod+"/"+fn, res)
}

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(tailscale.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:                    p,
		Name:                 "tailscale",
		DisplayName:          "Tailscale",
		Description:          "A Pulumi package for creating and managing Tailscale cloud resources.",
		Keywords:             []string{"pulumi", "tailscale"},
		License:              "Apache-2.0",
		Homepage:             "https://pulumi.io",
		GitHubOrg:            "tailscale",
		Repository:           "https://github.com/pulumi/pulumi-tailscale",
		Config:               map[string]*tfbridge.SchemaInfo{},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"tailscale_acl": {
				Tok: makeResource(mainMod, "Acl"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"acl": {
						CSharpName: "AclJson",
					},
				},
			},
			"tailscale_device_subnet_routes": {Tok: makeResource(mainMod, "DeviceSubnetRoutes")},
			"tailscale_dns_nameservers":      {Tok: makeResource(mainMod, "DnsNameservers")},
			"tailscale_dns_preferences":      {Tok: makeResource(mainMod, "DnsPreferences")},
			"tailscale_dns_search_paths":     {Tok: makeResource(mainMod, "DnsSearchPaths")},
			"tailscale_device_authorization": {Tok: makeResource(mainMod, "DeviceAuthorization")},
			"tailscale_device_tags":          {Tok: makeResource(mainMod, "DeviceTags")},
			"tailscale_tailnet_key":          {Tok: makeResource(mainMod, "TailnetKey")},
			"tailscale_device_key":           {Tok: makeResource(mainMod, "DeviceKey")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"tailscale_device":  {Tok: makeDataSource(mainMod, "getDevice")},
			"tailscale_devices": {Tok: makeDataSource(mainMod, "getDevices")},
			"tailscale_4via6":   {Tok: makeDataSource(mainMod, "get4Via6")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
