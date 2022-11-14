package ezkube

const ClusterAnnotation = "cluster.solo.io/cluster"

// ResourceId represents a global identifier for a k8s resource.
type ResourceId interface {
	GetName() string
	GetNamespace() string
}

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
