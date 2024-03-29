// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tailscale;

import com.pulumi.core.TypeShape;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;

public final class Config {

    private static final com.pulumi.Config config = com.pulumi.Config.of("tailscale");
/**
 * The API key to use for authenticating requests to the API. Can be set via the TAILSCALE_API_KEY environment variable.
 * Conflicts with &#39;oauth_client_id&#39; and &#39;oauth_client_secret&#39;.
 * 
 */
    public Optional<String> apiKey() {
        return Codegen.stringProp("apiKey").config(config).get();
    }
/**
 * The base URL of the Tailscale API. Defaults to https://api.tailscale.com. Can be set via the TAILSCALE_BASE_URL
 * environment variable.
 * 
 */
    public Optional<String> baseUrl() {
        return Codegen.stringProp("baseUrl").config(config).get();
    }
/**
 * The OAuth application&#39;s ID when using OAuth client credentials. Can be set via the TAILSCALE_OAUTH_CLIENT_ID environment
 * variable. Both &#39;oauth_client_id&#39; and &#39;oauth_client_secret&#39; must be set. Conflicts with &#39;api_key&#39;.
 * 
 */
    public Optional<String> oauthClientId() {
        return Codegen.stringProp("oauthClientId").config(config).get();
    }
/**
 * The OAuth application&#39;s secret when using OAuth client credentials. Can be set via the TAILSCALE_OAUTH_CLIENT_SECRET
 * environment variable. Both &#39;oauth_client_id&#39; and &#39;oauth_client_secret&#39; must be set. Conflicts with &#39;api_key&#39;.
 * 
 */
    public Optional<String> oauthClientSecret() {
        return Codegen.stringProp("oauthClientSecret").config(config).get();
    }
/**
 * The OAuth 2.0 scopes to request when for the access token generated using the supplied OAuth client credentials. See
 * https://tailscale.com/kb/1215/oauth-clients/#scopes for available scopes. Only valid when both &#39;oauth_client_id&#39; and
 * &#39;oauth_client_secret&#39; are set.
 * 
 */
    public Optional<List<String>> scopes() {
        return Codegen.objectProp("scopes", TypeShape.<List<String>>builder(List.class).addParameter(String.class).build()).config(config).get();
    }
/**
 * The organization name of the Tailnet in which to perform actions. Can be set via the TAILSCALE_TAILNET environment
 * variable. Default is the tailnet that owns API credentials passed to the provider.
 * 
 */
    public Optional<String> tailnet() {
        return Codegen.stringProp("tailnet").config(config).get();
    }
/**
 * User-Agent header for API requests.
 * 
 */
    public Optional<String> userAgent() {
        return Codegen.stringProp("userAgent").config(config).get();
    }
}
