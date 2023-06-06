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

### Node.js (Java/TypeScript)

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

## Configuration

The following configuration points are available:

- `tailscale:apiKey` - (Required) API key to authenticate with the Tailscale API. It must be provided, but it can also be sourced
  from the `TAILSCALE_API_KEY` environment variable.
- `tailscale:tailnet` - (Required) Tailscale tailnet to manage resources for. It must be provided, but it can also be
  sourced from the `TAILSCALE_TAILNET` variable. A tailnet is the name of your Tailscale network. You can find it in 
  the top left corner of the Admin Panel beside the Tailscale logo.
- `tailscale:oauthClientId` - The OAuth application's ID when using OAuth client credentials. Can be set via the OAUTH_CLIENT_ID environment variable. Both 'oauthClientId' and 'oauthClientSecret' must be set. Conflicts with 'apiKey'.
- `oauthClientSecret` - The OAuth application's secret when using OAuth client credentials. Can be set via the OAUTH_CLIENT_SECRET environment variable. Both 'oauthClientId' and 'oauthClientSecret' must be set. Conflicts with 'apiKey'.
- `scopes` - The OAuth 2.0 scopes to request when for the access token generated using the supplied OAuth client credentials. See https://tailscale.com/kb/1215/oauth-clients/#scopes for available scopes. Only valid when both 'oauthClientId' and 'oauthClientSecret' are set.

## Reference

For further information, please visit [the Tailscale provider docs](https://www.pulumi.com/registry/packages/tailscale)
or for detailed reference documentation, please visit [the API docs](https://www.pulumi.com/registry/packages/tailscale/api-docs/).
