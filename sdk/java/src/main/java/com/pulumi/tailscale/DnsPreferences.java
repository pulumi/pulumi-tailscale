// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tailscale;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.tailscale.DnsPreferencesArgs;
import com.pulumi.tailscale.Utilities;
import com.pulumi.tailscale.inputs.DnsPreferencesState;
import java.lang.Boolean;
import javax.annotation.Nullable;

/**
 * The dns_preferences resource allows you to configure DNS preferences for your Tailscale network. See https://tailscale.com/kb/1054/dns for more information.
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
 * import com.pulumi.tailscale.DnsPreferences;
 * import com.pulumi.tailscale.DnsPreferencesArgs;
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
 *         var samplePreferences = new DnsPreferences("samplePreferences", DnsPreferencesArgs.builder()
 *             .magicDns(true)
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
 * ID doesn&#39;t matter.
 * 
 * ```sh
 * $ pulumi import tailscale:index/dnsPreferences:DnsPreferences sample_preferences dns_preferences
 * ```
 * 
 */
@ResourceType(type="tailscale:index/dnsPreferences:DnsPreferences")
public class DnsPreferences extends com.pulumi.resources.CustomResource {
    /**
     * Whether or not to enable magic DNS
     * 
     */
    @Export(name="magicDns", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> magicDns;

    /**
     * @return Whether or not to enable magic DNS
     * 
     */
    public Output<Boolean> magicDns() {
        return this.magicDns;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DnsPreferences(java.lang.String name) {
        this(name, DnsPreferencesArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DnsPreferences(java.lang.String name, DnsPreferencesArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DnsPreferences(java.lang.String name, DnsPreferencesArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("tailscale:index/dnsPreferences:DnsPreferences", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private DnsPreferences(java.lang.String name, Output<java.lang.String> id, @Nullable DnsPreferencesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("tailscale:index/dnsPreferences:DnsPreferences", name, state, makeResourceOptions(options, id), false);
    }

    private static DnsPreferencesArgs makeArgs(DnsPreferencesArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? DnsPreferencesArgs.Empty : args;
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
    public static DnsPreferences get(java.lang.String name, Output<java.lang.String> id, @Nullable DnsPreferencesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DnsPreferences(name, id, state, options);
    }
}
