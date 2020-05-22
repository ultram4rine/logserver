// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.12.0
// source: log.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DHCPLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MAC  string `protobuf:"bytes,1,opt,name=MAC,proto3" json:"MAC,omitempty"`
	From string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To   string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *DHCPLogsRequest) Reset() {
	*x = DHCPLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DHCPLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DHCPLogsRequest) ProtoMessage() {}

func (x *DHCPLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DHCPLogsRequest.ProtoReflect.Descriptor instead.
func (*DHCPLogsRequest) Descriptor() ([]byte, []int) {
	return file_log_proto_rawDescGZIP(), []int{0}
}

func (x *DHCPLogsRequest) GetMAC() string {
	if x != nil {
		return x.MAC
	}
	return ""
}

func (x *DHCPLogsRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *DHCPLogsRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

type DHCPLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip        string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Timestamp string `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Message   string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DHCPLog) Reset() {
	*x = DHCPLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DHCPLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DHCPLog) ProtoMessage() {}

func (x *DHCPLog) ProtoReflect() protoreflect.Message {
	mi := &file_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DHCPLog.ProtoReflect.Descriptor instead.
func (*DHCPLog) Descriptor() ([]byte, []int) {
	return file_log_proto_rawDescGZIP(), []int{1}
}

func (x *DHCPLog) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *DHCPLog) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *DHCPLog) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DHCPLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Log []*DHCPLog `protobuf:"bytes,1,rep,name=log,proto3" json:"log,omitempty"`
	Err string     `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *DHCPLogsResponse) Reset() {
	*x = DHCPLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DHCPLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DHCPLogsResponse) ProtoMessage() {}

func (x *DHCPLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DHCPLogsResponse.ProtoReflect.Descriptor instead.
func (*DHCPLogsResponse) Descriptor() ([]byte, []int) {
	return file_log_proto_rawDescGZIP(), []int{2}
}

func (x *DHCPLogsResponse) GetLog() []*DHCPLog {
	if x != nil {
		return x.Log
	}
	return nil
}

func (x *DHCPLogsResponse) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type SwitchLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	From string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To   string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *SwitchLogsRequest) Reset() {
	*x = SwitchLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwitchLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwitchLogsRequest) ProtoMessage() {}

func (x *SwitchLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_log_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwitchLogsRequest.ProtoReflect.Descriptor instead.
func (*SwitchLogsRequest) Descriptor() ([]byte, []int) {
	return file_log_proto_rawDescGZIP(), []int{3}
}

func (x *SwitchLogsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SwitchLogsRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *SwitchLogsRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

type SwitchLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip      string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Ts      string `protobuf:"bytes,3,opt,name=ts,proto3" json:"ts,omitempty"`
	Message string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SwitchLog) Reset() {
	*x = SwitchLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwitchLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwitchLog) ProtoMessage() {}

func (x *SwitchLog) ProtoReflect() protoreflect.Message {
	mi := &file_log_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwitchLog.ProtoReflect.Descriptor instead.
func (*SwitchLog) Descriptor() ([]byte, []int) {
	return file_log_proto_rawDescGZIP(), []int{4}
}

func (x *SwitchLog) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *SwitchLog) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SwitchLog) GetTs() string {
	if x != nil {
		return x.Ts
	}
	return ""
}

func (x *SwitchLog) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SwitchLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Log []*SwitchLog `protobuf:"bytes,1,rep,name=log,proto3" json:"log,omitempty"`
	Err string       `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *SwitchLogsResponse) Reset() {
	*x = SwitchLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwitchLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwitchLogsResponse) ProtoMessage() {}

func (x *SwitchLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_log_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwitchLogsResponse.ProtoReflect.Descriptor instead.
func (*SwitchLogsResponse) Descriptor() ([]byte, []int) {
	return file_log_proto_rawDescGZIP(), []int{5}
}

func (x *SwitchLogsResponse) GetLog() []*SwitchLog {
	if x != nil {
		return x.Log
	}
	return nil
}

func (x *SwitchLogsResponse) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type SimilarSwitchesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SimilarSwitchesRequest) Reset() {
	*x = SimilarSwitchesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimilarSwitchesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimilarSwitchesRequest) ProtoMessage() {}

func (x *SimilarSwitchesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_log_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimilarSwitchesRequest.ProtoReflect.Descriptor instead.
func (*SimilarSwitchesRequest) Descriptor() ([]byte, []int) {
	return file_log_proto_rawDescGZIP(), []int{6}
}

func (x *SimilarSwitchesRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SimilarSwitch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	IP   string `protobuf:"bytes,2,opt,name=IP,proto3" json:"IP,omitempty"`
}

func (x *SimilarSwitch) Reset() {
	*x = SimilarSwitch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimilarSwitch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimilarSwitch) ProtoMessage() {}

func (x *SimilarSwitch) ProtoReflect() protoreflect.Message {
	mi := &file_log_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimilarSwitch.ProtoReflect.Descriptor instead.
func (*SimilarSwitch) Descriptor() ([]byte, []int) {
	return file_log_proto_rawDescGZIP(), []int{7}
}

func (x *SimilarSwitch) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SimilarSwitch) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

type SimilarSwitchesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Switch []*SimilarSwitch `protobuf:"bytes,1,rep,name=switch,proto3" json:"switch,omitempty"`
	Err    string           `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *SimilarSwitchesResponse) Reset() {
	*x = SimilarSwitchesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimilarSwitchesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimilarSwitchesResponse) ProtoMessage() {}

