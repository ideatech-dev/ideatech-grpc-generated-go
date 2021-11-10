// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.2
// source: web/survey/v1/survey_page_service.proto

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

type PageGetOneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageId int32 `protobuf:"varint,1,opt,name=page_id,json=pageId,proto3" json:"page_id,omitempty"`
	// Preloads
	WithQuestion       bool `protobuf:"varint,2,opt,name=with_question,json=withQuestion,proto3" json:"with_question,omitempty"`
	WithQuestionChoice bool `protobuf:"varint,3,opt,name=with_question_choice,json=withQuestionChoice,proto3" json:"with_question_choice,omitempty"`
}

func (x *PageGetOneRequest) Reset() {
	*x = PageGetOneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_survey_v1_survey_page_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageGetOneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageGetOneRequest) ProtoMessage() {}

func (x *PageGetOneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_survey_v1_survey_page_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageGetOneRequest.ProtoReflect.Descriptor instead.
func (*PageGetOneRequest) Descriptor() ([]byte, []int) {
	return file_web_survey_v1_survey_page_service_proto_rawDescGZIP(), []int{0}
}

func (x *PageGetOneRequest) GetPageId() int32 {
	if x != nil {
		return x.PageId
	}
	return 0
}

func (x *PageGetOneRequest) GetWithQuestion() bool {
	if x != nil {
		return x.WithQuestion
	}
	return false
}

func (x *PageGetOneRequest) GetWithQuestionChoice() bool {
	if x != nil {
		return x.WithQuestionChoice
	}
	return false
}

type PageGetOneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *SurveyPage `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *PageGetOneResponse) Reset() {
	*x = PageGetOneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_survey_v1_survey_page_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageGetOneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageGetOneResponse) ProtoMessage() {}

func (x *PageGetOneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_survey_v1_survey_page_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageGetOneResponse.ProtoReflect.Descriptor instead.
func (*PageGetOneResponse) Descriptor() ([]byte, []int) {
	return file_web_survey_v1_survey_page_service_proto_rawDescGZIP(), []int{1}
}

func (x *PageGetOneResponse) GetPage() *SurveyPage {
	if x != nil {
		return x.Page
	}
	return nil
}

type PageGetAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SurveyId int32 `protobuf:"varint,1,opt,name=survey_id,json=surveyId,proto3" json:"survey_id,omitempty"`
	// Preloads
	WithQuestion       bool `protobuf:"varint,2,opt,name=with_question,json=withQuestion,proto3" json:"with_question,omitempty"`
	WithQuestionOption bool `protobuf:"varint,3,opt,name=with_question_option,json=withQuestionOption,proto3" json:"with_question_option,omitempty"`
	WithQuestionChoice bool `protobuf:"varint,4,opt,name=with_question_choice,json=withQuestionChoice,proto3" json:"with_question_choice,omitempty"`
}

func (x *PageGetAllRequest) Reset() {
	*x = PageGetAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_survey_v1_survey_page_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageGetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageGetAllRequest) ProtoMessage() {}

func (x *PageGetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_survey_v1_survey_page_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageGetAllRequest.ProtoReflect.Descriptor instead.
func (*PageGetAllRequest) Descriptor() ([]byte, []int) {
	return file_web_survey_v1_survey_page_service_proto_rawDescGZIP(), []int{2}
}

func (x *PageGetAllRequest) GetSurveyId() int32 {
	if x != nil {
		return x.SurveyId
	}
	return 0
}

func (x *PageGetAllRequest) GetWithQuestion() bool {
	if x != nil {
		return x.WithQuestion
	}
	return false
}

func (x *PageGetAllRequest) GetWithQuestionOption() bool {
	if x != nil {
		return x.WithQuestionOption
	}
	return false
}

func (x *PageGetAllRequest) GetWithQuestionChoice() bool {
	if x != nil {
		return x.WithQuestionChoice
	}
	return false
}

type PageGetAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pages []*SurveyPage `protobuf:"bytes,1,rep,name=pages,proto3" json:"pages,omitempty"`
}

func (x *PageGetAllResponse) Reset() {
	*x = PageGetAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_survey_v1_survey_page_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageGetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageGetAllResponse) ProtoMessage() {}

func (x *PageGetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_survey_v1_survey_page_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageGetAllResponse.ProtoReflect.Descriptor instead.
func (*PageGetAllResponse) Descriptor() ([]byte, []int) {
	return file_web_survey_v1_survey_page_service_proto_rawDescGZIP(), []int{3}
}

func (x *PageGetAllResponse) GetPages() []*SurveyPage {
	if x != nil {
		return x.Pages
	}
	return nil
}

type PageCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SurveyId int32 `protobuf:"varint,1,opt,name=survey_id,json=surveyId,proto3" json:"survey_id,omitempty"`
}

func (x *PageCreateRequest) Reset() {
	*x = PageCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_survey_v1_survey_page_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageCreateRequest) ProtoMessage() {}

func (x *PageCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_survey_v1_survey_page_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageCreateRequest.ProtoReflect.Descriptor instead.
func (*PageCreateRequest) Descriptor() ([]byte, []int) {
	return file_web_survey_v1_survey_page_service_proto_rawDescGZIP(), []int{4}
}

func (x *PageCreateRequest) GetSurveyId() int32 {
	if x != nil {
		return x.SurveyId
	}
	return 0
}

type PageCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *SurveyPage `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *PageCreateResponse) Reset() {
	*x = PageCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_survey_v1_survey_page_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageCreateResponse) ProtoMessage() {}

