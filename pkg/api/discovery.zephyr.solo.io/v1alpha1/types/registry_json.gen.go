// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/mesh-projects/api/discovery/v1alpha1/registry.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/mesh-projects/pkg/api/core.zephyr.solo.io/v1alpha1/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for MeshServiceSpec
func (this *MeshServiceSpec) MarshalJSON() ([]byte, error) {
	str, err := RegistryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MeshServiceSpec
func (this *MeshServiceSpec) UnmarshalJSON(b []byte) error {
	return RegistryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for MeshServiceSpec_Subset
func (this *MeshServiceSpec_Subset) MarshalJSON() ([]byte, error) {
	str, err := RegistryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MeshServiceSpec_Subset
func (this *MeshServiceSpec_Subset) UnmarshalJSON(b []byte) error {
	return RegistryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Federation
func (this *Federation) MarshalJSON() ([]byte, error) {
	str, err := RegistryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Federation
func (this *Federation) UnmarshalJSON(b []byte) error {
	return RegistryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for KubeService
func (this *KubeService) MarshalJSON() ([]byte, error) {
	str, err := RegistryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for KubeService
func (this *KubeService) UnmarshalJSON(b []byte) error {
	return RegistryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for KubeServicePort
func (this *KubeServicePort) MarshalJSON() ([]byte, error) {
	str, err := RegistryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for KubeServicePort
func (this *KubeServicePort) UnmarshalJSON(b []byte) error {
	return RegistryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for MeshServiceStatus
func (this *MeshServiceStatus) MarshalJSON() ([]byte, error) {
	str, err := RegistryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MeshServiceStatus
func (this *MeshServiceStatus) UnmarshalJSON(b []byte) error {
	return RegistryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for MeshWorkloadSpec
func (this *MeshWorkloadSpec) MarshalJSON() ([]byte, error) {
	str, err := RegistryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MeshWorkloadSpec
func (this *MeshWorkloadSpec) UnmarshalJSON(b []byte) error {
	return RegistryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for KubePod
func (this *KubePod) MarshalJSON() ([]byte, error) {
	str, err := RegistryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for KubePod
func (this *KubePod) UnmarshalJSON(b []byte) error {
	return RegistryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for MeshWorkloadStatus
func (this *MeshWorkloadStatus) MarshalJSON() ([]byte, error) {
	str, err := RegistryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MeshWorkloadStatus
func (this *MeshWorkloadStatus) UnmarshalJSON(b []byte) error {
	return RegistryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	RegistryMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	RegistryUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)