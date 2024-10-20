// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: proto/recommender.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RecommendService_AddEvent_FullMethodName               = "/proto.RecommendService/AddEvent"
	RecommendService_RemoveEvent_FullMethodName            = "/proto.RecommendService/RemoveEvent"
	RecommendService_GetFoodRecommendations_FullMethodName = "/proto.RecommendService/GetFoodRecommendations"
)

// RecommendServiceClient is the client API for RecommendService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecommendServiceClient interface {
	AddEvent(ctx context.Context, in *AddEventReq, opts ...grpc.CallOption) (*Empty, error)
	RemoveEvent(ctx context.Context, in *RemoveEventReq, opts ...grpc.CallOption) (*Empty, error)
	GetFoodRecommendations(ctx context.Context, in *GetRecommendationsRequest, opts ...grpc.CallOption) (*GetRecommendationsResponse, error)
}

type recommendServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRecommendServiceClient(cc grpc.ClientConnInterface) RecommendServiceClient {
	return &recommendServiceClient{cc}
}

func (c *recommendServiceClient) AddEvent(ctx context.Context, in *AddEventReq, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, RecommendService_AddEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommendServiceClient) RemoveEvent(ctx context.Context, in *RemoveEventReq, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, RecommendService_RemoveEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommendServiceClient) GetFoodRecommendations(ctx context.Context, in *GetRecommendationsRequest, opts ...grpc.CallOption) (*GetRecommendationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRecommendationsResponse)
	err := c.cc.Invoke(ctx, RecommendService_GetFoodRecommendations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecommendServiceServer is the server API for RecommendService service.
// All implementations must embed UnimplementedRecommendServiceServer
// for forward compatibility.
type RecommendServiceServer interface {
	AddEvent(context.Context, *AddEventReq) (*Empty, error)
	RemoveEvent(context.Context, *RemoveEventReq) (*Empty, error)
	GetFoodRecommendations(context.Context, *GetRecommendationsRequest) (*GetRecommendationsResponse, error)
	mustEmbedUnimplementedRecommendServiceServer()
}

// UnimplementedRecommendServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRecommendServiceServer struct{}

func (UnimplementedRecommendServiceServer) AddEvent(context.Context, *AddEventReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEvent not implemented")
}
func (UnimplementedRecommendServiceServer) RemoveEvent(context.Context, *RemoveEventReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveEvent not implemented")
}
func (UnimplementedRecommendServiceServer) GetFoodRecommendations(context.Context, *GetRecommendationsRequest) (*GetRecommendationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFoodRecommendations not implemented")
}
func (UnimplementedRecommendServiceServer) mustEmbedUnimplementedRecommendServiceServer() {}
func (UnimplementedRecommendServiceServer) testEmbeddedByValue()                          {}

// UnsafeRecommendServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecommendServiceServer will
// result in compilation errors.
type UnsafeRecommendServiceServer interface {
	mustEmbedUnimplementedRecommendServiceServer()
}

func RegisterRecommendServiceServer(s grpc.ServiceRegistrar, srv RecommendServiceServer) {
	// If the following call pancis, it indicates UnimplementedRecommendServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RecommendService_ServiceDesc, srv)
}

func _RecommendService_AddEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddEventReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendServiceServer).AddEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecommendService_AddEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendServiceServer).AddEvent(ctx, req.(*AddEventReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecommendService_RemoveEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveEventReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendServiceServer).RemoveEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecommendService_RemoveEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendServiceServer).RemoveEvent(ctx, req.(*RemoveEventReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecommendService_GetFoodRecommendations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecommendationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendServiceServer).GetFoodRecommendations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecommendService_GetFoodRecommendations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendServiceServer).GetFoodRecommendations(ctx, req.(*GetRecommendationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RecommendService_ServiceDesc is the grpc.ServiceDesc for RecommendService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RecommendService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.RecommendService",
	HandlerType: (*RecommendServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddEvent",
			Handler:    _RecommendService_AddEvent_Handler,
		},
		{
			MethodName: "RemoveEvent",
			Handler:    _RecommendService_RemoveEvent_Handler,
		},
		{
			MethodName: "GetFoodRecommendations",
			Handler:    _RecommendService_GetFoodRecommendations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/recommender.proto",
}