// Code generated by protoc-gen-go.
// source: meta/extensions.proto
// DO NOT EDIT!

/*
Package meta is a generated protocol buffer package.

# Metadata for protoc-gen-template

This provides a set of options for annotating protobuf source code which
are parse by `protoc-gen-template` and exposed in the data available to
templates.

It is generated from these files:
	meta/extensions.proto

It has these top-level messages:
	FileMetadata
	MessageMetadata
	FieldMetadata
	EnumMetadata
	EnumValueMetadata
	ServiceMetadata
	MethodMetadata
*/
package meta

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Visibility int32

const (
	Visibility_PUBLIC  Visibility = 0
	Visibility_PRIVATE Visibility = 1
)

var Visibility_name = map[int32]string{
	0: "PUBLIC",
	1: "PRIVATE",
}
var Visibility_value = map[string]int32{
	"PUBLIC":  0,
	"PRIVATE": 1,
}

func (x Visibility) String() string {
	return proto.EnumName(Visibility_name, int32(x))
}
func (Visibility) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type FileMetadata struct {
	Visibility Visibility `protobuf:"varint,1,opt,name=visibility,enum=meta.Visibility" json:"visibility,omitempty"`
	// Arbitrary string tags
	Tags []string `protobuf:"bytes,2048,rep,name=tags" json:"tags,omitempty"`
	// Arbitrary key/value metadata
	Extra map[string]string `protobuf:"bytes,2047,rep,name=extra" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *FileMetadata) Reset()                    { *m = FileMetadata{} }
func (m *FileMetadata) String() string            { return proto.CompactTextString(m) }
func (*FileMetadata) ProtoMessage()               {}
func (*FileMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FileMetadata) GetVisibility() Visibility {
	if m != nil {
		return m.Visibility
	}
	return Visibility_PUBLIC
}

func (m *FileMetadata) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *FileMetadata) GetExtra() map[string]string {
	if m != nil {
		return m.Extra
	}
	return nil
}

type MessageMetadata struct {
	Visibility Visibility `protobuf:"varint,1,opt,name=visibility,enum=meta.Visibility" json:"visibility,omitempty"`
	// Arbitrary string tags
	Tags []string `protobuf:"bytes,2048,rep,name=tags" json:"tags,omitempty"`
	// Arbitrary key/value metadata
	Extra map[string]string `protobuf:"bytes,2047,rep,name=extra" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *MessageMetadata) Reset()                    { *m = MessageMetadata{} }
func (m *MessageMetadata) String() string            { return proto.CompactTextString(m) }
func (*MessageMetadata) ProtoMessage()               {}
func (*MessageMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MessageMetadata) GetVisibility() Visibility {
	if m != nil {
		return m.Visibility
	}
	return Visibility_PUBLIC
}

func (m *MessageMetadata) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *MessageMetadata) GetExtra() map[string]string {
	if m != nil {
		return m.Extra
	}
	return nil
}

