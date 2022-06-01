// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: managementpb/backup/artifacts.proto

package backupv1beta1

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

// ArtifactsClient is the client API for Artifacts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArtifactsClient interface {
	// ListArtifacts returns a list of all backup artifacts.
	ListArtifacts(ctx context.Context, in *ListArtifactsRequest, opts ...grpc.CallOption) (*ListArtifactsResponse, error)
	// DeleteArtifact deletes specified artifact.
	DeleteArtifact(ctx context.Context, in *DeleteArtifactRequest, opts ...grpc.CallOption) (*DeleteArtifactResponse, error)
}

type artifactsClient struct {
	cc grpc.ClientConnInterface
}

func NewArtifactsClient(cc grpc.ClientConnInterface) ArtifactsClient {
	return &artifactsClient{cc}
}

func (c *artifactsClient) ListArtifacts(ctx context.Context, in *ListArtifactsRequest, opts ...grpc.CallOption) (*ListArtifactsResponse, error) {
	out := new(ListArtifactsResponse)
	err := c.cc.Invoke(ctx, "/backup.v1beta1.Artifacts/ListArtifacts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactsClient) DeleteArtifact(ctx context.Context, in *DeleteArtifactRequest, opts ...grpc.CallOption) (*DeleteArtifactResponse, error) {
	out := new(DeleteArtifactResponse)
	err := c.cc.Invoke(ctx, "/backup.v1beta1.Artifacts/DeleteArtifact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArtifactsServer is the server API for Artifacts service.
// All implementations must embed UnimplementedArtifactsServer
// for forward compatibility
type ArtifactsServer interface {
	// ListArtifacts returns a list of all backup artifacts.
	ListArtifacts(context.Context, *ListArtifactsRequest) (*ListArtifactsResponse, error)
	// DeleteArtifact deletes specified artifact.
	DeleteArtifact(context.Context, *DeleteArtifactRequest) (*DeleteArtifactResponse, error)
	mustEmbedUnimplementedArtifactsServer()
}

// UnimplementedArtifactsServer must be embedded to have forward compatible implementations.
type UnimplementedArtifactsServer struct{}

func (UnimplementedArtifactsServer) ListArtifacts(context.Context, *ListArtifactsRequest) (*ListArtifactsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArtifacts not implemented")
}

func (UnimplementedArtifactsServer) DeleteArtifact(context.Context, *DeleteArtifactRequest) (*DeleteArtifactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArtifact not implemented")
}
func (UnimplementedArtifactsServer) mustEmbedUnimplementedArtifactsServer() {}

// UnsafeArtifactsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArtifactsServer will
// result in compilation errors.
type UnsafeArtifactsServer interface {
	mustEmbedUnimplementedArtifactsServer()
}

func RegisterArtifactsServer(s grpc.ServiceRegistrar, srv ArtifactsServer) {
	s.RegisterService(&Artifacts_ServiceDesc, srv)
}

func _Artifacts_ListArtifacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListArtifactsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactsServer).ListArtifacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backup.v1beta1.Artifacts/ListArtifacts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactsServer).ListArtifacts(ctx, req.(*ListArtifactsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Artifacts_DeleteArtifact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArtifactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactsServer).DeleteArtifact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backup.v1beta1.Artifacts/DeleteArtifact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactsServer).DeleteArtifact(ctx, req.(*DeleteArtifactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Artifacts_ServiceDesc is the grpc.ServiceDesc for Artifacts service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Artifacts_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "backup.v1beta1.Artifacts",
	HandlerType: (*ArtifactsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListArtifacts",
			Handler:    _Artifacts_ListArtifacts_Handler,
		},
		{
			MethodName: "DeleteArtifact",
			Handler:    _Artifacts_DeleteArtifact_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/backup/artifacts.proto",
}
