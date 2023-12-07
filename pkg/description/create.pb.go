// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: pbf/description/create.proto

package description

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

// CreateI is the input for creating descriptions.
//
//	{
//	    "object": [
//	        {
//	            "public": {
//	                "evnt": "778237",
//	                "text": "In this Hackathon we build awesome stuff!"
//	            }
//	        }
//	    ]
//	}
type CreateI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *CreateI_Filter   `protobuf:"bytes,100,opt,name=filter,proto3" json:"filter,omitempty"`
	Object []*CreateI_Object `protobuf:"bytes,200,rep,name=object,proto3" json:"object,omitempty"`
}

func (x *CreateI) Reset() {
	*x = CreateI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_description_create_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateI) ProtoMessage() {}

func (x *CreateI) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_description_create_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateI.ProtoReflect.Descriptor instead.
func (*CreateI) Descriptor() ([]byte, []int) {
	return file_pbf_description_create_proto_rawDescGZIP(), []int{0}
}

func (x *CreateI) GetFilter() *CreateI_Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *CreateI) GetObject() []*CreateI_Object {
	if x != nil {
		return x.Object
	}
	return nil
}

type CreateI_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateI_Filter) Reset() {
	*x = CreateI_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_description_create_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateI_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateI_Filter) ProtoMessage() {}

func (x *CreateI_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_description_create_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateI_Filter.ProtoReflect.Descriptor instead.
func (*CreateI_Filter) Descriptor() ([]byte, []int) {
	return file_pbf_description_create_proto_rawDescGZIP(), []int{1}
}

type CreateI_Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Intern *CreateI_Object_Intern `protobuf:"bytes,100,opt,name=intern,proto3" json:"intern,omitempty"`
	Public *CreateI_Object_Public `protobuf:"bytes,200,opt,name=public,proto3" json:"public,omitempty"`
}

func (x *CreateI_Object) Reset() {
	*x = CreateI_Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_description_create_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateI_Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateI_Object) ProtoMessage() {}

func (x *CreateI_Object) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_description_create_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateI_Object.ProtoReflect.Descriptor instead.
func (*CreateI_Object) Descriptor() ([]byte, []int) {
	return file_pbf_description_create_proto_rawDescGZIP(), []int{2}
}

func (x *CreateI_Object) GetIntern() *CreateI_Object_Intern {
	if x != nil {
		return x.Intern
	}
	return nil
}

func (x *CreateI_Object) GetPublic() *CreateI_Object_Public {
	if x != nil {
		return x.Public
	}
	return nil
}

type CreateI_Object_Intern struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateI_Object_Intern) Reset() {
	*x = CreateI_Object_Intern{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_description_create_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateI_Object_Intern) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateI_Object_Intern) ProtoMessage() {}

func (x *CreateI_Object_Intern) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_description_create_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateI_Object_Intern.ProtoReflect.Descriptor instead.
func (*CreateI_Object_Intern) Descriptor() ([]byte, []int) {
	return file_pbf_description_create_proto_rawDescGZIP(), []int{3}
}

