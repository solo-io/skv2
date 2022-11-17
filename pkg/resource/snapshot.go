package resource

import (
	"golang.org/x/exp/maps"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type GVKSelectorFunc = func(GVK schema.GroupVersionKind) bool

// Deprecated: TypedObject is not needed. use `GetObjectKind().SetGroupVersionKind` instead.
type TypedObject = client.Object

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
	return s.cloneInternal(true, selectors...)
}

func (s Snapshot) ShallowCopy(selectors ...GVKSelectorFunc) Snapshot {
	return s.cloneInternal(false, selectors...)
}

func (s Snapshot) cloneInternal(deepCopy bool, selectors ...GVKSelectorFunc) Snapshot {
	clone := Snapshot{}
	for k, v := range s {
		if len(selectors) == 0 {
			if deepCopy {
				clone[k] = copyNnsMap(v)
			} else {
				clone[k] = map[types.NamespacedName]client.Object{}
				maps.Copy(clone[k], v)
			}
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
			if deepCopy {
				clone[k] = copyNnsMap(v)
			} else {
				clone[k] = map[types.NamespacedName]client.Object{}
				maps.Copy(clone[k], v)
			}
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
		obj client.Object,
	),
) {
	if s == nil {
		return
	}
	for cluster, snap := range s {
		snap.ForEachObject(func(gvk schema.GroupVersionKind, obj client.Object) {
			handleObject(cluster, gvk, obj)
		})
	}
}

func copyNnsMap(m map[types.NamespacedName]client.Object) map[types.NamespacedName]client.Object {
	nnsMapCopy := map[types.NamespacedName]client.Object{}
	for k, v := range m {
		nnsMapCopy[k] = v.DeepCopyObject().(client.Object)
	}
	return nnsMapCopy
}

func (cs ClusterSnapshot) Insert(cluster string, gvk schema.GroupVersionKind, obj client.Object) {
	snapshot, ok := cs[cluster]
	if !ok {
		snapshot = Snapshot{}
	}
	snapshot.Insert(gvk, obj)
	cs[cluster] = snapshot
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

func (cs ClusterSnapshot) ShallowCopy(selectors ...GVKSelectorFunc) ClusterSnapshot {
	clone := ClusterSnapshot{}
	for k, v := range cs {
		clone[k] = v.ShallowCopy(selectors...)
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
