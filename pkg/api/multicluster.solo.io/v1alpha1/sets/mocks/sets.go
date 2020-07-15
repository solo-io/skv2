// Code generated by MockGen. DO NOT EDIT.
// Source: ./sets.go

// Package mock_v1alpha1sets is a generated GoMock package.
package mock_v1alpha1sets

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1"
	v1alpha1sets "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1/sets"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
	sets "k8s.io/apimachinery/pkg/util/sets"
)

// MockKubernetesClusterSet is a mock of KubernetesClusterSet interface.
type MockKubernetesClusterSet struct {
	ctrl     *gomock.Controller
	recorder *MockKubernetesClusterSetMockRecorder
}

// MockKubernetesClusterSetMockRecorder is the mock recorder for MockKubernetesClusterSet.
type MockKubernetesClusterSetMockRecorder struct {
	mock *MockKubernetesClusterSet
}

// NewMockKubernetesClusterSet creates a new mock instance.
func NewMockKubernetesClusterSet(ctrl *gomock.Controller) *MockKubernetesClusterSet {
	mock := &MockKubernetesClusterSet{ctrl: ctrl}
	mock.recorder = &MockKubernetesClusterSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKubernetesClusterSet) EXPECT() *MockKubernetesClusterSetMockRecorder {
	return m.recorder
}

// Keys mocks base method.
func (m *MockKubernetesClusterSet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockKubernetesClusterSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockKubernetesClusterSet)(nil).Keys))
}

// List mocks base method.
func (m *MockKubernetesClusterSet) List() []*v1alpha1.KubernetesCluster {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*v1alpha1.KubernetesCluster)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockKubernetesClusterSetMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockKubernetesClusterSet)(nil).List))
}

// Map mocks base method.
func (m *MockKubernetesClusterSet) Map() map[string]*v1alpha1.KubernetesCluster {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1alpha1.KubernetesCluster)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockKubernetesClusterSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockKubernetesClusterSet)(nil).Map))
}

// Insert mocks base method.
func (m *MockKubernetesClusterSet) Insert(kubernetesCluster ...*v1alpha1.KubernetesCluster) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range kubernetesCluster {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockKubernetesClusterSetMockRecorder) Insert(kubernetesCluster ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockKubernetesClusterSet)(nil).Insert), kubernetesCluster...)
}

// Equal mocks base method.
func (m *MockKubernetesClusterSet) Equal(kubernetesClusterSet v1alpha1sets.KubernetesClusterSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", kubernetesClusterSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockKubernetesClusterSetMockRecorder) Equal(kubernetesClusterSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockKubernetesClusterSet)(nil).Equal), kubernetesClusterSet)
}

// Has mocks base method.
func (m *MockKubernetesClusterSet) Has(kubernetesCluster *v1alpha1.KubernetesCluster) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", kubernetesCluster)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockKubernetesClusterSetMockRecorder) Has(kubernetesCluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockKubernetesClusterSet)(nil).Has), kubernetesCluster)
}

// Delete mocks base method.
func (m *MockKubernetesClusterSet) Delete(kubernetesCluster *v1alpha1.KubernetesCluster) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", kubernetesCluster)
}

// Delete indicates an expected call of Delete.
func (mr *MockKubernetesClusterSetMockRecorder) Delete(kubernetesCluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockKubernetesClusterSet)(nil).Delete), kubernetesCluster)
}

// Union mocks base method.
func (m *MockKubernetesClusterSet) Union(set v1alpha1sets.KubernetesClusterSet) v1alpha1sets.KubernetesClusterSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1alpha1sets.KubernetesClusterSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockKubernetesClusterSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockKubernetesClusterSet)(nil).Union), set)
}

// Difference mocks base method.
func (m *MockKubernetesClusterSet) Difference(set v1alpha1sets.KubernetesClusterSet) v1alpha1sets.KubernetesClusterSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1alpha1sets.KubernetesClusterSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockKubernetesClusterSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockKubernetesClusterSet)(nil).Difference), set)
}

// Intersection mocks base method.
func (m *MockKubernetesClusterSet) Intersection(set v1alpha1sets.KubernetesClusterSet) v1alpha1sets.KubernetesClusterSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1alpha1sets.KubernetesClusterSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockKubernetesClusterSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockKubernetesClusterSet)(nil).Intersection), set)
}

// Find mocks base method.
func (m *MockKubernetesClusterSet) Find(id ezkube.ResourceId) (*v1alpha1.KubernetesCluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1alpha1.KubernetesCluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockKubernetesClusterSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockKubernetesClusterSet)(nil).Find), id)
}

// Length mocks base method.
func (m *MockKubernetesClusterSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockKubernetesClusterSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockKubernetesClusterSet)(nil).Length))
}
