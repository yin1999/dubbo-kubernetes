// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.0
// source: api/system/v1alpha1/zone_insight.proto

package v1alpha1

import (
	reflect "reflect"
	sync "sync"
)

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"

	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

import (
	_ "github.com/apache/dubbo-kubernetes/api/mesh"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ZoneInsight struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of DDS subscriptions created by a given Zone Dubbo CP.
	Subscriptions []*DDSSubscription `protobuf:"bytes,1,rep,name=subscriptions,proto3" json:"subscriptions,omitempty"`
	// Statistics about Envoy Admin Streams
	EnvoyAdminStreams *EnvoyAdminStreams `protobuf:"bytes,2,opt,name=envoy_admin_streams,json=envoyAdminStreams,proto3" json:"envoy_admin_streams,omitempty"`
	HealthCheck       *HealthCheck       `protobuf:"bytes,3,opt,name=health_check,json=healthCheck,proto3" json:"health_check,omitempty"`
}

func (x *ZoneInsight) Reset() {
	*x = ZoneInsight{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_system_v1alpha1_zone_insight_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZoneInsight) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZoneInsight) ProtoMessage() {}

func (x *ZoneInsight) ProtoReflect() protoreflect.Message {
	mi := &file_api_system_v1alpha1_zone_insight_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZoneInsight.ProtoReflect.Descriptor instead.
func (*ZoneInsight) Descriptor() ([]byte, []int) {
	return file_api_system_v1alpha1_zone_insight_proto_rawDescGZIP(), []int{0}
}

func (x *ZoneInsight) GetSubscriptions() []*DDSSubscription {
	if x != nil {
		return x.Subscriptions
	}
	return nil
}

func (x *ZoneInsight) GetEnvoyAdminStreams() *EnvoyAdminStreams {
	if x != nil {
		return x.EnvoyAdminStreams
	}
	return nil
}

func (x *ZoneInsight) GetHealthCheck() *HealthCheck {
	if x != nil {
		return x.HealthCheck
	}
	return nil
}

type EnvoyAdminStreams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Global instance ID that handles XDS Config Dump streams.
	ConfigDumpGlobalInstanceId string `protobuf:"bytes,1,opt,name=config_dump_global_instance_id,json=configDumpGlobalInstanceId,proto3" json:"config_dump_global_instance_id,omitempty"`
	// Global instance ID that handles Stats streams.
	StatsGlobalInstanceId string `protobuf:"bytes,2,opt,name=stats_global_instance_id,json=statsGlobalInstanceId,proto3" json:"stats_global_instance_id,omitempty"`
	// Global instance ID that handles Clusters streams.
	ClustersGlobalInstanceId string `protobuf:"bytes,3,opt,name=clusters_global_instance_id,json=clustersGlobalInstanceId,proto3" json:"clusters_global_instance_id,omitempty"`
}

func (x *EnvoyAdminStreams) Reset() {
	*x = EnvoyAdminStreams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_system_v1alpha1_zone_insight_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvoyAdminStreams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvoyAdminStreams) ProtoMessage() {}

func (x *EnvoyAdminStreams) ProtoReflect() protoreflect.Message {
	mi := &file_api_system_v1alpha1_zone_insight_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvoyAdminStreams.ProtoReflect.Descriptor instead.
func (*EnvoyAdminStreams) Descriptor() ([]byte, []int) {
	return file_api_system_v1alpha1_zone_insight_proto_rawDescGZIP(), []int{1}
}

func (x *EnvoyAdminStreams) GetConfigDumpGlobalInstanceId() string {
	if x != nil {
		return x.ConfigDumpGlobalInstanceId
	}
	return ""
}

func (x *EnvoyAdminStreams) GetStatsGlobalInstanceId() string {
	if x != nil {
		return x.StatsGlobalInstanceId
	}
	return ""
}

func (x *EnvoyAdminStreams) GetClustersGlobalInstanceId() string {
	if x != nil {
		return x.ClustersGlobalInstanceId
	}
	return ""
}

// DDSSubscription describes a single DDS subscription
// created by a Zone to the Global.
// Ideally, there should be only one such subscription per Zone lifecycle.
// Presence of multiple subscriptions might indicate one of the following
// events:
// - transient loss of network connection between Zone and Global Control
// Planes
// - Zone Dubbo CP restarts (i.e. hot restart or crash)
// - Global Dubbo CP restarts (i.e. rolling update or crash)
// - etc
type DDSSubscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique id per DDS subscription.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Global CP instance that handled given subscription.
	GlobalInstanceId string `protobuf:"bytes,2,opt,name=global_instance_id,json=globalInstanceId,proto3" json:"global_instance_id,omitempty"`
	// Time when a given Zone connected to the Global.
	ConnectTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=connect_time,json=connectTime,proto3" json:"connect_time,omitempty"`
	// Time when a given Zone disconnected from the Global.
	DisconnectTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=disconnect_time,json=disconnectTime,proto3" json:"disconnect_time,omitempty"`
	// Status of the DDS subscription.
	Status *DDSSubscriptionStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	// Generation is an integer number which is periodically increased by the
	// status sink
	Generation uint32 `protobuf:"varint,7,opt,name=generation,proto3" json:"generation,omitempty"`
	// Config of Zone Dubbo CP
	Config string `protobuf:"bytes,8,opt,name=config,proto3" json:"config,omitempty"`
	// Indicates if subscription provided auth token
	AuthTokenProvided bool `protobuf:"varint,9,opt,name=auth_token_provided,json=authTokenProvided,proto3" json:"auth_token_provided,omitempty"`
	// Zone CP instance that handled the given subscription (This is the leader at
	// time of connection).
	ZoneInstanceId string `protobuf:"bytes,10,opt,name=zone_instance_id,json=zoneInstanceId,proto3" json:"zone_instance_id,omitempty"`
}

