// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tailscale

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The tailnetKey resource allows you to create pre-authentication keys that can register new nodes without needing to sign in via a web browser. See https://tailscale.com/kb/1085/auth-keys for more information
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tailscale/sdk/go/tailscale"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := tailscale.NewTailnetKey(ctx, "sampleKey", &tailscale.TailnetKeyArgs{
// 			Ephemeral:     pulumi.Bool(false),
// 			Preauthorized: pulumi.Bool(true),
// 			Reusable:      pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type TailnetKey struct {
	pulumi.CustomResourceState

	// Indicates if the key is ephemeral.
	Ephemeral pulumi.BoolPtrOutput `pulumi:"ephemeral"`
	// The authentication key
	Key pulumi.StringOutput `pulumi:"key"`
	// Determines whether or not the machines authenticated by the key will be authorized for the tailnet by default.
	Preauthorized pulumi.BoolPtrOutput `pulumi:"preauthorized"`
	// Indicates if the key is reusable or single-use.
	Reusable pulumi.BoolPtrOutput `pulumi:"reusable"`
	// List of tags to apply to the machines authenticated by the key.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
}

// NewTailnetKey registers a new resource with the given unique name, arguments, and options.
func NewTailnetKey(ctx *pulumi.Context,
	name string, args *TailnetKeyArgs, opts ...pulumi.ResourceOption) (*TailnetKey, error) {
	if args == nil {
		args = &TailnetKeyArgs{}
	}

	var resource TailnetKey
	err := ctx.RegisterResource("tailscale:index/tailnetKey:TailnetKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTailnetKey gets an existing TailnetKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTailnetKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TailnetKeyState, opts ...pulumi.ResourceOption) (*TailnetKey, error) {
	var resource TailnetKey
	err := ctx.ReadResource("tailscale:index/tailnetKey:TailnetKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TailnetKey resources.
type tailnetKeyState struct {
	// Indicates if the key is ephemeral.
	Ephemeral *bool `pulumi:"ephemeral"`
	// The authentication key
	Key *string `pulumi:"key"`
	// Determines whether or not the machines authenticated by the key will be authorized for the tailnet by default.
	Preauthorized *bool `pulumi:"preauthorized"`
	// Indicates if the key is reusable or single-use.
	Reusable *bool `pulumi:"reusable"`
	// List of tags to apply to the machines authenticated by the key.
	Tags []string `pulumi:"tags"`
}

type TailnetKeyState struct {
	// Indicates if the key is ephemeral.
	Ephemeral pulumi.BoolPtrInput
	// The authentication key
	Key pulumi.StringPtrInput
	// Determines whether or not the machines authenticated by the key will be authorized for the tailnet by default.
	Preauthorized pulumi.BoolPtrInput
	// Indicates if the key is reusable or single-use.
	Reusable pulumi.BoolPtrInput
	// List of tags to apply to the machines authenticated by the key.
	Tags pulumi.StringArrayInput
}

func (TailnetKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*tailnetKeyState)(nil)).Elem()
}

type tailnetKeyArgs struct {
	// Indicates if the key is ephemeral.
	Ephemeral *bool `pulumi:"ephemeral"`
	// Determines whether or not the machines authenticated by the key will be authorized for the tailnet by default.
	Preauthorized *bool `pulumi:"preauthorized"`
	// Indicates if the key is reusable or single-use.
	Reusable *bool `pulumi:"reusable"`
	// List of tags to apply to the machines authenticated by the key.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a TailnetKey resource.
type TailnetKeyArgs struct {
	// Indicates if the key is ephemeral.
	Ephemeral pulumi.BoolPtrInput
	// Determines whether or not the machines authenticated by the key will be authorized for the tailnet by default.
	Preauthorized pulumi.BoolPtrInput
	// Indicates if the key is reusable or single-use.
	Reusable pulumi.BoolPtrInput
	// List of tags to apply to the machines authenticated by the key.
	Tags pulumi.StringArrayInput
}

func (TailnetKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tailnetKeyArgs)(nil)).Elem()
}

type TailnetKeyInput interface {
	pulumi.Input

	ToTailnetKeyOutput() TailnetKeyOutput
	ToTailnetKeyOutputWithContext(ctx context.Context) TailnetKeyOutput
}

func (*TailnetKey) ElementType() reflect.Type {
	return reflect.TypeOf((**TailnetKey)(nil)).Elem()
}

