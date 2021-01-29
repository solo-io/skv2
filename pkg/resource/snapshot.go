package resource

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Snapshot represents a generic snapshot of client.Objects scoped to a single cluster
type Snapshot map[schema.GroupVersionKind]map[types.NamespacedName]client.Object

func (s Snapshot) Insert(gvk schema.GroupVersionKind, obj client.Object) {
	objects, ok := s[gvk]
	if !ok {
		objects = map[types.NamespacedName]client.Object{}
	}
	objects[types.NamespacedName{
		Namespace: obj.GetNamespace(),
		Name:      obj.GetName(),
	}] = obj
	s[gvk] = objects
}

func (s Snapshot) Delete(gvk schema.GroupVersionKind, id types.NamespacedName) {
	resources, ok := s[gvk]
	if !ok {
		return
	}
	delete(resources, id)
	s[gvk] = resources
}

func (s Snapshot) ForEachObject(handleObject func(gvk schema.GroupVersionKind, obj client.Object)) {
	for gvk, objs := range s {
		for _, obj := range objs {
			handleObject(gvk, obj)
		}
	}
}

func (s ClusterSnapshot) ForEachObject(handleObject func(cluster string, gvk schema.GroupVersionKind, obj client.Object)) {
	for cluster, snap := range s {
		snap.ForEachObject(func(gvk schema.GroupVersionKind, obj client.Object) {
			handleObject(cluster, gvk, obj)
		})
	}
}

func (cs ClusterSnapshot) Insert(cluster string, gvk schema.GroupVersionKind, obj client.Object) {
	snapshot, ok := cs[cluster]
	if !ok {
		snapshot = Snapshot{}
	}
	snapshot.Insert(gvk, obj)
	cs[cluster] = snapshot
}

func (cs ClusterSnapshot) Delete(cluster string, gvk schema.GroupVersionKind, id types.NamespacedName) {
	snapshot, ok := cs[cluster]
	if !ok {
		snapshot = Snapshot{}
	}
	snapshot.Delete(gvk, id)
	cs[cluster] = snapshot
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
