// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package xuperp2p

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

type XuperMessage_MessageType int32

const (
	XuperMessage_SENDBLOCK                XuperMessage_MessageType = 0
	XuperMessage_POSTTX                   XuperMessage_MessageType = 1
	XuperMessage_BATCHPOSTTX              XuperMessage_MessageType = 2
	XuperMessage_GET_BLOCK                XuperMessage_MessageType = 3
	XuperMessage_PING                     XuperMessage_MessageType = 4
	XuperMessage_GET_BLOCKCHAINSTATUS     XuperMessage_MessageType = 5
	XuperMessage_GET_BLOCK_RES            XuperMessage_MessageType = 6
	XuperMessage_GET_BLOCKCHAINSTATUS_RES XuperMessage_MessageType = 7
	// 向邻近确认区块是否为最新状态区块
	XuperMessage_CONFIRM_BLOCKCHAINSTATUS     XuperMessage_MessageType = 8
	XuperMessage_CONFIRM_BLOCKCHAINSTATUS_RES XuperMessage_MessageType = 9
	XuperMessage_MSG_TYPE_NONE                XuperMessage_MessageType = 10
	// query RPC port information
	XuperMessage_GET_RPC_PORT     XuperMessage_MessageType = 11
	XuperMessage_GET_RPC_PORT_RES XuperMessage_MessageType = 12
	// get authentication information
	XuperMessage_GET_AUTHENTICATION     XuperMessage_MessageType = 13
	XuperMessage_GET_AUTHENTICATION_RES XuperMessage_MessageType = 14
	// chained-bft NEW_VIEW message
	XuperMessage_CHAINED_BFT_NEW_VIEW_MSG XuperMessage_MessageType = 15
	// chained-bft NEW_PROPOSAL message
	XuperMessage_CHAINED_BFT_NEW_PROPOSAL_MSG XuperMessage_MessageType = 16
	// chained-bft vote message
	XuperMessage_CHAINED_BFT_VOTE_MSG XuperMessage_MessageType = 17
	// broadcast new block id to other node
	XuperMessage_NEW_BLOCKID XuperMessage_MessageType = 18
	// new node used to add to network automatic
	XuperMessage_NEW_NODE XuperMessage_MessageType = 19
	// 消息头同步对(GET_HEADERS <-> HEADERS),
	// 发送方通过GET_HEADERS消息询问区间范围内的所有区块哈希信息,
	// 接受方发送HEADERS信息, 该消息携带其所知的区间范围内的BlockId列表
	XuperMessage_GET_HEADERS XuperMessage_MessageType = 20
	XuperMessage_HEADERS     XuperMessage_MessageType = 21
	// 消息对(GET_BLOCKS <-> BLOCKS),
	// 发送方通过GET_BLOCKS消息询问BlockId列表内的所有对应区块信息,
	// 接受方发送BLOCKS信息, 该消息携带具体Block
	XuperMessage_GET_BLOCKS XuperMessage_MessageType = 22
	XuperMessage_BLOCKS     XuperMessage_MessageType = 23
)

var XuperMessage_MessageType_name = map[int32]string{
	0:  "SENDBLOCK",
	1:  "POSTTX",
	2:  "BATCHPOSTTX",
	3:  "GET_BLOCK",
	4:  "PING",
	5:  "GET_BLOCKCHAINSTATUS",
	6:  "GET_BLOCK_RES",
	7:  "GET_BLOCKCHAINSTATUS_RES",
	8:  "CONFIRM_BLOCKCHAINSTATUS",
	9:  "CONFIRM_BLOCKCHAINSTATUS_RES",
	10: "MSG_TYPE_NONE",
	11: "GET_RPC_PORT",
	12: "GET_RPC_PORT_RES",
	13: "GET_AUTHENTICATION",
	14: "GET_AUTHENTICATION_RES",
	15: "CHAINED_BFT_NEW_VIEW_MSG",
	16: "CHAINED_BFT_NEW_PROPOSAL_MSG",
	17: "CHAINED_BFT_VOTE_MSG",
	18: "NEW_BLOCKID",
	19: "NEW_NODE",
	20: "GET_HEADERS",
	21: "HEADERS",
	22: "GET_BLOCKS",
	23: "BLOCKS",
}

