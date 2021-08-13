// Code generated by MockGen. DO NOT EDIT.
// Source: ./reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1"
	controller "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockPaintReconciler is a mock of PaintReconciler interface.
type MockPaintReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockPaintReconcilerMockRecorder
}

// MockPaintReconcilerMockRecorder is the mock recorder for MockPaintReconciler.
type MockPaintReconcilerMockRecorder struct {
	mock *MockPaintReconciler
}

// NewMockPaintReconciler creates a new mock instance.
func NewMockPaintReconciler(ctrl *gomock.Controller) *MockPaintReconciler {
	mock := &MockPaintReconciler{ctrl: ctrl}
	mock.recorder = &MockPaintReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaintReconciler) EXPECT() *MockPaintReconcilerMockRecorder {
	return m.recorder
}

// ReconcilePaint mocks base method.
func (m *MockPaintReconciler) ReconcilePaint(obj *v1.Paint) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcilePaint", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcilePaint indicates an expected call of ReconcilePaint.
func (mr *MockPaintReconcilerMockRecorder) ReconcilePaint(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcilePaint", reflect.TypeOf((*MockPaintReconciler)(nil).ReconcilePaint), obj)
}

// MockPaintDeletionReconciler is a mock of PaintDeletionReconciler interface.
type MockPaintDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockPaintDeletionReconcilerMockRecorder
}

// MockPaintDeletionReconcilerMockRecorder is the mock recorder for MockPaintDeletionReconciler.
type MockPaintDeletionReconcilerMockRecorder struct {
	mock *MockPaintDeletionReconciler
}

// NewMockPaintDeletionReconciler creates a new mock instance.
func NewMockPaintDeletionReconciler(ctrl *gomock.Controller) *MockPaintDeletionReconciler {
	mock := &MockPaintDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockPaintDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaintDeletionReconciler) EXPECT() *MockPaintDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcilePaintDeletion mocks base method.
func (m *MockPaintDeletionReconciler) ReconcilePaintDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcilePaintDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcilePaintDeletion indicates an expected call of ReconcilePaintDeletion.
func (mr *MockPaintDeletionReconcilerMockRecorder) ReconcilePaintDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcilePaintDeletion", reflect.TypeOf((*MockPaintDeletionReconciler)(nil).ReconcilePaintDeletion), req)
}

// MockPaintFinalizer is a mock of PaintFinalizer interface.
type MockPaintFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockPaintFinalizerMockRecorder
}

// MockPaintFinalizerMockRecorder is the mock recorder for MockPaintFinalizer.
type MockPaintFinalizerMockRecorder struct {
	mock *MockPaintFinalizer
}

// NewMockPaintFinalizer creates a new mock instance.
func NewMockPaintFinalizer(ctrl *gomock.Controller) *MockPaintFinalizer {
	mock := &MockPaintFinalizer{ctrl: ctrl}
	mock.recorder = &MockPaintFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaintFinalizer) EXPECT() *MockPaintFinalizerMockRecorder {
	return m.recorder
}

// FinalizePaint mocks base method.
func (m *MockPaintFinalizer) FinalizePaint(obj *v1.Paint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizePaint", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizePaint indicates an expected call of FinalizePaint.
func (mr *MockPaintFinalizerMockRecorder) FinalizePaint(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizePaint", reflect.TypeOf((*MockPaintFinalizer)(nil).FinalizePaint), obj)
}

// PaintFinalizerName mocks base method.
func (m *MockPaintFinalizer) PaintFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PaintFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// PaintFinalizerName indicates an expected call of PaintFinalizerName.
func (mr *MockPaintFinalizerMockRecorder) PaintFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaintFinalizerName", reflect.TypeOf((*MockPaintFinalizer)(nil).PaintFinalizerName))
}

// ReconcilePaint mocks base method.
func (m *MockPaintFinalizer) ReconcilePaint(obj *v1.Paint) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcilePaint", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcilePaint indicates an expected call of ReconcilePaint.
func (mr *MockPaintFinalizerMockRecorder) ReconcilePaint(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcilePaint", reflect.TypeOf((*MockPaintFinalizer)(nil).ReconcilePaint), obj)
}

// MockPaintReconcileLoop is a mock of PaintReconcileLoop interface.
type MockPaintReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockPaintReconcileLoopMockRecorder
}

// MockPaintReconcileLoopMockRecorder is the mock recorder for MockPaintReconcileLoop.
type MockPaintReconcileLoopMockRecorder struct {
	mock *MockPaintReconcileLoop
}

// NewMockPaintReconcileLoop creates a new mock instance.
func NewMockPaintReconcileLoop(ctrl *gomock.Controller) *MockPaintReconcileLoop {
	mock := &MockPaintReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockPaintReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaintReconcileLoop) EXPECT() *MockPaintReconcileLoopMockRecorder {
	return m.recorder
}

// RunPaintReconciler mocks base method.
func (m *MockPaintReconcileLoop) RunPaintReconciler(ctx context.Context, rec controller.PaintReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunPaintReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunPaintReconciler indicates an expected call of RunPaintReconciler.
func (mr *MockPaintReconcileLoopMockRecorder) RunPaintReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunPaintReconciler", reflect.TypeOf((*MockPaintReconcileLoop)(nil).RunPaintReconciler), varargs...)
}

// MockClusterResourceReconciler is a mock of ClusterResourceReconciler interface.
type MockClusterResourceReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockClusterResourceReconcilerMockRecorder
}

// MockClusterResourceReconcilerMockRecorder is the mock recorder for MockClusterResourceReconciler.
type MockClusterResourceReconcilerMockRecorder struct {
	mock *MockClusterResourceReconciler
}

