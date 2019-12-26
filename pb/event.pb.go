// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
	return fileDescriptor_2d17a9d3f0ddf27e, []int{0}
}

///// 状态信息
type EventType int32

const (
	EventType_NULL_EVENT        EventType = 0
	EventType_BLOCK_EVENT       EventType = 1
	EventType_TRANSACTION_EVENT EventType = 2
	EventType_ACCOUNT_EVENT     EventType = 3
	EventType_UNSUBSCRIBE_EVENT EventType = 4
)

var EventType_name = map[int32]string{
	0: "NULL_EVENT",
	1: "BLOCK_EVENT",
	2: "TRANSACTION_EVENT",
	3: "ACCOUNT_EVENT",
	4: "UNSUBSCRIBE_EVENT",
}

var EventType_value = map[string]int32{
	"NULL_EVENT":        0,
	"BLOCK_EVENT":       1,
	"TRANSACTION_EVENT": 2,
	"ACCOUNT_EVENT":     3,
	"UNSUBSCRIBE_EVENT": 4,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}

func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{1}
}

type UnSubscribeStatusInfo int32

const (
	UnSubscribeStatusInfo_UNSUBSCRIBE_UNDEFINED UnSubscribeStatusInfo = 0
	UnSubscribeStatusInfo_UNSUBSCRIBE_SUCCESS   UnSubscribeStatusInfo = 1
	UnSubscribeStatusInfo_UNSUBSCRIBE_FAILED    UnSubscribeStatusInfo = 2
)

var UnSubscribeStatusInfo_name = map[int32]string{
	0: "UNSUBSCRIBE_UNDEFINED",
	1: "UNSUBSCRIBE_SUCCESS",
	2: "UNSUBSCRIBE_FAILED",
}

var UnSubscribeStatusInfo_value = map[string]int32{
	"UNSUBSCRIBE_UNDEFINED": 0,
	"UNSUBSCRIBE_SUCCESS":   1,
	"UNSUBSCRIBE_FAILED":    2,
}

func (x UnSubscribeStatusInfo) String() string {
	return proto.EnumName(UnSubscribeStatusInfo_name, int32(x))
}

func (UnSubscribeStatusInfo) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{2}
}

