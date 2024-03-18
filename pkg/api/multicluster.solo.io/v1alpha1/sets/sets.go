// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1alpha1sets

import (
	multicluster_solo_io_v1alpha1 "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1"

	"github.com/rotisserie/eris"
	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
)

type KubernetesClusterSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	List(filterResource ...func(*multicluster_solo_io_v1alpha1.KubernetesCluster) bool) []*multicluster_solo_io_v1alpha1.KubernetesCluster
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	UnsortedList(filterResource ...func(*multicluster_solo_io_v1alpha1.KubernetesCluster) bool) []*multicluster_solo_io_v1alpha1.KubernetesCluster
	// Return the Set as a map of key to resource.
	Map() map[string]*multicluster_solo_io_v1alpha1.KubernetesCluster
	// Insert a resource into the set.
	Insert(kubernetesCluster ...*multicluster_solo_io_v1alpha1.KubernetesCluster)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(kubernetesClusterSet KubernetesClusterSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(kubernetesCluster ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(kubernetesCluster ezkube.ResourceId)
	// Return the union with the provided set
	Union(set KubernetesClusterSet) KubernetesClusterSet
	// Return the difference with the provided set
	Difference(set KubernetesClusterSet) KubernetesClusterSet
	// Return the intersection with the provided set
	Intersection(set KubernetesClusterSet) KubernetesClusterSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*multicluster_solo_io_v1alpha1.KubernetesCluster, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another KubernetesClusterSet
	Delta(newSet KubernetesClusterSet) sksets.ResourceDelta
	// Create a deep copy of the current KubernetesClusterSet
	Clone() KubernetesClusterSet
}

func makeGenericKubernetesClusterSet(kubernetesClusterList []*multicluster_solo_io_v1alpha1.KubernetesCluster) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range kubernetesClusterList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type kubernetesClusterSet struct {
	set sksets.ResourceSet
}

func NewKubernetesClusterSet(kubernetesClusterList ...*multicluster_solo_io_v1alpha1.KubernetesCluster) KubernetesClusterSet {
	return &kubernetesClusterSet{set: makeGenericKubernetesClusterSet(kubernetesClusterList)}
}

func NewKubernetesClusterSetFromList(kubernetesClusterList *multicluster_solo_io_v1alpha1.KubernetesClusterList) KubernetesClusterSet {
	list := make([]*multicluster_solo_io_v1alpha1.KubernetesCluster, 0, len(kubernetesClusterList.Items))
	for idx := range kubernetesClusterList.Items {
		list = append(list, &kubernetesClusterList.Items[idx])
	}
	return &kubernetesClusterSet{set: makeGenericKubernetesClusterSet(list)}
}

func (s *kubernetesClusterSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *kubernetesClusterSet) List(filterResource ...func(*multicluster_solo_io_v1alpha1.KubernetesCluster) bool) []*multicluster_solo_io_v1alpha1.KubernetesCluster {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*multicluster_solo_io_v1alpha1.KubernetesCluster))
		})
	}

	objs := s.Generic().List(genericFilters...)
	kubernetesClusterList := make([]*multicluster_solo_io_v1alpha1.KubernetesCluster, 0, len(objs))
	for _, obj := range objs {
		kubernetesClusterList = append(kubernetesClusterList, obj.(*multicluster_solo_io_v1alpha1.KubernetesCluster))
	}
	return kubernetesClusterList
}

func (s *kubernetesClusterSet) UnsortedList(filterResource ...func(*multicluster_solo_io_v1alpha1.KubernetesCluster) bool) []*multicluster_solo_io_v1alpha1.KubernetesCluster {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*multicluster_solo_io_v1alpha1.KubernetesCluster))
		})
	}

	var kubernetesClusterList []*multicluster_solo_io_v1alpha1.KubernetesCluster
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		kubernetesClusterList = append(kubernetesClusterList, obj.(*multicluster_solo_io_v1alpha1.KubernetesCluster))
	}
	return kubernetesClusterList
}

