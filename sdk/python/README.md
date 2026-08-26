[![Actions Status](https://github.com/pulumi/pulumi-tailscale/workflows/master/badge.svg)](https://github.com/pulumi/pulumi-tailscale/actions)
[![Slack](http://www.pulumi.com/images/docs/badges/slack.svg)](https://slack.pulumi.com)
[![NPM version](https://badge.fury.io/js/%40pulumi%2Ftailscale.svg)](https://www.npmjs.com/package/@pulumi/tailscale)
[![Python version](https://badge.fury.io/py/pulumi-tailscale.svg)](https://pypi.org/project/pulumi-tailscale)
[![NuGet version](https://badge.fury.io/nu/pulumi.tailscale.svg)](https://badge.fury.io/nu/pulumi.tailscale)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pulumi/pulumi-tailscale/sdk)](https://pkg.go.dev/github.com/pulumi/pulumi-tailscale/sdk)
[![License](https://img.shields.io/npm/l/%40pulumi%2Fpulumi.svg)](https://github.com/pulumi/pulumi-tailscale/blob/master/LICENSE)

# Tailscale Resource Provider

The Tailscale Resource Provider lets you manage [Tailscale](https://tailscale.com/) resources.

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/tailscale

or `yarn`:

    $ yarn add @pulumi/tailscale

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_tailscale

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-tailscale/sdk

### .NET

To use from .NET, install using `dotnet add package`:

    $ dotnet add package Pulumi.Tailscale

### Java

To use from Java, add the following to the dependencies of your `pom.xml`:

    <dependency>
        <groupId>com.pulumi</groupId>
        <artifactId>tailscale</artifactId>
    </dependency>

## Configuration

The following configuration points are available. None is required: credentials and the
tailnet can equally be supplied through the environment, and `tailnet` defaults to the
tailnet that owns the credentials you authenticate with.

Set them with `pulumi config set`, using `--secret` for the sensitive ones:

    $ pulumi config set tailscale:oauthClientId my_client_id
    $ pulumi config set --secret tailscale:oauthClientSecret my_client_secret

- `tailscale:apiKey` - (Sensitive) API key to authenticate with the Tailscale API. Can be set via the
  `TAILSCALE_API_KEY` environment variable. If the value starts with `file:` it is treated as a path to a file on
  disk that contains the API key. Conflicts with `tailscale:oauthClientId` and `tailscale:oauthClientSecret`.
- `tailscale:audience` - The OIDC audience to request when discovering an identity token from the runtime
  (GitHub Actions, AWS, or Google Cloud) for workload identity federation. Can be set via the `TAILSCALE_AUDIENCE`
  environment variable. Requires `tailscale:oauthClientId`.
- `tailscale:baseUrl` - The base URL of the Tailscale API. Defaults to `https://api.tailscale.com`. Can be set via
  the `TAILSCALE_BASE_URL` environment variable.
- `tailscale:identityToken` - (Sensitive) The JWT identity token to exchange for a Tailscale API token when using a
  federated identity. Can be set via the `TAILSCALE_IDENTITY_TOKEN` environment variable. Conflicts with
  `tailscale:apiKey` and `tailscale:oauthClientSecret`.
- `tailscale:identityTokenEnvironmentVariableName` - The name of an environment variable to read the identity token
  from, for when an external system supplies it in a variable you do not control. Conflicts with
  `tailscale:identityToken`.
- `tailscale:oauthClientId` - The OAuth application or federated identity's ID when using OAuth client credentials
  or workload identity federation. Can be set via the `TAILSCALE_OAUTH_CLIENT_ID` environment variable. Either
  `tailscale:oauthClientSecret` or `tailscale:identityToken` must be set alongside it. Conflicts with
  `tailscale:apiKey`.
- `tailscale:oauthClientSecret` - (Sensitive) The OAuth application's secret when using OAuth client credentials.
  Can be set via the `TAILSCALE_OAUTH_CLIENT_SECRET` environment variable. Conflicts with `tailscale:apiKey` and
  `tailscale:identityToken`.
- `tailscale:scopes` - The OAuth 2.0 scopes to request when generating the access token using the supplied OAuth
  client credentials. See https://tailscale.com/kb/1623/trust-credentials#scopes for available scopes. Only valid
  when both `tailscale:oauthClientId` and `tailscale:oauthClientSecret` are set.
- `tailscale:tailnet` - The tailnet to manage resources for. Tailnets created before October 2025 can still use the
  legacy ID, but the tailnet ID is the preferred identifier. Can be set via the `TAILSCALE_TAILNET` environment
  variable. Defaults to the tailnet that owns the credentials passed to the provider.
- `tailscale:userAgent` - User-Agent header for API requests. Defaults to a Pulumi-specific value.

## Reference

For further information, please visit [the Tailscale provider docs](https://www.pulumi.com/registry/packages/tailscale)
or for detailed reference documentation, please visit [the API docs](https://www.pulumi.com/registry/packages/tailscale/api-docs/).
