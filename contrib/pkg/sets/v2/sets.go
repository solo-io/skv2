package sets_v2

import (
	"sync"

	sk_sets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/controllerutils"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ResourceSet[T client.Object] interface {
	Keys() sets.String
	List(filterResource ...func(T) bool) []T
	UnsortedList(filterResource ...func(T) bool) []T
	Map() map[string]T
	Insert(resource ...T)
	Equal(set ResourceSet[T]) bool
	Has(resource T) bool
	Delete(resource T)
	Union(set ResourceSet[T]) ResourceSet[T]
	Difference(set ResourceSet[T]) ResourceSet[T]
	Intersection(set ResourceSet[T]) ResourceSet[T]
	Find(resourceType, id T) (T, error)
	Length() int
	// returns the delta between this and and another ResourceSet[T]
	Delta(newSet ResourceSet[T]) ResourceDelta[T]
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

type resourceSet[T client.Object] struct {
	lock    sync.RWMutex
	set     sets.String
	mapping map[string]T
}

func NewResourceSet[T client.Object](resources ...T) ResourceSet[T] {
	set := sets.NewString()
	mapping := map[string]T{}
	for _, resource := range resources {
		key := sk_sets.Key(resource)
		set.Insert(key)
		mapping[key] = resource
	}
	return &resourceSet[T]{set: set, mapping: mapping}
}

func (s *resourceSet[T]) Keys() sets.String {
	return sets.NewString(s.set.List()...)
}

func (s *resourceSet[T]) List(filterResource ...func(T) bool) []T {
	s.lock.RLock()
	defer s.lock.RUnlock()
	var resources []T
	for _, key := range s.set.List() {
		var filtered bool
		for _, filter := range filterResource {
			if filter(s.mapping[key]) {
				filtered = true
				break
			}
		}
		if !filtered {
			resources = append(resources, s.mapping[key])
		}
	}
	return resources
}

func (s *resourceSet[T]) UnsortedList(filterResource ...func(T) bool) []T {
	s.lock.RLock()
	defer s.lock.RUnlock()

	keys := s.set.UnsortedList()
	resources := make([]T, 0, len(keys))

	for _, key := range keys {
		var filtered bool
		for _, filter := range filterResource {
			if filter(s.mapping[key]) {
				filtered = true
				break
			}
		}
		if !filtered {
			resources = append(resources, s.mapping[key])
		}
	}
	return resources
}

func (s *resourceSet[T]) Map() map[string]T {
	s.lock.RLock()
	defer s.lock.RUnlock()
	newMap := map[string]T{}
	for k, v := range s.mapping {
		newMap[k] = v
	}
	return newMap
}

func (s *resourceSet[T]) Insert(
	resources ...T,
) {
	s.lock.Lock()
	defer s.lock.Unlock()
	for _, resource := range resources {
		key := sk_sets.Key(resource)
		s.mapping[key] = resource
		s.set.Insert(key)
	}
}

func (s *resourceSet[T]) Has(resource T) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.set.Has(sk_sets.Key(resource))
}

func (s *resourceSet[T]) Equal(
	set ResourceSet[T],
) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.set.Equal(set.Keys())
}

func (s *resourceSet[T]) Delete(resource T) {
	s.lock.Lock()
	defer s.lock.Unlock()
	key := sk_sets.Key(resource)
	delete(s.mapping, key)
	s.set.Delete(key)
}

func (s *resourceSet[T]) Union(set ResourceSet[T]) ResourceSet[T] {
	return NewResourceSet[T](append(s.List(), set.List()...)...)
}

func (s *resourceSet[T]) Difference(set ResourceSet[T]) ResourceSet[T] {
	s.lock.RLock()
	defer s.lock.RUnlock()
	newSet := s.set.Difference(set.Keys())
	var newResources []T
	for key, _ := range newSet {
		val, _ := s.mapping[key]
		newResources = append(newResources, val)
	}
	return NewResourceSet[T](newResources...)
}

func (s *resourceSet[T]) Intersection(set ResourceSet[T]) ResourceSet[T] {
	s.lock.RLock()
	defer s.lock.RUnlock()
	newSet := s.set.Intersection(set.Keys())
	var newResources []T
	for key, _ := range newSet {
		val, _ := s.mapping[key]
		newResources = append(newResources, val)
	}
	return NewResourceSet[T](newResources...)
}

func (s *resourceSet[T]) Find(
	resourceType,
	id T,
) (T, error) {
	var resource T
	s.lock.RLock()
	defer s.lock.RUnlock()
	key := sk_sets.Key(id)
	resource, ok := s.mapping[key]
	if !ok {
		return resource, sk_sets.NotFoundErr(resourceType, id)
	}

	return resource, nil
}

func (s *resourceSet[T]) Length() int {

	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.mapping)
}

// note that this function will currently panic if called for a ResourceSet[T] containing non-runtime.Objects
func (oldSet *resourceSet[T]) Delta(newSet ResourceSet[T]) ResourceDelta[T] {
	updated, removed := NewResourceSet[T](), NewResourceSet[T]()

	// find objects updated or removed
	oldSet.List(func(oldObj T) bool {
		newObj, err := newSet.Find(oldObj, oldObj)
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
		if _, err := oldSet.Find(newObj, newObj); err != nil {
			// obj added
			updated.Insert(newObj)
		}
		return true // return value ignored
	})
	return ResourceDelta[T]{
		Inserted: updated,
		Removed:  removed,
	}

}

// Create a clone of the current set
// note that this function will currently panic if called for a ResourceSet[T] containing non-runtime.Objects
func (oldSet *resourceSet[T]) Clone() ResourceSet[T] {
	new := NewResourceSet[T]()
	oldSet.List(func(oldObj T) bool {
		copy := oldObj.DeepCopyObject().(T)
		new.Insert(copy)
		return true
	})
	return new
}
