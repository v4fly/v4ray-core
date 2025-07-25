package freedom

import (
	protocol "github.com/v4fly/v4ray-core/v0/common/protocol"
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

type ProtocolReplacement int32

const (
	ProtocolReplacement_IDENTITY  ProtocolReplacement = 0
	ProtocolReplacement_FORCE_TCP ProtocolReplacement = 1
	ProtocolReplacement_FORCE_UDP ProtocolReplacement = 2
)

// Enum value maps for ProtocolReplacement.
var (
	ProtocolReplacement_name = map[int32]string{
		0: "IDENTITY",
		1: "FORCE_TCP",
		2: "FORCE_UDP",
	}
	ProtocolReplacement_value = map[string]int32{
		"IDENTITY":  0,
		"FORCE_TCP": 1,
		"FORCE_UDP": 2,
	}
)

func (x ProtocolReplacement) Enum() *ProtocolReplacement {
	p := new(ProtocolReplacement)
	*p = x
	return p
}

func (x ProtocolReplacement) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProtocolReplacement) Descriptor() protoreflect.EnumDescriptor {
	return file_proxy_freedom_config_proto_enumTypes[0].Descriptor()
}

func (ProtocolReplacement) Type() protoreflect.EnumType {
	return &file_proxy_freedom_config_proto_enumTypes[0]
}

func (x ProtocolReplacement) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProtocolReplacement.Descriptor instead.
func (ProtocolReplacement) EnumDescriptor() ([]byte, []int) {
	return file_proxy_freedom_config_proto_rawDescGZIP(), []int{0}
}

type Config_DomainStrategy int32

const (
	Config_AS_IS   Config_DomainStrategy = 0
	Config_USE_IP  Config_DomainStrategy = 1
	Config_USE_IP4 Config_DomainStrategy = 2
	Config_USE_IP6 Config_DomainStrategy = 3
)

// Enum value maps for Config_DomainStrategy.
var (
	Config_DomainStrategy_name = map[int32]string{
		0: "AS_IS",
		1: "USE_IP",
		2: "USE_IP4",
		3: "USE_IP6",
	}
	Config_DomainStrategy_value = map[string]int32{
		"AS_IS":   0,
		"USE_IP":  1,
		"USE_IP4": 2,
		"USE_IP6": 3,
	}
)

func (x Config_DomainStrategy) Enum() *Config_DomainStrategy {
	p := new(Config_DomainStrategy)
	*p = x
	return p
}

func (x Config_DomainStrategy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Config_DomainStrategy) Descriptor() protoreflect.EnumDescriptor {
	return file_proxy_freedom_config_proto_enumTypes[1].Descriptor()
}

func (Config_DomainStrategy) Type() protoreflect.EnumType {
	return &file_proxy_freedom_config_proto_enumTypes[1]
}

