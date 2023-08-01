// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: pbf/label/delete.proto

package label

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

// DeleteI is the input for deleting labels.
//
//	{
//	    "object": [
//	        {
//	            "intern": {
//	                "labl": "863826"
//	            }
//	        }
//	    ]
//	}
type DeleteI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *DeleteI_Filter   `protobuf:"bytes,100,opt,name=filter,proto3" json:"filter,omitempty"`
	Object []*DeleteI_Object `protobuf:"bytes,200,rep,name=object,proto3" json:"object,omitempty"`
}

func (x *DeleteI) Reset() {
	*x = DeleteI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_label_delete_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteI) ProtoMessage() {}

func (x *DeleteI) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_label_delete_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteI.ProtoReflect.Descriptor instead.
func (*DeleteI) Descriptor() ([]byte, []int) {
	return file_pbf_label_delete_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteI) GetFilter() *DeleteI_Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *DeleteI) GetObject() []*DeleteI_Object {
	if x != nil {
		return x.Object
	}
	return nil
}

type DeleteI_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteI_Filter) Reset() {
	*x = DeleteI_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_label_delete_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteI_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteI_Filter) ProtoMessage() {}

func (x *DeleteI_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_label_delete_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteI_Filter.ProtoReflect.Descriptor instead.
func (*DeleteI_Filter) Descriptor() ([]byte, []int) {
	return file_pbf_label_delete_proto_rawDescGZIP(), []int{1}
}

type DeleteI_Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Intern *DeleteI_Object_Intern `protobuf:"bytes,100,opt,name=intern,proto3" json:"intern,omitempty"`
	Public *DeleteI_Object_Public `protobuf:"bytes,200,opt,name=public,proto3" json:"public,omitempty"`
}

func (x *DeleteI_Object) Reset() {
	*x = DeleteI_Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_label_delete_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteI_Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteI_Object) ProtoMessage() {}

func (x *DeleteI_Object) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_label_delete_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteI_Object.ProtoReflect.Descriptor instead.
func (*DeleteI_Object) Descriptor() ([]byte, []int) {
	return file_pbf_label_delete_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteI_Object) GetIntern() *DeleteI_Object_Intern {
	if x != nil {
		return x.Intern
	}
	return nil
}

func (x *DeleteI_Object) GetPublic() *DeleteI_Object_Public {
	if x != nil {
		return x.Public
	}
	return nil
}

type DeleteI_Object_Intern struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// labl is the ID of the label being deleted.
	Labl string `protobuf:"bytes,100,opt,name=labl,proto3" json:"labl,omitempty"`
}

func (x *DeleteI_Object_Intern) Reset() {
	*x = DeleteI_Object_Intern{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_label_delete_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteI_Object_Intern) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteI_Object_Intern) ProtoMessage() {}

func (x *DeleteI_Object_Intern) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_label_delete_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteI_Object_Intern.ProtoReflect.Descriptor instead.
func (*DeleteI_Object_Intern) Descriptor() ([]byte, []int) {
	return file_pbf_label_delete_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteI_Object_Intern) GetLabl() string {
	if x != nil {
		return x.Labl
	}
	return ""
}

type DeleteI_Object_Public struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteI_Object_Public) Reset() {
	*x = DeleteI_Object_Public{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_label_delete_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteI_Object_Public) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteI_Object_Public) ProtoMessage() {}

func (x *DeleteI_Object_Public) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_label_delete_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteI_Object_Public.ProtoReflect.Descriptor instead.
func (*DeleteI_Object_Public) Descriptor() ([]byte, []int) {
	return file_pbf_label_delete_proto_rawDescGZIP(), []int{4}
}

// DeleteO is the output for deleting labels.
//
//	{
//	    "object": [
//	        {
//	            "intern": {
//	                "stts": "deleted"
//	            }
//	        }
//	    ]
//	}
type DeleteO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *DeleteO_Filter   `protobuf:"bytes,100,opt,name=filter,proto3" json:"filter,omitempty"`
	Object []*DeleteO_Object `protobuf:"bytes,200,rep,name=object,proto3" json:"object,omitempty"`
}