func (x *SimilarSwitchesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_log_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimilarSwitchesResponse.ProtoReflect.Descriptor instead.
func (*SimilarSwitchesResponse) Descriptor() ([]byte, []int) {
	return file_log_proto_rawDescGZIP(), []int{8}
}

func (x *SimilarSwitchesResponse) GetSwitch() []*SimilarSwitch {
	if x != nil {
		return x.Switch
	}
	return nil
}

func (x *SimilarSwitchesResponse) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

var File_log_proto protoreflect.FileDescriptor

var file_log_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22,
	0x47, 0x0a, 0x0f, 0x44, 0x48, 0x43, 0x50, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x41, 0x43, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x4d, 0x41, 0x43, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x07, 0x44, 0x48, 0x43, 0x50,
	0x4c, 0x6f, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x43, 0x0a, 0x10, 0x44,
	0x48, 0x43, 0x50, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1d, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70,
	0x62, 0x2e, 0x44, 0x48, 0x43, 0x50, 0x4c, 0x6f, 0x67, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x12, 0x10,
	0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72,
	0x22, 0x4b, 0x0a, 0x11, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x22, 0x59, 0x0a,
	0x09, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x4c, 0x6f, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x47, 0x0a, 0x12, 0x53, 0x77, 0x69, 0x74,
	0x63, 0x68, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62,
	0x2e, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x4c, 0x6f, 0x67, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x12,
	0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72,
	0x72, 0x22, 0x2c, 0x0a, 0x16, 0x53, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72, 0x53, 0x77, 0x69, 0x74,
	0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x33, 0x0a, 0x0d, 0x53, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x50, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x50, 0x22, 0x56, 0x0a, 0x17, 0x53, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72, 0x53,
	0x77, 0x69, 0x74, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x29, 0x0a, 0x06, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72, 0x53, 0x77, 0x69, 0x74,
	0x63, 0x68, 0x52, 0x06, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x32, 0xd4, 0x01, 0x0a,
	0x03, 0x4c, 0x6f, 0x67, 0x12, 0x3a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x44, 0x48, 0x43, 0x50, 0x4c,
	0x6f, 0x67, 0x73, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x48, 0x43, 0x50, 0x4c, 0x6f, 0x67,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x48,
	0x43, 0x50, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x40, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x4c, 0x6f, 0x67,
	0x73, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x4c, 0x6f, 0x67,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x77,
	0x69, 0x74, 0x63, 0x68, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72,
	0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x65, 0x73, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x69,
	0x6d, 0x69, 0x6c, 0x61, 0x72, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x69, 0x6d, 0x69, 0x6c, 0x61,
	0x72, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_log_proto_rawDescOnce sync.Once
	file_log_proto_rawDescData = file_log_proto_rawDesc
)

func file_log_proto_rawDescGZIP() []byte {
	file_log_proto_rawDescOnce.Do(func() {
		file_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_log_proto_rawDescData)
	})
	return file_log_proto_rawDescData
}

