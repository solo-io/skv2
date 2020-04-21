package multicluster

import (
	"context"
	"sync"

	"github.com/rotisserie/eris"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// ManagerSet maintains a manager.Manager for every cluster registered in the system.
// It is implemented as an interface with private methods to enforce that the provided implementation is used.
type ManagerSet interface {
	getManager(cluster string) (manager.Manager, error)
	setManager(cluster string, mgr manager.Manager, cancel context.CancelFunc)
	deleteManager(cluster string)
}

// managerWithCancel contains a manager and a cancel function to stop it.
type managerWithCancel struct {
	cancel context.CancelFunc
	mgr    manager.Manager
}

// managerSet maintains a set of managers and cancel functions.
type managerSet struct {
	mutex    sync.RWMutex
	managers map[string]managerWithCancel
}

var _ ManagerSet = &managerSet{}

// NewManagerSet returns a ManagerSet. One ManagerSet should be used for every active clusterWatcher.
func NewManagerSet() *managerSet {
	return &managerSet{
		mutex:    sync.RWMutex{},
		managers: make(map[string]managerWithCancel),
	}
}

func (s *managerSet) getManager(cluster string) (manager.Manager, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	mgrCancel, ok := s.managers[cluster]
	if !ok {
		return nil, eris.Errorf("Failed to get manager for cluster %v", cluster)
	}
	return mgrCancel.mgr, nil
}

func (s *managerSet) setManager(cluster string, mgr manager.Manager, cancel context.CancelFunc) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.managers[cluster] = managerWithCancel{
		cancel: cancel,
		mgr:    mgr,
	}
}

func (s *managerSet) deleteManager(cluster string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	mgrCancel, ok := s.managers[cluster]
	if !ok {
		return
	}
	mgrCancel.cancel()
	delete(s.managers, cluster)
}
