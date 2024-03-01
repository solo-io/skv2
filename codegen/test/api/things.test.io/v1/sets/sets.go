// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1sets

import (
	things_test_io_v1 "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1"

	"github.com/rotisserie/eris"
	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
)

type PaintSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	List(filterResource ...func(*things_test_io_v1.Paint) bool) []*things_test_io_v1.Paint
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	UnsortedList(filterResource ...func(*things_test_io_v1.Paint) bool) []*things_test_io_v1.Paint
	// Return the Set as a map of key to resource.
	Map() map[string]*things_test_io_v1.Paint
	// Insert a resource into the set.
	Insert(paint ...*things_test_io_v1.Paint)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(paintSet PaintSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(paint ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(paint ezkube.ResourceId)
	// Return the union with the provided set
	Union(set PaintSet) PaintSet
	// Return the difference with the provided set
	Difference(set PaintSet) PaintSet
	// Return the intersection with the provided set
	Intersection(set PaintSet) PaintSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*things_test_io_v1.Paint, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another PaintSet
	Delta(newSet PaintSet) sksets.ResourceDelta
	// Create a deep copy of the current PaintSet
	Clone() PaintSet
}

func makeGenericPaintSet(paintList []*things_test_io_v1.Paint) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range paintList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type paintSet struct {
	set sksets.ResourceSet
}

func NewPaintSet(paintList ...*things_test_io_v1.Paint) PaintSet {
	return &paintSet{set: makeGenericPaintSet(paintList)}
}

func NewPaintSetFromList(paintList *things_test_io_v1.PaintList) PaintSet {
	list := make([]*things_test_io_v1.Paint, 0, len(paintList.Items))
	for idx := range paintList.Items {
		list = append(list, &paintList.Items[idx])
	}
	return &paintSet{set: makeGenericPaintSet(list)}
}

func (s *paintSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *paintSet) List(filterResource ...func(*things_test_io_v1.Paint) bool) []*things_test_io_v1.Paint {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*things_test_io_v1.Paint))
		})
	}

	objs := s.Generic().List(genericFilters...)
	paintList := make([]*things_test_io_v1.Paint, 0, len(objs))
	for _, obj := range objs {
		paintList = append(paintList, obj.(*things_test_io_v1.Paint))
	}
	return paintList
}

func (s *paintSet) UnsortedList(filterResource ...func(*things_test_io_v1.Paint) bool) []*things_test_io_v1.Paint {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*things_test_io_v1.Paint))
		})
	}

	var paintList []*things_test_io_v1.Paint
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		paintList = append(paintList, obj.(*things_test_io_v1.Paint))
	}
	return paintList
}

func (s *paintSet) Map() map[string]*things_test_io_v1.Paint {
	if s == nil {
		return nil
	}

	newMap := map[string]*things_test_io_v1.Paint{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*things_test_io_v1.Paint)
	}
	return newMap
}

func (s *paintSet) Insert(
	paintList ...*things_test_io_v1.Paint,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range paintList {
		s.Generic().Insert(obj)
	}
}

func (s *paintSet) Has(paint ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(paint)
}

func (s *paintSet) Equal(
	paintSet PaintSet,
) bool {
	if s == nil {
		return paintSet == nil
	}
	return s.Generic().Equal(paintSet.Generic())
}

func (s *paintSet) Delete(Paint ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(Paint)
}

func (s *paintSet) Union(set PaintSet) PaintSet {
	if s == nil {
		return set
	}
	return &paintMergedSet{sets: []sksets.ResourceSet{s.Generic(), set.Generic()}}
}

func (s *paintSet) Difference(set PaintSet) PaintSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &paintSet{set: newSet}
}

func (s *paintSet) Intersection(set PaintSet) PaintSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var paintList []*things_test_io_v1.Paint
	for _, obj := range newSet.List() {
		paintList = append(paintList, obj.(*things_test_io_v1.Paint))
	}
	return NewPaintSet(paintList...)
}

func (s *paintSet) Find(id ezkube.ResourceId) (*things_test_io_v1.Paint, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find Paint %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&things_test_io_v1.Paint{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*things_test_io_v1.Paint), nil
}

func (s *paintSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *paintSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *paintSet) Delta(newSet PaintSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *paintSet) Clone() PaintSet {
	if s == nil {
		return nil
	}
	return &paintMergedSet{sets: []sksets.ResourceSet{s.Generic()}}
}

type paintMergedSet struct {
	sets []sksets.ResourceSet
}

func NewPaintMergedSet(paintList ...*things_test_io_v1.Paint) PaintSet {
	return &paintMergedSet{sets: []sksets.ResourceSet{makeGenericPaintSet(paintList)}}
}