func (s *kubernetesClusterSet) Map() map[string]*multicluster_solo_io_v1alpha1.KubernetesCluster {
	if s == nil {
		return nil
	}

	newMap := map[string]*multicluster_solo_io_v1alpha1.KubernetesCluster{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*multicluster_solo_io_v1alpha1.KubernetesCluster)
	}
	return newMap
}

func (s *kubernetesClusterSet) Insert(
	kubernetesClusterList ...*multicluster_solo_io_v1alpha1.KubernetesCluster,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range kubernetesClusterList {
		s.Generic().Insert(obj)
	}
}

func (s *kubernetesClusterSet) Has(kubernetesCluster ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(kubernetesCluster)
}

func (s *kubernetesClusterSet) Equal(
	kubernetesClusterSet KubernetesClusterSet,
) bool {
	if s == nil {
		return kubernetesClusterSet == nil
	}
	return s.Generic().Equal(kubernetesClusterSet.Generic())
}

func (s *kubernetesClusterSet) Delete(KubernetesCluster ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(KubernetesCluster)
}

func (s *kubernetesClusterSet) Union(set KubernetesClusterSet) KubernetesClusterSet {
	if s == nil {
		return set
	}
	return &kubernetesClusterMergedSet{sets: []sksets.ResourceSet{s.Generic(), set.Generic()}}
}

func (s *kubernetesClusterSet) Difference(set KubernetesClusterSet) KubernetesClusterSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &kubernetesClusterSet{set: newSet}
}

func (s *kubernetesClusterSet) Intersection(set KubernetesClusterSet) KubernetesClusterSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var kubernetesClusterList []*multicluster_solo_io_v1alpha1.KubernetesCluster
	for _, obj := range newSet.List() {
		kubernetesClusterList = append(kubernetesClusterList, obj.(*multicluster_solo_io_v1alpha1.KubernetesCluster))
	}
	return NewKubernetesClusterSet(kubernetesClusterList...)
}

func (s *kubernetesClusterSet) Find(id ezkube.ResourceId) (*multicluster_solo_io_v1alpha1.KubernetesCluster, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find KubernetesCluster %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&multicluster_solo_io_v1alpha1.KubernetesCluster{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*multicluster_solo_io_v1alpha1.KubernetesCluster), nil
}

func (s *kubernetesClusterSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *kubernetesClusterSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *kubernetesClusterSet) Delta(newSet KubernetesClusterSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *kubernetesClusterSet) Clone() KubernetesClusterSet {
	if s == nil {
		return nil
	}
	return &kubernetesClusterMergedSet{sets: []sksets.ResourceSet{s.Generic()}}
}

type kubernetesClusterMergedSet struct {
	sets []sksets.ResourceSet
}

func NewKubernetesClusterMergedSet(kubernetesClusterList ...*multicluster_solo_io_v1alpha1.KubernetesCluster) KubernetesClusterSet {
	return &kubernetesClusterMergedSet{sets: []sksets.ResourceSet{makeGenericKubernetesClusterSet(kubernetesClusterList)}}
}

func NewKubernetesClusterMergedSetFromList(kubernetesClusterList *multicluster_solo_io_v1alpha1.KubernetesClusterList) KubernetesClusterSet {
	list := make([]*multicluster_solo_io_v1alpha1.KubernetesCluster, 0, len(kubernetesClusterList.Items))
	for idx := range kubernetesClusterList.Items {
		list = append(list, &kubernetesClusterList.Items[idx])
	}
	return &kubernetesClusterMergedSet{sets: []sksets.ResourceSet{makeGenericKubernetesClusterSet(list)}}
}

func (s *kubernetesClusterMergedSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	toRet := sets.String{}
	for _, set := range s.sets {
		toRet = toRet.Union(set.Keys())
	}
	return toRet
}

func (s *kubernetesClusterMergedSet) List(filterResource ...func(*multicluster_solo_io_v1alpha1.KubernetesCluster) bool) []*multicluster_solo_io_v1alpha1.KubernetesCluster {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*multicluster_solo_io_v1alpha1.KubernetesCluster))
		})
	}
	kubernetesClusterList := []*multicluster_solo_io_v1alpha1.KubernetesCluster{}
	for _, set := range s.sets {
		for _, obj := range set.List(genericFilters...) {
			kubernetesClusterList = append(kubernetesClusterList, obj.(*multicluster_solo_io_v1alpha1.KubernetesCluster))
		}
	}
	return kubernetesClusterList
}

