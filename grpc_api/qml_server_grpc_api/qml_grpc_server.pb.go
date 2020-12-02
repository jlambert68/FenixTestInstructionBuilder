// Code generated by protoc-gen-go. DO NOT EDIT.
// source: qml_grpc_server.proto

package qml_server_grpc_api

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
	return fileDescriptor_f4a7182eb9dd55f4, []int{0}
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
	return fileDescriptor_f4a7182eb9dd55f4, []int{1}
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

type BackendServerInformation struct {
	BackendServerIp      string   `protobuf:"bytes,1,opt,name=backend_server_ip,json=backendServerIp,proto3" json:"backend_server_ip,omitempty"`
	BackendServerPort    string   `protobuf:"bytes,2,opt,name=backend_server_port,json=backendServerPort,proto3" json:"backend_server_port,omitempty"`
	BackendServerUuid    string   `protobuf:"bytes,3,opt,name=backend_server_uuid,json=backendServerUuid,proto3" json:"backend_server_uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BackendServerInformation) Reset()         { *m = BackendServerInformation{} }
func (m *BackendServerInformation) String() string { return proto.CompactTextString(m) }
func (*BackendServerInformation) ProtoMessage()    {}
func (*BackendServerInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4a7182eb9dd55f4, []int{2}
}

func (m *BackendServerInformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BackendServerInformation.Unmarshal(m, b)
}
func (m *BackendServerInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BackendServerInformation.Marshal(b, m, deterministic)
}
func (m *BackendServerInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackendServerInformation.Merge(m, src)
}
func (m *BackendServerInformation) XXX_Size() int {
	return xxx_messageInfo_BackendServerInformation.Size(m)
}
func (m *BackendServerInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_BackendServerInformation.DiscardUnknown(m)
}

var xxx_messageInfo_BackendServerInformation proto.InternalMessageInfo

func (m *BackendServerInformation) GetBackendServerIp() string {
	if m != nil {
		return m.BackendServerIp
	}
	return ""
}

func (m *BackendServerInformation) GetBackendServerPort() string {
	if m != nil {
		return m.BackendServerPort
	}
	return ""
}

func (m *BackendServerInformation) GetBackendServerUuid() string {
	if m != nil {
		return m.BackendServerUuid
	}
	return ""
}

func init() {
	proto.RegisterType((*EmptyParameter)(nil), "qml_server_grpc_api.EmptyParameter")
	proto.RegisterType((*AckNackResponse)(nil), "qml_server_grpc_api.AckNackResponse")
	proto.RegisterType((*BackendServerInformation)(nil), "qml_server_grpc_api.BackendServerInformation")
}

func init() { proto.RegisterFile("qml_grpc_server.proto", fileDescriptor_f4a7182eb9dd55f4) }

var fileDescriptor_f4a7182eb9dd55f4 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x8d, 0x82, 0xb6, 0x7b, 0x30, 0x76, 0x8b, 0x10, 0x7a, 0x2a, 0x41, 0xa1, 0x08, 0xe6,
	0xa0, 0x4f, 0xa0, 0x20, 0xa5, 0x17, 0x89, 0x51, 0xcf, 0x65, 0x33, 0x19, 0x65, 0xd9, 0xee, 0x1f,
	0x67, 0x37, 0x82, 0x17, 0xdf, 0xc1, 0x97, 0xf0, 0x39, 0x25, 0x89, 0x15, 0x52, 0xe2, 0x71, 0xbe,
	0xf9, 0xed, 0x7e, 0x33, 0xdf, 0xb0, 0xd3, 0x37, 0xbd, 0x59, 0xbf, 0x92, 0x83, 0xb5, 0x47, 0x7a,
	0x47, 0xca, 0x1c, 0xd9, 0x60, 0xf9, 0xb4, 0x91, 0x3b, 0xa5, 0xeb, 0x0a, 0x27, 0xd3, 0x13, 0x76,
	0x7c, 0xa7, 0x5d, 0xf8, 0xc8, 0x05, 0x09, 0x8d, 0x01, 0x29, 0x5d, 0xb2, 0xf8, 0x06, 0xd4, 0xbd,
	0x00, 0x55, 0xa0, 0x77, 0xd6, 0x78, 0xe4, 0x09, 0x3b, 0x12, 0xa0, 0x8c, 0x00, 0x95, 0x44, 0xf3,
	0x68, 0x31, 0x2a, 0xb6, 0x25, 0x9f, 0xb1, 0x11, 0x58, 0xad, 0xd1, 0x04, 0x9f, 0xec, 0xcf, 0xa3,
	0xc5, 0xb8, 0xf8, 0xab, 0xd3, 0xef, 0x88, 0x25, 0xb7, 0x02, 0x14, 0x9a, 0xea, 0xb1, 0x75, 0x5d,
	0x99, 0x17, 0x4b, 0x5a, 0x04, 0x69, 0x0d, 0xbf, 0x60, 0x93, 0xb2, 0xeb, 0x6d, 0x47, 0x92, 0xae,
	0xfd, 0x7c, 0x5c, 0xc4, 0x65, 0xef, 0x91, 0xe3, 0x19, 0x9b, 0xee, 0xb0, 0xce, 0x52, 0xf8, 0xf5,
	0x9b, 0xf4, 0xe8, 0xdc, 0x52, 0x18, 0xe0, 0xeb, 0x5a, 0x56, 0xc9, 0xc1, 0x00, 0xff, 0x5c, 0xcb,
	0xea, 0xea, 0x2b, 0x62, 0xf1, 0x83, 0xde, 0x2c, 0xc9, 0x41, 0xa3, 0x4a, 0x40, 0xcf, 0x3f, 0xd9,
	0xf9, 0x13, 0xfa, 0xb0, 0x32, 0x3e, 0x50, 0x0d, 0xcd, 0xc8, 0xfd, 0x55, 0x72, 0x61, 0xaa, 0xd6,
	0xec, 0x32, 0x1b, 0x88, 0x35, 0xfb, 0x6f, 0xef, 0xd9, 0xd9, 0x20, 0xbe, 0x13, 0x78, 0xba, 0x57,
	0x1e, 0xb6, 0x37, 0xbb, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xad, 0x71, 0x4b, 0xa3, 0xcc, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QmlGrpcServicesClient is the client API for QmlGrpcServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QmlGrpcServicesClient interface {
	//Main coordinator sends task to worker with this service
	TestInstructionBackendServerIPandPort(ctx context.Context, in *BackendServerInformation, opts ...grpc.CallOption) (*AckNackResponse, error)
}

type qmlGrpcServicesClient struct {
	cc *grpc.ClientConn
}

func NewQmlGrpcServicesClient(cc *grpc.ClientConn) QmlGrpcServicesClient {
	return &qmlGrpcServicesClient{cc}
}

func (c *qmlGrpcServicesClient) TestInstructionBackendServerIPandPort(ctx context.Context, in *BackendServerInformation, opts ...grpc.CallOption) (*AckNackResponse, error) {
	out := new(AckNackResponse)
	err := c.cc.Invoke(ctx, "/qml_server_grpc_api.QmlGrpcServices/TestInstructionBackendServerIPandPort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QmlGrpcServicesServer is the server API for QmlGrpcServices service.
type QmlGrpcServicesServer interface {
	//Main coordinator sends task to worker with this service
	TestInstructionBackendServerIPandPort(context.Context, *BackendServerInformation) (*AckNackResponse, error)
}

// UnimplementedQmlGrpcServicesServer can be embedded to have forward compatible implementations.
type UnimplementedQmlGrpcServicesServer struct {
}

func (*UnimplementedQmlGrpcServicesServer) TestInstructionBackendServerIPandPort(ctx context.Context, req *BackendServerInformation) (*AckNackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestInstructionBackendServerIPandPort not implemented")
}

func RegisterQmlGrpcServicesServer(s *grpc.Server, srv QmlGrpcServicesServer) {
	s.RegisterService(&_QmlGrpcServices_serviceDesc, srv)
}

func _QmlGrpcServices_TestInstructionBackendServerIPandPort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackendServerInformation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QmlGrpcServicesServer).TestInstructionBackendServerIPandPort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qml_server_grpc_api.QmlGrpcServices/TestInstructionBackendServerIPandPort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QmlGrpcServicesServer).TestInstructionBackendServerIPandPort(ctx, req.(*BackendServerInformation))
	}
	return interceptor(ctx, in, info, handler)
}

var _QmlGrpcServices_serviceDesc = grpc.ServiceDesc{
	ServiceName: "qml_server_grpc_api.QmlGrpcServices",
	HandlerType: (*QmlGrpcServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestInstructionBackendServerIPandPort",
			Handler:    _QmlGrpcServices_TestInstructionBackendServerIPandPort_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "qml_grpc_server.proto",
}
