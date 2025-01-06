// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: grpc/engine.proto

package grpc

import (
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

type Objects struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Objects       []*Object              `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Objects) Reset() {
	*x = Objects{}
	mi := &file_grpc_engine_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Objects) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Objects) ProtoMessage() {}

func (x *Objects) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_engine_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Objects.ProtoReflect.Descriptor instead.
func (*Objects) Descriptor() ([]byte, []int) {
	return file_grpc_engine_proto_rawDescGZIP(), []int{0}
}

func (x *Objects) GetObjects() []*Object {
	if x != nil {
		return x.Objects
	}
	return nil
}

type Object struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id            uint32                 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Location      *Location              `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Object) Reset() {
	*x = Object{}
	mi := &file_grpc_engine_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_engine_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_grpc_engine_proto_rawDescGZIP(), []int{1}
}

func (x *Object) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Object) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Object) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

type Location struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Position      *Vector3               `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	Rotation      *Vector4               `protobuf:"bytes,2,opt,name=rotation,proto3" json:"rotation,omitempty"`
	Scale         *Vector3               `protobuf:"bytes,3,opt,name=scale,proto3" json:"scale,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Location) Reset() {
	*x = Location{}
	mi := &file_grpc_engine_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_engine_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_grpc_engine_proto_rawDescGZIP(), []int{2}
}

func (x *Location) GetPosition() *Vector3 {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *Location) GetRotation() *Vector4 {
	if x != nil {
		return x.Rotation
	}
	return nil
}

func (x *Location) GetScale() *Vector3 {
	if x != nil {
		return x.Scale
	}
	return nil
}

type Vector4 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	X             float32                `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y             float32                `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	Z             float32                `protobuf:"fixed32,3,opt,name=z,proto3" json:"z,omitempty"`
	W             float32                `protobuf:"fixed32,4,opt,name=w,proto3" json:"w,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Vector4) Reset() {
	*x = Vector4{}
	mi := &file_grpc_engine_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Vector4) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vector4) ProtoMessage() {}

func (x *Vector4) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_engine_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vector4.ProtoReflect.Descriptor instead.
func (*Vector4) Descriptor() ([]byte, []int) {
	return file_grpc_engine_proto_rawDescGZIP(), []int{3}
}

func (x *Vector4) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Vector4) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Vector4) GetZ() float32 {
	if x != nil {
		return x.Z
	}
	return 0
}

func (x *Vector4) GetW() float32 {
	if x != nil {
		return x.W
	}
	return 0
}

type Vector3 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	X             float32                `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y             float32                `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	Z             float32                `protobuf:"fixed32,3,opt,name=z,proto3" json:"z,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Vector3) Reset() {
	*x = Vector3{}
	mi := &file_grpc_engine_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Vector3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vector3) ProtoMessage() {}

func (x *Vector3) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_engine_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vector3.ProtoReflect.Descriptor instead.
func (*Vector3) Descriptor() ([]byte, []int) {
	return file_grpc_engine_proto_rawDescGZIP(), []int{4}
}

func (x *Vector3) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Vector3) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Vector3) GetZ() float32 {
	if x != nil {
		return x.Z
	}
	return 0
}

var File_grpc_engine_proto protoreflect.FileDescriptor

var file_grpc_engine_proto_rawDesc = []byte{
	0x0a, 0x11, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x07, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x12, 0x26, 0x0a, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0x58, 0x0a, 0x06, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x85, 0x01, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x29, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x33, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x08, 0x72,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x34, 0x52, 0x08, 0x72, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x05, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x33, 0x52, 0x05, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x22, 0x41, 0x0a, 0x07, 0x56,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x34, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x7a,
	0x12, 0x0c, 0x0a, 0x01, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x77, 0x22, 0x33,
	0x0a, 0x07, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x01, 0x7a, 0x32, 0xfb, 0x02, 0x0a, 0x06, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x33,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x12, 0x31, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x0c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x34, 0x0a, 0x0c, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x32, 0x0a, 0x0a,
	0x4d, 0x6f, 0x76, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0c, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x34, 0x0a, 0x0c, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x0c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x33, 0x0a, 0x0b, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x34, 0x0a, 0x0c, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0c, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x42, 0x10, 0x5a, 0x0e, 0x33, 0x64, 0x2d, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_engine_proto_rawDescOnce sync.Once
	file_grpc_engine_proto_rawDescData = file_grpc_engine_proto_rawDesc
)

func file_grpc_engine_proto_rawDescGZIP() []byte {
	file_grpc_engine_proto_rawDescOnce.Do(func() {
		file_grpc_engine_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_engine_proto_rawDescData)
	})
	return file_grpc_engine_proto_rawDescData
}

var file_grpc_engine_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_grpc_engine_proto_goTypes = []any{
	(*Objects)(nil),       // 0: grpc.Objects
	(*Object)(nil),        // 1: grpc.Object
	(*Location)(nil),      // 2: grpc.Location
	(*Vector4)(nil),       // 3: grpc.Vector4
	(*Vector3)(nil),       // 4: grpc.Vector3
	(*emptypb.Empty)(nil), // 5: google.protobuf.Empty
}
var file_grpc_engine_proto_depIdxs = []int32{
	1,  // 0: grpc.Objects.objects:type_name -> grpc.Object
	2,  // 1: grpc.Object.location:type_name -> grpc.Location
	4,  // 2: grpc.Location.position:type_name -> grpc.Vector3
	3,  // 3: grpc.Location.rotation:type_name -> grpc.Vector4
	4,  // 4: grpc.Location.scale:type_name -> grpc.Vector3
	5,  // 5: grpc.Engine.GetObjects:input_type -> google.protobuf.Empty
	1,  // 6: grpc.Engine.AddObject:input_type -> grpc.Object
	1,  // 7: grpc.Engine.RemoveObject:input_type -> grpc.Object
	1,  // 8: grpc.Engine.MoveObject:input_type -> grpc.Object
	1,  // 9: grpc.Engine.RotateObject:input_type -> grpc.Object
	1,  // 10: grpc.Engine.ScaleObject:input_type -> grpc.Object
	1,  // 11: grpc.Engine.UpdateObject:input_type -> grpc.Object
	0,  // 12: grpc.Engine.GetObjects:output_type -> grpc.Objects
	5,  // 13: grpc.Engine.AddObject:output_type -> google.protobuf.Empty
	5,  // 14: grpc.Engine.RemoveObject:output_type -> google.protobuf.Empty
	5,  // 15: grpc.Engine.MoveObject:output_type -> google.protobuf.Empty
	5,  // 16: grpc.Engine.RotateObject:output_type -> google.protobuf.Empty
	5,  // 17: grpc.Engine.ScaleObject:output_type -> google.protobuf.Empty
	5,  // 18: grpc.Engine.UpdateObject:output_type -> google.protobuf.Empty
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_grpc_engine_proto_init() }
func file_grpc_engine_proto_init() {
	if File_grpc_engine_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_engine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_engine_proto_goTypes,
		DependencyIndexes: file_grpc_engine_proto_depIdxs,
		MessageInfos:      file_grpc_engine_proto_msgTypes,
	}.Build()
	File_grpc_engine_proto = out.File
	file_grpc_engine_proto_rawDesc = nil
	file_grpc_engine_proto_goTypes = nil
	file_grpc_engine_proto_depIdxs = nil
}