var XuperMessage_MessageType_value = map[string]int32{
	"SENDBLOCK":                    0,
	"POSTTX":                       1,
	"BATCHPOSTTX":                  2,
	"GET_BLOCK":                    3,
	"PING":                         4,
	"GET_BLOCKCHAINSTATUS":         5,
	"GET_BLOCK_RES":                6,
	"GET_BLOCKCHAINSTATUS_RES":     7,
	"CONFIRM_BLOCKCHAINSTATUS":     8,
	"CONFIRM_BLOCKCHAINSTATUS_RES": 9,
	"MSG_TYPE_NONE":                10,
	"GET_RPC_PORT":                 11,
	"GET_RPC_PORT_RES":             12,
	"GET_AUTHENTICATION":           13,
	"GET_AUTHENTICATION_RES":       14,
	"CHAINED_BFT_NEW_VIEW_MSG":     15,
	"CHAINED_BFT_NEW_PROPOSAL_MSG": 16,
	"CHAINED_BFT_VOTE_MSG":         17,
	"NEW_BLOCKID":                  18,
	"NEW_NODE":                     19,
	"GET_HEADERS":                  20,
	"HEADERS":                      21,
	"GET_BLOCKS":                   22,
	"BLOCKS":                       23,
}

func (x XuperMessage_MessageType) String() string {
	return proto.EnumName(XuperMessage_MessageType_name, int32(x))
}

func (XuperMessage_MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 0}
}

type XuperMessage_ErrorType int32

const (
	// success
	XuperMessage_SUCCESS XuperMessage_ErrorType = 0
	XuperMessage_NONE    XuperMessage_ErrorType = 1
	// common error
	XuperMessage_UNKNOW_ERROR             XuperMessage_ErrorType = 2
	XuperMessage_CHECK_SUM_ERROR          XuperMessage_ErrorType = 3
	XuperMessage_UNMARSHAL_MSG_BODY_ERROR XuperMessage_ErrorType = 4
	XuperMessage_CONNECT_REFUSE           XuperMessage_ErrorType = 5
	// block error
	XuperMessage_GET_BLOCKCHAIN_ERROR           XuperMessage_ErrorType = 6
	XuperMessage_BLOCKCHAIN_NOTEXIST            XuperMessage_ErrorType = 7
	XuperMessage_GET_BLOCK_ERROR                XuperMessage_ErrorType = 8
	XuperMessage_CONFIRM_BLOCKCHAINSTATUS_ERROR XuperMessage_ErrorType = 9
	XuperMessage_GET_AUTHENTICATION_ERROR       XuperMessage_ErrorType = 10
	XuperMessage_GET_AUTHENTICATION_NOT_PASS    XuperMessage_ErrorType = 11
)

var XuperMessage_ErrorType_name = map[int32]string{
	0:  "SUCCESS",
	1:  "NONE",
	2:  "UNKNOW_ERROR",
	3:  "CHECK_SUM_ERROR",
	4:  "UNMARSHAL_MSG_BODY_ERROR",
	5:  "CONNECT_REFUSE",
	6:  "GET_BLOCKCHAIN_ERROR",
	7:  "BLOCKCHAIN_NOTEXIST",
	8:  "GET_BLOCK_ERROR",
	9:  "CONFIRM_BLOCKCHAINSTATUS_ERROR",
	10: "GET_AUTHENTICATION_ERROR",
	11: "GET_AUTHENTICATION_NOT_PASS",
}

var XuperMessage_ErrorType_value = map[string]int32{
	"SUCCESS":                        0,
	"NONE":                           1,
	"UNKNOW_ERROR":                   2,
	"CHECK_SUM_ERROR":                3,
	"UNMARSHAL_MSG_BODY_ERROR":       4,
	"CONNECT_REFUSE":                 5,
	"GET_BLOCKCHAIN_ERROR":           6,
	"BLOCKCHAIN_NOTEXIST":            7,
	"GET_BLOCK_ERROR":                8,
	"CONFIRM_BLOCKCHAINSTATUS_ERROR": 9,
	"GET_AUTHENTICATION_ERROR":       10,
	"GET_AUTHENTICATION_NOT_PASS":    11,
}

