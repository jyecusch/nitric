// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: nitric/proto/topics/v1/topics.proto

package topicspb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ClientMessage is the message sent from the service to the nitric server
type ClientMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// globally unique ID of the request/response pair
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are assignable to Content:
	//
	//	*ClientMessage_RegistrationRequest
	//	*ClientMessage_MessageResponse
	Content isClientMessage_Content `protobuf_oneof:"content"`
}

func (x *ClientMessage) Reset() {
	*x = ClientMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nitric_proto_topics_v1_topics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientMessage) ProtoMessage() {}

func (x *ClientMessage) ProtoReflect() protoreflect.Message {
	mi := &file_nitric_proto_topics_v1_topics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientMessage.ProtoReflect.Descriptor instead.
func (*ClientMessage) Descriptor() ([]byte, []int) {
	return file_nitric_proto_topics_v1_topics_proto_rawDescGZIP(), []int{0}
}

func (x *ClientMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (m *ClientMessage) GetContent() isClientMessage_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *ClientMessage) GetRegistrationRequest() *RegistrationRequest {
	if x, ok := x.GetContent().(*ClientMessage_RegistrationRequest); ok {
		return x.RegistrationRequest
	}
	return nil
}

func (x *ClientMessage) GetMessageResponse() *MessageResponse {
	if x, ok := x.GetContent().(*ClientMessage_MessageResponse); ok {
		return x.MessageResponse
	}
	return nil
}

type isClientMessage_Content interface {
	isClientMessage_Content()
}

type ClientMessage_RegistrationRequest struct {
	// Register a subscription to a topic
	RegistrationRequest *RegistrationRequest `protobuf:"bytes,2,opt,name=registration_request,json=registrationRequest,proto3,oneof"`
}

type ClientMessage_MessageResponse struct {
	// Handle a message received from a topic
	MessageResponse *MessageResponse `protobuf:"bytes,3,opt,name=message_response,json=messageResponse,proto3,oneof"`
}

func (*ClientMessage_RegistrationRequest) isClientMessage_Content() {}

func (*ClientMessage_MessageResponse) isClientMessage_Content() {}

type MessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopicName string `protobuf:"bytes,1,opt,name=topic_name,json=topicName,proto3" json:"topic_name,omitempty"`
	// Message Type
	Message *TopicMessage `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *MessageRequest) Reset() {
	*x = MessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nitric_proto_topics_v1_topics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageRequest) ProtoMessage() {}

func (x *MessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nitric_proto_topics_v1_topics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageRequest.ProtoReflect.Descriptor instead.
func (*MessageRequest) Descriptor() ([]byte, []int) {
	return file_nitric_proto_topics_v1_topics_proto_rawDescGZIP(), []int{1}
}

func (x *MessageRequest) GetTopicName() string {
	if x != nil {
		return x.TopicName
	}
	return ""
}

func (x *MessageRequest) GetMessage() *TopicMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

type MessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *MessageResponse) Reset() {
	*x = MessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nitric_proto_topics_v1_topics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageResponse) ProtoMessage() {}

func (x *MessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nitric_proto_topics_v1_topics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageResponse.ProtoReflect.Descriptor instead.
func (*MessageResponse) Descriptor() ([]byte, []int) {
	return file_nitric_proto_topics_v1_topics_proto_rawDescGZIP(), []int{2}
}

func (x *MessageResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// ServerMessage is the message sent from the nitric server to the service
type ServerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// globally unique ID of the request/response pair
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are assignable to Content:
	//
	//	*ServerMessage_RegistrationResponse
	//	*ServerMessage_MessageRequest
	Content isServerMessage_Content `protobuf_oneof:"content"`
}

func (x *ServerMessage) Reset() {
	*x = ServerMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nitric_proto_topics_v1_topics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerMessage) ProtoMessage() {}

func (x *ServerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_nitric_proto_topics_v1_topics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerMessage.ProtoReflect.Descriptor instead.
func (*ServerMessage) Descriptor() ([]byte, []int) {
	return file_nitric_proto_topics_v1_topics_proto_rawDescGZIP(), []int{3}
}

func (x *ServerMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (m *ServerMessage) GetContent() isServerMessage_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *ServerMessage) GetRegistrationResponse() *RegistrationResponse {
	if x, ok := x.GetContent().(*ServerMessage_RegistrationResponse); ok {
		return x.RegistrationResponse
	}
	return nil
}

func (x *ServerMessage) GetMessageRequest() *MessageRequest {
	if x, ok := x.GetContent().(*ServerMessage_MessageRequest); ok {
		return x.MessageRequest
	}
	return nil
}

type isServerMessage_Content interface {
	isServerMessage_Content()
}

type ServerMessage_RegistrationResponse struct {
	// Response to a topic subscription request
	RegistrationResponse *RegistrationResponse `protobuf:"bytes,2,opt,name=registration_response,json=registrationResponse,proto3,oneof"`
}

type ServerMessage_MessageRequest struct {
	// Response to a topic message request
	MessageRequest *MessageRequest `protobuf:"bytes,3,opt,name=message_request,json=messageRequest,proto3,oneof"`
}

func (*ServerMessage_RegistrationResponse) isServerMessage_Content() {}

func (*ServerMessage_MessageRequest) isServerMessage_Content() {}

type RegistrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopicName string `protobuf:"bytes,1,opt,name=topic_name,json=topicName,proto3" json:"topic_name,omitempty"`
}

func (x *RegistrationRequest) Reset() {
	*x = RegistrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nitric_proto_topics_v1_topics_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationRequest) ProtoMessage() {}

func (x *RegistrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nitric_proto_topics_v1_topics_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrationRequest.ProtoReflect.Descriptor instead.
func (*RegistrationRequest) Descriptor() ([]byte, []int) {
	return file_nitric_proto_topics_v1_topics_proto_rawDescGZIP(), []int{4}
}

func (x *RegistrationRequest) GetTopicName() string {
	if x != nil {
		return x.TopicName
	}
	return ""
}

type RegistrationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegistrationResponse) Reset() {
	*x = RegistrationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nitric_proto_topics_v1_topics_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistrationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationResponse) ProtoMessage() {}

func (x *RegistrationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nitric_proto_topics_v1_topics_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrationResponse.ProtoReflect.Descriptor instead.
func (*RegistrationResponse) Descriptor() ([]byte, []int) {
	return file_nitric_proto_topics_v1_topics_proto_rawDescGZIP(), []int{5}
}

type TopicMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The topic message contents
	//
	// Types that are assignable to Content:
	//
	//	*TopicMessage_StructPayload
	Content isTopicMessage_Content `protobuf_oneof:"content"`
}

func (x *TopicMessage) Reset() {
	*x = TopicMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nitric_proto_topics_v1_topics_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopicMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopicMessage) ProtoMessage() {}

func (x *TopicMessage) ProtoReflect() protoreflect.Message {
	mi := &file_nitric_proto_topics_v1_topics_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopicMessage.ProtoReflect.Descriptor instead.
func (*TopicMessage) Descriptor() ([]byte, []int) {
	return file_nitric_proto_topics_v1_topics_proto_rawDescGZIP(), []int{6}
}

func (m *TopicMessage) GetContent() isTopicMessage_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *TopicMessage) GetStructPayload() *structpb.Struct {
	if x, ok := x.GetContent().(*TopicMessage_StructPayload); ok {
		return x.StructPayload
	}
	return nil
}

type isTopicMessage_Content interface {
	isTopicMessage_Content()
}

type TopicMessage_StructPayload struct {
	StructPayload *structpb.Struct `protobuf:"bytes,1,opt,name=struct_payload,json=structPayload,proto3,oneof"`
}

func (*TopicMessage_StructPayload) isTopicMessage_Content() {}

// Request to publish a message to a topic
type TopicPublishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the topic to publish the topic to
	TopicName string `protobuf:"bytes,1,opt,name=topic_name,json=topicName,proto3" json:"topic_name,omitempty"`
	// The message to be published
	Message *TopicMessage `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// An optional delay specified in seconds (minimum 10 seconds)
	Delay *durationpb.Duration `protobuf:"bytes,3,opt,name=delay,proto3" json:"delay,omitempty"`
}

