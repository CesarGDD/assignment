// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: rgpb.proto

package rgpb

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

type Vehicle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Make   string `protobuf:"bytes,2,opt,name=make,proto3" json:"make,omitempty"`
	Model  string `protobuf:"bytes,3,opt,name=model,proto3" json:"model,omitempty"`
	Colour string `protobuf:"bytes,4,opt,name=colour,proto3" json:"colour,omitempty"`
}

func (x *Vehicle) Reset() {
	*x = Vehicle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rgpb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vehicle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vehicle) ProtoMessage() {}

func (x *Vehicle) ProtoReflect() protoreflect.Message {
	mi := &file_rgpb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vehicle.ProtoReflect.Descriptor instead.
func (*Vehicle) Descriptor() ([]byte, []int) {
	return file_rgpb_proto_rawDescGZIP(), []int{0}
}

func (x *Vehicle) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Vehicle) GetMake() string {
	if x != nil {
		return x.Make
	}
	return ""
}

func (x *Vehicle) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *Vehicle) GetColour() string {
	if x != nil {
		return x.Colour
	}
	return ""
}

type Insurer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Code int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Insurer) Reset() {
	*x = Insurer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rgpb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Insurer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Insurer) ProtoMessage() {}

func (x *Insurer) ProtoReflect() protoreflect.Message {
	mi := &file_rgpb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Insurer.ProtoReflect.Descriptor instead.
func (*Insurer) Descriptor() ([]byte, []int) {
	return file_rgpb_proto_rawDescGZIP(), []int{1}
}

func (x *Insurer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Insurer) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type Registration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumberPlate string   `protobuf:"bytes,1,opt,name=number_plate,json=numberPlate,proto3" json:"number_plate,omitempty"`
	Vehicle     *Vehicle `protobuf:"bytes,2,opt,name=vehicle,proto3" json:"vehicle,omitempty"`
	Insurer     *Insurer `protobuf:"bytes,3,opt,name=insurer,proto3" json:"insurer,omitempty"`
}

func (x *Registration) Reset() {
	*x = Registration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rgpb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registration) ProtoMessage() {}

func (x *Registration) ProtoReflect() protoreflect.Message {
	mi := &file_rgpb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registration.ProtoReflect.Descriptor instead.
func (*Registration) Descriptor() ([]byte, []int) {
	return file_rgpb_proto_rawDescGZIP(), []int{2}
}

func (x *Registration) GetNumberPlate() string {
	if x != nil {
		return x.NumberPlate
	}
	return ""
}

func (x *Registration) GetVehicle() *Vehicle {
	if x != nil {
		return x.Vehicle
	}
	return nil
}

func (x *Registration) GetInsurer() *Insurer {
	if x != nil {
		return x.Insurer
	}
	return nil
}

type CreateRegistrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Registration *Registration `protobuf:"bytes,1,opt,name=registration,proto3" json:"registration,omitempty"`
}

func (x *CreateRegistrationRequest) Reset() {
	*x = CreateRegistrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rgpb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRegistrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRegistrationRequest) ProtoMessage() {}

func (x *CreateRegistrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rgpb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRegistrationRequest.ProtoReflect.Descriptor instead.
func (*CreateRegistrationRequest) Descriptor() ([]byte, []int) {
	return file_rgpb_proto_rawDescGZIP(), []int{3}
}

func (x *CreateRegistrationRequest) GetRegistration() *Registration {
	if x != nil {
		return x.Registration
	}
	return nil
}

type CreateRegistrationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Registration *Registration `protobuf:"bytes,1,opt,name=registration,proto3" json:"registration,omitempty"`
}

func (x *CreateRegistrationResponse) Reset() {
	*x = CreateRegistrationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rgpb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRegistrationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRegistrationResponse) ProtoMessage() {}

