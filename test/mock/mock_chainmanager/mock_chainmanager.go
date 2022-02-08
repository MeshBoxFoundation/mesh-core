// Code generated by MockGen. DO NOT EDIT.
// Source: ./action/protocol/managers.go

// Package mock_chainmanager is a generated GoMock package.
package mock_chainmanager

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	protocol "github.com/MeshBoxFoundation/mesh-core/action/protocol"
	state "github.com/MeshBoxFoundation/mesh-core/state"
)

// MockStateReader is a mock of StateReader interface.
type MockStateReader struct {
	ctrl     *gomock.Controller
	recorder *MockStateReaderMockRecorder
}

// MockStateReaderMockRecorder is the mock recorder for MockStateReader.
type MockStateReaderMockRecorder struct {
	mock *MockStateReader
}

// NewMockStateReader creates a new mock instance.
func NewMockStateReader(ctrl *gomock.Controller) *MockStateReader {
	mock := &MockStateReader{ctrl: ctrl}
	mock.recorder = &MockStateReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStateReader) EXPECT() *MockStateReaderMockRecorder {
	return m.recorder
}

// Height mocks base method.
func (m *MockStateReader) Height() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Height")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Height indicates an expected call of Height.
func (mr *MockStateReaderMockRecorder) Height() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Height", reflect.TypeOf((*MockStateReader)(nil).Height))
}

// ReadView mocks base method.
func (m *MockStateReader) ReadView(arg0 string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadView", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadView indicates an expected call of ReadView.
func (mr *MockStateReaderMockRecorder) ReadView(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadView", reflect.TypeOf((*MockStateReader)(nil).ReadView), arg0)
}

// State mocks base method.
func (m *MockStateReader) State(arg0 interface{}, arg1 ...protocol.StateOption) (uint64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "State", varargs...)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// State indicates an expected call of State.
func (mr *MockStateReaderMockRecorder) State(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockStateReader)(nil).State), varargs...)
}

// States mocks base method.
func (m *MockStateReader) States(arg0 ...protocol.StateOption) (uint64, state.Iterator, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "States", varargs...)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(state.Iterator)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// States indicates an expected call of States.
func (mr *MockStateReaderMockRecorder) States(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "States", reflect.TypeOf((*MockStateReader)(nil).States), arg0...)
}

// MockStateManager is a mock of StateManager interface.
type MockStateManager struct {
	ctrl     *gomock.Controller
	recorder *MockStateManagerMockRecorder
}

// MockStateManagerMockRecorder is the mock recorder for MockStateManager.
type MockStateManagerMockRecorder struct {
	mock *MockStateManager
}

// NewMockStateManager creates a new mock instance.
func NewMockStateManager(ctrl *gomock.Controller) *MockStateManager {
	mock := &MockStateManager{ctrl: ctrl}
	mock.recorder = &MockStateManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStateManager) EXPECT() *MockStateManagerMockRecorder {
	return m.recorder
}

// DelState mocks base method.
func (m *MockStateManager) DelState(arg0 ...protocol.StateOption) (uint64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DelState", varargs...)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DelState indicates an expected call of DelState.
func (mr *MockStateManagerMockRecorder) DelState(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelState", reflect.TypeOf((*MockStateManager)(nil).DelState), arg0...)
}

// Height mocks base method.
func (m *MockStateManager) Height() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Height")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Height indicates an expected call of Height.
func (mr *MockStateManagerMockRecorder) Height() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Height", reflect.TypeOf((*MockStateManager)(nil).Height))
}

// Load mocks base method.
func (m *MockStateManager) Load(arg0, arg1 string, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Load indicates an expected call of Load.
func (mr *MockStateManagerMockRecorder) Load(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockStateManager)(nil).Load), arg0, arg1, arg2)
}

