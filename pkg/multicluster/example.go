package multicluster

import (
	"context"

	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/contextutils"
	skv2_corev1 "github.com/solo-io/skv2/pkg/api/kube/core/v1"
	"github.com/solo-io/skv2/pkg/api/kube/core/v1/controller"
	"go.uber.org/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

/*
Example of how the watcher could fit into an app setup flow.
*/
func example(local manager.Manager) {
	loop := controller.NewSecretReconcileLoop("cluster controller", local)
	clusterController := NewClusterWatcher(
		context.TODO(),
		local,
		multiclusterConfigmapReconcileLoop{}.HandleAddCluster,
	)

	err := loop.RunSecretReconciler(context.TODO(), clusterController)
	if err != nil {
		// oh no
	}

	var getter ClientGetter = clusterController
	multiclusterClients := NewMCClientSet(getter)
	fooSet, err := multiclusterClients.Cluster("foo")
	if err != nil {
		// uh oh
	}

	fooSet.Secrets().DeleteAllOfSecret(context.TODO())

}

/**
Rough sketch of a typed multicluster reconcile loop
*/
// TODO generate
type multiclusterConfigmapReconcileLoop struct {
	rec controller.ConfigMapReconciler
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

/**
Rough sketch of a typed multicluster clientset

Alternative is to have a structure like "set.Resource().Cluster().Action()", wdyt?
*/
// TODO generate

type multiclusterClientSet interface {
	Cluster(cluster string) (skv2_corev1.Clientset, error)
}

type mccs struct{ getter ClientGetter }

func (m mccs) Cluster(cluster string) (skv2_corev1.Clientset, error) {
	c, err := m.getter.Cluster(cluster)
	if err != nil {
		return nil, eris.Wrapf(err, "Failed to get client for cluster %v")
	}
	return skv2_corev1.NewClientset(c), nil
}

func NewMCClientSet(getter ClientGetter) multiclusterClientSet {
	return mccs{getter: getter}
}