// NewMockClusterResourceReconciler creates a new mock instance.
func NewMockClusterResourceReconciler(ctrl *gomock.Controller) *MockClusterResourceReconciler {
	mock := &MockClusterResourceReconciler{ctrl: ctrl}
	mock.recorder = &MockClusterResourceReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterResourceReconciler) EXPECT() *MockClusterResourceReconcilerMockRecorder {
	return m.recorder
}

// ReconcileClusterResource mocks base method.
func (m *MockClusterResourceReconciler) ReconcileClusterResource(obj *v1.ClusterResource) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileClusterResource", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileClusterResource indicates an expected call of ReconcileClusterResource.
func (mr *MockClusterResourceReconcilerMockRecorder) ReconcileClusterResource(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileClusterResource", reflect.TypeOf((*MockClusterResourceReconciler)(nil).ReconcileClusterResource), obj)
}

// MockClusterResourceDeletionReconciler is a mock of ClusterResourceDeletionReconciler interface.
type MockClusterResourceDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockClusterResourceDeletionReconcilerMockRecorder
}

// MockClusterResourceDeletionReconcilerMockRecorder is the mock recorder for MockClusterResourceDeletionReconciler.
type MockClusterResourceDeletionReconcilerMockRecorder struct {
	mock *MockClusterResourceDeletionReconciler
}

// NewMockClusterResourceDeletionReconciler creates a new mock instance.
func NewMockClusterResourceDeletionReconciler(ctrl *gomock.Controller) *MockClusterResourceDeletionReconciler {
	mock := &MockClusterResourceDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockClusterResourceDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterResourceDeletionReconciler) EXPECT() *MockClusterResourceDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileClusterResourceDeletion mocks base method.
func (m *MockClusterResourceDeletionReconciler) ReconcileClusterResourceDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileClusterResourceDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileClusterResourceDeletion indicates an expected call of ReconcileClusterResourceDeletion.
func (mr *MockClusterResourceDeletionReconcilerMockRecorder) ReconcileClusterResourceDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileClusterResourceDeletion", reflect.TypeOf((*MockClusterResourceDeletionReconciler)(nil).ReconcileClusterResourceDeletion), req)
}

// MockClusterResourceFinalizer is a mock of ClusterResourceFinalizer interface.
type MockClusterResourceFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockClusterResourceFinalizerMockRecorder
}

// MockClusterResourceFinalizerMockRecorder is the mock recorder for MockClusterResourceFinalizer.
type MockClusterResourceFinalizerMockRecorder struct {
	mock *MockClusterResourceFinalizer
}

// NewMockClusterResourceFinalizer creates a new mock instance.
func NewMockClusterResourceFinalizer(ctrl *gomock.Controller) *MockClusterResourceFinalizer {
	mock := &MockClusterResourceFinalizer{ctrl: ctrl}
	mock.recorder = &MockClusterResourceFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterResourceFinalizer) EXPECT() *MockClusterResourceFinalizerMockRecorder {
	return m.recorder
}

// ClusterResourceFinalizerName mocks base method.
func (m *MockClusterResourceFinalizer) ClusterResourceFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterResourceFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClusterResourceFinalizerName indicates an expected call of ClusterResourceFinalizerName.
func (mr *MockClusterResourceFinalizerMockRecorder) ClusterResourceFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterResourceFinalizerName", reflect.TypeOf((*MockClusterResourceFinalizer)(nil).ClusterResourceFinalizerName))
}

// FinalizeClusterResource mocks base method.
func (m *MockClusterResourceFinalizer) FinalizeClusterResource(obj *v1.ClusterResource) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeClusterResource", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeClusterResource indicates an expected call of FinalizeClusterResource.
func (mr *MockClusterResourceFinalizerMockRecorder) FinalizeClusterResource(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeClusterResource", reflect.TypeOf((*MockClusterResourceFinalizer)(nil).FinalizeClusterResource), obj)
}

// ReconcileClusterResource mocks base method.
func (m *MockClusterResourceFinalizer) ReconcileClusterResource(obj *v1.ClusterResource) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileClusterResource", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileClusterResource indicates an expected call of ReconcileClusterResource.
func (mr *MockClusterResourceFinalizerMockRecorder) ReconcileClusterResource(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileClusterResource", reflect.TypeOf((*MockClusterResourceFinalizer)(nil).ReconcileClusterResource), obj)
}

// MockClusterResourceReconcileLoop is a mock of ClusterResourceReconcileLoop interface.
type MockClusterResourceReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockClusterResourceReconcileLoopMockRecorder
}

// MockClusterResourceReconcileLoopMockRecorder is the mock recorder for MockClusterResourceReconcileLoop.
type MockClusterResourceReconcileLoopMockRecorder struct {
	mock *MockClusterResourceReconcileLoop
}

// NewMockClusterResourceReconcileLoop creates a new mock instance.
func NewMockClusterResourceReconcileLoop(ctrl *gomock.Controller) *MockClusterResourceReconcileLoop {
	mock := &MockClusterResourceReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockClusterResourceReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterResourceReconcileLoop) EXPECT() *MockClusterResourceReconcileLoopMockRecorder {
	return m.recorder
}

// RunClusterResourceReconciler mocks base method.
func (m *MockClusterResourceReconcileLoop) RunClusterResourceReconciler(ctx context.Context, rec controller.ClusterResourceReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunClusterResourceReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunClusterResourceReconciler indicates an expected call of RunClusterResourceReconciler.
func (mr *MockClusterResourceReconcileLoopMockRecorder) RunClusterResourceReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunClusterResourceReconciler", reflect.TypeOf((*MockClusterResourceReconcileLoop)(nil).RunClusterResourceReconciler), varargs...)
}
