package subscription

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

type ImportSource struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	Name                 string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Url                  string                 `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	TagPrefix            string                 `protobuf:"bytes,3,opt,name=tag_prefix,json=tagPrefix,proto3" json:"tag_prefix,omitempty"`
	ImportUsingTag       string                 `protobuf:"bytes,4,opt,name=import_using_tag,json=importUsingTag,proto3" json:"import_using_tag,omitempty"`
	DefaultExpireSeconds uint64                 `protobuf:"varint,5,opt,name=default_expire_seconds,json=defaultExpireSeconds,proto3" json:"default_expire_seconds,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *ImportSource) Reset() {
	*x = ImportSource{}
	mi := &file_app_subscription_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImportSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportSource) ProtoMessage() {}

func (x *ImportSource) ProtoReflect() protoreflect.Message {
	mi := &file_app_subscription_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportSource.ProtoReflect.Descriptor instead.
func (*ImportSource) Descriptor() ([]byte, []int) {
	return file_app_subscription_config_proto_rawDescGZIP(), []int{0}
}

func (x *ImportSource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ImportSource) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ImportSource) GetTagPrefix() string {
	if x != nil {
		return x.TagPrefix
	}
	return ""
}

func (x *ImportSource) GetImportUsingTag() string {
	if x != nil {
		return x.ImportUsingTag
	}
	return ""
}

func (x *ImportSource) GetDefaultExpireSeconds() uint64 {
	if x != nil {
		return x.DefaultExpireSeconds
	}
	return 0
}

// Config is the settings for Subscription Manager.
type Config struct {
	state                         protoimpl.MessageState `protogen:"open.v1"`
	Imports                       []*ImportSource        `protobuf:"bytes,1,rep,name=imports,proto3" json:"imports,omitempty"`
	NonnativeConverterOverlay     []byte                 `protobuf:"bytes,2,opt,name=nonnative_converter_overlay,json=nonnativeConverterOverlay,proto3" json:"nonnative_converter_overlay,omitempty"`
	NonnativeConverterOverlayFile string                 `protobuf:"bytes,96002,opt,name=nonnative_converter_overlay_file,json=nonnativeConverterOverlayFile,proto3" json:"nonnative_converter_overlay_file,omitempty"`
	Persistence                   bool                   `protobuf:"varint,3,opt,name=persistence,proto3" json:"persistence,omitempty"`
	unknownFields                 protoimpl.UnknownFields
	sizeCache                     protoimpl.SizeCache
}

func (x *Config) Reset() {
	*x = Config{}
	mi := &file_app_subscription_config_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_app_subscription_config_proto_msgTypes[1]
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
	return file_app_subscription_config_proto_rawDescGZIP(), []int{1}
}

func (x *Config) GetImports() []*ImportSource {
	if x != nil {
		return x.Imports
	}
	return nil
}

func (x *Config) GetNonnativeConverterOverlay() []byte {
	if x != nil {
		return x.NonnativeConverterOverlay
	}
	return nil
}

func (x *Config) GetNonnativeConverterOverlayFile() string {
	if x != nil {
		return x.NonnativeConverterOverlayFile
	}
	return ""
}

func (x *Config) GetPersistence() bool {
	if x != nil {
		return x.Persistence
	}
	return false
}

var File_app_subscription_config_proto protoreflect.FileDescriptor

const file_app_subscription_config_proto_rawDesc = "" +
	"\n" +
	"\x1dapp/subscription/config.proto\x12\x1bv2ray.core.app.subscription\x1a common/protoext/extensions.proto\"\xb3\x01\n" +
	"\fImportSource\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x10\n" +
	"\x03url\x18\x02 \x01(\tR\x03url\x12\x1d\n" +
	"\n" +
	"tag_prefix\x18\x03 \x01(\tR\ttagPrefix\x12(\n" +
	"\x10import_using_tag\x18\x04 \x01(\tR\x0eimportUsingTag\x124\n" +
	"\x16default_expire_seconds\x18\x05 \x01(\x04R\x14defaultExpireSeconds\"\xba\x02\n" +
	"\x06Config\x12C\n" +
	"\aimports\x18\x01 \x03(\v2).v2ray.core.app.subscription.ImportSourceR\aimports\x12>\n" +
	"\x1bnonnative_converter_overlay\x18\x02 \x01(\fR\x19nonnativeConverterOverlay\x12l\n" +
	" nonnative_converter_overlay_file\x18\x82\xee\x05 \x01(\tB!\x82\xb5\x18\x1d\"\x1bnonnative_converter_overlayR\x1dnonnativeConverterOverlayFile\x12 \n" +
	"\vpersistence\x18\x03 \x01(\bR\vpersistence:\x1b\x82\xb5\x18\x17\n" +
	"\aservice\x12\fsubscriptionBr\n" +
	"\x1fcom.v2ray.core.app.subscriptionP\x01Z/github.com/v4fly/v4ray-core/v0/app/subscription\xaa\x02\x1bV2Ray.Core.App.Subscriptionb\x06proto3"

var (
	file_app_subscription_config_proto_rawDescOnce sync.Once
	file_app_subscription_config_proto_rawDescData []byte
)

func file_app_subscription_config_proto_rawDescGZIP() []byte {
	file_app_subscription_config_proto_rawDescOnce.Do(func() {
		file_app_subscription_config_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_app_subscription_config_proto_rawDesc), len(file_app_subscription_config_proto_rawDesc)))
	})
	return file_app_subscription_config_proto_rawDescData
}

var file_app_subscription_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_app_subscription_config_proto_goTypes = []any{
	(*ImportSource)(nil), // 0: v2ray.core.app.subscription.ImportSource
	(*Config)(nil),       // 1: v2ray.core.app.subscription.Config
}
var file_app_subscription_config_proto_depIdxs = []int32{
	0, // 0: v2ray.core.app.subscription.Config.imports:type_name -> v2ray.core.app.subscription.ImportSource
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_app_subscription_config_proto_init() }
func file_app_subscription_config_proto_init() {
	if File_app_subscription_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_app_subscription_config_proto_rawDesc), len(file_app_subscription_config_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_subscription_config_proto_goTypes,
		DependencyIndexes: file_app_subscription_config_proto_depIdxs,
		MessageInfos:      file_app_subscription_config_proto_msgTypes,
	}.Build()
	File_app_subscription_config_proto = out.File
	file_app_subscription_config_proto_goTypes = nil
	file_app_subscription_config_proto_depIdxs = nil
}
