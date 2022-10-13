// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.7
// source: protos/genericsapi.proto

package genericsapiv1

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

type Foo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Value       int64                  `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
	FooSpecific string                 `protobuf:"bytes,4,opt,name=foo_specific,json=fooSpecific,proto3" json:"foo_specific,omitempty"`
	Ts          *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=ts,proto3" json:"ts,omitempty"`
}

func (x *Foo) Reset() {
	*x = Foo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_genericsapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Foo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Foo) ProtoMessage() {}

func (x *Foo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_genericsapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Foo.ProtoReflect.Descriptor instead.
func (*Foo) Descriptor() ([]byte, []int) {
	return file_protos_genericsapi_proto_rawDescGZIP(), []int{0}
}

func (x *Foo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Foo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Foo) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Foo) GetFooSpecific() string {
	if x != nil {
		return x.FooSpecific
	}
	return ""
}

func (x *Foo) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

type Bar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name               string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Value              int64                  `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
	BarSpecific        string                 `protobuf:"bytes,4,opt,name=bar_specific,json=barSpecific,proto3" json:"bar_specific,omitempty"`
	BarAnotherSpecific string                 `protobuf:"bytes,5,opt,name=bar_another_specific,json=barAnotherSpecific,proto3" json:"bar_another_specific,omitempty"`
	Ts                 *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=ts,proto3" json:"ts,omitempty"`
}

func (x *Bar) Reset() {
	*x = Bar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_genericsapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bar) ProtoMessage() {}

func (x *Bar) ProtoReflect() protoreflect.Message {
	mi := &file_protos_genericsapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bar.ProtoReflect.Descriptor instead.
func (*Bar) Descriptor() ([]byte, []int) {
	return file_protos_genericsapi_proto_rawDescGZIP(), []int{1}
}

func (x *Bar) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Bar) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Bar) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Bar) GetBarSpecific() string {
	if x != nil {
		return x.BarSpecific
	}
	return ""
}

func (x *Bar) GetBarAnotherSpecific() string {
	if x != nil {
		return x.BarAnotherSpecific
	}
	return ""
}

