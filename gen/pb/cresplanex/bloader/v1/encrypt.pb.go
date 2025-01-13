// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        (unknown)
// source: cresplanex/bloader/v1/encrypt.proto

package bloaderv1

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

type Encryption struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Enabled       bool                   `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	EncryptId     string                 `protobuf:"bytes,2,opt,name=encrypt_id,json=encryptId,proto3" json:"encrypt_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Encryption) Reset() {
	*x = Encryption{}
	mi := &file_cresplanex_bloader_v1_encrypt_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Encryption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Encryption) ProtoMessage() {}

func (x *Encryption) ProtoReflect() protoreflect.Message {
	mi := &file_cresplanex_bloader_v1_encrypt_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Encryption.ProtoReflect.Descriptor instead.
func (*Encryption) Descriptor() ([]byte, []int) {
	return file_cresplanex_bloader_v1_encrypt_proto_rawDescGZIP(), []int{0}
}

func (x *Encryption) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Encryption) GetEncryptId() string {
	if x != nil {
		return x.EncryptId
	}
	return ""
}

var File_cresplanex_bloader_v1_encrypt_proto protoreflect.FileDescriptor

var file_cresplanex_bloader_v1_encrypt_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x72, 0x65, 0x73, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x78, 0x2f, 0x62, 0x6c, 0x6f,
	0x61, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63, 0x72, 0x65, 0x73, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x78, 0x2e, 0x62, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x45, 0x0a, 0x0a,
	0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x49, 0x64, 0x42, 0xe5, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x72, 0x65, 0x73,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x78, 0x2e, 0x62, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x42, 0x0c, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x72,
	0x65, 0x73, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x78, 0x2f, 0x62, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x72, 0x65, 0x73, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x78, 0x2f, 0x62, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x6c,
	0x6f, 0x61, 0x64, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x42, 0x58, 0xaa, 0x02, 0x15,
	0x43, 0x72, 0x65, 0x73, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x78, 0x2e, 0x42, 0x6c, 0x6f, 0x61, 0x64,
	0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x15, 0x43, 0x72, 0x65, 0x73, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x78, 0x5c, 0x42, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x21,
	0x43, 0x72, 0x65, 0x73, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x78, 0x5c, 0x42, 0x6c, 0x6f, 0x61, 0x64,
	0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x17, 0x43, 0x72, 0x65, 0x73, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x78, 0x3a, 0x3a,
	0x42, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_cresplanex_bloader_v1_encrypt_proto_rawDescOnce sync.Once
	file_cresplanex_bloader_v1_encrypt_proto_rawDescData = file_cresplanex_bloader_v1_encrypt_proto_rawDesc
)

func file_cresplanex_bloader_v1_encrypt_proto_rawDescGZIP() []byte {
	file_cresplanex_bloader_v1_encrypt_proto_rawDescOnce.Do(func() {
		file_cresplanex_bloader_v1_encrypt_proto_rawDescData = protoimpl.X.CompressGZIP(file_cresplanex_bloader_v1_encrypt_proto_rawDescData)
	})
	return file_cresplanex_bloader_v1_encrypt_proto_rawDescData
}

var file_cresplanex_bloader_v1_encrypt_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cresplanex_bloader_v1_encrypt_proto_goTypes = []any{
	(*Encryption)(nil), // 0: cresplanex.bloader.v1.Encryption
}
var file_cresplanex_bloader_v1_encrypt_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cresplanex_bloader_v1_encrypt_proto_init() }
func file_cresplanex_bloader_v1_encrypt_proto_init() {
	if File_cresplanex_bloader_v1_encrypt_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cresplanex_bloader_v1_encrypt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cresplanex_bloader_v1_encrypt_proto_goTypes,
		DependencyIndexes: file_cresplanex_bloader_v1_encrypt_proto_depIdxs,
		MessageInfos:      file_cresplanex_bloader_v1_encrypt_proto_msgTypes,
	}.Build()
	File_cresplanex_bloader_v1_encrypt_proto = out.File
	file_cresplanex_bloader_v1_encrypt_proto_rawDesc = nil
	file_cresplanex_bloader_v1_encrypt_proto_goTypes = nil
	file_cresplanex_bloader_v1_encrypt_proto_depIdxs = nil
}