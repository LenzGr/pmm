// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: managementpb/filter.proto

package managementpb

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Checks for results whose key is equal to the provided scalar value.
// Valid for (int32 | int64 | string ).
type EqualsFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*EqualsFilter_IntValue
	//	*EqualsFilter_LongValue
	//	*EqualsFilter_StringValue
	Value isEqualsFilter_Value `protobuf_oneof:"value"`
}

func (x *EqualsFilter) Reset() {
	*x = EqualsFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_filter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EqualsFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EqualsFilter) ProtoMessage() {}

func (x *EqualsFilter) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_filter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EqualsFilter.ProtoReflect.Descriptor instead.
func (*EqualsFilter) Descriptor() ([]byte, []int) {
	return file_managementpb_filter_proto_rawDescGZIP(), []int{0}
}

func (m *EqualsFilter) GetValue() isEqualsFilter_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *EqualsFilter) GetIntValue() int32 {
	if x, ok := x.GetValue().(*EqualsFilter_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (x *EqualsFilter) GetLongValue() int64 {
	if x, ok := x.GetValue().(*EqualsFilter_LongValue); ok {
		return x.LongValue
	}
	return 0
}

func (x *EqualsFilter) GetStringValue() string {
	if x, ok := x.GetValue().(*EqualsFilter_StringValue); ok {
		return x.StringValue
	}
	return ""
}

type isEqualsFilter_Value interface {
	isEqualsFilter_Value()
}

type EqualsFilter_IntValue struct {
	IntValue int32 `protobuf:"varint,1,opt,name=int_value,json=intValue,proto3,oneof"`
}

type EqualsFilter_LongValue struct {
	LongValue int64 `protobuf:"varint,2,opt,name=long_value,json=longValue,proto3,oneof"`
}

type EqualsFilter_StringValue struct {
	StringValue string `protobuf:"bytes,3,opt,name=string_value,json=stringValue,proto3,oneof"`
}

func (*EqualsFilter_IntValue) isEqualsFilter_Value() {}

func (*EqualsFilter_LongValue) isEqualsFilter_Value() {}

func (*EqualsFilter_StringValue) isEqualsFilter_Value() {}

// Checks for results whose key is not equal to the provided scalar value.
// Valid for (int32 | int64 | string).
type NotEqualsFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*NotEqualsFilter_IntValue
	//	*NotEqualsFilter_LongValue
	//	*NotEqualsFilter_StringValue
	Value isNotEqualsFilter_Value `protobuf_oneof:"value"`
}

func (x *NotEqualsFilter) Reset() {
	*x = NotEqualsFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_filter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotEqualsFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotEqualsFilter) ProtoMessage() {}

func (x *NotEqualsFilter) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_filter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotEqualsFilter.ProtoReflect.Descriptor instead.
func (*NotEqualsFilter) Descriptor() ([]byte, []int) {
	return file_managementpb_filter_proto_rawDescGZIP(), []int{1}
}

func (m *NotEqualsFilter) GetValue() isNotEqualsFilter_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *NotEqualsFilter) GetIntValue() int32 {
	if x, ok := x.GetValue().(*NotEqualsFilter_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (x *NotEqualsFilter) GetLongValue() int64 {
	if x, ok := x.GetValue().(*NotEqualsFilter_LongValue); ok {
		return x.LongValue
	}
	return 0
}

func (x *NotEqualsFilter) GetStringValue() string {
	if x, ok := x.GetValue().(*NotEqualsFilter_StringValue); ok {
		return x.StringValue
	}
	return ""
}

type isNotEqualsFilter_Value interface {
	isNotEqualsFilter_Value()
}

type NotEqualsFilter_IntValue struct {
	IntValue int32 `protobuf:"varint,1,opt,name=int_value,json=intValue,proto3,oneof"`
}

type NotEqualsFilter_LongValue struct {
	LongValue int64 `protobuf:"varint,2,opt,name=long_value,json=longValue,proto3,oneof"`
}

type NotEqualsFilter_StringValue struct {
	StringValue string `protobuf:"bytes,3,opt,name=string_value,json=stringValue,proto3,oneof"`
}

func (*NotEqualsFilter_IntValue) isNotEqualsFilter_Value() {}

func (*NotEqualsFilter_LongValue) isNotEqualsFilter_Value() {}

func (*NotEqualsFilter_StringValue) isNotEqualsFilter_Value() {}

// Returns results whose key lies within the inclusive range of the provided minimum and maximum values.
// Valid for (IntRangeValues).
type BetweenFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IntRangeValues *IntRangeValues `protobuf:"bytes,1,opt,name=int_range_values,json=intRangeValues,proto3" json:"int_range_values,omitempty"`
}

func (x *BetweenFilter) Reset() {
	*x = BetweenFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_filter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BetweenFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BetweenFilter) ProtoMessage() {}

func (x *BetweenFilter) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_filter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BetweenFilter.ProtoReflect.Descriptor instead.
func (*BetweenFilter) Descriptor() ([]byte, []int) {
	return file_managementpb_filter_proto_rawDescGZIP(), []int{2}
}

func (x *BetweenFilter) GetIntRangeValues() *IntRangeValues {
	if x != nil {
		return x.IntRangeValues
	}
	return nil
}

// Checks for results whose key exists in a given array.
// Valid for (int_values | long_values | string_values).
type InFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*InFilter_IntValues
	//	*InFilter_LongValues
	//	*InFilter_StringValues
	Value isInFilter_Value `protobuf_oneof:"value"`
}

