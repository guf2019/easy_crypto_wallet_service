// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: easy_crypto_wallet_service.proto

package easy_crypto_wallet_service

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request messages
type GetEthereumBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *GetEthereumBalanceRequest) Reset() {
	*x = GetEthereumBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_easy_crypto_wallet_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEthereumBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEthereumBalanceRequest) ProtoMessage() {}

func (x *GetEthereumBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_easy_crypto_wallet_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEthereumBalanceRequest.ProtoReflect.Descriptor instead.
func (*GetEthereumBalanceRequest) Descriptor() ([]byte, []int) {
	return file_easy_crypto_wallet_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetEthereumBalanceRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

// Response messages
type GetEthereumBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string  `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Balance float64 `protobuf:"fixed64,2,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *GetEthereumBalanceResponse) Reset() {
	*x = GetEthereumBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_easy_crypto_wallet_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEthereumBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEthereumBalanceResponse) ProtoMessage() {}

func (x *GetEthereumBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_easy_crypto_wallet_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEthereumBalanceResponse.ProtoReflect.Descriptor instead.
func (*GetEthereumBalanceResponse) Descriptor() ([]byte, []int) {
	return file_easy_crypto_wallet_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetEthereumBalanceResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GetEthereumBalanceResponse) GetBalance() float64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

var File_easy_crypto_wallet_service_proto protoreflect.FileDescriptor

var file_easy_crypto_wallet_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x65, 0x61, 0x73, 0x79, 0x5f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1a, 0x65, 0x61, 0x73, 0x79, 0x5f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a,
	0x19, 0x47, 0x65, 0x74, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x22, 0x50, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x45, 0x74, 0x68, 0x65, 0x72,
	0x65, 0x75, 0x6d, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x32, 0xad, 0x02, 0x0a, 0x17, 0x45, 0x61, 0x73, 0x79, 0x43,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x91, 0x02, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65,
	0x75, 0x6d, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x35, 0x2e, 0x65, 0x61, 0x73, 0x79,
	0x5f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65,
	0x75, 0x6d, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x36, 0x2e, 0x65, 0x61, 0x73, 0x79, 0x5f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8b, 0x01, 0x92, 0x41, 0x64, 0x0a, 0x08,
	0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x12, 0x15, 0x47, 0x65, 0x74, 0x20, 0x65, 0x74,
	0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x20, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x1a,
	0x41, 0x54, 0x68, 0x69, 0x73, 0x20, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x20, 0x68, 0x65, 0x6c,
	0x70, 0x73, 0x20, 0x67, 0x65, 0x74, 0x20, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x20, 0x66,
	0x6f, 0x72, 0x20, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x20, 0x69, 0x6e, 0x20, 0x65, 0x74,
	0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x20, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x5f, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x42, 0xd4, 0x01, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x75, 0x66, 0x32, 0x30, 0x31, 0x39, 0x2f, 0x65, 0x61,
	0x73, 0x79, 0x5f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x65,
	0x61, 0x73, 0x79, 0x5f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x92, 0x41, 0x80, 0x01, 0x12, 0x1e, 0x0a,
	0x17, 0x45, 0x61, 0x73, 0x79, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x01, 0x02,
	0x72, 0x5b, 0x0a, 0x22, 0x45, 0x61, 0x73, 0x79, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x57, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x72, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x35, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x75, 0x66, 0x32, 0x30,
	0x31, 0x39, 0x2f, 0x65, 0x61, 0x73, 0x79, 0x5f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_easy_crypto_wallet_service_proto_rawDescOnce sync.Once
	file_easy_crypto_wallet_service_proto_rawDescData = file_easy_crypto_wallet_service_proto_rawDesc
)

func file_easy_crypto_wallet_service_proto_rawDescGZIP() []byte {
	file_easy_crypto_wallet_service_proto_rawDescOnce.Do(func() {
		file_easy_crypto_wallet_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_easy_crypto_wallet_service_proto_rawDescData)
	})
	return file_easy_crypto_wallet_service_proto_rawDescData
}

var file_easy_crypto_wallet_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_easy_crypto_wallet_service_proto_goTypes = []interface{}{
	(*GetEthereumBalanceRequest)(nil),  // 0: easy_crypto_wallet_service.GetEthereumBalanceRequest
	(*GetEthereumBalanceResponse)(nil), // 1: easy_crypto_wallet_service.GetEthereumBalanceResponse
}
var file_easy_crypto_wallet_service_proto_depIdxs = []int32{
	0, // 0: easy_crypto_wallet_service.EasyCryptoWalletService.GetEthereumBalance:input_type -> easy_crypto_wallet_service.GetEthereumBalanceRequest
	1, // 1: easy_crypto_wallet_service.EasyCryptoWalletService.GetEthereumBalance:output_type -> easy_crypto_wallet_service.GetEthereumBalanceResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_easy_crypto_wallet_service_proto_init() }
func file_easy_crypto_wallet_service_proto_init() {
	if File_easy_crypto_wallet_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_easy_crypto_wallet_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEthereumBalanceRequest); i {
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
		file_easy_crypto_wallet_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEthereumBalanceResponse); i {
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
			RawDescriptor: file_easy_crypto_wallet_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_easy_crypto_wallet_service_proto_goTypes,
		DependencyIndexes: file_easy_crypto_wallet_service_proto_depIdxs,
		MessageInfos:      file_easy_crypto_wallet_service_proto_msgTypes,
	}.Build()
	File_easy_crypto_wallet_service_proto = out.File
	file_easy_crypto_wallet_service_proto_rawDesc = nil
	file_easy_crypto_wallet_service_proto_goTypes = nil
	file_easy_crypto_wallet_service_proto_depIdxs = nil
}