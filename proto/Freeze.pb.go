// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/Freeze.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

// Set the freezing period in which the platform will stop creating events and accepting transactions. This is used before safely shut down the platform for maintenance.
type FreezeTransactionBody struct {
	StartHour            int32    `protobuf:"varint,1,opt,name=startHour,proto3" json:"startHour,omitempty"`
	StartMin             int32    `protobuf:"varint,2,opt,name=startMin,proto3" json:"startMin,omitempty"`
	EndHour              int32    `protobuf:"varint,3,opt,name=endHour,proto3" json:"endHour,omitempty"`
	EndMin               int32    `protobuf:"varint,4,opt,name=endMin,proto3" json:"endMin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FreezeTransactionBody) Reset()         { *m = FreezeTransactionBody{} }
func (m *FreezeTransactionBody) String() string { return proto.CompactTextString(m) }
func (*FreezeTransactionBody) ProtoMessage()    {}
func (*FreezeTransactionBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_bba56fa2d360a17a, []int{0}
}

func (m *FreezeTransactionBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FreezeTransactionBody.Unmarshal(m, b)
}
func (m *FreezeTransactionBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FreezeTransactionBody.Marshal(b, m, deterministic)
}
func (m *FreezeTransactionBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FreezeTransactionBody.Merge(m, src)
}
func (m *FreezeTransactionBody) XXX_Size() int {
	return xxx_messageInfo_FreezeTransactionBody.Size(m)
}
func (m *FreezeTransactionBody) XXX_DiscardUnknown() {
	xxx_messageInfo_FreezeTransactionBody.DiscardUnknown(m)
}

var xxx_messageInfo_FreezeTransactionBody proto.InternalMessageInfo

func (m *FreezeTransactionBody) GetStartHour() int32 {
	if m != nil {
		return m.StartHour
	}
	return 0
}

func (m *FreezeTransactionBody) GetStartMin() int32 {
	if m != nil {
		return m.StartMin
	}
	return 0
}

func (m *FreezeTransactionBody) GetEndHour() int32 {
	if m != nil {
		return m.EndHour
	}
	return 0
}

func (m *FreezeTransactionBody) GetEndMin() int32 {
	if m != nil {
		return m.EndMin
	}
	return 0
}

func init() {
	proto.RegisterType((*FreezeTransactionBody)(nil), "proto.FreezeTransactionBody")
}

func init() {
	proto.RegisterFile("proto/Freeze.proto", fileDescriptor_bba56fa2d360a17a)
}

var fileDescriptor_bba56fa2d360a17a = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8e, 0xb1, 0xca, 0xc2, 0x30,
	0x14, 0x85, 0xc9, 0xff, 0xdb, 0xaa, 0x19, 0x03, 0x4a, 0x29, 0x0e, 0xe2, 0xd4, 0xa5, 0xc9, 0xe0,
	0x1b, 0x74, 0x90, 0x2e, 0x82, 0x88, 0x93, 0x5b, 0xda, 0x84, 0xa6, 0x48, 0x73, 0x4b, 0x92, 0x0e,
	0xba, 0xfa, 0xe2, 0xe2, 0xad, 0xd6, 0xe9, 0xf2, 0x9d, 0xef, 0x1c, 0xb8, 0x94, 0xf5, 0x0e, 0x02,
	0x88, 0x83, 0xd3, 0xfa, 0xa1, 0x39, 0x02, 0x8b, 0xf0, 0xec, 0x9e, 0x84, 0xae, 0xc6, 0xfc, 0xe2,
	0xa4, 0xf5, 0xb2, 0x0e, 0x2d, 0xd8, 0x02, 0xd4, 0x9d, 0x6d, 0xe8, 0xd2, 0x07, 0xe9, 0x42, 0x09,
	0x83, 0x4b, 0xc8, 0x96, 0x64, 0xd1, 0xf9, 0x17, 0xb0, 0x94, 0x2e, 0x10, 0x8e, 0xad, 0x4d, 0xfe,
	0x50, 0x4e, 0xcc, 0x12, 0x3a, 0xd7, 0x56, 0xe1, 0xee, 0x1f, 0xd5, 0x17, 0xd9, 0x9a, 0xc6, 0xda,
	0xaa, 0xf7, 0x66, 0x86, 0xe2, 0x43, 0x45, 0x49, 0xd3, 0x1a, 0x3a, 0x6e, 0xb4, 0xd2, 0x4e, 0x72,
	0x23, 0xbd, 0x69, 0x9c, 0xec, 0xcd, 0xf8, 0xea, 0x89, 0x5c, 0xb3, 0xa6, 0x0d, 0x66, 0xa8, 0x78,
	0x0d, 0x9d, 0x98, 0xac, 0x18, 0xeb, 0xb9, 0x57, 0xb7, 0xbc, 0x01, 0x81, 0xdd, 0x2a, 0xc6, 0xb3,
	0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x89, 0x49, 0x68, 0x23, 0xf3, 0x00, 0x00, 0x00,
}
