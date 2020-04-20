package multicluster

import (
	"context"

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

// ClusterHandler is passed to RunClusterWatcher to handle select cluster events.
type ClusterHandler interface {
	// ClusterHandler is called when a new cluster is identified by a cluster watch.
	// The provided context is cancelled when a cluster is removed, so any teardown behavior for removed clusters
	// should take place when ctx is cancelled.
	HandleAddCluster(ctx context.Context, cluster string, mgr manager.Manager) error
}

type clusterWatcher struct {
	ctx      context.Context
	handlers []ClusterHandler
	clients  Client
}

// RunClusterWatcher initializes and runs a reconciler for KubeConfig secrets.
func RunClusterWatcher(ctx context.Context, localManager manager.Manager, clients Client, handlers ...ClusterHandler) error {
	watcher := &clusterWatcher{
		ctx:      ctx,
		handlers: handlers,
		clients:  clients,
	}

	err := watcher.registerManager(LocalCluster, localManager)
	if err != nil {
		return eris.Wrap(err, "Failed to register local kube config with multicluster watcher")
	}

	loop := controller.NewSecretReconcileLoop("cluster watcher", localManager)
	return loop.RunSecretReconciler(ctx, watcher, kubeconfig.SecretPredicate{})
}

func (c *clusterWatcher) ReconcileSecret(obj *v1.Secret) (reconcile.Result, error) {
	clusterName, cfg, err := kubeconfig.SecretToConfig(obj)
	if err != nil {
		return reconcile.Result{}, eris.Wrap(err, "failed to extract kubeconfig from secret")
	}

	if _, err := c.clients.getManager(clusterName); err != nil {
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

	c.clients.setManager(clusterName, mgr, cancel)

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
	// TODO we have to enforce that the cluster name is the resource name
	// we can't lookup the deleted resource to find the name on the spec, because the resource is deleted.
	clusterName := req.Name
	_, err := c.clients.getManager(clusterName)
	if err != nil {
		contextutils.LoggerFrom(c.ctx).Debugw("reconciled deleteManager on cluster secret for nonexistent cluster %v", clusterName)
		return
	}

	c.clients.deleteManager(clusterName)
}
