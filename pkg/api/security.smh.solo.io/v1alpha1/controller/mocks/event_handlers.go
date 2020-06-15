// Code generated by MockGen. DO NOT EDIT.
// Source: ./event_handlers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/security.smh.solo.io/v1alpha1"
	controller "github.com/solo-io/service-mesh-hub/pkg/api/security.smh.solo.io/v1alpha1/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockVirtualMeshCertificateSigningRequestEventHandler is a mock of VirtualMeshCertificateSigningRequestEventHandler interface.
type MockVirtualMeshCertificateSigningRequestEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualMeshCertificateSigningRequestEventHandlerMockRecorder
}

// MockVirtualMeshCertificateSigningRequestEventHandlerMockRecorder is the mock recorder for MockVirtualMeshCertificateSigningRequestEventHandler.
type MockVirtualMeshCertificateSigningRequestEventHandlerMockRecorder struct {
	mock *MockVirtualMeshCertificateSigningRequestEventHandler
}

// NewMockVirtualMeshCertificateSigningRequestEventHandler creates a new mock instance.
func NewMockVirtualMeshCertificateSigningRequestEventHandler(ctrl *gomock.Controller) *MockVirtualMeshCertificateSigningRequestEventHandler {
	mock := &MockVirtualMeshCertificateSigningRequestEventHandler{ctrl: ctrl}
	mock.recorder = &MockVirtualMeshCertificateSigningRequestEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualMeshCertificateSigningRequestEventHandler) EXPECT() *MockVirtualMeshCertificateSigningRequestEventHandlerMockRecorder {
	return m.recorder
}

// CreateVirtualMeshCertificateSigningRequest mocks base method.
func (m *MockVirtualMeshCertificateSigningRequestEventHandler) CreateVirtualMeshCertificateSigningRequest(obj *v1alpha1.VirtualMeshCertificateSigningRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVirtualMeshCertificateSigningRequest", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateVirtualMeshCertificateSigningRequest indicates an expected call of CreateVirtualMeshCertificateSigningRequest.
func (mr *MockVirtualMeshCertificateSigningRequestEventHandlerMockRecorder) CreateVirtualMeshCertificateSigningRequest(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVirtualMeshCertificateSigningRequest", reflect.TypeOf((*MockVirtualMeshCertificateSigningRequestEventHandler)(nil).CreateVirtualMeshCertificateSigningRequest), obj)
}

// UpdateVirtualMeshCertificateSigningRequest mocks base method.
func (m *MockVirtualMeshCertificateSigningRequestEventHandler) UpdateVirtualMeshCertificateSigningRequest(old, new *v1alpha1.VirtualMeshCertificateSigningRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVirtualMeshCertificateSigningRequest", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateVirtualMeshCertificateSigningRequest indicates an expected call of UpdateVirtualMeshCertificateSigningRequest.
func (mr *MockVirtualMeshCertificateSigningRequestEventHandlerMockRecorder) UpdateVirtualMeshCertificateSigningRequest(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVirtualMeshCertificateSigningRequest", reflect.TypeOf((*MockVirtualMeshCertificateSigningRequestEventHandler)(nil).UpdateVirtualMeshCertificateSigningRequest), old, new)
}

// DeleteVirtualMeshCertificateSigningRequest mocks base method.
func (m *MockVirtualMeshCertificateSigningRequestEventHandler) DeleteVirtualMeshCertificateSigningRequest(obj *v1alpha1.VirtualMeshCertificateSigningRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVirtualMeshCertificateSigningRequest", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVirtualMeshCertificateSigningRequest indicates an expected call of DeleteVirtualMeshCertificateSigningRequest.
func (mr *MockVirtualMeshCertificateSigningRequestEventHandlerMockRecorder) DeleteVirtualMeshCertificateSigningRequest(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVirtualMeshCertificateSigningRequest", reflect.TypeOf((*MockVirtualMeshCertificateSigningRequestEventHandler)(nil).DeleteVirtualMeshCertificateSigningRequest), obj)
}

// GenericVirtualMeshCertificateSigningRequest mocks base method.
func (m *MockVirtualMeshCertificateSigningRequestEventHandler) GenericVirtualMeshCertificateSigningRequest(obj *v1alpha1.VirtualMeshCertificateSigningRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericVirtualMeshCertificateSigningRequest", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericVirtualMeshCertificateSigningRequest indicates an expected call of GenericVirtualMeshCertificateSigningRequest.
func (mr *MockVirtualMeshCertificateSigningRequestEventHandlerMockRecorder) GenericVirtualMeshCertificateSigningRequest(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericVirtualMeshCertificateSigningRequest", reflect.TypeOf((*MockVirtualMeshCertificateSigningRequestEventHandler)(nil).GenericVirtualMeshCertificateSigningRequest), obj)
}

// MockVirtualMeshCertificateSigningRequestEventWatcher is a mock of VirtualMeshCertificateSigningRequestEventWatcher interface.
type MockVirtualMeshCertificateSigningRequestEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualMeshCertificateSigningRequestEventWatcherMockRecorder
}

// MockVirtualMeshCertificateSigningRequestEventWatcherMockRecorder is the mock recorder for MockVirtualMeshCertificateSigningRequestEventWatcher.
type MockVirtualMeshCertificateSigningRequestEventWatcherMockRecorder struct {
	mock *MockVirtualMeshCertificateSigningRequestEventWatcher
}

// NewMockVirtualMeshCertificateSigningRequestEventWatcher creates a new mock instance.
func NewMockVirtualMeshCertificateSigningRequestEventWatcher(ctrl *gomock.Controller) *MockVirtualMeshCertificateSigningRequestEventWatcher {
	mock := &MockVirtualMeshCertificateSigningRequestEventWatcher{ctrl: ctrl}
	mock.recorder = &MockVirtualMeshCertificateSigningRequestEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualMeshCertificateSigningRequestEventWatcher) EXPECT() *MockVirtualMeshCertificateSigningRequestEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockVirtualMeshCertificateSigningRequestEventWatcher) AddEventHandler(ctx context.Context, h controller.VirtualMeshCertificateSigningRequestEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler.
func (mr *MockVirtualMeshCertificateSigningRequestEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockVirtualMeshCertificateSigningRequestEventWatcher)(nil).AddEventHandler), varargs...)
}