/////// 事件订阅请求数据结构
// BlockEventRequest 订阅区块请求
type BlockEventRequest struct {
	BcName               string   `protobuf:"bytes,1,opt,name=bc_name,json=bcName,proto3" json:"bc_name,omitempty"`
	Proposer             string   `protobuf:"bytes,2,opt,name=proposer,proto3" json:"proposer,omitempty"`
	StartHeight          int64    `protobuf:"varint,3,opt,name=start_height,json=startHeight,proto3" json:"start_height,omitempty"`
	EndHeight            int64    `protobuf:"varint,4,opt,name=end_height,json=endHeight,proto3" json:"end_height,omitempty"`
	NeedContent          bool     `protobuf:"varint,5,opt,name=need_content,json=needContent,proto3" json:"need_content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockEventRequest) Reset()         { *m = BlockEventRequest{} }
func (m *BlockEventRequest) String() string { return proto.CompactTextString(m) }
func (*BlockEventRequest) ProtoMessage()    {}
func (*BlockEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{0}
}

func (m *BlockEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockEventRequest.Unmarshal(m, b)
}
func (m *BlockEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockEventRequest.Marshal(b, m, deterministic)
}
func (m *BlockEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockEventRequest.Merge(m, src)
}
func (m *BlockEventRequest) XXX_Size() int {
	return xxx_messageInfo_BlockEventRequest.Size(m)
}
func (m *BlockEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BlockEventRequest proto.InternalMessageInfo

func (m *BlockEventRequest) GetBcName() string {
	if m != nil {
		return m.BcName
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
	BcName               string   `protobuf:"bytes,1,opt,name=bc_name,json=bcName,proto3" json:"bc_name,omitempty"`
	Initiator            string   `protobuf:"bytes,2,opt,name=initiator,proto3" json:"initiator,omitempty"`
	AuthRequire          []string `protobuf:"bytes,3,rep,name=auth_require,json=authRequire,proto3" json:"auth_require,omitempty"`
	NeedContent          bool     `protobuf:"varint,4,opt,name=need_content,json=needContent,proto3" json:"need_content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionEventRequest) Reset()         { *m = TransactionEventRequest{} }
func (m *TransactionEventRequest) String() string { return proto.CompactTextString(m) }
func (*TransactionEventRequest) ProtoMessage()    {}
func (*TransactionEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{1}
}

func (m *TransactionEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionEventRequest.Unmarshal(m, b)
}
func (m *TransactionEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionEventRequest.Marshal(b, m, deterministic)
}
func (m *TransactionEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionEventRequest.Merge(m, src)
}
func (m *TransactionEventRequest) XXX_Size() int {
	return xxx_messageInfo_TransactionEventRequest.Size(m)
}
func (m *TransactionEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionEventRequest proto.InternalMessageInfo

func (m *TransactionEventRequest) GetBcName() string {
	if m != nil {
		return m.BcName
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
	FromAddr             string   `protobuf:"bytes,1,opt,name=from_addr,json=fromAddr,proto3" json:"from_addr,omitempty"`
	ToAddr               string   `protobuf:"bytes,2,opt,name=to_addr,json=toAddr,proto3" json:"to_addr,omitempty"`
	NeedContent          bool     `protobuf:"varint,3,opt,name=need_content,json=needContent,proto3" json:"need_content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountEventRequest) Reset()         { *m = AccountEventRequest{} }
func (m *AccountEventRequest) String() string { return proto.CompactTextString(m) }
func (*AccountEventRequest) ProtoMessage()    {}
func (*AccountEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{2}
}

func (m *AccountEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountEventRequest.Unmarshal(m, b)
}
func (m *AccountEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountEventRequest.Marshal(b, m, deterministic)
}
func (m *AccountEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountEventRequest.Merge(m, src)
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

// UnSubscribeRequest 取消事件订阅请求
type UnSubscribeRequest struct {
	EventId              string   `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnSubscribeRequest) Reset()         { *m = UnSubscribeRequest{} }
func (m *UnSubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*UnSubscribeRequest) ProtoMessage()    {}
func (*UnSubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{3}
}

func (m *UnSubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnSubscribeRequest.Unmarshal(m, b)
}
func (m *UnSubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnSubscribeRequest.Marshal(b, m, deterministic)
}
func (m *UnSubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnSubscribeRequest.Merge(m, src)
}
func (m *UnSubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_UnSubscribeRequest.Size(m)
}
func (m *UnSubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnSubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnSubscribeRequest proto.InternalMessageInfo

func (m *UnSubscribeRequest) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

// EventRequest 将上述几种事件订阅统一归为EventRequest
type EventRequest struct {
	EventType            EventType `protobuf:"varint,1,opt,name=event_type,json=eventType,proto3,enum=pb.EventType" json:"event_type,omitempty"`
	EventRequest         []byte    `protobuf:"bytes,2,opt,name=event_request,json=eventRequest,proto3" json:"event_request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *EventRequest) Reset()         { *m = EventRequest{} }
func (m *EventRequest) String() string { return proto.CompactTextString(m) }
func (*EventRequest) ProtoMessage()    {}
func (*EventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{4}
}

func (m *EventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventRequest.Unmarshal(m, b)
}
func (m *EventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventRequest.Marshal(b, m, deterministic)
}
func (m *EventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRequest.Merge(m, src)
}
func (m *EventRequest) XXX_Size() int {
	return xxx_messageInfo_EventRequest.Size(m)
}
func (m *EventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EventRequest proto.InternalMessageInfo

func (m *EventRequest) GetEventType() EventType {
	if m != nil {
		return m.EventType
	}
	return EventType_NULL_EVENT
}

func (m *EventRequest) GetEventRequest() []byte {
	if m != nil {
		return m.EventRequest
	}
	return nil
}

//////// 事件订阅返回数据结构
// SubscribeResponse 订阅产生的event_id
type SubscribeResponse struct {
	EventRequest         *EventRequest `protobuf:"bytes,1,opt,name=event_request,json=eventRequest,proto3" json:"event_request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SubscribeResponse) Reset()         { *m = SubscribeResponse{} }
func (m *SubscribeResponse) String() string { return proto.CompactTextString(m) }
func (*SubscribeResponse) ProtoMessage()    {}
func (*SubscribeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{5}
}

func (m *SubscribeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeResponse.Unmarshal(m, b)
}
func (m *SubscribeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeResponse.Marshal(b, m, deterministic)
}
func (m *SubscribeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeResponse.Merge(m, src)
}
func (m *SubscribeResponse) XXX_Size() int {
	return xxx_messageInfo_SubscribeResponse.Size(m)
}
func (m *SubscribeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeResponse proto.InternalMessageInfo

func (m *SubscribeResponse) GetEventRequest() *EventRequest {
	if m != nil {
		return m.EventRequest
	}
	return nil
}

type UnSubscribeResponse struct {
	Status               UnSubscribeStatusInfo `protobuf:"varint,1,opt,name=status,proto3,enum=pb.UnSubscribeStatusInfo" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UnSubscribeResponse) Reset()         { *m = UnSubscribeResponse{} }
func (m *UnSubscribeResponse) String() string { return proto.CompactTextString(m) }
func (*UnSubscribeResponse) ProtoMessage()    {}
func (*UnSubscribeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{6}
}

func (m *UnSubscribeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnSubscribeResponse.Unmarshal(m, b)
}
func (m *UnSubscribeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnSubscribeResponse.Marshal(b, m, deterministic)
}
func (m *UnSubscribeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnSubscribeResponse.Merge(m, src)
}
func (m *UnSubscribeResponse) XXX_Size() int {
	return xxx_messageInfo_UnSubscribeResponse.Size(m)
}
func (m *UnSubscribeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UnSubscribeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UnSubscribeResponse proto.InternalMessageInfo

func (m *UnSubscribeResponse) GetStatus() UnSubscribeStatusInfo {
	if m != nil {
		return m.Status
	}
	return UnSubscribeStatusInfo_UNSUBSCRIBE_UNDEFINED
}

// BlockEventResponse 订阅区块返回
type BlockEventResponse struct {
	Status               *BlockStatusInfo `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Block                *InternalBlock   `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BlockEventResponse) Reset()         { *m = BlockEventResponse{} }
func (m *BlockEventResponse) String() string { return proto.CompactTextString(m) }
func (*BlockEventResponse) ProtoMessage()    {}
func (*BlockEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{7}
}

func (m *BlockEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockEventResponse.Unmarshal(m, b)
}
func (m *BlockEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockEventResponse.Marshal(b, m, deterministic)
}
func (m *BlockEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockEventResponse.Merge(m, src)
}
func (m *BlockEventResponse) XXX_Size() int {
	return xxx_messageInfo_BlockEventResponse.Size(m)
}
func (m *BlockEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BlockEventResponse proto.InternalMessageInfo

func (m *BlockEventResponse) GetStatus() *BlockStatusInfo {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *BlockEventResponse) GetBlock() *InternalBlock {
	if m != nil {
		return m.Block
	}
	return nil
}

// TransactionEventResponse 订阅交易返回
type TransactionEventResponse struct {
	Status               *TransactionStatusInfo `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Tx                   *Transaction           `protobuf:"bytes,2,opt,name=tx,proto3" json:"tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *TransactionEventResponse) Reset()         { *m = TransactionEventResponse{} }
func (m *TransactionEventResponse) String() string { return proto.CompactTextString(m) }
func (*TransactionEventResponse) ProtoMessage()    {}
func (*TransactionEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{8}
}

func (m *TransactionEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionEventResponse.Unmarshal(m, b)
}
func (m *TransactionEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionEventResponse.Marshal(b, m, deterministic)
}
func (m *TransactionEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionEventResponse.Merge(m, src)
}
func (m *TransactionEventResponse) XXX_Size() int {
	return xxx_messageInfo_TransactionEventResponse.Size(m)
}
func (m *TransactionEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionEventResponse proto.InternalMessageInfo

func (m *TransactionEventResponse) GetStatus() *TransactionStatusInfo {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *TransactionEventResponse) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

// AccountEventResponse 订阅账户返回
type AccountEventResponse struct {
	Status               *AccountStatusInfo `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Tx                   *Transaction       `protobuf:"bytes,2,opt,name=tx,proto3" json:"tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AccountEventResponse) Reset()         { *m = AccountEventResponse{} }
func (m *AccountEventResponse) String() string { return proto.CompactTextString(m) }
func (*AccountEventResponse) ProtoMessage()    {}
func (*AccountEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{9}
}

func (m *AccountEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountEventResponse.Unmarshal(m, b)
}
func (m *AccountEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountEventResponse.Marshal(b, m, deterministic)
}
func (m *AccountEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountEventResponse.Merge(m, src)
}
func (m *AccountEventResponse) XXX_Size() int {
	return xxx_messageInfo_AccountEventResponse.Size(m)
}
func (m *AccountEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountEventResponse proto.InternalMessageInfo

func (m *AccountEventResponse) GetStatus() *AccountStatusInfo {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *AccountEventResponse) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

////// EventResponse 将上述几类统一归为EventResponse
type EventResponse struct {
	EventId              string    `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	EventType            EventType `protobuf:"varint,2,opt,name=event_type,json=eventType,proto3,enum=pb.EventType" json:"event_type,omitempty"`
	EventResponse        []byte    `protobuf:"bytes,3,opt,name=event_response,json=eventResponse,proto3" json:"event_response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *EventResponse) Reset()         { *m = EventResponse{} }
func (m *EventResponse) String() string { return proto.CompactTextString(m) }
func (*EventResponse) ProtoMessage()    {}
func (*EventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{10}
}

func (m *EventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventResponse.Unmarshal(m, b)
}
func (m *EventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventResponse.Marshal(b, m, deterministic)
}
func (m *EventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventResponse.Merge(m, src)
}
func (m *EventResponse) XXX_Size() int {
	return xxx_messageInfo_EventResponse.Size(m)
}
func (m *EventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventResponse proto.InternalMessageInfo

func (m *EventResponse) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *EventResponse) GetEventType() EventType {
	if m != nil {
		return m.EventType
	}
	return EventType_NULL_EVENT
}

func (m *EventResponse) GetEventResponse() []byte {
	if m != nil {
		return m.EventResponse
	}
	return nil
}

// BlockStatusInfo 区块元数据
type BlockStatusInfo struct {
	BcName               string      `protobuf:"bytes,1,opt,name=bc_name,json=bcName,proto3" json:"bc_name,omitempty"`
	Status               BlockStatus `protobuf:"varint,2,opt,name=status,proto3,enum=pb.BlockStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *BlockStatusInfo) Reset()         { *m = BlockStatusInfo{} }
func (m *BlockStatusInfo) String() string { return proto.CompactTextString(m) }
func (*BlockStatusInfo) ProtoMessage()    {}
func (*BlockStatusInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{11}
}

func (m *BlockStatusInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockStatusInfo.Unmarshal(m, b)
}
func (m *BlockStatusInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockStatusInfo.Marshal(b, m, deterministic)
}
func (m *BlockStatusInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockStatusInfo.Merge(m, src)
}
func (m *BlockStatusInfo) XXX_Size() int {
	return xxx_messageInfo_BlockStatusInfo.Size(m)
}
func (m *BlockStatusInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockStatusInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BlockStatusInfo proto.InternalMessageInfo

func (m *BlockStatusInfo) GetBcName() string {
	if m != nil {
		return m.BcName
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
	BcName               string            `protobuf:"bytes,1,opt,name=bc_name,json=bcName,proto3" json:"bc_name,omitempty"`
	Initiator            string            `protobuf:"bytes,2,opt,name=initiator,proto3" json:"initiator,omitempty"`
	Status               TransactionStatus `protobuf:"varint,3,opt,name=status,proto3,enum=pb.TransactionStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TransactionStatusInfo) Reset()         { *m = TransactionStatusInfo{} }
func (m *TransactionStatusInfo) String() string { return proto.CompactTextString(m) }
func (*TransactionStatusInfo) ProtoMessage()    {}
func (*TransactionStatusInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{12}
}

func (m *TransactionStatusInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionStatusInfo.Unmarshal(m, b)
}
func (m *TransactionStatusInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionStatusInfo.Marshal(b, m, deterministic)
}
func (m *TransactionStatusInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionStatusInfo.Merge(m, src)
}
func (m *TransactionStatusInfo) XXX_Size() int {
	return xxx_messageInfo_TransactionStatusInfo.Size(m)
}
func (m *TransactionStatusInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionStatusInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionStatusInfo proto.InternalMessageInfo

func (m *TransactionStatusInfo) GetBcName() string {
	if m != nil {
		return m.BcName
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
	BcName               string            `protobuf:"bytes,1,opt,name=bc_name,json=bcName,proto3" json:"bc_name,omitempty"`
	Status               TransactionStatus `protobuf:"varint,2,opt,name=status,proto3,enum=pb.TransactionStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AccountStatusInfo) Reset()         { *m = AccountStatusInfo{} }
func (m *AccountStatusInfo) String() string { return proto.CompactTextString(m) }
func (*AccountStatusInfo) ProtoMessage()    {}
func (*AccountStatusInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{13}
}

func (m *AccountStatusInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountStatusInfo.Unmarshal(m, b)
}
func (m *AccountStatusInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountStatusInfo.Marshal(b, m, deterministic)
}
func (m *AccountStatusInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountStatusInfo.Merge(m, src)
}
func (m *AccountStatusInfo) XXX_Size() int {
	return xxx_messageInfo_AccountStatusInfo.Size(m)
}
func (m *AccountStatusInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountStatusInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AccountStatusInfo proto.InternalMessageInfo

func (m *AccountStatusInfo) GetBcName() string {
	if m != nil {
		return m.BcName
	}
	return ""
}

func (m *AccountStatusInfo) GetStatus() TransactionStatus {
	if m != nil {
		return m.Status
	}
	return TransactionStatus_UNDEFINE
}

func init() {
	proto.RegisterEnum("pb.BlockStatus", BlockStatus_name, BlockStatus_value)
	proto.RegisterEnum("pb.EventType", EventType_name, EventType_value)
	proto.RegisterEnum("pb.UnSubscribeStatusInfo", UnSubscribeStatusInfo_name, UnSubscribeStatusInfo_value)
	proto.RegisterType((*BlockEventRequest)(nil), "pb.BlockEventRequest")
	proto.RegisterType((*TransactionEventRequest)(nil), "pb.TransactionEventRequest")
	proto.RegisterType((*AccountEventRequest)(nil), "pb.AccountEventRequest")
	proto.RegisterType((*UnSubscribeRequest)(nil), "pb.UnSubscribeRequest")
	proto.RegisterType((*EventRequest)(nil), "pb.EventRequest")
	proto.RegisterType((*SubscribeResponse)(nil), "pb.SubscribeResponse")
	proto.RegisterType((*UnSubscribeResponse)(nil), "pb.UnSubscribeResponse")
	proto.RegisterType((*BlockEventResponse)(nil), "pb.BlockEventResponse")
	proto.RegisterType((*TransactionEventResponse)(nil), "pb.TransactionEventResponse")
	proto.RegisterType((*AccountEventResponse)(nil), "pb.AccountEventResponse")
	proto.RegisterType((*EventResponse)(nil), "pb.EventResponse")
	proto.RegisterType((*BlockStatusInfo)(nil), "pb.BlockStatusInfo")
	proto.RegisterType((*TransactionStatusInfo)(nil), "pb.TransactionStatusInfo")
	proto.RegisterType((*AccountStatusInfo)(nil), "pb.AccountStatusInfo")
}

func init() { proto.RegisterFile("event.proto", fileDescriptor_2d17a9d3f0ddf27e) }

var fileDescriptor_2d17a9d3f0ddf27e = []byte{
	// 790 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x5d, 0x6f, 0xe3, 0x44,
	0x14, 0xad, 0xed, 0x36, 0x1b, 0x5f, 0x3b, 0xad, 0x33, 0x25, 0x34, 0x2d, 0x20, 0xba, 0x46, 0x68,
	0xa3, 0xc2, 0x16, 0x08, 0xe2, 0x91, 0x87, 0xc4, 0xf5, 0xaa, 0x66, 0x23, 0x07, 0x8d, 0x6d, 0x84,
	0xc4, 0x43, 0xe4, 0x8f, 0x29, 0x31, 0x6c, 0xc6, 0x5e, 0x7b, 0xbc, 0x74, 0x1f, 0x10, 0xff, 0x82,
	0x7f, 0xc1, 0x7f, 0x44, 0x1e, 0x3b, 0x59, 0xc7, 0x49, 0x61, 0xd5, 0xb7, 0xf8, 0xdc, 0x73, 0xef,
	0x39, 0x33, 0xf7, 0xd8, 0x01, 0x85, 0xbc, 0x21, 0x94, 0x5d, 0xa7, 0x59, 0xc2, 0x12, 0x24, 0xa6,
	0xc1, 0x85, 0x7a, 0x1f, 0x2e, 0xfd, 0x98, 0x56, 0x88, 0xfe, 0x8f, 0x00, 0xfd, 0xe9, 0xab, 0x24,
	0xfc, 0xdd, 0x2c, 0x69, 0x98, 0xbc, 0x2e, 0x48, 0xce, 0xd0, 0x19, 0x3c, 0x09, 0xc2, 0x05, 0xf5,
	0x57, 0x64, 0x28, 0x5c, 0x0a, 0x23, 0x19, 0x77, 0x82, 0xd0, 0xf6, 0x57, 0x04, 0x5d, 0x40, 0x37,
	0xcd, 0x92, 0x34, 0xc9, 0x49, 0x36, 0x14, 0x79, 0x65, 0xf3, 0x8c, 0x9e, 0x82, 0x9a, 0x33, 0x3f,
	0x63, 0x8b, 0x25, 0x89, 0x7f, 0x5d, 0xb2, 0xa1, 0x74, 0x29, 0x8c, 0x24, 0xac, 0x70, 0xec, 0x96,
	0x43, 0xe8, 0x13, 0x00, 0x42, 0xa3, 0x35, 0xe1, 0x90, 0x13, 0x64, 0x42, 0xa3, 0xba, 0xfc, 0x14,
	0x54, 0x4a, 0x48, 0xb4, 0x08, 0x13, 0xca, 0x08, 0x65, 0xc3, 0xa3, 0x4b, 0x61, 0xd4, 0xc5, 0x4a,
	0x89, 0x19, 0x15, 0xa4, 0xff, 0x2d, 0xc0, 0x99, 0x9b, 0xf9, 0x34, 0xf7, 0x43, 0x16, 0x27, 0xf4,
	0xfd, 0x5c, 0x7f, 0x0c, 0x72, 0x4c, 0x63, 0x16, 0xfb, 0x2c, 0x59, 0xdb, 0x7e, 0x07, 0x94, 0xaa,
	0x7e, 0xc1, 0x96, 0x8b, 0x8c, 0xbc, 0x2e, 0xe2, 0x8c, 0x0c, 0xa5, 0x4b, 0x69, 0x24, 0x63, 0xa5,
	0xc4, 0x70, 0x05, 0xed, 0x18, 0x3b, 0xdc, 0x35, 0x46, 0xe1, 0x74, 0x12, 0x86, 0x49, 0x41, 0xd9,
	0x96, 0xa7, 0x8f, 0x40, 0xbe, 0xcb, 0x92, 0xd5, 0xc2, 0x8f, 0xa2, 0xac, 0x76, 0xd5, 0x2d, 0x81,
	0x49, 0x14, 0x65, 0xa5, 0x61, 0x96, 0x54, 0xa5, 0xca, 0x55, 0x87, 0x25, 0xbc, 0xd0, 0xd6, 0x93,
	0x76, 0xf5, 0xbe, 0x02, 0xe4, 0x51, 0xa7, 0x08, 0xf2, 0x30, 0x8b, 0x03, 0xb2, 0x96, 0x3b, 0x87,
	0x2e, 0xdf, 0xf7, 0x22, 0x8e, 0x6a, 0xb5, 0x27, 0xfc, 0xd9, 0x8a, 0x74, 0x1f, 0xd4, 0x2d, 0x67,
	0x5f, 0x02, 0x54, 0x54, 0xf6, 0x36, 0xad, 0x2e, 0xec, 0x78, 0xdc, 0xbb, 0x4e, 0x83, 0x6b, 0xce,
	0x72, 0xdf, 0xa6, 0x04, 0xcb, 0x64, 0xfd, 0x13, 0x7d, 0x06, 0xbd, 0x8a, 0x9d, 0x55, 0xed, 0xdc,
	0xb0, 0x8a, 0x55, 0xd2, 0x18, 0xa9, 0xff, 0x00, 0xfd, 0x86, 0xa3, 0x3c, 0x4d, 0x68, 0x4e, 0xd0,
	0x77, 0xed, 0xce, 0x52, 0x4a, 0x19, 0x6b, 0x1b, 0xa9, 0xba, 0xbb, 0x35, 0xeb, 0x16, 0x4e, 0xb7,
	0xce, 0x57, 0x4f, 0xfb, 0x06, 0x3a, 0x39, 0xf3, 0x59, 0x91, 0xd7, 0x8e, 0xcf, 0xcb, 0x31, 0x0d,
	0xa2, 0xc3, 0x8b, 0x16, 0xbd, 0x4b, 0x70, 0x4d, 0xd4, 0x7f, 0x03, 0xd4, 0x4c, 0x78, 0x3d, 0xe8,
	0x8b, 0xad, 0x41, 0xca, 0xf8, 0xb4, 0x1c, 0xc4, 0x79, 0xbb, 0x23, 0xd0, 0x33, 0x38, 0x0a, 0xca,
	0x12, 0x3f, 0xb5, 0x32, 0xee, 0x97, 0x5c, 0x8b, 0x32, 0x92, 0x51, 0xff, 0x15, 0xef, 0xc1, 0x55,
	0x5d, 0xa7, 0x30, 0xdc, 0x4d, 0xe7, 0x5e, 0xeb, 0x4a, 0x65, 0xbd, 0xc1, 0xde, 0xa3, 0xfb, 0x29,
	0x88, 0xec, 0xbe, 0x16, 0x3d, 0x69, 0xd1, 0xb1, 0xc8, 0xee, 0xf5, 0x3b, 0xf8, 0x60, 0x3b, 0x75,
	0xb5, 0xd6, 0xf3, 0x96, 0xd6, 0xa0, 0x6c, 0xae, 0x99, 0x8f, 0xd1, 0xf9, 0x0b, 0x7a, 0xdb, 0x02,
	0x0f, 0x07, 0xad, 0x15, 0x2c, 0xf1, 0x7f, 0x82, 0xf5, 0x39, 0x1c, 0xaf, 0xe3, 0x51, 0x8d, 0xe6,
	0x61, 0x57, 0x71, 0x8f, 0x34, 0xf5, 0x74, 0x07, 0x4e, 0x5a, 0xcb, 0x79, 0xf8, 0x75, 0x7f, 0xb6,
	0x39, 0x7c, 0x25, 0x7e, 0xd2, 0x5a, 0xed, 0x26, 0x19, 0x7f, 0xc2, 0x60, 0xef, 0xfd, 0x3f, 0xf6,
	0x4b, 0xf2, 0xee, 0xd6, 0x25, 0x2e, 0x3c, 0xd8, 0xbb, 0xe1, 0x8d, 0xfc, 0x2f, 0xd0, 0xdf, 0x59,
	0xc9, 0xc3, 0xd2, 0xcf, 0x5b, 0xa7, 0xfa, 0xef, 0xe1, 0x57, 0xdf, 0x83, 0xd2, 0x38, 0x32, 0x92,
	0xe1, 0xc8, 0xc4, 0x78, 0x8e, 0xb5, 0x83, 0xf2, 0xa7, 0x8b, 0x3d, 0xfb, 0xa5, 0x26, 0x20, 0x80,
	0xce, 0x14, 0x4f, 0x6c, 0xe3, 0x56, 0x13, 0x91, 0x0a, 0x5d, 0x7b, 0x6e, 0x9b, 0x3f, 0x5b, 0x8e,
	0xab, 0x49, 0x57, 0x2b, 0x90, 0x37, 0xeb, 0x42, 0xc7, 0x00, 0xb6, 0x37, 0x9b, 0x2d, 0xcc, 0x9f,
	0x4c, 0xdb, 0xd5, 0x0e, 0xd0, 0x09, 0x28, 0xd3, 0xd9, 0xdc, 0x78, 0x59, 0x03, 0x02, 0x1a, 0x40,
	0xdf, 0xc5, 0x13, 0xdb, 0x99, 0x18, 0xae, 0x35, 0xb7, 0x6b, 0x58, 0x44, 0x7d, 0xe8, 0x4d, 0x0c,
	0x63, 0xee, 0xd9, 0x6e, 0x0d, 0x49, 0x25, 0xd3, 0xb3, 0x1d, 0x6f, 0xea, 0x18, 0xd8, 0x9a, 0x9a,
	0x35, 0x7c, 0x78, 0x15, 0xc2, 0x60, 0xef, 0x4b, 0x8c, 0xce, 0x61, 0xd0, 0xe4, 0x7b, 0xf6, 0x8d,
	0xf9, 0xc2, 0xb2, 0xcd, 0x1b, 0xed, 0x00, 0x9d, 0xc1, 0x69, 0xb3, 0xe4, 0x78, 0x86, 0x61, 0x3a,
	0x8e, 0x26, 0xa0, 0x0f, 0x01, 0x35, 0x0b, 0x2f, 0x26, 0xd6, 0xcc, 0xbc, 0xd1, 0xc4, 0xf1, 0x1f,
	0xd0, 0xfb, 0xb1, 0x08, 0xf2, 0x22, 0x70, 0x48, 0xf6, 0x26, 0x0e, 0x09, 0x1a, 0x83, 0xbc, 0xd1,
	0x44, 0x3b, 0x1f, 0xa4, 0x8b, 0x7e, 0x03, 0xa9, 0x62, 0xf8, 0xb5, 0x80, 0xc6, 0xa0, 0x34, 0x9c,
	0xbe, 0x57, 0x57, 0xd0, 0xe1, 0xff, 0xb5, 0xdf, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xab, 0x35,
	0x36, 0x5b, 0x8c, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PubsubServiceClient is the client API for PubsubService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PubsubServiceClient interface {
	Subscribe(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (PubsubService_SubscribeClient, error)
	UnSubscribe(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error)
}

type pubsubServiceClient struct {
	cc *grpc.ClientConn
}

func NewPubsubServiceClient(cc *grpc.ClientConn) PubsubServiceClient {
	return &pubsubServiceClient{cc}
}

func (c *pubsubServiceClient) Subscribe(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (PubsubService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PubsubService_serviceDesc.Streams[0], "/pb.PubsubService/Subscribe", opts...)
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
	Recv() (*EventResponse, error)
	grpc.ClientStream
}

type pubsubServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *pubsubServiceSubscribeClient) Recv() (*EventResponse, error) {
	m := new(EventResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pubsubServiceClient) UnSubscribe(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, "/pb.PubsubService/UnSubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PubsubServiceServer is the server API for PubsubService service.
type PubsubServiceServer interface {
	Subscribe(*EventRequest, PubsubService_SubscribeServer) error
	UnSubscribe(context.Context, *EventRequest) (*EventResponse, error)
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
	Send(*EventResponse) error
	grpc.ServerStream
}

type pubsubServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *pubsubServiceSubscribeServer) Send(m *EventResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _PubsubService_UnSubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubsubServiceServer).UnSubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PubsubService/UnSubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubsubServiceServer).UnSubscribe(ctx, req.(*EventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PubsubService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.PubsubService",
	HandlerType: (*PubsubServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnSubscribe",
			Handler:    _PubsubService_UnSubscribe_Handler,
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