func (x *InFilter) Reset() {
	*x = InFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_filter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InFilter) ProtoMessage() {}

func (x *InFilter) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_filter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InFilter.ProtoReflect.Descriptor instead.
func (*InFilter) Descriptor() ([]byte, []int) {
	return file_managementpb_filter_proto_rawDescGZIP(), []int{3}
}

func (m *InFilter) GetValue() isInFilter_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *InFilter) GetIntValues() *IntValues {
	if x, ok := x.GetValue().(*InFilter_IntValues); ok {
		return x.IntValues
	}
	return nil
}

func (x *InFilter) GetLongValues() *LongValues {
	if x, ok := x.GetValue().(*InFilter_LongValues); ok {
		return x.LongValues
	}
	return nil
}

func (x *InFilter) GetStringValues() *StringValues {
	if x, ok := x.GetValue().(*InFilter_StringValues); ok {
		return x.StringValues
	}
	return nil
}

type isInFilter_Value interface {
	isInFilter_Value()
}

type InFilter_IntValues struct {
	IntValues *IntValues `protobuf:"bytes,1,opt,name=int_values,json=intValues,proto3,oneof"`
}

type InFilter_LongValues struct {
	LongValues *LongValues `protobuf:"bytes,2,opt,name=long_values,json=longValues,proto3,oneof"`
}

type InFilter_StringValues struct {
	StringValues *StringValues `protobuf:"bytes,3,opt,name=string_values,json=stringValues,proto3,oneof"`
}

func (*InFilter_IntValues) isInFilter_Value() {}

func (*InFilter_LongValues) isInFilter_Value() {}

func (*InFilter_StringValues) isInFilter_Value() {}

// Checks for results whose key contains value as a substring
// Valid for (string_values).
type ContainsFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Value to search results in.
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ContainsFilter) Reset() {
	*x = ContainsFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_filter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainsFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainsFilter) ProtoMessage() {}

func (x *ContainsFilter) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_filter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainsFilter.ProtoReflect.Descriptor instead.
func (*ContainsFilter) Descriptor() ([]byte, []int) {
	return file_managementpb_filter_proto_rawDescGZIP(), []int{4}
}

func (x *ContainsFilter) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type IntValues struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []int32 `protobuf:"varint,1,rep,packed,name=values,proto3" json:"values,omitempty"`
}

func (x *IntValues) Reset() {
	*x = IntValues{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_filter_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntValues) ProtoMessage() {}

func (x *IntValues) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_filter_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntValues.ProtoReflect.Descriptor instead.
func (*IntValues) Descriptor() ([]byte, []int) {
	return file_managementpb_filter_proto_rawDescGZIP(), []int{5}
}

func (x *IntValues) GetValues() []int32 {
	if x != nil {
		return x.Values
	}
	return nil
}

type LongValues struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []int64 `protobuf:"varint,1,rep,packed,name=values,proto3" json:"values,omitempty"`
}

func (x *LongValues) Reset() {
	*x = LongValues{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_filter_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongValues) ProtoMessage() {}

func (x *LongValues) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_filter_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongValues.ProtoReflect.Descriptor instead.
func (*LongValues) Descriptor() ([]byte, []int) {
	return file_managementpb_filter_proto_rawDescGZIP(), []int{6}
}

func (x *LongValues) GetValues() []int64 {
	if x != nil {
		return x.Values
	}
	return nil
}

type StringValues struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *StringValues) Reset() {
	*x = StringValues{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_filter_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringValues) ProtoMessage() {}

func (x *StringValues) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_filter_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringValues.ProtoReflect.Descriptor instead.
func (*StringValues) Descriptor() ([]byte, []int) {
	return file_managementpb_filter_proto_rawDescGZIP(), []int{7}
}

func (x *StringValues) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

type IntRangeValues struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Minimum value in the range.
	Minimum int32 `protobuf:"varint,1,opt,name=minimum,proto3" json:"minimum,omitempty"`
	// Maximum value in the range.
	Maximum int32 `protobuf:"varint,2,opt,name=maximum,proto3" json:"maximum,omitempty"`
}

