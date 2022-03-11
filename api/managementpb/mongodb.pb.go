// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: managementpb/mongodb.proto

package managementpb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/mwitkow/go-proto-validators"
	inventorypb "github.com/percona/pmm/api/inventorypb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AddMongoDBRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Node identifier on which a service is been running.
	// Exactly one of these parameters should be present: node_id, node_name, add_node.
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Node name on which a service is been running.
	// Exactly one of these parameters should be present: node_id, node_name, add_node.
	NodeName string `protobuf:"bytes,2,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	// Create a new Node with those parameters.
	// Exactly one of these parameters should be present: node_id, node_name, add_node.
	AddNode *AddNodeParams `protobuf:"bytes,3,opt,name=add_node,json=addNode,proto3" json:"add_node,omitempty"`
	// Unique across all Services user-defined name. Required.
	ServiceName string `protobuf:"bytes,4,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// Node and Service access address (DNS name or IP).
	// Address (and port) or socket is required.
	Address string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	// Service Access port.
	// Port is required when the address present.
	Port uint32 `protobuf:"varint,6,opt,name=port,proto3" json:"port,omitempty"`
	// Service Access socket.
	// Address (and port) or socket is required.
	Socket string `protobuf:"bytes,19,opt,name=socket,proto3" json:"socket,omitempty"`
	// The "pmm-agent" identifier which should run agents. Required.
	PmmAgentId string `protobuf:"bytes,7,opt,name=pmm_agent_id,json=pmmAgentId,proto3" json:"pmm_agent_id,omitempty"`
	// Environment name.
	Environment string `protobuf:"bytes,8,opt,name=environment,proto3" json:"environment,omitempty"`
	// Cluster name.
	Cluster string `protobuf:"bytes,9,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// Replication set name.
	ReplicationSet string `protobuf:"bytes,10,opt,name=replication_set,json=replicationSet,proto3" json:"replication_set,omitempty"`
	// MongoDB username for exporter and QAN agent access.
	Username string `protobuf:"bytes,11,opt,name=username,proto3" json:"username,omitempty"`
	// MongoDB password for exporter and QAN agent access.
	Password string `protobuf:"bytes,12,opt,name=password,proto3" json:"password,omitempty"`
	// If true, adds qan-mongodb-profiler-agent for provided service.
	QanMongodbProfiler bool `protobuf:"varint,13,opt,name=qan_mongodb_profiler,json=qanMongodbProfiler,proto3" json:"qan_mongodb_profiler,omitempty"`
	// Custom user-assigned labels for Service.
	CustomLabels map[string]string `protobuf:"bytes,14,rep,name=custom_labels,json=customLabels,proto3" json:"custom_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Skip connection check.
	SkipConnectionCheck bool `protobuf:"varint,15,opt,name=skip_connection_check,json=skipConnectionCheck,proto3" json:"skip_connection_check,omitempty"`
	// Use TLS for database connections.
	Tls bool `protobuf:"varint,17,opt,name=tls,proto3" json:"tls,omitempty"`
	// Skip TLS certificate and hostname validation.
	TlsSkipVerify bool `protobuf:"varint,18,opt,name=tls_skip_verify,json=tlsSkipVerify,proto3" json:"tls_skip_verify,omitempty"`
	// Client certificate and key.
	TlsCertificateKey string `protobuf:"bytes,21,opt,name=tls_certificate_key,json=tlsCertificateKey,proto3" json:"tls_certificate_key,omitempty"`
	// Password for decrypting tls_certificate_key.
	TlsCertificateKeyFilePassword string `protobuf:"bytes,22,opt,name=tls_certificate_key_file_password,json=tlsCertificateKeyFilePassword,proto3" json:"tls_certificate_key_file_password,omitempty"`
	// Certificate Authority certificate chain.
	TlsCa string `protobuf:"bytes,23,opt,name=tls_ca,json=tlsCa,proto3" json:"tls_ca,omitempty"`
	// Defines metrics flow model for this exporter.
	// Metrics could be pushed to the server with vmagent,
	// pulled by the server, or the server could choose behavior automatically.
	MetricsMode MetricsMode `protobuf:"varint,20,opt,name=metrics_mode,json=metricsMode,proto3,enum=management.MetricsMode" json:"metrics_mode,omitempty"`
	// List of collector names to disable in this exporter.
	DisableCollectors []string `protobuf:"bytes,24,rep,name=disable_collectors,json=disableCollectors,proto3" json:"disable_collectors,omitempty"`
	// Authentication mechanism.
	// See https://docs.mongodb.com/manual/reference/connection-string/#mongodb-urioption-urioption.authMechanism
	// for details.
	AuthenticationMechanism string `protobuf:"bytes,25,opt,name=authentication_mechanism,json=authenticationMechanism,proto3" json:"authentication_mechanism,omitempty"`
	// Authentication database.
	AuthenticationDatabase string `protobuf:"bytes,26,opt,name=authentication_database,json=authenticationDatabase,proto3" json:"authentication_database,omitempty"`
	// Custom password for exporter endpoint /metrics.
	AgentPassword string `protobuf:"bytes,27,opt,name=agent_password,json=agentPassword,proto3" json:"agent_password,omitempty"`
	// List of collections to get stats from. Can use * .
	StatsCollections []string `protobuf:"bytes,28,rep,name=stats_collections,json=statsCollections,proto3" json:"stats_collections,omitempty"`
	// Collections limit. Only get Databases and collection stats if the total number of collections in the server
	// is less than this value. 0: no limit
	CollectionsLimit int32 `protobuf:"varint,29,opt,name=collections_limit,json=collectionsLimit,proto3" json:"collections_limit,omitempty"`
	// Enable all collectors
	EnableAllCollectors bool `protobuf:"varint,30,opt,name=enable_all_collectors,json=enableAllCollectors,proto3" json:"enable_all_collectors,omitempty"`
}

