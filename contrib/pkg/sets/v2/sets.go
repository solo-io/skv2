package sets_v2

import (
	"iter"
	"slices"
	"sort"
	"sync"

	sk_sets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ResourceSet[T client.Object] interface {
	// Get the set stored keys
	Keys() sets.Set[string]
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(T) bool) iter.Seq2[int, T]
	// Return the Set as a map of key to resource.
	Map() map[string]T
	// Insert a resource into the set.
	Insert(resource ...T)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(set ResourceSet[T]) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(resource T) bool
	// Delete the key matching the resource
	Delete(resource T)
	// Return the union with the provided set
	Union(set ResourceSet[T]) ResourceSet[T]
	// Return the difference with the provided set
	Difference(set ResourceSet[T]) ResourceSet[T]
	// Return the intersection with the provided set
	Intersection(set ResourceSet[T]) ResourceSet[T]
	// Find the resource with the given ID
	Find(id T) (T, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sk_sets.ResourceSet
	// returns the delta between this and and another ResourceSet[T]
	Delta(newSet ResourceSet[T]) sk_sets.ResourceDelta
	// Clone returns a deep copy of the set
	Clone() ResourceSet[T]
}

// ResourceDelta represents the set of changes between two ResourceSets.
type ResourceDelta[T client.Object] struct {
	// the resources inserted into the set
	Inserted ResourceSet[T]
	// the resources removed from the set
	Removed ResourceSet[T]
}

func (r *ResourceDelta[T]) DeltaV1() sk_sets.ResourceDelta {
	return sk_sets.ResourceDelta{
		Inserted: r.Inserted.Generic(),
		Removed:  r.Removed.Generic(),
	}
}

type resourceSet[T client.Object] struct {
	lock        sync.RWMutex
	set         []T
	sortFunc    func(toInsert, existing client.Object) bool
	compareFunc func(a, b client.Object) int
}

func NewResourceSet[T client.Object](
	sortFunc func(toInsert, existing client.Object) bool,
	compareFunc func(a, b client.Object) int,
	resources ...T,
) ResourceSet[T] {
	rs := &resourceSet[T]{
		set:         []T{},
		sortFunc:    sortFunc,
		compareFunc: compareFunc,
	}
	rs.Insert(resources...)
	return rs
}

func (s *resourceSet[T]) Keys() sets.Set[string] {
	s.lock.RLock()
	defer s.lock.RUnlock()
	keys := sets.Set[string]{}
	for _, resource := range s.set {
		keys.Insert(sk_sets.Key(resource))
	}
	return sets.Set[string]{}
}

func (s *resourceSet[T]) List(filterResource ...func(T) bool) iter.Seq2[int, T] {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return func(yield func(int, T) bool) {
	OUTER:
		for i, resource := range s.set {
			for _, filter := range filterResource {
				if filter(resource) {
					continue OUTER
				}
			}
			if !yield(i, resource) {
				break
			}
		}
	}
}

func (s *resourceSet[T]) Map() map[string]T {
	s.lock.RLock()
	defer s.lock.RUnlock()
	newMap := map[string]T{}
	for _, resource := range s.set {
		newMap[sk_sets.Key(resource)] = resource
	}
	return newMap
}

// Insert adds items to the set.
// If an item is already in the set, it is overwritten.
// The set is sorted based on the sortFunc. If sortFunc is nil, the set will be unsorted.
func (s *resourceSet[T]) Insert(resources ...T) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	for _, resource := range resources {
		insertIndex := sort.Search(len(s.set), func(i int) bool { return s.sortFunc(resource, s.set[i]) })

		// if the resource is already in the set, replace it
		if insertIndex < len(s.set) && s.compareFunc(resource, s.set[insertIndex]) == 0 {
			s.set[insertIndex] = resource
			return
		}
		if s.sortFunc == nil {
			s.set = append(s.set, resource)
			return
		}

		// insert the resource at the determined index
		s.set = slices.Insert(s.set, insertIndex, resource)
	}
}

func (s *resourceSet[T]) Has(resource T) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	i := sort.Search(len(s.set), func(i int) bool {
		return ezkube.CompareResourceIds(s.set[i], resource) >= 0
	})
	return i < len(s.set) && s.compareFunc(s.set[i], resource) == 0
}

func (s *resourceSet[T]) Equal(
	set ResourceSet[T],
) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.Generic().Equal(set.Generic())
}