func (x XuperMessage_ErrorType) String() string {
	return proto.EnumName(XuperMessage_ErrorType_name, int32(x))
}

func (XuperMessage_ErrorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 1}
}

// XuperMessage is the message of Xuper p2p server
type XuperMessage struct {
	Header               *XuperMessage_MessageHeader `protobuf:"bytes,1,opt,name=Header,proto3" json:"Header,omitempty"`
	Data                 *XuperMessage_MessageData   `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *XuperMessage) Reset()         { *m = XuperMessage{} }
func (m *XuperMessage) String() string { return proto.CompactTextString(m) }
func (*XuperMessage) ProtoMessage()    {}
func (*XuperMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *XuperMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XuperMessage.Unmarshal(m, b)
}
func (m *XuperMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XuperMessage.Marshal(b, m, deterministic)
}
func (m *XuperMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XuperMessage.Merge(m, src)
}
func (m *XuperMessage) XXX_Size() int {
	return xxx_messageInfo_XuperMessage.Size(m)
}
func (m *XuperMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_XuperMessage.DiscardUnknown(m)
}

var xxx_messageInfo_XuperMessage proto.InternalMessageInfo

func (m *XuperMessage) GetHeader() *XuperMessage_MessageHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *XuperMessage) GetData() *XuperMessage_MessageData {
	if m != nil {
		return m.Data
	}
	return nil
}

// MessageHeader is the message header of Xuper p2p server
type XuperMessage_MessageHeader struct {
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// dataCheckSum is the message data checksum, it can be used check where the message have been received
	Logid                string                   `protobuf:"bytes,2,opt,name=logid,proto3" json:"logid,omitempty"`
	From                 string                   `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	Bcname               string                   `protobuf:"bytes,4,opt,name=bcname,proto3" json:"bcname,omitempty"`
	Type                 XuperMessage_MessageType `protobuf:"varint,5,opt,name=type,proto3,enum=xuperp2p.XuperMessage_MessageType" json:"type,omitempty"`
	DataCheckSum         uint32                   `protobuf:"varint,6,opt,name=dataCheckSum,proto3" json:"dataCheckSum,omitempty"`
	ErrorType            XuperMessage_ErrorType   `protobuf:"varint,7,opt,name=errorType,proto3,enum=xuperp2p.XuperMessage_ErrorType" json:"errorType,omitempty"`
	EnableCompress       bool                     `protobuf:"varint,8,opt,name=enableCompress,proto3" json:"enableCompress,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *XuperMessage_MessageHeader) Reset()         { *m = XuperMessage_MessageHeader{} }
func (m *XuperMessage_MessageHeader) String() string { return proto.CompactTextString(m) }
func (*XuperMessage_MessageHeader) ProtoMessage()    {}
func (*XuperMessage_MessageHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 0}
}

func (m *XuperMessage_MessageHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XuperMessage_MessageHeader.Unmarshal(m, b)
}
func (m *XuperMessage_MessageHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XuperMessage_MessageHeader.Marshal(b, m, deterministic)
}
func (m *XuperMessage_MessageHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XuperMessage_MessageHeader.Merge(m, src)
}
func (m *XuperMessage_MessageHeader) XXX_Size() int {
	return xxx_messageInfo_XuperMessage_MessageHeader.Size(m)
}
func (m *XuperMessage_MessageHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_XuperMessage_MessageHeader.DiscardUnknown(m)
}

var xxx_messageInfo_XuperMessage_MessageHeader proto.InternalMessageInfo

func (m *XuperMessage_MessageHeader) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *XuperMessage_MessageHeader) GetLogid() string {
	if m != nil {
		return m.Logid
	}
	return ""
}

func (m *XuperMessage_MessageHeader) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *XuperMessage_MessageHeader) GetBcname() string {
	if m != nil {
		return m.Bcname
	}
	return ""
}

func (m *XuperMessage_MessageHeader) GetType() XuperMessage_MessageType {
	if m != nil {
		return m.Type
	}
	return XuperMessage_SENDBLOCK
}

func (m *XuperMessage_MessageHeader) GetDataCheckSum() uint32 {
	if m != nil {
		return m.DataCheckSum
	}
	return 0
}

func (m *XuperMessage_MessageHeader) GetErrorType() XuperMessage_ErrorType {
	if m != nil {
		return m.ErrorType
	}
	return XuperMessage_SUCCESS
}

func (m *XuperMessage_MessageHeader) GetEnableCompress() bool {
	if m != nil {
		return m.EnableCompress
	}
	return false
}

// MessageData is the message data of Xuper p2p server
type XuperMessage_MessageData struct {
	// msgInfo is the message infomation, use protobuf coding style
	MsgInfo              []byte   `protobuf:"bytes,3,opt,name=msgInfo,proto3" json:"msgInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XuperMessage_MessageData) Reset()         { *m = XuperMessage_MessageData{} }
func (m *XuperMessage_MessageData) String() string { return proto.CompactTextString(m) }
func (*XuperMessage_MessageData) ProtoMessage()    {}
func (*XuperMessage_MessageData) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 1}
}