func (s *kubernetesClusterMergedSet) UnsortedList(filterResource ...func(*multicluster_solo_io_v1alpha1.KubernetesCluster) bool) []*multicluster_solo_io_v1alpha1.KubernetesCluster {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*multicluster_solo_io_v1alpha1.KubernetesCluster))
		})
	}

	kubernetesClusterList := []*multicluster_solo_io_v1alpha1.KubernetesCluster{}
	for _, set := range s.sets {
		for _, obj := range set.UnsortedList(genericFilters...) {
			kubernetesClusterList = append(kubernetesClusterList, obj.(*multicluster_solo_io_v1alpha1.KubernetesCluster))
		}
	}
	return kubernetesClusterList
}

func (s *kubernetesClusterMergedSet) Map() map[string]*multicluster_solo_io_v1alpha1.KubernetesCluster {
	if s == nil {
		return nil
	}

	newMap := map[string]*multicluster_solo_io_v1alpha1.KubernetesCluster{}
	for _, set := range s.sets {
		for k, v := range set.Map() {
			newMap[k] = v.(*multicluster_solo_io_v1alpha1.KubernetesCluster)
		}
	}
	return newMap
}

func (s *kubernetesClusterMergedSet) Insert(
	kubernetesClusterList ...*multicluster_solo_io_v1alpha1.KubernetesCluster,
) {
	if s == nil {
	}
	if len(s.sets) == 0 {
		s.sets = append(s.sets, makeGenericKubernetesClusterSet(kubernetesClusterList))
	}
	for _, obj := range kubernetesClusterList {
		s.sets[0].Insert(obj)
	}
}

func (s *kubernetesClusterMergedSet) Has(kubernetesCluster ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	for _, set := range s.sets {
		if set.Has(kubernetesCluster) {
			return true
		}
	}
	return false
}

func (s *kubernetesClusterMergedSet) Equal(
	kubernetesClusterSet KubernetesClusterSet,
) bool {
	panic("unimplemented")
}

func (s *kubernetesClusterMergedSet) Delete(KubernetesCluster ezkube.ResourceId) {
	for _, set := range s.sets {
		set.Delete(KubernetesCluster)
	}
}

func (s *kubernetesClusterMergedSet) Union(set KubernetesClusterSet) KubernetesClusterSet {
	return &kubernetesClusterMergedSet{sets: append(s.sets, set.Generic())}
}

func (s *kubernetesClusterMergedSet) Difference(set KubernetesClusterSet) KubernetesClusterSet {
	panic("unimplemented")
}

func (s *kubernetesClusterMergedSet) Intersection(set KubernetesClusterSet) KubernetesClusterSet {
	panic("unimplemented")
}

func (s *kubernetesClusterMergedSet) Find(id ezkube.ResourceId) (*multicluster_solo_io_v1alpha1.KubernetesCluster, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find KubernetesCluster %v", sksets.Key(id))
	}

	var err error
	for _, set := range s.sets {
		var obj ezkube.ResourceId
		obj, err = set.Find(&multicluster_solo_io_v1alpha1.KubernetesCluster{}, id)
		if err == nil {
			return obj.(*multicluster_solo_io_v1alpha1.KubernetesCluster), nil
		}
	}

	return nil, err
}

func (s *kubernetesClusterMergedSet) Length() int {
	if s == nil {
		return 0
	}
	totalLen := 0
	for _, set := range s.sets {
		totalLen += set.Length()
	}
	return totalLen
}

func (s *kubernetesClusterMergedSet) Generic() sksets.ResourceSet {
	panic("unimplemented")
}

func (s *kubernetesClusterMergedSet) Delta(newSet KubernetesClusterSet) sksets.ResourceDelta {
	panic("unimplemented")
}

func (s *kubernetesClusterMergedSet) Clone() KubernetesClusterSet {
	if s == nil {
		return nil
	}
	return &kubernetesClusterMergedSet{sets: s.sets[:]}
}
