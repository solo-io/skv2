package sets_v2

import (
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
	List(filterResource ...func(T) bool) []T
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	UnsortedList(filterResource ...func(T) bool) []T
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
	// GetSortFunc() func(toInsert, existing T) bool
	// GetEqualityFunc() func(a, b T) bool
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
	lock         sync.RWMutex
	set          []T
	sortFunc     func(toInsert, existing T) bool
	equalityFunc func(a, b T) bool
}

func NewResourceSet[T client.Object](
	sortFunc func(toInsert, existing T) bool,
	equalityFunc func(a, b T) bool,
	resources ...T,
) ResourceSet[T] {
	rs := &resourceSet[T]{
		set:          []T{},
		sortFunc:     sortFunc,
		equalityFunc: equalityFunc,
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

func (s *resourceSet[T]) List(filterResource ...func(T) bool) []T {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.set
}

func (s *resourceSet[T]) UnsortedList(filterResource ...func(T) bool) []T {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.List()
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
	for _, objToInsert := range resources {
		s.insert(objToInsert)
	}
}

func (s *resourceSet[T]) insert(resource T) {
	insertIndex := sort.Search(len(s.set), func(i int) bool { return s.sortFunc(resource, s.set[i]) })

	// if the resource is already in the set, replace it
	if insertIndex < len(s.set) && s.equalityFunc(resource, s.set[insertIndex]) {
		s.set[insertIndex] = resource
		return
	}
	if s.sortFunc == nil {
		s.set = append(s.set, resource)
		return
	}

	// insert the resource at the determined index
	newSet := make([]T, len(s.set)+1)
	copy(newSet, s.set[:insertIndex])
	newSet[insertIndex] = resource
	copy(newSet[insertIndex+1:], s.set[insertIndex:])
	s.set = newSet
}

func (s *resourceSet[T]) Has(resource T) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	insertIndex := sort.Search(len(s.set), func(i int) bool { return s.sortFunc(resource, s.set[i]) })
	return insertIndex < len(s.set) && s.equalityFunc(resource, s.set[insertIndex])
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
	i := sort.Search(len(s.set), func(i int) bool { return s.equalityFunc(s.set[i], resource) })
	if i == len(s.set) {
		return
	}
	newSet := make([]T, len(s.set)-1)
	copy(newSet, s.set[:i])
	copy(newSet[i:], s.set[i+1:])
	s.set = newSet
}

func (s *resourceSet[T]) Union(set ResourceSet[T]) ResourceSet[T] {
	return NewResourceSet[T](s.sortFunc, s.equalityFunc, append(s.List(), set.List()...)...)
}

func (s *resourceSet[T]) Difference(set ResourceSet[T]) ResourceSet[T] {
	s.lock.RLock()
	defer s.lock.RUnlock()
	result := NewResourceSet[T](s.sortFunc, s.equalityFunc)
	for _, resource := range s.set {
		if !set.Has(resource) {
			result.Insert(resource)
		}
	}
	return result
}

func (s *resourceSet[T]) Intersection(set ResourceSet[T]) ResourceSet[T] {
	var walk, other ResourceSet[T]
	result := NewResourceSet(s.sortFunc, s.equalityFunc)
	if s.Length() < set.Length() {
		walk = NewResourceSet(s.sortFunc, s.equalityFunc, s.List()...)
		other = set
	} else {
		walk = set
		other = NewResourceSet(s.sortFunc, s.equalityFunc, s.List()...)
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
	i := sort.Search(s.Length(), func(i int) bool {
		return s.equalityFunc(s.set[i], id)
	})
	resource := s.set[i]
	if i == s.Length() {
		return resource, sk_sets.NotFoundErr(resource, id)
	}

	return resource, nil
}

func (s *resourceSet[T]) Length() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.set)
}

// note that this function will currently panic if called for a ResourceSet[T] containing non-runtime.Objects
func (oldSet *resourceSet[T]) Delta(newSet ResourceSet[T]) sk_sets.ResourceDelta {
	updated, removed := NewResourceSet[T](nil, nil), NewResourceSet[T](nil, nil)

	// find objects updated or removed
	oldSet.List(func(oldObj T) bool {
		newObj, err := newSet.Find(oldObj)
		switch {
		case err != nil:
			// obj removed
			removed.Insert(oldObj)
		case !controllerutils.ObjectsEqual(oldObj, newObj):
			// obj updated
			updated.Insert(newObj)
		default:
			// obj the same
		}
		return true // return value ignored
	})

	// find objects added
	newSet.List(func(newObj T) bool {
		if _, err := oldSet.Find(newObj); err != nil {
			// obj added
			updated.Insert(newObj)
		}
		return true // return value ignored
	})
	delta := &ResourceDelta[T]{
		Inserted: updated,
		Removed:  removed,
	}
	return delta.DeltaV1()
}

// Create a clone of the current set
// note that this function will currently panic if called for a ResourceSet[T] containing non-runtime.Objects
func (oldSet *resourceSet[T]) Clone() ResourceSet[T] {
	new := NewResourceSet[T](oldSet.sortFunc, oldSet.equalityFunc)
	oldSet.List(func(oldObj T) bool {
		copy := oldObj.DeepCopyObject().(T)
		new.Insert(copy)
		return true
	})
	return new
}

func (s *resourceSet[T]) Generic() sk_sets.ResourceSet {
	genericSortFunc := func(toInsert, existing ezkube.ResourceId) bool {
		return s.sortFunc(toInsert.(T), existing.(T))
	}
	genericEqualityFunc := func(a, b ezkube.ResourceId) bool {
		return s.equalityFunc(a.(T), b.(T))
	}
	set := sk_sets.NewResourceSet(genericSortFunc, genericEqualityFunc)
	for _, v := range s.List() {
		set.Insert(v)
	}
	return set
}

func (s *resourceSet[T]) GetSortFunc() func(toInsert, existing T) bool {
	return s.sortFunc
}

func (s *resourceSet[T]) GetEqualityFunc() func(a, b T) bool {
	return s.equalityFunc
}
