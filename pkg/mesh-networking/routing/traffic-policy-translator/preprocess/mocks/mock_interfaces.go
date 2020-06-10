// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_preprocess is a generated GoMock package.
package mock_preprocess

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1"
	v1alpha10 "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1"
	selection "github.com/solo-io/service-mesh-hub/pkg/common/kube/selection"
)

// MockTrafficPolicyPreprocessor is a mock of TrafficPolicyPreprocessor interface.
type MockTrafficPolicyPreprocessor struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficPolicyPreprocessorMockRecorder
}

// MockTrafficPolicyPreprocessorMockRecorder is the mock recorder for MockTrafficPolicyPreprocessor.
type MockTrafficPolicyPreprocessorMockRecorder struct {
	mock *MockTrafficPolicyPreprocessor
}

// NewMockTrafficPolicyPreprocessor creates a new mock instance.
func NewMockTrafficPolicyPreprocessor(ctrl *gomock.Controller) *MockTrafficPolicyPreprocessor {
	mock := &MockTrafficPolicyPreprocessor{ctrl: ctrl}
	mock.recorder = &MockTrafficPolicyPreprocessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficPolicyPreprocessor) EXPECT() *MockTrafficPolicyPreprocessorMockRecorder {
	return m.recorder
}

// PreprocessTrafficPolicy mocks base method.
func (m *MockTrafficPolicyPreprocessor) PreprocessTrafficPolicy(ctx context.Context, trafficPolicy *v1alpha10.TrafficPolicy) (map[selection.MeshServiceId][]*v1alpha10.TrafficPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PreprocessTrafficPolicy", ctx, trafficPolicy)
	ret0, _ := ret[0].(map[selection.MeshServiceId][]*v1alpha10.TrafficPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PreprocessTrafficPolicy indicates an expected call of PreprocessTrafficPolicy.
func (mr *MockTrafficPolicyPreprocessorMockRecorder) PreprocessTrafficPolicy(ctx, trafficPolicy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreprocessTrafficPolicy", reflect.TypeOf((*MockTrafficPolicyPreprocessor)(nil).PreprocessTrafficPolicy), ctx, trafficPolicy)
}

// PreprocessTrafficPoliciesForMeshService mocks base method.
func (m *MockTrafficPolicyPreprocessor) PreprocessTrafficPoliciesForMeshService(ctx context.Context, meshService *v1alpha1.MeshService) (map[selection.MeshServiceId][]*v1alpha10.TrafficPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PreprocessTrafficPoliciesForMeshService", ctx, meshService)
	ret0, _ := ret[0].(map[selection.MeshServiceId][]*v1alpha10.TrafficPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PreprocessTrafficPoliciesForMeshService indicates an expected call of PreprocessTrafficPoliciesForMeshService.
func (mr *MockTrafficPolicyPreprocessorMockRecorder) PreprocessTrafficPoliciesForMeshService(ctx, meshService interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreprocessTrafficPoliciesForMeshService", reflect.TypeOf((*MockTrafficPolicyPreprocessor)(nil).PreprocessTrafficPoliciesForMeshService), ctx, meshService)
}

// MockTrafficPolicyMerger is a mock of TrafficPolicyMerger interface.
type MockTrafficPolicyMerger struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficPolicyMergerMockRecorder
}

// MockTrafficPolicyMergerMockRecorder is the mock recorder for MockTrafficPolicyMerger.
type MockTrafficPolicyMergerMockRecorder struct {
	mock *MockTrafficPolicyMerger
}

// NewMockTrafficPolicyMerger creates a new mock instance.
func NewMockTrafficPolicyMerger(ctrl *gomock.Controller) *MockTrafficPolicyMerger {
	mock := &MockTrafficPolicyMerger{ctrl: ctrl}
	mock.recorder = &MockTrafficPolicyMergerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficPolicyMerger) EXPECT() *MockTrafficPolicyMergerMockRecorder {
	return m.recorder
}

// MergeTrafficPoliciesForMeshServices mocks base method.
func (m *MockTrafficPolicyMerger) MergeTrafficPoliciesForMeshServices(ctx context.Context, meshServices []*v1alpha1.MeshService) (map[selection.MeshServiceId][]*v1alpha10.TrafficPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MergeTrafficPoliciesForMeshServices", ctx, meshServices)
	ret0, _ := ret[0].(map[selection.MeshServiceId][]*v1alpha10.TrafficPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MergeTrafficPoliciesForMeshServices indicates an expected call of MergeTrafficPoliciesForMeshServices.
func (mr *MockTrafficPolicyMergerMockRecorder) MergeTrafficPoliciesForMeshServices(ctx, meshServices interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MergeTrafficPoliciesForMeshServices", reflect.TypeOf((*MockTrafficPolicyMerger)(nil).MergeTrafficPoliciesForMeshServices), ctx, meshServices)
}

// MockTrafficPolicyValidator is a mock of TrafficPolicyValidator interface.
type MockTrafficPolicyValidator struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficPolicyValidatorMockRecorder
}

// MockTrafficPolicyValidatorMockRecorder is the mock recorder for MockTrafficPolicyValidator.
type MockTrafficPolicyValidatorMockRecorder struct {
	mock *MockTrafficPolicyValidator
}

// NewMockTrafficPolicyValidator creates a new mock instance.
func NewMockTrafficPolicyValidator(ctrl *gomock.Controller) *MockTrafficPolicyValidator {
	mock := &MockTrafficPolicyValidator{ctrl: ctrl}
	mock.recorder = &MockTrafficPolicyValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficPolicyValidator) EXPECT() *MockTrafficPolicyValidatorMockRecorder {
	return m.recorder
}

// Validate mocks base method.
func (m *MockTrafficPolicyValidator) Validate(ctx context.Context, trafficPolicy *v1alpha10.TrafficPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", ctx, trafficPolicy)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockTrafficPolicyValidatorMockRecorder) Validate(ctx, trafficPolicy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockTrafficPolicyValidator)(nil).Validate), ctx, trafficPolicy)
}