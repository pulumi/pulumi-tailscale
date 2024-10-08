// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tailscale;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.tailscale.DnsNameserversArgs;
import com.pulumi.tailscale.Utilities;
import com.pulumi.tailscale.inputs.DnsNameserversState;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * The dns_nameservers resource allows you to configure DNS nameservers for your Tailscale network. See https://tailscale.com/kb/1054/dns for more information.
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
 * import com.pulumi.tailscale.DnsNameservers;
 * import com.pulumi.tailscale.DnsNameserversArgs;
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
 *         var sampleNameservers = new DnsNameservers("sampleNameservers", DnsNameserversArgs.builder()
 *             .nameservers(            
 *                 "8.8.8.8",
 *                 "8.8.4.4")
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
 * $ pulumi import tailscale:index/dnsNameservers:DnsNameservers sample dns_nameservers
 * ```
 * 
 */
@ResourceType(type="tailscale:index/dnsNameservers:DnsNameservers")
public class DnsNameservers extends com.pulumi.resources.CustomResource {
    /**
     * Devices on your network will use these nameservers to resolve DNS names. IPv4 or IPv6 addresses are accepted.
     * 
     */
    @Export(name="nameservers", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> nameservers;

    /**
     * @return Devices on your network will use these nameservers to resolve DNS names. IPv4 or IPv6 addresses are accepted.
     * 
     */
    public Output<List<String>> nameservers() {
        return this.nameservers;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DnsNameservers(java.lang.String name) {
        this(name, DnsNameserversArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DnsNameservers(java.lang.String name, DnsNameserversArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DnsNameservers(java.lang.String name, DnsNameserversArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("tailscale:index/dnsNameservers:DnsNameservers", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private DnsNameservers(java.lang.String name, Output<java.lang.String> id, @Nullable DnsNameserversState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("tailscale:index/dnsNameservers:DnsNameservers", name, state, makeResourceOptions(options, id), false);
    }

    private static DnsNameserversArgs makeArgs(DnsNameserversArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? DnsNameserversArgs.Empty : args;
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
    public static DnsNameservers get(java.lang.String name, Output<java.lang.String> id, @Nullable DnsNameserversState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DnsNameservers(name, id, state, options);
    }
}
