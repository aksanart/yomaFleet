// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.0
// source: pkg/grpc_client/vehicle/contract/grpc.proto

package contract

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
	VehicleService_HealthCheck_FullMethodName   = "/vehicleservice.VehicleService/HealthCheck"
	VehicleService_CreateVehicle_FullMethodName = "/vehicleservice.VehicleService/CreateVehicle"
	VehicleService_UpdateVehicle_FullMethodName = "/vehicleservice.VehicleService/UpdateVehicle"
	VehicleService_ListVehicle_FullMethodName   = "/vehicleservice.VehicleService/ListVehicle"
	VehicleService_DetailVehicle_FullMethodName = "/vehicleservice.VehicleService/DetailVehicle"
	VehicleService_DeleteVehicle_FullMethodName = "/vehicleservice.VehicleService/DeleteVehicle"
)

// VehicleServiceClient is the client API for VehicleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VehicleServiceClient interface {
	HealthCheck(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*DefaultResponse, error)
	CreateVehicle(ctx context.Context, in *CreateVehicleReq, opts ...grpc.CallOption) (*CreateVehicleResponse, error)
	UpdateVehicle(ctx context.Context, in *UpdateVehicleReq, opts ...grpc.CallOption) (*DefaultResponse, error)
	ListVehicle(ctx context.Context, in *ListVehicleReq, opts ...grpc.CallOption) (*ListVehicleResponse, error)
	DetailVehicle(ctx context.Context, in *IDVehicleReq, opts ...grpc.CallOption) (*DetailVehicleResponse, error)
	DeleteVehicle(ctx context.Context, in *IDVehicleReq, opts ...grpc.CallOption) (*DefaultResponse, error)
}

type vehicleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVehicleServiceClient(cc grpc.ClientConnInterface) VehicleServiceClient {
	return &vehicleServiceClient{cc}
}

func (c *vehicleServiceClient) HealthCheck(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*DefaultResponse, error) {
	out := new(DefaultResponse)
	err := c.cc.Invoke(ctx, VehicleService_HealthCheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleServiceClient) CreateVehicle(ctx context.Context, in *CreateVehicleReq, opts ...grpc.CallOption) (*CreateVehicleResponse, error) {
	out := new(CreateVehicleResponse)
	err := c.cc.Invoke(ctx, VehicleService_CreateVehicle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleServiceClient) UpdateVehicle(ctx context.Context, in *UpdateVehicleReq, opts ...grpc.CallOption) (*DefaultResponse, error) {
	out := new(DefaultResponse)
	err := c.cc.Invoke(ctx, VehicleService_UpdateVehicle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleServiceClient) ListVehicle(ctx context.Context, in *ListVehicleReq, opts ...grpc.CallOption) (*ListVehicleResponse, error) {
	out := new(ListVehicleResponse)
	err := c.cc.Invoke(ctx, VehicleService_ListVehicle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleServiceClient) DetailVehicle(ctx context.Context, in *IDVehicleReq, opts ...grpc.CallOption) (*DetailVehicleResponse, error) {
	out := new(DetailVehicleResponse)
	err := c.cc.Invoke(ctx, VehicleService_DetailVehicle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleServiceClient) DeleteVehicle(ctx context.Context, in *IDVehicleReq, opts ...grpc.CallOption) (*DefaultResponse, error) {
	out := new(DefaultResponse)
	err := c.cc.Invoke(ctx, VehicleService_DeleteVehicle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VehicleServiceServer is the server API for VehicleService service.
// All implementations must embed UnimplementedVehicleServiceServer
// for forward compatibility
type VehicleServiceServer interface {
	HealthCheck(context.Context, *EmptyRequest) (*DefaultResponse, error)
	CreateVehicle(context.Context, *CreateVehicleReq) (*CreateVehicleResponse, error)
	UpdateVehicle(context.Context, *UpdateVehicleReq) (*DefaultResponse, error)
	ListVehicle(context.Context, *ListVehicleReq) (*ListVehicleResponse, error)
	DetailVehicle(context.Context, *IDVehicleReq) (*DetailVehicleResponse, error)
	DeleteVehicle(context.Context, *IDVehicleReq) (*DefaultResponse, error)
	mustEmbedUnimplementedVehicleServiceServer()
}

// UnimplementedVehicleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVehicleServiceServer struct {
}

func (UnimplementedVehicleServiceServer) HealthCheck(context.Context, *EmptyRequest) (*DefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedVehicleServiceServer) CreateVehicle(context.Context, *CreateVehicleReq) (*CreateVehicleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVehicle not implemented")
}
func (UnimplementedVehicleServiceServer) UpdateVehicle(context.Context, *UpdateVehicleReq) (*DefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVehicle not implemented")
}
func (UnimplementedVehicleServiceServer) ListVehicle(context.Context, *ListVehicleReq) (*ListVehicleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVehicle not implemented")
}
func (UnimplementedVehicleServiceServer) DetailVehicle(context.Context, *IDVehicleReq) (*DetailVehicleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetailVehicle not implemented")
}
func (UnimplementedVehicleServiceServer) DeleteVehicle(context.Context, *IDVehicleReq) (*DefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVehicle not implemented")
}
func (UnimplementedVehicleServiceServer) mustEmbedUnimplementedVehicleServiceServer() {}

// UnsafeVehicleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VehicleServiceServer will
// result in compilation errors.
type UnsafeVehicleServiceServer interface {
	mustEmbedUnimplementedVehicleServiceServer()
}

func RegisterVehicleServiceServer(s grpc.ServiceRegistrar, srv VehicleServiceServer) {
	s.RegisterService(&VehicleService_ServiceDesc, srv)
}

func _VehicleService_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleServiceServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VehicleService_HealthCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleServiceServer).HealthCheck(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleService_CreateVehicle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVehicleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleServiceServer).CreateVehicle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VehicleService_CreateVehicle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleServiceServer).CreateVehicle(ctx, req.(*CreateVehicleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleService_UpdateVehicle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateVehicleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleServiceServer).UpdateVehicle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VehicleService_UpdateVehicle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleServiceServer).UpdateVehicle(ctx, req.(*UpdateVehicleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleService_ListVehicle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVehicleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleServiceServer).ListVehicle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VehicleService_ListVehicle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleServiceServer).ListVehicle(ctx, req.(*ListVehicleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleService_DetailVehicle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDVehicleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleServiceServer).DetailVehicle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VehicleService_DetailVehicle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleServiceServer).DetailVehicle(ctx, req.(*IDVehicleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleService_DeleteVehicle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDVehicleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleServiceServer).DeleteVehicle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VehicleService_DeleteVehicle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleServiceServer).DeleteVehicle(ctx, req.(*IDVehicleReq))
	}
	return interceptor(ctx, in, info, handler)
}

// VehicleService_ServiceDesc is the grpc.ServiceDesc for VehicleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VehicleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vehicleservice.VehicleService",
	HandlerType: (*VehicleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _VehicleService_HealthCheck_Handler,
		},
		{
			MethodName: "CreateVehicle",
			Handler:    _VehicleService_CreateVehicle_Handler,
		},
		{
			MethodName: "UpdateVehicle",
			Handler:    _VehicleService_UpdateVehicle_Handler,
		},
		{
			MethodName: "ListVehicle",
			Handler:    _VehicleService_ListVehicle_Handler,
		},
		{
			MethodName: "DetailVehicle",
			Handler:    _VehicleService_DetailVehicle_Handler,
		},
		{
			MethodName: "DeleteVehicle",
			Handler:    _VehicleService_DeleteVehicle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/grpc_client/vehicle/contract/grpc.proto",
}
