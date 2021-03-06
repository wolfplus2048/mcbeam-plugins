// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.3
// source: session/proto/gate.proto

package gate

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type MsgType int32

const (
	MsgType_MsgRequest  MsgType = 0
	MsgType_MsgNotify   MsgType = 1
	MsgType_MsgResponse MsgType = 2
	MsgType_MsgPush     MsgType = 3
)

// Enum value maps for MsgType.
var (
	MsgType_name = map[int32]string{
		0: "MsgRequest",
		1: "MsgNotify",
		2: "MsgResponse",
		3: "MsgPush",
	}
	MsgType_value = map[string]int32{
		"MsgRequest":  0,
		"MsgNotify":   1,
		"MsgResponse": 2,
		"MsgPush":     3,
	}
)

func (x MsgType) Enum() *MsgType {
	p := new(MsgType)
	*p = x
	return p
}

func (x MsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_session_proto_gate_proto_enumTypes[0].Descriptor()
}

func (MsgType) Type() protoreflect.EnumType {
	return &file_session_proto_gate_proto_enumTypes[0]
}

func (x MsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgType.Descriptor instead.
func (MsgType) EnumDescriptor() ([]byte, []int) {
	return file_session_proto_gate_proto_rawDescGZIP(), []int{0}
}

type RPCType int32

const (
	RPCType_Sys  RPCType = 0
	RPCType_User RPCType = 1
)

// Enum value maps for RPCType.
var (
	RPCType_name = map[int32]string{
		0: "Sys",
		1: "User",
	}
	RPCType_value = map[string]int32{
		"Sys":  0,
		"User": 1,
	}
)

func (x RPCType) Enum() *RPCType {
	p := new(RPCType)
	*p = x
	return p
}

func (x RPCType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RPCType) Descriptor() protoreflect.EnumDescriptor {
	return file_session_proto_gate_proto_enumTypes[1].Descriptor()
}

func (RPCType) Type() protoreflect.EnumType {
	return &file_session_proto_gate_proto_enumTypes[1]
}

func (x RPCType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RPCType.Descriptor instead.
func (RPCType) EnumDescriptor() ([]byte, []int) {
	return file_session_proto_gate_proto_rawDescGZIP(), []int{1}
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code   int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Detail string `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"`
	Status string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_gate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_gate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_session_proto_gate_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Error) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

func (x *Error) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid  string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_gate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_gate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_session_proto_gate_proto_rawDescGZIP(), []int{1}
}

func (x *Session) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Session) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Session) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type SessionClose struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *SessionClose) Reset() {
	*x = SessionClose{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_gate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionClose) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionClose) ProtoMessage() {}

func (x *SessionClose) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_gate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionClose.ProtoReflect.Descriptor instead.
func (*SessionClose) Descriptor() ([]byte, []int) {
	return file_session_proto_gate_proto_rawDescGZIP(), []int{2}
}

func (x *SessionClose) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Route string  `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
	Data  []byte  `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Reply string  `protobuf:"bytes,4,opt,name=reply,proto3" json:"reply,omitempty"`
	Type  MsgType `protobuf:"varint,5,opt,name=type,proto3,enum=gate.MsgType" json:"type,omitempty"`
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_gate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_gate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_session_proto_gate_proto_rawDescGZIP(), []int{3}
}

func (x *Msg) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Msg) GetRoute() string {
	if x != nil {
		return x.Route
	}
	return ""
}

func (x *Msg) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Msg) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

func (x *Msg) GetType() MsgType {
	if x != nil {
		return x.Type
	}
	return MsgType_MsgRequest
}

type KickMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *KickMsg) Reset() {
	*x = KickMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_gate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KickMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KickMsg) ProtoMessage() {}

func (x *KickMsg) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_gate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KickMsg.ProtoReflect.Descriptor instead.
func (*KickMsg) Descriptor() ([]byte, []int) {
	return file_session_proto_gate_proto_rawDescGZIP(), []int{4}
}

