package v2

import (
	"slices"
	"strings"

	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/resource"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type GVKSelectorFunc = resource.GVKSelectorFunc

// Snapshot represents a generic snapshot of client.Objects scoped to a single cluster
type Snapshot map[schema.GroupVersionKind][]client.Object

func (s Snapshot) Insert(gvk schema.GroupVersionKind, obj client.Object) {
	objects, ok := s[gvk]
	if !ok {
		objects = []client.Object{}
	}
	insertIndex, found := slices.BinarySearchFunc(
		objects,
		obj,
		func(a, b client.Object) int { return ezkube.ResourceIdsCompare(a, b) },
	)
	if found {
		objects[insertIndex] = obj
	} else {
		objects = slices.Insert(objects, insertIndex, obj)
	}
	s[gvk] = objects
}

func (s Snapshot) Delete(gvk schema.GroupVersionKind, id types.NamespacedName) {
	resources, ok := s[gvk]
	if !ok {
		return
	}

	i, found := slices.BinarySearchFunc(
		resources,
		id,
		func(a client.Object, b types.NamespacedName) int {
			// compare names
			if cmp := strings.Compare(a.GetName(), b.Name); cmp != 0 {
				return cmp
			}

			// compare namespaces
			return strings.Compare(a.GetNamespace(), b.Namespace)
		},
	)
	if found {
		resources = slices.Delete(resources, i, i+1)
	}
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
				clone[k] = copyNnsSlice(v)
			} else {
				clone[k] = make([]client.Object, len(v))
				copy(clone[k], v)
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
				clone[k] = copyNnsSlice(v)
			} else {
				clone[k] = make([]client.Object, len(v))
				copy(clone[k], v)
			}
			continue
		}
	}
	return clone
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

func copyNnsSlice(m []client.Object) []client.Object {
	nsSliceCopy := make([]client.Object, len(m))
	for k, v := range m {
		nsSliceCopy[k] = v.DeepCopyObject().(client.Object)
	}
	return nsSliceCopy
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