func (x *DeleteO) Reset() {
	*x = DeleteO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_label_delete_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteO) ProtoMessage() {}

func (x *DeleteO) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_label_delete_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteO.ProtoReflect.Descriptor instead.
func (*DeleteO) Descriptor() ([]byte, []int) {
	return file_pbf_label_delete_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteO) GetFilter() *DeleteO_Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *DeleteO) GetObject() []*DeleteO_Object {
	if x != nil {
		return x.Object
	}
	return nil
}

type DeleteO_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteO_Filter) Reset() {
	*x = DeleteO_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_label_delete_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteO_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteO_Filter) ProtoMessage() {}

func (x *DeleteO_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_label_delete_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteO_Filter.ProtoReflect.Descriptor instead.
func (*DeleteO_Filter) Descriptor() ([]byte, []int) {
	return file_pbf_label_delete_proto_rawDescGZIP(), []int{6}
}

type DeleteO_Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Intern *DeleteO_Object_Intern `protobuf:"bytes,100,opt,name=intern,proto3" json:"intern,omitempty"`
	Public *DeleteO_Object_Public `protobuf:"bytes,200,opt,name=public,proto3" json:"public,omitempty"`
}

func (x *DeleteO_Object) Reset() {
	*x = DeleteO_Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_label_delete_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteO_Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteO_Object) ProtoMessage() {}

func (x *DeleteO_Object) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_label_delete_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteO_Object.ProtoReflect.Descriptor instead.
func (*DeleteO_Object) Descriptor() ([]byte, []int) {
	return file_pbf_label_delete_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteO_Object) GetIntern() *DeleteO_Object_Intern {
	if x != nil {
		return x.Intern
	}
	return nil
}

func (x *DeleteO_Object) GetPublic() *DeleteO_Object_Public {
	if x != nil {
		return x.Public
	}
	return nil
}

type DeleteO_Object_Intern struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// stts is the resource status upon successful label deletion.
	Stts string `protobuf:"bytes,100,opt,name=stts,proto3" json:"stts,omitempty"`
}

func (x *DeleteO_Object_Intern) Reset() {
	*x = DeleteO_Object_Intern{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_label_delete_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteO_Object_Intern) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteO_Object_Intern) ProtoMessage() {}

func (x *DeleteO_Object_Intern) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_label_delete_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteO_Object_Intern.ProtoReflect.Descriptor instead.
func (*DeleteO_Object_Intern) Descriptor() ([]byte, []int) {
	return file_pbf_label_delete_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteO_Object_Intern) GetStts() string {
	if x != nil {
		return x.Stts
	}
	return ""
}

type DeleteO_Object_Public struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteO_Object_Public) Reset() {
	*x = DeleteO_Object_Public{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_label_delete_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteO_Object_Public) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteO_Object_Public) ProtoMessage() {}

func (x *DeleteO_Object_Public) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_label_delete_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteO_Object_Public.ProtoReflect.Descriptor instead.
func (*DeleteO_Object_Public) Descriptor() ([]byte, []int) {
	return file_pbf_label_delete_proto_rawDescGZIP(), []int{9}
}

var File_pbf_label_delete_proto protoreflect.FileDescriptor

