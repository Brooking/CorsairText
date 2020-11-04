// Code generated by MockGen. DO NOT EDIT.
// Source: spot.go

// Package mockspot is a generated GoMock package.
package mockspot

import (
	action "corsairtext/action"
	spot "corsairtext/universe/spot"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSpot is a mock of Spot interface
type MockSpot struct {
	ctrl     *gomock.Controller
	recorder *MockSpotMockRecorder
}

// MockSpotMockRecorder is the mock recorder for MockSpot
type MockSpotMockRecorder struct {
	mock *MockSpot
}

// NewMockSpot creates a new mock instance
func NewMockSpot(ctrl *gomock.Controller) *MockSpot {
	mock := &MockSpot{ctrl: ctrl}
	mock.recorder = &MockSpotMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSpot) EXPECT() *MockSpotMockRecorder {
	return m.recorder
}

// Actions mocks base method
func (m *MockSpot) Actions() action.List {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Actions")
	ret0, _ := ret[0].(action.List)
	return ret0
}

// Actions indicates an expected call of Actions
func (mr *MockSpotMockRecorder) Actions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Actions", reflect.TypeOf((*MockSpot)(nil).Actions))
}

// Description mocks base method
func (m *MockSpot) Description() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Description")
	ret0, _ := ret[0].(string)
	return ret0
}

// Description indicates an expected call of Description
func (mr *MockSpotMockRecorder) Description() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Description", reflect.TypeOf((*MockSpot)(nil).Description))
}

// Name mocks base method
func (m *MockSpot) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockSpotMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockSpot)(nil).Name))
}

// Path mocks base method
func (m *MockSpot) Path() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Path")
	ret0, _ := ret[0].(string)
	return ret0
}

// Path indicates an expected call of Path
func (mr *MockSpotMockRecorder) Path() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Path", reflect.TypeOf((*MockSpot)(nil).Path))
}

// ListAdjacent mocks base method
func (m *MockSpot) ListAdjacent() []spot.Spot {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAdjacent")
	ret0, _ := ret[0].([]spot.Spot)
	return ret0
}

// ListAdjacent indicates an expected call of ListAdjacent
func (mr *MockSpotMockRecorder) ListAdjacent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAdjacent", reflect.TypeOf((*MockSpot)(nil).ListAdjacent))
}

// AddChild mocks base method
func (m *MockSpot) AddChild(child spot.Spot) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddChild", child)
}

// AddChild indicates an expected call of AddChild
func (mr *MockSpotMockRecorder) AddChild(child interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddChild", reflect.TypeOf((*MockSpot)(nil).AddChild), child)
}

// Parent mocks base method
func (m *MockSpot) Parent() spot.Spot {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parent")
	ret0, _ := ret[0].(spot.Spot)
	return ret0
}

// Parent indicates an expected call of Parent
func (mr *MockSpotMockRecorder) Parent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parent", reflect.TypeOf((*MockSpot)(nil).Parent))
}

// Children mocks base method
func (m *MockSpot) Children() []spot.Spot {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Children")
	ret0, _ := ret[0].([]spot.Spot)
	return ret0
}

// Children indicates an expected call of Children
func (mr *MockSpotMockRecorder) Children() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Children", reflect.TypeOf((*MockSpot)(nil).Children))
}
