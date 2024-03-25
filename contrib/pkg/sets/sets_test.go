// This file is a copy paste of the sets package from k8s.io/apimachinery/pkg/util/sets
// It makes sure that the behavior is the exact same, without needing a separate allocated set.

package sets

import (
	"testing"

	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	"github.com/solo-io/skv2/pkg/ezkube"
	"google.golang.org/protobuf/proto"
)

func TestResourcesSet(t *testing.T) {
	s := Resources{}
	s2 := Resources{}
	if s.Length() != 0 {
		t.Errorf("Expected len=0: %d", s.Length())
	}
	s.Insert(&v1.ObjectRef{Name: "a"}, &v1.ObjectRef{Name: "b"})
	if s.Length() != 2 {
		t.Errorf("Expected len=2: %d", s.Length())
	}
	s.Insert(&v1.ObjectRef{Name: "c"})
	if s.Has(&v1.ObjectRef{Name: "d"}) {
		t.Errorf("Unexpected contents: %#v", s)
	}
	if !s.Has(&v1.ObjectRef{Name: "a"}) {
		t.Errorf("Missing contents: %#v", s)
	}
	s.Delete(&v1.ObjectRef{Name: "a"})
	if s.Has(&v1.ObjectRef{Name: "a"}) {
		t.Errorf("Unexpected contents: %#v", s)
	}
	s.Insert(&v1.ObjectRef{Name: "a"})
	if s.HasAll(&v1.ObjectRef{Name: "a"}, &v1.ObjectRef{Name: "b"}, &v1.ObjectRef{Name: "d"}) {
		t.Errorf("Unexpected contents: %#v", s)
	}
	if !s.HasAll(&v1.ObjectRef{Name: "a"}, &v1.ObjectRef{Name: "b"}) {
		t.Errorf("Missing contents: %#v", s)
	}
	s2.Insert(&v1.ObjectRef{Name: "a"}, &v1.ObjectRef{Name: "b"}, &v1.ObjectRef{Name: "d"})
	if s.IsSuperset(s2) {
		t.Errorf("Unexpected contents: %#v", s)
	}
	s2.Delete(&v1.ObjectRef{Name: "d"})
	if !s.IsSuperset(s2) {
		t.Errorf("Missing contents: %#v", s)
	}
}

func TestResourcesSetDeleteMultiples(t *testing.T) {
	s := Resources{}
	s.Insert(&v1.ObjectRef{Name: "a"}, &v1.ObjectRef{Name: "b"}, &v1.ObjectRef{Name: "c"})
	if s.Length() != 3 {
		t.Errorf("Expected len=3: %d", s.Length())
	}

	s.Delete(&v1.ObjectRef{Name: "a"}, &v1.ObjectRef{Name: "c"})
	if s.Length() != 1 {
		t.Errorf("Expected len=1: %d", s.Length())
	}
	if s.Has(&v1.ObjectRef{Name: "a"}) {
		t.Errorf("Unexpected contents: %#v", s)
	}
	if s.Has(&v1.ObjectRef{Name: "c"}) {
		t.Errorf("Unexpected contents: %#v", s)
	}
	if !s.Has(&v1.ObjectRef{Name: "b"}) {
		t.Errorf("Missing contents: %#v", s)
	}

}

func TestNewStringSet(t *testing.T) {
	s := newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds, &v1.ObjectRef{Name: "a"}, &v1.ObjectRef{Name: "b"}, &v1.ObjectRef{Name: "c"})
	if s.Length() != 3 {
		t.Errorf("Expected len=3: %d", s.Length())
	}
	if !s.Has(&v1.ObjectRef{Name: "a"}) || !s.Has(&v1.ObjectRef{Name: "b"}) || !s.Has(&v1.ObjectRef{Name: "c"}) {
		t.Errorf("Unexpected contents: %#v", s)
	}
}
func TestResourcesSetList(t *testing.T) {
	s := newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds,
		&v1.ObjectRef{Name: "z"},
		&v1.ObjectRef{Name: "y"},
		&v1.ObjectRef{Name: "x"},
		&v1.ObjectRef{Name: "a"},
	)
	list := s.List()
	expected := []*v1.ObjectRef{
		{Name: "a"},
		{Name: "x"},
		{Name: "y"},
		{Name: "z"},
	}
	for idx := range list {
		if !proto.Equal(list[idx].(proto.Message), expected[idx]) {
			t.Errorf("List gave unexpected result: %#v", s.List())
		}
	}
}

