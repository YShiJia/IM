// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.4
// source: apps/status/rpc/status.proto

package statusmodel

import (
	pbmessage "github.com/YShiJia/IM/pbmodel/pbmessage"
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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ping string `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_status_rpc_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_apps_status_rpc_status_proto_msgTypes[0]
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
	return file_apps_status_rpc_status_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetPing() string {
	if x != nil {
		return x.Ping
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_status_rpc_status_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_apps_status_rpc_status_proto_msgTypes[1]
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
	return file_apps_status_rpc_status_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

type UserOnlineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SocialId []string `protobuf:"bytes,1,rep,name=socialId,proto3" json:"socialId,omitempty"`
}

func (x *UserOnlineRequest) Reset() {
	*x = UserOnlineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_status_rpc_status_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserOnlineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserOnlineRequest) ProtoMessage() {}

func (x *UserOnlineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_status_rpc_status_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserOnlineRequest.ProtoReflect.Descriptor instead.
func (*UserOnlineRequest) Descriptor() ([]byte, []int) {
	return file_apps_status_rpc_status_proto_rawDescGZIP(), []int{2}
}

func (x *UserOnlineRequest) GetSocialId() []string {
	if x != nil {
		return x.SocialId
	}
	return nil
}

type UserOnlineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SocialId []string `protobuf:"bytes,1,rep,name=socialId,proto3" json:"socialId,omitempty"`
}

func (x *UserOnlineResponse) Reset() {
	*x = UserOnlineResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_status_rpc_status_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserOnlineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserOnlineResponse) ProtoMessage() {}

func (x *UserOnlineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apps_status_rpc_status_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserOnlineResponse.ProtoReflect.Descriptor instead.
func (*UserOnlineResponse) Descriptor() ([]byte, []int) {
	return file_apps_status_rpc_status_proto_rawDescGZIP(), []int{3}
}

func (x *UserOnlineResponse) GetSocialId() []string {
	if x != nil {
		return x.SocialId
	}
	return nil
}

type ClientConnAddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SocialId string `protobuf:"bytes,1,opt,name=socialId,proto3" json:"socialId,omitempty"`
}

func (x *ClientConnAddressRequest) Reset() {
	*x = ClientConnAddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_status_rpc_status_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientConnAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientConnAddressRequest) ProtoMessage() {}

func (x *ClientConnAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_status_rpc_status_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientConnAddressRequest.ProtoReflect.Descriptor instead.
func (*ClientConnAddressRequest) Descriptor() ([]byte, []int) {
	return file_apps_status_rpc_status_proto_rawDescGZIP(), []int{4}
}

func (x *ClientConnAddressRequest) GetSocialId() string {
	if x != nil {
		return x.SocialId
	}
	return ""
}

type ClientConnAddressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *ClientConnAddressResponse) Reset() {
	*x = ClientConnAddressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_status_rpc_status_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientConnAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientConnAddressResponse) ProtoMessage() {}

func (x *ClientConnAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apps_status_rpc_status_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientConnAddressResponse.ProtoReflect.Descriptor instead.
func (*ClientConnAddressResponse) Descriptor() ([]byte, []int) {
	return file_apps_status_rpc_status_proto_rawDescGZIP(), []int{5}
}

func (x *ClientConnAddressResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type ClientMsgSyncRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SocialId string `protobuf:"bytes,1,opt,name=socialId,proto3" json:"socialId,omitempty"`
	// 区分是私聊消息还是群聊消息
	MsgType int32  `protobuf:"varint,2,opt,name=msgType,proto3" json:"msgType,omitempty"`
	FromId  string `protobuf:"bytes,3,opt,name=fromId,proto3" json:"fromId,omitempty"`
	// 消息起始seq号
	Begin int64 `protobuf:"varint,4,opt,name=begin,proto3" json:"begin,omitempty"`
	// 消息结束seq号
	End int64 `protobuf:"varint,5,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *ClientMsgSyncRequest) Reset() {
	*x = ClientMsgSyncRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_status_rpc_status_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientMsgSyncRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientMsgSyncRequest) ProtoMessage() {}

