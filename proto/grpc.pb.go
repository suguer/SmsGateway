// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: proto/grpc.proto

package proto

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

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID     string `protobuf:"bytes,1,opt,name=AppID,proto3" json:"AppID,omitempty"`
	AppSecret string `protobuf:"bytes,2,opt,name=AppSecret,proto3" json:"AppSecret,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_proto_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *Config) GetAppSecret() string {
	if x != nil {
		return x.AppSecret
	}
	return ""
}

type MessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Platform     string            `protobuf:"bytes,1,opt,name=Platform,proto3" json:"Platform,omitempty"`
	Config       *Config           `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	Content      string            `protobuf:"bytes,3,opt,name=Content,proto3" json:"Content,omitempty"`
	Mobile       string            `protobuf:"bytes,4,opt,name=Mobile,proto3" json:"Mobile,omitempty"`
	SignName     string            `protobuf:"bytes,5,opt,name=SignName,proto3" json:"SignName,omitempty"`
	TemplateCode string            `protobuf:"bytes,6,opt,name=TemplateCode,proto3" json:"TemplateCode,omitempty"`
	Param        map[string]string `protobuf:"bytes,7,rep,name=Param,proto3" json:"Param,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SdkAppId     string            `protobuf:"bytes,8,opt,name=SdkAppId,proto3" json:"SdkAppId,omitempty"`
}

func (x *MessageRequest) Reset() {
	*x = MessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageRequest) ProtoMessage() {}

func (x *MessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_proto_msgTypes[1]
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
	return file_proto_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *MessageRequest) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *MessageRequest) GetConfig() *Config {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *MessageRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *MessageRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *MessageRequest) GetSignName() string {
	if x != nil {
		return x.SignName
	}
	return ""
}

func (x *MessageRequest) GetTemplateCode() string {
	if x != nil {
		return x.TemplateCode
	}
	return ""
}

func (x *MessageRequest) GetParam() map[string]string {
	if x != nil {
		return x.Param
	}
	return nil
}

func (x *MessageRequest) GetSdkAppId() string {
	if x != nil {
		return x.SdkAppId
	}
	return ""
}

// The response message containing the greetings
type MessageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error      string `protobuf:"bytes,1,opt,name=Error,proto3" json:"Error,omitempty"`
	Code       string `protobuf:"bytes,2,opt,name=Code,proto3" json:"Code,omitempty"`
	RequestId  string `protobuf:"bytes,3,opt,name=RequestId,proto3" json:"RequestId,omitempty"`
	BizId      string `protobuf:"bytes,4,opt,name=BizId,proto3" json:"BizId,omitempty"`
	RawContent string `protobuf:"bytes,5,opt,name=RawContent,proto3" json:"RawContent,omitempty"`
}

func (x *MessageReply) Reset() {
	*x = MessageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageReply) ProtoMessage() {}

func (x *MessageReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageReply.ProtoReflect.Descriptor instead.
func (*MessageReply) Descriptor() ([]byte, []int) {
	return file_proto_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *MessageReply) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *MessageReply) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *MessageReply) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *MessageReply) GetBizId() string {
	if x != nil {
		return x.BizId
	}
	return ""
}

func (x *MessageReply) GetRawContent() string {
	if x != nil {
		return x.RawContent
	}
	return ""
}

type TemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Platform      string  `protobuf:"bytes,1,opt,name=Platform,proto3" json:"Platform,omitempty"`
	Config        *Config `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	Name          string  `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Content       string  `protobuf:"bytes,4,opt,name=Content,proto3" json:"Content,omitempty"`
	TemplateType  string  `protobuf:"bytes,5,opt,name=TemplateType,proto3" json:"TemplateType,omitempty"`
	Remark        string  `protobuf:"bytes,6,opt,name=Remark,proto3" json:"Remark,omitempty"`
	International string  `protobuf:"bytes,7,opt,name=International,proto3" json:"International,omitempty"`
}

func (x *TemplateRequest) Reset() {
	*x = TemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateRequest) ProtoMessage() {}

func (x *TemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateRequest.ProtoReflect.Descriptor instead.
func (*TemplateRequest) Descriptor() ([]byte, []int) {
	return file_proto_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *TemplateRequest) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *TemplateRequest) GetConfig() *Config {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *TemplateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TemplateRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *TemplateRequest) GetTemplateType() string {
	if x != nil {
		return x.TemplateType
	}
	return ""
}

func (x *TemplateRequest) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *TemplateRequest) GetInternational() string {
	if x != nil {
		return x.International
	}
	return ""
}

type QuerySmsTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Platform     string  `protobuf:"bytes,1,opt,name=Platform,proto3" json:"Platform,omitempty"`
	Config       *Config `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	TemplateCode string  `protobuf:"bytes,3,opt,name=TemplateCode,proto3" json:"TemplateCode,omitempty"`
}

