// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.1
// source: proto/v2/fasilitas-service.proto

package v2

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
	FasilitasService_GetMasterRegionByParam_FullMethodName   = "/service.fasilitas.master.proto.v2.Fasilitas_service/GetMasterRegionByParam"
	FasilitasService_GetMasterRegionByParamId_FullMethodName = "/service.fasilitas.master.proto.v2.Fasilitas_service/GetMasterRegionByParamId"
)

// FasilitasServiceClient is the client API for FasilitasService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FasilitasServiceClient interface {
	GetMasterRegionByParam(ctx context.Context, in *GetMasterRegionByParamReq, opts ...grpc.CallOption) (*MasterRegion, error)
	GetMasterRegionByParamId(ctx context.Context, in *GetMasterRegionByParamReq, opts ...grpc.CallOption) (*MasterRegion, error)
}

type fasilitasServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFasilitasServiceClient(cc grpc.ClientConnInterface) FasilitasServiceClient {
	return &fasilitasServiceClient{cc}
}

func (c *fasilitasServiceClient) GetMasterRegionByParam(ctx context.Context, in *GetMasterRegionByParamReq, opts ...grpc.CallOption) (*MasterRegion, error) {
	out := new(MasterRegion)
	err := c.cc.Invoke(ctx, FasilitasService_GetMasterRegionByParam_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fasilitasServiceClient) GetMasterRegionByParamId(ctx context.Context, in *GetMasterRegionByParamReq, opts ...grpc.CallOption) (*MasterRegion, error) {
	out := new(MasterRegion)
	err := c.cc.Invoke(ctx, FasilitasService_GetMasterRegionByParamId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FasilitasServiceServer is the server API for FasilitasService service.
// All implementations must embed UnimplementedFasilitasServiceServer
// for forward compatibility
type FasilitasServiceServer interface {
	GetMasterRegionByParam(context.Context, *GetMasterRegionByParamReq) (*MasterRegion, error)
	GetMasterRegionByParamId(context.Context, *GetMasterRegionByParamReq) (*MasterRegion, error)
	mustEmbedUnimplementedFasilitasServiceServer()
}

// UnimplementedFasilitasServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFasilitasServiceServer struct {
}

func (UnimplementedFasilitasServiceServer) GetMasterRegionByParam(context.Context, *GetMasterRegionByParamReq) (*MasterRegion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMasterRegionByParam not implemented")
}
func (UnimplementedFasilitasServiceServer) GetMasterRegionByParamId(context.Context, *GetMasterRegionByParamReq) (*MasterRegion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMasterRegionByParamId not implemented")
}
func (UnimplementedFasilitasServiceServer) mustEmbedUnimplementedFasilitasServiceServer() {}

// UnsafeFasilitasServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FasilitasServiceServer will
// result in compilation errors.
type UnsafeFasilitasServiceServer interface {
	mustEmbedUnimplementedFasilitasServiceServer()
}

func RegisterFasilitasServiceServer(s grpc.ServiceRegistrar, srv FasilitasServiceServer) {
	s.RegisterService(&FasilitasService_ServiceDesc, srv)
}

func _FasilitasService_GetMasterRegionByParam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMasterRegionByParamReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FasilitasServiceServer).GetMasterRegionByParam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FasilitasService_GetMasterRegionByParam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FasilitasServiceServer).GetMasterRegionByParam(ctx, req.(*GetMasterRegionByParamReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FasilitasService_GetMasterRegionByParamId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMasterRegionByParamReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FasilitasServiceServer).GetMasterRegionByParamId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FasilitasService_GetMasterRegionByParamId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FasilitasServiceServer).GetMasterRegionByParamId(ctx, req.(*GetMasterRegionByParamReq))
	}
	return interceptor(ctx, in, info, handler)
}

// FasilitasService_ServiceDesc is the grpc.ServiceDesc for FasilitasService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FasilitasService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.fasilitas.master.proto.v2.Fasilitas_service",
	HandlerType: (*FasilitasServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMasterRegionByParam",
			Handler:    _FasilitasService_GetMasterRegionByParam_Handler,
		},
		{
			MethodName: "GetMasterRegionByParamId",
			Handler:    _FasilitasService_GetMasterRegionByParamId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v2/fasilitas-service.proto",
}
