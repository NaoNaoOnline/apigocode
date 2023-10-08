// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: pbf/policy/search.proto

package policy

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

// SearchI is the input for searching policies.
//
//	{
//	    "filter": {
//	        "paging": {
//	            "perpage": "50",
//	            "pointer": "100"
//	        }
//	    },
//	    "object": [
//	        {
//	            "symbol": {
//	                "ltst": "default"
//	            }
//	        }
//	    ]
//	}
//
// For more information about the policy contracts, record types and SMAs, see
// the contracts repository on Github.
//
//	https://github.com/NaoNaoOnline/contracts
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
		mi := &file_pbf_policy_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchI) ProtoMessage() {}

func (x *SearchI) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_policy_search_proto_msgTypes[0]
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
	return file_pbf_policy_search_proto_rawDescGZIP(), []int{0}
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

	Paging *SearchI_Filter_Paging `protobuf:"bytes,100,opt,name=paging,proto3" json:"paging,omitempty"`
}

func (x *SearchI_Filter) Reset() {
	*x = SearchI_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_policy_search_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchI_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchI_Filter) ProtoMessage() {}

func (x *SearchI_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_policy_search_proto_msgTypes[1]
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
	return file_pbf_policy_search_proto_rawDescGZIP(), []int{1}
}

func (x *SearchI_Filter) GetPaging() *SearchI_Filter_Paging {
	if x != nil {
		return x.Paging
	}
	return nil
}

type SearchI_Filter_Paging struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Perpage string `protobuf:"bytes,100,opt,name=perpage,proto3" json:"perpage,omitempty"`
	Pointer string `protobuf:"bytes,200,opt,name=pointer,proto3" json:"pointer,omitempty"`
}

func (x *SearchI_Filter_Paging) Reset() {
	*x = SearchI_Filter_Paging{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_policy_search_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchI_Filter_Paging) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchI_Filter_Paging) ProtoMessage() {}

func (x *SearchI_Filter_Paging) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_policy_search_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchI_Filter_Paging.ProtoReflect.Descriptor instead.
func (*SearchI_Filter_Paging) Descriptor() ([]byte, []int) {
	return file_pbf_policy_search_proto_rawDescGZIP(), []int{2}
}

func (x *SearchI_Filter_Paging) GetPerpage() string {
	if x != nil {
		return x.Perpage
	}
	return ""
}