func NewPaintMergedSetFromList(paintList *things_test_io_v1.PaintList) PaintSet {
	list := make([]*things_test_io_v1.Paint, 0, len(paintList.Items))
	for idx := range paintList.Items {
		list = append(list, &paintList.Items[idx])
	}
	return &paintMergedSet{sets: []sksets.ResourceSet{makeGenericPaintSet(list)}}
}

func (s *paintMergedSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	toRet := sets.String{}
	for _, set := range s.sets {
		toRet = toRet.Union(set.Keys())
	}
	return toRet
}

func (s *paintMergedSet) List(filterResource ...func(*things_test_io_v1.Paint) bool) []*things_test_io_v1.Paint {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return !filter(obj.(*things_test_io_v1.Paint))
		})
	}
	paintList := []*things_test_io_v1.Paint{}
	for _, set := range s.sets {
		for _, obj := range set.List(genericFilters...) {
			paintList = append(paintList, obj.(*things_test_io_v1.Paint))
		}
	}
	return paintList
}

func (s *paintMergedSet) UnsortedList(filterResource ...func(*things_test_io_v1.Paint) bool) []*things_test_io_v1.Paint {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return !filter(obj.(*things_test_io_v1.Paint))
		})
	}

	paintList := []*things_test_io_v1.Paint{}
	for _, set := range s.sets {
		for _, obj := range set.UnsortedList(genericFilters...) {
			paintList = append(paintList, obj.(*things_test_io_v1.Paint))
		}
	}
	return paintList
}

func (s *paintMergedSet) Map() map[string]*things_test_io_v1.Paint {
	if s == nil {
		return nil
	}

	newMap := map[string]*things_test_io_v1.Paint{}
	for _, set := range s.sets {
		for k, v := range set.Map() {
			newMap[k] = v.(*things_test_io_v1.Paint)
		}
	}
	return newMap
}

func (s *paintMergedSet) Insert(
	paintList ...*things_test_io_v1.Paint,
) {
	if s == nil {
	}
	if len(s.sets) == 0 {
		s.sets = append(s.sets, makeGenericPaintSet(paintList))
	}
	for _, obj := range paintList {
		s.sets[0].Insert(obj)
	}
}

func (s *paintMergedSet) Has(paint ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	for _, set := range s.sets {
		if set.Has(paint) {
			return true
		}
	}
	return false
}

func (s *paintMergedSet) Equal(
	paintSet PaintSet,
) bool {
	panic("unimplemented")
}

func (s *paintMergedSet) Delete(Paint ezkube.ResourceId) {
	panic("unimplemented")
}

func (s *paintMergedSet) Union(set PaintSet) PaintSet {
	return &paintMergedSet{sets: append(s.sets, set.Generic())}
}

func (s *paintMergedSet) Difference(set PaintSet) PaintSet {
	panic("unimplemented")
}

func (s *paintMergedSet) Intersection(set PaintSet) PaintSet {
	panic("unimplemented")
}

func (s *paintMergedSet) Find(id ezkube.ResourceId) (*things_test_io_v1.Paint, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find Paint %v", sksets.Key(id))
	}

	var err error
	for _, set := range s.sets {
		var obj ezkube.ResourceId
		obj, err = set.Find(&things_test_io_v1.Paint{}, id)
		if err == nil {
			return obj.(*things_test_io_v1.Paint), nil
		}
	}

	return nil, err
}

func (s *paintMergedSet) Length() int {
	if s == nil {
		return 0
	}
	totalLen := 0
	for _, set := range s.sets {
		totalLen += set.Length()
	}
	return totalLen
}

func (s *paintMergedSet) Generic() sksets.ResourceSet {
	panic("unimplemented")
}

func (s *paintMergedSet) Delta(newSet PaintSet) sksets.ResourceDelta {
	panic("unimplemented")
}

func (s *paintMergedSet) Clone() PaintSet {
	if s == nil {
		return nil
	}
	return &paintMergedSet{sets: s.sets[:]}
}

type ClusterResourceSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	List(filterResource ...func(*things_test_io_v1.ClusterResource) bool) []*things_test_io_v1.ClusterResource
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	// The filter function should return false to keep the resource, true to drop it.
	UnsortedList(filterResource ...func(*things_test_io_v1.ClusterResource) bool) []*things_test_io_v1.ClusterResource
	// Return the Set as a map of key to resource.
	Map() map[string]*things_test_io_v1.ClusterResource
	// Insert a resource into the set.
	Insert(clusterResource ...*things_test_io_v1.ClusterResource)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(clusterResourceSet ClusterResourceSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(clusterResource ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(clusterResource ezkube.ResourceId)
	// Return the union with the provided set
	Union(set ClusterResourceSet) ClusterResourceSet
	// Return the difference with the provided set
	Difference(set ClusterResourceSet) ClusterResourceSet
	// Return the intersection with the provided set
	Intersection(set ClusterResourceSet) ClusterResourceSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*things_test_io_v1.ClusterResource, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another ClusterResourceSet
	Delta(newSet ClusterResourceSet) sksets.ResourceDelta
	// Create a deep copy of the current ClusterResourceSet
	Clone() ClusterResourceSet
}

func makeGenericClusterResourceSet(clusterResourceList []*things_test_io_v1.ClusterResource) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range clusterResourceList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type clusterResourceSet struct {
	set sksets.ResourceSet
}

func NewClusterResourceSet(clusterResourceList ...*things_test_io_v1.ClusterResource) ClusterResourceSet {
	return &clusterResourceSet{set: makeGenericClusterResourceSet(clusterResourceList)}
}

func NewClusterResourceSetFromList(clusterResourceList *things_test_io_v1.ClusterResourceList) ClusterResourceSet {
	list := make([]*things_test_io_v1.ClusterResource, 0, len(clusterResourceList.Items))
	for idx := range clusterResourceList.Items {
		list = append(list, &clusterResourceList.Items[idx])
	}
	return &clusterResourceSet{set: makeGenericClusterResourceSet(list)}
}

func (s *clusterResourceSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *clusterResourceSet) List(filterResource ...func(*things_test_io_v1.ClusterResource) bool) []*things_test_io_v1.ClusterResource {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*things_test_io_v1.ClusterResource))
		})
	}

	objs := s.Generic().List(genericFilters...)
	clusterResourceList := make([]*things_test_io_v1.ClusterResource, 0, len(objs))
	for _, obj := range objs {
		clusterResourceList = append(clusterResourceList, obj.(*things_test_io_v1.ClusterResource))
	}
	return clusterResourceList
}

func (s *clusterResourceSet) UnsortedList(filterResource ...func(*things_test_io_v1.ClusterResource) bool) []*things_test_io_v1.ClusterResource {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*things_test_io_v1.ClusterResource))
		})
	}

	var clusterResourceList []*things_test_io_v1.ClusterResource
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		clusterResourceList = append(clusterResourceList, obj.(*things_test_io_v1.ClusterResource))
	}
	return clusterResourceList
}

func (s *clusterResourceSet) Map() map[string]*things_test_io_v1.ClusterResource {
	if s == nil {
		return nil
	}

	newMap := map[string]*things_test_io_v1.ClusterResource{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*things_test_io_v1.ClusterResource)
	}
	return newMap
}

func (s *clusterResourceSet) Insert(
	clusterResourceList ...*things_test_io_v1.ClusterResource,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range clusterResourceList {
		s.Generic().Insert(obj)
	}
}

func (s *clusterResourceSet) Has(clusterResource ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(clusterResource)
}

func (s *clusterResourceSet) Equal(
	clusterResourceSet ClusterResourceSet,
) bool {
	if s == nil {
		return clusterResourceSet == nil
	}
	return s.Generic().Equal(clusterResourceSet.Generic())
}

func (s *clusterResourceSet) Delete(ClusterResource ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(ClusterResource)
}

func (s *clusterResourceSet) Union(set ClusterResourceSet) ClusterResourceSet {
	if s == nil {
		return set
	}
	return &clusterResourceMergedSet{sets: []sksets.ResourceSet{s.Generic(), set.Generic()}}
}

func (s *clusterResourceSet) Difference(set ClusterResourceSet) ClusterResourceSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &clusterResourceSet{set: newSet}
}

func (s *clusterResourceSet) Intersection(set ClusterResourceSet) ClusterResourceSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var clusterResourceList []*things_test_io_v1.ClusterResource
	for _, obj := range newSet.List() {
		clusterResourceList = append(clusterResourceList, obj.(*things_test_io_v1.ClusterResource))
	}
	return NewClusterResourceSet(clusterResourceList...)
}

func (s *clusterResourceSet) Find(id ezkube.ResourceId) (*things_test_io_v1.ClusterResource, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find ClusterResource %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&things_test_io_v1.ClusterResource{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*things_test_io_v1.ClusterResource), nil
}

func (s *clusterResourceSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *clusterResourceSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *clusterResourceSet) Delta(newSet ClusterResourceSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *clusterResourceSet) Clone() ClusterResourceSet {
	if s == nil {
		return nil
	}
	return &clusterResourceMergedSet{sets: []sksets.ResourceSet{s.Generic()}}
}

type clusterResourceMergedSet struct {
	sets []sksets.ResourceSet
}

func NewClusterResourceMergedSet(clusterResourceList ...*things_test_io_v1.ClusterResource) ClusterResourceSet {
	return &clusterResourceMergedSet{sets: []sksets.ResourceSet{makeGenericClusterResourceSet(clusterResourceList)}}
}