// ProtocolDirty mocks base method.
func (m *MockStateManager) ProtocolDirty(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProtocolDirty", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ProtocolDirty indicates an expected call of ProtocolDirty.
func (mr *MockStateManagerMockRecorder) ProtocolDirty(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProtocolDirty", reflect.TypeOf((*MockStateManager)(nil).ProtocolDirty), arg0)
}

// PutState mocks base method.
func (m *MockStateManager) PutState(arg0 interface{}, arg1 ...protocol.StateOption) (uint64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutState", varargs...)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutState indicates an expected call of PutState.
func (mr *MockStateManagerMockRecorder) PutState(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutState", reflect.TypeOf((*MockStateManager)(nil).PutState), varargs...)
}

// ReadView mocks base method.
func (m *MockStateManager) ReadView(arg0 string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadView", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadView indicates an expected call of ReadView.
func (mr *MockStateManagerMockRecorder) ReadView(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadView", reflect.TypeOf((*MockStateManager)(nil).ReadView), arg0)
}

// Reset mocks base method.
func (m *MockStateManager) Reset() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset")
}

// Reset indicates an expected call of Reset.
func (mr *MockStateManagerMockRecorder) Reset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockStateManager)(nil).Reset))
}

// Revert mocks base method.
func (m *MockStateManager) Revert(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Revert", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Revert indicates an expected call of Revert.
func (mr *MockStateManagerMockRecorder) Revert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Revert", reflect.TypeOf((*MockStateManager)(nil).Revert), arg0)
}

// Snapshot mocks base method.
func (m *MockStateManager) Snapshot() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshot")
	ret0, _ := ret[0].(int)
	return ret0
}

// Snapshot indicates an expected call of Snapshot.
func (mr *MockStateManagerMockRecorder) Snapshot() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshot", reflect.TypeOf((*MockStateManager)(nil).Snapshot))
}

// State mocks base method.
func (m *MockStateManager) State(arg0 interface{}, arg1 ...protocol.StateOption) (uint64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "State", varargs...)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// State indicates an expected call of State.
func (mr *MockStateManagerMockRecorder) State(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockStateManager)(nil).State), varargs...)
}

// States mocks base method.
func (m *MockStateManager) States(arg0 ...protocol.StateOption) (uint64, state.Iterator, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "States", varargs...)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(state.Iterator)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// States indicates an expected call of States.
func (mr *MockStateManagerMockRecorder) States(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "States", reflect.TypeOf((*MockStateManager)(nil).States), arg0...)
}

// Unload mocks base method.
func (m *MockStateManager) Unload(arg0, arg1 string, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unload", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unload indicates an expected call of Unload.
func (mr *MockStateManagerMockRecorder) Unload(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unload", reflect.TypeOf((*MockStateManager)(nil).Unload), arg0, arg1, arg2)
}

// WriteView mocks base method.
func (m *MockStateManager) WriteView(arg0 string, arg1 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteView", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteView indicates an expected call of WriteView.
func (mr *MockStateManagerMockRecorder) WriteView(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteView", reflect.TypeOf((*MockStateManager)(nil).WriteView), arg0, arg1)
}

// MockDock is a mock of Dock interface.
type MockDock struct {
	ctrl     *gomock.Controller
	recorder *MockDockMockRecorder
}

// MockDockMockRecorder is the mock recorder for MockDock.
type MockDockMockRecorder struct {
	mock *MockDock
}

// NewMockDock creates a new mock instance.
func NewMockDock(ctrl *gomock.Controller) *MockDock {
	mock := &MockDock{ctrl: ctrl}
	mock.recorder = &MockDockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDock) EXPECT() *MockDockMockRecorder {
	return m.recorder
}

// Load mocks base method.
func (m *MockDock) Load(arg0, arg1 string, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Load indicates an expected call of Load.
func (mr *MockDockMockRecorder) Load(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockDock)(nil).Load), arg0, arg1, arg2)
}

// ProtocolDirty mocks base method.
func (m *MockDock) ProtocolDirty(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProtocolDirty", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ProtocolDirty indicates an expected call of ProtocolDirty.
func (mr *MockDockMockRecorder) ProtocolDirty(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProtocolDirty", reflect.TypeOf((*MockDock)(nil).ProtocolDirty), arg0)
}

// Reset mocks base method.
func (m *MockDock) Reset() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset")
}

// Reset indicates an expected call of Reset.
func (mr *MockDockMockRecorder) Reset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockDock)(nil).Reset))
}

// Unload mocks base method.
func (m *MockDock) Unload(arg0, arg1 string, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unload", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unload indicates an expected call of Unload.
func (mr *MockDockMockRecorder) Unload(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unload", reflect.TypeOf((*MockDock)(nil).Unload), arg0, arg1, arg2)
}
