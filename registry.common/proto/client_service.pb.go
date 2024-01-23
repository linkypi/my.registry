// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: client_service.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SubType int32

const (
	SubType_NoneSubPub  SubType = 0
	SubType_Subscribe   SubType = 1
	SubType_UnSubscribe SubType = 2
)

// Enum value maps for SubType.
var (
	SubType_name = map[int32]string{
		0: "NoneSubPub",
		1: "Subscribe",
		2: "UnSubscribe",
	}
	SubType_value = map[string]int32{
		"NoneSubPub":  0,
		"Subscribe":   1,
		"UnSubscribe": 2,
	}
)

func (x SubType) Enum() *SubType {
	p := new(SubType)
	*p = x
	return p
}

func (x SubType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SubType) Descriptor() protoreflect.EnumDescriptor {
	return file_client_service_proto_enumTypes[0].Descriptor()
}

func (SubType) Type() protoreflect.EnumType {
	return &file_client_service_proto_enumTypes[0]
}

func (x SubType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SubType.Descriptor instead.
func (SubType) EnumDescriptor() ([]byte, []int) {
	return file_client_service_proto_rawDescGZIP(), []int{0}
}

type FetchMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shards    string            `protobuf:"bytes,1,opt,name=shards,proto3" json:"shards,omitempty"`
	Nodes     map[string]string `protobuf:"bytes,2,rep,name=nodes,proto3" json:"nodes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ErrorType ErrorType         `protobuf:"varint,3,opt,name=errorType,proto3,enum=ErrorType" json:"errorType,omitempty"`
}

func (x *FetchMetadataResponse) Reset() {
	*x = FetchMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchMetadataResponse) ProtoMessage() {}

func (x *FetchMetadataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchMetadataResponse.ProtoReflect.Descriptor instead.
func (*FetchMetadataResponse) Descriptor() ([]byte, []int) {
	return file_client_service_proto_rawDescGZIP(), []int{0}
}

func (x *FetchMetadataResponse) GetShards() string {
	if x != nil {
		return x.Shards
	}
	return ""
}

func (x *FetchMetadataResponse) GetNodes() map[string]string {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *FetchMetadataResponse) GetErrorType() ErrorType {
	if x != nil {
		return x.ErrorType
	}
	return ErrorType_None
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName string `protobuf:"bytes,1,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	ServiceIp   string `protobuf:"bytes,2,opt,name=serviceIp,proto3" json:"serviceIp,omitempty"`
	ServicePort int32  `protobuf:"varint,3,opt,name=servicePort,proto3" json:"servicePort,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_client_service_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *RegisterRequest) GetServiceIp() string {
	if x != nil {
		return x.ServiceIp
	}
	return ""
}

func (x *RegisterRequest) GetServicePort() int32 {
	if x != nil {
		return x.ServicePort
	}
	return 0
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorType ErrorType `protobuf:"varint,1,opt,name=errorType,proto3,enum=ErrorType" json:"errorType,omitempty"`
	Success   bool      `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_client_service_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterResponse) GetErrorType() ErrorType {
	if x != nil {
		return x.ErrorType
	}
	return ErrorType_None
}

func (x *RegisterResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type SubRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubType     SubType `protobuf:"varint,1,opt,name=subType,proto3,enum=SubType" json:"subType,omitempty"`
	ServiceName string  `protobuf:"bytes,2,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	ServiceAddr string  `protobuf:"bytes,3,opt,name=serviceAddr,proto3" json:"serviceAddr,omitempty"`
}

func (x *SubRequest) Reset() {
	*x = SubRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubRequest) ProtoMessage() {}

func (x *SubRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubRequest.ProtoReflect.Descriptor instead.
func (*SubRequest) Descriptor() ([]byte, []int) {
	return file_client_service_proto_rawDescGZIP(), []int{3}
}

func (x *SubRequest) GetSubType() SubType {
	if x != nil {
		return x.SubType
	}
	return SubType_NoneSubPub
}

func (x *SubRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *SubRequest) GetServiceAddr() string {
	if x != nil {
		return x.ServiceAddr
	}
	return ""
}

type ServiceChangedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName string `protobuf:"bytes,1,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
}

func (x *ServiceChangedRequest) Reset() {
	*x = ServiceChangedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceChangedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceChangedRequest) ProtoMessage() {}

