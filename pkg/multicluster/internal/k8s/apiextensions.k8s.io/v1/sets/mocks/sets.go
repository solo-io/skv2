// Code generated by MockGen. DO NOT EDIT.
// Source: ./sets.go

// Package mock_v1sets is a generated GoMock package.
package mock_v1sets

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sets "github.com/solo-io/skv2/contrib/pkg/sets"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
	v1sets "github.com/solo-io/skv2/pkg/multicluster/internal/k8s/apiextensions.k8s.io/v1/sets"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	sets0 "k8s.io/apimachinery/pkg/util/sets"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockCustomResourceDefinitionSet is a mock of CustomResourceDefinitionSet interface
type MockCustomResourceDefinitionSet struct {
	ctrl     *gomock.Controller
	recorder *MockCustomResourceDefinitionSetMockRecorder
}

// MockCustomResourceDefinitionSetMockRecorder is the mock recorder for MockCustomResourceDefinitionSet
type MockCustomResourceDefinitionSetMockRecorder struct {
	mock *MockCustomResourceDefinitionSet
}

// NewMockCustomResourceDefinitionSet creates a new mock instance
func NewMockCustomResourceDefinitionSet(ctrl *gomock.Controller) *MockCustomResourceDefinitionSet {
	mock := &MockCustomResourceDefinitionSet{ctrl: ctrl}
	mock.recorder = &MockCustomResourceDefinitionSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCustomResourceDefinitionSet) EXPECT() *MockCustomResourceDefinitionSetMockRecorder {
	return m.recorder
}

// Keys mocks base method
func (m *MockCustomResourceDefinitionSet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys
func (mr *MockCustomResourceDefinitionSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockCustomResourceDefinitionSet)(nil).Keys))
}

// List mocks base method
func (m *MockCustomResourceDefinitionSet) List(filterResource ...func(*v1.CustomResourceDefinition) bool) []*v1.CustomResourceDefinition {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1.CustomResourceDefinition)
	return ret0
}

// List indicates an expected call of List
func (mr *MockCustomResourceDefinitionSetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCustomResourceDefinitionSet)(nil).List), filterResource...)
}

// UnsortedList mocks base method
func (m *MockCustomResourceDefinitionSet) UnsortedList(filterResource ...func(*v1.CustomResourceDefinition) bool) []*v1.CustomResourceDefinition {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnsortedList", varargs...)
	ret0, _ := ret[0].([]*v1.CustomResourceDefinition)
	return ret0
}

// UnsortedList indicates an expected call of UnsortedList
func (mr *MockCustomResourceDefinitionSetMockRecorder) UnsortedList(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsortedList", reflect.TypeOf((*MockCustomResourceDefinitionSet)(nil).UnsortedList), filterResource...)
}

// Map mocks base method
func (m *MockCustomResourceDefinitionSet) Map() map[string]*v1.CustomResourceDefinition {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1.CustomResourceDefinition)
	return ret0
}

// Map indicates an expected call of Map
func (mr *MockCustomResourceDefinitionSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockCustomResourceDefinitionSet)(nil).Map))
}

// Insert mocks base method
func (m *MockCustomResourceDefinitionSet) Insert(customResourceDefinition ...*v1.CustomResourceDefinition) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range customResourceDefinition {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert
func (mr *MockCustomResourceDefinitionSetMockRecorder) Insert(customResourceDefinition ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockCustomResourceDefinitionSet)(nil).Insert), customResourceDefinition...)
}

// Equal mocks base method
func (m *MockCustomResourceDefinitionSet) Equal(customResourceDefinitionSet v1sets.CustomResourceDefinitionSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", customResourceDefinitionSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (mr *MockCustomResourceDefinitionSetMockRecorder) Equal(customResourceDefinitionSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockCustomResourceDefinitionSet)(nil).Equal), customResourceDefinitionSet)
}

// Has mocks base method
func (m *MockCustomResourceDefinitionSet) Has(customResourceDefinition ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", customResourceDefinition)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockCustomResourceDefinitionSetMockRecorder) Has(customResourceDefinition interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockCustomResourceDefinitionSet)(nil).Has), customResourceDefinition)
}

