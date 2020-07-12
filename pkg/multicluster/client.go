package multicluster

import (
	"github.com/rotisserie/eris"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Client exposes client.Client for multiple clusters.
type Client interface {
	// List available clusters
	ClusterSet

	// Cluster returns a client.Client for the given cluster.
	Cluster(name string) (client.Client, error)

}

type mcClient struct {
	managers ManagerSet
}

var _ Client = &mcClient{}

func NewClient(managers ManagerSet) *mcClient {
	return &mcClient{managers: managers}
}

func (c *mcClient) Cluster(name string) (client.Client, error) {
	mgr, err := c.managers.Cluster(name)
	if err != nil {
		return nil, eris.Wrapf(err, "failed to get client for cluster %v", name)
	}
	return mgr.GetClient(), nil
}

func (c *mcClient) 	ListClusters() []string {
	return c.managers.ListClusters()
}
