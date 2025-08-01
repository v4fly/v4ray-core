package tls

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

type PacketConfig struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PacketConfig) Reset() {
	*x = PacketConfig{}
	mi := &file_transport_internet_headers_tls_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PacketConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PacketConfig) ProtoMessage() {}

func (x *PacketConfig) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_headers_tls_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PacketConfig.ProtoReflect.Descriptor instead.
func (*PacketConfig) Descriptor() ([]byte, []int) {
	return file_transport_internet_headers_tls_config_proto_rawDescGZIP(), []int{0}
}

var File_transport_internet_headers_tls_config_proto protoreflect.FileDescriptor

const file_transport_internet_headers_tls_config_proto_rawDesc = "" +
	"\n" +
	"+transport/internet/headers/tls/config.proto\x12)v2ray.core.transport.internet.headers.tls\"\x0e\n" +
	"\fPacketConfigB\x9c\x01\n" +
	"-com.v2ray.core.transport.internet.headers.tlsP\x01Z=github.com/v4fly/v4ray-core/v0/transport/internet/headers/tls\xaa\x02)V2Ray.Core.Transport.Internet.Headers.Tlsb\x06proto3"

var (
	file_transport_internet_headers_tls_config_proto_rawDescOnce sync.Once
	file_transport_internet_headers_tls_config_proto_rawDescData []byte
)

func file_transport_internet_headers_tls_config_proto_rawDescGZIP() []byte {
	file_transport_internet_headers_tls_config_proto_rawDescOnce.Do(func() {
		file_transport_internet_headers_tls_config_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_transport_internet_headers_tls_config_proto_rawDesc), len(file_transport_internet_headers_tls_config_proto_rawDesc)))
	})
	return file_transport_internet_headers_tls_config_proto_rawDescData
}

var file_transport_internet_headers_tls_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_transport_internet_headers_tls_config_proto_goTypes = []any{
	(*PacketConfig)(nil), // 0: v2ray.core.transport.internet.headers.tls.PacketConfig
}
var file_transport_internet_headers_tls_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_transport_internet_headers_tls_config_proto_init() }
func file_transport_internet_headers_tls_config_proto_init() {
	if File_transport_internet_headers_tls_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_transport_internet_headers_tls_config_proto_rawDesc), len(file_transport_internet_headers_tls_config_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transport_internet_headers_tls_config_proto_goTypes,
		DependencyIndexes: file_transport_internet_headers_tls_config_proto_depIdxs,
		MessageInfos:      file_transport_internet_headers_tls_config_proto_msgTypes,
	}.Build()
	File_transport_internet_headers_tls_config_proto = out.File
	file_transport_internet_headers_tls_config_proto_goTypes = nil
	file_transport_internet_headers_tls_config_proto_depIdxs = nil
}
