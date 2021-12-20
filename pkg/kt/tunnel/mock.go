// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/kt/connect/types.go

// Package connect is a generated GoMock package.
package tunnel

import (
	reflect "reflect"

	exec "github.com/alibaba/kt-connect/pkg/kt/exec"
	util "github.com/alibaba/kt-connect/pkg/kt/util"
	gomock "github.com/golang/mock/gomock"
)

// MockShadowInterface is a mock of ShadowInterface interface.
type MockShadowInterface struct {
	ctrl     *gomock.Controller
	recorder *MockShadowInterfaceMockRecorder
}

// MockShadowInterfaceMockRecorder is the mock recorder for MockShadowInterface.
type MockShadowInterfaceMockRecorder struct {
	mock *MockShadowInterface
}

// NewMockShadowInterface creates a new mock instance.
func NewMockShadowInterface(ctrl *gomock.Controller) *MockShadowInterface {
	mock := &MockShadowInterface{ctrl: ctrl}
	mock.recorder = &MockShadowInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShadowInterface) EXPECT() *MockShadowInterfaceMockRecorder {
	return m.recorder
}

// Inbound mocks base method.
func (m *MockShadowInterface) Inbound(exposePort, podName string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Inbound", exposePort, podName)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Inbound indicates an expected call of Inbound.
func (mr *MockShadowInterfaceMockRecorder) Inbound(exposePort, podName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Inbound", reflect.TypeOf((*MockShadowInterface)(nil).Inbound), exposePort, podName)
}

// Outbound mocks base method.
func (m *MockShadowInterface) Outbound(name, podIP string, credential *util.SSHCredential, cidrs []string, exec exec.CliInterface) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Outbound", name, podIP, credential, cidrs, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Outbound indicates an expected call of Outbound.
func (mr *MockShadowInterfaceMockRecorder) Outbound(name, podIP, credential, cidrs, exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Outbound", reflect.TypeOf((*MockShadowInterface)(nil).Outbound), name, podIP, credential, cidrs, exec)
}