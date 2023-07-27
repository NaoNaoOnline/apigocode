// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: pbf/user/create.proto

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

// CreateI is the input for creating users.
//
//	{
//	    "object": [
//	        {
//	            "public": {
//	                "name": "xh3b4sd"
//	            }
//	        }
//	    ]
//	}
type CreateI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *CreateI_Filter   `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	Object []*CreateI_Object `protobuf:"bytes,2,rep,name=object,proto3" json:"object,omitempty"`
}

func (x *CreateI) Reset() {
	*x = CreateI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_user_create_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateI) ProtoMessage() {}

func (x *CreateI) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_user_create_proto_msgTypes[0]
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
	return file_pbf_user_create_proto_rawDescGZIP(), []int{0}
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
		mi := &file_pbf_user_create_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateI_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateI_Filter) ProtoMessage() {}

func (x *CreateI_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_user_create_proto_msgTypes[1]
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
	return file_pbf_user_create_proto_rawDescGZIP(), []int{1}
}

type CreateI_Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Intern map[string]string      `protobuf:"bytes,1,rep,name=intern,proto3" json:"intern,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Public *CreateI_Object_Public `protobuf:"bytes,2,opt,name=public,proto3" json:"public,omitempty"`
}

func (x *CreateI_Object) Reset() {
	*x = CreateI_Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_user_create_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateI_Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateI_Object) ProtoMessage() {}

func (x *CreateI_Object) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_user_create_proto_msgTypes[2]
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
	return file_pbf_user_create_proto_rawDescGZIP(), []int{2}
}

func (x *CreateI_Object) GetIntern() map[string]string {
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

type CreateI_Object_Public struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the user name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateI_Object_Public) Reset() {
	*x = CreateI_Object_Public{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_user_create_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateI_Object_Public) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateI_Object_Public) ProtoMessage() {}

func (x *CreateI_Object_Public) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_user_create_proto_msgTypes[3]
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
	return file_pbf_user_create_proto_rawDescGZIP(), []int{3}
}

func (x *CreateI_Object_Public) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// CreateO is the output for creating users.
//
//	{
//	    "object": [
//	        {
//	            "intern": {
//	                "user.naonao.online/id": "<id>",
//	                "user.naonao.online/created": "<unix>"
//	            }
//	        }
//	    ]
//	}
type CreateO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *CreateO_Filter   `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	Object []*CreateO_Object `protobuf:"bytes,2,rep,name=object,proto3" json:"object,omitempty"`
}

func (x *CreateO) Reset() {
	*x = CreateO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_user_create_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateO) ProtoMessage() {}

func (x *CreateO) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_user_create_proto_msgTypes[4]
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
	return file_pbf_user_create_proto_rawDescGZIP(), []int{4}
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
		mi := &file_pbf_user_create_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateO_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateO_Filter) ProtoMessage() {}

func (x *CreateO_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_user_create_proto_msgTypes[5]
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
	return file_pbf_user_create_proto_rawDescGZIP(), []int{5}
}

type CreateO_Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Intern map[string]string      `protobuf:"bytes,1,rep,name=intern,proto3" json:"intern,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Public *CreateO_Object_Public `protobuf:"bytes,2,opt,name=public,proto3" json:"public,omitempty"`
}

func (x *CreateO_Object) Reset() {
	*x = CreateO_Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_user_create_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateO_Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateO_Object) ProtoMessage() {}

func (x *CreateO_Object) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_user_create_proto_msgTypes[6]
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
	return file_pbf_user_create_proto_rawDescGZIP(), []int{6}
}

func (x *CreateO_Object) GetIntern() map[string]string {
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

type CreateO_Object_Public struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateO_Object_Public) Reset() {
	*x = CreateO_Object_Public{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_user_create_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateO_Object_Public) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateO_Object_Public) ProtoMessage() {}

func (x *CreateO_Object_Public) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_user_create_proto_msgTypes[7]
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
	return file_pbf_user_create_proto_rawDescGZIP(), []int{7}
}

var File_pbf_user_create_proto protoreflect.FileDescriptor

var file_pbf_user_create_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x62, 0x66, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x65, 0x0a,
	0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x12, 0x2c, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x5f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x5f,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0xba, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x38, 0x0a, 0x06, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x12, 0x33, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x1a, 0x39, 0x0a, 0x0b, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x2b, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x5f, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x65, 0x0a, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x12, 0x2c, 0x0a, 0x06, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x5f, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x06, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4f, 0x5f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0xba, 0x01, 0x0a, 0x0e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x38, 0x0a, 0x06,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x33, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x1a, 0x39, 0x0a, 0x0b, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x17, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x42,
	0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pbf_user_create_proto_rawDescOnce sync.Once
	file_pbf_user_create_proto_rawDescData = file_pbf_user_create_proto_rawDesc
)

func file_pbf_user_create_proto_rawDescGZIP() []byte {
	file_pbf_user_create_proto_rawDescOnce.Do(func() {
		file_pbf_user_create_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbf_user_create_proto_rawDescData)
	})
	return file_pbf_user_create_proto_rawDescData
}

var file_pbf_user_create_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pbf_user_create_proto_goTypes = []interface{}{
	(*CreateI)(nil),               // 0: user.CreateI
	(*CreateI_Filter)(nil),        // 1: user.CreateI_Filter
	(*CreateI_Object)(nil),        // 2: user.CreateI_Object
	(*CreateI_Object_Public)(nil), // 3: user.CreateI_Object_Public
	(*CreateO)(nil),               // 4: user.CreateO
	(*CreateO_Filter)(nil),        // 5: user.CreateO_Filter
	(*CreateO_Object)(nil),        // 6: user.CreateO_Object
	(*CreateO_Object_Public)(nil), // 7: user.CreateO_Object_Public
	nil,                           // 8: user.CreateI_Object.InternEntry
	nil,                           // 9: user.CreateO_Object.InternEntry
}
var file_pbf_user_create_proto_depIdxs = []int32{
	1, // 0: user.CreateI.filter:type_name -> user.CreateI_Filter
	2, // 1: user.CreateI.object:type_name -> user.CreateI_Object
	8, // 2: user.CreateI_Object.intern:type_name -> user.CreateI_Object.InternEntry
	3, // 3: user.CreateI_Object.public:type_name -> user.CreateI_Object_Public
	5, // 4: user.CreateO.filter:type_name -> user.CreateO_Filter
	6, // 5: user.CreateO.object:type_name -> user.CreateO_Object
	9, // 6: user.CreateO_Object.intern:type_name -> user.CreateO_Object.InternEntry
	7, // 7: user.CreateO_Object.public:type_name -> user.CreateO_Object_Public
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_pbf_user_create_proto_init() }
func file_pbf_user_create_proto_init() {
	if File_pbf_user_create_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbf_user_create_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_pbf_user_create_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_pbf_user_create_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_pbf_user_create_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_pbf_user_create_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_pbf_user_create_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_pbf_user_create_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_pbf_user_create_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_pbf_user_create_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbf_user_create_proto_goTypes,
		DependencyIndexes: file_pbf_user_create_proto_depIdxs,
		MessageInfos:      file_pbf_user_create_proto_msgTypes,
	}.Build()
	File_pbf_user_create_proto = out.File
	file_pbf_user_create_proto_rawDesc = nil
	file_pbf_user_create_proto_goTypes = nil
	file_pbf_user_create_proto_depIdxs = nil
}
