//go:build python || all
// +build python all

package examples

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestAccKeyPy(t *testing.T) {
	checkTokens(t)
	test := getPythonBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "py-tailnet-key"),
			Secrets: map[string]string{
				"tailscale:oauthClientSecret": os.Getenv("TAILSCALE_OAUTH_CLIENT_SECRET"),
				"tailscale:oauthClientId":     os.Getenv("TAILSCALE_OAUTH_CLIENT_ID"),
			},
		})

	integration.ProgramTest(t, &test)
}

func getPythonBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	basePython := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			filepath.Join("..", "sdk", "python", "bin"),
		},
	})

	return basePython
}
