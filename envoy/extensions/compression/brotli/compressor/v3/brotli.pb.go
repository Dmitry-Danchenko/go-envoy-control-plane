// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: envoy/extensions/compression/brotli/compressor/v3/brotli.proto

package envoy_extensions_compression_brotli_compressor_v3

import (
	_ "github.com/cncf/xds/go/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type Brotli_EncoderMode int32

const (
	Brotli_DEFAULT Brotli_EncoderMode = 0
	Brotli_GENERIC Brotli_EncoderMode = 1
	Brotli_TEXT    Brotli_EncoderMode = 2
	Brotli_FONT    Brotli_EncoderMode = 3
)

// Enum value maps for Brotli_EncoderMode.
var (
	Brotli_EncoderMode_name = map[int32]string{
		0: "DEFAULT",
		1: "GENERIC",
		2: "TEXT",
		3: "FONT",
	}
	Brotli_EncoderMode_value = map[string]int32{
		"DEFAULT": 0,
		"GENERIC": 1,
		"TEXT":    2,
		"FONT":    3,
	}
)

func (x Brotli_EncoderMode) Enum() *Brotli_EncoderMode {
	p := new(Brotli_EncoderMode)
	*p = x
	return p
}

func (x Brotli_EncoderMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Brotli_EncoderMode) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_compression_brotli_compressor_v3_brotli_proto_enumTypes[0].Descriptor()
}

func (Brotli_EncoderMode) Type() protoreflect.EnumType {
	return &file_envoy_extensions_compression_brotli_compressor_v3_brotli_proto_enumTypes[0]
}

func (x Brotli_EncoderMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Brotli_EncoderMode.Descriptor instead.
func (Brotli_EncoderMode) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_compression_brotli_compressor_v3_brotli_proto_rawDescGZIP(), []int{0, 0}
}

// [#next-free-field: 7]
type Brotli struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Value from 0 to 11 that controls the main compression speed-density lever.
	// The higher quality, the slower compression. The default value is 3.
	Quality *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=quality,proto3" json:"quality,omitempty"`
	// A value used to tune encoder for specific input. For more information about modes,
	// please refer to brotli manual: https://brotli.org/encode.html#aa6f
	// This field will be set to "DEFAULT" if not specified.
	EncoderMode Brotli_EncoderMode `protobuf:"varint,2,opt,name=encoder_mode,json=encoderMode,proto3,enum=envoy.extensions.compression.brotli.compressor.v3.Brotli_EncoderMode" json:"encoder_mode,omitempty"`
	// Value from 10 to 24 that represents the base two logarithmic of the compressor's window size.
	// Larger window results in better compression at the expense of memory usage. The default is 18.
	// For more details about this parameter, please refer to brotli manual:
	// https://brotli.org/encode.html#a9a8
	WindowBits *wrappers.UInt32Value `protobuf:"bytes,3,opt,name=window_bits,json=windowBits,proto3" json:"window_bits,omitempty"`
	// Value from 16 to 24 that represents the base two logarithmic of the compressor's input block
	// size. Larger input block results in better compression at the expense of memory usage. The
	// default is 24. For more details about this parameter, please refer to brotli manual:
	// https://brotli.org/encode.html#a9a8
	InputBlockBits *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=input_block_bits,json=inputBlockBits,proto3" json:"input_block_bits,omitempty"`
	// Value for compressor's next output buffer. If not set, defaults to 4096.
	ChunkSize *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=chunk_size,json=chunkSize,proto3" json:"chunk_size,omitempty"`
	// If true, disables "literal context modeling" format feature.
	// This flag is a "decoding-speed vs compression ratio" trade-off.
	DisableLiteralContextModeling bool `protobuf:"varint,6,opt,name=disable_literal_context_modeling,json=disableLiteralContextModeling,proto3" json:"disable_literal_context_modeling,omitempty"`
}

func (x *Brotli) Reset() {
	*x = Brotli{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_compression_brotli_compressor_v3_brotli_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Brotli) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Brotli) ProtoMessage() {}

func (x *Brotli) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_compression_brotli_compressor_v3_brotli_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Brotli.ProtoReflect.Descriptor instead.
func (*Brotli) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_compression_brotli_compressor_v3_brotli_proto_rawDescGZIP(), []int{0}
}

func (x *Brotli) GetQuality() *wrappers.UInt32Value {
	if x != nil {
		return x.Quality
	}
	return nil
}

func (x *Brotli) GetEncoderMode() Brotli_EncoderMode {
	if x != nil {
		return x.EncoderMode
	}
	return Brotli_DEFAULT
}

func (x *Brotli) GetWindowBits() *wrappers.UInt32Value {
	if x != nil {
		return x.WindowBits
	}
	return nil
}

func (x *Brotli) GetInputBlockBits() *wrappers.UInt32Value {
	if x != nil {
		return x.InputBlockBits
	}
	return nil
}

func (x *Brotli) GetChunkSize() *wrappers.UInt32Value {
	if x != nil {
		return x.ChunkSize
	}
	return nil
}

func (x *Brotli) GetDisableLiteralContextModeling() bool {
	if x != nil {
		return x.DisableLiteralContextModeling
	}
	return false
}

var File_envoy_extensions_compression_brotli_compressor_v3_brotli_proto protoreflect.FileDescriptor

