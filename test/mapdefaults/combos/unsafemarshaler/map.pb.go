// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: combos/unsafemarshaler/map.proto

/*
Package mapdefaults is a generated protocol buffer package.

It is generated from these files:
	combos/unsafemarshaler/map.proto

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
		0xdb, 0x00, 0xad, 0xd3, 0xb4, 0x69, 0x93, 0xfe, 0xa4, 0x41, 0x9f, 0xfa, 0x92, 0xd6, 0x4f, 0x85,
		0xf3, 0x50, 0xa0, 0x0f, 0x7d, 0xf0, 0xaa, 0x06, 0x9a, 0xb6, 0x6e, 0xeb, 0xb6, 0x06, 0x1a, 0xc0,
		0x0f, 0x0d, 0xee, 0xdf, 0x70, 0x86, 0xa4, 0x76, 0xa8, 0x00, 0x4e, 0x9e, 0xa4, 0x7b, 0xee, 0xf9,
		0xbe, 0x39, 0x73, 0xee, 0xb9, 0xe7, 0x9c, 0x7b, 0x39, 0xf0, 0x85, 0xcb, 0xb0, 0xda, 0xb3, 0xed,
		0x9e, 0x89, 0x36, 0x1c, 0xd7, 0xf6, 0xed, 0x83, 0x41, 0x77, 0xa3, 0x83, 0x3c, 0xdd, 0x35, 0x1c,
		0xdf, 0x76, 0xcb, 0x44, 0x26, 0x2d, 0x50, 0x8d, 0x32, 0xd7, 0x58, 0xdb, 0x85, 0xc5, 0xeb, 0x86,
		0x89, 0xb6, 0x02, 0xc5, 0x16, 0xf2, 0xa5, 0x2b, 0x90, 0xea, 0x1a, 0x26, 0x2a, 0x0a, 0xab, 0xc9,
		0xf5, 0xdc, 0xc5, 0x47, 0xcb, 0x23, 0xa0, 0x72, 0x14, 0xd1, 0xc4, 0x62, 0x85, 0x20, 0xd6, 0xde,
		0x49, 0xc1, 0xd2, 0x84, 0x59, 0x49, 0x82, 0x94, 0xa5, 0xf5, 0x31, 0xa3, 0xb0, 0x9e, 0x55, 0xc8,
		0xff, 0x52, 0x11, 0xe6, 0x1c, 0x4d, 0xbf, 0xad, 0xf5, 0x50, 0x31, 0x41, 0xc4, 0x7c, 0x28, 0x95,
		0x00, 0x3a, 0xc8, 0x41, 0x56, 0x07, 0x59, 0xfa, 0x51, 0x31, 0xb9, 0x9a, 0x5c, 0xcf, 0x2a, 0x21,
		0x89, 0xf4, 0x14, 0x2c, 0x3a, 0x83, 0x03, 0xd3, 0xd0, 0xd5, 0x90, 0x1a, 0xac, 0x26, 0xd7, 0xd3,
		0x8a, 0x48, 0x27, 0xb6, 0x86, 0xca, 0x8f, 0xc3, 0xc2, 0x5d, 0xa4, 0xdd, 0x0e, 0xab, 0xe6, 0x88,
		0x6a, 0x01, 0x8b, 0x43, 0x8a, 0x55, 0xc8, 0xf7, 0x91, 0xe7, 0x69, 0x3d, 0xa4, 0xfa, 0x47, 0x0e,
		0x2a, 0xa6, 0xc8, 0xdb, 0xaf, 0x8e, 0xbd, 0xfd, 0xe8, 0x9b, 0xe7, 0x18, 0xaa, 0x7d, 0xe4, 0x20,
		0xa9, 0x02, 0x59, 0x64, 0x0d, 0xfa, 0x94, 0x21, 0x7d, 0x82, 0xff, 0x6a, 0xd6, 0xa0, 0x3f, 0xca,
		0x92, 0xc1, 0x30, 0x46, 0x31, 0xe7, 0x21, 0xf7, 0x8e, 0xa1, 0xa3, 0xe2, 0x2c, 0x21, 0x78, 0x7c,
		0x8c, 0xa0, 0x45, 0xe7, 0x47, 0x39, 0x38, 0x4e, 0xaa, 0x42, 0x16, 0xbd, 0xe4, 0x23, 0xcb, 0x33,
		0x6c, 0xab, 0x38, 0x47, 0x48, 0x1e, 0x9b, 0xb0, 0x8a, 0xc8, 0xec, 0x8c, 0x52, 0x0c, 0x71, 0xd2,
		0x65, 0x98, 0xb3, 0x1d, 0xdf, 0xb0, 0x2d, 0xaf, 0x98, 0x59, 0x15, 0xd6, 0x73, 0x17, 0x3f, 0x32,
		0x31, 0x10, 0x1a, 0x54, 0x47, 0xe1, 0xca, 0x52, 0x1d, 0x44, 0xcf, 0x1e, 0xb8, 0x3a, 0x52, 0x75,
		0xbb, 0x83, 0x54, 0xc3, 0xea, 0xda, 0xc5, 0x2c, 0x21, 0x38, 0x3f, 0xfe, 0x22, 0x44, 0xb1, 0x6a,
		0x77, 0x50, 0xdd, 0xea, 0xda, 0x4a, 0xc1, 0x8b, 0x8c, 0xa5, 0x33, 0x30, 0xeb, 0x1d, 0x59, 0xbe,
		0xf6, 0x52, 0x31, 0x4f, 0x22, 0x84, 0x8d, 0xd6, 0xfe, 0x2f, 0x0d, 0x0b, 0xd3, 0x84, 0xd8, 0x35,
		0x48, 0x77, 0xf1, 0x5b, 0x16, 0x13, 0xa7, 0xf1, 0x01, 0xc5, 0x44, 0x9d, 0x38, 0xfb, 0x43, 0x3a,
		0xb1, 0x02, 0x39, 0x0b, 0x79, 0x3e, 0xea, 0xd0, 0x88, 0x48, 0x4e, 0x19, 0x53, 0x40, 0x41, 0xe3,
		0x21, 0x95, 0xfa, 0xa1, 0x42, 0xea, 0x26, 0x2c, 0x04, 0x26, 0xa9, 0xae, 0x66, 0xf5, 0x78, 0x6c,
		0x6e, 0xc4, 0x59, 0x52, 0xae, 0x71, 0x9c, 0x82, 0x61, 0x4a, 0x01, 0x45, 0xc6, 0xd2, 0x16, 0x80,
		0x6d, 0x21, 0xbb, 0xab, 0x76, 0x90, 0x6e, 0x16, 0x33, 0x27, 0x78, 0xa9, 0x81, 0x55, 0xc6, 0xbc,
		0x64, 0x53, 0xa9, 0x6e, 0x4a, 0x57, 0x87, 0xa1, 0x36, 0x77, 0x42, 0xa4, 0xec, 0xd2, 0x4d, 0x36,
		0x16, 0x6d, 0xfb, 0x50, 0x70, 0x11, 0x8e, 0x7b, 0xd4, 0x61, 0x6f, 0x96, 0x25, 0x46, 0x94, 0x63,
		0xdf, 0x4c, 0x61, 0x30, 0xfa, 0x62, 0xf3, 0x6e, 0x78, 0x28, 0x7d, 0x14, 0x02, 0x81, 0x4a, 0xc2,
		0x0a, 0x48, 0x16, 0xca, 0x73, 0xe1, 0x9e, 0xd6, 0x47, 0x2b, 0x57, 0xa0, 0x10, 0x75, 0x8f, 0xb4,
		0x0c, 0x69, 0xcf, 0xd7, 0x5c, 0x9f, 0x44, 0x61, 0x5a, 0xa1, 0x03, 0x49, 0x84, 0x24, 0xb2, 0x3a,
		0x24, 0xcb, 0xa5, 0x15, 0xfc, 0xef, 0xca, 0x73, 0x30, 0x1f, 0x79, 0xfc, 0xb4, 0xc0, 0xb5, 0x2f,
		0xce, 0xc2, 0xf2, 0xa4, 0x98, 0x9b, 0x18, 0xfe, 0x67, 0x60, 0xd6, 0x1a, 0xf4, 0x0f, 0x90, 0x5b,
		0x4c, 0x12, 0x06, 0x36, 0x92, 0x2a, 0x90, 0x36, 0xb5, 0x03, 0x64, 0x16, 0x53, 0xab, 0xc2, 0x7a,
		0xe1, 0xe2, 0x53, 0x53, 0x45, 0x75, 0x79, 0x07, 0x43, 0x14, 0x8a, 0x94, 0x3e, 0x09, 0x29, 0x96,
		0xe2, 0x30, 0xc3, 0x93, 0xd3, 0x31, 0xe0, 0x58, 0x54, 0x08, 0x4e, 0x7a, 0x18, 0xb2, 0xf8, 0x2f,
		0xf5, 0xed, 0x2c, 0xb1, 0x39, 0x83, 0x05, 0xd8, 0xaf, 0xd2, 0x0a, 0x64, 0x48, 0x98, 0x75, 0x10,
		0x2f, 0x0d, 0xc1, 0x18, 0x2f, 0x4c, 0x07, 0x75, 0xb5, 0x81, 0xe9, 0xab, 0x77, 0x34, 0x73, 0x80,
		0x48, 0xc0, 0x64, 0x95, 0x3c, 0x13, 0x7e, 0x06, 0xcb, 0xa4, 0xf3, 0x90, 0xa3, 0x51, 0x69, 0x58,
		0x1d, 0xf4, 0x12, 0xc9, 0x3e, 0x69, 0x85, 0x06, 0x6a, 0x1d, 0x4b, 0xf0, 0xe3, 0x6f, 0x79, 0xb6,
		0xc5, 0x97, 0x96, 0x3c, 0x02, 0x0b, 0xc8, 0xe3, 0x9f, 0x1b, 0x4d, 0x7c, 0x8f, 0x4c, 0x7e, 0xbd,
		0xd1, 0x58, 0x5c, 0xfb, 0x56, 0x02, 0x52, 0x64, 0xbf, 0x2d, 0x40, 0xae, 0xfd, 0x62, 0xb3, 0xa6,
		0x6e, 0x35, 0xf6, 0x37, 0x77, 0x6a, 0xa2, 0x20, 0x15, 0x00, 0x88, 0xe0, 0xfa, 0x4e, 0xa3, 0xd2,
		0x16, 0x13, 0xc1, 0xb8, 0xbe, 0xd7, 0xbe, 0xfc, 0xac, 0x98, 0x0c, 0x00, 0xfb, 0x54, 0x90, 0x0a,
		0x2b, 0x3c, 0x73, 0x51, 0x4c, 0x4b, 0x22, 0xe4, 0x29, 0x41, 0xfd, 0x66, 0x6d, 0xeb, 0xf2, 0xb3,
		0xe2, 0x6c, 0x54, 0xf2, 0xcc, 0x45, 0x71, 0x4e, 0x9a, 0x87, 0x2c, 0x91, 0x6c, 0x36, 0x1a, 0x3b,
		0x62, 0x26, 0xe0, 0x6c, 0xb5, 0x95, 0xfa, 0xde, 0xb6, 0x98, 0x0d, 0x38, 0xb7, 0x95, 0xc6, 0x7e,
		0x53, 0x84, 0x80, 0x61, 0xb7, 0xd6, 0x6a, 0x55, 0xb6, 0x6b, 0x62, 0x2e, 0xd0, 0xd8, 0x7c, 0xb1,
		0x5d, 0x6b, 0x89, 0xf9, 0x88, 0x59, 0xcf, 0x5c, 0x14, 0xe7, 0x83, 0x47, 0xd4, 0xf6, 0xf6, 0x77,
		0xc5, 0x82, 0xb4, 0x08, 0xf3, 0xf4, 0x11, 0xdc, 0x88, 0x85, 0x11, 0xd1, 0xe5, 0x67, 0x45, 0x71,
		0x68, 0x08, 0x65, 0x59, 0x8c, 0x08, 0x2e, 0x3f, 0x2b, 0x4a, 0x6b, 0x55, 0x48, 0x93, 0xe8, 0x92,
		0x24, 0x28, 0xec, 0x54, 0x36, 0x6b, 0x3b, 0x6a, 0xa3, 0xd9, 0xae, 0x37, 0xf6, 0x2a, 0x3b, 0xa2,
		0x30, 0x94, 0x29, 0xb5, 0x4f, 0xef, 0xd7, 0x95, 0xda, 0x96, 0x98, 0x08, 0xcb, 0x9a, 0xb5, 0x4a,
		0xbb, 0xb6, 0x25, 0x26, 0xd7, 0x74, 0x58, 0x9e, 0x94, 0x67, 0x26, 0xee, 0x8c, 0xd0, 0x12, 0x27,
		0x4e, 0x58, 0x62, 0xc2, 0x35, 0xb6, 0xc4, 0x5f, 0x11, 0x60, 0x69, 0x42, 0xae, 0x9d, 0xf8, 0x90,
		0xe7, 0x21, 0x4d, 0x43, 0x94, 0x56, 0x9f, 0x27, 0x26, 0x26, 0x6d, 0x12, 0xb0, 0x63, 0x15, 0x88,
		0xe0, 0xc2, 0x15, 0x38, 0x79, 0x42, 0x05, 0xc6, 0x14, 0x63, 0x46, 0xbe, 0x22, 0x40, 0xf1, 0x24,
		0xee, 0x98, 0x44, 0x91, 0x88, 0x24, 0x8a, 0x6b, 0xa3, 0x06, 0x5c, 0x38, 0xf9, 0x1d, 0xc6, 0xac,
		0x78, 0x43, 0x80, 0x33, 0x93, 0x1b, 0x95, 0x89, 0x36, 0x7c, 0x12, 0x66, 0xfb, 0xc8, 0x3f, 0xb4,
		0x79, 0xb1, 0xfe, 0xd8, 0x84, 0x12, 0x80, 0xa7, 0x47, 0x7d, 0xc5, 0x50, 0xe1, 0x1a, 0x92, 0x3c,
		0xa9, 0xdb, 0xa0, 0xd6, 0x8c, 0x59, 0xfa, 0xf9, 0x04, 0x3c, 0x34, 0x91, 0x7c, 0xa2, 0xa1, 0x8f,
		0x00, 0x18, 0x96, 0x33, 0xf0, 0x69, 0x41, 0xa6, 0xf9, 0x29, 0x4b, 0x24, 0x64, 0xef, 0xe3, 0xdc,
		0x33, 0xf0, 0x83, 0xf9, 0x24, 0x99, 0x07, 0x2a, 0x22, 0x0a, 0x57, 0x86, 0x86, 0xa6, 0x88, 0xa1,
		0xa5, 0x13, 0xde, 0x74, 0xac, 0xd6, 0x7d, 0x1c, 0x44, 0xdd, 0x34, 0x90, 0xe5, 0xab, 0x9e, 0xef,
		0x22, 0xad, 0x6f, 0x58, 0x3d, 0x92, 0x80, 0x33, 0x72, 0xba, 0xab, 0x99, 0x1e, 0x52, 0x16, 0xe8,
		0x74, 0x8b, 0xcf, 0x62, 0x04, 0xa9, 0x32, 0x6e, 0x08, 0x31, 0x1b, 0x41, 0xd0, 0xe9, 0x00, 0xb1,
		0xf6, 0x0f, 0x73, 0x90, 0x0b, 0xb5, 0x75, 0xd2, 0x05, 0xc8, 0xdf, 0xd2, 0xee, 0x68, 0x2a, 0x6f,
		0xd5, 0xa9, 0x27, 0x72, 0x58, 0xd6, 0x64, 0xed, 0xfa, 0xc7, 0x61, 0x99, 0xa8, 0xd8, 0x03, 0x1f,
		0xb9, 0xaa, 0x6e, 0x6a, 0x9e, 0x47, 0x9c, 0x96, 0x21, 0xaa, 0x12, 0x9e, 0x6b, 0xe0, 0xa9, 0x2a,
		0x9f, 0x91, 0x2e, 0xc1, 0x12, 0x41, 0xf4, 0x07, 0xa6, 0x6f, 0x38, 0x26, 0x52, 0xf1, 0xe1, 0xc1,
		0x23, 0x89, 0x38, 0xb0, 0x6c, 0x11, 0x6b, 0xec, 0x32, 0x05, 0x6c, 0x91, 0x27, 0x6d, 0xc1, 0x23,
		0x04, 0xd6, 0x43, 0x16, 0x72, 0x35, 0x1f, 0xa9, 0xe8, 0x73, 0x03, 0xcd, 0xf4, 0x54, 0xcd, 0xea,
		0xa8, 0x87, 0x9a, 0x77, 0x58, 0x5c, 0xc6, 0x04, 0x9b, 0x89, 0xa2, 0xa0, 0x9c, 0xc3, 0x8a, 0xdb,
		0x4c, 0xaf, 0x46, 0xd4, 0x2a, 0x56, 0xe7, 0x53, 0x9a, 0x77, 0x28, 0xc9, 0x70, 0x86, 0xb0, 0x78,
		0xbe, 0x6b, 0x58, 0x3d, 0x55, 0x3f, 0x44, 0xfa, 0x6d, 0x75, 0xe0, 0x77, 0xaf, 0x14, 0x1f, 0x0e,
		0x3f, 0x9f, 0x58, 0xd8, 0x22, 0x3a, 0x55, 0xac, 0xb2, 0xef, 0x77, 0xaf, 0x48, 0x2d, 0xc8, 0xe3,
		0xc5, 0xe8, 0x1b, 0x2f, 0x23, 0xb5, 0x6b, 0xbb, 0xa4, 0xb2, 0x14, 0x26, 0xec, 0xec, 0x90, 0x07,
		0xcb, 0x0d, 0x06, 0xd8, 0xb5, 0x3b, 0x48, 0x4e, 0xb7, 0x9a, 0xb5, 0xda, 0x96, 0x92, 0xe3, 0x2c,
		0xd7, 0x6d, 0x17, 0x07, 0x54, 0xcf, 0x0e, 0x1c, 0x9c, 0xa3, 0x01, 0xd5, 0xb3, 0xb9, 0x7b, 0x2f,
		0xc1, 0x92, 0xae, 0xd3, 0x77, 0x36, 0x74, 0x95, 0xb5, 0xf8, 0x5e, 0x51, 0x8c, 0x38, 0x4b, 0xd7,
		0xb7, 0xa9, 0x02, 0x8b, 0x71, 0x4f, 0xba, 0x0a, 0x0f, 0x0d, 0x9d, 0x15, 0x06, 0x2e, 0x8e, 0xbd,
		0xe5, 0x28, 0xf4, 0x12, 0x2c, 0x39, 0x47, 0xe3, 0x40, 0x29, 0xf2, 0x44, 0xe7, 0x68, 0x14, 0xf6,
		0x18, 0x39, 0xb6, 0xb9, 0x48, 0xd7, 0x7c, 0xd4, 0x29, 0x9e, 0x0d, 0x6b, 0x87, 0x26, 0xa4, 0x0d,
		0x10, 0x75, 0x5d, 0x45, 0x96, 0x76, 0x60, 0x22, 0x55, 0x73, 0x91, 0xa5, 0x79, 0xc5, 0xf3, 0x61,
		0xe5, 0x82, 0xae, 0xd7, 0xc8, 0x6c, 0x85, 0x4c, 0x4a, 0x4f, 0xc2, 0xa2, 0x7d, 0x70, 0x4b, 0xa7,
		0x91, 0xa5, 0x3a, 0x2e, 0xea, 0x1a, 0x2f, 0x15, 0x1f, 0x25, 0x6e, 0x5a, 0xc0, 0x13, 0x24, 0xae,
		0x9a, 0x44, 0x2c, 0x3d, 0x01, 0xa2, 0xee, 0x1d, 0x6a, 0xae, 0x43, 0x4a, 0xbb, 0xe7, 0x68, 0x3a,
		0x2a, 0x3e, 0x46, 0x55, 0xa9, 0x7c, 0x8f, 0x8b, 0x71, 0x64, 0x7b, 0x77, 0x8d, 0xae, 0xcf, 0x19,
		0x1f, 0xa7, 0x91, 0x4d, 0x64, 0x8c, 0x6d, 0x1d, 0x44, 0xe7, 0xd0, 0x89, 0x3e, 0x78, 0x9d, 0xa8,
		0x15, 0x9c, 0x43, 0x27, 0xfc, 0xdc, 0x9b, 0xb0, 0x3c, 0xb0, 0x0c, 0xcb, 0x47, 0xae, 0xe3, 0x22,
		0xdc, 0xee, 0xd3, 0x3d, 0x5b, 0xfc, 0x97, 0xb9, 0x13, 0x1a, 0xf6, 0xfd, 0xb0, 0x36, 0x0d, 0x15,
		0x65, 0x69, 0x30, 0x2e, 0x5c, 0x93, 0x21, 0x1f, 0x8e, 0x20, 0x29, 0x0b, 0x34, 0x86, 0x44, 0x01,
		0x57, 0xe3, 0x6a, 0x63, 0x0b, 0xd7, 0xd1, 0xcf, 0xd6, 0xc4, 0x04, 0xae, 0xe7, 0x3b, 0xf5, 0x76,
		0x4d, 0x55, 0xf6, 0xf7, 0xda, 0xf5, 0xdd, 0x9a, 0x98, 0x7c, 0x32, 0x9b, 0xf9, 0xde, 0x9c, 0x78,
		0xef, 0xde, 0xbd, 0x7b, 0x89, 0xb5, 0xef, 0x24, 0xa0, 0x10, 0xed, 0xa1, 0xa5, 0x9f, 0x84, 0xb3,
		0xfc, 0xc0, 0xeb, 0x21, 0x5f, 0xbd, 0x6b, 0xb8, 0x24, 0xa8, 0xfb, 0x1a, 0xed, 0x42, 0x83, 0xf5,
		0x58, 0x66, 0x5a, 0x2d, 0xe4, 0xbf, 0x60, 0xb8, 0x38, 0x64, 0xfb, 0x9a, 0x2f, 0xed, 0xc0, 0x79,
		0xcb, 0x56, 0x3d, 0x5f, 0xb3, 0x3a, 0x9a, 0xdb, 0x51, 0x87, 0x57, 0x0d, 0xaa, 0xa6, 0xeb, 0xc8,
		0xf3, 0x6c, 0x5a, 0x4c, 0x02, 0x96, 0x8f, 0x58, 0x76, 0x8b, 0x29, 0x0f, 0xb3, 0x6c, 0x85, 0xa9,
		0x8e, 0xc4, 0x4e, 0xf2, 0xa4, 0xd8, 0x79, 0x18, 0xb2, 0x7d, 0xcd, 0x51, 0x91, 0xe5, 0xbb, 0x47,
		0xa4, 0xf3, 0xcb, 0x28, 0x99, 0xbe, 0xe6, 0xd4, 0xf0, 0xf8, 0xc3, 0x5b, 0x83, 0xb0, 0x1f, 0xff,
		0x29, 0x09, 0xf9, 0x70, 0xf7, 0x87, 0x9b, 0x69, 0x9d, 0x64, 0x7a, 0x81, 0xe4, 0x82, 0x8f, 0x3e,
		0xb0, 0x57, 0x2c, 0x57, 0x71, 0x09, 0x90, 0x67, 0x69, 0x4f, 0xa6, 0x50, 0x24, 0x2e, 0xbf, 0x78,
		0xf7, 0x23, 0xda, 0xe9, 0x67, 0x14, 0x36, 0x92, 0xb6, 0x61, 0xf6, 0x96, 0x47, 0xb8, 0x67, 0x09,
		0xf7, 0xa3, 0x0f, 0xe6, 0xbe, 0xd1, 0x22, 0xe4, 0xd9, 0x1b, 0x2d, 0x75, 0xaf, 0xa1, 0xec, 0x56,
		0x76, 0x14, 0x06, 0x97, 0xce, 0x41, 0xca, 0xd4, 0x5e, 0x3e, 0x8a, 0x16, 0x0b, 0x22, 0x9a, 0xd6,
		0xf1, 0xe7, 0x20, 0x75, 0x17, 0x69, 0xb7, 0xa3, 0x29, 0x9a, 0x88, 0x3e, 0xc4, 0xd0, 0xdf, 0x80,
		0x34, 0xf1, 0x97, 0x04, 0xc0, 0x3c, 0x26, 0xce, 0x48, 0x19, 0x48, 0x55, 0x1b, 0x0a, 0x0e, 0x7f,
		0x11, 0xf2, 0x54, 0xaa, 0x36, 0xeb, 0xb5, 0x6a, 0x4d, 0x4c, 0xac, 0x5d, 0x82, 0x59, 0xea, 0x04,
		0xbc, 0x35, 0x02, 0x37, 0x88, 0x33, 0x6c, 0xc8, 0x38, 0x04, 0x3e, 0xbb, 0xbf, 0xbb, 0x59, 0x53,
		0xc4, 0x44, 0x78, 0x79, 0x3d, 0xc8, 0x87, 0x1b, 0xbf, 0x1f, 0x4d, 0x4c, 0xfd, 0xb5, 0x00, 0xb9,
		0x50, 0x23, 0x87, 0x5b, 0x08, 0xcd, 0x34, 0xed, 0xbb, 0xaa, 0x66, 0x1a, 0x9a, 0xc7, 0x82, 0x02,
		0x88, 0xa8, 0x82, 0x25, 0xd3, 0x2e, 0xda, 0x8f, 0xc4, 0xf8, 0xd7, 0x05, 0x10, 0x47, 0x9b, 0xc0,
		0x11, 0x03, 0x85, 0x1f, 0xab, 0x81, 0xaf, 0x09, 0x50, 0x88, 0x76, 0x7e, 0x23, 0xe6, 0x5d, 0xf8,
		0xb1, 0x9a, 0xf7, 0x76, 0x02, 0xe6, 0x23, 0xfd, 0xde, 0xb4, 0xd6, 0x7d, 0x0e, 0x16, 0x8d, 0x0e,
		0xea, 0x3b, 0xb6, 0x8f, 0x2c, 0xfd, 0x48, 0x35, 0xd1, 0x1d, 0x64, 0x16, 0xd7, 0x48, 0xa2, 0xd8,
		0x78, 0x70, 0x47, 0x59, 0xae, 0x0f, 0x71, 0x3b, 0x18, 0x26, 0x2f, 0xd5, 0xb7, 0x6a, 0xbb, 0xcd,
		0x46, 0xbb, 0xb6, 0x57, 0x7d, 0x51, 0xdd, 0xdf, 0xfb, 0xa9, 0xbd, 0xc6, 0x0b, 0x7b, 0x8a, 0x68,
		0x8c, 0xa8, 0x7d, 0x88, 0x5b, 0xbd, 0x09, 0xe2, 0xa8, 0x51, 0xd2, 0x59, 0x98, 0x64, 0x96, 0x38,
		0x23, 0x2d, 0xc1, 0xc2, 0x5e, 0x43, 0x6d, 0xd5, 0xb7, 0x6a, 0x6a, 0xed, 0xfa, 0xf5, 0x5a, 0xb5,
		0xdd, 0xa2, 0x47, 0xec, 0x40, 0xbb, 0x1d, 0xdd, 0xd4, 0xaf, 0x26, 0x61, 0x69, 0x82, 0x25, 0x52,
		0x85, 0x75, 0xf7, 0xf4, 0xc0, 0xf1, 0xf4, 0x34, 0xd6, 0x97, 0x71, 0xff, 0xd0, 0xd4, 0x5c, 0x9f,
		0x1d, 0x06, 0x9e, 0x00, 0xec, 0x25, 0xcb, 0x37, 0xba, 0x06, 0x72, 0xd9, 0x8d, 0x04, 0x6d, 0xf9,
		0x17, 0x86, 0x72, 0x7a, 0x29, 0xf1, 0x13, 0x20, 0x39, 0xb6, 0x67, 0xf8, 0xc6, 0x1d, 0xa4, 0x1a,
		0x16, 0xbf, 0xbe, 0xc0, 0x47, 0x80, 0x94, 0x22, 0xf2, 0x99, 0xba, 0xe5, 0x07, 0xda, 0x16, 0xea,
		0x69, 0x23, 0xda, 0x38, 0x81, 0x27, 0x15, 0x91, 0xcf, 0x04, 0xda, 0x17, 0x20, 0xdf, 0xb1, 0x07,
		0xb8, 0xa1, 0xa2, 0x7a, 0xb8, 0x5e, 0x08, 0x4a, 0x8e, 0xca, 0x02, 0x15, 0xd6, 0xf1, 0x0e, 0xef,
		0x4d, 0xf2, 0x4a, 0x8e, 0xca, 0xa8, 0xca, 0xe3, 0xb0, 0xa0, 0xf5, 0x7a, 0x2e, 0x26, 0xe7, 0x44,
		0xb4, 0x87, 0x2f, 0x04, 0x62, 0xa2, 0xb8, 0x72, 0x03, 0x32, 0xdc, 0x0f, 0xb8, 0x24, 0x63, 0x4f,
		0xa8, 0x0e, 0xbd, 0xbd, 0x4a, 0xac, 0x67, 0x95, 0x8c, 0xc5, 0x27, 0x2f, 0x40, 0xde, 0xf0, 0xd4,
		0xe1, 0x35, 0x6a, 0x62, 0x35, 0xb1, 0x9e, 0x51, 0x72, 0x86, 0x17, 0xdc, 0x9b, 0xad, 0xbd, 0x91,
		0x80, 0x42, 0xf4, 0x1a, 0x58, 0xda, 0x82, 0x8c, 0x69, 0xeb, 0x1a, 0x09, 0x2d, 0xfa, 0x1b, 0xc4,
		0x7a, 0xcc, 0xcd, 0x71, 0x79, 0x87, 0xe9, 0x2b, 0x01, 0x72, 0xe5, 0xef, 0x05, 0xc8, 0x70, 0xb1,
		0x74, 0x06, 0x52, 0x8e, 0xe6, 0x1f, 0x12, 0xba, 0xf4, 0x66, 0x42, 0x14, 0x14, 0x32, 0xc6, 0x72,
		0xcf, 0xd1, 0x2c, 0x12, 0x02, 0x4c, 0x8e, 0xc7, 0x78, 0x5d, 0x4d, 0xa4, 0x75, 0xc8, 0x01, 0xc1,
		0xee, 0xf7, 0x91, 0xe5, 0x7b, 0x7c, 0x5d, 0x99, 0xbc, 0xca, 0xc4, 0xd2, 0x53, 0xb0, 0xe8, 0xbb,
		0x9a, 0x61, 0x46, 0x74, 0x53, 0x44, 0x57, 0xe4, 0x13, 0x81, 0xb2, 0x0c, 0xe7, 0x38, 0x6f, 0x07,
		0xf9, 0x9a, 0x7e, 0x88, 0x3a, 0x43, 0xd0, 0x2c, 0xb9, 0x63, 0x3c, 0xcb, 0x14, 0xb6, 0xd8, 0x3c,
		0xc7, 0xae, 0x7d, 0x57, 0x80, 0x45, 0x7e, 0xa4, 0xe9, 0x04, 0xce, 0xda, 0x05, 0xd0, 0x2c, 0xcb,
		0xf6, 0xc3, 0xee, 0x1a, 0x0f, 0xe5, 0x31, 0x5c, 0xb9, 0x12, 0x80, 0x94, 0x10, 0xc1, 0x4a, 0x1f,
		0x60, 0x38, 0x73, 0xa2, 0xdb, 0xce, 0x43, 0x8e, 0xdd, 0xf1, 0x93, 0x1f, 0x8a, 0xe8, 0x21, 0x18,
		0xa8, 0x08, 0x9f, 0x7d, 0xa4, 0x65, 0x48, 0x1f, 0xa0, 0x9e, 0x61, 0xb1, 0x9b, 0x47, 0x3a, 0xe0,
		0xf7, 0x99, 0xa9, 0xe0, 0x3e, 0x73, 0xf3, 0x26, 0x2c, 0xe9, 0x76, 0x7f, 0xd4, 0xdc, 0x4d, 0x71,
		0xe4, 0x20, 0xee, 0x7d, 0x4a, 0xf8, 0x2c, 0x0c, 0x5b, 0xcc, 0xaf, 0x24, 0x92, 0xdb, 0xcd, 0xcd,
		0xaf, 0x25, 0x56, 0xb6, 0x29, 0xae, 0xc9, 0x5f, 0x53, 0x41, 0x5d, 0x13, 0xe9, 0xd8, 0x74, 0xf8,
		0xff, 0x8f, 0xc1, 0xd3, 0x3d, 0xc3, 0x3f, 0x1c, 0x1c, 0x94, 0x75, 0xbb, 0xbf, 0xd1, 0xb3, 0x7b,
		0xf6, 0xf0, 0x87, 0x31, 0x3c, 0x22, 0x03, 0xf2, 0x1f, 0xfb, 0x71, 0x2c, 0x1b, 0x48, 0x57, 0x62,
		0x7f, 0x49, 0x93, 0xf7, 0x60, 0x89, 0x29, 0xab, 0xe4, 0x76, 0x9e, 0x9e, 0x0e, 0xa4, 0x07, 0xde,
		0xd0, 0x14, 0xbf, 0xf9, 0x0e, 0xa9, 0xd5, 0xca, 0x22, 0x83, 0xe2, 0x39, 0x7a, 0x80, 0x90, 0x15,
		0x78, 0x28, 0xc2, 0x47, 0xf7, 0x25, 0x72, 0x63, 0x18, 0xbf, 0xc3, 0x18, 0x97, 0x42, 0x8c, 0x2d,
		0x06, 0x95, 0xab, 0x30, 0x7f, 0x1a, 0xae, 0xbf, 0x65, 0x5c, 0x79, 0x14, 0x26, 0xd9, 0x86, 0x05,
		0x42, 0xa2, 0x0f, 0x3c, 0xdf, 0xee, 0x93, 0xa4, 0xf7, 0x60, 0x9a, 0xbf, 0x7b, 0x87, 0x6e, 0x94,
		0x02, 0x86, 0x55, 0x03, 0x94, 0x2c, 0x03, 0xf9, 0x41, 0xa2, 0x83, 0x74, 0x33, 0x86, 0xe1, 0x4d,
		0x66, 0x48, 0xa0, 0x2f, 0x7f, 0x06, 0x96, 0xf1, 0xff, 0x24, 0x27, 0x85, 0x2d, 0x89, 0xbf, 0x8f,
		0x2a, 0x7e, 0xf7, 0x15, 0xba, 0x17, 0x97, 0x02, 0x82, 0x90, 0x4d, 0xa1, 0x55, 0xec, 0x21, 0xdf,
		0x47, 0xae, 0xa7, 0x6a, 0xe6, 0x24, 0xf3, 0x42, 0x07, 0xfa, 0xe2, 0x97, 0xde, 0x8d, 0xae, 0xe2,
		0x36, 0x45, 0x56, 0x4c, 0x53, 0xde, 0x87, 0xb3, 0x13, 0xa2, 0x62, 0x0a, 0xce, 0x57, 0x19, 0xe7,
		0xf2, 0x58, 0x64, 0x60, 0xda, 0x26, 0x70, 0x79, 0xb0, 0x96, 0x53, 0x70, 0xfe, 0x36, 0xe3, 0x94,
		0x18, 0x96, 0x2f, 0x29, 0x66, 0xbc, 0x01, 0x8b, 0x77, 0x90, 0x7b, 0x60, 0x7b, 0xec, 0x12, 0x65,
		0x0a, 0xba, 0xd7, 0x18, 0xdd, 0x02, 0x03, 0x92, 0x5b, 0x15, 0xcc, 0x75, 0x15, 0x32, 0x5d, 0x4d,
		0x47, 0x53, 0x50, 0x7c, 0x99, 0x51, 0xcc, 0x61, 0x7d, 0x0c, 0xad, 0x40, 0xbe, 0x67, 0xb3, 0xb2,
		0x14, 0x0f, 0x7f, 0x9d, 0xc1, 0x73, 0x1c, 0xc3, 0x28, 0x1c, 0xdb, 0x19, 0x98, 0xb8, 0x66, 0xc5,
		0x53, 0xfc, 0x0e, 0xa7, 0xe0, 0x18, 0x46, 0x71, 0x0a, 0xb7, 0xfe, 0x2e, 0xa7, 0xf0, 0x42, 0xfe,
		0x7c, 0x1e, 0x72, 0xb6, 0x65, 0x1e, 0xd9, 0xd6, 0x34, 0x46, 0xfc, 0x1e, 0x63, 0x00, 0x06, 0xc1,
		0x04, 0xd7, 0x20, 0x3b, 0xed, 0x42, 0xfc, 0xfe, 0xbb, 0x7c, 0x7b, 0xf0, 0x15, 0xd8, 0x86, 0x05,
		0x9e, 0xa0, 0x0c, 0xdb, 0x9a, 0x82, 0xe2, 0x0f, 0x18, 0x45, 0x21, 0x04, 0x63, 0xaf, 0xe1, 0x23,
		0xcf, 0xef, 0xa1, 0x69, 0x48, 0xde, 0xe0, 0xaf, 0xc1, 0x20, 0xcc, 0x95, 0x07, 0xc8, 0xd2, 0x0f,
		0xa7, 0x63, 0xf8, 0x2a, 0x77, 0x25, 0xc7, 0x60, 0x8a, 0x2a, 0xcc, 0xf7, 0x35, 0xd7, 0x3b, 0xd4,
		0xcc, 0xa9, 0x96, 0xe3, 0x0f, 0x19, 0x47, 0x3e, 0x00, 0x31, 0x8f, 0x0c, 0xac, 0xd3, 0xd0, 0x7c,
		0x8d, 0x7b, 0x24, 0x04, 0x63, 0x5b, 0xcf, 0xf3, 0xc9, 0x55, 0xd5, 0x69, 0xd8, 0xfe, 0x88, 0x6f,
		0x3d, 0x8a, 0xdd, 0x0d, 0x33, 0x5e, 0x83, 0xac, 0x67, 0xbc, 0x3c, 0x15, 0xcd, 0x1f, 0xf3, 0x95,
		0x26, 0x00, 0x0c, 0x7e, 0x11, 0xce, 0x4d, 0x2c, 0x13, 0x53, 0x90, 0x7d, 0x9d, 0x91, 0x9d, 0x99,
		0x50, 0x2a, 0x58, 0x4a, 0x38, 0x2d, 0xe5, 0x9f, 0xf0, 0x94, 0x80, 0x46, 0xb8, 0x9a, 0xf8, 0xa0,
		0xe0, 0x69, 0xdd, 0xd3, 0x79, 0xed, 0x4f, 0xb9, 0xd7, 0x28, 0x36, 0xe2, 0xb5, 0x36, 0x9c, 0x61,
		0x8c, 0xa7, 0x5b, 0xd7, 0x6f, 0xf0, 0xc4, 0x4a, 0xd1, 0xfb, 0xd1, 0xd5, 0xfd, 0x69, 0x58, 0x09,
		0xdc, 0xc9, 0x3b, 0x52, 0x4f, 0xed, 0x6b, 0xce, 0x14, 0xcc, 0xdf, 0x64, 0xcc, 0x3c, 0xe3, 0x07,
		0x2d, 0xad, 0xb7, 0xab, 0x39, 0x98, 0xfc, 0x26, 0x14, 0x39, 0xf9, 0xc0, 0x72, 0x91, 0x6e, 0xf7,
		0x2c, 0xe3, 0x65, 0xd4, 0x99, 0x82, 0xfa, 0xcf, 0x46, 0x96, 0x6a, 0x3f, 0x04, 0xc7, 0xcc, 0x75,
		0x10, 0x83, 0x5e, 0x45, 0x35, 0xfa, 0x8e, 0xed, 0xfa, 0x31, 0x8c, 0x7f, 0xce, 0x57, 0x2a, 0xc0,
		0xd5, 0x09, 0x4c, 0xae, 0x41, 0x81, 0x0c, 0xa7, 0x0d, 0xc9, 0xbf, 0x60, 0x44, 0xf3, 0x43, 0x14,
		0x4b, 0x1c, 0xba, 0xdd, 0x77, 0x34, 0x77, 0x9a, 0xfc, 0xf7, 0x97, 0x3c, 0x71, 0x30, 0x08, 0x4b,
		0x1c, 0xfe, 0x91, 0x83, 0x70, 0xb5, 0x9f, 0x82, 0xe1, 0x5b, 0x3c, 0x71, 0x70, 0x0c, 0xa3, 0xe0,
		0x0d, 0xc3, 0x14, 0x14, 0x7f, 0xc5, 0x29, 0x38, 0x06, 0x53, 0x7c, 0x7a, 0x58, 0x68, 0x5d, 0xd4,
		0x33, 0x3c, 0xdf, 0xa5, 0x7d, 0xf0, 0x83, 0xa9, 0xbe, 0xfd, 0x6e, 0xb4, 0x09, 0x53, 0x42, 0x50,
		0xf9, 0x06, 0x2c, 0x8c, 0xb4, 0x18, 0x52, 0xdc, 0xd7, 0x0d, 0xc5, 0x9f, 0x7d, 0x9f, 0x25, 0xa3,
		0x68, 0x87, 0x21, 0xef, 0xe0, 0x75, 0x8f, 0xf6, 0x01, 0xf1, 0x64, 0xaf, 0xbc, 0x1f, 0x2c, 0x7d,
		0xa4, 0x0d, 0x90, 0xaf, 0xc3, 0x7c, 0xa4, 0x07, 0x88, 0xa7, 0xfa, 0x39, 0x46, 0x95, 0x0f, 0xb7,
		0x00, 0xf2, 0x25, 0x48, 0xe1, 0x7a, 0x1e, 0x0f, 0xff, 0x79, 0x06, 0x27, 0xea, 0xf2, 0x27, 0x20,
		0xc3, 0xeb, 0x78, 0x3c, 0xf4, 0x17, 0x18, 0x34, 0x80, 0x60, 0x38, 0xaf, 0xe1, 0xf1, 0xf0, 0x5f,
		0xe4, 0x70, 0x0e, 0xc1, 0xf0, 0xe9, 0x5d, 0xf8, 0x37, 0xbf, 0x94, 0x62, 0x79, 0x98, 0xfb, 0xee,
		0x1a, 0xcc, 0xb1, 0xe2, 0x1d, 0x8f, 0xfe, 0x3c, 0x7b, 0x38, 0x47, 0xc8, 0xcf, 0x41, 0x7a, 0x4a,
		0x87, 0x7f, 0x81, 0x41, 0xa9, 0xbe, 0x5c, 0x85, 0x5c, 0xa8, 0x60, 0xc7, 0xc3, 0x7f, 0x99, 0xc1,
		0xc3, 0x28, 0x6c, 0x3a, 0x2b, 0xd8, 0xf1, 0x04, 0xbf, 0xc2, 0x4d, 0x67, 0x08, 0xec, 0x36, 0x5e,
		0xab, 0xe3, 0xd1, 0xbf, 0xca, 0xbd, 0xce, 0x21, 0xf2, 0xf3, 0x90, 0x0d, 0xf2, 0x6f, 0x3c, 0xfe,
		0xd7, 0x18, 0x7e, 0x88, 0xc1, 0x1e, 0x08, 0xe5, 0xff, 0x78, 0x8a, 0x5f, 0xe7, 0x1e, 0x08, 0xa1,
		0xf0, 0x36, 0x1a, 0xad, 0xe9, 0xf1, 0x4c, 0xbf, 0xc1, 0xb7, 0xd1, 0x48, 0x49, 0xc7, 0xab, 0x49,
		0xd2, 0x60, 0x3c, 0xc5, 0x6f, 0xf2, 0xd5, 0x24, 0xfa, 0xd8, 0x8c, 0xd1, 0x22, 0x19, 0xcf, 0xf1,
		0x5b, 0xdc, 0x8c, 0x91, 0x1a, 0x29, 0x37, 0x41, 0x1a, 0x2f, 0x90, 0xf1, 0x7c, 0x5f, 0x64, 0x7c,
		0x8b, 0x63, 0xf5, 0x51, 0x7e, 0x01, 0xce, 0x4c, 0x2e, 0x8e, 0xf1, 0xac, 0x5f, 0x7a, 0x7f, 0xe4,
		0x38, 0x13, 0xae, 0x8d, 0x72, 0x7b, 0x98, 0x65, 0xc3, 0x85, 0x31, 0x9e, 0xf6, 0xd5, 0xf7, 0xa3,
		0x89, 0x36, 0x5c, 0x17, 0xe5, 0x0a, 0xc0, 0xb0, 0x26, 0xc5, 0x73, 0xbd, 0xc6, 0xb8, 0x42, 0x20,
		0xbc, 0x35, 0x58, 0x49, 0x8a, 0xc7, 0x7f, 0x99, 0x6f, 0x0d, 0x86, 0xc0, 0x5b, 0x83, 0x57, 0xa3,
		0x78, 0xf4, 0xeb, 0x7c, 0x6b, 0x70, 0x88, 0x7c, 0x0d, 0x32, 0xd6, 0xc0, 0x34, 0x71, 0x6c, 0x49,
		0x0f, 0xfe, 0xe0, 0xa8, 0xf8, 0xaf, 0x1f, 0x30, 0x30, 0x07, 0xc8, 0x97, 0x20, 0x8d, 0xfa, 0x07,
		0xa8, 0x13, 0x87, 0xfc, 0xb7, 0x0f, 0x78, 0x3e, 0xc1, 0xda, 0xf2, 0xf3, 0x00, 0xf4, 0x30, 0x4d,
		0x7e, 0x25, 0x8a, 0xc1, 0xfe, 0xfb, 0x07, 0xec, 0x5b, 0x86, 0x21, 0x64, 0x48, 0x40, 0xbf, 0x8c,
		0x78, 0x30, 0xc1, 0xbb, 0x51, 0x02, 0x72, 0x00, 0xbf, 0x0a, 0x73, 0xb7, 0x3c, 0xdb, 0xf2, 0xb5,
		0x5e, 0x1c, 0xfa, 0x3f, 0x18, 0x9a, 0xeb, 0x63, 0x87, 0xf5, 0x6d, 0x17, 0xf9, 0x5a, 0xcf, 0x8b,
		0xc3, 0xfe, 0x27, 0xc3, 0x06, 0x00, 0x0c, 0xd6, 0x35, 0xcf, 0x9f, 0xe6, 0xbd, 0xff, 0x8b, 0x83,
		0x39, 0x00, 0x1b, 0x8d, 0xff, 0xbf, 0x8d, 0x8e, 0xe2, 0xb0, 0xef, 0x71, 0xa3, 0x99, 0xbe, 0xfc,
		0x09, 0xc8, 0xe2, 0x7f, 0xe9, 0xf7, 0x3d, 0x31, 0xe0, 0xff, 0x66, 0xe0, 0x21, 0x02, 0x3f, 0xd9,
		0xf3, 0x3b, 0xbe, 0x11, 0xef, 0xec, 0xff, 0x61, 0x2b, 0xcd, 0xf5, 0xe5, 0x0a, 0xe4, 0x3c, 0xbf,
		0xd3, 0x19, 0xb0, 0x8e, 0x26, 0x06, 0xfe, 0xbf, 0x1f, 0x04, 0x87, 0xdc, 0x00, 0xb3, 0xf9, 0xf4,
		0x9b, 0xc7, 0x25, 0xe1, 0xad, 0xe3, 0x92, 0xf0, 0xf6, 0x71, 0x49, 0x98, 0x7c, 0x71, 0x07, 0xdb,
		0xf6, 0xb6, 0x4d, 0xaf, 0xec, 0xe0, 0xeb, 0x09, 0x58, 0xd5, 0xed, 0xfe, 0x81, 0xed, 0x6d, 0xd0,
		0xe4, 0x12, 0xa4, 0x96, 0x8d, 0xbe, 0xe6, 0xb0, 0x4b, 0xb7, 0x5c, 0x5f, 0x73, 0xd8, 0xb7, 0x7b,
		0xde, 0xca, 0xe9, 0x2e, 0xec, 0xd6, 0x7e, 0x06, 0xe6, 0x76, 0x35, 0xa7, 0x8d, 0x3c, 0x5f, 0x22,
		0x8e, 0x21, 0x5f, 0xb9, 0xb0, 0x2b, 0xd0, 0xd5, 0x72, 0x88, 0xb8, 0xcc, 0xd4, 0xca, 0x2d, 0xdf,
		0x6d, 0xf9, 0x2e, 0xf9, 0xb9, 0x58, 0x99, 0xf5, 0xc8, 0x60, 0xe5, 0x2a, 0xe4, 0x42, 0x62, 0x49,
		0x84, 0xe4, 0x6d, 0x74, 0xc4, 0xbe, 0x73, 0xc1, 0xff, 0x4a, 0xcb, 0xc3, 0xef, 0xb8, 0xb0, 0x8c,
		0x0e, 0xe4, 0xc4, 0x15, 0x61, 0x73, 0xeb, 0xcd, 0xfb, 0xa5, 0x99, 0xb7, 0xee, 0x97, 0x66, 0xfe,
		0xf1, 0x7e, 0x69, 0xe6, 0xed, 0xfb, 0x25, 0xe1, 0xbd, 0xfb, 0x25, 0xe1, 0xfb, 0xf7, 0x4b, 0xc2,
		0xbd, 0xe3, 0x92, 0xf0, 0xd5, 0xe3, 0x92, 0xf0, 0x8d, 0xe3, 0x92, 0xf0, 0xed, 0xe3, 0x92, 0xf0,
		0xe6, 0x71, 0x69, 0xe6, 0xad, 0xe3, 0xd2, 0x0c, 0x76, 0xdc, 0xf7, 0x8e, 0x4b, 0x33, 0xef, 0x1d,
		0x97, 0x84, 0xef, 0x1f, 0x97, 0x84, 0x7b, 0xff, 0x5c, 0x9a, 0x39, 0x98, 0x25, 0x6f, 0xf3, 0xcc,
		0x0f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x88, 0xd9, 0x6f, 0x99, 0xc2, 0x2f, 0x00, 0x00,
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
func (m *MapTest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MapTest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.StrStr) > 0 {
		for k := range m.StrStr {
			dAtA[i] = 0xa
			i++
			v := m.StrStr[k]
			mapSize := 1 + len(k) + sovMap(uint64(len(k))) + 1 + len(v) + sovMap(uint64(len(v)))
			i = encodeVarintMap(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintMap(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintMap(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	return i, nil
}

func encodeFixed64Map(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Map(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintMap(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}

func init() { proto.RegisterFile("combos/unsafemarshaler/map.proto", fileDescriptorMap) }

var fileDescriptorMap = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x2f, 0x4f, 0x03, 0x41,
	0x10, 0xc5, 0x77, 0xda, 0xd0, 0x86, 0x3d, 0x43, 0x2e, 0x88, 0x4b, 0xc5, 0xe4, 0x82, 0xaa, 0xe1,
	0x2e, 0x01, 0x43, 0x91, 0x04, 0x24, 0xa6, 0xc5, 0x93, 0xbd, 0xb2, 0x77, 0x25, 0xdc, 0x75, 0x2f,
	0xfb, 0x87, 0xa4, 0x8a, 0x7e, 0x1c, 0x24, 0x92, 0x8f, 0x50, 0x59, 0x89, 0xec, 0x0e, 0x06, 0x59,
	0x59, 0x49, 0xd8, 0x22, 0xea, 0xde, 0xef, 0xe5, 0x37, 0x93, 0x3c, 0x9e, 0x4e, 0x55, 0x53, 0x28,
	0x93, 0xbb, 0xb9, 0x11, 0xa5, 0x6c, 0x84, 0x36, 0x33, 0x51, 0x4b, 0x9d, 0x37, 0xa2, 0xcd, 0x5a,
	0xad, 0xac, 0x8a, 0xa3, 0x46, 0xb4, 0x4f, 0xb2, 0x14, 0xae, 0xb6, 0x66, 0x70, 0x5e, 0x3d, 0xdb,
	0x99, 0x2b, 0xb2, 0xa9, 0x6a, 0xf2, 0x4a, 0x55, 0x2a, 0x0f, 0x4e, 0xe1, 0xca, 0x40, 0x01, 0x42,
	0xda, 0xdf, 0x9e, 0xbd, 0xf1, 0xfe, 0xbd, 0x68, 0x1f, 0xa4, 0xb1, 0xf1, 0x88, 0xf7, 0x8d, 0xd5,
	0x8f, 0xc6, 0xea, 0x04, 0xd2, 0xee, 0x30, 0xba, 0x48, 0xb3, 0x83, 0xc7, 0xd9, 0xbf, 0x96, 0x4d,
	0xac, 0x9e, 0x58, 0x7d, 0x37, 0xb7, 0x7a, 0x31, 0xee, 0x99, 0x00, 0x83, 0x11, 0x8f, 0x0e, 0xea,
	0xf8, 0x84, 0x77, 0x5f, 0xe4, 0x22, 0x81, 0x14, 0x86, 0xc7, 0xe3, 0xbf, 0x18, 0x9f, 0xf2, 0xa3,
	0x57, 0x51, 0x3b, 0x99, 0x74, 0x42, 0xb7, 0x87, 0xeb, 0xce, 0x15, 0xdc, 0xdc, 0xae, 0x3c, 0xb2,
	0xb5, 0x47, 0xf6, 0xe5, 0x91, 0x6d, 0x3c, 0xc2, 0xd6, 0x23, 0xec, 0x3c, 0xc2, 0x92, 0x10, 0xde,
	0x09, 0xe1, 0x83, 0x10, 0x3e, 0x09, 0x61, 0x45, 0xc8, 0xd6, 0x84, 0x6c, 0x43, 0x08, 0x3f, 0x84,
	0x6c, 0x4b, 0x08, 0x3b, 0x42, 0x58, 0x7e, 0x23, 0x2b, 0x7a, 0x61, 0xcd, 0xe5, 0x6f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xd3, 0x9a, 0xd4, 0x14, 0x2d, 0x01, 0x00, 0x00,
}
