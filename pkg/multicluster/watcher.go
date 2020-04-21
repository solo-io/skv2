package multicluster

import (
	"context"
	"sync"

	"github.com/hashicorp/go-multierror"
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

type clusterWatcher struct {
	ctx      context.Context
	handlers []ClusterHandler
	cancels  *cancelSet
}

// RunClusterWatcher initializes and runs a reconciler for KubeConfig secrets.
// It starts and runs ClusterHandlers for the localManager as if it were discovered by the watcher.
func RunClusterWatcher(ctx context.Context, localManager manager.Manager, handlers ...ClusterHandler) error {
	watcher := &clusterWatcher{
		ctx:      ctx,
		handlers: handlers,
		cancels:  newCancelSet(),
	}

	err := watcher.startManager(LocalCluster, localManager)
	if err != nil {
		return eris.Wrap(err, "failed to register local kube config with multicluster watcher")
	}

	loop := controller.NewSecretReconcileLoop("cluster watcher", localManager)
	return loop.RunSecretReconciler(ctx, watcher, kubeconfig.SecretPredicate{})
}

func (c *clusterWatcher) ReconcileSecret(obj *v1.Secret) (reconcile.Result, error) {
	clusterName, cfg, err := kubeconfig.SecretToConfig(obj)
	if err != nil {
		return reconcile.Result{}, eris.Wrap(err, "failed to extract kubeconfig from secret")
	}

	if _, err := c.cancels.get(clusterName); err == nil {
		return reconcile.Result{}, eris.Errorf("cluster %v already initialized", clusterName)
	}

	mgr, err := manager.New(cfg.RestConfig, manager.Options{
		// TODO these should be configurable, disable for now
		MetricsBindAddress:     "0",
		HealthProbeBindAddress: "0",
	})
	if err != nil {
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, c.startManager(clusterName, mgr)
}

func (c *clusterWatcher) startManager(clusterName string, mgr manager.Manager) error {
	ctx, cancel := context.WithCancel(
		contextutils.WithLoggerValues(context.WithValue(c.ctx, "cluster", clusterName), zap.String("cluster", clusterName)))
	go func() {
		err := mgr.Start(ctx.Done())
		if err != nil {
			cancel()
			contextutils.LoggerFrom(ctx).DPanicw("manager start failed for cluster %v", clusterName)
		}
	}()

	c.cancels.set(clusterName, cancel)

	errs := &multierror.Error{}
	for _, handler := range c.handlers {
		err := handler.HandleAddCluster(ctx, clusterName, mgr)
		if err != nil {
			errs.Errors = append(errs.Errors, err)
		}
	}

	return errs.ErrorOrNil()
}

func (c *clusterWatcher) ReconcileSecretDeletion(req reconcile.Request) {
	clusterName := req.Name
	cancel, err := c.cancels.get(clusterName)
	if err != nil {
		contextutils.LoggerFrom(c.ctx).Debugw("reconciled delete on cluster secret for uninitialized cluster %v", clusterName)
		return
	}

	cancel()
	c.cancels.delete(clusterName)
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
