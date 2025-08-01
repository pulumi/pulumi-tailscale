// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The dnsNameservers resource allows you to configure DNS nameservers for your Tailscale network. See https://tailscale.com/kb/1054/dns for more information.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tailscale from "@pulumi/tailscale";
 *
 * const sampleNameservers = new tailscale.DnsNameservers("sample_nameservers", {nameservers: [
 *     "8.8.8.8",
 *     "8.8.4.4",
 * ]});
 * ```
 *
 * ## Import
 *
 * ID doesn't matter.
 *
 * ```sh
 * $ pulumi import tailscale:index/dnsNameservers:DnsNameservers sample dns_nameservers
 * ```
 */
export class DnsNameservers extends pulumi.CustomResource {
    /**
     * Get an existing DnsNameservers resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DnsNameserversState, opts?: pulumi.CustomResourceOptions): DnsNameservers {
        return new DnsNameservers(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tailscale:index/dnsNameservers:DnsNameservers';

    /**
     * Returns true if the given object is an instance of DnsNameservers.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DnsNameservers {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DnsNameservers.__pulumiType;
    }

    /**
     * Devices on your network will use these nameservers to resolve DNS names. IPv4 or IPv6 addresses are accepted.
     */
    public readonly nameservers!: pulumi.Output<string[]>;

    /**
     * Create a DnsNameservers resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DnsNameserversArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DnsNameserversArgs | DnsNameserversState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DnsNameserversState | undefined;
            resourceInputs["nameservers"] = state ? state.nameservers : undefined;
        } else {
            const args = argsOrState as DnsNameserversArgs | undefined;
            if ((!args || args.nameservers === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nameservers'");
            }
            resourceInputs["nameservers"] = args ? args.nameservers : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DnsNameservers.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DnsNameservers resources.
 */
export interface DnsNameserversState {
    /**
     * Devices on your network will use these nameservers to resolve DNS names. IPv4 or IPv6 addresses are accepted.
     */
    nameservers?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a DnsNameservers resource.
 */
export interface DnsNameserversArgs {
    /**
     * Devices on your network will use these nameservers to resolve DNS names. IPv4 or IPv6 addresses are accepted.
     */
    nameservers: pulumi.Input<pulumi.Input<string>[]>;
}
