package multicluster

import "sigs.k8s.io/controller-runtime/pkg/client"

// Client provides access to a client.Client for any registered cluster.
// Only one Client should be created within a multicluster system.
type Client interface {
	// Cluster returns a client.Client for the given cluster if one is available, else errors.
	Cluster(name string) (client.Client, error)
}

type mcClient struct {
	managers ManagerSet
}

var _ Client = &mcClient{}

func NewClient(managers ManagerSet) *mcClient {
	return &mcClient{
		managers: managers,
	}
}

func (s *mcClient) Cluster(name string) (client.Client, error) {
	mgr, err := s.managers.getManager(name)
	if err != nil {
		return nil, err
	}

	return mgr.GetClient(), nil
}