var file_envoy_extensions_compression_brotli_compressor_v3_brotli_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x62,
	0x72, 0x6f, 0x74, 0x6c, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72,
	0x2f, 0x76, 0x33, 0x2f, 0x62, 0x72, 0x6f, 0x74, 0x6c, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x31, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x62,
	0x72, 0x6f, 0x74, 0x6c, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72,
	0x2e, 0x76, 0x33, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x04, 0x0a, 0x06,
	0x42, 0x72, 0x6f, 0x74, 0x6c, 0x69, 0x12, 0x3f, 0x0a, 0x07, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x18, 0x0b, 0x52, 0x07,
	0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x72, 0x0a, 0x0c, 0x65, 0x6e, 0x63, 0x6f, 0x64,
	0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x45, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x62, 0x72, 0x6f,
	0x74, 0x6c, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x76,
	0x33, 0x2e, 0x42, 0x72, 0x6f, 0x74, 0x6c, 0x69, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x72,
	0x4d, 0x6f, 0x64, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0b,
	0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x77,
	0x69, 0x6e, 0x64, 0x6f, 0x77, 0x5f, 0x62, 0x69, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x09,
	0xfa, 0x42, 0x06, 0x2a, 0x04, 0x18, 0x18, 0x28, 0x0a, 0x52, 0x0a, 0x77, 0x69, 0x6e, 0x64, 0x6f,
	0x77, 0x42, 0x69, 0x74, 0x73, 0x12, 0x51, 0x0a, 0x10, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x62, 0x69, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x09, 0xfa,
	0x42, 0x06, 0x2a, 0x04, 0x18, 0x18, 0x28, 0x10, 0x52, 0x0e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x69, 0x74, 0x73, 0x12, 0x49, 0x0a, 0x0a, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55,
	0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0c, 0xfa, 0x42, 0x09, 0x2a,
	0x07, 0x18, 0x80, 0x80, 0x04, 0x28, 0x80, 0x20, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x47, 0x0a, 0x20, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6c,
	0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1d, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x22, 0x3b, 0x0a, 0x0b,
	0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x44,
	0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x47, 0x45, 0x4e, 0x45,
	0x52, 0x49, 0x43, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x45, 0x58, 0x54, 0x10, 0x02, 0x12,
	0x08, 0x0a, 0x04, 0x46, 0x4f, 0x4e, 0x54, 0x10, 0x03, 0x42, 0x58, 0x0a, 0x3f, 0x69, 0x6f, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x62, 0x72, 0x6f, 0x74, 0x6c, 0x69, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x33, 0x42, 0x0b, 0x42, 0x72,
	0x6f, 0x74, 0x6c, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06,
	0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_compression_brotli_compressor_v3_brotli_proto_rawDescOnce sync.Once
	file_envoy_extensions_compression_brotli_compressor_v3_brotli_proto_rawDescData = file_envoy_extensions_compression_brotli_compressor_v3_brotli_proto_rawDesc
)

func file_envoy_extensions_compression_brotli_compressor_v3_brotli_proto_rawDescGZIP() []byte {
	file_envoy_extensions_compression_brotli_compressor_v3_brotli_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_compression_brotli_compressor_v3_brotli_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_compression_brotli_compressor_v3_brotli_proto_rawDescData)
	})
	return file_envoy_extensions_compression_brotli_compressor_v3_brotli_proto_rawDescData
}

var file_envoy_extensions_compression_brotli_compressor_v3_brotli_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_envoy_extensions_compression_brotli_compressor_v3_brotli_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_compression_brotli_compressor_v3_brotli_proto_goTypes = []interface{}{
	(Brotli_EncoderMode)(0),      // 0: envoy.extensions.compression.brotli.compressor.v3.Brotli.EncoderMode
	(*Brotli)(nil),               // 1: envoy.extensions.compression.brotli.compressor.v3.Brotli
	(*wrappers.UInt32Value)(nil), // 2: google.protobuf.UInt32Value
}
var file_envoy_extensions_compression_brotli_compressor_v3_brotli_proto_depIdxs = []int32{
	2, // 0: envoy.extensions.compression.brotli.compressor.v3.Brotli.quality:type_name -> google.protobuf.UInt32Value
	0, // 1: envoy.extensions.compression.brotli.compressor.v3.Brotli.encoder_mode:type_name -> envoy.extensions.compression.brotli.compressor.v3.Brotli.EncoderMode
	2, // 2: envoy.extensions.compression.brotli.compressor.v3.Brotli.window_bits:type_name -> google.protobuf.UInt32Value
	2, // 3: envoy.extensions.compression.brotli.compressor.v3.Brotli.input_block_bits:type_name -> google.protobuf.UInt32Value
	2, // 4: envoy.extensions.compression.brotli.compressor.v3.Brotli.chunk_size:type_name -> google.protobuf.UInt32Value
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_envoy_extensions_compression_brotli_compressor_v3_brotli_proto_init() }
func file_envoy_extensions_compression_brotli_compressor_v3_brotli_proto_init() {
	if File_envoy_extensions_compression_brotli_compressor_v3_brotli_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_compression_brotli_compressor_v3_brotli_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Brotli); i {
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
			RawDescriptor: file_envoy_extensions_compression_brotli_compressor_v3_brotli_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_compression_brotli_compressor_v3_brotli_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_compression_brotli_compressor_v3_brotli_proto_depIdxs,
		EnumInfos:         file_envoy_extensions_compression_brotli_compressor_v3_brotli_proto_enumTypes,
		MessageInfos:      file_envoy_extensions_compression_brotli_compressor_v3_brotli_proto_msgTypes,
	}.Build()
	File_envoy_extensions_compression_brotli_compressor_v3_brotli_proto = out.File
	file_envoy_extensions_compression_brotli_compressor_v3_brotli_proto_rawDesc = nil
	file_envoy_extensions_compression_brotli_compressor_v3_brotli_proto_goTypes = nil
	file_envoy_extensions_compression_brotli_compressor_v3_brotli_proto_depIdxs = nil
}