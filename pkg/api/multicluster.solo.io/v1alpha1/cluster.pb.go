// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: github.com/solo-io/skv2/api/multicluster/v1alpha1/cluster.proto

package v1alpha1

import (
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
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

//
//Representation of a Kubernetes cluster that has been registered.
type KubernetesClusterSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the secret which contains the kubeconfig with information to connect to the remote cluster.
	SecretName string `protobuf:"bytes,1,opt,name=secret_name,json=secretName,proto3" json:"secret_name,omitempty"`
	// name local DNS suffix used by the cluster.
	// used for building FQDNs for in-cluster services
	// defaults to 'cluster.local'
	ClusterDomain string `protobuf:"bytes,2,opt,name=cluster_domain,json=clusterDomain,proto3" json:"cluster_domain,omitempty"`
	// Metadata for clusters provisioned from cloud providers.
	ProviderInfo *KubernetesClusterSpec_ProviderInfo `protobuf:"bytes,3,opt,name=provider_info,json=providerInfo,proto3" json:"provider_info,omitempty"`
}

func (x *KubernetesClusterSpec) Reset() {
	*x = KubernetesClusterSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesClusterSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesClusterSpec) ProtoMessage() {}

func (x *KubernetesClusterSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesClusterSpec.ProtoReflect.Descriptor instead.
func (*KubernetesClusterSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *KubernetesClusterSpec) GetSecretName() string {
	if x != nil {
		return x.SecretName
	}
	return ""
}

func (x *KubernetesClusterSpec) GetClusterDomain() string {
	if x != nil {
		return x.ClusterDomain
	}
	return ""
}

func (x *KubernetesClusterSpec) GetProviderInfo() *KubernetesClusterSpec_ProviderInfo {
	if x != nil {
		return x.ProviderInfo
	}
	return nil
}

type KubernetesClusterStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of statuses about the kubernetes cluster.
	// This list allows for multiple applications/pods to record their connection status.
	Status []*v1.Status `protobuf:"bytes,1,rep,name=status,proto3" json:"status,omitempty"`
	// The namespace in which cluster registration resources were created.
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// The set of PolicyRules attached to ClusterRoles when this cluster was registered.
	PolicyRules []*PolicyRule `protobuf:"bytes,3,rep,name=policy_rules,json=policyRules,proto3" json:"policy_rules,omitempty"`
}

func (x *KubernetesClusterStatus) Reset() {
	*x = KubernetesClusterStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesClusterStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesClusterStatus) ProtoMessage() {}

func (x *KubernetesClusterStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesClusterStatus.ProtoReflect.Descriptor instead.
func (*KubernetesClusterStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *KubernetesClusterStatus) GetStatus() []*v1.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *KubernetesClusterStatus) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *KubernetesClusterStatus) GetPolicyRules() []*PolicyRule {
	if x != nil {
		return x.PolicyRules
	}
	return nil
}

//
//Copy pasted from the official kubernetes definition:
//https://github.com/kubernetes/api/blob/697df40f2d58d7d48b180b83d7b9b2b5ff812923/rbac/v1alpha1/generated.proto#L98
//PolicyRule holds information that describes a policy rule, but does not contain information
//about who the rule applies to or which namespace the rule applies to.
type PolicyRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

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
	NonResourceUrls []string `protobuf:"bytes,5,rep,name=non_resource_urls,json=nonResourceUrls,proto3" json:"non_resource_urls,omitempty"`
}

func (x *PolicyRule) Reset() {
	*x = PolicyRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicyRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyRule) ProtoMessage() {}

func (x *PolicyRule) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyRule.ProtoReflect.Descriptor instead.
func (*PolicyRule) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_rawDescGZIP(), []int{2}
}

func (x *PolicyRule) GetVerbs() []string {
	if x != nil {
		return x.Verbs
	}
	return nil
}

func (x *PolicyRule) GetApiGroups() []string {
	if x != nil {
		return x.ApiGroups
	}
	return nil
}

func (x *PolicyRule) GetResources() []string {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *PolicyRule) GetResourceNames() []string {
	if x != nil {
		return x.ResourceNames
	}
	return nil
}

func (x *PolicyRule) GetNonResourceUrls() []string {
	if x != nil {
		return x.NonResourceUrls
	}
	return nil
}

// Metadata for clusters provisioned from cloud providers.
type KubernetesClusterSpec_ProviderInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Metadata specific to a cloud provider.
	//
	// Types that are assignable to ProviderInfoType:
	//	*KubernetesClusterSpec_ProviderInfo_Eks
	ProviderInfoType isKubernetesClusterSpec_ProviderInfo_ProviderInfoType `protobuf_oneof:"provider_info_type"`
}

func (x *KubernetesClusterSpec_ProviderInfo) Reset() {
	*x = KubernetesClusterSpec_ProviderInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesClusterSpec_ProviderInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesClusterSpec_ProviderInfo) ProtoMessage() {}

func (x *KubernetesClusterSpec_ProviderInfo) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesClusterSpec_ProviderInfo.ProtoReflect.Descriptor instead.
func (*KubernetesClusterSpec_ProviderInfo) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_rawDescGZIP(), []int{0, 0}
}

