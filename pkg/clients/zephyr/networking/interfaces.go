package zephyr_networking

import (
	"context"

	networkingv1alpha1 "github.com/solo-io/mesh-projects/pkg/api/networking.zephyr.solo.io/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

//go:generate mockgen -source ./interfaces.go -destination mocks/mock_clients.go

type VirtualMeshClient interface {
	Get(ctx context.Context, name, namespace string) (*networkingv1alpha1.VirtualMesh, error)
	List(ctx context.Context, opts ...client.ListOption) (*networkingv1alpha1.VirtualMeshList, error)
	UpdateStatus(ctx context.Context, virtualMesh *networkingv1alpha1.VirtualMesh, opts ...client.UpdateOption) error
}

type TrafficPolicyClient interface {
	Get(ctx context.Context, name string, namespace string) (*networkingv1alpha1.TrafficPolicy, error)
	Create(ctx context.Context, trafficPolicy *networkingv1alpha1.TrafficPolicy, options ...client.CreateOption) error
	Update(ctx context.Context, trafficPolicy *networkingv1alpha1.TrafficPolicy, options ...client.UpdateOption) error
	UpdateStatus(ctx context.Context, trafficPolicy *networkingv1alpha1.TrafficPolicy, options ...client.UpdateOption) error
	List(ctx context.Context, opts ...client.ListOption) (*networkingv1alpha1.TrafficPolicyList, error)
}