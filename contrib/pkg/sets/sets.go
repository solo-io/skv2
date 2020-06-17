package sets

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
)

// k8s resources are uniquely identified by their name and namespace
func Key(objectMeta metav1.Object) string {
	return objectMeta.GetName() + "." + objectMeta.GetNamespace() + "." + objectMeta.GetClusterName()
}

type ResourceSet interface {
	Keys() sets.String
	List() []metav1.Object
	Map() map[string]metav1.Object
	Insert(resource ...metav1.Object)
	Equal(resourceSet ResourceSet) bool
	Has(resource metav1.Object) bool
	Delete(resource metav1.Object)
	Union(set ResourceSet) ResourceSet
	Difference(set ResourceSet) ResourceSet
	Intersection(set ResourceSet) ResourceSet
}

type resourceSet struct {
	set     sets.String
	mapping map[string]metav1.Object
}

func NewResourceSet(resources ...metav1.Object) ResourceSet {
	set := sets.NewString()
	mapping := map[string]metav1.Object{}
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

func (s resourceSet) List() []metav1.Object {
	var Resources []metav1.Object
	for _, key := range s.set.List() {
		Resources = append(Resources, s.mapping[key])
	}
	return Resources
}

func (s resourceSet) Map() map[string]metav1.Object {
	newMap := map[string]metav1.Object{}
	for k, v := range s.mapping {
		newMap[k] = v
	}
	return newMap
}

func (s resourceSet) Insert(
	Resources ...metav1.Object,
) {
	for _, Resource := range Resources {
		key := Key(Resource)
		s.mapping[key] = Resource
		s.set.Insert(key)
	}
}

func (s resourceSet) Has(resource metav1.Object) bool {
	return s.set.Has(Key(resource))
}

func (s resourceSet) Equal(
	ResourceSet ResourceSet,
) bool {
	return s.set.Equal(ResourceSet.Keys())
}

func (s resourceSet) Delete(resource metav1.Object) {
	key := Key(resource)
	delete(s.mapping, key)
	s.set.Delete(key)
}

func (s resourceSet) Union(set ResourceSet) ResourceSet {
	return NewResourceSet(append(s.List(), set.List()...)...)
}

func (s resourceSet) Difference(set ResourceSet) ResourceSet {
	newSet := s.set.Difference(set.Keys())
	var newResources []metav1.Object
	for key, _ := range newSet {
		val, _ := s.mapping[key]
		newResources = append(newResources, val)
	}
	return NewResourceSet(newResources...)
}

func (s resourceSet) Intersection(set ResourceSet) ResourceSet {
	newSet := s.set.Intersection(set.Keys())
	var newResources []metav1.Object
	for key, _ := range newSet {
		val, _ := s.mapping[key]
		newResources = append(newResources, val)
	}
	return NewResourceSet(newResources...)
}
