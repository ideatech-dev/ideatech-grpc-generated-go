// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.2
// source: web/survey/v1/webhooks.proto

package grpcweb_survey_v1

import (
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

type WebhookGetOneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WebhookId    int32 `protobuf:"varint,1,opt,name=webhook_id,json=webhookId,proto3" json:"webhook_id,omitempty"`
	SurveyPageId int32 `protobuf:"varint,2,opt,name=survey_page_id,json=surveyPageId,proto3" json:"survey_page_id,omitempty"`
}

func (x *WebhookGetOneRequest) Reset() {
	*x = WebhookGetOneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_survey_v1_webhooks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebhookGetOneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebhookGetOneRequest) ProtoMessage() {}

func (x *WebhookGetOneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_survey_v1_webhooks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebhookGetOneRequest.ProtoReflect.Descriptor instead.
func (*WebhookGetOneRequest) Descriptor() ([]byte, []int) {
	return file_web_survey_v1_webhooks_proto_rawDescGZIP(), []int{0}
}

func (x *WebhookGetOneRequest) GetWebhookId() int32 {
	if x != nil {
		return x.WebhookId
	}
	return 0
}

func (x *WebhookGetOneRequest) GetSurveyPageId() int32 {
	if x != nil {
		return x.SurveyPageId
	}
	return 0
}

type WebhookGetOneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Webhook *Webhook `protobuf:"bytes,1,opt,name=webhook,proto3" json:"webhook,omitempty"`
}

func (x *WebhookGetOneResponse) Reset() {
	*x = WebhookGetOneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_survey_v1_webhooks_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebhookGetOneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebhookGetOneResponse) ProtoMessage() {}

func (x *WebhookGetOneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_survey_v1_webhooks_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebhookGetOneResponse.ProtoReflect.Descriptor instead.
func (*WebhookGetOneResponse) Descriptor() ([]byte, []int) {
	return file_web_survey_v1_webhooks_proto_rawDescGZIP(), []int{1}
}

func (x *WebhookGetOneResponse) GetWebhook() *Webhook {
	if x != nil {
		return x.Webhook
	}
	return nil
}

type WebhookCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SurveyPageId int32 `protobuf:"varint,1,opt,name=survey_page_id,json=surveyPageId,proto3" json:"survey_page_id,omitempty"`
	// Payloads
	Link   string     `protobuf:"bytes,2,opt,name=link,proto3" json:"link,omitempty"`
	Method MethodType `protobuf:"varint,3,opt,name=method,proto3,enum=web.survey.v1.MethodType" json:"method,omitempty"`
}

func (x *WebhookCreateRequest) Reset() {
	*x = WebhookCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_survey_v1_webhooks_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebhookCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebhookCreateRequest) ProtoMessage() {}

func (x *WebhookCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_survey_v1_webhooks_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebhookCreateRequest.ProtoReflect.Descriptor instead.
func (*WebhookCreateRequest) Descriptor() ([]byte, []int) {
	return file_web_survey_v1_webhooks_proto_rawDescGZIP(), []int{2}
}

func (x *WebhookCreateRequest) GetSurveyPageId() int32 {
	if x != nil {
		return x.SurveyPageId
	}
	return 0
}

func (x *WebhookCreateRequest) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *WebhookCreateRequest) GetMethod() MethodType {
	if x != nil {
		return x.Method
	}
	return MethodType_ALL_UNSPECIFIED
}

type WebhookCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Webhook *Webhook `protobuf:"bytes,1,opt,name=webhook,proto3" json:"webhook,omitempty"`
}

func (x *WebhookCreateResponse) Reset() {
	*x = WebhookCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_survey_v1_webhooks_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebhookCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebhookCreateResponse) ProtoMessage() {}

func (x *WebhookCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_survey_v1_webhooks_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebhookCreateResponse.ProtoReflect.Descriptor instead.
func (*WebhookCreateResponse) Descriptor() ([]byte, []int) {
	return file_web_survey_v1_webhooks_proto_rawDescGZIP(), []int{3}
}

func (x *WebhookCreateResponse) GetWebhook() *Webhook {
	if x != nil {
		return x.Webhook
	}
	return nil
}

type WebhookUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WebhookId int32 `protobuf:"varint,1,opt,name=webhook_id,json=webhookId,proto3" json:"webhook_id,omitempty"`
	// Payloads
	NewLink   string     `protobuf:"bytes,2,opt,name=new_link,json=newLink,proto3" json:"new_link,omitempty"`
	NewMethod MethodType `protobuf:"varint,3,opt,name=new_method,json=newMethod,proto3,enum=web.survey.v1.MethodType" json:"new_method,omitempty"`
}

func (x *WebhookUpdateRequest) Reset() {
	*x = WebhookUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_survey_v1_webhooks_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebhookUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebhookUpdateRequest) ProtoMessage() {}

