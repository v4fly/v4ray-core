package log

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

type Severity int32

const (
	Severity_Unknown Severity = 0
	Severity_Error   Severity = 1
	Severity_Warning Severity = 2
	Severity_Info    Severity = 3
	Severity_Debug   Severity = 4
)

// Enum value maps for Severity.
var (
	Severity_name = map[int32]string{
		0: "Unknown",
		1: "Error",
		2: "Warning",
		3: "Info",
		4: "Debug",
	}
	Severity_value = map[string]int32{
		"Unknown": 0,
		"Error":   1,
		"Warning": 2,
		"Info":    3,
		"Debug":   4,
	}
)

func (x Severity) Enum() *Severity {
	p := new(Severity)
	*p = x
	return p
}

func (x Severity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Severity) Descriptor() protoreflect.EnumDescriptor {
	return file_common_log_log_proto_enumTypes[0].Descriptor()
}

func (Severity) Type() protoreflect.EnumType {
	return &file_common_log_log_proto_enumTypes[0]
}

func (x Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Severity.Descriptor instead.
func (Severity) EnumDescriptor() ([]byte, []int) {
	return file_common_log_log_proto_rawDescGZIP(), []int{0}
}

var File_common_log_log_proto protoreflect.FileDescriptor

const file_common_log_log_proto_rawDesc = "" +
	"\n" +
	"\x14common/log/log.proto\x12\x15v2ray.core.common.log*D\n" +
	"\bSeverity\x12\v\n" +
	"\aUnknown\x10\x00\x12\t\n" +
	"\x05Error\x10\x01\x12\v\n" +
	"\aWarning\x10\x02\x12\b\n" +
	"\x04Info\x10\x03\x12\t\n" +
	"\x05Debug\x10\x04B`\n" +
	"\x19com.v2ray.core.common.logP\x01Z)github.com/v4fly/v4ray-core/v0/common/log\xaa\x02\x15V2Ray.Core.Common.Logb\x06proto3"

var (
	file_common_log_log_proto_rawDescOnce sync.Once
	file_common_log_log_proto_rawDescData []byte
)

func file_common_log_log_proto_rawDescGZIP() []byte {
	file_common_log_log_proto_rawDescOnce.Do(func() {
		file_common_log_log_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_common_log_log_proto_rawDesc), len(file_common_log_log_proto_rawDesc)))
	})
	return file_common_log_log_proto_rawDescData
}

var file_common_log_log_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_log_log_proto_goTypes = []any{
	(Severity)(0), // 0: v2ray.core.common.log.Severity
}
var file_common_log_log_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_log_log_proto_init() }
func file_common_log_log_proto_init() {
	if File_common_log_log_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_common_log_log_proto_rawDesc), len(file_common_log_log_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_log_log_proto_goTypes,
		DependencyIndexes: file_common_log_log_proto_depIdxs,
		EnumInfos:         file_common_log_log_proto_enumTypes,
	}.Build()
	File_common_log_log_proto = out.File
	file_common_log_log_proto_goTypes = nil
	file_common_log_log_proto_depIdxs = nil
}
