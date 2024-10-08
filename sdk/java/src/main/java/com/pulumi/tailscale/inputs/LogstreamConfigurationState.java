// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tailscale.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LogstreamConfigurationState extends com.pulumi.resources.ResourceArgs {

    public static final LogstreamConfigurationState Empty = new LogstreamConfigurationState();

    /**
     * The type of system to which logs are being streamed.
     * 
     */
    @Import(name="destinationType")
    private @Nullable Output<String> destinationType;

    /**
     * @return The type of system to which logs are being streamed.
     * 
     */
    public Optional<Output<String>> destinationType() {
        return Optional.ofNullable(this.destinationType);
    }

    /**
     * The type of log that is streamed to this endpoint.
     * 
     */
    @Import(name="logType")
    private @Nullable Output<String> logType;

    /**
     * @return The type of log that is streamed to this endpoint.
     * 
     */
    public Optional<Output<String>> logType() {
        return Optional.ofNullable(this.logType);
    }

    /**
     * The token/password with which log streams to this endpoint should be authenticated.
     * 
     */
    @Import(name="token")
    private @Nullable Output<String> token;

    /**
     * @return The token/password with which log streams to this endpoint should be authenticated.
     * 
     */
    public Optional<Output<String>> token() {
        return Optional.ofNullable(this.token);
    }

    /**
     * The URL to which log streams are being posted.
     * 
     */
    @Import(name="url")
    private @Nullable Output<String> url;

    /**
     * @return The URL to which log streams are being posted.
     * 
     */
    public Optional<Output<String>> url() {
        return Optional.ofNullable(this.url);
    }

    /**
     * The username with which log streams to this endpoint are authenticated. Only required if destination_type is &#39;elastic&#39;, defaults to &#39;user&#39; if not set.
     * 
     */
    @Import(name="user")
    private @Nullable Output<String> user;

    /**
     * @return The username with which log streams to this endpoint are authenticated. Only required if destination_type is &#39;elastic&#39;, defaults to &#39;user&#39; if not set.
     * 
     */
    public Optional<Output<String>> user() {
        return Optional.ofNullable(this.user);
    }

    private LogstreamConfigurationState() {}

    private LogstreamConfigurationState(LogstreamConfigurationState $) {
        this.destinationType = $.destinationType;
        this.logType = $.logType;
        this.token = $.token;
        this.url = $.url;
        this.user = $.user;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LogstreamConfigurationState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LogstreamConfigurationState $;

        public Builder() {
            $ = new LogstreamConfigurationState();
        }

        public Builder(LogstreamConfigurationState defaults) {
            $ = new LogstreamConfigurationState(Objects.requireNonNull(defaults));
        }

        /**
         * @param destinationType The type of system to which logs are being streamed.
         * 
         * @return builder
         * 
         */
        public Builder destinationType(@Nullable Output<String> destinationType) {
            $.destinationType = destinationType;
            return this;
        }

        /**
         * @param destinationType The type of system to which logs are being streamed.
         * 
         * @return builder
         * 
         */
        public Builder destinationType(String destinationType) {
            return destinationType(Output.of(destinationType));
        }

        /**
         * @param logType The type of log that is streamed to this endpoint.
         * 
         * @return builder
         * 
         */
        public Builder logType(@Nullable Output<String> logType) {
            $.logType = logType;
            return this;
        }

        /**
         * @param logType The type of log that is streamed to this endpoint.
         * 
         * @return builder
         * 
         */
        public Builder logType(String logType) {
            return logType(Output.of(logType));
        }

        /**
         * @param token The token/password with which log streams to this endpoint should be authenticated.
         * 
         * @return builder
         * 
         */
        public Builder token(@Nullable Output<String> token) {
            $.token = token;
            return this;
        }

        /**
         * @param token The token/password with which log streams to this endpoint should be authenticated.
         * 
         * @return builder
         * 
         */
        public Builder token(String token) {
            return token(Output.of(token));
        }

        /**
         * @param url The URL to which log streams are being posted.
         * 
         * @return builder
         * 
         */
        public Builder url(@Nullable Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url The URL to which log streams are being posted.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        /**
         * @param user The username with which log streams to this endpoint are authenticated. Only required if destination_type is &#39;elastic&#39;, defaults to &#39;user&#39; if not set.
         * 
         * @return builder
         * 
         */
        public Builder user(@Nullable Output<String> user) {
            $.user = user;
            return this;
        }

        /**
         * @param user The username with which log streams to this endpoint are authenticated. Only required if destination_type is &#39;elastic&#39;, defaults to &#39;user&#39; if not set.
         * 
         * @return builder
         * 
         */
        public Builder user(String user) {
            return user(Output.of(user));
        }

        public LogstreamConfigurationState build() {
            return $;
        }
    }

}