func (x *KickMsg) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type KickAnswer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kicked bool `protobuf:"varint,1,opt,name=kicked,proto3" json:"kicked,omitempty"`
}

func (x *KickAnswer) Reset() {
	*x = KickAnswer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_gate_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KickAnswer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KickAnswer) ProtoMessage() {}

func (x *KickAnswer) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_gate_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KickAnswer.ProtoReflect.Descriptor instead.
func (*KickAnswer) Descriptor() ([]byte, []int) {
	return file_session_proto_gate_proto_rawDescGZIP(), []int{5}
}

func (x *KickAnswer) GetKicked() bool {
	if x != nil {
		return x.Kicked
	}
	return false
}

type PushMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Route string `protobuf:"bytes,1,opt,name=route,proto3" json:"route,omitempty"`
	Uid   string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Data  []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PushMsg) Reset() {
	*x = PushMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_gate_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMsg) ProtoMessage() {}

func (x *PushMsg) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_gate_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMsg.ProtoReflect.Descriptor instead.
func (*PushMsg) Descriptor() ([]byte, []int) {
	return file_session_proto_gate_proto_rawDescGZIP(), []int{6}
}

func (x *PushMsg) GetRoute() string {
	if x != nil {
		return x.Route
	}
	return ""
}

func (x *PushMsg) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *PushMsg) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       RPCType  `protobuf:"varint,1,opt,name=type,proto3,enum=gate.RPCType" json:"type,omitempty"`
	Session    *Session `protobuf:"bytes,2,opt,name=session,proto3" json:"session,omitempty"`
	Msg        *Msg     `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	FrontendID string   `protobuf:"bytes,4,opt,name=frontendID,proto3" json:"frontendID,omitempty"`
	Metadata   []byte   `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_gate_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_gate_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_session_proto_gate_proto_rawDescGZIP(), []int{7}
}

func (x *Request) GetType() RPCType {
	if x != nil {
		return x.Type
	}
	return RPCType_Sys
}

func (x *Request) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

func (x *Request) GetMsg() *Msg {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *Request) GetFrontendID() string {
	if x != nil {
		return x.FrontendID
	}
	return ""
}

func (x *Request) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Error *Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_gate_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_gate_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_session_proto_gate_proto_rawDescGZIP(), []int{8}
}

func (x *Response) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Response) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_session_proto_gate_proto protoreflect.FileDescriptor

var file_session_proto_gate_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x61, 0x74, 0x65,
	0x22, 0x5b, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3f, 0x0a,
	0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x20,
	0x0a, 0x0c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64,
	0x22, 0x78, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x21, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x4d, 0x73, 0x67,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x21, 0x0a, 0x07, 0x4b, 0x69,
	0x63, 0x6b, 0x4d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x24, 0x0a,
	0x0a, 0x4b, 0x69, 0x63, 0x6b, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6b,
	0x69, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6b, 0x69, 0x63,
	0x6b, 0x65, 0x64, 0x22, 0x45, 0x0a, 0x07, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x12, 0x14,
	0x0a, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xae, 0x01, 0x0a, 0x07, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x50, 0x43, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12,
	0x1e, 0x0a, 0x0a, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x49, 0x44, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x49, 0x44, 0x12,
	0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x41, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2a, 0x46,
	0x0a, 0x07, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x73, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x73, 0x67,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x73, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x73, 0x67,
	0x50, 0x75, 0x73, 0x68, 0x10, 0x03, 0x2a, 0x1c, 0x0a, 0x07, 0x52, 0x50, 0x43, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x79, 0x73, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x10, 0x01, 0x32, 0x31, 0x0a, 0x06, 0x4d, 0x63, 0x62, 0x41, 0x70, 0x70, 0x12, 0x27,
	0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x0d, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0xb6, 0x01, 0x0a, 0x07, 0x4d, 0x63, 0x62, 0x47,
	0x61, 0x74, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x50, 0x75, 0x73, 0x68, 0x12, 0x0d, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x1a, 0x0e, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x0b,
	0x50, 0x75, 0x73, 0x68, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0d, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x0e, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x04,
	0x42, 0x69, 0x6e, 0x64, 0x12, 0x0d, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x1a, 0x0e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x04, 0x4b, 0x69, 0x63, 0x6b, 0x12, 0x0d, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x2e, 0x4b, 0x69, 0x63, 0x6b, 0x4d, 0x73, 0x67, 0x1a, 0x10, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x2e, 0x4b, 0x69, 0x63, 0x6b, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x00,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_session_proto_gate_proto_rawDescOnce sync.Once
	file_session_proto_gate_proto_rawDescData = file_session_proto_gate_proto_rawDesc
)

