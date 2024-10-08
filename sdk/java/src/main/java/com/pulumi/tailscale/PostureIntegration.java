// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tailscale;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.tailscale.PostureIntegrationArgs;
import com.pulumi.tailscale.Utilities;
import com.pulumi.tailscale.inputs.PostureIntegrationState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The posture_integration resource allows you to manage integrations with device posture data providers. See https://tailscale.com/kb/1288/device-posture for more information.
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
 * import com.pulumi.tailscale.PostureIntegration;
 * import com.pulumi.tailscale.PostureIntegrationArgs;
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
 *         var samplePostureIntegration = new PostureIntegration("samplePostureIntegration", PostureIntegrationArgs.builder()
 *             .postureProvider("falcon")
 *             .cloudId("us-1")
 *             .clientId("clientid1")
 *             .clientSecret("test-secret1")
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
 * Posture integration can be imported using the posture integration id, e.g.,
 * 
 * ```sh
 * $ pulumi import tailscale:index/postureIntegration:PostureIntegration sample_posture_integration 123456789
 * ```
 * 
 */
@ResourceType(type="tailscale:index/postureIntegration:PostureIntegration")
public class PostureIntegration extends com.pulumi.resources.CustomResource {
    /**
     * Unique identifier for your client.
     * 
     */
    @Export(name="clientId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientId;

    /**
     * @return Unique identifier for your client.
     * 
     */
    public Output<Optional<String>> clientId() {
        return Codegen.optional(this.clientId);
    }
    /**
     * The secret (auth key, token, etc.) used to authenticate with the provider.
     * 
     */
    @Export(name="clientSecret", refs={String.class}, tree="[0]")
    private Output<String> clientSecret;

    /**
     * @return The secret (auth key, token, etc.) used to authenticate with the provider.
     * 
     */
    public Output<String> clientSecret() {
        return this.clientSecret;
    }
    /**
     * Identifies which of the provider&#39;s clouds to integrate with.
     * 
     */
    @Export(name="cloudId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> cloudId;

    /**
     * @return Identifies which of the provider&#39;s clouds to integrate with.
     * 
     */
    public Output<Optional<String>> cloudId() {
        return Codegen.optional(this.cloudId);
    }
    /**
     * The type of posture integration data provider.
     * 
     */
    @Export(name="postureProvider", refs={String.class}, tree="[0]")
    private Output<String> postureProvider;

    /**
     * @return The type of posture integration data provider.
     * 
     */
    public Output<String> postureProvider() {
        return this.postureProvider;
    }
    /**
     * The Microsoft Intune directory (tenant) ID. For other providers, this is left blank.
     * 
     */
    @Export(name="tenantId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> tenantId;

    /**
     * @return The Microsoft Intune directory (tenant) ID. For other providers, this is left blank.
     * 
     */
    public Output<Optional<String>> tenantId() {
        return Codegen.optional(this.tenantId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PostureIntegration(java.lang.String name) {
        this(name, PostureIntegrationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PostureIntegration(java.lang.String name, PostureIntegrationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PostureIntegration(java.lang.String name, PostureIntegrationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("tailscale:index/postureIntegration:PostureIntegration", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private PostureIntegration(java.lang.String name, Output<java.lang.String> id, @Nullable PostureIntegrationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("tailscale:index/postureIntegration:PostureIntegration", name, state, makeResourceOptions(options, id), false);
    }

    private static PostureIntegrationArgs makeArgs(PostureIntegrationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? PostureIntegrationArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
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
    public static PostureIntegration get(java.lang.String name, Output<java.lang.String> id, @Nullable PostureIntegrationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PostureIntegration(name, id, state, options);
    }
}
