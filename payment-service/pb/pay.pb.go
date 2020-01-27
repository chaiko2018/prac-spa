// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pay.proto

package paymentservice

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

type PayRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Amount               int64    `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PayRequest) Reset()         { *m = PayRequest{} }
func (m *PayRequest) String() string { return proto.CompactTextString(m) }
func (*PayRequest) ProtoMessage()    {}
func (*PayRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0564d675d5c516e0, []int{0}
}

func (m *PayRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayRequest.Unmarshal(m, b)
}
func (m *PayRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayRequest.Marshal(b, m, deterministic)
}
func (m *PayRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayRequest.Merge(m, src)
}
func (m *PayRequest) XXX_Size() int {
	return xxx_messageInfo_PayRequest.Size(m)
}
func (m *PayRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PayRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PayRequest proto.InternalMessageInfo

func (m *PayRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PayRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PayRequest) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *PayRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PayRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type PayResponse struct {
	Paid                 bool     `protobuf:"varint,1,opt,name=paid,proto3" json:"paid,omitempty"`
	Captured             bool     `protobuf:"varint,3,opt,name=captured,proto3" json:"captured,omitempty"`
	Amount               int64    `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PayResponse) Reset()         { *m = PayResponse{} }
func (m *PayResponse) String() string { return proto.CompactTextString(m) }
func (*PayResponse) ProtoMessage()    {}
func (*PayResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0564d675d5c516e0, []int{1}
}

func (m *PayResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayResponse.Unmarshal(m, b)
}
func (m *PayResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayResponse.Marshal(b, m, deterministic)
}
func (m *PayResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayResponse.Merge(m, src)
}
func (m *PayResponse) XXX_Size() int {
	return xxx_messageInfo_PayResponse.Size(m)
}
func (m *PayResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PayResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PayResponse proto.InternalMessageInfo

func (m *PayResponse) GetPaid() bool {
	if m != nil {
		return m.Paid
	}
	return false
}

func (m *PayResponse) GetCaptured() bool {
	if m != nil {
		return m.Captured
	}
	return false
}

func (m *PayResponse) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func init() {
	proto.RegisterType((*PayRequest)(nil), "paymentservice.PayRequest")
	proto.RegisterType((*PayResponse)(nil), "paymentservice.PayResponse")
}

func init() { proto.RegisterFile("pay.proto", fileDescriptor_0564d675d5c516e0) }

var fileDescriptor_0564d675d5c516e0 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x6d, 0x77, 0xb7, 0x74, 0x67, 0x61, 0x0f, 0x83, 0x48, 0xa8, 0x97, 0xd2, 0xd3, 0x9e,
	0x7a, 0xd0, 0x9f, 0xb0, 0x67, 0x41, 0x03, 0xfe, 0x80, 0xb1, 0x19, 0xd6, 0x20, 0x4d, 0x62, 0x92,
	0x0a, 0xbd, 0xf9, 0xd3, 0xc5, 0xd9, 0x45, 0x2b, 0x78, 0x9b, 0x97, 0x79, 0x64, 0xbe, 0xf7, 0x60,
	0x1b, 0x68, 0xee, 0x43, 0xf4, 0xd9, 0xe3, 0x3e, 0xd0, 0x3c, 0xb2, 0xcb, 0x89, 0xe3, 0x87, 0x1d,
	0xb8, 0xfb, 0x2c, 0x00, 0x1e, 0x69, 0xd6, 0xfc, 0x3e, 0x71, 0xca, 0xb8, 0x87, 0xd2, 0x1a, 0x55,
	0xb4, 0xc5, 0x61, 0xa5, 0x4b, 0x6b, 0xf0, 0x1a, 0x36, 0xd9, 0xbf, 0xb1, 0x53, 0x65, 0x5b, 0x1c,
	0xb6, 0xfa, 0x2c, 0xf0, 0x06, 0x2a, 0x1a, 0xfd, 0xe4, 0xb2, 0x5a, 0x89, 0xf3, 0xa2, 0x10, 0x61,
	0xed, 0x68, 0x64, 0xb5, 0x16, 0xb3, 0xcc, 0xd8, 0xc2, 0xce, 0x70, 0x1a, 0xa2, 0x0d, 0xd9, 0x7a,
	0xa7, 0x36, 0xb2, 0x5a, 0x3e, 0x75, 0xcf, 0xb0, 0x13, 0x82, 0x14, 0xbc, 0x4b, 0xfc, 0xfd, 0x49,
	0xa0, 0x0b, 0x44, 0xad, 0x65, 0xc6, 0x06, 0xea, 0x81, 0x42, 0x9e, 0x22, 0x1b, 0x39, 0x59, 0xeb,
	0x1f, 0xbd, 0x80, 0x29, 0x97, 0x30, 0x77, 0x4f, 0x12, 0xec, 0x81, 0x1c, 0x9d, 0x38, 0xe2, 0x11,
	0xaa, 0xe3, 0x2b, 0xc5, 0x13, 0x63, 0xd3, 0xff, 0xad, 0xa0, 0xff, 0x8d, 0xdf, 0xdc, 0xfe, 0xbb,
	0x3b, 0x83, 0x75, 0x57, 0x2f, 0x95, 0x74, 0x78, 0xff, 0x15, 0x00, 0x00, 0xff, 0xff, 0x62, 0x6b,
	0x09, 0x50, 0x50, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PayManagerClient is the client API for PayManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PayManagerClient interface {
	Charge(ctx context.Context, in *PayRequest, opts ...grpc.CallOption) (*PayResponse, error)
}

type payManagerClient struct {
	cc *grpc.ClientConn
}

func NewPayManagerClient(cc *grpc.ClientConn) PayManagerClient {
	return &payManagerClient{cc}
}

func (c *payManagerClient) Charge(ctx context.Context, in *PayRequest, opts ...grpc.CallOption) (*PayResponse, error) {
	out := new(PayResponse)
	err := c.cc.Invoke(ctx, "/paymentservice.PayManager/Charge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PayManagerServer is the server API for PayManager service.
type PayManagerServer interface {
	Charge(context.Context, *PayRequest) (*PayResponse, error)
}

// UnimplementedPayManagerServer can be embedded to have forward compatible implementations.
type UnimplementedPayManagerServer struct {
}

func (*UnimplementedPayManagerServer) Charge(ctx context.Context, req *PayRequest) (*PayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Charge not implemented")
}

func RegisterPayManagerServer(s *grpc.Server, srv PayManagerServer) {
	s.RegisterService(&_PayManager_serviceDesc, srv)
}

func _PayManager_Charge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayManagerServer).Charge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paymentservice.PayManager/Charge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayManagerServer).Charge(ctx, req.(*PayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PayManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "paymentservice.PayManager",
	HandlerType: (*PayManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Charge",
			Handler:    _PayManager_Charge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pay.proto",
}