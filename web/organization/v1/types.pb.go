// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.2
// source: web/organization/v1/types.proto

package grpcweb_organization_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Organization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DatetimeCreated *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=datetime_created,json=datetimeCreated,proto3" json:"datetime_created,omitempty"`
	DatetimeUpdated *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=datetime_updated,json=datetimeUpdated,proto3" json:"datetime_updated,omitempty"`
	Name            string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Organization) Reset() {
	*x = Organization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_organization_v1_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Organization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Organization) ProtoMessage() {}

func (x *Organization) ProtoReflect() protoreflect.Message {
	mi := &file_web_organization_v1_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Organization.ProtoReflect.Descriptor instead.
func (*Organization) Descriptor() ([]byte, []int) {
	return file_web_organization_v1_types_proto_rawDescGZIP(), []int{0}
}

func (x *Organization) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Organization) GetDatetimeCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.DatetimeCreated
	}
	return nil
}

func (x *Organization) GetDatetimeUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.DatetimeUpdated
	}
	return nil
}

func (x *Organization) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_web_organization_v1_types_proto protoreflect.FileDescriptor

var file_web_organization_v1_types_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x77, 0x65, 0x62, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x77, 0x65, 0x62, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x01, 0x0a, 0x0c, 0x4f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x45, 0x0a, 0x10, 0x64, 0x61, 0x74, 0x65,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f,
	0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x45, 0x0a, 0x10, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x60, 0x5a, 0x5e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x64, 0x65, 0x61, 0x74, 0x65, 0x63,
	0x68, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x69, 0x64, 0x65, 0x61, 0x74, 0x65, 0x63, 0x68, 0x2d, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f,
	0x2f, 0x77, 0x65, 0x62, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x77, 0x65, 0x62, 0x5f, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_web_organization_v1_types_proto_rawDescOnce sync.Once
	file_web_organization_v1_types_proto_rawDescData = file_web_organization_v1_types_proto_rawDesc
)

func file_web_organization_v1_types_proto_rawDescGZIP() []byte {
	file_web_organization_v1_types_proto_rawDescOnce.Do(func() {
		file_web_organization_v1_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_web_organization_v1_types_proto_rawDescData)
	})
	return file_web_organization_v1_types_proto_rawDescData
}

var file_web_organization_v1_types_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_web_organization_v1_types_proto_goTypes = []interface{}{
	(*Organization)(nil),          // 0: web.organization.v1.Organization
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_web_organization_v1_types_proto_depIdxs = []int32{
	1, // 0: web.organization.v1.Organization.datetime_created:type_name -> google.protobuf.Timestamp
	1, // 1: web.organization.v1.Organization.datetime_updated:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_web_organization_v1_types_proto_init() }
func file_web_organization_v1_types_proto_init() {
	if File_web_organization_v1_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_web_organization_v1_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Organization); i {
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
			RawDescriptor: file_web_organization_v1_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_web_organization_v1_types_proto_goTypes,
		DependencyIndexes: file_web_organization_v1_types_proto_depIdxs,
		MessageInfos:      file_web_organization_v1_types_proto_msgTypes,
	}.Build()
	File_web_organization_v1_types_proto = out.File
	file_web_organization_v1_types_proto_rawDesc = nil
	file_web_organization_v1_types_proto_goTypes = nil
	file_web_organization_v1_types_proto_depIdxs = nil
}
