// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orc8r/protos/redis.proto

package protos

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

type RedisState struct {
	SerializedMsg        []byte   `protobuf:"bytes,1,opt,name=serialized_msg,json=serializedMsg,proto3" json:"serialized_msg,omitempty"`
	Version              uint64   `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RedisState) Reset()         { *m = RedisState{} }
func (m *RedisState) String() string { return proto.CompactTextString(m) }
func (*RedisState) ProtoMessage()    {}
func (*RedisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6de3b144d47785a, []int{0}
}

func (m *RedisState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisState.Unmarshal(m, b)
}
func (m *RedisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisState.Marshal(b, m, deterministic)
}
func (m *RedisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisState.Merge(m, src)
}
func (m *RedisState) XXX_Size() int {
	return xxx_messageInfo_RedisState.Size(m)
}
func (m *RedisState) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisState.DiscardUnknown(m)
}

var xxx_messageInfo_RedisState proto.InternalMessageInfo

func (m *RedisState) GetSerializedMsg() []byte {
	if m != nil {
		return m.SerializedMsg
	}
	return nil
}

func (m *RedisState) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func init() {
	proto.RegisterType((*RedisState)(nil), "magma.orc8r.RedisState")
}

func init() {
	proto.RegisterFile("orc8r/protos/redis.proto", fileDescriptor_e6de3b144d47785a)
}

var fileDescriptor_e6de3b144d47785a = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc8, 0x2f, 0x4a, 0xb6,
	0x28, 0xd2, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x2f, 0xd6, 0x2f, 0x4a, 0x4d, 0xc9, 0x2c, 0xd6, 0x03,
	0x73, 0x84, 0xb8, 0x73, 0x13, 0xd3, 0x73, 0x13, 0xf5, 0xc0, 0xf2, 0x4a, 0xbe, 0x5c, 0x5c, 0x41,
	0x20, 0xb9, 0xe0, 0x92, 0xc4, 0x92, 0x54, 0x21, 0x55, 0x2e, 0xbe, 0xe2, 0xd4, 0xa2, 0xcc, 0xc4,
	0x9c, 0xcc, 0xaa, 0xd4, 0x94, 0xf8, 0xdc, 0xe2, 0x74, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x9e, 0x20,
	0x5e, 0x84, 0xa8, 0x6f, 0x71, 0xba, 0x90, 0x04, 0x17, 0x7b, 0x59, 0x6a, 0x51, 0x71, 0x66, 0x7e,
	0x9e, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x4b, 0x10, 0x8c, 0xeb, 0x24, 0x1b, 0x25, 0x0d, 0x36, 0x5d,
	0x1f, 0x62, 0x7b, 0x72, 0x4e, 0x7e, 0x69, 0x8a, 0x7e, 0x7a, 0x3e, 0xd4, 0x19, 0x49, 0x6c, 0x60,
	0xda, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xc8, 0xf2, 0xd7, 0x0f, 0x9d, 0x00, 0x00, 0x00,
}
