package output

import (
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/solo-io/skv2/contrib/pkg/sets"
)
// TODO remove
// represents a generic delta between two snapshots
type SnapshotDelta struct {
	Inserted map[schema.GroupVersionKind]sets.ResourceSet
	Removed  map[schema.GroupVersionKind]sets.ResourceSet
}

func NewSnapshotDelta() SnapshotDelta {
	return SnapshotDelta{
		Inserted: make(map[schema.GroupVersionKind]sets.ResourceSet),
		Removed:  make(map[schema.GroupVersionKind]sets.ResourceSet),
	}
}

func (d SnapshotDelta) AddInserted(gvk schema.GroupVersionKind, set sets.ResourceSet) {
	d.Inserted[gvk] = set
}

func (d SnapshotDelta) AddRemoved(gvk schema.GroupVersionKind, set sets.ResourceSet) {
	d.Removed[gvk] = set
}
