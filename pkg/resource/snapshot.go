package resource

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Snapshot represents a generic snapshot of client.Objects scoped to a single cluster
type Snapshot map[schema.GroupVersionKind]map[types.NamespacedName]client.Object

func (s Snapshot) Insert(gvk schema.GroupVersionKind, ref types.NamespacedName, obj client.Object) {
	objects, ok := s[gvk]
	if !ok {
		objects = map[types.NamespacedName]client.Object{}
	}
	objects[ref] = obj
	s[gvk] = objects
}

// ClusterSnapshot represents a set of snapshots partitioned by cluster
type ClusterSnapshot map[string]Snapshot

func (s Snapshot) Clone() Snapshot {
	clone := Snapshot{}
	for k, v := range s {
		clone[k] = copyNnsMap(v)
	}
	return clone
}

func copyNnsMap(m map[types.NamespacedName]client.Object) map[types.NamespacedName]client.Object {
	nnsMapCopy := map[types.NamespacedName]client.Object{}
	for k, v := range m {
		nnsMapCopy[k] = v
	}
	return nnsMapCopy
}