var file_log_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_log_proto_goTypes = []interface{}{
	(*DHCPLogsRequest)(nil),         // 0: pb.DHCPLogsRequest
	(*DHCPLog)(nil),                 // 1: pb.DHCPLog
	(*DHCPLogsResponse)(nil),        // 2: pb.DHCPLogsResponse
	(*SwitchLogsRequest)(nil),       // 3: pb.SwitchLogsRequest
	(*SwitchLog)(nil),               // 4: pb.SwitchLog
	(*SwitchLogsResponse)(nil),      // 5: pb.SwitchLogsResponse
	(*SimilarSwitchesRequest)(nil),  // 6: pb.SimilarSwitchesRequest
	(*SimilarSwitch)(nil),           // 7: pb.SimilarSwitch
	(*SimilarSwitchesResponse)(nil), // 8: pb.SimilarSwitchesResponse
}
var file_log_proto_depIdxs = []int32{
	1, // 0: pb.DHCPLogsResponse.log:type_name -> pb.DHCPLog
	4, // 1: pb.SwitchLogsResponse.log:type_name -> pb.SwitchLog
	7, // 2: pb.SimilarSwitchesResponse.switch:type_name -> pb.SimilarSwitch
	0, // 3: pb.Log.GetDHCPLogs:input_type -> pb.DHCPLogsRequest
	3, // 4: pb.Log.GetSwitchLogs:input_type -> pb.SwitchLogsRequest
	6, // 5: pb.Log.GetSimilarSwitches:input_type -> pb.SimilarSwitchesRequest
	2, // 6: pb.Log.GetDHCPLogs:output_type -> pb.DHCPLogsResponse
	5, // 7: pb.Log.GetSwitchLogs:output_type -> pb.SwitchLogsResponse
	8, // 8: pb.Log.GetSimilarSwitches:output_type -> pb.SimilarSwitchesResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_log_proto_init() }
func file_log_proto_init() {
	if File_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DHCPLogsRequest); i {
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
		file_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DHCPLog); i {
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
		file_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DHCPLogsResponse); i {
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
		file_log_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwitchLogsRequest); i {
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
		file_log_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwitchLog); i {
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
		file_log_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwitchLogsResponse); i {
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
		file_log_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimilarSwitchesRequest); i {
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
		file_log_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimilarSwitch); i {
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
		file_log_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimilarSwitchesResponse); i {
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
			RawDescriptor: file_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_log_proto_goTypes,
		DependencyIndexes: file_log_proto_depIdxs,
		MessageInfos:      file_log_proto_msgTypes,
	}.Build()
	File_log_proto = out.File
	file_log_proto_rawDesc = nil
	file_log_proto_goTypes = nil
	file_log_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LogClient is the client API for Log service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogClient interface {
	GetDHCPLogs(ctx context.Context, in *DHCPLogsRequest, opts ...grpc.CallOption) (*DHCPLogsResponse, error)
	GetSwitchLogs(ctx context.Context, in *SwitchLogsRequest, opts ...grpc.CallOption) (*SwitchLogsResponse, error)
	GetSimilarSwitches(ctx context.Context, in *SimilarSwitchesRequest, opts ...grpc.CallOption) (*SimilarSwitchesResponse, error)
}

type logClient struct {
	cc grpc.ClientConnInterface
}

func NewLogClient(cc grpc.ClientConnInterface) LogClient {
	return &logClient{cc}
}

func (c *logClient) GetDHCPLogs(ctx context.Context, in *DHCPLogsRequest, opts ...grpc.CallOption) (*DHCPLogsResponse, error) {
	out := new(DHCPLogsResponse)
	err := c.cc.Invoke(ctx, "/pb.Log/GetDHCPLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logClient) GetSwitchLogs(ctx context.Context, in *SwitchLogsRequest, opts ...grpc.CallOption) (*SwitchLogsResponse, error) {
	out := new(SwitchLogsResponse)
	err := c.cc.Invoke(ctx, "/pb.Log/GetSwitchLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logClient) GetSimilarSwitches(ctx context.Context, in *SimilarSwitchesRequest, opts ...grpc.CallOption) (*SimilarSwitchesResponse, error) {
	out := new(SimilarSwitchesResponse)
	err := c.cc.Invoke(ctx, "/pb.Log/GetSimilarSwitches", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogServer is the server API for Log service.
type LogServer interface {
	GetDHCPLogs(context.Context, *DHCPLogsRequest) (*DHCPLogsResponse, error)
	GetSwitchLogs(context.Context, *SwitchLogsRequest) (*SwitchLogsResponse, error)
	GetSimilarSwitches(context.Context, *SimilarSwitchesRequest) (*SimilarSwitchesResponse, error)
}

// UnimplementedLogServer can be embedded to have forward compatible implementations.
type UnimplementedLogServer struct {
}

func (*UnimplementedLogServer) GetDHCPLogs(context.Context, *DHCPLogsRequest) (*DHCPLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDHCPLogs not implemented")
}
func (*UnimplementedLogServer) GetSwitchLogs(context.Context, *SwitchLogsRequest) (*SwitchLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSwitchLogs not implemented")
}
func (*UnimplementedLogServer) GetSimilarSwitches(context.Context, *SimilarSwitchesRequest) (*SimilarSwitchesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSimilarSwitches not implemented")
}

func RegisterLogServer(s *grpc.Server, srv LogServer) {
	s.RegisterService(&_Log_serviceDesc, srv)
}

func _Log_GetDHCPLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DHCPLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServer).GetDHCPLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Log/GetDHCPLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServer).GetDHCPLogs(ctx, req.(*DHCPLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Log_GetSwitchLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SwitchLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServer).GetSwitchLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Log/GetSwitchLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServer).GetSwitchLogs(ctx, req.(*SwitchLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Log_GetSimilarSwitches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimilarSwitchesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServer).GetSimilarSwitches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Log/GetSimilarSwitches",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServer).GetSimilarSwitches(ctx, req.(*SimilarSwitchesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Log_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Log",
	HandlerType: (*LogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDHCPLogs",
			Handler:    _Log_GetDHCPLogs_Handler,
		},
		{
			MethodName: "GetSwitchLogs",
			Handler:    _Log_GetSwitchLogs_Handler,
		},
		{
			MethodName: "GetSimilarSwitches",
			Handler:    _Log_GetSimilarSwitches_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "log.proto",
}
