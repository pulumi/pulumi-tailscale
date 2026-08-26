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
	"regexp"
	"strings"

	// Allow embedding metadata
	_ "embed"

	"github.com/tailscale/terraform-provider-tailscale/tailscale"

	pfbridge "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/info"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"

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
	p := pfbridge.ShimProvider(tailscale.NewFrameworkProvider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        mainPkg,
		Version:     version.Version,
		DisplayName: "Tailscale",
		Description: "A Pulumi package for creating and managing Tailscale cloud resources.",
		Keywords:    []string{"pulumi", mainPkg},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		GitHubOrg:   mainPkg,
		Repository:  "https://github.com/pulumi/pulumi-tailscale",
		Config: map[string]*tfbridge.SchemaInfo{
			"user_agent": {
				Default: &tfbridge.DefaultInfo{
					Value: fmt.Sprintf("Pulumi/3.0 (https://www.pulumi.com) pulumi-tailscale/%s", version.Version),
				},
			},
		},
		MetadataInfo: tfbridge.NewProviderMetadata(metadata),
		DocRules:     &tfbridge.DocRuleInfo{EditRules: editRules},
		// The upstream descriptions are carried verbatim into schema.json and from
		// there into every SDK's docstrings, so the fixups below have to be applied to
		// the schema as well as to the registry docs.
		SchemaPostProcessor: fixUpstreamConfigDescriptions,
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

func editRules(defaults []tfbridge.DocsEdit) []tfbridge.DocsEdit {
	return append(
		defaults,
		fixScopesSentence,
		stripTerraformBlock,
		camelCaseQuotedNames,
		configExamplesAsStackConfig,
	)
}

// configNames maps the upstream provider's attribute names to the names Pulumi exposes.
//
// The bridge already renames these where upstream writes them in backticks. It misses
// them where upstream writes them in single quotes, because the renamer's delimiter set
// (pkg/tfgen/docs.go, codeLikeSingleWord) covers whitespace, double quotes, backticks
// and brackets but not apostrophes. Attributes whose name is a single word need no
// entry, since snake_case is only ambiguous for compound names.
//
//nolint:gosec // G101 fires on the attribute names below; they are field names, not credentials.
var configNames = map[string]string{
	"api_key":        "apiKey",
	"base_url":       "baseUrl",
	"identity_token": "identityToken",
	"identity_token_environment_variable_name": "identityTokenEnvironmentVariableName",
	"oauth_client_id":                          "oauthClientId",
	"oauth_client_secret":                      "oauthClientSecret",
	"user_agent":                               "userAgent",
}

// secretConfigKeys are the config keys schema.json marks sensitive. Examples that set
// them have to use `pulumi config set --secret`.
var secretConfigKeys = map[string]bool{
	"apiKey":            true,
	"identityToken":     true,
	"oauthClientSecret": true,
}

// The upstream `scopes` description ends in a sentence that lost a clause somewhere:
// "Only valid when both 'oauth_client_id' and 'oauth_client_secret', or both are set."
//
// Both halves are patched here rather than only in the docs, because the description is
// also copied into schema.json and from there into the SDKs. The rule is a no-op once
// the sentence is fixed upstream.
const (
	garbledScopesSentence = "Only valid when both 'oauth_client_id' and 'oauth_client_secret', or both are set."
	fixedScopesSentence   = "Only valid when both 'oauth_client_id' and 'oauth_client_secret' are set."
)

var fixScopesSentence = tfbridge.DocsEdit{
	Path: "*",
	Edit: func(_ string, content []byte) ([]byte, error) {
		return []byte(strings.ReplaceAll(string(content), garbledScopesSentence, fixedScopesSentence)), nil
	},
}

// tfplugindocs renders the upstream provider example inside a `terraform` block pinning
// the provider version. `pulumi convert` cannot translate that block; the failure is
// swallowed (pkg/tfgen/convert_cli.go, singleExampleFromHCLToPCL) and the resulting empty
// example takes the whole "Example Usage" section with it
// (pkg/tfgen/installation_docs.go, removeEmptySection). Strip the block so only the
// `provider` block reaches the converter.
//
// Anchoring the closing brace at column 0 is enough to find the end of the block:
// tfplugindocs output is gofmt-like, so nothing else in the example starts a line with
// a closing brace.
var terraformBlock = regexp.MustCompile(`(?m)^terraform[ \t]*\{\n(?s:.*?)^\}\n+`)

var stripTerraformBlock = tfbridge.DocsEdit{
	Path: "index.md",
	Edit: func(_ string, content []byte) ([]byte, error) {
		return terraformBlock.ReplaceAll(content, nil), nil
	},
}

// camelCaseQuotedNames rewrites 'oauth_client_id' as `oauthClientId`, which the bridge's
// own renamer skips. See configNames.
var singleQuotedWord = regexp.MustCompile(`'([a-z][a-z0-9_]*)'`)

var camelCaseQuotedNames = tfbridge.DocsEdit{
	Path:  "*",
	Phase: info.PostCodeTranslation,
	Edit: func(_ string, content []byte) ([]byte, error) {
		return camelCaseQuotedNamesIn(content), nil
	},
}

func camelCaseQuotedNamesIn(content []byte) []byte {
	return singleQuotedWord.ReplaceAllFunc(content, func(match []byte) []byte {
		camel, ok := configNames[string(match[1:len(match)-1])]
		if !ok {
			return match
		}
		return []byte("`" + camel + "`")
	})
}

// `pulumi convert` renders a Terraform `provider` block as project-level configuration in
// Pulumi.yaml, using the `value:` form. That is valid, but it is the wrong advice for
// credentials: Pulumi.yaml is committed and shared by every stack, and the generated
// entries carry no `secret: true`, so the page would show API keys and client secrets as
// plaintext in source control. The runtime field is empty as well, because the bridge has
// no language to fill in for a provider-only example, which makes the file invalid as
// printed.
//
// Rewrite each block as the `pulumi config set` commands that put the same values in
// Pulumi.<stack>.yaml, encrypting the ones schema.json marks sensitive.
var (
	configYAMLExample = regexp.MustCompile("(?s)```yaml\n# Pulumi\\.yaml provider configuration file\n.*?```\n")
	configYAMLEntry   = regexp.MustCompile(`(?m)^[ \t]+tailscale:(\w+):\n[ \t]+value: (.+)$`)
)

var configExamplesAsStackConfig = tfbridge.DocsEdit{
	Path:  "index.md",
	Phase: info.PostCodeTranslation,
	Edit: func(_ string, content []byte) ([]byte, error) {
		return configYAMLExample.ReplaceAllFunc(content, func(block []byte) []byte {
			entries := configYAMLEntry.FindAllSubmatch(block, -1)
			if len(entries) == 0 {
				// Not a shape we recognize; leave it alone rather than drop it.
				return block
			}
			var b strings.Builder
			b.WriteString("```bash\n")
			for _, entry := range entries {
				key, value := string(entry[1]), strings.TrimSpace(string(entry[2]))
				secret := ""
				if secretConfigKeys[key] {
					secret = "--secret "
				}
				fmt.Fprintf(&b, "$ pulumi config set %stailscale:%s %s\n", secret, key, value)
			}
			b.WriteString("```\n")
			return []byte(b.String())
		}), nil
	},
}

// fixUpstreamConfigDescriptions applies the fixups the docs get to the config
// descriptions in the schema, which the SDKs render as docstrings.
//
// The registry docs get the Terraform Cloud rename for free, from the bridge's blanket
// Terraform -> Pulumi replacement; the schema does not, and a Pulumi SDK pointing its
// users at Terraform Cloud is worse than unhelpful.
func fixUpstreamConfigDescriptions(spec *pschema.PackageSpec) {
	fix := func(description string) string {
		description = strings.ReplaceAll(description, garbledScopesSentence, fixedScopesSentence)
		description = strings.ReplaceAll(description, "Terraform Cloud workload identity", "Pulumi Cloud workload identity")
		return string(camelCaseQuotedNamesIn([]byte(description)))
	}
	for name, variable := range spec.Config.Variables {
		variable.Description = fix(variable.Description)
		spec.Config.Variables[name] = variable
	}
	for name, property := range spec.Provider.InputProperties {
		property.Description = fix(property.Description)
		spec.Provider.InputProperties[name] = property
	}
	for name, property := range spec.Provider.Properties {
		property.Description = fix(property.Description)
		spec.Provider.Properties[name] = property
	}
}
