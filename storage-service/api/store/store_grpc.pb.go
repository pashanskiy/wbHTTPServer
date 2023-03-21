// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: store/store.proto

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
	StoreStorageService_GetStores_FullMethodName   = "/store_service.StoreStorageService/GetStores"
	StoreStorageService_CreateStore_FullMethodName = "/store_service.StoreStorageService/CreateStore"
	StoreStorageService_UpdateStore_FullMethodName = "/store_service.StoreStorageService/UpdateStore"
	StoreStorageService_DeleteStore_FullMethodName = "/store_service.StoreStorageService/DeleteStore"
)

// StoreStorageServiceClient is the client API for StoreStorageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StoreStorageServiceClient interface {
	GetStores(ctx context.Context, in *StoreInfo, opts ...grpc.CallOption) (*GetStoresResponse, error)
	CreateStore(ctx context.Context, in *CreateStoreRequest, opts ...grpc.CallOption) (*core.EmptyMessage, error)
	UpdateStore(ctx context.Context, in *UpdateStoreRequest, opts ...grpc.CallOption) (*core.EmptyMessage, error)
	DeleteStore(ctx context.Context, in *DeleteStoreRequest, opts ...grpc.CallOption) (*core.EmptyMessage, error)
}

type storeStorageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStoreStorageServiceClient(cc grpc.ClientConnInterface) StoreStorageServiceClient {
	return &storeStorageServiceClient{cc}
}

func (c *storeStorageServiceClient) GetStores(ctx context.Context, in *StoreInfo, opts ...grpc.CallOption) (*GetStoresResponse, error) {
	out := new(GetStoresResponse)
	err := c.cc.Invoke(ctx, StoreStorageService_GetStores_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeStorageServiceClient) CreateStore(ctx context.Context, in *CreateStoreRequest, opts ...grpc.CallOption) (*core.EmptyMessage, error) {
	out := new(core.EmptyMessage)
	err := c.cc.Invoke(ctx, StoreStorageService_CreateStore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeStorageServiceClient) UpdateStore(ctx context.Context, in *UpdateStoreRequest, opts ...grpc.CallOption) (*core.EmptyMessage, error) {
	out := new(core.EmptyMessage)
	err := c.cc.Invoke(ctx, StoreStorageService_UpdateStore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeStorageServiceClient) DeleteStore(ctx context.Context, in *DeleteStoreRequest, opts ...grpc.CallOption) (*core.EmptyMessage, error) {
	out := new(core.EmptyMessage)
	err := c.cc.Invoke(ctx, StoreStorageService_DeleteStore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreStorageServiceServer is the server API for StoreStorageService service.
// All implementations must embed UnimplementedStoreStorageServiceServer
// for forward compatibility
type StoreStorageServiceServer interface {
	GetStores(context.Context, *StoreInfo) (*GetStoresResponse, error)
	CreateStore(context.Context, *CreateStoreRequest) (*core.EmptyMessage, error)
	UpdateStore(context.Context, *UpdateStoreRequest) (*core.EmptyMessage, error)
	DeleteStore(context.Context, *DeleteStoreRequest) (*core.EmptyMessage, error)
	mustEmbedUnimplementedStoreStorageServiceServer()
}

// UnimplementedStoreStorageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStoreStorageServiceServer struct {
}

func (UnimplementedStoreStorageServiceServer) GetStores(context.Context, *StoreInfo) (*GetStoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStores not implemented")
}
func (UnimplementedStoreStorageServiceServer) CreateStore(context.Context, *CreateStoreRequest) (*core.EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStore not implemented")
}
func (UnimplementedStoreStorageServiceServer) UpdateStore(context.Context, *UpdateStoreRequest) (*core.EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStore not implemented")
}
func (UnimplementedStoreStorageServiceServer) DeleteStore(context.Context, *DeleteStoreRequest) (*core.EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStore not implemented")
}
func (UnimplementedStoreStorageServiceServer) mustEmbedUnimplementedStoreStorageServiceServer() {}

// UnsafeStoreStorageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StoreStorageServiceServer will
// result in compilation errors.
type UnsafeStoreStorageServiceServer interface {
	mustEmbedUnimplementedStoreStorageServiceServer()
}

func RegisterStoreStorageServiceServer(s grpc.ServiceRegistrar, srv StoreStorageServiceServer) {
	s.RegisterService(&StoreStorageService_ServiceDesc, srv)
}

func _StoreStorageService_GetStores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreStorageServiceServer).GetStores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StoreStorageService_GetStores_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreStorageServiceServer).GetStores(ctx, req.(*StoreInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreStorageService_CreateStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreStorageServiceServer).CreateStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StoreStorageService_CreateStore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreStorageServiceServer).CreateStore(ctx, req.(*CreateStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreStorageService_UpdateStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreStorageServiceServer).UpdateStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StoreStorageService_UpdateStore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreStorageServiceServer).UpdateStore(ctx, req.(*UpdateStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreStorageService_DeleteStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreStorageServiceServer).DeleteStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StoreStorageService_DeleteStore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreStorageServiceServer).DeleteStore(ctx, req.(*DeleteStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StoreStorageService_ServiceDesc is the grpc.ServiceDesc for StoreStorageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StoreStorageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "store_service.StoreStorageService",
	HandlerType: (*StoreStorageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStores",
			Handler:    _StoreStorageService_GetStores_Handler,
		},
		{
			MethodName: "CreateStore",
			Handler:    _StoreStorageService_CreateStore_Handler,
		},
		{
			MethodName: "UpdateStore",
			Handler:    _StoreStorageService_UpdateStore_Handler,
		},
		{
			MethodName: "DeleteStore",
			Handler:    _StoreStorageService_DeleteStore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "store/store.proto",
}