func file_session_proto_gate_proto_rawDescGZIP() []byte {
	file_session_proto_gate_proto_rawDescOnce.Do(func() {
		file_session_proto_gate_proto_rawDescData = protoimpl.X.CompressGZIP(file_session_proto_gate_proto_rawDescData)
	})
	return file_session_proto_gate_proto_rawDescData
}

var file_session_proto_gate_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_session_proto_gate_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_session_proto_gate_proto_goTypes = []interface{}{
	(MsgType)(0),         // 0: gate.MsgType
	(RPCType)(0),         // 1: gate.RPCType
	(*Error)(nil),        // 2: gate.Error
	(*Session)(nil),      // 3: gate.Session
	(*SessionClose)(nil), // 4: gate.SessionClose
	(*Msg)(nil),          // 5: gate.Msg
	(*KickMsg)(nil),      // 6: gate.KickMsg
	(*KickAnswer)(nil),   // 7: gate.KickAnswer
	(*PushMsg)(nil),      // 8: gate.PushMsg
	(*Request)(nil),      // 9: gate.Request
	(*Response)(nil),     // 10: gate.Response
}
var file_session_proto_gate_proto_depIdxs = []int32{
	0,  // 0: gate.Msg.type:type_name -> gate.MsgType
	1,  // 1: gate.Request.type:type_name -> gate.RPCType
	3,  // 2: gate.Request.session:type_name -> gate.Session
	5,  // 3: gate.Request.msg:type_name -> gate.Msg
	2,  // 4: gate.Response.error:type_name -> gate.Error
	9,  // 5: gate.McbApp.Call:input_type -> gate.Request
	8,  // 6: gate.McbGate.Push:input_type -> gate.PushMsg
	3,  // 7: gate.McbGate.PushSession:input_type -> gate.Session
	3,  // 8: gate.McbGate.Bind:input_type -> gate.Session
	6,  // 9: gate.McbGate.Kick:input_type -> gate.KickMsg
	10, // 10: gate.McbApp.Call:output_type -> gate.Response
	10, // 11: gate.McbGate.Push:output_type -> gate.Response
	10, // 12: gate.McbGate.PushSession:output_type -> gate.Response
	10, // 13: gate.McbGate.Bind:output_type -> gate.Response
	7,  // 14: gate.McbGate.Kick:output_type -> gate.KickAnswer
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_session_proto_gate_proto_init() }
func file_session_proto_gate_proto_init() {
	if File_session_proto_gate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_session_proto_gate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_session_proto_gate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_session_proto_gate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionClose); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_session_proto_gate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_session_proto_gate_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KickMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_session_proto_gate_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KickAnswer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_session_proto_gate_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_session_proto_gate_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_session_proto_gate_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_session_proto_gate_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_session_proto_gate_proto_goTypes,
		DependencyIndexes: file_session_proto_gate_proto_depIdxs,
		EnumInfos:         file_session_proto_gate_proto_enumTypes,
		MessageInfos:      file_session_proto_gate_proto_msgTypes,
	}.Build()
	File_session_proto_gate_proto = out.File
	file_session_proto_gate_proto_rawDesc = nil
	file_session_proto_gate_proto_goTypes = nil
	file_session_proto_gate_proto_depIdxs = nil
}
