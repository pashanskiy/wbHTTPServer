// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: store/store.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	core "wbHTTPServer/storage-service/api/core"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetStoresResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoresInfo []*StoreInfo `protobuf:"bytes,1,rep,name=StoresInfo,proto3" json:"StoresInfo,omitempty"`
}

func (x *GetStoresResponse) Reset() {
	*x = GetStoresResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStoresResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStoresResponse) ProtoMessage() {}

func (x *GetStoresResponse) ProtoReflect() protoreflect.Message {
	mi := &file_store_store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStoresResponse.ProtoReflect.Descriptor instead.
func (*GetStoresResponse) Descriptor() ([]byte, []int) {
	return file_store_store_proto_rawDescGZIP(), []int{0}
}

func (x *GetStoresResponse) GetStoresInfo() []*StoreInfo {
	if x != nil {
		return x.StoresInfo
	}
	return nil
}

type StoreInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid     *string `protobuf:"bytes,1,opt,name=uid,proto3,oneof" json:"uid,omitempty"`
	Name    *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Address *string `protobuf:"bytes,3,opt,name=address,proto3,oneof" json:"address,omitempty"`
	Working *bool   `protobuf:"varint,4,opt,name=working,proto3,oneof" json:"working,omitempty"`
	Owner   *string `protobuf:"bytes,5,opt,name=owner,proto3,oneof" json:"owner,omitempty"`
}

func (x *StoreInfo) Reset() {
	*x = StoreInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_store_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreInfo) ProtoMessage() {}

func (x *StoreInfo) ProtoReflect() protoreflect.Message {
	mi := &file_store_store_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreInfo.ProtoReflect.Descriptor instead.
func (*StoreInfo) Descriptor() ([]byte, []int) {
	return file_store_store_proto_rawDescGZIP(), []int{1}
}

func (x *StoreInfo) GetUid() string {
	if x != nil && x.Uid != nil {
		return *x.Uid
	}
	return ""
}

func (x *StoreInfo) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *StoreInfo) GetAddress() string {
	if x != nil && x.Address != nil {
		return *x.Address
	}
	return ""
}

func (x *StoreInfo) GetWorking() bool {
	if x != nil && x.Working != nil {
		return *x.Working
	}
	return false
}

func (x *StoreInfo) GetOwner() string {
	if x != nil && x.Owner != nil {
		return *x.Owner
	}
	return ""
}

type CreateStoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address *string `protobuf:"bytes,2,opt,name=address,proto3,oneof" json:"address,omitempty"`
	Working *bool   `protobuf:"varint,3,opt,name=working,proto3,oneof" json:"working,omitempty"`
	Owner   *string `protobuf:"bytes,4,opt,name=owner,proto3,oneof" json:"owner,omitempty"`
}

func (x *CreateStoreRequest) Reset() {
	*x = CreateStoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_store_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStoreRequest) ProtoMessage() {}