type CreateI_Object_Public struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// evnt is the event ID this description is mapped to.
	Evnt string `protobuf:"bytes,100,opt,name=evnt,proto3" json:"evnt,omitempty"`
	// text is the description explaining what this event is about.
	Text string `protobuf:"bytes,200,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *CreateI_Object_Public) Reset() {
	*x = CreateI_Object_Public{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_description_create_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateI_Object_Public) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateI_Object_Public) ProtoMessage() {}

func (x *CreateI_Object_Public) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_description_create_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateI_Object_Public.ProtoReflect.Descriptor instead.
func (*CreateI_Object_Public) Descriptor() ([]byte, []int) {
	return file_pbf_description_create_proto_rawDescGZIP(), []int{4}
}

func (x *CreateI_Object_Public) GetEvnt() string {
	if x != nil {
		return x.Evnt
	}
	return ""
}

func (x *CreateI_Object_Public) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// CreateO is the output for creating descriptions.
//
//	{
//	    "object": [
//	        {
//	            "intern": {
//	                "crtd": "1689001255",
//	                "desc": "98972893"
//	            }
//	        }
//	    ]
//	}
type CreateO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *CreateO_Filter   `protobuf:"bytes,100,opt,name=filter,proto3" json:"filter,omitempty"`
	Object []*CreateO_Object `protobuf:"bytes,200,rep,name=object,proto3" json:"object,omitempty"`
}

func (x *CreateO) Reset() {
	*x = CreateO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_description_create_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateO) ProtoMessage() {}

func (x *CreateO) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_description_create_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateO.ProtoReflect.Descriptor instead.
func (*CreateO) Descriptor() ([]byte, []int) {
	return file_pbf_description_create_proto_rawDescGZIP(), []int{5}
}

func (x *CreateO) GetFilter() *CreateO_Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *CreateO) GetObject() []*CreateO_Object {
	if x != nil {
		return x.Object
	}
	return nil
}

type CreateO_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateO_Filter) Reset() {
	*x = CreateO_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_description_create_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateO_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateO_Filter) ProtoMessage() {}

func (x *CreateO_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_description_create_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateO_Filter.ProtoReflect.Descriptor instead.
func (*CreateO_Filter) Descriptor() ([]byte, []int) {
	return file_pbf_description_create_proto_rawDescGZIP(), []int{6}
}

type CreateO_Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Intern *CreateO_Object_Intern `protobuf:"bytes,100,opt,name=intern,proto3" json:"intern,omitempty"`
	Public *CreateO_Object_Public `protobuf:"bytes,200,opt,name=public,proto3" json:"public,omitempty"`
}

func (x *CreateO_Object) Reset() {
	*x = CreateO_Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_description_create_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateO_Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateO_Object) ProtoMessage() {}

func (x *CreateO_Object) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_description_create_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateO_Object.ProtoReflect.Descriptor instead.
func (*CreateO_Object) Descriptor() ([]byte, []int) {
	return file_pbf_description_create_proto_rawDescGZIP(), []int{7}
}

func (x *CreateO_Object) GetIntern() *CreateO_Object_Intern {
	if x != nil {
		return x.Intern
	}
	return nil
}

func (x *CreateO_Object) GetPublic() *CreateO_Object_Public {
	if x != nil {
		return x.Public
	}
	return nil
}

type CreateO_Object_Intern struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// crtd is the unix timestamp in seconds at which the description got created.
	Crtd string `protobuf:"bytes,100,opt,name=crtd,proto3" json:"crtd,omitempty"`
	// desc is the ID of the description being created.
	Desc string `protobuf:"bytes,200,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *CreateO_Object_Intern) Reset() {
	*x = CreateO_Object_Intern{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_description_create_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateO_Object_Intern) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateO_Object_Intern) ProtoMessage() {}

func (x *CreateO_Object_Intern) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_description_create_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateO_Object_Intern.ProtoReflect.Descriptor instead.
func (*CreateO_Object_Intern) Descriptor() ([]byte, []int) {
	return file_pbf_description_create_proto_rawDescGZIP(), []int{8}
}

func (x *CreateO_Object_Intern) GetCrtd() string {
	if x != nil {
		return x.Crtd
	}
	return ""
}

func (x *CreateO_Object_Intern) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type CreateO_Object_Public struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateO_Object_Public) Reset() {
	*x = CreateO_Object_Public{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_description_create_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateO_Object_Public) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateO_Object_Public) ProtoMessage() {}

func (x *CreateO_Object_Public) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_description_create_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateO_Object_Public.ProtoReflect.Descriptor instead.
func (*CreateO_Object_Public) Descriptor() ([]byte, []int) {
	return file_pbf_description_create_proto_rawDescGZIP(), []int{9}
}

var File_pbf_description_create_proto protoreflect.FileDescriptor

