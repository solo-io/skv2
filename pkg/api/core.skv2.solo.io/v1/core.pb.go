// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: github.com/solo-io/skv2/api/core/v1/core.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The State of a reconciled object
type Status_State int32

const (
	// Waiting to be processed.
	Status_PENDING Status_State = 0
	// Currently processing.
	Status_PROCESSING Status_State = 1
	// Invalid parameters supplied, will not continue.
	Status_INVALID Status_State = 2
	// Failed during processing.
	Status_FAILED Status_State = 3
	// Finished processing successfully.
	Status_ACCEPTED Status_State = 4
)

// Enum value maps for Status_State.
var (
	Status_State_name = map[int32]string{
		0: "PENDING",
		1: "PROCESSING",
		2: "INVALID",
		3: "FAILED",
		4: "ACCEPTED",
	}
	Status_State_value = map[string]int32{
		"PENDING":    0,
		"PROCESSING": 1,
		"INVALID":    2,
		"FAILED":     3,
		"ACCEPTED":   4,
	}
)

func (x Status_State) Enum() *Status_State {
	p := new(Status_State)
	*p = x
	return p
}

func (x Status_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status_State) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_skv2_api_core_v1_core_proto_enumTypes[0].Descriptor()
}

func (Status_State) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_skv2_api_core_v1_core_proto_enumTypes[0]
}

func (x Status_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status_State.Descriptor instead.
func (Status_State) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_skv2_api_core_v1_core_proto_rawDescGZIP(), []int{4, 0}
}

// Resource reference for an object
type ObjectRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the resource being referenced
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// namespace of the resource being referenced
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *ObjectRef) Reset() {
	*x = ObjectRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_skv2_api_core_v1_core_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectRef) ProtoMessage() {}

func (x *ObjectRef) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_skv2_api_core_v1_core_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectRef.ProtoReflect.Descriptor instead.
func (*ObjectRef) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_skv2_api_core_v1_core_proto_rawDescGZIP(), []int{0}
}

func (x *ObjectRef) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ObjectRef) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

// Resource reference for a cross-cluster-scoped object
type ClusterObjectRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the resource being referenced
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// namespace of the resource being referenced
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// name of the cluster in which the resource exists
	ClusterName string `protobuf:"bytes,3,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
}

func (x *ClusterObjectRef) Reset() {
	*x = ClusterObjectRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_skv2_api_core_v1_core_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterObjectRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterObjectRef) ProtoMessage() {}

func (x *ClusterObjectRef) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_skv2_api_core_v1_core_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterObjectRef.ProtoReflect.Descriptor instead.
func (*ClusterObjectRef) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_skv2_api_core_v1_core_proto_rawDescGZIP(), []int{1}
}

func (x *ClusterObjectRef) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClusterObjectRef) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ClusterObjectRef) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

// Resource reference for a typed object
type TypedObjectRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// API group of the resource being referenced
	ApiGroup *wrappers.StringValue `protobuf:"bytes,1,opt,name=api_group,json=apiGroup,proto3" json:"api_group,omitempty"`
	// Kind of the resource being referenced
	Kind *wrappers.StringValue `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// name of the resource being referenced
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// namespace of the resource being referenced
	Namespace string `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *TypedObjectRef) Reset() {
	*x = TypedObjectRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_skv2_api_core_v1_core_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypedObjectRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypedObjectRef) ProtoMessage() {}

func (x *TypedObjectRef) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_skv2_api_core_v1_core_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypedObjectRef.ProtoReflect.Descriptor instead.
func (*TypedObjectRef) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_skv2_api_core_v1_core_proto_rawDescGZIP(), []int{2}
}

func (x *TypedObjectRef) GetApiGroup() *wrappers.StringValue {
	if x != nil {
		return x.ApiGroup
	}
	return nil
}

func (x *TypedObjectRef) GetKind() *wrappers.StringValue {
	if x != nil {
		return x.Kind
	}
	return nil
}

func (x *TypedObjectRef) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TypedObjectRef) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

// Resource reference for a typed, cross-cluster-scoped object
type TypedClusterObjectRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// API group of the resource being referenced
	ApiGroup *wrappers.StringValue `protobuf:"bytes,1,opt,name=api_group,json=apiGroup,proto3" json:"api_group,omitempty"`
	// Kind of the resource being referenced
	Kind *wrappers.StringValue `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// name of the resource being referenced
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// namespace of the resource being referenced
	Namespace string `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// name of the cluster in which the resource exists
	ClusterName string `protobuf:"bytes,5,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
}

func (x *TypedClusterObjectRef) Reset() {
	*x = TypedClusterObjectRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_skv2_api_core_v1_core_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypedClusterObjectRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypedClusterObjectRef) ProtoMessage() {}

func (x *TypedClusterObjectRef) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_skv2_api_core_v1_core_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypedClusterObjectRef.ProtoReflect.Descriptor instead.
func (*TypedClusterObjectRef) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_skv2_api_core_v1_core_proto_rawDescGZIP(), []int{3}
}

func (x *TypedClusterObjectRef) GetApiGroup() *wrappers.StringValue {
	if x != nil {
		return x.ApiGroup
	}
	return nil
}

func (x *TypedClusterObjectRef) GetKind() *wrappers.StringValue {
	if x != nil {
		return x.Kind
	}
	return nil
}

func (x *TypedClusterObjectRef) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TypedClusterObjectRef) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *TypedClusterObjectRef) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

