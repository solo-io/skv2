package ezkube

import (
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const ClusterAnnotation = "cluster.solo.io/cluster"

// ResourceId represents a global identifier for a k8s resource.
type ResourceId interface {
	GetName() string
	GetNamespace() string
}

// ClusterResourceId represents a global identifier for a k8s resource.
type ClusterResourceId interface {
	GetName() string
	GetNamespace() string
	GetAnnotations() map[string]string
}

// internal struct needed to create helper func that converts ref to struct that satisfies ClusterResourceId interface
type clusterResourceId struct {
	Name, Namespace string
	Annotations     map[string]string
}

func (c clusterResourceId) GetName() string {
	return c.Name
}

func (c clusterResourceId) GetNamespace() string {
	return c.Namespace
}

func (c clusterResourceId) GetAnnotations() map[string]string {
	return c.Annotations
}

type deprecatedClusterResourceId interface {
	GetName() string
	GetNamespace() string
	GetClusterName() string
}

// ConvertRefToId converts a ClusterObjectRef to a struct that implements the ClusterResourceId interface
func ConvertRefToId(ref deprecatedClusterResourceId) ClusterResourceId {
	return clusterResourceId{
		Name:        ref.GetName(),
		Namespace:   ref.GetNamespace(),
		Annotations: map[string]string{ClusterAnnotation: ref.GetClusterName()},
	}
}

func GetDeprecatedClusterName(id ResourceId) string {
	if id, ok := id.(deprecatedClusterResourceId); ok {
		return id.GetClusterName()
	}
	return ""
}

func GetClusterName(id ClusterResourceId) string {
	if id.GetAnnotations() == nil {
		return ""
	}
	return id.GetAnnotations()[ClusterAnnotation]
}

func SetClusterName(obj client.Object, cluster string) {
	if obj.GetAnnotations() == nil {
		obj.SetAnnotations(map[string]string{})
	}
	obj.GetAnnotations()[ClusterAnnotation] = cluster
}
