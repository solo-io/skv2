// Code generated by MockGen. DO NOT EDIT.
// Source: ./multicluster_reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	controller "github.com/solo-io/skv2/pkg/multicluster/internal/k8s/apiextensions.k8s.io/v1beta1/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	reflect "reflect"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMulticlusterCustomResourceDefinitionReconciler is a mock of MulticlusterCustomResourceDefinitionReconciler interface
type MockMulticlusterCustomResourceDefinitionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterCustomResourceDefinitionReconcilerMockRecorder
}

// MockMulticlusterCustomResourceDefinitionReconcilerMockRecorder is the mock recorder for MockMulticlusterCustomResourceDefinitionReconciler
type MockMulticlusterCustomResourceDefinitionReconcilerMockRecorder struct {
	mock *MockMulticlusterCustomResourceDefinitionReconciler
}

// NewMockMulticlusterCustomResourceDefinitionReconciler creates a new mock instance
func NewMockMulticlusterCustomResourceDefinitionReconciler(ctrl *gomock.Controller) *MockMulticlusterCustomResourceDefinitionReconciler {
	mock := &MockMulticlusterCustomResourceDefinitionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterCustomResourceDefinitionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterCustomResourceDefinitionReconciler) EXPECT() *MockMulticlusterCustomResourceDefinitionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileCustomResourceDefinition mocks base method
func (m *MockMulticlusterCustomResourceDefinitionReconciler) ReconcileCustomResourceDefinition(clusterName string, obj *v1.CustomResourceDefinition) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileCustomResourceDefinition", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileCustomResourceDefinition indicates an expected call of ReconcileCustomResourceDefinition
func (mr *MockMulticlusterCustomResourceDefinitionReconcilerMockRecorder) ReconcileCustomResourceDefinition(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileCustomResourceDefinition", reflect.TypeOf((*MockMulticlusterCustomResourceDefinitionReconciler)(nil).ReconcileCustomResourceDefinition), clusterName, obj)
}

// MockMulticlusterCustomResourceDefinitionDeletionReconciler is a mock of MulticlusterCustomResourceDefinitionDeletionReconciler interface
type MockMulticlusterCustomResourceDefinitionDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterCustomResourceDefinitionDeletionReconcilerMockRecorder
}

// MockMulticlusterCustomResourceDefinitionDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterCustomResourceDefinitionDeletionReconciler
type MockMulticlusterCustomResourceDefinitionDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterCustomResourceDefinitionDeletionReconciler
}

// NewMockMulticlusterCustomResourceDefinitionDeletionReconciler creates a new mock instance
func NewMockMulticlusterCustomResourceDefinitionDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterCustomResourceDefinitionDeletionReconciler {
	mock := &MockMulticlusterCustomResourceDefinitionDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterCustomResourceDefinitionDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterCustomResourceDefinitionDeletionReconciler) EXPECT() *MockMulticlusterCustomResourceDefinitionDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileCustomResourceDefinitionDeletion mocks base method
func (m *MockMulticlusterCustomResourceDefinitionDeletionReconciler) ReconcileCustomResourceDefinitionDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileCustomResourceDefinitionDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileCustomResourceDefinitionDeletion indicates an expected call of ReconcileCustomResourceDefinitionDeletion
func (mr *MockMulticlusterCustomResourceDefinitionDeletionReconcilerMockRecorder) ReconcileCustomResourceDefinitionDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileCustomResourceDefinitionDeletion", reflect.TypeOf((*MockMulticlusterCustomResourceDefinitionDeletionReconciler)(nil).ReconcileCustomResourceDefinitionDeletion), clusterName, req)
}

// MockMulticlusterCustomResourceDefinitionReconcileLoop is a mock of MulticlusterCustomResourceDefinitionReconcileLoop interface
type MockMulticlusterCustomResourceDefinitionReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterCustomResourceDefinitionReconcileLoopMockRecorder
}

// MockMulticlusterCustomResourceDefinitionReconcileLoopMockRecorder is the mock recorder for MockMulticlusterCustomResourceDefinitionReconcileLoop
type MockMulticlusterCustomResourceDefinitionReconcileLoopMockRecorder struct {
	mock *MockMulticlusterCustomResourceDefinitionReconcileLoop
}

// NewMockMulticlusterCustomResourceDefinitionReconcileLoop creates a new mock instance
func NewMockMulticlusterCustomResourceDefinitionReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterCustomResourceDefinitionReconcileLoop {
	mock := &MockMulticlusterCustomResourceDefinitionReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterCustomResourceDefinitionReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterCustomResourceDefinitionReconcileLoop) EXPECT() *MockMulticlusterCustomResourceDefinitionReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterCustomResourceDefinitionReconciler mocks base method
func (m *MockMulticlusterCustomResourceDefinitionReconcileLoop) AddMulticlusterCustomResourceDefinitionReconciler(ctx context.Context, rec controller.MulticlusterCustomResourceDefinitionReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterCustomResourceDefinitionReconciler", varargs...)
}

// AddMulticlusterCustomResourceDefinitionReconciler indicates an expected call of AddMulticlusterCustomResourceDefinitionReconciler
func (mr *MockMulticlusterCustomResourceDefinitionReconcileLoopMockRecorder) AddMulticlusterCustomResourceDefinitionReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterCustomResourceDefinitionReconciler", reflect.TypeOf((*MockMulticlusterCustomResourceDefinitionReconcileLoop)(nil).AddMulticlusterCustomResourceDefinitionReconciler), varargs...)
}