func (x *WebhookUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_survey_v1_webhooks_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebhookUpdateRequest.ProtoReflect.Descriptor instead.
func (*WebhookUpdateRequest) Descriptor() ([]byte, []int) {
	return file_web_survey_v1_webhooks_proto_rawDescGZIP(), []int{4}
}

func (x *WebhookUpdateRequest) GetWebhookId() int32 {
	if x != nil {
		return x.WebhookId
	}
	return 0
}

func (x *WebhookUpdateRequest) GetNewLink() string {
	if x != nil {
		return x.NewLink
	}
	return ""
}

func (x *WebhookUpdateRequest) GetNewMethod() MethodType {
	if x != nil {
		return x.NewMethod
	}
	return MethodType_ALL_UNSPECIFIED
}

type WebhookUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Webhook *Webhook `protobuf:"bytes,1,opt,name=webhook,proto3" json:"webhook,omitempty"`
}

func (x *WebhookUpdateResponse) Reset() {
	*x = WebhookUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_survey_v1_webhooks_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebhookUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebhookUpdateResponse) ProtoMessage() {}

func (x *WebhookUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_survey_v1_webhooks_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebhookUpdateResponse.ProtoReflect.Descriptor instead.
func (*WebhookUpdateResponse) Descriptor() ([]byte, []int) {
	return file_web_survey_v1_webhooks_proto_rawDescGZIP(), []int{5}
}

func (x *WebhookUpdateResponse) GetWebhook() *Webhook {
	if x != nil {
		return x.Webhook
	}
	return nil
}

type WebhookDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WebhookId int32 `protobuf:"varint,1,opt,name=webhook_id,json=webhookId,proto3" json:"webhook_id,omitempty"`
}

func (x *WebhookDeleteRequest) Reset() {
	*x = WebhookDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_survey_v1_webhooks_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebhookDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebhookDeleteRequest) ProtoMessage() {}

func (x *WebhookDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_survey_v1_webhooks_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebhookDeleteRequest.ProtoReflect.Descriptor instead.
func (*WebhookDeleteRequest) Descriptor() ([]byte, []int) {
	return file_web_survey_v1_webhooks_proto_rawDescGZIP(), []int{6}
}

func (x *WebhookDeleteRequest) GetWebhookId() int32 {
	if x != nil {
		return x.WebhookId
	}
	return 0
}

type WebhookDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WebhookDeleteResponse) Reset() {
	*x = WebhookDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_survey_v1_webhooks_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebhookDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebhookDeleteResponse) ProtoMessage() {}

func (x *WebhookDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_survey_v1_webhooks_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebhookDeleteResponse.ProtoReflect.Descriptor instead.
func (*WebhookDeleteResponse) Descriptor() ([]byte, []int) {
	return file_web_survey_v1_webhooks_proto_rawDescGZIP(), []int{7}
}

var File_web_survey_v1_webhooks_proto protoreflect.FileDescriptor

var file_web_survey_v1_webhooks_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x77, 0x65, 0x62, 0x2f, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2f, 0x76, 0x31, 0x2f,
	0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x77, 0x65, 0x62, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x77,
	0x65, 0x62, 0x2f, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x14, 0x57, 0x65, 0x62, 0x68,
	0x6f, 0x6f, 0x6b, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12,
	0x24, 0x0a, 0x0e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x50,
	0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x15, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30,
	0x0a, 0x07, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x07, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x22, 0x83, 0x01, 0x0a, 0x14, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x75, 0x72,
	0x76, 0x65, 0x79, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x50, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c,
	0x69, 0x6e, 0x6b, 0x12, 0x31, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x49, 0x0a, 0x15, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f,
	0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x30, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x07, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f,
	0x6b, 0x22, 0x8a, 0x01, 0x0a, 0x14, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x65,
	0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x65, 0x77,
	0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x77,
	0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x38, 0x0a, 0x0a, 0x6e, 0x65, 0x77, 0x5f, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73,
	0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x09, 0x6e, 0x65, 0x77, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x49,
	0x0a, 0x15, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x68, 0x6f,
	0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73,
	0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x52, 0x07, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x22, 0x35, 0x0a, 0x14, 0x57, 0x65, 0x62,
	0x68, 0x6f, 0x6f, 0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x49, 0x64,
	0x22, 0x17, 0x0a, 0x15, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x80, 0x03, 0x0a, 0x0e, 0x57, 0x65,
	0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x0d,
	0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x23, 0x2e,
	0x77, 0x65, 0x62, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65,
	0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x24, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0d, 0x57, 0x65, 0x62, 0x68,
	0x6f, 0x6f, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x77, 0x65, 0x62, 0x2e,
	0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f,
	0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x57,
	0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0d, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75, 0x72, 0x76,
	0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x77, 0x65, 0x62,
	0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f,
	0x6f, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x5a, 0x0a, 0x0d, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x23, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75, 0x72,
	0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x54, 0x5a, 0x52,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x64, 0x65, 0x61, 0x74,
	0x65, 0x63, 0x68, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x69, 0x64, 0x65, 0x61, 0x74, 0x65, 0x63, 0x68,
	0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d,
	0x67, 0x6f, 0x2f, 0x77, 0x65, 0x62, 0x2f, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2f, 0x76, 0x31,
	0x3b, 0x67, 0x72, 0x70, 0x63, 0x77, 0x65, 0x62, 0x5f, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_web_survey_v1_webhooks_proto_rawDescOnce sync.Once
	file_web_survey_v1_webhooks_proto_rawDescData = file_web_survey_v1_webhooks_proto_rawDesc
)

