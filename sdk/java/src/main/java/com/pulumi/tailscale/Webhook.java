// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tailscale;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.tailscale.Utilities;
import com.pulumi.tailscale.WebhookArgs;
import com.pulumi.tailscale.inputs.WebhookState;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The webhook resource allows you to configure webhook endpoints for your Tailscale network. See https://tailscale.com/kb/1213/webhooks for more information.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.tailscale.Webhook;
 * import com.pulumi.tailscale.WebhookArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var sampleWebhook = new Webhook("sampleWebhook", WebhookArgs.builder()
 *             .endpointUrl("https://example.com/webhook/endpoint")
 *             .providerType("slack")
 *             .subscriptions(            
 *                 "nodeCreated",
 *                 "userDeleted")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Webhooks can be imported using the endpoint id, e.g.,
 * 
 * ```sh
 * $ pulumi import tailscale:index/webhook:Webhook sample_webhook 123456789
 * ```
 * 
 */
@ResourceType(type="tailscale:index/webhook:Webhook")
public class Webhook extends com.pulumi.resources.CustomResource {
    /**
     * The endpoint to send webhook events to.
     * 
     */
    @Export(name="endpointUrl", refs={String.class}, tree="[0]")
    private Output<String> endpointUrl;

    /**
     * @return The endpoint to send webhook events to.
     * 
     */
    public Output<String> endpointUrl() {
        return this.endpointUrl;
    }
    /**
     * The provider type of the endpoint URL. Also referred to as the &#39;destination&#39; for the webhook in the admin panel. Webhook event payloads are formatted according to the provider type if it is set to a known value. Must be one of `slack`, `mattermost`, `googlechat`, or `discord` if set.
     * 
     */
    @Export(name="providerType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> providerType;

    /**
     * @return The provider type of the endpoint URL. Also referred to as the &#39;destination&#39; for the webhook in the admin panel. Webhook event payloads are formatted according to the provider type if it is set to a known value. Must be one of `slack`, `mattermost`, `googlechat`, or `discord` if set.
     * 
     */
    public Output<Optional<String>> providerType() {
        return Codegen.optional(this.providerType);
    }
    /**
     * The secret used for signing webhook payloads. Only set on resource creation. See https://tailscale.com/kb/1213/webhooks#webhook-secret for more information.
     * 
     */
    @Export(name="secret", refs={String.class}, tree="[0]")
    private Output<String> secret;

    /**
     * @return The secret used for signing webhook payloads. Only set on resource creation. See https://tailscale.com/kb/1213/webhooks#webhook-secret for more information.
     * 
     */
    public Output<String> secret() {
        return this.secret;
    }
    /**
     * The Tailscale events to subscribe this webhook to. See https://tailscale.com/kb/1213/webhooks#events for the list of valid events.
     * 
     */
    @Export(name="subscriptions", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> subscriptions;

    /**
     * @return The Tailscale events to subscribe this webhook to. See https://tailscale.com/kb/1213/webhooks#events for the list of valid events.
     * 
     */
    public Output<List<String>> subscriptions() {
        return this.subscriptions;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Webhook(java.lang.String name) {
        this(name, WebhookArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Webhook(java.lang.String name, WebhookArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Webhook(java.lang.String name, WebhookArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("tailscale:index/webhook:Webhook", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Webhook(java.lang.String name, Output<java.lang.String> id, @Nullable WebhookState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("tailscale:index/webhook:Webhook", name, state, makeResourceOptions(options, id), false);
    }

    private static WebhookArgs makeArgs(WebhookArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? WebhookArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "secret"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Webhook get(java.lang.String name, Output<java.lang.String> id, @Nullable WebhookState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Webhook(name, id, state, options);
    }
}