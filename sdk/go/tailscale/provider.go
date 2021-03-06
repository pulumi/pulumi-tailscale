// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tailscale

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provider type for the tailscale package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	// The API key to use for authenticating requests to the API. Can be set via the TAILSCALE_API_KEY environment variable.
	ApiKey pulumi.StringPtrOutput `pulumi:"apiKey"`
	// The base URL of the Tailscale API. Defaults to https://api.tailscale.com. Can be set via the TAILSCALE_BASE_URL
	// environment variable.
	BaseUrl pulumi.StringPtrOutput `pulumi:"baseUrl"`
	// The Tailnet to perform actions in. Can be set via the TAILSCALE_TAILNET environment variable.
	Tailnet pulumi.StringPtrOutput `pulumi:"tailnet"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}

	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:tailscale", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// The API key to use for authenticating requests to the API. Can be set via the TAILSCALE_API_KEY environment variable.
	ApiKey *string `pulumi:"apiKey"`
	// The base URL of the Tailscale API. Defaults to https://api.tailscale.com. Can be set via the TAILSCALE_BASE_URL
	// environment variable.
	BaseUrl *string `pulumi:"baseUrl"`
	// The Tailnet to perform actions in. Can be set via the TAILSCALE_TAILNET environment variable.
	Tailnet *string `pulumi:"tailnet"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// The API key to use for authenticating requests to the API. Can be set via the TAILSCALE_API_KEY environment variable.
	ApiKey pulumi.StringPtrInput
	// The base URL of the Tailscale API. Defaults to https://api.tailscale.com. Can be set via the TAILSCALE_BASE_URL
	// environment variable.
	BaseUrl pulumi.StringPtrInput
	// The Tailnet to perform actions in. Can be set via the TAILSCALE_TAILNET environment variable.
	Tailnet pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}
