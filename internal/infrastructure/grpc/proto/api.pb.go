// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: api.proto

package __

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StatisticsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ram              *Ram              `protobuf:"bytes,1,opt,name=ram,proto3" json:"ram,omitempty"`
	Storage          *Storage          `protobuf:"bytes,2,opt,name=storage,proto3" json:"storage,omitempty"`
	Loadavg          *LoadAvg          `protobuf:"bytes,3,opt,name=loadavg,proto3" json:"loadavg,omitempty"`
	Uptime           *Uptime           `protobuf:"bytes,4,opt,name=uptime,proto3" json:"uptime,omitempty"`
	DockerContainers map[string]string `protobuf:"bytes,5,rep,name=DockerContainers,proto3" json:"DockerContainers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	UpdatedAt        uint64            `protobuf:"varint,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *StatisticsResponse) Reset() {
	*x = StatisticsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatisticsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatisticsResponse) ProtoMessage() {}

func (x *StatisticsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatisticsResponse.ProtoReflect.Descriptor instead.
func (*StatisticsResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *StatisticsResponse) GetRam() *Ram {
	if x != nil {
		return x.Ram
	}
	return nil
}

func (x *StatisticsResponse) GetStorage() *Storage {
	if x != nil {
		return x.Storage
	}
	return nil
}

func (x *StatisticsResponse) GetLoadavg() *LoadAvg {
	if x != nil {
		return x.Loadavg
	}
	return nil
}

func (x *StatisticsResponse) GetUptime() *Uptime {
	if x != nil {
		return x.Uptime
	}
	return nil
}

func (x *StatisticsResponse) GetDockerContainers() map[string]string {
	if x != nil {
		return x.DockerContainers
	}
	return nil
}

func (x *StatisticsResponse) GetUpdatedAt() uint64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type Ram struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total     uint64 `protobuf:"varint,1,opt,name=Total,proto3" json:"Total,omitempty"`
	Available uint64 `protobuf:"varint,2,opt,name=Available,proto3" json:"Available,omitempty"`
}

func (x *Ram) Reset() {
	*x = Ram{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ram) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ram) ProtoMessage() {}

func (x *Ram) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ram.ProtoReflect.Descriptor instead.
func (*Ram) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *Ram) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Ram) GetAvailable() uint64 {
	if x != nil {
		return x.Available
	}
	return 0
}

type Storage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total     uint64 `protobuf:"varint,1,opt,name=Total,proto3" json:"Total,omitempty"`
	Available uint64 `protobuf:"varint,2,opt,name=Available,proto3" json:"Available,omitempty"`
	Free      uint64 `protobuf:"varint,3,opt,name=Free,proto3" json:"Free,omitempty"`
}

func (x *Storage) Reset() {
	*x = Storage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Storage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Storage) ProtoMessage() {}

func (x *Storage) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Storage.ProtoReflect.Descriptor instead.
func (*Storage) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *Storage) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Storage) GetAvailable() uint64 {
	if x != nil {
		return x.Available
	}
	return 0
}

func (x *Storage) GetFree() uint64 {
	if x != nil {
		return x.Free
	}
	return 0
}

type LoadAvg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Load1  float32 `protobuf:"fixed32,1,opt,name=Load1,proto3" json:"Load1,omitempty"`
	Load5  float32 `protobuf:"fixed32,2,opt,name=Load5,proto3" json:"Load5,omitempty"`
	Load15 float32 `protobuf:"fixed32,3,opt,name=Load15,proto3" json:"Load15,omitempty"`
}

func (x *LoadAvg) Reset() {
	*x = LoadAvg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadAvg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadAvg) ProtoMessage() {}

func (x *LoadAvg) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadAvg.ProtoReflect.Descriptor instead.
func (*LoadAvg) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *LoadAvg) GetLoad1() float32 {
	if x != nil {
		return x.Load1
	}
	return 0
}

func (x *LoadAvg) GetLoad5() float32 {
	if x != nil {
		return x.Load5
	}
	return 0
}

func (x *LoadAvg) GetLoad15() float32 {
	if x != nil {
		return x.Load15
	}
	return 0
}

