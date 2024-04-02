// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: doctors_service.proto

package doctors_protobuf

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

type Doctor_Gender int32

const (
	Doctor_UNSPECIFIED Doctor_Gender = 0
	Doctor_MALE        Doctor_Gender = 1
	Doctor_FEMALE      Doctor_Gender = 2
)

// Enum value maps for Doctor_Gender.
var (
	Doctor_Gender_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "MALE",
		2: "FEMALE",
	}
	Doctor_Gender_value = map[string]int32{
		"UNSPECIFIED": 0,
		"MALE":        1,
		"FEMALE":      2,
	}
)

func (x Doctor_Gender) Enum() *Doctor_Gender {
	p := new(Doctor_Gender)
	*p = x
	return p
}

func (x Doctor_Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Doctor_Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_doctors_service_proto_enumTypes[0].Descriptor()
}

func (Doctor_Gender) Type() protoreflect.EnumType {
	return &file_doctors_service_proto_enumTypes[0]
}

func (x Doctor_Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Doctor_Gender.Descriptor instead.
func (Doctor_Gender) EnumDescriptor() ([]byte, []int) {
	return file_doctors_service_proto_rawDescGZIP(), []int{5, 0}
}

type DoctorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Id    int32  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DoctorRequest) Reset() {
	*x = DoctorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doctors_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoctorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoctorRequest) ProtoMessage() {}

