//go:build nodejs || all
// +build nodejs all

package examples

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestAccDnsTs(t *testing.T) {
	checkTokens(t)
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "dns-ts"),
			Secrets: map[string]string{
				"tailscale:oauthClientSecret": os.Getenv("TAILSCALE_OAUTH_CLIENT_SECRET"),
				"tailscale:oauthClientId":     os.Getenv("TAILSCALE_OAUTH_CLIENT_ID"),
			},
		})

	integration.ProgramTest(t, &test)
}

func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	baseJS := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"@pulumi/tailscale",
		},
	})

	return baseJS
}
