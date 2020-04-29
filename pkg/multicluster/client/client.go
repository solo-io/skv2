package client

import (
	"github.com/rotisserie/eris"
	"github.com/solo-io/skv2/pkg/multicluster"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type mcClient struct {
	managers multicluster.ManagerSet
}

var _ multicluster.Client = &mcClient{}

func NewClient(managers multicluster.ManagerSet) *mcClient {
	return &mcClient{managers: managers}
}

func (c *mcClient) Cluster(name string) (client.Client, error) {
	mgr, err := c.managers.Cluster(name)
	if err != nil {
		return nil, eris.Wrapf(err, "failed to get client for cluster %v", name)
	}
	return mgr.GetClient(), nil
}
