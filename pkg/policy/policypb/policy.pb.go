// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: pkg/policy/policypb/policy.proto

package policypb

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

type Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllowUnauthenticatedUsers                 *wrapperspb.BoolValue `protobuf:"bytes,1,opt,name=allow_unauthenticated_users,json=allowUnauthenticatedUsers,proto3" json:"allow_unauthenticated_users,omitempty"`
	AllowChangingUsernamesAndPasswords        *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=allow_changing_usernames_and_passwords,json=allowChangingUsernamesAndPasswords,proto3" json:"allow_changing_usernames_and_passwords,omitempty"`
	ShowPoints                                *wrapperspb.BoolValue `protobuf:"bytes,3,opt,name=show_points,json=showPoints,proto3" json:"show_points,omitempty"`
	ShowAddresses                             *wrapperspb.BoolValue `protobuf:"bytes,4,opt,name=show_addresses,json=showAddresses,proto3" json:"show_addresses,omitempty"`
	AllowRedTeamLaunchingServiceTestsManually *wrapperspb.BoolValue `protobuf:"bytes,5,opt,name=allow_red_team_launching_service_tests_manually,json=allowRedTeamLaunchingServiceTestsManually,proto3" json:"allow_red_team_launching_service_tests_manually,omitempty"`
}

func (x *Policy) Reset() {
	*x = Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_policy_policypb_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy) ProtoMessage() {}

func (x *Policy) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_policy_policypb_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy.ProtoReflect.Descriptor instead.
func (*Policy) Descriptor() ([]byte, []int) {
	return file_pkg_policy_policypb_policy_proto_rawDescGZIP(), []int{0}
}

func (x *Policy) GetAllowUnauthenticatedUsers() *wrapperspb.BoolValue {
	if x != nil {
		return x.AllowUnauthenticatedUsers
	}
	return nil
}

func (x *Policy) GetAllowChangingUsernamesAndPasswords() *wrapperspb.BoolValue {
	if x != nil {
		return x.AllowChangingUsernamesAndPasswords
	}
	return nil
}

func (x *Policy) GetShowPoints() *wrapperspb.BoolValue {
	if x != nil {
		return x.ShowPoints
	}
	return nil
}

func (x *Policy) GetShowAddresses() *wrapperspb.BoolValue {
	if x != nil {
		return x.ShowAddresses
	}
	return nil
}

func (x *Policy) GetAllowRedTeamLaunchingServiceTestsManually() *wrapperspb.BoolValue {
	if x != nil {
		return x.AllowRedTeamLaunchingServiceTestsManually
	}
	return nil
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_policy_policypb_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_policy_policypb_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_pkg_policy_policypb_policy_proto_rawDescGZIP(), []int{1}
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Policy *Policy `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_policy_policypb_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_policy_policypb_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_pkg_policy_policypb_policy_proto_rawDescGZIP(), []int{2}
}

func (x *GetResponse) GetPolicy() *Policy {
	if x != nil {
		return x.Policy
	}
	return nil
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Policy *Policy `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_policy_policypb_policy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_policy_policypb_policy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_pkg_policy_policypb_policy_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateRequest) GetPolicy() *Policy {
	if x != nil {
		return x.Policy
	}
	return nil
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_policy_policypb_policy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_policy_policypb_policy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_pkg_policy_policypb_policy_proto_rawDescGZIP(), []int{4}
}

var File_pkg_policy_policypb_policy_proto protoreflect.FileDescriptor