func TestResourcesSetDifference(t *testing.T) {
	a := newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds, &v1.ObjectRef{Name: "1"}, &v1.ObjectRef{Name: "2"}, &v1.ObjectRef{Name: "3"})
	b := newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds,
		&v1.ObjectRef{Name: "1"},
		&v1.ObjectRef{Name: "2"},
		&v1.ObjectRef{Name: "4"},
		&v1.ObjectRef{Name: "5"},
	)
	c := a.Difference(b)
	d := b.Difference(a)
	if c.Length() != 1 {
		t.Errorf("Expected len=1: %d", c.Length())
	}
	if !c.Has(&v1.ObjectRef{Name: "3"}) {
		t.Errorf("Unexpected contents: %#v", c.List())
	}
	if d.Length() != 2 {
		t.Errorf("Expected len=2: %d", d.Length())
	}
	if !d.Has(&v1.ObjectRef{Name: "4"}) || !d.Has(&v1.ObjectRef{Name: "5"}) {
		t.Errorf("Unexpected contents: %#v", d.List())
	}
}

func TestResourcesSetHasAny(t *testing.T) {
	a := newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds, &v1.ObjectRef{Name: "1"}, &v1.ObjectRef{Name: "2"}, &v1.ObjectRef{Name: "3"})

	if !a.HasAny(&v1.ObjectRef{Name: "1"}, &v1.ObjectRef{Name: "4"}) {
		t.Errorf("expected true, got false")
	}

	if a.HasAny(&v1.ObjectRef{Name: "0"}, &v1.ObjectRef{Name: "4"}) {
		t.Errorf("expected false, got true")
	}
}

func TestResourcesSetEquals(t *testing.T) {
	// Simple case (order doesn't matter)
	a := newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds, &v1.ObjectRef{Name: "1"}, &v1.ObjectRef{Name: "2"})
	b := newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds, &v1.ObjectRef{Name: "2"}, &v1.ObjectRef{Name: "1"})
	if !a.Equal(b) {
		t.Errorf("Expected to be equal: %v vs %v", a, b)
	}

	// It is a set; duplicates are ignored
	b = newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds, &v1.ObjectRef{Name: "2"}, &v1.ObjectRef{Name: "2"}, &v1.ObjectRef{Name: "1"})
	if !a.Equal(b) {
		t.Errorf("Expected to be equal: %v vs %v", a, b)
	}

	// Edge cases around empty sets / empty strings
	a = newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds)
	b = newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds)
	if !a.Equal(b) {
		t.Errorf("Expected to be equal: %v vs %v", a, b)
	}

	b = newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds, &v1.ObjectRef{Name: "1"}, &v1.ObjectRef{Name: "2"}, &v1.ObjectRef{Name: "3"})
	if a.Equal(b) {
		t.Errorf("Expected to be not-equal: %v vs %v", a, b)
	}

	b = newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds, &v1.ObjectRef{Name: "1"}, &v1.ObjectRef{Name: "2"}, &v1.ObjectRef{})
	if a.Equal(b) {
		t.Errorf("Expected to be not-equal: %v vs %v", a, b)
	}

	// Check for equality after mutation
	a = newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds)
	a.Insert(&v1.ObjectRef{Name: "1"})
	if a.Equal(b) {
		t.Errorf("Expected to be not-equal: %v vs %v", a, b)
	}

	a.Insert(&v1.ObjectRef{Name: "2"})
	if a.Equal(b) {
		t.Errorf("Expected to be not-equal: %v vs %v", a, b)
	}

	a.Insert(&v1.ObjectRef{})
	if !a.Equal(b) {
		t.Errorf("Expected to be equal: %v vs %v", a, b)
	}

	a.Delete(&v1.ObjectRef{})
	if a.Equal(b) {
		t.Errorf("Expected to be not-equal: %v vs %v", a, b)
	}
}

