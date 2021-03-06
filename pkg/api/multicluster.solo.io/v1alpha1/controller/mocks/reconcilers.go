// Code generated by MockGen. DO NOT EDIT.
// Source: ./reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1"
	controller "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockKubernetesClusterReconciler is a mock of KubernetesClusterReconciler interface
type MockKubernetesClusterReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockKubernetesClusterReconcilerMockRecorder
}

// MockKubernetesClusterReconcilerMockRecorder is the mock recorder for MockKubernetesClusterReconciler
type MockKubernetesClusterReconcilerMockRecorder struct {
	mock *MockKubernetesClusterReconciler
}

// NewMockKubernetesClusterReconciler creates a new mock instance
func NewMockKubernetesClusterReconciler(ctrl *gomock.Controller) *MockKubernetesClusterReconciler {
	mock := &MockKubernetesClusterReconciler{ctrl: ctrl}
	mock.recorder = &MockKubernetesClusterReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKubernetesClusterReconciler) EXPECT() *MockKubernetesClusterReconcilerMockRecorder {
	return m.recorder
}

// ReconcileKubernetesCluster mocks base method
func (m *MockKubernetesClusterReconciler) ReconcileKubernetesCluster(obj *v1alpha1.KubernetesCluster) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileKubernetesCluster", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileKubernetesCluster indicates an expected call of ReconcileKubernetesCluster
func (mr *MockKubernetesClusterReconcilerMockRecorder) ReconcileKubernetesCluster(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileKubernetesCluster", reflect.TypeOf((*MockKubernetesClusterReconciler)(nil).ReconcileKubernetesCluster), obj)
}

// MockKubernetesClusterDeletionReconciler is a mock of KubernetesClusterDeletionReconciler interface
type MockKubernetesClusterDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockKubernetesClusterDeletionReconcilerMockRecorder
}

// MockKubernetesClusterDeletionReconcilerMockRecorder is the mock recorder for MockKubernetesClusterDeletionReconciler
type MockKubernetesClusterDeletionReconcilerMockRecorder struct {
	mock *MockKubernetesClusterDeletionReconciler
}

// NewMockKubernetesClusterDeletionReconciler creates a new mock instance
func NewMockKubernetesClusterDeletionReconciler(ctrl *gomock.Controller) *MockKubernetesClusterDeletionReconciler {
	mock := &MockKubernetesClusterDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockKubernetesClusterDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKubernetesClusterDeletionReconciler) EXPECT() *MockKubernetesClusterDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileKubernetesClusterDeletion mocks base method
func (m *MockKubernetesClusterDeletionReconciler) ReconcileKubernetesClusterDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileKubernetesClusterDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileKubernetesClusterDeletion indicates an expected call of ReconcileKubernetesClusterDeletion
func (mr *MockKubernetesClusterDeletionReconcilerMockRecorder) ReconcileKubernetesClusterDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileKubernetesClusterDeletion", reflect.TypeOf((*MockKubernetesClusterDeletionReconciler)(nil).ReconcileKubernetesClusterDeletion), req)
}

// MockKubernetesClusterFinalizer is a mock of KubernetesClusterFinalizer interface
type MockKubernetesClusterFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockKubernetesClusterFinalizerMockRecorder
}

// MockKubernetesClusterFinalizerMockRecorder is the mock recorder for MockKubernetesClusterFinalizer
type MockKubernetesClusterFinalizerMockRecorder struct {
	mock *MockKubernetesClusterFinalizer
}

// NewMockKubernetesClusterFinalizer creates a new mock instance
func NewMockKubernetesClusterFinalizer(ctrl *gomock.Controller) *MockKubernetesClusterFinalizer {
	mock := &MockKubernetesClusterFinalizer{ctrl: ctrl}
	mock.recorder = &MockKubernetesClusterFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKubernetesClusterFinalizer) EXPECT() *MockKubernetesClusterFinalizerMockRecorder {
	return m.recorder
}

// ReconcileKubernetesCluster mocks base method
func (m *MockKubernetesClusterFinalizer) ReconcileKubernetesCluster(obj *v1alpha1.KubernetesCluster) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileKubernetesCluster", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileKubernetesCluster indicates an expected call of ReconcileKubernetesCluster
func (mr *MockKubernetesClusterFinalizerMockRecorder) ReconcileKubernetesCluster(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileKubernetesCluster", reflect.TypeOf((*MockKubernetesClusterFinalizer)(nil).ReconcileKubernetesCluster), obj)
}

// KubernetesClusterFinalizerName mocks base method
func (m *MockKubernetesClusterFinalizer) KubernetesClusterFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KubernetesClusterFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// KubernetesClusterFinalizerName indicates an expected call of KubernetesClusterFinalizerName
func (mr *MockKubernetesClusterFinalizerMockRecorder) KubernetesClusterFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KubernetesClusterFinalizerName", reflect.TypeOf((*MockKubernetesClusterFinalizer)(nil).KubernetesClusterFinalizerName))
}

// FinalizeKubernetesCluster mocks base method
func (m *MockKubernetesClusterFinalizer) FinalizeKubernetesCluster(obj *v1alpha1.KubernetesCluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeKubernetesCluster", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeKubernetesCluster indicates an expected call of FinalizeKubernetesCluster
func (mr *MockKubernetesClusterFinalizerMockRecorder) FinalizeKubernetesCluster(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeKubernetesCluster", reflect.TypeOf((*MockKubernetesClusterFinalizer)(nil).FinalizeKubernetesCluster), obj)
}

// MockKubernetesClusterReconcileLoop is a mock of KubernetesClusterReconcileLoop interface
type MockKubernetesClusterReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockKubernetesClusterReconcileLoopMockRecorder
}

// MockKubernetesClusterReconcileLoopMockRecorder is the mock recorder for MockKubernetesClusterReconcileLoop
type MockKubernetesClusterReconcileLoopMockRecorder struct {
	mock *MockKubernetesClusterReconcileLoop
}

// NewMockKubernetesClusterReconcileLoop creates a new mock instance
func NewMockKubernetesClusterReconcileLoop(ctrl *gomock.Controller) *MockKubernetesClusterReconcileLoop {
	mock := &MockKubernetesClusterReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockKubernetesClusterReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKubernetesClusterReconcileLoop) EXPECT() *MockKubernetesClusterReconcileLoopMockRecorder {
	return m.recorder
}

// RunKubernetesClusterReconciler mocks base method
func (m *MockKubernetesClusterReconcileLoop) RunKubernetesClusterReconciler(ctx context.Context, rec controller.KubernetesClusterReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunKubernetesClusterReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunKubernetesClusterReconciler indicates an expected call of RunKubernetesClusterReconciler
func (mr *MockKubernetesClusterReconcileLoopMockRecorder) RunKubernetesClusterReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunKubernetesClusterReconciler", reflect.TypeOf((*MockKubernetesClusterReconcileLoop)(nil).RunKubernetesClusterReconciler), varargs...)
}
