package resource

import "sigs.k8s.io/controller-runtime/pkg/client"

//go:generate mockgen -source ./client.go -destination ./mocks/client.go

// right now this just allows us to generate mocks for the client.
// (ilackarms): eventually it might make sense to put clients for non-kube backends in this package
type Client interface {
	client.Client
}