func (i *TailnetKey) ToTailnetKeyOutput() TailnetKeyOutput {
	return i.ToTailnetKeyOutputWithContext(context.Background())
}

func (i *TailnetKey) ToTailnetKeyOutputWithContext(ctx context.Context) TailnetKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TailnetKeyOutput)
}

// TailnetKeyArrayInput is an input type that accepts TailnetKeyArray and TailnetKeyArrayOutput values.
// You can construct a concrete instance of `TailnetKeyArrayInput` via:
//
//          TailnetKeyArray{ TailnetKeyArgs{...} }
type TailnetKeyArrayInput interface {
	pulumi.Input

	ToTailnetKeyArrayOutput() TailnetKeyArrayOutput
	ToTailnetKeyArrayOutputWithContext(context.Context) TailnetKeyArrayOutput
}

type TailnetKeyArray []TailnetKeyInput

func (TailnetKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TailnetKey)(nil)).Elem()
}

func (i TailnetKeyArray) ToTailnetKeyArrayOutput() TailnetKeyArrayOutput {
	return i.ToTailnetKeyArrayOutputWithContext(context.Background())
}

func (i TailnetKeyArray) ToTailnetKeyArrayOutputWithContext(ctx context.Context) TailnetKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TailnetKeyArrayOutput)
}

// TailnetKeyMapInput is an input type that accepts TailnetKeyMap and TailnetKeyMapOutput values.
// You can construct a concrete instance of `TailnetKeyMapInput` via:
//
//          TailnetKeyMap{ "key": TailnetKeyArgs{...} }
type TailnetKeyMapInput interface {
	pulumi.Input

	ToTailnetKeyMapOutput() TailnetKeyMapOutput
	ToTailnetKeyMapOutputWithContext(context.Context) TailnetKeyMapOutput
}

type TailnetKeyMap map[string]TailnetKeyInput

func (TailnetKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TailnetKey)(nil)).Elem()
}

func (i TailnetKeyMap) ToTailnetKeyMapOutput() TailnetKeyMapOutput {
	return i.ToTailnetKeyMapOutputWithContext(context.Background())
}

func (i TailnetKeyMap) ToTailnetKeyMapOutputWithContext(ctx context.Context) TailnetKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TailnetKeyMapOutput)
}

type TailnetKeyOutput struct{ *pulumi.OutputState }

func (TailnetKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TailnetKey)(nil)).Elem()
}

func (o TailnetKeyOutput) ToTailnetKeyOutput() TailnetKeyOutput {
	return o
}

func (o TailnetKeyOutput) ToTailnetKeyOutputWithContext(ctx context.Context) TailnetKeyOutput {
	return o
}

type TailnetKeyArrayOutput struct{ *pulumi.OutputState }

func (TailnetKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TailnetKey)(nil)).Elem()
}

func (o TailnetKeyArrayOutput) ToTailnetKeyArrayOutput() TailnetKeyArrayOutput {
	return o
}

func (o TailnetKeyArrayOutput) ToTailnetKeyArrayOutputWithContext(ctx context.Context) TailnetKeyArrayOutput {
	return o
}

func (o TailnetKeyArrayOutput) Index(i pulumi.IntInput) TailnetKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TailnetKey {
		return vs[0].([]*TailnetKey)[vs[1].(int)]
	}).(TailnetKeyOutput)
}

type TailnetKeyMapOutput struct{ *pulumi.OutputState }

func (TailnetKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TailnetKey)(nil)).Elem()
}

func (o TailnetKeyMapOutput) ToTailnetKeyMapOutput() TailnetKeyMapOutput {
	return o
}

func (o TailnetKeyMapOutput) ToTailnetKeyMapOutputWithContext(ctx context.Context) TailnetKeyMapOutput {
	return o
}

func (o TailnetKeyMapOutput) MapIndex(k pulumi.StringInput) TailnetKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TailnetKey {
		return vs[0].(map[string]*TailnetKey)[vs[1].(string)]
	}).(TailnetKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TailnetKeyInput)(nil)).Elem(), &TailnetKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*TailnetKeyArrayInput)(nil)).Elem(), TailnetKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TailnetKeyMapInput)(nil)).Elem(), TailnetKeyMap{})
	pulumi.RegisterOutputType(TailnetKeyOutput{})
	pulumi.RegisterOutputType(TailnetKeyArrayOutput{})
	pulumi.RegisterOutputType(TailnetKeyMapOutput{})
}
