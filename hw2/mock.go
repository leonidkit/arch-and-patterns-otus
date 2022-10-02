// Code generated by MockGen. DO NOT EDIT.
// Source: hw2.go

// Package hw2 is a generated GoMock package.
package hw2

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRotable is a mock of Rotable interface.
type MockRotable struct {
	ctrl     *gomock.Controller
	recorder *MockRotableMockRecorder
}

// MockRotableMockRecorder is the mock recorder for MockRotable.
type MockRotableMockRecorder struct {
	mock *MockRotable
}

// NewMockRotable creates a new mock instance.
func NewMockRotable(ctrl *gomock.Controller) *MockRotable {
	mock := &MockRotable{ctrl: ctrl}
	mock.recorder = &MockRotableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRotable) EXPECT() *MockRotableMockRecorder {
	return m.recorder
}

// AngurlarVelocity mocks base method.
func (m *MockRotable) AngurlarVelocity() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AngurlarVelocity")
	ret0, _ := ret[0].(int64)
	return ret0
}

// AngurlarVelocity indicates an expected call of AngurlarVelocity.
func (mr *MockRotableMockRecorder) AngurlarVelocity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AngurlarVelocity", reflect.TypeOf((*MockRotable)(nil).AngurlarVelocity))
}

// Direction mocks base method.
func (m *MockRotable) Direction() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Direction")
	ret0, _ := ret[0].(int64)
	return ret0
}

// Direction indicates an expected call of Direction.
func (mr *MockRotableMockRecorder) Direction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Direction", reflect.TypeOf((*MockRotable)(nil).Direction))
}

// DirectionNumber mocks base method.
func (m *MockRotable) DirectionNumber() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DirectionNumber")
	ret0, _ := ret[0].(int64)
	return ret0
}

// DirectionNumber indicates an expected call of DirectionNumber.
func (mr *MockRotableMockRecorder) DirectionNumber() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DirectionNumber", reflect.TypeOf((*MockRotable)(nil).DirectionNumber))
}

// SetDirection mocks base method.
func (m *MockRotable) SetDirection(arg0 int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDirection", arg0)
}

// SetDirection indicates an expected call of SetDirection.
func (mr *MockRotableMockRecorder) SetDirection(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDirection", reflect.TypeOf((*MockRotable)(nil).SetDirection), arg0)
}

// MockMovable is a mock of Movable interface.
type MockMovable struct {
	ctrl     *gomock.Controller
	recorder *MockMovableMockRecorder
}

// MockMovableMockRecorder is the mock recorder for MockMovable.
type MockMovableMockRecorder struct {
	mock *MockMovable
}

// NewMockMovable creates a new mock instance.
func NewMockMovable(ctrl *gomock.Controller) *MockMovable {
	mock := &MockMovable{ctrl: ctrl}
	mock.recorder = &MockMovableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMovable) EXPECT() *MockMovableMockRecorder {
	return m.recorder
}

// Position mocks base method.
func (m *MockMovable) Position() (int64, int64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Position")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(int64)
	return ret0, ret1
}

// Position indicates an expected call of Position.
func (mr *MockMovableMockRecorder) Position() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Position", reflect.TypeOf((*MockMovable)(nil).Position))
}

// SetPosition mocks base method.
func (m *MockMovable) SetPosition(x, y int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPosition", x, y)
}

// SetPosition indicates an expected call of SetPosition.
func (mr *MockMovableMockRecorder) SetPosition(x, y interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPosition", reflect.TypeOf((*MockMovable)(nil).SetPosition), x, y)
}

// Velocity mocks base method.
func (m *MockMovable) Velocity() (int64, int64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Velocity")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(int64)
	return ret0, ret1
}

// Velocity indicates an expected call of Velocity.
func (mr *MockMovableMockRecorder) Velocity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Velocity", reflect.TypeOf((*MockMovable)(nil).Velocity))
}