func (x *AddMongoDBRequest) Reset() {
	*x = AddMongoDBRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_mongodb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMongoDBRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMongoDBRequest) ProtoMessage() {}

func (x *AddMongoDBRequest) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_mongodb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMongoDBRequest.ProtoReflect.Descriptor instead.
func (*AddMongoDBRequest) Descriptor() ([]byte, []int) {
	return file_managementpb_mongodb_proto_rawDescGZIP(), []int{0}
}

func (x *AddMongoDBRequest) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *AddMongoDBRequest) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *AddMongoDBRequest) GetAddNode() *AddNodeParams {
	if x != nil {
		return x.AddNode
	}
	return nil
}

func (x *AddMongoDBRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *AddMongoDBRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AddMongoDBRequest) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *AddMongoDBRequest) GetSocket() string {
	if x != nil {
		return x.Socket
	}
	return ""
}

func (x *AddMongoDBRequest) GetPmmAgentId() string {
	if x != nil {
		return x.PmmAgentId
	}
	return ""
}

func (x *AddMongoDBRequest) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *AddMongoDBRequest) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *AddMongoDBRequest) GetReplicationSet() string {
	if x != nil {
		return x.ReplicationSet
	}
	return ""
}

func (x *AddMongoDBRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AddMongoDBRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AddMongoDBRequest) GetQanMongodbProfiler() bool {
	if x != nil {
		return x.QanMongodbProfiler
	}
	return false
}

func (x *AddMongoDBRequest) GetCustomLabels() map[string]string {
	if x != nil {
		return x.CustomLabels
	}
	return nil
}

func (x *AddMongoDBRequest) GetSkipConnectionCheck() bool {
	if x != nil {
		return x.SkipConnectionCheck
	}
	return false
}

func (x *AddMongoDBRequest) GetTls() bool {
	if x != nil {
		return x.Tls
	}
	return false
}

func (x *AddMongoDBRequest) GetTlsSkipVerify() bool {
	if x != nil {
		return x.TlsSkipVerify
	}
	return false
}

func (x *AddMongoDBRequest) GetTlsCertificateKey() string {
	if x != nil {
		return x.TlsCertificateKey
	}
	return ""
}

func (x *AddMongoDBRequest) GetTlsCertificateKeyFilePassword() string {
	if x != nil {
		return x.TlsCertificateKeyFilePassword
	}
	return ""
}

func (x *AddMongoDBRequest) GetTlsCa() string {
	if x != nil {
		return x.TlsCa
	}
	return ""
}

func (x *AddMongoDBRequest) GetMetricsMode() MetricsMode {
	if x != nil {
		return x.MetricsMode
	}
	return MetricsMode_AUTO
}

func (x *AddMongoDBRequest) GetDisableCollectors() []string {
	if x != nil {
		return x.DisableCollectors
	}
	return nil
}

