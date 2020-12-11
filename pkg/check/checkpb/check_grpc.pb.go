// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package checkpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CheckServiceClient is the client API for CheckService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CheckServiceClient interface {
	GetAllByRoundID(ctx context.Context, in *GetAllByRoundIDRequest, opts ...grpc.CallOption) (*GetAllByRoundIDResponse, error)
	GetByRoundServiceID(ctx context.Context, in *GetByRoundServiceIDRequest, opts ...grpc.CallOption) (*GetByRoundServiceIDResponse, error)
	GetAllByServiceID(ctx context.Context, in *GetAllByServiceIDRequest, opts ...grpc.CallOption) (*GetAllByServiceIDResponse, error)
}

type checkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCheckServiceClient(cc grpc.ClientConnInterface) CheckServiceClient {
	return &checkServiceClient{cc}
}

func (c *checkServiceClient) GetAllByRoundID(ctx context.Context, in *GetAllByRoundIDRequest, opts ...grpc.CallOption) (*GetAllByRoundIDResponse, error) {
	out := new(GetAllByRoundIDResponse)
	err := c.cc.Invoke(ctx, "/pkg.check.checkpb.CheckService/GetAllByRoundID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkServiceClient) GetByRoundServiceID(ctx context.Context, in *GetByRoundServiceIDRequest, opts ...grpc.CallOption) (*GetByRoundServiceIDResponse, error) {
	out := new(GetByRoundServiceIDResponse)
	err := c.cc.Invoke(ctx, "/pkg.check.checkpb.CheckService/GetByRoundServiceID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkServiceClient) GetAllByServiceID(ctx context.Context, in *GetAllByServiceIDRequest, opts ...grpc.CallOption) (*GetAllByServiceIDResponse, error) {
	out := new(GetAllByServiceIDResponse)
	err := c.cc.Invoke(ctx, "/pkg.check.checkpb.CheckService/GetAllByServiceID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckServiceServer is the server API for CheckService service.
// All implementations should embed UnimplementedCheckServiceServer
// for forward compatibility
type CheckServiceServer interface {
	GetAllByRoundID(context.Context, *GetAllByRoundIDRequest) (*GetAllByRoundIDResponse, error)
	GetByRoundServiceID(context.Context, *GetByRoundServiceIDRequest) (*GetByRoundServiceIDResponse, error)
	GetAllByServiceID(context.Context, *GetAllByServiceIDRequest) (*GetAllByServiceIDResponse, error)
}

// UnimplementedCheckServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCheckServiceServer struct {
}

func (*UnimplementedCheckServiceServer) GetAllByRoundID(context.Context, *GetAllByRoundIDRequest) (*GetAllByRoundIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllByRoundID not implemented")
}
func (*UnimplementedCheckServiceServer) GetByRoundServiceID(context.Context, *GetByRoundServiceIDRequest) (*GetByRoundServiceIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByRoundServiceID not implemented")
}
func (*UnimplementedCheckServiceServer) GetAllByServiceID(context.Context, *GetAllByServiceIDRequest) (*GetAllByServiceIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllByServiceID not implemented")
}

func RegisterCheckServiceServer(s *grpc.Server, srv CheckServiceServer) {
	s.RegisterService(&_CheckService_serviceDesc, srv)
}

func _CheckService_GetAllByRoundID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllByRoundIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckServiceServer).GetAllByRoundID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.check.checkpb.CheckService/GetAllByRoundID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckServiceServer).GetAllByRoundID(ctx, req.(*GetAllByRoundIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CheckService_GetByRoundServiceID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByRoundServiceIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckServiceServer).GetByRoundServiceID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.check.checkpb.CheckService/GetByRoundServiceID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckServiceServer).GetByRoundServiceID(ctx, req.(*GetByRoundServiceIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CheckService_GetAllByServiceID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllByServiceIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckServiceServer).GetAllByServiceID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.check.checkpb.CheckService/GetAllByServiceID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckServiceServer).GetAllByServiceID(ctx, req.(*GetAllByServiceIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CheckService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pkg.check.checkpb.CheckService",
	HandlerType: (*CheckServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllByRoundID",
			Handler:    _CheckService_GetAllByRoundID_Handler,
		},
		{
			MethodName: "GetByRoundServiceID",
			Handler:    _CheckService_GetByRoundServiceID_Handler,
		},
		{
			MethodName: "GetAllByServiceID",
			Handler:    _CheckService_GetAllByServiceID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/check/checkpb/check.proto",
}
