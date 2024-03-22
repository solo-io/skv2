// This file is a copy paste of the sets package from k8s.io/apimachinery/pkg/util/sets
// It makes sure that the behavior is the exact same, without needing a separate allocated set.

package sets_test

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/rotisserie/eris"
	things_test_io_v1 "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1"
	v1sets "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func BenchmarkResourcesSet100(b *testing.B)    { benchmarkResourcesSet(100, b) }
func BenchmarkResourcesSet1000(b *testing.B)   { benchmarkResourcesSet(1000, b) }
func BenchmarkResourcesSet10000(b *testing.B)  { benchmarkResourcesSet(10000, b) }
func BenchmarkResourcesSet20000(b *testing.B)  { benchmarkResourcesSet(20000, b) }
func BenchmarkResourcesSet30000(b *testing.B)  { benchmarkResourcesSet(30000, b) }
func BenchmarkResourcesSet40000(b *testing.B)  { benchmarkResourcesSet(40000, b) }
func BenchmarkResourcesSet50000(b *testing.B)  { benchmarkResourcesSet(50000, b) }
func BenchmarkResourcesSet100000(b *testing.B) { benchmarkResourcesSet(100000, b) }

func benchmarkResourcesSet(count int, b *testing.B) {
	resources := make([]*things_test_io_v1.Paint, count)
	for i := 0; i < count; i++ {
		resources[i] = &things_test_io_v1.Paint{
			ObjectMeta: metav1.ObjectMeta{
				Name:              fmt.Sprintf("name-%d", i),
				Namespace:         fmt.Sprintf("namespace-%d", i),
				Annotations:       map[string]string{ezkube.ClusterAnnotation: "random-cluster"},
				CreationTimestamp: metav1.Time{Time: metav1.Now().Time.Add(time.Duration(i))},
			},
		}

	}

	for n := 0; n < b.N; n++ {
		s := v1sets.NewPaintSet(ezkube.CreationTimestampAscending, ezkube.CreationTimestampsCompare)
		// for _, resource := range resources {
		s.Insert(resources...)
		// }
		l := s.List()
		// l := s.List(filterResource)
		// SortByCreationTime(l) // only for map implementation
		for _, r := range l {
			r.GetName()
		}
	}
}

func filterResource(resource *things_test_io_v1.Paint) bool {
	i, _ := strconv.Atoi(strings.Split(resource.GetName(), "-")[1])
	return i < 20001
}

// SortByCreationTime accepts a slice of client.Object instances and sorts it by creation timestamp in ascending order.
// It panics if the argument isn't a slice, or if it is a slice of a type that does not implement client.Object.
func SortByCreationTime(objs interface{}) {

	// Validate the argument
	if err := validate(objs); err != nil {
		panic(err)
	}

	// If we got past validation, the argument is either an empty slice or a slice of client.Object.
	// In the former case the comparison function will not be invoked, in the latter we can safely use getCreationTime.
	sort.SliceStable(objs, func(i, j int) bool {
		iTime := getCreationTime(objs, i)
		jTime := getCreationTime(objs, j)

		if iTime.Equal(&jTime) {
			iWorkload := getObjectName(objs, i)
			jWorkload := getObjectName(objs, j)
			return strings.Compare(iWorkload, jWorkload) > 0
		}

		return iTime.Before(&jTime)
	})
}
func validate(objs interface{}) error {
	s := reflect.ValueOf(objs)

	if s.Kind() != reflect.Slice {
		return eris.Errorf("argument must be a slice")
	}

	if s.IsNil() {
		// Zero value slice
		return nil
	}

	for i := 0; i < s.Len(); i++ {
		el := s.Index(i).Interface()
		if _, ok := el.(client.Object); !ok {
			return eris.Errorf("input slice contains element of unexpected type %T; all elements must implement client.Object", el)
		}
	}

	return nil
}
func getCreationTime(objs interface{}, index int) metav1.Time {
	return reflect.ValueOf(objs).Index(index).Interface().(client.Object).GetCreationTimestamp()
}
func getObjectName(objs interface{}, index int) string {
	obj := reflect.ValueOf(objs).Index(index).Interface().(client.Object)

	return IdFromObject(obj)
}
func IdFromObject(obj client.Object) string {
	return fmt.Sprintf("%s.%s.%s", obj.GetName(), obj.GetNamespace(), ezkube.GetClusterName(obj))
}
