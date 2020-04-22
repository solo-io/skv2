// Definitions for the multicluster Kubernetes Controllers
package multicluster

import (
	"context"

	apps_v1 "k8s.io/api/apps/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the Deployment Resource across clusters.
// implemented by the user
type MulticlusterDeploymentReconciler interface {
	ReconcileDeployment(clusterName string, obj *apps_v1.Deployment) (reconcile.Result, error)
}

// Reconcile deletion events for the Deployment Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterDeploymentDeletionReconciler interface {
	ReconcileDeploymentDeletion(clusterName string, req reconcile.Request)
}

type MulticlusterDeploymentReconcilerFuncs struct {
	OnReconcileDeployment         func(clusterName string, obj *apps_v1.Deployment) (reconcile.Result, error)
	OnReconcileDeploymentDeletion func(clusterName string, req reconcile.Request)
}

func (f *MulticlusterDeploymentReconcilerFuncs) ReconcileDeployment(clusterName string, obj *apps_v1.Deployment) (reconcile.Result, error) {
	if f.OnReconcileDeployment == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileDeployment(clusterName, obj)
}

func (f *MulticlusterDeploymentReconcilerFuncs) ReconcileDeploymentDeletion(clusterName string, req reconcile.Request) {
	if f.OnReconcileDeploymentDeletion == nil {
		return
	}
	f.OnReconcileDeploymentDeletion(clusterName, req)
}

type MulticlusterDeploymentReconcileLoop interface {
	// AddMulticlusterDeploymentReconciler adds a MulticlusterDeploymentReconciler to the MulticlusterDeploymentReconcileLoop.
	AddMulticlusterDeploymentReconciler(ctx context.Context, rec MulticlusterDeploymentReconciler, predicates ...predicate.Predicate)
}

type multiclusterDeploymentReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterDeploymentReconcileLoop) AddMulticlusterDeploymentReconciler(ctx context.Context, rec MulticlusterDeploymentReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericDeploymentMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterDeploymentReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterDeploymentReconcileLoop {
	return &multiclusterDeploymentReconcileLoop{loop: multicluster.NewLoop(name, cw, &apps_v1.Deployment{})}
}

type genericDeploymentMulticlusterReconciler struct {
	reconciler MulticlusterDeploymentReconciler
}

func (g genericDeploymentMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) {
	if deletionReconciler, ok := g.reconciler.(MulticlusterDeploymentDeletionReconciler); ok {
		deletionReconciler.ReconcileDeploymentDeletion(cluster, req)
	}
}

func (g genericDeploymentMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*apps_v1.Deployment)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: Deployment handler received event for %T", object)
	}
	return g.reconciler.ReconcileDeployment(cluster, obj)
}

// Reconcile Upsert events for the ReplicaSet Resource across clusters.
// implemented by the user
type MulticlusterReplicaSetReconciler interface {
	ReconcileReplicaSet(clusterName string, obj *apps_v1.ReplicaSet) (reconcile.Result, error)
}

// Reconcile deletion events for the ReplicaSet Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterReplicaSetDeletionReconciler interface {
	ReconcileReplicaSetDeletion(clusterName string, req reconcile.Request)
}

type MulticlusterReplicaSetReconcilerFuncs struct {
	OnReconcileReplicaSet         func(clusterName string, obj *apps_v1.ReplicaSet) (reconcile.Result, error)
	OnReconcileReplicaSetDeletion func(clusterName string, req reconcile.Request)
}

func (f *MulticlusterReplicaSetReconcilerFuncs) ReconcileReplicaSet(clusterName string, obj *apps_v1.ReplicaSet) (reconcile.Result, error) {
	if f.OnReconcileReplicaSet == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileReplicaSet(clusterName, obj)
}

func (f *MulticlusterReplicaSetReconcilerFuncs) ReconcileReplicaSetDeletion(clusterName string, req reconcile.Request) {
	if f.OnReconcileReplicaSetDeletion == nil {
		return
	}
	f.OnReconcileReplicaSetDeletion(clusterName, req)
}

type MulticlusterReplicaSetReconcileLoop interface {
	// AddMulticlusterReplicaSetReconciler adds a MulticlusterReplicaSetReconciler to the MulticlusterReplicaSetReconcileLoop.
	AddMulticlusterReplicaSetReconciler(ctx context.Context, rec MulticlusterReplicaSetReconciler, predicates ...predicate.Predicate)
}

type multiclusterReplicaSetReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterReplicaSetReconcileLoop) AddMulticlusterReplicaSetReconciler(ctx context.Context, rec MulticlusterReplicaSetReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericReplicaSetMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterReplicaSetReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterReplicaSetReconcileLoop {
	return &multiclusterReplicaSetReconcileLoop{loop: multicluster.NewLoop(name, cw, &apps_v1.ReplicaSet{})}
}

type genericReplicaSetMulticlusterReconciler struct {
	reconciler MulticlusterReplicaSetReconciler
}

func (g genericReplicaSetMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) {
	if deletionReconciler, ok := g.reconciler.(MulticlusterReplicaSetDeletionReconciler); ok {
		deletionReconciler.ReconcileReplicaSetDeletion(cluster, req)
	}
}

func (g genericReplicaSetMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*apps_v1.ReplicaSet)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: ReplicaSet handler received event for %T", object)
	}
	return g.reconciler.ReconcileReplicaSet(cluster, obj)
}