func (x *AddMongoDBRequest) GetAuthenticationMechanism() string {
	if x != nil {
		return x.AuthenticationMechanism
	}
	return ""
}

func (x *AddMongoDBRequest) GetAuthenticationDatabase() string {
	if x != nil {
		return x.AuthenticationDatabase
	}
	return ""
}

func (x *AddMongoDBRequest) GetAgentPassword() string {
	if x != nil {
		return x.AgentPassword
	}
	return ""
}

func (x *AddMongoDBRequest) GetStatsCollections() []string {
	if x != nil {
		return x.StatsCollections
	}
	return nil
}

func (x *AddMongoDBRequest) GetCollectionsLimit() int32 {
	if x != nil {
		return x.CollectionsLimit
	}
	return 0
}

func (x *AddMongoDBRequest) GetEnableAllCollectors() bool {
	if x != nil {
		return x.EnableAllCollectors
	}
	return false
}

type AddMongoDBResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service            *inventorypb.MongoDBService          `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	MongodbExporter    *inventorypb.MongoDBExporter         `protobuf:"bytes,2,opt,name=mongodb_exporter,json=mongodbExporter,proto3" json:"mongodb_exporter,omitempty"`
	QanMongodbProfiler *inventorypb.QANMongoDBProfilerAgent `protobuf:"bytes,3,opt,name=qan_mongodb_profiler,json=qanMongodbProfiler,proto3" json:"qan_mongodb_profiler,omitempty"`
}

func (x *AddMongoDBResponse) Reset() {
	*x = AddMongoDBResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_mongodb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMongoDBResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMongoDBResponse) ProtoMessage() {}

func (x *AddMongoDBResponse) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_mongodb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMongoDBResponse.ProtoReflect.Descriptor instead.
func (*AddMongoDBResponse) Descriptor() ([]byte, []int) {
	return file_managementpb_mongodb_proto_rawDescGZIP(), []int{1}
}

func (x *AddMongoDBResponse) GetService() *inventorypb.MongoDBService {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *AddMongoDBResponse) GetMongodbExporter() *inventorypb.MongoDBExporter {
	if x != nil {
		return x.MongodbExporter
	}
	return nil
}

func (x *AddMongoDBResponse) GetQanMongodbProfiler() *inventorypb.QANMongoDBProfilerAgent {
	if x != nil {
		return x.QanMongodbProfiler
	}
	return nil
}

var File_managementpb_mongodb_proto protoreflect.FileDescriptor

var file_managementpb_mongodb_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x6d,
	0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61, 0x67,
	0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f,
	0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x70, 0x62,
	0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x93, 0x0a, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x44, 0x42,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a,
	0x08, 0x61, 0x64, 0x64, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64,
	0x4e, 0x6f, 0x64, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x4e,
	0x6f, 0x64, 0x65, 0x12, 0x29, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58,
	0x01, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f,
	0x63, 0x6b, 0x65, 0x74, 0x12, 0x28, 0x0a, 0x0c, 0x70, 0x6d, 0x6d, 0x5f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02,
	0x58, 0x01, 0x52, 0x0a, 0x70, 0x6d, 0x6d, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x71,
	0x61, 0x6e, 0x5f, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x71, 0x61, 0x6e, 0x4d, 0x6f,
	0x6e, 0x67, 0x6f, 0x64, 0x62, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x12, 0x54, 0x0a,
	0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0e,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x41, 0x64, 0x64, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x44, 0x42, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x13, 0x73, 0x6b, 0x69, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6c, 0x73, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x74, 0x6c, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x6c, 0x73,
	0x5f, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x18, 0x12, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x74, 0x6c, 0x73, 0x53, 0x6b, 0x69, 0x70, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x12, 0x2e, 0x0a, 0x13, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x74, 0x6c, 0x73, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x4b, 0x65,
	0x79, 0x12, 0x48, 0x0a, 0x21, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1d, 0x74, 0x6c,
	0x73, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x46,
	0x69, 0x6c, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x74,
	0x6c, 0x73, 0x5f, 0x63, 0x61, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6c, 0x73,
	0x43, 0x61, 0x12, 0x3a, 0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x4d, 0x6f, 0x64,
	0x65, 0x52, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x2d,
	0x0a, 0x12, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x73, 0x18, 0x18, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x64, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x39, 0x0a,
	0x18, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x73, 0x6d, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x17, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x73, 0x6d, 0x12, 0x37, 0x0a, 0x17, 0x61, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x61, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x1c, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x10, 0x73, 0x74, 0x61, 0x74, 0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x10, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x32, 0x0a, 0x15, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x61, 0x6c, 0x6c,
	0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x13, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x1a, 0x3f, 0x0a, 0x11, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xe6, 0x01, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x4d,
	0x6f, 0x6e, 0x67, 0x6f, 0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33,
	0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x4d, 0x6f, 0x6e, 0x67,
	0x6f, 0x44, 0x42, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x10, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x5f, 0x65,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x44,
	0x42, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x52, 0x0f, 0x6d, 0x6f, 0x6e, 0x67, 0x6f,
	0x64, 0x62, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x12, 0x54, 0x0a, 0x14, 0x71, 0x61,
	0x6e, 0x5f, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x51, 0x41, 0x4e, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x44, 0x42, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x12, 0x71, 0x61,
	0x6e, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72,
	0x32, 0x93, 0x03, 0x0a, 0x07, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x44, 0x42, 0x12, 0x87, 0x03, 0x0a,
	0x0a, 0x41, 0x64, 0x64, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x44, 0x42, 0x12, 0x1d, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x4d, 0x6f, 0x6e, 0x67,
	0x6f, 0x44, 0x42, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x4d, 0x6f, 0x6e, 0x67, 0x6f,
	0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xb9, 0x02, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1f, 0x22, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2f, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x44, 0x42, 0x2f, 0x41, 0x64, 0x64, 0x3a,
	0x01, 0x2a, 0x92, 0x41, 0x90, 0x02, 0x12, 0x0b, 0x41, 0x64, 0x64, 0x20, 0x4d, 0x6f, 0x6e, 0x67,
	0x6f, 0x44, 0x42, 0x1a, 0x80, 0x02, 0x41, 0x64, 0x64, 0x73, 0x20, 0x4d, 0x6f, 0x6e, 0x67, 0x6f,
	0x44, 0x42, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x73, 0x20, 0x73, 0x65, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x20, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x20, 0x49, 0x74, 0x20, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x69, 0x63, 0x61, 0x6c, 0x6c, 0x79, 0x20, 0x61, 0x64, 0x64, 0x73, 0x20, 0x61, 0x20, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x74, 0x6f, 0x20, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2c, 0x20, 0x77, 0x68, 0x69, 0x63, 0x68, 0x20, 0x69, 0x73, 0x20, 0x72, 0x75, 0x6e,
	0x6e, 0x69, 0x6e, 0x67, 0x20, 0x6f, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x64, 0x20, 0x22, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x2c, 0x20,
	0x74, 0x68, 0x65, 0x6e, 0x20, 0x61, 0x64, 0x64, 0x73, 0x20, 0x22, 0x6d, 0x6f, 0x6e, 0x67, 0x6f,
	0x64, 0x62, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x22, 0x2c, 0x20, 0x61, 0x6e,
	0x64, 0x20, 0x22, 0x71, 0x61, 0x6e, 0x5f, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x5f, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x22, 0x20, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x20,
	0x77, 0x69, 0x74, 0x68, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x64, 0x20, 0x22, 0x70, 0x6d, 0x6d, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x22,
	0x20, 0x61, 0x6e, 0x64, 0x20, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x20, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x42, 0x1f, 0x5a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x3b, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_managementpb_mongodb_proto_rawDescOnce sync.Once
	file_managementpb_mongodb_proto_rawDescData = file_managementpb_mongodb_proto_rawDesc
)

func file_managementpb_mongodb_proto_rawDescGZIP() []byte {
	file_managementpb_mongodb_proto_rawDescOnce.Do(func() {
		file_managementpb_mongodb_proto_rawDescData = protoimpl.X.CompressGZIP(file_managementpb_mongodb_proto_rawDescData)
	})
	return file_managementpb_mongodb_proto_rawDescData
}

var file_managementpb_mongodb_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_managementpb_mongodb_proto_goTypes = []interface{}{
	(*AddMongoDBRequest)(nil),                   // 0: management.AddMongoDBRequest
	(*AddMongoDBResponse)(nil),                  // 1: management.AddMongoDBResponse
	nil,                                         // 2: management.AddMongoDBRequest.CustomLabelsEntry
	(*AddNodeParams)(nil),                       // 3: management.AddNodeParams
	(MetricsMode)(0),                            // 4: management.MetricsMode
	(*inventorypb.MongoDBService)(nil),          // 5: inventory.MongoDBService
	(*inventorypb.MongoDBExporter)(nil),         // 6: inventory.MongoDBExporter
	(*inventorypb.QANMongoDBProfilerAgent)(nil), // 7: inventory.QANMongoDBProfilerAgent
}
var file_managementpb_mongodb_proto_depIdxs = []int32{
	3, // 0: management.AddMongoDBRequest.add_node:type_name -> management.AddNodeParams
	2, // 1: management.AddMongoDBRequest.custom_labels:type_name -> management.AddMongoDBRequest.CustomLabelsEntry
	4, // 2: management.AddMongoDBRequest.metrics_mode:type_name -> management.MetricsMode
	5, // 3: management.AddMongoDBResponse.service:type_name -> inventory.MongoDBService
	6, // 4: management.AddMongoDBResponse.mongodb_exporter:type_name -> inventory.MongoDBExporter
	7, // 5: management.AddMongoDBResponse.qan_mongodb_profiler:type_name -> inventory.QANMongoDBProfilerAgent
	0, // 6: management.MongoDB.AddMongoDB:input_type -> management.AddMongoDBRequest
	1, // 7: management.MongoDB.AddMongoDB:output_type -> management.AddMongoDBResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_managementpb_mongodb_proto_init() }
func file_managementpb_mongodb_proto_init() {
	if File_managementpb_mongodb_proto != nil {
		return
	}
	file_managementpb_metrics_proto_init()
	file_managementpb_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_managementpb_mongodb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddMongoDBRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managementpb_mongodb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddMongoDBResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_managementpb_mongodb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_managementpb_mongodb_proto_goTypes,
		DependencyIndexes: file_managementpb_mongodb_proto_depIdxs,
		MessageInfos:      file_managementpb_mongodb_proto_msgTypes,
	}.Build()
	File_managementpb_mongodb_proto = out.File
	file_managementpb_mongodb_proto_rawDesc = nil
	file_managementpb_mongodb_proto_goTypes = nil
	file_managementpb_mongodb_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MongoDBClient is the client API for MongoDB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MongoDBClient interface {
	// AddMongoDB adds MongoDB Service and starts several Agents.
	// It automatically adds a service to inventory, which is running on provided "node_id",
	// then adds "mongodb_exporter", and "qan_mongodb_profiler" agents
	// with provided "pmm_agent_id" and other parameters.
	AddMongoDB(ctx context.Context, in *AddMongoDBRequest, opts ...grpc.CallOption) (*AddMongoDBResponse, error)
}

type mongoDBClient struct {
	cc grpc.ClientConnInterface
}

func NewMongoDBClient(cc grpc.ClientConnInterface) MongoDBClient {
	return &mongoDBClient{cc}
}

func (c *mongoDBClient) AddMongoDB(ctx context.Context, in *AddMongoDBRequest, opts ...grpc.CallOption) (*AddMongoDBResponse, error) {
	out := new(AddMongoDBResponse)
	err := c.cc.Invoke(ctx, "/management.MongoDB/AddMongoDB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MongoDBServer is the server API for MongoDB service.
type MongoDBServer interface {
	// AddMongoDB adds MongoDB Service and starts several Agents.
	// It automatically adds a service to inventory, which is running on provided "node_id",
	// then adds "mongodb_exporter", and "qan_mongodb_profiler" agents
	// with provided "pmm_agent_id" and other parameters.
	AddMongoDB(context.Context, *AddMongoDBRequest) (*AddMongoDBResponse, error)
}

// UnimplementedMongoDBServer can be embedded to have forward compatible implementations.
type UnimplementedMongoDBServer struct {
}

func (*UnimplementedMongoDBServer) AddMongoDB(context.Context, *AddMongoDBRequest) (*AddMongoDBResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMongoDB not implemented")
}

func RegisterMongoDBServer(s *grpc.Server, srv MongoDBServer) {
	s.RegisterService(&_MongoDB_serviceDesc, srv)
}

func _MongoDB_AddMongoDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMongoDBRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoDBServer).AddMongoDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.MongoDB/AddMongoDB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoDBServer).AddMongoDB(ctx, req.(*AddMongoDBRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MongoDB_serviceDesc = grpc.ServiceDesc{
	ServiceName: "management.MongoDB",
	HandlerType: (*MongoDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddMongoDB",
			Handler:    _MongoDB_AddMongoDB_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/mongodb.proto",
}
