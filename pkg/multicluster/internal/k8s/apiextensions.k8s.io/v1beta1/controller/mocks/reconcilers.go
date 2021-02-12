// Code generated by MockGen. DO NOT EDIT.
// Source: ./reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	controller "github.com/solo-io/skv2/pkg/multicluster/internal/k8s/apiextensions.k8s.io/v1beta1/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockCustomResourceDefinitionReconciler is a mock of CustomResourceDefinitionReconciler interface.
type MockCustomResourceDefinitionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockCustomResourceDefinitionReconcilerMockRecorder
}

// MockCustomResourceDefinitionReconcilerMockRecorder is the mock recorder for MockCustomResourceDefinitionReconciler.
type MockCustomResourceDefinitionReconcilerMockRecorder struct {
	mock *MockCustomResourceDefinitionReconciler
}

// NewMockCustomResourceDefinitionReconciler creates a new mock instance.
func NewMockCustomResourceDefinitionReconciler(ctrl *gomock.Controller) *MockCustomResourceDefinitionReconciler {
	mock := &MockCustomResourceDefinitionReconciler{ctrl: ctrl}
	mock.recorder = &MockCustomResourceDefinitionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomResourceDefinitionReconciler) EXPECT() *MockCustomResourceDefinitionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileCustomResourceDefinition mocks base method.
func (m *MockCustomResourceDefinitionReconciler) ReconcileCustomResourceDefinition(obj *v1beta1.CustomResourceDefinition) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileCustomResourceDefinition", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileCustomResourceDefinition indicates an expected call of ReconcileCustomResourceDefinition.
func (mr *MockCustomResourceDefinitionReconcilerMockRecorder) ReconcileCustomResourceDefinition(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileCustomResourceDefinition", reflect.TypeOf((*MockCustomResourceDefinitionReconciler)(nil).ReconcileCustomResourceDefinition), obj)
}

// MockCustomResourceDefinitionDeletionReconciler is a mock of CustomResourceDefinitionDeletionReconciler interface.
type MockCustomResourceDefinitionDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockCustomResourceDefinitionDeletionReconcilerMockRecorder
}

// MockCustomResourceDefinitionDeletionReconcilerMockRecorder is the mock recorder for MockCustomResourceDefinitionDeletionReconciler.
type MockCustomResourceDefinitionDeletionReconcilerMockRecorder struct {
	mock *MockCustomResourceDefinitionDeletionReconciler
}

// NewMockCustomResourceDefinitionDeletionReconciler creates a new mock instance.
func NewMockCustomResourceDefinitionDeletionReconciler(ctrl *gomock.Controller) *MockCustomResourceDefinitionDeletionReconciler {
	mock := &MockCustomResourceDefinitionDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockCustomResourceDefinitionDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomResourceDefinitionDeletionReconciler) EXPECT() *MockCustomResourceDefinitionDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileCustomResourceDefinitionDeletion mocks base method.
func (m *MockCustomResourceDefinitionDeletionReconciler) ReconcileCustomResourceDefinitionDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileCustomResourceDefinitionDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileCustomResourceDefinitionDeletion indicates an expected call of ReconcileCustomResourceDefinitionDeletion.
func (mr *MockCustomResourceDefinitionDeletionReconcilerMockRecorder) ReconcileCustomResourceDefinitionDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileCustomResourceDefinitionDeletion", reflect.TypeOf((*MockCustomResourceDefinitionDeletionReconciler)(nil).ReconcileCustomResourceDefinitionDeletion), req)
}

// MockCustomResourceDefinitionFinalizer is a mock of CustomResourceDefinitionFinalizer interface.
type MockCustomResourceDefinitionFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockCustomResourceDefinitionFinalizerMockRecorder
}

// MockCustomResourceDefinitionFinalizerMockRecorder is the mock recorder for MockCustomResourceDefinitionFinalizer.
type MockCustomResourceDefinitionFinalizerMockRecorder struct {
	mock *MockCustomResourceDefinitionFinalizer
}

