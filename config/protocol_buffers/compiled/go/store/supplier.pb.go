// Code generated by protoc-gen-go. DO NOT EDIT.
// source: supplier.proto

package store

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

type Supplier struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Country              string   `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Supplier) Reset()         { *m = Supplier{} }
func (m *Supplier) String() string { return proto.CompactTextString(m) }
func (*Supplier) ProtoMessage()    {}
func (*Supplier) Descriptor() ([]byte, []int) {
	return fileDescriptor_86038919225469b2, []int{0}
}

func (m *Supplier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Supplier.Unmarshal(m, b)
}
func (m *Supplier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Supplier.Marshal(b, m, deterministic)
}
func (m *Supplier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Supplier.Merge(m, src)
}
func (m *Supplier) XXX_Size() int {
	return xxx_messageInfo_Supplier.Size(m)
}
func (m *Supplier) XXX_DiscardUnknown() {
	xxx_messageInfo_Supplier.DiscardUnknown(m)
}

var xxx_messageInfo_Supplier proto.InternalMessageInfo

func (m *Supplier) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Supplier) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func init() {
	proto.RegisterType((*Supplier)(nil), "store.Supplier")
}

func init() { proto.RegisterFile("supplier.proto", fileDescriptor_86038919225469b2) }

var fileDescriptor_86038919225469b2 = []byte{
	// 92 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x2e, 0x2d, 0x28,
	0xc8, 0xc9, 0x4c, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d, 0x2e, 0xc9, 0x2f,
	0x4a, 0x55, 0xb2, 0xe0, 0xe2, 0x08, 0x86, 0x4a, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6,
	0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x42, 0x12, 0x5c, 0xec, 0xc9, 0xf9, 0xa5,
	0x79, 0x25, 0x45, 0x95, 0x12, 0x4c, 0x60, 0x61, 0x18, 0x37, 0x89, 0x0d, 0x6c, 0x8e, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x47, 0xf6, 0x35, 0x34, 0x59, 0x00, 0x00, 0x00,
}