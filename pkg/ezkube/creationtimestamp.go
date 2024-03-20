package ezkube

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreationTimestampsEqual(a, b ResourceId) bool {
	return a.(metav1.Object).GetCreationTimestamp().Time.Equal(b.(metav1.Object).GetCreationTimestamp().Time)
}

// CreationTimestampAscending returns true if a was created before b (ascending order)
func CreationTimestampAscending(a, b ResourceId) bool {
	return a.(metav1.Object).GetCreationTimestamp().Time.Before(b.(metav1.Object).GetCreationTimestamp().Time)
}