func TestResourcesUnion(t *testing.T) {
	tests := []struct {
		s1       Resources
		s2       Resources
		expected Resources
	}{
		{
			newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds,
				&v1.ObjectRef{Name: "1"},
				&v1.ObjectRef{Name: "2"},
				&v1.ObjectRef{Name: "3"},
				&v1.ObjectRef{Name: "4"},
			),
			newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds,
				&v1.ObjectRef{Name: "3"},
				&v1.ObjectRef{Name: "4"},
				&v1.ObjectRef{Name: "5"},
				&v1.ObjectRef{Name: "6"},
			),
			newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds,
				&v1.ObjectRef{Name: "1"},
				&v1.ObjectRef{Name: "2"},
				&v1.ObjectRef{Name: "3"},
				&v1.ObjectRef{Name: "4"},
				&v1.ObjectRef{Name: "5"},
				&v1.ObjectRef{Name: "6"},
			),
		},
		{
			newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds,
				&v1.ObjectRef{Name: "1"},
				&v1.ObjectRef{Name: "2"},
				&v1.ObjectRef{Name: "3"},
				&v1.ObjectRef{Name: "4"},
			),
			newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds),
			newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds,
				&v1.ObjectRef{Name: "1"},
				&v1.ObjectRef{Name: "2"},
				&v1.ObjectRef{Name: "3"},
				&v1.ObjectRef{Name: "4"},
			),
		},
		{
			newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds),
			newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds,
				&v1.ObjectRef{Name: "1"},
				&v1.ObjectRef{Name: "2"},
				&v1.ObjectRef{Name: "3"},
				&v1.ObjectRef{Name: "4"},
			),
			newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds,
				&v1.ObjectRef{Name: "1"},
				&v1.ObjectRef{Name: "2"},
				&v1.ObjectRef{Name: "3"},
				&v1.ObjectRef{Name: "4"},
			),
		},
		{
			newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds),
			newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds),
			newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds),
		},
	}

	for _, test := range tests {
		union := test.s1.Union(test.s2)
		if union.Length() != test.expected.Length() {
			t.Errorf("Expected union.Length()=%d but got %d", test.expected.Length(), union.Length())
		}

		if !union.Equal(test.expected) {
			t.Errorf("Expected union.Equal(expected) but not true.  union:%v expected:%v", union.List(), test.expected.List())
		}
	}
}

func TestResourcesIntersection(t *testing.T) {
	tests := []struct {
		s1       Resources
		s2       Resources
		expected Resources
	}{
		{
			newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds,
				&v1.ObjectRef{Name: "1"},
				&v1.ObjectRef{Name: "2"},
				&v1.ObjectRef{Name: "3"},
				&v1.ObjectRef{Name: "4"},
			),
			newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds,
				&v1.ObjectRef{Name: "3"},
				&v1.ObjectRef{Name: "4"},
				&v1.ObjectRef{Name: "5"},
				&v1.ObjectRef{Name: "6"},
			),
			newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds, &v1.ObjectRef{Name: "3"}, &v1.ObjectRef{Name: "4"}),
		},
		{
			newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds,
				&v1.ObjectRef{Name: "1"},
				&v1.ObjectRef{Name: "2"},
				&v1.ObjectRef{Name: "3"},
				&v1.ObjectRef{Name: "4"},
			),
			newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds,
				&v1.ObjectRef{Name: "1"},
				&v1.ObjectRef{Name: "2"},
				&v1.ObjectRef{Name: "3"},
				&v1.ObjectRef{Name: "4"},
			),
			newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds,
				&v1.ObjectRef{Name: "1"},
				&v1.ObjectRef{Name: "2"},
				&v1.ObjectRef{Name: "3"},
				&v1.ObjectRef{Name: "4"},
			),
		},
		{
			newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds,
				&v1.ObjectRef{Name: "1"},
				&v1.ObjectRef{Name: "2"},
				&v1.ObjectRef{Name: "3"},
				&v1.ObjectRef{Name: "4"},
			),
			newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds),
			newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds),
		},
		{
			newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds),
			newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds,
				&v1.ObjectRef{Name: "1"},
				&v1.ObjectRef{Name: "2"},
				&v1.ObjectRef{Name: "3"},
				&v1.ObjectRef{Name: "4"},
			),
			newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds),
		},
		{
			newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds),
			newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds),
			newResources(ezkube.ResourceIdsAscending, ezkube.CompareResourceIds),
		},
	}

	for _, test := range tests {
		intersection := test.s1.Intersection(test.s2)
		if intersection.Length() != test.expected.Length() {
			t.Errorf("Expected intersection.Length()=%d but got %d", test.expected.Length(), intersection.Length())
		}

		if !intersection.Equal(test.expected) {
			t.Errorf(
				"Expected intersection.Equal(expected) but not true.  intersection:%v expected:%v",
				intersection.List(),
				test.expected.List(),
			)
		}
	}
}
