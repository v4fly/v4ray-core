package quic

import (
	protocol "github.com/v4fly/v4ray-core/v0/common/protocol"
	_ "github.com/v4fly/v4ray-core/v0/common/protoext"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
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
	state         protoimpl.MessageState   `protogen:"open.v1"`
	Key           string                   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Security      *protocol.SecurityConfig `protobuf:"bytes,2,opt,name=security,proto3" json:"security,omitempty"`
	Header        *anypb.Any               `protobuf:"bytes,3,opt,name=header,proto3" json:"header,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Config) Reset() {
	*x = Config{}
	mi := &file_transport_internet_quic_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_quic_config_proto_msgTypes[0]
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
	return file_transport_internet_quic_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Config) GetSecurity() *protocol.SecurityConfig {
	if x != nil {
		return x.Security
	}
	return nil
}

func (x *Config) GetHeader() *anypb.Any {
	if x != nil {
		return x.Header
	}
	return nil
}

var File_transport_internet_quic_config_proto protoreflect.FileDescriptor

const file_transport_internet_quic_config_proto_rawDesc = "" +
	"\n" +
	"$transport/internet/quic/config.proto\x12\"v2ray.core.transport.internet.quic\x1a\x19google/protobuf/any.proto\x1a\x1dcommon/protocol/headers.proto\x1a common/protoext/extensions.proto\"\xa7\x01\n" +
	"\x06Config\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12F\n" +
	"\bsecurity\x18\x02 \x01(\v2*.v2ray.core.common.protocol.SecurityConfigR\bsecurity\x12,\n" +
	"\x06header\x18\x03 \x01(\v2\x14.google.protobuf.AnyR\x06header:\x15\x82\xb5\x18\x11\n" +
	"\ttransport\x12\x04quicB\x87\x01\n" +
	"&com.v2ray.core.transport.internet.quicP\x01Z6github.com/v4fly/v4ray-core/v0/transport/internet/quic\xaa\x02\"V2Ray.Core.Transport.Internet.Quicb\x06proto3"

var (
	file_transport_internet_quic_config_proto_rawDescOnce sync.Once
	file_transport_internet_quic_config_proto_rawDescData []byte
)

func file_transport_internet_quic_config_proto_rawDescGZIP() []byte {
	file_transport_internet_quic_config_proto_rawDescOnce.Do(func() {
		file_transport_internet_quic_config_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_transport_internet_quic_config_proto_rawDesc), len(file_transport_internet_quic_config_proto_rawDesc)))
	})
	return file_transport_internet_quic_config_proto_rawDescData
}

var file_transport_internet_quic_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_transport_internet_quic_config_proto_goTypes = []any{
	(*Config)(nil),                  // 0: v2ray.core.transport.internet.quic.Config
	(*protocol.SecurityConfig)(nil), // 1: v2ray.core.common.protocol.SecurityConfig
	(*anypb.Any)(nil),               // 2: google.protobuf.Any
}
var file_transport_internet_quic_config_proto_depIdxs = []int32{
	1, // 0: v2ray.core.transport.internet.quic.Config.security:type_name -> v2ray.core.common.protocol.SecurityConfig
	2, // 1: v2ray.core.transport.internet.quic.Config.header:type_name -> google.protobuf.Any
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_transport_internet_quic_config_proto_init() }
func file_transport_internet_quic_config_proto_init() {
	if File_transport_internet_quic_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_transport_internet_quic_config_proto_rawDesc), len(file_transport_internet_quic_config_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transport_internet_quic_config_proto_goTypes,
		DependencyIndexes: file_transport_internet_quic_config_proto_depIdxs,
		MessageInfos:      file_transport_internet_quic_config_proto_msgTypes,
	}.Build()
	File_transport_internet_quic_config_proto = out.File
	file_transport_internet_quic_config_proto_goTypes = nil
	file_transport_internet_quic_config_proto_depIdxs = nil
}