func (m *KubernetesClusterSpec_ProviderInfo) GetProviderInfoType() isKubernetesClusterSpec_ProviderInfo_ProviderInfoType {
	if m != nil {
		return m.ProviderInfoType
	}
	return nil
}

func (x *KubernetesClusterSpec_ProviderInfo) GetEks() *KubernetesClusterSpec_Eks {
	if x, ok := x.GetProviderInfoType().(*KubernetesClusterSpec_ProviderInfo_Eks); ok {
		return x.Eks
	}
	return nil
}

type isKubernetesClusterSpec_ProviderInfo_ProviderInfoType interface {
	isKubernetesClusterSpec_ProviderInfo_ProviderInfoType()
}

type KubernetesClusterSpec_ProviderInfo_Eks struct {
	// Provider info for an AWS EKS provisioned cluster.
	Eks *KubernetesClusterSpec_Eks `protobuf:"bytes,1,opt,name=eks,proto3,oneof"`
}

func (*KubernetesClusterSpec_ProviderInfo_Eks) isKubernetesClusterSpec_ProviderInfo_ProviderInfoType() {
}

// AWS metadata associated with an EKS provisioned cluster.
type KubernetesClusterSpec_Eks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// AWS ARN.
	Arn string `protobuf:"bytes,1,opt,name=arn,proto3" json:"arn,omitempty"`
	// AWS 12 digit account ID.
	AccountId string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// AWS region.
	Region string `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	// EKS resource name.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *KubernetesClusterSpec_Eks) Reset() {
	*x = KubernetesClusterSpec_Eks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesClusterSpec_Eks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesClusterSpec_Eks) ProtoMessage() {}

func (x *KubernetesClusterSpec_Eks) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesClusterSpec_Eks.ProtoReflect.Descriptor instead.
func (*KubernetesClusterSpec_Eks) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_rawDescGZIP(), []int{0, 1}
}

func (x *KubernetesClusterSpec_Eks) GetArn() string {
	if x != nil {
		return x.Arn
	}
	return ""
}

func (x *KubernetesClusterSpec_Eks) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *KubernetesClusterSpec_Eks) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *KubernetesClusterSpec_Eks) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto protoreflect.FileDescriptor

var file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6b, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x14, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f,
	0x73, 0x6b, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x03, 0x0a, 0x15,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x5d, 0x0a,
	0x0d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x70, 0x65,
	0x63, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x69, 0x0a, 0x0c,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x43, 0x0a, 0x03,
	0x65, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x45, 0x6b, 0x73, 0x48, 0x00, 0x52, 0x03, 0x65, 0x6b,
	0x73, 0x42, 0x14, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x62, 0x0a, 0x03, 0x45, 0x6b, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x61, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x72, 0x6e,
	0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xaf, 0x01, 0x0a, 0x17,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73,
	0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x75, 0x6c, 0x65,
	0x52, 0x0b, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x22, 0xb2, 0x01,
	0x0a, 0x0a, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x65, 0x72, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x76, 0x65, 0x72,
	0x62, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x69, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12,
	0x25, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x6e, 0x6f, 0x6e, 0x5f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0f, 0x6e, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x72,
	0x6c, 0x73, 0x42, 0x4b, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6b, 0x76, 0x32, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_rawDescOnce sync.Once
	file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_rawDescData = file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_rawDesc
)

func file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_rawDescData)
	})
	return file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_rawDescData
}

var file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_goTypes = []interface{}{
	(*KubernetesClusterSpec)(nil),              // 0: multicluster.solo.io.KubernetesClusterSpec
	(*KubernetesClusterStatus)(nil),            // 1: multicluster.solo.io.KubernetesClusterStatus
	(*PolicyRule)(nil),                         // 2: multicluster.solo.io.PolicyRule
	(*KubernetesClusterSpec_ProviderInfo)(nil), // 3: multicluster.solo.io.KubernetesClusterSpec.ProviderInfo
	(*KubernetesClusterSpec_Eks)(nil),          // 4: multicluster.solo.io.KubernetesClusterSpec.Eks
	(*v1.Status)(nil),                          // 5: core.skv2.solo.io.Status
}
var file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_depIdxs = []int32{
	3, // 0: multicluster.solo.io.KubernetesClusterSpec.provider_info:type_name -> multicluster.solo.io.KubernetesClusterSpec.ProviderInfo
	5, // 1: multicluster.solo.io.KubernetesClusterStatus.status:type_name -> core.skv2.solo.io.Status
	2, // 2: multicluster.solo.io.KubernetesClusterStatus.policy_rules:type_name -> multicluster.solo.io.PolicyRule
	4, // 3: multicluster.solo.io.KubernetesClusterSpec.ProviderInfo.eks:type_name -> multicluster.solo.io.KubernetesClusterSpec.Eks
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_init() }
func file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_init() {
	if File_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesClusterSpec); i {
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
		file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesClusterStatus); i {
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
		file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolicyRule); i {
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
		file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesClusterSpec_ProviderInfo); i {
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
		file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesClusterSpec_Eks); i {
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
	file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*KubernetesClusterSpec_ProviderInfo_Eks)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto = out.File
	file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_rawDesc = nil
	file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_goTypes = nil
	file_github_com_solo_io_skv2_api_multicluster_v1alpha1_cluster_proto_depIdxs = nil
}
