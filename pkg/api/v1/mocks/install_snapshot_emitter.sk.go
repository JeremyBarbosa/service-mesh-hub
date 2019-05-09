// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/api/v1/install_snapshot_emitter.sk.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	clients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
	v1 "github.com/solo-io/supergloo/pkg/api/v1"
)

// MockInstallEmitter is a mock of InstallEmitter interface
type MockInstallEmitter struct {
	ctrl     *gomock.Controller
	recorder *MockInstallEmitterMockRecorder
}

// MockInstallEmitterMockRecorder is the mock recorder for MockInstallEmitter
type MockInstallEmitterMockRecorder struct {
	mock *MockInstallEmitter
}

// NewMockInstallEmitter creates a new mock instance
func NewMockInstallEmitter(ctrl *gomock.Controller) *MockInstallEmitter {
	mock := &MockInstallEmitter{ctrl: ctrl}
	mock.recorder = &MockInstallEmitterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInstallEmitter) EXPECT() *MockInstallEmitterMockRecorder {
	return m.recorder
}

// Register mocks base method
func (m *MockInstallEmitter) Register() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register")
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockInstallEmitterMockRecorder) Register() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockInstallEmitter)(nil).Register))
}

// Install mocks base method
func (m *MockInstallEmitter) Install() v1.InstallClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Install")
	ret0, _ := ret[0].(v1.InstallClient)
	return ret0
}

// Install indicates an expected call of Install
func (mr *MockInstallEmitterMockRecorder) Install() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockInstallEmitter)(nil).Install))
}

// Mesh mocks base method
func (m *MockInstallEmitter) Mesh() v1.MeshClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mesh")
	ret0, _ := ret[0].(v1.MeshClient)
	return ret0
}

// Mesh indicates an expected call of Mesh
func (mr *MockInstallEmitterMockRecorder) Mesh() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mesh", reflect.TypeOf((*MockInstallEmitter)(nil).Mesh))
}

// MeshIngress mocks base method
func (m *MockInstallEmitter) MeshIngress() v1.MeshIngressClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MeshIngress")
	ret0, _ := ret[0].(v1.MeshIngressClient)
	return ret0
}

// MeshIngress indicates an expected call of MeshIngress
func (mr *MockInstallEmitterMockRecorder) MeshIngress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MeshIngress", reflect.TypeOf((*MockInstallEmitter)(nil).MeshIngress))
}

// Snapshots mocks base method
func (m *MockInstallEmitter) Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *v1.InstallSnapshot, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshots", watchNamespaces, opts)
	ret0, _ := ret[0].(<-chan *v1.InstallSnapshot)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Snapshots indicates an expected call of Snapshots
func (mr *MockInstallEmitterMockRecorder) Snapshots(watchNamespaces, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshots", reflect.TypeOf((*MockInstallEmitter)(nil).Snapshots), watchNamespaces, opts)
}