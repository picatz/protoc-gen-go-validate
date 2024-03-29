// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: example.proto

package example

import (
	_ "github.com/picatz/protoc-gen-go-validate/pkg/validate"
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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName     string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName      string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email         string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Nickname      *string  `protobuf:"bytes,4,opt,name=nickname,proto3,oneof" json:"nickname,omitempty"`
	Team          *string  `protobuf:"bytes,5,opt,name=team,proto3,oneof" json:"team,omitempty"`
	Points        uint64   `protobuf:"varint,6,opt,name=points,proto3" json:"points,omitempty"`
	ExtraPoints   int64    `protobuf:"varint,7,opt,name=extraPoints,proto3" json:"extraPoints,omitempty"`
	Something     float32  `protobuf:"fixed32,8,opt,name=something,proto3" json:"something,omitempty"`
	Key           []byte   `protobuf:"bytes,9,opt,name=key,proto3,oneof" json:"key,omitempty"`
	Friends       []string `protobuf:"bytes,10,rep,name=friends,proto3" json:"friends,omitempty"`
	SomethingElse float64  `protobuf:"fixed64,11,opt,name=somethingElse,proto3" json:"somethingElse,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Request) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Request) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Request) GetNickname() string {
	if x != nil && x.Nickname != nil {
		return *x.Nickname
	}
	return ""
}

func (x *Request) GetTeam() string {
	if x != nil && x.Team != nil {
		return *x.Team
	}
	return ""
}

func (x *Request) GetPoints() uint64 {
	if x != nil {
		return x.Points
	}
	return 0
}

func (x *Request) GetExtraPoints() int64 {
	if x != nil {
		return x.ExtraPoints
	}
	return 0
}

func (x *Request) GetSomething() float32 {
	if x != nil {
		return x.Something
	}
	return 0
}

func (x *Request) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *Request) GetFriends() []string {
	if x != nil {
		return x.Friends
	}
	return nil
}

func (x *Request) GetSomethingElse() float64 {
	if x != nil {
		return x.SomethingElse
	}
	return 0
}

type Request2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email     string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Nickname  *string  `protobuf:"bytes,4,opt,name=nickname,proto3,oneof" json:"nickname,omitempty"`
	Nested    *Request `protobuf:"bytes,5,opt,name=nested,proto3" json:"nested,omitempty"`
}

func (x *Request2) Reset() {
	*x = Request2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request2) ProtoMessage() {}

func (x *Request2) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request2.ProtoReflect.Descriptor instead.
func (*Request2) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{1}
}

func (x *Request2) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Request2) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Request2) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Request2) GetNickname() string {
	if x != nil && x.Nickname != nil {
		return *x.Nickname
	}
	return ""
}

func (x *Request2) GetNested() *Request {
	if x != nil {
		return x.Nested
	}
	return nil
}

type Request3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Something string `protobuf:"bytes,1,opt,name=something,proto3" json:"something,omitempty"`
}

func (x *Request3) Reset() {
	*x = Request3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request3) ProtoMessage() {}

func (x *Request3) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request3.ProtoReflect.Descriptor instead.
func (*Request3) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{2}
}

func (x *Request3) GetSomething() string {
	if x != nil {
		return x.Something
	}
	return ""
}

var File_example_proto protoreflect.FileDescriptor

var file_example_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x04, 0x0a, 0x07, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x16, 0xca, 0x43, 0x13, 0x0a, 0x11, 0x10,
	0x00, 0x18, 0x80, 0x02, 0x32, 0x01, 0x2e, 0x48, 0x00, 0x50, 0x01, 0x58, 0x01, 0x60, 0xff, 0x01,
	0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x09, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x16,
	0xca, 0x43, 0x13, 0x0a, 0x11, 0x10, 0x00, 0x18, 0x80, 0x02, 0x32, 0x01, 0x2e, 0x48, 0x00, 0x50,
	0x01, 0x58, 0x01, 0x60, 0xff, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x33, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x1d, 0xca, 0x43, 0x1a, 0x0a, 0x18, 0x10, 0x00, 0x18, 0x80, 0x02, 0x42, 0x0a, 0x40, 0x67, 0x6d,
	0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x50, 0x01, 0x58, 0x01, 0x60, 0xff, 0x01, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x3d, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1c, 0xca, 0x43, 0x19, 0x0a, 0x17, 0x58, 0x01,
	0x60, 0xff, 0x01, 0x6a, 0x02, 0x5e, 0x2e, 0x72, 0x0c, 0x5b, 0x5b, 0x3a, 0x5e, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x3a, 0x5d, 0x5d, 0x48, 0x00, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0e, 0xca, 0x43, 0x0b, 0x0a, 0x09, 0x50, 0x01, 0x58, 0x01, 0x60, 0xff, 0x01,
	0x78, 0x01, 0x48, 0x01, 0x52, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a,
	0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0a, 0xca,
	0x43, 0x07, 0x22, 0x05, 0x20, 0x01, 0x28, 0xe8, 0x07, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x12, 0x29, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x72, 0x61, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xca, 0x43, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52,
	0x0b, 0x65, 0x78, 0x74, 0x72, 0x61, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x28, 0x0a, 0x09,
	0x73, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x42,
	0x0a, 0xca, 0x43, 0x07, 0x3a, 0x05, 0x15, 0x00, 0x00, 0x00, 0x00, 0x52, 0x09, 0x73, 0x6f, 0x6d,
	0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0c, 0x42, 0x0a, 0xca, 0x43, 0x07, 0x12, 0x05, 0x08, 0x80, 0x10, 0x50, 0x01, 0x48,
	0x02, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x07, 0x66, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x42, 0x07, 0xca, 0x43, 0x04, 0x0a,
	0x02, 0x78, 0x01, 0x52, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x34, 0x0a, 0x0d,
	0x73, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x45, 0x6c, 0x73, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x01, 0x42, 0x0e, 0xca, 0x43, 0x0b, 0x42, 0x09, 0x11, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x52, 0x0d, 0x73, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x45, 0x6c,
	0x73, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6b, 0x65, 0x79,
	0x22, 0xb5, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x12, 0x1d, 0x0a,
	0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1f, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x29, 0x0a, 0x06, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x07, 0xca, 0x43, 0x04, 0x4a,
	0x02, 0x10, 0x01, 0x52, 0x06, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f,
	0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x28, 0x0a, 0x08, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x33, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69,
	0x6e, 0x67, 0x42, 0x15, 0x5a, 0x13, 0x70, 0x6b, 0x67, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x3b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_example_proto_rawDescOnce sync.Once
	file_example_proto_rawDescData = file_example_proto_rawDesc
)

func file_example_proto_rawDescGZIP() []byte {
	file_example_proto_rawDescOnce.Do(func() {
		file_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_proto_rawDescData)
	})
	return file_example_proto_rawDescData
}

var file_example_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_example_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: Request
	(*Request2)(nil), // 1: Request2
	(*Request3)(nil), // 2: Request3
}
var file_example_proto_depIdxs = []int32{
	0, // 0: Request2.nested:type_name -> Request
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_example_proto_init() }
func file_example_proto_init() {
	if File_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request2); i {
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
		file_example_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request3); i {
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
	file_example_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_example_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_example_proto_goTypes,
		DependencyIndexes: file_example_proto_depIdxs,
		MessageInfos:      file_example_proto_msgTypes,
	}.Build()
	File_example_proto = out.File
	file_example_proto_rawDesc = nil
	file_example_proto_goTypes = nil
	file_example_proto_depIdxs = nil
}