func file_web_survey_v1_webhooks_proto_rawDescGZIP() []byte {
	file_web_survey_v1_webhooks_proto_rawDescOnce.Do(func() {
		file_web_survey_v1_webhooks_proto_rawDescData = protoimpl.X.CompressGZIP(file_web_survey_v1_webhooks_proto_rawDescData)
	})
	return file_web_survey_v1_webhooks_proto_rawDescData
}

var file_web_survey_v1_webhooks_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_web_survey_v1_webhooks_proto_goTypes = []interface{}{
	(*WebhookGetOneRequest)(nil),  // 0: web.survey.v1.WebhookGetOneRequest
	(*WebhookGetOneResponse)(nil), // 1: web.survey.v1.WebhookGetOneResponse
	(*WebhookCreateRequest)(nil),  // 2: web.survey.v1.WebhookCreateRequest
	(*WebhookCreateResponse)(nil), // 3: web.survey.v1.WebhookCreateResponse
	(*WebhookUpdateRequest)(nil),  // 4: web.survey.v1.WebhookUpdateRequest
	(*WebhookUpdateResponse)(nil), // 5: web.survey.v1.WebhookUpdateResponse
	(*WebhookDeleteRequest)(nil),  // 6: web.survey.v1.WebhookDeleteRequest
	(*WebhookDeleteResponse)(nil), // 7: web.survey.v1.WebhookDeleteResponse
	(*Webhook)(nil),               // 8: web.survey.v1.Webhook
	(MethodType)(0),               // 9: web.survey.v1.MethodType
}
var file_web_survey_v1_webhooks_proto_depIdxs = []int32{
	8, // 0: web.survey.v1.WebhookGetOneResponse.webhook:type_name -> web.survey.v1.Webhook
	9, // 1: web.survey.v1.WebhookCreateRequest.method:type_name -> web.survey.v1.MethodType
	8, // 2: web.survey.v1.WebhookCreateResponse.webhook:type_name -> web.survey.v1.Webhook
	9, // 3: web.survey.v1.WebhookUpdateRequest.new_method:type_name -> web.survey.v1.MethodType
	8, // 4: web.survey.v1.WebhookUpdateResponse.webhook:type_name -> web.survey.v1.Webhook
	0, // 5: web.survey.v1.WebhookService.WebhookGetOne:input_type -> web.survey.v1.WebhookGetOneRequest
	2, // 6: web.survey.v1.WebhookService.WebhookCreate:input_type -> web.survey.v1.WebhookCreateRequest
	4, // 7: web.survey.v1.WebhookService.WebhookUpdate:input_type -> web.survey.v1.WebhookUpdateRequest
	6, // 8: web.survey.v1.WebhookService.WebhookDelete:input_type -> web.survey.v1.WebhookDeleteRequest
	1, // 9: web.survey.v1.WebhookService.WebhookGetOne:output_type -> web.survey.v1.WebhookGetOneResponse
	3, // 10: web.survey.v1.WebhookService.WebhookCreate:output_type -> web.survey.v1.WebhookCreateResponse
	5, // 11: web.survey.v1.WebhookService.WebhookUpdate:output_type -> web.survey.v1.WebhookUpdateResponse
	7, // 12: web.survey.v1.WebhookService.WebhookDelete:output_type -> web.survey.v1.WebhookDeleteResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_web_survey_v1_webhooks_proto_init() }
func file_web_survey_v1_webhooks_proto_init() {
	if File_web_survey_v1_webhooks_proto != nil {
		return
	}
	file_web_survey_v1_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_web_survey_v1_webhooks_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebhookGetOneRequest); i {
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
		file_web_survey_v1_webhooks_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebhookGetOneResponse); i {
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
		file_web_survey_v1_webhooks_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebhookCreateRequest); i {
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
		file_web_survey_v1_webhooks_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebhookCreateResponse); i {
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
		file_web_survey_v1_webhooks_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebhookUpdateRequest); i {
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
		file_web_survey_v1_webhooks_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebhookUpdateResponse); i {
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
		file_web_survey_v1_webhooks_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebhookDeleteRequest); i {
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
		file_web_survey_v1_webhooks_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebhookDeleteResponse); i {
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
			RawDescriptor: file_web_survey_v1_webhooks_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_web_survey_v1_webhooks_proto_goTypes,
		DependencyIndexes: file_web_survey_v1_webhooks_proto_depIdxs,
		MessageInfos:      file_web_survey_v1_webhooks_proto_msgTypes,
	}.Build()
	File_web_survey_v1_webhooks_proto = out.File
	file_web_survey_v1_webhooks_proto_rawDesc = nil
	file_web_survey_v1_webhooks_proto_goTypes = nil
	file_web_survey_v1_webhooks_proto_depIdxs = nil
}