func (s *resourceSet[T]) Delete(resource T) {
	s.lock.Lock()
	defer s.lock.Unlock()
	i := sort.Search(len(s.set), func(i int) bool {
		return s.compareFunc(s.set[i], resource) >= 0
	})
	if found := i < len(s.set) && s.compareFunc(s.set[i], resource) == 0; !found {
		return
	}

	s.set = slices.Delete(s.set, i, i+1)
}

func (s *resourceSet[T]) Union(set ResourceSet[T]) ResourceSet[T] {
	list := []T{}
	for _, resource := range s.Generic().Union(set.Generic()).List() {
		list = append(list, resource.(T))
	}
	return NewResourceSet[T](
		s.sortFunc,
		s.compareFunc,
		list...,
	)
}

func (s *resourceSet[T]) Difference(set ResourceSet[T]) ResourceSet[T] {
	s.lock.RLock()
	defer s.lock.RUnlock()
	result := NewResourceSet[T](s.sortFunc, s.compareFunc)
	for _, resource := range s.set {
		if !set.Has(resource) {
			result.Insert(resource)
		}
	}
	return result
}

func (s *resourceSet[T]) Intersection(set ResourceSet[T]) ResourceSet[T] {
	s.lock.RLock()
	defer s.lock.RUnlock()
	var walk, other ResourceSet[T]
	result := NewResourceSet[T](s.sortFunc, s.compareFunc)
	if len(s.set) < set.Length() {
		walk = NewResourceSet(s.sortFunc, s.compareFunc, s.set...)
		other = set
	} else {
		walk = set
		other = NewResourceSet(s.sortFunc, s.compareFunc, s.set...)
	}
	for _, key := range walk.List() {
		if other.Has(key) {
			result.Insert(key)
		}
	}
	return result
}

func (s *resourceSet[T]) Find(id T) (T, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	key := sk_sets.Key(id)
	i := sort.Search(len(s.set), func(i int) bool {
		return key <= sk_sets.Key(s.set[i])
	})
	var resource T
	if i < len(s.set) {
		resource = s.set[i]
	}
	if i != len(s.set) && sk_sets.Key(resource) == key {
		return resource, nil
	}

	return resource, sk_sets.NotFoundErr(resource, id)
}

func (s *resourceSet[T]) Length() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.set)
}

// note that this function will currently panic if called for a ResourceSet[T] containing non-runtime.Objects
func (oldSet *resourceSet[T]) Delta(newSet ResourceSet[T]) sk_sets.ResourceDelta {
	updated, removed := NewResourceSet[T](oldSet.sortFunc, oldSet.compareFunc), NewResourceSet[T](oldSet.sortFunc, oldSet.compareFunc)

	// find objects updated or removed
	for _, resource := range oldSet.set {
		newObj, err := newSet.Find(resource)
		switch {
		case err != nil:
			// obj removed
			removed.Insert(resource)
		case !controllerutils.ObjectsEqual(resource, newObj):
			// obj updated
			updated.Insert(newObj)
		default:
			// obj the same
		}
	}

	// find objects added
	for _, newObj := range newSet.Generic().List() {
		if _, err := oldSet.Find(newObj.(T)); err != nil {
			// obj added
			updated.Insert(newObj.(T))
		}
	}

	delta := &ResourceDelta[T]{
		Inserted: updated,
		Removed:  removed,
	}
	return delta.DeltaV1()
}

// Create a clone of the current set
// note that this function will currently panic if called for a ResourceSet[T] containing non-runtime.Objects
func (oldSet *resourceSet[T]) Clone() ResourceSet[T] {
	new := NewResourceSet[T](oldSet.sortFunc, oldSet.compareFunc)
	oldSet.List(func(oldObj T) bool {
		copy := oldObj.DeepCopyObject().(T)
		new.Insert(copy)
		return true
	})
	return new
}

func (s *resourceSet[T]) Generic() sk_sets.ResourceSet {
	genericSortFunc := func(toInsert, existing interface{}) bool {
		return s.sortFunc(toInsert.(T), existing.(T))
	}
	genericCompareFunc := func(a, b interface{}) int {
		return s.compareFunc(a.(T), b.(T))
	}
	set := sk_sets.NewResourceSet(genericSortFunc, genericCompareFunc)
	for _, v := range s.List() {
		set.Insert(v)
	}
	return set
}
