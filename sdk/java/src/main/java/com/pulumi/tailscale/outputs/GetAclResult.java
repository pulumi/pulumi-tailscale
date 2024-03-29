// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tailscale.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetAclResult {
    /**
     * @return The contents of Tailscale ACL as a HuJSON string
     * 
     */
    private String hujson;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The contents of Tailscale ACL as a JSON string
     * 
     */
    private String json;

    private GetAclResult() {}
    /**
     * @return The contents of Tailscale ACL as a HuJSON string
     * 
     */
    public String hujson() {
        return this.hujson;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The contents of Tailscale ACL as a JSON string
     * 
     */
    public String json() {
        return this.json;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAclResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String hujson;
        private String id;
        private String json;
        public Builder() {}
        public Builder(GetAclResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.hujson = defaults.hujson;
    	      this.id = defaults.id;
    	      this.json = defaults.json;
        }

        @CustomType.Setter
        public Builder hujson(String hujson) {
            if (hujson == null) {
              throw new MissingRequiredPropertyException("GetAclResult", "hujson");
            }
            this.hujson = hujson;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetAclResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder json(String json) {
            if (json == null) {
              throw new MissingRequiredPropertyException("GetAclResult", "json");
            }
            this.json = json;
            return this;
        }
        public GetAclResult build() {
            final var _resultValue = new GetAclResult();
            _resultValue.hujson = hujson;
            _resultValue.id = id;
            _resultValue.json = json;
            return _resultValue;
        }
    }
}