func (x Config_DomainStrategy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Config_DomainStrategy.Descriptor instead.
func (Config_DomainStrategy) EnumDescriptor() ([]byte, []int) {
	return file_proxy_freedom_config_proto_rawDescGZIP(), []int{1, 0}
}

type DestinationOverride struct {
	state         protoimpl.MessageState   `protogen:"open.v1"`
	Server        *protocol.ServerEndpoint `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DestinationOverride) Reset() {
	*x = DestinationOverride{}
	mi := &file_proxy_freedom_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DestinationOverride) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestinationOverride) ProtoMessage() {}

func (x *DestinationOverride) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_freedom_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestinationOverride.ProtoReflect.Descriptor instead.
func (*DestinationOverride) Descriptor() ([]byte, []int) {
	return file_proxy_freedom_config_proto_rawDescGZIP(), []int{0}
}

func (x *DestinationOverride) GetServer() *protocol.ServerEndpoint {
	if x != nil {
		return x.Server
	}
	return nil
}

type Config struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	DomainStrategy Config_DomainStrategy  `protobuf:"varint,1,opt,name=domain_strategy,json=domainStrategy,proto3,enum=v2ray.core.proxy.freedom.Config_DomainStrategy" json:"domain_strategy,omitempty"`
	// Deprecated: Marked as deprecated in proxy/freedom/config.proto.
	Timeout             uint32               `protobuf:"varint,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
	DestinationOverride *DestinationOverride `protobuf:"bytes,3,opt,name=destination_override,json=destinationOverride,proto3" json:"destination_override,omitempty"`
	UserLevel           uint32               `protobuf:"varint,4,opt,name=user_level,json=userLevel,proto3" json:"user_level,omitempty"`
	ProtocolReplacement ProtocolReplacement  `protobuf:"varint,5,opt,name=protocol_replacement,json=protocolReplacement,proto3,enum=v2ray.core.proxy.freedom.ProtocolReplacement" json:"protocol_replacement,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *Config) Reset() {
	*x = Config{}
	mi := &file_proxy_freedom_config_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_freedom_config_proto_msgTypes[1]
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
	return file_proxy_freedom_config_proto_rawDescGZIP(), []int{1}
}

func (x *Config) GetDomainStrategy() Config_DomainStrategy {
	if x != nil {
		return x.DomainStrategy
	}
	return Config_AS_IS
}

// Deprecated: Marked as deprecated in proxy/freedom/config.proto.
func (x *Config) GetTimeout() uint32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *Config) GetDestinationOverride() *DestinationOverride {
	if x != nil {
		return x.DestinationOverride
	}
	return nil
}

func (x *Config) GetUserLevel() uint32 {
	if x != nil {
		return x.UserLevel
	}
	return 0
}

func (x *Config) GetProtocolReplacement() ProtocolReplacement {
	if x != nil {
		return x.ProtocolReplacement
	}
	return ProtocolReplacement_IDENTITY
}

type SimplifiedConfig struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	DestinationOverride *DestinationOverride   `protobuf:"bytes,3,opt,name=destination_override,json=destinationOverride,proto3" json:"destination_override,omitempty"`
	ProtocolReplacement ProtocolReplacement    `protobuf:"varint,5,opt,name=protocol_replacement,json=protocolReplacement,proto3,enum=v2ray.core.proxy.freedom.ProtocolReplacement" json:"protocol_replacement,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *SimplifiedConfig) Reset() {
	*x = SimplifiedConfig{}
	mi := &file_proxy_freedom_config_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SimplifiedConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimplifiedConfig) ProtoMessage() {}

func (x *SimplifiedConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_freedom_config_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimplifiedConfig.ProtoReflect.Descriptor instead.
func (*SimplifiedConfig) Descriptor() ([]byte, []int) {
	return file_proxy_freedom_config_proto_rawDescGZIP(), []int{2}
}

func (x *SimplifiedConfig) GetDestinationOverride() *DestinationOverride {
	if x != nil {
		return x.DestinationOverride
	}
	return nil
}

func (x *SimplifiedConfig) GetProtocolReplacement() ProtocolReplacement {
	if x != nil {
		return x.ProtocolReplacement
	}
	return ProtocolReplacement_IDENTITY
}

var File_proxy_freedom_config_proto protoreflect.FileDescriptor

const file_proxy_freedom_config_proto_rawDesc = "" +
	"\n" +
	"\x1aproxy/freedom/config.proto\x12\x18v2ray.core.proxy.freedom\x1a!common/protocol/server_spec.proto\x1a common/protoext/extensions.proto\"Y\n" +
	"\x13DestinationOverride\x12B\n" +
	"\x06server\x18\x01 \x01(\v2*.v2ray.core.common.protocol.ServerEndpointR\x06server\"\xa6\x03\n" +
	"\x06Config\x12X\n" +
	"\x0fdomain_strategy\x18\x01 \x01(\x0e2/.v2ray.core.proxy.freedom.Config.DomainStrategyR\x0edomainStrategy\x12\x1c\n" +
	"\atimeout\x18\x02 \x01(\rB\x02\x18\x01R\atimeout\x12`\n" +
	"\x14destination_override\x18\x03 \x01(\v2-.v2ray.core.proxy.freedom.DestinationOverrideR\x13destinationOverride\x12\x1d\n" +
	"\n" +
	"user_level\x18\x04 \x01(\rR\tuserLevel\x12`\n" +
	"\x14protocol_replacement\x18\x05 \x01(\x0e2-.v2ray.core.proxy.freedom.ProtocolReplacementR\x13protocolReplacement\"A\n" +
	"\x0eDomainStrategy\x12\t\n" +
	"\x05AS_IS\x10\x00\x12\n" +
	"\n" +
	"\x06USE_IP\x10\x01\x12\v\n" +
	"\aUSE_IP4\x10\x02\x12\v\n" +
	"\aUSE_IP6\x10\x03\"\xef\x01\n" +
	"\x10SimplifiedConfig\x12`\n" +
	"\x14destination_override\x18\x03 \x01(\v2-.v2ray.core.proxy.freedom.DestinationOverrideR\x13destinationOverride\x12`\n" +
	"\x14protocol_replacement\x18\x05 \x01(\x0e2-.v2ray.core.proxy.freedom.ProtocolReplacementR\x13protocolReplacement:\x17\x82\xb5\x18\x13\n" +
	"\boutbound\x12\afreedom*A\n" +
	"\x13ProtocolReplacement\x12\f\n" +
	"\bIDENTITY\x10\x00\x12\r\n" +
	"\tFORCE_TCP\x10\x01\x12\r\n" +
	"\tFORCE_UDP\x10\x02Bi\n" +
	"\x1ccom.v2ray.core.proxy.freedomP\x01Z,github.com/v4fly/v4ray-core/v0/proxy/freedom\xaa\x02\x18V2Ray.Core.Proxy.Freedomb\x06proto3"

var (
	file_proxy_freedom_config_proto_rawDescOnce sync.Once
	file_proxy_freedom_config_proto_rawDescData []byte
)

func file_proxy_freedom_config_proto_rawDescGZIP() []byte {
	file_proxy_freedom_config_proto_rawDescOnce.Do(func() {
		file_proxy_freedom_config_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proxy_freedom_config_proto_rawDesc), len(file_proxy_freedom_config_proto_rawDesc)))
	})
	return file_proxy_freedom_config_proto_rawDescData
}

