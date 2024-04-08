package sets_v2

import (
	"slices"
	"sync"

	sk_sets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ResourceSet[T client.Object] interface {
	// Get the set stored keys
	Keys() sets.Set[string]

	// Filter returns an iterator that will iterate over the set of elements
	// that match the provided filter. If the filter returns true, the resource will be included in the iteration.
	// The index and resource are passed to the provided function for every element in the *filtered set*.
	// The index is the index of the resource in the *filtered* set.
	// The iteration can be stopped by returning false from the function. This can be thought of as a "break" statement in a loop.
	// Returning true will continue the iteration. This can be thought of as a "continue" statement in a loop.
	// For iteration that does not need to be filtered, use Iter.
	Filter(filterResource func(T) bool) func(yield func(int, T) bool)

	// Iter iterates over the set, passing the index and resource to the provided function for every element in the set.
	// The iteration can be stopped by returning false from the function. This can be thought of as a "break" statement in a loop.
	// Returning true will continue the iteration. This can be thought of as a "continue" statement in a loop.
	Iter(func(int, T) bool)

	// FilterOutAndCreateList constructs a list of resource that do not match any of the provided filters.
	// Use of this function should be limited to only when a filtered list is needed.
	// For iteration that does not require creating a new list, use FilterOutAndIterate.
	FilterOutAndCreateList(filterResource ...func(T) bool) []T
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
	Length() int
	// returns the generic implementation of the set
	Generic() sk_sets.ResourceSet
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

func (s *resourceSet[T]) FilterOutAndIterate(filterResource ...func(T) bool) func(yield func(int, T) bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return func(yield func(int, T) bool) {
		i := 0
	OUTER:
		for _, resource := range s.set {
			for _, filter := range filterResource {
				if filter(resource) {
					continue OUTER
				}
			}
			if !yield(i, resource) {
				break
			}
			i += 1
		}
	}
}

func (s *resourceSet[T]) Filter(filterResource func(T) bool) func(yield func(int, T) bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return func(yield func(int, T) bool) {
		i := 0
		for _, resource := range s.set {
			if filterResource(resource) {
				if !yield(i, resource) {
					break
				}
				i += 1
			}
		}
	}
}

func (s *resourceSet[T]) FilterOutAndCreateList(filterResource ...func(T) bool) []T {
	s.lock.RLock()
	defer s.lock.RUnlock()
	var ret []T
	for _, resource := range s.set {
		filtered := false
		for _, filter := range filterResource {
			if filter(resource) {
				filtered = true
				break
			}
		}
		if !filtered {
			ret = append(ret, resource)
		}
	}
	return ret
}

func (s *resourceSet[T]) Iter(yield func(int, T) bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	for i, resource := range s.set {
		if !yield(i, resource) {
			break
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
	walk.Iter(func(_ int, key T) bool {
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

func (s *resourceSet[T]) Length() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.set)
}

// Create a clone of the current set
// note that this function will currently panic if called for a ResourceSet[T] containing non-runtime.Objects
func (oldSet *resourceSet[T]) Clone() ResourceSet[T] {
	new := NewResourceSet[T]()

	oldSet.Iter(func(_ int, oldObj T) bool {
		copy := oldObj.DeepCopyObject().(T)
		new.Insert(copy)
		return true
	})
	return new
}

func (s *resourceSet[T]) Generic() sk_sets.ResourceSet {
	set := sk_sets.NewResourceSet(nil)
	s.Iter(func(_ int, v T) bool {
		set.Insert(v)
		return true
	})
	return set
}
