// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: protos/clients.proto

package proto

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

type ListClientsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListClientsRequest) Reset() {
	*x = ListClientsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_clients_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListClientsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClientsRequest) ProtoMessage() {}

func (x *ListClientsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_clients_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClientsRequest.ProtoReflect.Descriptor instead.
func (*ListClientsRequest) Descriptor() ([]byte, []int) {
	return file_protos_clients_proto_rawDescGZIP(), []int{0}
}

type ListClientsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clients []*ClientInfo `protobuf:"bytes,1,rep,name=clients,proto3" json:"clients,omitempty"`
}

func (x *ListClientsReply) Reset() {
	*x = ListClientsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_clients_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListClientsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClientsReply) ProtoMessage() {}

func (x *ListClientsReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_clients_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClientsReply.ProtoReflect.Descriptor instead.
func (*ListClientsReply) Descriptor() ([]byte, []int) {
	return file_protos_clients_proto_rawDescGZIP(), []int{1}
}

func (x *ListClientsReply) GetClients() []*ClientInfo {
	if x != nil {
		return x.Clients
	}
	return nil
}

type ClientInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string           `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	User     string           `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Dbname   string           `protobuf:"bytes,3,opt,name=dbname,proto3" json:"dbname,omitempty"`
	Shards   []*UsedShardInfo `protobuf:"bytes,4,rep,name=shards,proto3" json:"shards,omitempty"`
}

func (x *ClientInfo) Reset() {
	*x = ClientInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_clients_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientInfo) ProtoMessage() {}

func (x *ClientInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_clients_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientInfo.ProtoReflect.Descriptor instead.
func (*ClientInfo) Descriptor() ([]byte, []int) {
	return file_protos_clients_proto_rawDescGZIP(), []int{2}
}

func (x *ClientInfo) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *ClientInfo) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *ClientInfo) GetDbname() string {
	if x != nil {
		return x.Dbname
	}
	return ""
}

func (x *ClientInfo) GetShards() []*UsedShardInfo {
	if x != nil {
		return x.Shards
	}
	return nil
}

type UsedShardInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instance *DBInstaceInfo `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
}

func (x *UsedShardInfo) Reset() {
	*x = UsedShardInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_clients_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsedShardInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsedShardInfo) ProtoMessage() {}

func (x *UsedShardInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_clients_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsedShardInfo.ProtoReflect.Descriptor instead.
func (*UsedShardInfo) Descriptor() ([]byte, []int) {
	return file_protos_clients_proto_rawDescGZIP(), []int{3}
}

func (x *UsedShardInfo) GetInstance() *DBInstaceInfo {
	if x != nil {
		return x.Instance
	}
	return nil
}

type DBInstaceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
}

func (x *DBInstaceInfo) Reset() {
	*x = DBInstaceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_clients_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBInstaceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBInstaceInfo) ProtoMessage() {}

func (x *DBInstaceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_clients_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBInstaceInfo.ProtoReflect.Descriptor instead.
func (*DBInstaceInfo) Descriptor() ([]byte, []int) {
	return file_protos_clients_proto_rawDescGZIP(), []int{4}
}

func (x *DBInstaceInfo) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

var File_protos_clients_proto protoreflect.FileDescriptor

var file_protos_clients_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x73, 0x70, 0x71, 0x72, 0x22, 0x14, 0x0a, 0x12,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x3e, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2a, 0x0a, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x70, 0x71, 0x72, 0x2e, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x82, 0x01, 0x0a, 0x0a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x62, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x64, 0x62, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x68,
	0x61, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x70, 0x71,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x64, 0x53, 0x68, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x06, 0x73, 0x68, 0x61, 0x72, 0x64, 0x73, 0x22, 0x40, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x64, 0x53,
	0x68, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2f, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x70, 0x71,
	0x72, 0x2e, 0x44, 0x42, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x2b, 0x0a, 0x0d, 0x44, 0x42, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x56, 0x0a, 0x11, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0b, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x18, 0x2e, 0x73, 0x70, 0x71,
	0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x70, 0x71, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x0c,
	0x5a, 0x0a, 0x73, 0x70, 0x71, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_clients_proto_rawDescOnce sync.Once
	file_protos_clients_proto_rawDescData = file_protos_clients_proto_rawDesc
)

func file_protos_clients_proto_rawDescGZIP() []byte {
	file_protos_clients_proto_rawDescOnce.Do(func() {
		file_protos_clients_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_clients_proto_rawDescData)
	})
	return file_protos_clients_proto_rawDescData
}

var file_protos_clients_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_clients_proto_goTypes = []interface{}{
	(*ListClientsRequest)(nil), // 0: spqr.ListClientsRequest
	(*ListClientsReply)(nil),   // 1: spqr.ListClientsReply
	(*ClientInfo)(nil),         // 2: spqr.ClientInfo
	(*UsedShardInfo)(nil),      // 3: spqr.UsedShardInfo
	(*DBInstaceInfo)(nil),      // 4: spqr.DBInstaceInfo
}
var file_protos_clients_proto_depIdxs = []int32{
	2, // 0: spqr.ListClientsReply.clients:type_name -> spqr.ClientInfo
	3, // 1: spqr.ClientInfo.shards:type_name -> spqr.UsedShardInfo
	4, // 2: spqr.UsedShardInfo.instance:type_name -> spqr.DBInstaceInfo
	0, // 3: spqr.ClientInfoService.ListClients:input_type -> spqr.ListClientsRequest
	1, // 4: spqr.ClientInfoService.ListClients:output_type -> spqr.ListClientsReply
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protos_clients_proto_init() }
func file_protos_clients_proto_init() {
	if File_protos_clients_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_clients_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListClientsRequest); i {
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
		file_protos_clients_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListClientsReply); i {
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
		file_protos_clients_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientInfo); i {
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
		file_protos_clients_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsedShardInfo); i {
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
		file_protos_clients_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBInstaceInfo); i {
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
			RawDescriptor: file_protos_clients_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_clients_proto_goTypes,
		DependencyIndexes: file_protos_clients_proto_depIdxs,
		MessageInfos:      file_protos_clients_proto_msgTypes,
	}.Build()
	File_protos_clients_proto = out.File
	file_protos_clients_proto_rawDesc = nil
	file_protos_clients_proto_goTypes = nil
	file_protos_clients_proto_depIdxs = nil
}
