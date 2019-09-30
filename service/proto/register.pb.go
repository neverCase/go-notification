// Code generated by protoc-gen-go. DO NOT EDIT.
// source: register.proto

package register_push_v1

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

type RegisterRequest struct {
	ExternalAddr         string   `protobuf:"bytes,1,opt,name=externalAddr,proto3" json:"externalAddr,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1303fe8288f4efb6, []int{0}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetExternalAddr() string {
	if m != nil {
		return m.ExternalAddr
	}
	return ""
}

func (m *RegisterRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type RegisterResponse struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1303fe8288f4efb6, []int{1}
}

func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RegisterResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type PingRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1303fe8288f4efb6, []int{2}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

func (m *PingRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type PongResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PongResponse) Reset()         { *m = PongResponse{} }
func (m *PongResponse) String() string { return proto.CompactTextString(m) }
func (*PongResponse) ProtoMessage()    {}
func (*PongResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1303fe8288f4efb6, []int{3}
}

func (m *PongResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PongResponse.Unmarshal(m, b)
}
func (m *PongResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PongResponse.Marshal(b, m, deterministic)
}
func (m *PongResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PongResponse.Merge(m, src)
}
func (m *PongResponse) XXX_Size() int {
	return xxx_messageInfo_PongResponse.Size(m)
}
func (m *PongResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PongResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PongResponse proto.InternalMessageInfo

func (m *PongResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type CommonRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Addr                 string   `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommonRequest) Reset()         { *m = CommonRequest{} }
func (m *CommonRequest) String() string { return proto.CompactTextString(m) }
func (*CommonRequest) ProtoMessage()    {}
func (*CommonRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1303fe8288f4efb6, []int{4}
}

