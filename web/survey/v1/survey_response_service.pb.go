// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.2
// source: web/survey/v1/survey_response_service.proto

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

type ResponseGetAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SurveyId int32 `protobuf:"varint,1,opt,name=survey_id,json=surveyId,proto3" json:"survey_id,omitempty"`
	// Pagination
	Page           int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	ContentPerPage int32 `protobuf:"varint,3,opt,name=content_per_page,json=contentPerPage,proto3" json:"content_per_page,omitempty"`
}

func (x *ResponseGetAllRequest) Reset() {
	*x = ResponseGetAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_survey_v1_survey_response_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseGetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseGetAllRequest) ProtoMessage() {}

func (x *ResponseGetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_survey_v1_survey_response_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseGetAllRequest.ProtoReflect.Descriptor instead.
func (*ResponseGetAllRequest) Descriptor() ([]byte, []int) {
	return file_web_survey_v1_survey_response_service_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseGetAllRequest) GetSurveyId() int32 {
	if x != nil {
		return x.SurveyId
	}
	return 0
}

func (x *ResponseGetAllRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ResponseGetAllRequest) GetContentPerPage() int32 {
	if x != nil {
		return x.ContentPerPage
	}
	return 0
}

type ResponseGetAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response []*SurveyResponse `protobuf:"bytes,1,rep,name=response,proto3" json:"response,omitempty"`
	// Pagination
	Count   int32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	MaxPage int32 `protobuf:"varint,3,opt,name=max_page,json=maxPage,proto3" json:"max_page,omitempty"`
}

func (x *ResponseGetAllResponse) Reset() {
	*x = ResponseGetAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_survey_v1_survey_response_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseGetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseGetAllResponse) ProtoMessage() {}

func (x *ResponseGetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_survey_v1_survey_response_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseGetAllResponse.ProtoReflect.Descriptor instead.
func (*ResponseGetAllResponse) Descriptor() ([]byte, []int) {
	return file_web_survey_v1_survey_response_service_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseGetAllResponse) GetResponse() []*SurveyResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *ResponseGetAllResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ResponseGetAllResponse) GetMaxPage() int32 {
	if x != nil {
		return x.MaxPage
	}
	return 0
}

type ResponseGetOneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseId int32 `protobuf:"varint,1,opt,name=response_id,json=responseId,proto3" json:"response_id,omitempty"`
}

func (x *ResponseGetOneRequest) Reset() {
	*x = ResponseGetOneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_survey_v1_survey_response_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseGetOneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseGetOneRequest) ProtoMessage() {}

func (x *ResponseGetOneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_survey_v1_survey_response_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseGetOneRequest.ProtoReflect.Descriptor instead.
func (*ResponseGetOneRequest) Descriptor() ([]byte, []int) {
	return file_web_survey_v1_survey_response_service_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseGetOneRequest) GetResponseId() int32 {
	if x != nil {
		return x.ResponseId
	}
	return 0
}

type ResponseGetOneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *SurveyResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *ResponseGetOneResponse) Reset() {
	*x = ResponseGetOneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_survey_v1_survey_response_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseGetOneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseGetOneResponse) ProtoMessage() {}

func (x *ResponseGetOneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_survey_v1_survey_response_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseGetOneResponse.ProtoReflect.Descriptor instead.
func (*ResponseGetOneResponse) Descriptor() ([]byte, []int) {
	return file_web_survey_v1_survey_response_service_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseGetOneResponse) GetResponse() *SurveyResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

type ResponseAnswerGetListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseId int32 `protobuf:"varint,1,opt,name=response_id,json=responseId,proto3" json:"response_id,omitempty"`
	QuestionId int32 `protobuf:"varint,2,opt,name=question_id,json=questionId,proto3" json:"question_id,omitempty"`
}

func (x *ResponseAnswerGetListRequest) Reset() {
	*x = ResponseAnswerGetListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_survey_v1_survey_response_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseAnswerGetListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseAnswerGetListRequest) ProtoMessage() {}

func (x *ResponseAnswerGetListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_survey_v1_survey_response_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseAnswerGetListRequest.ProtoReflect.Descriptor instead.
func (*ResponseAnswerGetListRequest) Descriptor() ([]byte, []int) {
	return file_web_survey_v1_survey_response_service_proto_rawDescGZIP(), []int{4}
}

func (x *ResponseAnswerGetListRequest) GetResponseId() int32 {
	if x != nil {
		return x.ResponseId
	}
	return 0
}

func (x *ResponseAnswerGetListRequest) GetQuestionId() int32 {
	if x != nil {
		return x.QuestionId
	}
	return 0
}

type ResponseAnswerGetListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Answers []*SurveyResponseAnswer `protobuf:"bytes,1,rep,name=answers,proto3" json:"answers,omitempty"`
}

func (x *ResponseAnswerGetListResponse) Reset() {
	*x = ResponseAnswerGetListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_survey_v1_survey_response_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseAnswerGetListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseAnswerGetListResponse) ProtoMessage() {}

