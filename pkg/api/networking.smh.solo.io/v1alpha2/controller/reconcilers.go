// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./reconcilers.go -destination mocks/reconcilers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	networking_smh_solo_io_v1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha2"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the TrafficPolicy Resource.
// implemented by the user
type TrafficPolicyReconciler interface {
	ReconcileTrafficPolicy(obj *networking_smh_solo_io_v1alpha2.TrafficPolicy) (reconcile.Result, error)
}

// Reconcile deletion events for the TrafficPolicy Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type TrafficPolicyDeletionReconciler interface {
	ReconcileTrafficPolicyDeletion(req reconcile.Request) error
}

type TrafficPolicyReconcilerFuncs struct {
	OnReconcileTrafficPolicy         func(obj *networking_smh_solo_io_v1alpha2.TrafficPolicy) (reconcile.Result, error)
	OnReconcileTrafficPolicyDeletion func(req reconcile.Request) error
}

func (f *TrafficPolicyReconcilerFuncs) ReconcileTrafficPolicy(obj *networking_smh_solo_io_v1alpha2.TrafficPolicy) (reconcile.Result, error) {
	if f.OnReconcileTrafficPolicy == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileTrafficPolicy(obj)
}

func (f *TrafficPolicyReconcilerFuncs) ReconcileTrafficPolicyDeletion(req reconcile.Request) error {
	if f.OnReconcileTrafficPolicyDeletion == nil {
		return nil
	}
	return f.OnReconcileTrafficPolicyDeletion(req)
}

// Reconcile and finalize the TrafficPolicy Resource
// implemented by the user
type TrafficPolicyFinalizer interface {
	TrafficPolicyReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	TrafficPolicyFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeTrafficPolicy(obj *networking_smh_solo_io_v1alpha2.TrafficPolicy) error
}

type TrafficPolicyReconcileLoop interface {
	RunTrafficPolicyReconciler(ctx context.Context, rec TrafficPolicyReconciler, predicates ...predicate.Predicate) error
}

type trafficPolicyReconcileLoop struct {
	loop reconcile.Loop
}

func NewTrafficPolicyReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) TrafficPolicyReconcileLoop {
	return &trafficPolicyReconcileLoop{
		loop: reconcile.NewLoop(name, mgr, &networking_smh_solo_io_v1alpha2.TrafficPolicy{}, options),
	}
}

func (c *trafficPolicyReconcileLoop) RunTrafficPolicyReconciler(ctx context.Context, reconciler TrafficPolicyReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericTrafficPolicyReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(TrafficPolicyFinalizer); ok {
		reconcilerWrapper = genericTrafficPolicyFinalizer{
			genericTrafficPolicyReconciler: genericReconciler,
			finalizingReconciler:           finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericTrafficPolicyHandler implements a generic reconcile.Reconciler
type genericTrafficPolicyReconciler struct {
	reconciler TrafficPolicyReconciler
}

func (r genericTrafficPolicyReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_smh_solo_io_v1alpha2.TrafficPolicy)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: TrafficPolicy handler received event for %T", object)
	}
	return r.reconciler.ReconcileTrafficPolicy(obj)
}

func (r genericTrafficPolicyReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(TrafficPolicyDeletionReconciler); ok {
		return deletionReconciler.ReconcileTrafficPolicyDeletion(request)
	}
	return nil
}

// genericTrafficPolicyFinalizer implements a generic reconcile.FinalizingReconciler
type genericTrafficPolicyFinalizer struct {
	genericTrafficPolicyReconciler
	finalizingReconciler TrafficPolicyFinalizer
}

func (r genericTrafficPolicyFinalizer) FinalizerName() string {
	return r.finalizingReconciler.TrafficPolicyFinalizerName()
}

func (r genericTrafficPolicyFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*networking_smh_solo_io_v1alpha2.TrafficPolicy)
	if !ok {
		return errors.Errorf("internal error: TrafficPolicy handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeTrafficPolicy(obj)
}

// Reconcile Upsert events for the AccessPolicy Resource.
// implemented by the user
type AccessPolicyReconciler interface {
	ReconcileAccessPolicy(obj *networking_smh_solo_io_v1alpha2.AccessPolicy) (reconcile.Result, error)
}

// Reconcile deletion events for the AccessPolicy Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type AccessPolicyDeletionReconciler interface {
	ReconcileAccessPolicyDeletion(req reconcile.Request) error
}

type AccessPolicyReconcilerFuncs struct {
	OnReconcileAccessPolicy         func(obj *networking_smh_solo_io_v1alpha2.AccessPolicy) (reconcile.Result, error)
	OnReconcileAccessPolicyDeletion func(req reconcile.Request) error
}

func (f *AccessPolicyReconcilerFuncs) ReconcileAccessPolicy(obj *networking_smh_solo_io_v1alpha2.AccessPolicy) (reconcile.Result, error) {
	if f.OnReconcileAccessPolicy == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileAccessPolicy(obj)
}

func (f *AccessPolicyReconcilerFuncs) ReconcileAccessPolicyDeletion(req reconcile.Request) error {
	if f.OnReconcileAccessPolicyDeletion == nil {
		return nil
	}
	return f.OnReconcileAccessPolicyDeletion(req)
}

// Reconcile and finalize the AccessPolicy Resource
// implemented by the user
type AccessPolicyFinalizer interface {
	AccessPolicyReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	AccessPolicyFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeAccessPolicy(obj *networking_smh_solo_io_v1alpha2.AccessPolicy) error
}

type AccessPolicyReconcileLoop interface {
	RunAccessPolicyReconciler(ctx context.Context, rec AccessPolicyReconciler, predicates ...predicate.Predicate) error
}

type accessPolicyReconcileLoop struct {
	loop reconcile.Loop
}

func NewAccessPolicyReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) AccessPolicyReconcileLoop {
	return &accessPolicyReconcileLoop{
		loop: reconcile.NewLoop(name, mgr, &networking_smh_solo_io_v1alpha2.AccessPolicy{}, options),
	}
}

