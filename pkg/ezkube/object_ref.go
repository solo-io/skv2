package ezkube

import (
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func MakeObjectRef(resource ResourceId) *v1.ObjectRef {
	if resource == nil {
		return nil
	}
	return &v1.ObjectRef{
		Name:      resource.GetName(),
		Namespace: resource.GetNamespace(),
	}
}

func MakeClusterObjectRef(resource ClusterResourceId) *v1.ClusterObjectRef {
	if resource == nil {
		return nil
	}
	return &v1.ClusterObjectRef{
		Name:        resource.GetName(),
		Namespace:   resource.GetNamespace(),
		ClusterName: GetClusterName(resource),
	}
}

func RefsMatch(ref1, ref2 ResourceId) bool {
	// If either ref is nil, return true if they are both nil
	if ref1 == nil || ref2 == nil {
		return ref1 == nil && ref2 == nil
	}
	return ref1.GetNamespace() == ref2.GetNamespace() &&
		ref1.GetName() == ref2.GetName()
}

func ClusterRefsMatch(ref1, ref2 ClusterResourceId) bool {
	// If either ref is nil, return true if they are both nil
	if ref1 == nil || ref2 == nil {
		return ref1 == nil && ref2 == nil
	}
	return ref1.GetNamespace() == ref2.GetNamespace() &&
		ref1.GetName() == ref2.GetName() &&
		GetClusterName(ref1) == GetClusterName(ref2)
}

func MakeClientObjectKey(ref ResourceId) client.ObjectKey {
	if ref == nil {
		return client.ObjectKey{}
	}
	return client.ObjectKey{
		Namespace: ref.GetNamespace(),
		Name:      ref.GetName(),
	}
}
