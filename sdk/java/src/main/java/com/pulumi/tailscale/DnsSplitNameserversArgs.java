// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tailscale;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class DnsSplitNameserversArgs extends com.pulumi.resources.ResourceArgs {

    public static final DnsSplitNameserversArgs Empty = new DnsSplitNameserversArgs();

    /**
     * Domain to configure split DNS for. Requests for this domain will be resolved using the provided nameservers.
     * 
     */
    @Import(name="domain", required=true)
    private Output<String> domain;

    /**
     * @return Domain to configure split DNS for. Requests for this domain will be resolved using the provided nameservers.
     * 
     */
    public Output<String> domain() {
        return this.domain;
    }

    /**
     * Devices on your network will use these nameservers to resolve DNS names. IPv4 or IPv6 addresses are accepted.
     * 
     */
    @Import(name="nameservers", required=true)
    private Output<List<String>> nameservers;

    /**
     * @return Devices on your network will use these nameservers to resolve DNS names. IPv4 or IPv6 addresses are accepted.
     * 
     */
    public Output<List<String>> nameservers() {
        return this.nameservers;
    }

    private DnsSplitNameserversArgs() {}

    private DnsSplitNameserversArgs(DnsSplitNameserversArgs $) {
        this.domain = $.domain;
        this.nameservers = $.nameservers;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DnsSplitNameserversArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DnsSplitNameserversArgs $;

        public Builder() {
            $ = new DnsSplitNameserversArgs();
        }

        public Builder(DnsSplitNameserversArgs defaults) {
            $ = new DnsSplitNameserversArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param domain Domain to configure split DNS for. Requests for this domain will be resolved using the provided nameservers.
         * 
         * @return builder
         * 
         */
        public Builder domain(Output<String> domain) {
            $.domain = domain;
            return this;
        }

        /**
         * @param domain Domain to configure split DNS for. Requests for this domain will be resolved using the provided nameservers.
         * 
         * @return builder
         * 
         */
        public Builder domain(String domain) {
            return domain(Output.of(domain));
        }

        /**
         * @param nameservers Devices on your network will use these nameservers to resolve DNS names. IPv4 or IPv6 addresses are accepted.
         * 
         * @return builder
         * 
         */
        public Builder nameservers(Output<List<String>> nameservers) {
            $.nameservers = nameservers;
            return this;
        }

        /**
         * @param nameservers Devices on your network will use these nameservers to resolve DNS names. IPv4 or IPv6 addresses are accepted.
         * 
         * @return builder
         * 
         */
        public Builder nameservers(List<String> nameservers) {
            return nameservers(Output.of(nameservers));
        }

        /**
         * @param nameservers Devices on your network will use these nameservers to resolve DNS names. IPv4 or IPv6 addresses are accepted.
         * 
         * @return builder
         * 
         */
        public Builder nameservers(String... nameservers) {
            return nameservers(List.of(nameservers));
        }

        public DnsSplitNameserversArgs build() {
            if ($.domain == null) {
                throw new MissingRequiredPropertyException("DnsSplitNameserversArgs", "domain");
            }
            if ($.nameservers == null) {
                throw new MissingRequiredPropertyException("DnsSplitNameserversArgs", "nameservers");
            }
            return $;
        }
    }

}