func (x *CreateRegistrationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rgpb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRegistrationResponse.ProtoReflect.Descriptor instead.
func (*CreateRegistrationResponse) Descriptor() ([]byte, []int) {
	return file_rgpb_proto_rawDescGZIP(), []int{4}
}

func (x *CreateRegistrationResponse) GetRegistration() *Registration {
	if x != nil {
		return x.Registration
	}
	return nil
}

type GetRegistrarionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumberPlate string `protobuf:"bytes,1,opt,name=number_plate,json=numberPlate,proto3" json:"number_plate,omitempty"`
}

func (x *GetRegistrarionRequest) Reset() {
	*x = GetRegistrarionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rgpb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRegistrarionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRegistrarionRequest) ProtoMessage() {}

func (x *GetRegistrarionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rgpb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRegistrarionRequest.ProtoReflect.Descriptor instead.
func (*GetRegistrarionRequest) Descriptor() ([]byte, []int) {
	return file_rgpb_proto_rawDescGZIP(), []int{5}
}

func (x *GetRegistrarionRequest) GetNumberPlate() string {
	if x != nil {
		return x.NumberPlate
	}
	return ""
}

type GetRegistrationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Registration *Registration `protobuf:"bytes,1,opt,name=registration,proto3" json:"registration,omitempty"`
}

func (x *GetRegistrationResponse) Reset() {
	*x = GetRegistrationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rgpb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRegistrationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRegistrationResponse) ProtoMessage() {}

func (x *GetRegistrationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rgpb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRegistrationResponse.ProtoReflect.Descriptor instead.
func (*GetRegistrationResponse) Descriptor() ([]byte, []int) {
	return file_rgpb_proto_rawDescGZIP(), []int{6}
}

func (x *GetRegistrationResponse) GetRegistration() *Registration {
	if x != nil {
		return x.Registration
	}
	return nil
}

type ListRegistrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRegistrationRequest) Reset() {
	*x = ListRegistrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rgpb_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRegistrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRegistrationRequest) ProtoMessage() {}

func (x *ListRegistrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rgpb_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRegistrationRequest.ProtoReflect.Descriptor instead.
func (*ListRegistrationRequest) Descriptor() ([]byte, []int) {
	return file_rgpb_proto_rawDescGZIP(), []int{7}
}

type ListRegistrationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Registration []*Registration `protobuf:"bytes,1,rep,name=registration,proto3" json:"registration,omitempty"`
}

func (x *ListRegistrationResponse) Reset() {
	*x = ListRegistrationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rgpb_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRegistrationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRegistrationResponse) ProtoMessage() {}

func (x *ListRegistrationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rgpb_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRegistrationResponse.ProtoReflect.Descriptor instead.
func (*ListRegistrationResponse) Descriptor() ([]byte, []int) {
	return file_rgpb_proto_rawDescGZIP(), []int{8}
}

func (x *ListRegistrationResponse) GetRegistration() []*Registration {
	if x != nil {
		return x.Registration
	}
	return nil
}

var File_rgpb_proto protoreflect.FileDescriptor

var file_rgpb_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x67, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72, 0x67,
	0x70, 0x62, 0x22, 0x5f, 0x0a, 0x07, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x6b, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6d, 0x61, 0x6b, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x6f, 0x6c, 0x6f, 0x75, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6c,
	0x6f, 0x75, 0x72, 0x22, 0x31, 0x0a, 0x07, 0x49, 0x6e, 0x73, 0x75, 0x72, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x83, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x76, 0x65,
	0x68, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x67,
	0x70, 0x62, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x07, 0x76, 0x65, 0x68, 0x69,
	0x63, 0x6c, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x69, 0x6e, 0x73, 0x75, 0x72, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x67, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x73, 0x75,
	0x72, 0x65, 0x72, 0x52, 0x07, 0x69, 0x6e, 0x73, 0x75, 0x72, 0x65, 0x72, 0x22, 0x53, 0x0a, 0x19,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0c, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x72, 0x67, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x54, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x36, 0x0a, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72, 0x67, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3b, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x50,
	0x6c, 0x61, 0x74, 0x65, 0x22, 0x51, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x36, 0x0a, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72, 0x67, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x19, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x52, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36,
	0x0a, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72, 0x67, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x91, 0x02, 0x0a, 0x13, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x72, 0x67, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x72, 0x67, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x72, 0x67, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x72, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x72, 0x67, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x72, 0x67,
	0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x72, 0x67, 0x70,
	0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2f, 0x72,
	0x67, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rgpb_proto_rawDescOnce sync.Once
	file_rgpb_proto_rawDescData = file_rgpb_proto_rawDesc
)

