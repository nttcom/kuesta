// Code generated by MockGen. DO NOT EDIT.
// Source: gitrepo.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gitrepo "github.com/nttcom/kuesta/internal/gitrepo"
)

// MockGitRepoClient is a mock of GitRepoClient interface.
type MockGitRepoClient struct {
	ctrl     *gomock.Controller
	recorder *MockGitRepoClientMockRecorder
}

// MockGitRepoClientMockRecorder is the mock recorder for MockGitRepoClient.
type MockGitRepoClientMockRecorder struct {
	mock *MockGitRepoClient
}

// NewMockGitRepoClient creates a new mock instance.
func NewMockGitRepoClient(ctrl *gomock.Controller) *MockGitRepoClient {
	mock := &MockGitRepoClient{ctrl: ctrl}
	mock.recorder = &MockGitRepoClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGitRepoClient) EXPECT() *MockGitRepoClientMockRecorder {
	return m.recorder
}

// CreatePullRequest mocks base method.
func (m *MockGitRepoClient) CreatePullRequest(ctx context.Context, payload gitrepo.GitPullRequestPayload) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePullRequest", ctx, payload)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePullRequest indicates an expected call of CreatePullRequest.
func (mr *MockGitRepoClientMockRecorder) CreatePullRequest(ctx, payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePullRequest", reflect.TypeOf((*MockGitRepoClient)(nil).CreatePullRequest), ctx, payload)
}

// HealthCheck mocks base method.
func (m *MockGitRepoClient) HealthCheck() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck")
	ret0, _ := ret[0].(error)
	return ret0
}

// HealthCheck indicates an expected call of HealthCheck.
func (mr *MockGitRepoClientMockRecorder) HealthCheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockGitRepoClient)(nil).HealthCheck))
}

// Kind mocks base method.
func (m *MockGitRepoClient) Kind() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Kind")
	ret0, _ := ret[0].(string)
	return ret0
}

// Kind indicates an expected call of Kind.
func (mr *MockGitRepoClientMockRecorder) Kind() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kind", reflect.TypeOf((*MockGitRepoClient)(nil).Kind))
}
