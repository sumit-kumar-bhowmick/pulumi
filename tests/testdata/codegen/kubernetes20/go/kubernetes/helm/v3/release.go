// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A non-overlay, non-component, non-Kubernetes resource.
type Release struct {
	pulumi.CustomResourceState

	// Chart name to be installed. A path may be used.
	Chart pulumi.StringOutput `pulumi:"chart"`
	// List of assets (raw yaml files). Content is read and merged with values (with values taking precedence).
	ValueYamlFiles pulumi.AssetOrArchiveArrayOutput `pulumi:"valueYamlFiles"`
	// Custom values set for the release.
	Values pulumi.MapOutput `pulumi:"values"`
}

// NewRelease registers a new resource with the given unique name, arguments, and options.
func NewRelease(ctx *pulumi.Context,
	name string, args *ReleaseArgs, opts ...pulumi.ResourceOption) (*Release, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Chart == nil {
		return nil, errors.New("invalid value for required argument 'Chart'")
	}
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Release
	err := ctx.RegisterResource("kubernetes:helm.sh/v3:Release", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRelease gets an existing Release resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRelease(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReleaseState, opts ...pulumi.ResourceOption) (*Release, error) {
	var resource Release
	err := ctx.ReadResource("kubernetes:helm.sh/v3:Release", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Release resources.
type releaseState struct {
}

type ReleaseState struct {
}

func (ReleaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*releaseState)(nil)).Elem()
}

type releaseArgs struct {
	// Chart name to be installed. A path may be used.
	Chart string `pulumi:"chart"`
	// List of assets (raw yaml files). Content is read and merged with values.
	ValueYamlFiles []pulumi.AssetOrArchive `pulumi:"valueYamlFiles"`
	// Custom values set for the release.
	Values map[string]interface{} `pulumi:"values"`
}

// The set of arguments for constructing a Release resource.
type ReleaseArgs struct {
	// Chart name to be installed. A path may be used.
	Chart pulumi.StringInput
	// List of assets (raw yaml files). Content is read and merged with values.
	ValueYamlFiles pulumi.AssetOrArchiveArrayInput
	// Custom values set for the release.
	Values pulumi.MapInput
}

func (ReleaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*releaseArgs)(nil)).Elem()
}

type ReleaseInput interface {
	pulumi.Input

	ToReleaseOutput() ReleaseOutput
	ToReleaseOutputWithContext(ctx context.Context) ReleaseOutput
}

func (*Release) ElementType() reflect.Type {
	return reflect.TypeOf((**Release)(nil)).Elem()
}

func (i *Release) ToReleaseOutput() ReleaseOutput {
	return i.ToReleaseOutputWithContext(context.Background())
}

func (i *Release) ToReleaseOutputWithContext(ctx context.Context) ReleaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseOutput)
}

// ReleaseArrayInput is an input type that accepts ReleaseArray and ReleaseArrayOutput values.
// You can construct a concrete instance of `ReleaseArrayInput` via:
//
//	ReleaseArray{ ReleaseArgs{...} }
type ReleaseArrayInput interface {
	pulumi.Input

	ToReleaseArrayOutput() ReleaseArrayOutput
	ToReleaseArrayOutputWithContext(context.Context) ReleaseArrayOutput
}

type ReleaseArray []ReleaseInput

func (ReleaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Release)(nil)).Elem()
}

func (i ReleaseArray) ToReleaseArrayOutput() ReleaseArrayOutput {
	return i.ToReleaseArrayOutputWithContext(context.Background())
}

func (i ReleaseArray) ToReleaseArrayOutputWithContext(ctx context.Context) ReleaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseArrayOutput)
}

// ReleaseMapInput is an input type that accepts ReleaseMap and ReleaseMapOutput values.
// You can construct a concrete instance of `ReleaseMapInput` via:
//
//	ReleaseMap{ "key": ReleaseArgs{...} }
type ReleaseMapInput interface {
	pulumi.Input

	ToReleaseMapOutput() ReleaseMapOutput
	ToReleaseMapOutputWithContext(context.Context) ReleaseMapOutput
}

type ReleaseMap map[string]ReleaseInput

func (ReleaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Release)(nil)).Elem()
}

func (i ReleaseMap) ToReleaseMapOutput() ReleaseMapOutput {
	return i.ToReleaseMapOutputWithContext(context.Background())
}

func (i ReleaseMap) ToReleaseMapOutputWithContext(ctx context.Context) ReleaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseMapOutput)
}

type ReleaseOutput struct{ *pulumi.OutputState }

func (ReleaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Release)(nil)).Elem()
}

func (o ReleaseOutput) ToReleaseOutput() ReleaseOutput {
	return o
}

func (o ReleaseOutput) ToReleaseOutputWithContext(ctx context.Context) ReleaseOutput {
	return o
}

// Chart name to be installed. A path may be used.
func (o ReleaseOutput) Chart() pulumi.StringOutput {
	return o.ApplyT(func(v *Release) pulumi.StringOutput { return v.Chart }).(pulumi.StringOutput)
}

// List of assets (raw yaml files). Content is read and merged with values (with values taking precedence).
func (o ReleaseOutput) ValueYamlFiles() pulumi.AssetOrArchiveArrayOutput {
	return o.ApplyT(func(v *Release) pulumi.AssetOrArchiveArrayOutput { return v.ValueYamlFiles }).(pulumi.AssetOrArchiveArrayOutput)
}

// Custom values set for the release.
func (o ReleaseOutput) Values() pulumi.MapOutput {
	return o.ApplyT(func(v *Release) pulumi.MapOutput { return v.Values }).(pulumi.MapOutput)
}

type ReleaseArrayOutput struct{ *pulumi.OutputState }

func (ReleaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Release)(nil)).Elem()
}

func (o ReleaseArrayOutput) ToReleaseArrayOutput() ReleaseArrayOutput {
	return o
}

func (o ReleaseArrayOutput) ToReleaseArrayOutputWithContext(ctx context.Context) ReleaseArrayOutput {
	return o
}

func (o ReleaseArrayOutput) Index(i pulumi.IntInput) ReleaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Release {
		return vs[0].([]*Release)[vs[1].(int)]
	}).(ReleaseOutput)
}

type ReleaseMapOutput struct{ *pulumi.OutputState }

func (ReleaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Release)(nil)).Elem()
}

func (o ReleaseMapOutput) ToReleaseMapOutput() ReleaseMapOutput {
	return o
}

func (o ReleaseMapOutput) ToReleaseMapOutputWithContext(ctx context.Context) ReleaseMapOutput {
	return o
}

func (o ReleaseMapOutput) MapIndex(k pulumi.StringInput) ReleaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Release {
		return vs[0].(map[string]*Release)[vs[1].(string)]
	}).(ReleaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseInput)(nil)).Elem(), &Release{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseArrayInput)(nil)).Elem(), ReleaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseMapInput)(nil)).Elem(), ReleaseMap{})
	pulumi.RegisterOutputType(ReleaseOutput{})
	pulumi.RegisterOutputType(ReleaseArrayOutput{})
	pulumi.RegisterOutputType(ReleaseMapOutput{})
}
