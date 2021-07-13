// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/kt/cluster/types.go

// Package cluster is a generated GoMock package.
package cluster

import (
	reflect "reflect"

	options "github.com/alibaba/kt-connect/pkg/kt/options"
	util "github.com/alibaba/kt-connect/pkg/kt/util"
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/apps/v1"
	v10 "k8s.io/api/core/v1"
)

// MockKubernetesInterface is a mock of KubernetesInterface interface.
type MockKubernetesInterface struct {
	ctrl     *gomock.Controller
	recorder *MockKubernetesInterfaceMockRecorder
}

// MockKubernetesInterfaceMockRecorder is the mock recorder for MockKubernetesInterface.
type MockKubernetesInterfaceMockRecorder struct {
	mock *MockKubernetesInterface
}

// NewMockKubernetesInterface creates a new mock instance.
func NewMockKubernetesInterface(ctrl *gomock.Controller) *MockKubernetesInterface {
	mock := &MockKubernetesInterface{ctrl: ctrl}
	mock.recorder = &MockKubernetesInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKubernetesInterface) EXPECT() *MockKubernetesInterfaceMockRecorder {
	return m.recorder
}

// ClusterCrids mocks base method.
func (m *MockKubernetesInterface) ClusterCrids(namespace string, connectOptions *options.ConnectOptions) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterCrids", namespace, connectOptions)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClusterCrids indicates an expected call of ClusterCrids.
func (mr *MockKubernetesInterfaceMockRecorder) ClusterCrids(namespace, connectOptions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterCrids", reflect.TypeOf((*MockKubernetesInterface)(nil).ClusterCrids), namespace, connectOptions)
}

