// Code generated by MockGen. DO NOT EDIT.
// Source: ./clients.go

// Package mock_v1beta1 is a generated GoMock package.
package mock_v1beta1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "github.com/solo-io/skv2/pkg/multicluster/internal/k8s/apiextensions.k8s.io/v1beta1"
	v1beta10 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockMulticlusterClientset is a mock of MulticlusterClientset interface
type MockMulticlusterClientset struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterClientsetMockRecorder
}

// MockMulticlusterClientsetMockRecorder is the mock recorder for MockMulticlusterClientset
type MockMulticlusterClientsetMockRecorder struct {
	mock *MockMulticlusterClientset
}

// NewMockMulticlusterClientset creates a new mock instance
func NewMockMulticlusterClientset(ctrl *gomock.Controller) *MockMulticlusterClientset {
	mock := &MockMulticlusterClientset{ctrl: ctrl}
	mock.recorder = &MockMulticlusterClientsetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterClientset) EXPECT() *MockMulticlusterClientsetMockRecorder {
	return m.recorder
}

// Cluster mocks base method
func (m *MockMulticlusterClientset) Cluster(cluster string) (v1beta1.Clientset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", cluster)
	ret0, _ := ret[0].(v1beta1.Clientset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cluster indicates an expected call of Cluster
func (mr *MockMulticlusterClientsetMockRecorder) Cluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockMulticlusterClientset)(nil).Cluster), cluster)
}

// MockClientset is a mock of Clientset interface
type MockClientset struct {
	ctrl     *gomock.Controller
	recorder *MockClientsetMockRecorder
}

// MockClientsetMockRecorder is the mock recorder for MockClientset
type MockClientsetMockRecorder struct {
	mock *MockClientset
}

// NewMockClientset creates a new mock instance
func NewMockClientset(ctrl *gomock.Controller) *MockClientset {
	mock := &MockClientset{ctrl: ctrl}
	mock.recorder = &MockClientsetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClientset) EXPECT() *MockClientsetMockRecorder {
	return m.recorder
}

// CustomResourceDefinitions mocks base method
func (m *MockClientset) CustomResourceDefinitions() v1beta1.CustomResourceDefinitionClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CustomResourceDefinitions")
	ret0, _ := ret[0].(v1beta1.CustomResourceDefinitionClient)
	return ret0
}

// CustomResourceDefinitions indicates an expected call of CustomResourceDefinitions
func (mr *MockClientsetMockRecorder) CustomResourceDefinitions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CustomResourceDefinitions", reflect.TypeOf((*MockClientset)(nil).CustomResourceDefinitions))
}

// MockCustomResourceDefinitionReader is a mock of CustomResourceDefinitionReader interface
type MockCustomResourceDefinitionReader struct {
	ctrl     *gomock.Controller
	recorder *MockCustomResourceDefinitionReaderMockRecorder
}

// MockCustomResourceDefinitionReaderMockRecorder is the mock recorder for MockCustomResourceDefinitionReader
type MockCustomResourceDefinitionReaderMockRecorder struct {
	mock *MockCustomResourceDefinitionReader
}

// NewMockCustomResourceDefinitionReader creates a new mock instance
func NewMockCustomResourceDefinitionReader(ctrl *gomock.Controller) *MockCustomResourceDefinitionReader {
	mock := &MockCustomResourceDefinitionReader{ctrl: ctrl}
	mock.recorder = &MockCustomResourceDefinitionReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCustomResourceDefinitionReader) EXPECT() *MockCustomResourceDefinitionReaderMockRecorder {
	return m.recorder
}

