// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package tailscale

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The deviceTags resource is used to apply tags to a device within a Tailnet. For more information on ACL tags, see
// the [ACL tags documentation](https://tailscale.com/kb/1068/acl-tags/) for more details.
type DeviceTags struct {
	pulumi.CustomResourceState

	// The device to apply tags to.
	DeviceId pulumi.StringOutput `pulumi:"deviceId"`
	// The tags to apply to the device.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
}

// NewDeviceTags registers a new resource with the given unique name, arguments, and options.
func NewDeviceTags(ctx *pulumi.Context,
	name string, args *DeviceTagsArgs, opts ...pulumi.ResourceOption) (*DeviceTags, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceId == nil {
		return nil, errors.New("invalid value for required argument 'DeviceId'")
	}
	if args.Tags == nil {
		return nil, errors.New("invalid value for required argument 'Tags'")
	}
	var resource DeviceTags
	err := ctx.RegisterResource("tailscale:index/deviceTags:DeviceTags", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeviceTags gets an existing DeviceTags resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeviceTags(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeviceTagsState, opts ...pulumi.ResourceOption) (*DeviceTags, error) {
	var resource DeviceTags
	err := ctx.ReadResource("tailscale:index/deviceTags:DeviceTags", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeviceTags resources.
type deviceTagsState struct {
	// The device to apply tags to.
	DeviceId *string `pulumi:"deviceId"`
	// The tags to apply to the device.
	Tags []string `pulumi:"tags"`
}

type DeviceTagsState struct {
	// The device to apply tags to.
	DeviceId pulumi.StringPtrInput
	// The tags to apply to the device.
	Tags pulumi.StringArrayInput
}

func (DeviceTagsState) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceTagsState)(nil)).Elem()
}

type deviceTagsArgs struct {
	// The device to apply tags to.
	DeviceId string `pulumi:"deviceId"`
	// The tags to apply to the device.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a DeviceTags resource.
type DeviceTagsArgs struct {
	// The device to apply tags to.
	DeviceId pulumi.StringInput
	// The tags to apply to the device.
	Tags pulumi.StringArrayInput
}

func (DeviceTagsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceTagsArgs)(nil)).Elem()
}

type DeviceTagsInput interface {
	pulumi.Input

	ToDeviceTagsOutput() DeviceTagsOutput
	ToDeviceTagsOutputWithContext(ctx context.Context) DeviceTagsOutput
}

func (*DeviceTags) ElementType() reflect.Type {
	return reflect.TypeOf((**DeviceTags)(nil)).Elem()
}

func (i *DeviceTags) ToDeviceTagsOutput() DeviceTagsOutput {
	return i.ToDeviceTagsOutputWithContext(context.Background())
}

func (i *DeviceTags) ToDeviceTagsOutputWithContext(ctx context.Context) DeviceTagsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceTagsOutput)
}

// DeviceTagsArrayInput is an input type that accepts DeviceTagsArray and DeviceTagsArrayOutput values.
// You can construct a concrete instance of `DeviceTagsArrayInput` via:
//
//          DeviceTagsArray{ DeviceTagsArgs{...} }
type DeviceTagsArrayInput interface {
	pulumi.Input

	ToDeviceTagsArrayOutput() DeviceTagsArrayOutput
	ToDeviceTagsArrayOutputWithContext(context.Context) DeviceTagsArrayOutput
}

type DeviceTagsArray []DeviceTagsInput

func (DeviceTagsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DeviceTags)(nil)).Elem()
}

func (i DeviceTagsArray) ToDeviceTagsArrayOutput() DeviceTagsArrayOutput {
	return i.ToDeviceTagsArrayOutputWithContext(context.Background())
}

func (i DeviceTagsArray) ToDeviceTagsArrayOutputWithContext(ctx context.Context) DeviceTagsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceTagsArrayOutput)
}

// DeviceTagsMapInput is an input type that accepts DeviceTagsMap and DeviceTagsMapOutput values.
// You can construct a concrete instance of `DeviceTagsMapInput` via:
//
//          DeviceTagsMap{ "key": DeviceTagsArgs{...} }
type DeviceTagsMapInput interface {
	pulumi.Input

	ToDeviceTagsMapOutput() DeviceTagsMapOutput
	ToDeviceTagsMapOutputWithContext(context.Context) DeviceTagsMapOutput
}

type DeviceTagsMap map[string]DeviceTagsInput

func (DeviceTagsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DeviceTags)(nil)).Elem()
}

func (i DeviceTagsMap) ToDeviceTagsMapOutput() DeviceTagsMapOutput {
	return i.ToDeviceTagsMapOutputWithContext(context.Background())
}

func (i DeviceTagsMap) ToDeviceTagsMapOutputWithContext(ctx context.Context) DeviceTagsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceTagsMapOutput)
}

type DeviceTagsOutput struct{ *pulumi.OutputState }

func (DeviceTagsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeviceTags)(nil)).Elem()
}

func (o DeviceTagsOutput) ToDeviceTagsOutput() DeviceTagsOutput {
	return o
}

func (o DeviceTagsOutput) ToDeviceTagsOutputWithContext(ctx context.Context) DeviceTagsOutput {
	return o
}

type DeviceTagsArrayOutput struct{ *pulumi.OutputState }

func (DeviceTagsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DeviceTags)(nil)).Elem()
}

func (o DeviceTagsArrayOutput) ToDeviceTagsArrayOutput() DeviceTagsArrayOutput {
	return o
}

func (o DeviceTagsArrayOutput) ToDeviceTagsArrayOutputWithContext(ctx context.Context) DeviceTagsArrayOutput {
	return o
}

func (o DeviceTagsArrayOutput) Index(i pulumi.IntInput) DeviceTagsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DeviceTags {
		return vs[0].([]*DeviceTags)[vs[1].(int)]
	}).(DeviceTagsOutput)
}

type DeviceTagsMapOutput struct{ *pulumi.OutputState }

func (DeviceTagsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DeviceTags)(nil)).Elem()
}

func (o DeviceTagsMapOutput) ToDeviceTagsMapOutput() DeviceTagsMapOutput {
	return o
}

func (o DeviceTagsMapOutput) ToDeviceTagsMapOutputWithContext(ctx context.Context) DeviceTagsMapOutput {
	return o
}

func (o DeviceTagsMapOutput) MapIndex(k pulumi.StringInput) DeviceTagsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DeviceTags {
		return vs[0].(map[string]*DeviceTags)[vs[1].(string)]
	}).(DeviceTagsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceTagsInput)(nil)).Elem(), &DeviceTags{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceTagsArrayInput)(nil)).Elem(), DeviceTagsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceTagsMapInput)(nil)).Elem(), DeviceTagsMap{})
	pulumi.RegisterOutputType(DeviceTagsOutput{})
	pulumi.RegisterOutputType(DeviceTagsArrayOutput{})
	pulumi.RegisterOutputType(DeviceTagsMapOutput{})
}