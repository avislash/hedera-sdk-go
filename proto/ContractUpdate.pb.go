// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/ContractUpdate.proto

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

// Modify a smart contract instance to have the given parameter values. Any null field is ignored (left unchanged). If only the contractInstanceExpirationTime is being modified, then no signature is needed on this transaction other than for the account paying for the transaction itself. But if any of the other fields are being modified, then it must be signed by the adminKey. The use of adminKey is not currently supported in this API, but in the future will be implemented to allow these fields to be modified, and also to make modifications to the state of the instance. If the contract is created with no admin key, then none of the fields can be changed that need an admin signature, and therefore no admin key can ever be added. So if there is no admin key, then things like the bytecode are immutable. But if there is an admin key, then they can be changed. For example, the admin key might be a threshold key, which requires 3 of 5 binding arbitration judges to agree before the bytecode can be changed. This can be used to add flexibility to the management of smart contract behavior. But this is optional. If the smart contract is created without an admin key, then such a key can never be added, and its bytecode will be immutable.
type ContractUpdateTransactionBody struct {
	ContractID           *ContractID `protobuf:"bytes,1,opt,name=contractID,proto3" json:"contractID,omitempty"`
	ExpirationTime       *Timestamp  `protobuf:"bytes,2,opt,name=expirationTime,proto3" json:"expirationTime,omitempty"`
	AdminKey             *Key        `protobuf:"bytes,3,opt,name=adminKey,proto3" json:"adminKey,omitempty"`
	ProxyAccountID       *AccountID  `protobuf:"bytes,6,opt,name=proxyAccountID,proto3" json:"proxyAccountID,omitempty"`
	AutoRenewPeriod      *Duration   `protobuf:"bytes,7,opt,name=autoRenewPeriod,proto3" json:"autoRenewPeriod,omitempty"`
	FileID               *FileID     `protobuf:"bytes,8,opt,name=fileID,proto3" json:"fileID,omitempty"`
	Memo                 string      `protobuf:"bytes,9,opt,name=memo,proto3" json:"memo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ContractUpdateTransactionBody) Reset()         { *m = ContractUpdateTransactionBody{} }
func (m *ContractUpdateTransactionBody) String() string { return proto.CompactTextString(m) }
func (*ContractUpdateTransactionBody) ProtoMessage()    {}
func (*ContractUpdateTransactionBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_c32c440d581b3cd8, []int{0}
}

func (m *ContractUpdateTransactionBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractUpdateTransactionBody.Unmarshal(m, b)
}
func (m *ContractUpdateTransactionBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractUpdateTransactionBody.Marshal(b, m, deterministic)
}
func (m *ContractUpdateTransactionBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractUpdateTransactionBody.Merge(m, src)
}
func (m *ContractUpdateTransactionBody) XXX_Size() int {
	return xxx_messageInfo_ContractUpdateTransactionBody.Size(m)
}
func (m *ContractUpdateTransactionBody) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractUpdateTransactionBody.DiscardUnknown(m)
}

var xxx_messageInfo_ContractUpdateTransactionBody proto.InternalMessageInfo

func (m *ContractUpdateTransactionBody) GetContractID() *ContractID {
	if m != nil {
		return m.ContractID
	}
	return nil
}

func (m *ContractUpdateTransactionBody) GetExpirationTime() *Timestamp {
	if m != nil {
		return m.ExpirationTime
	}
	return nil
}

func (m *ContractUpdateTransactionBody) GetAdminKey() *Key {
	if m != nil {
		return m.AdminKey
	}
	return nil
}

func (m *ContractUpdateTransactionBody) GetProxyAccountID() *AccountID {
	if m != nil {
		return m.ProxyAccountID
	}
	return nil
}

func (m *ContractUpdateTransactionBody) GetAutoRenewPeriod() *Duration {
	if m != nil {
		return m.AutoRenewPeriod
	}
	return nil
}

func (m *ContractUpdateTransactionBody) GetFileID() *FileID {
	if m != nil {
		return m.FileID
	}
	return nil
}

func (m *ContractUpdateTransactionBody) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func init() {
	proto.RegisterType((*ContractUpdateTransactionBody)(nil), "proto.ContractUpdateTransactionBody")
}

func init() {
	proto.RegisterFile("proto/ContractUpdate.proto", fileDescriptor_c32c440d581b3cd8)
}

var fileDescriptor_c32c440d581b3cd8 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x1c, 0xc5, 0xe9, 0xd4, 0xb9, 0x45, 0x74, 0x1a, 0x54, 0x4a, 0x41, 0x18, 0x82, 0xb2, 0xcb, 0x5a,
	0xd4, 0x8b, 0x1e, 0xdd, 0x86, 0x38, 0x76, 0x19, 0x61, 0x5e, 0xbc, 0x65, 0x69, 0x5c, 0x83, 0xa6,
	0xff, 0x90, 0xa6, 0xb8, 0x7e, 0x5f, 0x3f, 0x88, 0x34, 0x89, 0xc5, 0xf5, 0x94, 0xf2, 0xde, 0xef,
	0xd7, 0x17, 0x82, 0x22, 0xa5, 0xc1, 0x40, 0x32, 0x85, 0xdc, 0x68, 0xca, 0xcc, 0x9b, 0x4a, 0xa9,
	0xe1, 0xb1, 0x0d, 0xf1, 0x81, 0x3d, 0xa2, 0x4b, 0x87, 0x4c, 0x68, 0x21, 0xd8, 0xaa, 0x52, 0xbc,
	0x70, 0x75, 0x74, 0xee, 0xf2, 0x59, 0xa9, 0xa9, 0x11, 0x90, 0xfb, 0xf4, 0xc2, 0xa5, 0x2b, 0x21,
	0x79, 0x61, 0xa8, 0x54, 0x2e, 0xbe, 0xfe, 0xe9, 0xa0, 0xab, 0xdd, 0x91, 0x95, 0xa6, 0x79, 0x41,
	0x59, 0xad, 0x4e, 0x20, 0xad, 0xf0, 0x1d, 0x42, 0xcc, 0x03, 0xf3, 0x59, 0x18, 0x0c, 0x83, 0xd1,
	0xd1, 0xfd, 0x99, 0xb3, 0xe3, 0x69, 0x53, 0x90, 0x7f, 0x10, 0x7e, 0x44, 0x27, 0x7c, 0xab, 0x84,
	0xdb, 0xaf, 0x17, 0xc3, 0x8e, 0xd5, 0x4e, 0xbd, 0xd6, 0x5c, 0x82, 0xb4, 0x38, 0x7c, 0x8b, 0x7a,
	0x34, 0x95, 0x22, 0x5f, 0xf0, 0x2a, 0xdc, 0xb3, 0x0e, 0xf2, 0xce, 0x82, 0x57, 0xa4, 0xe9, 0xea,
	0x05, 0xa5, 0x61, 0x5b, 0x3d, 0x33, 0x06, 0x65, 0x5e, 0x5f, 0xac, 0xbb, 0xb3, 0xd0, 0xe4, 0xa4,
	0xc5, 0xe1, 0x27, 0x34, 0xa0, 0xa5, 0x01, 0xc2, 0x73, 0xfe, 0xbd, 0xe4, 0x5a, 0x40, 0x1a, 0x1e,
	0x5a, 0x75, 0xe0, 0xd5, 0xbf, 0x77, 0x23, 0x6d, 0x0e, 0xdf, 0xa0, 0xee, 0x87, 0xf8, 0xe2, 0xf3,
	0x59, 0xd8, 0xb3, 0xc6, 0xb1, 0x37, 0x5e, 0x6c, 0x48, 0x7c, 0x89, 0x31, 0xda, 0x97, 0x5c, 0x42,
	0xd8, 0x1f, 0x06, 0xa3, 0x3e, 0xb1, 0xdf, 0x93, 0x57, 0x14, 0x31, 0x90, 0x71, 0xc6, 0x53, 0xae,
	0x69, 0x9c, 0xd1, 0x22, 0xdb, 0x68, 0xaa, 0x32, 0xf7, 0x83, 0x65, 0xf0, 0x3e, 0xda, 0x08, 0x93,
	0x95, 0xeb, 0x98, 0x81, 0x4c, 0x9a, 0x36, 0x71, 0xf8, 0xb8, 0x48, 0x3f, 0xc7, 0x1b, 0x48, 0x2c,
	0xbb, 0xee, 0xda, 0xe3, 0xe1, 0x37, 0x00, 0x00, 0xff, 0xff, 0x17, 0x3c, 0xd7, 0x8d, 0x21, 0x02,
	0x00, 0x00,
}
