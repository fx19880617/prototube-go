// Code generated by protoc-gen-go. DO NOT EDIT.
// source: prototube.proto

package prototube

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Server side message decoration
type PrototubeMessageHeader struct {
	// Wire version of the Prototube message
	Version int32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Schema version of the message
	SchemaVersion int32 `protobuf:"varint,2,opt,name=schema_version,json=schemaVersion,proto3" json:"schema_version,omitempty"`
	// Time when the message is produced since UNIX epoch
	Ts int64 `protobuf:"varint,3,opt,name=ts,proto3" json:"ts,omitempty"`
	// 16 bytes UUID
	Uuid                 []byte   `protobuf:"bytes,4,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrototubeMessageHeader) Reset()         { *m = PrototubeMessageHeader{} }
func (m *PrototubeMessageHeader) String() string { return proto.CompactTextString(m) }
func (*PrototubeMessageHeader) ProtoMessage()    {}
func (*PrototubeMessageHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb72f4170c9dd46a, []int{0}
}

func (m *PrototubeMessageHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrototubeMessageHeader.Unmarshal(m, b)
}
func (m *PrototubeMessageHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrototubeMessageHeader.Marshal(b, m, deterministic)
}
func (m *PrototubeMessageHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrototubeMessageHeader.Merge(m, src)
}
func (m *PrototubeMessageHeader) XXX_Size() int {
	return xxx_messageInfo_PrototubeMessageHeader.Size(m)
}
func (m *PrototubeMessageHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_PrototubeMessageHeader.DiscardUnknown(m)
}

var xxx_messageInfo_PrototubeMessageHeader proto.InternalMessageInfo

func (m *PrototubeMessageHeader) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *PrototubeMessageHeader) GetSchemaVersion() int32 {
	if m != nil {
		return m.SchemaVersion
	}
	return 0
}

func (m *PrototubeMessageHeader) GetTs() int64 {
	if m != nil {
		return m.Ts
	}
	return 0
}

func (m *PrototubeMessageHeader) GetUuid() []byte {
	if m != nil {
		return m.Uuid
	}
	return nil
}

func init() {
	proto.RegisterType((*PrototubeMessageHeader)(nil), "PrototubeMessageHeader")
}

func init() { proto.RegisterFile("prototube.proto", fileDescriptor_bb72f4170c9dd46a) }

var fileDescriptor_bb72f4170c9dd46a = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0x29, 0x4d, 0x4a, 0xd5, 0x03, 0xb3, 0x94, 0x6a, 0xb9, 0xc4, 0x02, 0x60, 0x42, 0xbe,
	0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x1e, 0xa9, 0x89, 0x29, 0xa9, 0x45, 0x42, 0x12, 0x5c, 0xec,
	0x65, 0xa9, 0x45, 0xc5, 0x99, 0xf9, 0x79, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x30, 0xae,
	0x90, 0x2a, 0x17, 0x5f, 0x71, 0x72, 0x46, 0x6a, 0x6e, 0x62, 0x3c, 0x4c, 0x01, 0x13, 0x58, 0x01,
	0x2f, 0x44, 0x34, 0x0c, 0xaa, 0x8c, 0x8f, 0x8b, 0xa9, 0xa4, 0x58, 0x82, 0x59, 0x81, 0x51, 0x83,
	0x39, 0x88, 0xa9, 0xa4, 0x58, 0x48, 0x88, 0x8b, 0xa5, 0xb4, 0x34, 0x33, 0x45, 0x82, 0x45, 0x81,
	0x51, 0x83, 0x27, 0x08, 0xcc, 0x76, 0x12, 0xe3, 0x12, 0xc9, 0xcc, 0xd7, 0x4b, 0xcc, 0x29, 0x49,
	0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x4b, 0x49, 0x2c, 0x49, 0xd4, 0xcb, 0x4c, 0xc9, 0x49, 0x62, 0x03,
	0xbb, 0xce, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x27, 0xe1, 0x91, 0xb0, 0x00, 0x00, 0x00,
}
