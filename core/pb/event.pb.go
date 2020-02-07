// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// BlockStatus block status
type BlockStatus int32

const (
	BlockStatus_ERROR    BlockStatus = 0
	BlockStatus_TRUNK    BlockStatus = 1
	BlockStatus_BRANCH   BlockStatus = 2
	BlockStatus_NONEXIST BlockStatus = 3
)

var BlockStatus_name = map[int32]string{
	0: "ERROR",
	1: "TRUNK",
	2: "BRANCH",
	3: "NONEXIST",
}
var BlockStatus_value = map[string]int32{
	"ERROR":    0,
	"TRUNK":    1,
	"BRANCH":   2,
	"NONEXIST": 3,
}

func (x BlockStatus) String() string {
	return proto.EnumName(BlockStatus_name, int32(x))
}
func (BlockStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_event_28ecaf4b97a0a4ec, []int{0}
}

// /// EventType 事件类型
type EventType int32

const (
	EventType_UNDEFINED          EventType = 0
	EventType_BLOCK              EventType = 1
	EventType_TRANSACTION        EventType = 2
	EventType_ACCOUNT            EventType = 3
	EventType_SUBSCRIBE_RESPONSE EventType = 4
)

var EventType_name = map[int32]string{
	0: "UNDEFINED",
	1: "BLOCK",
	2: "TRANSACTION",
	3: "ACCOUNT",
	4: "SUBSCRIBE_RESPONSE",
}
var EventType_value = map[string]int32{
	"UNDEFINED":          0,
	"BLOCK":              1,
	"TRANSACTION":        2,
	"ACCOUNT":            3,
	"SUBSCRIBE_RESPONSE": 4,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}
func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_event_28ecaf4b97a0a4ec, []int{1}
}

type UnsubscribeStatusInfo int32

const (
	UnsubscribeStatusInfo_UNSUBSCRIBE_UNDEFINED UnsubscribeStatusInfo = 0
	UnsubscribeStatusInfo_UNSUBSCRIBE_SUCCESS   UnsubscribeStatusInfo = 1
	UnsubscribeStatusInfo_UNSUBSCRIBE_FAILED    UnsubscribeStatusInfo = 2
)

var UnsubscribeStatusInfo_name = map[int32]string{
	0: "UNSUBSCRIBE_UNDEFINED",
	1: "UNSUBSCRIBE_SUCCESS",
	2: "UNSUBSCRIBE_FAILED",
}
var UnsubscribeStatusInfo_value = map[string]int32{
	"UNSUBSCRIBE_UNDEFINED": 0,
	"UNSUBSCRIBE_SUCCESS":   1,
	"UNSUBSCRIBE_FAILED":    2,
}

func (x UnsubscribeStatusInfo) String() string {
	return proto.EnumName(UnsubscribeStatusInfo_name, int32(x))
}
func (UnsubscribeStatusInfo) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_event_28ecaf4b97a0a4ec, []int{2}
}

