package meek

import (
	_ "github.com/v4fly/v4ray-core/v0/common/protoext"
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
	Url           string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Config) Reset() {
	*x = Config{}
	mi := &file_transport_internet_request_stereotype_meek_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_request_stereotype_meek_config_proto_msgTypes[0]
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
	return file_transport_internet_request_stereotype_meek_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_transport_internet_request_stereotype_meek_config_proto protoreflect.FileDescriptor

const file_transport_internet_request_stereotype_meek_config_proto_rawDesc = "" +
	"\n" +
	"7transport/internet/request/stereotype/meek/config.proto\x125v2ray.core.transport.internet.request.stereotype.meek\x1a common/protoext/extensions.proto\"5\n" +
	"\x06Config\x12\x10\n" +
	"\x03url\x18\x01 \x01(\tR\x03url:\x19\x82\xb5\x18\x15\n" +
	"\ttransport\x12\x04meek\x90\xff)\x01B\xc0\x01\n" +
	"9com.v2ray.core.transport.internet.request.stereotype.meekP\x01ZIgithub.com/v4fly/v4ray-core/v0/transport/internet/request/stereotype/meek\xaa\x025V2Ray.Core.Transport.Internet.Request.Stereotype.Meekb\x06proto3"

var (
	file_transport_internet_request_stereotype_meek_config_proto_rawDescOnce sync.Once
	file_transport_internet_request_stereotype_meek_config_proto_rawDescData []byte
)

func file_transport_internet_request_stereotype_meek_config_proto_rawDescGZIP() []byte {
	file_transport_internet_request_stereotype_meek_config_proto_rawDescOnce.Do(func() {
		file_transport_internet_request_stereotype_meek_config_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_transport_internet_request_stereotype_meek_config_proto_rawDesc), len(file_transport_internet_request_stereotype_meek_config_proto_rawDesc)))
	})
	return file_transport_internet_request_stereotype_meek_config_proto_rawDescData
}

var file_transport_internet_request_stereotype_meek_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_transport_internet_request_stereotype_meek_config_proto_goTypes = []any{
	(*Config)(nil), // 0: v2ray.core.transport.internet.request.stereotype.meek.Config
}
var file_transport_internet_request_stereotype_meek_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_transport_internet_request_stereotype_meek_config_proto_init() }
func file_transport_internet_request_stereotype_meek_config_proto_init() {
	if File_transport_internet_request_stereotype_meek_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_transport_internet_request_stereotype_meek_config_proto_rawDesc), len(file_transport_internet_request_stereotype_meek_config_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transport_internet_request_stereotype_meek_config_proto_goTypes,
		DependencyIndexes: file_transport_internet_request_stereotype_meek_config_proto_depIdxs,
		MessageInfos:      file_transport_internet_request_stereotype_meek_config_proto_msgTypes,
	}.Build()
	File_transport_internet_request_stereotype_meek_config_proto = out.File
	file_transport_internet_request_stereotype_meek_config_proto_goTypes = nil
	file_transport_internet_request_stereotype_meek_config_proto_depIdxs = nil
}
