// Code generated by MockGen. DO NOT EDIT.
// Source: ./sets.go

// Package mock_v1sets is a generated GoMock package.
package mock_v1sets

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sets "github.com/solo-io/skv2/contrib/pkg/sets"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
	v1sets "github.com/solo-io/skv2/pkg/multicluster/internal/k8s/core/v1/sets"
	v1 "k8s.io/api/core/v1"
	sets0 "k8s.io/apimachinery/pkg/util/sets"
)

// MockSecretSet is a mock of SecretSet interface.
type MockSecretSet struct {
	ctrl     *gomock.Controller
	recorder *MockSecretSetMockRecorder
}

// MockSecretSetMockRecorder is the mock recorder for MockSecretSet.
type MockSecretSetMockRecorder struct {
	mock *MockSecretSet
}

// NewMockSecretSet creates a new mock instance.
func NewMockSecretSet(ctrl *gomock.Controller) *MockSecretSet {
	mock := &MockSecretSet{ctrl: ctrl}
	mock.recorder = &MockSecretSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretSet) EXPECT() *MockSecretSetMockRecorder {
	return m.recorder
}

// Keys mocks base method.
func (m *MockSecretSet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockSecretSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockSecretSet)(nil).Keys))
}

// List mocks base method.
func (m *MockSecretSet) List(filterResource ...func(*v1.Secret) bool) []*v1.Secret {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1.Secret)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockSecretSetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSecretSet)(nil).List), filterResource...)
}

// Map mocks base method.
func (m *MockSecretSet) Map() map[string]*v1.Secret {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1.Secret)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockSecretSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockSecretSet)(nil).Map))
}

// Insert mocks base method.
func (m *MockSecretSet) Insert(secret ...*v1.Secret) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range secret {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockSecretSetMockRecorder) Insert(secret ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockSecretSet)(nil).Insert), secret...)
}

// Equal mocks base method.
func (m *MockSecretSet) Equal(secretSet v1sets.SecretSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", secretSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockSecretSetMockRecorder) Equal(secretSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockSecretSet)(nil).Equal), secretSet)
}

// Has mocks base method.
func (m *MockSecretSet) Has(secret ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", secret)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockSecretSetMockRecorder) Has(secret interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockSecretSet)(nil).Has), secret)
}

// Delete mocks base method.
func (m *MockSecretSet) Delete(secret ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", secret)
}

// Delete indicates an expected call of Delete.
func (mr *MockSecretSetMockRecorder) Delete(secret interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSecretSet)(nil).Delete), secret)
}

// Union mocks base method.
func (m *MockSecretSet) Union(set v1sets.SecretSet) v1sets.SecretSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1sets.SecretSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockSecretSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockSecretSet)(nil).Union), set)
}

// Difference mocks base method.
func (m *MockSecretSet) Difference(set v1sets.SecretSet) v1sets.SecretSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1sets.SecretSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockSecretSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockSecretSet)(nil).Difference), set)
}

// Intersection mocks base method.
func (m *MockSecretSet) Intersection(set v1sets.SecretSet) v1sets.SecretSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1sets.SecretSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockSecretSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockSecretSet)(nil).Intersection), set)
}

// Find mocks base method.
func (m *MockSecretSet) Find(id ezkube.ResourceId) (*v1.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockSecretSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockSecretSet)(nil).Find), id)
}

// Length mocks base method.
func (m *MockSecretSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockSecretSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockSecretSet)(nil).Length))
}

// Generic mocks base method.
func (m *MockSecretSet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic.
func (mr *MockSecretSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockSecretSet)(nil).Generic))
}

// Delta mocks base method.
func (m *MockSecretSet) Delta(newSet v1sets.SecretSet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta.
func (mr *MockSecretSetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockSecretSet)(nil).Delta), newSet)
}

// MockServiceAccountSet is a mock of ServiceAccountSet interface.
type MockServiceAccountSet struct {
	ctrl     *gomock.Controller
	recorder *MockServiceAccountSetMockRecorder
}

