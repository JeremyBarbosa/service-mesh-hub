// Code generated by skv2. DO NOT EDIT.

/*
	Utility for manually building input snapshots. Used primarily in tests.
*/
package input

import (
	certificates_smh_solo_io_v1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/certificates.smh.solo.io/v1alpha2"
	certificates_smh_solo_io_v1alpha2_sets "github.com/solo-io/service-mesh-hub/pkg/api/certificates.smh.solo.io/v1alpha2/sets"

	v1_sets "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/sets"
	v1 "k8s.io/api/core/v1"
)

type InputSnapshotManualBuilder struct {
	name string

	issuedCertificates  certificates_smh_solo_io_v1alpha2_sets.IssuedCertificateSet
	certificateRequests certificates_smh_solo_io_v1alpha2_sets.CertificateRequestSet

	secrets v1_sets.SecretSet
	pods    v1_sets.PodSet
}

func NewInputSnapshotManualBuilder(name string) *InputSnapshotManualBuilder {
	return &InputSnapshotManualBuilder{
		name: name,

		issuedCertificates:  certificates_smh_solo_io_v1alpha2_sets.NewIssuedCertificateSet(),
		certificateRequests: certificates_smh_solo_io_v1alpha2_sets.NewCertificateRequestSet(),

		secrets: v1_sets.NewSecretSet(),
		pods:    v1_sets.NewPodSet(),
	}
}

func (i *InputSnapshotManualBuilder) Build() Snapshot {
	return NewSnapshot(
		i.name,

		i.issuedCertificates,
		i.certificateRequests,

		i.secrets,
		i.pods,
	)
}
func (i *InputSnapshotManualBuilder) AddIssuedCertificates(issuedCertificates []*certificates_smh_solo_io_v1alpha2.IssuedCertificate) *InputSnapshotManualBuilder {
	i.issuedCertificates.Insert(issuedCertificates...)
	return i
}
func (i *InputSnapshotManualBuilder) AddCertificateRequests(certificateRequests []*certificates_smh_solo_io_v1alpha2.CertificateRequest) *InputSnapshotManualBuilder {
	i.certificateRequests.Insert(certificateRequests...)
	return i
}
func (i *InputSnapshotManualBuilder) AddSecrets(secrets []*v1.Secret) *InputSnapshotManualBuilder {
	i.secrets.Insert(secrets...)
	return i
}
func (i *InputSnapshotManualBuilder) AddPods(pods []*v1.Pod) *InputSnapshotManualBuilder {
	i.pods.Insert(pods...)
	return i
}