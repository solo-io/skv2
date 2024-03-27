package sets

import (
	"fmt"
	"sync"

	"github.com/rotisserie/eris"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
)

var NotFoundErr = func(resourceType ezkube.ResourceId, id ezkube.ResourceId) error {
	return eris.Errorf("%T with id %v not found", resourceType, Key(id))
}

const defaultSeparator = "."

// k8s resources are uniquely identified by their name and namespace
func Key(id ezkube.ResourceId) string {
	return ezkube.KeyWithSeparator(id, defaultSeparator)
}

// typed keys are helpful for logging; currently unused in the Set implementation but placed here for convenience
func TypedKey(id ezkube.ResourceId) string {
	return fmt.Sprintf("%s.%T", Key(id), id)
}

type ResourceSet interface {
	Keys() sets.String
	List(filterResource ...func(ezkube.ResourceId) bool) []ezkube.ResourceId
	UnsortedList(filterResource ...func(ezkube.ResourceId) bool) []ezkube.ResourceId
	Map() Resources
	Insert(resource ...ezkube.ResourceId)
	Equal(set ResourceSet) bool
	Has(resource ezkube.ResourceId) bool
	Delete(resource ezkube.ResourceId)
	Union(set ResourceSet) ResourceSet
	Difference(set ResourceSet) ResourceSet
	Intersection(set ResourceSet) ResourceSet
	IsSuperset(set ResourceSet) bool
	Find(resourceType, id ezkube.ResourceId) (ezkube.ResourceId, error)
	Length() int
	// returns the delta between this and and another ResourceSet
	Delta(newSet ResourceSet) ResourceDelta
	// Clone returns a deep copy of the set
	Clone() ResourceSet
}

// ResourceDelta represents the set of changes between two ResourceSets.
type ResourceDelta struct {
	// the resources inserted into the set
	Inserted ResourceSet
	// the resources removed from the set
	Removed ResourceSet
}

type threadSafeResourceSet struct {
	lock sync.RWMutex
	set  Resources
}

func NewResourceSet(
	sortFunc func(toInsert, existing interface{}) bool,
	equalityFunc func(a, b interface{}) int,
	resources ...ezkube.ResourceId,
) ResourceSet {
	return &threadSafeResourceSet{
		set: newResources(sortFunc, equalityFunc, resources...),
	}
}

func (t *threadSafeResourceSet) Keys() sets.String {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.set.Keys()
}

func (t *threadSafeResourceSet) List(filterResource ...func(ezkube.ResourceId) bool) []ezkube.ResourceId {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.set.List(filterResource...)
}

func (t *threadSafeResourceSet) UnsortedList(filterResource ...func(ezkube.ResourceId) bool) []ezkube.ResourceId {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.set.UnsortedList(filterResource...)
}

func (t *threadSafeResourceSet) Map() Resources {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.set
}

func (t *threadSafeResourceSet) Insert(resources ...ezkube.ResourceId) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.set.Insert(resources...)
}

func (t *threadSafeResourceSet) Has(resource ezkube.ResourceId) bool {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.set.Has(resource)
}

func (t *threadSafeResourceSet) IsSuperset(set ResourceSet) bool {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.set.IsSuperset(set.Map())
}

func (t *threadSafeResourceSet) Equal(set ResourceSet) bool {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.set.Equal(set.Map())
}

func (t *threadSafeResourceSet) Delete(resource ezkube.ResourceId) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.set.Delete(resource)
}

func (t *threadSafeResourceSet) Union(set ResourceSet) ResourceSet {
	t.lock.Lock()
	defer t.lock.Unlock()
	return &threadSafeResourceSet{set: t.set.Union(set.Map())}
}

func (t *threadSafeResourceSet) Difference(set ResourceSet) ResourceSet {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return &threadSafeResourceSet{set: t.set.Difference(set.Map())}
}

func (t *threadSafeResourceSet) Intersection(set ResourceSet) ResourceSet {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return &threadSafeResourceSet{set: t.set.Intersection(set.Map())}
}

func (t *threadSafeResourceSet) Find(
	resourceType,
	id ezkube.ResourceId,
) (ezkube.ResourceId, error) {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.set.Find(resourceType, id)
}

func (t *threadSafeResourceSet) Length() int {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.set.Length()
}

// note that this function will currently panic if called for a ResourceSet containing non-runtime.Objects
func (t *threadSafeResourceSet) Delta(newSet ResourceSet) ResourceDelta {
	return t.set.Delta(newSet)
}

// Create a clone of the current set
// note that this function will currently panic if called for a ResourceSet containing non-runtime.Objects
func (t *threadSafeResourceSet) Clone() ResourceSet {
	return t.set.Clone()
}

// // must have GOEXPERIMENT=rangefunc enabled
// // example -> for k, v := r.All2 { ... }
// func (r *threadSafeResourceSet) All2() iter.Seq2[int, ezkube.ResourceId] {
// 	return func(yield func(int, ezkube.ResourceId) bool) {
// 		for i, resource := range r.set.set {
// 			if !yield(i, resource) {
// 				break
// 			}
// 		}
// 	}
// }

// // must have GOEXPERIMENT=rangefunc enabled
// // example -> for v := r.All1 { ... }
// func (r *threadSafeResourceSet) All1() iter.Seq[ezkube.ResourceId] {
// 	return func(yield func(ezkube.ResourceId) bool) {
// 		for _, resource := range r.set.set {
// 			if !yield(resource) {
// 				break
// 			}
// 		}
// 	}
// }