// MockServiceAccountSetMockRecorder is the mock recorder for MockServiceAccountSet.
type MockServiceAccountSetMockRecorder struct {
	mock *MockServiceAccountSet
}

// NewMockServiceAccountSet creates a new mock instance.
func NewMockServiceAccountSet(ctrl *gomock.Controller) *MockServiceAccountSet {
	mock := &MockServiceAccountSet{ctrl: ctrl}
	mock.recorder = &MockServiceAccountSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceAccountSet) EXPECT() *MockServiceAccountSetMockRecorder {
	return m.recorder
}

// Keys mocks base method.
func (m *MockServiceAccountSet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockServiceAccountSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockServiceAccountSet)(nil).Keys))
}

// List mocks base method.
func (m *MockServiceAccountSet) List(filterResource ...func(*v1.ServiceAccount) bool) []*v1.ServiceAccount {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1.ServiceAccount)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockServiceAccountSetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockServiceAccountSet)(nil).List), filterResource...)
}

// Map mocks base method.
func (m *MockServiceAccountSet) Map() map[string]*v1.ServiceAccount {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1.ServiceAccount)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockServiceAccountSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockServiceAccountSet)(nil).Map))
}

// Insert mocks base method.
func (m *MockServiceAccountSet) Insert(serviceAccount ...*v1.ServiceAccount) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range serviceAccount {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockServiceAccountSetMockRecorder) Insert(serviceAccount ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockServiceAccountSet)(nil).Insert), serviceAccount...)
}

// Equal mocks base method.
func (m *MockServiceAccountSet) Equal(serviceAccountSet v1sets.ServiceAccountSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", serviceAccountSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockServiceAccountSetMockRecorder) Equal(serviceAccountSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockServiceAccountSet)(nil).Equal), serviceAccountSet)
}

// Has mocks base method.
func (m *MockServiceAccountSet) Has(serviceAccount ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", serviceAccount)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockServiceAccountSetMockRecorder) Has(serviceAccount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockServiceAccountSet)(nil).Has), serviceAccount)
}

// Delete mocks base method.
func (m *MockServiceAccountSet) Delete(serviceAccount ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", serviceAccount)
}

// Delete indicates an expected call of Delete.
func (mr *MockServiceAccountSetMockRecorder) Delete(serviceAccount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockServiceAccountSet)(nil).Delete), serviceAccount)
}

// Union mocks base method.
func (m *MockServiceAccountSet) Union(set v1sets.ServiceAccountSet) v1sets.ServiceAccountSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1sets.ServiceAccountSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockServiceAccountSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockServiceAccountSet)(nil).Union), set)
}

// Difference mocks base method.
func (m *MockServiceAccountSet) Difference(set v1sets.ServiceAccountSet) v1sets.ServiceAccountSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1sets.ServiceAccountSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockServiceAccountSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockServiceAccountSet)(nil).Difference), set)
}

// Intersection mocks base method.
func (m *MockServiceAccountSet) Intersection(set v1sets.ServiceAccountSet) v1sets.ServiceAccountSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1sets.ServiceAccountSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockServiceAccountSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockServiceAccountSet)(nil).Intersection), set)
}

// Find mocks base method.
func (m *MockServiceAccountSet) Find(id ezkube.ResourceId) (*v1.ServiceAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1.ServiceAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockServiceAccountSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockServiceAccountSet)(nil).Find), id)
}

// Length mocks base method.
func (m *MockServiceAccountSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockServiceAccountSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockServiceAccountSet)(nil).Length))
}

// Generic mocks base method.
func (m *MockServiceAccountSet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic.
func (mr *MockServiceAccountSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockServiceAccountSet)(nil).Generic))
}

// Delta mocks base method.
func (m *MockServiceAccountSet) Delta(newSet v1sets.ServiceAccountSet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta.
func (mr *MockServiceAccountSetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockServiceAccountSet)(nil).Delta), newSet)
}

// MockNamespaceSet is a mock of NamespaceSet interface.
type MockNamespaceSet struct {
	ctrl     *gomock.Controller
	recorder *MockNamespaceSetMockRecorder
}

