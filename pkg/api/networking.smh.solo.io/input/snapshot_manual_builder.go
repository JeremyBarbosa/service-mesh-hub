// Code generated by skv2. DO NOT EDIT.

/*
	Utility for manually building input snapshots. Used primarily in tests.
*/
package input

import (
	discovery_smh_solo_io_v1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2"
	discovery_smh_solo_io_v1alpha2_sets "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2/sets"

	networking_smh_solo_io_v1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha2"
	networking_smh_solo_io_v1alpha2_sets "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha2/sets"

	multicluster_solo_io_v1alpha1 "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1"
	multicluster_solo_io_v1alpha1_sets "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1/sets"
)

type InputSnapshotManualBuilder struct {
	name string

	meshServices  discovery_smh_solo_io_v1alpha2_sets.MeshServiceSet
	meshWorkloads discovery_smh_solo_io_v1alpha2_sets.MeshWorkloadSet
	meshes        discovery_smh_solo_io_v1alpha2_sets.MeshSet

	trafficPolicies  networking_smh_solo_io_v1alpha2_sets.TrafficPolicySet
	accessPolicies   networking_smh_solo_io_v1alpha2_sets.AccessPolicySet
	virtualMeshes    networking_smh_solo_io_v1alpha2_sets.VirtualMeshSet
	failoverServices networking_smh_solo_io_v1alpha2_sets.FailoverServiceSet

	kubernetesClusters multicluster_solo_io_v1alpha1_sets.KubernetesClusterSet
}

func NewInputSnapshotManualBuilder(name string) *InputSnapshotManualBuilder {
	return &InputSnapshotManualBuilder{
		name: name,

		meshServices:  discovery_smh_solo_io_v1alpha2_sets.NewMeshServiceSet(),
		meshWorkloads: discovery_smh_solo_io_v1alpha2_sets.NewMeshWorkloadSet(),
		meshes:        discovery_smh_solo_io_v1alpha2_sets.NewMeshSet(),

		trafficPolicies:  networking_smh_solo_io_v1alpha2_sets.NewTrafficPolicySet(),
		accessPolicies:   networking_smh_solo_io_v1alpha2_sets.NewAccessPolicySet(),
		virtualMeshes:    networking_smh_solo_io_v1alpha2_sets.NewVirtualMeshSet(),
		failoverServices: networking_smh_solo_io_v1alpha2_sets.NewFailoverServiceSet(),

		kubernetesClusters: multicluster_solo_io_v1alpha1_sets.NewKubernetesClusterSet(),
	}
}

func (i *InputSnapshotManualBuilder) Build() Snapshot {
	return NewSnapshot(
		i.name,

		i.meshServices,
		i.meshWorkloads,
		i.meshes,

		i.trafficPolicies,
		i.accessPolicies,
		i.virtualMeshes,
		i.failoverServices,

		i.kubernetesClusters,
	)
}
func (i *InputSnapshotManualBuilder) AddMeshServices(meshServices []*discovery_smh_solo_io_v1alpha2.MeshService) *InputSnapshotManualBuilder {
	i.meshServices.Insert(meshServices...)
	return i
}
func (i *InputSnapshotManualBuilder) AddMeshWorkloads(meshWorkloads []*discovery_smh_solo_io_v1alpha2.MeshWorkload) *InputSnapshotManualBuilder {
	i.meshWorkloads.Insert(meshWorkloads...)
	return i
}
func (i *InputSnapshotManualBuilder) AddMeshes(meshes []*discovery_smh_solo_io_v1alpha2.Mesh) *InputSnapshotManualBuilder {
	i.meshes.Insert(meshes...)
	return i
}
func (i *InputSnapshotManualBuilder) AddTrafficPolicies(trafficPolicies []*networking_smh_solo_io_v1alpha2.TrafficPolicy) *InputSnapshotManualBuilder {
	i.trafficPolicies.Insert(trafficPolicies...)
	return i
}
func (i *InputSnapshotManualBuilder) AddAccessPolicies(accessPolicies []*networking_smh_solo_io_v1alpha2.AccessPolicy) *InputSnapshotManualBuilder {
	i.accessPolicies.Insert(accessPolicies...)
	return i
}
func (i *InputSnapshotManualBuilder) AddVirtualMeshes(virtualMeshes []*networking_smh_solo_io_v1alpha2.VirtualMesh) *InputSnapshotManualBuilder {
	i.virtualMeshes.Insert(virtualMeshes...)
	return i
}
func (i *InputSnapshotManualBuilder) AddFailoverServices(failoverServices []*networking_smh_solo_io_v1alpha2.FailoverService) *InputSnapshotManualBuilder {
	i.failoverServices.Insert(failoverServices...)
	return i
}
func (i *InputSnapshotManualBuilder) AddKubernetesClusters(kubernetesClusters []*multicluster_solo_io_v1alpha1.KubernetesCluster) *InputSnapshotManualBuilder {
	i.kubernetesClusters.Insert(kubernetesClusters...)
	return i
}