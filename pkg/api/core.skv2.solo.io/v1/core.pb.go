// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: github.com/solo-io/skv2/api/core/v1/core.proto

package v1

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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
	return file_github_com_solo_io_skv2_api_core_v1_core_proto_rawDescGZIP(), []int{5, 0}
}

// Object Selector expression operator, while the set-based syntax differs from Kubernetes (kubernetes: `key: !mylabel`, gloo: `key: mylabel, operator: "!"` | kubernetes: `key: mylabel`, gloo: `key: mylabel, operator: exists`), the functionality remains the same.
type ObjectSelector_Expression_Operator int32

const (
	// =
	ObjectSelector_Expression_Equals ObjectSelector_Expression_Operator = 0
	// ==
	ObjectSelector_Expression_DoubleEquals ObjectSelector_Expression_Operator = 1
	// !=
	ObjectSelector_Expression_NotEquals ObjectSelector_Expression_Operator = 2
	// in
	ObjectSelector_Expression_In ObjectSelector_Expression_Operator = 3
	// notin
	ObjectSelector_Expression_NotIn ObjectSelector_Expression_Operator = 4
	// exists
	ObjectSelector_Expression_Exists ObjectSelector_Expression_Operator = 5
	// !
	ObjectSelector_Expression_DoesNotExist ObjectSelector_Expression_Operator = 6
	// gt
	ObjectSelector_Expression_GreaterThan ObjectSelector_Expression_Operator = 7
	// lt
	ObjectSelector_Expression_LessThan ObjectSelector_Expression_Operator = 8
)

// Enum value maps for ObjectSelector_Expression_Operator.
var (
	ObjectSelector_Expression_Operator_name = map[int32]string{
		0: "Equals",
		1: "DoubleEquals",
		2: "NotEquals",
		3: "In",
		4: "NotIn",
		5: "Exists",
		6: "DoesNotExist",
		7: "GreaterThan",
		8: "LessThan",
	}
	ObjectSelector_Expression_Operator_value = map[string]int32{
		"Equals":       0,
		"DoubleEquals": 1,
		"NotEquals":    2,
		"In":           3,
		"NotIn":        4,
		"Exists":       5,
		"DoesNotExist": 6,
		"GreaterThan":  7,
		"LessThan":     8,
	}
)

func (x ObjectSelector_Expression_Operator) Enum() *ObjectSelector_Expression_Operator {
	p := new(ObjectSelector_Expression_Operator)
	*p = x
	return p
}

func (x ObjectSelector_Expression_Operator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ObjectSelector_Expression_Operator) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_skv2_api_core_v1_core_proto_enumTypes[1].Descriptor()
}

func (ObjectSelector_Expression_Operator) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_skv2_api_core_v1_core_proto_enumTypes[1]
}

func (x ObjectSelector_Expression_Operator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ObjectSelector_Expression_Operator.Descriptor instead.
func (ObjectSelector_Expression_Operator) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_skv2_api_core_v1_core_proto_rawDescGZIP(), []int{6, 1, 0}
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

// Object providing a list of object refs.
// Used to store lists of refs inside a map.
type ObjectRefList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Refs []*ObjectRef `protobuf:"bytes,1,rep,name=refs,proto3" json:"refs,omitempty"`
}

func (x *ObjectRefList) Reset() {
	*x = ObjectRefList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_skv2_api_core_v1_core_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectRefList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectRefList) ProtoMessage() {}

func (x *ObjectRefList) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ObjectRefList.ProtoReflect.Descriptor instead.
func (*ObjectRefList) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_skv2_api_core_v1_core_proto_rawDescGZIP(), []int{1}
}

func (x *ObjectRefList) GetRefs() []*ObjectRef {
	if x != nil {
		return x.Refs
	}
	return nil
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
		mi := &file_github_com_solo_io_skv2_api_core_v1_core_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterObjectRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterObjectRef) ProtoMessage() {}

func (x *ClusterObjectRef) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ClusterObjectRef.ProtoReflect.Descriptor instead.
func (*ClusterObjectRef) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_skv2_api_core_v1_core_proto_rawDescGZIP(), []int{2}
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
		mi := &file_github_com_solo_io_skv2_api_core_v1_core_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypedObjectRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypedObjectRef) ProtoMessage() {}

func (x *TypedObjectRef) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use TypedObjectRef.ProtoReflect.Descriptor instead.
func (*TypedObjectRef) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_skv2_api_core_v1_core_proto_rawDescGZIP(), []int{3}
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
		mi := &file_github_com_solo_io_skv2_api_core_v1_core_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypedClusterObjectRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypedClusterObjectRef) ProtoMessage() {}

