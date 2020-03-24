// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_mesh is a generated GoMock package.
package mock_mesh

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/mesh-projects/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	controller "github.com/solo-io/mesh-projects/services/common/cluster/apps/v1/controller"
	v1 "k8s.io/api/apps/v1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMeshFinder is a mock of MeshFinder interface.
type MockMeshFinder struct {
	ctrl     *gomock.Controller
	recorder *MockMeshFinderMockRecorder
}

// MockMeshFinderMockRecorder is the mock recorder for MockMeshFinder.
type MockMeshFinderMockRecorder struct {
	mock *MockMeshFinder
}

// NewMockMeshFinder creates a new mock instance.
func NewMockMeshFinder(ctrl *gomock.Controller) *MockMeshFinder {
	mock := &MockMeshFinder{ctrl: ctrl}
	mock.recorder = &MockMeshFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMeshFinder) EXPECT() *MockMeshFinderMockRecorder {
	return m.recorder
}

// StartDiscovery mocks base method.
func (m *MockMeshFinder) StartDiscovery(deploymentController controller.DeploymentController, predicates []predicate.Predicate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartDiscovery", deploymentController, predicates)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartDiscovery indicates an expected call of StartDiscovery.
func (mr *MockMeshFinderMockRecorder) StartDiscovery(deploymentController, predicates interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartDiscovery", reflect.TypeOf((*MockMeshFinder)(nil).StartDiscovery), deploymentController, predicates)
}

// Create mocks base method.
func (m *MockMeshFinder) Create(obj *v1.Deployment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockMeshFinderMockRecorder) Create(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMeshFinder)(nil).Create), obj)
}

// Update mocks base method.
func (m *MockMeshFinder) Update(old, new *v1.Deployment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockMeshFinderMockRecorder) Update(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMeshFinder)(nil).Update), old, new)
}

// Delete mocks base method.
func (m *MockMeshFinder) Delete(obj *v1.Deployment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMeshFinderMockRecorder) Delete(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMeshFinder)(nil).Delete), obj)
}

// Generic mocks base method.
func (m *MockMeshFinder) Generic(obj *v1.Deployment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// Generic indicates an expected call of Generic.
func (mr *MockMeshFinderMockRecorder) Generic(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockMeshFinder)(nil).Generic), obj)
}

// MockMeshScanner is a mock of MeshScanner interface.
type MockMeshScanner struct {
	ctrl     *gomock.Controller
	recorder *MockMeshScannerMockRecorder
}

// MockMeshScannerMockRecorder is the mock recorder for MockMeshScanner.
type MockMeshScannerMockRecorder struct {
	mock *MockMeshScanner
}

// NewMockMeshScanner creates a new mock instance.
func NewMockMeshScanner(ctrl *gomock.Controller) *MockMeshScanner {
	mock := &MockMeshScanner{ctrl: ctrl}
	mock.recorder = &MockMeshScannerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMeshScanner) EXPECT() *MockMeshScannerMockRecorder {
	return m.recorder
}

// ScanDeployment mocks base method.
func (m *MockMeshScanner) ScanDeployment(arg0 context.Context, arg1 *v1.Deployment, arg2 client.Client) (*v1alpha1.Mesh, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScanDeployment", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1alpha1.Mesh)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScanDeployment indicates an expected call of ScanDeployment.
func (mr *MockMeshScannerMockRecorder) ScanDeployment(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanDeployment", reflect.TypeOf((*MockMeshScanner)(nil).ScanDeployment), arg0, arg1, arg2)
}