func (x *PageCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_survey_v1_survey_page_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageCreateResponse.ProtoReflect.Descriptor instead.
func (*PageCreateResponse) Descriptor() ([]byte, []int) {
	return file_web_survey_v1_survey_page_service_proto_rawDescGZIP(), []int{5}
}

func (x *PageCreateResponse) GetPage() *SurveyPage {
	if x != nil {
		return x.Page
	}
	return nil
}

type PageDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageId int32 `protobuf:"varint,1,opt,name=page_id,json=pageId,proto3" json:"page_id,omitempty"`
}

func (x *PageDeleteRequest) Reset() {
	*x = PageDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_survey_v1_survey_page_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageDeleteRequest) ProtoMessage() {}

func (x *PageDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_survey_v1_survey_page_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageDeleteRequest.ProtoReflect.Descriptor instead.
func (*PageDeleteRequest) Descriptor() ([]byte, []int) {
	return file_web_survey_v1_survey_page_service_proto_rawDescGZIP(), []int{6}
}

func (x *PageDeleteRequest) GetPageId() int32 {
	if x != nil {
		return x.PageId
	}
	return 0
}

type PageDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PageDeleteResponse) Reset() {
	*x = PageDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_survey_v1_survey_page_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageDeleteResponse) ProtoMessage() {}

func (x *PageDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_survey_v1_survey_page_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageDeleteResponse.ProtoReflect.Descriptor instead.
func (*PageDeleteResponse) Descriptor() ([]byte, []int) {
	return file_web_survey_v1_survey_page_service_proto_rawDescGZIP(), []int{7}
}

var File_web_survey_v1_survey_page_service_proto protoreflect.FileDescriptor

var file_web_survey_v1_survey_page_service_proto_rawDesc = []byte{
	0x0a, 0x27, 0x77, 0x65, 0x62, 0x2f, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x77, 0x65, 0x62, 0x2e, 0x73,
	0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x77, 0x65, 0x62, 0x2f, 0x73, 0x75,
	0x72, 0x76, 0x65, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x11, 0x50, 0x61, 0x67, 0x65, 0x47, 0x65, 0x74, 0x4f,
	0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x77, 0x69, 0x74, 0x68, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x14, 0x77, 0x69, 0x74, 0x68, 0x5f,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x77, 0x69, 0x74, 0x68, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x22, 0x43, 0x0a, 0x12, 0x50, 0x61, 0x67,
	0x65, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2d, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x77, 0x65, 0x62, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75,
	0x72, 0x76, 0x65, 0x79, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0xb9,
	0x01, 0x0a, 0x11, 0x50, 0x61, 0x67, 0x65, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x49,
	0x64, 0x12, 0x23, 0x0a, 0x0d, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x77, 0x69, 0x74, 0x68, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x14, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x77, 0x69, 0x74, 0x68, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x14, 0x77, 0x69, 0x74, 0x68,
	0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x77, 0x69, 0x74, 0x68, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x22, 0x45, 0x0a, 0x12, 0x50, 0x61,
	0x67, 0x65, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2f, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x50, 0x61, 0x67, 0x65, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x73, 0x22, 0x30, 0x0a, 0x11, 0x50, 0x61, 0x67, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x75, 0x72, 0x76, 0x65,
	0x79, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x12, 0x50, 0x61, 0x67, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75,
	0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x2c, 0x0a, 0x11, 0x50, 0x61, 0x67, 0x65,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x70, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x50, 0x61, 0x67, 0x65, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xdf, 0x02, 0x0a,
	0x11, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x50, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x51, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x65, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65,
	0x12, 0x20, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x65, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x12, 0x20, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75, 0x72, 0x76,
	0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x65,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75, 0x72,
	0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73,
	0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0a, 0x50,
	0x61, 0x67, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x77, 0x65, 0x62, 0x2e,
	0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x77, 0x65,
	0x62, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x54,
	0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x64, 0x65,
	0x61, 0x74, 0x65, 0x63, 0x68, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x69, 0x64, 0x65, 0x61, 0x74, 0x65,
	0x63, 0x68, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x77, 0x65, 0x62, 0x2f, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2f,
	0x76, 0x31, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x77, 0x65, 0x62, 0x5f, 0x73, 0x75, 0x72, 0x76, 0x65,
	0x79, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_web_survey_v1_survey_page_service_proto_rawDescOnce sync.Once
	file_web_survey_v1_survey_page_service_proto_rawDescData = file_web_survey_v1_survey_page_service_proto_rawDesc
)

func file_web_survey_v1_survey_page_service_proto_rawDescGZIP() []byte {
	file_web_survey_v1_survey_page_service_proto_rawDescOnce.Do(func() {
		file_web_survey_v1_survey_page_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_web_survey_v1_survey_page_service_proto_rawDescData)
	})
	return file_web_survey_v1_survey_page_service_proto_rawDescData
}