// Delete mocks base method
func (m *MockCustomResourceDefinitionSet) Delete(customResourceDefinition ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", customResourceDefinition)
}

// Delete indicates an expected call of Delete
func (mr *MockCustomResourceDefinitionSetMockRecorder) Delete(customResourceDefinition interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCustomResourceDefinitionSet)(nil).Delete), customResourceDefinition)
}

// Union mocks base method
func (m *MockCustomResourceDefinitionSet) Union(set v1sets.CustomResourceDefinitionSet) v1sets.CustomResourceDefinitionSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1sets.CustomResourceDefinitionSet)
	return ret0
}

// Union indicates an expected call of Union
func (mr *MockCustomResourceDefinitionSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockCustomResourceDefinitionSet)(nil).Union), set)
}

// Difference mocks base method
func (m *MockCustomResourceDefinitionSet) Difference(set v1sets.CustomResourceDefinitionSet) v1sets.CustomResourceDefinitionSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1sets.CustomResourceDefinitionSet)
	return ret0
}

// Difference indicates an expected call of Difference
func (mr *MockCustomResourceDefinitionSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockCustomResourceDefinitionSet)(nil).Difference), set)
}

// Intersection mocks base method
func (m *MockCustomResourceDefinitionSet) Intersection(set v1sets.CustomResourceDefinitionSet) v1sets.CustomResourceDefinitionSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1sets.CustomResourceDefinitionSet)
	return ret0
}

// Intersection indicates an expected call of Intersection
func (mr *MockCustomResourceDefinitionSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockCustomResourceDefinitionSet)(nil).Intersection), set)
}

// Find mocks base method
func (m *MockCustomResourceDefinitionSet) Find(id ezkube.ResourceId) (*v1.CustomResourceDefinition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1.CustomResourceDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockCustomResourceDefinitionSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockCustomResourceDefinitionSet)(nil).Find), id)
}

// Length mocks base method
func (m *MockCustomResourceDefinitionSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length
func (mr *MockCustomResourceDefinitionSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockCustomResourceDefinitionSet)(nil).Length))
}

// Generic mocks base method
func (m *MockCustomResourceDefinitionSet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic
func (mr *MockCustomResourceDefinitionSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockCustomResourceDefinitionSet)(nil).Generic))
}

// Delta mocks base method
func (m *MockCustomResourceDefinitionSet) Delta(newSet v1sets.CustomResourceDefinitionSet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta
func (mr *MockCustomResourceDefinitionSetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockCustomResourceDefinitionSet)(nil).Delta), newSet)
}

// Clone mocks base method
func (m *MockCustomResourceDefinitionSet) Clone() v1sets.CustomResourceDefinitionSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(v1sets.CustomResourceDefinitionSet)
	return ret0
}

// Clone indicates an expected call of Clone
func (mr *MockCustomResourceDefinitionSetMockRecorder) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockCustomResourceDefinitionSet)(nil).Clone))
}

// GetSortFunc mocks base method
func (m *MockCustomResourceDefinitionSet) GetSortFunc() func(client.Object, client.Object) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSortFunc")
	ret0, _ := ret[0].(func(client.Object, client.Object) bool)
	return ret0
}

// GetSortFunc indicates an expected call of GetSortFunc
func (mr *MockCustomResourceDefinitionSetMockRecorder) GetSortFunc() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSortFunc", reflect.TypeOf((*MockCustomResourceDefinitionSet)(nil).GetSortFunc))
}

// GetEqualityFunc mocks base method
func (m *MockCustomResourceDefinitionSet) GetEqualityFunc() func(client.Object, client.Object) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEqualityFunc")
	ret0, _ := ret[0].(func(client.Object, client.Object) bool)
	return ret0
}

// GetEqualityFunc indicates an expected call of GetEqualityFunc
func (mr *MockCustomResourceDefinitionSetMockRecorder) GetEqualityFunc() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEqualityFunc", reflect.TypeOf((*MockCustomResourceDefinitionSet)(nil).GetEqualityFunc))
}
