// Code generated by MockGen. DO NOT EDIT.
// Source: command.go

// Package mockcommandprocessor is a generated GoMock package.
package mockcommandprocessor

import (
	match "corsairtext/textui/commandprocessor/match"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCommandProcessor is a mock of CommandProcessor interface
type MockCommandProcessor struct {
	ctrl     *gomock.Controller
	recorder *MockCommandProcessorMockRecorder
}

// MockCommandProcessorMockRecorder is the mock recorder for MockCommandProcessor
type MockCommandProcessorMockRecorder struct {
	mock *MockCommandProcessor
}

// NewMockCommandProcessor creates a new mock instance
func NewMockCommandProcessor(ctrl *gomock.Controller) *MockCommandProcessor {
	mock := &MockCommandProcessor{ctrl: ctrl}
	mock.recorder = &MockCommandProcessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCommandProcessor) EXPECT() *MockCommandProcessorMockRecorder {
	return m.recorder
}

// CommandMatcher mocks base method
func (m *MockCommandProcessor) CommandMatcher() match.Matcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommandMatcher")
	ret0, _ := ret[0].(match.Matcher)
	return ret0
}

// CommandMatcher indicates an expected call of CommandMatcher
func (mr *MockCommandProcessorMockRecorder) CommandMatcher() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommandMatcher", reflect.TypeOf((*MockCommandProcessor)(nil).CommandMatcher))
}

// ShowAdjacency mocks base method
func (m *MockCommandProcessor) ShowAdjacency() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowAdjacency")
	ret0, _ := ret[0].(error)
	return ret0
}

// ShowAdjacency indicates an expected call of ShowAdjacency
func (mr *MockCommandProcessorMockRecorder) ShowAdjacency() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowAdjacency", reflect.TypeOf((*MockCommandProcessor)(nil).ShowAdjacency))
}

// ShowAllHelp mocks base method
func (m *MockCommandProcessor) ShowAllHelp() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowAllHelp")
	ret0, _ := ret[0].(error)
	return ret0
}

// ShowAllHelp indicates an expected call of ShowAllHelp
func (mr *MockCommandProcessorMockRecorder) ShowAllHelp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowAllHelp", reflect.TypeOf((*MockCommandProcessor)(nil).ShowAllHelp))
}

// ShowHelp mocks base method
func (m *MockCommandProcessor) ShowHelp(command string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowHelp", command)
	ret0, _ := ret[0].(error)
	return ret0
}

// ShowHelp indicates an expected call of ShowHelp
func (mr *MockCommandProcessorMockRecorder) ShowHelp(command interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowHelp", reflect.TypeOf((*MockCommandProcessor)(nil).ShowHelp), command)
}

// ShowLook mocks base method
func (m *MockCommandProcessor) ShowLook() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowLook")
	ret0, _ := ret[0].(error)
	return ret0
}

// ShowLook indicates an expected call of ShowLook
func (mr *MockCommandProcessorMockRecorder) ShowLook() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowLook", reflect.TypeOf((*MockCommandProcessor)(nil).ShowLook))
}

// Obey mocks base method
func (m *MockCommandProcessor) Obey(commandLine string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Obey", commandLine)
	ret0, _ := ret[0].(error)
	return ret0
}

// Obey indicates an expected call of Obey
func (mr *MockCommandProcessorMockRecorder) Obey(commandLine interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Obey", reflect.TypeOf((*MockCommandProcessor)(nil).Obey), commandLine)
}