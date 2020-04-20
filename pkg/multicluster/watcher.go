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
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

const (
	LocalCluster = ""
)

// ClientGetter provides access to a client.Client for any registered cluster
type ClientGetter interface {
	// Cluster returns a client.Client for the given cluster if one is available, else errors.
	Cluster(name string) (client.Client, error)
}

// AddClusterHandler is called when a new cluster is added.
// The provided context is cancelled when a cluster is removed, so any teardown behavior for removed clusters
// should take place when ctx is cancelled.
type AddClusterHandler func(ctx context.Context, cluster string, mgr manager.Manager) error

// ClusterWatcher calls ClusterHandlers and maintains a set of active cluster credentials.
type ClusterWatcher interface {
	controller.SecretReconciler
	ClientGetter
}

type clusterWatcher struct {
	ctx      context.Context
	handlers []AddClusterHandler
	managers managerSet
}

// NewClusterWatcher returns an implementation of ClusterWatcher with localManager already registered.
func NewClusterWatcher(ctx context.Context, localManager manager.Manager, handlers ...AddClusterHandler) ClusterWatcher {
	watcher := &clusterWatcher{
		ctx:      ctx,
		handlers: handlers,
		managers: managerSet{
			mutex:    sync.RWMutex{},
			managers: make(map[string]managerWithCancel),
		},
	}

	err := watcher.registerManager(LocalCluster, localManager)
	if err != nil {
		contextutils.LoggerFrom(ctx).Panicw("Failed to register local kube config with multicluster watcher", zap.Error(err))
	}

	return watcher
}

func (c *clusterWatcher) ReconcileSecret(obj *v1.Secret) (reconcile.Result, error) {
	clusterName, cfg, err := kubeconfig.SecretToConfig(obj)
	if err != nil {
		return reconcile.Result{}, eris.Wrap(err, "failed to extract kubeconfig from secret")
	}

	if _, err := c.managers.get(clusterName); err != nil {
		return reconcile.Result{}, err
	}

	mgr, err := manager.New(cfg.RestConfig, manager.Options{
		// TODO these should be configurable, disable for now
		MetricsBindAddress:     "0",
		HealthProbeBindAddress: "0",
	})
	if err != nil {
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, c.registerManager(clusterName, mgr)
}

func (c *clusterWatcher) registerManager(clusterName string, mgr manager.Manager) error {
	ctx, cancel := context.WithCancel(
		contextutils.WithLoggerValues(context.WithValue(c.ctx, "cluster", clusterName), zap.String("cluster", clusterName)))
	go func() {
		err := mgr.Start(ctx.Done())
		if err != nil {
			cancel()
			contextutils.LoggerFrom(ctx).DPanicw("manager start failed for cluster %v", clusterName)
		}
	}()

	c.managers.set(clusterName, mgr, cancel)

	errs := &multierror.Error{}
	for _, handler := range c.handlers {
		err := handler(ctx, clusterName, mgr)
		if err != nil {
			errs.Errors = append(errs.Errors, err)
		}
	}

	return errs.ErrorOrNil()
}

func (c *clusterWatcher) ReconcileSecretDeletion(req reconcile.Request) {
	// TODO we have to enforce that the cluster name is the resource name
	// we can't lookup the deleted resource to find the name on the spec, because the resource is deleted.
	clusterName := req.Name
	_, err := c.managers.get(clusterName)
	if err != nil {
		contextutils.LoggerFrom(c.ctx).Debugw("reconciled delete on cluster secret for nonexistent cluster %v", clusterName)
		return
	}

	c.managers.delete(clusterName)
}

func (c *clusterWatcher) Cluster(name string) (client.Client, error) {
	mgr, err := c.managers.get(name)
	if err != nil {
		return nil, err
	}

	return mgr.GetClient(), nil
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

func (m managerSet) get(cluster string) (manager.Manager, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	mgrCancel, ok := m.managers[cluster]
	if !ok {
		return nil, eris.Errorf("Failed to get manager for cluster %v", cluster)
	}
	return mgrCancel.mgr, nil
}

func (m managerSet) set(cluster string, mgr manager.Manager, cancel context.CancelFunc) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.managers[cluster] = managerWithCancel{
		cancel: cancel,
		mgr:    mgr,
	}
}

func (m managerSet) delete(cluster string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	mgrCancel, ok := m.managers[cluster]
	if !ok {
		return
	}
	mgrCancel.cancel()
	delete(m.managers, cluster)
}