// BlockStatusInfo 区块元数据
type BlockStatusInfo struct {
	Bcname               string      `protobuf:"bytes,1,opt,name=bcname" json:"bcname,omitempty"`
	Status               BlockStatus `protobuf:"varint,2,opt,name=status,enum=pb.BlockStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *BlockStatusInfo) Reset()         { *m = BlockStatusInfo{} }
func (m *BlockStatusInfo) String() string { return proto.CompactTextString(m) }
func (*BlockStatusInfo) ProtoMessage()    {}
func (*BlockStatusInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_28ecaf4b97a0a4ec, []int{0}
}
func (m *BlockStatusInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockStatusInfo.Unmarshal(m, b)
}
func (m *BlockStatusInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockStatusInfo.Marshal(b, m, deterministic)
}
func (dst *BlockStatusInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockStatusInfo.Merge(dst, src)
}
func (m *BlockStatusInfo) XXX_Size() int {
	return xxx_messageInfo_BlockStatusInfo.Size(m)
}
func (m *BlockStatusInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockStatusInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BlockStatusInfo proto.InternalMessageInfo

func (m *BlockStatusInfo) GetBcname() string {
	if m != nil {
		return m.Bcname
	}
	return ""
}

func (m *BlockStatusInfo) GetStatus() BlockStatus {
	if m != nil {
		return m.Status
	}
	return BlockStatus_ERROR
}

type TransactionStatusInfo struct {
	Bcname               string            `protobuf:"bytes,1,opt,name=bcname" json:"bcname,omitempty"`
	Initiator            string            `protobuf:"bytes,2,opt,name=initiator" json:"initiator,omitempty"`
	Status               TransactionStatus `protobuf:"varint,3,opt,name=status,enum=pb.TransactionStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TransactionStatusInfo) Reset()         { *m = TransactionStatusInfo{} }
func (m *TransactionStatusInfo) String() string { return proto.CompactTextString(m) }
func (*TransactionStatusInfo) ProtoMessage()    {}
func (*TransactionStatusInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_28ecaf4b97a0a4ec, []int{1}
}
func (m *TransactionStatusInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionStatusInfo.Unmarshal(m, b)
}
func (m *TransactionStatusInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionStatusInfo.Marshal(b, m, deterministic)
}
func (dst *TransactionStatusInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionStatusInfo.Merge(dst, src)
}
func (m *TransactionStatusInfo) XXX_Size() int {
	return xxx_messageInfo_TransactionStatusInfo.Size(m)
}
func (m *TransactionStatusInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionStatusInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionStatusInfo proto.InternalMessageInfo

func (m *TransactionStatusInfo) GetBcname() string {
	if m != nil {
		return m.Bcname
	}
	return ""
}

func (m *TransactionStatusInfo) GetInitiator() string {
	if m != nil {
		return m.Initiator
	}
	return ""
}

func (m *TransactionStatusInfo) GetStatus() TransactionStatus {
	if m != nil {
		return m.Status
	}
	return TransactionStatus_UNDEFINE
}

type AccountStatusInfo struct {
	Bcname               string            `protobuf:"bytes,1,opt,name=bcname" json:"bcname,omitempty"`
	Status               TransactionStatus `protobuf:"varint,2,opt,name=status,enum=pb.TransactionStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AccountStatusInfo) Reset()         { *m = AccountStatusInfo{} }
func (m *AccountStatusInfo) String() string { return proto.CompactTextString(m) }
func (*AccountStatusInfo) ProtoMessage()    {}
func (*AccountStatusInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_28ecaf4b97a0a4ec, []int{2}
}
func (m *AccountStatusInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountStatusInfo.Unmarshal(m, b)
}
func (m *AccountStatusInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountStatusInfo.Marshal(b, m, deterministic)
}
func (dst *AccountStatusInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountStatusInfo.Merge(dst, src)
}
func (m *AccountStatusInfo) XXX_Size() int {
	return xxx_messageInfo_AccountStatusInfo.Size(m)
}
func (m *AccountStatusInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountStatusInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AccountStatusInfo proto.InternalMessageInfo

func (m *AccountStatusInfo) GetBcname() string {
	if m != nil {
		return m.Bcname
	}
	return ""
}

func (m *AccountStatusInfo) GetStatus() TransactionStatus {
	if m != nil {
		return m.Status
	}
	return TransactionStatus_UNDEFINE
}

// ///// 事件订阅请求数据结构
// EventRequest 将事件订阅统一归为EventRequest
type EventRequest struct {
	Type                 EventType `protobuf:"varint,1,opt,name=type,enum=pb.EventType" json:"type,omitempty"`
	Payload              []byte    `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *EventRequest) Reset()         { *m = EventRequest{} }
func (m *EventRequest) String() string { return proto.CompactTextString(m) }
func (*EventRequest) ProtoMessage()    {}
func (*EventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_28ecaf4b97a0a4ec, []int{3}
}
func (m *EventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventRequest.Unmarshal(m, b)
}
func (m *EventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventRequest.Marshal(b, m, deterministic)
}
func (dst *EventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRequest.Merge(dst, src)
}
func (m *EventRequest) XXX_Size() int {
	return xxx_messageInfo_EventRequest.Size(m)
}
func (m *EventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EventRequest proto.InternalMessageInfo

func (m *EventRequest) GetType() EventType {
	if m != nil {
		return m.Type
	}
	return EventType_UNDEFINED
}

func (m *EventRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// BlockEventRequest 订阅区块请求
type BlockEventRequest struct {
	Bcname               string   `protobuf:"bytes,1,opt,name=bcname" json:"bcname,omitempty"`
	Proposer             string   `protobuf:"bytes,2,opt,name=proposer" json:"proposer,omitempty"`
	StartHeight          int64    `protobuf:"varint,3,opt,name=start_height,json=startHeight" json:"start_height,omitempty"`
	EndHeight            int64    `protobuf:"varint,4,opt,name=end_height,json=endHeight" json:"end_height,omitempty"`
	NeedContent          bool     `protobuf:"varint,5,opt,name=need_content,json=needContent" json:"need_content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockEventRequest) Reset()         { *m = BlockEventRequest{} }
func (m *BlockEventRequest) String() string { return proto.CompactTextString(m) }
func (*BlockEventRequest) ProtoMessage()    {}
func (*BlockEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_28ecaf4b97a0a4ec, []int{4}
}
func (m *BlockEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockEventRequest.Unmarshal(m, b)
}
func (m *BlockEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockEventRequest.Marshal(b, m, deterministic)
}
func (dst *BlockEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockEventRequest.Merge(dst, src)
}
func (m *BlockEventRequest) XXX_Size() int {
	return xxx_messageInfo_BlockEventRequest.Size(m)
}
func (m *BlockEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BlockEventRequest proto.InternalMessageInfo

func (m *BlockEventRequest) GetBcname() string {
	if m != nil {
		return m.Bcname
	}
	return ""
}

func (m *BlockEventRequest) GetProposer() string {
	if m != nil {
		return m.Proposer
	}
	return ""
}

func (m *BlockEventRequest) GetStartHeight() int64 {
	if m != nil {
		return m.StartHeight
	}
	return 0
}

func (m *BlockEventRequest) GetEndHeight() int64 {
	if m != nil {
		return m.EndHeight
	}
	return 0
}

func (m *BlockEventRequest) GetNeedContent() bool {
	if m != nil {
		return m.NeedContent
	}
	return false
}

// TransactionEventRequest 订阅交易请求
type TransactionEventRequest struct {
	Bcname               string   `protobuf:"bytes,1,opt,name=bcname" json:"bcname,omitempty"`
	Initiator            string   `protobuf:"bytes,2,opt,name=initiator" json:"initiator,omitempty"`
	AuthRequire          []string `protobuf:"bytes,3,rep,name=auth_require,json=authRequire" json:"auth_require,omitempty"`
	NeedContent          bool     `protobuf:"varint,4,opt,name=need_content,json=needContent" json:"need_content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionEventRequest) Reset()         { *m = TransactionEventRequest{} }
func (m *TransactionEventRequest) String() string { return proto.CompactTextString(m) }
func (*TransactionEventRequest) ProtoMessage()    {}
func (*TransactionEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_28ecaf4b97a0a4ec, []int{5}
}
func (m *TransactionEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionEventRequest.Unmarshal(m, b)
}
func (m *TransactionEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionEventRequest.Marshal(b, m, deterministic)
}
func (dst *TransactionEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionEventRequest.Merge(dst, src)
}
func (m *TransactionEventRequest) XXX_Size() int {
	return xxx_messageInfo_TransactionEventRequest.Size(m)
}
func (m *TransactionEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionEventRequest proto.InternalMessageInfo

func (m *TransactionEventRequest) GetBcname() string {
	if m != nil {
		return m.Bcname
	}
	return ""
}

func (m *TransactionEventRequest) GetInitiator() string {
	if m != nil {
		return m.Initiator
	}
	return ""
}

func (m *TransactionEventRequest) GetAuthRequire() []string {
	if m != nil {
		return m.AuthRequire
	}
	return nil
}

func (m *TransactionEventRequest) GetNeedContent() bool {
	if m != nil {
		return m.NeedContent
	}
	return false
}

// AccountEventRequest 订阅账户请求
type AccountEventRequest struct {
	FromAddr             string   `protobuf:"bytes,1,opt,name=from_addr,json=fromAddr" json:"from_addr,omitempty"`
	ToAddr               string   `protobuf:"bytes,2,opt,name=to_addr,json=toAddr" json:"to_addr,omitempty"`
	NeedContent          bool     `protobuf:"varint,3,opt,name=need_content,json=needContent" json:"need_content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountEventRequest) Reset()         { *m = AccountEventRequest{} }
func (m *AccountEventRequest) String() string { return proto.CompactTextString(m) }
func (*AccountEventRequest) ProtoMessage()    {}
func (*AccountEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_28ecaf4b97a0a4ec, []int{6}
}
func (m *AccountEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountEventRequest.Unmarshal(m, b)
}
func (m *AccountEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountEventRequest.Marshal(b, m, deterministic)
}
func (dst *AccountEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountEventRequest.Merge(dst, src)
}
func (m *AccountEventRequest) XXX_Size() int {
	return xxx_messageInfo_AccountEventRequest.Size(m)
}
func (m *AccountEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountEventRequest proto.InternalMessageInfo

func (m *AccountEventRequest) GetFromAddr() string {
	if m != nil {
		return m.FromAddr
	}
	return ""
}

func (m *AccountEventRequest) GetToAddr() string {
	if m != nil {
		return m.ToAddr
	}
	return ""
}

func (m *AccountEventRequest) GetNeedContent() bool {
	if m != nil {
		return m.NeedContent
	}
	return false
}

// UnsubscribeRequest 取消事件订阅请求
type UnsubscribeRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnsubscribeRequest) Reset()         { *m = UnsubscribeRequest{} }
func (m *UnsubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*UnsubscribeRequest) ProtoMessage()    {}
func (*UnsubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_28ecaf4b97a0a4ec, []int{7}
}
func (m *UnsubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnsubscribeRequest.Unmarshal(m, b)
}
func (m *UnsubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnsubscribeRequest.Marshal(b, m, deterministic)
}
func (dst *UnsubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnsubscribeRequest.Merge(dst, src)
}
func (m *UnsubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_UnsubscribeRequest.Size(m)
}
func (m *UnsubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnsubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnsubscribeRequest proto.InternalMessageInfo

func (m *UnsubscribeRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// ////// 事件订阅返回数据结构
// ///// UnsubscribeResponse 取消订阅pb接口
type UnsubscribeResponse struct {
	Id                   string                `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Status               UnsubscribeStatusInfo `protobuf:"varint,3,opt,name=status,enum=pb.UnsubscribeStatusInfo" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UnsubscribeResponse) Reset()         { *m = UnsubscribeResponse{} }
func (m *UnsubscribeResponse) String() string { return proto.CompactTextString(m) }
func (*UnsubscribeResponse) ProtoMessage()    {}
func (*UnsubscribeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_28ecaf4b97a0a4ec, []int{8}
}
func (m *UnsubscribeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnsubscribeResponse.Unmarshal(m, b)
}
func (m *UnsubscribeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnsubscribeResponse.Marshal(b, m, deterministic)
}
func (dst *UnsubscribeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnsubscribeResponse.Merge(dst, src)
}
func (m *UnsubscribeResponse) XXX_Size() int {
	return xxx_messageInfo_UnsubscribeResponse.Size(m)
}
func (m *UnsubscribeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UnsubscribeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UnsubscribeResponse proto.InternalMessageInfo

func (m *UnsubscribeResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UnsubscribeResponse) GetStatus() UnsubscribeStatusInfo {
	if m != nil {
		return m.Status
	}
	return UnsubscribeStatusInfo_UNSUBSCRIBE_UNDEFINED
}

// /// Event 将多种事件订阅回复统一为Event, 易于扩展
type Event struct {
	Id                   string    `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type                 EventType `protobuf:"varint,2,opt,name=type,enum=pb.EventType" json:"type,omitempty"`
	Payload              []byte    `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_28ecaf4b97a0a4ec, []int{9}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (dst *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(dst, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetType() EventType {
	if m != nil {
		return m.Type
	}
	return EventType_UNDEFINED
}

func (m *Event) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// BlockEvent 订阅区块返回
type BlockEvent struct {
	Status               *BlockStatusInfo `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Block                *InternalBlock   `protobuf:"bytes,2,opt,name=block" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BlockEvent) Reset()         { *m = BlockEvent{} }
func (m *BlockEvent) String() string { return proto.CompactTextString(m) }
func (*BlockEvent) ProtoMessage()    {}
func (*BlockEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_28ecaf4b97a0a4ec, []int{10}
}
func (m *BlockEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockEvent.Unmarshal(m, b)
}
func (m *BlockEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockEvent.Marshal(b, m, deterministic)
}
func (dst *BlockEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockEvent.Merge(dst, src)
}
func (m *BlockEvent) XXX_Size() int {
	return xxx_messageInfo_BlockEvent.Size(m)
}
func (m *BlockEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockEvent.DiscardUnknown(m)
}

var xxx_messageInfo_BlockEvent proto.InternalMessageInfo

func (m *BlockEvent) GetStatus() *BlockStatusInfo {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *BlockEvent) GetBlock() *InternalBlock {
	if m != nil {
		return m.Block
	}
	return nil
}

// TransactionEvent 订阅交易返回
type TransactionEvent struct {
	Status               *TransactionStatusInfo `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Tx                   *Transaction           `protobuf:"bytes,2,opt,name=tx" json:"tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *TransactionEvent) Reset()         { *m = TransactionEvent{} }
func (m *TransactionEvent) String() string { return proto.CompactTextString(m) }
func (*TransactionEvent) ProtoMessage()    {}
func (*TransactionEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_28ecaf4b97a0a4ec, []int{11}
}
func (m *TransactionEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionEvent.Unmarshal(m, b)
}
func (m *TransactionEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionEvent.Marshal(b, m, deterministic)
}
func (dst *TransactionEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionEvent.Merge(dst, src)
}
func (m *TransactionEvent) XXX_Size() int {
	return xxx_messageInfo_TransactionEvent.Size(m)
}
func (m *TransactionEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionEvent.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionEvent proto.InternalMessageInfo

func (m *TransactionEvent) GetStatus() *TransactionStatusInfo {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *TransactionEvent) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

// AccountEvent 订阅账户返回
type AccountEvent struct {
	Status               *AccountStatusInfo `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Tx                   *Transaction       `protobuf:"bytes,2,opt,name=tx" json:"tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AccountEvent) Reset()         { *m = AccountEvent{} }
func (m *AccountEvent) String() string { return proto.CompactTextString(m) }
func (*AccountEvent) ProtoMessage()    {}
func (*AccountEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_28ecaf4b97a0a4ec, []int{12}
}
func (m *AccountEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountEvent.Unmarshal(m, b)
}
func (m *AccountEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountEvent.Marshal(b, m, deterministic)
}
func (dst *AccountEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountEvent.Merge(dst, src)
}
func (m *AccountEvent) XXX_Size() int {
	return xxx_messageInfo_AccountEvent.Size(m)
}
func (m *AccountEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountEvent.DiscardUnknown(m)
}

var xxx_messageInfo_AccountEvent proto.InternalMessageInfo

func (m *AccountEvent) GetStatus() *AccountStatusInfo {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *AccountEvent) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

func init() {
	proto.RegisterType((*BlockStatusInfo)(nil), "pb.BlockStatusInfo")
	proto.RegisterType((*TransactionStatusInfo)(nil), "pb.TransactionStatusInfo")
	proto.RegisterType((*AccountStatusInfo)(nil), "pb.AccountStatusInfo")
	proto.RegisterType((*EventRequest)(nil), "pb.EventRequest")
	proto.RegisterType((*BlockEventRequest)(nil), "pb.BlockEventRequest")
	proto.RegisterType((*TransactionEventRequest)(nil), "pb.TransactionEventRequest")
	proto.RegisterType((*AccountEventRequest)(nil), "pb.AccountEventRequest")
	proto.RegisterType((*UnsubscribeRequest)(nil), "pb.UnsubscribeRequest")
	proto.RegisterType((*UnsubscribeResponse)(nil), "pb.UnsubscribeResponse")
	proto.RegisterType((*Event)(nil), "pb.Event")
	proto.RegisterType((*BlockEvent)(nil), "pb.BlockEvent")
	proto.RegisterType((*TransactionEvent)(nil), "pb.TransactionEvent")
	proto.RegisterType((*AccountEvent)(nil), "pb.AccountEvent")
	proto.RegisterEnum("pb.BlockStatus", BlockStatus_name, BlockStatus_value)
	proto.RegisterEnum("pb.EventType", EventType_name, EventType_value)
	proto.RegisterEnum("pb.UnsubscribeStatusInfo", UnsubscribeStatusInfo_name, UnsubscribeStatusInfo_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PubsubService service

type PubsubServiceClient interface {
	Subscribe(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (PubsubService_SubscribeClient, error)
	Unsubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*UnsubscribeResponse, error)
}

type pubsubServiceClient struct {
	cc *grpc.ClientConn
}

func NewPubsubServiceClient(cc *grpc.ClientConn) PubsubServiceClient {
	return &pubsubServiceClient{cc}
}

func (c *pubsubServiceClient) Subscribe(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (PubsubService_SubscribeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PubsubService_serviceDesc.Streams[0], c.cc, "/pb.PubsubService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &pubsubServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PubsubService_SubscribeClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type pubsubServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *pubsubServiceSubscribeClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pubsubServiceClient) Unsubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*UnsubscribeResponse, error) {
	out := new(UnsubscribeResponse)
	err := grpc.Invoke(ctx, "/pb.PubsubService/Unsubscribe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PubsubService service

type PubsubServiceServer interface {
	Subscribe(*EventRequest, PubsubService_SubscribeServer) error
	Unsubscribe(context.Context, *UnsubscribeRequest) (*UnsubscribeResponse, error)
}

func RegisterPubsubServiceServer(s *grpc.Server, srv PubsubServiceServer) {
	s.RegisterService(&_PubsubService_serviceDesc, srv)
}

func _PubsubService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EventRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PubsubServiceServer).Subscribe(m, &pubsubServiceSubscribeServer{stream})
}

type PubsubService_SubscribeServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type pubsubServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *pubsubServiceSubscribeServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _PubsubService_Unsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubsubServiceServer).Unsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PubsubService/Unsubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubsubServiceServer).Unsubscribe(ctx, req.(*UnsubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PubsubService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.PubsubService",
	HandlerType: (*PubsubServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Unsubscribe",
			Handler:    _PubsubService_Unsubscribe_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _PubsubService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "event.proto",
}

func init() { proto.RegisterFile("event.proto", fileDescriptor_event_28ecaf4b97a0a4ec) }

var fileDescriptor_event_28ecaf4b97a0a4ec = []byte{
	// 754 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xed, 0x6e, 0xe2, 0x46,
	0x14, 0xad, 0x6d, 0x20, 0xf8, 0xda, 0x09, 0xce, 0x20, 0x02, 0x49, 0x5b, 0x15, 0x50, 0xa5, 0x45,
	0x54, 0x1b, 0xb5, 0xf4, 0x77, 0x2b, 0x81, 0xe3, 0xd5, 0xa2, 0x5d, 0x99, 0xd5, 0xd8, 0x48, 0xab,
	0x56, 0x2a, 0xf2, 0xc7, 0xa4, 0x58, 0xcd, 0x8e, 0xbd, 0xf6, 0x78, 0x95, 0xa8, 0x7d, 0x8d, 0xbe,
	0x45, 0x1f, 0xb2, 0xf2, 0xd8, 0x18, 0x63, 0xd2, 0x2e, 0xff, 0x98, 0x73, 0x3f, 0xce, 0x99, 0x3b,
	0xe7, 0x62, 0x50, 0xc8, 0x27, 0x42, 0xd9, 0x6d, 0x14, 0x87, 0x2c, 0x44, 0x62, 0xe4, 0xde, 0xa8,
	0x8f, 0xde, 0xd6, 0x09, 0x68, 0x8e, 0x8c, 0x31, 0x74, 0x16, 0x0f, 0xa1, 0xf7, 0x87, 0xc5, 0x1c,
	0x96, 0x26, 0x4b, 0x7a, 0x1f, 0xa2, 0x2b, 0x68, 0xb9, 0x1e, 0x75, 0x3e, 0x90, 0x81, 0x30, 0x14,
	0x26, 0x32, 0x2e, 0x4e, 0xe8, 0x05, 0xb4, 0x12, 0x9e, 0x35, 0x10, 0x87, 0xc2, 0xe4, 0x62, 0xd6,
	0xb9, 0x8d, 0xdc, 0xdb, 0x4a, 0x31, 0x2e, 0xc2, 0xe3, 0xbf, 0xa0, 0x67, 0xc7, 0x0e, 0x4d, 0x1c,
	0x8f, 0x05, 0x21, 0x3d, 0xa1, 0xf3, 0x57, 0x20, 0x07, 0x34, 0x60, 0x81, 0xc3, 0xc2, 0x98, 0x37,
	0x97, 0xf1, 0x1e, 0x40, 0x2f, 0x4b, 0x5e, 0x89, 0xf3, 0xf6, 0x32, 0xde, 0x23, 0x82, 0x92, 0xfd,
	0x17, 0xb8, 0x9c, 0x7b, 0x5e, 0x98, 0x52, 0x76, 0x02, 0xf3, 0xcb, 0xda, 0x9d, 0x3e, 0xd3, 0xfb,
	0x0d, 0xa8, 0x46, 0x36, 0x4e, 0x4c, 0x3e, 0xa6, 0x24, 0x61, 0x68, 0x04, 0x0d, 0xf6, 0x14, 0xe5,
	0x4d, 0x2f, 0x66, 0xe7, 0x59, 0x31, 0x8f, 0xdb, 0x4f, 0x11, 0xc1, 0x3c, 0x84, 0x06, 0x70, 0x16,
	0x39, 0x4f, 0x0f, 0xa1, 0xe3, 0x73, 0x0a, 0x15, 0xef, 0x8e, 0xe3, 0x7f, 0x04, 0xb8, 0xe4, 0xe3,
	0x3b, 0x68, 0xf9, 0x5f, 0x4a, 0x6f, 0xa0, 0x1d, 0xc5, 0x61, 0x14, 0x26, 0x64, 0x37, 0xa2, 0xf2,
	0x8c, 0x46, 0xa0, 0x26, 0xcc, 0x89, 0xd9, 0x66, 0x4b, 0x82, 0xdf, 0xb7, 0x8c, 0xcf, 0x49, 0xc2,
	0x0a, 0xc7, 0x5e, 0x73, 0x08, 0x7d, 0x0d, 0x40, 0xa8, 0xbf, 0x4b, 0x68, 0xf0, 0x04, 0x99, 0x50,
	0xbf, 0x08, 0x8f, 0x40, 0xa5, 0x84, 0xf8, 0x1b, 0x2f, 0xa4, 0x8c, 0x50, 0x36, 0x68, 0x0e, 0x85,
	0x49, 0x1b, 0x2b, 0x19, 0xa6, 0xe7, 0xd0, 0xf8, 0x6f, 0x01, 0xfa, 0x95, 0xc9, 0x9c, 0x24, 0xfa,
	0xff, 0x1f, 0x76, 0x04, 0xaa, 0x93, 0xb2, 0xed, 0x26, 0x26, 0x1f, 0xd3, 0x20, 0x26, 0x03, 0x69,
	0x28, 0x4d, 0x64, 0xac, 0x64, 0x18, 0xce, 0xa1, 0x23, 0x5d, 0x8d, 0x63, 0x5d, 0x14, 0xba, 0xc5,
	0x7b, 0x1f, 0x48, 0xfa, 0x12, 0xe4, 0xfb, 0x38, 0xfc, 0xb0, 0x71, 0x7c, 0x3f, 0x2e, 0x54, 0xb5,
	0x33, 0x60, 0xee, 0xfb, 0x31, 0xea, 0xc3, 0x19, 0x0b, 0xf3, 0x50, 0xae, 0xaa, 0xc5, 0x42, 0x1e,
	0xa8, 0xf3, 0x49, 0xc7, 0x7c, 0xdf, 0x02, 0x5a, 0xd3, 0x24, 0x75, 0x13, 0x2f, 0x0e, 0x5c, 0xb2,
	0xa3, 0xbb, 0x00, 0x31, 0xf0, 0x0b, 0x1e, 0x31, 0xf0, 0xc7, 0xef, 0xa1, 0x7b, 0x90, 0x95, 0x44,
	0x21, 0x4d, 0x48, 0x3d, 0x0d, 0xfd, 0x50, 0xf3, 0xf6, 0x75, 0x66, 0xa1, 0x4a, 0xe1, 0xde, 0xc2,
	0xa5, 0x07, 0x6d, 0x68, 0xf2, 0x8b, 0x1e, 0xf5, 0xda, 0x99, 0x51, 0x3c, 0xc9, 0x8c, 0xd2, 0xa1,
	0x19, 0x5d, 0x80, 0xbd, 0x17, 0xd1, 0x77, 0xa5, 0xac, 0xac, 0xbd, 0x32, 0xeb, 0xd6, 0x56, 0xbd,
	0x2a, 0x08, 0xbd, 0x80, 0xa6, 0x9b, 0x85, 0x38, 0xb1, 0x32, 0xbb, 0xcc, 0x72, 0x97, 0x94, 0x91,
	0x98, 0x3a, 0x0f, 0xbc, 0x06, 0xe7, 0xf1, 0xf1, 0x3d, 0x68, 0x75, 0x03, 0x55, 0x06, 0x90, 0x33,
	0x5d, 0x3f, 0xbb, 0x80, 0x07, 0x7c, 0xdf, 0x80, 0xc8, 0x1e, 0x0b, 0xb2, 0x4e, 0x2d, 0x1d, 0x8b,
	0xec, 0x71, 0xfc, 0x1b, 0xa8, 0x55, 0x47, 0x54, 0x96, 0x3c, 0xe7, 0xe0, 0x4b, 0x7e, 0xf4, 0x1f,
	0x71, 0x72, 0xff, 0xe9, 0x4f, 0xa0, 0x54, 0x66, 0x81, 0x64, 0x68, 0x1a, 0x18, 0xaf, 0xb0, 0xf6,
	0x45, 0xf6, 0xd3, 0xc6, 0x6b, 0xf3, 0x8d, 0x26, 0x20, 0x80, 0xd6, 0x02, 0xcf, 0x4d, 0xfd, 0xb5,
	0x26, 0x22, 0x15, 0xda, 0xe6, 0xca, 0x34, 0xde, 0x2f, 0x2d, 0x5b, 0x93, 0xa6, 0xbf, 0x82, 0x5c,
	0xbe, 0x0b, 0x3a, 0x07, 0x79, 0x6d, 0xde, 0x19, 0xaf, 0x96, 0xa6, 0x71, 0x97, 0x37, 0x58, 0xbc,
	0x5d, 0xe9, 0x59, 0x83, 0x0e, 0x28, 0x36, 0x9e, 0x9b, 0xd6, 0x5c, 0xb7, 0x97, 0x2b, 0x53, 0x13,
	0x91, 0x02, 0x67, 0x73, 0x5d, 0x5f, 0xad, 0x4d, 0x5b, 0x93, 0xd0, 0x15, 0x20, 0x6b, 0xbd, 0xb0,
	0x74, 0xbc, 0x5c, 0x18, 0x1b, 0x6c, 0x58, 0xef, 0x56, 0xa6, 0x65, 0x68, 0x8d, 0xa9, 0x07, 0xbd,
	0x67, 0xed, 0x83, 0xae, 0xa1, 0xb7, 0x36, 0xf7, 0x25, 0x55, 0xd2, 0x3e, 0x74, 0xab, 0x21, 0x6b,
	0xad, 0xeb, 0x86, 0x65, 0x69, 0x42, 0x46, 0x52, 0x0d, 0xbc, 0x9a, 0x2f, 0xdf, 0x1a, 0x77, 0x9a,
	0x38, 0xfb, 0x13, 0xce, 0xdf, 0xa5, 0x6e, 0x92, 0xba, 0x16, 0x89, 0x3f, 0x05, 0x1e, 0x41, 0x53,
	0x90, 0xad, 0x1d, 0x27, 0xd2, 0x4a, 0xe7, 0x15, 0xcb, 0x71, 0x23, 0x97, 0xc8, 0xf7, 0x02, 0xfa,
	0x19, 0x94, 0x8a, 0x42, 0x74, 0x55, 0x73, 0xfc, 0xae, 0xa6, 0x7f, 0x84, 0xe7, 0x2b, 0xe4, 0xb6,
	0xf8, 0x87, 0xeb, 0xc7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x91, 0x57, 0xf4, 0xa4, 0xd9, 0x06,
	0x00, 0x00,
}