func (c *accessPolicyReconcileLoop) RunAccessPolicyReconciler(ctx context.Context, reconciler AccessPolicyReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericAccessPolicyReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(AccessPolicyFinalizer); ok {
		reconcilerWrapper = genericAccessPolicyFinalizer{
			genericAccessPolicyReconciler: genericReconciler,
			finalizingReconciler:          finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericAccessPolicyHandler implements a generic reconcile.Reconciler
type genericAccessPolicyReconciler struct {
	reconciler AccessPolicyReconciler
}

func (r genericAccessPolicyReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_smh_solo_io_v1alpha2.AccessPolicy)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: AccessPolicy handler received event for %T", object)
	}
	return r.reconciler.ReconcileAccessPolicy(obj)
}

func (r genericAccessPolicyReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(AccessPolicyDeletionReconciler); ok {
		return deletionReconciler.ReconcileAccessPolicyDeletion(request)
	}
	return nil
}

// genericAccessPolicyFinalizer implements a generic reconcile.FinalizingReconciler
type genericAccessPolicyFinalizer struct {
	genericAccessPolicyReconciler
	finalizingReconciler AccessPolicyFinalizer
}

func (r genericAccessPolicyFinalizer) FinalizerName() string {
	return r.finalizingReconciler.AccessPolicyFinalizerName()
}

func (r genericAccessPolicyFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*networking_smh_solo_io_v1alpha2.AccessPolicy)
	if !ok {
		return errors.Errorf("internal error: AccessPolicy handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeAccessPolicy(obj)
}

// Reconcile Upsert events for the VirtualMesh Resource.
// implemented by the user
type VirtualMeshReconciler interface {
	ReconcileVirtualMesh(obj *networking_smh_solo_io_v1alpha2.VirtualMesh) (reconcile.Result, error)
}

// Reconcile deletion events for the VirtualMesh Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type VirtualMeshDeletionReconciler interface {
	ReconcileVirtualMeshDeletion(req reconcile.Request) error
}

type VirtualMeshReconcilerFuncs struct {
	OnReconcileVirtualMesh         func(obj *networking_smh_solo_io_v1alpha2.VirtualMesh) (reconcile.Result, error)
	OnReconcileVirtualMeshDeletion func(req reconcile.Request) error
}

func (f *VirtualMeshReconcilerFuncs) ReconcileVirtualMesh(obj *networking_smh_solo_io_v1alpha2.VirtualMesh) (reconcile.Result, error) {
	if f.OnReconcileVirtualMesh == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileVirtualMesh(obj)
}

func (f *VirtualMeshReconcilerFuncs) ReconcileVirtualMeshDeletion(req reconcile.Request) error {
	if f.OnReconcileVirtualMeshDeletion == nil {
		return nil
	}
	return f.OnReconcileVirtualMeshDeletion(req)
}

// Reconcile and finalize the VirtualMesh Resource
// implemented by the user
type VirtualMeshFinalizer interface {
	VirtualMeshReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	VirtualMeshFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeVirtualMesh(obj *networking_smh_solo_io_v1alpha2.VirtualMesh) error
}

type VirtualMeshReconcileLoop interface {
	RunVirtualMeshReconciler(ctx context.Context, rec VirtualMeshReconciler, predicates ...predicate.Predicate) error
}

type virtualMeshReconcileLoop struct {
	loop reconcile.Loop
}

func NewVirtualMeshReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) VirtualMeshReconcileLoop {
	return &virtualMeshReconcileLoop{
		loop: reconcile.NewLoop(name, mgr, &networking_smh_solo_io_v1alpha2.VirtualMesh{}, options),
	}
}

func (c *virtualMeshReconcileLoop) RunVirtualMeshReconciler(ctx context.Context, reconciler VirtualMeshReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericVirtualMeshReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(VirtualMeshFinalizer); ok {
		reconcilerWrapper = genericVirtualMeshFinalizer{
			genericVirtualMeshReconciler: genericReconciler,
			finalizingReconciler:         finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericVirtualMeshHandler implements a generic reconcile.Reconciler
type genericVirtualMeshReconciler struct {
	reconciler VirtualMeshReconciler
}

func (r genericVirtualMeshReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_smh_solo_io_v1alpha2.VirtualMesh)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: VirtualMesh handler received event for %T", object)
	}
	return r.reconciler.ReconcileVirtualMesh(obj)
}

func (r genericVirtualMeshReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(VirtualMeshDeletionReconciler); ok {
		return deletionReconciler.ReconcileVirtualMeshDeletion(request)
	}
	return nil
}

// genericVirtualMeshFinalizer implements a generic reconcile.FinalizingReconciler
type genericVirtualMeshFinalizer struct {
	genericVirtualMeshReconciler
	finalizingReconciler VirtualMeshFinalizer
}

func (r genericVirtualMeshFinalizer) FinalizerName() string {
	return r.finalizingReconciler.VirtualMeshFinalizerName()
}

func (r genericVirtualMeshFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*networking_smh_solo_io_v1alpha2.VirtualMesh)
	if !ok {
		return errors.Errorf("internal error: VirtualMesh handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeVirtualMesh(obj)
}