var file_web_survey_v1_survey_page_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_web_survey_v1_survey_page_service_proto_goTypes = []interface{}{
	(*PageGetOneRequest)(nil),  // 0: web.survey.v1.PageGetOneRequest
	(*PageGetOneResponse)(nil), // 1: web.survey.v1.PageGetOneResponse
	(*PageGetAllRequest)(nil),  // 2: web.survey.v1.PageGetAllRequest
	(*PageGetAllResponse)(nil), // 3: web.survey.v1.PageGetAllResponse
	(*PageCreateRequest)(nil),  // 4: web.survey.v1.PageCreateRequest
	(*PageCreateResponse)(nil), // 5: web.survey.v1.PageCreateResponse
	(*PageDeleteRequest)(nil),  // 6: web.survey.v1.PageDeleteRequest
	(*PageDeleteResponse)(nil), // 7: web.survey.v1.PageDeleteResponse
	(*SurveyPage)(nil),         // 8: web.survey.v1.SurveyPage
}
var file_web_survey_v1_survey_page_service_proto_depIdxs = []int32{
	8, // 0: web.survey.v1.PageGetOneResponse.page:type_name -> web.survey.v1.SurveyPage
	8, // 1: web.survey.v1.PageGetAllResponse.pages:type_name -> web.survey.v1.SurveyPage
	8, // 2: web.survey.v1.PageCreateResponse.page:type_name -> web.survey.v1.SurveyPage
	0, // 3: web.survey.v1.SurveyPageService.PageGetOne:input_type -> web.survey.v1.PageGetOneRequest
	2, // 4: web.survey.v1.SurveyPageService.PageGetAll:input_type -> web.survey.v1.PageGetAllRequest
	4, // 5: web.survey.v1.SurveyPageService.PageCreate:input_type -> web.survey.v1.PageCreateRequest
	6, // 6: web.survey.v1.SurveyPageService.PageDelete:input_type -> web.survey.v1.PageDeleteRequest
	1, // 7: web.survey.v1.SurveyPageService.PageGetOne:output_type -> web.survey.v1.PageGetOneResponse
	3, // 8: web.survey.v1.SurveyPageService.PageGetAll:output_type -> web.survey.v1.PageGetAllResponse
	5, // 9: web.survey.v1.SurveyPageService.PageCreate:output_type -> web.survey.v1.PageCreateResponse
	7, // 10: web.survey.v1.SurveyPageService.PageDelete:output_type -> web.survey.v1.PageDeleteResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_web_survey_v1_survey_page_service_proto_init() }
func file_web_survey_v1_survey_page_service_proto_init() {
	if File_web_survey_v1_survey_page_service_proto != nil {
		return
	}
	file_web_survey_v1_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_web_survey_v1_survey_page_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageGetOneRequest); i {
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
		file_web_survey_v1_survey_page_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageGetOneResponse); i {
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
		file_web_survey_v1_survey_page_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageGetAllRequest); i {
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
		file_web_survey_v1_survey_page_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageGetAllResponse); i {
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
		file_web_survey_v1_survey_page_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageCreateRequest); i {
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
		file_web_survey_v1_survey_page_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageCreateResponse); i {
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
		file_web_survey_v1_survey_page_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageDeleteRequest); i {
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
		file_web_survey_v1_survey_page_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageDeleteResponse); i {
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
			RawDescriptor: file_web_survey_v1_survey_page_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_web_survey_v1_survey_page_service_proto_goTypes,
		DependencyIndexes: file_web_survey_v1_survey_page_service_proto_depIdxs,
		MessageInfos:      file_web_survey_v1_survey_page_service_proto_msgTypes,
	}.Build()
	File_web_survey_v1_survey_page_service_proto = out.File
	file_web_survey_v1_survey_page_service_proto_rawDesc = nil
	file_web_survey_v1_survey_page_service_proto_goTypes = nil
	file_web_survey_v1_survey_page_service_proto_depIdxs = nil
}
