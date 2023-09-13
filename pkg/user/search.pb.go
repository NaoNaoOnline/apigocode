// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: pbf/user/search.proto

package user

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

// SearchI is the input for searching users.
//
//	{
//	    "filter": [
//	        "chunking": {
//	            "perpage": "50",
//	            "pointer": "100"
//	        }
//	    ],
//	    "object": [
//	        {
//	            "intern": {
//	                "user": "551265"
//	            }
//	        }
//	    ]
//	}
type SearchI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *SearchI_Filter   `protobuf:"bytes,100,opt,name=filter,proto3" json:"filter,omitempty"`
	Object []*SearchI_Object `protobuf:"bytes,200,rep,name=object,proto3" json:"object,omitempty"`
}

func (x *SearchI) Reset() {
	*x = SearchI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_user_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchI) ProtoMessage() {}

func (x *SearchI) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_user_search_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchI.ProtoReflect.Descriptor instead.
func (*SearchI) Descriptor() ([]byte, []int) {
	return file_pbf_user_search_proto_rawDescGZIP(), []int{0}
}

func (x *SearchI) GetFilter() *SearchI_Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *SearchI) GetObject() []*SearchI_Object {
	if x != nil {
		return x.Object
	}
	return nil
}

type SearchI_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunking *SearchI_Filter_Chunking `protobuf:"bytes,100,opt,name=chunking,proto3" json:"chunking,omitempty"`
}

func (x *SearchI_Filter) Reset() {
	*x = SearchI_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_user_search_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchI_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchI_Filter) ProtoMessage() {}

func (x *SearchI_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_user_search_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchI_Filter.ProtoReflect.Descriptor instead.
func (*SearchI_Filter) Descriptor() ([]byte, []int) {
	return file_pbf_user_search_proto_rawDescGZIP(), []int{1}
}

func (x *SearchI_Filter) GetChunking() *SearchI_Filter_Chunking {
	if x != nil {
		return x.Chunking
	}
	return nil
}

type SearchI_Filter_Chunking struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Perpage string `protobuf:"bytes,100,opt,name=perpage,proto3" json:"perpage,omitempty"`
	Pointer string `protobuf:"bytes,200,opt,name=pointer,proto3" json:"pointer,omitempty"`
}

func (x *SearchI_Filter_Chunking) Reset() {
	*x = SearchI_Filter_Chunking{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_user_search_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchI_Filter_Chunking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchI_Filter_Chunking) ProtoMessage() {}

func (x *SearchI_Filter_Chunking) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_user_search_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchI_Filter_Chunking.ProtoReflect.Descriptor instead.
func (*SearchI_Filter_Chunking) Descriptor() ([]byte, []int) {
	return file_pbf_user_search_proto_rawDescGZIP(), []int{2}
}

func (x *SearchI_Filter_Chunking) GetPerpage() string {
	if x != nil {
		return x.Perpage
	}
	return ""
}

func (x *SearchI_Filter_Chunking) GetPointer() string {
	if x != nil {
		return x.Pointer
	}
	return ""
}

type SearchI_Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Intern *SearchI_Object_Intern `protobuf:"bytes,100,opt,name=intern,proto3" json:"intern,omitempty"`
	Public *SearchI_Object_Public `protobuf:"bytes,200,opt,name=public,proto3" json:"public,omitempty"`
}

func (x *SearchI_Object) Reset() {
	*x = SearchI_Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_user_search_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchI_Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchI_Object) ProtoMessage() {}

func (x *SearchI_Object) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_user_search_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchI_Object.ProtoReflect.Descriptor instead.
func (*SearchI_Object) Descriptor() ([]byte, []int) {
	return file_pbf_user_search_proto_rawDescGZIP(), []int{3}
}

func (x *SearchI_Object) GetIntern() *SearchI_Object_Intern {
	if x != nil {
		return x.Intern
	}
	return nil
}

func (x *SearchI_Object) GetPublic() *SearchI_Object_Public {
	if x != nil {
		return x.Public
	}
	return nil
}

type SearchI_Object_Intern struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// user is the ID of the user being searched. If user is left empty, the
	// current user sending the request is looked up and returned based on the
	// given authorization parameters. Given an OAuth access token, its encoded
	// subject claim may refer to the internally persisted user object.
	User string `protobuf:"bytes,100,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *SearchI_Object_Intern) Reset() {
	*x = SearchI_Object_Intern{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_user_search_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchI_Object_Intern) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchI_Object_Intern) ProtoMessage() {}

func (x *SearchI_Object_Intern) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_user_search_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchI_Object_Intern.ProtoReflect.Descriptor instead.
func (*SearchI_Object_Intern) Descriptor() ([]byte, []int) {
	return file_pbf_user_search_proto_rawDescGZIP(), []int{4}
}

func (x *SearchI_Object_Intern) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type SearchI_Object_Public struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the name of the user being searched.
	Name string `protobuf:"bytes,100,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SearchI_Object_Public) Reset() {
	*x = SearchI_Object_Public{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_user_search_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchI_Object_Public) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchI_Object_Public) ProtoMessage() {}

