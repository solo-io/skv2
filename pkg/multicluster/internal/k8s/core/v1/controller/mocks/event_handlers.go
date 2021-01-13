// Code generated by MockGen. DO NOT EDIT.
// Source: ./event_handlers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	controller "github.com/solo-io/skv2/pkg/multicluster/internal/k8s/core/v1/controller"
	v1 "k8s.io/api/core/v1"
	reflect "reflect"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockSecretEventHandler is a mock of SecretEventHandler interface
type MockSecretEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockSecretEventHandlerMockRecorder
}

// MockSecretEventHandlerMockRecorder is the mock recorder for MockSecretEventHandler
type MockSecretEventHandlerMockRecorder struct {
	mock *MockSecretEventHandler
}

// NewMockSecretEventHandler creates a new mock instance
func NewMockSecretEventHandler(ctrl *gomock.Controller) *MockSecretEventHandler {
	mock := &MockSecretEventHandler{ctrl: ctrl}
	mock.recorder = &MockSecretEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSecretEventHandler) EXPECT() *MockSecretEventHandlerMockRecorder {
	return m.recorder
}

// CreateSecret mocks base method
func (m *MockSecretEventHandler) CreateSecret(obj *v1.Secret) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSecret", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSecret indicates an expected call of CreateSecret
func (mr *MockSecretEventHandlerMockRecorder) CreateSecret(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecret", reflect.TypeOf((*MockSecretEventHandler)(nil).CreateSecret), obj)
}

// UpdateSecret mocks base method
func (m *MockSecretEventHandler) UpdateSecret(old, new *v1.Secret) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSecret", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSecret indicates an expected call of UpdateSecret
func (mr *MockSecretEventHandlerMockRecorder) UpdateSecret(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSecret", reflect.TypeOf((*MockSecretEventHandler)(nil).UpdateSecret), old, new)
}

// DeleteSecret mocks base method
func (m *MockSecretEventHandler) DeleteSecret(obj *v1.Secret) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecret", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSecret indicates an expected call of DeleteSecret
func (mr *MockSecretEventHandlerMockRecorder) DeleteSecret(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecret", reflect.TypeOf((*MockSecretEventHandler)(nil).DeleteSecret), obj)
}

// GenericSecret mocks base method
func (m *MockSecretEventHandler) GenericSecret(obj *v1.Secret) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericSecret", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericSecret indicates an expected call of GenericSecret
func (mr *MockSecretEventHandlerMockRecorder) GenericSecret(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericSecret", reflect.TypeOf((*MockSecretEventHandler)(nil).GenericSecret), obj)
}

// MockSecretEventWatcher is a mock of SecretEventWatcher interface
type MockSecretEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockSecretEventWatcherMockRecorder
}

// MockSecretEventWatcherMockRecorder is the mock recorder for MockSecretEventWatcher
type MockSecretEventWatcherMockRecorder struct {
	mock *MockSecretEventWatcher
}

// NewMockSecretEventWatcher creates a new mock instance
func NewMockSecretEventWatcher(ctrl *gomock.Controller) *MockSecretEventWatcher {
	mock := &MockSecretEventWatcher{ctrl: ctrl}
	mock.recorder = &MockSecretEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSecretEventWatcher) EXPECT() *MockSecretEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method