var file_pkg_policy_policypb_policy_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x70, 0x62, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x70, 0x62, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x03, 0x0a, 0x06, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x12, 0x5a, 0x0a, 0x1b, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x75, 0x6e, 0x61, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x19, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x6e,
	0x0a, 0x26, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x69, 0x6e, 0x67,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x22, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x69, 0x6e, 0x67, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x41, 0x6e, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x3b,
	0x0a, 0x0b, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x0a, 0x73, 0x68, 0x6f, 0x77, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x41, 0x0a, 0x0e, 0x73,
	0x68, 0x6f, 0x77, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x0d, 0x73, 0x68, 0x6f, 0x77, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x7e,
	0x0a, 0x2f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x65, 0x61, 0x6d,
	0x5f, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x6c,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x29, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x64, 0x54, 0x65, 0x61,
	0x6d, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x54, 0x65, 0x73, 0x74, 0x73, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x6c, 0x79, 0x22, 0x0c,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x42, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x6b,
	0x67, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x70,
	0x62, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x22, 0x44, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x33, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x06,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x10, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xb2, 0x01, 0x0a, 0x0d, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x1f, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x53, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x22, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x34, 0x5a,
	0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x54, 0x72, 0x61, 0x6b, 0x2f, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x54, 0x72, 0x61, 0x6b, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_policy_policypb_policy_proto_rawDescOnce sync.Once
	file_pkg_policy_policypb_policy_proto_rawDescData = file_pkg_policy_policypb_policy_proto_rawDesc
)

func file_pkg_policy_policypb_policy_proto_rawDescGZIP() []byte {
	file_pkg_policy_policypb_policy_proto_rawDescOnce.Do(func() {
		file_pkg_policy_policypb_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_policy_policypb_policy_proto_rawDescData)
	})
	return file_pkg_policy_policypb_policy_proto_rawDescData
}

var file_pkg_policy_policypb_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_policy_policypb_policy_proto_goTypes = []interface{}{
	(*Policy)(nil),               // 0: pkg.policy.policypb.Policy
	(*GetRequest)(nil),           // 1: pkg.policy.policypb.GetRequest
	(*GetResponse)(nil),          // 2: pkg.policy.policypb.GetResponse
	(*UpdateRequest)(nil),        // 3: pkg.policy.policypb.UpdateRequest
	(*UpdateResponse)(nil),       // 4: pkg.policy.policypb.UpdateResponse
	(*wrapperspb.BoolValue)(nil), // 5: google.protobuf.BoolValue
}
var file_pkg_policy_policypb_policy_proto_depIdxs = []int32{
	5, // 0: pkg.policy.policypb.Policy.allow_unauthenticated_users:type_name -> google.protobuf.BoolValue
	5, // 1: pkg.policy.policypb.Policy.allow_changing_usernames_and_passwords:type_name -> google.protobuf.BoolValue
	5, // 2: pkg.policy.policypb.Policy.show_points:type_name -> google.protobuf.BoolValue
	5, // 3: pkg.policy.policypb.Policy.show_addresses:type_name -> google.protobuf.BoolValue
	5, // 4: pkg.policy.policypb.Policy.allow_red_team_launching_service_tests_manually:type_name -> google.protobuf.BoolValue
	0, // 5: pkg.policy.policypb.GetResponse.policy:type_name -> pkg.policy.policypb.Policy
	0, // 6: pkg.policy.policypb.UpdateRequest.policy:type_name -> pkg.policy.policypb.Policy
	1, // 7: pkg.policy.policypb.PolicyService.Get:input_type -> pkg.policy.policypb.GetRequest
	3, // 8: pkg.policy.policypb.PolicyService.Update:input_type -> pkg.policy.policypb.UpdateRequest
	2, // 9: pkg.policy.policypb.PolicyService.Get:output_type -> pkg.policy.policypb.GetResponse
	4, // 10: pkg.policy.policypb.PolicyService.Update:output_type -> pkg.policy.policypb.UpdateResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_pkg_policy_policypb_policy_proto_init() }
func file_pkg_policy_policypb_policy_proto_init() {
	if File_pkg_policy_policypb_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_policy_policypb_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy); i {
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
		file_pkg_policy_policypb_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_pkg_policy_policypb_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_pkg_policy_policypb_policy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_pkg_policy_policypb_policy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
			RawDescriptor: file_pkg_policy_policypb_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_policy_policypb_policy_proto_goTypes,
		DependencyIndexes: file_pkg_policy_policypb_policy_proto_depIdxs,
		MessageInfos:      file_pkg_policy_policypb_policy_proto_msgTypes,
	}.Build()
	File_pkg_policy_policypb_policy_proto = out.File
	file_pkg_policy_policypb_policy_proto_rawDesc = nil
	file_pkg_policy_policypb_policy_proto_goTypes = nil
	file_pkg_policy_policypb_policy_proto_depIdxs = nil
}