func (x *SearchI_Filter_Paging) GetPointer() string {
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
	Symbol *SearchI_Object_Symbol `protobuf:"bytes,300,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (x *SearchI_Object) Reset() {
	*x = SearchI_Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_policy_search_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchI_Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchI_Object) ProtoMessage() {}

func (x *SearchI_Object) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_policy_search_proto_msgTypes[3]
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
	return file_pbf_policy_search_proto_rawDescGZIP(), []int{3}
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

func (x *SearchI_Object) GetSymbol() *SearchI_Object_Symbol {
	if x != nil {
		return x.Symbol
	}
	return nil
}

type SearchI_Object_Intern struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// kind is the record type for which policies are being searched. Providing
	// kind may only be allowed if ltst is set to "raw".
	//
	//	CreateMember for records of members being created within a system
	//	CreateSystem for records of systems being created
	//	DeleteMember for records of members being deleted within a system
	//	DeleteSystem for records of systems being deleted
	Kind string `protobuf:"bytes,100,opt,name=kind,proto3" json:"kind,omitempty"`
}

func (x *SearchI_Object_Intern) Reset() {
	*x = SearchI_Object_Intern{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_policy_search_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchI_Object_Intern) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchI_Object_Intern) ProtoMessage() {}

func (x *SearchI_Object_Intern) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_policy_search_proto_msgTypes[4]
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
	return file_pbf_policy_search_proto_rawDescGZIP(), []int{4}
}

func (x *SearchI_Object_Intern) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

type SearchI_Object_Public struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SearchI_Object_Public) Reset() {
	*x = SearchI_Object_Public{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_policy_search_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchI_Object_Public) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchI_Object_Public) ProtoMessage() {}

func (x *SearchI_Object_Public) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_policy_search_proto_msgTypes[5]
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
	return file_pbf_policy_search_proto_rawDescGZIP(), []int{5}
}

type SearchI_Object_Symbol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ltst set to "default" returns the latest aggregated version of cached
	// policy records indexed from all onchain smart contracts configured. That
	// is, the list of aggregated records representing the currently active
	// authorization states, plus the list of records that have been removed so
	// far.
	//
	// ltst set to "aggregated" returns the latest aggregated version of cached
	// policy records indexed from all onchain smart contracts configured. That
	// is, the list of aggregated records representing the currently active
	// authorization states, minus the list of records that have been removed so
	// far.
	//
	// ltst set to "raw" returns the complete list of records submitted so far.
	// This raw data is not aggregated, transparently representing the onchain
	// state of emitted smart contract events currently cached internally.
	//
	// Note that indexing happens periodically in a background process, which can
	// be triggered by policy members to update the cached state on demand using
	// policy.API/Update.
	Ltst string `protobuf:"bytes,100,opt,name=ltst,proto3" json:"ltst,omitempty"`
}

func (x *SearchI_Object_Symbol) Reset() {
	*x = SearchI_Object_Symbol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_policy_search_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchI_Object_Symbol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchI_Object_Symbol) ProtoMessage() {}

func (x *SearchI_Object_Symbol) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_policy_search_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchI_Object_Symbol.ProtoReflect.Descriptor instead.
func (*SearchI_Object_Symbol) Descriptor() ([]byte, []int) {
	return file_pbf_policy_search_proto_rawDescGZIP(), []int{6}
}

func (x *SearchI_Object_Symbol) GetLtst() string {
	if x != nil {
		return x.Ltst
	}
	return ""
}

// SearchO is the output for searching policies.
//
//	{
//	    "filter": {
//	        "paging": {
//	            "intotal": "750",
//	            "pointer": "150"
//	        }
//	    },
//	    "object": [
//	        {
//	            "intern": {
//	                "chid": "42161",
//	                "from": "0x1234",
//	                "hash": "0x2345",
//	                "kind": "CreateMember"
//	            },
//	            "public": {
//	                "acce": "2",
//	                "memb": "0x3456",
//	                "syst": "0"
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
		mi := &file_pbf_policy_search_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchO) ProtoMessage() {}

func (x *SearchO) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_policy_search_proto_msgTypes[7]
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
	return file_pbf_policy_search_proto_rawDescGZIP(), []int{7}
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

	Paging *SearchO_Filter_Paging `protobuf:"bytes,100,opt,name=paging,proto3" json:"paging,omitempty"`
}

func (x *SearchO_Filter) Reset() {
	*x = SearchO_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_policy_search_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchO_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchO_Filter) ProtoMessage() {}

func (x *SearchO_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_policy_search_proto_msgTypes[8]
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
	return file_pbf_policy_search_proto_rawDescGZIP(), []int{8}
}

func (x *SearchO_Filter) GetPaging() *SearchO_Filter_Paging {
	if x != nil {
		return x.Paging
	}
	return nil
}

type SearchO_Filter_Paging struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Intotal string `protobuf:"bytes,100,opt,name=intotal,proto3" json:"intotal,omitempty"`
	Pointer string `protobuf:"bytes,200,opt,name=pointer,proto3" json:"pointer,omitempty"`
}

func (x *SearchO_Filter_Paging) Reset() {
	*x = SearchO_Filter_Paging{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_policy_search_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchO_Filter_Paging) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchO_Filter_Paging) ProtoMessage() {}

func (x *SearchO_Filter_Paging) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_policy_search_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchO_Filter_Paging.ProtoReflect.Descriptor instead.
func (*SearchO_Filter_Paging) Descriptor() ([]byte, []int) {
	return file_pbf_policy_search_proto_rawDescGZIP(), []int{9}
}

func (x *SearchO_Filter_Paging) GetIntotal() string {
	if x != nil {
		return x.Intotal
	}
	return ""
}

func (x *SearchO_Filter_Paging) GetPointer() string {
	if x != nil {
		return x.Pointer
	}
	return ""
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
		mi := &file_pbf_policy_search_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchO_Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchO_Object) ProtoMessage() {}

func (x *SearchO_Object) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_policy_search_proto_msgTypes[10]
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
	return file_pbf_policy_search_proto_rawDescGZIP(), []int{10}
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

	// chid is the chain ID, the unique identifier representing the blockchain
	// network on which this record is located.
	Chid string `protobuf:"bytes,100,opt,name=chid,proto3" json:"chid,omitempty"`
	// from is the record creator, the sender of the transaction that submitted
	// this record.
	From string `protobuf:"bytes,200,opt,name=from,proto3" json:"from,omitempty"`
	// hash is the onchain transaction hash that submitted this record.
	Hash string `protobuf:"bytes,300,opt,name=hash,proto3" json:"hash,omitempty"`
	// kind is the record type.
	//
	//	CreateMember for records of members being created within a system
	//	CreateSystem for records of systems being created
	//	DeleteMember for records of members being deleted within a system
	//	DeleteSystem for records of systems being deleted
	Kind string `protobuf:"bytes,400,opt,name=kind,proto3" json:"kind,omitempty"`
}

func (x *SearchO_Object_Intern) Reset() {
	*x = SearchO_Object_Intern{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_policy_search_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchO_Object_Intern) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchO_Object_Intern) ProtoMessage() {}

func (x *SearchO_Object_Intern) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_policy_search_proto_msgTypes[11]
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
	return file_pbf_policy_search_proto_rawDescGZIP(), []int{11}
}

func (x *SearchO_Object_Intern) GetChid() string {
	if x != nil {
		return x.Chid
	}
	return ""
}

func (x *SearchO_Object_Intern) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *SearchO_Object_Intern) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *SearchO_Object_Intern) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

type SearchO_Object_Public struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// acce is the record level, permission or role.
	Acce string `protobuf:"bytes,100,opt,name=acce,proto3" json:"acce,omitempty"`
	// memb is the record account, identity or user.
	Memb string `protobuf:"bytes,200,opt,name=memb,proto3" json:"memb,omitempty"`
	// syst is the record context, resource or scope.
	Syst string `protobuf:"bytes,300,opt,name=syst,proto3" json:"syst,omitempty"`
}

func (x *SearchO_Object_Public) Reset() {
	*x = SearchO_Object_Public{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbf_policy_search_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchO_Object_Public) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchO_Object_Public) ProtoMessage() {}

func (x *SearchO_Object_Public) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_policy_search_proto_msgTypes[12]
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
	return file_pbf_policy_search_proto_rawDescGZIP(), []int{12}
}

func (x *SearchO_Object_Public) GetAcce() string {
	if x != nil {
		return x.Acce
	}
	return ""
}

func (x *SearchO_Object_Public) GetMemb() string {
	if x != nil {
		return x.Memb
	}
	return ""
}

func (x *SearchO_Object_Public) GetSyst() string {
	if x != nil {
		return x.Syst
	}
	return ""
}

var File_pbf_policy_search_proto protoreflect.FileDescriptor

var file_pbf_policy_search_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x62, 0x66, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x22, 0x6a, 0x0a, 0x07, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x12, 0x2e, 0x0a, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x5f, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x06,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0xc8, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x5f, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x47, 0x0a,
	0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x5f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x35, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49,
	0x5f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x06,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x22, 0x4c, 0x0a, 0x15, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x49, 0x5f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x70, 0x61, 0x67, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x65, 0x72, 0x70, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x07, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x22, 0xb7, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49,
	0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x35, 0x0a, 0x06, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x52, 0x06, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x36,
	0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49,
	0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x06,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x18, 0xac, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x22, 0x2b,
	0x0a, 0x15, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18,
	0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x22, 0x2b, 0x0a, 0x15, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x5f,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x6c, 0x74, 0x73, 0x74, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x74, 0x73,
	0x74, 0x22, 0x6a, 0x0a, 0x07, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x12, 0x2e, 0x0a, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x5f, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x06,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0xc8, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x5f, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x47, 0x0a,
	0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x5f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x35, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f,
	0x5f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x06,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x22, 0x4c, 0x0a, 0x15, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x4f, 0x5f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x12,
	0x18, 0x0a, 0x07, 0x69, 0x6e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x69, 0x6e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x19, 0x0a, 0x07, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x22, 0x7f, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x5f,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x35, 0x0a, 0x06, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x52, 0x06, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x36, 0x0a,
	0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x5f,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x06, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x22, 0x6a, 0x0a, 0x15, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f,
	0x5f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x68, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x68,
	0x69, 0x64, 0x12, 0x13, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x13, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18,
	0xac, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x13, 0x0a, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x18, 0x90, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x22, 0x55, 0x0a, 0x15, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x5f, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x63,
	0x63, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x63, 0x63, 0x65, 0x12, 0x13,
	0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x62, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d,
	0x65, 0x6d, 0x62, 0x12, 0x13, 0x0a, 0x04, 0x73, 0x79, 0x73, 0x74, 0x18, 0xac, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x73, 0x79, 0x73, 0x74, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x3b, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbf_policy_search_proto_rawDescOnce sync.Once
	file_pbf_policy_search_proto_rawDescData = file_pbf_policy_search_proto_rawDesc
)

func file_pbf_policy_search_proto_rawDescGZIP() []byte {
	file_pbf_policy_search_proto_rawDescOnce.Do(func() {
		file_pbf_policy_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbf_policy_search_proto_rawDescData)
	})
	return file_pbf_policy_search_proto_rawDescData
}

var file_pbf_policy_search_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_pbf_policy_search_proto_goTypes = []interface{}{
	(*SearchI)(nil),               // 0: policy.SearchI
	(*SearchI_Filter)(nil),        // 1: policy.SearchI_Filter
	(*SearchI_Filter_Paging)(nil), // 2: policy.SearchI_Filter_Paging
	(*SearchI_Object)(nil),        // 3: policy.SearchI_Object
	(*SearchI_Object_Intern)(nil), // 4: policy.SearchI_Object_Intern
	(*SearchI_Object_Public)(nil), // 5: policy.SearchI_Object_Public
	(*SearchI_Object_Symbol)(nil), // 6: policy.SearchI_Object_Symbol
	(*SearchO)(nil),               // 7: policy.SearchO
	(*SearchO_Filter)(nil),        // 8: policy.SearchO_Filter
	(*SearchO_Filter_Paging)(nil), // 9: policy.SearchO_Filter_Paging
	(*SearchO_Object)(nil),        // 10: policy.SearchO_Object
	(*SearchO_Object_Intern)(nil), // 11: policy.SearchO_Object_Intern
	(*SearchO_Object_Public)(nil), // 12: policy.SearchO_Object_Public
}
var file_pbf_policy_search_proto_depIdxs = []int32{
	1,  // 0: policy.SearchI.filter:type_name -> policy.SearchI_Filter
	3,  // 1: policy.SearchI.object:type_name -> policy.SearchI_Object
	2,  // 2: policy.SearchI_Filter.paging:type_name -> policy.SearchI_Filter_Paging
	4,  // 3: policy.SearchI_Object.intern:type_name -> policy.SearchI_Object_Intern
	5,  // 4: policy.SearchI_Object.public:type_name -> policy.SearchI_Object_Public
	6,  // 5: policy.SearchI_Object.symbol:type_name -> policy.SearchI_Object_Symbol
	8,  // 6: policy.SearchO.filter:type_name -> policy.SearchO_Filter
	10, // 7: policy.SearchO.object:type_name -> policy.SearchO_Object
	9,  // 8: policy.SearchO_Filter.paging:type_name -> policy.SearchO_Filter_Paging
	11, // 9: policy.SearchO_Object.intern:type_name -> policy.SearchO_Object_Intern
	12, // 10: policy.SearchO_Object.public:type_name -> policy.SearchO_Object_Public
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_pbf_policy_search_proto_init() }
func file_pbf_policy_search_proto_init() {
	if File_pbf_policy_search_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbf_policy_search_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_pbf_policy_search_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_pbf_policy_search_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchI_Filter_Paging); i {
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
		file_pbf_policy_search_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_pbf_policy_search_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_pbf_policy_search_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_pbf_policy_search_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchI_Object_Symbol); i {
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
		file_pbf_policy_search_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_pbf_policy_search_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
		file_pbf_policy_search_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchO_Filter_Paging); i {
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
		file_pbf_policy_search_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
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
		file_pbf_policy_search_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
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
		file_pbf_policy_search_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_pbf_policy_search_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbf_policy_search_proto_goTypes,
		DependencyIndexes: file_pbf_policy_search_proto_depIdxs,
		MessageInfos:      file_pbf_policy_search_proto_msgTypes,
	}.Build()
	File_pbf_policy_search_proto = out.File
	file_pbf_policy_search_proto_rawDesc = nil
	file_pbf_policy_search_proto_goTypes = nil
	file_pbf_policy_search_proto_depIdxs = nil
}
