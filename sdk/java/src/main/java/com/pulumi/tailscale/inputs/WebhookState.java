// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tailscale.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class WebhookState extends com.pulumi.resources.ResourceArgs {

    public static final WebhookState Empty = new WebhookState();

    /**
     * The endpoint to send webhook events to.
     * 
     */
    @Import(name="endpointUrl")
    private @Nullable Output<String> endpointUrl;

    /**
     * @return The endpoint to send webhook events to.
     * 
     */
    public Optional<Output<String>> endpointUrl() {
        return Optional.ofNullable(this.endpointUrl);
    }

    /**
     * The provider type of the endpoint URL. Also referred to as the &#39;destination&#39; for the webhook in the admin panel. Webhook event payloads are formatted according to the provider type if it is set to a known value. Must be one of `slack`, `mattermost`, `googlechat`, or `discord` if set.
     * 
     */
    @Import(name="providerType")
    private @Nullable Output<String> providerType;

    /**
     * @return The provider type of the endpoint URL. Also referred to as the &#39;destination&#39; for the webhook in the admin panel. Webhook event payloads are formatted according to the provider type if it is set to a known value. Must be one of `slack`, `mattermost`, `googlechat`, or `discord` if set.
     * 
     */
    public Optional<Output<String>> providerType() {
        return Optional.ofNullable(this.providerType);
    }

    /**
     * The secret used for signing webhook payloads. Only set on resource creation. See https://tailscale.com/kb/1213/webhooks#webhook-secret for more information.
     * 
     */
    @Import(name="secret")
    private @Nullable Output<String> secret;

    /**
     * @return The secret used for signing webhook payloads. Only set on resource creation. See https://tailscale.com/kb/1213/webhooks#webhook-secret for more information.
     * 
     */
    public Optional<Output<String>> secret() {
        return Optional.ofNullable(this.secret);
    }

    /**
     * The Tailscale events to subscribe this webhook to. See https://tailscale.com/kb/1213/webhooks#events for the list of valid events.
     * 
     */
    @Import(name="subscriptions")
    private @Nullable Output<List<String>> subscriptions;

    /**
     * @return The Tailscale events to subscribe this webhook to. See https://tailscale.com/kb/1213/webhooks#events for the list of valid events.
     * 
     */
    public Optional<Output<List<String>>> subscriptions() {
        return Optional.ofNullable(this.subscriptions);
    }

    private WebhookState() {}

    private WebhookState(WebhookState $) {
        this.endpointUrl = $.endpointUrl;
        this.providerType = $.providerType;
        this.secret = $.secret;
        this.subscriptions = $.subscriptions;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(WebhookState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private WebhookState $;

        public Builder() {
            $ = new WebhookState();
        }

        public Builder(WebhookState defaults) {
            $ = new WebhookState(Objects.requireNonNull(defaults));
        }

        /**
         * @param endpointUrl The endpoint to send webhook events to.
         * 
         * @return builder
         * 
         */
        public Builder endpointUrl(@Nullable Output<String> endpointUrl) {
            $.endpointUrl = endpointUrl;
            return this;
        }

        /**
         * @param endpointUrl The endpoint to send webhook events to.
         * 
         * @return builder
         * 
         */
        public Builder endpointUrl(String endpointUrl) {
            return endpointUrl(Output.of(endpointUrl));
        }

        /**
         * @param providerType The provider type of the endpoint URL. Also referred to as the &#39;destination&#39; for the webhook in the admin panel. Webhook event payloads are formatted according to the provider type if it is set to a known value. Must be one of `slack`, `mattermost`, `googlechat`, or `discord` if set.
         * 
         * @return builder
         * 
         */
        public Builder providerType(@Nullable Output<String> providerType) {
            $.providerType = providerType;
            return this;
        }

        /**
         * @param providerType The provider type of the endpoint URL. Also referred to as the &#39;destination&#39; for the webhook in the admin panel. Webhook event payloads are formatted according to the provider type if it is set to a known value. Must be one of `slack`, `mattermost`, `googlechat`, or `discord` if set.
         * 
         * @return builder
         * 
         */
        public Builder providerType(String providerType) {
            return providerType(Output.of(providerType));
        }

        /**
         * @param secret The secret used for signing webhook payloads. Only set on resource creation. See https://tailscale.com/kb/1213/webhooks#webhook-secret for more information.
         * 
         * @return builder
         * 
         */
        public Builder secret(@Nullable Output<String> secret) {
            $.secret = secret;
            return this;
        }

        /**
         * @param secret The secret used for signing webhook payloads. Only set on resource creation. See https://tailscale.com/kb/1213/webhooks#webhook-secret for more information.
         * 
         * @return builder
         * 
         */
        public Builder secret(String secret) {
            return secret(Output.of(secret));
        }

        /**
         * @param subscriptions The Tailscale events to subscribe this webhook to. See https://tailscale.com/kb/1213/webhooks#events for the list of valid events.
         * 
         * @return builder
         * 
         */
        public Builder subscriptions(@Nullable Output<List<String>> subscriptions) {
            $.subscriptions = subscriptions;
            return this;
        }

        /**
         * @param subscriptions The Tailscale events to subscribe this webhook to. See https://tailscale.com/kb/1213/webhooks#events for the list of valid events.
         * 
         * @return builder
         * 
         */
        public Builder subscriptions(List<String> subscriptions) {
            return subscriptions(Output.of(subscriptions));
        }

        /**
         * @param subscriptions The Tailscale events to subscribe this webhook to. See https://tailscale.com/kb/1213/webhooks#events for the list of valid events.
         * 
         * @return builder
         * 
         */
        public Builder subscriptions(String... subscriptions) {
            return subscriptions(List.of(subscriptions));
        }

        public WebhookState build() {
            return $;
        }
    }

}