// CreateService mocks base method.
func (m *MockKubernetesInterface) CreateService(name, namespace string, external bool, port int, labels map[string]string) (*v10.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateService", name, namespace, external, port, labels)
	ret0, _ := ret[0].(*v10.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateService indicates an expected call of CreateService.
func (mr *MockKubernetesInterfaceMockRecorder) CreateService(name, namespace, external, port, labels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateService", reflect.TypeOf((*MockKubernetesInterface)(nil).CreateService), name, namespace, external, port, labels)
}

// DecreaseRef mocks base method.
func (m *MockKubernetesInterface) DecreaseRef(namespace, deployment string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecreaseRef", namespace, deployment)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecreaseRef indicates an expected call of DecreaseRef.
func (mr *MockKubernetesInterfaceMockRecorder) DecreaseRef(namespace, deployment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecreaseRef", reflect.TypeOf((*MockKubernetesInterface)(nil).DecreaseRef), namespace, deployment)
}

// Deployment mocks base method.
func (m *MockKubernetesInterface) Deployment(name, namespace string) (*v1.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deployment", name, namespace)
	ret0, _ := ret[0].(*v1.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Deployment indicates an expected call of Deployment.
func (mr *MockKubernetesInterfaceMockRecorder) Deployment(name, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deployment", reflect.TypeOf((*MockKubernetesInterface)(nil).Deployment), name, namespace)
}

// GetAllExistingShadowDeployments mocks base method.
func (m *MockKubernetesInterface) GetAllExistingShadowDeployments(namespace string) ([]v1.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllExistingShadowDeployments", namespace)
	ret0, _ := ret[0].([]v1.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllExistingShadowDeployments indicates an expected call of GetAllExistingShadowDeployments.
func (mr *MockKubernetesInterfaceMockRecorder) GetAllExistingShadowDeployments(namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllExistingShadowDeployments", reflect.TypeOf((*MockKubernetesInterface)(nil).GetAllExistingShadowDeployments), namespace)
}

// GetDeployment mocks base method.
func (m *MockKubernetesInterface) GetDeployment(name, namespace string) (*v1.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeployment", name, namespace)
	ret0, _ := ret[0].(*v1.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeployment indicates an expected call of GetDeployment.
func (mr *MockKubernetesInterfaceMockRecorder) GetDeployment(name, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeployment", reflect.TypeOf((*MockKubernetesInterface)(nil).GetDeployment), name, namespace)
}

// GetOrCreateShadow mocks base method.
func (m *MockKubernetesInterface) GetOrCreateShadow(name, namespace, image string, labels, annotations, envs map[string]string, debug, reuseShadow bool) (string, string, string, *util.SSHCredential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrCreateShadow", name, namespace, image, labels, annotations, envs, debug, reuseShadow)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(*util.SSHCredential)
	ret4, _ := ret[4].(error)
	return ret0, ret1, ret2, ret3, ret4
}

// GetOrCreateShadow indicates an expected call of GetOrCreateShadow.
func (mr *MockKubernetesInterfaceMockRecorder) GetOrCreateShadow(name, namespace, image, labels, annotations, envs, debug, reuseShadow interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrCreateShadow", reflect.TypeOf((*MockKubernetesInterface)(nil).GetOrCreateShadow), name, namespace, image, labels, annotations, envs, debug, reuseShadow)
}

// RemoveConfigMap mocks base method.
func (m *MockKubernetesInterface) RemoveConfigMap(name, namespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveConfigMap", name, namespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveConfigMap indicates an expected call of RemoveConfigMap.
func (mr *MockKubernetesInterfaceMockRecorder) RemoveConfigMap(name, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveConfigMap", reflect.TypeOf((*MockKubernetesInterface)(nil).RemoveConfigMap), name, namespace)
}

// RemoveDeployment mocks base method.
func (m *MockKubernetesInterface) RemoveDeployment(name, namespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveDeployment", name, namespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveDeployment indicates an expected call of RemoveDeployment.
func (mr *MockKubernetesInterfaceMockRecorder) RemoveDeployment(name, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveDeployment", reflect.TypeOf((*MockKubernetesInterface)(nil).RemoveDeployment), name, namespace)
}

// RemoveService mocks base method.
func (m *MockKubernetesInterface) RemoveService(name, namespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveService", name, namespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveService indicates an expected call of RemoveService.
func (mr *MockKubernetesInterfaceMockRecorder) RemoveService(name, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveService", reflect.TypeOf((*MockKubernetesInterface)(nil).RemoveService), name, namespace)
}

// Scale mocks base method.
func (m *MockKubernetesInterface) Scale(deployment *v1.Deployment, replicas *int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scale", deployment, replicas)
	ret0, _ := ret[0].(error)
	return ret0
}

// Scale indicates an expected call of Scale.
func (mr *MockKubernetesInterfaceMockRecorder) Scale(deployment, replicas interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scale", reflect.TypeOf((*MockKubernetesInterface)(nil).Scale), deployment, replicas)
}

// ScaleTo mocks base method.
func (m *MockKubernetesInterface) ScaleTo(deployment, namespace string, replicas *int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScaleTo", deployment, namespace, replicas)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScaleTo indicates an expected call of ScaleTo.
func (mr *MockKubernetesInterfaceMockRecorder) ScaleTo(deployment, namespace, replicas interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScaleTo", reflect.TypeOf((*MockKubernetesInterface)(nil).ScaleTo), deployment, namespace, replicas)
}

// ServiceHosts mocks base method.
func (m *MockKubernetesInterface) ServiceHosts(namespace string) map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceHosts", namespace)
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// ServiceHosts indicates an expected call of ServiceHosts.
func (mr *MockKubernetesInterfaceMockRecorder) ServiceHosts(namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceHosts", reflect.TypeOf((*MockKubernetesInterface)(nil).ServiceHosts), namespace)
}

// UpdateDeployment mocks base method.
func (m *MockKubernetesInterface) UpdateDeployment(namespace string, deployment *v1.Deployment) (*v1.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDeployment", namespace, deployment)
	ret0, _ := ret[0].(*v1.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateDeployment indicates an expected call of UpdateDeployment.
func (mr *MockKubernetesInterfaceMockRecorder) UpdateDeployment(namespace, deployment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDeployment", reflect.TypeOf((*MockKubernetesInterface)(nil).UpdateDeployment), namespace, deployment)
}