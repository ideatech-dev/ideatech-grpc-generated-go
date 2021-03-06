// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.2
// source: web/survey/v1/survey_service.proto

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

type SurveyGetOneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SurveyId int32 `protobuf:"varint,1,opt,name=survey_id,json=surveyId,proto3" json:"survey_id,omitempty"`
}

func (x *SurveyGetOneRequest) Reset() {
	*x = SurveyGetOneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_survey_v1_survey_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SurveyGetOneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SurveyGetOneRequest) ProtoMessage() {}

func (x *SurveyGetOneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_survey_v1_survey_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SurveyGetOneRequest.ProtoReflect.Descriptor instead.
func (*SurveyGetOneRequest) Descriptor() ([]byte, []int) {
	return file_web_survey_v1_survey_service_proto_rawDescGZIP(), []int{0}
}

func (x *SurveyGetOneRequest) GetSurveyId() int32 {
	if x != nil {
		return x.SurveyId
	}
	return 0
}

type SurveyGetOneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Survey *Survey `protobuf:"bytes,1,opt,name=survey,proto3" json:"survey,omitempty"`
}

func (x *SurveyGetOneResponse) Reset() {
	*x = SurveyGetOneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_survey_v1_survey_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SurveyGetOneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SurveyGetOneResponse) ProtoMessage() {}

func (x *SurveyGetOneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_survey_v1_survey_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SurveyGetOneResponse.ProtoReflect.Descriptor instead.
func (*SurveyGetOneResponse) Descriptor() ([]byte, []int) {
	return file_web_survey_v1_survey_service_proto_rawDescGZIP(), []int{1}
}

func (x *SurveyGetOneResponse) GetSurvey() *Survey {
	if x != nil {
		return x.Survey
	}
	return nil
}

type SurveyGetAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrganizationId int32      `protobuf:"varint,1,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	FilterType     SurveyType `protobuf:"varint,2,opt,name=filter_type,json=filterType,proto3,enum=web.survey.v1.SurveyType" json:"filter_type,omitempty"`
	// Filtering
	// pagination
	Page           int32 `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	ContentPerPage int32 `protobuf:"varint,4,opt,name=content_per_page,json=contentPerPage,proto3" json:"content_per_page,omitempty"`
	// searching
	SearchByName string `protobuf:"bytes,5,opt,name=search_by_name,json=searchByName,proto3" json:"search_by_name,omitempty"`
}

func (x *SurveyGetAllRequest) Reset() {
	*x = SurveyGetAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_survey_v1_survey_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SurveyGetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SurveyGetAllRequest) ProtoMessage() {}

func (x *SurveyGetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_survey_v1_survey_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SurveyGetAllRequest.ProtoReflect.Descriptor instead.
func (*SurveyGetAllRequest) Descriptor() ([]byte, []int) {
	return file_web_survey_v1_survey_service_proto_rawDescGZIP(), []int{2}
}

func (x *SurveyGetAllRequest) GetOrganizationId() int32 {
	if x != nil {
		return x.OrganizationId
	}
	return 0
}

func (x *SurveyGetAllRequest) GetFilterType() SurveyType {
	if x != nil {
		return x.FilterType
	}
	return SurveyType_SURVEY_TYPE_ALL_UNSPECIFIED
}

func (x *SurveyGetAllRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SurveyGetAllRequest) GetContentPerPage() int32 {
	if x != nil {
		return x.ContentPerPage
	}
	return 0
}

func (x *SurveyGetAllRequest) GetSearchByName() string {
	if x != nil {
		return x.SearchByName
	}
	return ""
}

type SurveyGetAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Surveys []*Survey `protobuf:"bytes,1,rep,name=surveys,proto3" json:"surveys,omitempty"`
	// Pagination
	Count   int32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	MaxPage int32 `protobuf:"varint,3,opt,name=max_page,json=maxPage,proto3" json:"max_page,omitempty"`
}

func (x *SurveyGetAllResponse) Reset() {
	*x = SurveyGetAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_survey_v1_survey_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SurveyGetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SurveyGetAllResponse) ProtoMessage() {}

