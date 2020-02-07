// Code generated by protoc-gen-go. DO NOT EDIT.
// source: xcheck.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ComplianceCheckResponse struct {
	Header               *Header        `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Signature            *SignatureInfo `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ComplianceCheckResponse) Reset()         { *m = ComplianceCheckResponse{} }
func (m *ComplianceCheckResponse) String() string { return proto.CompactTextString(m) }
func (*ComplianceCheckResponse) ProtoMessage()    {}
func (*ComplianceCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_xcheck_00c6de706df97d9a, []int{0}
}
func (m *ComplianceCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComplianceCheckResponse.Unmarshal(m, b)
}
func (m *ComplianceCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComplianceCheckResponse.Marshal(b, m, deterministic)
}
func (dst *ComplianceCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComplianceCheckResponse.Merge(dst, src)
}
func (m *ComplianceCheckResponse) XXX_Size() int {
	return xxx_messageInfo_ComplianceCheckResponse.Size(m)
}
func (m *ComplianceCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ComplianceCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ComplianceCheckResponse proto.InternalMessageInfo

func (m *ComplianceCheckResponse) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ComplianceCheckResponse) GetSignature() *SignatureInfo {
	if m != nil {
		return m.Signature
	}
	return nil
}

type TransferCheckResponse struct {
	Header               *Header        `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Signature            *SignatureInfo `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TransferCheckResponse) Reset()         { *m = TransferCheckResponse{} }
func (m *TransferCheckResponse) String() string { return proto.CompactTextString(m) }
func (*TransferCheckResponse) ProtoMessage()    {}
func (*TransferCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_xcheck_00c6de706df97d9a, []int{1}
}
func (m *TransferCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferCheckResponse.Unmarshal(m, b)
}
func (m *TransferCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferCheckResponse.Marshal(b, m, deterministic)
}
func (dst *TransferCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferCheckResponse.Merge(dst, src)
}
func (m *TransferCheckResponse) XXX_Size() int {
	return xxx_messageInfo_TransferCheckResponse.Size(m)
}
func (m *TransferCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransferCheckResponse proto.InternalMessageInfo

func (m *TransferCheckResponse) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *TransferCheckResponse) GetSignature() *SignatureInfo {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*ComplianceCheckResponse)(nil), "pb.ComplianceCheckResponse")
	proto.RegisterType((*TransferCheckResponse)(nil), "pb.TransferCheckResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Xcheck service

type XcheckClient interface {
	ComplianceCheck(ctx context.Context, in *TxStatus, opts ...grpc.CallOption) (*ComplianceCheckResponse, error)
	TransferCheck(ctx context.Context, in *TxStatus, opts ...grpc.CallOption) (*TransferCheckResponse, error)
}

type xcheckClient struct {
	cc *grpc.ClientConn
}

func NewXcheckClient(cc *grpc.ClientConn) XcheckClient {
	return &xcheckClient{cc}
}

func (c *xcheckClient) ComplianceCheck(ctx context.Context, in *TxStatus, opts ...grpc.CallOption) (*ComplianceCheckResponse, error) {
	out := new(ComplianceCheckResponse)
	err := grpc.Invoke(ctx, "/pb.Xcheck/ComplianceCheck", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xcheckClient) TransferCheck(ctx context.Context, in *TxStatus, opts ...grpc.CallOption) (*TransferCheckResponse, error) {
	out := new(TransferCheckResponse)
	err := grpc.Invoke(ctx, "/pb.Xcheck/TransferCheck", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Xcheck service

type XcheckServer interface {
	ComplianceCheck(context.Context, *TxStatus) (*ComplianceCheckResponse, error)
	TransferCheck(context.Context, *TxStatus) (*TransferCheckResponse, error)
}

func RegisterXcheckServer(s *grpc.Server, srv XcheckServer) {
	s.RegisterService(&_Xcheck_serviceDesc, srv)
}

func _Xcheck_ComplianceCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XcheckServer).ComplianceCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Xcheck/ComplianceCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XcheckServer).ComplianceCheck(ctx, req.(*TxStatus))
	}
	return interceptor(ctx, in, info, handler)
}

func _Xcheck_TransferCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XcheckServer).TransferCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Xcheck/TransferCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XcheckServer).TransferCheck(ctx, req.(*TxStatus))
	}
	return interceptor(ctx, in, info, handler)
}

var _Xcheck_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Xcheck",
	HandlerType: (*XcheckServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ComplianceCheck",
			Handler:    _Xcheck_ComplianceCheck_Handler,
		},
		{
			MethodName: "TransferCheck",
			Handler:    _Xcheck_TransferCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "xcheck.proto",
}

func init() { proto.RegisterFile("xcheck.proto", fileDescriptor_xcheck_00c6de706df97d9a) }

var fileDescriptor_xcheck_00c6de706df97d9a = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xa9, 0x48, 0xce, 0x48,
	0x4d, 0xce, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x92, 0x92, 0x49, 0xcf,
	0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c,
	0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0xa8, 0x90, 0x02, 0xa9, 0x4f, 0xcc, 0xcc, 0x83, 0xf0, 0x94, 0xf2,
	0xb8, 0xc4, 0x9d, 0xf3, 0x73, 0x0b, 0x72, 0x32, 0x13, 0xf3, 0x92, 0x53, 0x9d, 0x41, 0x06, 0x05,
	0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x29, 0x71, 0xb1, 0x65, 0xa4, 0x26, 0xa6, 0xa4,
	0x16, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x71, 0xe9, 0x15, 0x24, 0xe9, 0x79, 0x80, 0x45,
	0x82, 0xa0, 0x32, 0x42, 0xfa, 0x5c, 0x9c, 0xc5, 0x99, 0xe9, 0x79, 0x89, 0x25, 0xa5, 0x45, 0xa9,
	0x12, 0x4c, 0x60, 0x65, 0x82, 0x20, 0x65, 0xc1, 0x30, 0x41, 0xcf, 0xbc, 0xb4, 0xfc, 0x20, 0x84,
	0x1a, 0xa5, 0x1c, 0x2e, 0xd1, 0x90, 0xa2, 0xc4, 0xbc, 0xe2, 0xb4, 0xd4, 0x22, 0xda, 0xdb, 0x66,
	0xd4, 0xc0, 0xc8, 0xc5, 0x16, 0x01, 0x0e, 0x1e, 0x21, 0x1b, 0x2e, 0x7e, 0x34, 0x8f, 0x0a, 0xf1,
	0x80, 0xf4, 0x86, 0x54, 0x04, 0x97, 0x24, 0x96, 0x94, 0x16, 0x4b, 0x49, 0x83, 0x78, 0xb8, 0xc2,
	0xc2, 0x82, 0x8b, 0x17, 0xc5, 0xd9, 0x68, 0x7a, 0x25, 0xc1, 0x3c, 0x6c, 0xfe, 0x4a, 0x62, 0x03,
	0x87, 0xb3, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x8a, 0x44, 0xb4, 0xa7, 0x01, 0x00, 0x00,
}