func (x *TypedClusterObjectRef) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use TypedClusterObjectRef.ProtoReflect.Descriptor instead.
func (*TypedClusterObjectRef) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_skv2_api_core_v1_core_proto_rawDescGZIP(), []int{4}
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
		mi := &file_github_com_solo_io_skv2_api_core_v1_core_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_skv2_api_core_v1_core_proto_msgTypes[5]
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
	return file_github_com_solo_io_skv2_api_core_v1_core_proto_rawDescGZIP(), []int{5}
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

// Select K8s Objects by namespace, labels, or both.
type ObjectSelector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Select Objects in these namespaces. If omitted, Gloo Mesh will only select Objects in the same namespace
	// as the parent resource (e.g. VirtualGateway) that owns this selector.
	// The reserved value "*" can be used to select objects in all namespaces watched by Gloo Mesh.
	Namespaces []string `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	// Select objects whose labels match the ones specified here.
	Labels map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Expressions allow for more flexible object label matching, such as equality-based requirements, set-based requirements, or a combination of both.
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#equality-based-requirement
	Expressions []*ObjectSelector_Expression `protobuf:"bytes,3,rep,name=expressions,proto3" json:"expressions,omitempty"`
}

func (x *ObjectSelector) Reset() {
	*x = ObjectSelector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_skv2_api_core_v1_core_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectSelector) ProtoMessage() {}

func (x *ObjectSelector) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_skv2_api_core_v1_core_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectSelector.ProtoReflect.Descriptor instead.
func (*ObjectSelector) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_skv2_api_core_v1_core_proto_rawDescGZIP(), []int{6}
}

func (x *ObjectSelector) GetNamespaces() []string {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

func (x *ObjectSelector) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *ObjectSelector) GetExpressions() []*ObjectSelector_Expression {
	if x != nil {
		return x.Expressions
	}
	return nil
}

type ObjectSelector_Expression struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Kubernetes label key, must conform to Kubernetes syntax requirements
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#syntax-and-character-set
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The operator can only be in, notin, =, ==, !=, exists, ! (DoesNotExist), gt (GreaterThan), lt (LessThan).
	Operator ObjectSelector_Expression_Operator `protobuf:"varint,2,opt,name=operator,proto3,enum=core.skv2.solo.io.ObjectSelector_Expression_Operator" json:"operator,omitempty"`
	Values   []string                           `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *ObjectSelector_Expression) Reset() {
	*x = ObjectSelector_Expression{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_skv2_api_core_v1_core_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectSelector_Expression) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectSelector_Expression) ProtoMessage() {}

func (x *ObjectSelector_Expression) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_skv2_api_core_v1_core_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectSelector_Expression.ProtoReflect.Descriptor instead.
func (*ObjectSelector_Expression) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_skv2_api_core_v1_core_proto_rawDescGZIP(), []int{6, 1}
}

func (x *ObjectSelector_Expression) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ObjectSelector_Expression) GetOperator() ObjectSelector_Expression_Operator {
	if x != nil {
		return x.Operator
	}
	return ObjectSelector_Expression_Equals
}

