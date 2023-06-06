package examples

import (
	"os"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}

func checkTokens(t *testing.T) {
	clientSecret := os.Getenv("TAILSCALE_OAUTH_CLIENT_SECRET")
	if clientSecret == "" {
		t.Log("Failing due to missing TAILSCALE_OAUTH_CLIENT_SECRET variable")
		t.FailNow()

	}
	clientId := os.Getenv("TAILSCALE_OAUTH_CLIENT_ID")
	if clientId == "" {
		t.Log("Failing due to missing TAILSCALE_OAUTH_CLIENT_ID variable")
		t.FailNow()
	}
}

func getBaseOptions() integration.ProgramTestOptions {
	return integration.ProgramTestOptions{
		RunUpdateTest:        false,
		ExpectRefreshChanges: true,
	}
}
