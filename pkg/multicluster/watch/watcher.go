package watch

import (
	"context"
	"sort"
	"sync"

	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	"github.com/solo-io/skv2/pkg/multicluster/internal/k8s/core/v1/controller"
	"github.com/solo-io/skv2/pkg/multicluster/kubeconfig"
	"github.com/solo-io/skv2/pkg/reconcile"
	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

type clusterWatcher struct {
	ctx      context.Context
	handlers *handlerList
	managers *managerSet
	options  manager.Options
}

var _ multicluster.ClusterWatcher = &clusterWatcher{}
var _ multicluster.ManagerSet = &clusterWatcher{}
var _ multicluster.ClusterSet = &clusterWatcher{}

// NewClusterWatcher returns a *clusterWatcher.
// When ctx is cancelled, all cluster managers started by the clusterWatcher are stopped.
// Provided manager.Options are applied to all managers started by the clusterWatcher.
func NewClusterWatcher(ctx context.Context, options manager.Options) *clusterWatcher {
	return &clusterWatcher{
		ctx:      ctx,
		handlers: newHandlerList(),
		managers: newManagerSet(),
		options:  options,
	}
}

func (c *clusterWatcher) Run(master manager.Manager) error {
	loop := controller.NewSecretReconcileLoop("cluster watcher", master, reconcile.Options{})
	return loop.RunSecretReconciler(c.ctx, c, kubeconfig.Predicate)
}

func (c *clusterWatcher) ReconcileSecret(obj *v1.Secret) (reconcile.Result, error) {
	clusterName, clientCfg, err := kubeconfig.SecretToConfig(obj)
	if err != nil {
		return reconcile.Result{}, eris.Wrap(err, "failed to extract kubeconfig from secret")
	}

	restCfg, err := clientCfg.ClientConfig()
	if err != nil {
		return reconcile.Result{}, eris.Wrap(err, "failed to create rest config from kubeconfig")
	}

	// If the cluster already has a manager, remove the existing instance and start again.
	if _, err := c.managers.get(clusterName); err == nil {
		c.removeCluster(clusterName)
	}

	mgr, err := manager.New(restCfg, c.optionsWithDefaults())
	if err != nil {
		return reconcile.Result{}, err
	}

	c.startManager(clusterName, mgr)

	return reconcile.Result{}, nil
}

func (c *clusterWatcher) ReconcileSecretDeletion(req reconcile.Request) {
	c.removeCluster(req.Name)
}

func (c *clusterWatcher) RegisterClusterHandler(handler multicluster.ClusterHandler) {
	// Call the handler on all previously discovered clusters.
	c.managers.applyHandler(handler)
	// Add the handler to the list of active handlers to be called on clusters discovered later.
	c.handlers.add(handler)
}

func (c *clusterWatcher) Cluster(cluster string) (manager.Manager, error) {
	return c.managers.get(cluster)
}

func (s *clusterWatcher) ListClusters() []string {
	return s.managers.list()
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

	c.managers.set(clusterName, mgr, ctx, cancel)
	c.handlers.AddCluster(ctx, clusterName, mgr)
}

func (c *clusterWatcher) removeCluster(clusterName string) {
	c.managers.delete(clusterName)
	c.handlers.RemoveCluster(clusterName)
}

func (c *clusterWatcher) optionsWithDefaults() manager.Options {
	options := c.options
	options.HealthProbeBindAddress = "0"
	options.MetricsBindAddress = "0"
	return options
}

type asyncManager struct {
	manager manager.Manager
	ctx     context.Context
	cancel  context.CancelFunc
}

// managerSet maintains a set of managers.
type managerSet struct {
	mutex         sync.RWMutex
	asyncManagers map[string]asyncManager
}

func newManagerSet() *managerSet {
	return &managerSet{
		mutex:         sync.RWMutex{},
		asyncManagers: make(map[string]asyncManager),
	}
}

func (s *managerSet) get(cluster string) (manager.Manager, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	am, ok := s.asyncManagers[cluster]
	if !ok {
		return nil, eris.Errorf("Failed to get manager for cluster %v", cluster)
	}
	return am.manager, nil
}

func (s *managerSet) set(cluster string, manager manager.Manager, ctx context.Context, cancel context.CancelFunc) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.asyncManagers[cluster] = asyncManager{
		manager: manager,
		ctx:     ctx,
		cancel:  cancel,
	}
}

func (s *managerSet) delete(cluster string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	am, ok := s.asyncManagers[cluster]
	if !ok {
		return
	}
	am.cancel()
	delete(s.asyncManagers, cluster)
}

func (s *managerSet) list() []string {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var output []string
	for cluster := range s.asyncManagers {
		output = append(output, cluster)
	}
	sort.Strings(output)
	return output
}

func (s *managerSet) applyHandler(h multicluster.ClusterHandler) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for cluster, am := range s.asyncManagers {
		h.AddCluster(am.ctx, cluster, am.manager)
	}
}

type handlerList struct {
	mutex    sync.RWMutex
	handlers []multicluster.ClusterHandler
}

func newHandlerList() *handlerList {
	return &handlerList{
		mutex:    sync.RWMutex{},
		handlers: make([]multicluster.ClusterHandler, 0),
	}
}

func (h *handlerList) add(handler multicluster.ClusterHandler) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	h.handlers = append(h.handlers, handler)
}

func (h *handlerList) AddCluster(ctx context.Context, cluster string, mgr manager.Manager) {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	for _, handler := range h.handlers {
		handler.AddCluster(ctx, cluster, mgr)
	}
}

func (h *handlerList) RemoveCluster(cluster string) {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	for _, handler := range h.handlers {
		removeHandler, ok := handler.(multicluster.ClusterRemovedHandler)
		if ok {
			removeHandler.RemoveCluster(cluster)
		}
	}
}