func (x *SearchI_Object_Public) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_user_search_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchI_Object_Public.ProtoReflect.Descriptor instead.
func (*SearchI_Object_Public) Descriptor() ([]byte, []int) {
	return file_pbf_user_search_proto_rawDescGZIP(), []int{5}
}

func (x *SearchI_Object_Public) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// SearchO is the output for searching users.
//
//	{
//	    "filter": [
//	        "chunking": {
//	            "perpage": "50",
//	            "pointer": "150"
//	        }
//	    ],
//	    "object": [
//	        {
//	            "intern": {
//	                "crtd": "1689001255",
//	                "user": "551265"
//	            },
//	            "public": {
//	                "imag": "https://gravatar.com/9050f8ea",
//	                "name": "xh3b4sd"
//	            }
//	        },
//	        ...
//	    ]
//	}
type SearchO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *SearchO_Filter   `protobuf:"bytes,100,opt,name=filter,proto3" json:"filter,omitempty"`
	Object []*SearchO_Object `protobuf:"bytes,200,rep,name=object,proto3" json:"object,omitempty"`
}

func (x *SearchO) Reset() {
	*x = SearchO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_user_search_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchO) ProtoMessage() {}

func (x *SearchO) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_user_search_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchO.ProtoReflect.Descriptor instead.
func (*SearchO) Descriptor() ([]byte, []int) {
	return file_pbf_user_search_proto_rawDescGZIP(), []int{6}
}

func (x *SearchO) GetFilter() *SearchO_Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *SearchO) GetObject() []*SearchO_Object {
	if x != nil {
		return x.Object
	}
	return nil
}

type SearchO_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SearchO_Filter) Reset() {
	*x = SearchO_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_user_search_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchO_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchO_Filter) ProtoMessage() {}

func (x *SearchO_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_user_search_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchO_Filter.ProtoReflect.Descriptor instead.
func (*SearchO_Filter) Descriptor() ([]byte, []int) {
	return file_pbf_user_search_proto_rawDescGZIP(), []int{7}
}

type SearchO_Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Intern *SearchO_Object_Intern `protobuf:"bytes,100,opt,name=intern,proto3" json:"intern,omitempty"`
	Public *SearchO_Object_Public `protobuf:"bytes,200,opt,name=public,proto3" json:"public,omitempty"`
}

func (x *SearchO_Object) Reset() {
	*x = SearchO_Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_user_search_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchO_Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchO_Object) ProtoMessage() {}

func (x *SearchO_Object) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_user_search_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchO_Object.ProtoReflect.Descriptor instead.
func (*SearchO_Object) Descriptor() ([]byte, []int) {
	return file_pbf_user_search_proto_rawDescGZIP(), []int{8}
}

func (x *SearchO_Object) GetIntern() *SearchO_Object_Intern {
	if x != nil {
		return x.Intern
	}
	return nil
}

func (x *SearchO_Object) GetPublic() *SearchO_Object_Public {
	if x != nil {
		return x.Public
	}
	return nil
}

type SearchO_Object_Intern struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// crtd is the unix timestamp in seconds at which the event got created.
	Crtd string `protobuf:"bytes,100,opt,name=crtd,proto3" json:"crtd,omitempty"`
	// user is the ID of the user being searched.
	User string `protobuf:"bytes,200,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *SearchO_Object_Intern) Reset() {
	*x = SearchO_Object_Intern{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_user_search_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchO_Object_Intern) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchO_Object_Intern) ProtoMessage() {}

func (x *SearchO_Object_Intern) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_user_search_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchO_Object_Intern.ProtoReflect.Descriptor instead.
func (*SearchO_Object_Intern) Descriptor() ([]byte, []int) {
	return file_pbf_user_search_proto_rawDescGZIP(), []int{9}
}

func (x *SearchO_Object_Intern) GetCrtd() string {
	if x != nil {
		return x.Crtd
	}
	return ""
}

func (x *SearchO_Object_Intern) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type SearchO_Object_Public struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// imag is the URL pointing to the user's profile picture.
	Imag string `protobuf:"bytes,100,opt,name=imag,proto3" json:"imag,omitempty"`
	// name is the user name.
	Name string `protobuf:"bytes,200,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SearchO_Object_Public) Reset() {
	*x = SearchO_Object_Public{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_user_search_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchO_Object_Public) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchO_Object_Public) ProtoMessage() {}

func (x *SearchO_Object_Public) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_user_search_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchO_Object_Public.ProtoReflect.Descriptor instead.
func (*SearchO_Object_Public) Descriptor() ([]byte, []int) {
	return file_pbf_user_search_proto_rawDescGZIP(), []int{10}
}

func (x *SearchO_Object_Public) GetImag() string {
	if x != nil {
		return x.Imag
	}
	return ""
}

func (x *SearchO_Object_Public) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_pbf_user_search_proto protoreflect.FileDescriptor

