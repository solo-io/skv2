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

type ClusterResourceId interface {
	ResourceId
	GetClusterName() string
}

func GetClusterName(obj client.Object) string {
	if obj.GetAnnotations() == nil {
		return ""
	}
	return obj.GetAnnotations()[ClusterAnnotation]
}

func SetClusterName(obj client.Object, cluster string) {
	if obj.GetAnnotations() == nil {
		obj.SetAnnotations(map[string]string{})
	}
	obj.GetAnnotations()[ClusterAnnotation] = cluster
}
