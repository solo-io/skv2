package ezkube

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CompareResourceId returns an integer comparing two ResourceIds lexicographically.
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
// a, b must be of type metav1.Object
func CreationTimestampsCompare(a, b interface{}) int {
	aa, aok := a.(metav1.Object)
	if !aok {
		panic("a is not a metav1.Object")
	}
	bb, bok := b.(metav1.Object)
	if !bok {
		panic("b is not a metav1.Object")
	}
	if creationTimestampsEqual(aa, bb) {
		return 0
	}
	if creationTimestampLessThan(aa, bb) {
		return -1
	}
	return 1
}

// CreationTimestampAscending returns true if a was created before b (ascending order)
// a & b must k8s metav1.Object interface
func CreationTimestampAscending(a, b interface{}) bool {
	aa, aok := a.(metav1.Object)
	if !aok {
		panic("a is not a metav1.Object")
	}
	bb, bok := b.(metav1.Object)
	if !bok {
		panic("b is not a metav1.Object")
	}
	return creationTimestampLessThan(aa, bb)
}

func creationTimestampLessThan(a, b metav1.Object) bool {
	return a.GetCreationTimestamp().Time.Before(b.GetCreationTimestamp().Time)
}

// CreationTimestampsEqual returns true if the creation timestamps of a and b are equal
// a & b must k8s metav1.Object interface
func creationTimestampsEqual(a, b metav1.Object) bool {
	return a.GetCreationTimestamp().Time.Equal(b.GetCreationTimestamp().Time)
}