func (m *XuperMessage_MessageData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XuperMessage_MessageData.Unmarshal(m, b)
}
func (m *XuperMessage_MessageData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XuperMessage_MessageData.Marshal(b, m, deterministic)
}
func (m *XuperMessage_MessageData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XuperMessage_MessageData.Merge(m, src)
}
func (m *XuperMessage_MessageData) XXX_Size() int {
	return xxx_messageInfo_XuperMessage_MessageData.Size(m)
}
func (m *XuperMessage_MessageData) XXX_DiscardUnknown() {
	xxx_messageInfo_XuperMessage_MessageData.DiscardUnknown(m)
}

var xxx_messageInfo_XuperMessage_MessageData proto.InternalMessageInfo

func (m *XuperMessage_MessageData) GetMsgInfo() []byte {
	if m != nil {
		return m.MsgInfo
	}
	return nil
}

func init() {
	proto.RegisterEnum("xuperp2p.XuperMessage_MessageType", XuperMessage_MessageType_name, XuperMessage_MessageType_value)
	proto.RegisterEnum("xuperp2p.XuperMessage_ErrorType", XuperMessage_ErrorType_name, XuperMessage_ErrorType_value)
	proto.RegisterType((*XuperMessage)(nil), "xuperp2p.XuperMessage")
	proto.RegisterType((*XuperMessage_MessageHeader)(nil), "xuperp2p.XuperMessage.MessageHeader")
	proto.RegisterType((*XuperMessage_MessageData)(nil), "xuperp2p.XuperMessage.MessageData")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 687 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0x4b, 0x6f, 0xe2, 0x48,
	0x10, 0xc7, 0xc3, 0x1b, 0x8a, 0x47, 0x3a, 0x15, 0x96, 0x58, 0xd9, 0x68, 0x17, 0xa1, 0xd5, 0x2e,
	0x27, 0x0e, 0x59, 0x69, 0x4f, 0xab, 0x95, 0x8c, 0xe9, 0x80, 0x95, 0xd0, 0x6d, 0x75, 0x37, 0x79,
	0x9c, 0x2c, 0x27, 0x71, 0xb2, 0xd1, 0x04, 0x8c, 0x0c, 0x19, 0x4d, 0xbe, 0xc2, 0x7c, 0x90, 0xb9,
	0xce, 0x47, 0x9c, 0x51, 0xb7, 0x0d, 0x21, 0xaf, 0x99, 0x13, 0xae, 0xff, 0xff, 0x57, 0x5d, 0xd5,
	0xd5, 0x25, 0xa0, 0x3e, 0x0d, 0x17, 0x8b, 0xe0, 0x36, 0xec, 0xcd, 0xe3, 0x68, 0x19, 0x61, 0xf9,
	0xd3, 0xc3, 0x3c, 0x8c, 0xe7, 0x87, 0xf3, 0xce, 0x67, 0x80, 0xda, 0xb9, 0x0e, 0xc6, 0x09, 0x80,
	0xff, 0x42, 0x71, 0x14, 0x06, 0xd7, 0x61, 0x6c, 0x65, 0xda, 0x99, 0x6e, 0xf5, 0xf0, 0x8f, 0xde,
	0x8a, 0xed, 0x6d, 0x72, 0xbd, 0xf4, 0x37, 0x61, 0x45, 0x9a, 0x83, 0xff, 0x40, 0x7e, 0x10, 0x2c,
	0x03, 0x2b, 0x6b, 0x72, 0x3b, 0x3f, 0xce, 0xd5, 0xa4, 0x30, 0xfc, 0xfe, 0xd7, 0x2c, 0xd4, 0x9f,
	0x9d, 0x88, 0x16, 0x94, 0x3e, 0x86, 0xf1, 0xe2, 0x2e, 0x9a, 0x99, 0x46, 0x2a, 0x62, 0x15, 0x62,
	0x13, 0x0a, 0xf7, 0xd1, 0xed, 0xdd, 0xb5, 0x29, 0x52, 0x11, 0x49, 0x80, 0x08, 0xf9, 0x9b, 0x38,
	0x9a, 0x5a, 0x39, 0x23, 0x9a, 0x6f, 0x6c, 0x41, 0xf1, 0xf2, 0x6a, 0x16, 0x4c, 0x43, 0x2b, 0x6f,
	0xd4, 0x34, 0xd2, 0x5d, 0x2e, 0x1f, 0xe7, 0xa1, 0x55, 0x68, 0x67, 0xba, 0x8d, 0x9f, 0x75, 0xa9,
	0x1e, 0xe7, 0xa1, 0x30, 0x3c, 0x76, 0xa0, 0x76, 0x1d, 0x2c, 0x03, 0xe7, 0xff, 0xf0, 0xea, 0x83,
	0x7c, 0x98, 0x5a, 0xc5, 0x76, 0xa6, 0x5b, 0x17, 0xcf, 0x34, 0xfc, 0x0f, 0x2a, 0x61, 0x1c, 0x47,
	0xb1, 0x4e, 0xb3, 0x4a, 0xa6, 0x40, 0xfb, 0x9d, 0x02, 0x74, 0xc5, 0x89, 0xa7, 0x14, 0xfc, 0x13,
	0x1a, 0xe1, 0x2c, 0xb8, 0xbc, 0x0f, 0x9d, 0x68, 0x3a, 0x8f, 0xc3, 0xc5, 0xc2, 0x2a, 0xb7, 0x33,
	0xdd, 0xb2, 0x78, 0xa1, 0xee, 0xff, 0x05, 0xd5, 0x8d, 0x31, 0xea, 0x71, 0x4d, 0x17, 0xb7, 0xee,
	0xec, 0x26, 0x32, 0x13, 0xa8, 0x89, 0x55, 0xd8, 0xf9, 0x96, 0x5b, 0x93, 0xa6, 0x40, 0x1d, 0x2a,
	0x92, 0xb2, 0x41, 0xff, 0x84, 0x3b, 0xc7, 0x64, 0x0b, 0x01, 0x8a, 0x1e, 0x97, 0x4a, 0x9d, 0x93,
	0x0c, 0x6e, 0x43, 0xb5, 0x6f, 0x2b, 0x67, 0x94, 0x0a, 0x59, 0xcd, 0x0e, 0xa9, 0xf2, 0x13, 0x36,
	0x87, 0x65, 0xc8, 0x7b, 0x2e, 0x1b, 0x92, 0x3c, 0x5a, 0xd0, 0x5c, 0x1b, 0xce, 0xc8, 0x76, 0x99,
	0x54, 0xb6, 0x9a, 0x48, 0x52, 0xc0, 0x1d, 0xa8, 0xaf, 0x1d, 0x5f, 0x50, 0x49, 0x8a, 0x78, 0x00,
	0xd6, 0x5b, 0xb0, 0x71, 0x4b, 0xda, 0x75, 0x38, 0x3b, 0x72, 0xc5, 0xf8, 0xf5, 0x71, 0x65, 0x6c,
	0xc3, 0xc1, 0x7b, 0xae, 0xc9, 0xaf, 0xe8, 0x82, 0x63, 0x39, 0xf4, 0xd5, 0x85, 0x47, 0x7d, 0xc6,
	0x19, 0x25, 0x80, 0x04, 0x6a, 0xba, 0xa0, 0xf0, 0x1c, 0xdf, 0xe3, 0x42, 0x91, 0x2a, 0x36, 0x81,
	0x6c, 0x2a, 0x26, 0xb5, 0x86, 0x2d, 0x40, 0xad, 0xda, 0x13, 0x35, 0xa2, 0x4c, 0xb9, 0x8e, 0xad,
	0x5c, 0xce, 0x48, 0x1d, 0xf7, 0xa1, 0xf5, 0x5a, 0x37, 0x39, 0x0d, 0xd3, 0xae, 0xee, 0x81, 0x0e,
	0xfc, 0xfe, 0x91, 0xf2, 0x19, 0x3d, 0xf3, 0x4f, 0x5d, 0x7a, 0xe6, 0x8f, 0xe5, 0x90, 0x6c, 0x9b,
	0x76, 0x5f, 0xb8, 0x9e, 0xe0, 0x1e, 0x97, 0xf6, 0x89, 0x21, 0x88, 0x9e, 0xdc, 0x26, 0x71, 0xca,
	0x15, 0x35, 0xce, 0x8e, 0x9e, 0xbe, 0xe6, 0xcd, 0x35, 0xdd, 0x01, 0x41, 0xac, 0x41, 0x59, 0x0b,
	0x8c, 0x0f, 0x28, 0xd9, 0xd5, 0xb6, 0x6e, 0x6a, 0x44, 0xed, 0x01, 0x15, 0x92, 0x34, 0xb1, 0x0a,
	0xa5, 0x55, 0xf0, 0x0b, 0x36, 0x00, 0xd6, 0x33, 0x96, 0xa4, 0xa5, 0x9f, 0x35, 0xfd, 0xde, 0xeb,
	0x7c, 0xc9, 0x42, 0x65, 0xbd, 0x6b, 0x3a, 0x4d, 0x4e, 0x1c, 0x87, 0x4a, 0x49, 0xb6, 0xf4, 0x8b,
	0x9a, 0x99, 0x65, 0xf4, 0xcc, 0x26, 0xec, 0x98, 0xf1, 0x33, 0x9f, 0x0a, 0xc1, 0x05, 0xc9, 0xe2,
	0x2e, 0x6c, 0x3b, 0x23, 0xea, 0x1c, 0xfb, 0x72, 0x32, 0x4e, 0xc5, 0x9c, 0xbe, 0xfe, 0x84, 0x8d,
	0x6d, 0x21, 0x47, 0xc9, 0x8d, 0xfc, 0x3e, 0x1f, 0x5c, 0xa4, 0x6e, 0x1e, 0x11, 0x1a, 0x0e, 0x67,
	0x8c, 0x3a, 0x7a, 0xc2, 0x47, 0x13, 0x49, 0x49, 0xe1, 0xf5, 0xaa, 0xa4, 0x74, 0x11, 0xf7, 0x60,
	0x77, 0x43, 0x65, 0x5c, 0xd1, 0x73, 0x57, 0x2a, 0x52, 0xd2, 0x95, 0x9f, 0x76, 0x28, 0xa1, 0xcb,
	0xd8, 0x81, 0xdf, 0xde, 0xdd, 0x84, 0x84, 0xa9, 0xac, 0x36, 0xed, 0xc5, 0xc3, 0x25, 0x2e, 0xe0,
	0xef, 0xf0, 0xeb, 0x1b, 0x2e, 0xe3, 0xca, 0xf7, 0x6c, 0x29, 0x49, 0xf5, 0xb2, 0x68, 0xfe, 0x1d,
	0xff, 0xfe, 0x1e, 0x00, 0x00, 0xff, 0xff, 0x38, 0xdf, 0xfc, 0x48, 0x2e, 0x05, 0x00, 0x00,
}
