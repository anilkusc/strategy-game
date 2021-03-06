// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protos

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

// StrategyGameClient is the client API for StrategyGame service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StrategyGameClient interface {
	CreateGame(ctx context.Context, in *CreateGameInputs, opts ...grpc.CallOption) (*CreateGameOutputs, error)
	JoinGame(ctx context.Context, in *JoinGameInputs, opts ...grpc.CallOption) (*JoinGameOutputs, error)
	MakeMove(ctx context.Context, in *MoveInputs, opts ...grpc.CallOption) (*MoveOutputs, error)
	GetLastMoves(ctx context.Context, in *LastMovesInputs, opts ...grpc.CallOption) (*LastMovesOutputs, error)
	GetMaps(ctx context.Context, in *GetMapsInput, opts ...grpc.CallOption) (*GetMapsOutput, error)
}

type strategyGameClient struct {
	cc grpc.ClientConnInterface
}

func NewStrategyGameClient(cc grpc.ClientConnInterface) StrategyGameClient {
	return &strategyGameClient{cc}
}

func (c *strategyGameClient) CreateGame(ctx context.Context, in *CreateGameInputs, opts ...grpc.CallOption) (*CreateGameOutputs, error) {
	out := new(CreateGameOutputs)
	err := c.cc.Invoke(ctx, "/api.StrategyGame/CreateGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *strategyGameClient) JoinGame(ctx context.Context, in *JoinGameInputs, opts ...grpc.CallOption) (*JoinGameOutputs, error) {
	out := new(JoinGameOutputs)
	err := c.cc.Invoke(ctx, "/api.StrategyGame/JoinGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *strategyGameClient) MakeMove(ctx context.Context, in *MoveInputs, opts ...grpc.CallOption) (*MoveOutputs, error) {
	out := new(MoveOutputs)
	err := c.cc.Invoke(ctx, "/api.StrategyGame/MakeMove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *strategyGameClient) GetLastMoves(ctx context.Context, in *LastMovesInputs, opts ...grpc.CallOption) (*LastMovesOutputs, error) {
	out := new(LastMovesOutputs)
	err := c.cc.Invoke(ctx, "/api.StrategyGame/GetLastMoves", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *strategyGameClient) GetMaps(ctx context.Context, in *GetMapsInput, opts ...grpc.CallOption) (*GetMapsOutput, error) {
	out := new(GetMapsOutput)
	err := c.cc.Invoke(ctx, "/api.StrategyGame/GetMaps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StrategyGameServer is the server API for StrategyGame service.
// All implementations must embed UnimplementedStrategyGameServer
// for forward compatibility
type StrategyGameServer interface {
	CreateGame(context.Context, *CreateGameInputs) (*CreateGameOutputs, error)
	JoinGame(context.Context, *JoinGameInputs) (*JoinGameOutputs, error)
	MakeMove(context.Context, *MoveInputs) (*MoveOutputs, error)
	GetLastMoves(context.Context, *LastMovesInputs) (*LastMovesOutputs, error)
	GetMaps(context.Context, *GetMapsInput) (*GetMapsOutput, error)
	mustEmbedUnimplementedStrategyGameServer()
}

// UnimplementedStrategyGameServer must be embedded to have forward compatible implementations.
type UnimplementedStrategyGameServer struct {
}

func (UnimplementedStrategyGameServer) CreateGame(context.Context, *CreateGameInputs) (*CreateGameOutputs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGame not implemented")
}
func (UnimplementedStrategyGameServer) JoinGame(context.Context, *JoinGameInputs) (*JoinGameOutputs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinGame not implemented")
}
func (UnimplementedStrategyGameServer) MakeMove(context.Context, *MoveInputs) (*MoveOutputs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeMove not implemented")
}
func (UnimplementedStrategyGameServer) GetLastMoves(context.Context, *LastMovesInputs) (*LastMovesOutputs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastMoves not implemented")
}
func (UnimplementedStrategyGameServer) GetMaps(context.Context, *GetMapsInput) (*GetMapsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMaps not implemented")
}
func (UnimplementedStrategyGameServer) mustEmbedUnimplementedStrategyGameServer() {}

// UnsafeStrategyGameServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StrategyGameServer will
// result in compilation errors.
type UnsafeStrategyGameServer interface {
	mustEmbedUnimplementedStrategyGameServer()
}

func RegisterStrategyGameServer(s grpc.ServiceRegistrar, srv StrategyGameServer) {
	s.RegisterService(&StrategyGame_ServiceDesc, srv)
}

func _StrategyGame_CreateGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGameInputs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StrategyGameServer).CreateGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.StrategyGame/CreateGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StrategyGameServer).CreateGame(ctx, req.(*CreateGameInputs))
	}
	return interceptor(ctx, in, info, handler)
}

func _StrategyGame_JoinGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinGameInputs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StrategyGameServer).JoinGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.StrategyGame/JoinGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StrategyGameServer).JoinGame(ctx, req.(*JoinGameInputs))
	}
	return interceptor(ctx, in, info, handler)
}

func _StrategyGame_MakeMove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveInputs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StrategyGameServer).MakeMove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.StrategyGame/MakeMove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StrategyGameServer).MakeMove(ctx, req.(*MoveInputs))
	}
	return interceptor(ctx, in, info, handler)
}

func _StrategyGame_GetLastMoves_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LastMovesInputs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StrategyGameServer).GetLastMoves(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.StrategyGame/GetLastMoves",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StrategyGameServer).GetLastMoves(ctx, req.(*LastMovesInputs))
	}
	return interceptor(ctx, in, info, handler)
}

func _StrategyGame_GetMaps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMapsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StrategyGameServer).GetMaps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.StrategyGame/GetMaps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StrategyGameServer).GetMaps(ctx, req.(*GetMapsInput))
	}
	return interceptor(ctx, in, info, handler)
}

// StrategyGame_ServiceDesc is the grpc.ServiceDesc for StrategyGame service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StrategyGame_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.StrategyGame",
	HandlerType: (*StrategyGameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGame",
			Handler:    _StrategyGame_CreateGame_Handler,
		},
		{
			MethodName: "JoinGame",
			Handler:    _StrategyGame_JoinGame_Handler,
		},
		{
			MethodName: "MakeMove",
			Handler:    _StrategyGame_MakeMove_Handler,
		},
		{
			MethodName: "GetLastMoves",
			Handler:    _StrategyGame_GetLastMoves_Handler,
		},
		{
			MethodName: "GetMaps",
			Handler:    _StrategyGame_GetMaps_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
