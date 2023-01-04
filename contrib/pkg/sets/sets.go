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

const defaultSeparator = "."

var builderPool = sync.Pool{
	New: func() any {
		return &strings.Builder{}
	},
}

// internal struct that satisfies ezkube.ResourceId interface
type resourceId struct {
	name      string
	namespace string
}

func (id resourceId) GetName() string {
	return id.name
}

func (id resourceId) GetNamespace() string {
	return id.namespace
}

// internal struct that satisfies ezkube.ClusterResourceId interface
type clusterResourceId struct {
	name        string
	namespace   string
	annotations map[string]string
}

func (id clusterResourceId) GetName() string {
	return id.name
}

func (id clusterResourceId) GetNamespace() string {
	return id.namespace
}

func (id clusterResourceId) GetAnnotations() map[string]string {
	return id.annotations
}

// k8s resources are uniquely identified by their name and namespace
// Key constructs a string consisting of the field values of the given resource id, separated by '.'
func Key(id ezkube.ResourceId) string {
	return KeyWithSeparator(id, defaultSeparator)
}

// KeyWithSeparator constructs a string consisting of the field values of the given resource id, separated by
// the given separator
func KeyWithSeparator(id ezkube.ResourceId, separator string) string {
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
	// handle the possibility that clusterName could be set either as an annotation (new way)
	// or as a field (old way pre-k8s 1.25)
	if clusterId, ok := id.(ezkube.ClusterResourceId); ok {
		clusterNameByAnnotation := ezkube.GetClusterName(clusterId)
		if clusterNameByAnnotation != "" {
			b.WriteString(clusterNameByAnnotation)
			return b.String()
		}
	}
	if deprecatedClusterId, ok := id.(interface{ GetClusterName() string }); ok {
		b.WriteString(deprecatedClusterId.GetClusterName())
	}
	return b.String()
}

// ResourceIdFromKeyWithSeparator converts a key back into a ResourceId, using the separator '.'
// Returns an error if it cannot be converted.
func ResourceIdFromKey(key string) (ezkube.ResourceId, error) {
	return ResourceIdFromKeyWithSeparator(key, defaultSeparator)
}

// ResourceIdFromKeyWithSeparator converts a key back into a ResourceId, using the given separator.
// Returns an error if it cannot be converted.
func ResourceIdFromKeyWithSeparator(key string, separator string) (ezkube.ResourceId, error) {
	parts := strings.Split(key, separator)
	if len(parts) == 2 {
		return resourceId{
			name:      parts[0],
			namespace: parts[1],
		}, nil

	} else if len(parts) == 3 {
		return clusterResourceId{
			name:      parts[0],
			namespace: parts[1],
			annotations: map[string]string{
				ezkube.ClusterAnnotation: parts[2],
			},
		}, nil
	} else {
		return nil, eris.Errorf("could not convert key %s with separator %s into resource id; unexpected number of parts: %d", key, separator, len(parts))
	}
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