func (x *QuerySmsTemplateRequest) Reset() {
	*x = QuerySmsTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuerySmsTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuerySmsTemplateRequest) ProtoMessage() {}

func (x *QuerySmsTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuerySmsTemplateRequest.ProtoReflect.Descriptor instead.
func (*QuerySmsTemplateRequest) Descriptor() ([]byte, []int) {
	return file_proto_grpc_proto_rawDescGZIP(), []int{4}
}

func (x *QuerySmsTemplateRequest) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *QuerySmsTemplateRequest) GetConfig() *Config {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *QuerySmsTemplateRequest) GetTemplateCode() string {
	if x != nil {
		return x.TemplateCode
	}
	return ""
}

type TemplateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error        string `protobuf:"bytes,1,opt,name=Error,proto3" json:"Error,omitempty"`
	Code         string `protobuf:"bytes,2,opt,name=Code,proto3" json:"Code,omitempty"`
	RequestId    string `protobuf:"bytes,3,opt,name=RequestId,proto3" json:"RequestId,omitempty"`
	TemplateCode string `protobuf:"bytes,4,opt,name=TemplateCode,proto3" json:"TemplateCode,omitempty"`
	RawContent   string `protobuf:"bytes,5,opt,name=RawContent,proto3" json:"RawContent,omitempty"`
}

func (x *TemplateReply) Reset() {
	*x = TemplateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateReply) ProtoMessage() {}

func (x *TemplateReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateReply.ProtoReflect.Descriptor instead.
func (*TemplateReply) Descriptor() ([]byte, []int) {
	return file_proto_grpc_proto_rawDescGZIP(), []int{5}
}

func (x *TemplateReply) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *TemplateReply) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *TemplateReply) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *TemplateReply) GetTemplateCode() string {
	if x != nil {
		return x.TemplateCode
	}
	return ""
}

func (x *TemplateReply) GetRawContent() string {
	if x != nil {
		return x.RawContent
	}
	return ""
}

type CommonReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error      string `protobuf:"bytes,1,opt,name=Error,proto3" json:"Error,omitempty"`
	Code       string `protobuf:"bytes,2,opt,name=Code,proto3" json:"Code,omitempty"`
	RequestId  string `protobuf:"bytes,3,opt,name=RequestId,proto3" json:"RequestId,omitempty"`
	RawContent string `protobuf:"bytes,4,opt,name=RawContent,proto3" json:"RawContent,omitempty"`
}

func (x *CommonReply) Reset() {
	*x = CommonReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonReply) ProtoMessage() {}

func (x *CommonReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonReply.ProtoReflect.Descriptor instead.
func (*CommonReply) Descriptor() ([]byte, []int) {
	return file_proto_grpc_proto_rawDescGZIP(), []int{6}
}

func (x *CommonReply) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *CommonReply) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CommonReply) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *CommonReply) GetRawContent() string {
	if x != nil {
		return x.RawContent
	}
	return ""
}

var File_proto_grpc_proto protoreflect.FileDescriptor

