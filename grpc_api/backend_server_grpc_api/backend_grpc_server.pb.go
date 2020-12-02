// Code generated by protoc-gen-go. DO NOT EDIT.
// source: backend_grpc_server.proto

package backend_server_grpc_api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Parameter used for Empty inputs
type EmptyParameter struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyParameter) Reset()         { *m = EmptyParameter{} }
func (m *EmptyParameter) String() string { return proto.CompactTextString(m) }
func (*EmptyParameter) ProtoMessage()    {}
func (*EmptyParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_83ae52c3330092bb, []int{0}
}

func (m *EmptyParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyParameter.Unmarshal(m, b)
}
func (m *EmptyParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyParameter.Marshal(b, m, deterministic)
}
func (m *EmptyParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyParameter.Merge(m, src)
}
func (m *EmptyParameter) XXX_Size() int {
	return xxx_messageInfo_EmptyParameter.Size(m)
}
func (m *EmptyParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyParameter.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyParameter proto.InternalMessageInfo

// Ack/Nack- Response message with comment
type AckNackResponse struct {
	Acknack              bool     `protobuf:"varint,1,opt,name=acknack,proto3" json:"acknack,omitempty"`
	Comments             string   `protobuf:"bytes,2,opt,name=comments,proto3" json:"comments,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AckNackResponse) Reset()         { *m = AckNackResponse{} }
func (m *AckNackResponse) String() string { return proto.CompactTextString(m) }
func (*AckNackResponse) ProtoMessage()    {}
func (*AckNackResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83ae52c3330092bb, []int{1}
}

func (m *AckNackResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AckNackResponse.Unmarshal(m, b)
}
func (m *AckNackResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AckNackResponse.Marshal(b, m, deterministic)
}
func (m *AckNackResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AckNackResponse.Merge(m, src)
}
func (m *AckNackResponse) XXX_Size() int {
	return xxx_messageInfo_AckNackResponse.Size(m)
}
func (m *AckNackResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AckNackResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AckNackResponse proto.InternalMessageInfo

func (m *AckNackResponse) GetAcknack() bool {
	if m != nil {
		return m.Acknack
	}
	return false
}

func (m *AckNackResponse) GetComments() string {
	if m != nil {
		return m.Comments
	}
	return ""
}

// Message for GUID
type GuidResponse struct {
	Guid                 string   `protobuf:"bytes,1,opt,name=guid,proto3" json:"guid,omitempty"`
	Acknack              bool     `protobuf:"varint,2,opt,name=acknack,proto3" json:"acknack,omitempty"`
	Comments             string   `protobuf:"bytes,3,opt,name=comments,proto3" json:"comments,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GuidResponse) Reset()         { *m = GuidResponse{} }
func (m *GuidResponse) String() string { return proto.CompactTextString(m) }
func (*GuidResponse) ProtoMessage()    {}
func (*GuidResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83ae52c3330092bb, []int{2}
}

func (m *GuidResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GuidResponse.Unmarshal(m, b)
}
func (m *GuidResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GuidResponse.Marshal(b, m, deterministic)
}
func (m *GuidResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GuidResponse.Merge(m, src)
}
func (m *GuidResponse) XXX_Size() int {
	return xxx_messageInfo_GuidResponse.Size(m)
}
func (m *GuidResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GuidResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GuidResponse proto.InternalMessageInfo

func (m *GuidResponse) GetGuid() string {
	if m != nil {
		return m.Guid
	}
	return ""
}

func (m *GuidResponse) GetAcknack() bool {
	if m != nil {
		return m.Acknack
	}
	return false
}

func (m *GuidResponse) GetComments() string {
	if m != nil {
		return m.Comments
	}
	return ""
}

// Json string message with all Plugins that was saved by Server
type PluginQmlModelFromServerResponse struct {
	JsonStringForPluginQmlModel string   `protobuf:"bytes,1,opt,name=jsonStringForPluginQmlModel,proto3" json:"jsonStringForPluginQmlModel,omitempty"`
	Acknack                     bool     `protobuf:"varint,2,opt,name=acknack,proto3" json:"acknack,omitempty"`
	Comments                    string   `protobuf:"bytes,3,opt,name=comments,proto3" json:"comments,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *PluginQmlModelFromServerResponse) Reset()         { *m = PluginQmlModelFromServerResponse{} }
func (m *PluginQmlModelFromServerResponse) String() string { return proto.CompactTextString(m) }
func (*PluginQmlModelFromServerResponse) ProtoMessage()    {}
func (*PluginQmlModelFromServerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83ae52c3330092bb, []int{3}
}

func (m *PluginQmlModelFromServerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginQmlModelFromServerResponse.Unmarshal(m, b)
}
func (m *PluginQmlModelFromServerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginQmlModelFromServerResponse.Marshal(b, m, deterministic)
}
func (m *PluginQmlModelFromServerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginQmlModelFromServerResponse.Merge(m, src)
}
func (m *PluginQmlModelFromServerResponse) XXX_Size() int {
	return xxx_messageInfo_PluginQmlModelFromServerResponse.Size(m)
}
func (m *PluginQmlModelFromServerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginQmlModelFromServerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PluginQmlModelFromServerResponse proto.InternalMessageInfo

func (m *PluginQmlModelFromServerResponse) GetJsonStringForPluginQmlModel() string {
	if m != nil {
		return m.JsonStringForPluginQmlModel
	}
	return ""
}

func (m *PluginQmlModelFromServerResponse) GetAcknack() bool {
	if m != nil {
		return m.Acknack
	}
	return false
}

func (m *PluginQmlModelFromServerResponse) GetComments() string {
	if m != nil {
		return m.Comments
	}
	return ""
}

// Json string message with all Domains that was saved by Server
type DomainQmlModelFromServerResponse struct {
	JsonStringForDomainQmlModel string   `protobuf:"bytes,1,opt,name=jsonStringForDomainQmlModel,proto3" json:"jsonStringForDomainQmlModel,omitempty"`
	Acknack                     bool     `protobuf:"varint,2,opt,name=acknack,proto3" json:"acknack,omitempty"`
	Comments                    string   `protobuf:"bytes,3,opt,name=comments,proto3" json:"comments,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *DomainQmlModelFromServerResponse) Reset()         { *m = DomainQmlModelFromServerResponse{} }
func (m *DomainQmlModelFromServerResponse) String() string { return proto.CompactTextString(m) }
func (*DomainQmlModelFromServerResponse) ProtoMessage()    {}
func (*DomainQmlModelFromServerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83ae52c3330092bb, []int{4}
}

func (m *DomainQmlModelFromServerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DomainQmlModelFromServerResponse.Unmarshal(m, b)
}
func (m *DomainQmlModelFromServerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DomainQmlModelFromServerResponse.Marshal(b, m, deterministic)
}
func (m *DomainQmlModelFromServerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DomainQmlModelFromServerResponse.Merge(m, src)
}
func (m *DomainQmlModelFromServerResponse) XXX_Size() int {
	return xxx_messageInfo_DomainQmlModelFromServerResponse.Size(m)
}
func (m *DomainQmlModelFromServerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DomainQmlModelFromServerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DomainQmlModelFromServerResponse proto.InternalMessageInfo

func (m *DomainQmlModelFromServerResponse) GetJsonStringForDomainQmlModel() string {
	if m != nil {
		return m.JsonStringForDomainQmlModel
	}
	return ""
}

func (m *DomainQmlModelFromServerResponse) GetAcknack() bool {
	if m != nil {
		return m.Acknack
	}
	return false
}

func (m *DomainQmlModelFromServerResponse) GetComments() string {
	if m != nil {
		return m.Comments
	}
	return ""
}

func init() {
	proto.RegisterType((*EmptyParameter)(nil), "backend_server_grpc_api.EmptyParameter")
	proto.RegisterType((*AckNackResponse)(nil), "backend_server_grpc_api.AckNackResponse")
	proto.RegisterType((*GuidResponse)(nil), "backend_server_grpc_api.GuidResponse")
	proto.RegisterType((*PluginQmlModelFromServerResponse)(nil), "backend_server_grpc_api.PluginQmlModelFromServerResponse")
	proto.RegisterType((*DomainQmlModelFromServerResponse)(nil), "backend_server_grpc_api.DomainQmlModelFromServerResponse")
}

func init() { proto.RegisterFile("backend_grpc_server.proto", fileDescriptor_83ae52c3330092bb) }

var fileDescriptor_83ae52c3330092bb = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x41, 0x4f, 0xc2, 0x30,
	0x14, 0xc7, 0x19, 0x18, 0x85, 0x4a, 0xd4, 0xf4, 0x22, 0xe0, 0x85, 0x34, 0x31, 0x72, 0xe2, 0xa0,
	0x27, 0x6f, 0x62, 0x94, 0xc5, 0x44, 0x0d, 0x0e, 0x0f, 0x7a, 0xc2, 0xd2, 0xbd, 0x2c, 0x75, 0x6b,
	0xbb, 0xb4, 0x1d, 0x89, 0x47, 0xbf, 0x86, 0x1f, 0xd3, 0x4f, 0x60, 0xd8, 0x84, 0x30, 0xe2, 0x04,
	0x35, 0xde, 0xf6, 0xd6, 0xff, 0xfb, 0xbf, 0x7f, 0xf3, 0x7e, 0x29, 0x6a, 0x8e, 0x29, 0x0b, 0x41,
	0xfa, 0xa3, 0x40, 0xc7, 0x6c, 0x64, 0x40, 0x4f, 0x40, 0x77, 0x63, 0xad, 0xac, 0xc2, 0xfb, 0xb3,
	0xa3, 0xec, 0x6f, 0xa6, 0xa0, 0x31, 0x27, 0x7b, 0x68, 0xe7, 0x52, 0xc4, 0xf6, 0x65, 0x40, 0x35,
	0x15, 0x60, 0x41, 0x13, 0x17, 0xed, 0xf6, 0x58, 0x78, 0x4b, 0x59, 0xe8, 0x81, 0x89, 0x95, 0x34,
	0x80, 0x1b, 0x68, 0x8b, 0xb2, 0x50, 0x52, 0x16, 0x36, 0x9c, 0xb6, 0xd3, 0xa9, 0x7a, 0xb3, 0x12,
	0xb7, 0x50, 0x95, 0x29, 0x21, 0x40, 0x5a, 0xd3, 0x28, 0xb7, 0x9d, 0x4e, 0xcd, 0x9b, 0xd7, 0xe4,
	0x01, 0xd5, 0xdd, 0x84, 0xfb, 0x73, 0x17, 0x8c, 0x36, 0x82, 0x84, 0xfb, 0xa9, 0x45, 0xcd, 0x4b,
	0xbf, 0x17, 0x9d, 0xcb, 0xc5, 0xce, 0x95, 0x25, 0xe7, 0x37, 0x07, 0xb5, 0x07, 0x51, 0x12, 0x70,
	0x79, 0x27, 0xa2, 0x1b, 0xe5, 0x43, 0xd4, 0xd7, 0x4a, 0x0c, 0xd3, 0xab, 0xcd, 0xc7, 0x9d, 0xa1,
	0x83, 0x67, 0xa3, 0xe4, 0xd0, 0x6a, 0x2e, 0x83, 0xbe, 0xd2, 0xf9, 0x86, 0xcf, 0x14, 0xdf, 0x49,
	0xfe, 0x10, 0xee, 0x42, 0x09, 0xfa, 0xa3, 0x70, 0xf9, 0x86, 0x2f, 0xc3, 0xe5, 0x25, 0xbf, 0x0b,
	0x77, 0xfc, 0x5e, 0x41, 0xe4, 0x1e, 0x8c, 0xbd, 0x92, 0xc6, 0xea, 0x84, 0x59, 0xae, 0xe4, 0x79,
	0x46, 0x86, 0xab, 0x63, 0x36, 0x0d, 0xc9, 0x19, 0x18, 0x3c, 0x46, 0xdb, 0x3d, 0x0d, 0x8f, 0x2a,
	0xe9, 0x45, 0x7c, 0x02, 0xf8, 0xa8, 0x5b, 0x80, 0x4f, 0x37, 0xcf, 0x4e, 0xab, 0x53, 0x28, 0x5c,
	0x42, 0x8a, 0x94, 0xf0, 0x13, 0xaa, 0xbb, 0x20, 0x41, 0x53, 0x0b, 0x53, 0x4c, 0xd6, 0x1f, 0x72,
	0x58, 0x28, 0x5c, 0xc4, 0x8d, 0x94, 0xf0, 0xab, 0x83, 0x9a, 0xd7, 0x8a, 0xfa, 0xd9, 0x5a, 0x97,
	0x56, 0xb1, 0xfe, 0xbc, 0xd3, 0x42, 0xe1, 0x2a, 0x06, 0x17, 0x32, 0x64, 0xdb, 0xfb, 0x87, 0x0c,
	0xab, 0x50, 0x23, 0xa5, 0xf1, 0x66, 0xfa, 0x06, 0x9c, 0x7c, 0x04, 0x00, 0x00, 0xff, 0xff, 0x6c,
	0x4b, 0x68, 0x70, 0x20, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestInstructionBackendGrpcServicesClient is the client API for TestInstructionBackendGrpcServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestInstructionBackendGrpcServicesClient interface {
	//QML Server can check if Backend server is still alive with this service
	AreYouAlive(ctx context.Context, in *EmptyParameter, opts ...grpc.CallOption) (*AckNackResponse, error)
	// ** QML Server forwards this questions to Backend Server **
	// Generates an new GUID, used for identification
	GenerateGuid(ctx context.Context, in *EmptyParameter, opts ...grpc.CallOption) (*GuidResponse, error)
	// ** QML Server forwards this questions to Backend Server **
	// Return Plugins saved by the server
	LoadPluginModelFromServer(ctx context.Context, in *EmptyParameter, opts ...grpc.CallOption) (*PluginQmlModelFromServerResponse, error)
	// ** QML Server forwards this questions to Backend Server **
	// Return Domains saved by the server
	LoadDomainModelFromServer(ctx context.Context, in *EmptyParameter, opts ...grpc.CallOption) (*DomainQmlModelFromServerResponse, error)
}

type testInstructionBackendGrpcServicesClient struct {
	cc *grpc.ClientConn
}

func NewTestInstructionBackendGrpcServicesClient(cc *grpc.ClientConn) TestInstructionBackendGrpcServicesClient {
	return &testInstructionBackendGrpcServicesClient{cc}
}

func (c *testInstructionBackendGrpcServicesClient) AreYouAlive(ctx context.Context, in *EmptyParameter, opts ...grpc.CallOption) (*AckNackResponse, error) {
	out := new(AckNackResponse)
	err := c.cc.Invoke(ctx, "/backend_server_grpc_api.TestInstructionBackendGrpcServices/AreYouAlive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testInstructionBackendGrpcServicesClient) GenerateGuid(ctx context.Context, in *EmptyParameter, opts ...grpc.CallOption) (*GuidResponse, error) {
	out := new(GuidResponse)
	err := c.cc.Invoke(ctx, "/backend_server_grpc_api.TestInstructionBackendGrpcServices/GenerateGuid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testInstructionBackendGrpcServicesClient) LoadPluginModelFromServer(ctx context.Context, in *EmptyParameter, opts ...grpc.CallOption) (*PluginQmlModelFromServerResponse, error) {
	out := new(PluginQmlModelFromServerResponse)
	err := c.cc.Invoke(ctx, "/backend_server_grpc_api.TestInstructionBackendGrpcServices/LoadPluginModelFromServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testInstructionBackendGrpcServicesClient) LoadDomainModelFromServer(ctx context.Context, in *EmptyParameter, opts ...grpc.CallOption) (*DomainQmlModelFromServerResponse, error) {
	out := new(DomainQmlModelFromServerResponse)
	err := c.cc.Invoke(ctx, "/backend_server_grpc_api.TestInstructionBackendGrpcServices/LoadDomainModelFromServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestInstructionBackendGrpcServicesServer is the server API for TestInstructionBackendGrpcServices service.
type TestInstructionBackendGrpcServicesServer interface {
	//QML Server can check if Backend server is still alive with this service
	AreYouAlive(context.Context, *EmptyParameter) (*AckNackResponse, error)
	// ** QML Server forwards this questions to Backend Server **
	// Generates an new GUID, used for identification
	GenerateGuid(context.Context, *EmptyParameter) (*GuidResponse, error)
	// ** QML Server forwards this questions to Backend Server **
	// Return Plugins saved by the server
	LoadPluginModelFromServer(context.Context, *EmptyParameter) (*PluginQmlModelFromServerResponse, error)
	// ** QML Server forwards this questions to Backend Server **
	// Return Domains saved by the server
	LoadDomainModelFromServer(context.Context, *EmptyParameter) (*DomainQmlModelFromServerResponse, error)
}

// UnimplementedTestInstructionBackendGrpcServicesServer can be embedded to have forward compatible implementations.
type UnimplementedTestInstructionBackendGrpcServicesServer struct {
}

func (*UnimplementedTestInstructionBackendGrpcServicesServer) AreYouAlive(ctx context.Context, req *EmptyParameter) (*AckNackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AreYouAlive not implemented")
}
func (*UnimplementedTestInstructionBackendGrpcServicesServer) GenerateGuid(ctx context.Context, req *EmptyParameter) (*GuidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateGuid not implemented")
}
func (*UnimplementedTestInstructionBackendGrpcServicesServer) LoadPluginModelFromServer(ctx context.Context, req *EmptyParameter) (*PluginQmlModelFromServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadPluginModelFromServer not implemented")
}
func (*UnimplementedTestInstructionBackendGrpcServicesServer) LoadDomainModelFromServer(ctx context.Context, req *EmptyParameter) (*DomainQmlModelFromServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadDomainModelFromServer not implemented")
}

func RegisterTestInstructionBackendGrpcServicesServer(s *grpc.Server, srv TestInstructionBackendGrpcServicesServer) {
	s.RegisterService(&_TestInstructionBackendGrpcServices_serviceDesc, srv)
}

func _TestInstructionBackendGrpcServices_AreYouAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestInstructionBackendGrpcServicesServer).AreYouAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backend_server_grpc_api.TestInstructionBackendGrpcServices/AreYouAlive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestInstructionBackendGrpcServicesServer).AreYouAlive(ctx, req.(*EmptyParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestInstructionBackendGrpcServices_GenerateGuid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestInstructionBackendGrpcServicesServer).GenerateGuid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backend_server_grpc_api.TestInstructionBackendGrpcServices/GenerateGuid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestInstructionBackendGrpcServicesServer).GenerateGuid(ctx, req.(*EmptyParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestInstructionBackendGrpcServices_LoadPluginModelFromServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestInstructionBackendGrpcServicesServer).LoadPluginModelFromServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backend_server_grpc_api.TestInstructionBackendGrpcServices/LoadPluginModelFromServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestInstructionBackendGrpcServicesServer).LoadPluginModelFromServer(ctx, req.(*EmptyParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestInstructionBackendGrpcServices_LoadDomainModelFromServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestInstructionBackendGrpcServicesServer).LoadDomainModelFromServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backend_server_grpc_api.TestInstructionBackendGrpcServices/LoadDomainModelFromServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestInstructionBackendGrpcServicesServer).LoadDomainModelFromServer(ctx, req.(*EmptyParameter))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestInstructionBackendGrpcServices_serviceDesc = grpc.ServiceDesc{
	ServiceName: "backend_server_grpc_api.TestInstructionBackendGrpcServices",
	HandlerType: (*TestInstructionBackendGrpcServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AreYouAlive",
			Handler:    _TestInstructionBackendGrpcServices_AreYouAlive_Handler,
		},
		{
			MethodName: "GenerateGuid",
			Handler:    _TestInstructionBackendGrpcServices_GenerateGuid_Handler,
		},
		{
			MethodName: "LoadPluginModelFromServer",
			Handler:    _TestInstructionBackendGrpcServices_LoadPluginModelFromServer_Handler,
		},
		{
			MethodName: "LoadDomainModelFromServer",
			Handler:    _TestInstructionBackendGrpcServices_LoadDomainModelFromServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backend_grpc_server.proto",
}