func (x *DoctorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doctors_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoctorRequest.ProtoReflect.Descriptor instead.
func (*DoctorRequest) Descriptor() ([]byte, []int) {
	return file_doctors_service_proto_rawDescGZIP(), []int{0}
}

func (x *DoctorRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *DoctorRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DoctorsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Limit  int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *DoctorsRequest) Reset() {
	*x = DoctorsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doctors_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoctorsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoctorsRequest) ProtoMessage() {}

func (x *DoctorsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doctors_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoctorsRequest.ProtoReflect.Descriptor instead.
func (*DoctorsRequest) Descriptor() ([]byte, []int) {
	return file_doctors_service_proto_rawDescGZIP(), []int{1}
}

func (x *DoctorsRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *DoctorsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *DoctorsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type CreateDoctorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token        string        `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Name         string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Gender       Doctor_Gender `protobuf:"varint,3,opt,name=gender,proto3,enum=doctors.Doctor_Gender" json:"gender,omitempty"`
	PhoneNumber  string        `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Specialities string        `protobuf:"bytes,5,opt,name=specialities,proto3" json:"specialities,omitempty"`
	SpecialNote  string        `protobuf:"bytes,10,opt,name=special_note,json=specialNote,proto3" json:"special_note,omitempty"`
}

func (x *CreateDoctorRequest) Reset() {
	*x = CreateDoctorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doctors_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDoctorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDoctorRequest) ProtoMessage() {}

func (x *CreateDoctorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doctors_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDoctorRequest.ProtoReflect.Descriptor instead.
func (*CreateDoctorRequest) Descriptor() ([]byte, []int) {
	return file_doctors_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateDoctorRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CreateDoctorRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateDoctorRequest) GetGender() Doctor_Gender {
	if x != nil {
		return x.Gender
	}
	return Doctor_UNSPECIFIED
}

func (x *CreateDoctorRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *CreateDoctorRequest) GetSpecialities() string {
	if x != nil {
		return x.Specialities
	}
	return ""
}

func (x *CreateDoctorRequest) GetSpecialNote() string {
	if x != nil {
		return x.SpecialNote
	}
	return ""
}

type PaginatedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count   int32   `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Results []int32 `protobuf:"varint,2,rep,packed,name=results,proto3" json:"results,omitempty"`
}

func (x *PaginatedResponse) Reset() {
	*x = PaginatedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doctors_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginatedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginatedResponse) ProtoMessage() {}

func (x *PaginatedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_doctors_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginatedResponse.ProtoReflect.Descriptor instead.
func (*PaginatedResponse) Descriptor() ([]byte, []int) {
	return file_doctors_service_proto_rawDescGZIP(), []int{3}
}

func (x *PaginatedResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *PaginatedResponse) GetResults() []int32 {
	if x != nil {
		return x.Results
	}
	return nil
}

type DoctorID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DoctorID) Reset() {
	*x = DoctorID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doctors_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoctorID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoctorID) ProtoMessage() {}

func (x *DoctorID) ProtoReflect() protoreflect.Message {
	mi := &file_doctors_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoctorID.ProtoReflect.Descriptor instead.
func (*DoctorID) Descriptor() ([]byte, []int) {
	return file_doctors_service_proto_rawDescGZIP(), []int{4}
}

func (x *DoctorID) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Doctor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Active       bool          `protobuf:"varint,2,opt,name=active,proto3" json:"active,omitempty"`
	Name         string        `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Gender       Doctor_Gender `protobuf:"varint,4,opt,name=gender,proto3,enum=doctors.Doctor_Gender" json:"gender,omitempty"`
	PhoneNumber  string        `protobuf:"bytes,5,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Specialities string        `protobuf:"bytes,6,opt,name=specialities,proto3" json:"specialities,omitempty"`
	SpecialNote  string        `protobuf:"bytes,7,opt,name=special_note,json=specialNote,proto3" json:"special_note,omitempty"`
}

func (x *Doctor) Reset() {
	*x = Doctor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doctors_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Doctor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Doctor) ProtoMessage() {}

func (x *Doctor) ProtoReflect() protoreflect.Message {
	mi := &file_doctors_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Doctor.ProtoReflect.Descriptor instead.
func (*Doctor) Descriptor() ([]byte, []int) {
	return file_doctors_service_proto_rawDescGZIP(), []int{5}
}

func (x *Doctor) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Doctor) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *Doctor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Doctor) GetGender() Doctor_Gender {
	if x != nil {
		return x.Gender
	}
	return Doctor_UNSPECIFIED
}

func (x *Doctor) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *Doctor) GetSpecialities() string {
	if x != nil {
		return x.Specialities
	}
	return ""
}

func (x *Doctor) GetSpecialNote() string {
	if x != nil {
		return x.SpecialNote
	}
	return ""
}

var File_doctors_service_proto protoreflect.FileDescriptor

var file_doctors_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73,
	0x22, 0x35, 0x0a, 0x0d, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x54, 0x0a, 0x0e, 0x44, 0x6f, 0x63, 0x74, 0x6f,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0xd9, 0x01,
	0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x2e, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x16, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61,
	0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61,
	0x6c, 0x5f, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x61, 0x6c, 0x4e, 0x6f, 0x74, 0x65, 0x22, 0x43, 0x0a, 0x11, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x1a,
	0x0a, 0x08, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x8f, 0x02, 0x0a, 0x06, 0x44,
	0x6f, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x2e, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x16, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x44, 0x6f, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x70, 0x65, 0x63,
	0x69, 0x61, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x70, 0x65, 0x63,
	0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x4e, 0x6f, 0x74, 0x65, 0x22, 0x2f, 0x0a, 0x06, 0x47,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x32, 0xcc, 0x01, 0x0a,
	0x0d, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x2e, 0x64, 0x6f,
	0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x44, 0x6f,
	0x63, 0x74, 0x6f, 0x72, 0x12, 0x44, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x63, 0x74, 0x6f,
	0x72, 0x73, 0x49, 0x44, 0x73, 0x12, 0x17, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e,
	0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x1c, 0x2e, 0x64, 0x6f, 0x63,
	0x74, 0x6f, 0x72, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f,
	0x72, 0x73, 0x2e, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x44, 0x42, 0x3c, 0x5a, 0x3a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x65, 0x6b, 0x43, 0x6c, 0x69,
	0x6e, 0x69, 0x63, 0x2f, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2d, 0x4d, 0x69, 0x63, 0x72,
	0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_doctors_service_proto_rawDescOnce sync.Once
	file_doctors_service_proto_rawDescData = file_doctors_service_proto_rawDesc
)

func file_doctors_service_proto_rawDescGZIP() []byte {
	file_doctors_service_proto_rawDescOnce.Do(func() {
		file_doctors_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_doctors_service_proto_rawDescData)
	})
	return file_doctors_service_proto_rawDescData
}

var file_doctors_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_doctors_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_doctors_service_proto_goTypes = []interface{}{
	(Doctor_Gender)(0),          // 0: doctors.Doctor.Gender
	(*DoctorRequest)(nil),       // 1: doctors.DoctorRequest
	(*DoctorsRequest)(nil),      // 2: doctors.DoctorsRequest
	(*CreateDoctorRequest)(nil), // 3: doctors.CreateDoctorRequest
	(*PaginatedResponse)(nil),   // 4: doctors.PaginatedResponse
	(*DoctorID)(nil),            // 5: doctors.DoctorID
	(*Doctor)(nil),              // 6: doctors.Doctor
}
var file_doctors_service_proto_depIdxs = []int32{
	0, // 0: doctors.CreateDoctorRequest.gender:type_name -> doctors.Doctor.Gender
	0, // 1: doctors.Doctor.gender:type_name -> doctors.Doctor.Gender
	1, // 2: doctors.DoctorService.GetDoctor:input_type -> doctors.DoctorRequest
	2, // 3: doctors.DoctorService.GetDoctorsIDs:input_type -> doctors.DoctorsRequest
	3, // 4: doctors.DoctorService.CreateDoctor:input_type -> doctors.CreateDoctorRequest
	6, // 5: doctors.DoctorService.GetDoctor:output_type -> doctors.Doctor
	4, // 6: doctors.DoctorService.GetDoctorsIDs:output_type -> doctors.PaginatedResponse
	5, // 7: doctors.DoctorService.CreateDoctor:output_type -> doctors.DoctorID
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_doctors_service_proto_init() }
func file_doctors_service_proto_init() {
	if File_doctors_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_doctors_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoctorRequest); i {
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
		file_doctors_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoctorsRequest); i {
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
		file_doctors_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDoctorRequest); i {
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
		file_doctors_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginatedResponse); i {
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
		file_doctors_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoctorID); i {
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
		file_doctors_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Doctor); i {
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
			RawDescriptor: file_doctors_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_doctors_service_proto_goTypes,
		DependencyIndexes: file_doctors_service_proto_depIdxs,
		EnumInfos:         file_doctors_service_proto_enumTypes,
		MessageInfos:      file_doctors_service_proto_msgTypes,
	}.Build()
	File_doctors_service_proto = out.File
	file_doctors_service_proto_rawDesc = nil
	file_doctors_service_proto_goTypes = nil
	file_doctors_service_proto_depIdxs = nil
}
