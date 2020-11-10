// Code generated by MockGen. DO NOT EDIT.
// Source: base.go

// Package mockbase is a generated GoMock package.
package mockbase

import (
	action "corsairtext/universe/action"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBase is a mock of Base interface
type MockBase struct {
	ctrl     *gomock.Controller
	recorder *MockBaseMockRecorder
}

// MockBaseMockRecorder is the mock recorder for MockBase
type MockBaseMockRecorder struct {
	mock *MockBase
}

// NewMockBase creates a new mock instance
func NewMockBase(ctrl *gomock.Controller) *MockBase {
	mock := &MockBase{ctrl: ctrl}
	mock.recorder = &MockBaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBase) EXPECT() *MockBaseMockRecorder {
	return m.recorder
}

// Actions mocks base method
func (m *MockBase) Actions() action.List {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Actions")
	ret0, _ := ret[0].(action.List)
	return ret0
}

// Actions indicates an expected call of Actions
func (mr *MockBaseMockRecorder) Actions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Actions", reflect.TypeOf((*MockBase)(nil).Actions))
}