var file_pbf_label_delete_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x62, 0x66, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2f, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22,
	0x68, 0x0a, 0x07, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x12, 0x2d, 0x0a, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x5f, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x06, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x18, 0xc8, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x49, 0x5f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x7d, 0x0a, 0x0e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x34, 0x0a,
	0x06, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x5f, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x52, 0x06, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x12, 0x35, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0xc8, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x22, 0x2b, 0x0a, 0x15, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x62, 0x6c, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6c, 0x61, 0x62, 0x6c, 0x22, 0x17, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x22, 0x68, 0x0a, 0x07, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x12, 0x2d, 0x0a, 0x06, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x5f, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x06, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0xc8, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4f, 0x5f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x7d, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x34,
	0x0a, 0x06, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x5f, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x52, 0x06, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x12, 0x35, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0xc8,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x22, 0x2b, 0x0a, 0x15, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x74, 0x73, 0x18, 0x64, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x73, 0x74, 0x74, 0x73, 0x22, 0x17, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbf_label_delete_proto_rawDescOnce sync.Once
	file_pbf_label_delete_proto_rawDescData = file_pbf_label_delete_proto_rawDesc
)

func file_pbf_label_delete_proto_rawDescGZIP() []byte {
	file_pbf_label_delete_proto_rawDescOnce.Do(func() {
		file_pbf_label_delete_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbf_label_delete_proto_rawDescData)
	})
	return file_pbf_label_delete_proto_rawDescData
}

var file_pbf_label_delete_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pbf_label_delete_proto_goTypes = []interface{}{
	(*DeleteI)(nil),               // 0: label.DeleteI
	(*DeleteI_Filter)(nil),        // 1: label.DeleteI_Filter
	(*DeleteI_Object)(nil),        // 2: label.DeleteI_Object
	(*DeleteI_Object_Intern)(nil), // 3: label.DeleteI_Object_Intern
	(*DeleteI_Object_Public)(nil), // 4: label.DeleteI_Object_Public
	(*DeleteO)(nil),               // 5: label.DeleteO
	(*DeleteO_Filter)(nil),        // 6: label.DeleteO_Filter
	(*DeleteO_Object)(nil),        // 7: label.DeleteO_Object
	(*DeleteO_Object_Intern)(nil), // 8: label.DeleteO_Object_Intern
	(*DeleteO_Object_Public)(nil), // 9: label.DeleteO_Object_Public
}
var file_pbf_label_delete_proto_depIdxs = []int32{
	1, // 0: label.DeleteI.filter:type_name -> label.DeleteI_Filter
	2, // 1: label.DeleteI.object:type_name -> label.DeleteI_Object
	3, // 2: label.DeleteI_Object.intern:type_name -> label.DeleteI_Object_Intern
	4, // 3: label.DeleteI_Object.public:type_name -> label.DeleteI_Object_Public
	6, // 4: label.DeleteO.filter:type_name -> label.DeleteO_Filter
	7, // 5: label.DeleteO.object:type_name -> label.DeleteO_Object
	8, // 6: label.DeleteO_Object.intern:type_name -> label.DeleteO_Object_Intern
	9, // 7: label.DeleteO_Object.public:type_name -> label.DeleteO_Object_Public
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_pbf_label_delete_proto_init() }
func file_pbf_label_delete_proto_init() {
	if File_pbf_label_delete_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbf_label_delete_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteI); i {
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
		file_pbf_label_delete_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteI_Filter); i {
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
		file_pbf_label_delete_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteI_Object); i {
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
		file_pbf_label_delete_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteI_Object_Intern); i {
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
		file_pbf_label_delete_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteI_Object_Public); i {
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
		file_pbf_label_delete_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteO); i {
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
		file_pbf_label_delete_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteO_Filter); i {
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
		file_pbf_label_delete_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteO_Object); i {
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
		file_pbf_label_delete_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteO_Object_Intern); i {
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
		file_pbf_label_delete_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteO_Object_Public); i {
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
			RawDescriptor: file_pbf_label_delete_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbf_label_delete_proto_goTypes,
		DependencyIndexes: file_pbf_label_delete_proto_depIdxs,
		MessageInfos:      file_pbf_label_delete_proto_msgTypes,
	}.Build()
	File_pbf_label_delete_proto = out.File
	file_pbf_label_delete_proto_rawDesc = nil
	file_pbf_label_delete_proto_goTypes = nil
	file_pbf_label_delete_proto_depIdxs = nil
}
