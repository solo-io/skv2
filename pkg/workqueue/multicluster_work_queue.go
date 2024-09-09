package workqueue

import (
	"sync"

	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// MultiClusterQueues multiplexes queues across
// multiple k8s clusters.
type MultiClusterQueues struct {
	queues map[string]workqueue.TypedRateLimitingInterface[reconcile.Request]
	lock   sync.RWMutex
}

// sets the queue for a cluster
func (s *MultiClusterQueues) Set(cluster string, queue workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.queues == nil {
		s.queues = make(map[string]workqueue.TypedRateLimitingInterface[reconcile.Request])
	}
	s.queues[cluster] = queue
}

// removes the queue for a cluster
func (s *MultiClusterQueues) Remove(cluster string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.queues, cluster)
}

// get the stored queues for the cluster
func (s *MultiClusterQueues) Get(cluster string) workqueue.TypedRateLimitingInterface[reconcile.Request] {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.queues[cluster]
}

// currently unused, useful for debugging
func (s *MultiClusterQueues) All() []workqueue.TypedRateLimitingInterface[reconcile.Request] {
	s.lock.RLock()
	defer s.lock.RUnlock()
	var queues []workqueue.TypedRateLimitingInterface[reconcile.Request]
	for _, queue := range s.queues {
		queues = append(queues, queue)
	}
	return queues
}
