// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/skv2/api/multicluster/v1alpha1/cluster.proto

package v1alpha1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//Representation of a Kubernetes cluster that has been registered.
type KubernetesClusterSpec struct {
	// name of the secret which contains the kubeconfig with information to connect to the remote cluster.
	SecretName string `protobuf:"bytes,1,opt,name=secret_name,json=secretName,proto3" json:"secret_name,omitempty"`
	// name local DNS suffix used by the cluster.
	// used for building FQDNs for in-cluster services
	// defaults to 'cluster.local'
	ClusterDomain string `protobuf:"bytes,2,opt,name=cluster_domain,json=clusterDomain,proto3" json:"cluster_domain,omitempty"`
	// Metadata for clusters provisioned from cloud providers.
	ProviderInfo         *KubernetesClusterSpec_ProviderInfo `protobuf:"bytes,3,opt,name=provider_info,json=providerInfo,proto3" json:"provider_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *KubernetesClusterSpec) Reset()         { *m = KubernetesClusterSpec{} }
func (m *KubernetesClusterSpec) String() string { return proto.CompactTextString(m) }
func (*KubernetesClusterSpec) ProtoMessage()    {}
func (*KubernetesClusterSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_533baf1d8d7a56a7, []int{0}
}
func (m *KubernetesClusterSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KubernetesClusterSpec.Unmarshal(m, b)
}
func (m *KubernetesClusterSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KubernetesClusterSpec.Marshal(b, m, deterministic)
}
func (m *KubernetesClusterSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KubernetesClusterSpec.Merge(m, src)
}
func (m *KubernetesClusterSpec) XXX_Size() int {
	return xxx_messageInfo_KubernetesClusterSpec.Size(m)
}
func (m *KubernetesClusterSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_KubernetesClusterSpec.DiscardUnknown(m)
}

var xxx_messageInfo_KubernetesClusterSpec proto.InternalMessageInfo

func (m *KubernetesClusterSpec) GetSecretName() string {
	if m != nil {
		return m.SecretName
	}
	return ""
}

func (m *KubernetesClusterSpec) GetClusterDomain() string {
	if m != nil {
		return m.ClusterDomain
	}
	return ""
}

func (m *KubernetesClusterSpec) GetProviderInfo() *KubernetesClusterSpec_ProviderInfo {
	if m != nil {
		return m.ProviderInfo
	}
	return nil
}

// Metadata for clusters provisioned from cloud providers.
type KubernetesClusterSpec_ProviderInfo struct {
	// Metadata specific to a cloud provider.
	//
	// Types that are valid to be assigned to ProviderInfoType:
	//	*KubernetesClusterSpec_ProviderInfo_Eks
	ProviderInfoType     isKubernetesClusterSpec_ProviderInfo_ProviderInfoType `protobuf_oneof:"provider_info_type"`
	XXX_NoUnkeyedLiteral struct{}                                              `json:"-"`
	XXX_unrecognized     []byte                                                `json:"-"`
	XXX_sizecache        int32                                                 `json:"-"`
}

func (m *KubernetesClusterSpec_ProviderInfo) Reset()         { *m = KubernetesClusterSpec_ProviderInfo{} }
func (m *KubernetesClusterSpec_ProviderInfo) String() string { return proto.CompactTextString(m) }
func (*KubernetesClusterSpec_ProviderInfo) ProtoMessage()    {}
func (*KubernetesClusterSpec_ProviderInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_533baf1d8d7a56a7, []int{0, 0}
}
func (m *KubernetesClusterSpec_ProviderInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KubernetesClusterSpec_ProviderInfo.Unmarshal(m, b)
}
func (m *KubernetesClusterSpec_ProviderInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KubernetesClusterSpec_ProviderInfo.Marshal(b, m, deterministic)
}
func (m *KubernetesClusterSpec_ProviderInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KubernetesClusterSpec_ProviderInfo.Merge(m, src)
}
func (m *KubernetesClusterSpec_ProviderInfo) XXX_Size() int {
	return xxx_messageInfo_KubernetesClusterSpec_ProviderInfo.Size(m)
}
func (m *KubernetesClusterSpec_ProviderInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_KubernetesClusterSpec_ProviderInfo.DiscardUnknown(m)
}

var xxx_messageInfo_KubernetesClusterSpec_ProviderInfo proto.InternalMessageInfo

type isKubernetesClusterSpec_ProviderInfo_ProviderInfoType interface {
	isKubernetesClusterSpec_ProviderInfo_ProviderInfoType()
	Equal(interface{}) bool
}

type KubernetesClusterSpec_ProviderInfo_Eks struct {
	Eks *KubernetesClusterSpec_Eks `protobuf:"bytes,1,opt,name=eks,proto3,oneof" json:"eks,omitempty"`
}

func (*KubernetesClusterSpec_ProviderInfo_Eks) isKubernetesClusterSpec_ProviderInfo_ProviderInfoType() {
}

func (m *KubernetesClusterSpec_ProviderInfo) GetProviderInfoType() isKubernetesClusterSpec_ProviderInfo_ProviderInfoType {
	if m != nil {
		return m.ProviderInfoType
	}
	return nil
}

func (m *KubernetesClusterSpec_ProviderInfo) GetEks() *KubernetesClusterSpec_Eks {
	if x, ok := m.GetProviderInfoType().(*KubernetesClusterSpec_ProviderInfo_Eks); ok {
		return x.Eks
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*KubernetesClusterSpec_ProviderInfo) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*KubernetesClusterSpec_ProviderInfo_Eks)(nil),
	}
}

// AWS metadata associated with an EKS provisioned cluster.
type KubernetesClusterSpec_Eks struct {
	// AWS ARN.
	Arn string `protobuf:"bytes,1,opt,name=arn,proto3" json:"arn,omitempty"`
	// AWS 12 digit account ID.
	AccountId string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// AWS region.
	Region string `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	// EKS resource name.
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KubernetesClusterSpec_Eks) Reset()         { *m = KubernetesClusterSpec_Eks{} }
func (m *KubernetesClusterSpec_Eks) String() string { return proto.CompactTextString(m) }
func (*KubernetesClusterSpec_Eks) ProtoMessage()    {}
func (*KubernetesClusterSpec_Eks) Descriptor() ([]byte, []int) {
	return fileDescriptor_533baf1d8d7a56a7, []int{0, 1}
}
func (m *KubernetesClusterSpec_Eks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KubernetesClusterSpec_Eks.Unmarshal(m, b)
}
func (m *KubernetesClusterSpec_Eks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KubernetesClusterSpec_Eks.Marshal(b, m, deterministic)
}
func (m *KubernetesClusterSpec_Eks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KubernetesClusterSpec_Eks.Merge(m, src)
}
func (m *KubernetesClusterSpec_Eks) XXX_Size() int {
	return xxx_messageInfo_KubernetesClusterSpec_Eks.Size(m)
}
func (m *KubernetesClusterSpec_Eks) XXX_DiscardUnknown() {
	xxx_messageInfo_KubernetesClusterSpec_Eks.DiscardUnknown(m)
}

