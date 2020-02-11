// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contract_service.proto

package pbrpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	pb "github.com/xuperchain/xuperchain/core/contractsdk/go/pb"
	grpc "google.golang.org/grpc"
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

func init() { proto.RegisterFile("contract_service.proto", fileDescriptor_e663a77702825514) }

var fileDescriptor_e663a77702825514 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x85, 0x98, 0xd8, 0xfa, 0x36, 0x71, 0xf0, 0x24, 0x0e, 0x93, 0x10, 0x03, 0xc6, 0x06,
	0x97, 0x04, 0x8d, 0x1b, 0x07, 0xa4, 0xae, 0x48, 0x01, 0x21, 0x65, 0x83, 0x16, 0x55, 0xaa, 0x84,
	0x50, 0xe2, 0x3c, 0xd2, 0xd0, 0x10, 0x07, 0xfb, 0xb9, 0xa4, 0x5f, 0x8b, 0xaf, 0xc1, 0x97, 0x42,
	0x4d, 0xed, 0xb6, 0x02, 0x27, 0xe9, 0x65, 0xb7, 0x24, 0xef, 0xf7, 0xff, 0xf9, 0xe5, 0xc5, 0x0e,
	0x3c, 0xe0, 0xa2, 0x20, 0x19, 0x71, 0xfa, 0xaa, 0x50, 0xce, 0x33, 0x8e, 0x5e, 0x29, 0x05, 0x09,
	0x76, 0x5c, 0xf1, 0x69, 0x94, 0x15, 0x9e, 0x2d, 0x7b, 0x6a, 0xce, 0x4f, 0xee, 0xaf, 0xef, 0x6a,
	0xe8, 0xf2, 0xf7, 0x1d, 0x80, 0x30, 0xa2, 0x6c, 0x8e, 0x03, 0x91, 0x20, 0x1b, 0xc3, 0xde, 0x20,
	0xca, 0x73, 0x76, 0xee, 0xfd, 0x17, 0x4e, 0x66, 0x9e, 0x01, 0xa3, 0x3c, 0xff, 0x84, 0x3f, 0x35,
	0x2a, 0x3a, 0xb9, 0xe8, 0xe4, 0x54, 0x29, 0x0a, 0x85, 0xec, 0x03, 0xec, 0xdd, 0x64, 0x45, 0xca,
	0x4e, 0x9d, 0x81, 0x65, 0xc9, 0x2a, 0x1f, 0xb7, 0x10, 0x2b, 0xd9, 0xe5, 0x9f, 0x03, 0xd8, 0x1f,
	0x2e, 0x14, 0x5f, 0x76, 0x1a, 0x42, 0xef, 0x46, 0xd3, 0x75, 0xfc, 0x1d, 0x39, 0xb1, 0x47, 0xee,
	0xac, 0x26, 0x2b, 0x3f, 0x6d, 0x06, 0x4c, 0xa3, 0x21, 0xf4, 0x02, 0x6c, 0xf7, 0x05, 0xd8, 0xe1,
	0xab, 0x01, 0xe3, 0x1b, 0xc3, 0xd1, 0x5b, 0xcc, 0x91, 0xd0, 0x28, 0x9f, 0x38, 0x13, 0x2b, 0xc4,
	0x5a, 0x9f, 0xb6, 0x32, 0x46, 0x3c, 0x81, 0xc3, 0x10, 0x7f, 0xbd, 0x27, 0x94, 0x11, 0x09, 0xc9,
	0xce, 0x9c, 0x19, 0x5b, 0xb6, 0xe6, 0x67, 0x1d, 0x94, 0x71, 0x8f, 0x60, 0xff, 0xa3, 0x46, 0xb9,
	0x18, 0x55, 0xcc, 0xdd, 0x8b, 0xa9, 0x5a, 0xed, 0x59, 0x3b, 0x64, 0xac, 0x5f, 0x00, 0xea, 0x47,
	0x57, 0xb9, 0xe0, 0xb3, 0x86, 0x2d, 0xb6, 0x01, 0xda, 0xb7, 0xd8, 0x36, 0xb7, 0x9e, 0xf4, 0xc1,
	0x48, 0x46, 0x85, 0xfa, 0x86, 0x4d, 0xd3, 0xb0, 0xe5, 0xf6, 0x69, 0x6c, 0x28, 0x23, 0xe6, 0x70,
	0x34, 0x30, 0x40, 0x7d, 0x38, 0x9e, 0x3b, 0x63, 0xdb, 0x88, 0x5d, 0xe0, 0xc5, 0x0e, 0xa4, 0x59,
	0xa4, 0x82, 0xe3, 0x00, 0xa9, 0xcf, 0xb9, 0xd0, 0x05, 0xf5, 0x93, 0x44, 0xa2, 0x52, 0xa8, 0x98,
	0xdf, 0xb4, 0xc1, 0xfe, 0x25, 0xed, 0x92, 0x2f, 0x77, 0x0f, 0xdc, 0xc2, 0xd1, 0x64, 0x9f, 0xe1,
	0x30, 0xc0, 0xfa, 0xcd, 0xfa, 0x32, 0x55, 0xec, 0xa2, 0xa9, 0x1b, 0x4b, 0x58, 0xf5, 0x43, 0xf7,
	0xa4, 0xac, 0x67, 0x02, 0xbd, 0x21, 0xd2, 0xb5, 0xa6, 0x52, 0x13, 0x73, 0x7f, 0xb6, 0x75, 0xdd,
	0x2a, 0xcf, 0xbb, 0xb0, 0x55, 0xcb, 0x57, 0x6f, 0xde, 0xdd, 0x9d, 0xbc, 0x4e, 0x33, 0x9a, 0xea,
	0xd8, 0xe3, 0xe2, 0x87, 0x5f, 0xe9, 0x12, 0x65, 0x9d, 0xdc, 0xbe, 0xe4, 0x42, 0xa2, 0x6f, 0x4d,
	0x2a, 0x99, 0xf9, 0xa9, 0xf0, 0xcb, 0x58, 0x96, 0x3c, 0xbe, 0x57, 0xff, 0x49, 0x5f, 0xfd, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0xf9, 0x0a, 0x58, 0x9d, 0x88, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NativeCodeClient is the client API for NativeCode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NativeCodeClient interface {
	Call(ctx context.Context, in *pb.NativeCallRequest, opts ...grpc.CallOption) (*pb.NativeCallResponse, error)
	Ping(ctx context.Context, in *pb.PingRequest, opts ...grpc.CallOption) (*pb.PingResponse, error)
}

