// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: rgpb.proto

package rgpb

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

// RegistrationServiceClient is the client API for RegistrationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegistrationServiceClient interface {
	CreateRegistration(ctx context.Context, in *CreateRegistrationRequest, opts ...grpc.CallOption) (*CreateRegistrationResponse, error)
	GetRegistrarion(ctx context.Context, in *GetRegistrarionRequest, opts ...grpc.CallOption) (*GetRegistrationResponse, error)
	ListRegistration(ctx context.Context, in *ListRegistrationRequest, opts ...grpc.CallOption) (*ListRegistrationResponse, error)
}

type registrationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRegistrationServiceClient(cc grpc.ClientConnInterface) RegistrationServiceClient {
	return &registrationServiceClient{cc}
}

func (c *registrationServiceClient) CreateRegistration(ctx context.Context, in *CreateRegistrationRequest, opts ...grpc.CallOption) (*CreateRegistrationResponse, error) {
	out := new(CreateRegistrationResponse)
	err := c.cc.Invoke(ctx, "/rgpb.RegistrationService/CreateRegistration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationServiceClient) GetRegistrarion(ctx context.Context, in *GetRegistrarionRequest, opts ...grpc.CallOption) (*GetRegistrationResponse, error) {
	out := new(GetRegistrationResponse)
	err := c.cc.Invoke(ctx, "/rgpb.RegistrationService/GetRegistrarion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationServiceClient) ListRegistration(ctx context.Context, in *ListRegistrationRequest, opts ...grpc.CallOption) (*ListRegistrationResponse, error) {
	out := new(ListRegistrationResponse)
	err := c.cc.Invoke(ctx, "/rgpb.RegistrationService/ListRegistration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistrationServiceServer is the server API for RegistrationService service.
// All implementations must embed UnimplementedRegistrationServiceServer
// for forward compatibility
type RegistrationServiceServer interface {
	CreateRegistration(context.Context, *CreateRegistrationRequest) (*CreateRegistrationResponse, error)
	GetRegistrarion(context.Context, *GetRegistrarionRequest) (*GetRegistrationResponse, error)
	ListRegistration(context.Context, *ListRegistrationRequest) (*ListRegistrationResponse, error)
	mustEmbedUnimplementedRegistrationServiceServer()
}

// UnimplementedRegistrationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRegistrationServiceServer struct {
}

func (UnimplementedRegistrationServiceServer) CreateRegistration(context.Context, *CreateRegistrationRequest) (*CreateRegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRegistration not implemented")
}
func (UnimplementedRegistrationServiceServer) GetRegistrarion(context.Context, *GetRegistrarionRequest) (*GetRegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegistrarion not implemented")
}
func (UnimplementedRegistrationServiceServer) ListRegistration(context.Context, *ListRegistrationRequest) (*ListRegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRegistration not implemented")
}
func (UnimplementedRegistrationServiceServer) mustEmbedUnimplementedRegistrationServiceServer() {}

// UnsafeRegistrationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegistrationServiceServer will
// result in compilation errors.
type UnsafeRegistrationServiceServer interface {
	mustEmbedUnimplementedRegistrationServiceServer()
}

func RegisterRegistrationServiceServer(s grpc.ServiceRegistrar, srv RegistrationServiceServer) {
	s.RegisterService(&RegistrationService_ServiceDesc, srv)
}

func _RegistrationService_CreateRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServiceServer).CreateRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rgpb.RegistrationService/CreateRegistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServiceServer).CreateRegistration(ctx, req.(*CreateRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistrationService_GetRegistrarion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegistrarionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServiceServer).GetRegistrarion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rgpb.RegistrationService/GetRegistrarion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServiceServer).GetRegistrarion(ctx, req.(*GetRegistrarionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistrationService_ListRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServiceServer).ListRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rgpb.RegistrationService/ListRegistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServiceServer).ListRegistration(ctx, req.(*ListRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegistrationService_ServiceDesc is the grpc.ServiceDesc for RegistrationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RegistrationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rgpb.RegistrationService",
	HandlerType: (*RegistrationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRegistration",
			Handler:    _RegistrationService_CreateRegistration_Handler,
		},
		{
			MethodName: "GetRegistrarion",
			Handler:    _RegistrationService_GetRegistrarion_Handler,
		},
		{
			MethodName: "ListRegistration",
			Handler:    _RegistrationService_ListRegistration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rgpb.proto",
}
