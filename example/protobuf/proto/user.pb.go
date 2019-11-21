// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user.proto

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

type Msg1 struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   int32    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Msg1) Reset()         { *m = Msg1{} }
func (m *Msg1) String() string { return proto.CompactTextString(m) }
func (*Msg1) ProtoMessage()    {}
func (*Msg1) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{0}
}

func (m *Msg1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Msg1.Unmarshal(m, b)
}
func (m *Msg1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Msg1.Marshal(b, m, deterministic)
}
func (m *Msg1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg1.Merge(m, src)
}
func (m *Msg1) XXX_Size() int {
	return xxx_messageInfo_Msg1.Size(m)
}
func (m *Msg1) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg1.DiscardUnknown(m)
}

var xxx_messageInfo_Msg1 proto.InternalMessageInfo

func (m *Msg1) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Msg1) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Msg2 struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Alias                string   `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`
	Id                   int32    `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Msg2) Reset()         { *m = Msg2{} }
func (m *Msg2) String() string { return proto.CompactTextString(m) }
func (*Msg2) ProtoMessage()    {}
func (*Msg2) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{1}
}

func (m *Msg2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Msg2.Unmarshal(m, b)
}
func (m *Msg2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Msg2.Marshal(b, m, deterministic)
}
func (m *Msg2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg2.Merge(m, src)
}
func (m *Msg2) XXX_Size() int {
	return xxx_messageInfo_Msg2.Size(m)
}
func (m *Msg2) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg2.DiscardUnknown(m)
}

var xxx_messageInfo_Msg2 proto.InternalMessageInfo

func (m *Msg2) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Msg2) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *Msg2) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Response struct {
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*Msg1)(nil), "proto.Msg1")
	proto.RegisterType((*Msg2)(nil), "proto.Msg2")
	proto.RegisterType((*Response)(nil), "proto.Response")
}

func init() { proto.RegisterFile("proto/user.proto", fileDescriptor_d570e3e37e5899c5) }

var fileDescriptor_d570e3e37e5899c5 = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x2d, 0x4e, 0x2d, 0xd2, 0x03, 0x33, 0x85, 0x58, 0xc1, 0x94, 0x92, 0x16, 0x17,
	0x8b, 0x6f, 0x71, 0xba, 0xa1, 0x90, 0x10, 0x17, 0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02,
	0xa3, 0x06, 0x67, 0x10, 0x98, 0x2d, 0xc4, 0xc7, 0xc5, 0x94, 0x99, 0x22, 0xc1, 0xa4, 0xc0, 0xa8,
	0xc1, 0x1a, 0xc4, 0x94, 0x99, 0xa2, 0xe4, 0x00, 0x56, 0x6b, 0x84, 0x55, 0xad, 0x08, 0x17, 0x6b,
	0x62, 0x4e, 0x66, 0x62, 0x31, 0x58, 0x39, 0x67, 0x10, 0x84, 0x03, 0x35, 0x81, 0x19, 0x6e, 0x82,
	0x0a, 0x17, 0x47, 0x50, 0x6a, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x04, 0x17, 0x7b, 0x72,
	0x7e, 0x5e, 0x49, 0x6a, 0x5e, 0x09, 0xd4, 0x20, 0x18, 0x37, 0x89, 0x0d, 0xec, 0x34, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x7e, 0xf7, 0x80, 0x5f, 0xb5, 0x00, 0x00, 0x00,
}