type nativeCodeClient struct {
	cc *grpc.ClientConn
}

func NewNativeCodeClient(cc *grpc.ClientConn) NativeCodeClient {
	return &nativeCodeClient{cc}
}

func (c *nativeCodeClient) Call(ctx context.Context, in *pb.NativeCallRequest, opts ...grpc.CallOption) (*pb.NativeCallResponse, error) {
	out := new(pb.NativeCallResponse)
	err := c.cc.Invoke(ctx, "/xchain.contract.svc.NativeCode/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nativeCodeClient) Ping(ctx context.Context, in *pb.PingRequest, opts ...grpc.CallOption) (*pb.PingResponse, error) {
	out := new(pb.PingResponse)
	err := c.cc.Invoke(ctx, "/xchain.contract.svc.NativeCode/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NativeCodeServer is the server API for NativeCode service.
type NativeCodeServer interface {
	Call(context.Context, *pb.NativeCallRequest) (*pb.NativeCallResponse, error)
	Ping(context.Context, *pb.PingRequest) (*pb.PingResponse, error)
}

func RegisterNativeCodeServer(s *grpc.Server, srv NativeCodeServer) {
	s.RegisterService(&_NativeCode_serviceDesc, srv)
}

func _NativeCode_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.NativeCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NativeCodeServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xchain.contract.svc.NativeCode/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NativeCodeServer).Call(ctx, req.(*pb.NativeCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NativeCode_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NativeCodeServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xchain.contract.svc.NativeCode/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NativeCodeServer).Ping(ctx, req.(*pb.PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NativeCode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "xchain.contract.svc.NativeCode",
	HandlerType: (*NativeCodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _NativeCode_Call_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _NativeCode_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contract_service.proto",
}

// SyscallClient is the client API for Syscall service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SyscallClient interface {
	// KV service
	PutObject(ctx context.Context, in *pb.PutRequest, opts ...grpc.CallOption) (*pb.PutResponse, error)
	GetObject(ctx context.Context, in *pb.GetRequest, opts ...grpc.CallOption) (*pb.GetResponse, error)
	DeleteObject(ctx context.Context, in *pb.DeleteRequest, opts ...grpc.CallOption) (*pb.DeleteResponse, error)
	NewIterator(ctx context.Context, in *pb.IteratorRequest, opts ...grpc.CallOption) (*pb.IteratorResponse, error)
	// Chain service
	QueryTx(ctx context.Context, in *pb.QueryTxRequest, opts ...grpc.CallOption) (*pb.QueryTxResponse, error)
	QueryBlock(ctx context.Context, in *pb.QueryBlockRequest, opts ...grpc.CallOption) (*pb.QueryBlockResponse, error)
	Transfer(ctx context.Context, in *pb.TransferRequest, opts ...grpc.CallOption) (*pb.TransferResponse, error)
	ContractCall(ctx context.Context, in *pb.ContractCallRequest, opts ...grpc.CallOption) (*pb.ContractCallResponse, error)
	GetAccountAddresses(ctx context.Context, in *pb.GetAccountAddressesRequest, opts ...grpc.CallOption) (*pb.GetAccountAddressesResponse, error)
	// Heartbeat
	Ping(ctx context.Context, in *pb.PingRequest, opts ...grpc.CallOption) (*pb.PingResponse, error)
	GetCallArgs(ctx context.Context, in *pb.GetCallArgsRequest, opts ...grpc.CallOption) (*pb.CallArgs, error)
	SetOutput(ctx context.Context, in *pb.SetOutputRequest, opts ...grpc.CallOption) (*pb.SetOutputResponse, error)
}

type syscallClient struct {
	cc *grpc.ClientConn
}

func NewSyscallClient(cc *grpc.ClientConn) SyscallClient {
	return &syscallClient{cc}
}

func (c *syscallClient) PutObject(ctx context.Context, in *pb.PutRequest, opts ...grpc.CallOption) (*pb.PutResponse, error) {
	out := new(pb.PutResponse)
	err := c.cc.Invoke(ctx, "/xchain.contract.svc.Syscall/PutObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syscallClient) GetObject(ctx context.Context, in *pb.GetRequest, opts ...grpc.CallOption) (*pb.GetResponse, error) {
	out := new(pb.GetResponse)
	err := c.cc.Invoke(ctx, "/xchain.contract.svc.Syscall/GetObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syscallClient) DeleteObject(ctx context.Context, in *pb.DeleteRequest, opts ...grpc.CallOption) (*pb.DeleteResponse, error) {
	out := new(pb.DeleteResponse)
	err := c.cc.Invoke(ctx, "/xchain.contract.svc.Syscall/DeleteObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syscallClient) NewIterator(ctx context.Context, in *pb.IteratorRequest, opts ...grpc.CallOption) (*pb.IteratorResponse, error) {
	out := new(pb.IteratorResponse)
	err := c.cc.Invoke(ctx, "/xchain.contract.svc.Syscall/NewIterator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syscallClient) QueryTx(ctx context.Context, in *pb.QueryTxRequest, opts ...grpc.CallOption) (*pb.QueryTxResponse, error) {
	out := new(pb.QueryTxResponse)
	err := c.cc.Invoke(ctx, "/xchain.contract.svc.Syscall/QueryTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syscallClient) QueryBlock(ctx context.Context, in *pb.QueryBlockRequest, opts ...grpc.CallOption) (*pb.QueryBlockResponse, error) {
	out := new(pb.QueryBlockResponse)
	err := c.cc.Invoke(ctx, "/xchain.contract.svc.Syscall/QueryBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syscallClient) Transfer(ctx context.Context, in *pb.TransferRequest, opts ...grpc.CallOption) (*pb.TransferResponse, error) {
	out := new(pb.TransferResponse)
	err := c.cc.Invoke(ctx, "/xchain.contract.svc.Syscall/Transfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syscallClient) ContractCall(ctx context.Context, in *pb.ContractCallRequest, opts ...grpc.CallOption) (*pb.ContractCallResponse, error) {
	out := new(pb.ContractCallResponse)
	err := c.cc.Invoke(ctx, "/xchain.contract.svc.Syscall/ContractCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syscallClient) GetAccountAddresses(ctx context.Context, in *pb.GetAccountAddressesRequest, opts ...grpc.CallOption) (*pb.GetAccountAddressesResponse, error) {
	out := new(pb.GetAccountAddressesResponse)
	err := c.cc.Invoke(ctx, "/xchain.contract.svc.Syscall/GetAccountAddresses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syscallClient) Ping(ctx context.Context, in *pb.PingRequest, opts ...grpc.CallOption) (*pb.PingResponse, error) {
	out := new(pb.PingResponse)
	err := c.cc.Invoke(ctx, "/xchain.contract.svc.Syscall/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syscallClient) GetCallArgs(ctx context.Context, in *pb.GetCallArgsRequest, opts ...grpc.CallOption) (*pb.CallArgs, error) {
	out := new(pb.CallArgs)
	err := c.cc.Invoke(ctx, "/xchain.contract.svc.Syscall/GetCallArgs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syscallClient) SetOutput(ctx context.Context, in *pb.SetOutputRequest, opts ...grpc.CallOption) (*pb.SetOutputResponse, error) {
	out := new(pb.SetOutputResponse)
	err := c.cc.Invoke(ctx, "/xchain.contract.svc.Syscall/SetOutput", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SyscallServer is the server API for Syscall service.
type SyscallServer interface {
	// KV service
	PutObject(context.Context, *pb.PutRequest) (*pb.PutResponse, error)
	GetObject(context.Context, *pb.GetRequest) (*pb.GetResponse, error)
	DeleteObject(context.Context, *pb.DeleteRequest) (*pb.DeleteResponse, error)
	NewIterator(context.Context, *pb.IteratorRequest) (*pb.IteratorResponse, error)
	// Chain service
	QueryTx(context.Context, *pb.QueryTxRequest) (*pb.QueryTxResponse, error)
	QueryBlock(context.Context, *pb.QueryBlockRequest) (*pb.QueryBlockResponse, error)
	Transfer(context.Context, *pb.TransferRequest) (*pb.TransferResponse, error)
	ContractCall(context.Context, *pb.ContractCallRequest) (*pb.ContractCallResponse, error)
	GetAccountAddresses(context.Context, *pb.GetAccountAddressesRequest) (*pb.GetAccountAddressesResponse, error)
	// Heartbeat
	Ping(context.Context, *pb.PingRequest) (*pb.PingResponse, error)
	GetCallArgs(context.Context, *pb.GetCallArgsRequest) (*pb.CallArgs, error)
	SetOutput(context.Context, *pb.SetOutputRequest) (*pb.SetOutputResponse, error)
}

func RegisterSyscallServer(s *grpc.Server, srv SyscallServer) {
	s.RegisterService(&_Syscall_serviceDesc, srv)
}

func _Syscall_PutObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyscallServer).PutObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xchain.contract.svc.Syscall/PutObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyscallServer).PutObject(ctx, req.(*pb.PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Syscall_GetObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyscallServer).GetObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xchain.contract.svc.Syscall/GetObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyscallServer).GetObject(ctx, req.(*pb.GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Syscall_DeleteObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyscallServer).DeleteObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xchain.contract.svc.Syscall/DeleteObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyscallServer).DeleteObject(ctx, req.(*pb.DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Syscall_NewIterator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.IteratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyscallServer).NewIterator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xchain.contract.svc.Syscall/NewIterator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyscallServer).NewIterator(ctx, req.(*pb.IteratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Syscall_QueryTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.QueryTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyscallServer).QueryTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xchain.contract.svc.Syscall/QueryTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyscallServer).QueryTx(ctx, req.(*pb.QueryTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Syscall_QueryBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.QueryBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyscallServer).QueryBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xchain.contract.svc.Syscall/QueryBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyscallServer).QueryBlock(ctx, req.(*pb.QueryBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Syscall_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.TransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyscallServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xchain.contract.svc.Syscall/Transfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyscallServer).Transfer(ctx, req.(*pb.TransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Syscall_ContractCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.ContractCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyscallServer).ContractCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xchain.contract.svc.Syscall/ContractCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyscallServer).ContractCall(ctx, req.(*pb.ContractCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Syscall_GetAccountAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.GetAccountAddressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyscallServer).GetAccountAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xchain.contract.svc.Syscall/GetAccountAddresses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyscallServer).GetAccountAddresses(ctx, req.(*pb.GetAccountAddressesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Syscall_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyscallServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xchain.contract.svc.Syscall/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyscallServer).Ping(ctx, req.(*pb.PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Syscall_GetCallArgs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.GetCallArgsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyscallServer).GetCallArgs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xchain.contract.svc.Syscall/GetCallArgs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyscallServer).GetCallArgs(ctx, req.(*pb.GetCallArgsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Syscall_SetOutput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.SetOutputRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyscallServer).SetOutput(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xchain.contract.svc.Syscall/SetOutput",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyscallServer).SetOutput(ctx, req.(*pb.SetOutputRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Syscall_serviceDesc = grpc.ServiceDesc{
	ServiceName: "xchain.contract.svc.Syscall",
	HandlerType: (*SyscallServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutObject",
			Handler:    _Syscall_PutObject_Handler,
		},
		{
			MethodName: "GetObject",
			Handler:    _Syscall_GetObject_Handler,
		},
		{
			MethodName: "DeleteObject",
			Handler:    _Syscall_DeleteObject_Handler,
		},
		{
			MethodName: "NewIterator",
			Handler:    _Syscall_NewIterator_Handler,
		},
		{
			MethodName: "QueryTx",
			Handler:    _Syscall_QueryTx_Handler,
		},
		{
			MethodName: "QueryBlock",
			Handler:    _Syscall_QueryBlock_Handler,
		},
		{
			MethodName: "Transfer",
			Handler:    _Syscall_Transfer_Handler,
		},
		{
			MethodName: "ContractCall",
			Handler:    _Syscall_ContractCall_Handler,
		},
		{
			MethodName: "GetAccountAddresses",
			Handler:    _Syscall_GetAccountAddresses_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Syscall_Ping_Handler,
		},
		{
			MethodName: "GetCallArgs",
			Handler:    _Syscall_GetCallArgs_Handler,
		},
		{
			MethodName: "SetOutput",
			Handler:    _Syscall_SetOutput_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contract_service.proto",
}
