// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/mesh-projects/api/security/v1alpha1/certificates.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/mesh-projects/pkg/api/core.zephyr.solo.io/v1alpha1/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for VirtualMeshCertificateSigningRequestSpec
func (this *VirtualMeshCertificateSigningRequestSpec) MarshalJSON() ([]byte, error) {
	str, err := CertificatesMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for VirtualMeshCertificateSigningRequestSpec
func (this *VirtualMeshCertificateSigningRequestSpec) UnmarshalJSON(b []byte) error {
	return CertificatesUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for CertConfig
func (this *CertConfig) MarshalJSON() ([]byte, error) {
	str, err := CertificatesMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for CertConfig
func (this *CertConfig) UnmarshalJSON(b []byte) error {
	return CertificatesUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for VirtualMeshCertificateSigningResponse
func (this *VirtualMeshCertificateSigningResponse) MarshalJSON() ([]byte, error) {
	str, err := CertificatesMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for VirtualMeshCertificateSigningResponse
func (this *VirtualMeshCertificateSigningResponse) UnmarshalJSON(b []byte) error {
	return CertificatesUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ThirdPartyApprovalWorkflow
func (this *ThirdPartyApprovalWorkflow) MarshalJSON() ([]byte, error) {
	str, err := CertificatesMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ThirdPartyApprovalWorkflow
func (this *ThirdPartyApprovalWorkflow) UnmarshalJSON(b []byte) error {
	return CertificatesUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for VirtualMeshCertificateSigningRequestStatus
func (this *VirtualMeshCertificateSigningRequestStatus) MarshalJSON() ([]byte, error) {
	str, err := CertificatesMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for VirtualMeshCertificateSigningRequestStatus
func (this *VirtualMeshCertificateSigningRequestStatus) UnmarshalJSON(b []byte) error {
	return CertificatesUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	CertificatesMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	CertificatesUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)