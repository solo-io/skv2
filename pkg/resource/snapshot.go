package resource

import (
	"golang.org/x/exp/maps"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type GVKSelectorFunc = func(GVK schema.GroupVersionKind) bool

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

func (s Snapshot) Clear() {
	// Clear each map
	for gvk := range s {
		maps.Clear(s[gvk])
	}
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

func (s Snapshot) Clone(selectors ...GVKSelectorFunc) Snapshot {
	clone := Snapshot{}
	for k, v := range s {
		if len(selectors) == 0 {
			clone[k] = copyNnsMap(v)
			continue
		}
		selected := false
		for _, selector := range selectors {
			if selector(k) {
				selected = true
				break
			}
		}
		if selected {
			clone[k] = copyNnsMap(v)
			continue
		}
	}
	return clone
}

// Merges the Snapshot with a Snapshot passed in as an argument. The values
// in the passed in Snapshot will take precedence when there is an object mapped
// to the same gvk and name in both Snapshots.
func (s Snapshot) Merge(toMerge Snapshot) Snapshot {
	merged := s.Clone()
	for gvk, objectsMap := range toMerge {
		if _, ok := merged[gvk]; ok {
			for name, object := range objectsMap {
				// If there is already an object specified here, the object from toMerge
				// will replace it
				merged[gvk][name] = object
			}
		} else {
			merged[gvk] = objectsMap
		}
	}
	return merged
}

// ClusterSnapshot represents a set of snapshots partitioned by cluster
type ClusterSnapshot map[string]Snapshot

func (s ClusterSnapshot) ForEachObject(
	handleObject func(
		cluster string,
		gvk schema.GroupVersionKind,
		obj TypedObject,
	),
) {
	if s == nil {
		return
	}
	for cluster, snap := range s {
		snap.ForEachObject(func(gvk schema.GroupVersionKind, obj TypedObject) {
			handleObject(cluster, gvk, obj)
		})
	}
}

func copyNnsMap(m map[types.NamespacedName]TypedObject) map[types.NamespacedName]TypedObject {
	nnsMapCopy := map[types.NamespacedName]TypedObject{}
	for k, v := range m {
		nnsMapCopy[k] = v.DeepCopyObject().(TypedObject)
	}
	return nnsMapCopy
}

func (cs ClusterSnapshot) Insert(cluster string, gvk schema.GroupVersionKind, obj TypedObject) {
	snapshot, ok := cs[cluster]
	if !ok {
		snapshot = Snapshot{}
	}
	snapshot.Insert(gvk, obj)
	cs[cluster] = snapshot
}

func (cs ClusterSnapshot) Clear() {
	for _, snapshot := range cs {
		snapshot.Clear()
	}
}

func (cs ClusterSnapshot) Delete(
	cluster string,
	gvk schema.GroupVersionKind,
	id types.NamespacedName,
) {
	snapshot, ok := cs[cluster]
	if !ok {
		return
	}
	snapshot.Delete(gvk, id)
	cs[cluster] = snapshot
}

func (cs ClusterSnapshot) Clone(selectors ...GVKSelectorFunc) ClusterSnapshot {
	clone := ClusterSnapshot{}
	for k, v := range cs {
		clone[k] = v.Clone(selectors...)
	}
	return clone
}

// Merges the ClusterSnapshot with a ClusterSnapshot passed in as an argument.
// If a cluster exists in both ClusterSnapshots, then both Snapshots for the
// cluster is merged; with the passed in ClusterSnapshot's corresponding Snapshot
// taking precedence in case of conflicts.
func (cs ClusterSnapshot) Merge(toMerge ClusterSnapshot) ClusterSnapshot {
	merged := cs.Clone()
	for cluster, snapshot := range toMerge {
		if baseSnap, ok := merged[cluster]; ok {
			merged[cluster] = baseSnap.Merge(snapshot)
		} else {
			merged[cluster] = snapshot
		}
	}
	return merged
}
