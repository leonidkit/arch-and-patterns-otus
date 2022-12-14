// Code generated by MockGen. DO NOT EDIT.
// Source: hw5.go

// Package hw5 is a generated GoMock package.
package hw5

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockICommand is a mock of ICommand interface.
type MockICommand struct {
	ctrl     *gomock.Controller
	recorder *MockICommandMockRecorder
}

// MockICommandMockRecorder is the mock recorder for MockICommand.
type MockICommandMockRecorder struct {
	mock *MockICommand
}

// NewMockICommand creates a new mock instance.
func NewMockICommand(ctrl *gomock.Controller) *MockICommand {
	mock := &MockICommand{ctrl: ctrl}
	mock.recorder = &MockICommandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICommand) EXPECT() *MockICommandMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockICommand) Execute() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute")
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute.
func (mr *MockICommandMockRecorder) Execute() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockICommand)(nil).Execute))
}
