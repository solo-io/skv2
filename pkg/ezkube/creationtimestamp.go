package ezkube

import "sigs.k8s.io/controller-runtime/pkg/client"

// CreationTimestampCompare returns an integer comparing two resources lexicographically.
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
func CreationTimestampCompare(a, b client.Object) int {
	if a.GetCreationTimestamp().Time.Equal(b.GetCreationTimestamp().Time) {
		return 0
	}
	if a.GetCreationTimestamp().Time.Before(b.GetCreationTimestamp().Time) {
		return -1
	}
	return 1
}
