// Code generated by MockGen. DO NOT EDIT.
// Source: ./dependencies.go

// Package mock_internal is a generated GoMock package.
package mock_internal

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1sets "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/sets"
	v1alpha2sets "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2/sets"
	mesh "github.com/solo-io/service-mesh-hub/pkg/mesh-networking/translation/istio/mesh"
	meshservice "github.com/solo-io/service-mesh-hub/pkg/mesh-networking/translation/istio/meshservice"
	v1alpha1sets "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1/sets"
)

// MockDependencyFactory is a mock of DependencyFactory interface
type MockDependencyFactory struct {
	ctrl     *gomock.Controller
	recorder *MockDependencyFactoryMockRecorder
}

// MockDependencyFactoryMockRecorder is the mock recorder for MockDependencyFactory
type MockDependencyFactoryMockRecorder struct {
	mock *MockDependencyFactory
}

// NewMockDependencyFactory creates a new mock instance
func NewMockDependencyFactory(ctrl *gomock.Controller) *MockDependencyFactory {
	mock := &MockDependencyFactory{ctrl: ctrl}
	mock.recorder = &MockDependencyFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDependencyFactory) EXPECT() *MockDependencyFactoryMockRecorder {
	return m.recorder
}

// MakeMeshServiceTranslator mocks base method
func (m *MockDependencyFactory) MakeMeshServiceTranslator(clusters v1alpha1sets.KubernetesClusterSet, meshServices v1alpha2sets.MeshServiceSet) meshservice.Translator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeMeshServiceTranslator", clusters, meshServices)
	ret0, _ := ret[0].(meshservice.Translator)
	return ret0
}

// MakeMeshServiceTranslator indicates an expected call of MakeMeshServiceTranslator
func (mr *MockDependencyFactoryMockRecorder) MakeMeshServiceTranslator(clusters, meshServices interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeMeshServiceTranslator", reflect.TypeOf((*MockDependencyFactory)(nil).MakeMeshServiceTranslator), clusters, meshServices)
}

// MakeMeshTranslator mocks base method
func (m *MockDependencyFactory) MakeMeshTranslator(ctx context.Context, clusters v1alpha1sets.KubernetesClusterSet, secrets v1sets.SecretSet, meshWorkloads v1alpha2sets.MeshWorkloadSet, meshServices v1alpha2sets.MeshServiceSet) mesh.Translator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeMeshTranslator", ctx, clusters, secrets, meshWorkloads, meshServices)
	ret0, _ := ret[0].(mesh.Translator)
	return ret0
}

// MakeMeshTranslator indicates an expected call of MakeMeshTranslator
func (mr *MockDependencyFactoryMockRecorder) MakeMeshTranslator(ctx, clusters, secrets, meshWorkloads, meshServices interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeMeshTranslator", reflect.TypeOf((*MockDependencyFactory)(nil).MakeMeshTranslator), ctx, clusters, secrets, meshWorkloads, meshServices)
}