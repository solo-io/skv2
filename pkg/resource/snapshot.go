package resource

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// a typed object is a client.Object with a TypeMeta
type TypedObject interface {
	client.Object
	SetGroupVersionKind(gvk schema.GroupVersionKind)
}

// Snapshot represents a generic snapshot of client.Objects scoped to a single cluster
type Snapshot map[schema.GroupVersionKind]map[types.NamespacedName]TypedObject

func (s Snapshot) Insert(gvk schema.GroupVersionKind, obj TypedObject) {
	objects, ok := s[gvk]
	if !ok {
		objects = map[types.NamespacedName]TypedObject{}
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

func (s Snapshot) ForEachObject(handleObject func(gvk schema.GroupVersionKind, obj TypedObject)) {
	if s == nil {
		return
	}
	for gvk, objs := range s {
		for _, obj := range objs {
			handleObject(gvk, obj)
		}
	}
}

func (s ClusterSnapshot) ForEachObject(handleObject func(cluster string, gvk schema.GroupVersionKind, obj TypedObject)) {
	if s == nil {
		return
	}
	for cluster, snap := range s {
		snap.ForEachObject(func(gvk schema.GroupVersionKind, obj TypedObject) {
			handleObject(cluster, gvk, obj)
		})
	}
}

func (cs ClusterSnapshot) Insert(cluster string, gvk schema.GroupVersionKind, obj TypedObject) {
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

func copyNnsMap(m map[types.NamespacedName]TypedObject) map[types.NamespacedName]TypedObject {
	nnsMapCopy := map[types.NamespacedName]TypedObject{}
	for k, v := range m {
		nnsMapCopy[k] = v
	}
	return nnsMapCopy
}
