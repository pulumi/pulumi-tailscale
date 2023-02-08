// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tailscale.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetDevicesDevice {
    private List<String> addresses;
    /**
     * @return The ID of this resource.
     * 
     */
    private String id;
    private String name;
    private List<String> tags;
    private String user;

    private GetDevicesDevice() {}
    public List<String> addresses() {
        return this.addresses;
    }
    /**
     * @return The ID of this resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String name() {
        return this.name;
    }
    public List<String> tags() {
        return this.tags;
    }
    public String user() {
        return this.user;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDevicesDevice defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> addresses;
        private String id;
        private String name;
        private List<String> tags;
        private String user;
        public Builder() {}
        public Builder(GetDevicesDevice defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.addresses = defaults.addresses;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.tags = defaults.tags;
    	      this.user = defaults.user;
        }

        @CustomType.Setter
        public Builder addresses(List<String> addresses) {
            this.addresses = Objects.requireNonNull(addresses);
            return this;
        }
        public Builder addresses(String... addresses) {
            return addresses(List.of(addresses));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder tags(List<String> tags) {
            this.tags = Objects.requireNonNull(tags);
            return this;
        }
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }
        @CustomType.Setter
        public Builder user(String user) {
            this.user = Objects.requireNonNull(user);
            return this;
        }
        public GetDevicesDevice build() {
            final var o = new GetDevicesDevice();
            o.addresses = addresses;
            o.id = id;
            o.name = name;
            o.tags = tags;
            o.user = user;
            return o;
        }
    }
}