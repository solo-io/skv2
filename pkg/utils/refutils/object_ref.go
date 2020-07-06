package refutils

import (
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	"github.com/solo-io/skv2/pkg/ezkube"
)

func MakeObjectRef(resource ezkube.ResourceId) *v1.ObjectRef {
	return &v1.ObjectRef{
		Name:      resource.GetName(),
		Namespace: resource.GetNamespace(),
	}
}

func MakeClusterObjectRef(resource ezkube.ClusterResourceId) *v1.ClusterObjectRef {
	return &v1.ClusterObjectRef{
		Name:        resource.GetName(),
		Namespace:   resource.GetNamespace(),
		ClusterName: resource.GetClusterName(),
	}
}
