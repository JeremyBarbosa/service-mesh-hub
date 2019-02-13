// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/supergloo/api/v1/policy.proto

package v1 // import "github.com/solo-io/supergloo/pkg2/api/v1"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Policy struct {
	Rules                []*Rule  `protobuf:"bytes,2,rep,name=rules" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_5e42b8cbb1d1d488, []int{0}
}
func (m *Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy.Unmarshal(m, b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
}
func (dst *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(dst, src)
}
func (m *Policy) XXX_Size() int {
	return xxx_messageInfo_Policy.Size(m)
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetRules() []*Rule {
	if m != nil {
		return m.Rules
	}
	return nil
}

type Rule struct {
	Source               *core.ResourceRef `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	Destination          *core.ResourceRef `protobuf:"bytes,2,opt,name=destination" json:"destination,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Rule) Reset()         { *m = Rule{} }
func (m *Rule) String() string { return proto.CompactTextString(m) }
func (*Rule) ProtoMessage()    {}
func (*Rule) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_5e42b8cbb1d1d488, []int{1}
}
func (m *Rule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rule.Unmarshal(m, b)
}
func (m *Rule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rule.Marshal(b, m, deterministic)
}
func (dst *Rule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rule.Merge(dst, src)
}
func (m *Rule) XXX_Size() int {
	return xxx_messageInfo_Rule.Size(m)
}
func (m *Rule) XXX_DiscardUnknown() {
	xxx_messageInfo_Rule.DiscardUnknown(m)
}

var xxx_messageInfo_Rule proto.InternalMessageInfo

func (m *Rule) GetSource() *core.ResourceRef {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *Rule) GetDestination() *core.ResourceRef {
	if m != nil {
		return m.Destination
	}
	return nil
}

func init() {
	proto.RegisterType((*Policy)(nil), "supergloo.solo.io.Policy")
	proto.RegisterType((*Rule)(nil), "supergloo.solo.io.Rule")
}
func (this *Policy) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Policy)
	if !ok {
		that2, ok := that.(Policy)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Rules) != len(that1.Rules) {
		return false
	}
	for i := range this.Rules {
		if !this.Rules[i].Equal(that1.Rules[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Rule) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Rule)
	if !ok {
		that2, ok := that.(Rule)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Source.Equal(that1.Source) {
		return false
	}
	if !this.Destination.Equal(that1.Destination) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

func init() {
	proto.RegisterFile("github.com/solo-io/supergloo/api/v1/policy.proto", fileDescriptor_policy_5e42b8cbb1d1d488)
}

var fileDescriptor_policy_5e42b8cbb1d1d488 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x48, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xce, 0xcf, 0xc9, 0xd7, 0xcd, 0xcc, 0xd7, 0x2f,
	0x2e, 0x2d, 0x48, 0x2d, 0x4a, 0xcf, 0xc9, 0xcf, 0xd7, 0x4f, 0x2c, 0xc8, 0xd4, 0x2f, 0x33, 0xd4,
	0x2f, 0xc8, 0xcf, 0xc9, 0x4c, 0xae, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x84, 0x4b,
	0xeb, 0x81, 0x34, 0xe8, 0x65, 0xe6, 0x4b, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x65, 0xf5, 0x41,
	0x2c, 0x88, 0x42, 0x29, 0x1d, 0x6c, 0x46, 0x83, 0xe8, 0xec, 0xcc, 0x12, 0x98, 0xc9, 0x45, 0xa9,
	0x69, 0x10, 0xd5, 0x4a, 0xe6, 0x5c, 0x6c, 0x01, 0x60, 0x6b, 0x84, 0x74, 0xb9, 0x58, 0x8b, 0x4a,
	0x73, 0x52, 0x8b, 0x25, 0x98, 0x14, 0x98, 0x35, 0xb8, 0x8d, 0xc4, 0xf5, 0x30, 0x2c, 0xd4, 0x0b,
	0x2a, 0xcd, 0x49, 0x0d, 0x82, 0xa8, 0x52, 0x2a, 0xe3, 0x62, 0x01, 0x71, 0x85, 0x0c, 0xb9, 0xd8,
	0x8a, 0xf3, 0x4b, 0x8b, 0x92, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x24, 0xf5, 0x92,
	0xf3, 0x8b, 0x52, 0x11, 0x5a, 0x52, 0x21, 0xb2, 0x41, 0xa9, 0x69, 0x41, 0x50, 0x85, 0x42, 0xd6,
	0x5c, 0xdc, 0x29, 0xa9, 0xc5, 0x25, 0x99, 0x79, 0x89, 0x25, 0x99, 0xf9, 0x79, 0x12, 0x4c, 0x84,
	0xf4, 0x21, 0xab, 0x76, 0xd2, 0x5d, 0xf1, 0x48, 0x8e, 0x31, 0x4a, 0x1d, 0x6f, 0xf8, 0x15, 0x64,
	0xa7, 0x43, 0x7d, 0x9a, 0xc4, 0x06, 0xf6, 0xa6, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x20, 0x41,
	0x1a, 0x1a, 0x71, 0x01, 0x00, 0x00,
}