type FieldMetadata struct {
	Visibility Visibility `protobuf:"varint,1,opt,name=visibility,enum=meta.Visibility" json:"visibility,omitempty"`
	// A string identifying a random-value generator for "example" code
	Generator       string  `protobuf:"bytes,2,opt,name=generator" json:"generator,omitempty"`
	ExampleDouble   float64 `protobuf:"fixed64,3,opt,name=example_double,json=exampleDouble" json:"example_double,omitempty"`
	ExampleFloat    float32 `protobuf:"fixed32,4,opt,name=example_float,json=exampleFloat" json:"example_float,omitempty"`
	ExampleInt32    int32   `protobuf:"varint,5,opt,name=example_int32,json=exampleInt32" json:"example_int32,omitempty"`
	ExampleInt64    int64   `protobuf:"varint,6,opt,name=example_int64,json=exampleInt64" json:"example_int64,omitempty"`
	ExampleUint32   uint32  `protobuf:"varint,7,opt,name=example_uint32,json=exampleUint32" json:"example_uint32,omitempty"`
	ExampleUint64   uint64  `protobuf:"varint,8,opt,name=example_uint64,json=exampleUint64" json:"example_uint64,omitempty"`
	ExampleSint32   int32   `protobuf:"zigzag32,9,opt,name=example_sint32,json=exampleSint32" json:"example_sint32,omitempty"`
	ExampleSint64   int64   `protobuf:"zigzag64,10,opt,name=example_sint64,json=exampleSint64" json:"example_sint64,omitempty"`
	ExampleFixed32  uint32  `protobuf:"fixed32,11,opt,name=example_fixed32,json=exampleFixed32" json:"example_fixed32,omitempty"`
	ExampleFixed64  uint64  `protobuf:"fixed64,12,opt,name=example_fixed64,json=exampleFixed64" json:"example_fixed64,omitempty"`
	ExampleSfixed32 int32   `protobuf:"fixed32,13,opt,name=example_sfixed32,json=exampleSfixed32" json:"example_sfixed32,omitempty"`
	ExampleSfixed64 int64   `protobuf:"fixed64,14,opt,name=example_sfixed64,json=exampleSfixed64" json:"example_sfixed64,omitempty"`
	ExampleBool     bool    `protobuf:"varint,15,opt,name=example_bool,json=exampleBool" json:"example_bool,omitempty"`
	ExampleString   string  `protobuf:"bytes,16,opt,name=example_string,json=exampleString" json:"example_string,omitempty"`
	ExampleBytes    []byte  `protobuf:"bytes,17,opt,name=example_bytes,json=exampleBytes,proto3" json:"example_bytes,omitempty"`
	// Arbitrary string tags
	Tags []string `protobuf:"bytes,2048,rep,name=tags" json:"tags,omitempty"`
	// Arbitrary key/value metadata
	Extra map[string]string `protobuf:"bytes,2047,rep,name=extra" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *FieldMetadata) Reset()                    { *m = FieldMetadata{} }
func (m *FieldMetadata) String() string            { return proto.CompactTextString(m) }
func (*FieldMetadata) ProtoMessage()               {}
func (*FieldMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FieldMetadata) GetVisibility() Visibility {
	if m != nil {
		return m.Visibility
	}
	return Visibility_PUBLIC
}

func (m *FieldMetadata) GetGenerator() string {
	if m != nil {
		return m.Generator
	}
	return ""
}

func (m *FieldMetadata) GetExampleDouble() float64 {
	if m != nil {
		return m.ExampleDouble
	}
	return 0
}

func (m *FieldMetadata) GetExampleFloat() float32 {
	if m != nil {
		return m.ExampleFloat
	}
	return 0
}

func (m *FieldMetadata) GetExampleInt32() int32 {
	if m != nil {
		return m.ExampleInt32
	}
	return 0
}

func (m *FieldMetadata) GetExampleInt64() int64 {
	if m != nil {
		return m.ExampleInt64
	}
	return 0
}

func (m *FieldMetadata) GetExampleUint32() uint32 {
	if m != nil {
		return m.ExampleUint32
	}
	return 0
}

func (m *FieldMetadata) GetExampleUint64() uint64 {
	if m != nil {
		return m.ExampleUint64
	}
	return 0
}

func (m *FieldMetadata) GetExampleSint32() int32 {
	if m != nil {
		return m.ExampleSint32
	}
	return 0
}

func (m *FieldMetadata) GetExampleSint64() int64 {
	if m != nil {
		return m.ExampleSint64
	}
	return 0
}

func (m *FieldMetadata) GetExampleFixed32() uint32 {
	if m != nil {
		return m.ExampleFixed32
	}
	return 0
}

func (m *FieldMetadata) GetExampleFixed64() uint64 {
	if m != nil {
		return m.ExampleFixed64
	}
	return 0
}

func (m *FieldMetadata) GetExampleSfixed32() int32 {
	if m != nil {
		return m.ExampleSfixed32
	}
	return 0
}

func (m *FieldMetadata) GetExampleSfixed64() int64 {
	if m != nil {
		return m.ExampleSfixed64
	}
	return 0
}

func (m *FieldMetadata) GetExampleBool() bool {
	if m != nil {
		return m.ExampleBool
	}
	return false
}

func (m *FieldMetadata) GetExampleString() string {
	if m != nil {
		return m.ExampleString
	}
	return ""
}

func (m *FieldMetadata) GetExampleBytes() []byte {
	if m != nil {
		return m.ExampleBytes
	}
	return nil
}

func (m *FieldMetadata) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *FieldMetadata) GetExtra() map[string]string {
	if m != nil {
		return m.Extra
	}
	return nil
}

type EnumMetadata struct {
	Visibility Visibility `protobuf:"varint,1,opt,name=visibility,enum=meta.Visibility" json:"visibility,omitempty"`
	// Arbitrary string tags
	Tags []string `protobuf:"bytes,2048,rep,name=tags" json:"tags,omitempty"`
	// Arbitrary key/value metadata
	Extra map[string]string `protobuf:"bytes,2047,rep,name=extra" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *EnumMetadata) Reset()                    { *m = EnumMetadata{} }
