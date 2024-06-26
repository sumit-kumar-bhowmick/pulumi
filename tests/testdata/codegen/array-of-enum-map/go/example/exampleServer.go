// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"array-of-enum-map/example/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ExampleServer struct {
	pulumi.CustomResourceState

	MapArrayEnum AnnotationStoreSchemaValueTypeMapArrayOutput `pulumi:"mapArrayEnum"`
}

// NewExampleServer registers a new resource with the given unique name, arguments, and options.
func NewExampleServer(ctx *pulumi.Context,
	name string, args *ExampleServerArgs, opts ...pulumi.ResourceOption) (*ExampleServer, error) {
	if args == nil {
		args = &ExampleServerArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ExampleServer
	err := ctx.RegisterResource("example:index:ExampleServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExampleServer gets an existing ExampleServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExampleServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExampleServerState, opts ...pulumi.ResourceOption) (*ExampleServer, error) {
	var resource ExampleServer
	err := ctx.ReadResource("example:index:ExampleServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExampleServer resources.
type exampleServerState struct {
}

type ExampleServerState struct {
}

func (ExampleServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*exampleServerState)(nil)).Elem()
}

type exampleServerArgs struct {
	MapArrayEnum []map[string]AnnotationStoreSchemaValueType `pulumi:"mapArrayEnum"`
}

// The set of arguments for constructing a ExampleServer resource.
type ExampleServerArgs struct {
	MapArrayEnum AnnotationStoreSchemaValueTypeMapArrayInput
}

func (ExampleServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*exampleServerArgs)(nil)).Elem()
}

type ExampleServerInput interface {
	pulumi.Input

	ToExampleServerOutput() ExampleServerOutput
	ToExampleServerOutputWithContext(ctx context.Context) ExampleServerOutput
}

func (*ExampleServer) ElementType() reflect.Type {
	return reflect.TypeOf((**ExampleServer)(nil)).Elem()
}

func (i *ExampleServer) ToExampleServerOutput() ExampleServerOutput {
	return i.ToExampleServerOutputWithContext(context.Background())
}

func (i *ExampleServer) ToExampleServerOutputWithContext(ctx context.Context) ExampleServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExampleServerOutput)
}

type ExampleServerOutput struct{ *pulumi.OutputState }

func (ExampleServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExampleServer)(nil)).Elem()
}

func (o ExampleServerOutput) ToExampleServerOutput() ExampleServerOutput {
	return o
}

func (o ExampleServerOutput) ToExampleServerOutputWithContext(ctx context.Context) ExampleServerOutput {
	return o
}

func (o ExampleServerOutput) MapArrayEnum() AnnotationStoreSchemaValueTypeMapArrayOutput {
	return o.ApplyT(func(v *ExampleServer) AnnotationStoreSchemaValueTypeMapArrayOutput { return v.MapArrayEnum }).(AnnotationStoreSchemaValueTypeMapArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExampleServerInput)(nil)).Elem(), &ExampleServer{})
	pulumi.RegisterOutputType(ExampleServerOutput{})
}