func NewClusterResourceMergedSetFromList(clusterResourceList *things_test_io_v1.ClusterResourceList) ClusterResourceSet {
	list := make([]*things_test_io_v1.ClusterResource, 0, len(clusterResourceList.Items))
	for idx := range clusterResourceList.Items {
		list = append(list, &clusterResourceList.Items[idx])
	}
	return &clusterResourceMergedSet{sets: []sksets.ResourceSet{makeGenericClusterResourceSet(list)}}
}

func (s *clusterResourceMergedSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	toRet := sets.String{}
	for _, set := range s.sets {
		toRet = toRet.Union(set.Keys())
	}
	return toRet
}

func (s *clusterResourceMergedSet) List(filterResource ...func(*things_test_io_v1.ClusterResource) bool) []*things_test_io_v1.ClusterResource {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*things_test_io_v1.ClusterResource))
		})
	}
	clusterResourceList := []*things_test_io_v1.ClusterResource{}
	for _, set := range s.sets {
		for _, obj := range set.List(genericFilters...) {
			clusterResourceList = append(clusterResourceList, obj.(*things_test_io_v1.ClusterResource))
		}
	}
	return clusterResourceList
}

func (s *clusterResourceMergedSet) UnsortedList(filterResource ...func(*things_test_io_v1.ClusterResource) bool) []*things_test_io_v1.ClusterResource {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*things_test_io_v1.ClusterResource))
		})
	}

	clusterResourceList := []*things_test_io_v1.ClusterResource{}
	for _, set := range s.sets {
		for _, obj := range set.UnsortedList(genericFilters...) {
			clusterResourceList = append(clusterResourceList, obj.(*things_test_io_v1.ClusterResource))
		}
	}
	return clusterResourceList
}

func (s *clusterResourceMergedSet) Map() map[string]*things_test_io_v1.ClusterResource {
	if s == nil {
		return nil
	}

	newMap := map[string]*things_test_io_v1.ClusterResource{}
	for _, set := range s.sets {
		for k, v := range set.Map() {
			newMap[k] = v.(*things_test_io_v1.ClusterResource)
		}
	}
	return newMap
}

func (s *clusterResourceMergedSet) Insert(
	clusterResourceList ...*things_test_io_v1.ClusterResource,
) {
	if s == nil {
	}
	if len(s.sets) == 0 {
		s.sets = append(s.sets, makeGenericClusterResourceSet(clusterResourceList))
	}
	for _, obj := range clusterResourceList {
		s.sets[0].Insert(obj)
	}
}

func (s *clusterResourceMergedSet) Has(clusterResource ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	for _, set := range s.sets {
		if set.Has(clusterResource) {
			return true
		}
	}
	return false
}

func (s *clusterResourceMergedSet) Equal(
	clusterResourceSet ClusterResourceSet,
) bool {
	panic("unimplemented")
}

func (s *clusterResourceMergedSet) Delete(ClusterResource ezkube.ResourceId) {
	panic("unimplemented")
}

func (s *clusterResourceMergedSet) Union(set ClusterResourceSet) ClusterResourceSet {
	return &clusterResourceMergedSet{sets: append(s.sets, set.Generic())}
}

func (s *clusterResourceMergedSet) Difference(set ClusterResourceSet) ClusterResourceSet {
	panic("unimplemented")
}

func (s *clusterResourceMergedSet) Intersection(set ClusterResourceSet) ClusterResourceSet {
	panic("unimplemented")
}

func (s *clusterResourceMergedSet) Find(id ezkube.ResourceId) (*things_test_io_v1.ClusterResource, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find ClusterResource %v", sksets.Key(id))
	}

	var err error
	for _, set := range s.sets {
		var obj ezkube.ResourceId
		obj, err = set.Find(&things_test_io_v1.ClusterResource{}, id)
		if err == nil {
			return obj.(*things_test_io_v1.ClusterResource), nil
		}
	}

	return nil, err
}

func (s *clusterResourceMergedSet) Length() int {
	if s == nil {
		return 0
	}
	totalLen := 0
	for _, set := range s.sets {
		totalLen += set.Length()
	}
	return totalLen
}

func (s *clusterResourceMergedSet) Generic() sksets.ResourceSet {
	panic("unimplemented")
}

func (s *clusterResourceMergedSet) Delta(newSet ClusterResourceSet) sksets.ResourceDelta {
	panic("unimplemented")
}

func (s *clusterResourceMergedSet) Clone() ClusterResourceSet {
	if s == nil {
		return nil
	}
	return &clusterResourceMergedSet{sets: s.sets[:]}
}
