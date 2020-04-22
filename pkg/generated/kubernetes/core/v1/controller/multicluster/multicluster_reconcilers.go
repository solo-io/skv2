// Definitions for the multicluster Kubernetes Controllers
package multicluster

import (
	"context"

	core_v1 "k8s.io/api/core/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the Secret Resource across clusters.
// implemented by the user
type MulticlusterSecretReconciler interface {
	ReconcileSecret(clusterName string, obj *core_v1.Secret) (reconcile.Result, error)
}

// Reconcile deletion events for the Secret Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterSecretDeletionReconciler interface {
	ReconcileSecretDeletion(clusterName string, req reconcile.Request)
}

type MulticlusterSecretReconcilerFuncs struct {
	OnReconcileSecret         func(clusterName string, obj *core_v1.Secret) (reconcile.Result, error)
	OnReconcileSecretDeletion func(clusterName string, req reconcile.Request)
}

func (f *MulticlusterSecretReconcilerFuncs) ReconcileSecret(clusterName string, obj *core_v1.Secret) (reconcile.Result, error) {
	if f.OnReconcileSecret == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileSecret(clusterName, obj)
}

func (f *MulticlusterSecretReconcilerFuncs) ReconcileSecretDeletion(clusterName string, req reconcile.Request) {
	if f.OnReconcileSecretDeletion == nil {
		return
	}
	f.OnReconcileSecretDeletion(clusterName, req)
}

type MulticlusterSecretReconcileLoop interface {
	// AddMulticlusterSecretReconciler adds a MulticlusterSecretReconciler to the MulticlusterSecretReconcileLoop.
	AddMulticlusterSecretReconciler(ctx context.Context, rec MulticlusterSecretReconciler, predicates ...predicate.Predicate)
}

type multiclusterSecretReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterSecretReconcileLoop) AddMulticlusterSecretReconciler(ctx context.Context, rec MulticlusterSecretReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericSecretMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterSecretReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterSecretReconcileLoop {
	return &multiclusterSecretReconcileLoop{loop: multicluster.NewLoop(name, cw, &core_v1.Secret{})}
}

type genericSecretMulticlusterReconciler struct {
	reconciler MulticlusterSecretReconciler
}

func (g genericSecretMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) {
	if deletionReconciler, ok := g.reconciler.(MulticlusterSecretDeletionReconciler); ok {
		deletionReconciler.ReconcileSecretDeletion(cluster, req)
	}
}

func (g genericSecretMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*core_v1.Secret)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: Secret handler received event for %T", object)
	}
	return g.reconciler.ReconcileSecret(cluster, obj)
}

// Reconcile Upsert events for the ServiceAccount Resource across clusters.
// implemented by the user
type MulticlusterServiceAccountReconciler interface {
	ReconcileServiceAccount(clusterName string, obj *core_v1.ServiceAccount) (reconcile.Result, error)
}

// Reconcile deletion events for the ServiceAccount Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterServiceAccountDeletionReconciler interface {
	ReconcileServiceAccountDeletion(clusterName string, req reconcile.Request)
}

type MulticlusterServiceAccountReconcilerFuncs struct {
	OnReconcileServiceAccount         func(clusterName string, obj *core_v1.ServiceAccount) (reconcile.Result, error)
	OnReconcileServiceAccountDeletion func(clusterName string, req reconcile.Request)
}

func (f *MulticlusterServiceAccountReconcilerFuncs) ReconcileServiceAccount(clusterName string, obj *core_v1.ServiceAccount) (reconcile.Result, error) {
	if f.OnReconcileServiceAccount == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileServiceAccount(clusterName, obj)
}

func (f *MulticlusterServiceAccountReconcilerFuncs) ReconcileServiceAccountDeletion(clusterName string, req reconcile.Request) {
	if f.OnReconcileServiceAccountDeletion == nil {
		return
	}
	f.OnReconcileServiceAccountDeletion(clusterName, req)
}

type MulticlusterServiceAccountReconcileLoop interface {
	// AddMulticlusterServiceAccountReconciler adds a MulticlusterServiceAccountReconciler to the MulticlusterServiceAccountReconcileLoop.
	AddMulticlusterServiceAccountReconciler(ctx context.Context, rec MulticlusterServiceAccountReconciler, predicates ...predicate.Predicate)
}

type multiclusterServiceAccountReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterServiceAccountReconcileLoop) AddMulticlusterServiceAccountReconciler(ctx context.Context, rec MulticlusterServiceAccountReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericServiceAccountMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterServiceAccountReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterServiceAccountReconcileLoop {
	return &multiclusterServiceAccountReconcileLoop{loop: multicluster.NewLoop(name, cw, &core_v1.ServiceAccount{})}
}

type genericServiceAccountMulticlusterReconciler struct {
	reconciler MulticlusterServiceAccountReconciler
}

func (g genericServiceAccountMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) {
	if deletionReconciler, ok := g.reconciler.(MulticlusterServiceAccountDeletionReconciler); ok {
		deletionReconciler.ReconcileServiceAccountDeletion(cluster, req)
	}
}

func (g genericServiceAccountMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*core_v1.ServiceAccount)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: ServiceAccount handler received event for %T", object)
	}
	return g.reconciler.ReconcileServiceAccount(cluster, obj)
}

// Reconcile Upsert events for the ConfigMap Resource across clusters.
// implemented by the user
type MulticlusterConfigMapReconciler interface {
	ReconcileConfigMap(clusterName string, obj *core_v1.ConfigMap) (reconcile.Result, error)
}

// Reconcile deletion events for the ConfigMap Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterConfigMapDeletionReconciler interface {
	ReconcileConfigMapDeletion(clusterName string, req reconcile.Request)
}

type MulticlusterConfigMapReconcilerFuncs struct {
	OnReconcileConfigMap         func(clusterName string, obj *core_v1.ConfigMap) (reconcile.Result, error)
	OnReconcileConfigMapDeletion func(clusterName string, req reconcile.Request)
}

func (f *MulticlusterConfigMapReconcilerFuncs) ReconcileConfigMap(clusterName string, obj *core_v1.ConfigMap) (reconcile.Result, error) {
	if f.OnReconcileConfigMap == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileConfigMap(clusterName, obj)
}

func (f *MulticlusterConfigMapReconcilerFuncs) ReconcileConfigMapDeletion(clusterName string, req reconcile.Request) {
	if f.OnReconcileConfigMapDeletion == nil {
		return
	}
	f.OnReconcileConfigMapDeletion(clusterName, req)
}

type MulticlusterConfigMapReconcileLoop interface {
	// AddMulticlusterConfigMapReconciler adds a MulticlusterConfigMapReconciler to the MulticlusterConfigMapReconcileLoop.
	AddMulticlusterConfigMapReconciler(ctx context.Context, rec MulticlusterConfigMapReconciler, predicates ...predicate.Predicate)
}

type multiclusterConfigMapReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterConfigMapReconcileLoop) AddMulticlusterConfigMapReconciler(ctx context.Context, rec MulticlusterConfigMapReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericConfigMapMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterConfigMapReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterConfigMapReconcileLoop {
	return &multiclusterConfigMapReconcileLoop{loop: multicluster.NewLoop(name, cw, &core_v1.ConfigMap{})}
}

type genericConfigMapMulticlusterReconciler struct {
	reconciler MulticlusterConfigMapReconciler
}

func (g genericConfigMapMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) {
	if deletionReconciler, ok := g.reconciler.(MulticlusterConfigMapDeletionReconciler); ok {
		deletionReconciler.ReconcileConfigMapDeletion(cluster, req)
	}
}

func (g genericConfigMapMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*core_v1.ConfigMap)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: ConfigMap handler received event for %T", object)
	}
	return g.reconciler.ReconcileConfigMap(cluster, obj)
}

// Reconcile Upsert events for the Service Resource across clusters.
// implemented by the user
type MulticlusterServiceReconciler interface {
	ReconcileService(clusterName string, obj *core_v1.Service) (reconcile.Result, error)
}

// Reconcile deletion events for the Service Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterServiceDeletionReconciler interface {
	ReconcileServiceDeletion(clusterName string, req reconcile.Request)
}

type MulticlusterServiceReconcilerFuncs struct {
	OnReconcileService         func(clusterName string, obj *core_v1.Service) (reconcile.Result, error)
	OnReconcileServiceDeletion func(clusterName string, req reconcile.Request)
}

func (f *MulticlusterServiceReconcilerFuncs) ReconcileService(clusterName string, obj *core_v1.Service) (reconcile.Result, error) {
	if f.OnReconcileService == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileService(clusterName, obj)
}

