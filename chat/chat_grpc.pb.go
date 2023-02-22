// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: chat/chat.proto

package chat

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

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatClient interface {
	Messaging(ctx context.Context, opts ...grpc.CallOption) (Chat_MessagingClient, error)
	CheckUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error)
	GetUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Chat_GetUsersClient, error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) Messaging(ctx context.Context, opts ...grpc.CallOption) (Chat_MessagingClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chat_ServiceDesc.Streams[0], "/chat.Chat/Messaging", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatMessagingClient{stream}
	return x, nil
}

type Chat_MessagingClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type chatMessagingClient struct {
	grpc.ClientStream
}

func (x *chatMessagingClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatMessagingClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatClient) CheckUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/chat.Chat/CheckUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) GetUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Chat_GetUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chat_ServiceDesc.Streams[1], "/chat.Chat/GetUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatGetUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chat_GetUsersClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type chatGetUsersClient struct {
	grpc.ClientStream
}

func (x *chatGetUsersClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServer is the server API for Chat service.
// All implementations must embed UnimplementedChatServer
// for forward compatibility
type ChatServer interface {
	Messaging(Chat_MessagingServer) error
	CheckUser(context.Context, *User) (*Response, error)
	GetUsers(*Empty, Chat_GetUsersServer) error
	mustEmbedUnimplementedChatServer()
}

// UnimplementedChatServer must be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (UnimplementedChatServer) Messaging(Chat_MessagingServer) error {
	return status.Errorf(codes.Unimplemented, "method Messaging not implemented")
}
func (UnimplementedChatServer) CheckUser(context.Context, *User) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUser not implemented")
}
func (UnimplementedChatServer) GetUsers(*Empty, Chat_GetUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedChatServer) mustEmbedUnimplementedChatServer() {}

// UnsafeChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServer will
// result in compilation errors.
type UnsafeChatServer interface {
	mustEmbedUnimplementedChatServer()
}

func RegisterChatServer(s grpc.ServiceRegistrar, srv ChatServer) {
	s.RegisterService(&Chat_ServiceDesc, srv)
}

func _Chat_Messaging_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServer).Messaging(&chatMessagingServer{stream})
}

type Chat_MessagingServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type chatMessagingServer struct {
	grpc.ServerStream
}

func (x *chatMessagingServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatMessagingServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Chat_CheckUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).CheckUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.Chat/CheckUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).CheckUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_GetUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServer).GetUsers(m, &chatGetUsersServer{stream})
}

type Chat_GetUsersServer interface {
	Send(*User) error
	grpc.ServerStream
}

type chatGetUsersServer struct {
	grpc.ServerStream
}

func (x *chatGetUsersServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

// Chat_ServiceDesc is the grpc.ServiceDesc for Chat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckUser",
			Handler:    _Chat_CheckUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Messaging",
			Handler:       _Chat_Messaging_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GetUsers",
			Handler:       _Chat_GetUsers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chat/chat.proto",
}
