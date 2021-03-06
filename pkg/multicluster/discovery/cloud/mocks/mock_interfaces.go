// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_cloud is a generated GoMock package.
package mock_cloud

import (
	context "context"
	reflect "reflect"

	eks "github.com/aws/aws-sdk-go/service/eks"
	gomock "github.com/golang/mock/gomock"
	oauth2 "golang.org/x/oauth2"
	container "google.golang.org/genproto/googleapis/container/v1"
	token "sigs.k8s.io/aws-iam-authenticator/pkg/token"
)

// MockEksClient is a mock of EksClient interface
type MockEksClient struct {
	ctrl     *gomock.Controller
	recorder *MockEksClientMockRecorder
}

// MockEksClientMockRecorder is the mock recorder for MockEksClient
type MockEksClientMockRecorder struct {
	mock *MockEksClient
}

// NewMockEksClient creates a new mock instance
func NewMockEksClient(ctrl *gomock.Controller) *MockEksClient {
	mock := &MockEksClient{ctrl: ctrl}
	mock.recorder = &MockEksClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEksClient) EXPECT() *MockEksClientMockRecorder {
	return m.recorder
}

// DescribeCluster mocks base method
func (m *MockEksClient) DescribeCluster(ctx context.Context, name string) (*eks.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCluster", ctx, name)
	ret0, _ := ret[0].(*eks.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCluster indicates an expected call of DescribeCluster
func (mr *MockEksClientMockRecorder) DescribeCluster(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCluster", reflect.TypeOf((*MockEksClient)(nil).DescribeCluster), ctx, name)
}

// ListClusters mocks base method
func (m *MockEksClient) ListClusters(ctx context.Context, fn func(*eks.ListClustersOutput)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListClusters", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListClusters indicates an expected call of ListClusters
func (mr *MockEksClientMockRecorder) ListClusters(ctx, fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClusters", reflect.TypeOf((*MockEksClient)(nil).ListClusters), ctx, fn)
}

// Token mocks base method
func (m *MockEksClient) Token(ctx context.Context, name string) (token.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Token", ctx, name)
	ret0, _ := ret[0].(token.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Token indicates an expected call of Token
func (mr *MockEksClientMockRecorder) Token(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Token", reflect.TypeOf((*MockEksClient)(nil).Token), ctx, name)
}

// MockGkeClient is a mock of GkeClient interface
type MockGkeClient struct {
	ctrl     *gomock.Controller
	recorder *MockGkeClientMockRecorder
}

// MockGkeClientMockRecorder is the mock recorder for MockGkeClient
type MockGkeClientMockRecorder struct {
	mock *MockGkeClient
}

// NewMockGkeClient creates a new mock instance
func NewMockGkeClient(ctrl *gomock.Controller) *MockGkeClient {
	mock := &MockGkeClient{ctrl: ctrl}
	mock.recorder = &MockGkeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGkeClient) EXPECT() *MockGkeClientMockRecorder {
	return m.recorder
}

// Token mocks base method
func (m *MockGkeClient) Token(ctx context.Context) (*oauth2.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Token", ctx)
	ret0, _ := ret[0].(*oauth2.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Token indicates an expected call of Token
func (mr *MockGkeClientMockRecorder) Token(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Token", reflect.TypeOf((*MockGkeClient)(nil).Token), ctx)
}

// ListClusters mocks base method
func (m *MockGkeClient) ListClusters(ctx context.Context) ([]*container.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListClusters", ctx)
	ret0, _ := ret[0].([]*container.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClusters indicates an expected call of ListClusters
func (mr *MockGkeClientMockRecorder) ListClusters(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClusters", reflect.TypeOf((*MockGkeClient)(nil).ListClusters), ctx)
}
