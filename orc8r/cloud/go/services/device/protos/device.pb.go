// Code generated by protoc-gen-go. DO NOT EDIT.
// source: device.proto

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

type PhysicalEntity struct {
	// Globally unique identifier per type (MAC/SN)
	DeviceID string `protobuf:"bytes,1,opt,name=deviceID,proto3" json:"deviceID,omitempty"`
	// Used to deserialize info
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// Any other information (manufacturer, location, owner, etc)
	Info                 []byte   `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PhysicalEntity) Reset()         { *m = PhysicalEntity{} }
func (m *PhysicalEntity) String() string { return proto.CompactTextString(m) }
func (*PhysicalEntity) ProtoMessage()    {}
func (*PhysicalEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_870276a56ac00da5, []int{0}
}

func (m *PhysicalEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PhysicalEntity.Unmarshal(m, b)
}
func (m *PhysicalEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PhysicalEntity.Marshal(b, m, deterministic)
}
func (m *PhysicalEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhysicalEntity.Merge(m, src)
}
func (m *PhysicalEntity) XXX_Size() int {
	return xxx_messageInfo_PhysicalEntity.Size(m)
}
func (m *PhysicalEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_PhysicalEntity.DiscardUnknown(m)
}

var xxx_messageInfo_PhysicalEntity proto.InternalMessageInfo

func (m *PhysicalEntity) GetDeviceID() string {
	if m != nil {
		return m.DeviceID
	}
	return ""
}

func (m *PhysicalEntity) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PhysicalEntity) GetInfo() []byte {
	if m != nil {
		return m.Info
	}
	return nil
}

type RegisterOrUpdateDevicesRequest struct {
	NetworkID            string            `protobuf:"bytes,1,opt,name=networkID,proto3" json:"networkID,omitempty"`
	Entities             []*PhysicalEntity `protobuf:"bytes,2,rep,name=entities,proto3" json:"entities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RegisterOrUpdateDevicesRequest) Reset()         { *m = RegisterOrUpdateDevicesRequest{} }
func (m *RegisterOrUpdateDevicesRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterOrUpdateDevicesRequest) ProtoMessage()    {}
func (*RegisterOrUpdateDevicesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_870276a56ac00da5, []int{1}
}

func (m *RegisterOrUpdateDevicesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterOrUpdateDevicesRequest.Unmarshal(m, b)
}
func (m *RegisterOrUpdateDevicesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterOrUpdateDevicesRequest.Marshal(b, m, deterministic)
}
func (m *RegisterOrUpdateDevicesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterOrUpdateDevicesRequest.Merge(m, src)
}
func (m *RegisterOrUpdateDevicesRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterOrUpdateDevicesRequest.Size(m)
}
func (m *RegisterOrUpdateDevicesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterOrUpdateDevicesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterOrUpdateDevicesRequest proto.InternalMessageInfo

func (m *RegisterOrUpdateDevicesRequest) GetNetworkID() string {
	if m != nil {
		return m.NetworkID
	}
	return ""
}

func (m *RegisterOrUpdateDevicesRequest) GetEntities() []*PhysicalEntity {
	if m != nil {
		return m.Entities
	}
	return nil
}

type DeviceID struct {
	DeviceID             string   `protobuf:"bytes,1,opt,name=deviceID,proto3" json:"deviceID,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceID) Reset()         { *m = DeviceID{} }
func (m *DeviceID) String() string { return proto.CompactTextString(m) }
func (*DeviceID) ProtoMessage()    {}
func (*DeviceID) Descriptor() ([]byte, []int) {
	return fileDescriptor_870276a56ac00da5, []int{2}
}

func (m *DeviceID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceID.Unmarshal(m, b)
}
func (m *DeviceID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceID.Marshal(b, m, deterministic)
}
func (m *DeviceID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceID.Merge(m, src)
}
func (m *DeviceID) XXX_Size() int {
	return xxx_messageInfo_DeviceID.Size(m)
}
func (m *DeviceID) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceID.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceID proto.InternalMessageInfo

func (m *DeviceID) GetDeviceID() string {
	if m != nil {
		return m.DeviceID
	}
	return ""
}

func (m *DeviceID) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type GetDeviceInfoRequest struct {
	NetworkID            string      `protobuf:"bytes,1,opt,name=networkID,proto3" json:"networkID,omitempty"`
	DeviceIDs            []*DeviceID `protobuf:"bytes,2,rep,name=deviceIDs,proto3" json:"deviceIDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetDeviceInfoRequest) Reset()         { *m = GetDeviceInfoRequest{} }
func (m *GetDeviceInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetDeviceInfoRequest) ProtoMessage()    {}
func (*GetDeviceInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_870276a56ac00da5, []int{3}
}

func (m *GetDeviceInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeviceInfoRequest.Unmarshal(m, b)
}
func (m *GetDeviceInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeviceInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetDeviceInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeviceInfoRequest.Merge(m, src)
}
func (m *GetDeviceInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetDeviceInfoRequest.Size(m)
}
func (m *GetDeviceInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeviceInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeviceInfoRequest proto.InternalMessageInfo

func (m *GetDeviceInfoRequest) GetNetworkID() string {
	if m != nil {
		return m.NetworkID
	}
	return ""
}

func (m *GetDeviceInfoRequest) GetDeviceIDs() []*DeviceID {
	if m != nil {
		return m.DeviceIDs
	}
	return nil
}

type GetDeviceInfoResponse struct {
	// A map of device IDs to corresponding PhysicalEntity structure
	DeviceMap            map[string]*PhysicalEntity `protobuf:"bytes,1,rep,name=deviceMap,proto3" json:"deviceMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *GetDeviceInfoResponse) Reset()         { *m = GetDeviceInfoResponse{} }
func (m *GetDeviceInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetDeviceInfoResponse) ProtoMessage()    {}
func (*GetDeviceInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_870276a56ac00da5, []int{4}
}

func (m *GetDeviceInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeviceInfoResponse.Unmarshal(m, b)
}
func (m *GetDeviceInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeviceInfoResponse.Marshal(b, m, deterministic)
}
func (m *GetDeviceInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeviceInfoResponse.Merge(m, src)
}
func (m *GetDeviceInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetDeviceInfoResponse.Size(m)
}
func (m *GetDeviceInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeviceInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeviceInfoResponse proto.InternalMessageInfo

func (m *GetDeviceInfoResponse) GetDeviceMap() map[string]*PhysicalEntity {
	if m != nil {
		return m.DeviceMap
	}
	return nil
}

type DeleteDevicesRequest struct {
	NetworkID            string      `protobuf:"bytes,1,opt,name=networkID,proto3" json:"networkID,omitempty"`
	DeviceIDs            []*DeviceID `protobuf:"bytes,2,rep,name=deviceIDs,proto3" json:"deviceIDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DeleteDevicesRequest) Reset()         { *m = DeleteDevicesRequest{} }
func (m *DeleteDevicesRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteDevicesRequest) ProtoMessage()    {}
func (*DeleteDevicesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_870276a56ac00da5, []int{5}
}

func (m *DeleteDevicesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteDevicesRequest.Unmarshal(m, b)
}
func (m *DeleteDevicesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteDevicesRequest.Marshal(b, m, deterministic)
}
func (m *DeleteDevicesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteDevicesRequest.Merge(m, src)
}
func (m *DeleteDevicesRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteDevicesRequest.Size(m)
}
func (m *DeleteDevicesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteDevicesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteDevicesRequest proto.InternalMessageInfo

func (m *DeleteDevicesRequest) GetNetworkID() string {
	if m != nil {
		return m.NetworkID
	}
	return ""
}

func (m *DeleteDevicesRequest) GetDeviceIDs() []*DeviceID {
	if m != nil {
		return m.DeviceIDs
	}
	return nil
}

func init() {
	proto.RegisterType((*PhysicalEntity)(nil), "magma.orc8r.device.PhysicalEntity")
	proto.RegisterType((*RegisterOrUpdateDevicesRequest)(nil), "magma.orc8r.device.RegisterOrUpdateDevicesRequest")
	proto.RegisterType((*DeviceID)(nil), "magma.orc8r.device.DeviceID")
	proto.RegisterType((*GetDeviceInfoRequest)(nil), "magma.orc8r.device.GetDeviceInfoRequest")
	proto.RegisterType((*GetDeviceInfoResponse)(nil), "magma.orc8r.device.GetDeviceInfoResponse")
	proto.RegisterMapType((map[string]*PhysicalEntity)(nil), "magma.orc8r.device.GetDeviceInfoResponse.DeviceMapEntry")
	proto.RegisterType((*DeleteDevicesRequest)(nil), "magma.orc8r.device.DeleteDevicesRequest")
}

func init() {
	proto.RegisterFile("device.proto", fileDescriptor_870276a56ac00da5)
}

var fileDescriptor_870276a56ac00da5 = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x41, 0x8f, 0x9a, 0x40,
	0x14, 0x16, 0x6c, 0x0d, 0xbe, 0xaa, 0x6d, 0x27, 0x6d, 0x42, 0x89, 0x69, 0xcc, 0x9c, 0xe8, 0x05,
	0x13, 0x7b, 0x21, 0x1e, 0x7a, 0x68, 0x30, 0x4d, 0x0f, 0xb6, 0x0d, 0x69, 0x4d, 0xe3, 0xa9, 0x14,
	0x9f, 0x2e, 0x51, 0x18, 0x96, 0x19, 0xdd, 0x70, 0xd9, 0xff, 0xb8, 0x7f, 0x66, 0xcf, 0x1b, 0x18,
	0xd0, 0xb0, 0x4b, 0x36, 0x98, 0xdd, 0x13, 0x33, 0x6f, 0xde, 0xf7, 0xbd, 0xef, 0xbd, 0xef, 0x05,
	0xe8, 0xad, 0xf0, 0x10, 0xf8, 0x68, 0xc5, 0x09, 0x13, 0x8c, 0x90, 0xd0, 0xdb, 0x84, 0x9e, 0xc5,
	0x12, 0xdf, 0x4e, 0x2c, 0xf9, 0x62, 0x7c, 0xc8, 0x6f, 0xe3, 0x3c, 0x81, 0x8f, 0x7d, 0x16, 0x86,
	0x2c, 0x92, 0xe9, 0xf4, 0x37, 0x0c, 0x7e, 0x5d, 0xa4, 0x3c, 0xf0, 0xbd, 0xdd, 0x2c, 0x12, 0x81,
	0x48, 0x89, 0x01, 0x9a, 0x84, 0x7d, 0x77, 0x74, 0x65, 0xa4, 0x98, 0x5d, 0xf7, 0x78, 0x27, 0x04,
	0x5e, 0x88, 0x34, 0x46, 0x5d, 0xcd, 0xe3, 0xf9, 0x39, 0x8b, 0x05, 0xd1, 0x9a, 0xe9, 0xed, 0x91,
	0x62, 0xf6, 0xdc, 0xfc, 0x4c, 0xaf, 0xe1, 0xa3, 0x8b, 0x9b, 0x80, 0x0b, 0x4c, 0x7e, 0x26, 0x7f,
	0xe2, 0x95, 0x27, 0xd0, 0xc9, 0x39, 0xb8, 0x8b, 0x97, 0x7b, 0xe4, 0x82, 0x0c, 0xa1, 0x1b, 0xa1,
	0xb8, 0x62, 0xc9, 0xf6, 0x58, 0xe6, 0x14, 0x20, 0x5f, 0x40, 0xc3, 0x4c, 0x4d, 0x80, 0x5c, 0x57,
	0x47, 0x6d, 0xf3, 0xd5, 0x84, 0x5a, 0x0f, 0xfb, 0xb2, 0xaa, 0xca, 0xdd, 0x23, 0x86, 0x4e, 0x41,
	0x73, 0x4a, 0xcd, 0x67, 0xf6, 0x43, 0x63, 0x78, 0xf7, 0x0d, 0x45, 0x01, 0x8f, 0xd6, 0xac, 0x99,
	0xe2, 0x29, 0x74, 0x4b, 0xd6, 0x52, 0xf2, 0xb0, 0x4e, 0x72, 0x29, 0xcb, 0x3d, 0xa5, 0xd3, 0x1b,
	0x05, 0xde, 0xdf, 0x2b, 0xc9, 0x63, 0x16, 0x71, 0x24, 0x8b, 0x92, 0x75, 0xee, 0xc5, 0xba, 0x92,
	0xb3, 0xda, 0x75, 0xac, 0xb5, 0xe8, 0xa2, 0xd6, 0xdc, 0x8b, 0x67, 0x91, 0x48, 0x52, 0xf7, 0x44,
	0x65, 0xfc, 0x83, 0x41, 0xf5, 0x91, 0xbc, 0x81, 0xf6, 0x16, 0xd3, 0xa2, 0xaf, 0xec, 0x48, 0x6c,
	0x78, 0x79, 0xf0, 0x76, 0x7b, 0x39, 0x9c, 0x66, 0x06, 0x48, 0xc0, 0x54, 0xb5, 0x95, 0x6c, 0x8a,
	0x0e, 0xee, 0xf0, 0x4c, 0xdf, 0x9f, 0x30, 0xc5, 0xc9, 0xad, 0x0a, 0x1d, 0x19, 0x27, 0x4b, 0x78,
	0x5d, 0xae, 0x5f, 0x51, 0x9e, 0x4c, 0xea, 0x68, 0x1e, 0xdf, 0x51, 0xe3, 0x6d, 0x05, 0xb3, 0x60,
	0xc1, 0x8a, 0xb6, 0xc8, 0x5f, 0xe8, 0x57, 0x92, 0x9f, 0x8f, 0x79, 0x0d, 0xfd, 0x8a, 0x8f, 0xc4,
	0x6c, 0x60, 0xb5, 0xe4, 0xfb, 0xd4, 0x78, 0x29, 0x68, 0x8b, 0xfc, 0x80, 0x7e, 0xc5, 0x9a, 0xfa,
	0x3a, 0x75, 0xee, 0xd5, 0xea, 0xfe, 0xaa, 0x2d, 0x3b, 0xf2, 0xcf, 0xf2, 0x5f, 0x7e, 0x3f, 0xdf,
	0x05, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x49, 0xca, 0xd2, 0x92, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DeviceClient is the client API for Device service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeviceClient interface {
	RegisterDevices(ctx context.Context, in *RegisterOrUpdateDevicesRequest, opts ...grpc.CallOption) (*protos.Void, error)
	UpdateDevices(ctx context.Context, in *RegisterOrUpdateDevicesRequest, opts ...grpc.CallOption) (*protos.Void, error)
	GetDeviceInfo(ctx context.Context, in *GetDeviceInfoRequest, opts ...grpc.CallOption) (*GetDeviceInfoResponse, error)
	DeleteDevices(ctx context.Context, in *DeleteDevicesRequest, opts ...grpc.CallOption) (*protos.Void, error)
}

type deviceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceClient(cc grpc.ClientConnInterface) DeviceClient {
	return &deviceClient{cc}
}

func (c *deviceClient) RegisterDevices(ctx context.Context, in *RegisterOrUpdateDevicesRequest, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.device.Device/RegisterDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceClient) UpdateDevices(ctx context.Context, in *RegisterOrUpdateDevicesRequest, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.device.Device/UpdateDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceClient) GetDeviceInfo(ctx context.Context, in *GetDeviceInfoRequest, opts ...grpc.CallOption) (*GetDeviceInfoResponse, error) {
	out := new(GetDeviceInfoResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.device.Device/GetDeviceInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceClient) DeleteDevices(ctx context.Context, in *DeleteDevicesRequest, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.device.Device/DeleteDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceServer is the server API for Device service.
type DeviceServer interface {
	RegisterDevices(context.Context, *RegisterOrUpdateDevicesRequest) (*protos.Void, error)
	UpdateDevices(context.Context, *RegisterOrUpdateDevicesRequest) (*protos.Void, error)
	GetDeviceInfo(context.Context, *GetDeviceInfoRequest) (*GetDeviceInfoResponse, error)
	DeleteDevices(context.Context, *DeleteDevicesRequest) (*protos.Void, error)
}

// UnimplementedDeviceServer can be embedded to have forward compatible implementations.
type UnimplementedDeviceServer struct {
}

func (*UnimplementedDeviceServer) RegisterDevices(ctx context.Context, req *RegisterOrUpdateDevicesRequest) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterDevices not implemented")
}
func (*UnimplementedDeviceServer) UpdateDevices(ctx context.Context, req *RegisterOrUpdateDevicesRequest) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDevices not implemented")
}
func (*UnimplementedDeviceServer) GetDeviceInfo(ctx context.Context, req *GetDeviceInfoRequest) (*GetDeviceInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceInfo not implemented")
}
func (*UnimplementedDeviceServer) DeleteDevices(ctx context.Context, req *DeleteDevicesRequest) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDevices not implemented")
}

func RegisterDeviceServer(s *grpc.Server, srv DeviceServer) {
	s.RegisterService(&_Device_serviceDesc, srv)
}

func _Device_RegisterDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterOrUpdateDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).RegisterDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.device.Device/RegisterDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).RegisterDevices(ctx, req.(*RegisterOrUpdateDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Device_UpdateDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterOrUpdateDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).UpdateDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.device.Device/UpdateDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).UpdateDevices(ctx, req.(*RegisterOrUpdateDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Device_GetDeviceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).GetDeviceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.device.Device/GetDeviceInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).GetDeviceInfo(ctx, req.(*GetDeviceInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Device_DeleteDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).DeleteDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.device.Device/DeleteDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).DeleteDevices(ctx, req.(*DeleteDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Device_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.device.Device",
	HandlerType: (*DeviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterDevices",
			Handler:    _Device_RegisterDevices_Handler,
		},
		{
			MethodName: "UpdateDevices",
			Handler:    _Device_UpdateDevices_Handler,
		},
		{
			MethodName: "GetDeviceInfo",
			Handler:    _Device_GetDeviceInfo_Handler,
		},
		{
			MethodName: "DeleteDevices",
			Handler:    _Device_DeleteDevices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "device.proto",
}
