// Code generated by MockGen. DO NOT EDIT.
// Source: hw3.go

// Package hw3 is a generated GoMock package.
package hw3

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

// Name mocks base method.
func (m *MockICommand) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockICommandMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockICommand)(nil).Name))
}

// MockQueue is a mock of Queue interface.
type MockQueue struct {
	ctrl     *gomock.Controller
	recorder *MockQueueMockRecorder
}

// MockQueueMockRecorder is the mock recorder for MockQueue.
type MockQueueMockRecorder struct {
	mock *MockQueue
}

// NewMockQueue creates a new mock instance.
func NewMockQueue(ctrl *gomock.Controller) *MockQueue {
	mock := &MockQueue{ctrl: ctrl}
	mock.recorder = &MockQueueMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueue) EXPECT() *MockQueueMockRecorder {
	return m.recorder
}

// IsEmpty mocks base method.
func (m *MockQueue) IsEmpty() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEmpty")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsEmpty indicates an expected call of IsEmpty.
func (mr *MockQueueMockRecorder) IsEmpty() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEmpty", reflect.TypeOf((*MockQueue)(nil).IsEmpty))
}

// Pop mocks base method.
func (m *MockQueue) Pop() (ICommand, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pop")
	ret0, _ := ret[0].(ICommand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Pop indicates an expected call of Pop.
func (mr *MockQueueMockRecorder) Pop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pop", reflect.TypeOf((*MockQueue)(nil).Pop))
}

// Push mocks base method.
func (m *MockQueue) Push(arg0 ICommand) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Push", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Push indicates an expected call of Push.
func (mr *MockQueueMockRecorder) Push(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockQueue)(nil).Push), arg0)
}