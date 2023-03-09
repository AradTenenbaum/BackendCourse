// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: rpc_transfer.proto

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

type TransferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromAccountId int64  `protobuf:"varint,1,opt,name=from_account_id,json=fromAccountId,proto3" json:"from_account_id,omitempty"`
	ToAccountId   int64  `protobuf:"varint,2,opt,name=to_account_id,json=toAccountId,proto3" json:"to_account_id,omitempty"`
	Amount        int64  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency      string `protobuf:"bytes,4,opt,name=currency,proto3" json:"currency,omitempty"`
}

func (x *TransferRequest) Reset() {
	*x = TransferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_transfer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferRequest) ProtoMessage() {}

func (x *TransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_transfer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferRequest.ProtoReflect.Descriptor instead.
func (*TransferRequest) Descriptor() ([]byte, []int) {
	return file_rpc_transfer_proto_rawDescGZIP(), []int{0}
}

func (x *TransferRequest) GetFromAccountId() int64 {
	if x != nil {
		return x.FromAccountId
	}
	return 0
}

func (x *TransferRequest) GetToAccountId() int64 {
	if x != nil {
		return x.ToAccountId
	}
	return 0
}

func (x *TransferRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *TransferRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

type TransferResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transfer    *Transfer `protobuf:"bytes,1,opt,name=transfer,proto3" json:"transfer,omitempty"`
	FromAccount *Account  `protobuf:"bytes,2,opt,name=from_account,json=fromAccount,proto3" json:"from_account,omitempty"`
	ToAccount   *Account  `protobuf:"bytes,3,opt,name=to_account,json=toAccount,proto3" json:"to_account,omitempty"`
	FromEntry   *Entry    `protobuf:"bytes,4,opt,name=from_entry,json=fromEntry,proto3" json:"from_entry,omitempty"`
	ToEntry     *Entry    `protobuf:"bytes,5,opt,name=to_entry,json=toEntry,proto3" json:"to_entry,omitempty"`
}

func (x *TransferResponse) Reset() {
	*x = TransferResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_transfer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferResponse) ProtoMessage() {}

func (x *TransferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_transfer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferResponse.ProtoReflect.Descriptor instead.
func (*TransferResponse) Descriptor() ([]byte, []int) {
	return file_rpc_transfer_proto_rawDescGZIP(), []int{1}
}

func (x *TransferResponse) GetTransfer() *Transfer {
	if x != nil {
		return x.Transfer
	}
	return nil
}

func (x *TransferResponse) GetFromAccount() *Account {
	if x != nil {
		return x.FromAccount
	}
	return nil
}

func (x *TransferResponse) GetToAccount() *Account {
	if x != nil {
		return x.ToAccount
	}
	return nil
}

func (x *TransferResponse) GetFromEntry() *Entry {
	if x != nil {
		return x.FromEntry
	}
	return nil
}

func (x *TransferResponse) GetToEntry() *Entry {
	if x != nil {
		return x.ToEntry
	}
	return nil
}

var File_rpc_transfer_proto protoreflect.FileDescriptor

var file_rpc_transfer_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x01, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x66, 0x72, 0x6f, 0x6d,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0d, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x22, 0x0a, 0x0d, 0x74, 0x6f, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x6f, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0xe8, 0x01, 0x0a, 0x10, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a,
	0x08, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x08, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x70, 0x62, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x09, 0x74, 0x6f, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x65, 0x6e, 0x74, 0x72,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x24, 0x0a,
	0x08, 0x74, 0x6f, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x74, 0x6f, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x41, 0x72, 0x61, 0x64, 0x54, 0x65, 0x6e, 0x65, 0x6e, 0x62, 0x61, 0x75, 0x6d, 0x2f,
	0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_transfer_proto_rawDescOnce sync.Once
	file_rpc_transfer_proto_rawDescData = file_rpc_transfer_proto_rawDesc
)

func file_rpc_transfer_proto_rawDescGZIP() []byte {
	file_rpc_transfer_proto_rawDescOnce.Do(func() {
		file_rpc_transfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_transfer_proto_rawDescData)
	})
	return file_rpc_transfer_proto_rawDescData
}

var file_rpc_transfer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_transfer_proto_goTypes = []interface{}{
	(*TransferRequest)(nil),  // 0: pb.TransferRequest
	(*TransferResponse)(nil), // 1: pb.TransferResponse
	(*Transfer)(nil),         // 2: pb.Transfer
	(*Account)(nil),          // 3: pb.Account
	(*Entry)(nil),            // 4: pb.Entry
}
var file_rpc_transfer_proto_depIdxs = []int32{
	2, // 0: pb.TransferResponse.transfer:type_name -> pb.Transfer
	3, // 1: pb.TransferResponse.from_account:type_name -> pb.Account
	3, // 2: pb.TransferResponse.to_account:type_name -> pb.Account
	4, // 3: pb.TransferResponse.from_entry:type_name -> pb.Entry
	4, // 4: pb.TransferResponse.to_entry:type_name -> pb.Entry
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_rpc_transfer_proto_init() }
func file_rpc_transfer_proto_init() {
	if File_rpc_transfer_proto != nil {
		return
	}
	file_transfer_proto_init()
	file_account_proto_init()
	file_entry_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_transfer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferRequest); i {
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
		file_rpc_transfer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferResponse); i {
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
			RawDescriptor: file_rpc_transfer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_transfer_proto_goTypes,
		DependencyIndexes: file_rpc_transfer_proto_depIdxs,
		MessageInfos:      file_rpc_transfer_proto_msgTypes,
	}.Build()
	File_rpc_transfer_proto = out.File
	file_rpc_transfer_proto_rawDesc = nil
	file_rpc_transfer_proto_goTypes = nil
	file_rpc_transfer_proto_depIdxs = nil
}
