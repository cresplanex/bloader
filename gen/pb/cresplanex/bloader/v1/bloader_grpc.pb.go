// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: cresplanex/bloader/v1/bloader.proto

package bloaderv1

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
	BloaderSlaveService_Connect_FullMethodName                  = "/cresplanex.bloader.v1.BloaderSlaveService/Connect"
	BloaderSlaveService_Disconnect_FullMethodName               = "/cresplanex.bloader.v1.BloaderSlaveService/Disconnect"
	BloaderSlaveService_SlaveCommand_FullMethodName             = "/cresplanex.bloader.v1.BloaderSlaveService/SlaveCommand"
	BloaderSlaveService_SlaveCommandDefaultStore_FullMethodName = "/cresplanex.bloader.v1.BloaderSlaveService/SlaveCommandDefaultStore"
	BloaderSlaveService_CallExec_FullMethodName                 = "/cresplanex.bloader.v1.BloaderSlaveService/CallExec"
	BloaderSlaveService_ReceiveChanelConnect_FullMethodName     = "/cresplanex.bloader.v1.BloaderSlaveService/ReceiveChanelConnect"
	BloaderSlaveService_SendLoader_FullMethodName               = "/cresplanex.bloader.v1.BloaderSlaveService/SendLoader"
	BloaderSlaveService_SendAuth_FullMethodName                 = "/cresplanex.bloader.v1.BloaderSlaveService/SendAuth"
	BloaderSlaveService_SendStoreData_FullMethodName            = "/cresplanex.bloader.v1.BloaderSlaveService/SendStoreData"
	BloaderSlaveService_SendStoreOk_FullMethodName              = "/cresplanex.bloader.v1.BloaderSlaveService/SendStoreOk"
	BloaderSlaveService_SendTarget_FullMethodName               = "/cresplanex.bloader.v1.BloaderSlaveService/SendTarget"
	BloaderSlaveService_ReceiveLoadTermChannel_FullMethodName   = "/cresplanex.bloader.v1.BloaderSlaveService/ReceiveLoadTermChannel"
)

// BloaderSlaveServiceClient is the client API for BloaderSlaveService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BloaderSlaveServiceClient interface {
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error)
	Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error)
	SlaveCommand(ctx context.Context, in *SlaveCommandRequest, opts ...grpc.CallOption) (*SlaveCommandResponse, error)
	SlaveCommandDefaultStore(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[SlaveCommandDefaultStoreRequest, SlaveCommandDefaultStoreResponse], error)
	CallExec(ctx context.Context, in *CallExecRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[CallExecResponse], error)
	ReceiveChanelConnect(ctx context.Context, in *ReceiveChanelConnectRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ReceiveChanelConnectResponse], error)
	SendLoader(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[SendLoaderRequest, SendLoaderResponse], error)
	SendAuth(ctx context.Context, in *SendAuthRequest, opts ...grpc.CallOption) (*SendAuthResponse, error)
	SendStoreData(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[SendStoreDataRequest, SendStoreDataResponse], error)
	SendStoreOk(ctx context.Context, in *SendStoreOkRequest, opts ...grpc.CallOption) (*SendStoreOkResponse, error)
	SendTarget(ctx context.Context, in *SendTargetRequest, opts ...grpc.CallOption) (*SendTargetResponse, error)
	ReceiveLoadTermChannel(ctx context.Context, in *ReceiveLoadTermChannelRequest, opts ...grpc.CallOption) (*ReceiveLoadTermChannelResponse, error)
}

type bloaderSlaveServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBloaderSlaveServiceClient(cc grpc.ClientConnInterface) BloaderSlaveServiceClient {
	return &bloaderSlaveServiceClient{cc}
}

