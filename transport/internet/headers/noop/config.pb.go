package noop

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

type Config struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Config) Reset() {
	*x = Config{}
	mi := &file_transport_internet_headers_noop_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_headers_noop_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_transport_internet_headers_noop_config_proto_rawDescGZIP(), []int{0}
}

type ConnectionConfig struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConnectionConfig) Reset() {
	*x = ConnectionConfig{}
	mi := &file_transport_internet_headers_noop_config_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionConfig) ProtoMessage() {}

func (x *ConnectionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_headers_noop_config_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionConfig.ProtoReflect.Descriptor instead.
func (*ConnectionConfig) Descriptor() ([]byte, []int) {
	return file_transport_internet_headers_noop_config_proto_rawDescGZIP(), []int{1}
}

var File_transport_internet_headers_noop_config_proto protoreflect.FileDescriptor

const file_transport_internet_headers_noop_config_proto_rawDesc = "" +
	"\n" +
	",transport/internet/headers/noop/config.proto\x12*v2ray.core.transport.internet.headers.noop\"\b\n" +
	"\x06Config\"\x12\n" +
	"\x10ConnectionConfigB\x9f\x01\n" +
	".com.v2ray.core.transport.internet.headers.noopP\x01Z>github.com/v4fly/v4ray-core/v0/transport/internet/headers/noop\xaa\x02*V2Ray.Core.Transport.Internet.Headers.Noopb\x06proto3"

var (
	file_transport_internet_headers_noop_config_proto_rawDescOnce sync.Once
	file_transport_internet_headers_noop_config_proto_rawDescData []byte
)

func file_transport_internet_headers_noop_config_proto_rawDescGZIP() []byte {
	file_transport_internet_headers_noop_config_proto_rawDescOnce.Do(func() {
		file_transport_internet_headers_noop_config_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_transport_internet_headers_noop_config_proto_rawDesc), len(file_transport_internet_headers_noop_config_proto_rawDesc)))
	})
	return file_transport_internet_headers_noop_config_proto_rawDescData
}

var file_transport_internet_headers_noop_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_transport_internet_headers_noop_config_proto_goTypes = []any{
	(*Config)(nil),           // 0: v2ray.core.transport.internet.headers.noop.Config
	(*ConnectionConfig)(nil), // 1: v2ray.core.transport.internet.headers.noop.ConnectionConfig
}
var file_transport_internet_headers_noop_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_transport_internet_headers_noop_config_proto_init() }
func file_transport_internet_headers_noop_config_proto_init() {
	if File_transport_internet_headers_noop_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_transport_internet_headers_noop_config_proto_rawDesc), len(file_transport_internet_headers_noop_config_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transport_internet_headers_noop_config_proto_goTypes,
		DependencyIndexes: file_transport_internet_headers_noop_config_proto_depIdxs,
		MessageInfos:      file_transport_internet_headers_noop_config_proto_msgTypes,
	}.Build()
	File_transport_internet_headers_noop_config_proto = out.File
	file_transport_internet_headers_noop_config_proto_goTypes = nil
	file_transport_internet_headers_noop_config_proto_depIdxs = nil
}
