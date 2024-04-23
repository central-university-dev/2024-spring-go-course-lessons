// Code generated by MockGen. DO NOT EDIT.
// Source: gotestexample/internal/app/command (interfaces: Repository,IDProvider)
//
// Generated by this command:
//
//	mockgen -destination=../../../mocks/create_short_url_cmd.go -package=mocks -mock_names=Repository=MockCmdRepo,IDProvider=StubIDProvider gotestexample/internal/app/command Repository,IDProvider
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCmdRepo is a mock of Repository interface.
type MockCmdRepo struct {
	ctrl     *gomock.Controller
	recorder *MockCmdRepoMockRecorder
}

// MockCmdRepoMockRecorder is the mock recorder for MockCmdRepo.
type MockCmdRepoMockRecorder struct {
	mock *MockCmdRepo
}

// NewMockCmdRepo creates a new mock instance.
func NewMockCmdRepo(ctrl *gomock.Controller) *MockCmdRepo {
	mock := &MockCmdRepo{ctrl: ctrl}
	mock.recorder = &MockCmdRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCmdRepo) EXPECT() *MockCmdRepoMockRecorder {
	return m.recorder
}

// Save mocks base method.
func (m *MockCmdRepo) Save(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockCmdRepoMockRecorder) Save(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockCmdRepo)(nil).Save), arg0, arg1, arg2)
}

// StubIDProvider is a mock of IDProvider interface.
type StubIDProvider struct {
	ctrl     *gomock.Controller
	recorder *StubIDProviderMockRecorder
}

// StubIDProviderMockRecorder is the mock recorder for StubIDProvider.
type StubIDProviderMockRecorder struct {
	mock *StubIDProvider
}

// NewStubIDProvider creates a new mock instance.
func NewStubIDProvider(ctrl *gomock.Controller) *StubIDProvider {
	mock := &StubIDProvider{ctrl: ctrl}
	mock.recorder = &StubIDProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *StubIDProvider) EXPECT() *StubIDProviderMockRecorder {
	return m.recorder
}

// Provide mocks base method.
func (m *StubIDProvider) Provide() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Provide")
	ret0, _ := ret[0].(string)
	return ret0
}

// Provide indicates an expected call of Provide.
func (mr *StubIDProviderMockRecorder) Provide() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Provide", reflect.TypeOf((*StubIDProvider)(nil).Provide))
}
