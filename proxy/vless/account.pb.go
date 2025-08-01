package vless

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Account struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the account, in the form of a UUID, e.g., "66ad4540-b58c-4ad2-9926-ea63445a9b57".
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Flow settings.
	Flow string `protobuf:"bytes,2,opt,name=flow,proto3" json:"flow,omitempty"`
	// Encryption settings. Only applies to client side, and only accepts "none" for now.
	Encryption    string `protobuf:"bytes,3,opt,name=encryption,proto3" json:"encryption,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Account) Reset() {
	*x = Account{}
	mi := &file_proxy_vless_account_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_vless_account_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_proxy_vless_account_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Account) GetFlow() string {
	if x != nil {
		return x.Flow
	}
	return ""
}

func (x *Account) GetEncryption() string {
	if x != nil {
		return x.Encryption
	}
	return ""
}

var File_proxy_vless_account_proto protoreflect.FileDescriptor

const file_proxy_vless_account_proto_rawDesc = "" +
	"\n" +
	"\x19proxy/vless/account.proto\x12\x16v2ray.core.proxy.vless\"M\n" +
	"\aAccount\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04flow\x18\x02 \x01(\tR\x04flow\x12\x1e\n" +
	"\n" +
	"encryption\x18\x03 \x01(\tR\n" +
	"encryptionBc\n" +
	"\x1acom.v2ray.core.proxy.vlessP\x01Z*github.com/v4fly/v4ray-core/v0/proxy/vless\xaa\x02\x16V2Ray.Core.Proxy.Vlessb\x06proto3"

var (
	file_proxy_vless_account_proto_rawDescOnce sync.Once
	file_proxy_vless_account_proto_rawDescData []byte
)

func file_proxy_vless_account_proto_rawDescGZIP() []byte {
	file_proxy_vless_account_proto_rawDescOnce.Do(func() {
		file_proxy_vless_account_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proxy_vless_account_proto_rawDesc), len(file_proxy_vless_account_proto_rawDesc)))
	})
	return file_proxy_vless_account_proto_rawDescData
}

var file_proxy_vless_account_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proxy_vless_account_proto_goTypes = []any{
	(*Account)(nil), // 0: v2ray.core.proxy.vless.Account
}
var file_proxy_vless_account_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proxy_vless_account_proto_init() }
func file_proxy_vless_account_proto_init() {
	if File_proxy_vless_account_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proxy_vless_account_proto_rawDesc), len(file_proxy_vless_account_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proxy_vless_account_proto_goTypes,
		DependencyIndexes: file_proxy_vless_account_proto_depIdxs,
		MessageInfos:      file_proxy_vless_account_proto_msgTypes,
	}.Build()
	File_proxy_vless_account_proto = out.File
	file_proxy_vless_account_proto_goTypes = nil
	file_proxy_vless_account_proto_depIdxs = nil
}
