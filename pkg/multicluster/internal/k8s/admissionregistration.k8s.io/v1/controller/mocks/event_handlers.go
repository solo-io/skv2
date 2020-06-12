// Code generated by MockGen. DO NOT EDIT.
// Source: ./event_handlers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	controller "github.com/solo-io/skv2/pkg/multicluster/internal/k8s/admissionregistration.k8s.io/v1/controller"
	v1 "k8s.io/api/admissionregistration/v1"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockValidatingWebhookConfigurationEventHandler is a mock of ValidatingWebhookConfigurationEventHandler interface.
type MockValidatingWebhookConfigurationEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockValidatingWebhookConfigurationEventHandlerMockRecorder
}

// MockValidatingWebhookConfigurationEventHandlerMockRecorder is the mock recorder for MockValidatingWebhookConfigurationEventHandler.
type MockValidatingWebhookConfigurationEventHandlerMockRecorder struct {
	mock *MockValidatingWebhookConfigurationEventHandler
}

// NewMockValidatingWebhookConfigurationEventHandler creates a new mock instance.
func NewMockValidatingWebhookConfigurationEventHandler(ctrl *gomock.Controller) *MockValidatingWebhookConfigurationEventHandler {
	mock := &MockValidatingWebhookConfigurationEventHandler{ctrl: ctrl}
	mock.recorder = &MockValidatingWebhookConfigurationEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidatingWebhookConfigurationEventHandler) EXPECT() *MockValidatingWebhookConfigurationEventHandlerMockRecorder {
	return m.recorder
}

// CreateValidatingWebhookConfiguration mocks base method.
func (m *MockValidatingWebhookConfigurationEventHandler) CreateValidatingWebhookConfiguration(obj *v1.ValidatingWebhookConfiguration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateValidatingWebhookConfiguration", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateValidatingWebhookConfiguration indicates an expected call of CreateValidatingWebhookConfiguration.
func (mr *MockValidatingWebhookConfigurationEventHandlerMockRecorder) CreateValidatingWebhookConfiguration(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateValidatingWebhookConfiguration", reflect.TypeOf((*MockValidatingWebhookConfigurationEventHandler)(nil).CreateValidatingWebhookConfiguration), obj)
}

// UpdateValidatingWebhookConfiguration mocks base method.
func (m *MockValidatingWebhookConfigurationEventHandler) UpdateValidatingWebhookConfiguration(old, new *v1.ValidatingWebhookConfiguration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateValidatingWebhookConfiguration", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateValidatingWebhookConfiguration indicates an expected call of UpdateValidatingWebhookConfiguration.
func (mr *MockValidatingWebhookConfigurationEventHandlerMockRecorder) UpdateValidatingWebhookConfiguration(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateValidatingWebhookConfiguration", reflect.TypeOf((*MockValidatingWebhookConfigurationEventHandler)(nil).UpdateValidatingWebhookConfiguration), old, new)
}

// DeleteValidatingWebhookConfiguration mocks base method.
func (m *MockValidatingWebhookConfigurationEventHandler) DeleteValidatingWebhookConfiguration(obj *v1.ValidatingWebhookConfiguration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteValidatingWebhookConfiguration", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteValidatingWebhookConfiguration indicates an expected call of DeleteValidatingWebhookConfiguration.
func (mr *MockValidatingWebhookConfigurationEventHandlerMockRecorder) DeleteValidatingWebhookConfiguration(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteValidatingWebhookConfiguration", reflect.TypeOf((*MockValidatingWebhookConfigurationEventHandler)(nil).DeleteValidatingWebhookConfiguration), obj)
}

// GenericValidatingWebhookConfiguration mocks base method.
func (m *MockValidatingWebhookConfigurationEventHandler) GenericValidatingWebhookConfiguration(obj *v1.ValidatingWebhookConfiguration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericValidatingWebhookConfiguration", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericValidatingWebhookConfiguration indicates an expected call of GenericValidatingWebhookConfiguration.
func (mr *MockValidatingWebhookConfigurationEventHandlerMockRecorder) GenericValidatingWebhookConfiguration(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericValidatingWebhookConfiguration", reflect.TypeOf((*MockValidatingWebhookConfigurationEventHandler)(nil).GenericValidatingWebhookConfiguration), obj)
}

// MockValidatingWebhookConfigurationEventWatcher is a mock of ValidatingWebhookConfigurationEventWatcher interface.
type MockValidatingWebhookConfigurationEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockValidatingWebhookConfigurationEventWatcherMockRecorder
}

// MockValidatingWebhookConfigurationEventWatcherMockRecorder is the mock recorder for MockValidatingWebhookConfigurationEventWatcher.
type MockValidatingWebhookConfigurationEventWatcherMockRecorder struct {
	mock *MockValidatingWebhookConfigurationEventWatcher
}

// NewMockValidatingWebhookConfigurationEventWatcher creates a new mock instance.
func NewMockValidatingWebhookConfigurationEventWatcher(ctrl *gomock.Controller) *MockValidatingWebhookConfigurationEventWatcher {
	mock := &MockValidatingWebhookConfigurationEventWatcher{ctrl: ctrl}
	mock.recorder = &MockValidatingWebhookConfigurationEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidatingWebhookConfigurationEventWatcher) EXPECT() *MockValidatingWebhookConfigurationEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockValidatingWebhookConfigurationEventWatcher) AddEventHandler(ctx context.Context, h controller.ValidatingWebhookConfigurationEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler.
func (mr *MockValidatingWebhookConfigurationEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockValidatingWebhookConfigurationEventWatcher)(nil).AddEventHandler), varargs...)
}
