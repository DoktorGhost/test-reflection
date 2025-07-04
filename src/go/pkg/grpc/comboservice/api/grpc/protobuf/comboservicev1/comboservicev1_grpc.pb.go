// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.29.3
// source: api/grpc/protobuf/comboservicev1/comboservicev1.proto

package comboservicev1

import (
	context "context"
	v1 "github.com/DoktorGhost/test-reflection/src/go/pkg/grpc/common/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ComboService_GetSelectionProducts_FullMethodName = "/combo.ComboService/GetSelectionProducts"
	ComboService_GetResources_FullMethodName         = "/combo.ComboService/GetResources"
)

// ComboServiceClient is the client API for ComboService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComboServiceClient interface {
	GetSelectionProducts(ctx context.Context, in *v1.SelectionRequest, opts ...grpc.CallOption) (*v1.SelectionResponse, error)
	GetResources(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetResourcesResponse, error)
}

type comboServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewComboServiceClient(cc grpc.ClientConnInterface) ComboServiceClient {
	return &comboServiceClient{cc}
}

func (c *comboServiceClient) GetSelectionProducts(ctx context.Context, in *v1.SelectionRequest, opts ...grpc.CallOption) (*v1.SelectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.SelectionResponse)
	err := c.cc.Invoke(ctx, ComboService_GetSelectionProducts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comboServiceClient) GetResources(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetResourcesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetResourcesResponse)
	err := c.cc.Invoke(ctx, ComboService_GetResources_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComboServiceServer is the server API for ComboService service.
// All implementations must embed UnimplementedComboServiceServer
// for forward compatibility
type ComboServiceServer interface {
	GetSelectionProducts(context.Context, *v1.SelectionRequest) (*v1.SelectionResponse, error)
	GetResources(context.Context, *emptypb.Empty) (*GetResourcesResponse, error)
	mustEmbedUnimplementedComboServiceServer()
}

// UnimplementedComboServiceServer must be embedded to have forward compatible implementations.
type UnimplementedComboServiceServer struct {
}

func (UnimplementedComboServiceServer) GetSelectionProducts(context.Context, *v1.SelectionRequest) (*v1.SelectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSelectionProducts not implemented")
}
func (UnimplementedComboServiceServer) GetResources(context.Context, *emptypb.Empty) (*GetResourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResources not implemented")
}
func (UnimplementedComboServiceServer) mustEmbedUnimplementedComboServiceServer() {}

// UnsafeComboServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComboServiceServer will
// result in compilation errors.
type UnsafeComboServiceServer interface {
	mustEmbedUnimplementedComboServiceServer()
}

func RegisterComboServiceServer(s grpc.ServiceRegistrar, srv ComboServiceServer) {
	s.RegisterService(&ComboService_ServiceDesc, srv)
}

func _ComboService_GetSelectionProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.SelectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComboServiceServer).GetSelectionProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComboService_GetSelectionProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComboServiceServer).GetSelectionProducts(ctx, req.(*v1.SelectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComboService_GetResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComboServiceServer).GetResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComboService_GetResources_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComboServiceServer).GetResources(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ComboService_ServiceDesc is the grpc.ServiceDesc for ComboService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ComboService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "combo.ComboService",
	HandlerType: (*ComboServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSelectionProducts",
			Handler:    _ComboService_GetSelectionProducts_Handler,
		},
		{
			MethodName: "GetResources",
			Handler:    _ComboService_GetResources_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/grpc/protobuf/comboservicev1/comboservicev1.proto",
}

const (
	ComboCatalogService_GetProducts_FullMethodName = "/combo.ComboCatalogService/GetProducts"
)

// ComboCatalogServiceClient is the client API for ComboCatalogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComboCatalogServiceClient interface {
	GetProducts(ctx context.Context, in *v1.DiscountRequest, opts ...grpc.CallOption) (*v1.DiscountResponse, error)
}

type comboCatalogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewComboCatalogServiceClient(cc grpc.ClientConnInterface) ComboCatalogServiceClient {
	return &comboCatalogServiceClient{cc}
}