// GetCustomResourceDefinition mocks base method
func (m *MockCustomResourceDefinitionReader) GetCustomResourceDefinition(ctx context.Context, name string) (*v1beta10.CustomResourceDefinition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCustomResourceDefinition", ctx, name)
	ret0, _ := ret[0].(*v1beta10.CustomResourceDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCustomResourceDefinition indicates an expected call of GetCustomResourceDefinition
func (mr *MockCustomResourceDefinitionReaderMockRecorder) GetCustomResourceDefinition(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCustomResourceDefinition", reflect.TypeOf((*MockCustomResourceDefinitionReader)(nil).GetCustomResourceDefinition), ctx, name)
}

// ListCustomResourceDefinition mocks base method
func (m *MockCustomResourceDefinitionReader) ListCustomResourceDefinition(ctx context.Context, opts ...client.ListOption) (*v1beta10.CustomResourceDefinitionList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCustomResourceDefinition", varargs...)
	ret0, _ := ret[0].(*v1beta10.CustomResourceDefinitionList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCustomResourceDefinition indicates an expected call of ListCustomResourceDefinition
func (mr *MockCustomResourceDefinitionReaderMockRecorder) ListCustomResourceDefinition(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCustomResourceDefinition", reflect.TypeOf((*MockCustomResourceDefinitionReader)(nil).ListCustomResourceDefinition), varargs...)
}

// MockCustomResourceDefinitionWriter is a mock of CustomResourceDefinitionWriter interface
type MockCustomResourceDefinitionWriter struct {
	ctrl     *gomock.Controller
	recorder *MockCustomResourceDefinitionWriterMockRecorder
}

// MockCustomResourceDefinitionWriterMockRecorder is the mock recorder for MockCustomResourceDefinitionWriter
type MockCustomResourceDefinitionWriterMockRecorder struct {
	mock *MockCustomResourceDefinitionWriter
}

// NewMockCustomResourceDefinitionWriter creates a new mock instance
func NewMockCustomResourceDefinitionWriter(ctrl *gomock.Controller) *MockCustomResourceDefinitionWriter {
	mock := &MockCustomResourceDefinitionWriter{ctrl: ctrl}
	mock.recorder = &MockCustomResourceDefinitionWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCustomResourceDefinitionWriter) EXPECT() *MockCustomResourceDefinitionWriterMockRecorder {
	return m.recorder
}

// CreateCustomResourceDefinition mocks base method
func (m *MockCustomResourceDefinitionWriter) CreateCustomResourceDefinition(ctx context.Context, obj *v1beta10.CustomResourceDefinition, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateCustomResourceDefinition", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCustomResourceDefinition indicates an expected call of CreateCustomResourceDefinition
func (mr *MockCustomResourceDefinitionWriterMockRecorder) CreateCustomResourceDefinition(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCustomResourceDefinition", reflect.TypeOf((*MockCustomResourceDefinitionWriter)(nil).CreateCustomResourceDefinition), varargs...)
}

// DeleteCustomResourceDefinition mocks base method
func (m *MockCustomResourceDefinitionWriter) DeleteCustomResourceDefinition(ctx context.Context, name string, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteCustomResourceDefinition", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCustomResourceDefinition indicates an expected call of DeleteCustomResourceDefinition
func (mr *MockCustomResourceDefinitionWriterMockRecorder) DeleteCustomResourceDefinition(ctx, name interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCustomResourceDefinition", reflect.TypeOf((*MockCustomResourceDefinitionWriter)(nil).DeleteCustomResourceDefinition), varargs...)
}

// UpdateCustomResourceDefinition mocks base method
func (m *MockCustomResourceDefinitionWriter) UpdateCustomResourceDefinition(ctx context.Context, obj *v1beta10.CustomResourceDefinition, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateCustomResourceDefinition", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCustomResourceDefinition indicates an expected call of UpdateCustomResourceDefinition
func (mr *MockCustomResourceDefinitionWriterMockRecorder) UpdateCustomResourceDefinition(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCustomResourceDefinition", reflect.TypeOf((*MockCustomResourceDefinitionWriter)(nil).UpdateCustomResourceDefinition), varargs...)
}

// PatchCustomResourceDefinition mocks base method
func (m *MockCustomResourceDefinitionWriter) PatchCustomResourceDefinition(ctx context.Context, obj *v1beta10.CustomResourceDefinition, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchCustomResourceDefinition", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchCustomResourceDefinition indicates an expected call of PatchCustomResourceDefinition
func (mr *MockCustomResourceDefinitionWriterMockRecorder) PatchCustomResourceDefinition(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchCustomResourceDefinition", reflect.TypeOf((*MockCustomResourceDefinitionWriter)(nil).PatchCustomResourceDefinition), varargs...)
}

// DeleteAllOfCustomResourceDefinition mocks base method
func (m *MockCustomResourceDefinitionWriter) DeleteAllOfCustomResourceDefinition(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfCustomResourceDefinition", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfCustomResourceDefinition indicates an expected call of DeleteAllOfCustomResourceDefinition
func (mr *MockCustomResourceDefinitionWriterMockRecorder) DeleteAllOfCustomResourceDefinition(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfCustomResourceDefinition", reflect.TypeOf((*MockCustomResourceDefinitionWriter)(nil).DeleteAllOfCustomResourceDefinition), varargs...)
}

// UpsertCustomResourceDefinition mocks base method
func (m *MockCustomResourceDefinitionWriter) UpsertCustomResourceDefinition(ctx context.Context, obj *v1beta10.CustomResourceDefinition, transitionFuncs ...v1beta1.CustomResourceDefinitionTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertCustomResourceDefinition", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertCustomResourceDefinition indicates an expected call of UpsertCustomResourceDefinition
func (mr *MockCustomResourceDefinitionWriterMockRecorder) UpsertCustomResourceDefinition(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertCustomResourceDefinition", reflect.TypeOf((*MockCustomResourceDefinitionWriter)(nil).UpsertCustomResourceDefinition), varargs...)
}

// MockCustomResourceDefinitionStatusWriter is a mock of CustomResourceDefinitionStatusWriter interface
type MockCustomResourceDefinitionStatusWriter struct {
	ctrl     *gomock.Controller
	recorder *MockCustomResourceDefinitionStatusWriterMockRecorder
}

// MockCustomResourceDefinitionStatusWriterMockRecorder is the mock recorder for MockCustomResourceDefinitionStatusWriter
type MockCustomResourceDefinitionStatusWriterMockRecorder struct {
	mock *MockCustomResourceDefinitionStatusWriter
}

// NewMockCustomResourceDefinitionStatusWriter creates a new mock instance
func NewMockCustomResourceDefinitionStatusWriter(ctrl *gomock.Controller) *MockCustomResourceDefinitionStatusWriter {
	mock := &MockCustomResourceDefinitionStatusWriter{ctrl: ctrl}
	mock.recorder = &MockCustomResourceDefinitionStatusWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCustomResourceDefinitionStatusWriter) EXPECT() *MockCustomResourceDefinitionStatusWriterMockRecorder {
	return m.recorder
}

// UpdateCustomResourceDefinitionStatus mocks base method
func (m *MockCustomResourceDefinitionStatusWriter) UpdateCustomResourceDefinitionStatus(ctx context.Context, obj *v1beta10.CustomResourceDefinition, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateCustomResourceDefinitionStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCustomResourceDefinitionStatus indicates an expected call of UpdateCustomResourceDefinitionStatus
func (mr *MockCustomResourceDefinitionStatusWriterMockRecorder) UpdateCustomResourceDefinitionStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCustomResourceDefinitionStatus", reflect.TypeOf((*MockCustomResourceDefinitionStatusWriter)(nil).UpdateCustomResourceDefinitionStatus), varargs...)
}

// PatchCustomResourceDefinitionStatus mocks base method
func (m *MockCustomResourceDefinitionStatusWriter) PatchCustomResourceDefinitionStatus(ctx context.Context, obj *v1beta10.CustomResourceDefinition, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchCustomResourceDefinitionStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchCustomResourceDefinitionStatus indicates an expected call of PatchCustomResourceDefinitionStatus
func (mr *MockCustomResourceDefinitionStatusWriterMockRecorder) PatchCustomResourceDefinitionStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchCustomResourceDefinitionStatus", reflect.TypeOf((*MockCustomResourceDefinitionStatusWriter)(nil).PatchCustomResourceDefinitionStatus), varargs...)
}

// MockCustomResourceDefinitionClient is a mock of CustomResourceDefinitionClient interface
type MockCustomResourceDefinitionClient struct {
	ctrl     *gomock.Controller
	recorder *MockCustomResourceDefinitionClientMockRecorder
}

// MockCustomResourceDefinitionClientMockRecorder is the mock recorder for MockCustomResourceDefinitionClient
type MockCustomResourceDefinitionClientMockRecorder struct {
	mock *MockCustomResourceDefinitionClient
}

// NewMockCustomResourceDefinitionClient creates a new mock instance
func NewMockCustomResourceDefinitionClient(ctrl *gomock.Controller) *MockCustomResourceDefinitionClient {
	mock := &MockCustomResourceDefinitionClient{ctrl: ctrl}
	mock.recorder = &MockCustomResourceDefinitionClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCustomResourceDefinitionClient) EXPECT() *MockCustomResourceDefinitionClientMockRecorder {
	return m.recorder
}

// GetCustomResourceDefinition mocks base method
func (m *MockCustomResourceDefinitionClient) GetCustomResourceDefinition(ctx context.Context, name string) (*v1beta10.CustomResourceDefinition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCustomResourceDefinition", ctx, name)
	ret0, _ := ret[0].(*v1beta10.CustomResourceDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCustomResourceDefinition indicates an expected call of GetCustomResourceDefinition
func (mr *MockCustomResourceDefinitionClientMockRecorder) GetCustomResourceDefinition(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCustomResourceDefinition", reflect.TypeOf((*MockCustomResourceDefinitionClient)(nil).GetCustomResourceDefinition), ctx, name)
}

// ListCustomResourceDefinition mocks base method
func (m *MockCustomResourceDefinitionClient) ListCustomResourceDefinition(ctx context.Context, opts ...client.ListOption) (*v1beta10.CustomResourceDefinitionList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCustomResourceDefinition", varargs...)
	ret0, _ := ret[0].(*v1beta10.CustomResourceDefinitionList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCustomResourceDefinition indicates an expected call of ListCustomResourceDefinition
func (mr *MockCustomResourceDefinitionClientMockRecorder) ListCustomResourceDefinition(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCustomResourceDefinition", reflect.TypeOf((*MockCustomResourceDefinitionClient)(nil).ListCustomResourceDefinition), varargs...)
}

// CreateCustomResourceDefinition mocks base method
func (m *MockCustomResourceDefinitionClient) CreateCustomResourceDefinition(ctx context.Context, obj *v1beta10.CustomResourceDefinition, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateCustomResourceDefinition", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCustomResourceDefinition indicates an expected call of CreateCustomResourceDefinition
func (mr *MockCustomResourceDefinitionClientMockRecorder) CreateCustomResourceDefinition(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCustomResourceDefinition", reflect.TypeOf((*MockCustomResourceDefinitionClient)(nil).CreateCustomResourceDefinition), varargs...)
}

// DeleteCustomResourceDefinition mocks base method
func (m *MockCustomResourceDefinitionClient) DeleteCustomResourceDefinition(ctx context.Context, name string, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteCustomResourceDefinition", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCustomResourceDefinition indicates an expected call of DeleteCustomResourceDefinition
func (mr *MockCustomResourceDefinitionClientMockRecorder) DeleteCustomResourceDefinition(ctx, name interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCustomResourceDefinition", reflect.TypeOf((*MockCustomResourceDefinitionClient)(nil).DeleteCustomResourceDefinition), varargs...)
}

// UpdateCustomResourceDefinition mocks base method
func (m *MockCustomResourceDefinitionClient) UpdateCustomResourceDefinition(ctx context.Context, obj *v1beta10.CustomResourceDefinition, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateCustomResourceDefinition", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCustomResourceDefinition indicates an expected call of UpdateCustomResourceDefinition
func (mr *MockCustomResourceDefinitionClientMockRecorder) UpdateCustomResourceDefinition(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCustomResourceDefinition", reflect.TypeOf((*MockCustomResourceDefinitionClient)(nil).UpdateCustomResourceDefinition), varargs...)
}

// PatchCustomResourceDefinition mocks base method
func (m *MockCustomResourceDefinitionClient) PatchCustomResourceDefinition(ctx context.Context, obj *v1beta10.CustomResourceDefinition, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchCustomResourceDefinition", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchCustomResourceDefinition indicates an expected call of PatchCustomResourceDefinition
func (mr *MockCustomResourceDefinitionClientMockRecorder) PatchCustomResourceDefinition(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchCustomResourceDefinition", reflect.TypeOf((*MockCustomResourceDefinitionClient)(nil).PatchCustomResourceDefinition), varargs...)
}

// DeleteAllOfCustomResourceDefinition mocks base method
func (m *MockCustomResourceDefinitionClient) DeleteAllOfCustomResourceDefinition(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfCustomResourceDefinition", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfCustomResourceDefinition indicates an expected call of DeleteAllOfCustomResourceDefinition
func (mr *MockCustomResourceDefinitionClientMockRecorder) DeleteAllOfCustomResourceDefinition(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfCustomResourceDefinition", reflect.TypeOf((*MockCustomResourceDefinitionClient)(nil).DeleteAllOfCustomResourceDefinition), varargs...)
}

// UpsertCustomResourceDefinition mocks base method
func (m *MockCustomResourceDefinitionClient) UpsertCustomResourceDefinition(ctx context.Context, obj *v1beta10.CustomResourceDefinition, transitionFuncs ...v1beta1.CustomResourceDefinitionTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertCustomResourceDefinition", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertCustomResourceDefinition indicates an expected call of UpsertCustomResourceDefinition
func (mr *MockCustomResourceDefinitionClientMockRecorder) UpsertCustomResourceDefinition(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertCustomResourceDefinition", reflect.TypeOf((*MockCustomResourceDefinitionClient)(nil).UpsertCustomResourceDefinition), varargs...)
}

// UpdateCustomResourceDefinitionStatus mocks base method
func (m *MockCustomResourceDefinitionClient) UpdateCustomResourceDefinitionStatus(ctx context.Context, obj *v1beta10.CustomResourceDefinition, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateCustomResourceDefinitionStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCustomResourceDefinitionStatus indicates an expected call of UpdateCustomResourceDefinitionStatus
func (mr *MockCustomResourceDefinitionClientMockRecorder) UpdateCustomResourceDefinitionStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCustomResourceDefinitionStatus", reflect.TypeOf((*MockCustomResourceDefinitionClient)(nil).UpdateCustomResourceDefinitionStatus), varargs...)
}

// PatchCustomResourceDefinitionStatus mocks base method
func (m *MockCustomResourceDefinitionClient) PatchCustomResourceDefinitionStatus(ctx context.Context, obj *v1beta10.CustomResourceDefinition, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchCustomResourceDefinitionStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchCustomResourceDefinitionStatus indicates an expected call of PatchCustomResourceDefinitionStatus
func (mr *MockCustomResourceDefinitionClientMockRecorder) PatchCustomResourceDefinitionStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchCustomResourceDefinitionStatus", reflect.TypeOf((*MockCustomResourceDefinitionClient)(nil).PatchCustomResourceDefinitionStatus), varargs...)
}

// MockMulticlusterCustomResourceDefinitionClient is a mock of MulticlusterCustomResourceDefinitionClient interface
type MockMulticlusterCustomResourceDefinitionClient struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterCustomResourceDefinitionClientMockRecorder
}

// MockMulticlusterCustomResourceDefinitionClientMockRecorder is the mock recorder for MockMulticlusterCustomResourceDefinitionClient
type MockMulticlusterCustomResourceDefinitionClientMockRecorder struct {
	mock *MockMulticlusterCustomResourceDefinitionClient
}

// NewMockMulticlusterCustomResourceDefinitionClient creates a new mock instance
func NewMockMulticlusterCustomResourceDefinitionClient(ctrl *gomock.Controller) *MockMulticlusterCustomResourceDefinitionClient {
	mock := &MockMulticlusterCustomResourceDefinitionClient{ctrl: ctrl}
	mock.recorder = &MockMulticlusterCustomResourceDefinitionClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterCustomResourceDefinitionClient) EXPECT() *MockMulticlusterCustomResourceDefinitionClientMockRecorder {
	return m.recorder
}

// Cluster mocks base method
func (m *MockMulticlusterCustomResourceDefinitionClient) Cluster(cluster string) (v1beta1.CustomResourceDefinitionClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", cluster)
	ret0, _ := ret[0].(v1beta1.CustomResourceDefinitionClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cluster indicates an expected call of Cluster
func (mr *MockMulticlusterCustomResourceDefinitionClientMockRecorder) Cluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockMulticlusterCustomResourceDefinitionClient)(nil).Cluster), cluster)
}
