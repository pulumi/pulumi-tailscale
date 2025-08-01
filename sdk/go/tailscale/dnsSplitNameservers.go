// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tailscale

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-tailscale/sdk/go/tailscale/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The dnsSplitNameservers resource allows you to configure split DNS nameservers for your Tailscale network. See https://tailscale.com/kb/1054/dns for more information.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tailscale/sdk/go/tailscale"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := tailscale.NewDnsSplitNameservers(ctx, "sample_split_nameservers", &tailscale.DnsSplitNameserversArgs{
//				Domain: pulumi.String("foo.example.com"),
//				Nameservers: pulumi.StringArray{
//					pulumi.String("1.1.1.1"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Split DNS nameservers can be imported using the domain name, e.g.
//
// ```sh
// $ pulumi import tailscale:index/dnsSplitNameservers:DnsSplitNameservers sample_split_nameservers example.com
// ```
type DnsSplitNameservers struct {
	pulumi.CustomResourceState

	// Domain to configure split DNS for. Requests for this domain will be resolved using the provided nameservers. Changing this will force the resource to be recreated.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Devices on your network will use these nameservers to resolve DNS names. IPv4 or IPv6 addresses are accepted.
	Nameservers pulumi.StringArrayOutput `pulumi:"nameservers"`
}

// NewDnsSplitNameservers registers a new resource with the given unique name, arguments, and options.
func NewDnsSplitNameservers(ctx *pulumi.Context,
	name string, args *DnsSplitNameserversArgs, opts ...pulumi.ResourceOption) (*DnsSplitNameservers, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.Nameservers == nil {
		return nil, errors.New("invalid value for required argument 'Nameservers'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DnsSplitNameservers
	err := ctx.RegisterResource("tailscale:index/dnsSplitNameservers:DnsSplitNameservers", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnsSplitNameservers gets an existing DnsSplitNameservers resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnsSplitNameservers(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnsSplitNameserversState, opts ...pulumi.ResourceOption) (*DnsSplitNameservers, error) {
	var resource DnsSplitNameservers
	err := ctx.ReadResource("tailscale:index/dnsSplitNameservers:DnsSplitNameservers", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DnsSplitNameservers resources.
type dnsSplitNameserversState struct {
	// Domain to configure split DNS for. Requests for this domain will be resolved using the provided nameservers. Changing this will force the resource to be recreated.
	Domain *string `pulumi:"domain"`
	// Devices on your network will use these nameservers to resolve DNS names. IPv4 or IPv6 addresses are accepted.
	Nameservers []string `pulumi:"nameservers"`
}

type DnsSplitNameserversState struct {
	// Domain to configure split DNS for. Requests for this domain will be resolved using the provided nameservers. Changing this will force the resource to be recreated.
	Domain pulumi.StringPtrInput
	// Devices on your network will use these nameservers to resolve DNS names. IPv4 or IPv6 addresses are accepted.
	Nameservers pulumi.StringArrayInput
}

func (DnsSplitNameserversState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsSplitNameserversState)(nil)).Elem()
}

type dnsSplitNameserversArgs struct {
	// Domain to configure split DNS for. Requests for this domain will be resolved using the provided nameservers. Changing this will force the resource to be recreated.
	Domain string `pulumi:"domain"`
	// Devices on your network will use these nameservers to resolve DNS names. IPv4 or IPv6 addresses are accepted.
	Nameservers []string `pulumi:"nameservers"`
}

// The set of arguments for constructing a DnsSplitNameservers resource.
type DnsSplitNameserversArgs struct {
	// Domain to configure split DNS for. Requests for this domain will be resolved using the provided nameservers. Changing this will force the resource to be recreated.
	Domain pulumi.StringInput
	// Devices on your network will use these nameservers to resolve DNS names. IPv4 or IPv6 addresses are accepted.
	Nameservers pulumi.StringArrayInput
}

func (DnsSplitNameserversArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsSplitNameserversArgs)(nil)).Elem()
}

type DnsSplitNameserversInput interface {
	pulumi.Input

	ToDnsSplitNameserversOutput() DnsSplitNameserversOutput
	ToDnsSplitNameserversOutputWithContext(ctx context.Context) DnsSplitNameserversOutput
}

func (*DnsSplitNameservers) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsSplitNameservers)(nil)).Elem()
}

func (i *DnsSplitNameservers) ToDnsSplitNameserversOutput() DnsSplitNameserversOutput {
	return i.ToDnsSplitNameserversOutputWithContext(context.Background())
}

func (i *DnsSplitNameservers) ToDnsSplitNameserversOutputWithContext(ctx context.Context) DnsSplitNameserversOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsSplitNameserversOutput)
}

// DnsSplitNameserversArrayInput is an input type that accepts DnsSplitNameserversArray and DnsSplitNameserversArrayOutput values.
// You can construct a concrete instance of `DnsSplitNameserversArrayInput` via:
//
//	DnsSplitNameserversArray{ DnsSplitNameserversArgs{...} }
type DnsSplitNameserversArrayInput interface {
	pulumi.Input

	ToDnsSplitNameserversArrayOutput() DnsSplitNameserversArrayOutput
	ToDnsSplitNameserversArrayOutputWithContext(context.Context) DnsSplitNameserversArrayOutput
}

