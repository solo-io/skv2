package multicluster

import (
	"context"
	"sync"

	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/skv2/pkg/api/kube/core/v1/controller"
	"github.com/solo-io/skv2/pkg/multicluster/kubeconfig"
	"github.com/solo-io/skv2/pkg/reconcile"
	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

const (
	LocalCluster = ""
)

type ClusterWatcher interface {
	// Run starts a watch for KubeConfig secrets on the cluster managed by the local manager.Manager.
	// Note that Run will call Start on the local manager and run all registered ClusterHandlers.
	Run(local manager.Manager) error
	// RegisterClusterHandler adds a ClusterHandler to the ClusterWatcher.
	RegisterClusterHandler(handler ClusterHandler)
}

type clusterWatcher struct {
	ctx      context.Context
	mutex    sync.RWMutex
	handlers *handlerList
	cancels  *cancelSet
}

var _ ClusterWatcher = &clusterWatcher{}

func NewClusterWatcher(ctx context.Context) *clusterWatcher {
	return &clusterWatcher{
		ctx:      ctx,
		mutex:    sync.RWMutex{},
		handlers: newHandlerList(),
		cancels:  newCancelSet(),
	}
}

func (c *clusterWatcher) Run(local manager.Manager) error {
	c.startManager(LocalCluster, local)
	loop := controller.NewSecretReconcileLoop("cluster watcher", local)
	return loop.RunSecretReconciler(c.ctx, c, kubeconfig.SecretPredicate{})
}

func (c *clusterWatcher) ReconcileSecret(obj *v1.Secret) (reconcile.Result, error) {
	clusterName, cfg, err := kubeconfig.SecretToConfig(obj)
	if err != nil {
		return reconcile.Result{}, eris.Wrap(err, "failed to extract kubeconfig from secret")
	}

	// If the cluster already has a manager, remove the existing instance and start again.
	if _, err := c.cancels.get(clusterName); err == nil {
		c.removeCluster(clusterName)
	}

	mgr, err := manager.New(cfg.RestConfig, manager.Options{
		// TODO these should be configurable, disable for now
		MetricsBindAddress:     "0",
		HealthProbeBindAddress: "0",
	})
	if err != nil {
		return reconcile.Result{}, err
	}

	c.startManager(clusterName, mgr)

	return reconcile.Result{}, nil
}

func (c *clusterWatcher) startManager(clusterName string, mgr manager.Manager) {
	ctx, cancel := context.WithCancel(
		contextutils.WithLoggerValues(context.WithValue(c.ctx, "cluster", clusterName), zap.String("cluster", clusterName)))
	go func() {
		err := mgr.Start(ctx.Done())
		if err != nil {
			contextutils.LoggerFrom(ctx).DPanicw("manager start failed for cluster %v", clusterName)
		}
	}()

	c.cancels.set(clusterName, cancel)
	c.handlers.HandleAddCluster(ctx, clusterName, mgr)
}

func (c *clusterWatcher) ReconcileSecretDeletion(req reconcile.Request) {
	// TODO update to namespace.name
	c.removeCluster(req.Name)
}

func (c *clusterWatcher) removeCluster(clusterName string) {
	cancel, err := c.cancels.get(clusterName)
	if err != nil {
		contextutils.LoggerFrom(c.ctx).Debugw("reconciled delete on cluster secret for uninitialized cluster %v", clusterName)
		return
	}

	cancel()
	c.cancels.delete(clusterName)
}

func (c *clusterWatcher) RegisterClusterHandler(handler ClusterHandler) {
	c.handlers.add(handler)
}

// cancelSet maintains a set of cancel functions.
type cancelSet struct {
	mutex   sync.RWMutex
	cancels map[string]context.CancelFunc
}

func newCancelSet() *cancelSet {
	return &cancelSet{
		mutex:   sync.RWMutex{},
		cancels: make(map[string]context.CancelFunc),
	}
}

func (s *cancelSet) get(cluster string) (context.CancelFunc, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	cancel, ok := s.cancels[cluster]
	if !ok {
		return nil, eris.Errorf("Failed to get cancel function for cluster %v", cluster)
	}
	return cancel, nil
}

func (s *cancelSet) set(cluster string, cancel context.CancelFunc) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.cancels[cluster] = cancel
}

func (s *cancelSet) delete(cluster string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.cancels, cluster)
}

type handlerList struct {
	mutex    sync.RWMutex
	handlers []ClusterHandler
}

func newHandlerList() *handlerList {
	return &handlerList{
		mutex:    sync.RWMutex{},
		handlers: make([]ClusterHandler, 0),
	}
}

func (h *handlerList) add(handler ClusterHandler) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	h.handlers = append(h.handlers, handler)
}

func (h *handlerList) HandleAddCluster(ctx context.Context, cluster string, mgr manager.Manager) {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	for _, handler := range h.handlers {
		handler.HandleAddCluster(ctx, cluster, mgr)
	}
}
