// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/networking/v1alpha2/http.proto

package types

import (
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type HttpMethodValue int32

const (
	HttpMethodValue_GET     HttpMethodValue = 0
	HttpMethodValue_POST    HttpMethodValue = 1
	HttpMethodValue_PUT     HttpMethodValue = 2
	HttpMethodValue_DELETE  HttpMethodValue = 3
	HttpMethodValue_HEAD    HttpMethodValue = 4
	HttpMethodValue_CONNECT HttpMethodValue = 5
	HttpMethodValue_OPTIONS HttpMethodValue = 6
	HttpMethodValue_TRACE   HttpMethodValue = 7
	HttpMethodValue_PATCH   HttpMethodValue = 8
)

var HttpMethodValue_name = map[int32]string{
	0: "GET",
	1: "POST",
	2: "PUT",
	3: "DELETE",
	4: "HEAD",
	5: "CONNECT",
	6: "OPTIONS",
	7: "TRACE",
	8: "PATCH",
}

var HttpMethodValue_value = map[string]int32{
	"GET":     0,
	"POST":    1,
	"PUT":     2,
	"DELETE":  3,
	"HEAD":    4,
	"CONNECT": 5,
	"OPTIONS": 6,
	"TRACE":   7,
	"PATCH":   8,
}

func (x HttpMethodValue) String() string {
	return proto.EnumName(HttpMethodValue_name, int32(x))
}

func (HttpMethodValue) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_910e869d37f284f8, []int{0}
}

func init() {
	proto.RegisterEnum("networking.smh.solo.io.HttpMethodValue", HttpMethodValue_name, HttpMethodValue_value)
}

func init() {
	proto.RegisterFile("github.com/solo-io/service-mesh-hub/api/networking/v1alpha2/http.proto", fileDescriptor_910e869d37f284f8)
}

var fileDescriptor_910e869d37f284f8 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xcf, 0xc1, 0x4e, 0x84, 0x30,
	0x10, 0x06, 0x60, 0x75, 0x77, 0x61, 0xad, 0x07, 0x9b, 0x1e, 0x7c, 0x08, 0x13, 0xda, 0xa8, 0x4f,
	0x80, 0x6c, 0x15, 0x13, 0x05, 0x94, 0xea, 0xc1, 0x1b, 0xac, 0x0d, 0x6d, 0x16, 0x76, 0x1a, 0x3a,
	0xac, 0xf1, 0xed, 0x4d, 0xbd, 0x60, 0x3c, 0x79, 0x9b, 0x99, 0x3f, 0x99, 0xfc, 0x1f, 0xb9, 0xeb,
	0x2c, 0x9a, 0xa9, 0xe5, 0x5b, 0x18, 0x84, 0x87, 0x1e, 0x12, 0x0b, 0xc2, 0xeb, 0xf1, 0x60, 0xb7,
	0x3a, 0x19, 0xb4, 0x37, 0x89, 0x99, 0x5a, 0xd1, 0x38, 0x2b, 0xf6, 0x1a, 0x3f, 0x61, 0xdc, 0xd9,
	0x7d, 0x27, 0x0e, 0x57, 0x4d, 0xef, 0x4c, 0x73, 0x2d, 0x0c, 0xa2, 0xe3, 0x6e, 0x04, 0x04, 0x76,
	0x31, 0xe7, 0xdc, 0x0f, 0x86, 0x87, 0x5f, 0xdc, 0xc2, 0xa5, 0x27, 0xe7, 0x39, 0xa2, 0x7b, 0xd2,
	0x68, 0xe0, 0xe3, 0xad, 0xe9, 0x27, 0xcd, 0x62, 0xb2, 0xb8, 0x97, 0x8a, 0x1e, 0xb1, 0x35, 0x59,
	0x56, 0x65, 0xad, 0xe8, 0x71, 0x38, 0x55, 0xaf, 0x8a, 0x9e, 0x30, 0x42, 0xa2, 0x8d, 0x7c, 0x94,
	0x4a, 0xd2, 0x45, 0x88, 0x73, 0x99, 0x6e, 0xe8, 0x92, 0x9d, 0x91, 0x38, 0x2b, 0x8b, 0x42, 0x66,
	0x8a, 0xae, 0xc2, 0x52, 0x56, 0xea, 0xa1, 0x2c, 0x6a, 0x1a, 0xb1, 0x53, 0xb2, 0x52, 0x2f, 0x69,
	0x26, 0x69, 0x1c, 0xc6, 0x2a, 0x55, 0x59, 0x4e, 0xd7, 0xb7, 0xf5, 0xfb, 0xf3, 0x7f, 0x58, 0x6e,
	0xd7, 0xfd, 0xa1, 0xfd, 0xae, 0x3e, 0x33, 0xf1, 0xcb, 0x69, 0xdf, 0x46, 0x3f, 0xd0, 0x9b, 0xef,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xbe, 0xab, 0x15, 0x3b, 0x32, 0x01, 0x00, 0x00,
}