func (m *EnumMetadata) String() string            { return proto.CompactTextString(m) }
func (*EnumMetadata) ProtoMessage()               {}
func (*EnumMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *EnumMetadata) GetVisibility() Visibility {
	if m != nil {
		return m.Visibility
	}
	return Visibility_PUBLIC
}

func (m *EnumMetadata) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *EnumMetadata) GetExtra() map[string]string {
	if m != nil {
		return m.Extra
	}
	return nil
}

type EnumValueMetadata struct {
	Visibility Visibility `protobuf:"varint,1,opt,name=visibility,enum=meta.Visibility" json:"visibility,omitempty"`
	// Arbitrary string tags
	Tags []string `protobuf:"bytes,2048,rep,name=tags" json:"tags,omitempty"`
	// Arbitrary key/value metadata
	Extra map[string]string `protobuf:"bytes,2047,rep,name=extra" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *EnumValueMetadata) Reset()                    { *m = EnumValueMetadata{} }
func (m *EnumValueMetadata) String() string            { return proto.CompactTextString(m) }
func (*EnumValueMetadata) ProtoMessage()               {}
func (*EnumValueMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *EnumValueMetadata) GetVisibility() Visibility {
	if m != nil {
		return m.Visibility
	}
	return Visibility_PUBLIC
}

func (m *EnumValueMetadata) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *EnumValueMetadata) GetExtra() map[string]string {
	if m != nil {
		return m.Extra
	}
	return nil
}

type ServiceMetadata struct {
	Visibility Visibility `protobuf:"varint,1,opt,name=visibility,enum=meta.Visibility" json:"visibility,omitempty"`
	// The service address, ie "grpc.example.com"
	Addr string `protobuf:"bytes,2,opt,name=addr" json:"addr,omitempty"`
	// Arbitrary string tags
	Tags []string `protobuf:"bytes,2048,rep,name=tags" json:"tags,omitempty"`
	// Arbitrary key/value metadata
	Extra map[string]string `protobuf:"bytes,2047,rep,name=extra" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ServiceMetadata) Reset()                    { *m = ServiceMetadata{} }
func (m *ServiceMetadata) String() string            { return proto.CompactTextString(m) }
func (*ServiceMetadata) ProtoMessage()               {}
func (*ServiceMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ServiceMetadata) GetVisibility() Visibility {
	if m != nil {
		return m.Visibility
	}
	return Visibility_PUBLIC
}

func (m *ServiceMetadata) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *ServiceMetadata) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *ServiceMetadata) GetExtra() map[string]string {
	if m != nil {
		return m.Extra
	}
	return nil
}

