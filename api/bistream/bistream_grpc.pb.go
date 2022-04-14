// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: bistream/bistream.proto

package bistream

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DoubleStreamClient is the client API for DoubleStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DoubleStreamClient interface {
	DoubleStream(ctx context.Context, opts ...grpc.CallOption) (DoubleStream_DoubleStreamClient, error)
}

type doubleStreamClient struct {
	cc grpc.ClientConnInterface
}

func NewDoubleStreamClient(cc grpc.ClientConnInterface) DoubleStreamClient {
	return &doubleStreamClient{cc}
}

func (c *doubleStreamClient) DoubleStream(ctx context.Context, opts ...grpc.CallOption) (DoubleStream_DoubleStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &DoubleStream_ServiceDesc.Streams[0], "/bistream.DoubleStream/DoubleStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &doubleStreamDoubleStreamClient{stream}
	return x, nil
}

type DoubleStream_DoubleStreamClient interface {
	Send(*DoubleStreamRequest) error
	Recv() (*DoubleStreamResponse, error)
	grpc.ClientStream
}

type doubleStreamDoubleStreamClient struct {
	grpc.ClientStream
}

func (x *doubleStreamDoubleStreamClient) Send(m *DoubleStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *doubleStreamDoubleStreamClient) Recv() (*DoubleStreamResponse, error) {
	m := new(DoubleStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DoubleStreamServer is the server API for DoubleStream service.
// All implementations must embed UnimplementedDoubleStreamServer
// for forward compatibility
type DoubleStreamServer interface {
	DoubleStream(DoubleStream_DoubleStreamServer) error
	mustEmbedUnimplementedDoubleStreamServer()
}

// UnimplementedDoubleStreamServer must be embedded to have forward compatible implementations.
type UnimplementedDoubleStreamServer struct {
}

func (UnimplementedDoubleStreamServer) DoubleStream(DoubleStream_DoubleStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method DoubleStream not implemented")
}
func (UnimplementedDoubleStreamServer) mustEmbedUnimplementedDoubleStreamServer() {}

// UnsafeDoubleStreamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DoubleStreamServer will
// result in compilation errors.
type UnsafeDoubleStreamServer interface {
	mustEmbedUnimplementedDoubleStreamServer()
}

func RegisterDoubleStreamServer(s grpc.ServiceRegistrar, srv DoubleStreamServer) {
	s.RegisterService(&DoubleStream_ServiceDesc, srv)
}

func _DoubleStream_DoubleStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DoubleStreamServer).DoubleStream(&doubleStreamDoubleStreamServer{stream})
}

type DoubleStream_DoubleStreamServer interface {
	Send(*DoubleStreamResponse) error
	Recv() (*DoubleStreamRequest, error)
	grpc.ServerStream
}

type doubleStreamDoubleStreamServer struct {
	grpc.ServerStream
}

func (x *doubleStreamDoubleStreamServer) Send(m *DoubleStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *doubleStreamDoubleStreamServer) Recv() (*DoubleStreamRequest, error) {
	m := new(DoubleStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DoubleStream_ServiceDesc is the grpc.ServiceDesc for DoubleStream service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DoubleStream_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bistream.DoubleStream",
	HandlerType: (*DoubleStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DoubleStream",
			Handler:       _DoubleStream_DoubleStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "bistream/bistream.proto",
}
