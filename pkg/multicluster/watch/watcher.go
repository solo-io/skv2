package watch

import (
	"context"
	"sort"
	"sync"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	"github.com/solo-io/skv2/pkg/multicluster/internal/k8s/core/v1/controller"
	"github.com/solo-io/skv2/pkg/multicluster/kubeconfig"
	"github.com/solo-io/skv2/pkg/reconcile"
	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// RetryOptions specify how to retry when a manager fails to be created or started.
type RetryOptions struct {
	// Delay is the initial retry delay.
	// Default is 1 second.
	Delay time.Duration

	// DelayType is the type of retry delay (fixed or exponential backoff).
	// Default is exponential backoff.
	DelayType RetryDelayType

	// MaxDelay is the maximum delay between retries.
	// Default is 0 (no max).
	MaxDelay time.Duration

	// MaxRetries is the maximum number of retries.
	// Default is 0 (retry forever).
	MaxRetries uint
}

// RetryDelayType is the type of delay to be used for manager creation retries.
type RetryDelayType int

const (
	// RetryDelayType_Backoff means retry with exponential backoff (with random jitter).
	RetryDelayType_Backoff RetryDelayType = iota

	// RetryDelayType_Fixed means retry at a fixed interval (with random jitter).
	RetryDelayType_Fixed
)

type clusterWatcher struct {
	ctx             context.Context
	handlers        *handlerList
	managers        *managerSet
	managerOptions  manager.Options
	watchNamespaces []string
	retryOptions    RetryOptions
}

var _ multicluster.Interface = &clusterWatcher{}

// NewClusterWatcher returns a *clusterWatcher, which watches for changes to kubeconfig secrets
// (which contain kubeconfigs for remote clusters).
// - When ctx is cancelled, all cluster managers started by the clusterWatcher are stopped.
// - Provided manager.Options are applied to all managers started by the clusterWatcher.
// - RetryOptions specify how to retry manager creation if it fails. Any fields not explicitly provided
//   in the retry options will take on the default values.
// - If watchNamespaces is not empty, only secrets in the given namespaces will be watched. If empty, secrets in
// all namespaces will be watched.
func NewClusterWatcher(ctx context.Context,
	managerOptions manager.Options,
	retryOptions RetryOptions,
	watchNamespaces []string) *clusterWatcher {
	return &clusterWatcher{
		ctx:             ctx,
		handlers:        newHandlerList(),
		managers:        newManagerSet(),
		managerOptions:  managerOptions,
		retryOptions:    retryOptions,
		watchNamespaces: watchNamespaces,
	}
}

func (c *clusterWatcher) Run(master manager.Manager) error {
	loop := controller.NewSecretReconcileLoop("cluster watcher", master, reconcile.Options{})
	return loop.RunSecretReconciler(c.ctx, c, kubeconfig.BuildPredicate(c.watchNamespaces))
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

	c.startManager(clusterName, restCfg)

	return reconcile.Result{}, nil
}

func (c *clusterWatcher) ReconcileSecretDeletion(req reconcile.Request) error {
	c.removeCluster(req.Name)
	return nil
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

func (c *clusterWatcher) startManager(clusterName string, restCfg *rest.Config) {
	retryOptions := c.retryOptionsWithDefaults()
	go func() { // this must be async because mgr.Start(ctx) is blocking
		retry.Do(func() error {
			mgr, err := manager.New(restCfg, c.managerOptionsWithDefaults())
			if err != nil {
				contextutils.LoggerFrom(c.ctx).Errorf("Manager creation failed for cluster %v: %v", clusterName, err)
				return err
			}

			ctx, cancel := context.WithCancel(contextutils.WithLoggerValues(c.ctx, zap.String("cluster", clusterName)))

			// add manager to managers+handlers.  It may fail to start and need to be removed ¯\_(ツ)_/¯
			c.managers.set(clusterName, mgr, ctx, cancel)
			c.handlers.AddCluster(ctx, clusterName, mgr)

			err = mgr.Start(ctx) // blocking until error is thrown
			if err != nil {
				contextutils.LoggerFrom(ctx).Errorf("Manager start failed for cluster %v: %v", clusterName, err)

				// remove failed manager from managers+handlers
				c.managers.delete(clusterName)
				c.handlers.RemoveCluster(clusterName)
			}

			// continue the exponentially-backing-off retry
			return err
		}, retryOptions...)
	}()
}

func (c *clusterWatcher) removeCluster(clusterName string) {
	c.managers.delete(clusterName)
	c.handlers.RemoveCluster(clusterName)
}

func (c *clusterWatcher) managerOptionsWithDefaults() manager.Options {
	managerOptions := c.managerOptions
	managerOptions.HealthProbeBindAddress = "0"
	managerOptions.MetricsBindAddress = "0"
	return managerOptions
}

func (c *clusterWatcher) retryOptionsWithDefaults() []retry.Option {
	// set delay to 1 second if not set
	delay := c.retryOptions.Delay
	if delay == 0 {
		delay = time.Second
	}

	// convert the delay type
	var delayType retry.DelayTypeFunc
	if c.retryOptions.DelayType == RetryDelayType_Fixed {
		delayType = retry.FixedDelay
	} else {
		delayType = retry.BackOffDelay
	}

	// the rest will default to their zero values if not set
	maxDelay := c.retryOptions.MaxDelay
	attempts := c.retryOptions.MaxRetries

	return []retry.Option{
		retry.DelayType(retry.CombineDelay(delayType, retry.RandomDelay)),
		retry.MaxJitter(delay),
		retry.Delay(delay),
		retry.MaxDelay(maxDelay),
		retry.Attempts(attempts),
	}
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
