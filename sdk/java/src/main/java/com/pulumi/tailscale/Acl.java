// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tailscale;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.tailscale.AclArgs;
import com.pulumi.tailscale.Utilities;
import com.pulumi.tailscale.inputs.AclState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The acl resource allows you to configure a Tailscale ACL. See https://tailscale.com/kb/1018/acls for more information. Note that this resource will completely overwrite existing ACL contents for a given tailnet.
 * 
 * If tests are defined in the ACL (the top-level &#34;tests&#34; section), ACL validation will occur before creation and update operations are applied.
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
 * import com.pulumi.tailscale.Acl;
 * import com.pulumi.tailscale.AclArgs;
 * import static com.pulumi.codegen.internal.Serialization.*;
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
 *         var asJson = new Acl("asJson", AclArgs.builder()
 *             .acl(serializeJson(
 *                 jsonObject(
 *                     jsonProperty("acls", jsonArray(jsonObject(
 *                         jsonProperty("action", "accept"),
 *                         jsonProperty("users", jsonArray("*")),
 *                         jsonProperty("ports", jsonArray("*:*"))
 *                     )))
 *                 )))
 *             .build());
 * 
 *         var asHujson = new Acl("asHujson", AclArgs.builder()
 *             .acl("""
 *   {
 *     // Comments in HuJSON policy are preserved when the policy is applied.
 *     "acls": [
 *       {
 *         // Allow all users access to all ports.
 *         action = "accept",
 *         users  = ["*"],
 *         ports  = ["*:*"],
 *       },
 *     ],
 *   }
 *             """)
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
 * $ pulumi import tailscale:index/acl:Acl sample_acl acl
 * ```
 * 
 */
@ResourceType(type="tailscale:index/acl:Acl")
public class Acl extends com.pulumi.resources.CustomResource {
    /**
     * The policy that defines which devices and users are allowed to connect in your network. Can be either a JSON or a HuJSON string.
     * 
     */
    @Export(name="acl", refs={String.class}, tree="[0]")
    private Output<String> acl;

    /**
     * @return The policy that defines which devices and users are allowed to connect in your network. Can be either a JSON or a HuJSON string.
     * 
     */
    public Output<String> acl() {
        return this.acl;
    }
    /**
     * If true, will skip requirement to import acl before allowing changes. Be careful, can cause ACL to be overwritten
     * 
     */
    @Export(name="overwriteExistingContent", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> overwriteExistingContent;

    /**
     * @return If true, will skip requirement to import acl before allowing changes. Be careful, can cause ACL to be overwritten
     * 
     */
    public Output<Optional<Boolean>> overwriteExistingContent() {
        return Codegen.optional(this.overwriteExistingContent);
    }
    /**
     * If true, will reset the ACL for the Tailnet to the default when this resource is destroyed
     * 
     */
    @Export(name="resetAclOnDestroy", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> resetAclOnDestroy;

    /**
     * @return If true, will reset the ACL for the Tailnet to the default when this resource is destroyed
     * 
     */
    public Output<Optional<Boolean>> resetAclOnDestroy() {
        return Codegen.optional(this.resetAclOnDestroy);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Acl(java.lang.String name) {
        this(name, AclArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Acl(java.lang.String name, AclArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Acl(java.lang.String name, AclArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("tailscale:index/acl:Acl", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Acl(java.lang.String name, Output<java.lang.String> id, @Nullable AclState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("tailscale:index/acl:Acl", name, state, makeResourceOptions(options, id), false);
    }

    private static AclArgs makeArgs(AclArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? AclArgs.Empty : args;
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
    public static Acl get(java.lang.String name, Output<java.lang.String> id, @Nullable AclState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Acl(name, id, state, options);
    }
}
