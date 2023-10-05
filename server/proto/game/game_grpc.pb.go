// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: game.proto

package game

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

const (
	GameService_ProxyGameboard_FullMethodName  = "/game.GameService/ProxyGameboard"
	GameService_StreamGameboard_FullMethodName = "/game.GameService/StreamGameboard"
)

// GameServiceClient is the client API for GameService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameServiceClient interface {
	ProxyGameboard(ctx context.Context, opts ...grpc.CallOption) (GameService_ProxyGameboardClient, error)
	StreamGameboard(ctx context.Context, opts ...grpc.CallOption) (GameService_StreamGameboardClient, error)
}

type gameServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGameServiceClient(cc grpc.ClientConnInterface) GameServiceClient {
	return &gameServiceClient{cc}
}

func (c *gameServiceClient) ProxyGameboard(ctx context.Context, opts ...grpc.CallOption) (GameService_ProxyGameboardClient, error) {
	stream, err := c.cc.NewStream(ctx, &GameService_ServiceDesc.Streams[0], GameService_ProxyGameboard_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &gameServiceProxyGameboardClient{stream}
	return x, nil
}

type GameService_ProxyGameboardClient interface {
	Send(*Move) error
	Recv() (*GameBoard, error)
	grpc.ClientStream
}

type gameServiceProxyGameboardClient struct {
	grpc.ClientStream
}

func (x *gameServiceProxyGameboardClient) Send(m *Move) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gameServiceProxyGameboardClient) Recv() (*GameBoard, error) {
	m := new(GameBoard)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gameServiceClient) StreamGameboard(ctx context.Context, opts ...grpc.CallOption) (GameService_StreamGameboardClient, error) {
	stream, err := c.cc.NewStream(ctx, &GameService_ServiceDesc.Streams[1], GameService_StreamGameboard_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &gameServiceStreamGameboardClient{stream}
	return x, nil
}

type GameService_StreamGameboardClient interface {
	Send(*Move) error
	Recv() (*GameBoard, error)
	grpc.ClientStream
}

type gameServiceStreamGameboardClient struct {
	grpc.ClientStream
}

func (x *gameServiceStreamGameboardClient) Send(m *Move) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gameServiceStreamGameboardClient) Recv() (*GameBoard, error) {
	m := new(GameBoard)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GameServiceServer is the server API for GameService service.
// All implementations must embed UnimplementedGameServiceServer
// for forward compatibility
type GameServiceServer interface {
	ProxyGameboard(GameService_ProxyGameboardServer) error
	StreamGameboard(GameService_StreamGameboardServer) error
	mustEmbedUnimplementedGameServiceServer()
}

// UnimplementedGameServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGameServiceServer struct {
}

func (UnimplementedGameServiceServer) ProxyGameboard(GameService_ProxyGameboardServer) error {
	return status.Errorf(codes.Unimplemented, "method ProxyGameboard not implemented")
}
func (UnimplementedGameServiceServer) StreamGameboard(GameService_StreamGameboardServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamGameboard not implemented")
}
func (UnimplementedGameServiceServer) mustEmbedUnimplementedGameServiceServer() {}

// UnsafeGameServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameServiceServer will
// result in compilation errors.
type UnsafeGameServiceServer interface {
	mustEmbedUnimplementedGameServiceServer()
}

func RegisterGameServiceServer(s grpc.ServiceRegistrar, srv GameServiceServer) {
	s.RegisterService(&GameService_ServiceDesc, srv)
}

func _GameService_ProxyGameboard_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GameServiceServer).ProxyGameboard(&gameServiceProxyGameboardServer{stream})
}

type GameService_ProxyGameboardServer interface {
	Send(*GameBoard) error
	Recv() (*Move, error)
	grpc.ServerStream
}

type gameServiceProxyGameboardServer struct {
	grpc.ServerStream
}

func (x *gameServiceProxyGameboardServer) Send(m *GameBoard) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gameServiceProxyGameboardServer) Recv() (*Move, error) {
	m := new(Move)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GameService_StreamGameboard_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GameServiceServer).StreamGameboard(&gameServiceStreamGameboardServer{stream})
}

type GameService_StreamGameboardServer interface {
	Send(*GameBoard) error
	Recv() (*Move, error)
	grpc.ServerStream
}

type gameServiceStreamGameboardServer struct {
	grpc.ServerStream
}

func (x *gameServiceStreamGameboardServer) Send(m *GameBoard) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gameServiceStreamGameboardServer) Recv() (*Move, error) {
	m := new(Move)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GameService_ServiceDesc is the grpc.ServiceDesc for GameService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GameService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "game.GameService",
	HandlerType: (*GameServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ProxyGameboard",
			Handler:       _GameService_ProxyGameboard_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamGameboard",
			Handler:       _GameService_StreamGameboard_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "game.proto",
}
