package sets

import (
	"fmt"
	"testing"

	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
)

// Define a global variable to prevent compiler optimizations
var result interface{}

scale := 10000

func BenchmarkResources_Insert(b *testing.B) {
	var r *Resources
	for i := 0; i < b.N; i++ {
		r = newResources() // Assume newResources is corrected to initialize correctly
		for j := 0; j < scale; j++ {
			resource := &v1.ObjectRef{Namespace: "namespace", Name: "name" + fmt.Sprint(j)}
			r.Insert(resource)
		}
	}
	// Store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

func BenchmarkResources_Find(b *testing.B) {
	r := newResources()
	for j := 0; j < scale; j++ {
		resource := &v1.ObjectRef{Namespace: "namespace", Name: "name" + fmt.Sprint(j)}
		r.Insert(resource)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.Find(&v1.TypedObjectRef{}, &v1.ObjectRef{Namespace: "namespace", Name: "name500"})
	}
}

func BenchmarkResources_Delete(b *testing.B) {
	var r *Resources
	for i := 0; i < b.N; i++ {
		r = newResources()
		for j := 0; j < scale; j++ {
			resource := &v1.ObjectRef{Namespace: "namespace", Name: "name" + fmt.Sprint(j)}
			r.Insert(resource)
		}
		b.ResetTimer()
		for j := 0; j < scale; j++ {
			r.Delete(&v1.ObjectRef{Namespace: "namespace", Name: "name" + fmt.Sprint(j)})
		}
		b.StopTimer()
	}
	result = r
}


func BenchmarkResources_List(b *testing.B) {
	r := newResources()
	for j := 0; j < scale ; j++ {
		resource := &ezkube.GenericResourceId{Namespace: "namespace", Name: "name" + string(j)}
		r.Insert(resource)
	}

	b.ResetTimer()
	var res []ezkube.ResourceId
	for i := 0; i < b.N; i++ {
		res = r.List()
	}
	// Store the result to prevent the compiler from optimizing the loop away.
	result = res
}

func BenchmarkResources_UnsortedList(b *testing.B) {
	r := newResources()
	for j := 0; j < scale; j++ {
		resource := &ezkube.GenericResourceId{Namespace: "namespace", Name: "name" + string(j)}
		r.Insert(resource)
	}

	b.ResetTimer()
	var res []ezkube.ResourceId
	for i := 0; i < b.N; i++ {
		res = r.UnsortedList()
	}
	// Store the result to prevent the compiler from optimizing the loop away.
	result = res
}