func (x *Bar) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ColumnName  string `protobuf:"bytes,1,opt,name=column_name,json=columnName,proto3" json:"column_name,omitempty"`
	StringValue string `protobuf:"bytes,2,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_genericsapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_protos_genericsapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_protos_genericsapi_proto_rawDescGZIP(), []int{2}
}

func (x *Filter) GetColumnName() string {
	if x != nil {
		return x.ColumnName
	}
	return ""
}

func (x *Filter) GetStringValue() string {
	if x != nil {
		return x.StringValue
	}
	return ""
}

type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filters []*Filter `protobuf:"bytes,2,rep,name=filters,proto3" json:"filters,omitempty"`
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_genericsapi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_genericsapi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_protos_genericsapi_proto_rawDescGZIP(), []int{3}
}

func (x *ReadRequest) GetFilters() []*Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type ReadFooReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foos []*Foo `protobuf:"bytes,1,rep,name=foos,proto3" json:"foos,omitempty"`
}

func (x *ReadFooReply) Reset() {
	*x = ReadFooReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_genericsapi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadFooReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadFooReply) ProtoMessage() {}

func (x *ReadFooReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_genericsapi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadFooReply.ProtoReflect.Descriptor instead.
func (*ReadFooReply) Descriptor() ([]byte, []int) {
	return file_protos_genericsapi_proto_rawDescGZIP(), []int{4}
}

func (x *ReadFooReply) GetFoos() []*Foo {
	if x != nil {
		return x.Foos
	}
	return nil
}

type ReadBarReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bars []*Bar `protobuf:"bytes,1,rep,name=bars,proto3" json:"bars,omitempty"`
}

func (x *ReadBarReply) Reset() {
	*x = ReadBarReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_genericsapi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadBarReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadBarReply) ProtoMessage() {}

func (x *ReadBarReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_genericsapi_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadBarReply.ProtoReflect.Descriptor instead.
func (*ReadBarReply) Descriptor() ([]byte, []int) {
	return file_protos_genericsapi_proto_rawDescGZIP(), []int{5}
}

func (x *ReadBarReply) GetBars() []*Bar {
	if x != nil {
		return x.Bars
	}
	return nil
}

var File_protos_genericsapi_proto protoreflect.FileDescriptor

var file_protos_genericsapi_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63,
	0x73, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x69, 0x63, 0x73, 0x61, 0x70, 0x69, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x01, 0x0a, 0x03, 0x46,
	0x6f, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x66, 0x6f, 0x6f, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x66, 0x6f, 0x6f, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x12,
	0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x22, 0xc0, 0x01, 0x0a, 0x03,
	0x42, 0x61, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x62, 0x61, 0x72, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x61, 0x72, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63,
	0x12, 0x30, 0x0a, 0x14, 0x62, 0x61, 0x72, 0x5f, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f,
	0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12,
	0x62, 0x61, 0x72, 0x41, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66,
	0x69, 0x63, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x22, 0x4c,
	0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3e, 0x0a, 0x0b,
	0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x07, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x73, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x22, 0x36, 0x0a, 0x0c,
	0x52, 0x65, 0x61, 0x64, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x26, 0x0a, 0x04,
	0x66, 0x6f, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x69, 0x63, 0x73, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x6f, 0x52, 0x04,
	0x66, 0x6f, 0x6f, 0x73, 0x22, 0x36, 0x0a, 0x0c, 0x52, 0x65, 0x61, 0x64, 0x42, 0x61, 0x72, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x26, 0x0a, 0x04, 0x62, 0x61, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x73, 0x61, 0x70, 0x69,
	0x76, 0x31, 0x2e, 0x42, 0x61, 0x72, 0x52, 0x04, 0x62, 0x61, 0x72, 0x73, 0x32, 0x4d, 0x0a, 0x0a,
	0x46, 0x6f, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x04, 0x52, 0x65,
	0x61, 0x64, 0x12, 0x1a, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x73, 0x61, 0x70, 0x69,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x73, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x61, 0x64, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x4d, 0x0a, 0x0a, 0x42,
	0x61, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x04, 0x52, 0x65, 0x61,
	0x64, 0x12, 0x1a, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x73, 0x61, 0x70, 0x69, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x73, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x61, 0x64, 0x42, 0x61, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x1e, 0x5a, 0x1c, 0x2f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x73, 0x61, 0x70, 0x69, 0x76, 0x31, 0x3b, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x69, 0x63, 0x73, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_protos_genericsapi_proto_rawDescOnce sync.Once
	file_protos_genericsapi_proto_rawDescData = file_protos_genericsapi_proto_rawDesc
)

func file_protos_genericsapi_proto_rawDescGZIP() []byte {
	file_protos_genericsapi_proto_rawDescOnce.Do(func() {
		file_protos_genericsapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_genericsapi_proto_rawDescData)
	})
	return file_protos_genericsapi_proto_rawDescData
}

var file_protos_genericsapi_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_genericsapi_proto_goTypes = []interface{}{
	(*Foo)(nil),                   // 0: genericsapiv1.Foo
	(*Bar)(nil),                   // 1: genericsapiv1.Bar
	(*Filter)(nil),                // 2: genericsapiv1.Filter
	(*ReadRequest)(nil),           // 3: genericsapiv1.ReadRequest
	(*ReadFooReply)(nil),          // 4: genericsapiv1.ReadFooReply
	(*ReadBarReply)(nil),          // 5: genericsapiv1.ReadBarReply
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_protos_genericsapi_proto_depIdxs = []int32{
	6, // 0: genericsapiv1.Foo.ts:type_name -> google.protobuf.Timestamp
	6, // 1: genericsapiv1.Bar.ts:type_name -> google.protobuf.Timestamp
	2, // 2: genericsapiv1.ReadRequest.filters:type_name -> genericsapiv1.Filter
	0, // 3: genericsapiv1.ReadFooReply.foos:type_name -> genericsapiv1.Foo
	1, // 4: genericsapiv1.ReadBarReply.bars:type_name -> genericsapiv1.Bar
	3, // 5: genericsapiv1.FooService.Read:input_type -> genericsapiv1.ReadRequest
	3, // 6: genericsapiv1.BarService.Read:input_type -> genericsapiv1.ReadRequest
	4, // 7: genericsapiv1.FooService.Read:output_type -> genericsapiv1.ReadFooReply
	5, // 8: genericsapiv1.BarService.Read:output_type -> genericsapiv1.ReadBarReply
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_protos_genericsapi_proto_init() }
func file_protos_genericsapi_proto_init() {
	if File_protos_genericsapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_genericsapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Foo); i {
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
		file_protos_genericsapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bar); i {
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
		file_protos_genericsapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
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
		file_protos_genericsapi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRequest); i {
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
		file_protos_genericsapi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadFooReply); i {
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
		file_protos_genericsapi_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadBarReply); i {
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
			RawDescriptor: file_protos_genericsapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_protos_genericsapi_proto_goTypes,
		DependencyIndexes: file_protos_genericsapi_proto_depIdxs,
		MessageInfos:      file_protos_genericsapi_proto_msgTypes,
	}.Build()
	File_protos_genericsapi_proto = out.File
	file_protos_genericsapi_proto_rawDesc = nil
	file_protos_genericsapi_proto_goTypes = nil
	file_protos_genericsapi_proto_depIdxs = nil
}
