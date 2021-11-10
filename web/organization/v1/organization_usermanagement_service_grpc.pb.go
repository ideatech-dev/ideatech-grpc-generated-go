// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpcweb_organization_v1

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

// OrganizationUserManagementServiceClient is the client API for OrganizationUserManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrganizationUserManagementServiceClient interface {
	GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error)
	InviteUser(ctx context.Context, in *InviteUserRequest, opts ...grpc.CallOption) (*InviteUserResponse, error)
	RemoveUser(ctx context.Context, in *RemoveUserRequest, opts ...grpc.CallOption) (*RemoveUserResponse, error)
}

type organizationUserManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrganizationUserManagementServiceClient(cc grpc.ClientConnInterface) OrganizationUserManagementServiceClient {
	return &organizationUserManagementServiceClient{cc}
}

func (c *organizationUserManagementServiceClient) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	out := new(GetUsersResponse)
	err := c.cc.Invoke(ctx, "/web.organization.v1.OrganizationUserManagementService/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationUserManagementServiceClient) InviteUser(ctx context.Context, in *InviteUserRequest, opts ...grpc.CallOption) (*InviteUserResponse, error) {
	out := new(InviteUserResponse)
	err := c.cc.Invoke(ctx, "/web.organization.v1.OrganizationUserManagementService/InviteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationUserManagementServiceClient) RemoveUser(ctx context.Context, in *RemoveUserRequest, opts ...grpc.CallOption) (*RemoveUserResponse, error) {
	out := new(RemoveUserResponse)
	err := c.cc.Invoke(ctx, "/web.organization.v1.OrganizationUserManagementService/RemoveUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrganizationUserManagementServiceServer is the server API for OrganizationUserManagementService service.
// All implementations must embed UnimplementedOrganizationUserManagementServiceServer
// for forward compatibility
type OrganizationUserManagementServiceServer interface {
	GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error)
	InviteUser(context.Context, *InviteUserRequest) (*InviteUserResponse, error)
	RemoveUser(context.Context, *RemoveUserRequest) (*RemoveUserResponse, error)
	mustEmbedUnimplementedOrganizationUserManagementServiceServer()
}

// UnimplementedOrganizationUserManagementServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrganizationUserManagementServiceServer struct {
}

func (UnimplementedOrganizationUserManagementServiceServer) GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedOrganizationUserManagementServiceServer) InviteUser(context.Context, *InviteUserRequest) (*InviteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InviteUser not implemented")
}
func (UnimplementedOrganizationUserManagementServiceServer) RemoveUser(context.Context, *RemoveUserRequest) (*RemoveUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUser not implemented")
}
func (UnimplementedOrganizationUserManagementServiceServer) mustEmbedUnimplementedOrganizationUserManagementServiceServer() {
}

// UnsafeOrganizationUserManagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrganizationUserManagementServiceServer will
// result in compilation errors.
type UnsafeOrganizationUserManagementServiceServer interface {
	mustEmbedUnimplementedOrganizationUserManagementServiceServer()
}

func RegisterOrganizationUserManagementServiceServer(s grpc.ServiceRegistrar, srv OrganizationUserManagementServiceServer) {
	s.RegisterService(&OrganizationUserManagementService_ServiceDesc, srv)
}

func _OrganizationUserManagementService_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationUserManagementServiceServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web.organization.v1.OrganizationUserManagementService/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationUserManagementServiceServer).GetUsers(ctx, req.(*GetUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationUserManagementService_InviteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationUserManagementServiceServer).InviteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web.organization.v1.OrganizationUserManagementService/InviteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationUserManagementServiceServer).InviteUser(ctx, req.(*InviteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationUserManagementService_RemoveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationUserManagementServiceServer).RemoveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web.organization.v1.OrganizationUserManagementService/RemoveUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationUserManagementServiceServer).RemoveUser(ctx, req.(*RemoveUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrganizationUserManagementService_ServiceDesc is the grpc.ServiceDesc for OrganizationUserManagementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrganizationUserManagementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "web.organization.v1.OrganizationUserManagementService",
	HandlerType: (*OrganizationUserManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUsers",
			Handler:    _OrganizationUserManagementService_GetUsers_Handler,
		},
		{
			MethodName: "InviteUser",
			Handler:    _OrganizationUserManagementService_InviteUser_Handler,
		},
		{
			MethodName: "RemoveUser",
			Handler:    _OrganizationUserManagementService_RemoveUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "web/organization/v1/organization_usermanagement_service.proto",
}
