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
	sortFunc func(toInsert, existing interface{}) bool
	// compareFunc is the function used to compare two resources.
	// compareFunc should return 0 if a == b, -1 if a < b, and 1 if a > b.
	compareFunc func(a, b interface{}) int
}

func newResources(
	sortFunc func(toInsert, existing interface{}) bool,
	equalityFunc func(a, b interface{}) int,
	resources ...ezkube.ResourceId,
) Resources {
	r := Resources{
		set:         []ezkube.ResourceId{},
		sortFunc:    sortFunc,
		compareFunc: equalityFunc,
	}
	for _, resource := range resources {
		if resource == nil {
			continue
		}
		r.Insert(resource)
	}
	return r
}

func (r *Resources) Find(resourceType, id ezkube.ResourceId) (ezkube.ResourceId, error) {
	key := Key(id)
	var key_i string
	i := sort.Search(r.Length(), func(i int) bool {
		key_i = Key(r.set[i])
		return key <= key_i
	})
	if i != r.Length() && key_i == key {
		return r.set[i], nil
	}
	return nil, NotFoundErr(resourceType, id)
}

func (r *Resources) Length() int {
	return len(r.set)
}

func (r *Resources) Delta(newSet ResourceSet) ResourceDelta {
	updated, removed := NewResourceSet(r.sortFunc, r.compareFunc), NewResourceSet(r.sortFunc, r.compareFunc)

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
			if !r.Has(newObj) {
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

func (r *Resources) Clone() ResourceSet {
	new := NewResourceSet(r.sortFunc, r.compareFunc, r.set...)
	r.List(
		func(oldObj ezkube.ResourceId) bool {
			copy := oldObj.(client.Object).DeepCopyObject().(client.Object)
			new.Insert(copy)
			return true
		},
	)
	return new
}

func (r *Resources) Keys() sets.String {
	keys := sets.NewString()
	for _, key := range r.set {
		keys.Insert(Key(key))
	}
	return keys
}

func (r *Resources) List(filterResource ...func(ezkube.ResourceId) bool) []ezkube.ResourceId {
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

func (r *Resources) UnsortedList(filterResource ...func(ezkube.ResourceId) bool) []ezkube.ResourceId {
	return r.List(filterResource...)
}

func (r Resources) Map() map[string]ezkube.ResourceId {
	shallowCopy := make(map[string]ezkube.ResourceId, len(r.set))
	for _, obj := range r.set {
		shallowCopy[Key(obj)] = obj
	}
	return shallowCopy
}

// Insert adds items to the set.
// If an item is already in the set, it is overwritten.
// The set is sorted based on the sortFunc. If sortFunc is nil, the set will be unsorted.
func (r *Resources) Insert(resources ...ezkube.ResourceId) {
	for _, resource := range resources {
		insertIndex := sort.Search(r.Length(), func(i int) bool { return r.sortFunc(resource, r.set[i]) })

		// if the resource is already in the set, replace it
		if insertIndex < len(r.set) && r.compareFunc(resource, r.set[insertIndex]) == 0 {
			r.set[insertIndex] = resource
			return
		}
		if r.sortFunc == nil {
			r.set = append(r.set, resource)
			return
		}

		// insert the resource at the determined index
		newSet := make([]ezkube.ResourceId, len(r.set)+1)
		copy(newSet, r.set[:insertIndex])
		newSet[insertIndex] = resource
		copy(newSet[insertIndex+1:], r.set[insertIndex:])
		r.set = newSet
	}
}

// Delete removes all items from the set.
func (r *Resources) Delete(items ...ezkube.ResourceId) {
	for _, item := range items {
		i := sort.Search(r.Length(), func(i int) bool {
			return r.compareFunc(r.set[i], item) >= 0
		})
		if found := i < r.Length() && r.compareFunc(r.set[i], item) == 0; !found {
			continue
		}

		// remove item the set
		newSet := make([]ezkube.ResourceId, len(r.set)-1)
		copy(newSet, r.set[:i])
		copy(newSet[i:], r.set[i+1:])
		r.set = newSet
	}
}

// HasAll returns true if and only if all items are contained in the set.
func (r *Resources) HasAll(items ...ezkube.ResourceId) bool {
	for _, item := range items {
		if !r.Has(item) {
			return false
		}
	}
	return true
}

// HasAny returns true if any items are contained in the set.
func (r *Resources) HasAny(items ...ezkube.ResourceId) bool {
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
func (r *Resources) Difference(s2 Resources) Resources {
	result := Resources{
		set:         []ezkube.ResourceId{},
		sortFunc:    r.sortFunc,
		compareFunc: r.compareFunc,
	}
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
func (r1 *Resources) Union(s2 Resources) Resources {
	result := Resources{
		set:         []ezkube.ResourceId{},
		sortFunc:    r1.sortFunc,
		compareFunc: r1.compareFunc,
	}
	for _, resource := range r1.set {
		result.Insert(resource)
	}
	for _, resource := range s2.set {
		if !result.Has(resource) {
			result.Insert(resource)
		}
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
	result := Resources{
		set:         []ezkube.ResourceId{},
		sortFunc:    r1.sortFunc,
		compareFunc: r1.compareFunc,
	}
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
func (r *Resources) Has(item ezkube.ResourceId) bool {
	i := sort.Search(r.Length(), func(i int) bool {
		return ezkube.CompareResourceIds(r.set[i], item) >= 0
	})
	return i < r.Length() && r.compareFunc(r.set[i], item) == 0
}

// IsSuperset returns true if and only if s1 is a superset of s2.
func (r1 *Resources) IsSuperset(r2 Resources) bool {
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
func (r1 *Resources) Equal(r2 Resources) bool {
	return r1.Length() == r2.Length() && r1.IsSuperset(r2)
}

// Returns a single element from the set.
func (r *Resources) PopAny() (ezkube.ResourceId, bool) {
	for _, key := range r.set {
		r.Delete(key)
		return key, true
	}
	return nil, false
}
