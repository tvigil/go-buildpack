// Code generated by MockGen. DO NOT EDIT.
// Source: finalize.go

package finalize_test

import (
	gomock "github.com/golang/mock/gomock"
	io "io"
)

// MockCommand is a mock of Command interface
type MockCommand struct {
	ctrl     *gomock.Controller
	recorder *MockCommandMockRecorder
}

// MockCommandMockRecorder is the mock recorder for MockCommand
type MockCommandMockRecorder struct {
	mock *MockCommand
}

// NewMockCommand creates a new mock instance
func NewMockCommand(ctrl *gomock.Controller) *MockCommand {
	mock := &MockCommand{ctrl: ctrl}
	mock.recorder = &MockCommandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockCommand) EXPECT() *MockCommandMockRecorder {
	return _m.recorder
}

// Execute mocks base method
func (_m *MockCommand) Execute(_param0 string, _param1 io.Writer, _param2 io.Writer, _param3 string, _param4 ...string) error {
	_s := []interface{}{_param0, _param1, _param2, _param3}
	for _, _x := range _param4 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Execute", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute
func (_mr *MockCommandMockRecorder) Execute(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Execute", _s...)
}

// MockStager is a mock of Stager interface
type MockStager struct {
	ctrl     *gomock.Controller
	recorder *MockStagerMockRecorder
}

// MockStagerMockRecorder is the mock recorder for MockStager
type MockStagerMockRecorder struct {
	mock *MockStager
}

// NewMockStager creates a new mock instance
func NewMockStager(ctrl *gomock.Controller) *MockStager {
	mock := &MockStager{ctrl: ctrl}
	mock.recorder = &MockStagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockStager) EXPECT() *MockStagerMockRecorder {
	return _m.recorder
}

// BuildDir mocks base method
func (_m *MockStager) BuildDir() string {
	ret := _m.ctrl.Call(_m, "BuildDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// BuildDir indicates an expected call of BuildDir
func (_mr *MockStagerMockRecorder) BuildDir() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BuildDir")
}

// ClearDepDir mocks base method
func (_m *MockStager) ClearDepDir() error {
	ret := _m.ctrl.Call(_m, "ClearDepDir")
	ret0, _ := ret[0].(error)
	return ret0
}

// ClearDepDir indicates an expected call of ClearDepDir
func (_mr *MockStagerMockRecorder) ClearDepDir() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ClearDepDir")
}

// DepDir mocks base method
func (_m *MockStager) DepDir() string {
	ret := _m.ctrl.Call(_m, "DepDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// DepDir indicates an expected call of DepDir
func (_mr *MockStagerMockRecorder) DepDir() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DepDir")
}

// DepsIdx mocks base method
func (_m *MockStager) DepsIdx() string {
	ret := _m.ctrl.Call(_m, "DepsIdx")
	ret0, _ := ret[0].(string)
	return ret0
}

// DepsIdx indicates an expected call of DepsIdx
func (_mr *MockStagerMockRecorder) DepsIdx() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DepsIdx")
}

// WriteProfileD mocks base method
func (_m *MockStager) WriteProfileD(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "WriteProfileD", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteProfileD indicates an expected call of WriteProfileD
func (_mr *MockStagerMockRecorder) WriteProfileD(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WriteProfileD", arg0, arg1)
}