var file_proto_grpc_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x53, 0x6d, 0x73, 0x42, 0x61, 0x73, 0x65, 0x22, 0x3c, 0x0a, 0x06, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x41,
	0x70, 0x70, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x41, 0x70, 0x70, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0xd7, 0x02, 0x0a, 0x0e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x27, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x53, 0x6d, 0x73, 0x42, 0x61,
	0x73, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x69, 0x67, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x69, 0x67, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x38, 0x0a, 0x05, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x53, 0x6d, 0x73, 0x42, 0x61, 0x73, 0x65, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x1a, 0x0a,
	0x08, 0x53, 0x64, 0x6b, 0x41, 0x70, 0x70, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x53, 0x64, 0x6b, 0x41, 0x70, 0x70, 0x49, 0x64, 0x1a, 0x38, 0x0a, 0x0a, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x8c, 0x01, 0x0a, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x42, 0x69, 0x7a, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x42, 0x69, 0x7a,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x61, 0x77, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x52, 0x61, 0x77, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x22, 0xe6, 0x01, 0x0a, 0x0f, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x12, 0x27, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x53, 0x6d, 0x73, 0x42, 0x61, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52,
	0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x24, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x22, 0x82, 0x01, 0x0a, 0x17,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x6d, 0x73, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x12, 0x27, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x53, 0x6d, 0x73, 0x42, 0x61, 0x73, 0x65, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x22, 0x0a, 0x0c,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x22, 0x9b, 0x01, 0x0a, 0x0d, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x52, 0x61, 0x77, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x52, 0x61, 0x77, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x75,
	0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x61, 0x77, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x52, 0x61, 0x77, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0xe6, 0x01, 0x0a, 0x0a, 0x53, 0x6d, 0x73, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x12, 0x41, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x17, 0x2e, 0x53, 0x6d, 0x73, 0x42, 0x61, 0x73, 0x65, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x53, 0x6d, 0x73, 0x42, 0x61, 0x73, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x6d, 0x73, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x53,
	0x6d, 0x73, 0x42, 0x61, 0x73, 0x65, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x53, 0x6d, 0x73, 0x42, 0x61, 0x73, 0x65,
	0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x4c, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x6d, 0x73, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x53, 0x6d, 0x73, 0x42, 0x61, 0x73, 0x65, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x53, 0x6d, 0x73, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x53, 0x6d, 0x73, 0x42, 0x61, 0x73, 0x65,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x31,
	0x0a, 0x0a, 0x53, 0x6d, 0x73, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x42, 0x0f, 0x53, 0x6d,
	0x73, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x10, 0x53, 0x6d, 0x73, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_grpc_proto_rawDescOnce sync.Once
	file_proto_grpc_proto_rawDescData = file_proto_grpc_proto_rawDesc
)

func file_proto_grpc_proto_rawDescGZIP() []byte {
	file_proto_grpc_proto_rawDescOnce.Do(func() {
		file_proto_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_grpc_proto_rawDescData)
	})
	return file_proto_grpc_proto_rawDescData
}

var file_proto_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_grpc_proto_goTypes = []interface{}{
	(*Config)(nil),                  // 0: SmsBase.Config
	(*MessageRequest)(nil),          // 1: SmsBase.MessageRequest
	(*MessageReply)(nil),            // 2: SmsBase.MessageReply
	(*TemplateRequest)(nil),         // 3: SmsBase.TemplateRequest
	(*QuerySmsTemplateRequest)(nil), // 4: SmsBase.QuerySmsTemplateRequest
	(*TemplateReply)(nil),           // 5: SmsBase.TemplateReply
	(*CommonReply)(nil),             // 6: SmsBase.CommonReply
	nil,                             // 7: SmsBase.MessageRequest.ParamEntry
}
var file_proto_grpc_proto_depIdxs = []int32{
	0, // 0: SmsBase.MessageRequest.config:type_name -> SmsBase.Config
	7, // 1: SmsBase.MessageRequest.Param:type_name -> SmsBase.MessageRequest.ParamEntry
	0, // 2: SmsBase.TemplateRequest.config:type_name -> SmsBase.Config
	0, // 3: SmsBase.QuerySmsTemplateRequest.config:type_name -> SmsBase.Config
	1, // 4: SmsBase.SmsGateway.CreateMessage:input_type -> SmsBase.MessageRequest
	3, // 5: SmsBase.SmsGateway.CreateSmsTemplate:input_type -> SmsBase.TemplateRequest
	4, // 6: SmsBase.SmsGateway.QuerySmsTemplate:input_type -> SmsBase.QuerySmsTemplateRequest
	2, // 7: SmsBase.SmsGateway.CreateMessage:output_type -> SmsBase.MessageReply
	5, // 8: SmsBase.SmsGateway.CreateSmsTemplate:output_type -> SmsBase.TemplateReply
	6, // 9: SmsBase.SmsGateway.QuerySmsTemplate:output_type -> SmsBase.CommonReply
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_grpc_proto_init() }
func file_proto_grpc_proto_init() {
	if File_proto_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_proto_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageReply); i {
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
		file_proto_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateRequest); i {
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
		file_proto_grpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuerySmsTemplateRequest); i {
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
		file_proto_grpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateReply); i {
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
		file_proto_grpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonReply); i {
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
			RawDescriptor: file_proto_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_grpc_proto_goTypes,
		DependencyIndexes: file_proto_grpc_proto_depIdxs,
		MessageInfos:      file_proto_grpc_proto_msgTypes,
	}.Build()
	File_proto_grpc_proto = out.File
	file_proto_grpc_proto_rawDesc = nil
	file_proto_grpc_proto_goTypes = nil
	file_proto_grpc_proto_depIdxs = nil
}