func (x *TopicPublishRequest) Reset() {
	*x = TopicPublishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nitric_proto_topics_v1_topics_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopicPublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopicPublishRequest) ProtoMessage() {}

func (x *TopicPublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nitric_proto_topics_v1_topics_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopicPublishRequest.ProtoReflect.Descriptor instead.
func (*TopicPublishRequest) Descriptor() ([]byte, []int) {
	return file_nitric_proto_topics_v1_topics_proto_rawDescGZIP(), []int{7}
}

func (x *TopicPublishRequest) GetTopicName() string {
	if x != nil {
		return x.TopicName
	}
	return ""
}

func (x *TopicPublishRequest) GetMessage() *TopicMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *TopicPublishRequest) GetDelay() *durationpb.Duration {
	if x != nil {
		return x.Delay
	}
	return nil
}

// Result of publishing an topic
type TopicPublishResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TopicPublishResponse) Reset() {
	*x = TopicPublishResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nitric_proto_topics_v1_topics_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopicPublishResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopicPublishResponse) ProtoMessage() {}

func (x *TopicPublishResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nitric_proto_topics_v1_topics_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopicPublishResponse.ProtoReflect.Descriptor instead.
func (*TopicPublishResponse) Descriptor() ([]byte, []int) {
	return file_nitric_proto_topics_v1_topics_proto_rawDescGZIP(), []int{8}
}

var File_nitric_proto_topics_v1_topics_proto protoreflect.FileDescriptor

var file_nitric_proto_topics_v1_topics_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe2, 0x01, 0x0a, 0x0d,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x60, 0x0a,
	0x14, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6e, 0x69,
	0x74, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x13, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x54, 0x0a, 0x10, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6e, 0x69, 0x74, 0x72,
	0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x22, 0x6f, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x3e, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x70, 0x69,
	0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x2b, 0x0a, 0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0xe2,
	0x01, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x63, 0x0a, 0x15, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52,
	0x14, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x22, 0x34, 0x0a, 0x13, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x5b, 0x0a, 0x0c, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x40, 0x0a, 0x0e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0xa5,
	0x01, 0x0a, 0x13, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x6f, 0x70, 0x69, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x05, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x22, 0x16, 0x0a, 0x14, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x6e,
	0x0a, 0x06, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x12, 0x64, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x12, 0x2b, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2c, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x6b,
	0x0a, 0x0a, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x12, 0x5d, 0x0a, 0x09,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x25, 0x2e, 0x6e, 0x69, 0x74, 0x72,
	0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x1a, 0x25, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x9e, 0x01, 0x0a, 0x19,
	0x69, 0x6f, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x74, 0x65, 0x63,
	0x68, 0x2f, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x2f, 0x76,
	0x31, 0x3b, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x70, 0x62, 0xaa, 0x02, 0x16, 0x4e, 0x69, 0x74,
	0x72, 0x69, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73,
	0x2e, 0x76, 0x31, 0xca, 0x02, 0x16, 0x4e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x5c, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x5c, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nitric_proto_topics_v1_topics_proto_rawDescOnce sync.Once
	file_nitric_proto_topics_v1_topics_proto_rawDescData = file_nitric_proto_topics_v1_topics_proto_rawDesc
)