func (c *comboCatalogServiceClient) GetProducts(ctx context.Context, in *v1.DiscountRequest, opts ...grpc.CallOption) (*v1.DiscountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.DiscountResponse)
	err := c.cc.Invoke(ctx, ComboCatalogService_GetProducts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComboCatalogServiceServer is the server API for ComboCatalogService service.
// All implementations must embed UnimplementedComboCatalogServiceServer
// for forward compatibility
type ComboCatalogServiceServer interface {
	GetProducts(context.Context, *v1.DiscountRequest) (*v1.DiscountResponse, error)
	mustEmbedUnimplementedComboCatalogServiceServer()
}

// UnimplementedComboCatalogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedComboCatalogServiceServer struct {
}

func (UnimplementedComboCatalogServiceServer) GetProducts(context.Context, *v1.DiscountRequest) (*v1.DiscountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProducts not implemented")
}
func (UnimplementedComboCatalogServiceServer) mustEmbedUnimplementedComboCatalogServiceServer() {}

// UnsafeComboCatalogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComboCatalogServiceServer will
// result in compilation errors.
type UnsafeComboCatalogServiceServer interface {
	mustEmbedUnimplementedComboCatalogServiceServer()
}

func RegisterComboCatalogServiceServer(s grpc.ServiceRegistrar, srv ComboCatalogServiceServer) {
	s.RegisterService(&ComboCatalogService_ServiceDesc, srv)
}

func _ComboCatalogService_GetProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.DiscountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComboCatalogServiceServer).GetProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComboCatalogService_GetProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComboCatalogServiceServer).GetProducts(ctx, req.(*v1.DiscountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ComboCatalogService_ServiceDesc is the grpc.ServiceDesc for ComboCatalogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ComboCatalogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "combo.ComboCatalogService",
	HandlerType: (*ComboCatalogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProducts",
			Handler:    _ComboCatalogService_GetProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/grpc/protobuf/comboservicev1/comboservicev1.proto",
}

const (
	ResourcesService_GetResources_FullMethodName = "/combo.ResourcesService/GetResources"
)

// ResourcesServiceClient is the client API for ResourcesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResourcesServiceClient interface {
	GetResources(ctx context.Context, in *v1.ResourcesRequest, opts ...grpc.CallOption) (*v1.ResourcesResponse, error)
}

type resourcesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResourcesServiceClient(cc grpc.ClientConnInterface) ResourcesServiceClient {
	return &resourcesServiceClient{cc}
}

func (c *resourcesServiceClient) GetResources(ctx context.Context, in *v1.ResourcesRequest, opts ...grpc.CallOption) (*v1.ResourcesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.ResourcesResponse)
	err := c.cc.Invoke(ctx, ResourcesService_GetResources_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourcesServiceServer is the server API for ResourcesService service.
// All implementations must embed UnimplementedResourcesServiceServer
// for forward compatibility
type ResourcesServiceServer interface {
	GetResources(context.Context, *v1.ResourcesRequest) (*v1.ResourcesResponse, error)
	mustEmbedUnimplementedResourcesServiceServer()
}

// UnimplementedResourcesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedResourcesServiceServer struct {
}

func (UnimplementedResourcesServiceServer) GetResources(context.Context, *v1.ResourcesRequest) (*v1.ResourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResources not implemented")
}
func (UnimplementedResourcesServiceServer) mustEmbedUnimplementedResourcesServiceServer() {}

// UnsafeResourcesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourcesServiceServer will
// result in compilation errors.
type UnsafeResourcesServiceServer interface {
	mustEmbedUnimplementedResourcesServiceServer()
}

func RegisterResourcesServiceServer(s grpc.ServiceRegistrar, srv ResourcesServiceServer) {
	s.RegisterService(&ResourcesService_ServiceDesc, srv)
}

func _ResourcesService_GetResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.ResourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServiceServer).GetResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourcesService_GetResources_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServiceServer).GetResources(ctx, req.(*v1.ResourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ResourcesService_ServiceDesc is the grpc.ServiceDesc for ResourcesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResourcesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "combo.ResourcesService",
	HandlerType: (*ResourcesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetResources",
			Handler:    _ResourcesService_GetResources_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/grpc/protobuf/comboservicev1/comboservicev1.proto",
}
