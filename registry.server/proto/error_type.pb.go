// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: error_type.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ErrorType int32

const (
	ErrorType_None                 ErrorType = 0
	ErrorType_ClusterDown          ErrorType = 1
	ErrorType_ClusterStateNotMatch ErrorType = 2
	ErrorType_ClusterIdNotMatch    ErrorType = 3
	ErrorType_LeaderIdNotMatch     ErrorType = 4
	ErrorType_TermNotMatch         ErrorType = 5
	ErrorType_MetaDataChanged      ErrorType = 6
	ErrorType_UnknownError         ErrorType = 7
)

// Enum value maps for ErrorType.
var (
	ErrorType_name = map[int32]string{
		0: "None",
		1: "ClusterDown",
		2: "ClusterStateNotMatch",
		3: "ClusterIdNotMatch",
		4: "LeaderIdNotMatch",
		5: "TermNotMatch",
		6: "MetaDataChanged",
		7: "UnknownError",
	}
	ErrorType_value = map[string]int32{
		"None":                 0,
		"ClusterDown":          1,
		"ClusterStateNotMatch": 2,
		"ClusterIdNotMatch":    3,
		"LeaderIdNotMatch":     4,
		"TermNotMatch":         5,
		"MetaDataChanged":      6,
		"UnknownError":         7,
	}
)

func (x ErrorType) Enum() *ErrorType {
	p := new(ErrorType)
	*p = x
	return p
}

func (x ErrorType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorType) Descriptor() protoreflect.EnumDescriptor {
	return file_error_type_proto_enumTypes[0].Descriptor()
}

func (ErrorType) Type() protoreflect.EnumType {
	return &file_error_type_proto_enumTypes[0]
}

func (x ErrorType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorType.Descriptor instead.
func (ErrorType) EnumDescriptor() ([]byte, []int) {
	return file_error_type_proto_rawDescGZIP(), []int{0}
}

var File_error_type_proto protoreflect.FileDescriptor

var file_error_type_proto_rawDesc = []byte{
	0x0a, 0x10, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a,
	0xa6, 0x01, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a,
	0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x44, 0x6f, 0x77, 0x6e, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x4e,
	0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x4c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x4e, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x10, 0x04, 0x12,
	0x10, 0x0a, 0x0c, 0x54, 0x65, 0x72, 0x6d, 0x4e, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x10,
	0x05, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x64, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77,
	0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x07, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_error_type_proto_rawDescOnce sync.Once
	file_error_type_proto_rawDescData = file_error_type_proto_rawDesc
)

func file_error_type_proto_rawDescGZIP() []byte {
	file_error_type_proto_rawDescOnce.Do(func() {
		file_error_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_error_type_proto_rawDescData)
	})
	return file_error_type_proto_rawDescData
}

var file_error_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_error_type_proto_goTypes = []interface{}{
	(ErrorType)(0), // 0: ErrorType
}
var file_error_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_error_type_proto_init() }
func file_error_type_proto_init() {
	if File_error_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_error_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_error_type_proto_goTypes,
		DependencyIndexes: file_error_type_proto_depIdxs,
		EnumInfos:         file_error_type_proto_enumTypes,
	}.Build()
	File_error_type_proto = out.File
	file_error_type_proto_rawDesc = nil
	file_error_type_proto_goTypes = nil
	file_error_type_proto_depIdxs = nil
}