// NewMockCustomResourceDefinitionFinalizer creates a new mock instance.
func NewMockCustomResourceDefinitionFinalizer(ctrl *gomock.Controller) *MockCustomResourceDefinitionFinalizer {
	mock := &MockCustomResourceDefinitionFinalizer{ctrl: ctrl}
	mock.recorder = &MockCustomResourceDefinitionFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomResourceDefinitionFinalizer) EXPECT() *MockCustomResourceDefinitionFinalizerMockRecorder {
	return m.recorder
}

// ReconcileCustomResourceDefinition mocks base method.
func (m *MockCustomResourceDefinitionFinalizer) ReconcileCustomResourceDefinition(obj *v1beta1.CustomResourceDefinition) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileCustomResourceDefinition", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileCustomResourceDefinition indicates an expected call of ReconcileCustomResourceDefinition.
func (mr *MockCustomResourceDefinitionFinalizerMockRecorder) ReconcileCustomResourceDefinition(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileCustomResourceDefinition", reflect.TypeOf((*MockCustomResourceDefinitionFinalizer)(nil).ReconcileCustomResourceDefinition), obj)
}

// CustomResourceDefinitionFinalizerName mocks base method.
func (m *MockCustomResourceDefinitionFinalizer) CustomResourceDefinitionFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CustomResourceDefinitionFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// CustomResourceDefinitionFinalizerName indicates an expected call of CustomResourceDefinitionFinalizerName.
func (mr *MockCustomResourceDefinitionFinalizerMockRecorder) CustomResourceDefinitionFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CustomResourceDefinitionFinalizerName", reflect.TypeOf((*MockCustomResourceDefinitionFinalizer)(nil).CustomResourceDefinitionFinalizerName))
}

// FinalizeCustomResourceDefinition mocks base method.
func (m *MockCustomResourceDefinitionFinalizer) FinalizeCustomResourceDefinition(obj *v1beta1.CustomResourceDefinition) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeCustomResourceDefinition", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeCustomResourceDefinition indicates an expected call of FinalizeCustomResourceDefinition.
func (mr *MockCustomResourceDefinitionFinalizerMockRecorder) FinalizeCustomResourceDefinition(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeCustomResourceDefinition", reflect.TypeOf((*MockCustomResourceDefinitionFinalizer)(nil).FinalizeCustomResourceDefinition), obj)
}

// MockCustomResourceDefinitionReconcileLoop is a mock of CustomResourceDefinitionReconcileLoop interface.
type MockCustomResourceDefinitionReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockCustomResourceDefinitionReconcileLoopMockRecorder
}

// MockCustomResourceDefinitionReconcileLoopMockRecorder is the mock recorder for MockCustomResourceDefinitionReconcileLoop.
type MockCustomResourceDefinitionReconcileLoopMockRecorder struct {
	mock *MockCustomResourceDefinitionReconcileLoop
}

// NewMockCustomResourceDefinitionReconcileLoop creates a new mock instance.
func NewMockCustomResourceDefinitionReconcileLoop(ctrl *gomock.Controller) *MockCustomResourceDefinitionReconcileLoop {
	mock := &MockCustomResourceDefinitionReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockCustomResourceDefinitionReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomResourceDefinitionReconcileLoop) EXPECT() *MockCustomResourceDefinitionReconcileLoopMockRecorder {
	return m.recorder
}

// RunCustomResourceDefinitionReconciler mocks base method.
func (m *MockCustomResourceDefinitionReconcileLoop) RunCustomResourceDefinitionReconciler(ctx context.Context, rec controller.CustomResourceDefinitionReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunCustomResourceDefinitionReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunCustomResourceDefinitionReconciler indicates an expected call of RunCustomResourceDefinitionReconciler.
func (mr *MockCustomResourceDefinitionReconcileLoopMockRecorder) RunCustomResourceDefinitionReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunCustomResourceDefinitionReconciler", reflect.TypeOf((*MockCustomResourceDefinitionReconcileLoop)(nil).RunCustomResourceDefinitionReconciler), varargs...)
}