func (m *MockSecretEventWatcher) AddEventHandler(ctx context.Context, h controller.SecretEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler
func (mr *MockSecretEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockSecretEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockServiceAccountEventHandler is a mock of ServiceAccountEventHandler interface
type MockServiceAccountEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockServiceAccountEventHandlerMockRecorder
}

// MockServiceAccountEventHandlerMockRecorder is the mock recorder for MockServiceAccountEventHandler
type MockServiceAccountEventHandlerMockRecorder struct {
	mock *MockServiceAccountEventHandler
}

// NewMockServiceAccountEventHandler creates a new mock instance
func NewMockServiceAccountEventHandler(ctrl *gomock.Controller) *MockServiceAccountEventHandler {
	mock := &MockServiceAccountEventHandler{ctrl: ctrl}
	mock.recorder = &MockServiceAccountEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceAccountEventHandler) EXPECT() *MockServiceAccountEventHandlerMockRecorder {
	return m.recorder
}

// CreateServiceAccount mocks base method
func (m *MockServiceAccountEventHandler) CreateServiceAccount(obj *v1.ServiceAccount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateServiceAccount", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateServiceAccount indicates an expected call of CreateServiceAccount
func (mr *MockServiceAccountEventHandlerMockRecorder) CreateServiceAccount(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateServiceAccount", reflect.TypeOf((*MockServiceAccountEventHandler)(nil).CreateServiceAccount), obj)
}

// UpdateServiceAccount mocks base method
func (m *MockServiceAccountEventHandler) UpdateServiceAccount(old, new *v1.ServiceAccount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateServiceAccount", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateServiceAccount indicates an expected call of UpdateServiceAccount
func (mr *MockServiceAccountEventHandlerMockRecorder) UpdateServiceAccount(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateServiceAccount", reflect.TypeOf((*MockServiceAccountEventHandler)(nil).UpdateServiceAccount), old, new)
}

// DeleteServiceAccount mocks base method
func (m *MockServiceAccountEventHandler) DeleteServiceAccount(obj *v1.ServiceAccount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteServiceAccount", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteServiceAccount indicates an expected call of DeleteServiceAccount
func (mr *MockServiceAccountEventHandlerMockRecorder) DeleteServiceAccount(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteServiceAccount", reflect.TypeOf((*MockServiceAccountEventHandler)(nil).DeleteServiceAccount), obj)
}

// GenericServiceAccount mocks base method
func (m *MockServiceAccountEventHandler) GenericServiceAccount(obj *v1.ServiceAccount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericServiceAccount", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericServiceAccount indicates an expected call of GenericServiceAccount
func (mr *MockServiceAccountEventHandlerMockRecorder) GenericServiceAccount(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericServiceAccount", reflect.TypeOf((*MockServiceAccountEventHandler)(nil).GenericServiceAccount), obj)
}

// MockServiceAccountEventWatcher is a mock of ServiceAccountEventWatcher interface
type MockServiceAccountEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockServiceAccountEventWatcherMockRecorder
}

// MockServiceAccountEventWatcherMockRecorder is the mock recorder for MockServiceAccountEventWatcher
type MockServiceAccountEventWatcherMockRecorder struct {
	mock *MockServiceAccountEventWatcher
}

// NewMockServiceAccountEventWatcher creates a new mock instance
func NewMockServiceAccountEventWatcher(ctrl *gomock.Controller) *MockServiceAccountEventWatcher {
	mock := &MockServiceAccountEventWatcher{ctrl: ctrl}
	mock.recorder = &MockServiceAccountEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceAccountEventWatcher) EXPECT() *MockServiceAccountEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method
func (m *MockServiceAccountEventWatcher) AddEventHandler(ctx context.Context, h controller.ServiceAccountEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler
func (mr *MockServiceAccountEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockServiceAccountEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockNamespaceEventHandler is a mock of NamespaceEventHandler interface
type MockNamespaceEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockNamespaceEventHandlerMockRecorder
}

// MockNamespaceEventHandlerMockRecorder is the mock recorder for MockNamespaceEventHandler
type MockNamespaceEventHandlerMockRecorder struct {
	mock *MockNamespaceEventHandler
}

// NewMockNamespaceEventHandler creates a new mock instance
func NewMockNamespaceEventHandler(ctrl *gomock.Controller) *MockNamespaceEventHandler {
	mock := &MockNamespaceEventHandler{ctrl: ctrl}
	mock.recorder = &MockNamespaceEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNamespaceEventHandler) EXPECT() *MockNamespaceEventHandlerMockRecorder {
	return m.recorder
}

// CreateNamespace mocks base method
func (m *MockNamespaceEventHandler) CreateNamespace(obj *v1.Namespace) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNamespace", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNamespace indicates an expected call of CreateNamespace
func (mr *MockNamespaceEventHandlerMockRecorder) CreateNamespace(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNamespace", reflect.TypeOf((*MockNamespaceEventHandler)(nil).CreateNamespace), obj)
}

// UpdateNamespace mocks base method
func (m *MockNamespaceEventHandler) UpdateNamespace(old, new *v1.Namespace) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNamespace", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateNamespace indicates an expected call of UpdateNamespace
func (mr *MockNamespaceEventHandlerMockRecorder) UpdateNamespace(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNamespace", reflect.TypeOf((*MockNamespaceEventHandler)(nil).UpdateNamespace), old, new)
}

// DeleteNamespace mocks base method
func (m *MockNamespaceEventHandler) DeleteNamespace(obj *v1.Namespace) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNamespace", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNamespace indicates an expected call of DeleteNamespace
func (mr *MockNamespaceEventHandlerMockRecorder) DeleteNamespace(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNamespace", reflect.TypeOf((*MockNamespaceEventHandler)(nil).DeleteNamespace), obj)
}

// GenericNamespace mocks base method
func (m *MockNamespaceEventHandler) GenericNamespace(obj *v1.Namespace) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericNamespace", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericNamespace indicates an expected call of GenericNamespace
func (mr *MockNamespaceEventHandlerMockRecorder) GenericNamespace(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericNamespace", reflect.TypeOf((*MockNamespaceEventHandler)(nil).GenericNamespace), obj)
}

// MockNamespaceEventWatcher is a mock of NamespaceEventWatcher interface
type MockNamespaceEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockNamespaceEventWatcherMockRecorder
}

// MockNamespaceEventWatcherMockRecorder is the mock recorder for MockNamespaceEventWatcher
type MockNamespaceEventWatcherMockRecorder struct {
	mock *MockNamespaceEventWatcher
}

// NewMockNamespaceEventWatcher creates a new mock instance
func NewMockNamespaceEventWatcher(ctrl *gomock.Controller) *MockNamespaceEventWatcher {
	mock := &MockNamespaceEventWatcher{ctrl: ctrl}
	mock.recorder = &MockNamespaceEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNamespaceEventWatcher) EXPECT() *MockNamespaceEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method
func (m *MockNamespaceEventWatcher) AddEventHandler(ctx context.Context, h controller.NamespaceEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler
func (mr *MockNamespaceEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockNamespaceEventWatcher)(nil).AddEventHandler), varargs...)
}
