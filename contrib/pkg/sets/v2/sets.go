package sets_v2

import (
	"slices"
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
	// List returns an iterator for the set.
	// Pass an optional filter function to skip iteration on specific entries; Note: index will still progress.
	List(filterResource ...func(T) bool) func(yield func(int, T) bool)
	// Iterate over the set, passing the index and resource to the provided function.
	// The iteration can be stopped by returning false from the function.
	// Returning true will continue the iteration.
	Iter(func(int, T) bool)
	// Return the Set as a map of key to resource.
	Map() map[string]T
	// Insert a resource into the set.
	Insert(resource ...T)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(set ResourceSet[T]) bool
	// Check if the set contains the resource.
	Has(resource T) bool
	// Delete the matching resource.
	Delete(resource ezkube.ResourceId)
	// Return the union with the provided set
	Union(set ResourceSet[T]) ResourceSet[T]
	// Return the difference with the provided set
	Difference(set ResourceSet[T]) ResourceSet[T]
	// Return the intersection with the provided set
	Intersection(set ResourceSet[T]) ResourceSet[T]
	// Find the resource with the given ID
	Find(resource ezkube.ResourceId) (T, error)
	// Get the length of the set
	Len() int
	// returns the generic implementation of the set
	Generic() sk_sets.ResourceSet
	// returns the delta between this and and another ResourceSet[T]
	Delta(newSet ResourceSet[T]) sk_sets.ResourceDelta
	// Clone returns a deep copy of the set
	Clone() ResourceSet[T]
	// Get the compare function used by the set
	GetCompareFunc() func(a, b ezkube.ResourceId) int
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
	compareFunc func(a, b ezkube.ResourceId) int
}

func NewResourceSet[T client.Object](
	resources ...T,
) ResourceSet[T] {
	rs := &resourceSet[T]{
		set:         []T{},
		compareFunc: ezkube.ResourceIdsCompare,
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

func (s *resourceSet[T]) List(filterResource ...func(T) bool) func(yield func(int, T) bool) {
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
// The set is sorted based on the compare func.
func (s *resourceSet[T]) Insert(resources ...T) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	for _, resource := range resources {
		insertIndex, found := slices.BinarySearchFunc(
			s.set,
			resource,
			func(a, b T) int { return s.compareFunc(a, b) },
		)
		if found {
			s.set[insertIndex] = resource
			continue
		}
		// insert the resource at the determined index
		s.set = slices.Insert(s.set, insertIndex, resource)
	}
}

func (s *resourceSet[T]) Has(resource T) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	_, found := slices.BinarySearchFunc(
		s.set,
		resource,
		func(a, b T) int { return s.compareFunc(a, b) },
	)
	return found
}

func (s *resourceSet[T]) Equal(
	set ResourceSet[T],
) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.Generic().Equal(set.Generic())
}

func (s *resourceSet[T]) Delete(resource ezkube.ResourceId) {
	s.lock.Lock()
	defer s.lock.Unlock()

	i, found := slices.BinarySearchFunc(
		s.set,
		resource,
		func(a T, b ezkube.ResourceId) int { return s.compareFunc(a, b) },
	)
	if found {
		s.set = slices.Delete(s.set, i, i+1)
	}
}

func (s *resourceSet[T]) Union(set ResourceSet[T]) ResourceSet[T] {
	list := []T{}
	for _, resource := range s.Generic().Union(set.Generic()).List() {
		list = append(list, resource.(T))
	}
	return NewResourceSet[T](
		list...,
	)
}

func (s *resourceSet[T]) Difference(set ResourceSet[T]) ResourceSet[T] {
	s.lock.RLock()
	defer s.lock.RUnlock()
	result := NewResourceSet[T]()
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
	result := NewResourceSet[T]()
	if len(s.set) < set.Len() {
		walk = NewResourceSet(s.set...)
		other = set
	} else {
		walk = set
		other = NewResourceSet(s.set...)
	}
	walk.List()(func(_ int, key T) bool {
		if other.Has(key) {
			result.Insert(key)
		}
		return true
	})
	return result
}

func (s *resourceSet[T]) Find(resource ezkube.ResourceId) (T, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	insertIndex, found := slices.BinarySearchFunc(
		s.set,
		resource,
		func(a T, b ezkube.ResourceId) int { return s.compareFunc(a, b) },
	)
	if found {
		return s.set[insertIndex], nil
	}

	var r T
	return r, sk_sets.NotFoundErr(r, resource)
}

func (s *resourceSet[T]) Len() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.set)
}

// note that this function will currently panic if called for a ResourceSet[T] containing non-runtime.Objects
func (oldSet *resourceSet[T]) Delta(newSet ResourceSet[T]) sk_sets.ResourceDelta {
	updated, removed := NewResourceSet[T](), NewResourceSet[T]()

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
	new := NewResourceSet[T]()

	oldSet.List()(func(_ int, oldObj T) bool {
		copy := oldObj.DeepCopyObject().(T)
		new.Insert(copy)
		return true
	})
	return new
}

func (s *resourceSet[T]) Generic() sk_sets.ResourceSet {
	set := sk_sets.NewResourceSet(nil)
	s.List()(func(_ int, v T) bool {
		set.Insert(v)
		return true
	})
	return set
}

func (s *resourceSet[T]) GetCompareFunc() func(a, b ezkube.ResourceId) int {
	return s.compareFunc
}
