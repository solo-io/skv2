package sets

import (
	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// sets.Resources is a set of strings, implemented via map[string]struct{} for minimal memory consumption.
type Resources struct {
	l []ezkube.ResourceId
	// map of hash to index
	m map[uint64]uint64
}

func newResources(resources ...ezkube.ResourceId) *Resources {
	m := &Resources{
		l: []ezkube.ResourceId{},
		m: make(map[uint64]uint64, len(resources)),
	}
	for idx, resource := range resources {
		if resource == nil {
			continue
		}
		hash := Hash(resource)
		if _, ok := m.m[hash]; ok {
			continue
		}
		m.l = append(m.l, resource)
		m.m[Hash(resource)] = uint64(idx)
	}
	return m
}

func (r *Resources) Find(resourceType, id ezkube.ResourceId) (ezkube.ResourceId, error) {
	key := Hash(id)
	if idx, ok := r.m[key]; ok {
		return r.l[idx], nil
	}
	return nil, NotFoundErr(resourceType, id)
}

func (r *Resources) Length() int {
	return r.Len()
}

func (r *Resources) Delta(newSet ResourceSet) ResourceDelta {
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

func (r *Resources) Clone() ResourceSet {
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

func (r *Resources) Keys() sets.String {
	keys := sets.NewString()
	for _, e := range r.l {
		key := Key(e)
		keys.Insert(key)
	}
	return keys
}

func (r *Resources) List(filterResource ...func(ezkube.ResourceId) bool) []ezkube.ResourceId {
	var resources []ezkube.ResourceId
	for _, e := range r.l {
		var filtered bool
		for _, filter := range filterResource {
			if filter(e) {
				filtered = true
				break
			}
		}
		if !filtered {
			resources = append(resources, e)
		}
	}
	return resources
}

func (r *Resources) UnsortedList(filterResource ...func(ezkube.ResourceId) bool) []ezkube.ResourceId {
	var resources []ezkube.ResourceId
	for _, e := range r.l {
		var filtered bool
		for _, filter := range filterResource {
			if filter(e) {
				filtered = true
				break
			}
		}
		if !filtered {
			resources = append(resources, e)
		}
	}
	return resources
}

func (r *Resources) Map() map[string]ezkube.ResourceId {
	res := make(map[string]ezkube.ResourceId)
	for _, resource := range r.l {
		if resource == nil {
			continue
		}
		res[Key(resource)] = resource
	}
	return res
}

func (r *Resources) Set() *Resources {
	return r
}

func (r *Resources) Insert(resources ...ezkube.ResourceId) {
	if r.m == nil {
		r.m = make(map[uint64]uint64, len(resources))
	}
	for _, resource := range resources {
		if resource == nil {
			continue
		}
		hash := Hash(resource)
		if idx, ok := r.m[hash]; ok {
			r.l[idx] = resource
		} else {
			r.l = append(r.l, resource)

			r.m[hash] = uint64(len(r.l) - 1)
		}
	}
}

// Delete removes all items from the set.
func (s *Resources) Delete(items ...ezkube.ResourceId) *Resources {
	for _, item := range items {
		for idx, e := range s.l {
			if Key(e) == Key(item) {
				s.l = append((s.l)[:idx], (s.l)[idx+1:]...)
				delete(s.m, Hash(item))
			}
		}
	}
	return s
}

// HasAll returns true if and only if all items are contained in the set.
func (s *Resources) HasAll(items ...ezkube.ResourceId) bool {
	for _, item := range items {
		if !s.Has(item) {
			return false
		}
	}
	return true
}

// HasAny returns true if any items are contained in the set.
func (s *Resources) HasAny(items ...ezkube.ResourceId) bool {
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
func (s *Resources) Difference(s2 *Resources) *Resources {
	result := &Resources{}
	for _, key := range s.List() {
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
func (s1 *Resources) Union(s2 *Resources) *Resources {
	result := &Resources{}
	for _, key := range s1.List() {
		result.Insert(key)
	}
	for _, key := range s2.List() {
		if !s1.Has(key) {
			result.Insert(key)
		}
	}
	return result
}

// Intersection returns a new set which includes the item in BOTH s1 and s2
// For example:
// s1 = {a1, a2}
// s2 = {a2, a3}
// s1.Intersection(s2) = {a2}
func (s1 *Resources) Intersection(s2 *Resources) *Resources {
	var walk, other *Resources
	result := &Resources{}
	if s1.Len() < s2.Len() {
		walk = s1
		other = s2
	} else {
		walk = s2
		other = s1
	}
	for _, key := range walk.List() {
		if other.Has(key) {
			result.Insert(key)
		}
	}
	return result
}

// IsSuperset returns true if and only if s1 is a superset of s2.
func (s1 *Resources) IsSuperset(s2 *Resources) bool {
	for _, item := range s2.l {
		if !s1.Has(item) {
			return false
		}
	}
	return true
}

// Equal returns true if and only if s1 is equal (as a set) to s2.
// Two sets are equal if their membership is identical.
// (In practice, this means same elements, order doesn't matter)
func (s1 *Resources) Equal(s2 *Resources) bool {
	return s1.Len() == s2.Len() && s1.IsSuperset(s2)
}

type sortableSliceOfString []string

func (s sortableSliceOfString) Len() int           { return len(s) }
func (s sortableSliceOfString) Less(i, j int) bool { return lessString(s[i], s[j]) }
func (s sortableSliceOfString) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// Returns a single element from the set.
func (s *Resources) PopAny() (ezkube.ResourceId, bool) {
	for _, key := range s.List() {
		s.Delete(key)
		return key, true
	}
	return nil, false
}

// Len returns the size of the set.
func (s *Resources) Len() int {
	return len(s.l)
}

func lessString(lhs, rhs string) bool {
	return lhs < rhs
}
