package sets

import (
	"fmt"
	"strings"
	"sync"

	"github.com/rotisserie/eris"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
)

var NotFoundErr = func(resourceType ezkube.ResourceId, id ezkube.ResourceId) error {
	return eris.Errorf("%T with id %v not found", resourceType, Key(id))
}

const separator = "."

var builderPool = sync.Pool{
	New: func() any {
		return &strings.Builder{}
	},
}

// k8s resources are uniquely identified by their name and namespace
func Key(id ezkube.ResourceId) string {
	b := builderPool.Get().(*strings.Builder)
	defer func() {
		b.Reset()
		builderPool.Put(b)
	}()
	// When kubernetes objects are passed in here, a call to the GetX() functions will panic, so
	// this will return "<unknown>" always if the input is nil.
	if id == nil {
		return "<unknown>"
	}
	b.WriteString(id.GetName())
	b.WriteString(separator)
	b.WriteString(id.GetNamespace())
	b.WriteString(separator)
	if clusterId, ok := id.(ezkube.ClusterResourceId); ok {
		b.WriteString(separator)
		b.WriteString(ezkube.GetClusterName(clusterId))
	}
	return b.String()
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

func NewResourceSet(resources ...ezkube.ResourceId) ResourceSet {
	return &threadSafeResourceSet{set: newResources(resources...)}
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
	return t.set.Map()
}

func (t *threadSafeResourceSet) Insert(
	resources ...ezkube.ResourceId,
) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.set.Insert(resources...)
}

func (t *threadSafeResourceSet) Has(resource ezkube.ResourceId) bool {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.set.Has(resource)
}

// Has returns true if and only if item is contained in the set.
func (s Resources) Has(item ezkube.ResourceId) bool {
	_, contained := s[Key(item)]
	return contained
}

func (t *threadSafeResourceSet) IsSuperset(
	set ResourceSet,
) bool {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.set.IsSuperset(set.Map())
}

func (t *threadSafeResourceSet) Equal(
	set ResourceSet,
) bool {
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
	return t.set.Len()
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
