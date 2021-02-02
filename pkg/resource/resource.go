package resource

import (
	"github.com/solo-io/skv2/pkg/ezkube"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func ToClientKey(res ezkube.ResourceId) client.ObjectKey {
	return client.ObjectKey{
		Namespace: res.GetNamespace(),
		Name:      res.GetName(),
	}
}
