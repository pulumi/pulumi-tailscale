package tailscale

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// The Example Usage section as tfplugindocs renders it in the upstream provider's
// docs/index.md at v0.29.2: a version pin and a provider block, and no resource.
const upstreamExampleUsage = "## Example Usage\n" +
	"\n" +
	"```terraform\n" +
	"terraform {\n" +
	"  required_providers {\n" +
	"    tailscale = {\n" +
	`      source  = "tailscale/tailscale"` + "\n" +
	`      version = "<version>"` + "\n" +
	"    }\n" +
	"  }\n" +
	"}\n" +
	"\n" +
	`provider "tailscale" {` + "\n" +
	`  oauth_client_id      = "my_client_id"` + "\n" +
	`  oauth_client_secret  = "my_client_secret"` + "\n" +
	`  tailnet              = "example.com"` + "\n" +
	"}\n" +
	"```\n" +
	"\n" +
	"## Authentication\n"

func TestOverrideExampleUsage(t *testing.T) {
	t.Parallel()

	actual, err := overrideExampleUsage.Edit("index.md", []byte(upstreamExampleUsage))
	require.NoError(t, err)

	expected := "## Example Usage\n\n" + exampleUsageReplacement + "\n## Authentication\n"
	require.Equal(t, expected, string(actual))

	// The credentials upstream puts in the example must not survive into the page.
	require.NotContains(t, string(actual), "my_client_secret")
	require.NotContains(t, string(actual), "terraform {")
}

func TestOverrideExampleUsageLeavesOtherSectionsAlone(t *testing.T) {
	t.Parallel()

	// The Authentication examples are bare provider blocks that convert fine; only the
	// block under "## Example Usage" is replaced.
	input := "## Authentication\n\n```terraform\n" +
		`provider "tailscale" {` + "\n" +
		`  api_key = "my_api_key"` + "\n" +
		"}\n```\n"
	actual, err := overrideExampleUsage.Edit("index.md", []byte(input))
	require.NoError(t, err)
	require.Equal(t, input, string(actual))
}

func TestCamelCaseQuotedNames(t *testing.T) {
	t.Parallel()

	input := []byte("Conflicts with 'oauth_client_id' and 'oauth_client_secret'. " +
		"Conflicts with 'identity_token' and 'identity_token_environment_variable_name'. " +
		"If the value starts with 'file:' then it is treated as a path. Requires 'tailnet'.")

	actual, err := camelCaseQuotedNames.Edit("index.md", input)
	require.NoError(t, err)

	require.Equal(t,
		"Conflicts with `oauthClientId` and `oauthClientSecret`. "+
			"Conflicts with `identityToken` and `identityTokenEnvironmentVariableName`. "+
			"If the value starts with 'file:' then it is treated as a path. Requires 'tailnet'.",
		string(actual))
}

func TestFixScopesSentence(t *testing.T) {
	t.Parallel()

	input := []byte("See the docs for available scopes. " + garbledScopesSentence)
	actual, err := fixScopesSentence.Edit("index.md", input)
	require.NoError(t, err)
	require.Equal(t, "See the docs for available scopes. "+fixedScopesSentence, string(actual))
}

func TestConfigExamplesAsStackConfig(t *testing.T) {
	t.Parallel()

	input := "See argument reference.\n\n" +
		"```yaml\n" +
		"# Pulumi.yaml provider configuration file\n" +
		"name: configuration-example\n" +
		"runtime:\n" +
		"config:\n" +
		"    tailscale:oauthClientId:\n" +
		"        value: my_client_id\n" +
		"    tailscale:oauthClientSecret:\n" +
		"        value: my_client_secret\n" +
		"    tailscale:tailnet:\n" +
		"        value: example.com\n" +
		"\n" +
		"```\n" +
		"More prose.\n"

	actual, err := configExamplesAsStackConfig.Edit("index.md", []byte(input))
	require.NoError(t, err)

	expected := "See argument reference.\n\n" +
		"```bash\n" +
		"$ pulumi config set tailscale:oauthClientId my_client_id\n" +
		"$ pulumi config set --secret tailscale:oauthClientSecret my_client_secret\n" +
		"$ pulumi config set tailscale:tailnet example.com\n" +
		"```\n" +
		"More prose.\n"
	require.Equal(t, expected, string(actual))
}

func TestConfigExamplesAsStackConfigLeavesUnrecognizedBlocksAlone(t *testing.T) {
	t.Parallel()

	input := "```yaml\n# Pulumi.yaml provider configuration file\nname: configuration-example\nruntime:\n\n```\n"
	actual, err := configExamplesAsStackConfig.Edit("index.md", []byte(input))
	require.NoError(t, err)
	require.Equal(t, input, string(actual))
}
