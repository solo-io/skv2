package multicluster

import (
	"context"
	"sync"

	"github.com/rotisserie/eris"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// ClientSet provides access to a client.Client for any registered cluster.
// Only one ClientSet should be created within a multicluster system.
type ClientSet interface {
	// Cluster returns a client.Client for the given cluster if one is available, else errors.
	Cluster(name string) (client.Client, error)
	getManager(cluster string) (manager.Manager, error)
	setManager(cluster string, mgr manager.Manager, cancel context.CancelFunc)
	deleteManager(cluster string)
}

// managerWithCancel contains a manager and a cancel function to stop it.
type managerWithCancel struct {
	cancel context.CancelFunc
	mgr    manager.Manager
}

// setManager maintains a setManager of managers and cancel functions.
type set struct {
	mutex    sync.RWMutex
	managers map[string]managerWithCancel
}

func NewClientSet() ClientSet {
	return set{
		mutex:    sync.RWMutex{},
		managers: make(map[string]managerWithCancel),
	}
}

func (s set) Cluster(name string) (client.Client, error) {
	mgr, err := s.getManager(name)
	if err != nil {
		return nil, err
	}

	return mgr.GetClient(), nil
}

func (s set) getManager(cluster string) (manager.Manager, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	mgrCancel, ok := s.managers[cluster]
	if !ok {
		return nil, eris.Errorf("Failed to getManager manager for cluster %v", cluster)
	}
	return mgrCancel.mgr, nil
}

func (s set) setManager(cluster string, mgr manager.Manager, cancel context.CancelFunc) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.managers[cluster] = managerWithCancel{
		cancel: cancel,
		mgr:    mgr,
	}
}

func (s set) deleteManager(cluster string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	mgrCancel, ok := s.managers[cluster]
	if !ok {
		return
	}
	mgrCancel.cancel()
	delete(s.managers, cluster)
}