var xxx_messageInfo_KubernetesClusterSpec_Eks proto.InternalMessageInfo

func (m *KubernetesClusterSpec_Eks) GetArn() string {
	if m != nil {
		return m.Arn
	}
	return ""
}

func (m *KubernetesClusterSpec_Eks) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *KubernetesClusterSpec_Eks) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *KubernetesClusterSpec_Eks) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type KubernetesClusterStatus struct {
	// List of statuses about the kubernetes cluster.
	// This list allows for multiple applications/pods to record their connection status.
	Status []*v1.Status `protobuf:"bytes,1,rep,name=status,proto3" json:"status,omitempty"`
	// The namespace on in which cluster registration resources are created.
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// JSON representation of the set of PolicyRules to attach to the newly created ClusterRole when registering clusters.
	PolicyRules          []*PolicyRule `protobuf:"bytes,3,rep,name=policy_rules,json=policyRules,proto3" json:"policy_rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *KubernetesClusterStatus) Reset()         { *m = KubernetesClusterStatus{} }
func (m *KubernetesClusterStatus) String() string { return proto.CompactTextString(m) }
func (*KubernetesClusterStatus) ProtoMessage()    {}
func (*KubernetesClusterStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_533baf1d8d7a56a7, []int{1}
}
func (m *KubernetesClusterStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KubernetesClusterStatus.Unmarshal(m, b)
}
func (m *KubernetesClusterStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KubernetesClusterStatus.Marshal(b, m, deterministic)
}
func (m *KubernetesClusterStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KubernetesClusterStatus.Merge(m, src)
}
func (m *KubernetesClusterStatus) XXX_Size() int {
	return xxx_messageInfo_KubernetesClusterStatus.Size(m)
}
func (m *KubernetesClusterStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_KubernetesClusterStatus.DiscardUnknown(m)
}

var xxx_messageInfo_KubernetesClusterStatus proto.InternalMessageInfo

func (m *KubernetesClusterStatus) GetStatus() []*v1.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *KubernetesClusterStatus) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *KubernetesClusterStatus) GetPolicyRules() []*PolicyRule {
	if m != nil {
		return m.PolicyRules
	}
	return nil
}

//
//Copy pasted from the official kubernetes definition:
//https://github.com/kubernetes/api/blob/697df40f2d58d7d48b180b83d7b9b2b5ff812923/rbac/v1alpha1/generated.proto#L98
//PolicyRule holds information that describes a policy rule, but does not contain information
//about who the rule applies to or which namespace the rule applies to.
type PolicyRule struct {
	// Verbs is a list of Verbs that apply to ALL the ResourceKinds and AttributeRestrictions contained in this rule.  VerbAll represents all kinds.
	Verbs []string `protobuf:"bytes,1,rep,name=verbs,proto3" json:"verbs,omitempty"`
	// APIGroups is the name of the APIGroup that contains the resources.  If multiple API groups are specified, any action requested against one of
	// the enumerated resources in any API group will be allowed.
	// +optional
	ApiGroups []string `protobuf:"bytes,2,rep,name=api_groups,json=apiGroups,proto3" json:"api_groups,omitempty"`
	// Resources is a list of resources this rule applies to.  ResourceAll represents all resources.
	// +optional
	Resources []string `protobuf:"bytes,3,rep,name=resources,proto3" json:"resources,omitempty"`
	// ResourceNames is an optional white list of names that the rule applies to.  An empty set means that everything is allowed.
	// +optional
	ResourceNames []string `protobuf:"bytes,4,rep,name=resource_names,json=resourceNames,proto3" json:"resource_names,omitempty"`
	// NonResourceURLs is a set of partial urls that a user should have access to.  *s are allowed, but only as the full, final step in the path
	// Since non-resource URLs are not namespaced, this field is only applicable for ClusterRoles referenced from a ClusterRoleBinding.
	// Rules can either apply to API resources (such as "pods" or "secrets") or non-resource URL paths (such as "/api"),  but not both.
	// +optional
	NonResourceUrls      []string `protobuf:"bytes,5,rep,name=non_resource_urls,json=nonResourceUrls,proto3" json:"non_resource_urls,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PolicyRule) Reset()         { *m = PolicyRule{} }
func (m *PolicyRule) String() string { return proto.CompactTextString(m) }
func (*PolicyRule) ProtoMessage()    {}
func (*PolicyRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_533baf1d8d7a56a7, []int{2}
}
func (m *PolicyRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PolicyRule.Unmarshal(m, b)
}
func (m *PolicyRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PolicyRule.Marshal(b, m, deterministic)
}
func (m *PolicyRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyRule.Merge(m, src)
}
func (m *PolicyRule) XXX_Size() int {
	return xxx_messageInfo_PolicyRule.Size(m)
}
func (m *PolicyRule) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyRule.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyRule proto.InternalMessageInfo

func (m *PolicyRule) GetVerbs() []string {
	if m != nil {
		return m.Verbs
	}
	return nil
}

func (m *PolicyRule) GetApiGroups() []string {
	if m != nil {
		return m.ApiGroups
	}
	return nil
}

func (m *PolicyRule) GetResources() []string {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *PolicyRule) GetResourceNames() []string {
	if m != nil {
		return m.ResourceNames
	}
	return nil
}

func (m *PolicyRule) GetNonResourceUrls() []string {
	if m != nil {
		return m.NonResourceUrls
	}
	return nil
}

func init() {
	proto.RegisterType((*KubernetesClusterSpec)(nil), "multicluster.solo.io.KubernetesClusterSpec")
	proto.RegisterType((*KubernetesClusterSpec_ProviderInfo)(nil), "multicluster.solo.io.KubernetesClusterSpec.ProviderInfo")
	proto.RegisterType((*KubernetesClusterSpec_Eks)(nil), "multicluster.solo.io.KubernetesClusterSpec.Eks")
	proto.RegisterType((*KubernetesClusterStatus)(nil), "multicluster.solo.io.KubernetesClusterStatus")
	proto.RegisterType((*PolicyRule)(nil), "multicluster.solo.io.PolicyRule")
}

func init() {
	proto.RegisterFile("github.com/solo-io/skv2/api/multicluster/v1alpha1/cluster.proto", fileDescriptor_533baf1d8d7a56a7)
}

var fileDescriptor_533baf1d8d7a56a7 = []byte{
	// 537 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xdd, 0x6a, 0x13, 0x41,
	0x14, 0xc7, 0xdd, 0x6e, 0x5a, 0xd8, 0x93, 0xd6, 0x8f, 0x21, 0xea, 0x1a, 0xbf, 0x42, 0x40, 0x08,
	0x82, 0xbb, 0x24, 0xde, 0x78, 0x23, 0x42, 0x63, 0xa9, 0x45, 0x90, 0xb2, 0xe2, 0x8d, 0x20, 0xcb,
	0x64, 0x33, 0xdd, 0x0e, 0xd9, 0xcc, 0x19, 0xe6, 0x23, 0xb4, 0x0f, 0xe1, 0x7b, 0x78, 0x27, 0x78,
	0xe5, 0xf3, 0xf8, 0x0e, 0xde, 0xcb, 0xce, 0x4e, 0xb6, 0x91, 0xc6, 0x0b, 0xaf, 0x72, 0xce, 0xef,
	0xfc, 0x73, 0xe6, 0x3f, 0x67, 0xe7, 0xc0, 0x9b, 0x92, 0x9b, 0x73, 0x3b, 0x4b, 0x0a, 0x5c, 0xa6,
	0x1a, 0x2b, 0x7c, 0xc1, 0x31, 0xd5, 0x8b, 0xd5, 0x24, 0xa5, 0x92, 0xa7, 0x4b, 0x5b, 0x19, 0x5e,
	0x54, 0x56, 0x1b, 0xa6, 0xd2, 0xd5, 0x98, 0x56, 0xf2, 0x9c, 0x8e, 0x53, 0x0f, 0x12, 0xa9, 0xd0,
	0x20, 0xe9, 0x6d, 0x8a, 0x92, 0xba, 0x45, 0xc2, 0xb1, 0xdf, 0x2b, 0xb1, 0x44, 0x27, 0x48, 0xeb,
	0xa8, 0xd1, 0xf6, 0x09, 0xbb, 0x30, 0x0d, 0x64, 0x17, 0xc6, 0xb3, 0x87, 0xed, 0x69, 0x05, 0x2a,
	0x96, 0xae, 0xc6, 0xee, 0xb7, 0x29, 0x0e, 0xbf, 0x86, 0x70, 0xf7, 0xbd, 0x9d, 0x31, 0x25, 0x98,
	0x61, 0x7a, 0xda, 0x1c, 0xf2, 0x51, 0xb2, 0x82, 0x3c, 0x85, 0xae, 0x66, 0x85, 0x62, 0x26, 0x17,
	0x74, 0xc9, 0xe2, 0x60, 0x10, 0x8c, 0xa2, 0x0c, 0x1a, 0xf4, 0x81, 0x2e, 0x19, 0x79, 0x06, 0x37,
	0xbd, 0xa9, 0x7c, 0x8e, 0x4b, 0xca, 0x45, 0xbc, 0xe3, 0x34, 0x07, 0x9e, 0xbe, 0x75, 0x90, 0x7c,
	0x81, 0x03, 0xa9, 0x70, 0xc5, 0xe7, 0x4c, 0xe5, 0x5c, 0x9c, 0x61, 0x1c, 0x0e, 0x82, 0x51, 0x77,
	0xf2, 0x2a, 0xd9, 0x76, 0xad, 0x64, 0xab, 0x97, 0xe4, 0xd4, 0x37, 0x38, 0x11, 0x67, 0x98, 0xed,
	0xcb, 0x8d, 0xac, 0xcf, 0x61, 0x7f, 0xb3, 0x4a, 0xa6, 0x10, 0xb2, 0x85, 0x76, 0x76, 0xbb, 0x93,
	0xf4, 0x7f, 0x0e, 0x39, 0x5a, 0xe8, 0x77, 0x37, 0xb2, 0xfa, 0xdf, 0x87, 0x3d, 0x20, 0x7f, 0x79,
	0xce, 0xcd, 0xa5, 0x64, 0xfd, 0x19, 0x84, 0x47, 0x0b, 0x4d, 0x6e, 0x43, 0x48, 0x95, 0xf0, 0x03,
	0xa9, 0x43, 0xf2, 0x18, 0x80, 0x16, 0x05, 0x5a, 0x61, 0x72, 0x3e, 0xf7, 0x53, 0x88, 0x3c, 0x39,
	0x99, 0x93, 0x7b, 0xb0, 0xa7, 0x58, 0xc9, 0x51, 0xb8, 0xab, 0x47, 0x99, 0xcf, 0x08, 0x81, 0x8e,
	0x1b, 0x6d, 0xc7, 0x51, 0x17, 0x0f, 0xbf, 0x07, 0x70, 0xff, 0xba, 0x3d, 0x43, 0x8d, 0xd5, 0x64,
	0x0c, 0x7b, 0xda, 0x45, 0x71, 0x30, 0x08, 0x47, 0xdd, 0xc9, 0x83, 0xc4, 0x7d, 0xc8, 0xfa, 0xf3,
	0xb6, 0x57, 0x6b, 0xa4, 0x99, 0x17, 0x92, 0x47, 0x10, 0xd5, 0x6d, 0xb5, 0xa4, 0x05, 0x5b, 0x1b,
	0x6b, 0x01, 0x99, 0xc2, 0xbe, 0xc4, 0x8a, 0x17, 0x97, 0xb9, 0xb2, 0x15, 0xd3, 0x71, 0xe8, 0xda,
	0x0e, 0xb6, 0x0f, 0xed, 0xd4, 0x29, 0x33, 0x5b, 0xb1, 0xac, 0x2b, 0xdb, 0x58, 0x0f, 0x7f, 0x04,
	0x00, 0x57, 0x35, 0xd2, 0x83, 0xdd, 0x15, 0x53, 0xb3, 0xc6, 0x63, 0x94, 0x35, 0x89, 0x9b, 0x90,
	0xe4, 0x79, 0xa9, 0xd0, 0x4a, 0x1d, 0xef, 0xb8, 0x52, 0x44, 0x25, 0x3f, 0x76, 0xa0, 0xb6, 0xa9,
	0x98, 0x46, 0xab, 0x0a, 0xef, 0x22, 0xca, 0xae, 0x40, 0xfd, 0xd0, 0xd6, 0x89, 0x7b, 0x8b, 0x3a,
	0xee, 0x38, 0xc9, 0xc1, 0x9a, 0xd6, 0xcf, 0x51, 0x93, 0xe7, 0x70, 0x47, 0xa0, 0xc8, 0x5b, 0xa9,
	0x55, 0x95, 0x8e, 0x77, 0x9d, 0xf2, 0x96, 0x40, 0x91, 0x79, 0xfe, 0x49, 0x55, 0xfa, 0xf0, 0xf8,
	0xe7, 0xef, 0x4e, 0xf0, 0xed, 0xd7, 0x93, 0xe0, 0xf3, 0xeb, 0x7f, 0xad, 0xa7, 0x5c, 0x94, 0xd7,
	0x56, 0x74, 0x3d, 0x8c, 0x76, 0x55, 0x67, 0x7b, 0x6e, 0x8d, 0x5e, 0xfe, 0x09, 0x00, 0x00, 0xff,
	0xff, 0x68, 0xc4, 0x21, 0x30, 0xe6, 0x03, 0x00, 0x00,
}

func (this *KubernetesClusterSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*KubernetesClusterSpec)
	if !ok {
		that2, ok := that.(KubernetesClusterSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.SecretName != that1.SecretName {
		return false
	}
	if this.ClusterDomain != that1.ClusterDomain {
		return false
	}
	if !this.ProviderInfo.Equal(that1.ProviderInfo) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *KubernetesClusterSpec_ProviderInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*KubernetesClusterSpec_ProviderInfo)
	if !ok {
		that2, ok := that.(KubernetesClusterSpec_ProviderInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.ProviderInfoType == nil {
		if this.ProviderInfoType != nil {
			return false
		}
	} else if this.ProviderInfoType == nil {
		return false
	} else if !this.ProviderInfoType.Equal(that1.ProviderInfoType) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *KubernetesClusterSpec_ProviderInfo_Eks) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*KubernetesClusterSpec_ProviderInfo_Eks)
	if !ok {
		that2, ok := that.(KubernetesClusterSpec_ProviderInfo_Eks)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Eks.Equal(that1.Eks) {
		return false
	}
	return true
}
func (this *KubernetesClusterSpec_Eks) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*KubernetesClusterSpec_Eks)
	if !ok {
		that2, ok := that.(KubernetesClusterSpec_Eks)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Arn != that1.Arn {
		return false
	}
	if this.AccountId != that1.AccountId {
		return false
	}
	if this.Region != that1.Region {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *KubernetesClusterStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*KubernetesClusterStatus)
	if !ok {
		that2, ok := that.(KubernetesClusterStatus)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Status) != len(that1.Status) {
		return false
	}
	for i := range this.Status {
		if !this.Status[i].Equal(that1.Status[i]) {
			return false
		}
	}
	if this.Namespace != that1.Namespace {
		return false
	}
	if len(this.PolicyRules) != len(that1.PolicyRules) {
		return false
	}
	for i := range this.PolicyRules {
		if !this.PolicyRules[i].Equal(that1.PolicyRules[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *PolicyRule) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PolicyRule)
	if !ok {
		that2, ok := that.(PolicyRule)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Verbs) != len(that1.Verbs) {
		return false
	}
	for i := range this.Verbs {
		if this.Verbs[i] != that1.Verbs[i] {
			return false
		}
	}
	if len(this.ApiGroups) != len(that1.ApiGroups) {
		return false
	}
	for i := range this.ApiGroups {
		if this.ApiGroups[i] != that1.ApiGroups[i] {
			return false
		}
	}
	if len(this.Resources) != len(that1.Resources) {
		return false
	}
	for i := range this.Resources {
		if this.Resources[i] != that1.Resources[i] {
			return false
		}
	}
	if len(this.ResourceNames) != len(that1.ResourceNames) {
		return false
	}
	for i := range this.ResourceNames {
		if this.ResourceNames[i] != that1.ResourceNames[i] {
			return false
		}
	}
	if len(this.NonResourceUrls) != len(that1.NonResourceUrls) {
		return false
	}
	for i := range this.NonResourceUrls {
		if this.NonResourceUrls[i] != that1.NonResourceUrls[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
