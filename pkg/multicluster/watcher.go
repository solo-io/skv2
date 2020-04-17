package multicluster

import (
	"context"
	"sync"

	"github.com/hashicorp/go-multierror"
	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/contextutils"
	skv2_corev1 "github.com/solo-io/skv2/pkg/api/kube/core/v1"
	"github.com/solo-io/skv2/pkg/api/kube/core/v1/controller"
	"github.com/solo-io/skv2/pkg/multicluster/kubeconfig"
	"github.com/solo-io/skv2/pkg/reconcile"
	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func example(local manager.Manager) {
	loop := controller.NewSecretReconcileLoop("cluster controller", local)
	clusterController := NewClusterWatcher(
		context.TODO(),
		multiclusterConfigmapReconcileLoop{},
	)

	// TODO predicate for kubeconfig secrets
	err := loop.RunSecretReconciler(context.TODO(), clusterController)
	if err != nil {
		// oh no
	}

	var getter ClientGetter = clusterController
	fooClient, err := getter.Cluster("foo")
	if err != nil {
		// oh no!
	}

	fooSecretClient := skv2_corev1.NewSecretClient(fooClient)
	err = fooSecretClient.DeleteAllOfSecret(context.TODO())
	if err != nil {
		// uh oh
	}
}

// TODO generate
type multiclusterConfigmapReconcileLoop struct {
	rec             controller.ConfigMapReconciler
	onRemoveCluster func(cluster string)
}

func (c multiclusterConfigmapReconcileLoop) HandleAddCluster(ctx context.Context, cluster string, mgr manager.Manager) error {
	go func() {
		err := controller.NewConfigMapReconcileLoop(cluster, mgr).RunConfigMapReconciler(ctx, c.rec)
		if err != nil {
			contextutils.LoggerFrom(ctx).DPanicw("ConfigMap reconcile loop stopped with error", zap.Error(err))
		}
	}()
	return nil
}

func (c multiclusterConfigmapReconcileLoop) HandleRemoveCluster(cluster string) {
	c.onRemoveCluster(cluster)
}

// ClientGetter provides access to a client.Client for any registered cluster
type ClientGetter interface {
	// Cluster returns a client.Client for the given cluster if one is available, else errors.
	Cluster(name string) (client.Client, error)
}

// ClusterHandler can be passed to the ClusterWatcher to allow components to respond to cluster events.
type ClusterHandler interface {
	// HandleAddCluster is called when a new cluster is added.
	// The provided context is cancelled when a cluster is removed.
	HandleAddCluster(ctx context.Context, cluster string, mgr manager.Manager) error
	// HandleRemoveCluster is called when a cluster is deleted.
	HandleRemoveCluster(cluster string)
}

type clusterWatcher struct {
	ctx      context.Context
	handlers []ClusterHandler
	managers managerSet
}

// ClusterWatcher calls ClusterHandlers and maintains a set of active cluster credentials.
type ClusterWatcher interface {
	controller.SecretReconciler
	ClientGetter
}

func NewClusterWatcher(ctx context.Context, handlers ...ClusterHandler) ClusterWatcher {
	return &clusterWatcher{
		ctx:      ctx,
		handlers: handlers,
		managers: managerSet{
			mutex:    sync.RWMutex{},
			managers: make(map[string]managerWithCancel),
		},
	}
}

func (c *clusterWatcher) ReconcileSecret(obj *v1.Secret) (reconcile.Result, error) {
	clusterName, cfg, err := kubeconfig.SecretToConfig(obj)
	if err != nil {
		return reconcile.Result{}, eris.Wrap(err, "failed to extract kubeconfig from secret")
	}

	if _, err := c.managers.getManager(clusterName); err != nil {
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

	ctx, cancel := context.WithCancel(contextutils.WithLoggerValues(context.WithValue(c.ctx, "cluster", clusterName)))
	go func() {
		err = mgr.Start(ctx.Done())
		if err != nil {
			cancel()
			contextutils.LoggerFrom(ctx).DPanicw("manager start failed for cluster %v", clusterName)
		}
	}()

	c.managers.setManager(clusterName, mgr, cancel)

	errs := &multierror.Error{}
	for _, handler := range c.handlers {
		err := handler.HandleAddCluster(ctx, clusterName, mgr)
		if err != nil {
			errs.Errors = append(errs.Errors, err)
		}
	}

	return reconcile.Result{}, errs.ErrorOrNil()
}

func (c *clusterWatcher) ReconcileSecretDeletion(req reconcile.Request) {
	// TODO we have to enforce that the cluster name is the resource name
	// we can't lookup the deleted resource to find the name on the spec, because the resource is deleted.
	clusterName := req.Name
	_, err := c.managers.getManager(clusterName)
	if err != nil {
		contextutils.LoggerFrom(c.ctx).Debugw("reconciled delete on cluster secret for nonexistent cluster %v", clusterName)
		return
	}

	for _, handler := range c.handlers {
		handler.HandleRemoveCluster(clusterName)
	}

	c.managers.deleteManager(clusterName)
}

func (c *clusterWatcher) Cluster(name string) (client.Client, error) {
	mgr, err := c.managers.getManager(name)
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

// TODO should this be extracted? we could add GetClient to this and pass it around instead of the controller
// managerSet maintains a set of managers and cancel functions
type managerSet struct {
	mutex    sync.RWMutex
	managers map[string]managerWithCancel
}

func (m managerSet) getManager(cluster string) (manager.Manager, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	mgrCancel, ok := m.managers[cluster]
	if !ok {
		return nil, eris.Errorf("Failed to get manager for cluster %v", cluster)
	}
	return mgrCancel.mgr, nil
}

func (m managerSet) setManager(cluster string, mgr manager.Manager, cancel context.CancelFunc) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.managers[cluster] = managerWithCancel{
		cancel: cancel,
		mgr:    mgr,
	}
}

func (m managerSet) deleteManager(cluster string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	mgrCancel, ok := m.managers[cluster]
	if !ok {
		return
	}
	mgrCancel.cancel()
	delete(m.managers, cluster)
}
