// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// VimUIClient is the client API for VimUI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VimUIClient interface {
	// ListNodeSpecs returns a list of Node(s) that can be started through the UI.
	ListNodeSpecs(ctx context.Context, in *ListNodeSpecsRequest, opts ...grpc.CallOption) (VimUI_ListNodeSpecsClient, error)
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error)
}

type vimUIClient struct {
	cc grpc.ClientConnInterface
}

func NewVimUIClient(cc grpc.ClientConnInterface) VimUIClient {
	return &vimUIClient{cc}
}

func (c *vimUIClient) ListNodeSpecs(ctx context.Context, in *ListNodeSpecsRequest, opts ...grpc.CallOption) (VimUI_ListNodeSpecsClient, error) {
	stream, err := c.cc.NewStream(ctx, &VimUI_ServiceDesc.Streams[0], "/v1.VimUI/ListNodeSpecs", opts...)
	if err != nil {
		return nil, err
	}
	x := &vimUIListNodeSpecsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type VimUI_ListNodeSpecsClient interface {
	Recv() (*ListNodeSpecsResponse, error)
	grpc.ClientStream
}

type vimUIListNodeSpecsClient struct {
	grpc.ClientStream
}

func (x *vimUIListNodeSpecsClient) Recv() (*ListNodeSpecsResponse, error) {
	m := new(ListNodeSpecsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *vimUIClient) IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error) {
	out := new(IsReadOnlyResponse)
	err := c.cc.Invoke(ctx, "/v1.VimUI/IsReadOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VimUIServer is the server API for VimUI service.
// All implementations must embed UnimplementedVimUIServer
// for forward compatibility
type VimUIServer interface {
	// ListNodeSpecs returns a list of Node(s) that can be started through the UI.
	ListNodeSpecs(*ListNodeSpecsRequest, VimUI_ListNodeSpecsServer) error
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error)
	mustEmbedUnimplementedVimUIServer()
}

// UnimplementedVimUIServer must be embedded to have forward compatible implementations.
type UnimplementedVimUIServer struct {
}

func (UnimplementedVimUIServer) ListNodeSpecs(*ListNodeSpecsRequest, VimUI_ListNodeSpecsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListNodeSpecs not implemented")
}
func (UnimplementedVimUIServer) IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsReadOnly not implemented")
}
func (UnimplementedVimUIServer) mustEmbedUnimplementedVimUIServer() {}

// UnsafeVimUIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VimUIServer will
// result in compilation errors.
type UnsafeVimUIServer interface {
	mustEmbedUnimplementedVimUIServer()
}

func RegisterVimUIServer(s grpc.ServiceRegistrar, srv VimUIServer) {
	s.RegisterService(&VimUI_ServiceDesc, srv)
}

func _VimUI_ListNodeSpecs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListNodeSpecsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VimUIServer).ListNodeSpecs(m, &vimUIListNodeSpecsServer{stream})
}

type VimUI_ListNodeSpecsServer interface {
	Send(*ListNodeSpecsResponse) error
	grpc.ServerStream
}

type vimUIListNodeSpecsServer struct {
	grpc.ServerStream
}

func (x *vimUIListNodeSpecsServer) Send(m *ListNodeSpecsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _VimUI_IsReadOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsReadOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VimUIServer).IsReadOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.VimUI/IsReadOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VimUIServer).IsReadOnly(ctx, req.(*IsReadOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VimUI_ServiceDesc is the grpc.ServiceDesc for VimUI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VimUI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.VimUI",
	HandlerType: (*VimUIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsReadOnly",
			Handler:    _VimUI_IsReadOnly_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListNodeSpecs",
			Handler:       _VimUI_ListNodeSpecs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "vim-ui.proto",
}
