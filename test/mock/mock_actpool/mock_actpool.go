// Code generated by MockGen. DO NOT EDIT.
// Source: ./actpool/actpool.go

// Package mock_actpool is a generated GoMock package.
package mock_actpool

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	hash "github.com/iotexproject/go-pkgs/hash"
	address "github.com/iotexproject/iotex-address/address"
	action "github.com/MeshBoxFoundation/mesh-core/action"
	block "github.com/MeshBoxFoundation/mesh-core/blockchain/block"
)

// MockActPool is a mock of ActPool interface.
type MockActPool struct {
	ctrl     *gomock.Controller
	recorder *MockActPoolMockRecorder
}

// MockActPoolMockRecorder is the mock recorder for MockActPool.
type MockActPoolMockRecorder struct {
	mock *MockActPool
}

// NewMockActPool creates a new mock instance.
func NewMockActPool(ctrl *gomock.Controller) *MockActPool {
	mock := &MockActPool{ctrl: ctrl}
	mock.recorder = &MockActPoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockActPool) EXPECT() *MockActPoolMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockActPool) Add(ctx context.Context, act action.SealedEnvelope) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, act)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockActPoolMockRecorder) Add(ctx, act interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockActPool)(nil).Add), ctx, act)
}

// AddActionEnvelopeValidators mocks base method.
func (m *MockActPool) AddActionEnvelopeValidators(arg0 ...action.SealedEnvelopeValidator) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddActionEnvelopeValidators", varargs...)
}

// AddActionEnvelopeValidators indicates an expected call of AddActionEnvelopeValidators.
func (mr *MockActPoolMockRecorder) AddActionEnvelopeValidators(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddActionEnvelopeValidators", reflect.TypeOf((*MockActPool)(nil).AddActionEnvelopeValidators), arg0...)
}

// DeleteAction mocks base method.
func (m *MockActPool) DeleteAction(arg0 address.Address) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteAction", arg0)
}

// DeleteAction indicates an expected call of DeleteAction.
func (mr *MockActPoolMockRecorder) DeleteAction(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAction", reflect.TypeOf((*MockActPool)(nil).DeleteAction), arg0)
}

// GetActionByHash mocks base method.
func (m *MockActPool) GetActionByHash(hash hash.Hash256) (action.SealedEnvelope, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActionByHash", hash)
	ret0, _ := ret[0].(action.SealedEnvelope)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActionByHash indicates an expected call of GetActionByHash.
func (mr *MockActPoolMockRecorder) GetActionByHash(hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActionByHash", reflect.TypeOf((*MockActPool)(nil).GetActionByHash), hash)
}

// GetCapacity mocks base method.
func (m *MockActPool) GetCapacity() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCapacity")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetCapacity indicates an expected call of GetCapacity.
func (mr *MockActPoolMockRecorder) GetCapacity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCapacity", reflect.TypeOf((*MockActPool)(nil).GetCapacity))
}

// GetGasCapacity mocks base method.
func (m *MockActPool) GetGasCapacity() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGasCapacity")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetGasCapacity indicates an expected call of GetGasCapacity.
func (mr *MockActPoolMockRecorder) GetGasCapacity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGasCapacity", reflect.TypeOf((*MockActPool)(nil).GetGasCapacity))
}

// GetGasSize mocks base method.
func (m *MockActPool) GetGasSize() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGasSize")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetGasSize indicates an expected call of GetGasSize.
func (mr *MockActPoolMockRecorder) GetGasSize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGasSize", reflect.TypeOf((*MockActPool)(nil).GetGasSize))
}

// GetPendingNonce mocks base method.
func (m *MockActPool) GetPendingNonce(addr string) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPendingNonce", addr)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPendingNonce indicates an expected call of GetPendingNonce.
func (mr *MockActPoolMockRecorder) GetPendingNonce(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPendingNonce", reflect.TypeOf((*MockActPool)(nil).GetPendingNonce), addr)
}

// GetSize mocks base method.
func (m *MockActPool) GetSize() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSize")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetSize indicates an expected call of GetSize.
func (mr *MockActPoolMockRecorder) GetSize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSize", reflect.TypeOf((*MockActPool)(nil).GetSize))
}

// GetUnconfirmedActs mocks base method.
func (m *MockActPool) GetUnconfirmedActs(addr string) []action.SealedEnvelope {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnconfirmedActs", addr)
	ret0, _ := ret[0].([]action.SealedEnvelope)
	return ret0
}

// GetUnconfirmedActs indicates an expected call of GetUnconfirmedActs.
func (mr *MockActPoolMockRecorder) GetUnconfirmedActs(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnconfirmedActs", reflect.TypeOf((*MockActPool)(nil).GetUnconfirmedActs), addr)
}

// PendingActionMap mocks base method.
func (m *MockActPool) PendingActionMap() map[string][]action.SealedEnvelope {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PendingActionMap")
	ret0, _ := ret[0].(map[string][]action.SealedEnvelope)
	return ret0
}

// PendingActionMap indicates an expected call of PendingActionMap.
func (mr *MockActPoolMockRecorder) PendingActionMap() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PendingActionMap", reflect.TypeOf((*MockActPool)(nil).PendingActionMap))
}

// ReceiveBlock mocks base method.
func (m *MockActPool) ReceiveBlock(arg0 *block.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReceiveBlock", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReceiveBlock indicates an expected call of ReceiveBlock.
func (mr *MockActPoolMockRecorder) ReceiveBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceiveBlock", reflect.TypeOf((*MockActPool)(nil).ReceiveBlock), arg0)
}

// Reset mocks base method.
func (m *MockActPool) Reset() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset")
}

// Reset indicates an expected call of Reset.
func (mr *MockActPoolMockRecorder) Reset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockActPool)(nil).Reset))
}

// Validate mocks base method.
func (m *MockActPool) Validate(arg0 context.Context, arg1 action.SealedEnvelope) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockActPoolMockRecorder) Validate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockActPool)(nil).Validate), arg0, arg1)
}
