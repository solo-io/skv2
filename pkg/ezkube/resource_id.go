package ezkube

import (
	"strings"
	"sync"

	"github.com/rotisserie/eris"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const ClusterAnnotation = "cluster.solo.io/cluster"

var (
	builderPool = sync.Pool{
		New: func() any {
			return &strings.Builder{}
		},
	}
	// global lock for get/set cluster name annotation
	clusterNameLock = &sync.Mutex{}
)

// ResourceId represents a global identifier for a k8s resource.
type ResourceId interface {
	GetName() string
	GetNamespace() string
}

// ClusterResourceId represents a global identifier for a k8s resource.
type ClusterResourceId interface {
	GetName() string
	GetNamespace() string
	GetAnnotations() map[string]string
}

// internal struct that satisfies ResourceId interface
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

// internal struct that satisfies ClusterResourceId interface
type clusterResourceId struct {
	name, namespace string
	annotations     map[string]string
}

func (c clusterResourceId) GetName() string {
	return c.name
}

func (c clusterResourceId) GetNamespace() string {
	return c.namespace
}

func (c clusterResourceId) GetAnnotations() map[string]string {
	return c.annotations
}

// this is specifically to support k8s <= 1.23
type deprecatedClusterResourceId interface {
	GetName() string
	GetNamespace() string
	GetClusterName() string
}

// this is specifically to support k8s 1.24
type deprecatedZZZClusterResourceId interface {
	GetName() string
	GetNamespace() string
	GetZZZ_DeprecatedClusterName() string
}

// ConvertRefToId converts a ClusterObjectRef to a struct that implements the ClusterResourceId interface.
// Will not set an empty cluster name over an existing cluster name.
func ConvertRefToId(ref deprecatedClusterResourceId) ClusterResourceId {
	// if ref already stores annotations then we need to store the updates
	annotations := map[string]string{}

	cri, isCri := ref.(ClusterResourceId)
	if isCri {
		annotations = cri.GetAnnotations()
	}

	clusterName := ref.GetClusterName()
	if clusterName != "" {
		if isCri {
			clusterNameLock.Lock()
			annotations[ClusterAnnotation] = clusterName
			clusterNameLock.Unlock()
		} else {
			annotations[ClusterAnnotation] = clusterName
		}
	}

	return clusterResourceId{
		name:        ref.GetName(),
		namespace:   ref.GetNamespace(),
		annotations: annotations,
	}
}

func getDeprecatedClusterName(id ResourceId) string {
	if depResourceId, ok := id.(deprecatedClusterResourceId); ok {
		return depResourceId.GetClusterName()
	} else if depZZZResourceId, ok := id.(deprecatedZZZClusterResourceId); ok {
		return depZZZResourceId.GetZZZ_DeprecatedClusterName()
	}
	return ""
}

// GetClusterName returns the cluster name defined in the given object's cluster annotation.
// If annotations is nil or the cluster annotation value is empty, then the object is checked for the deprecated metadata cluster fields.
func GetClusterName(id ClusterResourceId) string {
	annotations := id.GetAnnotations()

	clusterNameLock.Lock()
	clusterName := annotations[ClusterAnnotation]
	clusterNameLock.Unlock()

	if annotations == nil {
		return getDeprecatedClusterName(id)
	}
	if clusterName == "" {
		return getDeprecatedClusterName(id)
	}

	return clusterName
}

func SetClusterName(obj client.Object, cluster string) {
	if obj.GetAnnotations() == nil {
		obj.SetAnnotations(map[string]string{})
	}

	annotations := obj.GetAnnotations()

	clusterNameLock.Lock()
	annotations[ClusterAnnotation] = cluster
	clusterNameLock.Unlock()
}

// KeyWithSeparator constructs a string consisting of the field values of the given resource id,
// separated by the given separator. It can be used to construct a unique key for a resource.
func KeyWithSeparator(id ResourceId, separator string) string {
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
	if clusterId, ok := id.(ClusterResourceId); ok {
		clusterNameByAnnotation := GetClusterName(clusterId)
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

// ResourceIdFromKeyWithSeparator converts a key back into a ResourceId, using the given separator.
// Returns an error if it cannot be converted.
func ResourceIdFromKeyWithSeparator(key string, separator string) (ResourceId, error) {
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
				ClusterAnnotation: parts[2],
			},
		}, nil
	} else {
		return nil, eris.Errorf("could not convert key %s with separator %s into resource id; unexpected number of parts: %d", key, separator, len(parts))
	}
}

// ResourceIdsCompare compares two ResourceId instances, first by name, then by namespace, and finally by cluster name.
// If the names or namespaces differ, the comparison returns the result of strings.Compare on those values.
// If the names and namespaces are the same, it attempts to cast the ResourceId instances to ClusterResourceId
// and compares the cluster names. If the cast fails, it falls back to using the deprecated cluster name retrieval.
func ResourceIdsCompare(a, b ResourceId) int {
	// compare names
	if cmp := strings.Compare(a.GetName(), b.GetName()); cmp != 0 {
		return cmp
	}

	// compare namespaces
	if cmp := strings.Compare(a.GetNamespace(), b.GetNamespace()); cmp != 0 {
		return cmp
	}

	// compare cluster names
	// attempt to cast to ClusterResourceId
	// if fails, attempt cast to deprecatedClusterResourceId since we might be working with a ClusterObjectRef
	var (
		aCluster, bCluster string
	)

	if a_cri, ok := a.(ClusterResourceId); ok {
		aCluster = GetClusterName(a_cri)
	} else {
		aCluster = getDeprecatedClusterName(a)
	}

	if b_cri, ok := b.(ClusterResourceId); ok {
		bCluster = GetClusterName(b_cri)
	} else {
		bCluster = getDeprecatedClusterName(b)
	}

	return strings.Compare(aCluster, bCluster)
}