type DnsSplitNameserversArray []DnsSplitNameserversInput

func (DnsSplitNameserversArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsSplitNameservers)(nil)).Elem()
}

func (i DnsSplitNameserversArray) ToDnsSplitNameserversArrayOutput() DnsSplitNameserversArrayOutput {
	return i.ToDnsSplitNameserversArrayOutputWithContext(context.Background())
}

func (i DnsSplitNameserversArray) ToDnsSplitNameserversArrayOutputWithContext(ctx context.Context) DnsSplitNameserversArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsSplitNameserversArrayOutput)
}

// DnsSplitNameserversMapInput is an input type that accepts DnsSplitNameserversMap and DnsSplitNameserversMapOutput values.
// You can construct a concrete instance of `DnsSplitNameserversMapInput` via:
//
//	DnsSplitNameserversMap{ "key": DnsSplitNameserversArgs{...} }
type DnsSplitNameserversMapInput interface {
	pulumi.Input

	ToDnsSplitNameserversMapOutput() DnsSplitNameserversMapOutput
	ToDnsSplitNameserversMapOutputWithContext(context.Context) DnsSplitNameserversMapOutput
}

type DnsSplitNameserversMap map[string]DnsSplitNameserversInput

func (DnsSplitNameserversMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsSplitNameservers)(nil)).Elem()
}

func (i DnsSplitNameserversMap) ToDnsSplitNameserversMapOutput() DnsSplitNameserversMapOutput {
	return i.ToDnsSplitNameserversMapOutputWithContext(context.Background())
}

func (i DnsSplitNameserversMap) ToDnsSplitNameserversMapOutputWithContext(ctx context.Context) DnsSplitNameserversMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsSplitNameserversMapOutput)
}

type DnsSplitNameserversOutput struct{ *pulumi.OutputState }

func (DnsSplitNameserversOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsSplitNameservers)(nil)).Elem()
}

func (o DnsSplitNameserversOutput) ToDnsSplitNameserversOutput() DnsSplitNameserversOutput {
	return o
}

func (o DnsSplitNameserversOutput) ToDnsSplitNameserversOutputWithContext(ctx context.Context) DnsSplitNameserversOutput {
	return o
}

// Domain to configure split DNS for. Requests for this domain will be resolved using the provided nameservers. Changing this will force the resource to be recreated.
func (o DnsSplitNameserversOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsSplitNameservers) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// Devices on your network will use these nameservers to resolve DNS names. IPv4 or IPv6 addresses are accepted.
func (o DnsSplitNameserversOutput) Nameservers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DnsSplitNameservers) pulumi.StringArrayOutput { return v.Nameservers }).(pulumi.StringArrayOutput)
}

type DnsSplitNameserversArrayOutput struct{ *pulumi.OutputState }

func (DnsSplitNameserversArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsSplitNameservers)(nil)).Elem()
}

func (o DnsSplitNameserversArrayOutput) ToDnsSplitNameserversArrayOutput() DnsSplitNameserversArrayOutput {
	return o
}

func (o DnsSplitNameserversArrayOutput) ToDnsSplitNameserversArrayOutputWithContext(ctx context.Context) DnsSplitNameserversArrayOutput {
	return o
}

func (o DnsSplitNameserversArrayOutput) Index(i pulumi.IntInput) DnsSplitNameserversOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DnsSplitNameservers {
		return vs[0].([]*DnsSplitNameservers)[vs[1].(int)]
	}).(DnsSplitNameserversOutput)
}

type DnsSplitNameserversMapOutput struct{ *pulumi.OutputState }

func (DnsSplitNameserversMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsSplitNameservers)(nil)).Elem()
}

func (o DnsSplitNameserversMapOutput) ToDnsSplitNameserversMapOutput() DnsSplitNameserversMapOutput {
	return o
}

func (o DnsSplitNameserversMapOutput) ToDnsSplitNameserversMapOutputWithContext(ctx context.Context) DnsSplitNameserversMapOutput {
	return o
}

func (o DnsSplitNameserversMapOutput) MapIndex(k pulumi.StringInput) DnsSplitNameserversOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DnsSplitNameservers {
		return vs[0].(map[string]*DnsSplitNameservers)[vs[1].(string)]
	}).(DnsSplitNameserversOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DnsSplitNameserversInput)(nil)).Elem(), &DnsSplitNameservers{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsSplitNameserversArrayInput)(nil)).Elem(), DnsSplitNameserversArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsSplitNameserversMapInput)(nil)).Elem(), DnsSplitNameserversMap{})
	pulumi.RegisterOutputType(DnsSplitNameserversOutput{})
	pulumi.RegisterOutputType(DnsSplitNameserversArrayOutput{})
	pulumi.RegisterOutputType(DnsSplitNameserversMapOutput{})
}
