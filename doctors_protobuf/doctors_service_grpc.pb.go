// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: doctors_service.proto

package doctors_protobuf

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
	DoctorService_GetDoctor_FullMethodName     = "/doctors.DoctorService/GetDoctor"
	DoctorService_GetDoctorsIds_FullMethodName = "/doctors.DoctorService/GetDoctorsIds"
)

// DoctorServiceClient is the client API for DoctorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DoctorServiceClient interface {
	GetDoctor(ctx context.Context, in *DoctorRequest, opts ...grpc.CallOption) (*Doctor, error)
	GetDoctorsIds(ctx context.Context, in *DoctorsRequest, opts ...grpc.CallOption) (*PaginatedResponse, error)
}

type doctorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDoctorServiceClient(cc grpc.ClientConnInterface) DoctorServiceClient {
	return &doctorServiceClient{cc}
}

func (c *doctorServiceClient) GetDoctor(ctx context.Context, in *DoctorRequest, opts ...grpc.CallOption) (*Doctor, error) {
	out := new(Doctor)
	err := c.cc.Invoke(ctx, DoctorService_GetDoctor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) GetDoctorsIds(ctx context.Context, in *DoctorsRequest, opts ...grpc.CallOption) (*PaginatedResponse, error) {
	out := new(PaginatedResponse)
	err := c.cc.Invoke(ctx, DoctorService_GetDoctorsIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DoctorServiceServer is the server API for DoctorService service.
// All implementations must embed UnimplementedDoctorServiceServer
// for forward compatibility
type DoctorServiceServer interface {
	GetDoctor(context.Context, *DoctorRequest) (*Doctor, error)
	GetDoctorsIds(context.Context, *DoctorsRequest) (*PaginatedResponse, error)
	mustEmbedUnimplementedDoctorServiceServer()
}

// UnimplementedDoctorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDoctorServiceServer struct {
}

func (UnimplementedDoctorServiceServer) GetDoctor(context.Context, *DoctorRequest) (*Doctor, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDoctor not implemented")
}
func (UnimplementedDoctorServiceServer) GetDoctorsIds(context.Context, *DoctorsRequest) (*PaginatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDoctorsIds not implemented")
}
func (UnimplementedDoctorServiceServer) mustEmbedUnimplementedDoctorServiceServer() {}

// UnsafeDoctorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DoctorServiceServer will
// result in compilation errors.
type UnsafeDoctorServiceServer interface {
	mustEmbedUnimplementedDoctorServiceServer()
}

func RegisterDoctorServiceServer(s grpc.ServiceRegistrar, srv DoctorServiceServer) {
	s.RegisterService(&DoctorService_ServiceDesc, srv)
}

func _DoctorService_GetDoctor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoctorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).GetDoctor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_GetDoctor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).GetDoctor(ctx, req.(*DoctorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_GetDoctorsIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoctorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).GetDoctorsIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_GetDoctorsIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).GetDoctorsIds(ctx, req.(*DoctorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DoctorService_ServiceDesc is the grpc.ServiceDesc for DoctorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DoctorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "doctors.DoctorService",
	HandlerType: (*DoctorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDoctor",
			Handler:    _DoctorService_GetDoctor_Handler,
		},
		{
			MethodName: "GetDoctorsIds",
			Handler:    _DoctorService_GetDoctorsIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "doctors_service.proto",
}
