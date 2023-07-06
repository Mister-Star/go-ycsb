// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: kvdb_server.proto

package taas_proto

import (
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

type KvDBData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpType OpType `protobuf:"varint,1,opt,name=op_type,json=opType,proto3,enum=taas_proto.OpType" json:"op_type,omitempty"`
	Key    string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value  string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Csn    uint64 `protobuf:"varint,4,opt,name=csn,proto3" json:"csn,omitempty"`
}

func (x *KvDBData) Reset() {
	*x = KvDBData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kvdb_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KvDBData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KvDBData) ProtoMessage() {}

func (x *KvDBData) ProtoReflect() protoreflect.Message {
	mi := &file_kvdb_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KvDBData.ProtoReflect.Descriptor instead.
func (*KvDBData) Descriptor() ([]byte, []int) {
	return file_kvdb_server_proto_rawDescGZIP(), []int{0}
}

func (x *KvDBData) GetOpType() OpType {
	if x != nil {
		return x.OpType
	}
	return OpType_Read
}

func (x *KvDBData) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KvDBData) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *KvDBData) GetCsn() uint64 {
	if x != nil {
		return x.Csn
	}
	return 0
}

type KvDBRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*KvDBData `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *KvDBRequest) Reset() {
	*x = KvDBRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kvdb_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KvDBRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KvDBRequest) ProtoMessage() {}

func (x *KvDBRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kvdb_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KvDBRequest.ProtoReflect.Descriptor instead.
func (*KvDBRequest) Descriptor() ([]byte, []int) {
	return file_kvdb_server_proto_rawDescGZIP(), []int{1}
}

func (x *KvDBRequest) GetData() []*KvDBData {
	if x != nil {
		return x.Data
	}
	return nil
}

type KvDBResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool        `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Data   []*KvDBData `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *KvDBResponse) Reset() {
	*x = KvDBResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kvdb_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KvDBResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KvDBResponse) ProtoMessage() {}

func (x *KvDBResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kvdb_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KvDBResponse.ProtoReflect.Descriptor instead.
func (*KvDBResponse) Descriptor() ([]byte, []int) {
	return file_kvdb_server_proto_rawDescGZIP(), []int{2}
}

func (x *KvDBResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *KvDBResponse) GetData() []*KvDBData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_kvdb_server_proto protoreflect.FileDescriptor

var file_kvdb_server_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6b, 0x76, 0x64, 0x62, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x61, 0x61, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x71, 0x0a, 0x08, 0x4b, 0x76, 0x44, 0x42, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2b,
	0x0a, 0x07, 0x6f, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x12, 0x2e, 0x74, 0x61, 0x61, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x73, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x03, 0x63, 0x73, 0x6e, 0x22, 0x37, 0x0a, 0x0b, 0x4b, 0x76, 0x44, 0x42, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x61, 0x61, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4b, 0x76, 0x44, 0x42, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x50,
	0x0a, 0x0c, 0x4b, 0x76, 0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x61, 0x61, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4b, 0x76, 0x44, 0x42, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x32, 0x4a, 0x0a, 0x0e, 0x4b, 0x76, 0x44, 0x42, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x38, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x74, 0x61, 0x61, 0x73,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x76, 0x44, 0x42, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74, 0x61, 0x61, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4b, 0x76, 0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x4a, 0x0a, 0x0e,
	0x4b, 0x76, 0x44, 0x42, 0x50, 0x75, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38,
	0x0a, 0x03, 0x50, 0x75, 0x74, 0x12, 0x17, 0x2e, 0x74, 0x61, 0x61, 0x73, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4b, 0x76, 0x44, 0x42, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x74, 0x61, 0x61, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x76, 0x44, 0x42,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x74, 0x61,
	0x61, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kvdb_server_proto_rawDescOnce sync.Once
	file_kvdb_server_proto_rawDescData = file_kvdb_server_proto_rawDesc
)

func file_kvdb_server_proto_rawDescGZIP() []byte {
	file_kvdb_server_proto_rawDescOnce.Do(func() {
		file_kvdb_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_kvdb_server_proto_rawDescData)
	})
	return file_kvdb_server_proto_rawDescData
}

var file_kvdb_server_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_kvdb_server_proto_goTypes = []interface{}{
	(*KvDBData)(nil),     // 0: taas_proto.KvDBData
	(*KvDBRequest)(nil),  // 1: taas_proto.KvDBRequest
	(*KvDBResponse)(nil), // 2: taas_proto.KvDBResponse
	(OpType)(0),          // 3: taas_proto.OpType
}
var file_kvdb_server_proto_depIdxs = []int32{
	3, // 0: taas_proto.KvDBData.op_type:type_name -> taas_proto.OpType
	0, // 1: taas_proto.KvDBRequest.data:type_name -> taas_proto.KvDBData
	0, // 2: taas_proto.KvDBResponse.data:type_name -> taas_proto.KvDBData
	1, // 3: taas_proto.KvDBGetService.Get:input_type -> taas_proto.KvDBRequest
	1, // 4: taas_proto.KvDBPutService.Put:input_type -> taas_proto.KvDBRequest
	2, // 5: taas_proto.KvDBGetService.Get:output_type -> taas_proto.KvDBResponse
	2, // 6: taas_proto.KvDBPutService.Put:output_type -> taas_proto.KvDBResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_kvdb_server_proto_init() }
func file_kvdb_server_proto_init() {
	if File_kvdb_server_proto != nil {
		return
	}
	file_transaction_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_kvdb_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KvDBData); i {
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
		file_kvdb_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KvDBRequest); i {
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
		file_kvdb_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KvDBResponse); i {
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
			RawDescriptor: file_kvdb_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_kvdb_server_proto_goTypes,
		DependencyIndexes: file_kvdb_server_proto_depIdxs,
		MessageInfos:      file_kvdb_server_proto_msgTypes,
	}.Build()
	File_kvdb_server_proto = out.File
	file_kvdb_server_proto_rawDesc = nil
	file_kvdb_server_proto_goTypes = nil
	file_kvdb_server_proto_depIdxs = nil
}
