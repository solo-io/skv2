package resource

import "sigs.k8s.io/controller-runtime/pkg/client"

// provides a simple, extendable interface for dealing with resources
type Resource interface {
	GetName() string
	GetNamespace() string
}

func ToClientKey(res Resource) client.ObjectKey {
	return client.ObjectKey{
		Namespace: res.GetNamespace(),
		Name:      res.GetName(),
	}
}
