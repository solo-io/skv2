// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_auth is a generated GoMock package.
package mock_auth

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/rbac/v1"
	rest "k8s.io/client-go/rest"
)

// MockRemoteAuthorityConfigCreator is a mock of RemoteAuthorityConfigCreator interface.
type MockRemoteAuthorityConfigCreator struct {
	ctrl     *gomock.Controller
	recorder *MockRemoteAuthorityConfigCreatorMockRecorder
}

// MockRemoteAuthorityConfigCreatorMockRecorder is the mock recorder for MockRemoteAuthorityConfigCreator.
type MockRemoteAuthorityConfigCreatorMockRecorder struct {
	mock *MockRemoteAuthorityConfigCreator
}

// NewMockRemoteAuthorityConfigCreator creates a new mock instance.
func NewMockRemoteAuthorityConfigCreator(ctrl *gomock.Controller) *MockRemoteAuthorityConfigCreator {
	mock := &MockRemoteAuthorityConfigCreator{ctrl: ctrl}
	mock.recorder = &MockRemoteAuthorityConfigCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRemoteAuthorityConfigCreator) EXPECT() *MockRemoteAuthorityConfigCreatorMockRecorder {
	return m.recorder
}

// ConfigFromRemoteServiceAccount mocks base method.
func (m *MockRemoteAuthorityConfigCreator) ConfigFromRemoteServiceAccount(ctx context.Context, targetClusterCfg *rest.Config, name, namespace string) (*rest.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigFromRemoteServiceAccount", ctx, targetClusterCfg, name, namespace)
	ret0, _ := ret[0].(*rest.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigFromRemoteServiceAccount indicates an expected call of ConfigFromRemoteServiceAccount.
func (mr *MockRemoteAuthorityConfigCreatorMockRecorder) ConfigFromRemoteServiceAccount(ctx, targetClusterCfg, name, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigFromRemoteServiceAccount", reflect.TypeOf((*MockRemoteAuthorityConfigCreator)(nil).ConfigFromRemoteServiceAccount), ctx, targetClusterCfg, name, namespace)
}

// MockClusterAuthorization is a mock of ClusterAuthorization interface.
type MockClusterAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockClusterAuthorizationMockRecorder
}

// MockClusterAuthorizationMockRecorder is the mock recorder for MockClusterAuthorization.
type MockClusterAuthorizationMockRecorder struct {
	mock *MockClusterAuthorization
}

// NewMockClusterAuthorization creates a new mock instance.
func NewMockClusterAuthorization(ctrl *gomock.Controller) *MockClusterAuthorization {
	mock := &MockClusterAuthorization{ctrl: ctrl}
	mock.recorder = &MockClusterAuthorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterAuthorization) EXPECT() *MockClusterAuthorizationMockRecorder {
	return m.recorder
}

// BuildClusterScopedRemoteBearerToken mocks base method.
func (m *MockClusterAuthorization) BuildClusterScopedRemoteBearerToken(ctx context.Context, targetClusterCfg *rest.Config, name, namespace string, clusterRoles ...*v1.ClusterRole) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, targetClusterCfg, name, namespace}
	for _, a := range clusterRoles {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BuildClusterScopedRemoteBearerToken", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildClusterScopedRemoteBearerToken indicates an expected call of BuildClusterScopedRemoteBearerToken.
func (mr *MockClusterAuthorizationMockRecorder) BuildClusterScopedRemoteBearerToken(ctx, targetClusterCfg, name, namespace interface{}, clusterRoles ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, targetClusterCfg, name, namespace}, clusterRoles...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildClusterScopedRemoteBearerToken", reflect.TypeOf((*MockClusterAuthorization)(nil).BuildClusterScopedRemoteBearerToken), varargs...)
}

// BuildRemoteBearerToken mocks base method.
func (m *MockClusterAuthorization) BuildRemoteBearerToken(ctx context.Context, targetClusterCfg *rest.Config, name, namespace string, roles []*v1.Role) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildRemoteBearerToken", ctx, targetClusterCfg, name, namespace, roles)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildRemoteBearerToken indicates an expected call of BuildRemoteBearerToken.
func (mr *MockClusterAuthorizationMockRecorder) BuildRemoteBearerToken(ctx, targetClusterCfg, name, namespace, roles interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildRemoteBearerToken", reflect.TypeOf((*MockClusterAuthorization)(nil).BuildRemoteBearerToken), ctx, targetClusterCfg, name, namespace, roles)
}
