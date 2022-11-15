package ezkube

import "sigs.k8s.io/controller-runtime/pkg/client"

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
