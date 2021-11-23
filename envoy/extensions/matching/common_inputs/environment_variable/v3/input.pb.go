// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: envoy/extensions/matching/common_inputs/environment_variable/v3/input.proto

package envoy_extensions_matching_common_inputs_environment_variable_v3

import (
	_ "github.com/cncf/xds/go/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// Reads an environment variable to provide an input for matching.
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the environment variable to read from.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_matching_common_inputs_environment_variable_v3_input_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_matching_common_inputs_environment_variable_v3_input_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_envoy_extensions_matching_common_inputs_environment_variable_v3_input_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_envoy_extensions_matching_common_inputs_environment_variable_v3_input_proto protoreflect.FileDescriptor

var file_envoy_extensions_matching_common_inputs_environment_variable_v3_input_proto_rawDesc = []byte{
	0x0a, 0x4b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x76,
	0x33, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3f, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x33, 0x1a, 0x1d,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x65, 0x0a,
	0x4d, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x33, 0x42, 0x0a,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1,
	0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_matching_common_inputs_environment_variable_v3_input_proto_rawDescOnce sync.Once
	file_envoy_extensions_matching_common_inputs_environment_variable_v3_input_proto_rawDescData = file_envoy_extensions_matching_common_inputs_environment_variable_v3_input_proto_rawDesc
)

func file_envoy_extensions_matching_common_inputs_environment_variable_v3_input_proto_rawDescGZIP() []byte {
	file_envoy_extensions_matching_common_inputs_environment_variable_v3_input_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_matching_common_inputs_environment_variable_v3_input_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_matching_common_inputs_environment_variable_v3_input_proto_rawDescData)
	})
	return file_envoy_extensions_matching_common_inputs_environment_variable_v3_input_proto_rawDescData
}

var file_envoy_extensions_matching_common_inputs_environment_variable_v3_input_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_matching_common_inputs_environment_variable_v3_input_proto_goTypes = []interface{}{
	(*Config)(nil), // 0: envoy.extensions.matching.common_inputs.environment_variable.v3.Config
}
var file_envoy_extensions_matching_common_inputs_environment_variable_v3_input_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_envoy_extensions_matching_common_inputs_environment_variable_v3_input_proto_init() }
func file_envoy_extensions_matching_common_inputs_environment_variable_v3_input_proto_init() {
	if File_envoy_extensions_matching_common_inputs_environment_variable_v3_input_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_matching_common_inputs_environment_variable_v3_input_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
			RawDescriptor: file_envoy_extensions_matching_common_inputs_environment_variable_v3_input_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_matching_common_inputs_environment_variable_v3_input_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_matching_common_inputs_environment_variable_v3_input_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_matching_common_inputs_environment_variable_v3_input_proto_msgTypes,
	}.Build()
	File_envoy_extensions_matching_common_inputs_environment_variable_v3_input_proto = out.File
	file_envoy_extensions_matching_common_inputs_environment_variable_v3_input_proto_rawDesc = nil
	file_envoy_extensions_matching_common_inputs_environment_variable_v3_input_proto_goTypes = nil
	file_envoy_extensions_matching_common_inputs_environment_variable_v3_input_proto_depIdxs = nil
}