type Uptime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dur float32 `protobuf:"fixed32,1,opt,name=dur,proto3" json:"dur,omitempty"`
}

func (x *Uptime) Reset() {
	*x = Uptime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Uptime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Uptime) ProtoMessage() {}

func (x *Uptime) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Uptime.ProtoReflect.Descriptor instead.
func (*Uptime) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *Uptime) GetDur() float32 {
	if x != nil {
		return x.Dur
	}
	return 0
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70, 0x69,
	0x5f, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf2,
	0x02, 0x0a, 0x12, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x03, 0x72, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x70, 0x62, 0x2e, 0x52, 0x61, 0x6d, 0x52,
	0x03, 0x72, 0x61, 0x6d, 0x12, 0x29, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x70, 0x62, 0x2e, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12,
	0x29, 0x0a, 0x07, 0x6c, 0x6f, 0x61, 0x64, 0x61, 0x76, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x41, 0x76,
	0x67, 0x52, 0x07, 0x6c, 0x6f, 0x61, 0x64, 0x61, 0x76, 0x67, 0x12, 0x26, 0x0a, 0x06, 0x75, 0x70,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69,
	0x5f, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x06, 0x75, 0x70, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x5c, 0x0a, 0x10, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x61,
	0x70, 0x69, 0x5f, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x10,
	0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x1a, 0x43,
	0x0a, 0x15, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x39, 0x0a, 0x03, 0x52, 0x61, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x1c, 0x0a, 0x09, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x51,
	0x0a, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x1c, 0x0a, 0x09, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x09, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x46, 0x72, 0x65, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x46, 0x72, 0x65,
	0x65, 0x22, 0x4d, 0x0a, 0x07, 0x4c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x67, 0x12, 0x14, 0x0a, 0x05,
	0x4c, 0x6f, 0x61, 0x64, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x4c, 0x6f, 0x61,
	0x64, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x6f, 0x61, 0x64, 0x35, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x4c, 0x6f, 0x61, 0x64, 0x35, 0x12, 0x16, 0x0a, 0x06, 0x4c, 0x6f, 0x61, 0x64,
	0x31, 0x35, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x4c, 0x6f, 0x61, 0x64, 0x31, 0x35,
	0x22, 0x1a, 0x0a, 0x06, 0x55, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x75,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x64, 0x75, 0x72, 0x32, 0x6e, 0x0a, 0x12,
	0x48, 0x6f, 0x73, 0x74, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x58, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x61, 0x70,
	0x69, 0x5f, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12,
	0x0b, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x42, 0x04, 0x5a, 0x02,
	0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_proto_goTypes = []interface{}{
	(*StatisticsResponse)(nil), // 0: api_pb.StatisticsResponse
	(*Ram)(nil),                // 1: api_pb.Ram
	(*Storage)(nil),            // 2: api_pb.Storage
	(*LoadAvg)(nil),            // 3: api_pb.LoadAvg
	(*Uptime)(nil),             // 4: api_pb.Uptime
	nil,                        // 5: api_pb.StatisticsResponse.DockerContainersEntry
	(*emptypb.Empty)(nil),      // 6: google.protobuf.Empty
}
var file_api_proto_depIdxs = []int32{
	1, // 0: api_pb.StatisticsResponse.ram:type_name -> api_pb.Ram
	2, // 1: api_pb.StatisticsResponse.storage:type_name -> api_pb.Storage
	3, // 2: api_pb.StatisticsResponse.loadavg:type_name -> api_pb.LoadAvg
	4, // 3: api_pb.StatisticsResponse.uptime:type_name -> api_pb.Uptime
	5, // 4: api_pb.StatisticsResponse.DockerContainers:type_name -> api_pb.StatisticsResponse.DockerContainersEntry
	6, // 5: api_pb.HostMonitorService.GetStatistics:input_type -> google.protobuf.Empty
	0, // 6: api_pb.HostMonitorService.GetStatistics:output_type -> api_pb.StatisticsResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatisticsResponse); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ram); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Storage); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadAvg); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Uptime); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
