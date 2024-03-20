package sets

import (
	"sort"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Resources struct {
	// set is a list of unique entries sorted according to sortFunc
	set []ezkube.ResourceId
	// sortFunc is the function used to sort the set and returns true if toInsert should be inserted before existing.
	// If nil, the set will be unsorted.
	sortFunc func(toInsert, existing ezkube.ResourceId) bool
	// sortIndex is a map of resource keys to their index position in the set
	sortIndex map[string]int
}

func newResources(
	sortFunc func(toInsert, existing ezkube.ResourceId) bool,
	resources ...ezkube.ResourceId,
) Resources {
	r := Resources{
		set:       append([]ezkube.ResourceId{}, resources...),
		sortFunc:  sortFunc,
		sortIndex: make(map[string]int, len(resources)),
	}
	for _, resource := range resources {
		if resource == nil {
			continue
		}
		r.Insert(resource)
	}
	return r
}

func (r Resources) Find(resourceType, id ezkube.ResourceId) (ezkube.ResourceId, error) {
	key := Key(id)
	index, ok := r.sortIndex[key]
	if !ok {
		return nil, NotFoundErr(resourceType, id)
	}

	return r.set[index], nil
}

func (r Resources) Length() int {
	return len(r.set)
}

func (r Resources) Delta(newSet ResourceSet) ResourceDelta {
	updated, removed := NewResourceSet(r.sortFunc), NewResourceSet(r.sortFunc)

	// find objects updated or removed
	r.List(
		func(oldObj ezkube.ResourceId) bool {
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
		},
	)

	// find objects added
	newSet.List(
		func(newObj ezkube.ResourceId) bool {
			if _, err := r.Find(newObj, newObj); err != nil {
				// obj added
				updated.Insert(newObj)
			}
			return true // return value ignored
		},
	)
	return ResourceDelta{
		Inserted: updated,
		Removed:  removed,
	}
}

func (r Resources) Clone() ResourceSet {
	new := NewResourceSet(r.sortFunc, r.set...)
	r.List(
		func(oldObj ezkube.ResourceId) bool {
			copy := oldObj.(client.Object).DeepCopyObject().(client.Object)
			new.Insert(copy)
			return true
		},
	)
	return new
}

func (r Resources) Keys() sets.String {
	keys := sets.NewString()
	for key, _ := range r.sortIndex {
		keys.Insert(key)
	}
	return keys
}

func (r Resources) List(filterResource ...func(ezkube.ResourceId) bool) []ezkube.ResourceId {
	if len(filterResource) == 0 {
		return r.set
	}
	resources := make([]ezkube.ResourceId, 0, len(r.set))
	for _, resource := range r.set {
		var filtered bool
		for _, filter := range filterResource {
			if filter(resource) {
				filtered = true
				break
			}
		}
		if !filtered {
			resources = append(resources, resource)
		}
	}
	return resources
}

func (r Resources) UnsortedList(filterResource ...func(ezkube.ResourceId) bool) []ezkube.ResourceId {
	return r.List(filterResource...)
}

func (r Resources) Map() map[string]ezkube.ResourceId {
	shallowCopy := make(map[string]ezkube.ResourceId, len(r.sortIndex))
	for key, i := range r.sortIndex {
		shallowCopy[key] = r.set[i]
	}
	return shallowCopy
}

// Insert adds items to the set.
// If an item is already in the set, it is overwritten.
// The set is sorted based on the sortFunc. If sortFunc is nil, the set will be unsorted.
func (r Resources) Insert(resources ...ezkube.ResourceId) {
	for _, objToInsert := range resources {
		r.insert(objToInsert)
	}
}

func (r Resources) insert(objToInsert ezkube.ResourceId) {
	// index to start iterating set at to update sort order map
	earliestIndex := -1

	key := Key(objToInsert)
	currentIndex, exists := r.sortIndex[key]

	if r.sortFunc == nil {
		if exists {
			r.set[currentIndex] = objToInsert
			return
		}
		r.set = append(r.set, objToInsert)
		r.sortIndex[key] = len(r.set) - 1
		return
	}

	insertIndex := sort.Search(len(r.set), func(i int) bool { return r.sortFunc(objToInsert, r.set[i]) })

	// if the resource exists in the set, update the resource and determine if the sort order map needs to be updated
	if exists {
		r.set[r.sortIndex[key]] = objToInsert
		if insertIndex != currentIndex {
			earliestIndex = insertIndex
			if earliestIndex == len(r.set) {
				// decrement by 1 to avoid out of range
				earliestIndex--
			}

			for i := earliestIndex; i < len(r.set[earliestIndex:]); i++ {
				r.sortIndex[Key(r.set[i])] = i
			}
		}
		return
	}

	// insert the resource at the determined index and update the sort order map
	newSet := make([]ezkube.ResourceId, len(r.set)+1)
	copy(newSet, r.set[:insertIndex])
	newSet[insertIndex] = objToInsert
	copy(newSet[insertIndex+1:], r.set[insertIndex:])
	r.set = newSet
}

// Delete removes all items from the set.
func (r Resources) Delete(items ...ezkube.ResourceId) Resources {
	for _, item := range items {
		key := Key(item)
		index, exists := r.sortIndex[key]
		if !exists {
			continue
		}
		// delete the key from the index and the set
		delete(r.sortIndex, key)
		newSet := make([]ezkube.ResourceId, len(r.set)-1)
		copy(newSet, r.set[:index])
		copy(newSet[index:], r.set[index+1:])
		r.set = newSet
	}
	return r
}

// HasAll returns true if and only if all items are contained in the set.
func (r Resources) HasAll(items ...ezkube.ResourceId) bool {
	for _, item := range items {
		if !r.Has(item) {
			return false
		}
	}
	return true
}

// HasAny returns true if any items are contained in the set.
func (r Resources) HasAny(items ...ezkube.ResourceId) bool {
	for _, item := range items {
		if r.Has(item) {
			return true
		}
	}
	return false
}

// Difference returns a set of objects that are not in s2
// For example:
// s1 = {a1, a2, a3}
// s2 = {a1, a2, a4, a5}
// s1.Difference(s2) = {a3}
// s2.Difference(s1) = {a4, a5}
func (r Resources) Difference(s2 Resources) Resources {
	result := Resources{}
	for _, key := range r.set {
		if !s2.Has(key) {
			result.Insert(key)
		}
	}
	return result
}

// Union returns a new set which includes items in either s1 or s2.
// For example:
// s1 = {a1, a2}
// s2 = {a3, a4}
// s1.Union(s2) = {a1, a2, a3, a4}
// s2.Union(s1) = {a1, a2, a3, a4}
func (r1 Resources) Union(s2 Resources) Resources {
	result := Resources{}
	for _, key := range r1.set {
		result.Insert(key)
	}
	for _, key := range s2.set {
		result.Insert(key)
	}
	return result
}

// Intersection returns a new set which includes items in BOTH s1 and s2
// For example:
// s1 = {a1, a2}
// s2 = {a2, a3}
// s1.Intersection(s2) = {a2}
func (r1 Resources) Intersection(r2 Resources) Resources {
	var walk, other Resources
	result := Resources{}
	if r1.Length() < r2.Length() {
		walk = r1
		other = r2
	} else {
		walk = r2
		other = r1
	}
	for _, key := range walk.set {
		if other.Has(key) {
			result.Insert(key)
		}
	}
	return result
}

// Has returns true if and only if item is contained in the set.
func (r Resources) Has(item ezkube.ResourceId) bool {
	_, contained := r.sortIndex[Key(item)]
	return contained
}

// IsSuperset returns true if and only if s1 is a superset of s2.
func (r1 Resources) IsSuperset(r2 Resources) bool {
	for _, item := range r2.set {
		if !r1.Has(item) {
			return false
		}
	}
	return true
}

// Equal returns true if and only if s1 is equal (as a set) to s2.
// Two sets are equal if their membership is identical.
// (In practice, this means same elements, order doesn't matter)
func (r1 Resources) Equal(r2 Resources) bool {
	return r1.Length() == r2.Length() && r1.IsSuperset(r2)
}

// Returns a single element from the set.
func (r Resources) PopAny() (ezkube.ResourceId, bool) {
	for _, key := range r.set {
		r.Delete(key)
		return key, true
	}
	return nil, false
}
