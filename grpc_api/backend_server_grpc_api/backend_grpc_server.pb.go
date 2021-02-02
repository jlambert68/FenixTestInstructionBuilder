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
	return fileDescriptor_83ae52c3330092bb, []int{0}
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

// Json string message with QML model that will be load from Server
type QmlModeToLoadlFromServerResponse struct {
	JsonStringForPluginQmlModel  string                       `protobuf:"bytes,1,opt,name=jsonStringForPluginQmlModel,proto3" json:"jsonStringForPluginQmlModel,omitempty"`
	QmlModelTypeToAndFromBackend QMLModelTypeToAndFromBackend `protobuf:"varint,2,opt,name=qmlModelTypeToAndFromBackend,proto3,enum=grpc_common.QMLModelTypeToAndFromBackend" json:"qmlModelTypeToAndFromBackend,omitempty"`
	Acknack                      bool                         `protobuf:"varint,3,opt,name=acknack,proto3" json:"acknack,omitempty"`
	Comments                     string                       `protobuf:"bytes,4,opt,name=comments,proto3" json:"comments,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}                     `json:"-"`
	XXX_unrecognized             []byte                       `json:"-"`
	XXX_sizecache                int32                        `json:"-"`
}

func (m *QmlModeToLoadlFromServerResponse) Reset()         { *m = QmlModeToLoadlFromServerResponse{} }
func (m *QmlModeToLoadlFromServerResponse) String() string { return proto.CompactTextString(m) }
func (*QmlModeToLoadlFromServerResponse) ProtoMessage()    {}
func (*QmlModeToLoadlFromServerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83ae52c3330092bb, []int{1}
}

func (m *QmlModeToLoadlFromServerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QmlModeToLoadlFromServerResponse.Unmarshal(m, b)
}
func (m *QmlModeToLoadlFromServerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QmlModeToLoadlFromServerResponse.Marshal(b, m, deterministic)
}
func (m *QmlModeToLoadlFromServerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QmlModeToLoadlFromServerResponse.Merge(m, src)
}
func (m *QmlModeToLoadlFromServerResponse) XXX_Size() int {
	return xxx_messageInfo_QmlModeToLoadlFromServerResponse.Size(m)
}
func (m *QmlModeToLoadlFromServerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QmlModeToLoadlFromServerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QmlModeToLoadlFromServerResponse proto.InternalMessageInfo

func (m *QmlModeToLoadlFromServerResponse) GetJsonStringForPluginQmlModel() string {
	if m != nil {
		return m.JsonStringForPluginQmlModel
	}
	return ""
}

func (m *QmlModeToLoadlFromServerResponse) GetQmlModelTypeToAndFromBackend() QMLModelTypeToAndFromBackend {
	if m != nil {
		return m.QmlModelTypeToAndFromBackend
	}
	return QMLModelTypeToAndFromBackend_DomainMode
}

func (m *QmlModeToLoadlFromServerResponse) GetAcknack() bool {
	if m != nil {
		return m.Acknack
	}
	return false
}

func (m *QmlModeToLoadlFromServerResponse) GetComments() string {
	if m != nil {
		return m.Comments
	}
	return ""
}

// Json string message with QML model that will be saved at Server
type QmlModelToSaveAtServerRequest struct {
	JsonStringForPluginQmlModel  string                       `protobuf:"bytes,1,opt,name=jsonStringForPluginQmlModel,proto3" json:"jsonStringForPluginQmlModel,omitempty"`
	QmlModelTypeToAndFromBackend QMLModelTypeToAndFromBackend `protobuf:"varint,2,opt,name=qmlModelTypeToAndFromBackend,proto3,enum=grpc_common.QMLModelTypeToAndFromBackend" json:"qmlModelTypeToAndFromBackend,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}                     `json:"-"`
	XXX_unrecognized             []byte                       `json:"-"`
	XXX_sizecache                int32                        `json:"-"`
}

func (m *QmlModelToSaveAtServerRequest) Reset()         { *m = QmlModelToSaveAtServerRequest{} }
func (m *QmlModelToSaveAtServerRequest) String() string { return proto.CompactTextString(m) }
func (*QmlModelToSaveAtServerRequest) ProtoMessage()    {}
func (*QmlModelToSaveAtServerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_83ae52c3330092bb, []int{2}
}

func (m *QmlModelToSaveAtServerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QmlModelToSaveAtServerRequest.Unmarshal(m, b)
}
func (m *QmlModelToSaveAtServerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QmlModelToSaveAtServerRequest.Marshal(b, m, deterministic)
}
func (m *QmlModelToSaveAtServerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QmlModelToSaveAtServerRequest.Merge(m, src)
}
func (m *QmlModelToSaveAtServerRequest) XXX_Size() int {
	return xxx_messageInfo_QmlModelToSaveAtServerRequest.Size(m)
}
func (m *QmlModelToSaveAtServerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QmlModelToSaveAtServerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QmlModelToSaveAtServerRequest proto.InternalMessageInfo

func (m *QmlModelToSaveAtServerRequest) GetJsonStringForPluginQmlModel() string {
	if m != nil {
		return m.JsonStringForPluginQmlModel
	}
	return ""
}

func (m *QmlModelToSaveAtServerRequest) GetQmlModelTypeToAndFromBackend() QMLModelTypeToAndFromBackend {
	if m != nil {
		return m.QmlModelTypeToAndFromBackend
	}
	return QMLModelTypeToAndFromBackend_DomainMode
}

func init() {
	proto.RegisterType((*GuidResponse)(nil), "backend_server_grpc_api.GuidResponse")
	proto.RegisterType((*QmlModeToLoadlFromServerResponse)(nil), "backend_server_grpc_api.QmlModeToLoadlFromServerResponse")
	proto.RegisterType((*QmlModelToSaveAtServerRequest)(nil), "backend_server_grpc_api.QmlModelToSaveAtServerRequest")
}

func init() { proto.RegisterFile("backend_grpc_server.proto", fileDescriptor_83ae52c3330092bb) }

var fileDescriptor_83ae52c3330092bb = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x53, 0x4d, 0xab, 0xd3, 0x40,
	0x14, 0x6d, 0x5e, 0x1f, 0xfa, 0x1c, 0xcb, 0x03, 0x67, 0x63, 0xcc, 0x7b, 0x42, 0x09, 0x08, 0xcf,
	0x4d, 0x90, 0x0a, 0x82, 0x3b, 0x23, 0xd8, 0xa2, 0xb4, 0xd2, 0xa6, 0x59, 0xe8, 0xaa, 0x4c, 0x27,
	0xd7, 0x30, 0x26, 0x99, 0x9b, 0xce, 0x4c, 0x0a, 0x5d, 0xb9, 0xf6, 0xd7, 0xf8, 0x57, 0xfc, 0x49,
	0x92, 0xaf, 0xda, 0x8a, 0x69, 0xdd, 0xba, 0xbb, 0x33, 0xf7, 0xe3, 0x9c, 0x7b, 0x38, 0x97, 0x3c,
	0x59, 0x33, 0x9e, 0x80, 0x8c, 0x56, 0xb1, 0xca, 0xf9, 0x4a, 0x83, 0xda, 0x82, 0xf2, 0x72, 0x85,
	0x06, 0xe9, 0xe3, 0x36, 0x55, 0xff, 0xd6, 0x15, 0x2c, 0x17, 0xce, 0xa3, 0x2a, 0xe2, 0x98, 0x65,
	0x28, 0xeb, 0x5a, 0xf7, 0x13, 0x19, 0x4c, 0x0a, 0x11, 0x05, 0xa0, 0x73, 0x94, 0x1a, 0x28, 0x25,
	0x97, 0x71, 0x21, 0x22, 0xdb, 0x1a, 0x5a, 0x77, 0x0f, 0x82, 0x2a, 0xa6, 0x36, 0xb9, 0xcf, 0x78,
	0x22, 0x19, 0x4f, 0xec, 0x8b, 0xa1, 0x75, 0x77, 0x15, 0xb4, 0x4f, 0xea, 0x90, 0xab, 0x72, 0x1a,
	0x48, 0xa3, 0xed, 0x7e, 0xd5, 0xb1, 0x7f, 0xbb, 0xdf, 0x2f, 0xc8, 0x70, 0x91, 0xa5, 0x33, 0x8c,
	0x20, 0xc4, 0x29, 0xb2, 0x28, 0x1d, 0x2b, 0xcc, 0x96, 0x15, 0xa5, 0x3d, 0xdc, 0x1b, 0x72, 0xf3,
	0x55, 0xa3, 0x5c, 0x1a, 0x25, 0x64, 0x3c, 0x46, 0x35, 0x4f, 0x8b, 0x58, 0xc8, 0xa6, 0x2d, 0x6d,
	0x58, 0x9c, 0x2a, 0xa1, 0x19, 0xb9, 0xdd, 0x34, 0x71, 0xb8, 0xcb, 0x21, 0x44, 0x5f, 0x46, 0x25,
	0xd0, 0xdb, 0x5a, 0x83, 0x8a, 0xf1, 0xf5, 0xe8, 0xb9, 0x77, 0xb8, 0xfa, 0x62, 0x36, 0xed, 0x6c,
	0x08, 0x4e, 0x8e, 0x3b, 0xd4, 0xa2, 0xdf, 0xad, 0xc5, 0xe5, 0x1f, 0x5a, 0xfc, 0xb4, 0xc8, 0xd3,
	0x96, 0x71, 0x88, 0x4b, 0xb6, 0x05, 0xdf, 0xb4, 0x4a, 0x6c, 0x0a, 0xd0, 0xe6, 0xbf, 0x13, 0x62,
	0xf4, 0xa3, 0x4f, 0xdc, 0x10, 0xb4, 0x79, 0x2f, 0xb5, 0x51, 0x05, 0x37, 0x02, 0x65, 0x93, 0x9a,
	0xa8, 0x9c, 0x97, 0xcb, 0x09, 0x0e, 0x9a, 0x7e, 0x20, 0x0f, 0x7d, 0x05, 0x9f, 0xb1, 0xf0, 0x53,
	0xb1, 0x05, 0x7a, 0x73, 0x04, 0xff, 0x2e, 0xcb, 0xcd, 0x6e, 0xce, 0x14, 0xcb, 0xc0, 0x80, 0x72,
	0x6e, 0x8f, 0x92, 0x3e, 0x4f, 0x3e, 0x32, 0x9e, 0xb4, 0x56, 0x71, 0x7b, 0x34, 0x24, 0x83, 0x09,
	0x48, 0x50, 0xcc, 0x40, 0xe9, 0xd9, 0xd3, 0xc3, 0x9e, 0x79, 0x1d, 0x57, 0xe0, 0x1d, 0xfa, 0xdd,
	0xed, 0xd1, 0x6f, 0xe4, 0xba, 0x74, 0xe7, 0x6f, 0x73, 0xd2, 0x17, 0xff, 0xac, 0xd1, 0x0c, 0xb4,
	0x66, 0x31, 0x38, 0xaf, 0x3b, 0xc1, 0xce, 0x5d, 0x80, 0xdb, 0xa3, 0x5f, 0xc8, 0xa0, 0x74, 0x44,
	0x88, 0x0d, 0xfc, 0xab, 0x73, 0xc3, 0xfe, 0x6e, 0xa1, 0x73, 0xf2, 0xad, 0xef, 0x55, 0x17, 0xff,
	0xf2, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7a, 0xc3, 0x0f, 0x51, 0x3a, 0x04, 0x00, 0x00,
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
	// Loads data, defined by MessageTypeToAndFromBackend,  from the server
	LoadFromServer(ctx context.Context, in *QMLModelTypeToAndFromBackendMessage, opts ...grpc.CallOption) (*QmlModeToLoadlFromServerResponse, error)
	// ** QML Server forwards this questions to Backend Server **
	// Save data, defined by MessageTypeToAndFromBackend, at the server
	SaveToServer(ctx context.Context, in *QmlModelToSaveAtServerRequest, opts ...grpc.CallOption) (*AckNackResponse, error)
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

func (c *testInstructionBackendGrpcServicesClient) LoadFromServer(ctx context.Context, in *QMLModelTypeToAndFromBackendMessage, opts ...grpc.CallOption) (*QmlModeToLoadlFromServerResponse, error) {
	out := new(QmlModeToLoadlFromServerResponse)
	err := c.cc.Invoke(ctx, "/backend_server_grpc_api.TestInstructionBackendGrpcServices/LoadFromServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testInstructionBackendGrpcServicesClient) SaveToServer(ctx context.Context, in *QmlModelToSaveAtServerRequest, opts ...grpc.CallOption) (*AckNackResponse, error) {
	out := new(AckNackResponse)
	err := c.cc.Invoke(ctx, "/backend_server_grpc_api.TestInstructionBackendGrpcServices/SaveToServer", in, out, opts...)
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
	// Loads data, defined by MessageTypeToAndFromBackend,  from the server
	LoadFromServer(context.Context, *QMLModelTypeToAndFromBackendMessage) (*QmlModeToLoadlFromServerResponse, error)
	// ** QML Server forwards this questions to Backend Server **
	// Save data, defined by MessageTypeToAndFromBackend, at the server
	SaveToServer(context.Context, *QmlModelToSaveAtServerRequest) (*AckNackResponse, error)
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
func (*UnimplementedTestInstructionBackendGrpcServicesServer) LoadFromServer(ctx context.Context, req *QMLModelTypeToAndFromBackendMessage) (*QmlModeToLoadlFromServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadFromServer not implemented")
}
func (*UnimplementedTestInstructionBackendGrpcServicesServer) SaveToServer(ctx context.Context, req *QmlModelToSaveAtServerRequest) (*AckNackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveToServer not implemented")
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

func _TestInstructionBackendGrpcServices_LoadFromServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QMLModelTypeToAndFromBackendMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestInstructionBackendGrpcServicesServer).LoadFromServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backend_server_grpc_api.TestInstructionBackendGrpcServices/LoadFromServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestInstructionBackendGrpcServicesServer).LoadFromServer(ctx, req.(*QMLModelTypeToAndFromBackendMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestInstructionBackendGrpcServices_SaveToServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QmlModelToSaveAtServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestInstructionBackendGrpcServicesServer).SaveToServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backend_server_grpc_api.TestInstructionBackendGrpcServices/SaveToServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestInstructionBackendGrpcServicesServer).SaveToServer(ctx, req.(*QmlModelToSaveAtServerRequest))
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
			MethodName: "LoadFromServer",
			Handler:    _TestInstructionBackendGrpcServices_LoadFromServer_Handler,
		},
		{
			MethodName: "SaveToServer",
			Handler:    _TestInstructionBackendGrpcServices_SaveToServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backend_grpc_server.proto",
}