func (x *ServiceChangedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceChangedRequest.ProtoReflect.Descriptor instead.
func (*ServiceChangedRequest) Descriptor() ([]byte, []int) {
	return file_client_service_proto_rawDescGZIP(), []int{4}
}

func (x *ServiceChangedRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

type FetchServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName      string             `protobuf:"bytes,1,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	ServiceInstances []*ServiceInstance `protobuf:"bytes,2,rep,name=serviceInstances,proto3" json:"serviceInstances,omitempty"`
}

func (x *FetchServiceResponse) Reset() {
	*x = FetchServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchServiceResponse) ProtoMessage() {}

func (x *FetchServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchServiceResponse.ProtoReflect.Descriptor instead.
func (*FetchServiceResponse) Descriptor() ([]byte, []int) {
	return file_client_service_proto_rawDescGZIP(), []int{5}
}

func (x *FetchServiceResponse) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *FetchServiceResponse) GetServiceInstances() []*ServiceInstance {
	if x != nil {
		return x.ServiceInstances
	}
	return nil
}

type FetchServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName string `protobuf:"bytes,1,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
}

func (x *FetchServiceRequest) Reset() {
	*x = FetchServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchServiceRequest) ProtoMessage() {}

func (x *FetchServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchServiceRequest.ProtoReflect.Descriptor instead.
func (*FetchServiceRequest) Descriptor() ([]byte, []int) {
	return file_client_service_proto_rawDescGZIP(), []int{6}
}

func (x *FetchServiceRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

type ServiceInstance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName string `protobuf:"bytes,1,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	ServiceIp   string `protobuf:"bytes,2,opt,name=serviceIp,proto3" json:"serviceIp,omitempty"`
	ServicePort int32  `protobuf:"varint,3,opt,name=servicePort,proto3" json:"servicePort,omitempty"`
}

func (x *ServiceInstance) Reset() {
	*x = ServiceInstance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceInstance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceInstance) ProtoMessage() {}

func (x *ServiceInstance) ProtoReflect() protoreflect.Message {
	mi := &file_client_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceInstance.ProtoReflect.Descriptor instead.
func (*ServiceInstance) Descriptor() ([]byte, []int) {
	return file_client_service_proto_rawDescGZIP(), []int{7}
}

func (x *ServiceInstance) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *ServiceInstance) GetServiceIp() string {
	if x != nil {
		return x.ServiceIp
	}
	return ""
}

func (x *ServiceInstance) GetServicePort() int32 {
	if x != nil {
		return x.ServicePort
	}
	return 0
}

type HeartbeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName string `protobuf:"bytes,1,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	ServiceIp   string `protobuf:"bytes,2,opt,name=serviceIp,proto3" json:"serviceIp,omitempty"`
	ServicePort int32  `protobuf:"varint,3,opt,name=servicePort,proto3" json:"servicePort,omitempty"`
}

func (x *HeartbeatRequest) Reset() {
	*x = HeartbeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatRequest) ProtoMessage() {}

func (x *HeartbeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatRequest.ProtoReflect.Descriptor instead.
func (*HeartbeatRequest) Descriptor() ([]byte, []int) {
	return file_client_service_proto_rawDescGZIP(), []int{8}
}

func (x *HeartbeatRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *HeartbeatRequest) GetServiceIp() string {
	if x != nil {
		return x.ServiceIp
	}
	return ""
}

func (x *HeartbeatRequest) GetServicePort() int32 {
	if x != nil {
		return x.ServicePort
	}
	return 0
}

var File_client_service_proto protoreflect.FileDescriptor

var file_client_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x01, 0x0a, 0x15, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x68, 0x61, 0x72, 0x64, 0x73, 0x12, 0x37, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73,
	0x12, 0x28, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x38, 0x0a, 0x0a, 0x4e, 0x6f,
	0x64, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x73, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x56, 0x0a, 0x10, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a,
	0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0a, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x22, 0x74, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x22, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x08, 0x2e, 0x53, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x73, 0x75, 0x62, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x41, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x22, 0x39, 0x0a, 0x15, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0x76, 0x0a, 0x14, 0x46, 0x65, 0x74, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x10,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x22, 0x37, 0x0a, 0x13, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x73, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x74, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x70, 0x12, 0x20, 0x0a, 0x0b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x2a, 0x39,
	0x0a, 0x07, 0x53, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x6f, 0x6e,
	0x65, 0x53, 0x75, 0x62, 0x50, 0x75, 0x62, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x6e, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x10, 0x02, 0x32, 0x8a, 0x03, 0x0a, 0x0d, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32,
	0x0a, 0x09, 0x53, 0x75, 0x62, 0x53, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x0b, 0x2e, 0x53, 0x75,
	0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x49, 0x0a, 0x15, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x16, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x46, 0x0a,
	0x15, 0x46, 0x65, 0x74, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x14, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0d, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16,
	0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x64,
	0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x11, 0x2e, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_service_proto_rawDescOnce sync.Once
	file_client_service_proto_rawDescData = file_client_service_proto_rawDesc
)

func file_client_service_proto_rawDescGZIP() []byte {
	file_client_service_proto_rawDescOnce.Do(func() {
		file_client_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_service_proto_rawDescData)
	})
	return file_client_service_proto_rawDescData
}

var file_client_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_client_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_client_service_proto_goTypes = []interface{}{
	(SubType)(0),                  // 0: SubType
	(*FetchMetadataResponse)(nil), // 1: FetchMetadataResponse
	(*RegisterRequest)(nil),       // 2: RegisterRequest
	(*RegisterResponse)(nil),      // 3: RegisterResponse
	(*SubRequest)(nil),            // 4: SubRequest
	(*ServiceChangedRequest)(nil), // 5: ServiceChangedRequest
	(*FetchServiceResponse)(nil),  // 6: FetchServiceResponse
	(*FetchServiceRequest)(nil),   // 7: FetchServiceRequest
	(*ServiceInstance)(nil),       // 8: ServiceInstance
	(*HeartbeatRequest)(nil),      // 9: HeartbeatRequest
	nil,                           // 10: FetchMetadataResponse.NodesEntry
	(ErrorType)(0),                // 11: ErrorType
	(*emptypb.Empty)(nil),         // 12: google.protobuf.Empty
}
var file_client_service_proto_depIdxs = []int32{
	10, // 0: FetchMetadataResponse.nodes:type_name -> FetchMetadataResponse.NodesEntry
	11, // 1: FetchMetadataResponse.errorType:type_name -> ErrorType
	11, // 2: RegisterResponse.errorType:type_name -> ErrorType
	0,  // 3: SubRequest.subType:type_name -> SubType
	8,  // 4: FetchServiceResponse.serviceInstances:type_name -> ServiceInstance
	2,  // 5: ClientService.Register:input_type -> RegisterRequest
	4,  // 6: ClientService.SubScribe:input_type -> SubRequest
	5,  // 7: ClientService.PublishServiceChanged:input_type -> ServiceChangedRequest
	7,  // 8: ClientService.FetchServiceInstances:input_type -> FetchServiceRequest
	12, // 9: ClientService.FetchMetadata:input_type -> google.protobuf.Empty
	9,  // 10: ClientService.SendHeartbeat:input_type -> HeartbeatRequest
	3,  // 11: ClientService.Register:output_type -> RegisterResponse
	12, // 12: ClientService.SubScribe:output_type -> google.protobuf.Empty
	12, // 13: ClientService.PublishServiceChanged:output_type -> google.protobuf.Empty
	6,  // 14: ClientService.FetchServiceInstances:output_type -> FetchServiceResponse
	1,  // 15: ClientService.FetchMetadata:output_type -> FetchMetadataResponse
	12, // 16: ClientService.SendHeartbeat:output_type -> google.protobuf.Empty
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_client_service_proto_init() }
func file_client_service_proto_init() {
	if File_client_service_proto != nil {
		return
	}
	file_error_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_client_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchMetadataResponse); i {
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
		file_client_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_client_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_client_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubRequest); i {
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
		file_client_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceChangedRequest); i {
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
		file_client_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchServiceResponse); i {
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
		file_client_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchServiceRequest); i {
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
		file_client_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceInstance); i {
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
		file_client_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatRequest); i {
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
			RawDescriptor: file_client_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_client_service_proto_goTypes,
		DependencyIndexes: file_client_service_proto_depIdxs,
		EnumInfos:         file_client_service_proto_enumTypes,
		MessageInfos:      file_client_service_proto_msgTypes,
	}.Build()
	File_client_service_proto = out.File
	file_client_service_proto_rawDesc = nil
	file_client_service_proto_goTypes = nil
	file_client_service_proto_depIdxs = nil
}