func (x *ClientMsgSyncRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_status_rpc_status_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientMsgSyncRequest.ProtoReflect.Descriptor instead.
func (*ClientMsgSyncRequest) Descriptor() ([]byte, []int) {
	return file_apps_status_rpc_status_proto_rawDescGZIP(), []int{6}
}

func (x *ClientMsgSyncRequest) GetSocialId() string {
	if x != nil {
		return x.SocialId
	}
	return ""
}

func (x *ClientMsgSyncRequest) GetMsgType() int32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

func (x *ClientMsgSyncRequest) GetFromId() string {
	if x != nil {
		return x.FromId
	}
	return ""
}

func (x *ClientMsgSyncRequest) GetBegin() int64 {
	if x != nil {
		return x.Begin
	}
	return 0
}

func (x *ClientMsgSyncRequest) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

type ClientMsgSyncResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*pbmessage.PbMessage `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *ClientMsgSyncResponse) Reset() {
	*x = ClientMsgSyncResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_status_rpc_status_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientMsgSyncResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientMsgSyncResponse) ProtoMessage() {}

func (x *ClientMsgSyncResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apps_status_rpc_status_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientMsgSyncResponse.ProtoReflect.Descriptor instead.
func (*ClientMsgSyncResponse) Descriptor() ([]byte, []int) {
	return file_apps_status_rpc_status_proto_rawDescGZIP(), []int{7}
}

func (x *ClientMsgSyncResponse) GetMessages() []*pbmessage.PbMessage {
	if x != nil {
		return x.Messages
	}
	return nil
}

var File_apps_status_rpc_status_proto protoreflect.FileDescriptor

var file_apps_status_rpc_status_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x72, 0x70,
	0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x15, 0x70, 0x62, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a,
	0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x22, 0x1e, 0x0a, 0x08,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x22, 0x2f, 0x0a, 0x11,
	0x55, 0x73, 0x65, 0x72, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x22, 0x30, 0x0a,
	0x12, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x22,
	0x36, 0x0a, 0x18, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x6f, 0x63, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x6f, 0x63, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x19, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x8c,
	0x01, 0x0a, 0x14, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x53, 0x79, 0x6e, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6f, 0x63, 0x69, 0x61,
	0x6c, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x63, 0x69, 0x61,
	0x6c, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x66, 0x72, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66,
	0x72, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x65,
	0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x45, 0x0a,
	0x15, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x50, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x32, 0xa0, 0x02, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x29, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0f, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x55, 0x73,
	0x65, 0x72, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x19, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x58, 0x0a, 0x11, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0d, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x1c, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x53, 0x79, 0x6e,
	0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x53, 0x79, 0x6e, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_status_rpc_status_proto_rawDescOnce sync.Once
	file_apps_status_rpc_status_proto_rawDescData = file_apps_status_rpc_status_proto_rawDesc
)

func file_apps_status_rpc_status_proto_rawDescGZIP() []byte {
	file_apps_status_rpc_status_proto_rawDescOnce.Do(func() {
		file_apps_status_rpc_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_status_rpc_status_proto_rawDescData)
	})
	return file_apps_status_rpc_status_proto_rawDescData
}

var file_apps_status_rpc_status_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_apps_status_rpc_status_proto_goTypes = []any{
	(*Request)(nil),                   // 0: status.Request
	(*Response)(nil),                  // 1: status.Response
	(*UserOnlineRequest)(nil),         // 2: status.UserOnlineRequest
	(*UserOnlineResponse)(nil),        // 3: status.UserOnlineResponse
	(*ClientConnAddressRequest)(nil),  // 4: status.ClientConnAddressRequest
	(*ClientConnAddressResponse)(nil), // 5: status.ClientConnAddressResponse
	(*ClientMsgSyncRequest)(nil),      // 6: status.ClientMsgSyncRequest
	(*ClientMsgSyncResponse)(nil),     // 7: status.ClientMsgSyncResponse
	(*pbmessage.PbMessage)(nil),       // 8: model.PbMessage
}
var file_apps_status_rpc_status_proto_depIdxs = []int32{
	8, // 0: status.ClientMsgSyncResponse.messages:type_name -> model.PbMessage
	0, // 1: status.Status.Ping:input_type -> status.Request
	2, // 2: status.Status.UserOnline:input_type -> status.UserOnlineRequest
	4, // 3: status.Status.ClientConnAddress:input_type -> status.ClientConnAddressRequest
	6, // 4: status.Status.ClientMsgSync:input_type -> status.ClientMsgSyncRequest
	1, // 5: status.Status.Ping:output_type -> status.Response
	3, // 6: status.Status.UserOnline:output_type -> status.UserOnlineResponse
	5, // 7: status.Status.ClientConnAddress:output_type -> status.ClientConnAddressResponse
	7, // 8: status.Status.ClientMsgSync:output_type -> status.ClientMsgSyncResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_apps_status_rpc_status_proto_init() }
func file_apps_status_rpc_status_proto_init() {
	if File_apps_status_rpc_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_status_rpc_status_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_apps_status_rpc_status_proto_msgTypes[1].Exporter = func(v any, i int) any {
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
		file_apps_status_rpc_status_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UserOnlineRequest); i {
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
		file_apps_status_rpc_status_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UserOnlineResponse); i {
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
		file_apps_status_rpc_status_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ClientConnAddressRequest); i {
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
		file_apps_status_rpc_status_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ClientConnAddressResponse); i {
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
		file_apps_status_rpc_status_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ClientMsgSyncRequest); i {
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
		file_apps_status_rpc_status_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*ClientMsgSyncResponse); i {
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
			RawDescriptor: file_apps_status_rpc_status_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_status_rpc_status_proto_goTypes,
		DependencyIndexes: file_apps_status_rpc_status_proto_depIdxs,
		MessageInfos:      file_apps_status_rpc_status_proto_msgTypes,
	}.Build()
	File_apps_status_rpc_status_proto = out.File
	file_apps_status_rpc_status_proto_rawDesc = nil
	file_apps_status_rpc_status_proto_goTypes = nil
	file_apps_status_rpc_status_proto_depIdxs = nil
}
