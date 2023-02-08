// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tailscale;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class DeviceTagsArgs extends com.pulumi.resources.ResourceArgs {

    public static final DeviceTagsArgs Empty = new DeviceTagsArgs();

    /**
     * The device to set tags for
     * 
     */
    @Import(name="deviceId", required=true)
    private Output<String> deviceId;

    /**
     * @return The device to set tags for
     * 
     */
    public Output<String> deviceId() {
        return this.deviceId;
    }

    /**
     * The tags to apply to the device
     * 
     */
    @Import(name="tags", required=true)
    private Output<List<String>> tags;

    /**
     * @return The tags to apply to the device
     * 
     */
    public Output<List<String>> tags() {
        return this.tags;
    }

    private DeviceTagsArgs() {}

    private DeviceTagsArgs(DeviceTagsArgs $) {
        this.deviceId = $.deviceId;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DeviceTagsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DeviceTagsArgs $;

        public Builder() {
            $ = new DeviceTagsArgs();
        }

        public Builder(DeviceTagsArgs defaults) {
            $ = new DeviceTagsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param deviceId The device to set tags for
         * 
         * @return builder
         * 
         */
        public Builder deviceId(Output<String> deviceId) {
            $.deviceId = deviceId;
            return this;
        }

        /**
         * @param deviceId The device to set tags for
         * 
         * @return builder
         * 
         */
        public Builder deviceId(String deviceId) {
            return deviceId(Output.of(deviceId));
        }

        /**
         * @param tags The tags to apply to the device
         * 
         * @return builder
         * 
         */
        public Builder tags(Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags The tags to apply to the device
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags The tags to apply to the device
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        public DeviceTagsArgs build() {
            $.deviceId = Objects.requireNonNull($.deviceId, "expected parameter 'deviceId' to be non-null");
            $.tags = Objects.requireNonNull($.tags, "expected parameter 'tags' to be non-null");
            return $;
        }
    }

}