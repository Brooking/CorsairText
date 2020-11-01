// Code generated by MockGen. DO NOT EDIT.
// Source: keyboardreader.go

// Package mockkeyboardreader is a generated GoMock package.
package mockkeyboardreader

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockKeyboardReader is a mock of KeyboardReader interface
type MockKeyboardReader struct {
	ctrl     *gomock.Controller
	recorder *MockKeyboardReaderMockRecorder
}

// MockKeyboardReaderMockRecorder is the mock recorder for MockKeyboardReader
type MockKeyboardReaderMockRecorder struct {
	mock *MockKeyboardReader
}

// NewMockKeyboardReader creates a new mock instance
func NewMockKeyboardReader(ctrl *gomock.Controller) *MockKeyboardReader {
	mock := &MockKeyboardReader{ctrl: ctrl}
	mock.recorder = &MockKeyboardReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKeyboardReader) EXPECT() *MockKeyboardReaderMockRecorder {
	return m.recorder
}

// Read mocks base method
func (m *MockKeyboardReader) Read() (rune, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read")
	ret0, _ := ret[0].(rune)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockKeyboardReaderMockRecorder) Read() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockKeyboardReader)(nil).Read))
}

// ReadLn mocks base method
func (m *MockKeyboardReader) ReadLn() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadLn")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadLn indicates an expected call of ReadLn
func (mr *MockKeyboardReaderMockRecorder) ReadLn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadLn", reflect.TypeOf((*MockKeyboardReader)(nil).ReadLn))
}
