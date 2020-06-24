package sets

import (
	"sync"

	"github.com/rotisserie/eris"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
)

var NotFoundErr = func(resourceType ezkube.ResourceId, id ezkube.ResourceId) error {
	return eris.Errorf("%T with id %v not found", resourceType, Key(id))
}

// k8s resources are uniquely identified by their name and namespace
func Key(id ezkube.ResourceId) string {
	return id.GetName() + "." + id.GetNamespace() + "." + id.GetClusterName()
}

type ResourceSet interface {
	Keys() sets.String
	List() []ezkube.ResourceId
	Map() map[string]ezkube.ResourceId
	Insert(resource ...ezkube.ResourceId)
	Equal(resourceSet ResourceSet) bool
	Has(resource ezkube.ResourceId) bool
	Delete(resource ezkube.ResourceId)
	Union(set ResourceSet) ResourceSet
	Difference(set ResourceSet) ResourceSet
	Intersection(set ResourceSet) ResourceSet
	Find(resourceType, id ezkube.ResourceId) (ezkube.ResourceId, error)
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
		key := Key(resource)
		set.Insert(key)
		mapping[key] = resource
	}
	return &resourceSet{set: set, mapping: mapping}
}

func (s resourceSet) Keys() sets.String {
	return sets.NewString(s.set.List()...)
}

func (s resourceSet) List() []ezkube.ResourceId {
	s.lock.RLock()
	defer s.lock.RUnlock()
	var Resources []ezkube.ResourceId
	for _, key := range s.set.List() {
		Resources = append(Resources, s.mapping[key])
	}
	return Resources
}

func (s resourceSet) Map() map[string]ezkube.ResourceId {
	s.lock.RLock()
	defer s.lock.RUnlock()
	newMap := map[string]ezkube.ResourceId{}
	for k, v := range s.mapping {
		newMap[k] = v
	}
	return newMap
}

func (s resourceSet) Insert(
	Resources ...ezkube.ResourceId,
) {
	s.lock.Lock()
	defer s.lock.Unlock()
	for _, Resource := range Resources {
		key := Key(Resource)
		s.mapping[key] = Resource
		s.set.Insert(key)
	}
}

func (s resourceSet) Has(resource ezkube.ResourceId) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.set.Has(Key(resource))
}

func (s resourceSet) Equal(
	ResourceSet ResourceSet,
) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.set.Equal(ResourceSet.Keys())
}

func (s resourceSet) Delete(resource ezkube.ResourceId) {
	s.lock.Lock()
	defer s.lock.Unlock()
	key := Key(resource)
	delete(s.mapping, key)
	s.set.Delete(key)
}

func (s resourceSet) Union(set ResourceSet) ResourceSet {
	return NewResourceSet(append(s.List(), set.List()...)...)
}

func (s resourceSet) Difference(set ResourceSet) ResourceSet {
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

func (s resourceSet) Intersection(set ResourceSet) ResourceSet {
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

func (s resourceSet) Find(
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