func (x *IntRangeValues) Reset() {
	*x = IntRangeValues{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_filter_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntRangeValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntRangeValues) ProtoMessage() {}

func (x *IntRangeValues) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_filter_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntRangeValues.ProtoReflect.Descriptor instead.
func (*IntRangeValues) Descriptor() ([]byte, []int) {
	return file_managementpb_filter_proto_rawDescGZIP(), []int{8}
}

func (x *IntRangeValues) GetMinimum() int32 {
	if x != nil {
		return x.Minimum
	}
	return 0
}

func (x *IntRangeValues) GetMaximum() int32 {
	if x != nil {
		return x.Maximum
	}
	return 0
}

var File_managementpb_filter_proto protoreflect.FileDescriptor

var file_managementpb_filter_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7c, 0x0a, 0x0c, 0x45, 0x71, 0x75, 0x61, 0x6c, 0x73, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x7f, 0x0a, 0x0f, 0x4e, 0x6f, 0x74, 0x45, 0x71, 0x75, 0x61, 0x6c, 0x73,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x08, 0x69, 0x6e, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x09, 0x6c, 0x6f, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x55, 0x0a, 0x0d, 0x42, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x10, 0x69, 0x6e, 0x74, 0x5f, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x6e, 0x74,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x0e, 0x69, 0x6e, 0x74,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0xc7, 0x01, 0x0a, 0x08,
	0x49, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x0a, 0x69, 0x6e, 0x74, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x6e, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x48, 0x00, 0x52, 0x09, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x12, 0x39, 0x0a, 0x0b, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x4c, 0x6f, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x48, 0x00, 0x52,
	0x0a, 0x6c, 0x6f, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x3f, 0x0a, 0x0d, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x48, 0x00, 0x52, 0x0c,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x42, 0x07, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x26, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x23, 0x0a,
	0x09, 0x49, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x22, 0x24, 0x0a, 0x0a, 0x4c, 0x6f, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x26, 0x0a, 0x0c, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x22, 0x44, 0x0a, 0x0e, 0x49, 0x6e, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d,
	0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x42, 0x8e, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x0b, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2f, 0x70, 0x6d, 0x6d,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70,
	0x62, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0xca, 0x02, 0x0a, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0xe2, 0x02, 0x16, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_managementpb_filter_proto_rawDescOnce sync.Once
	file_managementpb_filter_proto_rawDescData = file_managementpb_filter_proto_rawDesc
)

func file_managementpb_filter_proto_rawDescGZIP() []byte {
	file_managementpb_filter_proto_rawDescOnce.Do(func() {
		file_managementpb_filter_proto_rawDescData = protoimpl.X.CompressGZIP(file_managementpb_filter_proto_rawDescData)
	})
	return file_managementpb_filter_proto_rawDescData
}

var file_managementpb_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_managementpb_filter_proto_goTypes = []interface{}{
	(*EqualsFilter)(nil),    // 0: management.EqualsFilter
	(*NotEqualsFilter)(nil), // 1: management.NotEqualsFilter
	(*BetweenFilter)(nil),   // 2: management.BetweenFilter
	(*InFilter)(nil),        // 3: management.InFilter
	(*ContainsFilter)(nil),  // 4: management.ContainsFilter
	(*IntValues)(nil),       // 5: management.IntValues
	(*LongValues)(nil),      // 6: management.LongValues
	(*StringValues)(nil),    // 7: management.StringValues
	(*IntRangeValues)(nil),  // 8: management.IntRangeValues
}
var file_managementpb_filter_proto_depIdxs = []int32{
	8, // 0: management.BetweenFilter.int_range_values:type_name -> management.IntRangeValues
	5, // 1: management.InFilter.int_values:type_name -> management.IntValues
	6, // 2: management.InFilter.long_values:type_name -> management.LongValues
	7, // 3: management.InFilter.string_values:type_name -> management.StringValues
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_managementpb_filter_proto_init() }
func file_managementpb_filter_proto_init() {
	if File_managementpb_filter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_managementpb_filter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EqualsFilter); i {
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
		file_managementpb_filter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotEqualsFilter); i {
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
		file_managementpb_filter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BetweenFilter); i {
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
		file_managementpb_filter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InFilter); i {
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
		file_managementpb_filter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainsFilter); i {
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
		file_managementpb_filter_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntValues); i {
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
		file_managementpb_filter_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongValues); i {
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
		file_managementpb_filter_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringValues); i {
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
		file_managementpb_filter_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntRangeValues); i {
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
	file_managementpb_filter_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*EqualsFilter_IntValue)(nil),
		(*EqualsFilter_LongValue)(nil),
		(*EqualsFilter_StringValue)(nil),
	}
	file_managementpb_filter_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*NotEqualsFilter_IntValue)(nil),
		(*NotEqualsFilter_LongValue)(nil),
		(*NotEqualsFilter_StringValue)(nil),
	}
	file_managementpb_filter_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*InFilter_IntValues)(nil),
		(*InFilter_LongValues)(nil),
		(*InFilter_StringValues)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_managementpb_filter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_managementpb_filter_proto_goTypes,
		DependencyIndexes: file_managementpb_filter_proto_depIdxs,
		MessageInfos:      file_managementpb_filter_proto_msgTypes,
	}.Build()
	File_managementpb_filter_proto = out.File
	file_managementpb_filter_proto_rawDesc = nil
	file_managementpb_filter_proto_goTypes = nil
	file_managementpb_filter_proto_depIdxs = nil
}