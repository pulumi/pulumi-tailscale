// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The logstreamConfiguration resource allows you to configure streaming configuration or network flow logs to a supported security information and event management (SIEM) system. See https://tailscale.com/kb/1255/log-streaming for more information.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tailscale from "@pulumi/tailscale";
 *
 * const sampleLogstreamConfiguration = new tailscale.LogstreamConfiguration("sample_logstream_configuration", {
 *     logType: "configuration",
 *     destinationType: "panther",
 *     url: "https://example.com",
 *     token: "some-token",
 * });
 * ```
 *
 * ## Import
 *
 * Logstream configuration can be imported using the logstream configuration id, e.g.,
 *
 * ```sh
 * $ pulumi import tailscale:index/logstreamConfiguration:LogstreamConfiguration sample_logstream_configuration 123456789
 * ```
 */
export class LogstreamConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing LogstreamConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LogstreamConfigurationState, opts?: pulumi.CustomResourceOptions): LogstreamConfiguration {
        return new LogstreamConfiguration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tailscale:index/logstreamConfiguration:LogstreamConfiguration';

    /**
     * Returns true if the given object is an instance of LogstreamConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LogstreamConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LogstreamConfiguration.__pulumiType;
    }

    /**
     * The type of system to which logs are being streamed.
     */
    public readonly destinationType!: pulumi.Output<string>;
    /**
     * The type of log that is streamed to this endpoint.
     */
    public readonly logType!: pulumi.Output<string>;
    /**
     * The token/password with which log streams to this endpoint should be authenticated.
     */
    public readonly token!: pulumi.Output<string>;
    /**
     * The URL to which log streams are being posted.
     */
    public readonly url!: pulumi.Output<string>;
    /**
     * The username with which log streams to this endpoint are authenticated. Only required if destinationType is 'elastic', defaults to 'user' if not set.
     */
    public readonly user!: pulumi.Output<string | undefined>;

    /**
     * Create a LogstreamConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LogstreamConfigurationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LogstreamConfigurationArgs | LogstreamConfigurationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LogstreamConfigurationState | undefined;
            resourceInputs["destinationType"] = state ? state.destinationType : undefined;
            resourceInputs["logType"] = state ? state.logType : undefined;
            resourceInputs["token"] = state ? state.token : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
            resourceInputs["user"] = state ? state.user : undefined;
        } else {
            const args = argsOrState as LogstreamConfigurationArgs | undefined;
            if ((!args || args.destinationType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destinationType'");
            }
            if ((!args || args.logType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'logType'");
            }
            if ((!args || args.token === undefined) && !opts.urn) {
                throw new Error("Missing required property 'token'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["destinationType"] = args ? args.destinationType : undefined;
            resourceInputs["logType"] = args ? args.logType : undefined;
            resourceInputs["token"] = args?.token ? pulumi.secret(args.token) : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["user"] = args ? args.user : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["token"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(LogstreamConfiguration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LogstreamConfiguration resources.
 */
export interface LogstreamConfigurationState {
    /**
     * The type of system to which logs are being streamed.
     */
    destinationType?: pulumi.Input<string>;
    /**
     * The type of log that is streamed to this endpoint.
     */
    logType?: pulumi.Input<string>;
    /**
     * The token/password with which log streams to this endpoint should be authenticated.
     */
    token?: pulumi.Input<string>;
    /**
     * The URL to which log streams are being posted.
     */
    url?: pulumi.Input<string>;
    /**
     * The username with which log streams to this endpoint are authenticated. Only required if destinationType is 'elastic', defaults to 'user' if not set.
     */
    user?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LogstreamConfiguration resource.
 */
export interface LogstreamConfigurationArgs {
    /**
     * The type of system to which logs are being streamed.
     */
    destinationType: pulumi.Input<string>;
    /**
     * The type of log that is streamed to this endpoint.
     */
    logType: pulumi.Input<string>;
    /**
     * The token/password with which log streams to this endpoint should be authenticated.
     */
    token: pulumi.Input<string>;
    /**
     * The URL to which log streams are being posted.
     */
    url: pulumi.Input<string>;
    /**
     * The username with which log streams to this endpoint are authenticated. Only required if destinationType is 'elastic', defaults to 'user' if not set.
     */
    user?: pulumi.Input<string>;
}