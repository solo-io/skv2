// Code generated by MockGen. DO NOT EDIT.
// Source: ./sets.go

// Package mock_v1sets is a generated GoMock package.
package mock_v1sets

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1"
	v1sets "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1/sets"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
	sets "k8s.io/apimachinery/pkg/util/sets"
)

// MockPaintSet is a mock of PaintSet interface.
type MockPaintSet struct {
	ctrl     *gomock.Controller
	recorder *MockPaintSetMockRecorder
}

// MockPaintSetMockRecorder is the mock recorder for MockPaintSet.
type MockPaintSetMockRecorder struct {
	mock *MockPaintSet
}

// NewMockPaintSet creates a new mock instance.
func NewMockPaintSet(ctrl *gomock.Controller) *MockPaintSet {
	mock := &MockPaintSet{ctrl: ctrl}
	mock.recorder = &MockPaintSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaintSet) EXPECT() *MockPaintSetMockRecorder {
	return m.recorder
}

// Keys mocks base method.
func (m *MockPaintSet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockPaintSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockPaintSet)(nil).Keys))
}

// List mocks base method.
func (m *MockPaintSet) List(filterResource ...func(*v1.Paint) bool) []*v1.Paint {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1.Paint)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockPaintSetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPaintSet)(nil).List), filterResource...)
}

// Map mocks base method.
func (m *MockPaintSet) Map() map[string]*v1.Paint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1.Paint)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockPaintSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockPaintSet)(nil).Map))
}

// Insert mocks base method.
func (m *MockPaintSet) Insert(paint ...*v1.Paint) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range paint {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockPaintSetMockRecorder) Insert(paint ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockPaintSet)(nil).Insert), paint...)
}

// Equal mocks base method.
func (m *MockPaintSet) Equal(paintSet v1sets.PaintSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", paintSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockPaintSetMockRecorder) Equal(paintSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockPaintSet)(nil).Equal), paintSet)
}

// Has mocks base method.
func (m *MockPaintSet) Has(paint ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", paint)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockPaintSetMockRecorder) Has(paint interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockPaintSet)(nil).Has), paint)
}

// Delete mocks base method.
func (m *MockPaintSet) Delete(paint ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", paint)
}

// Delete indicates an expected call of Delete.
func (mr *MockPaintSetMockRecorder) Delete(paint interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPaintSet)(nil).Delete), paint)
}

// Union mocks base method.
func (m *MockPaintSet) Union(set v1sets.PaintSet) v1sets.PaintSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1sets.PaintSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockPaintSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockPaintSet)(nil).Union), set)
}

// Difference mocks base method.
func (m *MockPaintSet) Difference(set v1sets.PaintSet) v1sets.PaintSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1sets.PaintSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockPaintSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockPaintSet)(nil).Difference), set)
}

// Intersection mocks base method.
func (m *MockPaintSet) Intersection(set v1sets.PaintSet) v1sets.PaintSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1sets.PaintSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockPaintSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockPaintSet)(nil).Intersection), set)
}

// Find mocks base method.
func (m *MockPaintSet) Find(id ezkube.ResourceId) (*v1.Paint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1.Paint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockPaintSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockPaintSet)(nil).Find), id)
}

// Length mocks base method.
func (m *MockPaintSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockPaintSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockPaintSet)(nil).Length))
}

// MockClusterResourceSet is a mock of ClusterResourceSet interface.
type MockClusterResourceSet struct {
	ctrl     *gomock.Controller
	recorder *MockClusterResourceSetMockRecorder
}

// MockClusterResourceSetMockRecorder is the mock recorder for MockClusterResourceSet.
type MockClusterResourceSetMockRecorder struct {
	mock *MockClusterResourceSet
}

// NewMockClusterResourceSet creates a new mock instance.
func NewMockClusterResourceSet(ctrl *gomock.Controller) *MockClusterResourceSet {
	mock := &MockClusterResourceSet{ctrl: ctrl}
	mock.recorder = &MockClusterResourceSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterResourceSet) EXPECT() *MockClusterResourceSetMockRecorder {
	return m.recorder
}

// Keys mocks base method.
func (m *MockClusterResourceSet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockClusterResourceSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockClusterResourceSet)(nil).Keys))
}

// List mocks base method.
func (m *MockClusterResourceSet) List(filterResource ...func(*v1.ClusterResource) bool) []*v1.ClusterResource {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1.ClusterResource)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockClusterResourceSetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockClusterResourceSet)(nil).List), filterResource...)
}

// Map mocks base method.
func (m *MockClusterResourceSet) Map() map[string]*v1.ClusterResource {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1.ClusterResource)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockClusterResourceSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockClusterResourceSet)(nil).Map))
}

// Insert mocks base method.
func (m *MockClusterResourceSet) Insert(clusterResource ...*v1.ClusterResource) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range clusterResource {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockClusterResourceSetMockRecorder) Insert(clusterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockClusterResourceSet)(nil).Insert), clusterResource...)
}

// Equal mocks base method.
func (m *MockClusterResourceSet) Equal(clusterResourceSet v1sets.ClusterResourceSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", clusterResourceSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockClusterResourceSetMockRecorder) Equal(clusterResourceSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockClusterResourceSet)(nil).Equal), clusterResourceSet)
}

// Has mocks base method.
func (m *MockClusterResourceSet) Has(clusterResource ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", clusterResource)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockClusterResourceSetMockRecorder) Has(clusterResource interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockClusterResourceSet)(nil).Has), clusterResource)
}

// Delete mocks base method.
func (m *MockClusterResourceSet) Delete(clusterResource ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", clusterResource)
}

// Delete indicates an expected call of Delete.
func (mr *MockClusterResourceSetMockRecorder) Delete(clusterResource interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClusterResourceSet)(nil).Delete), clusterResource)
}

// Union mocks base method.
func (m *MockClusterResourceSet) Union(set v1sets.ClusterResourceSet) v1sets.ClusterResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1sets.ClusterResourceSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockClusterResourceSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockClusterResourceSet)(nil).Union), set)
}

// Difference mocks base method.
func (m *MockClusterResourceSet) Difference(set v1sets.ClusterResourceSet) v1sets.ClusterResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1sets.ClusterResourceSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockClusterResourceSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockClusterResourceSet)(nil).Difference), set)
}

// Intersection mocks base method.
func (m *MockClusterResourceSet) Intersection(set v1sets.ClusterResourceSet) v1sets.ClusterResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1sets.ClusterResourceSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockClusterResourceSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockClusterResourceSet)(nil).Intersection), set)
}

// Find mocks base method.
func (m *MockClusterResourceSet) Find(id ezkube.ResourceId) (*v1.ClusterResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1.ClusterResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockClusterResourceSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockClusterResourceSet)(nil).Find), id)
}

// Length mocks base method.
func (m *MockClusterResourceSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockClusterResourceSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockClusterResourceSet)(nil).Length))
}
