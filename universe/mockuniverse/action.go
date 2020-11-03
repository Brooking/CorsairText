// Code generated by MockGen. DO NOT EDIT.
// Source: action.go

// Package mockuniverse is a generated GoMock package.
package mockuniverse

import (
	action "corsairtext/action"
	universe "corsairtext/universe"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAction is a mock of Action interface
type MockAction struct {
	ctrl     *gomock.Controller
	recorder *MockActionMockRecorder
}

// MockActionMockRecorder is the mock recorder for MockAction
type MockActionMockRecorder struct {
	mock *MockAction
}

// NewMockAction creates a new mock instance
func NewMockAction(ctrl *gomock.Controller) *MockAction {
	mock := &MockAction{ctrl: ctrl}
	mock.recorder = &MockActionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAction) EXPECT() *MockActionMockRecorder {
	return m.recorder
}

// Buy mocks base method
func (m *MockAction) Buy(arg0 int, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Buy", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Buy indicates an expected call of Buy
func (mr *MockActionMockRecorder) Buy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Buy", reflect.TypeOf((*MockAction)(nil).Buy), arg0, arg1)
}

// Go mocks base method
func (m *MockAction) Go(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Go", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Go indicates an expected call of Go
func (mr *MockActionMockRecorder) Go(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Go", reflect.TypeOf((*MockAction)(nil).Go), arg0)
}

// Help mocks base method
func (m *MockAction) Help() ([]action.Type, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Help")
	ret0, _ := ret[0].([]action.Type)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Help indicates an expected call of Help
func (mr *MockActionMockRecorder) Help() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Help", reflect.TypeOf((*MockAction)(nil).Help))
}

// Look mocks base method
func (m *MockAction) Look() (universe.View, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Look")
	ret0, _ := ret[0].(universe.View)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Look indicates an expected call of Look
func (mr *MockActionMockRecorder) Look() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Look", reflect.TypeOf((*MockAction)(nil).Look))
}

// Mine mocks base method
func (m *MockAction) Mine() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mine")
	ret0, _ := ret[0].(error)
	return ret0
}

// Mine indicates an expected call of Mine
func (mr *MockActionMockRecorder) Mine() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mine", reflect.TypeOf((*MockAction)(nil).Mine))
}

// Quit mocks base method
func (m *MockAction) Quit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Quit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Quit indicates an expected call of Quit
func (mr *MockActionMockRecorder) Quit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Quit", reflect.TypeOf((*MockAction)(nil).Quit))
}

// Sell mocks base method
func (m *MockAction) Sell(arg0 int, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sell", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Sell indicates an expected call of Sell
func (mr *MockActionMockRecorder) Sell(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sell", reflect.TypeOf((*MockAction)(nil).Sell), arg0, arg1)
}
