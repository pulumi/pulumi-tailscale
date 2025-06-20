// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The awsExternalId resource allows you to mint an AWS External ID that Tailscale can use to assume an AWS IAM role that you create for the purposes of allowing Tailscale to stream logs to your S3 bucket. See the logstreamConfiguration resource for more details.
 */
export class AwsExternalId extends pulumi.CustomResource {
    /**
     * Get an existing AwsExternalId resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AwsExternalIdState, opts?: pulumi.CustomResourceOptions): AwsExternalId {
        return new AwsExternalId(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tailscale:index/awsExternalId:AwsExternalId';

    /**
     * Returns true if the given object is an instance of AwsExternalId.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AwsExternalId {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AwsExternalId.__pulumiType;
    }

    /**
     * The External ID that Tailscale will supply when assuming your role. You must reference this in your IAM role's trust policy. See https://docs.aws.amazon.com/IAM/latest/UserGuide/id*roles*common-scenarios_third-party.html for more information on external IDs.
     */
    public /*out*/ readonly externalId!: pulumi.Output<string>;
    /**
     * The AWS account from which Tailscale will assume your role. You must reference this in your IAM role's trust policy. See https://docs.aws.amazon.com/IAM/latest/UserGuide/id*roles*common-scenarios_third-party.html for more information on external IDs.
     */
    public /*out*/ readonly tailscaleAwsAccountId!: pulumi.Output<string>;

    /**
     * Create a AwsExternalId resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AwsExternalIdArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AwsExternalIdArgs | AwsExternalIdState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AwsExternalIdState | undefined;
            resourceInputs["externalId"] = state ? state.externalId : undefined;
            resourceInputs["tailscaleAwsAccountId"] = state ? state.tailscaleAwsAccountId : undefined;
        } else {
            const args = argsOrState as AwsExternalIdArgs | undefined;
            resourceInputs["externalId"] = undefined /*out*/;
            resourceInputs["tailscaleAwsAccountId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AwsExternalId.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AwsExternalId resources.
 */
export interface AwsExternalIdState {
    /**
     * The External ID that Tailscale will supply when assuming your role. You must reference this in your IAM role's trust policy. See https://docs.aws.amazon.com/IAM/latest/UserGuide/id*roles*common-scenarios_third-party.html for more information on external IDs.
     */
    externalId?: pulumi.Input<string>;
    /**
     * The AWS account from which Tailscale will assume your role. You must reference this in your IAM role's trust policy. See https://docs.aws.amazon.com/IAM/latest/UserGuide/id*roles*common-scenarios_third-party.html for more information on external IDs.
     */
    tailscaleAwsAccountId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AwsExternalId resource.
 */
export interface AwsExternalIdArgs {
}
