// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/mesh-projects/api/config/v1alpha1/policy.proto

package v1alpha1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/mesh-projects/pkg/api/core.zephyr.solo.io/v1alpha1"
	_ "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for SecurityRuleSpec
func (this *SecurityRuleSpec) MarshalJSON() ([]byte, error) {
	str, err := PolicyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for SecurityRuleSpec
func (this *SecurityRuleSpec) UnmarshalJSON(b []byte) error {
	return PolicyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for SecurityRuleStatus
func (this *SecurityRuleStatus) MarshalJSON() ([]byte, error) {
	str, err := PolicyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for SecurityRuleStatus
func (this *SecurityRuleStatus) UnmarshalJSON(b []byte) error {
	return PolicyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	PolicyMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	PolicyUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)