func (c *bloaderSlaveServiceClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConnectResponse)
	err := c.cc.Invoke(ctx, BloaderSlaveService_Connect_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bloaderSlaveServiceClient) Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DisconnectResponse)
	err := c.cc.Invoke(ctx, BloaderSlaveService_Disconnect_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bloaderSlaveServiceClient) SlaveCommand(ctx context.Context, in *SlaveCommandRequest, opts ...grpc.CallOption) (*SlaveCommandResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SlaveCommandResponse)
	err := c.cc.Invoke(ctx, BloaderSlaveService_SlaveCommand_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bloaderSlaveServiceClient) SlaveCommandDefaultStore(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[SlaveCommandDefaultStoreRequest, SlaveCommandDefaultStoreResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &BloaderSlaveService_ServiceDesc.Streams[0], BloaderSlaveService_SlaveCommandDefaultStore_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SlaveCommandDefaultStoreRequest, SlaveCommandDefaultStoreResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type BloaderSlaveService_SlaveCommandDefaultStoreClient = grpc.ClientStreamingClient[SlaveCommandDefaultStoreRequest, SlaveCommandDefaultStoreResponse]

func (c *bloaderSlaveServiceClient) CallExec(ctx context.Context, in *CallExecRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[CallExecResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &BloaderSlaveService_ServiceDesc.Streams[1], BloaderSlaveService_CallExec_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[CallExecRequest, CallExecResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type BloaderSlaveService_CallExecClient = grpc.ServerStreamingClient[CallExecResponse]

func (c *bloaderSlaveServiceClient) ReceiveChanelConnect(ctx context.Context, in *ReceiveChanelConnectRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ReceiveChanelConnectResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &BloaderSlaveService_ServiceDesc.Streams[2], BloaderSlaveService_ReceiveChanelConnect_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ReceiveChanelConnectRequest, ReceiveChanelConnectResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type BloaderSlaveService_ReceiveChanelConnectClient = grpc.ServerStreamingClient[ReceiveChanelConnectResponse]

func (c *bloaderSlaveServiceClient) SendLoader(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[SendLoaderRequest, SendLoaderResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &BloaderSlaveService_ServiceDesc.Streams[3], BloaderSlaveService_SendLoader_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SendLoaderRequest, SendLoaderResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type BloaderSlaveService_SendLoaderClient = grpc.ClientStreamingClient[SendLoaderRequest, SendLoaderResponse]

func (c *bloaderSlaveServiceClient) SendAuth(ctx context.Context, in *SendAuthRequest, opts ...grpc.CallOption) (*SendAuthResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendAuthResponse)
	err := c.cc.Invoke(ctx, BloaderSlaveService_SendAuth_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bloaderSlaveServiceClient) SendStoreData(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[SendStoreDataRequest, SendStoreDataResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &BloaderSlaveService_ServiceDesc.Streams[4], BloaderSlaveService_SendStoreData_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SendStoreDataRequest, SendStoreDataResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type BloaderSlaveService_SendStoreDataClient = grpc.ClientStreamingClient[SendStoreDataRequest, SendStoreDataResponse]

func (c *bloaderSlaveServiceClient) SendStoreOk(ctx context.Context, in *SendStoreOkRequest, opts ...grpc.CallOption) (*SendStoreOkResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendStoreOkResponse)
	err := c.cc.Invoke(ctx, BloaderSlaveService_SendStoreOk_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bloaderSlaveServiceClient) SendTarget(ctx context.Context, in *SendTargetRequest, opts ...grpc.CallOption) (*SendTargetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendTargetResponse)
	err := c.cc.Invoke(ctx, BloaderSlaveService_SendTarget_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bloaderSlaveServiceClient) ReceiveLoadTermChannel(ctx context.Context, in *ReceiveLoadTermChannelRequest, opts ...grpc.CallOption) (*ReceiveLoadTermChannelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReceiveLoadTermChannelResponse)
	err := c.cc.Invoke(ctx, BloaderSlaveService_ReceiveLoadTermChannel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BloaderSlaveServiceServer is the server API for BloaderSlaveService service.
// All implementations must embed UnimplementedBloaderSlaveServiceServer
// for forward compatibility.
type BloaderSlaveServiceServer interface {
	Connect(context.Context, *ConnectRequest) (*ConnectResponse, error)
	Disconnect(context.Context, *DisconnectRequest) (*DisconnectResponse, error)
	SlaveCommand(context.Context, *SlaveCommandRequest) (*SlaveCommandResponse, error)
	SlaveCommandDefaultStore(grpc.ClientStreamingServer[SlaveCommandDefaultStoreRequest, SlaveCommandDefaultStoreResponse]) error
	CallExec(*CallExecRequest, grpc.ServerStreamingServer[CallExecResponse]) error
	ReceiveChanelConnect(*ReceiveChanelConnectRequest, grpc.ServerStreamingServer[ReceiveChanelConnectResponse]) error
	SendLoader(grpc.ClientStreamingServer[SendLoaderRequest, SendLoaderResponse]) error
	SendAuth(context.Context, *SendAuthRequest) (*SendAuthResponse, error)
	SendStoreData(grpc.ClientStreamingServer[SendStoreDataRequest, SendStoreDataResponse]) error
	SendStoreOk(context.Context, *SendStoreOkRequest) (*SendStoreOkResponse, error)
	SendTarget(context.Context, *SendTargetRequest) (*SendTargetResponse, error)
	ReceiveLoadTermChannel(context.Context, *ReceiveLoadTermChannelRequest) (*ReceiveLoadTermChannelResponse, error)
	mustEmbedUnimplementedBloaderSlaveServiceServer()
}

// UnimplementedBloaderSlaveServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBloaderSlaveServiceServer struct{}

func (UnimplementedBloaderSlaveServiceServer) Connect(context.Context, *ConnectRequest) (*ConnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedBloaderSlaveServiceServer) Disconnect(context.Context, *DisconnectRequest) (*DisconnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}
func (UnimplementedBloaderSlaveServiceServer) SlaveCommand(context.Context, *SlaveCommandRequest) (*SlaveCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SlaveCommand not implemented")
}
func (UnimplementedBloaderSlaveServiceServer) SlaveCommandDefaultStore(grpc.ClientStreamingServer[SlaveCommandDefaultStoreRequest, SlaveCommandDefaultStoreResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SlaveCommandDefaultStore not implemented")
}
func (UnimplementedBloaderSlaveServiceServer) CallExec(*CallExecRequest, grpc.ServerStreamingServer[CallExecResponse]) error {
	return status.Errorf(codes.Unimplemented, "method CallExec not implemented")
}
func (UnimplementedBloaderSlaveServiceServer) ReceiveChanelConnect(*ReceiveChanelConnectRequest, grpc.ServerStreamingServer[ReceiveChanelConnectResponse]) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveChanelConnect not implemented")
}
func (UnimplementedBloaderSlaveServiceServer) SendLoader(grpc.ClientStreamingServer[SendLoaderRequest, SendLoaderResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SendLoader not implemented")
}
func (UnimplementedBloaderSlaveServiceServer) SendAuth(context.Context, *SendAuthRequest) (*SendAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendAuth not implemented")
}
func (UnimplementedBloaderSlaveServiceServer) SendStoreData(grpc.ClientStreamingServer[SendStoreDataRequest, SendStoreDataResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SendStoreData not implemented")
}
func (UnimplementedBloaderSlaveServiceServer) SendStoreOk(context.Context, *SendStoreOkRequest) (*SendStoreOkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendStoreOk not implemented")
}
func (UnimplementedBloaderSlaveServiceServer) SendTarget(context.Context, *SendTargetRequest) (*SendTargetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTarget not implemented")
}
func (UnimplementedBloaderSlaveServiceServer) ReceiveLoadTermChannel(context.Context, *ReceiveLoadTermChannelRequest) (*ReceiveLoadTermChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveLoadTermChannel not implemented")
}
func (UnimplementedBloaderSlaveServiceServer) mustEmbedUnimplementedBloaderSlaveServiceServer() {}
func (UnimplementedBloaderSlaveServiceServer) testEmbeddedByValue()                             {}

// UnsafeBloaderSlaveServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BloaderSlaveServiceServer will
// result in compilation errors.
type UnsafeBloaderSlaveServiceServer interface {
	mustEmbedUnimplementedBloaderSlaveServiceServer()
}

func RegisterBloaderSlaveServiceServer(s grpc.ServiceRegistrar, srv BloaderSlaveServiceServer) {
	// If the following call pancis, it indicates UnimplementedBloaderSlaveServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BloaderSlaveService_ServiceDesc, srv)
}

func _BloaderSlaveService_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BloaderSlaveServiceServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BloaderSlaveService_Connect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BloaderSlaveServiceServer).Connect(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BloaderSlaveService_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BloaderSlaveServiceServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BloaderSlaveService_Disconnect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BloaderSlaveServiceServer).Disconnect(ctx, req.(*DisconnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BloaderSlaveService_SlaveCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SlaveCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BloaderSlaveServiceServer).SlaveCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BloaderSlaveService_SlaveCommand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BloaderSlaveServiceServer).SlaveCommand(ctx, req.(*SlaveCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BloaderSlaveService_SlaveCommandDefaultStore_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BloaderSlaveServiceServer).SlaveCommandDefaultStore(&grpc.GenericServerStream[SlaveCommandDefaultStoreRequest, SlaveCommandDefaultStoreResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type BloaderSlaveService_SlaveCommandDefaultStoreServer = grpc.ClientStreamingServer[SlaveCommandDefaultStoreRequest, SlaveCommandDefaultStoreResponse]

func _BloaderSlaveService_CallExec_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CallExecRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BloaderSlaveServiceServer).CallExec(m, &grpc.GenericServerStream[CallExecRequest, CallExecResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type BloaderSlaveService_CallExecServer = grpc.ServerStreamingServer[CallExecResponse]

func _BloaderSlaveService_ReceiveChanelConnect_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReceiveChanelConnectRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BloaderSlaveServiceServer).ReceiveChanelConnect(m, &grpc.GenericServerStream[ReceiveChanelConnectRequest, ReceiveChanelConnectResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type BloaderSlaveService_ReceiveChanelConnectServer = grpc.ServerStreamingServer[ReceiveChanelConnectResponse]

func _BloaderSlaveService_SendLoader_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BloaderSlaveServiceServer).SendLoader(&grpc.GenericServerStream[SendLoaderRequest, SendLoaderResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type BloaderSlaveService_SendLoaderServer = grpc.ClientStreamingServer[SendLoaderRequest, SendLoaderResponse]

func _BloaderSlaveService_SendAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BloaderSlaveServiceServer).SendAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BloaderSlaveService_SendAuth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BloaderSlaveServiceServer).SendAuth(ctx, req.(*SendAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BloaderSlaveService_SendStoreData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BloaderSlaveServiceServer).SendStoreData(&grpc.GenericServerStream[SendStoreDataRequest, SendStoreDataResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type BloaderSlaveService_SendStoreDataServer = grpc.ClientStreamingServer[SendStoreDataRequest, SendStoreDataResponse]

func _BloaderSlaveService_SendStoreOk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendStoreOkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BloaderSlaveServiceServer).SendStoreOk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BloaderSlaveService_SendStoreOk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BloaderSlaveServiceServer).SendStoreOk(ctx, req.(*SendStoreOkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BloaderSlaveService_SendTarget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTargetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BloaderSlaveServiceServer).SendTarget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BloaderSlaveService_SendTarget_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BloaderSlaveServiceServer).SendTarget(ctx, req.(*SendTargetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BloaderSlaveService_ReceiveLoadTermChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiveLoadTermChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BloaderSlaveServiceServer).ReceiveLoadTermChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BloaderSlaveService_ReceiveLoadTermChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BloaderSlaveServiceServer).ReceiveLoadTermChannel(ctx, req.(*ReceiveLoadTermChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BloaderSlaveService_ServiceDesc is the grpc.ServiceDesc for BloaderSlaveService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BloaderSlaveService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cresplanex.bloader.v1.BloaderSlaveService",
	HandlerType: (*BloaderSlaveServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _BloaderSlaveService_Connect_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _BloaderSlaveService_Disconnect_Handler,
		},
		{
			MethodName: "SlaveCommand",
			Handler:    _BloaderSlaveService_SlaveCommand_Handler,
		},
		{
			MethodName: "SendAuth",
			Handler:    _BloaderSlaveService_SendAuth_Handler,
		},
		{
			MethodName: "SendStoreOk",
			Handler:    _BloaderSlaveService_SendStoreOk_Handler,
		},
		{
			MethodName: "SendTarget",
			Handler:    _BloaderSlaveService_SendTarget_Handler,
		},
		{
			MethodName: "ReceiveLoadTermChannel",
			Handler:    _BloaderSlaveService_ReceiveLoadTermChannel_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SlaveCommandDefaultStore",
			Handler:       _BloaderSlaveService_SlaveCommandDefaultStore_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "CallExec",
			Handler:       _BloaderSlaveService_CallExec_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ReceiveChanelConnect",
			Handler:       _BloaderSlaveService_ReceiveChanelConnect_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendLoader",
			Handler:       _BloaderSlaveService_SendLoader_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SendStoreData",
			Handler:       _BloaderSlaveService_SendStoreData_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "cresplanex/bloader/v1/bloader.proto",
}
