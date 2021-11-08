// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package informertelegrambot

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

// InformerBotMessagingClient is the client API for InformerBotMessaging service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InformerBotMessagingClient interface {
	NewMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*MessageResponse, error)
}

type informerBotMessagingClient struct {
	cc grpc.ClientConnInterface
}

func NewInformerBotMessagingClient(cc grpc.ClientConnInterface) InformerBotMessagingClient {
	return &informerBotMessagingClient{cc}
}

func (c *informerBotMessagingClient) NewMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*MessageResponse, error) {
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, "/com.informertelegrambot.messaging.InformerBotMessaging/NewMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InformerBotMessagingServer is the server API for InformerBotMessaging service.
// All implementations must embed UnimplementedInformerBotMessagingServer
// for forward compatibility
type InformerBotMessagingServer interface {
	NewMessage(context.Context, *Message) (*MessageResponse, error)
	mustEmbedUnimplementedInformerBotMessagingServer()
}

// UnimplementedInformerBotMessagingServer must be embedded to have forward compatible implementations.
type UnimplementedInformerBotMessagingServer struct {
}

func (UnimplementedInformerBotMessagingServer) NewMessage(context.Context, *Message) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewMessage not implemented")
}
func (UnimplementedInformerBotMessagingServer) mustEmbedUnimplementedInformerBotMessagingServer() {}

// UnsafeInformerBotMessagingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InformerBotMessagingServer will
// result in compilation errors.
type UnsafeInformerBotMessagingServer interface {
	mustEmbedUnimplementedInformerBotMessagingServer()
}

func RegisterInformerBotMessagingServer(s grpc.ServiceRegistrar, srv InformerBotMessagingServer) {
	s.RegisterService(&InformerBotMessaging_ServiceDesc, srv)
}

func _InformerBotMessaging_NewMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformerBotMessagingServer).NewMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.informertelegrambot.messaging.InformerBotMessaging/NewMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformerBotMessagingServer).NewMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

// InformerBotMessaging_ServiceDesc is the grpc.ServiceDesc for InformerBotMessaging service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InformerBotMessaging_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.informertelegrambot.messaging.InformerBotMessaging",
	HandlerType: (*InformerBotMessagingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewMessage",
			Handler:    _InformerBotMessaging_NewMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/informertelegrambot/messaging.proto",
}