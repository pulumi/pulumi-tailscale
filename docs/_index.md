---
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Tailscale Provider
meta_desc: Provides an overview on how to configure the Pulumi Tailscale provider.
layout: package
---

## Installation

The Tailscale provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/tailscale`](https://www.npmjs.com/package/@pulumi/tailscale)
* Python: [`pulumi-tailscale`](https://pypi.org/project/pulumi-tailscale/)
* Go: [`github.com/pulumi/pulumi-tailscale/sdk/go/tailscale`](https://github.com/pulumi/pulumi-tailscale)
* .NET: [`Pulumi.Tailscale`](https://www.nuget.org/packages/Pulumi.Tailscale)
* Java: [`com.pulumi/tailscale`](https://central.sonatype.com/artifact/com.pulumi/tailscale)

## Overview

This provider is used to interact with resources supported by the [Tailscale API](https://tailscale.com/api).

Use the navigation to the left to read about the available resources and functions.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml,hcl" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as tailscale from "@pulumi/tailscale";

const sample = new tailscale.DnsNameservers("sample", {nameservers: [
    "1.1.1.1",
    "8.8.8.8",
]});
```

{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_tailscale as tailscale

sample = tailscale.DnsNameservers("sample", nameservers=[
    "1.1.1.1",
    "8.8.8.8",
])
```

{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Tailscale = Pulumi.Tailscale;

return await Deployment.RunAsync(() =>
{
    var sample = new Tailscale.DnsNameservers("sample", new()
    {
        Nameservers = new[]
        {
            "1.1.1.1",
            "8.8.8.8",
        },
    });

});

```

{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-tailscale/sdk/go/tailscale"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := tailscale.NewDnsNameservers(ctx, "sample", &tailscale.DnsNameserversArgs{
			Nameservers: pulumi.StringArray{
				pulumi.String("1.1.1.1"),
				pulumi.String("8.8.8.8"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
resources:
  sample:
    type: tailscale:DnsNameservers
    properties:
      nameservers:
        - 1.1.1.1
        - 8.8.8.8
```

{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.tailscale.DnsNameservers;
import com.pulumi.tailscale.DnsNameserversArgs;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var sample = new DnsNameservers("sample", DnsNameserversArgs.builder()
            .nameservers(
                "1.1.1.1",
                "8.8.8.8")
            .build());

    }
}
```

{{% /choosable %}}
{{% choosable language hcl %}}
```hcl
pulumi {
  required_providers {
    tailscale = {
      source = "pulumi/tailscale"
    }
  }
}

resource "tailscale_dnsnameservers" "sample" {
  nameservers = ["1.1.1.1", "8.8.8.8"]
}
```

{{% /choosable %}}
{{< /chooser >}}
## Authentication

There are several ways to authenticate the Tailscale provider with the Tailscale API.

Using a [trust credential](https://tailscale.com/kb/1623/trust-credentials) (an OAuth client or federated identity) is
recommended whenever possible as trust credentials can have granular access scopes applied to them whereas API keys cannot.

Available authentication methods are detailed below.
### OAuth clients

[OAuth clients](https://tailscale.com/kb/1215/oauth-clients) can be used for authentication by setting the `oauthClientId`
and `oauthClientSecret` arguments in the provider configuration to the client ID and client secret of a configured OAuth client
respectively:

```bash
$ pulumi config set tailscale:oauthClientId my_client_id
$ pulumi config set --secret tailscale:oauthClientSecret my_client_secret
$ pulumi config set tailscale:tailnet example.com
```

See argument reference for more details.
### Federated identities

[Workload identity federation](https://tailscale.com/kb/1581/workload-identity-federation) can be used for authentication
by setting the `oauthClientId` and `identityToken` in the provider configuration to the client ID of a configured
federated identity and a JWT identity token from a compatible issuer respectively:

```bash
$ pulumi config set --secret tailscale:identityToken my_identity_token
$ pulumi config set tailscale:oauthClientId my_client_id
$ pulumi config set tailscale:tailnet example.com
```

If Pulumi is running in a supported runtime (GitHub Actions, AWS via EC2 instance profile or ECS task role, or Google
Cloud), the provider can discover the OIDC token from the runtime automatically. Configure only the `oauthClientId`
and the `audience` expected by the federated identity:

```bash
$ pulumi config set tailscale:audience my_audience
$ pulumi config set tailscale:oauthClientId my_client_id
$ pulumi config set tailscale:tailnet example.com
```

For GitHub Actions, the workflow must declare `permissions: id-token: write`. For AWS, the runtime must have valid AWS
credentials available. For GCP, the runtime must be able to reach the metadata server.

See argument reference for more details.
### API keys

[API keys](https://tailscale.com/kb/1101/api#authentication) can be used for authentication by setting the `apiKey`
argument in the provider configuration:

```bash
$ pulumi config set --secret tailscale:apiKey my_api_key
$ pulumi config set tailscale:tailnet example.com
```

See argument reference for more details.
## Configuration Reference

- `apiKey` (String, Sensitive) The API key to use for authenticating requests to the API. Can be set via the TAILSCALE_API_KEY environment variable. If the value starts with 'file:' then it is treated as a path to a file on disk that contains the API key. Conflicts with `oauthClientId` and `oauthClientSecret`.
- `audience` (String) The OIDC audience to request when discovering an identity token from the runtime (GitHub Actions, AWS, or GCP) for workload identity federation. Can be set via the TAILSCALE_AUDIENCE environment variable. If the value starts with 'file:' then it is treated as a path to a file on disk that contains the audience. Requires `oauthClientId`. Conflicts with `apiKey`, `oauthClientSecret`, `identityToken`, and `identityTokenEnvironmentVariableName`.
- `baseUrl` (String) The base URL of the Tailscale API. Defaults to <https://api.tailscale.com>. Can be set via the TAILSCALE_BASE_URL environment variable.
- `identityToken` (String, Sensitive) The jwt identity token to exchange for a Tailscale API token when using a federated identity. Can be set via the TAILSCALE_IDENTITY_TOKEN environment variable. If the value starts with 'file:' then it is treated as a path to a file on disk that contains the identity token. Conflicts with `apiKey`, `oauthClientSecret`, and `identityTokenEnvironmentVariableName`.
- `identityTokenEnvironmentVariableName` (String) The name of an environment variable to read the identity token from. This is useful when the identity token is provided by an external system (such as Pulumi Cloud workload identity) in an environment variable you do not control. If the resolved value of the environment variable starts with 'file:' then it is treated as a path to a file on disk that contains identity token. Conflicts with `identityToken`.
- `oauthClientId` (String) The OAuth application or federated identity's ID when using OAuth client credentials or workload identity federation. Can be set via the TAILSCALE_OAUTH_CLIENT_ID environment variable. If the value starts with 'file:' then it is treated as a path to a file on disk that contains the client ID. Either `oauthClientSecret` or `identityToken` must be set alongside `oauthClientId`. Conflicts with `apiKey`.
- `oauthClientSecret` (String, Sensitive) The OAuth application's secret when using OAuth client credentials. Can be set via the TAILSCALE_OAUTH_CLIENT_SECRET environment variable. If the value starts with 'file:' then it is treated as a path to a file on disk that contains the client secret. Conflicts with `apiKey` and `identityToken`.
- `scopes` (List of String) The OAuth 2.0 scopes to request when generating the access token using the supplied OAuth client credentials. See <https://tailscale.com/kb/1623/trust-credentials#scopes> for available scopes. Only valid when both `oauthClientId` and `oauthClientSecret` are set.
- `tailnet` (String) The tailnet ID. Tailnets created before Oct 2025 can still use the legacy ID, but the Tailnet ID is the preferred identifier. Can be set via the TAILSCALE_TAILNET environment variable. Default is the tailnet that owns API credentials passed to the provider.
- `userAgent` (String) User-Agent header for API requests.