// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: protos/linkshorter.proto

package linkshorter_v1

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

// ShortenerClient is the client API for Shortener service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShortenerClient interface {
	Post(ctx context.Context, in *OriginalLink, opts ...grpc.CallOption) (*ShortLink, error)
	Get(ctx context.Context, in *ShortLink, opts ...grpc.CallOption) (*OriginalLink, error)
}

type shortenerClient struct {
	cc grpc.ClientConnInterface
}

func NewShortenerClient(cc grpc.ClientConnInterface) ShortenerClient {
	return &shortenerClient{cc}
}

func (c *shortenerClient) Post(ctx context.Context, in *OriginalLink, opts ...grpc.CallOption) (*ShortLink, error) {
	out := new(ShortLink)
	err := c.cc.Invoke(ctx, "/linkShorter.Shortener/Post", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenerClient) Get(ctx context.Context, in *ShortLink, opts ...grpc.CallOption) (*OriginalLink, error) {
	out := new(OriginalLink)
	err := c.cc.Invoke(ctx, "/linkShorter.Shortener/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShortenerServer is the server API for Shortener service.
// All implementations must embed UnimplementedShortenerServer
// for forward compatibility
type ShortenerServer interface {
	Post(context.Context, *OriginalLink) (*ShortLink, error)
	Get(context.Context, *ShortLink) (*OriginalLink, error)
	mustEmbedUnimplementedShortenerServer()
}

// UnimplementedShortenerServer must be embedded to have forward compatible implementations.
type UnimplementedShortenerServer struct {
}

func (UnimplementedShortenerServer) Post(context.Context, *OriginalLink) (*ShortLink, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
}
func (UnimplementedShortenerServer) Get(context.Context, *ShortLink) (*OriginalLink, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedShortenerServer) mustEmbedUnimplementedShortenerServer() {}

// UnsafeShortenerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShortenerServer will
// result in compilation errors.
type UnsafeShortenerServer interface {
	mustEmbedUnimplementedShortenerServer()
}

func RegisterShortenerServer(s grpc.ServiceRegistrar, srv ShortenerServer) {
	s.RegisterService(&Shortener_ServiceDesc, srv)
}

func _Shortener_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OriginalLink)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkShorter.Shortener/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).Post(ctx, req.(*OriginalLink))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shortener_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShortLink)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkShorter.Shortener/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).Get(ctx, req.(*ShortLink))
	}
	return interceptor(ctx, in, info, handler)
}

// Shortener_ServiceDesc is the grpc.ServiceDesc for Shortener service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Shortener_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "linkShorter.Shortener",
	HandlerType: (*ShortenerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Post",
			Handler:    _Shortener_Post_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Shortener_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/linkshorter.proto",
}
