package sets

import (
	"fmt"
	"sync"

	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/rotisserie/eris"
	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
)

var NotFoundErr = func(resourceType ezkube.ResourceId, id ezkube.ResourceId) error {
	return eris.Errorf("%T with id %v not found", resourceType, Key(id))
}

// k8s resources are uniquely identified by their name and namespace
func Key(id ezkube.ResourceId) string {
	if clusterId, ok := id.(ezkube.ClusterResourceId); ok {
		return clusterId.GetName() + "." + clusterId.GetNamespace() + "." + clusterId.GetClusterName()
	}
	return id.GetName() + "." + id.GetNamespace() + "."
}

// typed keys are helpful for logging; currently unused in the Set implementation but placed here for convenience
func TypedKey(id ezkube.ResourceId) string {
	return fmt.Sprintf("%v.%T", Key(id), id)
}

type ResourceSet interface {
	Keys() sets.String
	List(filterResource ...func(ezkube.ResourceId) bool) []ezkube.ResourceId
	Map() map[string]ezkube.ResourceId
	Insert(resource ...ezkube.ResourceId)
	Equal(set ResourceSet) bool
	Has(resource ezkube.ResourceId) bool
	Delete(resource ezkube.ResourceId)
	Union(set ResourceSet) ResourceSet
	Difference(set ResourceSet) ResourceSet
	Intersection(set ResourceSet) ResourceSet
	Find(resourceType, id ezkube.ResourceId) (ezkube.ResourceId, error)
	Length() int
	// returns the delta between this and and another ResourceSet
	Delta(newSet ResourceSet) ResourceDelta
}

// ResourceDelta represents the set of changes between two ResourceSets.
type ResourceDelta struct {
	// the resources inserted into the set
	Inserted ResourceSet
	// the resources removed from the set
	Removed ResourceSet
}

type resourceSet struct {
	lock    sync.RWMutex
	set     sets.String
	mapping map[string]ezkube.ResourceId
}

func NewResourceSet(resources ...ezkube.ResourceId) ResourceSet {
	set := sets.NewString()
	mapping := map[string]ezkube.ResourceId{}
	for _, resource := range resources {
		if resource == nil {
			continue
		}
		key := Key(resource)
		set.Insert(key)
		mapping[key] = resource
	}
	return &resourceSet{set: set, mapping: mapping}
}

func (s *resourceSet) Keys() sets.String {
	return sets.NewString(s.set.List()...)
}

func (s *resourceSet) List(filterResource ...func(ezkube.ResourceId) bool) []ezkube.ResourceId {
	s.lock.RLock()
	defer s.lock.RUnlock()
	var resources []ezkube.ResourceId
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

func (s *resourceSet) Map() map[string]ezkube.ResourceId {
	s.lock.RLock()
	defer s.lock.RUnlock()
	newMap := map[string]ezkube.ResourceId{}
	for k, v := range s.mapping {
		newMap[k] = v
	}
	return newMap
}

func (s *resourceSet) Insert(
	resources ...ezkube.ResourceId,
) {
	s.lock.Lock()
	defer s.lock.Unlock()
	for _, resource := range resources {
		if resource == nil {
			continue
		}
		key := Key(resource)
		s.mapping[key] = resource
		s.set.Insert(key)
	}
}

func (s *resourceSet) Has(resource ezkube.ResourceId) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.set.Has(Key(resource))
}

func (s *resourceSet) Equal(
	set ResourceSet,
) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.set.Equal(set.Keys())
}

func (s *resourceSet) Delete(resource ezkube.ResourceId) {
	s.lock.Lock()
	defer s.lock.Unlock()
	key := Key(resource)
	delete(s.mapping, key)
	s.set.Delete(key)
}

func (s *resourceSet) Union(set ResourceSet) ResourceSet {
	return NewResourceSet(append(s.List(), set.List()...)...)
}

func (s *resourceSet) Difference(set ResourceSet) ResourceSet {
	s.lock.RLock()
	defer s.lock.RUnlock()
	newSet := s.set.Difference(set.Keys())
	var newResources []ezkube.ResourceId
	for key, _ := range newSet {
		val, _ := s.mapping[key]
		newResources = append(newResources, val)
	}
	return NewResourceSet(newResources...)
}

func (s *resourceSet) Intersection(set ResourceSet) ResourceSet {
	s.lock.RLock()
	defer s.lock.RUnlock()
	newSet := s.set.Intersection(set.Keys())
	var newResources []ezkube.ResourceId
	for key, _ := range newSet {
		val, _ := s.mapping[key]
		newResources = append(newResources, val)
	}
	return NewResourceSet(newResources...)
}

func (s *resourceSet) Find(
	resourceType,
	id ezkube.ResourceId,
) (ezkube.ResourceId, error) {

	s.lock.RLock()
	defer s.lock.RUnlock()
	key := Key(id)
	resource, ok := s.mapping[key]
	if !ok {
		return nil, NotFoundErr(resourceType, id)
	}

	return resource, nil
}

func (s *resourceSet) Length() int {

	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.mapping)
}

// note that this function will currently panic if called for a ResourceSet containing non-runtime.Objects
func (oldSet *resourceSet) Delta(newSet ResourceSet) ResourceDelta {
	updated, removed := NewResourceSet(), NewResourceSet()

	// find objects updated or removed
	oldSet.List(func(oldObj ezkube.ResourceId) bool {
		newObj, err := newSet.Find(oldObj, oldObj)
		switch {
		case err != nil:
			// obj removed
			removed.Insert(oldObj)
		case !controllerutils.ObjectsEqual(oldObj.(client.Object), newObj.(client.Object)):
			// obj updated
			updated.Insert(newObj)
		default:
			// obj the same
		}
		return true // return value ignored
	})

	// find objects added
	newSet.List(func(newObj ezkube.ResourceId) bool {
		if _, err := oldSet.Find(newObj, newObj); err != nil {
			// obj added
			updated.Insert(newObj)
		}
		return true // return value ignored
	})
	return ResourceDelta{
		Inserted: updated,
		Removed:  removed,
	}

}