func (x *ObjectSelector_Expression) GetValues() []string {
	if x != nil {
		return x.Values
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
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x41, 0x0a, 0x0d, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x66, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x04, 0x72, 0x65, 0x66, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6b,
	0x76, 0x32, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x66, 0x52, 0x04, 0x72, 0x65, 0x66, 0x73, 0x22, 0x67, 0x0a, 0x10, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0xaf, 0x01, 0x0a, 0x0e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x12, 0x39, 0x0a, 0x09, 0x61, 0x70, 0x69, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x61, 0x70, 0x69, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x30, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0xd9, 0x01, 0x0a, 0x15, 0x54, 0x79, 0x70, 0x65, 0x64, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x12,
	0x39, 0x0a, 0x09, 0x61, 0x70, 0x69, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x08, 0x61, 0x70, 0x69, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x30, 0x0a, 0x04, 0x6b, 0x69,
	0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0xd0, 0x02, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x35, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x73, 0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2f, 0x0a,
	0x13, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x6f, 0x62, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x64, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x43,
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x4b, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0e, 0x0a,
	0x0a, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41,
	0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54,
	0x45, 0x44, 0x10, 0x04, 0x22, 0x98, 0x04, 0x0a, 0x0e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x45, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73,
	0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x4e,
	0x0a, 0x0b, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6b, 0x76, 0x32, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x39,
	0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x93, 0x02, 0x0a, 0x0a, 0x45, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x51, 0x0a, 0x08, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x87, 0x01, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x71, 0x75, 0x61, 0x6c, 0x73, 0x10, 0x00, 0x12, 0x10,
	0x0a, 0x0c, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x45, 0x71, 0x75, 0x61, 0x6c, 0x73, 0x10, 0x01,
	0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x6f, 0x74, 0x45, 0x71, 0x75, 0x61, 0x6c, 0x73, 0x10, 0x02, 0x12,
	0x06, 0x0a, 0x02, 0x49, 0x6e, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x4e, 0x6f, 0x74, 0x49, 0x6e,
	0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x10, 0x05, 0x12, 0x10,
	0x0a, 0x0c, 0x44, 0x6f, 0x65, 0x73, 0x4e, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x10, 0x06,
	0x12, 0x0f, 0x0a, 0x0b, 0x47, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x54, 0x68, 0x61, 0x6e, 0x10,
	0x07, 0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x65, 0x73, 0x73, 0x54, 0x68, 0x61, 0x6e, 0x10, 0x08, 0x42,
	0x42, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6b, 0x76, 0x32, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01, 0xd0,
	0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_github_com_solo_io_skv2_api_core_v1_core_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_github_com_solo_io_skv2_api_core_v1_core_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_github_com_solo_io_skv2_api_core_v1_core_proto_goTypes = []interface{}{
	(Status_State)(0),                       // 0: core.skv2.solo.io.Status.State
	(ObjectSelector_Expression_Operator)(0), // 1: core.skv2.solo.io.ObjectSelector.Expression.Operator
	(*ObjectRef)(nil),                       // 2: core.skv2.solo.io.ObjectRef
	(*ObjectRefList)(nil),                   // 3: core.skv2.solo.io.ObjectRefList
	(*ClusterObjectRef)(nil),                // 4: core.skv2.solo.io.ClusterObjectRef
	(*TypedObjectRef)(nil),                  // 5: core.skv2.solo.io.TypedObjectRef
	(*TypedClusterObjectRef)(nil),           // 6: core.skv2.solo.io.TypedClusterObjectRef
	(*Status)(nil),                          // 7: core.skv2.solo.io.Status
	(*ObjectSelector)(nil),                  // 8: core.skv2.solo.io.ObjectSelector
	nil,                                     // 9: core.skv2.solo.io.ObjectSelector.LabelsEntry
	(*ObjectSelector_Expression)(nil),       // 10: core.skv2.solo.io.ObjectSelector.Expression
	(*wrappers.StringValue)(nil),            // 11: google.protobuf.StringValue
	(*timestamp.Timestamp)(nil),             // 12: google.protobuf.Timestamp
}
var file_github_com_solo_io_skv2_api_core_v1_core_proto_depIdxs = []int32{
	2,  // 0: core.skv2.solo.io.ObjectRefList.refs:type_name -> core.skv2.solo.io.ObjectRef
	11, // 1: core.skv2.solo.io.TypedObjectRef.api_group:type_name -> google.protobuf.StringValue
	11, // 2: core.skv2.solo.io.TypedObjectRef.kind:type_name -> google.protobuf.StringValue
	11, // 3: core.skv2.solo.io.TypedClusterObjectRef.api_group:type_name -> google.protobuf.StringValue
	11, // 4: core.skv2.solo.io.TypedClusterObjectRef.kind:type_name -> google.protobuf.StringValue
	0,  // 5: core.skv2.solo.io.Status.state:type_name -> core.skv2.solo.io.Status.State
	12, // 6: core.skv2.solo.io.Status.processing_time:type_name -> google.protobuf.Timestamp
	11, // 7: core.skv2.solo.io.Status.owner:type_name -> google.protobuf.StringValue
	9,  // 8: core.skv2.solo.io.ObjectSelector.labels:type_name -> core.skv2.solo.io.ObjectSelector.LabelsEntry
	10, // 9: core.skv2.solo.io.ObjectSelector.expressions:type_name -> core.skv2.solo.io.ObjectSelector.Expression
	1,  // 10: core.skv2.solo.io.ObjectSelector.Expression.operator:type_name -> core.skv2.solo.io.ObjectSelector.Expression.Operator
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
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
			switch v := v.(*ObjectRefList); i {
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
		file_github_com_solo_io_skv2_api_core_v1_core_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_github_com_solo_io_skv2_api_core_v1_core_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_github_com_solo_io_skv2_api_core_v1_core_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_github_com_solo_io_skv2_api_core_v1_core_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectSelector); i {
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
		file_github_com_solo_io_skv2_api_core_v1_core_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectSelector_Expression); i {
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
			NumEnums:      2,
			NumMessages:   9,
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
