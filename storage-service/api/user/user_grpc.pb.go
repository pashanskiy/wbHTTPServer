// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: user/user.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	core "wbHTTPServer/storage-service/api/core"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	UserStorageService_GetUser_FullMethodName    = "/user_service.UserStorageService/GetUser"
	UserStorageService_CreateUser_FullMethodName = "/user_service.UserStorageService/CreateUser"
	UserStorageService_UpdateUser_FullMethodName = "/user_service.UserStorageService/UpdateUser"
	UserStorageService_DeleteUser_FullMethodName = "/user_service.UserStorageService/DeleteUser"
)

// UserStorageServiceClient is the client API for UserStorageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserStorageServiceClient interface {
	GetUser(ctx context.Context, in *GetUserInfo, opts ...grpc.CallOption) (*UserInfo, error)
	CreateUser(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*core.EmptyMessage, error)
	UpdateUser(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*core.EmptyMessage, error)
	DeleteUser(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*core.EmptyMessage, error)
}

type userStorageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserStorageServiceClient(cc grpc.ClientConnInterface) UserStorageServiceClient {
	return &userStorageServiceClient{cc}
}

func (c *userStorageServiceClient) GetUser(ctx context.Context, in *GetUserInfo, opts ...grpc.CallOption) (*UserInfo, error) {
	out := new(UserInfo)
	err := c.cc.Invoke(ctx, UserStorageService_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userStorageServiceClient) CreateUser(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*core.EmptyMessage, error) {
	out := new(core.EmptyMessage)
	err := c.cc.Invoke(ctx, UserStorageService_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userStorageServiceClient) UpdateUser(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*core.EmptyMessage, error) {
	out := new(core.EmptyMessage)
	err := c.cc.Invoke(ctx, UserStorageService_UpdateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userStorageServiceClient) DeleteUser(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*core.EmptyMessage, error) {
	out := new(core.EmptyMessage)
	err := c.cc.Invoke(ctx, UserStorageService_DeleteUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserStorageServiceServer is the server API for UserStorageService service.
// All implementations must embed UnimplementedUserStorageServiceServer
// for forward compatibility
type UserStorageServiceServer interface {
	GetUser(context.Context, *GetUserInfo) (*UserInfo, error)
	CreateUser(context.Context, *UserInfo) (*core.EmptyMessage, error)
	UpdateUser(context.Context, *UserInfo) (*core.EmptyMessage, error)
	DeleteUser(context.Context, *UserInfo) (*core.EmptyMessage, error)
	mustEmbedUnimplementedUserStorageServiceServer()
}

// UnimplementedUserStorageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserStorageServiceServer struct {
}

func (UnimplementedUserStorageServiceServer) GetUser(context.Context, *GetUserInfo) (*UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserStorageServiceServer) CreateUser(context.Context, *UserInfo) (*core.EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserStorageServiceServer) UpdateUser(context.Context, *UserInfo) (*core.EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserStorageServiceServer) DeleteUser(context.Context, *UserInfo) (*core.EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserStorageServiceServer) mustEmbedUnimplementedUserStorageServiceServer() {}

// UnsafeUserStorageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserStorageServiceServer will
// result in compilation errors.
type UnsafeUserStorageServiceServer interface {
	mustEmbedUnimplementedUserStorageServiceServer()
}

func RegisterUserStorageServiceServer(s grpc.ServiceRegistrar, srv UserStorageServiceServer) {
	s.RegisterService(&UserStorageService_ServiceDesc, srv)
}

func _UserStorageService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserStorageServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserStorageService_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserStorageServiceServer).GetUser(ctx, req.(*GetUserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserStorageService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserStorageServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserStorageService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserStorageServiceServer).CreateUser(ctx, req.(*UserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserStorageService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserStorageServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserStorageService_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserStorageServiceServer).UpdateUser(ctx, req.(*UserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserStorageService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserStorageServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserStorageService_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserStorageServiceServer).DeleteUser(ctx, req.(*UserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// UserStorageService_ServiceDesc is the grpc.ServiceDesc for UserStorageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserStorageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_service.UserStorageService",
	HandlerType: (*UserStorageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserStorageService_GetUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserStorageService_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserStorageService_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserStorageService_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}
