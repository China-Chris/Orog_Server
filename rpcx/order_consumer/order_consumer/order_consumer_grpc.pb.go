// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: order_consumer.proto

package order_consumer

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	OrderConsumer_Ping_FullMethodName = "/order_consumer.OrderConsumer/Ping"
)

// OrderConsumerClient is the client API for OrderConsumer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderConsumerClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type orderConsumerClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderConsumerClient(cc grpc.ClientConnInterface) OrderConsumerClient {
	return &orderConsumerClient{cc}
}

func (c *orderConsumerClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, OrderConsumer_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderConsumerServer is the server API for OrderConsumer service.
// All implementations must embed UnimplementedOrderConsumerServer
// for forward compatibility.
type OrderConsumerServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedOrderConsumerServer()
}

// UnimplementedOrderConsumerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOrderConsumerServer struct{}

func (UnimplementedOrderConsumerServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedOrderConsumerServer) mustEmbedUnimplementedOrderConsumerServer() {}
func (UnimplementedOrderConsumerServer) testEmbeddedByValue()                       {}

// UnsafeOrderConsumerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderConsumerServer will
// result in compilation errors.
type UnsafeOrderConsumerServer interface {
	mustEmbedUnimplementedOrderConsumerServer()
}

func RegisterOrderConsumerServer(s grpc.ServiceRegistrar, srv OrderConsumerServer) {
	// If the following call pancis, it indicates UnimplementedOrderConsumerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OrderConsumer_ServiceDesc, srv)
}

func _OrderConsumer_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderConsumerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderConsumer_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderConsumerServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderConsumer_ServiceDesc is the grpc.ServiceDesc for OrderConsumer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderConsumer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order_consumer.OrderConsumer",
	HandlerType: (*OrderConsumerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _OrderConsumer_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order_consumer.proto",
}
