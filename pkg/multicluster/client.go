package multicluster

import "sigs.k8s.io/controller-runtime/pkg/client"

// Client exposes client.Client for multiple clusters.
type Client interface {
	// Cluster returns a client.Client for the given cluster.
	Cluster(name string) (client.Client, error)
}
