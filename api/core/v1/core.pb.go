// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.7
// source: core/v1/core.proto

package corepb

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriberId string    `protobuf:"bytes,1,opt,name=subscriber_id,json=subscriberId,proto3" json:"subscriber_id,omitempty"`
	Name         string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Listener     *Listener `protobuf:"bytes,3,opt,name=listener,proto3" json:"listener,omitempty"`
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_core_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_core_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{0}
}

func (x *SubscribeRequest) GetSubscriberId() string {
	if x != nil {
		return x.SubscriberId
	}
	return ""
}

func (x *SubscribeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SubscribeRequest) GetListener() *Listener {
	if x != nil {
		return x.Listener
	}
	return nil
}

type SubscribeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriptionId string `protobuf:"bytes,1,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
}

func (x *SubscribeResponse) Reset() {
	*x = SubscribeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_core_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeResponse) ProtoMessage() {}

func (x *SubscribeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_core_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeResponse.ProtoReflect.Descriptor instead.
func (*SubscribeResponse) Descriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{1}
}

func (x *SubscribeResponse) GetSubscriptionId() string {
	if x != nil {
		return x.SubscriptionId
	}
	return ""
}

type Listener struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port int32 `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	// Types that are assignable to Match:
	//	*Listener_Triggers
	//	*Listener_TriggerRegex
	Match isListener_Match `protobuf_oneof:"match"`
}

func (x *Listener) Reset() {
	*x = Listener{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_core_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Listener) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Listener) ProtoMessage() {}

func (x *Listener) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_core_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Listener.ProtoReflect.Descriptor instead.
func (*Listener) Descriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{2}
}

func (x *Listener) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (m *Listener) GetMatch() isListener_Match {
	if m != nil {
		return m.Match
	}
	return nil
}

func (x *Listener) GetTriggers() *Triggers {
	if x, ok := x.GetMatch().(*Listener_Triggers); ok {
		return x.Triggers
	}
	return nil
}

func (x *Listener) GetTriggerRegex() string {
	if x, ok := x.GetMatch().(*Listener_TriggerRegex); ok {
		return x.TriggerRegex
	}
	return ""
}

type isListener_Match interface {
	isListener_Match()
}

type Listener_Triggers struct {
	Triggers *Triggers `protobuf:"bytes,2,opt,name=triggers,proto3,oneof"`
}

type Listener_TriggerRegex struct {
	TriggerRegex string `protobuf:"bytes,3,opt,name=trigger_regex,json=triggerRegex,proto3,oneof"`
}

func (*Listener_Triggers) isListener_Match() {}

func (*Listener_TriggerRegex) isListener_Match() {}

type Triggers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Triggers []string `protobuf:"bytes,1,rep,name=triggers,proto3" json:"triggers,omitempty"`
}

func (x *Triggers) Reset() {
	*x = Triggers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_core_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Triggers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Triggers) ProtoMessage() {}

func (x *Triggers) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_core_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Triggers.ProtoReflect.Descriptor instead.
func (*Triggers) Descriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{3}
}

func (x *Triggers) GetTriggers() []string {
	if x != nil {
		return x.Triggers
	}
	return nil
}

type EventMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to EventType:
	//	*EventMessage_Event
	//	*EventMessage_Response
	EventType isEventMessage_EventType `protobuf_oneof:"event_type"`
}

func (x *EventMessage) Reset() {
	*x = EventMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_core_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventMessage) ProtoMessage() {}

func (x *EventMessage) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_core_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventMessage.ProtoReflect.Descriptor instead.
func (*EventMessage) Descriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{4}
}

func (m *EventMessage) GetEventType() isEventMessage_EventType {
	if m != nil {
		return m.EventType
	}
	return nil
}

func (x *EventMessage) GetEvent() *Event {
	if x, ok := x.GetEventType().(*EventMessage_Event); ok {
		return x.Event
	}
	return nil
}

func (x *EventMessage) GetResponse() *EventResponse {
	if x, ok := x.GetEventType().(*EventMessage_Response); ok {
		return x.Response
	}
	return nil
}

type isEventMessage_EventType interface {
	isEventMessage_EventType()
}

type EventMessage_Event struct {
	Event *Event `protobuf:"bytes,1,opt,name=event,proto3,oneof"`
}

type EventMessage_Response struct {
	Response *EventResponse `protobuf:"bytes,2,opt,name=response,proto3,oneof"`
}

func (*EventMessage_Event) isEventMessage_EventType() {}

func (*EventMessage_Response) isEventMessage_EventType() {}

type EventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseId string `protobuf:"bytes,1,opt,name=response_id,json=responseId,proto3" json:"response_id,omitempty"`
	Event      *Event `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *EventResponse) Reset() {
	*x = EventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_core_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventResponse) ProtoMessage() {}

func (x *EventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_core_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventResponse.ProtoReflect.Descriptor instead.
func (*EventResponse) Descriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{5}
}

func (x *EventResponse) GetResponseId() string {
	if x != nil {
		return x.ResponseId
	}
	return ""
}

func (x *EventResponse) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Headers map[string]string      `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body    []byte                 `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Created *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_core_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_core_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{6}
}

func (x *Event) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *Event) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *Event) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

var File_core_v1_core_proto protoreflect.FileDescriptor

var file_core_v1_core_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x01, 0x0a, 0x10, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2d, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01,
	0x52, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x6c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x22, 0x3c, 0x0a, 0x11,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x95, 0x01, 0x0a, 0x08, 0x4c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0c, 0xfa, 0x42, 0x09, 0x1a, 0x07, 0x18, 0xc9, 0xfe, 0x03,
	0x20, 0xd0, 0x0f, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x37, 0x0a, 0x08, 0x74, 0x72, 0x69,
	0x67, 0x67, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x48, 0x00, 0x52, 0x08, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65,
	0x72, 0x73, 0x12, 0x25, 0x0a, 0x0d, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x72, 0x65,
	0x67, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x74, 0x72, 0x69,
	0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x67, 0x65, 0x78, 0x42, 0x07, 0x0a, 0x05, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x22, 0x26, 0x0a, 0x08, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x22, 0x8a, 0x01, 0x0a, 0x0c, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x48, 0x00, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x68, 0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x22, 0xcc, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x3d, 0x0a, 0x07, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x34,
	0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x32, 0xe0, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x71, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x21, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x3a, 0x01, 0x2a, 0x12, 0x5e, 0x0a, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x70, 0x69,
	0x63, 0x12, 0x1d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13,
	0x22, 0x0e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x3a, 0x01, 0x2a, 0x42, 0x14, 0x5a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x76, 0x31, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_core_v1_core_proto_rawDescOnce sync.Once
	file_core_v1_core_proto_rawDescData = file_core_v1_core_proto_rawDesc
)

func file_core_v1_core_proto_rawDescGZIP() []byte {
	file_core_v1_core_proto_rawDescOnce.Do(func() {
		file_core_v1_core_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_v1_core_proto_rawDescData)
	})
	return file_core_v1_core_proto_rawDescData
}

var file_core_v1_core_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_core_v1_core_proto_goTypes = []interface{}{
	(*SubscribeRequest)(nil),      // 0: control.core.v1.SubscribeRequest
	(*SubscribeResponse)(nil),     // 1: control.core.v1.SubscribeResponse
	(*Listener)(nil),              // 2: control.core.v1.Listener
	(*Triggers)(nil),              // 3: control.core.v1.Triggers
	(*EventMessage)(nil),          // 4: control.core.v1.EventMessage
	(*EventResponse)(nil),         // 5: control.core.v1.EventResponse
	(*Event)(nil),                 // 6: control.core.v1.Event
	nil,                           // 7: control.core.v1.Event.HeadersEntry
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 9: google.protobuf.Empty
}
var file_core_v1_core_proto_depIdxs = []int32{
	2, // 0: control.core.v1.SubscribeRequest.listener:type_name -> control.core.v1.Listener
	3, // 1: control.core.v1.Listener.triggers:type_name -> control.core.v1.Triggers
	6, // 2: control.core.v1.EventMessage.event:type_name -> control.core.v1.Event
	5, // 3: control.core.v1.EventMessage.response:type_name -> control.core.v1.EventResponse
	6, // 4: control.core.v1.EventResponse.event:type_name -> control.core.v1.Event
	7, // 5: control.core.v1.Event.headers:type_name -> control.core.v1.Event.HeadersEntry
	8, // 6: control.core.v1.Event.created:type_name -> google.protobuf.Timestamp
	0, // 7: control.core.v1.CoreService.Subscribe:input_type -> control.core.v1.SubscribeRequest
	4, // 8: control.core.v1.CoreService.EventTopic:input_type -> control.core.v1.EventMessage
	1, // 9: control.core.v1.CoreService.Subscribe:output_type -> control.core.v1.SubscribeResponse
	9, // 10: control.core.v1.CoreService.EventTopic:output_type -> google.protobuf.Empty
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_core_v1_core_proto_init() }
func file_core_v1_core_proto_init() {
	if File_core_v1_core_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_v1_core_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequest); i {
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
		file_core_v1_core_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeResponse); i {
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
		file_core_v1_core_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Listener); i {
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
		file_core_v1_core_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Triggers); i {
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
		file_core_v1_core_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventMessage); i {
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
		file_core_v1_core_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventResponse); i {
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
		file_core_v1_core_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
	file_core_v1_core_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Listener_Triggers)(nil),
		(*Listener_TriggerRegex)(nil),
	}
	file_core_v1_core_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*EventMessage_Event)(nil),
		(*EventMessage_Response)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_core_v1_core_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_core_v1_core_proto_goTypes,
		DependencyIndexes: file_core_v1_core_proto_depIdxs,
		MessageInfos:      file_core_v1_core_proto_msgTypes,
	}.Build()
	File_core_v1_core_proto = out.File
	file_core_v1_core_proto_rawDesc = nil
	file_core_v1_core_proto_goTypes = nil
	file_core_v1_core_proto_depIdxs = nil
}