func (f *MulticlusterServiceReconcilerFuncs) ReconcileServiceDeletion(clusterName string, req reconcile.Request) {
	if f.OnReconcileServiceDeletion == nil {
		return
	}
	f.OnReconcileServiceDeletion(clusterName, req)
}

type MulticlusterServiceReconcileLoop interface {
	// AddMulticlusterServiceReconciler adds a MulticlusterServiceReconciler to the MulticlusterServiceReconcileLoop.
	AddMulticlusterServiceReconciler(ctx context.Context, rec MulticlusterServiceReconciler, predicates ...predicate.Predicate)
}

type multiclusterServiceReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterServiceReconcileLoop) AddMulticlusterServiceReconciler(ctx context.Context, rec MulticlusterServiceReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericServiceMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterServiceReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterServiceReconcileLoop {
	return &multiclusterServiceReconcileLoop{loop: multicluster.NewLoop(name, cw, &core_v1.Service{})}
}

type genericServiceMulticlusterReconciler struct {
	reconciler MulticlusterServiceReconciler
}

func (g genericServiceMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) {
	if deletionReconciler, ok := g.reconciler.(MulticlusterServiceDeletionReconciler); ok {
		deletionReconciler.ReconcileServiceDeletion(cluster, req)
	}
}

func (g genericServiceMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*core_v1.Service)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: Service handler received event for %T", object)
	}
	return g.reconciler.ReconcileService(cluster, obj)
}

// Reconcile Upsert events for the Pod Resource across clusters.
// implemented by the user
type MulticlusterPodReconciler interface {
	ReconcilePod(clusterName string, obj *core_v1.Pod) (reconcile.Result, error)
}

// Reconcile deletion events for the Pod Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterPodDeletionReconciler interface {
	ReconcilePodDeletion(clusterName string, req reconcile.Request)
}

type MulticlusterPodReconcilerFuncs struct {
	OnReconcilePod         func(clusterName string, obj *core_v1.Pod) (reconcile.Result, error)
	OnReconcilePodDeletion func(clusterName string, req reconcile.Request)
}

func (f *MulticlusterPodReconcilerFuncs) ReconcilePod(clusterName string, obj *core_v1.Pod) (reconcile.Result, error) {
	if f.OnReconcilePod == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcilePod(clusterName, obj)
}

func (f *MulticlusterPodReconcilerFuncs) ReconcilePodDeletion(clusterName string, req reconcile.Request) {
	if f.OnReconcilePodDeletion == nil {
		return
	}
	f.OnReconcilePodDeletion(clusterName, req)
}

type MulticlusterPodReconcileLoop interface {
	// AddMulticlusterPodReconciler adds a MulticlusterPodReconciler to the MulticlusterPodReconcileLoop.
	AddMulticlusterPodReconciler(ctx context.Context, rec MulticlusterPodReconciler, predicates ...predicate.Predicate)
}

type multiclusterPodReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterPodReconcileLoop) AddMulticlusterPodReconciler(ctx context.Context, rec MulticlusterPodReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericPodMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterPodReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterPodReconcileLoop {
	return &multiclusterPodReconcileLoop{loop: multicluster.NewLoop(name, cw, &core_v1.Pod{})}
}

type genericPodMulticlusterReconciler struct {
	reconciler MulticlusterPodReconciler
}

func (g genericPodMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) {
	if deletionReconciler, ok := g.reconciler.(MulticlusterPodDeletionReconciler); ok {
		deletionReconciler.ReconcilePodDeletion(cluster, req)
	}
}

func (g genericPodMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*core_v1.Pod)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: Pod handler received event for %T", object)
	}
	return g.reconciler.ReconcilePod(cluster, obj)
}

// Reconcile Upsert events for the Namespace Resource across clusters.
// implemented by the user
type MulticlusterNamespaceReconciler interface {
	ReconcileNamespace(clusterName string, obj *core_v1.Namespace) (reconcile.Result, error)
}

// Reconcile deletion events for the Namespace Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterNamespaceDeletionReconciler interface {
	ReconcileNamespaceDeletion(clusterName string, req reconcile.Request)
}

