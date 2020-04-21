package multicluster

import (
	"context"
	"sync"

	"github.com/rotisserie/eris"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// Client provides access to a client.Client for any registered cluster.
// Only one Client should be created within a multicluster system.
type Client interface {
	// Cluster returns a client.Client for the given cluster if one is available, else errors.
	Cluster(name string) (client.Client, error)
}

var _ Client = &mcClient{}
var _ ClusterHandler = &mcClient{}

type mcClient struct {
	managers *managerSet
}

func NewClient() *mcClient {
	return &mcClient{
		managers: newManagerSet(),
	}
}

func (s *mcClient) Cluster(name string) (client.Client, error) {
	mgr, err := s.managers.get(name)
	if err != nil {
		return nil, err
	}

	return mgr.GetClient(), nil
}

func (s *mcClient) HandleAddCluster(ctx context.Context, cluster string, mgr manager.Manager) error {
	s.managers.set(cluster, mgr)

	go func() {
		<-ctx.Done()
		s.managers.delete(cluster)
	}()

	return nil
}

// managerSet maintains a set of managers.
type managerSet struct {
	mutex    sync.RWMutex
	managers map[string]manager.Manager
}

func newManagerSet() *managerSet {
	return &managerSet{
		mutex:    sync.RWMutex{},
		managers: make(map[string]manager.Manager),
	}
}

func (s *managerSet) get(cluster string) (manager.Manager, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	mgr, ok := s.managers[cluster]
	if !ok {
		return nil, eris.Errorf("Failed to get manager for cluster %v", cluster)
	}
	return mgr, nil
}

func (s *managerSet) set(cluster string, mgr manager.Manager) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.managers[cluster] = mgr
}

func (s *managerSet) delete(cluster string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.managers, cluster)
}
