// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: orbis/rabin/v1alpha1/deals.proto

package rabinv1alpha1

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

type EncryptedDeal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dhkey     []byte `protobuf:"bytes,1,opt,name=dhkey,proto3" json:"dhkey,omitempty"`
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	Nonce     []byte `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Cipher    []byte `protobuf:"bytes,4,opt,name=cipher,proto3" json:"cipher,omitempty"`
}

func (x *EncryptedDeal) Reset() {
	*x = EncryptedDeal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orbis_rabin_v1alpha1_deals_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptedDeal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptedDeal) ProtoMessage() {}

func (x *EncryptedDeal) ProtoReflect() protoreflect.Message {
	mi := &file_orbis_rabin_v1alpha1_deals_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptedDeal.ProtoReflect.Descriptor instead.
func (*EncryptedDeal) Descriptor() ([]byte, []int) {
	return file_orbis_rabin_v1alpha1_deals_proto_rawDescGZIP(), []int{0}
}

func (x *EncryptedDeal) GetDhkey() []byte {
	if x != nil {
		return x.Dhkey
	}
	return nil
}

func (x *EncryptedDeal) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *EncryptedDeal) GetNonce() []byte {
	if x != nil {
		return x.Nonce
	}
	return nil
}

func (x *EncryptedDeal) GetCipher() []byte {
	if x != nil {
		return x.Cipher
	}
	return nil
}

type Deal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index uint32         `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Deal  *EncryptedDeal `protobuf:"bytes,2,opt,name=deal,proto3" json:"deal,omitempty"`
}

func (x *Deal) Reset() {
	*x = Deal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orbis_rabin_v1alpha1_deals_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deal) ProtoMessage() {}

func (x *Deal) ProtoReflect() protoreflect.Message {
	mi := &file_orbis_rabin_v1alpha1_deals_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deal.ProtoReflect.Descriptor instead.
func (*Deal) Descriptor() ([]byte, []int) {
	return file_orbis_rabin_v1alpha1_deals_proto_rawDescGZIP(), []int{1}
}

func (x *Deal) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Deal) GetDeal() *EncryptedDeal {
	if x != nil {
		return x.Deal
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index    uint32              `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Response *VerifiableResponse `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orbis_rabin_v1alpha1_deals_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_orbis_rabin_v1alpha1_deals_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_orbis_rabin_v1alpha1_deals_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Response) GetResponse() *VerifiableResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

type VerifiableResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId []byte `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Index     uint32 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Approved  bool   `protobuf:"varint,3,opt,name=approved,proto3" json:"approved,omitempty"`
	Signature []byte `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *VerifiableResponse) Reset() {
	*x = VerifiableResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orbis_rabin_v1alpha1_deals_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifiableResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifiableResponse) ProtoMessage() {}

func (x *VerifiableResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orbis_rabin_v1alpha1_deals_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifiableResponse.ProtoReflect.Descriptor instead.
func (*VerifiableResponse) Descriptor() ([]byte, []int) {
	return file_orbis_rabin_v1alpha1_deals_proto_rawDescGZIP(), []int{3}
}

func (x *VerifiableResponse) GetSessionId() []byte {
	if x != nil {
		return x.SessionId
	}
	return nil
}

func (x *VerifiableResponse) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *VerifiableResponse) GetApproved() bool {
	if x != nil {
		return x.Approved
	}
	return false
}

func (x *VerifiableResponse) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_orbis_rabin_v1alpha1_deals_proto protoreflect.FileDescriptor

