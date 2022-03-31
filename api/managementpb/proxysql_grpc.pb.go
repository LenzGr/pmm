// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: managementpb/proxysql.proto

package managementpb

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

// ProxySQLClient is the client API for ProxySQL service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProxySQLClient interface {
	// AddProxySQL adds ProxySQL Service and starts several Agents.
	// It automatically adds a service to inventory, which is running on provided "node_id",
	// then adds "proxysql_exporter" with provided "pmm_agent_id" and other parameters.
	AddProxySQL(ctx context.Context, in *AddProxySQLRequest, opts ...grpc.CallOption) (*AddProxySQLResponse, error)
}

type proxySQLClient struct {
	cc grpc.ClientConnInterface
}

func NewProxySQLClient(cc grpc.ClientConnInterface) ProxySQLClient {
	return &proxySQLClient{cc}
}

func (c *proxySQLClient) AddProxySQL(ctx context.Context, in *AddProxySQLRequest, opts ...grpc.CallOption) (*AddProxySQLResponse, error) {
	out := new(AddProxySQLResponse)
	err := c.cc.Invoke(ctx, "/management.ProxySQL/AddProxySQL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProxySQLServer is the server API for ProxySQL service.
// All implementations must embed UnimplementedProxySQLServer
// for forward compatibility
type ProxySQLServer interface {
	// AddProxySQL adds ProxySQL Service and starts several Agents.
	// It automatically adds a service to inventory, which is running on provided "node_id",
	// then adds "proxysql_exporter" with provided "pmm_agent_id" and other parameters.
	AddProxySQL(context.Context, *AddProxySQLRequest) (*AddProxySQLResponse, error)
	mustEmbedUnimplementedProxySQLServer()
}

// UnimplementedProxySQLServer must be embedded to have forward compatible implementations.
type UnimplementedProxySQLServer struct {
}

func (UnimplementedProxySQLServer) AddProxySQL(context.Context, *AddProxySQLRequest) (*AddProxySQLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProxySQL not implemented")
}
func (UnimplementedProxySQLServer) mustEmbedUnimplementedProxySQLServer() {}

// UnsafeProxySQLServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProxySQLServer will
// result in compilation errors.
type UnsafeProxySQLServer interface {
	mustEmbedUnimplementedProxySQLServer()
}

func RegisterProxySQLServer(s grpc.ServiceRegistrar, srv ProxySQLServer) {
	s.RegisterService(&ProxySQL_ServiceDesc, srv)
}

func _ProxySQL_AddProxySQL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProxySQLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxySQLServer).AddProxySQL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.ProxySQL/AddProxySQL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxySQLServer).AddProxySQL(ctx, req.(*AddProxySQLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProxySQL_ServiceDesc is the grpc.ServiceDesc for ProxySQL service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProxySQL_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "management.ProxySQL",
	HandlerType: (*ProxySQLServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddProxySQL",
			Handler:    _ProxySQL_AddProxySQL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/proxysql.proto",
}