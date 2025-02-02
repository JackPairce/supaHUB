// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: superpeer.proto

package superpeer

import (
	types "supaHUB/backend/types"
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

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password     string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Peeeraddress string `protobuf:"bytes,3,opt,name=peeeraddress,proto3" json:"peeeraddress,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_superpeer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_superpeer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_superpeer_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterRequest) GetPeeeraddress() string {
	if x != nil {
		return x.Peeeraddress
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool  `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Id      int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_superpeer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_superpeer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_superpeer_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *RegisterResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SearchFilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Filename string `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *SearchFilesRequest) Reset() {
	*x = SearchFilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_superpeer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchFilesRequest) ProtoMessage() {}

func (x *SearchFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_superpeer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchFilesRequest.ProtoReflect.Descriptor instead.
func (*SearchFilesRequest) Descriptor() ([]byte, []int) {
	return file_superpeer_proto_rawDescGZIP(), []int{2}
}

func (x *SearchFilesRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SearchFilesRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type SearchFilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results *types.FileList `protobuf:"bytes,1,opt,name=results,proto3" json:"results,omitempty"`
}

func (x *SearchFilesResponse) Reset() {
	*x = SearchFilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_superpeer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchFilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchFilesResponse) ProtoMessage() {}

func (x *SearchFilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_superpeer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchFilesResponse.ProtoReflect.Descriptor instead.
func (*SearchFilesResponse) Descriptor() ([]byte, []int) {
	return file_superpeer_proto_rawDescGZIP(), []int{3}
}

func (x *SearchFilesResponse) GetResults() *types.FileList {
	if x != nil {
		return x.Results
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_superpeer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_superpeer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_superpeer_proto_rawDescGZIP(), []int{4}
}

type PeerId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PeerId) Reset() {
	*x = PeerId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_superpeer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerId) ProtoMessage() {}

func (x *PeerId) ProtoReflect() protoreflect.Message {
	mi := &file_superpeer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerId.ProtoReflect.Descriptor instead.
func (*PeerId) Descriptor() ([]byte, []int) {
	return file_superpeer_proto_rawDescGZIP(), []int{5}
}

func (x *PeerId) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PeerConnexion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peeraddress string `protobuf:"bytes,1,opt,name=peeraddress,proto3" json:"peeraddress,omitempty"`
}

func (x *PeerConnexion) Reset() {
	*x = PeerConnexion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_superpeer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerConnexion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerConnexion) ProtoMessage() {}

func (x *PeerConnexion) ProtoReflect() protoreflect.Message {
	mi := &file_superpeer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerConnexion.ProtoReflect.Descriptor instead.
func (*PeerConnexion) Descriptor() ([]byte, []int) {
	return file_superpeer_proto_rawDescGZIP(), []int{6}
}

func (x *PeerConnexion) GetPeeraddress() string {
	if x != nil {
		return x.Peeraddress
	}
	return ""
}

var File_superpeer_proto protoreflect.FileDescriptor

var file_superpeer_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x70, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a,
	0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x22, 0x0a, 0x0c, 0x70, 0x65, 0x65, 0x65, 0x72, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x65, 0x65, 0x65, 0x72, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x22, 0x3c, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x40, 0x0a, 0x12, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x0a, 0x13, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x69,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x18, 0x0a, 0x06, 0x50, 0x65, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x31, 0x0a, 0x0d, 0x50, 0x65, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x78, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x65, 0x65, 0x72, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x65, 0x65, 0x72, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0xe5, 0x02, 0x0a, 0x09, 0x53, 0x75, 0x70, 0x65, 0x72,
	0x50, 0x65, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x08,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a,
	0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0b, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x65, 0x65, 0x72,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x09, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x3b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x78, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50,
	0x65, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50,
	0x65, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x78, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x42, 0x0c,
	0x5a, 0x0a, 0x2f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x70, 0x65, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_superpeer_proto_rawDescOnce sync.Once
	file_superpeer_proto_rawDescData = file_superpeer_proto_rawDesc
)

func file_superpeer_proto_rawDescGZIP() []byte {
	file_superpeer_proto_rawDescOnce.Do(func() {
		file_superpeer_proto_rawDescData = protoimpl.X.CompressGZIP(file_superpeer_proto_rawDescData)
	})
	return file_superpeer_proto_rawDescData
}

var file_superpeer_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_superpeer_proto_goTypes = []interface{}{
	(*RegisterRequest)(nil),     // 0: protos.RegisterRequest
	(*RegisterResponse)(nil),    // 1: protos.RegisterResponse
	(*SearchFilesRequest)(nil),  // 2: protos.SearchFilesRequest
	(*SearchFilesResponse)(nil), // 3: protos.SearchFilesResponse
	(*Empty)(nil),               // 4: protos.Empty
	(*PeerId)(nil),              // 5: protos.PeerId
	(*PeerConnexion)(nil),       // 6: protos.PeerConnexion
	(*types.FileList)(nil),      // 7: FileList
}
var file_superpeer_proto_depIdxs = []int32{
	7, // 0: protos.SearchFilesResponse.results:type_name -> FileList
	4, // 1: protos.SuperPeer.Ping:input_type -> protos.Empty
	0, // 2: protos.SuperPeer.Register:input_type -> protos.RegisterRequest
	0, // 3: protos.SuperPeer.Login:input_type -> protos.RegisterRequest
	2, // 4: protos.SuperPeer.SearchFiles:input_type -> protos.SearchFilesRequest
	7, // 5: protos.SuperPeer.GetPeerFiles:input_type -> FileList
	5, // 6: protos.SuperPeer.GetPeerConnexion:input_type -> protos.PeerId
	4, // 7: protos.SuperPeer.Ping:output_type -> protos.Empty
	1, // 8: protos.SuperPeer.Register:output_type -> protos.RegisterResponse
	1, // 9: protos.SuperPeer.Login:output_type -> protos.RegisterResponse
	3, // 10: protos.SuperPeer.SearchFiles:output_type -> protos.SearchFilesResponse
	4, // 11: protos.SuperPeer.GetPeerFiles:output_type -> protos.Empty
	6, // 12: protos.SuperPeer.GetPeerConnexion:output_type -> protos.PeerConnexion
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_superpeer_proto_init() }
func file_superpeer_proto_init() {
	if File_superpeer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_superpeer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_superpeer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_superpeer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchFilesRequest); i {
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
		file_superpeer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchFilesResponse); i {
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
		file_superpeer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_superpeer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerId); i {
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
		file_superpeer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerConnexion); i {
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
			RawDescriptor: file_superpeer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_superpeer_proto_goTypes,
		DependencyIndexes: file_superpeer_proto_depIdxs,
		MessageInfos:      file_superpeer_proto_msgTypes,
	}.Build()
	File_superpeer_proto = out.File
	file_superpeer_proto_rawDesc = nil
	file_superpeer_proto_goTypes = nil
	file_superpeer_proto_depIdxs = nil
}