func (x *DDSSubscription) Reset() {
	*x = DDSSubscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_system_v1alpha1_zone_insight_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DDSSubscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DDSSubscription) ProtoMessage() {}

func (x *DDSSubscription) ProtoReflect() protoreflect.Message {
	mi := &file_api_system_v1alpha1_zone_insight_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DDSSubscription.ProtoReflect.Descriptor instead.
func (*DDSSubscription) Descriptor() ([]byte, []int) {
	return file_api_system_v1alpha1_zone_insight_proto_rawDescGZIP(), []int{2}
}

func (x *DDSSubscription) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DDSSubscription) GetGlobalInstanceId() string {
	if x != nil {
		return x.GlobalInstanceId
	}
	return ""
}

func (x *DDSSubscription) GetConnectTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ConnectTime
	}
	return nil
}

func (x *DDSSubscription) GetDisconnectTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DisconnectTime
	}
	return nil
}

func (x *DDSSubscription) GetStatus() *DDSSubscriptionStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *DDSSubscription) GetGeneration() uint32 {
	if x != nil {
		return x.Generation
	}
	return 0
}

func (x *DDSSubscription) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *DDSSubscription) GetAuthTokenProvided() bool {
	if x != nil {
		return x.AuthTokenProvided
	}
	return false
}

func (x *DDSSubscription) GetZoneInstanceId() string {
	if x != nil {
		return x.ZoneInstanceId
	}
	return ""
}