func (x *ResponseAnswerGetListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_survey_v1_survey_response_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseAnswerGetListResponse.ProtoReflect.Descriptor instead.
func (*ResponseAnswerGetListResponse) Descriptor() ([]byte, []int) {
	return file_web_survey_v1_survey_response_service_proto_rawDescGZIP(), []int{5}
}

func (x *ResponseAnswerGetListResponse) GetAnswers() []*SurveyResponseAnswer {
	if x != nil {
		return x.Answers
	}
	return nil
}

var File_web_survey_v1_survey_response_service_proto protoreflect.FileDescriptor

var file_web_survey_v1_survey_response_service_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x77, 0x65, 0x62, 0x2f, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x77,
	0x65, 0x62, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x77, 0x65,
	0x62, 0x2f, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a, 0x15, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x65, 0x72,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x22, 0x84, 0x01, 0x0a, 0x16,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73,
	0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x50, 0x61,
	0x67, 0x65, 0x22, 0x38, 0x0a, 0x15, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x47, 0x65,
	0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x64, 0x22, 0x53, 0x0a, 0x16,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73,
	0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x60, 0x0a, 0x1c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x22, 0x5e, 0x0a, 0x1d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75, 0x72, 0x76,
	0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x07, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x73, 0x32, 0xc9, 0x02, 0x0a, 0x15, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a,
	0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12,
	0x24, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75, 0x72, 0x76,
	0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x0e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x24,
	0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x47, 0x65, 0x74,
	0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x72, 0x0a, 0x15, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2c, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x64,
	0x65, 0x61, 0x74, 0x65, 0x63, 0x68, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x69, 0x64, 0x65, 0x61, 0x74,
	0x65, 0x63, 0x68, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x6f, 0x2f, 0x77, 0x65, 0x62, 0x2f,
	0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x77, 0x65,
	0x62, 0x5f, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_web_survey_v1_survey_response_service_proto_rawDescOnce sync.Once
	file_web_survey_v1_survey_response_service_proto_rawDescData = file_web_survey_v1_survey_response_service_proto_rawDesc
)

func file_web_survey_v1_survey_response_service_proto_rawDescGZIP() []byte {
	file_web_survey_v1_survey_response_service_proto_rawDescOnce.Do(func() {
		file_web_survey_v1_survey_response_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_web_survey_v1_survey_response_service_proto_rawDescData)
	})
	return file_web_survey_v1_survey_response_service_proto_rawDescData
}

var file_web_survey_v1_survey_response_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_web_survey_v1_survey_response_service_proto_goTypes = []interface{}{
	(*ResponseGetAllRequest)(nil),         // 0: web.survey.v1.ResponseGetAllRequest
	(*ResponseGetAllResponse)(nil),        // 1: web.survey.v1.ResponseGetAllResponse
	(*ResponseGetOneRequest)(nil),         // 2: web.survey.v1.ResponseGetOneRequest
	(*ResponseGetOneResponse)(nil),        // 3: web.survey.v1.ResponseGetOneResponse
	(*ResponseAnswerGetListRequest)(nil),  // 4: web.survey.v1.ResponseAnswerGetListRequest
	(*ResponseAnswerGetListResponse)(nil), // 5: web.survey.v1.ResponseAnswerGetListResponse
	(*SurveyResponse)(nil),                // 6: web.survey.v1.SurveyResponse
	(*SurveyResponseAnswer)(nil),          // 7: web.survey.v1.SurveyResponseAnswer
}
var file_web_survey_v1_survey_response_service_proto_depIdxs = []int32{
	6, // 0: web.survey.v1.ResponseGetAllResponse.response:type_name -> web.survey.v1.SurveyResponse
	6, // 1: web.survey.v1.ResponseGetOneResponse.response:type_name -> web.survey.v1.SurveyResponse
	7, // 2: web.survey.v1.ResponseAnswerGetListResponse.answers:type_name -> web.survey.v1.SurveyResponseAnswer
	0, // 3: web.survey.v1.SurveyResponseService.ResponseGetAll:input_type -> web.survey.v1.ResponseGetAllRequest
	2, // 4: web.survey.v1.SurveyResponseService.ResponseGetOne:input_type -> web.survey.v1.ResponseGetOneRequest
	4, // 5: web.survey.v1.SurveyResponseService.ResponseAnswerGetList:input_type -> web.survey.v1.ResponseAnswerGetListRequest
	1, // 6: web.survey.v1.SurveyResponseService.ResponseGetAll:output_type -> web.survey.v1.ResponseGetAllResponse
	3, // 7: web.survey.v1.SurveyResponseService.ResponseGetOne:output_type -> web.survey.v1.ResponseGetOneResponse
	5, // 8: web.survey.v1.SurveyResponseService.ResponseAnswerGetList:output_type -> web.survey.v1.ResponseAnswerGetListResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_web_survey_v1_survey_response_service_proto_init() }
func file_web_survey_v1_survey_response_service_proto_init() {
	if File_web_survey_v1_survey_response_service_proto != nil {
		return
	}
	file_web_survey_v1_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_web_survey_v1_survey_response_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseGetAllRequest); i {
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
		file_web_survey_v1_survey_response_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseGetAllResponse); i {
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
		file_web_survey_v1_survey_response_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseGetOneRequest); i {
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
		file_web_survey_v1_survey_response_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseGetOneResponse); i {
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
		file_web_survey_v1_survey_response_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseAnswerGetListRequest); i {
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
		file_web_survey_v1_survey_response_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseAnswerGetListResponse); i {
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
			RawDescriptor: file_web_survey_v1_survey_response_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_web_survey_v1_survey_response_service_proto_goTypes,
		DependencyIndexes: file_web_survey_v1_survey_response_service_proto_depIdxs,
		MessageInfos:      file_web_survey_v1_survey_response_service_proto_msgTypes,
	}.Build()
	File_web_survey_v1_survey_response_service_proto = out.File
	file_web_survey_v1_survey_response_service_proto_rawDesc = nil
	file_web_survey_v1_survey_response_service_proto_goTypes = nil
	file_web_survey_v1_survey_response_service_proto_depIdxs = nil
}