func file_nitric_proto_topics_v1_topics_proto_rawDescGZIP() []byte {
	file_nitric_proto_topics_v1_topics_proto_rawDescOnce.Do(func() {
		file_nitric_proto_topics_v1_topics_proto_rawDescData = protoimpl.X.CompressGZIP(file_nitric_proto_topics_v1_topics_proto_rawDescData)
	})
	return file_nitric_proto_topics_v1_topics_proto_rawDescData
}

var file_nitric_proto_topics_v1_topics_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_nitric_proto_topics_v1_topics_proto_goTypes = []interface{}{
	(*ClientMessage)(nil),        // 0: nitric.proto.topics.v1.ClientMessage
	(*MessageRequest)(nil),       // 1: nitric.proto.topics.v1.MessageRequest
	(*MessageResponse)(nil),      // 2: nitric.proto.topics.v1.MessageResponse
	(*ServerMessage)(nil),        // 3: nitric.proto.topics.v1.ServerMessage
	(*RegistrationRequest)(nil),  // 4: nitric.proto.topics.v1.RegistrationRequest
	(*RegistrationResponse)(nil), // 5: nitric.proto.topics.v1.RegistrationResponse
	(*TopicMessage)(nil),         // 6: nitric.proto.topics.v1.TopicMessage
	(*TopicPublishRequest)(nil),  // 7: nitric.proto.topics.v1.TopicPublishRequest
	(*TopicPublishResponse)(nil), // 8: nitric.proto.topics.v1.TopicPublishResponse
	(*structpb.Struct)(nil),      // 9: google.protobuf.Struct
	(*durationpb.Duration)(nil),  // 10: google.protobuf.Duration
}
var file_nitric_proto_topics_v1_topics_proto_depIdxs = []int32{
	4,  // 0: nitric.proto.topics.v1.ClientMessage.registration_request:type_name -> nitric.proto.topics.v1.RegistrationRequest
	2,  // 1: nitric.proto.topics.v1.ClientMessage.message_response:type_name -> nitric.proto.topics.v1.MessageResponse
	6,  // 2: nitric.proto.topics.v1.MessageRequest.message:type_name -> nitric.proto.topics.v1.TopicMessage
	5,  // 3: nitric.proto.topics.v1.ServerMessage.registration_response:type_name -> nitric.proto.topics.v1.RegistrationResponse
	1,  // 4: nitric.proto.topics.v1.ServerMessage.message_request:type_name -> nitric.proto.topics.v1.MessageRequest
	9,  // 5: nitric.proto.topics.v1.TopicMessage.struct_payload:type_name -> google.protobuf.Struct
	6,  // 6: nitric.proto.topics.v1.TopicPublishRequest.message:type_name -> nitric.proto.topics.v1.TopicMessage
	10, // 7: nitric.proto.topics.v1.TopicPublishRequest.delay:type_name -> google.protobuf.Duration
	7,  // 8: nitric.proto.topics.v1.Topics.Publish:input_type -> nitric.proto.topics.v1.TopicPublishRequest
	0,  // 9: nitric.proto.topics.v1.Subscriber.Subscribe:input_type -> nitric.proto.topics.v1.ClientMessage
	8,  // 10: nitric.proto.topics.v1.Topics.Publish:output_type -> nitric.proto.topics.v1.TopicPublishResponse
	3,  // 11: nitric.proto.topics.v1.Subscriber.Subscribe:output_type -> nitric.proto.topics.v1.ServerMessage
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_nitric_proto_topics_v1_topics_proto_init() }
func file_nitric_proto_topics_v1_topics_proto_init() {
	if File_nitric_proto_topics_v1_topics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nitric_proto_topics_v1_topics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientMessage); i {
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
		file_nitric_proto_topics_v1_topics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageRequest); i {
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
		file_nitric_proto_topics_v1_topics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageResponse); i {
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
		file_nitric_proto_topics_v1_topics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerMessage); i {
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
		file_nitric_proto_topics_v1_topics_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistrationRequest); i {
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
		file_nitric_proto_topics_v1_topics_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistrationResponse); i {
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
		file_nitric_proto_topics_v1_topics_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopicMessage); i {
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
		file_nitric_proto_topics_v1_topics_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopicPublishRequest); i {
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
		file_nitric_proto_topics_v1_topics_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopicPublishResponse); i {
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
	file_nitric_proto_topics_v1_topics_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ClientMessage_RegistrationRequest)(nil),
		(*ClientMessage_MessageResponse)(nil),
	}
	file_nitric_proto_topics_v1_topics_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*ServerMessage_RegistrationResponse)(nil),
		(*ServerMessage_MessageRequest)(nil),
	}
	file_nitric_proto_topics_v1_topics_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*TopicMessage_StructPayload)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nitric_proto_topics_v1_topics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_nitric_proto_topics_v1_topics_proto_goTypes,
		DependencyIndexes: file_nitric_proto_topics_v1_topics_proto_depIdxs,
		MessageInfos:      file_nitric_proto_topics_v1_topics_proto_msgTypes,
	}.Build()
	File_nitric_proto_topics_v1_topics_proto = out.File
	file_nitric_proto_topics_v1_topics_proto_rawDesc = nil
	file_nitric_proto_topics_v1_topics_proto_goTypes = nil
	file_nitric_proto_topics_v1_topics_proto_depIdxs = nil
}