// DDSSubscriptionStatus defines status of an DDS subscription.
type DDSSubscriptionStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Time when status of a given DDS subscription was most recently updated.
	LastUpdateTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=last_update_time,json=lastUpdateTime,proto3" json:"last_update_time,omitempty"`
	// Total defines an aggregate over individual DDS stats.
	Total *DDSServiceStats            `protobuf:"bytes,2,opt,name=total,proto3" json:"total,omitempty"`
	Stat  map[string]*DDSServiceStats `protobuf:"bytes,3,rep,name=stat,proto3" json:"stat,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DDSSubscriptionStatus) Reset() {
	*x = DDSSubscriptionStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_system_v1alpha1_zone_insight_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DDSSubscriptionStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DDSSubscriptionStatus) ProtoMessage() {}

func (x *DDSSubscriptionStatus) ProtoReflect() protoreflect.Message {
	mi := &file_api_system_v1alpha1_zone_insight_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DDSSubscriptionStatus.ProtoReflect.Descriptor instead.
func (*DDSSubscriptionStatus) Descriptor() ([]byte, []int) {
	return file_api_system_v1alpha1_zone_insight_proto_rawDescGZIP(), []int{3}
}

func (x *DDSSubscriptionStatus) GetLastUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdateTime
	}
	return nil
}

func (x *DDSSubscriptionStatus) GetTotal() *DDSServiceStats {
	if x != nil {
		return x.Total
	}
	return nil
}

func (x *DDSSubscriptionStatus) GetStat() map[string]*DDSServiceStats {
	if x != nil {
		return x.Stat
	}
	return nil
}

// DiscoveryServiceStats defines all stats over a single xDS service.
type DDSServiceStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Number of xDS responses sent to the Dataplane.
	ResponsesSent uint64 `protobuf:"varint,1,opt,name=responses_sent,json=responsesSent,proto3" json:"responses_sent,omitempty"`
	// Number of xDS responses ACKed by the Dataplane.
	ResponsesAcknowledged uint64 `protobuf:"varint,2,opt,name=responses_acknowledged,json=responsesAcknowledged,proto3" json:"responses_acknowledged,omitempty"`
	// Number of xDS responses NACKed by the Dataplane.
	ResponsesRejected uint64 `protobuf:"varint,3,opt,name=responses_rejected,json=responsesRejected,proto3" json:"responses_rejected,omitempty"`
}

func (x *DDSServiceStats) Reset() {
	*x = DDSServiceStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_system_v1alpha1_zone_insight_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DDSServiceStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DDSServiceStats) ProtoMessage() {}

func (x *DDSServiceStats) ProtoReflect() protoreflect.Message {
	mi := &file_api_system_v1alpha1_zone_insight_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DDSServiceStats.ProtoReflect.Descriptor instead.
func (*DDSServiceStats) Descriptor() ([]byte, []int) {
	return file_api_system_v1alpha1_zone_insight_proto_rawDescGZIP(), []int{4}
}

func (x *DDSServiceStats) GetResponsesSent() uint64 {
	if x != nil {
		return x.ResponsesSent
	}
	return 0
}

func (x *DDSServiceStats) GetResponsesAcknowledged() uint64 {
	if x != nil {
		return x.ResponsesAcknowledged
	}
	return 0
}

func (x *DDSServiceStats) GetResponsesRejected() uint64 {
	if x != nil {
		return x.ResponsesRejected
	}
	return 0
}

// HealthCheck holds information about the received zone health check
type HealthCheck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Time last health check received
	Time *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *HealthCheck) Reset() {
	*x = HealthCheck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_system_v1alpha1_zone_insight_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheck) ProtoMessage() {}

func (x *HealthCheck) ProtoReflect() protoreflect.Message {
	mi := &file_api_system_v1alpha1_zone_insight_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheck.ProtoReflect.Descriptor instead.
func (*HealthCheck) Descriptor() ([]byte, []int) {
	return file_api_system_v1alpha1_zone_insight_proto_rawDescGZIP(), []int{5}
}

func (x *HealthCheck) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

var File_api_system_v1alpha1_zone_insight_proto protoreflect.FileDescriptor

var file_api_system_v1alpha1_zone_insight_proto_rawDesc = []byte{
	0x0a, 0x26, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x69, 0x67,
	0x68, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a,
	0x16, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe2, 0x02, 0x0a, 0x0b, 0x5a, 0x6f, 0x6e,
	0x65, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x12, 0x4c, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x44, 0x53, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x58, 0x0a, 0x13, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x5f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x6e, 0x76, 0x6f,
	0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x11, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73,
	0x12, 0x45, 0x0a, 0x0c, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x0b, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x3a, 0x64, 0xaa, 0x8c, 0x89, 0xa6, 0x01, 0x15, 0x0a,
	0x13, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0xaa, 0x8c, 0x89, 0xa6, 0x01, 0x0d, 0x12, 0x0b, 0x5a, 0x6f, 0x6e, 0x65,
	0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0xaa, 0x8c, 0x89, 0xa6, 0x01, 0x08, 0x22, 0x06, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0xaa, 0x8c, 0x89, 0xa6, 0x01, 0x02, 0x18, 0x01, 0xaa, 0x8c, 0x89,
	0xa6, 0x01, 0x10, 0x3a, 0x0e, 0x0a, 0x0c, 0x7a, 0x6f, 0x6e, 0x65, 0x2d, 0x69, 0x6e, 0x73, 0x69,
	0x67, 0x68, 0x74, 0xaa, 0x8c, 0x89, 0xa6, 0x01, 0x04, 0x3a, 0x02, 0x18, 0x01, 0x22, 0xcf, 0x01,
	0x0a, 0x11, 0x45, 0x6e, 0x76, 0x6f, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x73, 0x12, 0x42, 0x0a, 0x1e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x64, 0x75,
	0x6d, 0x70, 0x5f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1a, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x44, 0x75, 0x6d, 0x70, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x18, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x5f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x3d, 0x0a, 0x1b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x5f, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x47,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x22,
	0xab, 0x03, 0x0a, 0x0f, 0x44, 0x44, 0x53, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x43, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x44,
	0x53, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x2e, 0x0a, 0x13, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x11, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x7a,
	0x6f, 0x6e, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x22, 0xc8, 0x02,
	0x0a, 0x15, 0x44, 0x44, 0x53, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x44, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x6c,
	0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3c, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x64,
	0x75, 0x62, 0x62, 0x6f, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x44, 0x53, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x4a, 0x0a, 0x04, 0x73,
	0x74, 0x61, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x64, 0x75, 0x62, 0x62,
	0x6f, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x44, 0x44, 0x53, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x04, 0x73, 0x74, 0x61, 0x74, 0x1a, 0x5f, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x44,
	0x53, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9e, 0x01, 0x0a, 0x0f, 0x44, 0x44, 0x53,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x0e,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x53,
	0x65, 0x6e, 0x74, 0x12, 0x35, 0x0a, 0x16, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73,
	0x5f, 0x61, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x15, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x41, 0x63,
	0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x64, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x73, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x22, 0x3d, 0x0a, 0x0b, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x64, 0x75,
	0x62, 0x62, 0x6f, 0x2d, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_system_v1alpha1_zone_insight_proto_rawDescOnce sync.Once
	file_api_system_v1alpha1_zone_insight_proto_rawDescData = file_api_system_v1alpha1_zone_insight_proto_rawDesc
)

func file_api_system_v1alpha1_zone_insight_proto_rawDescGZIP() []byte {
	file_api_system_v1alpha1_zone_insight_proto_rawDescOnce.Do(func() {
		file_api_system_v1alpha1_zone_insight_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_system_v1alpha1_zone_insight_proto_rawDescData)
	})
	return file_api_system_v1alpha1_zone_insight_proto_rawDescData
}

var file_api_system_v1alpha1_zone_insight_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_system_v1alpha1_zone_insight_proto_goTypes = []interface{}{
	(*ZoneInsight)(nil),           // 0: dubbo.system.v1alpha1.ZoneInsight
	(*EnvoyAdminStreams)(nil),     // 1: dubbo.system.v1alpha1.EnvoyAdminStreams
	(*DDSSubscription)(nil),       // 2: dubbo.system.v1alpha1.DDSSubscription
	(*DDSSubscriptionStatus)(nil), // 3: dubbo.system.v1alpha1.DDSSubscriptionStatus
	(*DDSServiceStats)(nil),       // 4: dubbo.system.v1alpha1.DDSServiceStats
	(*HealthCheck)(nil),           // 5: dubbo.system.v1alpha1.HealthCheck
	nil,                           // 6: dubbo.system.v1alpha1.DDSSubscriptionStatus.StatEntry
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_api_system_v1alpha1_zone_insight_proto_depIdxs = []int32{
	2,  // 0: dubbo.system.v1alpha1.ZoneInsight.subscriptions:type_name -> dubbo.system.v1alpha1.DDSSubscription
	1,  // 1: dubbo.system.v1alpha1.ZoneInsight.envoy_admin_streams:type_name -> dubbo.system.v1alpha1.EnvoyAdminStreams
	5,  // 2: dubbo.system.v1alpha1.ZoneInsight.health_check:type_name -> dubbo.system.v1alpha1.HealthCheck
	7,  // 3: dubbo.system.v1alpha1.DDSSubscription.connect_time:type_name -> google.protobuf.Timestamp
	7,  // 4: dubbo.system.v1alpha1.DDSSubscription.disconnect_time:type_name -> google.protobuf.Timestamp
	3,  // 5: dubbo.system.v1alpha1.DDSSubscription.status:type_name -> dubbo.system.v1alpha1.DDSSubscriptionStatus
	7,  // 6: dubbo.system.v1alpha1.DDSSubscriptionStatus.last_update_time:type_name -> google.protobuf.Timestamp
	4,  // 7: dubbo.system.v1alpha1.DDSSubscriptionStatus.total:type_name -> dubbo.system.v1alpha1.DDSServiceStats
	6,  // 8: dubbo.system.v1alpha1.DDSSubscriptionStatus.stat:type_name -> dubbo.system.v1alpha1.DDSSubscriptionStatus.StatEntry
	7,  // 9: dubbo.system.v1alpha1.HealthCheck.time:type_name -> google.protobuf.Timestamp
	4,  // 10: dubbo.system.v1alpha1.DDSSubscriptionStatus.StatEntry.value:type_name -> dubbo.system.v1alpha1.DDSServiceStats
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_api_system_v1alpha1_zone_insight_proto_init() }
func file_api_system_v1alpha1_zone_insight_proto_init() {
	if File_api_system_v1alpha1_zone_insight_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_system_v1alpha1_zone_insight_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZoneInsight); i {
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
		file_api_system_v1alpha1_zone_insight_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvoyAdminStreams); i {
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
		file_api_system_v1alpha1_zone_insight_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DDSSubscription); i {
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
		file_api_system_v1alpha1_zone_insight_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DDSSubscriptionStatus); i {
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
		file_api_system_v1alpha1_zone_insight_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DDSServiceStats); i {
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
		file_api_system_v1alpha1_zone_insight_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheck); i {
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
			RawDescriptor: file_api_system_v1alpha1_zone_insight_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_system_v1alpha1_zone_insight_proto_goTypes,
		DependencyIndexes: file_api_system_v1alpha1_zone_insight_proto_depIdxs,
		MessageInfos:      file_api_system_v1alpha1_zone_insight_proto_msgTypes,
	}.Build()
	File_api_system_v1alpha1_zone_insight_proto = out.File
	file_api_system_v1alpha1_zone_insight_proto_rawDesc = nil
	file_api_system_v1alpha1_zone_insight_proto_goTypes = nil
	file_api_system_v1alpha1_zone_insight_proto_depIdxs = nil
}