var file_pbf_description_create_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x62, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x74, 0x0a, 0x07, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x12, 0x33, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x5f, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x06, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0xc8, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x22, 0x10, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x5f, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x22, 0x89, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x5f,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x3a, 0x0a, 0x06, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x5f, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x52, 0x06, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x12, 0x3b, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0xc8, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x22,
	0x17, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x22, 0x40, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x76, 0x6e, 0x74, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x65, 0x76, 0x6e, 0x74, 0x12, 0x13, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0xc8, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x74, 0x0a, 0x07, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4f, 0x12, 0x33, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18,
	0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x5f, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x06, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0xc8, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x22, 0x10, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x5f, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x22, 0x89, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x5f, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x3a, 0x0a, 0x06, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x18,
	0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x52, 0x06, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x12, 0x3b, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0xc8, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x22, 0x40,
	0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x72, 0x74, 0x64, 0x18,
	0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x72, 0x74, 0x64, 0x12, 0x13, 0x0a, 0x04, 0x64,
	0x65, 0x73, 0x63, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63,
	0x22, 0x17, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x3b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pbf_description_create_proto_rawDescOnce sync.Once
	file_pbf_description_create_proto_rawDescData = file_pbf_description_create_proto_rawDesc
)

func file_pbf_description_create_proto_rawDescGZIP() []byte {
	file_pbf_description_create_proto_rawDescOnce.Do(func() {
		file_pbf_description_create_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbf_description_create_proto_rawDescData)
	})
	return file_pbf_description_create_proto_rawDescData
}

var file_pbf_description_create_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pbf_description_create_proto_goTypes = []interface{}{
	(*CreateI)(nil),               // 0: description.CreateI
	(*CreateI_Filter)(nil),        // 1: description.CreateI_Filter
	(*CreateI_Object)(nil),        // 2: description.CreateI_Object
	(*CreateI_Object_Intern)(nil), // 3: description.CreateI_Object_Intern
	(*CreateI_Object_Public)(nil), // 4: description.CreateI_Object_Public
	(*CreateO)(nil),               // 5: description.CreateO
	(*CreateO_Filter)(nil),        // 6: description.CreateO_Filter
	(*CreateO_Object)(nil),        // 7: description.CreateO_Object
	(*CreateO_Object_Intern)(nil), // 8: description.CreateO_Object_Intern
	(*CreateO_Object_Public)(nil), // 9: description.CreateO_Object_Public
}
var file_pbf_description_create_proto_depIdxs = []int32{
	1, // 0: description.CreateI.filter:type_name -> description.CreateI_Filter
	2, // 1: description.CreateI.object:type_name -> description.CreateI_Object
	3, // 2: description.CreateI_Object.intern:type_name -> description.CreateI_Object_Intern
	4, // 3: description.CreateI_Object.public:type_name -> description.CreateI_Object_Public
	6, // 4: description.CreateO.filter:type_name -> description.CreateO_Filter
	7, // 5: description.CreateO.object:type_name -> description.CreateO_Object
	8, // 6: description.CreateO_Object.intern:type_name -> description.CreateO_Object_Intern
	9, // 7: description.CreateO_Object.public:type_name -> description.CreateO_Object_Public
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_pbf_description_create_proto_init() }
func file_pbf_description_create_proto_init() {
	if File_pbf_description_create_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbf_description_create_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateI); i {
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
		file_pbf_description_create_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateI_Filter); i {
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
		file_pbf_description_create_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateI_Object); i {
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
		file_pbf_description_create_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateI_Object_Intern); i {
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
		file_pbf_description_create_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateI_Object_Public); i {
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
		file_pbf_description_create_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateO); i {
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
		file_pbf_description_create_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateO_Filter); i {
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
		file_pbf_description_create_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateO_Object); i {
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
		file_pbf_description_create_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateO_Object_Intern); i {
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
		file_pbf_description_create_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateO_Object_Public); i {
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
			RawDescriptor: file_pbf_description_create_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbf_description_create_proto_goTypes,
		DependencyIndexes: file_pbf_description_create_proto_depIdxs,
		MessageInfos:      file_pbf_description_create_proto_msgTypes,
	}.Build()
	File_pbf_description_create_proto = out.File
	file_pbf_description_create_proto_rawDesc = nil
	file_pbf_description_create_proto_goTypes = nil
	file_pbf_description_create_proto_depIdxs = nil
}
