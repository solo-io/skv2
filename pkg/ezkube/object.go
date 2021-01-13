package ezkube

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// the EzKube Object is a wrapper for a kubernetes runtime.Object
// which contains Kubernetes metadata
type Object interface {
	client.Object
	v1.Object
}

// the EzKube Object is a wrapper for a kubernetes List object
type List interface {
	client.Object
	v1.ListInterface
}
