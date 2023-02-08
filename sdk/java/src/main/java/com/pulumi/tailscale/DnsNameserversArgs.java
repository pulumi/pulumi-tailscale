// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tailscale;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class DnsNameserversArgs extends com.pulumi.resources.ResourceArgs {

    public static final DnsNameserversArgs Empty = new DnsNameserversArgs();

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

    private DnsNameserversArgs() {}

    private DnsNameserversArgs(DnsNameserversArgs $) {
        this.nameservers = $.nameservers;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DnsNameserversArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DnsNameserversArgs $;

        public Builder() {
            $ = new DnsNameserversArgs();
        }

        public Builder(DnsNameserversArgs defaults) {
            $ = new DnsNameserversArgs(Objects.requireNonNull(defaults));
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

        public DnsNameserversArgs build() {
            $.nameservers = Objects.requireNonNull($.nameservers, "expected parameter 'nameservers' to be non-null");
            return $;
        }
    }

}