var file_proxy_freedom_config_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proxy_freedom_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proxy_freedom_config_proto_goTypes = []any{
	(ProtocolReplacement)(0),        // 0: v2ray.core.proxy.freedom.ProtocolReplacement
	(Config_DomainStrategy)(0),      // 1: v2ray.core.proxy.freedom.Config.DomainStrategy
	(*DestinationOverride)(nil),     // 2: v2ray.core.proxy.freedom.DestinationOverride
	(*Config)(nil),                  // 3: v2ray.core.proxy.freedom.Config
	(*SimplifiedConfig)(nil),        // 4: v2ray.core.proxy.freedom.SimplifiedConfig
	(*protocol.ServerEndpoint)(nil), // 5: v2ray.core.common.protocol.ServerEndpoint
}
var file_proxy_freedom_config_proto_depIdxs = []int32{
	5, // 0: v2ray.core.proxy.freedom.DestinationOverride.server:type_name -> v2ray.core.common.protocol.ServerEndpoint
	1, // 1: v2ray.core.proxy.freedom.Config.domain_strategy:type_name -> v2ray.core.proxy.freedom.Config.DomainStrategy
	2, // 2: v2ray.core.proxy.freedom.Config.destination_override:type_name -> v2ray.core.proxy.freedom.DestinationOverride
	0, // 3: v2ray.core.proxy.freedom.Config.protocol_replacement:type_name -> v2ray.core.proxy.freedom.ProtocolReplacement
	2, // 4: v2ray.core.proxy.freedom.SimplifiedConfig.destination_override:type_name -> v2ray.core.proxy.freedom.DestinationOverride
	0, // 5: v2ray.core.proxy.freedom.SimplifiedConfig.protocol_replacement:type_name -> v2ray.core.proxy.freedom.ProtocolReplacement
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proxy_freedom_config_proto_init() }
func file_proxy_freedom_config_proto_init() {
	if File_proxy_freedom_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proxy_freedom_config_proto_rawDesc), len(file_proxy_freedom_config_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proxy_freedom_config_proto_goTypes,
		DependencyIndexes: file_proxy_freedom_config_proto_depIdxs,
		EnumInfos:         file_proxy_freedom_config_proto_enumTypes,
		MessageInfos:      file_proxy_freedom_config_proto_msgTypes,
	}.Build()
	File_proxy_freedom_config_proto = out.File
	file_proxy_freedom_config_proto_goTypes = nil
	file_proxy_freedom_config_proto_depIdxs = nil
}