type MulticlusterNamespaceReconcilerFuncs struct {
	OnReconcileNamespace         func(clusterName string, obj *core_v1.Namespace) (reconcile.Result, error)
	OnReconcileNamespaceDeletion func(clusterName string, req reconcile.Request)
}

func (f *MulticlusterNamespaceReconcilerFuncs) ReconcileNamespace(clusterName string, obj *core_v1.Namespace) (reconcile.Result, error) {
	if f.OnReconcileNamespace == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileNamespace(clusterName, obj)
}

func (f *MulticlusterNamespaceReconcilerFuncs) ReconcileNamespaceDeletion(clusterName string, req reconcile.Request) {
	if f.OnReconcileNamespaceDeletion == nil {
		return
	}
	f.OnReconcileNamespaceDeletion(clusterName, req)
}

type MulticlusterNamespaceReconcileLoop interface {
	// AddMulticlusterNamespaceReconciler adds a MulticlusterNamespaceReconciler to the MulticlusterNamespaceReconcileLoop.
	AddMulticlusterNamespaceReconciler(ctx context.Context, rec MulticlusterNamespaceReconciler, predicates ...predicate.Predicate)
}

type multiclusterNamespaceReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterNamespaceReconcileLoop) AddMulticlusterNamespaceReconciler(ctx context.Context, rec MulticlusterNamespaceReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericNamespaceMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterNamespaceReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterNamespaceReconcileLoop {
	return &multiclusterNamespaceReconcileLoop{loop: multicluster.NewLoop(name, cw, &core_v1.Namespace{})}
}

type genericNamespaceMulticlusterReconciler struct {
	reconciler MulticlusterNamespaceReconciler
}

func (g genericNamespaceMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) {
	if deletionReconciler, ok := g.reconciler.(MulticlusterNamespaceDeletionReconciler); ok {
		deletionReconciler.ReconcileNamespaceDeletion(cluster, req)
	}
}

func (g genericNamespaceMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*core_v1.Namespace)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: Namespace handler received event for %T", object)
	}
	return g.reconciler.ReconcileNamespace(cluster, obj)
}

// Reconcile Upsert events for the Node Resource across clusters.
// implemented by the user
type MulticlusterNodeReconciler interface {
	ReconcileNode(clusterName string, obj *core_v1.Node) (reconcile.Result, error)
}

// Reconcile deletion events for the Node Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterNodeDeletionReconciler interface {
	ReconcileNodeDeletion(clusterName string, req reconcile.Request)
}

type MulticlusterNodeReconcilerFuncs struct {
	OnReconcileNode         func(clusterName string, obj *core_v1.Node) (reconcile.Result, error)
	OnReconcileNodeDeletion func(clusterName string, req reconcile.Request)
}

func (f *MulticlusterNodeReconcilerFuncs) ReconcileNode(clusterName string, obj *core_v1.Node) (reconcile.Result, error) {
	if f.OnReconcileNode == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileNode(clusterName, obj)
}

func (f *MulticlusterNodeReconcilerFuncs) ReconcileNodeDeletion(clusterName string, req reconcile.Request) {
	if f.OnReconcileNodeDeletion == nil {
		return
	}
	f.OnReconcileNodeDeletion(clusterName, req)
}

type MulticlusterNodeReconcileLoop interface {
	// AddMulticlusterNodeReconciler adds a MulticlusterNodeReconciler to the MulticlusterNodeReconcileLoop.
	AddMulticlusterNodeReconciler(ctx context.Context, rec MulticlusterNodeReconciler, predicates ...predicate.Predicate)
}

type multiclusterNodeReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterNodeReconcileLoop) AddMulticlusterNodeReconciler(ctx context.Context, rec MulticlusterNodeReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericNodeMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterNodeReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterNodeReconcileLoop {
	return &multiclusterNodeReconcileLoop{loop: multicluster.NewLoop(name, cw, &core_v1.Node{})}
}

type genericNodeMulticlusterReconciler struct {
	reconciler MulticlusterNodeReconciler
}

func (g genericNodeMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) {
	if deletionReconciler, ok := g.reconciler.(MulticlusterNodeDeletionReconciler); ok {
		deletionReconciler.ReconcileNodeDeletion(cluster, req)
	}
}

func (g genericNodeMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*core_v1.Node)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: Node handler received event for %T", object)
	}
	return g.reconciler.ReconcileNode(cluster, obj)
}
