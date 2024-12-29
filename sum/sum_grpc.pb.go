// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: sum.proto

package sum

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
	Sum_AddNumbers_FullMethodName = "/sum.Sum/AddNumbers"
)

// SumClient is the client API for Sum service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SumClient interface {
	AddNumbers(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type sumClient struct {
	cc grpc.ClientConnInterface
}

func NewSumClient(cc grpc.ClientConnInterface) SumClient {
	return &sumClient{cc}
}

func (c *sumClient) AddNumbers(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, Sum_AddNumbers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SumServer is the server API for Sum service.
// All implementations must embed UnimplementedSumServer
// for forward compatibility.
type SumServer interface {
	AddNumbers(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedSumServer()
}

// UnimplementedSumServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSumServer struct{}

func (UnimplementedSumServer) AddNumbers(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNumbers not implemented")
}
func (UnimplementedSumServer) mustEmbedUnimplementedSumServer() {}
func (UnimplementedSumServer) testEmbeddedByValue()             {}

// UnsafeSumServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SumServer will
// result in compilation errors.
type UnsafeSumServer interface {
	mustEmbedUnimplementedSumServer()
}

func RegisterSumServer(s grpc.ServiceRegistrar, srv SumServer) {
	// If the following call pancis, it indicates UnimplementedSumServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Sum_ServiceDesc, srv)
}

func _Sum_AddNumbers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServer).AddNumbers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sum_AddNumbers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServer).AddNumbers(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Sum_ServiceDesc is the grpc.ServiceDesc for Sum service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sum_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sum.Sum",
	HandlerType: (*SumServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddNumbers",
			Handler:    _Sum_AddNumbers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sum.proto",
}