func (m *CommonRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonRequest.Unmarshal(m, b)
}
func (m *CommonRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonRequest.Marshal(b, m, deterministic)
}
func (m *CommonRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonRequest.Merge(m, src)
}
func (m *CommonRequest) XXX_Size() int {
	return xxx_messageInfo_CommonRequest.Size(m)
}
func (m *CommonRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommonRequest proto.InternalMessageInfo

func (m *CommonRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CommonRequest) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type GetAllocatedNodeResponse struct {
	Addr                 string   `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllocatedNodeResponse) Reset()         { *m = GetAllocatedNodeResponse{} }
func (m *GetAllocatedNodeResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllocatedNodeResponse) ProtoMessage()    {}
func (*GetAllocatedNodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1303fe8288f4efb6, []int{5}
}

func (m *GetAllocatedNodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllocatedNodeResponse.Unmarshal(m, b)
}
func (m *GetAllocatedNodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllocatedNodeResponse.Marshal(b, m, deterministic)
}
func (m *GetAllocatedNodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllocatedNodeResponse.Merge(m, src)
}
func (m *GetAllocatedNodeResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllocatedNodeResponse.Size(m)
}
func (m *GetAllocatedNodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllocatedNodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllocatedNodeResponse proto.InternalMessageInfo

func (m *GetAllocatedNodeResponse) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "register.push.v1.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "register.push.v1.RegisterResponse")
	proto.RegisterType((*PingRequest)(nil), "register.push.v1.PingRequest")
	proto.RegisterType((*PongResponse)(nil), "register.push.v1.PongResponse")
	proto.RegisterType((*CommonRequest)(nil), "register.push.v1.CommonRequest")
	proto.RegisterType((*GetAllocatedNodeResponse)(nil), "register.push.v1.GetAllocatedNodeResponse")
}

func init() { proto.RegisterFile("register.proto", fileDescriptor_1303fe8288f4efb6) }

var fileDescriptor_1303fe8288f4efb6 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4d, 0x4f, 0xb3, 0x40,
	0x10, 0xc7, 0x81, 0xf4, 0x79, 0xaa, 0x23, 0x56, 0x32, 0x27, 0x42, 0x52, 0x5f, 0xf6, 0xd4, 0x78,
	0x20, 0xd1, 0x5e, 0xbd, 0x34, 0x3d, 0x70, 0xd3, 0x8a, 0x9f, 0x60, 0x75, 0x27, 0x48, 0x02, 0x2c,
	0xb2, 0x5b, 0x5f, 0x3e, 0xac, 0xdf, 0xc5, 0x40, 0x8b, 0xbc, 0xd9, 0xea, 0x6d, 0x67, 0xe6, 0x3f,
	0xff, 0x99, 0xf9, 0x65, 0x61, 0x52, 0x50, 0x14, 0x2b, 0x4d, 0x85, 0x9f, 0x17, 0x52, 0x4b, 0x74,
	0x9a, 0x78, 0xad, 0x9e, 0xfd, 0xd7, 0x2b, 0x76, 0x07, 0x27, 0xe1, 0x36, 0x17, 0xd2, 0xcb, 0x9a,
	0x94, 0x46, 0x06, 0x36, 0xbd, 0x6b, 0x2a, 0x32, 0x9e, 0x2c, 0x84, 0x28, 0x5c, 0xf3, 0xdc, 0x9c,
	0x1d, 0x86, 0x9d, 0x1c, 0xba, 0x30, 0x4e, 0x49, 0x29, 0x1e, 0x91, 0x6b, 0x55, 0xe5, 0x3a, 0x64,
	0x37, 0xe0, 0x34, 0x86, 0x2a, 0x97, 0x99, 0x22, 0x9c, 0x80, 0x15, 0x8b, 0xca, 0xe7, 0x5f, 0x68,
	0xc5, 0x62, 0x4f, 0xf7, 0x14, 0x8e, 0x56, 0x71, 0x16, 0xd5, 0xab, 0xf4, 0x1a, 0xd9, 0x0c, 0xec,
	0x95, 0x2c, 0xcb, 0x5b, 0xe3, 0x96, 0x91, 0xd9, 0x35, 0x9a, 0xc3, 0xf1, 0x52, 0xa6, 0xa9, 0xcc,
	0x76, 0x58, 0x21, 0xc2, 0x88, 0x97, 0xd7, 0x6d, 0x16, 0xa8, 0xde, 0xcc, 0x07, 0x37, 0x20, 0xbd,
	0x48, 0x12, 0xf9, 0xc4, 0x35, 0x89, 0x5b, 0x29, 0xe8, 0x7b, 0x54, 0xad, 0x37, 0x1b, 0xfd, 0xf5,
	0xa7, 0x05, 0xe3, 0x80, 0x6b, 0x7a, 0xe3, 0x1f, 0xf8, 0x00, 0x07, 0xf5, 0xdd, 0x78, 0xe1, 0xf7,
	0x39, 0xfb, 0x3d, 0xc8, 0x1e, 0xdb, 0x27, 0xd9, 0x8c, 0x64, 0x06, 0x06, 0x30, 0x2a, 0x71, 0xe0,
	0x74, 0xa8, 0x6e, 0x61, 0xf2, 0x4e, 0x7f, 0x28, 0xb7, 0x30, 0x31, 0x03, 0x39, 0x38, 0xfd, 0xcb,
	0xf0, 0x6c, 0xd8, 0xd5, 0x41, 0xe6, 0x5d, 0x0e, 0x05, 0xbb, 0xf0, 0x30, 0x03, 0xef, 0xc1, 0x5e,
	0xca, 0x34, 0x4f, 0x48, 0xd3, 0xdf, 0xec, 0x7f, 0xdd, 0xfa, 0xf1, 0x7f, 0xf5, 0x6b, 0xe7, 0x5f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x10, 0xca, 0x59, 0x30, 0xc7, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GatewayClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongResponse, error)
	GetAllocatedNode(ctx context.Context, in *CommonRequest, opts ...grpc.CallOption) (*GetAllocatedNodeResponse, error)
	CompleteNode(ctx context.Context, in *CommonRequest, opts ...grpc.CallOption) (*PongResponse, error)
}

type gatewayClient struct {
	cc *grpc.ClientConn
}

func NewGatewayClient(cc *grpc.ClientConn) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/register.push.v1.Gateway/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongResponse, error) {
	out := new(PongResponse)
	err := c.cc.Invoke(ctx, "/register.push.v1.Gateway/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAllocatedNode(ctx context.Context, in *CommonRequest, opts ...grpc.CallOption) (*GetAllocatedNodeResponse, error) {
	out := new(GetAllocatedNodeResponse)
	err := c.cc.Invoke(ctx, "/register.push.v1.Gateway/GetAllocatedNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CompleteNode(ctx context.Context, in *CommonRequest, opts ...grpc.CallOption) (*PongResponse, error) {
	out := new(PongResponse)
	err := c.cc.Invoke(ctx, "/register.push.v1.Gateway/CompleteNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
type GatewayServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Ping(context.Context, *PingRequest) (*PongResponse, error)
	GetAllocatedNode(context.Context, *CommonRequest) (*GetAllocatedNodeResponse, error)
	CompleteNode(context.Context, *CommonRequest) (*PongResponse, error)
}

// UnimplementedGatewayServer can be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (*UnimplementedGatewayServer) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedGatewayServer) Ping(ctx context.Context, req *PingRequest) (*PongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedGatewayServer) GetAllocatedNode(ctx context.Context, req *CommonRequest) (*GetAllocatedNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllocatedNode not implemented")
}
func (*UnimplementedGatewayServer) CompleteNode(ctx context.Context, req *CommonRequest) (*PongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteNode not implemented")
}

func RegisterGatewayServer(s *grpc.Server, srv GatewayServer) {
	s.RegisterService(&_Gateway_serviceDesc, srv)
}

func _Gateway_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/register.push.v1.Gateway/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/register.push.v1.Gateway/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAllocatedNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAllocatedNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/register.push.v1.Gateway/GetAllocatedNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAllocatedNode(ctx, req.(*CommonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CompleteNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CompleteNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/register.push.v1.Gateway/CompleteNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CompleteNode(ctx, req.(*CommonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gateway_serviceDesc = grpc.ServiceDesc{
	ServiceName: "register.push.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Gateway_Register_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Gateway_Ping_Handler,
		},
		{
			MethodName: "GetAllocatedNode",
			Handler:    _Gateway_GetAllocatedNode_Handler,
		},
		{
			MethodName: "CompleteNode",
			Handler:    _Gateway_CompleteNode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "register.proto",
}
