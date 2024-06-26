// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tailscale.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class Get4Via6PlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final Get4Via6PlainArgs Empty = new Get4Via6PlainArgs();

    /**
     * The IPv4 CIDR to map
     * 
     */
    @Import(name="cidr", required=true)
    private String cidr;

    /**
     * @return The IPv4 CIDR to map
     * 
     */
    public String cidr() {
        return this.cidr;
    }

    /**
     * Site ID (between 0 and 65535)
     * 
     */
    @Import(name="site", required=true)
    private Integer site;

    /**
     * @return Site ID (between 0 and 65535)
     * 
     */
    public Integer site() {
        return this.site;
    }

    private Get4Via6PlainArgs() {}

    private Get4Via6PlainArgs(Get4Via6PlainArgs $) {
        this.cidr = $.cidr;
        this.site = $.site;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(Get4Via6PlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Get4Via6PlainArgs $;

        public Builder() {
            $ = new Get4Via6PlainArgs();
        }

        public Builder(Get4Via6PlainArgs defaults) {
            $ = new Get4Via6PlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param cidr The IPv4 CIDR to map
         * 
         * @return builder
         * 
         */
        public Builder cidr(String cidr) {
            $.cidr = cidr;
            return this;
        }

        /**
         * @param site Site ID (between 0 and 65535)
         * 
         * @return builder
         * 
         */
        public Builder site(Integer site) {
            $.site = site;
            return this;
        }

        public Get4Via6PlainArgs build() {
            if ($.cidr == null) {
                throw new MissingRequiredPropertyException("Get4Via6PlainArgs", "cidr");
            }
            if ($.site == null) {
                throw new MissingRequiredPropertyException("Get4Via6PlainArgs", "site");
            }
            return $;
        }
    }

}