// A generic status
type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The current state of the resource
	State Status_State `protobuf:"varint,1,opt,name=state,proto3,enum=core.skv2.solo.io.Status_State" json:"state,omitempty"`
	// A human readable message about the current state of the object
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// The most recently observed generation of the resource. This value corresponds to the `metadata.generation` of
	// a kubernetes resource
	ObservedGeneration int64 `protobuf:"varint,3,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// The time at which this status was recorded
	ProcessingTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=processing_time,json=processingTime,proto3" json:"processing_time,omitempty"`
	// (optional) The owner of the status, this value can be used to identify the entity which wrote this status.
	// This is useful in situations where a given resource may have multiple owners.
	Owner *wrappers.StringValue `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_skv2_api_core_v1_core_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_skv2_api_core_v1_core_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_skv2_api_core_v1_core_proto_rawDescGZIP(), []int{4}
}

func (x *Status) GetState() Status_State {
	if x != nil {
		return x.State
	}
	return Status_PENDING
}

func (x *Status) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Status) GetObservedGeneration() int64 {
	if x != nil {
		return x.ObservedGeneration
	}
	return 0
}

func (x *Status) GetProcessingTime() *timestamp.Timestamp {
	if x != nil {
		return x.ProcessingTime
	}
	return nil
}

func (x *Status) GetOwner() *wrappers.StringValue {
	if x != nil {
		return x.Owner
	}
	return nil
}

var File_github_com_solo_io_skv2_api_core_v1_core_proto protoreflect.FileDescriptor

var file_github_com_solo_io_skv2_api_core_v1_core_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6b, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65,
	0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x09, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x67, 0x0a, 0x10, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0xaf, 0x01, 0x0a, 0x0e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x66, 0x12, 0x39, 0x0a, 0x09, 0x61, 0x70, 0x69, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x61, 0x70, 0x69, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x30,
	0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x22, 0xd9, 0x01, 0x0a, 0x15, 0x54, 0x79, 0x70, 0x65, 0x64, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x12, 0x39, 0x0a, 0x09,
	0x61, 0x70, 0x69, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x61,
	0x70, 0x69, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x30, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xd0,
	0x02, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x35, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x73, 0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x6f, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x64, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x0f, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x32, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x22, 0x4b, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x52,
	0x4f, 0x43, 0x45, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45,
	0x44, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10,
	0x04, 0x42, 0x3e, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6b, 0x76, 0x32, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6b, 0x76, 0x32, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04,
	0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_skv2_api_core_v1_core_proto_rawDescOnce sync.Once
	file_github_com_solo_io_skv2_api_core_v1_core_proto_rawDescData = file_github_com_solo_io_skv2_api_core_v1_core_proto_rawDesc
)

func file_github_com_solo_io_skv2_api_core_v1_core_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_skv2_api_core_v1_core_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_skv2_api_core_v1_core_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_skv2_api_core_v1_core_proto_rawDescData)
	})
	return file_github_com_solo_io_skv2_api_core_v1_core_proto_rawDescData
}

var file_github_com_solo_io_skv2_api_core_v1_core_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_skv2_api_core_v1_core_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_solo_io_skv2_api_core_v1_core_proto_goTypes = []interface{}{
	(Status_State)(0),             // 0: core.skv2.solo.io.Status.State
	(*ObjectRef)(nil),             // 1: core.skv2.solo.io.ObjectRef
	(*ClusterObjectRef)(nil),      // 2: core.skv2.solo.io.ClusterObjectRef
	(*TypedObjectRef)(nil),        // 3: core.skv2.solo.io.TypedObjectRef
	(*TypedClusterObjectRef)(nil), // 4: core.skv2.solo.io.TypedClusterObjectRef
	(*Status)(nil),                // 5: core.skv2.solo.io.Status
	(*wrappers.StringValue)(nil),  // 6: google.protobuf.StringValue
	(*timestamp.Timestamp)(nil),   // 7: google.protobuf.Timestamp
}
var file_github_com_solo_io_skv2_api_core_v1_core_proto_depIdxs = []int32{
	6, // 0: core.skv2.solo.io.TypedObjectRef.api_group:type_name -> google.protobuf.StringValue
	6, // 1: core.skv2.solo.io.TypedObjectRef.kind:type_name -> google.protobuf.StringValue
	6, // 2: core.skv2.solo.io.TypedClusterObjectRef.api_group:type_name -> google.protobuf.StringValue
	6, // 3: core.skv2.solo.io.TypedClusterObjectRef.kind:type_name -> google.protobuf.StringValue
	0, // 4: core.skv2.solo.io.Status.state:type_name -> core.skv2.solo.io.Status.State
	7, // 5: core.skv2.solo.io.Status.processing_time:type_name -> google.protobuf.Timestamp
	6, // 6: core.skv2.solo.io.Status.owner:type_name -> google.protobuf.StringValue
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_skv2_api_core_v1_core_proto_init() }
func file_github_com_solo_io_skv2_api_core_v1_core_proto_init() {
	if File_github_com_solo_io_skv2_api_core_v1_core_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_skv2_api_core_v1_core_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectRef); i {
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
		file_github_com_solo_io_skv2_api_core_v1_core_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterObjectRef); i {
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
		file_github_com_solo_io_skv2_api_core_v1_core_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypedObjectRef); i {
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
		file_github_com_solo_io_skv2_api_core_v1_core_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypedClusterObjectRef); i {
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
		file_github_com_solo_io_skv2_api_core_v1_core_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_github_com_solo_io_skv2_api_core_v1_core_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_skv2_api_core_v1_core_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_skv2_api_core_v1_core_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_skv2_api_core_v1_core_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_skv2_api_core_v1_core_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_skv2_api_core_v1_core_proto = out.File
	file_github_com_solo_io_skv2_api_core_v1_core_proto_rawDesc = nil
	file_github_com_solo_io_skv2_api_core_v1_core_proto_goTypes = nil
	file_github_com_solo_io_skv2_api_core_v1_core_proto_depIdxs = nil
}
