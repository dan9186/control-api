// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package modulepb

import (
	v1 "api/core/v1"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ModuleServiceClient is the client API for ModuleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModuleServiceClient interface {
	EventTopic(ctx context.Context, in *v1.EventMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownResponse, error)
}

type moduleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewModuleServiceClient(cc grpc.ClientConnInterface) ModuleServiceClient {
	return &moduleServiceClient{cc}
}

func (c *moduleServiceClient) EventTopic(ctx context.Context, in *v1.EventMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/control.module.v1.ModuleService/EventTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moduleServiceClient) Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownResponse, error) {
	out := new(ShutdownResponse)
	err := c.cc.Invoke(ctx, "/control.module.v1.ModuleService/Shutdown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModuleServiceServer is the server API for ModuleService service.
// All implementations must embed UnimplementedModuleServiceServer
// for forward compatibility
type ModuleServiceServer interface {
	EventTopic(context.Context, *v1.EventMessage) (*emptypb.Empty, error)
	Shutdown(context.Context, *ShutdownRequest) (*ShutdownResponse, error)
	mustEmbedUnimplementedModuleServiceServer()
}

// UnimplementedModuleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedModuleServiceServer struct {
}

func (*UnimplementedModuleServiceServer) EventTopic(context.Context, *v1.EventMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EventTopic not implemented")
}
func (*UnimplementedModuleServiceServer) Shutdown(context.Context, *ShutdownRequest) (*ShutdownResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shutdown not implemented")
}
func (*UnimplementedModuleServiceServer) mustEmbedUnimplementedModuleServiceServer() {}

func RegisterModuleServiceServer(s *grpc.Server, srv ModuleServiceServer) {
	s.RegisterService(&_ModuleService_serviceDesc, srv)
}

func _ModuleService_EventTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.EventMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModuleServiceServer).EventTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control.module.v1.ModuleService/EventTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModuleServiceServer).EventTopic(ctx, req.(*v1.EventMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModuleService_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShutdownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModuleServiceServer).Shutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control.module.v1.ModuleService/Shutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModuleServiceServer).Shutdown(ctx, req.(*ShutdownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ModuleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "control.module.v1.ModuleService",
	HandlerType: (*ModuleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EventTopic",
			Handler:    _ModuleService_EventTopic_Handler,
		},
		{
			MethodName: "Shutdown",
			Handler:    _ModuleService_Shutdown_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "module/v1/module.proto",
}
