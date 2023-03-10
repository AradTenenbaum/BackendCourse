// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: rpc_withdraw.proto

package pb

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

type WithdrawRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId int64 `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Amount    int64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *WithdrawRequest) Reset() {
	*x = WithdrawRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_withdraw_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithdrawRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawRequest) ProtoMessage() {}

func (x *WithdrawRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_withdraw_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawRequest.ProtoReflect.Descriptor instead.
func (*WithdrawRequest) Descriptor() ([]byte, []int) {
	return file_rpc_withdraw_proto_rawDescGZIP(), []int{0}
}

func (x *WithdrawRequest) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *WithdrawRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type WithdrawResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewEntry *Entry   `protobuf:"bytes,1,opt,name=new_entry,json=newEntry,proto3" json:"new_entry,omitempty"`
	Account  *Account `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *WithdrawResponse) Reset() {
	*x = WithdrawResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_withdraw_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithdrawResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawResponse) ProtoMessage() {}

func (x *WithdrawResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_withdraw_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawResponse.ProtoReflect.Descriptor instead.
func (*WithdrawResponse) Descriptor() ([]byte, []int) {
	return file_rpc_withdraw_proto_rawDescGZIP(), []int{1}
}

func (x *WithdrawResponse) GetNewEntry() *Entry {
	if x != nil {
		return x.NewEntry
	}
	return nil
}

func (x *WithdrawResponse) GetAccount() *Account {
	if x != nil {
		return x.Account
	}
	return nil
}

var File_rpc_withdraw_proto protoreflect.FileDescriptor

var file_rpc_withdraw_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x70, 0x63, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x0f, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x61,
	0x0a, 0x10, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x26, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x08, 0x6e, 0x65, 0x77, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x25, 0x0a, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x41, 0x72, 0x61, 0x64, 0x54, 0x65, 0x6e, 0x65, 0x6e, 0x62, 0x61, 0x75, 0x6d, 0x2f, 0x42, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_withdraw_proto_rawDescOnce sync.Once
	file_rpc_withdraw_proto_rawDescData = file_rpc_withdraw_proto_rawDesc
)

func file_rpc_withdraw_proto_rawDescGZIP() []byte {
	file_rpc_withdraw_proto_rawDescOnce.Do(func() {
		file_rpc_withdraw_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_withdraw_proto_rawDescData)
	})
	return file_rpc_withdraw_proto_rawDescData
}

var file_rpc_withdraw_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_withdraw_proto_goTypes = []interface{}{
	(*WithdrawRequest)(nil),  // 0: pb.WithdrawRequest
	(*WithdrawResponse)(nil), // 1: pb.WithdrawResponse
	(*Entry)(nil),            // 2: pb.Entry
	(*Account)(nil),          // 3: pb.Account
}
var file_rpc_withdraw_proto_depIdxs = []int32{
	2, // 0: pb.WithdrawResponse.new_entry:type_name -> pb.Entry
	3, // 1: pb.WithdrawResponse.account:type_name -> pb.Account
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rpc_withdraw_proto_init() }
func file_rpc_withdraw_proto_init() {
	if File_rpc_withdraw_proto != nil {
		return
	}
	file_account_proto_init()
	file_entry_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_withdraw_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WithdrawRequest); i {
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
		file_rpc_withdraw_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WithdrawResponse); i {
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
			RawDescriptor: file_rpc_withdraw_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_withdraw_proto_goTypes,
		DependencyIndexes: file_rpc_withdraw_proto_depIdxs,
		MessageInfos:      file_rpc_withdraw_proto_msgTypes,
	}.Build()
	File_rpc_withdraw_proto = out.File
	file_rpc_withdraw_proto_rawDesc = nil
	file_rpc_withdraw_proto_goTypes = nil
	file_rpc_withdraw_proto_depIdxs = nil
}