type MethodMetadata struct {
	Visibility Visibility `protobuf:"varint,1,opt,name=visibility,enum=meta.Visibility" json:"visibility,omitempty"`
	// Arbitrary string tags
	Tags []string `protobuf:"bytes,2048,rep,name=tags" json:"tags,omitempty"`
	// Arbitrary key/value metadata
	Extra map[string]string `protobuf:"bytes,2047,rep,name=extra" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *MethodMetadata) Reset()                    { *m = MethodMetadata{} }
func (m *MethodMetadata) String() string            { return proto.CompactTextString(m) }
func (*MethodMetadata) ProtoMessage()               {}
func (*MethodMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *MethodMetadata) GetVisibility() Visibility {
	if m != nil {
		return m.Visibility
	}
	return Visibility_PUBLIC
}

func (m *MethodMetadata) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *MethodMetadata) GetExtra() map[string]string {
	if m != nil {
		return m.Extra
	}
	return nil
}

var E_FileMeta = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*FileMetadata)(nil),
	Field:         50001,
	Name:          "meta.file_meta",
	Tag:           "bytes,50001,opt,name=file_meta,json=fileMeta",
	Filename:      "meta/extensions.proto",
}

var E_MessageMeta = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*MessageMetadata)(nil),
	Field:         50001,
	Name:          "meta.message_meta",
	Tag:           "bytes,50001,opt,name=message_meta,json=messageMeta",
	Filename:      "meta/extensions.proto",
}

var E_FieldMeta = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*FieldMetadata)(nil),
	Field:         50001,
	Name:          "meta.field_meta",
	Tag:           "bytes,50001,opt,name=field_meta,json=fieldMeta",
	Filename:      "meta/extensions.proto",
}

var E_EnumMeta = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumOptions)(nil),
	ExtensionType: (*EnumMetadata)(nil),
	Field:         50001,
	Name:          "meta.enum_meta",
	Tag:           "bytes,50001,opt,name=enum_meta,json=enumMeta",
	Filename:      "meta/extensions.proto",
}

var E_EnumValueMeta = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumValueOptions)(nil),
	ExtensionType: (*EnumValueMetadata)(nil),
	Field:         50001,
	Name:          "meta.enum_value_meta",
	Tag:           "bytes,50001,opt,name=enum_value_meta,json=enumValueMeta",
	Filename:      "meta/extensions.proto",
}

var E_ServiceMeta = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.ServiceOptions)(nil),
	ExtensionType: (*ServiceMetadata)(nil),
	Field:         50001,
	Name:          "meta.service_meta",
	Tag:           "bytes,50001,opt,name=service_meta,json=serviceMeta",
	Filename:      "meta/extensions.proto",
}

var E_MethodMeta = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MethodOptions)(nil),
	ExtensionType: (*MethodMetadata)(nil),
	Field:         50001,
	Name:          "meta.method_meta",
	Tag:           "bytes,50001,opt,name=method_meta,json=methodMeta",
	Filename:      "meta/extensions.proto",
}

func init() {
	proto.RegisterType((*FileMetadata)(nil), "meta.FileMetadata")
	proto.RegisterType((*MessageMetadata)(nil), "meta.MessageMetadata")
	proto.RegisterType((*FieldMetadata)(nil), "meta.FieldMetadata")
	proto.RegisterType((*EnumMetadata)(nil), "meta.EnumMetadata")
	proto.RegisterType((*EnumValueMetadata)(nil), "meta.EnumValueMetadata")
	proto.RegisterType((*ServiceMetadata)(nil), "meta.ServiceMetadata")
	proto.RegisterType((*MethodMetadata)(nil), "meta.MethodMetadata")
	proto.RegisterEnum("meta.Visibility", Visibility_name, Visibility_value)
	proto.RegisterExtension(E_FileMeta)
	proto.RegisterExtension(E_MessageMeta)
	proto.RegisterExtension(E_FieldMeta)
	proto.RegisterExtension(E_EnumMeta)
	proto.RegisterExtension(E_EnumValueMeta)
	proto.RegisterExtension(E_ServiceMeta)
	proto.RegisterExtension(E_MethodMeta)
}

func init() { proto.RegisterFile("meta/extensions.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 864 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x4f, 0x6f, 0xe3, 0x44,
	0x14, 0x67, 0x9a, 0x34, 0x4d, 0x5e, 0x92, 0x26, 0x9d, 0xdd, 0x15, 0xa3, 0xd5, 0xfe, 0x99, 0x2d,
	0x5a, 0x31, 0x20, 0x6d, 0x82, 0x52, 0x63, 0x96, 0xdc, 0x28, 0xb4, 0x52, 0x25, 0xaa, 0xad, 0xa6,
	0x6c, 0x91, 0xe0, 0x50, 0x39, 0xf5, 0x24, 0xb5, 0xf0, 0x9f, 0xc8, 0x1e, 0x57, 0xe9, 0x8d, 0x13,
	0x07, 0x3e, 0x0c, 0x37, 0x3e, 0x02, 0x07, 0xc4, 0x01, 0xf1, 0x09, 0xf8, 0x28, 0xa0, 0xf1, 0xd8,
	0x8e, 0xe3, 0x24, 0x07, 0xb4, 0x16, 0xb7, 0xcc, 0xcf, 0xbf, 0xfc, 0xde, 0xfb, 0xbd, 0x99, 0xf7,
	0x66, 0xe0, 0x91, 0x27, 0xa4, 0x35, 0x14, 0x0b, 0x29, 0xfc, 0xc8, 0x09, 0xfc, 0x68, 0x30, 0x0f,
	0x03, 0x19, 0xe0, 0xba, 0x82, 0x1f, 0xd3, 0x59, 0x10, 0xcc, 0x5c, 0x31, 0x4c, 0xb0, 0x49, 0x3c,
	0x1d, 0xda, 0x22, 0xba, 0x09, 0x9d, 0xb9, 0x0c, 0x42, 0xcd, 0x3b, 0xfc, 0x0d, 0x41, 0xe7, 0xd4,
	0x71, 0xc5, 0xb9, 0x90, 0x96, 0x6d, 0x49, 0x0b, 0x7f, 0x02, 0x70, 0xe7, 0x44, 0xce, 0xc4, 0x71,
	0x1d, 0x79, 0x4f, 0x10, 0x45, 0x6c, 0x7f, 0xd4, 0x1f, 0x28, 0xb5, 0xc1, 0x55, 0x8e, 0xf3, 0x02,
	0x07, 0x3f, 0x80, 0xba, 0xb4, 0x66, 0x11, 0xf9, 0xb1, 0x4f, 0x6b, 0xac, 0xc5, 0x93, 0x05, 0x36,
	0x60, 0x57, 0x2c, 0x64, 0x68, 0x91, 0x7f, 0x7a, 0xb4, 0xc6, 0xda, 0xa3, 0xa7, 0x5a, 0xa2, 0x18,
	0x6a, 0x70, 0xa2, 0x08, 0x27, 0xbe, 0x0c, 0xef, 0xb9, 0x26, 0x3f, 0x7e, 0x0d, 0xb0, 0x04, 0x71,
	0x1f, 0x6a, 0x3f, 0x08, 0x9d, 0x43, 0x8b, 0xab, 0x9f, 0xf8, 0x21, 0xec, 0xde, 0x59, 0x6e, 0x2c,
	0xc8, 0x4e, 0x82, 0xe9, 0xc5, 0x78, 0xe7, 0x35, 0x3a, 0xfc, 0x03, 0x41, 0xef, 0x5c, 0x44, 0x91,
	0x35, 0xab, 0xdc, 0xca, 0x67, 0x25, 0x2b, 0x54, 0x4b, 0x94, 0xa2, 0x55, 0xea, 0xe6, 0x97, 0x06,
	0x74, 0x4f, 0x1d, 0xe1, 0xda, 0xef, 0xe0, 0xe5, 0x09, 0xb4, 0x66, 0xc2, 0x17, 0xa1, 0x25, 0x83,
	0x30, 0x8d, 0xb0, 0x04, 0xf0, 0x4b, 0xd8, 0x17, 0x0b, 0xcb, 0x9b, 0xbb, 0xe2, 0xda, 0x0e, 0xe2,
	0x89, 0x2b, 0x48, 0x8d, 0x22, 0x86, 0x78, 0x37, 0x45, 0xbf, 0x4a, 0x40, 0xfc, 0x01, 0x64, 0xc0,
	0xf5, 0xd4, 0x0d, 0x2c, 0x49, 0xea, 0x14, 0xb1, 0x1d, 0xde, 0x49, 0xc1, 0x53, 0x85, 0x15, 0x49,
	0x8e, 0x2f, 0x8f, 0x46, 0x64, 0x97, 0x22, 0xb6, 0x9b, 0x93, 0xce, 0x14, 0x56, 0x22, 0x99, 0x06,
	0x69, 0x50, 0xc4, 0x6a, 0x45, 0x92, 0x69, 0x14, 0xb3, 0x8a, 0xb5, 0xd4, 0x1e, 0x45, 0xac, 0x9b,
	0x67, 0xf5, 0x36, 0x01, 0xcb, 0x34, 0xd3, 0x20, 0x4d, 0x8a, 0x58, 0x7d, 0x85, 0xb6, 0xaa, 0x16,
	0x69, 0xb5, 0x16, 0x45, 0xec, 0x20, 0xa7, 0x5d, 0xae, 0xa9, 0x45, 0x5a, 0x0d, 0x28, 0x62, 0x78,
	0x85, 0x66, 0x1a, 0xf8, 0x43, 0xe8, 0xe5, 0xa5, 0x70, 0x16, 0xc2, 0x3e, 0x1a, 0x91, 0x36, 0x45,
	0x6c, 0x8f, 0x67, 0xff, 0x3e, 0xd5, 0xe8, 0x1a, 0xd1, 0x34, 0x48, 0x87, 0x22, 0xd6, 0x58, 0x25,
	0x9a, 0x06, 0xfe, 0x08, 0xfa, 0x79, 0xe0, 0x4c, 0xb2, 0x4b, 0x11, 0xeb, 0xf1, 0x4c, 0xe0, 0x32,
	0x85, 0xd7, 0xa9, 0xa6, 0x41, 0xf6, 0x29, 0x62, 0xfd, 0x12, 0xd5, 0x34, 0xf0, 0x0b, 0xc8, 0x6a,
	0x7a, 0x3d, 0x09, 0x02, 0x97, 0xf4, 0x28, 0x62, 0x4d, 0xde, 0x4e, 0xb1, 0xe3, 0x20, 0x70, 0x57,
	0x1c, 0xcb, 0xd0, 0xf1, 0x67, 0xa4, 0x9f, 0x9c, 0x8f, 0xdc, 0x71, 0x02, 0x16, 0xb7, 0x6c, 0x72,
	0x2f, 0x45, 0x44, 0x0e, 0x28, 0x62, 0x9d, 0x7c, 0xcb, 0x8e, 0x15, 0xb6, 0xb9, 0x65, 0x3e, 0x2d,
	0xb5, 0xcc, 0xb3, 0xac, 0xfb, 0x0b, 0x47, 0xba, 0xd2, 0x86, 0x51, 0x63, 0xec, 0xc4, 0x8f, 0xbd,
	0xff, 0x69, 0x8c, 0x15, 0x43, 0x55, 0xea, 0xe3, 0x4f, 0x04, 0x07, 0x4a, 0xfc, 0x4a, 0x21, 0x55,
	0x9b, 0xf9, 0xbc, 0x64, 0xe6, 0x70, 0x69, 0x66, 0x25, 0x5e, 0xa5, 0x8e, 0xfe, 0x46, 0xd0, 0xbb,
	0x14, 0xe1, 0x9d, 0x73, 0xf3, 0x2e, 0x7e, 0x30, 0xd4, 0x2d, 0xdb, 0xce, 0xe6, 0x58, 0xf2, 0xfb,
	0x3f, 0x0d, 0xeb, 0x52, 0x06, 0x95, 0x3a, 0xfc, 0x1d, 0xc1, 0xfe, 0xb9, 0x90, 0xb7, 0x81, 0x5d,
	0xf5, 0x86, 0x99, 0x25, 0x33, 0xcf, 0xb3, 0x9b, 0xa7, 0x18, 0xac, 0x4a, 0x2f, 0x1f, 0xbf, 0x04,
	0x58, 0x26, 0x88, 0x01, 0x1a, 0x17, 0x6f, 0x8f, 0xbf, 0x3e, 0xfb, 0xb2, 0xff, 0x1e, 0x6e, 0xc3,
	0xde, 0x05, 0x3f, 0xbb, 0xfa, 0xe2, 0x9b, 0x93, 0x3e, 0x1a, 0xbf, 0x81, 0xd6, 0xd4, 0x71, 0xc5,
	0xb5, 0x4a, 0x07, 0x3f, 0x19, 0xe8, 0x57, 0xc6, 0x20, 0x7b, 0x65, 0x24, 0xb7, 0xfc, 0x9b, 0xb9,
	0x54, 0xcf, 0x11, 0xf2, 0xd7, 0x4f, 0xea, 0x62, 0x69, 0x8f, 0xf0, 0xfa, 0x03, 0x80, 0x37, 0xa7,
	0xe9, 0x6a, 0xfc, 0x3d, 0x74, 0x3c, 0x7d, 0x9f, 0x6a, 0xcd, 0xe7, 0x6b, 0x9a, 0xe9, 0x75, 0x5b,
	0x96, 0x7d, 0xb4, 0xf1, 0x32, 0xe6, 0x6d, 0x6f, 0x09, 0x8c, 0x2f, 0x01, 0xa6, 0x6a, 0xf2, 0x68,
	0xe9, 0xa7, 0x1b, 0xd2, 0x15, 0xae, 0x5d, 0x16, 0x7e, 0xb0, 0x61, 0x64, 0xf1, 0xd6, 0x34, 0x5b,
	0xaa, 0x12, 0x08, 0x3f, 0xf6, 0xb6, 0x95, 0x40, 0x35, 0xd5, 0x96, 0x12, 0x14, 0x87, 0x07, 0x6f,
	0x8a, 0x74, 0x35, 0xb6, 0xa1, 0x97, 0x08, 0x26, 0x9b, 0xa1, 0x65, 0x5f, 0x6c, 0x94, 0x4d, 0x7a,
	0xb5, 0xac, 0xfd, 0xfe, 0x96, 0x5e, 0xe6, 0x5d, 0x51, 0x84, 0x54, 0xa1, 0x23, 0xdd, 0x0b, 0xdb,
	0x0a, 0x9d, 0xb6, 0xca, 0x96, 0x42, 0x97, 0x1a, 0x89, 0xb7, 0xa3, 0x25, 0x30, 0xfe, 0x16, 0xda,
	0x5e, 0x72, 0x36, 0xb5, 0xf6, 0xb3, 0x0d, 0x9b, 0xa8, 0xbe, 0x96, 0xa5, 0x1f, 0x6e, 0x3a, 0xd6,
	0x1c, 0xbc, 0x7c, 0x7d, 0x6c, 0xfe, 0xfc, 0x2b, 0xd9, 0x69, 0xa2, 0xef, 0x06, 0x33, 0x47, 0xde,
	0xc6, 0x93, 0xc1, 0x4d, 0xe0, 0x0d, 0xb9, 0x90, 0x71, 0xe8, 0x5f, 0x58, 0xf2, 0x56, 0x3f, 0x6f,
	0x6f, 0x5e, 0xcd, 0x84, 0xff, 0x4a, 0x0a, 0x6f, 0xee, 0x5a, 0x52, 0x0c, 0x95, 0xe4, 0xa4, 0x91,
	0x7c, 0x39, 0xfa, 0x37, 0x00, 0x00, 0xff, 0xff, 0x96, 0x8e, 0x06, 0x59, 0x25, 0x0b, 0x00, 0x00,
}