func (x *SurveyGetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_survey_v1_survey_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SurveyGetAllResponse.ProtoReflect.Descriptor instead.
func (*SurveyGetAllResponse) Descriptor() ([]byte, []int) {
	return file_web_survey_v1_survey_service_proto_rawDescGZIP(), []int{3}
}

func (x *SurveyGetAllResponse) GetSurveys() []*Survey {
	if x != nil {
		return x.Surveys
	}
	return nil
}

func (x *SurveyGetAllResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *SurveyGetAllResponse) GetMaxPage() int32 {
	if x != nil {
		return x.MaxPage
	}
	return 0
}

type SurveyCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrganizationId int32 `protobuf:"varint,1,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	// Payloads
	Name string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type SurveyType `protobuf:"varint,3,opt,name=type,proto3,enum=web.survey.v1.SurveyType" json:"type,omitempty"`
}

func (x *SurveyCreateRequest) Reset() {
	*x = SurveyCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_survey_v1_survey_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SurveyCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SurveyCreateRequest) ProtoMessage() {}

func (x *SurveyCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_survey_v1_survey_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SurveyCreateRequest.ProtoReflect.Descriptor instead.
func (*SurveyCreateRequest) Descriptor() ([]byte, []int) {
	return file_web_survey_v1_survey_service_proto_rawDescGZIP(), []int{4}
}

func (x *SurveyCreateRequest) GetOrganizationId() int32 {
	if x != nil {
		return x.OrganizationId
	}
	return 0
}

func (x *SurveyCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SurveyCreateRequest) GetType() SurveyType {
	if x != nil {
		return x.Type
	}
	return SurveyType_SURVEY_TYPE_ALL_UNSPECIFIED
}

type SurveyCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Survey *Survey `protobuf:"bytes,1,opt,name=survey,proto3" json:"survey,omitempty"`
}

func (x *SurveyCreateResponse) Reset() {
	*x = SurveyCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_survey_v1_survey_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SurveyCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SurveyCreateResponse) ProtoMessage() {}

func (x *SurveyCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_survey_v1_survey_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SurveyCreateResponse.ProtoReflect.Descriptor instead.
func (*SurveyCreateResponse) Descriptor() ([]byte, []int) {
	return file_web_survey_v1_survey_service_proto_rawDescGZIP(), []int{5}
}

func (x *SurveyCreateResponse) GetSurvey() *Survey {
	if x != nil {
		return x.Survey
	}
	return nil
}

type SurveyUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SurveyId int32 `protobuf:"varint,1,opt,name=survey_id,json=surveyId,proto3" json:"survey_id,omitempty"`
	// Payloads
	NewName string `protobuf:"bytes,2,opt,name=new_name,json=newName,proto3" json:"new_name,omitempty"`
}

func (x *SurveyUpdateRequest) Reset() {
	*x = SurveyUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_survey_v1_survey_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SurveyUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SurveyUpdateRequest) ProtoMessage() {}

func (x *SurveyUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_survey_v1_survey_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SurveyUpdateRequest.ProtoReflect.Descriptor instead.
func (*SurveyUpdateRequest) Descriptor() ([]byte, []int) {
	return file_web_survey_v1_survey_service_proto_rawDescGZIP(), []int{6}
}

func (x *SurveyUpdateRequest) GetSurveyId() int32 {
	if x != nil {
		return x.SurveyId
	}
	return 0
}

func (x *SurveyUpdateRequest) GetNewName() string {
	if x != nil {
		return x.NewName
	}
	return ""
}

type SurveyUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Survey *Survey `protobuf:"bytes,1,opt,name=survey,proto3" json:"survey,omitempty"`
}

func (x *SurveyUpdateResponse) Reset() {
	*x = SurveyUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_survey_v1_survey_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SurveyUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SurveyUpdateResponse) ProtoMessage() {}

func (x *SurveyUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_survey_v1_survey_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SurveyUpdateResponse.ProtoReflect.Descriptor instead.
func (*SurveyUpdateResponse) Descriptor() ([]byte, []int) {
	return file_web_survey_v1_survey_service_proto_rawDescGZIP(), []int{7}
}

func (x *SurveyUpdateResponse) GetSurvey() *Survey {
	if x != nil {
		return x.Survey
	}
	return nil
}

type SurveyDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SurveyId int32 `protobuf:"varint,1,opt,name=survey_id,json=surveyId,proto3" json:"survey_id,omitempty"`
}

func (x *SurveyDeleteRequest) Reset() {
	*x = SurveyDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_survey_v1_survey_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SurveyDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SurveyDeleteRequest) ProtoMessage() {}

func (x *SurveyDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_survey_v1_survey_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SurveyDeleteRequest.ProtoReflect.Descriptor instead.
func (*SurveyDeleteRequest) Descriptor() ([]byte, []int) {
	return file_web_survey_v1_survey_service_proto_rawDescGZIP(), []int{8}
}

func (x *SurveyDeleteRequest) GetSurveyId() int32 {
	if x != nil {
		return x.SurveyId
	}
	return 0
}

type SurveyDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SurveyDeleteResponse) Reset() {
	*x = SurveyDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_survey_v1_survey_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SurveyDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SurveyDeleteResponse) ProtoMessage() {}

func (x *SurveyDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_survey_v1_survey_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SurveyDeleteResponse.ProtoReflect.Descriptor instead.
func (*SurveyDeleteResponse) Descriptor() ([]byte, []int) {
	return file_web_survey_v1_survey_service_proto_rawDescGZIP(), []int{9}
}

var File_web_survey_v1_survey_service_proto protoreflect.FileDescriptor

var file_web_survey_v1_survey_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x77, 0x65, 0x62, 0x2f, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79,
	0x2e, 0x76, 0x31, 0x1a, 0x19, 0x77, 0x65, 0x62, 0x2f, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32,
	0x0a, 0x13, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79,
	0x49, 0x64, 0x22, 0x45, 0x0a, 0x14, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x47, 0x65, 0x74, 0x4f,
	0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x75,
	0x72, 0x76, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x77, 0x65, 0x62,
	0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x72, 0x76, 0x65,
	0x79, 0x52, 0x06, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x22, 0xde, 0x01, 0x0a, 0x13, 0x53, 0x75,
	0x72, 0x76, 0x65, 0x79, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x0b, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x19, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x50, 0x65, 0x72,
	0x50, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x62,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x78, 0x0a, 0x14, 0x53, 0x75,
	0x72, 0x76, 0x65, 0x79, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x07, 0x73, 0x75, 0x72, 0x76,
	0x65, 0x79, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x78,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x61, 0x78,
	0x50, 0x61, 0x67, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x13, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75,
	0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x45, 0x0a, 0x14, 0x53, 0x75, 0x72, 0x76,
	0x65, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2d, 0x0a, 0x06, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x06, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x22,
	0x4d, 0x0a, 0x13, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x75, 0x72, 0x76, 0x65,
	0x79, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x45,
	0x0a, 0x14, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75, 0x72,
	0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x06, 0x73,
	0x75, 0x72, 0x76, 0x65, 0x79, 0x22, 0x32, 0x0a, 0x13, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x49, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x53, 0x75, 0x72,
	0x76, 0x65, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0xcc, 0x03, 0x0a, 0x0d, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x0c, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x47, 0x65, 0x74,
	0x4f, 0x6e, 0x65, 0x12, 0x22, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75,
	0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x47, 0x65,
	0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x0c,
	0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x22, 0x2e, 0x77,
	0x65, 0x62, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x72,
	0x76, 0x65, 0x79, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x0c, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75, 0x72, 0x76,
	0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x77, 0x65, 0x62, 0x2e,
	0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57,
	0x0a, 0x0c, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x22,
	0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x75, 0x72, 0x76, 0x65, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x0c, 0x53, 0x75, 0x72, 0x76, 0x65,
	0x79, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x73, 0x75,
	0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x77, 0x65,
	0x62, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x72, 0x76,
	0x65, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x54, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x64, 0x65, 0x61, 0x74, 0x65, 0x63, 0x68, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x69, 0x64, 0x65, 0x61,
	0x74, 0x65, 0x63, 0x68, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x77, 0x65, 0x62, 0x2f, 0x73, 0x75, 0x72, 0x76, 0x65,
	0x79, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x77, 0x65, 0x62, 0x5f, 0x73, 0x75, 0x72,
	0x76, 0x65, 0x79, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_web_survey_v1_survey_service_proto_rawDescOnce sync.Once
	file_web_survey_v1_survey_service_proto_rawDescData = file_web_survey_v1_survey_service_proto_rawDesc
)

func file_web_survey_v1_survey_service_proto_rawDescGZIP() []byte {
	file_web_survey_v1_survey_service_proto_rawDescOnce.Do(func() {
		file_web_survey_v1_survey_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_web_survey_v1_survey_service_proto_rawDescData)
	})
	return file_web_survey_v1_survey_service_proto_rawDescData
}

var file_web_survey_v1_survey_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_web_survey_v1_survey_service_proto_goTypes = []interface{}{
	(*SurveyGetOneRequest)(nil),  // 0: web.survey.v1.SurveyGetOneRequest
	(*SurveyGetOneResponse)(nil), // 1: web.survey.v1.SurveyGetOneResponse
	(*SurveyGetAllRequest)(nil),  // 2: web.survey.v1.SurveyGetAllRequest
	(*SurveyGetAllResponse)(nil), // 3: web.survey.v1.SurveyGetAllResponse
	(*SurveyCreateRequest)(nil),  // 4: web.survey.v1.SurveyCreateRequest
	(*SurveyCreateResponse)(nil), // 5: web.survey.v1.SurveyCreateResponse
	(*SurveyUpdateRequest)(nil),  // 6: web.survey.v1.SurveyUpdateRequest
	(*SurveyUpdateResponse)(nil), // 7: web.survey.v1.SurveyUpdateResponse
	(*SurveyDeleteRequest)(nil),  // 8: web.survey.v1.SurveyDeleteRequest
	(*SurveyDeleteResponse)(nil), // 9: web.survey.v1.SurveyDeleteResponse
	(*Survey)(nil),               // 10: web.survey.v1.Survey
	(SurveyType)(0),              // 11: web.survey.v1.SurveyType
}
var file_web_survey_v1_survey_service_proto_depIdxs = []int32{
	10, // 0: web.survey.v1.SurveyGetOneResponse.survey:type_name -> web.survey.v1.Survey
	11, // 1: web.survey.v1.SurveyGetAllRequest.filter_type:type_name -> web.survey.v1.SurveyType
	10, // 2: web.survey.v1.SurveyGetAllResponse.surveys:type_name -> web.survey.v1.Survey
	11, // 3: web.survey.v1.SurveyCreateRequest.type:type_name -> web.survey.v1.SurveyType
	10, // 4: web.survey.v1.SurveyCreateResponse.survey:type_name -> web.survey.v1.Survey
	10, // 5: web.survey.v1.SurveyUpdateResponse.survey:type_name -> web.survey.v1.Survey
	0,  // 6: web.survey.v1.SurveyService.SurveyGetOne:input_type -> web.survey.v1.SurveyGetOneRequest
	2,  // 7: web.survey.v1.SurveyService.SurveyGetAll:input_type -> web.survey.v1.SurveyGetAllRequest
	4,  // 8: web.survey.v1.SurveyService.SurveyCreate:input_type -> web.survey.v1.SurveyCreateRequest
	6,  // 9: web.survey.v1.SurveyService.SurveyUpdate:input_type -> web.survey.v1.SurveyUpdateRequest
	8,  // 10: web.survey.v1.SurveyService.SurveyDelete:input_type -> web.survey.v1.SurveyDeleteRequest
	1,  // 11: web.survey.v1.SurveyService.SurveyGetOne:output_type -> web.survey.v1.SurveyGetOneResponse
	3,  // 12: web.survey.v1.SurveyService.SurveyGetAll:output_type -> web.survey.v1.SurveyGetAllResponse
	5,  // 13: web.survey.v1.SurveyService.SurveyCreate:output_type -> web.survey.v1.SurveyCreateResponse
	7,  // 14: web.survey.v1.SurveyService.SurveyUpdate:output_type -> web.survey.v1.SurveyUpdateResponse
	9,  // 15: web.survey.v1.SurveyService.SurveyDelete:output_type -> web.survey.v1.SurveyDeleteResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_web_survey_v1_survey_service_proto_init() }
func file_web_survey_v1_survey_service_proto_init() {
	if File_web_survey_v1_survey_service_proto != nil {
		return
	}
	file_web_survey_v1_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_web_survey_v1_survey_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SurveyGetOneRequest); i {
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
		file_web_survey_v1_survey_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SurveyGetOneResponse); i {
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
		file_web_survey_v1_survey_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SurveyGetAllRequest); i {
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
		file_web_survey_v1_survey_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SurveyGetAllResponse); i {
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
		file_web_survey_v1_survey_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SurveyCreateRequest); i {
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
		file_web_survey_v1_survey_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SurveyCreateResponse); i {
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
		file_web_survey_v1_survey_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SurveyUpdateRequest); i {
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
		file_web_survey_v1_survey_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SurveyUpdateResponse); i {
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
		file_web_survey_v1_survey_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SurveyDeleteRequest); i {
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
		file_web_survey_v1_survey_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SurveyDeleteResponse); i {
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
			RawDescriptor: file_web_survey_v1_survey_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_web_survey_v1_survey_service_proto_goTypes,
		DependencyIndexes: file_web_survey_v1_survey_service_proto_depIdxs,
		MessageInfos:      file_web_survey_v1_survey_service_proto_msgTypes,
	}.Build()
	File_web_survey_v1_survey_service_proto = out.File
	file_web_survey_v1_survey_service_proto_rawDesc = nil
	file_web_survey_v1_survey_service_proto_goTypes = nil
	file_web_survey_v1_survey_service_proto_depIdxs = nil
}
