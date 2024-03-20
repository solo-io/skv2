package ezkube

import (
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// CreationTimestampsEqual returns true if the creation timestamps of a and b are equal
func CreationTimestampsEqual(a, b client.Object) bool {
	return a.GetCreationTimestamp().Time.Equal(b.GetCreationTimestamp().Time)
}

// CreationTimestampAscending returns true if a was created before b (ascending order)
func CreationTimestampAscending(a, b client.Object) bool {
	return a.GetCreationTimestamp().Time.Before(b.GetCreationTimestamp().Time)
}
