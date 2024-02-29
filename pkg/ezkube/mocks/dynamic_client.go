// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/skv2/pkg/ezkube (interfaces: Ensurer)

// Package mock_ezkube is a generated GoMock package.
package mock_ezkube

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
	reflect "reflect"
	client "sigs.k8s.io/controller-runtime/pkg/client"
	manager "sigs.k8s.io/controller-runtime/pkg/manager"
)

// MockEnsurer is a mock of Ensurer interface
type MockEnsurer struct {
	ctrl     *gomock.Controller
	recorder *MockEnsurerMockRecorder
}

// MockEnsurerMockRecorder is the mock recorder for MockEnsurer
type MockEnsurerMockRecorder struct {
	mock *MockEnsurer
}

// NewMockEnsurer creates a new mock instance
func NewMockEnsurer(ctrl *gomock.Controller) *MockEnsurer {
	mock := &MockEnsurer{ctrl: ctrl}
	mock.recorder = &MockEnsurerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEnsurer) EXPECT() *MockEnsurerMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockEnsurer) Create(arg0 context.Context, arg1 ezkube.Object) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockEnsurerMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockEnsurer)(nil).Create), arg0, arg1)
}

// Delete mocks base method
func (m *MockEnsurer) Delete(arg0 context.Context, arg1 ezkube.Object) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockEnsurerMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockEnsurer)(nil).Delete), arg0, arg1)
}

// Ensure mocks base method
func (m *MockEnsurer) Ensure(arg0 context.Context, arg1, arg2 ezkube.Object, arg3 ...ezkube.ReconcileFunc) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Ensure", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ensure indicates an expected call of Ensure
func (mr *MockEnsurerMockRecorder) Ensure(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ensure", reflect.TypeOf((*MockEnsurer)(nil).Ensure), varargs...)
}

// Get mocks base method
func (m *MockEnsurer) Get(arg0 context.Context, arg1 ezkube.Object) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockEnsurerMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockEnsurer)(nil).Get), arg0, arg1)
}

// List mocks base method
func (m *MockEnsurer) List(arg0 context.Context, arg1 ezkube.List, arg2 ...client.ListOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// List indicates an expected call of List
func (mr *MockEnsurerMockRecorder) List(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockEnsurer)(nil).List), varargs...)
}

// Manager mocks base method
func (m *MockEnsurer) Manager() manager.Manager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Manager")
	ret0, _ := ret[0].(manager.Manager)
	return ret0
}

// Manager indicates an expected call of Manager
func (mr *MockEnsurerMockRecorder) Manager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Manager", reflect.TypeOf((*MockEnsurer)(nil).Manager))
}

// Update mocks base method
func (m *MockEnsurer) Update(arg0 context.Context, arg1 ezkube.Object, arg2 ...ezkube.ReconcileFunc) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockEnsurerMockRecorder) Update(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockEnsurer)(nil).Update), varargs...)
}

// UpdateStatus mocks base method
func (m *MockEnsurer) UpdateStatus(arg0 context.Context, arg1 ezkube.Object) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatus indicates an expected call of UpdateStatus
func (mr *MockEnsurerMockRecorder) UpdateStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockEnsurer)(nil).UpdateStatus), arg0, arg1)
}
