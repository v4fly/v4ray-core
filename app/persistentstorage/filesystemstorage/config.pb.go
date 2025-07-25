package filesystemstorage

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

type StateStorageRoot int32

const (
	StateStorageRoot_WorkDir StateStorageRoot = 0
)

// Enum value maps for StateStorageRoot.
var (
	StateStorageRoot_name = map[int32]string{
		0: "WorkDir",
	}
	StateStorageRoot_value = map[string]int32{
		"WorkDir": 0,
	}
)

func (x StateStorageRoot) Enum() *StateStorageRoot {
	p := new(StateStorageRoot)
	*p = x
	return p
}

func (x StateStorageRoot) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StateStorageRoot) Descriptor() protoreflect.EnumDescriptor {
	return file_app_persistentstorage_filesystemstorage_config_proto_enumTypes[0].Descriptor()
}

func (StateStorageRoot) Type() protoreflect.EnumType {
	return &file_app_persistentstorage_filesystemstorage_config_proto_enumTypes[0]
}

func (x StateStorageRoot) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StateStorageRoot.Descriptor instead.
func (StateStorageRoot) EnumDescriptor() ([]byte, []int) {
	return file_app_persistentstorage_filesystemstorage_config_proto_rawDescGZIP(), []int{0}
}

type Config struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	StateStorageRoot StateStorageRoot       `protobuf:"varint,1,opt,name=state_storage_root,json=stateStorageRoot,proto3,enum=v2ray.core.app.persistentstorage.filesystemstorage.StateStorageRoot" json:"state_storage_root,omitempty"`
	InstanceName     string                 `protobuf:"bytes,4,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	Protojson        bool                   `protobuf:"varint,5,opt,name=protojson,proto3" json:"protojson,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *Config) Reset() {
	*x = Config{}
	mi := &file_app_persistentstorage_filesystemstorage_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_app_persistentstorage_filesystemstorage_config_proto_msgTypes[0]
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
	return file_app_persistentstorage_filesystemstorage_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetStateStorageRoot() StateStorageRoot {
	if x != nil {
		return x.StateStorageRoot
	}
	return StateStorageRoot_WorkDir
}

func (x *Config) GetInstanceName() string {
	if x != nil {
		return x.InstanceName
	}
	return ""
}

func (x *Config) GetProtojson() bool {
	if x != nil {
		return x.Protojson
	}
	return false
}

var File_app_persistentstorage_filesystemstorage_config_proto protoreflect.FileDescriptor

const file_app_persistentstorage_filesystemstorage_config_proto_rawDesc = "" +
	"\n" +
	"4app/persistentstorage/filesystemstorage/config.proto\x122v2ray.core.app.persistentstorage.filesystemstorage\x1a common/protoext/extensions.proto\"\xe1\x01\n" +
	"\x06Config\x12r\n" +
	"\x12state_storage_root\x18\x01 \x01(\x0e2D.v2ray.core.app.persistentstorage.filesystemstorage.StateStorageRootR\x10stateStorageRoot\x12#\n" +
	"\rinstance_name\x18\x04 \x01(\tR\finstanceName\x12\x1c\n" +
	"\tprotojson\x18\x05 \x01(\bR\tprotojson: \x82\xb5\x18\x1c\n" +
	"\aservice\x12\x11filesystemstorage*\x1f\n" +
	"\x10StateStorageRoot\x12\v\n" +
	"\aWorkDir\x10\x00B\xb3\x01\n" +
	"2com.v2ray.core.persistentstorage.filesystemstorageP\x01ZFgithub.com/v4fly/v4ray-core/v0/app/persistentstorage/filesystemstorage\xaa\x022V2Ray.Core.App.Persistentstorage.Filesystemstorageb\x06proto3"

var (
	file_app_persistentstorage_filesystemstorage_config_proto_rawDescOnce sync.Once
	file_app_persistentstorage_filesystemstorage_config_proto_rawDescData []byte
)

func file_app_persistentstorage_filesystemstorage_config_proto_rawDescGZIP() []byte {
	file_app_persistentstorage_filesystemstorage_config_proto_rawDescOnce.Do(func() {
		file_app_persistentstorage_filesystemstorage_config_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_app_persistentstorage_filesystemstorage_config_proto_rawDesc), len(file_app_persistentstorage_filesystemstorage_config_proto_rawDesc)))
	})
	return file_app_persistentstorage_filesystemstorage_config_proto_rawDescData
}

var file_app_persistentstorage_filesystemstorage_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_app_persistentstorage_filesystemstorage_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_app_persistentstorage_filesystemstorage_config_proto_goTypes = []any{
	(StateStorageRoot)(0), // 0: v2ray.core.app.persistentstorage.filesystemstorage.StateStorageRoot
	(*Config)(nil),        // 1: v2ray.core.app.persistentstorage.filesystemstorage.Config
}
var file_app_persistentstorage_filesystemstorage_config_proto_depIdxs = []int32{
	0, // 0: v2ray.core.app.persistentstorage.filesystemstorage.Config.state_storage_root:type_name -> v2ray.core.app.persistentstorage.filesystemstorage.StateStorageRoot
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_app_persistentstorage_filesystemstorage_config_proto_init() }
func file_app_persistentstorage_filesystemstorage_config_proto_init() {
	if File_app_persistentstorage_filesystemstorage_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_app_persistentstorage_filesystemstorage_config_proto_rawDesc), len(file_app_persistentstorage_filesystemstorage_config_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_persistentstorage_filesystemstorage_config_proto_goTypes,
		DependencyIndexes: file_app_persistentstorage_filesystemstorage_config_proto_depIdxs,
		EnumInfos:         file_app_persistentstorage_filesystemstorage_config_proto_enumTypes,
		MessageInfos:      file_app_persistentstorage_filesystemstorage_config_proto_msgTypes,
	}.Build()
	File_app_persistentstorage_filesystemstorage_config_proto = out.File
	file_app_persistentstorage_filesystemstorage_config_proto_goTypes = nil
	file_app_persistentstorage_filesystemstorage_config_proto_depIdxs = nil
}
