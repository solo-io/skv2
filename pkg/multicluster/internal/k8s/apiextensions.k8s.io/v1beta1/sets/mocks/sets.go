// Code generated by MockGen. DO NOT EDIT.
// Source: ./sets.go

// Package mock_v1beta1sets is a generated GoMock package.
package mock_v1beta1sets

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
	v1beta1sets "github.com/solo-io/skv2/pkg/multicluster/internal/k8s/apiextensions.k8s.io/v1beta1/sets"
	v1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	sets "k8s.io/apimachinery/pkg/util/sets"
)

// MockCustomResourceDefinitionSet is a mock of CustomResourceDefinitionSet interface.
type MockCustomResourceDefinitionSet struct {
	ctrl     *gomock.Controller
	recorder *MockCustomResourceDefinitionSetMockRecorder
}

// MockCustomResourceDefinitionSetMockRecorder is the mock recorder for MockCustomResourceDefinitionSet.
type MockCustomResourceDefinitionSetMockRecorder struct {
	mock *MockCustomResourceDefinitionSet
}

// NewMockCustomResourceDefinitionSet creates a new mock instance.
func NewMockCustomResourceDefinitionSet(ctrl *gomock.Controller) *MockCustomResourceDefinitionSet {
	mock := &MockCustomResourceDefinitionSet{ctrl: ctrl}
	mock.recorder = &MockCustomResourceDefinitionSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomResourceDefinitionSet) EXPECT() *MockCustomResourceDefinitionSetMockRecorder {
	return m.recorder
}

// Keys mocks base method.
func (m *MockCustomResourceDefinitionSet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockCustomResourceDefinitionSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockCustomResourceDefinitionSet)(nil).Keys))
}

// List mocks base method.
func (m *MockCustomResourceDefinitionSet) List() []*v1beta1.CustomResourceDefinition {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*v1beta1.CustomResourceDefinition)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockCustomResourceDefinitionSetMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCustomResourceDefinitionSet)(nil).List))
}

// Map mocks base method.
func (m *MockCustomResourceDefinitionSet) Map() map[string]*v1beta1.CustomResourceDefinition {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1beta1.CustomResourceDefinition)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockCustomResourceDefinitionSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockCustomResourceDefinitionSet)(nil).Map))
}

// Insert mocks base method.
func (m *MockCustomResourceDefinitionSet) Insert(customResourceDefinition ...*v1beta1.CustomResourceDefinition) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range customResourceDefinition {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockCustomResourceDefinitionSetMockRecorder) Insert(customResourceDefinition ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockCustomResourceDefinitionSet)(nil).Insert), customResourceDefinition...)
}

// Equal mocks base method.
func (m *MockCustomResourceDefinitionSet) Equal(customResourceDefinitionSet v1beta1sets.CustomResourceDefinitionSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", customResourceDefinitionSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockCustomResourceDefinitionSetMockRecorder) Equal(customResourceDefinitionSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockCustomResourceDefinitionSet)(nil).Equal), customResourceDefinitionSet)
}

// Has mocks base method.
func (m *MockCustomResourceDefinitionSet) Has(customResourceDefinition *v1beta1.CustomResourceDefinition) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", customResourceDefinition)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockCustomResourceDefinitionSetMockRecorder) Has(customResourceDefinition interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockCustomResourceDefinitionSet)(nil).Has), customResourceDefinition)
}

// Delete mocks base method.
func (m *MockCustomResourceDefinitionSet) Delete(customResourceDefinition *v1beta1.CustomResourceDefinition) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", customResourceDefinition)
}

// Delete indicates an expected call of Delete.
func (mr *MockCustomResourceDefinitionSetMockRecorder) Delete(customResourceDefinition interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCustomResourceDefinitionSet)(nil).Delete), customResourceDefinition)
}

// Union mocks base method.
func (m *MockCustomResourceDefinitionSet) Union(set v1beta1sets.CustomResourceDefinitionSet) v1beta1sets.CustomResourceDefinitionSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1beta1sets.CustomResourceDefinitionSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockCustomResourceDefinitionSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockCustomResourceDefinitionSet)(nil).Union), set)
}

// Difference mocks base method.
func (m *MockCustomResourceDefinitionSet) Difference(set v1beta1sets.CustomResourceDefinitionSet) v1beta1sets.CustomResourceDefinitionSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1beta1sets.CustomResourceDefinitionSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockCustomResourceDefinitionSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockCustomResourceDefinitionSet)(nil).Difference), set)
}

// Intersection mocks base method.
func (m *MockCustomResourceDefinitionSet) Intersection(set v1beta1sets.CustomResourceDefinitionSet) v1beta1sets.CustomResourceDefinitionSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1beta1sets.CustomResourceDefinitionSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockCustomResourceDefinitionSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockCustomResourceDefinitionSet)(nil).Intersection), set)
}

// Find mocks base method.
func (m *MockCustomResourceDefinitionSet) Find(id ezkube.ResourceId) (*v1beta1.CustomResourceDefinition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1beta1.CustomResourceDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockCustomResourceDefinitionSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockCustomResourceDefinitionSet)(nil).Find), id)
}

// Length mocks base method.
func (m *MockCustomResourceDefinitionSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockCustomResourceDefinitionSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockCustomResourceDefinitionSet)(nil).Length))
}