func file_rgpb_proto_rawDescGZIP() []byte {
	file_rgpb_proto_rawDescOnce.Do(func() {
		file_rgpb_proto_rawDescData = protoimpl.X.CompressGZIP(file_rgpb_proto_rawDescData)
	})
	return file_rgpb_proto_rawDescData
}

var file_rgpb_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_rgpb_proto_goTypes = []interface{}{
	(*Vehicle)(nil),                    // 0: rgpb.Vehicle
	(*Insurer)(nil),                    // 1: rgpb.Insurer
	(*Registration)(nil),               // 2: rgpb.Registration
	(*CreateRegistrationRequest)(nil),  // 3: rgpb.CreateRegistrationRequest
	(*CreateRegistrationResponse)(nil), // 4: rgpb.CreateRegistrationResponse
	(*GetRegistrarionRequest)(nil),     // 5: rgpb.GetRegistrarionRequest
	(*GetRegistrationResponse)(nil),    // 6: rgpb.GetRegistrationResponse
	(*ListRegistrationRequest)(nil),    // 7: rgpb.ListRegistrationRequest
	(*ListRegistrationResponse)(nil),   // 8: rgpb.ListRegistrationResponse
}
var file_rgpb_proto_depIdxs = []int32{
	0, // 0: rgpb.Registration.vehicle:type_name -> rgpb.Vehicle
	1, // 1: rgpb.Registration.insurer:type_name -> rgpb.Insurer
	2, // 2: rgpb.CreateRegistrationRequest.registration:type_name -> rgpb.Registration
	2, // 3: rgpb.CreateRegistrationResponse.registration:type_name -> rgpb.Registration
	2, // 4: rgpb.GetRegistrationResponse.registration:type_name -> rgpb.Registration
	2, // 5: rgpb.ListRegistrationResponse.registration:type_name -> rgpb.Registration
	3, // 6: rgpb.RegistrationService.CreateRegistration:input_type -> rgpb.CreateRegistrationRequest
	5, // 7: rgpb.RegistrationService.GetRegistrarion:input_type -> rgpb.GetRegistrarionRequest
	7, // 8: rgpb.RegistrationService.ListRegistration:input_type -> rgpb.ListRegistrationRequest
	4, // 9: rgpb.RegistrationService.CreateRegistration:output_type -> rgpb.CreateRegistrationResponse
	6, // 10: rgpb.RegistrationService.GetRegistrarion:output_type -> rgpb.GetRegistrationResponse
	8, // 11: rgpb.RegistrationService.ListRegistration:output_type -> rgpb.ListRegistrationResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_rgpb_proto_init() }
func file_rgpb_proto_init() {
	if File_rgpb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rgpb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vehicle); i {
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
		file_rgpb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Insurer); i {
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
		file_rgpb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registration); i {
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
		file_rgpb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRegistrationRequest); i {
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
		file_rgpb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRegistrationResponse); i {
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
		file_rgpb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRegistrarionRequest); i {
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
		file_rgpb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRegistrationResponse); i {
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
		file_rgpb_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRegistrationRequest); i {
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
		file_rgpb_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRegistrationResponse); i {
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
			RawDescriptor: file_rgpb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rgpb_proto_goTypes,
		DependencyIndexes: file_rgpb_proto_depIdxs,
		MessageInfos:      file_rgpb_proto_msgTypes,
	}.Build()
	File_rgpb_proto = out.File
	file_rgpb_proto_rawDesc = nil
	file_rgpb_proto_goTypes = nil
	file_rgpb_proto_depIdxs = nil
}