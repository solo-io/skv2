package sets

import (
	"sort"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/ezkube"
	"golang.org/x/exp/maps"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// sets.Resources is a set of strings, implemented via map[string]struct{} for minimal memory consumption.
type Resources map[string]ezkube.ResourceId

func newResources(resources ...ezkube.ResourceId) Resources {
	mapping := make(Resources, len(resources))
	for _, resource := range resources {
		if resource == nil {
			continue
		}
		mapping.Insert(resource)
	}
	return mapping
}

func (r Resources) Find(resourceType, id ezkube.ResourceId) (ezkube.ResourceId, error) {
	key := Key(id)
	resource, ok := r[key]
	if !ok {
		return nil, NotFoundErr(resourceType, id)
	}

	return resource, nil
}

func (r Resources) Length() int {
	return r.Len()
}

func (r Resources) Delta(newSet ResourceSet) ResourceDelta {
	updated, removed := NewResourceSet(), NewResourceSet()

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
	new := NewResourceSet()
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
	for key, _ := range r {
		keys.Insert(key)
	}
	return keys
}

func (r Resources) List(filterResource ...func(ezkube.ResourceId) bool) []ezkube.ResourceId {
	res := make(sortableSliceOfString, 0, len(r))
	for key := range r {
		res = append(res, key)
	}
	sort.Sort(res)
	var resources []ezkube.ResourceId
	for _, key := range res {
		var filtered bool
		for _, filter := range filterResource {
			if filter(r[key]) {
				filtered = true
				break
			}
		}
		if !filtered {
			resources = append(resources, r[key])
		}
	}
	return resources
}

func (r Resources) UnsortedList(filterResource ...func(ezkube.ResourceId) bool) []ezkube.ResourceId {
	resources := make([]ezkube.ResourceId, 0, len(r))
	for _, val := range r {
		var filtered bool
		for _, filter := range filterResource {
			if filter(val) {
				filtered = true
				break
			}
		}
		if !filtered {
			resources = append(resources, val)
		}
	}
	return resources
}

func (r Resources) Map() Resources {
	shallowCopy := make(map[string]ezkube.ResourceId, len(r))
	maps.Copy(shallowCopy, r)
	return shallowCopy
}

func (r Resources) Insert(resources ...ezkube.ResourceId) {
	for _, resource := range resources {
		if resource == nil {
			continue
		}
		key := Key(resource)
		r[key] = resource
	}
}

// Delete removes all items from the set.
func (s Resources) Delete(items ...ezkube.ResourceId) Resources {
	for _, item := range items {
		delete(s, Key(item))
	}
	return s
}

// HasAll returns true if and only if all items are contained in the set.
func (s Resources) HasAll(items ...ezkube.ResourceId) bool {
	for _, item := range items {
		if !s.Has(item) {
			return false
		}
	}
	return true
}

// HasAny returns true if any items are contained in the set.
func (s Resources) HasAny(items ...ezkube.ResourceId) bool {
	for _, item := range items {
		if s.Has(item) {
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
func (s Resources) Difference(s2 Resources) Resources {
	result := Resources{}
	for _, key := range s {
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
func (s1 Resources) Union(s2 Resources) Resources {
	result := Resources{}
	for _, key := range s1 {
		result.Insert(key)
	}
	for _, key := range s2 {
		result.Insert(key)
	}
	return result
}

// Intersection returns a new set which includes the item in BOTH s1 and s2
// For example:
// s1 = {a1, a2}
// s2 = {a2, a3}
// s1.Intersection(s2) = {a2}
func (s1 Resources) Intersection(s2 Resources) Resources {
	var walk, other Resources
	result := Resources{}
	if s1.Len() < s2.Len() {
		walk = s1
		other = s2
	} else {
		walk = s2
		other = s1
	}
	for _, key := range walk {
		if other.Has(key) {
			result.Insert(key)
		}
	}
	return result
}

// IsSuperset returns true if and only if s1 is a superset of s2.
func (s1 Resources) IsSuperset(s2 Resources) bool {
	for _, item := range s2 {
		if !s1.Has(item) {
			return false
		}
	}
	return true
}

// Equal returns true if and only if s1 is equal (as a set) to s2.
// Two sets are equal if their membership is identical.
// (In practice, this means same elements, order doesn't matter)
func (s1 Resources) Equal(s2 Resources) bool {
	return len(s1) == len(s2) && s1.IsSuperset(s2)
}

type sortableSliceOfString []string

func (s sortableSliceOfString) Len() int           { return len(s) }
func (s sortableSliceOfString) Less(i, j int) bool { return lessString(s[i], s[j]) }
func (s sortableSliceOfString) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// Returns a single element from the set.
func (s Resources) PopAny() (ezkube.ResourceId, bool) {
	for _, key := range s {
		s.Delete(key)
		return key, true
	}
	return nil, false
}

// Len returns the size of the set.
func (s Resources) Len() int {
	return len(s)
}

func lessString(lhs, rhs string) bool {
	return lhs < rhs
}
