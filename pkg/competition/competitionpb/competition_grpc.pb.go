// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package competitionpb

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

// CompetitionServiceClient is the client API for CompetitionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompetitionServiceClient interface {
	LoadCompetition(ctx context.Context, in *LoadCompetitionRequest, opts ...grpc.CallOption) (*LoadCompetitionResponse, error)
	FetchCoreCompetition(ctx context.Context, in *FetchCoreCompetitionRequest, opts ...grpc.CallOption) (*FetchCoreCompetitionResponse, error)
	FetchEntireCompetition(ctx context.Context, in *FetchEntireCompetitionRequest, opts ...grpc.CallOption) (*FetchEntireCompetitionResponse, error)
	ResetScores(ctx context.Context, in *ResetScoresRequest, opts ...grpc.CallOption) (*ResetScoresResponse, error)
	DeleteCompetition(ctx context.Context, in *DeleteCompetitionRequest, opts ...grpc.CallOption) (*DeleteCompetitionResponse, error)
}

type competitionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCompetitionServiceClient(cc grpc.ClientConnInterface) CompetitionServiceClient {
	return &competitionServiceClient{cc}
}

func (c *competitionServiceClient) LoadCompetition(ctx context.Context, in *LoadCompetitionRequest, opts ...grpc.CallOption) (*LoadCompetitionResponse, error) {
	out := new(LoadCompetitionResponse)
	err := c.cc.Invoke(ctx, "/pkg.competition.competitionpb.CompetitionService/LoadCompetition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *competitionServiceClient) FetchCoreCompetition(ctx context.Context, in *FetchCoreCompetitionRequest, opts ...grpc.CallOption) (*FetchCoreCompetitionResponse, error) {
	out := new(FetchCoreCompetitionResponse)
	err := c.cc.Invoke(ctx, "/pkg.competition.competitionpb.CompetitionService/FetchCoreCompetition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *competitionServiceClient) FetchEntireCompetition(ctx context.Context, in *FetchEntireCompetitionRequest, opts ...grpc.CallOption) (*FetchEntireCompetitionResponse, error) {
	out := new(FetchEntireCompetitionResponse)
	err := c.cc.Invoke(ctx, "/pkg.competition.competitionpb.CompetitionService/FetchEntireCompetition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *competitionServiceClient) ResetScores(ctx context.Context, in *ResetScoresRequest, opts ...grpc.CallOption) (*ResetScoresResponse, error) {
	out := new(ResetScoresResponse)
	err := c.cc.Invoke(ctx, "/pkg.competition.competitionpb.CompetitionService/ResetScores", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *competitionServiceClient) DeleteCompetition(ctx context.Context, in *DeleteCompetitionRequest, opts ...grpc.CallOption) (*DeleteCompetitionResponse, error) {
	out := new(DeleteCompetitionResponse)
	err := c.cc.Invoke(ctx, "/pkg.competition.competitionpb.CompetitionService/DeleteCompetition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompetitionServiceServer is the server API for CompetitionService service.
// All implementations must embed UnimplementedCompetitionServiceServer
// for forward compatibility
type CompetitionServiceServer interface {
	LoadCompetition(context.Context, *LoadCompetitionRequest) (*LoadCompetitionResponse, error)
	FetchCoreCompetition(context.Context, *FetchCoreCompetitionRequest) (*FetchCoreCompetitionResponse, error)
	FetchEntireCompetition(context.Context, *FetchEntireCompetitionRequest) (*FetchEntireCompetitionResponse, error)
	ResetScores(context.Context, *ResetScoresRequest) (*ResetScoresResponse, error)
	DeleteCompetition(context.Context, *DeleteCompetitionRequest) (*DeleteCompetitionResponse, error)
	mustEmbedUnimplementedCompetitionServiceServer()
}

// UnimplementedCompetitionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCompetitionServiceServer struct {
}

func (UnimplementedCompetitionServiceServer) LoadCompetition(context.Context, *LoadCompetitionRequest) (*LoadCompetitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadCompetition not implemented")
}
func (UnimplementedCompetitionServiceServer) FetchCoreCompetition(context.Context, *FetchCoreCompetitionRequest) (*FetchCoreCompetitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchCoreCompetition not implemented")
}
func (UnimplementedCompetitionServiceServer) FetchEntireCompetition(context.Context, *FetchEntireCompetitionRequest) (*FetchEntireCompetitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchEntireCompetition not implemented")
}
func (UnimplementedCompetitionServiceServer) ResetScores(context.Context, *ResetScoresRequest) (*ResetScoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetScores not implemented")
}
func (UnimplementedCompetitionServiceServer) DeleteCompetition(context.Context, *DeleteCompetitionRequest) (*DeleteCompetitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCompetition not implemented")
}
func (UnimplementedCompetitionServiceServer) mustEmbedUnimplementedCompetitionServiceServer() {}

// UnsafeCompetitionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompetitionServiceServer will
// result in compilation errors.
type UnsafeCompetitionServiceServer interface {
	mustEmbedUnimplementedCompetitionServiceServer()
}

func RegisterCompetitionServiceServer(s grpc.ServiceRegistrar, srv CompetitionServiceServer) {
	s.RegisterService(&CompetitionService_ServiceDesc, srv)
}

func _CompetitionService_LoadCompetition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadCompetitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompetitionServiceServer).LoadCompetition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.competition.competitionpb.CompetitionService/LoadCompetition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompetitionServiceServer).LoadCompetition(ctx, req.(*LoadCompetitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompetitionService_FetchCoreCompetition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchCoreCompetitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompetitionServiceServer).FetchCoreCompetition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.competition.competitionpb.CompetitionService/FetchCoreCompetition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompetitionServiceServer).FetchCoreCompetition(ctx, req.(*FetchCoreCompetitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompetitionService_FetchEntireCompetition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchEntireCompetitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompetitionServiceServer).FetchEntireCompetition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.competition.competitionpb.CompetitionService/FetchEntireCompetition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompetitionServiceServer).FetchEntireCompetition(ctx, req.(*FetchEntireCompetitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompetitionService_ResetScores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetScoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompetitionServiceServer).ResetScores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.competition.competitionpb.CompetitionService/ResetScores",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompetitionServiceServer).ResetScores(ctx, req.(*ResetScoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompetitionService_DeleteCompetition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCompetitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompetitionServiceServer).DeleteCompetition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.competition.competitionpb.CompetitionService/DeleteCompetition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompetitionServiceServer).DeleteCompetition(ctx, req.(*DeleteCompetitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CompetitionService_ServiceDesc is the grpc.ServiceDesc for CompetitionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CompetitionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pkg.competition.competitionpb.CompetitionService",
	HandlerType: (*CompetitionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoadCompetition",
			Handler:    _CompetitionService_LoadCompetition_Handler,
		},
		{
			MethodName: "FetchCoreCompetition",
			Handler:    _CompetitionService_FetchCoreCompetition_Handler,
		},
		{
			MethodName: "FetchEntireCompetition",
			Handler:    _CompetitionService_FetchEntireCompetition_Handler,
		},
		{
			MethodName: "ResetScores",
			Handler:    _CompetitionService_ResetScores_Handler,
		},
		{
			MethodName: "DeleteCompetition",
			Handler:    _CompetitionService_DeleteCompetition_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/competition/competitionpb/competition.proto",
}
