// Code generated by protoc-gen-go. DO NOT EDIT.
// source: southbound.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protos "magma/orc8r/cloud/go/protos"
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

type GetMconfigRequest struct {
	HardwareID           string   `protobuf:"bytes,1,opt,name=hardwareID,proto3" json:"hardwareID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMconfigRequest) Reset()         { *m = GetMconfigRequest{} }
func (m *GetMconfigRequest) String() string { return proto.CompactTextString(m) }
func (*GetMconfigRequest) ProtoMessage()    {}
func (*GetMconfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_480661e00faacec1, []int{0}
}

func (m *GetMconfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMconfigRequest.Unmarshal(m, b)
}
func (m *GetMconfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMconfigRequest.Marshal(b, m, deterministic)
}
func (m *GetMconfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMconfigRequest.Merge(m, src)
}
func (m *GetMconfigRequest) XXX_Size() int {
	return xxx_messageInfo_GetMconfigRequest.Size(m)
}
func (m *GetMconfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMconfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMconfigRequest proto.InternalMessageInfo

func (m *GetMconfigRequest) GetHardwareID() string {
	if m != nil {
		return m.HardwareID
	}
	return ""
}

type GetMconfigResponse struct {
	Configs              *protos.GatewayConfigs `protobuf:"bytes,1,opt,name=configs,proto3" json:"configs,omitempty"`
	LogicalID            string                 `protobuf:"bytes,2,opt,name=logicalID,proto3" json:"logicalID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *GetMconfigResponse) Reset()         { *m = GetMconfigResponse{} }
func (m *GetMconfigResponse) String() string { return proto.CompactTextString(m) }
func (*GetMconfigResponse) ProtoMessage()    {}
func (*GetMconfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_480661e00faacec1, []int{1}
}

func (m *GetMconfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMconfigResponse.Unmarshal(m, b)
}
func (m *GetMconfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMconfigResponse.Marshal(b, m, deterministic)
}
func (m *GetMconfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMconfigResponse.Merge(m, src)
}
func (m *GetMconfigResponse) XXX_Size() int {
	return xxx_messageInfo_GetMconfigResponse.Size(m)
}
func (m *GetMconfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMconfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMconfigResponse proto.InternalMessageInfo

func (m *GetMconfigResponse) GetConfigs() *protos.GatewayConfigs {
	if m != nil {
		return m.Configs
	}
	return nil
}

func (m *GetMconfigResponse) GetLogicalID() string {
	if m != nil {
		return m.LogicalID
	}
	return ""
}

func init() {
	proto.RegisterType((*GetMconfigRequest)(nil), "magma.orc8r.configurator.GetMconfigRequest")
	proto.RegisterType((*GetMconfigResponse)(nil), "magma.orc8r.configurator.GetMconfigResponse")
}

func init() {
	proto.RegisterFile("southbound.proto", fileDescriptor_480661e00faacec1)
}

var fileDescriptor_480661e00faacec1 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x1b, 0x0f, 0x6a, 0xc7, 0x8b, 0xdd, 0x83, 0x84, 0x28, 0x5a, 0x72, 0x12, 0x94, 0x0d,
	0xb4, 0x08, 0x9e, 0x3c, 0xd8, 0x42, 0xc9, 0xc1, 0x4b, 0x04, 0x0f, 0xde, 0xa6, 0xc9, 0x9a, 0x06,
	0x92, 0x9d, 0x76, 0x77, 0x43, 0xf1, 0x33, 0xfa, 0xa5, 0x84, 0x1d, 0xb4, 0x2b, 0xfe, 0xc1, 0xd3,
	0xc2, 0x9b, 0xdf, 0xe3, 0xcd, 0xbc, 0x85, 0x63, 0x4b, 0xbd, 0x5b, 0x2d, 0xa9, 0xd7, 0x95, 0x5c,
	0x1b, 0x72, 0x24, 0xe2, 0x0e, 0xeb, 0x0e, 0x25, 0x99, 0xf2, 0xd6, 0xc8, 0x92, 0xf4, 0x4b, 0x53,
	0xf7, 0x06, 0x1d, 0x99, 0x64, 0xec, 0x27, 0x99, 0x9f, 0x64, 0x1e, 0xb6, 0x59, 0xc7, 0x04, 0x7b,
	0x93, 0x8b, 0x1f, 0x88, 0x92, 0xba, 0x8e, 0x34, 0x03, 0xe9, 0x14, 0x46, 0x0b, 0xe5, 0x1e, 0xd8,
	0x54, 0xa8, 0x4d, 0xaf, 0xac, 0x13, 0xe7, 0x00, 0x2b, 0x34, 0xd5, 0x16, 0x8d, 0xca, 0xe7, 0x71,
	0x34, 0x8e, 0x2e, 0x87, 0x45, 0xa0, 0xa4, 0x0d, 0x88, 0xd0, 0x64, 0xd7, 0xa4, 0xad, 0x12, 0x37,
	0x70, 0xc0, 0x8a, 0xf5, 0x96, 0xa3, 0xc9, 0xa9, 0x0c, 0x37, 0x5f, 0xa0, 0x53, 0x5b, 0x7c, 0x9d,
	0x31, 0x52, 0x7c, 0xb0, 0xe2, 0x0c, 0x86, 0x2d, 0xd5, 0x4d, 0x89, 0x6d, 0x3e, 0x8f, 0xf7, 0x7c,
	0xd6, 0x4e, 0x98, 0xbc, 0x45, 0x70, 0xf2, 0xf8, 0xd9, 0xc8, 0x2c, 0xb8, 0x5e, 0xdc, 0x01, 0xec,
	0xb6, 0x10, 0xa3, 0x2f, 0x61, 0x4f, 0xd4, 0x54, 0xc9, 0x5f, 0xf9, 0xe9, 0x40, 0x6c, 0xc2, 0x2b,
	0x72, 0xed, 0x94, 0xd1, 0xd8, 0x8a, 0x2b, 0xf9, 0x5b, 0xdd, 0xf2, 0x5b, 0x51, 0xc9, 0xf5, 0xff,
	0x60, 0x2e, 0x28, 0x1d, 0xdc, 0x1f, 0x3e, 0xef, 0xf3, 0x27, 0x2c, 0xf9, 0x9d, 0xbe, 0x07, 0x00,
	0x00, 0xff, 0xff, 0xb6, 0xde, 0x9a, 0xfe, 0xef, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SouthboundConfiguratorClient is the client API for SouthboundConfigurator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SouthboundConfiguratorClient interface {
	GetMconfig(ctx context.Context, in *protos.Void, opts ...grpc.CallOption) (*protos.GatewayConfigs, error)
	// GetMconfigInternal exists to support the existing streamer mconfig
	// policy. This should be removed when we migrate gateway mconfig updates
	// from streamer to this southbound configurator servicer.
	GetMconfigInternal(ctx context.Context, in *GetMconfigRequest, opts ...grpc.CallOption) (*GetMconfigResponse, error)
}

type southboundConfiguratorClient struct {
	cc grpc.ClientConnInterface
}

func NewSouthboundConfiguratorClient(cc grpc.ClientConnInterface) SouthboundConfiguratorClient {
	return &southboundConfiguratorClient{cc}
}

func (c *southboundConfiguratorClient) GetMconfig(ctx context.Context, in *protos.Void, opts ...grpc.CallOption) (*protos.GatewayConfigs, error) {
	out := new(protos.GatewayConfigs)
	err := c.cc.Invoke(ctx, "/magma.orc8r.configurator.SouthboundConfigurator/GetMconfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *southboundConfiguratorClient) GetMconfigInternal(ctx context.Context, in *GetMconfigRequest, opts ...grpc.CallOption) (*GetMconfigResponse, error) {
	out := new(GetMconfigResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.configurator.SouthboundConfigurator/GetMconfigInternal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SouthboundConfiguratorServer is the server API for SouthboundConfigurator service.
type SouthboundConfiguratorServer interface {
	GetMconfig(context.Context, *protos.Void) (*protos.GatewayConfigs, error)
	// GetMconfigInternal exists to support the existing streamer mconfig
	// policy. This should be removed when we migrate gateway mconfig updates
	// from streamer to this southbound configurator servicer.
	GetMconfigInternal(context.Context, *GetMconfigRequest) (*GetMconfigResponse, error)
}

// UnimplementedSouthboundConfiguratorServer can be embedded to have forward compatible implementations.
type UnimplementedSouthboundConfiguratorServer struct {
}

func (*UnimplementedSouthboundConfiguratorServer) GetMconfig(ctx context.Context, req *protos.Void) (*protos.GatewayConfigs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMconfig not implemented")
}
func (*UnimplementedSouthboundConfiguratorServer) GetMconfigInternal(ctx context.Context, req *GetMconfigRequest) (*GetMconfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMconfigInternal not implemented")
}

func RegisterSouthboundConfiguratorServer(s *grpc.Server, srv SouthboundConfiguratorServer) {
	s.RegisterService(&_SouthboundConfigurator_serviceDesc, srv)
}

func _SouthboundConfigurator_GetMconfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SouthboundConfiguratorServer).GetMconfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.configurator.SouthboundConfigurator/GetMconfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SouthboundConfiguratorServer).GetMconfig(ctx, req.(*protos.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _SouthboundConfigurator_GetMconfigInternal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMconfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SouthboundConfiguratorServer).GetMconfigInternal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.configurator.SouthboundConfigurator/GetMconfigInternal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SouthboundConfiguratorServer).GetMconfigInternal(ctx, req.(*GetMconfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SouthboundConfigurator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.configurator.SouthboundConfigurator",
	HandlerType: (*SouthboundConfiguratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMconfig",
			Handler:    _SouthboundConfigurator_GetMconfig_Handler,
		},
		{
			MethodName: "GetMconfigInternal",
			Handler:    _SouthboundConfigurator_GetMconfigInternal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "southbound.proto",
}
