// Code generated by MockGen. DO NOT EDIT.
// Source: ./meshworkload_translator.go

// Package mock_meshworkload is a generated GoMock package.
package mock_meshworkload

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1sets "github.com/solo-io/external-apis/pkg/api/k8s/apps/v1/sets"
	v1alpha2sets "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2/sets"
)

// MockTranslator is a mock of Translator interface
type MockTranslator struct {
	ctrl     *gomock.Controller
	recorder *MockTranslatorMockRecorder
}

// MockTranslatorMockRecorder is the mock recorder for MockTranslator
type MockTranslatorMockRecorder struct {
	mock *MockTranslator
}

// NewMockTranslator creates a new mock instance
func NewMockTranslator(ctrl *gomock.Controller) *MockTranslator {
	mock := &MockTranslator{ctrl: ctrl}
	mock.recorder = &MockTranslatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTranslator) EXPECT() *MockTranslatorMockRecorder {
	return m.recorder
}

// TranslateMeshWorkloads mocks base method
func (m *MockTranslator) TranslateMeshWorkloads(deployments v1sets.DeploymentSet, daemonSets v1sets.DaemonSetSet, statefulSets v1sets.StatefulSetSet, meshes v1alpha2sets.MeshSet) v1alpha2sets.MeshWorkloadSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TranslateMeshWorkloads", deployments, daemonSets, statefulSets, meshes)
	ret0, _ := ret[0].(v1alpha2sets.MeshWorkloadSet)
	return ret0
}

// TranslateMeshWorkloads indicates an expected call of TranslateMeshWorkloads
func (mr *MockTranslatorMockRecorder) TranslateMeshWorkloads(deployments, daemonSets, statefulSets, meshes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TranslateMeshWorkloads", reflect.TypeOf((*MockTranslator)(nil).TranslateMeshWorkloads), deployments, daemonSets, statefulSets, meshes)
}