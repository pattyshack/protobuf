// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: combos/neither/map.proto

/*
Package mapdefaults is a generated protocol buffer package.

It is generated from these files:
	combos/neither/map.proto

It has these top-level messages:
	MapTest
*/
package mapdefaults

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import github_com_gogo_protobuf_protoc_gen_gogo_descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import compress_gzip "compress/gzip"
import bytes "bytes"
import io_ioutil "io/ioutil"

import strings "strings"
import reflect "reflect"
import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type MapTest struct {
	StrStr map[string]string `protobuf:"bytes,1,rep,name=str_str,json=strStr" json:"str_str,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *MapTest) Reset()                    { *m = MapTest{} }
func (*MapTest) ProtoMessage()               {}
func (*MapTest) Descriptor() ([]byte, []int) { return fileDescriptorMap, []int{0} }

func init() {
	proto.RegisterType((*MapTest)(nil), "mapdefaults.MapTest")
}
func (this *MapTest) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return MapDescription()
}
func MapDescription() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	d := &github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet{}
	var gzipped = []byte{
		// 3742 bytes of a gzipped FileDescriptorSet
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x5a, 0x5d, 0x6c, 0x23, 0xd7,
		0x75, 0xd6, 0xf0, 0x47, 0x22, 0x0f, 0x29, 0x6a, 0x34, 0x92, 0x77, 0xb9, 0x72, 0xcc, 0xd5, 0x2a,
		0x76, 0x2c, 0xdb, 0x35, 0x15, 0xac, 0xbd, 0xeb, 0xdd, 0xd9, 0x26, 0x06, 0x45, 0x71, 0x15, 0x6e,
		0x25, 0x91, 0x19, 0x52, 0xf1, 0x3a, 0x7d, 0x18, 0x8c, 0x86, 0x97, 0xd4, 0xec, 0x0e, 0x67, 0x26,
		0x33, 0xc3, 0x5d, 0xcb, 0x2f, 0xdd, 0xc2, 0xfd, 0x41, 0x50, 0xa4, 0xff, 0x40, 0x13, 0xd7, 0x71,
		0xdb, 0x00, 0xa9, 0xd3, 0xf4, 0x2f, 0xe9, 0x4f, 0x1a, 0xf4, 0xa9, 0x2f, 0x69, 0xfd, 0x54, 0x38,
		0x0f, 0x05, 0xfa, 0xd0, 0x07, 0xaf, 0x6a, 0xa0, 0x69, 0xeb, 0xb6, 0x6e, 0x6b, 0xa0, 0x01, 0xfc,
		0xd0, 0xe0, 0xfe, 0x0d, 0x67, 0x48, 0x6a, 0x87, 0x0a, 0xe0, 0xe4, 0x49, 0xba, 0xe7, 0x9e, 0xef,
		0x9b, 0x33, 0xe7, 0x9e, 0x7b, 0xce, 0xb9, 0x97, 0x03, 0x5f, 0xb8, 0x0c, 0xab, 0x3d, 0xdb, 0xee,
		0x99, 0x68, 0xc3, 0x71, 0x6d, 0xdf, 0x3e, 0x18, 0x74, 0x37, 0x3a, 0xc8, 0xd3, 0x5d, 0xc3, 0xf1,
		0x6d, 0xb7, 0x4c, 0x64, 0xd2, 0x02, 0xd5, 0x28, 0x73, 0x8d, 0xb5, 0x5d, 0x58, 0xbc, 0x6e, 0x98,
		0x68, 0x2b, 0x50, 0x6c, 0x21, 0x5f, 0xba, 0x02, 0xa9, 0xae, 0x61, 0xa2, 0xa2, 0xb0, 0x9a, 0x5c,
		0xcf, 0x5d, 0x7c, 0xb4, 0x3c, 0x02, 0x2a, 0x47, 0x11, 0x4d, 0x2c, 0x56, 0x08, 0x62, 0xed, 0x9d,
		0x14, 0x2c, 0x4d, 0x98, 0x95, 0x24, 0x48, 0x59, 0x5a, 0x1f, 0x33, 0x0a, 0xeb, 0x59, 0x85, 0xfc,
		0x2f, 0x15, 0x61, 0xce, 0xd1, 0xf4, 0xdb, 0x5a, 0x0f, 0x15, 0x13, 0x44, 0xcc, 0x87, 0x52, 0x09,
		0xa0, 0x83, 0x1c, 0x64, 0x75, 0x90, 0xa5, 0x1f, 0x15, 0x93, 0xab, 0xc9, 0xf5, 0xac, 0x12, 0x92,
		0x48, 0x4f, 0xc1, 0xa2, 0x33, 0x38, 0x30, 0x0d, 0x5d, 0x0d, 0xa9, 0xc1, 0x6a, 0x72, 0x3d, 0xad,
		0x88, 0x74, 0x62, 0x6b, 0xa8, 0xfc, 0x38, 0x2c, 0xdc, 0x45, 0xda, 0xed, 0xb0, 0x6a, 0x8e, 0xa8,
		0x16, 0xb0, 0x38, 0xa4, 0x58, 0x85, 0x7c, 0x1f, 0x79, 0x9e, 0xd6, 0x43, 0xaa, 0x7f, 0xe4, 0xa0,
		0x62, 0x8a, 0xbc, 0xfd, 0xea, 0xd8, 0xdb, 0x8f, 0xbe, 0x79, 0x8e, 0xa1, 0xda, 0x47, 0x0e, 0x92,
		0x2a, 0x90, 0x45, 0xd6, 0xa0, 0x4f, 0x19, 0xd2, 0x27, 0xf8, 0xaf, 0x66, 0x0d, 0xfa, 0xa3, 0x2c,
		0x19, 0x0c, 0x63, 0x14, 0x73, 0x1e, 0x72, 0xef, 0x18, 0x3a, 0x2a, 0xce, 0x12, 0x82, 0xc7, 0xc7,
		0x08, 0x5a, 0x74, 0x7e, 0x94, 0x83, 0xe3, 0xa4, 0x2a, 0x64, 0xd1, 0x4b, 0x3e, 0xb2, 0x3c, 0xc3,
		0xb6, 0x8a, 0x73, 0x84, 0xe4, 0xb1, 0x09, 0xab, 0x88, 0xcc, 0xce, 0x28, 0xc5, 0x10, 0x27, 0x5d,
		0x86, 0x39, 0xdb, 0xf1, 0x0d, 0xdb, 0xf2, 0x8a, 0x99, 0x55, 0x61, 0x3d, 0x77, 0xf1, 0x23, 0x13,
		0x03, 0xa1, 0x41, 0x75, 0x14, 0xae, 0x2c, 0xd5, 0x41, 0xf4, 0xec, 0x81, 0xab, 0x23, 0x55, 0xb7,
		0x3b, 0x48, 0x35, 0xac, 0xae, 0x5d, 0xcc, 0x12, 0x82, 0xf3, 0xe3, 0x2f, 0x42, 0x14, 0xab, 0x76,
		0x07, 0xd5, 0xad, 0xae, 0xad, 0x14, 0xbc, 0xc8, 0x58, 0x3a, 0x03, 0xb3, 0xde, 0x91, 0xe5, 0x6b,
		0x2f, 0x15, 0xf3, 0x24, 0x42, 0xd8, 0x68, 0xed, 0xff, 0xd2, 0xb0, 0x30, 0x4d, 0x88, 0x5d, 0x83,
		0x74, 0x17, 0xbf, 0x65, 0x31, 0x71, 0x1a, 0x1f, 0x50, 0x4c, 0xd4, 0x89, 0xb3, 0x3f, 0xa4, 0x13,
		0x2b, 0x90, 0xb3, 0x90, 0xe7, 0xa3, 0x0e, 0x8d, 0x88, 0xe4, 0x94, 0x31, 0x05, 0x14, 0x34, 0x1e,
		0x52, 0xa9, 0x1f, 0x2a, 0xa4, 0x6e, 0xc2, 0x42, 0x60, 0x92, 0xea, 0x6a, 0x56, 0x8f, 0xc7, 0xe6,
		0x46, 0x9c, 0x25, 0xe5, 0x1a, 0xc7, 0x29, 0x18, 0xa6, 0x14, 0x50, 0x64, 0x2c, 0x6d, 0x01, 0xd8,
		0x16, 0xb2, 0xbb, 0x6a, 0x07, 0xe9, 0x66, 0x31, 0x73, 0x82, 0x97, 0x1a, 0x58, 0x65, 0xcc, 0x4b,
		0x36, 0x95, 0xea, 0xa6, 0x74, 0x75, 0x18, 0x6a, 0x73, 0x27, 0x44, 0xca, 0x2e, 0xdd, 0x64, 0x63,
		0xd1, 0xb6, 0x0f, 0x05, 0x17, 0xe1, 0xb8, 0x47, 0x1d, 0xf6, 0x66, 0x59, 0x62, 0x44, 0x39, 0xf6,
		0xcd, 0x14, 0x06, 0xa3, 0x2f, 0x36, 0xef, 0x86, 0x87, 0xd2, 0x47, 0x21, 0x10, 0xa8, 0x24, 0xac,
		0x80, 0x64, 0xa1, 0x3c, 0x17, 0xee, 0x69, 0x7d, 0xb4, 0x72, 0x05, 0x0a, 0x51, 0xf7, 0x48, 0xcb,
		0x90, 0xf6, 0x7c, 0xcd, 0xf5, 0x49, 0x14, 0xa6, 0x15, 0x3a, 0x90, 0x44, 0x48, 0x22, 0xab, 0x43,
		0xb2, 0x5c, 0x5a, 0xc1, 0xff, 0xae, 0x3c, 0x07, 0xf3, 0x91, 0xc7, 0x4f, 0x0b, 0x5c, 0xfb, 0xe2,
		0x2c, 0x2c, 0x4f, 0x8a, 0xb9, 0x89, 0xe1, 0x7f, 0x06, 0x66, 0xad, 0x41, 0xff, 0x00, 0xb9, 0xc5,
		0x24, 0x61, 0x60, 0x23, 0xa9, 0x02, 0x69, 0x53, 0x3b, 0x40, 0x66, 0x31, 0xb5, 0x2a, 0xac, 0x17,
		0x2e, 0x3e, 0x35, 0x55, 0x54, 0x97, 0x77, 0x30, 0x44, 0xa1, 0x48, 0xe9, 0x93, 0x90, 0x62, 0x29,
		0x0e, 0x33, 0x3c, 0x39, 0x1d, 0x03, 0x8e, 0x45, 0x85, 0xe0, 0xa4, 0x87, 0x21, 0x8b, 0xff, 0x52,
		0xdf, 0xce, 0x12, 0x9b, 0x33, 0x58, 0x80, 0xfd, 0x2a, 0xad, 0x40, 0x86, 0x84, 0x59, 0x07, 0xf1,
		0xd2, 0x10, 0x8c, 0xf1, 0xc2, 0x74, 0x50, 0x57, 0x1b, 0x98, 0xbe, 0x7a, 0x47, 0x33, 0x07, 0x88,
		0x04, 0x4c, 0x56, 0xc9, 0x33, 0xe1, 0x67, 0xb0, 0x4c, 0x3a, 0x0f, 0x39, 0x1a, 0x95, 0x86, 0xd5,
		0x41, 0x2f, 0x91, 0xec, 0x93, 0x56, 0x68, 0xa0, 0xd6, 0xb1, 0x04, 0x3f, 0xfe, 0x96, 0x67, 0x5b,
		0x7c, 0x69, 0xc9, 0x23, 0xb0, 0x80, 0x3c, 0xfe, 0xb9, 0xd1, 0xc4, 0xf7, 0xc8, 0xe4, 0xd7, 0x1b,
		0x8d, 0xc5, 0xb5, 0x6f, 0x25, 0x20, 0x45, 0xf6, 0xdb, 0x02, 0xe4, 0xda, 0x2f, 0x36, 0x6b, 0xea,
		0x56, 0x63, 0x7f, 0x73, 0xa7, 0x26, 0x0a, 0x52, 0x01, 0x80, 0x08, 0xae, 0xef, 0x34, 0x2a, 0x6d,
		0x31, 0x11, 0x8c, 0xeb, 0x7b, 0xed, 0xcb, 0xcf, 0x8a, 0xc9, 0x00, 0xb0, 0x4f, 0x05, 0xa9, 0xb0,
		0xc2, 0x33, 0x17, 0xc5, 0xb4, 0x24, 0x42, 0x9e, 0x12, 0xd4, 0x6f, 0xd6, 0xb6, 0x2e, 0x3f, 0x2b,
		0xce, 0x46, 0x25, 0xcf, 0x5c, 0x14, 0xe7, 0xa4, 0x79, 0xc8, 0x12, 0xc9, 0x66, 0xa3, 0xb1, 0x23,
		0x66, 0x02, 0xce, 0x56, 0x5b, 0xa9, 0xef, 0x6d, 0x8b, 0xd9, 0x80, 0x73, 0x5b, 0x69, 0xec, 0x37,
		0x45, 0x08, 0x18, 0x76, 0x6b, 0xad, 0x56, 0x65, 0xbb, 0x26, 0xe6, 0x02, 0x8d, 0xcd, 0x17, 0xdb,
		0xb5, 0x96, 0x98, 0x8f, 0x98, 0xf5, 0xcc, 0x45, 0x71, 0x3e, 0x78, 0x44, 0x6d, 0x6f, 0x7f, 0x57,
		0x2c, 0x48, 0x8b, 0x30, 0x4f, 0x1f, 0xc1, 0x8d, 0x58, 0x18, 0x11, 0x5d, 0x7e, 0x56, 0x14, 0x87,
		0x86, 0x50, 0x96, 0xc5, 0x88, 0xe0, 0xf2, 0xb3, 0xa2, 0xb4, 0x56, 0x85, 0x34, 0x89, 0x2e, 0x49,
		0x82, 0xc2, 0x4e, 0x65, 0xb3, 0xb6, 0xa3, 0x36, 0x9a, 0xed, 0x7a, 0x63, 0xaf, 0xb2, 0x23, 0x0a,
		0x43, 0x99, 0x52, 0xfb, 0xf4, 0x7e, 0x5d, 0xa9, 0x6d, 0x89, 0x89, 0xb0, 0xac, 0x59, 0xab, 0xb4,
		0x6b, 0x5b, 0x62, 0x72, 0x4d, 0x87, 0xe5, 0x49, 0x79, 0x66, 0xe2, 0xce, 0x08, 0x2d, 0x71, 0xe2,
		0x84, 0x25, 0x26, 0x5c, 0x63, 0x4b, 0xfc, 0x15, 0x01, 0x96, 0x26, 0xe4, 0xda, 0x89, 0x0f, 0x79,
		0x1e, 0xd2, 0x34, 0x44, 0x69, 0xf5, 0x79, 0x62, 0x62, 0xd2, 0x26, 0x01, 0x3b, 0x56, 0x81, 0x08,
		0x2e, 0x5c, 0x81, 0x93, 0x27, 0x54, 0x60, 0x4c, 0x31, 0x66, 0xe4, 0x2b, 0x02, 0x14, 0x4f, 0xe2,
		0x8e, 0x49, 0x14, 0x89, 0x48, 0xa2, 0xb8, 0x36, 0x6a, 0xc0, 0x85, 0x93, 0xdf, 0x61, 0xcc, 0x8a,
		0x37, 0x04, 0x38, 0x33, 0xb9, 0x51, 0x99, 0x68, 0xc3, 0x27, 0x61, 0xb6, 0x8f, 0xfc, 0x43, 0x9b,
		0x17, 0xeb, 0x8f, 0x4d, 0x28, 0x01, 0x78, 0x7a, 0xd4, 0x57, 0x0c, 0x15, 0xae, 0x21, 0xc9, 0x93,
		0xba, 0x0d, 0x6a, 0xcd, 0x98, 0xa5, 0x9f, 0x4f, 0xc0, 0x43, 0x13, 0xc9, 0x27, 0x1a, 0xfa, 0x08,
		0x80, 0x61, 0x39, 0x03, 0x9f, 0x16, 0x64, 0x9a, 0x9f, 0xb2, 0x44, 0x42, 0xf6, 0x3e, 0xce, 0x3d,
		0x03, 0x3f, 0x98, 0x4f, 0x92, 0x79, 0xa0, 0x22, 0xa2, 0x70, 0x65, 0x68, 0x68, 0x8a, 0x18, 0x5a,
		0x3a, 0xe1, 0x4d, 0xc7, 0x6a, 0xdd, 0xc7, 0x41, 0xd4, 0x4d, 0x03, 0x59, 0xbe, 0xea, 0xf9, 0x2e,
		0xd2, 0xfa, 0x86, 0xd5, 0x23, 0x09, 0x38, 0x23, 0xa7, 0xbb, 0x9a, 0xe9, 0x21, 0x65, 0x81, 0x4e,
		0xb7, 0xf8, 0x2c, 0x46, 0x90, 0x2a, 0xe3, 0x86, 0x10, 0xb3, 0x11, 0x04, 0x9d, 0x0e, 0x10, 0x6b,
		0xff, 0x30, 0x07, 0xb9, 0x50, 0x5b, 0x27, 0x5d, 0x80, 0xfc, 0x2d, 0xed, 0x8e, 0xa6, 0xf2, 0x56,
		0x9d, 0x7a, 0x22, 0x87, 0x65, 0x4d, 0xd6, 0xae, 0x7f, 0x1c, 0x96, 0x89, 0x8a, 0x3d, 0xf0, 0x91,
		0xab, 0xea, 0xa6, 0xe6, 0x79, 0xc4, 0x69, 0x19, 0xa2, 0x2a, 0xe1, 0xb9, 0x06, 0x9e, 0xaa, 0xf2,
		0x19, 0xe9, 0x12, 0x2c, 0x11, 0x44, 0x7f, 0x60, 0xfa, 0x86, 0x63, 0x22, 0x15, 0x1f, 0x1e, 0x3c,
		0x92, 0x88, 0x03, 0xcb, 0x16, 0xb1, 0xc6, 0x2e, 0x53, 0xc0, 0x16, 0x79, 0xd2, 0x16, 0x3c, 0x42,
		0x60, 0x3d, 0x64, 0x21, 0x57, 0xf3, 0x91, 0x8a, 0x3e, 0x37, 0xd0, 0x4c, 0x4f, 0xd5, 0xac, 0x8e,
		0x7a, 0xa8, 0x79, 0x87, 0xc5, 0x65, 0x4c, 0xb0, 0x99, 0x28, 0x0a, 0xca, 0x39, 0xac, 0xb8, 0xcd,
		0xf4, 0x6a, 0x44, 0xad, 0x62, 0x75, 0x3e, 0xa5, 0x79, 0x87, 0x92, 0x0c, 0x67, 0x08, 0x8b, 0xe7,
		0xbb, 0x86, 0xd5, 0x53, 0xf5, 0x43, 0xa4, 0xdf, 0x56, 0x07, 0x7e, 0xf7, 0x4a, 0xf1, 0xe1, 0xf0,
		0xf3, 0x89, 0x85, 0x2d, 0xa2, 0x53, 0xc5, 0x2a, 0xfb, 0x7e, 0xf7, 0x8a, 0xd4, 0x82, 0x3c, 0x5e,
		0x8c, 0xbe, 0xf1, 0x32, 0x52, 0xbb, 0xb6, 0x4b, 0x2a, 0x4b, 0x61, 0xc2, 0xce, 0x0e, 0x79, 0xb0,
		0xdc, 0x60, 0x80, 0x5d, 0xbb, 0x83, 0xe4, 0x74, 0xab, 0x59, 0xab, 0x6d, 0x29, 0x39, 0xce, 0x72,
		0xdd, 0x76, 0x71, 0x40, 0xf5, 0xec, 0xc0, 0xc1, 0x39, 0x1a, 0x50, 0x3d, 0x9b, 0xbb, 0xf7, 0x12,
		0x2c, 0xe9, 0x3a, 0x7d, 0x67, 0x43, 0x57, 0x59, 0x8b, 0xef, 0x15, 0xc5, 0x88, 0xb3, 0x74, 0x7d,
		0x9b, 0x2a, 0xb0, 0x18, 0xf7, 0xa4, 0xab, 0xf0, 0xd0, 0xd0, 0x59, 0x61, 0xe0, 0xe2, 0xd8, 0x5b,
		0x8e, 0x42, 0x2f, 0xc1, 0x92, 0x73, 0x34, 0x0e, 0x94, 0x22, 0x4f, 0x74, 0x8e, 0x46, 0x61, 0x8f,
		0x91, 0x63, 0x9b, 0x8b, 0x74, 0xcd, 0x47, 0x9d, 0xe2, 0xd9, 0xb0, 0x76, 0x68, 0x42, 0xda, 0x00,
		0x51, 0xd7, 0x55, 0x64, 0x69, 0x07, 0x26, 0x52, 0x35, 0x17, 0x59, 0x9a, 0x57, 0x3c, 0x1f, 0x56,
		0x2e, 0xe8, 0x7a, 0x8d, 0xcc, 0x56, 0xc8, 0xa4, 0xf4, 0x24, 0x2c, 0xda, 0x07, 0xb7, 0x74, 0x1a,
		0x59, 0xaa, 0xe3, 0xa2, 0xae, 0xf1, 0x52, 0xf1, 0x51, 0xe2, 0xa6, 0x05, 0x3c, 0x41, 0xe2, 0xaa,
		0x49, 0xc4, 0xd2, 0x13, 0x20, 0xea, 0xde, 0xa1, 0xe6, 0x3a, 0xa4, 0xb4, 0x7b, 0x8e, 0xa6, 0xa3,
		0xe2, 0x63, 0x54, 0x95, 0xca, 0xf7, 0xb8, 0x18, 0x47, 0xb6, 0x77, 0xd7, 0xe8, 0xfa, 0x9c, 0xf1,
		0x71, 0x1a, 0xd9, 0x44, 0xc6, 0xd8, 0xd6, 0x41, 0x74, 0x0e, 0x9d, 0xe8, 0x83, 0xd7, 0x89, 0x5a,
		0xc1, 0x39, 0x74, 0xc2, 0xcf, 0xbd, 0x09, 0xcb, 0x03, 0xcb, 0xb0, 0x7c, 0xe4, 0x3a, 0x2e, 0xc2,
		0xed, 0x3e, 0xdd, 0xb3, 0xc5, 0x7f, 0x99, 0x3b, 0xa1, 0x61, 0xdf, 0x0f, 0x6b, 0xd3, 0x50, 0x51,
		0x96, 0x06, 0xe3, 0xc2, 0x35, 0x19, 0xf2, 0xe1, 0x08, 0x92, 0xb2, 0x40, 0x63, 0x48, 0x14, 0x70,
		0x35, 0xae, 0x36, 0xb6, 0x70, 0x1d, 0xfd, 0x6c, 0x4d, 0x4c, 0xe0, 0x7a, 0xbe, 0x53, 0x6f, 0xd7,
		0x54, 0x65, 0x7f, 0xaf, 0x5d, 0xdf, 0xad, 0x89, 0xc9, 0x27, 0xb3, 0x99, 0xef, 0xcd, 0x89, 0xf7,
		0xee, 0xdd, 0xbb, 0x97, 0x58, 0xfb, 0x4e, 0x02, 0x0a, 0xd1, 0x1e, 0x5a, 0xfa, 0x49, 0x38, 0xcb,
		0x0f, 0xbc, 0x1e, 0xf2, 0xd5, 0xbb, 0x86, 0x4b, 0x82, 0xba, 0xaf, 0xd1, 0x2e, 0x34, 0x58, 0x8f,
		0x65, 0xa6, 0xd5, 0x42, 0xfe, 0x0b, 0x86, 0x8b, 0x43, 0xb6, 0xaf, 0xf9, 0xd2, 0x0e, 0x9c, 0xb7,
		0x6c, 0xd5, 0xf3, 0x35, 0xab, 0xa3, 0xb9, 0x1d, 0x75, 0x78, 0xd5, 0xa0, 0x6a, 0xba, 0x8e, 0x3c,
		0xcf, 0xa6, 0xc5, 0x24, 0x60, 0xf9, 0x88, 0x65, 0xb7, 0x98, 0xf2, 0x30, 0xcb, 0x56, 0x98, 0xea,
		0x48, 0xec, 0x24, 0x4f, 0x8a, 0x9d, 0x87, 0x21, 0xdb, 0xd7, 0x1c, 0x15, 0x59, 0xbe, 0x7b, 0x44,
		0x3a, 0xbf, 0x8c, 0x92, 0xe9, 0x6b, 0x4e, 0x0d, 0x8f, 0x3f, 0xbc, 0x35, 0x08, 0xfb, 0xf1, 0x9f,
		0x92, 0x90, 0x0f, 0x77, 0x7f, 0xb8, 0x99, 0xd6, 0x49, 0xa6, 0x17, 0x48, 0x2e, 0xf8, 0xe8, 0x03,
		0x7b, 0xc5, 0x72, 0x15, 0x97, 0x00, 0x79, 0x96, 0xf6, 0x64, 0x0a, 0x45, 0xe2, 0xf2, 0x8b, 0x77,
		0x3f, 0xa2, 0x9d, 0x7e, 0x46, 0x61, 0x23, 0x69, 0x1b, 0x66, 0x6f, 0x79, 0x84, 0x7b, 0x96, 0x70,
		0x3f, 0xfa, 0x60, 0xee, 0x1b, 0x2d, 0x42, 0x9e, 0xbd, 0xd1, 0x52, 0xf7, 0x1a, 0xca, 0x6e, 0x65,
		0x47, 0x61, 0x70, 0xe9, 0x1c, 0xa4, 0x4c, 0xed, 0xe5, 0xa3, 0x68, 0xb1, 0x20, 0xa2, 0x69, 0x1d,
		0x7f, 0x0e, 0x52, 0x77, 0x91, 0x76, 0x3b, 0x9a, 0xa2, 0x89, 0xe8, 0x43, 0x0c, 0xfd, 0x0d, 0x48,
		0x13, 0x7f, 0x49, 0x00, 0xcc, 0x63, 0xe2, 0x8c, 0x94, 0x81, 0x54, 0xb5, 0xa1, 0xe0, 0xf0, 0x17,
		0x21, 0x4f, 0xa5, 0x6a, 0xb3, 0x5e, 0xab, 0xd6, 0xc4, 0xc4, 0xda, 0x25, 0x98, 0xa5, 0x4e, 0xc0,
		0x5b, 0x23, 0x70, 0x83, 0x38, 0xc3, 0x86, 0x8c, 0x43, 0xe0, 0xb3, 0xfb, 0xbb, 0x9b, 0x35, 0x45,
		0x4c, 0x84, 0x97, 0xd7, 0x83, 0x7c, 0xb8, 0xf1, 0xfb, 0xd1, 0xc4, 0xd4, 0x5f, 0x0b, 0x90, 0x0b,
		0x35, 0x72, 0xb8, 0x85, 0xd0, 0x4c, 0xd3, 0xbe, 0xab, 0x6a, 0xa6, 0xa1, 0x79, 0x2c, 0x28, 0x80,
		0x88, 0x2a, 0x58, 0x32, 0xed, 0xa2, 0xfd, 0x48, 0x8c, 0x7f, 0x5d, 0x00, 0x71, 0xb4, 0x09, 0x1c,
		0x31, 0x50, 0xf8, 0xb1, 0x1a, 0xf8, 0x9a, 0x00, 0x85, 0x68, 0xe7, 0x37, 0x62, 0xde, 0x85, 0x1f,
		0xab, 0x79, 0x6f, 0x27, 0x60, 0x3e, 0xd2, 0xef, 0x4d, 0x6b, 0xdd, 0xe7, 0x60, 0xd1, 0xe8, 0xa0,
		0xbe, 0x63, 0xfb, 0xc8, 0xd2, 0x8f, 0x54, 0x13, 0xdd, 0x41, 0x66, 0x71, 0x8d, 0x24, 0x8a, 0x8d,
		0x07, 0x77, 0x94, 0xe5, 0xfa, 0x10, 0xb7, 0x83, 0x61, 0xf2, 0x52, 0x7d, 0xab, 0xb6, 0xdb, 0x6c,
		0xb4, 0x6b, 0x7b, 0xd5, 0x17, 0xd5, 0xfd, 0xbd, 0x9f, 0xda, 0x6b, 0xbc, 0xb0, 0xa7, 0x88, 0xc6,
		0x88, 0xda, 0x87, 0xb8, 0xd5, 0x9b, 0x20, 0x8e, 0x1a, 0x25, 0x9d, 0x85, 0x49, 0x66, 0x89, 0x33,
		0xd2, 0x12, 0x2c, 0xec, 0x35, 0xd4, 0x56, 0x7d, 0xab, 0xa6, 0xd6, 0xae, 0x5f, 0xaf, 0x55, 0xdb,
		0x2d, 0x7a, 0xc4, 0x0e, 0xb4, 0xdb, 0xd1, 0x4d, 0xfd, 0x6a, 0x12, 0x96, 0x26, 0x58, 0x22, 0x55,
		0x58, 0x77, 0x4f, 0x0f, 0x1c, 0x4f, 0x4f, 0x63, 0x7d, 0x19, 0xf7, 0x0f, 0x4d, 0xcd, 0xf5, 0xd9,
		0x61, 0xe0, 0x09, 0xc0, 0x5e, 0xb2, 0x7c, 0xa3, 0x6b, 0x20, 0x97, 0xdd, 0x48, 0xd0, 0x96, 0x7f,
		0x61, 0x28, 0xa7, 0x97, 0x12, 0x3f, 0x01, 0x92, 0x63, 0x7b, 0x86, 0x6f, 0xdc, 0x41, 0xaa, 0x61,
		0xf1, 0xeb, 0x0b, 0x7c, 0x04, 0x48, 0x29, 0x22, 0x9f, 0xa9, 0x5b, 0x7e, 0xa0, 0x6d, 0xa1, 0x9e,
		0x36, 0xa2, 0x8d, 0x13, 0x78, 0x52, 0x11, 0xf9, 0x4c, 0xa0, 0x7d, 0x01, 0xf2, 0x1d, 0x7b, 0x80,
		0x1b, 0x2a, 0xaa, 0x87, 0xeb, 0x85, 0xa0, 0xe4, 0xa8, 0x2c, 0x50, 0x61, 0x1d, 0xef, 0xf0, 0xde,
		0x24, 0xaf, 0xe4, 0xa8, 0x8c, 0xaa, 0x3c, 0x0e, 0x0b, 0x5a, 0xaf, 0xe7, 0x62, 0x72, 0x4e, 0x44,
		0x7b, 0xf8, 0x42, 0x20, 0x26, 0x8a, 0x2b, 0x37, 0x20, 0xc3, 0xfd, 0x80, 0x4b, 0x32, 0xf6, 0x84,
		0xea, 0xd0, 0xdb, 0xab, 0xc4, 0x7a, 0x56, 0xc9, 0x58, 0x7c, 0xf2, 0x02, 0xe4, 0x0d, 0x4f, 0x1d,
		0x5e, 0xa3, 0x26, 0x56, 0x13, 0xeb, 0x19, 0x25, 0x67, 0x78, 0xc1, 0xbd, 0xd9, 0xda, 0x1b, 0x09,
		0x28, 0x44, 0xaf, 0x81, 0xa5, 0x2d, 0xc8, 0x98, 0xb6, 0xae, 0x91, 0xd0, 0xa2, 0xbf, 0x41, 0xac,
		0xc7, 0xdc, 0x1c, 0x97, 0x77, 0x98, 0xbe, 0x12, 0x20, 0x57, 0xfe, 0x5e, 0x80, 0x0c, 0x17, 0x4b,
		0x67, 0x20, 0xe5, 0x68, 0xfe, 0x21, 0xa1, 0x4b, 0x6f, 0x26, 0x44, 0x41, 0x21, 0x63, 0x2c, 0xf7,
		0x1c, 0xcd, 0x22, 0x21, 0xc0, 0xe4, 0x78, 0x8c, 0xd7, 0xd5, 0x44, 0x5a, 0x87, 0x1c, 0x10, 0xec,
		0x7e, 0x1f, 0x59, 0xbe, 0xc7, 0xd7, 0x95, 0xc9, 0xab, 0x4c, 0x2c, 0x3d, 0x05, 0x8b, 0xbe, 0xab,
		0x19, 0x66, 0x44, 0x37, 0x45, 0x74, 0x45, 0x3e, 0x11, 0x28, 0xcb, 0x70, 0x8e, 0xf3, 0x76, 0x90,
		0xaf, 0xe9, 0x87, 0xa8, 0x33, 0x04, 0xcd, 0x92, 0x3b, 0xc6, 0xb3, 0x4c, 0x61, 0x8b, 0xcd, 0x73,
		0xec, 0xda, 0x77, 0x05, 0x58, 0xe4, 0x47, 0x9a, 0x4e, 0xe0, 0xac, 0x5d, 0x00, 0xcd, 0xb2, 0x6c,
		0x3f, 0xec, 0xae, 0xf1, 0x50, 0x1e, 0xc3, 0x95, 0x2b, 0x01, 0x48, 0x09, 0x11, 0xac, 0xf4, 0x01,
		0x86, 0x33, 0x27, 0xba, 0xed, 0x3c, 0xe4, 0xd8, 0x1d, 0x3f, 0xf9, 0xa1, 0x88, 0x1e, 0x82, 0x81,
		0x8a, 0xf0, 0xd9, 0x47, 0x5a, 0x86, 0xf4, 0x01, 0xea, 0x19, 0x16, 0xbb, 0x79, 0xa4, 0x03, 0x7e,
		0x9f, 0x99, 0x0a, 0xee, 0x33, 0x37, 0x6f, 0xc2, 0x92, 0x6e, 0xf7, 0x47, 0xcd, 0xdd, 0x14, 0x47,
		0x0e, 0xe2, 0xde, 0xa7, 0x84, 0xcf, 0xc2, 0xb0, 0xc5, 0xfc, 0x4a, 0x22, 0xb9, 0xdd, 0xdc, 0xfc,
		0x7a, 0x62, 0x65, 0x9b, 0xe2, 0x9a, 0xfc, 0x35, 0x15, 0xd4, 0x35, 0x91, 0x8e, 0x4d, 0x87, 0xff,
		0xff, 0x18, 0x3c, 0xdd, 0x33, 0xfc, 0xc3, 0xc1, 0x41, 0x59, 0xb7, 0xfb, 0x1b, 0x3d, 0xbb, 0x67,
		0x0f, 0x7f, 0x18, 0xc3, 0x23, 0x32, 0x20, 0xff, 0xb1, 0x1f, 0xc7, 0xb2, 0x81, 0x74, 0x25, 0xf6,
		0x97, 0x34, 0x79, 0x0f, 0x96, 0x98, 0xb2, 0x4a, 0x6e, 0xe7, 0xe9, 0xe9, 0x40, 0x7a, 0xe0, 0x0d,
		0x4d, 0xf1, 0x9b, 0xef, 0x90, 0x5a, 0xad, 0x2c, 0x32, 0x28, 0x9e, 0xa3, 0x07, 0x08, 0x59, 0x81,
		0x87, 0x22, 0x7c, 0x74, 0x5f, 0x22, 0x37, 0x86, 0xf1, 0x3b, 0x8c, 0x71, 0x29, 0xc4, 0xd8, 0x62,
		0x50, 0xb9, 0x0a, 0xf3, 0xa7, 0xe1, 0xfa, 0x5b, 0xc6, 0x95, 0x47, 0x61, 0x92, 0x6d, 0x58, 0x20,
		0x24, 0xfa, 0xc0, 0xf3, 0xed, 0x3e, 0x49, 0x7a, 0x0f, 0xa6, 0xf9, 0xbb, 0x77, 0xe8, 0x46, 0x29,
		0x60, 0x58, 0x35, 0x40, 0xc9, 0x32, 0x90, 0x1f, 0x24, 0x3a, 0x48, 0x37, 0x63, 0x18, 0xde, 0x64,
		0x86, 0x04, 0xfa, 0xf2, 0x67, 0x60, 0x19, 0xff, 0x4f, 0x72, 0x52, 0xd8, 0x92, 0xf8, 0xfb, 0xa8,
		0xe2, 0x77, 0x5f, 0xa1, 0x7b, 0x71, 0x29, 0x20, 0x08, 0xd9, 0x14, 0x5a, 0xc5, 0x1e, 0xf2, 0x7d,
		0xe4, 0x7a, 0xaa, 0x66, 0x4e, 0x32, 0x2f, 0x74, 0xa0, 0x2f, 0x7e, 0xe9, 0xdd, 0xe8, 0x2a, 0x6e,
		0x53, 0x64, 0xc5, 0x34, 0xe5, 0x7d, 0x38, 0x3b, 0x21, 0x2a, 0xa6, 0xe0, 0x7c, 0x95, 0x71, 0x2e,
		0x8f, 0x45, 0x06, 0xa6, 0x6d, 0x02, 0x97, 0x07, 0x6b, 0x39, 0x05, 0xe7, 0x6f, 0x33, 0x4e, 0x89,
		0x61, 0xf9, 0x92, 0x62, 0xc6, 0x1b, 0xb0, 0x78, 0x07, 0xb9, 0x07, 0xb6, 0xc7, 0x2e, 0x51, 0xa6,
		0xa0, 0x7b, 0x8d, 0xd1, 0x2d, 0x30, 0x20, 0xb9, 0x55, 0xc1, 0x5c, 0x57, 0x21, 0xd3, 0xd5, 0x74,
		0x34, 0x05, 0xc5, 0x97, 0x19, 0xc5, 0x1c, 0xd6, 0xc7, 0xd0, 0x0a, 0xe4, 0x7b, 0x36, 0x2b, 0x4b,
		0xf1, 0xf0, 0xd7, 0x19, 0x3c, 0xc7, 0x31, 0x8c, 0xc2, 0xb1, 0x9d, 0x81, 0x89, 0x6b, 0x56, 0x3c,
		0xc5, 0xef, 0x70, 0x0a, 0x8e, 0x61, 0x14, 0xa7, 0x70, 0xeb, 0xef, 0x72, 0x0a, 0x2f, 0xe4, 0xcf,
		0xe7, 0x21, 0x67, 0x5b, 0xe6, 0x91, 0x6d, 0x4d, 0x63, 0xc4, 0xef, 0x31, 0x06, 0x60, 0x10, 0x4c,
		0x70, 0x0d, 0xb2, 0xd3, 0x2e, 0xc4, 0x57, 0xdf, 0xe5, 0xdb, 0x83, 0xaf, 0xc0, 0x36, 0x2c, 0xf0,
		0x04, 0x65, 0xd8, 0xd6, 0x14, 0x14, 0xbf, 0xcf, 0x28, 0x0a, 0x21, 0x18, 0x7b, 0x0d, 0x1f, 0x79,
		0x7e, 0x0f, 0x4d, 0x43, 0xf2, 0x06, 0x7f, 0x0d, 0x06, 0x61, 0xae, 0x3c, 0x40, 0x96, 0x7e, 0x38,
		0x1d, 0xc3, 0xd7, 0xb8, 0x2b, 0x39, 0x06, 0x53, 0x54, 0x61, 0xbe, 0xaf, 0xb9, 0xde, 0xa1, 0x66,
		0x4e, 0xb5, 0x1c, 0x7f, 0xc0, 0x38, 0xf2, 0x01, 0x88, 0x79, 0x64, 0x60, 0x9d, 0x86, 0xe6, 0xeb,
		0xdc, 0x23, 0x21, 0x18, 0xdb, 0x7a, 0x9e, 0x4f, 0xae, 0xaa, 0x4e, 0xc3, 0xf6, 0x87, 0x7c, 0xeb,
		0x51, 0xec, 0x6e, 0x98, 0xf1, 0x1a, 0x64, 0x3d, 0xe3, 0xe5, 0xa9, 0x68, 0xfe, 0x88, 0xaf, 0x34,
		0x01, 0x60, 0xf0, 0x8b, 0x70, 0x6e, 0x62, 0x99, 0x98, 0x82, 0xec, 0x8f, 0x19, 0xd9, 0x99, 0x09,
		0xa5, 0x82, 0xa5, 0x84, 0xd3, 0x52, 0xfe, 0x09, 0x4f, 0x09, 0x68, 0x84, 0xab, 0x89, 0x0f, 0x0a,
		0x9e, 0xd6, 0x3d, 0x9d, 0xd7, 0xfe, 0x94, 0x7b, 0x8d, 0x62, 0x23, 0x5e, 0x6b, 0xc3, 0x19, 0xc6,
		0x78, 0xba, 0x75, 0xfd, 0x06, 0x4f, 0xac, 0x14, 0xbd, 0x1f, 0x5d, 0xdd, 0x9f, 0x86, 0x95, 0xc0,
		0x9d, 0xbc, 0x23, 0xf5, 0xd4, 0xbe, 0xe6, 0x4c, 0xc1, 0xfc, 0x4d, 0xc6, 0xcc, 0x33, 0x7e, 0xd0,
		0xd2, 0x7a, 0xbb, 0x9a, 0x83, 0xc9, 0x6f, 0x42, 0x91, 0x93, 0x0f, 0x2c, 0x17, 0xe9, 0x76, 0xcf,
		0x32, 0x5e, 0x46, 0x9d, 0x29, 0xa8, 0xff, 0x6c, 0x64, 0xa9, 0xf6, 0x43, 0x70, 0xcc, 0x5c, 0x07,
		0x31, 0xe8, 0x55, 0x54, 0xa3, 0xef, 0xd8, 0xae, 0x1f, 0xc3, 0xf8, 0xe7, 0x7c, 0xa5, 0x02, 0x5c,
		0x9d, 0xc0, 0xe4, 0x1a, 0x14, 0xc8, 0x70, 0xda, 0x90, 0xfc, 0x0b, 0x46, 0x34, 0x3f, 0x44, 0xb1,
		0xc4, 0xa1, 0xdb, 0x7d, 0x47, 0x73, 0xa7, 0xc9, 0x7f, 0x7f, 0xc9, 0x13, 0x07, 0x83, 0xb0, 0xc4,
		0xe1, 0x1f, 0x39, 0x08, 0x57, 0xfb, 0x29, 0x18, 0xbe, 0xc5, 0x13, 0x07, 0xc7, 0x30, 0x0a, 0xde,
		0x30, 0x4c, 0x41, 0xf1, 0x57, 0x9c, 0x82, 0x63, 0x30, 0xc5, 0xa7, 0x87, 0x85, 0xd6, 0x45, 0x3d,
		0xc3, 0xf3, 0x5d, 0xda, 0x07, 0x3f, 0x98, 0xea, 0xdb, 0xef, 0x46, 0x9b, 0x30, 0x25, 0x04, 0x95,
		0x6f, 0xc0, 0xc2, 0x48, 0x8b, 0x21, 0xc5, 0x7d, 0xdd, 0x50, 0xfc, 0xd9, 0xf7, 0x59, 0x32, 0x8a,
		0x76, 0x18, 0xf2, 0x0e, 0x5e, 0xf7, 0x68, 0x1f, 0x10, 0x4f, 0xf6, 0xca, 0xfb, 0xc1, 0xd2, 0x47,
		0xda, 0x00, 0xf9, 0x3a, 0xcc, 0x47, 0x7a, 0x80, 0x78, 0xaa, 0x9f, 0x63, 0x54, 0xf9, 0x70, 0x0b,
		0x20, 0x5f, 0x82, 0x14, 0xae, 0xe7, 0xf1, 0xf0, 0x9f, 0x67, 0x70, 0xa2, 0x2e, 0x7f, 0x02, 0x32,
		0xbc, 0x8e, 0xc7, 0x43, 0x7f, 0x81, 0x41, 0x03, 0x08, 0x86, 0xf3, 0x1a, 0x1e, 0x0f, 0xff, 0x45,
		0x0e, 0xe7, 0x10, 0x0c, 0x9f, 0xde, 0x85, 0x7f, 0xf3, 0x4b, 0x29, 0x96, 0x87, 0xb9, 0xef, 0xae,
		0xc1, 0x1c, 0x2b, 0xde, 0xf1, 0xe8, 0xcf, 0xb3, 0x87, 0x73, 0x84, 0xfc, 0x1c, 0xa4, 0xa7, 0x74,
		0xf8, 0x17, 0x18, 0x94, 0xea, 0xcb, 0x55, 0xc8, 0x85, 0x0a, 0x76, 0x3c, 0xfc, 0x97, 0x19, 0x3c,
		0x8c, 0xc2, 0xa6, 0xb3, 0x82, 0x1d, 0x4f, 0xf0, 0x2b, 0xdc, 0x74, 0x86, 0xc0, 0x6e, 0xe3, 0xb5,
		0x3a, 0x1e, 0xfd, 0xab, 0xdc, 0xeb, 0x1c, 0x22, 0x3f, 0x0f, 0xd9, 0x20, 0xff, 0xc6, 0xe3, 0x7f,
		0x8d, 0xe1, 0x87, 0x18, 0xec, 0x81, 0x50, 0xfe, 0x8f, 0xa7, 0xf8, 0x75, 0xee, 0x81, 0x10, 0x0a,
		0x6f, 0xa3, 0xd1, 0x9a, 0x1e, 0xcf, 0xf4, 0x1b, 0x7c, 0x1b, 0x8d, 0x94, 0x74, 0xbc, 0x9a, 0x24,
		0x0d, 0xc6, 0x53, 0xfc, 0x26, 0x5f, 0x4d, 0xa2, 0x8f, 0xcd, 0x18, 0x2d, 0x92, 0xf1, 0x1c, 0xbf,
		0xc5, 0xcd, 0x18, 0xa9, 0x91, 0x72, 0x13, 0xa4, 0xf1, 0x02, 0x19, 0xcf, 0xf7, 0x45, 0xc6, 0xb7,
		0x38, 0x56, 0x1f, 0xe5, 0x17, 0xe0, 0xcc, 0xe4, 0xe2, 0x18, 0xcf, 0xfa, 0xa5, 0xf7, 0x47, 0x8e,
		0x33, 0xe1, 0xda, 0x28, 0xb7, 0x87, 0x59, 0x36, 0x5c, 0x18, 0xe3, 0x69, 0x5f, 0x7d, 0x3f, 0x9a,
		0x68, 0xc3, 0x75, 0x51, 0xae, 0x00, 0x0c, 0x6b, 0x52, 0x3c, 0xd7, 0x6b, 0x8c, 0x2b, 0x04, 0xc2,
		0x5b, 0x83, 0x95, 0xa4, 0x78, 0xfc, 0x97, 0xf9, 0xd6, 0x60, 0x08, 0xbc, 0x35, 0x78, 0x35, 0x8a,
		0x47, 0xbf, 0xce, 0xb7, 0x06, 0x87, 0xc8, 0xd7, 0x20, 0x63, 0x0d, 0x4c, 0x13, 0xc7, 0x96, 0xf4,
		0xe0, 0x0f, 0x8e, 0x8a, 0xff, 0xfa, 0x01, 0x03, 0x73, 0x80, 0x7c, 0x09, 0xd2, 0xa8, 0x7f, 0x80,
		0x3a, 0x71, 0xc8, 0x7f, 0xfb, 0x80, 0xe7, 0x13, 0xac, 0x2d, 0x3f, 0x0f, 0x40, 0x0f, 0xd3, 0xe4,
		0x57, 0xa2, 0x18, 0xec, 0xbf, 0x7f, 0xc0, 0xbe, 0x65, 0x18, 0x42, 0x86, 0x04, 0xf4, 0xcb, 0x88,
		0x07, 0x13, 0xbc, 0x1b, 0x25, 0x20, 0x07, 0xf0, 0xab, 0x30, 0x77, 0xcb, 0xb3, 0x2d, 0x5f, 0xeb,
		0xc5, 0xa1, 0xff, 0x83, 0xa1, 0xb9, 0x3e, 0x76, 0x58, 0xdf, 0x76, 0x91, 0xaf, 0xf5, 0xbc, 0x38,
		0xec, 0x7f, 0x32, 0x6c, 0x00, 0xc0, 0x60, 0x5d, 0xf3, 0xfc, 0x69, 0xde, 0xfb, 0xbf, 0x38, 0x98,
		0x03, 0xb0, 0xd1, 0xf8, 0xff, 0xdb, 0xe8, 0x28, 0x0e, 0xfb, 0x1e, 0x37, 0x9a, 0xe9, 0xcb, 0x9f,
		0x80, 0x2c, 0xfe, 0x97, 0x7e, 0xdf, 0x13, 0x03, 0xfe, 0x6f, 0x06, 0x1e, 0x22, 0xf0, 0x93, 0x3d,
		0xbf, 0xe3, 0x1b, 0xf1, 0xce, 0xfe, 0x1f, 0xb6, 0xd2, 0x5c, 0x5f, 0xae, 0x40, 0xce, 0xf3, 0x3b,
		0x9d, 0x01, 0xeb, 0x68, 0x62, 0xe0, 0xff, 0xfb, 0x41, 0x70, 0xc8, 0x0d, 0x30, 0x9b, 0x4f, 0xbf,
		0x79, 0x5c, 0x12, 0xde, 0x3a, 0x2e, 0x09, 0x6f, 0x1f, 0x97, 0x84, 0xc9, 0x17, 0x77, 0xb0, 0x6d,
		0x6f, 0xdb, 0xf4, 0xca, 0x0e, 0xbe, 0x9a, 0x80, 0xa2, 0x6e, 0xf7, 0x0f, 0x6c, 0x6f, 0xc3, 0x42,
		0x86, 0x7f, 0x88, 0xdc, 0x8d, 0xbe, 0xe6, 0xb0, 0xcb, 0xb6, 0x5c, 0x5f, 0x73, 0xd8, 0x37, 0x7b,
		0xde, 0xca, 0xe9, 0x2e, 0xea, 0xd6, 0x7e, 0x06, 0xe6, 0x76, 0x35, 0xa7, 0x8d, 0x3c, 0x5f, 0x22,
		0x0e, 0x21, 0x5f, 0xb7, 0xb0, 0xab, 0xcf, 0xd5, 0x72, 0x88, 0xb8, 0xcc, 0xd4, 0xca, 0x2d, 0xdf,
		0x6d, 0xf9, 0x2e, 0xf9, 0x99, 0x58, 0x99, 0xf5, 0xc8, 0x60, 0xe5, 0x2a, 0xe4, 0x42, 0x62, 0x49,
		0x84, 0xe4, 0x6d, 0x74, 0xc4, 0xbe, 0x6f, 0xc1, 0xff, 0x4a, 0xcb, 0xc3, 0xef, 0xb7, 0xb0, 0x8c,
		0x0e, 0xe4, 0xc4, 0x15, 0x61, 0x73, 0xeb, 0xcd, 0xfb, 0xa5, 0x99, 0xb7, 0xee, 0x97, 0x66, 0xfe,
		0xf1, 0x7e, 0x69, 0xe6, 0xed, 0xfb, 0x25, 0xe1, 0xbd, 0xfb, 0x25, 0xe1, 0xfb, 0xf7, 0x4b, 0xc2,
		0xbd, 0xe3, 0x92, 0xf0, 0xb5, 0xe3, 0x92, 0xf0, 0x8d, 0xe3, 0x92, 0xf0, 0xed, 0xe3, 0x92, 0xf0,
		0xe6, 0x71, 0x69, 0xe6, 0xad, 0xe3, 0xd2, 0x0c, 0x76, 0xd8, 0xf7, 0x8e, 0x4b, 0x33, 0xef, 0x1d,
		0x97, 0x84, 0xef, 0x1f, 0x97, 0x66, 0xee, 0xfd, 0x73, 0x69, 0xe6, 0x60, 0x96, 0xbc, 0xcd, 0x33,
		0x3f, 0x08, 0x00, 0x00, 0xff, 0xff, 0x5d, 0xd9, 0x3a, 0x51, 0xba, 0x2f, 0x00, 0x00,
	}
	r := bytes.NewReader(gzipped)
	gzipr, err := compress_gzip.NewReader(r)
	if err != nil {
		panic(err)
	}
	ungzipped, err := io_ioutil.ReadAll(gzipr)
	if err != nil {
		panic(err)
	}
	if err := github_com_gogo_protobuf_proto.Unmarshal(ungzipped, d); err != nil {
		panic(err)
	}
	return d
}
func (this *MapTest) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*MapTest)
	if !ok {
		that2, ok := that.(MapTest)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *MapTest")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *MapTest but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *MapTest but is not nil && this == nil")
	}
	if len(this.StrStr) != len(that1.StrStr) {
		return fmt.Errorf("StrStr this(%v) Not Equal that(%v)", len(this.StrStr), len(that1.StrStr))
	}
	for i := range this.StrStr {
		if this.StrStr[i] != that1.StrStr[i] {
			return fmt.Errorf("StrStr this[%v](%v) Not Equal that[%v](%v)", i, this.StrStr[i], i, that1.StrStr[i])
		}
	}
	return nil
}
func (this *MapTest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*MapTest)
	if !ok {
		that2, ok := that.(MapTest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if len(this.StrStr) != len(that1.StrStr) {
		return false
	}
	for i := range this.StrStr {
		if this.StrStr[i] != that1.StrStr[i] {
			return false
		}
	}
	return true
}
func (this *MapTest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&mapdefaults.MapTest{")
	keysForStrStr := make([]string, 0, len(this.StrStr))
	for k := range this.StrStr {
		keysForStrStr = append(keysForStrStr, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForStrStr)
	mapStringForStrStr := "map[string]string{"
	for _, k := range keysForStrStr {
		mapStringForStrStr += fmt.Sprintf("%#v: %#v,", k, this.StrStr[k])
	}
	mapStringForStrStr += "}"
	if this.StrStr != nil {
		s = append(s, "StrStr: "+mapStringForStrStr+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMap(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func NewPopulatedMapTest(r randyMap, easy bool) *MapTest {
	this := &MapTest{}
	if r.Intn(10) != 0 {
		v1 := r.Intn(10)
		this.StrStr = make(map[string]string)
		for i := 0; i < v1; i++ {
			this.StrStr[randStringMap(r)] = randStringMap(r)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyMap interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneMap(r randyMap) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringMap(r randyMap) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneMap(r)
	}
	return string(tmps)
}
func randUnrecognizedMap(r randyMap, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldMap(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldMap(dAtA []byte, r randyMap, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateMap(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateMap(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateMap(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateMap(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateMap(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateMap(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateMap(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *MapTest) Size() (n int) {
	var l int
	_ = l
	if len(m.StrStr) > 0 {
		for k, v := range m.StrStr {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovMap(uint64(len(k))) + 1 + len(v) + sovMap(uint64(len(v)))
			n += mapEntrySize + 1 + sovMap(uint64(mapEntrySize))
		}
	}
	return n
}

func sovMap(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMap(x uint64) (n int) {
	return sovMap(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *MapTest) String() string {
	if this == nil {
		return "nil"
	}
	keysForStrStr := make([]string, 0, len(this.StrStr))
	for k := range this.StrStr {
		keysForStrStr = append(keysForStrStr, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForStrStr)
	mapStringForStrStr := "map[string]string{"
	for _, k := range keysForStrStr {
		mapStringForStrStr += fmt.Sprintf("%v: %v,", k, this.StrStr[k])
	}
	mapStringForStrStr += "}"
	s := strings.Join([]string{`&MapTest{`,
		`StrStr:` + mapStringForStrStr + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringMap(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}

func init() { proto.RegisterFile("combos/neither/map.proto", fileDescriptorMap) }

var fileDescriptorMap = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x31, 0x4f, 0x32, 0x41,
	0x10, 0x86, 0x77, 0x20, 0x1f, 0xe4, 0xdb, 0x6b, 0xcc, 0xc5, 0xe2, 0x42, 0x31, 0xb9, 0x58, 0xd1,
	0x78, 0x97, 0x68, 0x23, 0x96, 0x46, 0x4b, 0x1b, 0xb0, 0x37, 0x7b, 0xb8, 0x1c, 0x44, 0x8e, 0xbd,
	0xec, 0xce, 0x99, 0x50, 0xc9, 0xcf, 0xb1, 0xb4, 0xf4, 0x27, 0x50, 0x52, 0x5a, 0xb2, 0x63, 0x63,
	0x49, 0x49, 0x69, 0x5c, 0x2c, 0xe8, 0xde, 0xe7, 0xcd, 0x33, 0x93, 0xbc, 0x32, 0x19, 0x9b, 0xaa,
	0x30, 0x2e, 0x5f, 0xe8, 0x19, 0x4d, 0xb5, 0xcd, 0x2b, 0x55, 0x67, 0xb5, 0x35, 0x64, 0xe2, 0xa8,
	0x52, 0xf5, 0x93, 0x9e, 0xa8, 0x66, 0x4e, 0xae, 0x77, 0x5e, 0xce, 0x68, 0xda, 0x14, 0xd9, 0xd8,
	0x54, 0x79, 0x69, 0x4a, 0x93, 0x07, 0xa7, 0x68, 0x26, 0x81, 0x02, 0x84, 0x74, 0xb8, 0x3d, 0x7b,
	0x95, 0xdd, 0x7b, 0x55, 0x3f, 0x68, 0x47, 0xf1, 0x40, 0x76, 0x1d, 0xd9, 0x47, 0x47, 0x36, 0x81,
	0xb4, 0xdd, 0x8f, 0x2e, 0xd2, 0xec, 0xe8, 0x71, 0xf6, 0xa7, 0x65, 0x23, 0xb2, 0x23, 0xb2, 0x77,
	0x0b, 0xb2, 0xcb, 0x61, 0xc7, 0x05, 0xe8, 0x0d, 0x64, 0x74, 0x54, 0xc7, 0x27, 0xb2, 0xfd, 0xac,
	0x97, 0x09, 0xa4, 0xd0, 0xff, 0x3f, 0xfc, 0x8d, 0xf1, 0xa9, 0xfc, 0xf7, 0xa2, 0xe6, 0x8d, 0x4e,
	0x5a, 0xa1, 0x3b, 0xc0, 0x75, 0xeb, 0x0a, 0x6e, 0x6e, 0xd7, 0x1e, 0xc5, 0xc6, 0xa3, 0xf8, 0xf4,
	0x28, 0xb6, 0x1e, 0x61, 0xe7, 0x11, 0xf6, 0x1e, 0x61, 0xc5, 0x08, 0x6f, 0x8c, 0xf0, 0xce, 0x08,
	0x1f, 0x8c, 0xb0, 0x66, 0x14, 0x1b, 0x46, 0xb1, 0x65, 0x84, 0x6f, 0x46, 0xb1, 0x63, 0x84, 0x3d,
	0xa3, 0x58, 0x7d, 0xa1, 0x28, 0x3a, 0x61, 0xcd, 0xe5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd9,
	0x3d, 0x71, 0xd7, 0x25, 0x01, 0x00, 0x00,
}