var file_pbf_user_search_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x62, 0x66, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x66, 0x0a,
	0x07, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x12, 0x2c, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x5f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0xc8, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x4b, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49,
	0x5f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x08, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x69, 0x6e, 0x67, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x5f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x69,
	0x6e, 0x67, 0x22, 0x4e, 0x0a, 0x17, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x5f, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x5f, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x65, 0x72, 0x70, 0x61, 0x67, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x70, 0x65, 0x72, 0x70, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x07, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x22, 0x7b, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x5f, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x33, 0x0a, 0x06, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x64,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x52, 0x06, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x34, 0x0a, 0x06, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x22,
	0x2b, 0x0a, 0x15, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x2b, 0x0a, 0x15,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x64, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x66, 0x0a, 0x07, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x4f, 0x12, 0x2c, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x64,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x4f, 0x5f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x2d, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0xc8, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x22, 0x10, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x5f, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x22, 0x7b, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x5f, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x33, 0x0a, 0x06, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x18,
	0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x52, 0x06, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x34, 0x0a, 0x06, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x22, 0x40, 0x0a, 0x15, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x72, 0x74,
	0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x72, 0x74, 0x64, 0x12, 0x13, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x22, 0x40, 0x0a, 0x15, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x5f, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x69,
	0x6d, 0x61, 0x67, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6d, 0x61, 0x67, 0x12,
	0x13, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbf_user_search_proto_rawDescOnce sync.Once
	file_pbf_user_search_proto_rawDescData = file_pbf_user_search_proto_rawDesc
)

func file_pbf_user_search_proto_rawDescGZIP() []byte {
	file_pbf_user_search_proto_rawDescOnce.Do(func() {
		file_pbf_user_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbf_user_search_proto_rawDescData)
	})
	return file_pbf_user_search_proto_rawDescData
}

var file_pbf_user_search_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_pbf_user_search_proto_goTypes = []interface{}{
	(*SearchI)(nil),                 // 0: user.SearchI
	(*SearchI_Filter)(nil),          // 1: user.SearchI_Filter
	(*SearchI_Filter_Chunking)(nil), // 2: user.SearchI_Filter_Chunking
	(*SearchI_Object)(nil),          // 3: user.SearchI_Object
	(*SearchI_Object_Intern)(nil),   // 4: user.SearchI_Object_Intern
	(*SearchI_Object_Public)(nil),   // 5: user.SearchI_Object_Public
	(*SearchO)(nil),                 // 6: user.SearchO
	(*SearchO_Filter)(nil),          // 7: user.SearchO_Filter
	(*SearchO_Object)(nil),          // 8: user.SearchO_Object
	(*SearchO_Object_Intern)(nil),   // 9: user.SearchO_Object_Intern
	(*SearchO_Object_Public)(nil),   // 10: user.SearchO_Object_Public
}
var file_pbf_user_search_proto_depIdxs = []int32{
	1,  // 0: user.SearchI.filter:type_name -> user.SearchI_Filter
	3,  // 1: user.SearchI.object:type_name -> user.SearchI_Object
	2,  // 2: user.SearchI_Filter.chunking:type_name -> user.SearchI_Filter_Chunking
	4,  // 3: user.SearchI_Object.intern:type_name -> user.SearchI_Object_Intern
	5,  // 4: user.SearchI_Object.public:type_name -> user.SearchI_Object_Public
	7,  // 5: user.SearchO.filter:type_name -> user.SearchO_Filter
	8,  // 6: user.SearchO.object:type_name -> user.SearchO_Object
	9,  // 7: user.SearchO_Object.intern:type_name -> user.SearchO_Object_Intern
	10, // 8: user.SearchO_Object.public:type_name -> user.SearchO_Object_Public
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_pbf_user_search_proto_init() }
func file_pbf_user_search_proto_init() {
	if File_pbf_user_search_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbf_user_search_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchI); i {
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
		file_pbf_user_search_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchI_Filter); i {
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
		file_pbf_user_search_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchI_Filter_Chunking); i {
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
		file_pbf_user_search_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchI_Object); i {
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
		file_pbf_user_search_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchI_Object_Intern); i {
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
		file_pbf_user_search_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchI_Object_Public); i {
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
		file_pbf_user_search_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchO); i {
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
		file_pbf_user_search_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchO_Filter); i {
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
		file_pbf_user_search_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchO_Object); i {
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
		file_pbf_user_search_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchO_Object_Intern); i {
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
		file_pbf_user_search_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchO_Object_Public); i {
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
			RawDescriptor: file_pbf_user_search_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbf_user_search_proto_goTypes,
		DependencyIndexes: file_pbf_user_search_proto_depIdxs,
		MessageInfos:      file_pbf_user_search_proto_msgTypes,
	}.Build()
	File_pbf_user_search_proto = out.File
	file_pbf_user_search_proto_rawDesc = nil
	file_pbf_user_search_proto_goTypes = nil
	file_pbf_user_search_proto_depIdxs = nil
}