// MockNamespaceSetMockRecorder is the mock recorder for MockNamespaceSet.
type MockNamespaceSetMockRecorder struct {
	mock *MockNamespaceSet
}

// NewMockNamespaceSet creates a new mock instance.
func NewMockNamespaceSet(ctrl *gomock.Controller) *MockNamespaceSet {
	mock := &MockNamespaceSet{ctrl: ctrl}
	mock.recorder = &MockNamespaceSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNamespaceSet) EXPECT() *MockNamespaceSetMockRecorder {
	return m.recorder
}

// Keys mocks base method.
func (m *MockNamespaceSet) Keys() sets0.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets0.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockNamespaceSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockNamespaceSet)(nil).Keys))
}

// List mocks base method.
func (m *MockNamespaceSet) List(filterResource ...func(*v1.Namespace) bool) []*v1.Namespace {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1.Namespace)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockNamespaceSetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockNamespaceSet)(nil).List), filterResource...)
}

// Map mocks base method.
func (m *MockNamespaceSet) Map() map[string]*v1.Namespace {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1.Namespace)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockNamespaceSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockNamespaceSet)(nil).Map))
}

// Insert mocks base method.
func (m *MockNamespaceSet) Insert(namespace ...*v1.Namespace) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range namespace {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockNamespaceSetMockRecorder) Insert(namespace ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockNamespaceSet)(nil).Insert), namespace...)
}

// Equal mocks base method.
func (m *MockNamespaceSet) Equal(namespaceSet v1sets.NamespaceSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", namespaceSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockNamespaceSetMockRecorder) Equal(namespaceSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockNamespaceSet)(nil).Equal), namespaceSet)
}

// Has mocks base method.
func (m *MockNamespaceSet) Has(namespace ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", namespace)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockNamespaceSetMockRecorder) Has(namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockNamespaceSet)(nil).Has), namespace)
}

// Delete mocks base method.
func (m *MockNamespaceSet) Delete(namespace ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", namespace)
}

// Delete indicates an expected call of Delete.
func (mr *MockNamespaceSetMockRecorder) Delete(namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockNamespaceSet)(nil).Delete), namespace)
}

// Union mocks base method.
func (m *MockNamespaceSet) Union(set v1sets.NamespaceSet) v1sets.NamespaceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1sets.NamespaceSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockNamespaceSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockNamespaceSet)(nil).Union), set)
}

// Difference mocks base method.
func (m *MockNamespaceSet) Difference(set v1sets.NamespaceSet) v1sets.NamespaceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1sets.NamespaceSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockNamespaceSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockNamespaceSet)(nil).Difference), set)
}

// Intersection mocks base method.
func (m *MockNamespaceSet) Intersection(set v1sets.NamespaceSet) v1sets.NamespaceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1sets.NamespaceSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockNamespaceSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockNamespaceSet)(nil).Intersection), set)
}

// Find mocks base method.
func (m *MockNamespaceSet) Find(id ezkube.ResourceId) (*v1.Namespace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1.Namespace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockNamespaceSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockNamespaceSet)(nil).Find), id)
}

// Length mocks base method.
func (m *MockNamespaceSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockNamespaceSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockNamespaceSet)(nil).Length))
}

// Generic mocks base method.
func (m *MockNamespaceSet) Generic() sets.ResourceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(sets.ResourceSet)
	return ret0
}

// Generic indicates an expected call of Generic.
func (mr *MockNamespaceSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockNamespaceSet)(nil).Generic))
}

// Delta mocks base method.
func (m *MockNamespaceSet) Delta(newSet v1sets.NamespaceSet) sets.ResourceDelta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delta", newSet)
	ret0, _ := ret[0].(sets.ResourceDelta)
	return ret0
}

// Delta indicates an expected call of Delta.
func (mr *MockNamespaceSetMockRecorder) Delta(newSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delta", reflect.TypeOf((*MockNamespaceSet)(nil).Delta), newSet)
}
