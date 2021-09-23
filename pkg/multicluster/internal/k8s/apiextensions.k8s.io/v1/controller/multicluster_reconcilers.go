// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller

import (
	"context"

	apiextensions_k8s_io_v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the CustomResourceDefinition Resource across clusters.
// implemented by the user
type MulticlusterCustomResourceDefinitionReconciler interface {
	ReconcileCustomResourceDefinition(clusterName string, obj *apiextensions_k8s_io_v1.CustomResourceDefinition) (reconcile.Result, error)
}

// Reconcile deletion events for the CustomResourceDefinition Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterCustomResourceDefinitionDeletionReconciler interface {
	ReconcileCustomResourceDefinitionDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterCustomResourceDefinitionReconcilerFuncs struct {
	OnReconcileCustomResourceDefinition         func(clusterName string, obj *apiextensions_k8s_io_v1.CustomResourceDefinition) (reconcile.Result, error)
	OnReconcileCustomResourceDefinitionDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterCustomResourceDefinitionReconcilerFuncs) ReconcileCustomResourceDefinition(clusterName string, obj *apiextensions_k8s_io_v1.CustomResourceDefinition) (reconcile.Result, error) {
	if f.OnReconcileCustomResourceDefinition == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileCustomResourceDefinition(clusterName, obj)
}

func (f *MulticlusterCustomResourceDefinitionReconcilerFuncs) ReconcileCustomResourceDefinitionDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileCustomResourceDefinitionDeletion == nil {
		return nil
	}
	return f.OnReconcileCustomResourceDefinitionDeletion(clusterName, req)
}

type MulticlusterCustomResourceDefinitionReconcileLoop interface {
	// AddMulticlusterCustomResourceDefinitionReconciler adds a MulticlusterCustomResourceDefinitionReconciler to the MulticlusterCustomResourceDefinitionReconcileLoop.
	AddMulticlusterCustomResourceDefinitionReconciler(ctx context.Context, rec MulticlusterCustomResourceDefinitionReconciler, predicates ...predicate.Predicate)
}

type multiclusterCustomResourceDefinitionReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterCustomResourceDefinitionReconcileLoop) AddMulticlusterCustomResourceDefinitionReconciler(ctx context.Context, rec MulticlusterCustomResourceDefinitionReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericCustomResourceDefinitionMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterCustomResourceDefinitionReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterCustomResourceDefinitionReconcileLoop {
	return &multiclusterCustomResourceDefinitionReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &apiextensions_k8s_io_v1.CustomResourceDefinition{}, options)}
}

type genericCustomResourceDefinitionMulticlusterReconciler struct {
	reconciler MulticlusterCustomResourceDefinitionReconciler
}

func (g genericCustomResourceDefinitionMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterCustomResourceDefinitionDeletionReconciler); ok {
		return deletionReconciler.ReconcileCustomResourceDefinitionDeletion(cluster, req)
	}
	return nil
}

func (g genericCustomResourceDefinitionMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*apiextensions_k8s_io_v1.CustomResourceDefinition)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: CustomResourceDefinition handler received event for %T", object)
	}
	return g.reconciler.ReconcileCustomResourceDefinition(cluster, obj)
}