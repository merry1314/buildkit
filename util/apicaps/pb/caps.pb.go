// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.11.4
// source: github.com/moby/buildkit/util/apicaps/pb/caps.proto

package moby_buildkit_v1_apicaps

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

// APICap defines a capability supported by the service
type APICap struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	ID                  string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Enabled             bool                   `protobuf:"varint,2,opt,name=Enabled,proto3" json:"Enabled,omitempty"`
	Deprecated          bool                   `protobuf:"varint,3,opt,name=Deprecated,proto3" json:"Deprecated,omitempty"`                  // Unused. May be used for warnings in the future
	DisabledReason      string                 `protobuf:"bytes,4,opt,name=DisabledReason,proto3" json:"DisabledReason,omitempty"`           // Reason key for detection code
	DisabledReasonMsg   string                 `protobuf:"bytes,5,opt,name=DisabledReasonMsg,proto3" json:"DisabledReasonMsg,omitempty"`     // Message to the user
	DisabledAlternative string                 `protobuf:"bytes,6,opt,name=DisabledAlternative,proto3" json:"DisabledAlternative,omitempty"` // Identifier that updated client could catch.
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *APICap) Reset() {
	*x = APICap{}
	mi := &file_github_com_moby_buildkit_util_apicaps_pb_caps_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *APICap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APICap) ProtoMessage() {}

func (x *APICap) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_moby_buildkit_util_apicaps_pb_caps_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APICap.ProtoReflect.Descriptor instead.
func (*APICap) Descriptor() ([]byte, []int) {
	return file_github_com_moby_buildkit_util_apicaps_pb_caps_proto_rawDescGZIP(), []int{0}
}

func (x *APICap) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *APICap) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *APICap) GetDeprecated() bool {
	if x != nil {
		return x.Deprecated
	}
	return false
}

func (x *APICap) GetDisabledReason() string {
	if x != nil {
		return x.DisabledReason
	}
	return ""
}

func (x *APICap) GetDisabledReasonMsg() string {
	if x != nil {
		return x.DisabledReasonMsg
	}
	return ""
}

func (x *APICap) GetDisabledAlternative() string {
	if x != nil {
		return x.DisabledAlternative
	}
	return ""
}

var File_github_com_moby_buildkit_util_apicaps_pb_caps_proto protoreflect.FileDescriptor

const file_github_com_moby_buildkit_util_apicaps_pb_caps_proto_rawDesc = "" +
	"\n" +
	"3github.com/moby/buildkit/util/apicaps/pb/caps.proto\x12\x18moby.buildkit.v1.apicaps\"\xda\x01\n" +
	"\x06APICap\x12\x0e\n" +
	"\x02ID\x18\x01 \x01(\tR\x02ID\x12\x18\n" +
	"\aEnabled\x18\x02 \x01(\bR\aEnabled\x12\x1e\n" +
	"\n" +
	"Deprecated\x18\x03 \x01(\bR\n" +
	"Deprecated\x12&\n" +
	"\x0eDisabledReason\x18\x04 \x01(\tR\x0eDisabledReason\x12,\n" +
	"\x11DisabledReasonMsg\x18\x05 \x01(\tR\x11DisabledReasonMsg\x120\n" +
	"\x13DisabledAlternative\x18\x06 \x01(\tR\x13DisabledAlternativeBCZAgithub.com/moby/buildkit/util/apicaps/pb;moby_buildkit_v1_apicapsb\x06proto3"

var (
	file_github_com_moby_buildkit_util_apicaps_pb_caps_proto_rawDescOnce sync.Once
	file_github_com_moby_buildkit_util_apicaps_pb_caps_proto_rawDescData []byte
)

func file_github_com_moby_buildkit_util_apicaps_pb_caps_proto_rawDescGZIP() []byte {
	file_github_com_moby_buildkit_util_apicaps_pb_caps_proto_rawDescOnce.Do(func() {
		file_github_com_moby_buildkit_util_apicaps_pb_caps_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_github_com_moby_buildkit_util_apicaps_pb_caps_proto_rawDesc), len(file_github_com_moby_buildkit_util_apicaps_pb_caps_proto_rawDesc)))
	})
	return file_github_com_moby_buildkit_util_apicaps_pb_caps_proto_rawDescData
}

var file_github_com_moby_buildkit_util_apicaps_pb_caps_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_moby_buildkit_util_apicaps_pb_caps_proto_goTypes = []any{
	(*APICap)(nil), // 0: moby.buildkit.v1.apicaps.APICap
}
var file_github_com_moby_buildkit_util_apicaps_pb_caps_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_moby_buildkit_util_apicaps_pb_caps_proto_init() }
func file_github_com_moby_buildkit_util_apicaps_pb_caps_proto_init() {
	if File_github_com_moby_buildkit_util_apicaps_pb_caps_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_github_com_moby_buildkit_util_apicaps_pb_caps_proto_rawDesc), len(file_github_com_moby_buildkit_util_apicaps_pb_caps_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_moby_buildkit_util_apicaps_pb_caps_proto_goTypes,
		DependencyIndexes: file_github_com_moby_buildkit_util_apicaps_pb_caps_proto_depIdxs,
		MessageInfos:      file_github_com_moby_buildkit_util_apicaps_pb_caps_proto_msgTypes,
	}.Build()
	File_github_com_moby_buildkit_util_apicaps_pb_caps_proto = out.File
	file_github_com_moby_buildkit_util_apicaps_pb_caps_proto_goTypes = nil
	file_github_com_moby_buildkit_util_apicaps_pb_caps_proto_depIdxs = nil
}