var file_orbis_rabin_v1alpha1_deals_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6f, 0x72, 0x62, 0x69, 0x73, 0x2f, 0x72, 0x61, 0x62, 0x69, 0x6e, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x65, 0x61, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x6f, 0x72, 0x62, 0x69, 0x73, 0x2e, 0x72, 0x61, 0x62, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x71, 0x0a, 0x0d, 0x45, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x65, 0x64, 0x44, 0x65, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x68, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x64, 0x68, 0x6b, 0x65, 0x79, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6e, 0x6f,
	0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x22, 0x55, 0x0a, 0x04, 0x44,
	0x65, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x37, 0x0a, 0x04, 0x64, 0x65, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6f, 0x72, 0x62, 0x69, 0x73, 0x2e,
	0x72, 0x61, 0x62, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x44, 0x65, 0x61, 0x6c, 0x52, 0x04, 0x64, 0x65,
	0x61, 0x6c, 0x22, 0x66, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x44, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6f, 0x72, 0x62, 0x69, 0x73, 0x2e, 0x72,
	0x61, 0x62, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x83, 0x01, 0x0a, 0x12, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x42, 0xe8, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x72, 0x62, 0x69, 0x73, 0x2e, 0x72,
	0x61, 0x62, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0a, 0x44,
	0x65, 0x61, 0x6c, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6f, 0x72, 0x62, 0x69, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x62, 0x69, 0x73, 0x2f, 0x72,
	0x61, 0x62, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x72, 0x61,
	0x62, 0x69, 0x6e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x4f, 0x52,
	0x58, 0xaa, 0x02, 0x14, 0x4f, 0x72, 0x62, 0x69, 0x73, 0x2e, 0x52, 0x61, 0x62, 0x69, 0x6e, 0x2e,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x14, 0x4f, 0x72, 0x62, 0x69, 0x73,
	0x5c, 0x52, 0x61, 0x62, 0x69, 0x6e, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2,
	0x02, 0x20, 0x4f, 0x72, 0x62, 0x69, 0x73, 0x5c, 0x52, 0x61, 0x62, 0x69, 0x6e, 0x5c, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x16, 0x4f, 0x72, 0x62, 0x69, 0x73, 0x3a, 0x3a, 0x52, 0x61, 0x62, 0x69,
	0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_orbis_rabin_v1alpha1_deals_proto_rawDescOnce sync.Once
	file_orbis_rabin_v1alpha1_deals_proto_rawDescData = file_orbis_rabin_v1alpha1_deals_proto_rawDesc
)

func file_orbis_rabin_v1alpha1_deals_proto_rawDescGZIP() []byte {
	file_orbis_rabin_v1alpha1_deals_proto_rawDescOnce.Do(func() {
		file_orbis_rabin_v1alpha1_deals_proto_rawDescData = protoimpl.X.CompressGZIP(file_orbis_rabin_v1alpha1_deals_proto_rawDescData)
	})
	return file_orbis_rabin_v1alpha1_deals_proto_rawDescData
}

var file_orbis_rabin_v1alpha1_deals_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_orbis_rabin_v1alpha1_deals_proto_goTypes = []interface{}{
	(*EncryptedDeal)(nil),      // 0: orbis.rabin.v1alpha1.EncryptedDeal
	(*Deal)(nil),               // 1: orbis.rabin.v1alpha1.Deal
	(*Response)(nil),           // 2: orbis.rabin.v1alpha1.Response
	(*VerifiableResponse)(nil), // 3: orbis.rabin.v1alpha1.VerifiableResponse
}
var file_orbis_rabin_v1alpha1_deals_proto_depIdxs = []int32{
	0, // 0: orbis.rabin.v1alpha1.Deal.deal:type_name -> orbis.rabin.v1alpha1.EncryptedDeal
	3, // 1: orbis.rabin.v1alpha1.Response.response:type_name -> orbis.rabin.v1alpha1.VerifiableResponse
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_orbis_rabin_v1alpha1_deals_proto_init() }
func file_orbis_rabin_v1alpha1_deals_proto_init() {
	if File_orbis_rabin_v1alpha1_deals_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orbis_rabin_v1alpha1_deals_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptedDeal); i {
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
		file_orbis_rabin_v1alpha1_deals_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deal); i {
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
		file_orbis_rabin_v1alpha1_deals_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_orbis_rabin_v1alpha1_deals_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifiableResponse); i {
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
			RawDescriptor: file_orbis_rabin_v1alpha1_deals_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_orbis_rabin_v1alpha1_deals_proto_goTypes,
		DependencyIndexes: file_orbis_rabin_v1alpha1_deals_proto_depIdxs,
		MessageInfos:      file_orbis_rabin_v1alpha1_deals_proto_msgTypes,
	}.Build()
	File_orbis_rabin_v1alpha1_deals_proto = out.File
	file_orbis_rabin_v1alpha1_deals_proto_rawDesc = nil
	file_orbis_rabin_v1alpha1_deals_proto_goTypes = nil
	file_orbis_rabin_v1alpha1_deals_proto_depIdxs = nil
}