func (x *CreateStoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_store_store_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStoreRequest.ProtoReflect.Descriptor instead.
func (*CreateStoreRequest) Descriptor() ([]byte, []int) {
	return file_store_store_proto_rawDescGZIP(), []int{2}
}

func (x *CreateStoreRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateStoreRequest) GetAddress() string {
	if x != nil && x.Address != nil {
		return *x.Address
	}
	return ""
}

func (x *CreateStoreRequest) GetWorking() bool {
	if x != nil && x.Working != nil {
		return *x.Working
	}
	return false
}

func (x *CreateStoreRequest) GetOwner() string {
	if x != nil && x.Owner != nil {
		return *x.Owner
	}
	return ""
}

type UpdateStoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid     string  `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name    *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Address *string `protobuf:"bytes,3,opt,name=address,proto3,oneof" json:"address,omitempty"`
	Working *bool   `protobuf:"varint,4,opt,name=working,proto3,oneof" json:"working,omitempty"`
	Owner   *string `protobuf:"bytes,5,opt,name=owner,proto3,oneof" json:"owner,omitempty"`
}

func (x *UpdateStoreRequest) Reset() {
	*x = UpdateStoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_store_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStoreRequest) ProtoMessage() {}

func (x *UpdateStoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_store_store_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStoreRequest.ProtoReflect.Descriptor instead.
func (*UpdateStoreRequest) Descriptor() ([]byte, []int) {
	return file_store_store_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateStoreRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *UpdateStoreRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UpdateStoreRequest) GetAddress() string {
	if x != nil && x.Address != nil {
		return *x.Address
	}
	return ""
}

func (x *UpdateStoreRequest) GetWorking() bool {
	if x != nil && x.Working != nil {
		return *x.Working
	}
	return false
}

func (x *UpdateStoreRequest) GetOwner() string {
	if x != nil && x.Owner != nil {
		return *x.Owner
	}
	return ""
}

type DeleteStoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *DeleteStoreRequest) Reset() {
	*x = DeleteStoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_store_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStoreRequest) ProtoMessage() {}

func (x *DeleteStoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_store_store_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStoreRequest.ProtoReflect.Descriptor instead.
func (*DeleteStoreRequest) Descriptor() ([]byte, []int) {
	return file_store_store_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteStoreRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

var File_store_store_proto protoreflect.FileDescriptor

var file_store_store_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x1a, 0x0f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0xc7, 0x01, 0x0a, 0x09, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x15, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x1d, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x02, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x12,
	0x1d, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x03, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x12, 0x19,
	0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x75, 0x69,
	0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x69,
	0x6e, 0x67, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0xa3, 0x01, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x69,
	0x6e, 0x67, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x88, 0x01, 0x01,
	0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x22, 0xc3, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x02, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x88,
	0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x03, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x26, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64,
	0x32, 0xb0, 0x02, 0x0a, 0x13, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a,
	0x20, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x44, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x12, 0x21, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x21, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x44, 0x0a,
	0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x21, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x12, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x77, 0x62, 0x48, 0x54, 0x54, 0x50, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_store_proto_rawDescOnce sync.Once
	file_store_store_proto_rawDescData = file_store_store_proto_rawDesc
)

func file_store_store_proto_rawDescGZIP() []byte {
	file_store_store_proto_rawDescOnce.Do(func() {
		file_store_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_store_proto_rawDescData)
	})
	return file_store_store_proto_rawDescData
}

var file_store_store_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_store_store_proto_goTypes = []interface{}{
	(*GetStoresResponse)(nil),  // 0: store_service.GetStoresResponse
	(*StoreInfo)(nil),          // 1: store_service.StoreInfo
	(*CreateStoreRequest)(nil), // 2: store_service.CreateStoreRequest
	(*UpdateStoreRequest)(nil), // 3: store_service.UpdateStoreRequest
	(*DeleteStoreRequest)(nil), // 4: store_service.DeleteStoreRequest
	(*core.EmptyMessage)(nil),  // 5: core.EmptyMessage
}
var file_store_store_proto_depIdxs = []int32{
	1, // 0: store_service.GetStoresResponse.StoresInfo:type_name -> store_service.StoreInfo
	1, // 1: store_service.StoreStorageService.GetStores:input_type -> store_service.StoreInfo
	2, // 2: store_service.StoreStorageService.CreateStore:input_type -> store_service.CreateStoreRequest
	3, // 3: store_service.StoreStorageService.UpdateStore:input_type -> store_service.UpdateStoreRequest
	4, // 4: store_service.StoreStorageService.DeleteStore:input_type -> store_service.DeleteStoreRequest
	0, // 5: store_service.StoreStorageService.GetStores:output_type -> store_service.GetStoresResponse
	5, // 6: store_service.StoreStorageService.CreateStore:output_type -> core.EmptyMessage
	5, // 7: store_service.StoreStorageService.UpdateStore:output_type -> core.EmptyMessage
	5, // 8: store_service.StoreStorageService.DeleteStore:output_type -> core.EmptyMessage
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_store_store_proto_init() }
func file_store_store_proto_init() {
	if File_store_store_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_store_store_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStoresResponse); i {
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
		file_store_store_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreInfo); i {
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
		file_store_store_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStoreRequest); i {
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
		file_store_store_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStoreRequest); i {
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
		file_store_store_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStoreRequest); i {
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
	file_store_store_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_store_store_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_store_store_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_store_store_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_store_store_proto_goTypes,
		DependencyIndexes: file_store_store_proto_depIdxs,
		MessageInfos:      file_store_store_proto_msgTypes,
	}.Build()
	File_store_store_proto = out.File
	file_store_store_proto_rawDesc = nil
	file_store_store_proto_goTypes = nil
	file_store_store_proto_depIdxs = nil
}
