// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.0
// source: review/review-service.proto

package review

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RateHostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GuestId string  `protobuf:"bytes,1,opt,name=guestId,proto3" json:"guestId,omitempty"`
	HostId  string  `protobuf:"bytes,2,opt,name=hostId,proto3" json:"hostId,omitempty"`
	Rating  float64 `protobuf:"fixed64,3,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *RateHostRequest) Reset() {
	*x = RateHostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_review_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateHostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateHostRequest) ProtoMessage() {}

func (x *RateHostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_review_review_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateHostRequest.ProtoReflect.Descriptor instead.
func (*RateHostRequest) Descriptor() ([]byte, []int) {
	return file_review_review_service_proto_rawDescGZIP(), []int{0}
}

func (x *RateHostRequest) GetGuestId() string {
	if x != nil {
		return x.GuestId
	}
	return ""
}

func (x *RateHostRequest) GetHostId() string {
	if x != nil {
		return x.HostId
	}
	return ""
}

func (x *RateHostRequest) GetRating() float64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

type RateAccommodationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GuestId         string  `protobuf:"bytes,1,opt,name=guestId,proto3" json:"guestId,omitempty"`
	AccommodationId string  `protobuf:"bytes,2,opt,name=accommodationId,proto3" json:"accommodationId,omitempty"`
	Rating          float64 `protobuf:"fixed64,3,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *RateAccommodationRequest) Reset() {
	*x = RateAccommodationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_review_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateAccommodationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateAccommodationRequest) ProtoMessage() {}

func (x *RateAccommodationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_review_review_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateAccommodationRequest.ProtoReflect.Descriptor instead.
func (*RateAccommodationRequest) Descriptor() ([]byte, []int) {
	return file_review_review_service_proto_rawDescGZIP(), []int{1}
}

func (x *RateAccommodationRequest) GetGuestId() string {
	if x != nil {
		return x.GuestId
	}
	return ""
}

func (x *RateAccommodationRequest) GetAccommodationId() string {
	if x != nil {
		return x.AccommodationId
	}
	return ""
}

func (x *RateAccommodationRequest) GetRating() float64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

type RateHostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RateHostResponse) Reset() {
	*x = RateHostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_review_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateHostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateHostResponse) ProtoMessage() {}

func (x *RateHostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_review_review_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateHostResponse.ProtoReflect.Descriptor instead.
func (*RateHostResponse) Descriptor() ([]byte, []int) {
	return file_review_review_service_proto_rawDescGZIP(), []int{2}
}

type RateAccommodationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RateAccommodationResponse) Reset() {
	*x = RateAccommodationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_review_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateAccommodationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateAccommodationResponse) ProtoMessage() {}

func (x *RateAccommodationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_review_review_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateAccommodationResponse.ProtoReflect.Descriptor instead.
func (*RateAccommodationResponse) Descriptor() ([]byte, []int) {
	return file_review_review_service_proto_rawDescGZIP(), []int{3}
}

type GetRatingsForHostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostId string `protobuf:"bytes,1,opt,name=hostId,proto3" json:"hostId,omitempty"`
}

func (x *GetRatingsForHostRequest) Reset() {
	*x = GetRatingsForHostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_review_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRatingsForHostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRatingsForHostRequest) ProtoMessage() {}

func (x *GetRatingsForHostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_review_review_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRatingsForHostRequest.ProtoReflect.Descriptor instead.
func (*GetRatingsForHostRequest) Descriptor() ([]byte, []int) {
	return file_review_review_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetRatingsForHostRequest) GetHostId() string {
	if x != nil {
		return x.HostId
	}
	return ""
}

type GetRatingsForHostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ratings       []*HostRating `protobuf:"bytes,1,rep,name=ratings,proto3" json:"ratings,omitempty"`
	AverageRating float64       `protobuf:"fixed64,2,opt,name=averageRating,proto3" json:"averageRating,omitempty"`
}

func (x *GetRatingsForHostResponse) Reset() {
	*x = GetRatingsForHostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_review_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRatingsForHostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRatingsForHostResponse) ProtoMessage() {}

func (x *GetRatingsForHostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_review_review_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRatingsForHostResponse.ProtoReflect.Descriptor instead.
func (*GetRatingsForHostResponse) Descriptor() ([]byte, []int) {
	return file_review_review_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetRatingsForHostResponse) GetRatings() []*HostRating {
	if x != nil {
		return x.Ratings
	}
	return nil
}

func (x *GetRatingsForHostResponse) GetAverageRating() float64 {
	if x != nil {
		return x.AverageRating
	}
	return 0
}

type HostRating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string  `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Rating float64 `protobuf:"fixed64,2,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *HostRating) Reset() {
	*x = HostRating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_review_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostRating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostRating) ProtoMessage() {}

func (x *HostRating) ProtoReflect() protoreflect.Message {
	mi := &file_review_review_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostRating.ProtoReflect.Descriptor instead.
func (*HostRating) Descriptor() ([]byte, []int) {
	return file_review_review_service_proto_rawDescGZIP(), []int{6}
}

func (x *HostRating) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *HostRating) GetRating() float64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

type GetRatingsForAccommodationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccommodationId string `protobuf:"bytes,1,opt,name=accommodationId,proto3" json:"accommodationId,omitempty"`
}

func (x *GetRatingsForAccommodationRequest) Reset() {
	*x = GetRatingsForAccommodationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_review_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRatingsForAccommodationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRatingsForAccommodationRequest) ProtoMessage() {}

func (x *GetRatingsForAccommodationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_review_review_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRatingsForAccommodationRequest.ProtoReflect.Descriptor instead.
func (*GetRatingsForAccommodationRequest) Descriptor() ([]byte, []int) {
	return file_review_review_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetRatingsForAccommodationRequest) GetAccommodationId() string {
	if x != nil {
		return x.AccommodationId
	}
	return ""
}

type GetRatingsForAccommodationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ratings       []*AccommodationRating `protobuf:"bytes,1,rep,name=ratings,proto3" json:"ratings,omitempty"`
	AverageRating float64                `protobuf:"fixed64,2,opt,name=averageRating,proto3" json:"averageRating,omitempty"`
}

func (x *GetRatingsForAccommodationResponse) Reset() {
	*x = GetRatingsForAccommodationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_review_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRatingsForAccommodationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRatingsForAccommodationResponse) ProtoMessage() {}

func (x *GetRatingsForAccommodationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_review_review_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRatingsForAccommodationResponse.ProtoReflect.Descriptor instead.
func (*GetRatingsForAccommodationResponse) Descriptor() ([]byte, []int) {
	return file_review_review_service_proto_rawDescGZIP(), []int{8}
}

func (x *GetRatingsForAccommodationResponse) GetRatings() []*AccommodationRating {
	if x != nil {
		return x.Ratings
	}
	return nil
}

func (x *GetRatingsForAccommodationResponse) GetAverageRating() float64 {
	if x != nil {
		return x.AverageRating
	}
	return 0
}

type AccommodationRating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string  `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Rating float64 `protobuf:"fixed64,2,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *AccommodationRating) Reset() {
	*x = AccommodationRating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_review_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccommodationRating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccommodationRating) ProtoMessage() {}

func (x *AccommodationRating) ProtoReflect() protoreflect.Message {
	mi := &file_review_review_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccommodationRating.ProtoReflect.Descriptor instead.
func (*AccommodationRating) Descriptor() ([]byte, []int) {
	return file_review_review_service_proto_rawDescGZIP(), []int{9}
}

func (x *AccommodationRating) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AccommodationRating) GetRating() float64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

var File_review_review_service_proto protoreflect.FileDescriptor

var file_review_review_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x0f,
	0x52, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x67, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x67, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x6f, 0x73,
	0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x76, 0x0a, 0x18, 0x52, 0x61, 0x74,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12,
	0x28, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x22, 0x12, 0x0a, 0x10, 0x52, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x0a, 0x19, 0x52, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x32, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x46, 0x6f, 0x72, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x22, 0x68, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x46, 0x6f, 0x72, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x07, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x07, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x76,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0d, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x22, 0x3c, 0x0a, 0x0a, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x4d,
	0x0a, 0x21, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x46, 0x6f, 0x72, 0x41,
	0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x63,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x7a, 0x0a,
	0x22, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x46, 0x6f, 0x72, 0x41, 0x63,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x61, 0x76, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x45, 0x0a, 0x13, 0x41, 0x63, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x32, 0xce, 0x03, 0x0a, 0x10, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x50, 0x43, 0x12, 0x48, 0x0a, 0x08, 0x52, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x73,
	0x74, 0x12, 0x10, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01,
	0x2a, 0x22, 0x0c, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x12,
	0x6c, 0x0a, 0x11, 0x52, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x69, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x46, 0x6f, 0x72, 0x48, 0x6f,
	0x73, 0x74, 0x12, 0x19, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x46,
	0x6f, 0x72, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x46, 0x6f, 0x72, 0x48, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x17, 0x12, 0x15, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x2f,
	0x7b, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x7d, 0x12, 0x96, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74,
	0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x46, 0x6f, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x46, 0x6f, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x46, 0x6f, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x12, 0x27, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x7b, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x7d, 0x42, 0x0e, 0x5a, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_review_review_service_proto_rawDescOnce sync.Once
	file_review_review_service_proto_rawDescData = file_review_review_service_proto_rawDesc
)

func file_review_review_service_proto_rawDescGZIP() []byte {
	file_review_review_service_proto_rawDescOnce.Do(func() {
		file_review_review_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_review_review_service_proto_rawDescData)
	})
	return file_review_review_service_proto_rawDescData
}

var file_review_review_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_review_review_service_proto_goTypes = []interface{}{
	(*RateHostRequest)(nil),                    // 0: RateHostRequest
	(*RateAccommodationRequest)(nil),           // 1: RateAccommodationRequest
	(*RateHostResponse)(nil),                   // 2: RateHostResponse
	(*RateAccommodationResponse)(nil),          // 3: RateAccommodationResponse
	(*GetRatingsForHostRequest)(nil),           // 4: GetRatingsForHostRequest
	(*GetRatingsForHostResponse)(nil),          // 5: GetRatingsForHostResponse
	(*HostRating)(nil),                         // 6: HostRating
	(*GetRatingsForAccommodationRequest)(nil),  // 7: GetRatingsForAccommodationRequest
	(*GetRatingsForAccommodationResponse)(nil), // 8: GetRatingsForAccommodationResponse
	(*AccommodationRating)(nil),                // 9: AccommodationRating
}
var file_review_review_service_proto_depIdxs = []int32{
	6, // 0: GetRatingsForHostResponse.ratings:type_name -> HostRating
	9, // 1: GetRatingsForAccommodationResponse.ratings:type_name -> AccommodationRating
	0, // 2: RatingServiceRPC.RateHost:input_type -> RateHostRequest
	1, // 3: RatingServiceRPC.RateAccommodation:input_type -> RateAccommodationRequest
	4, // 4: RatingServiceRPC.GetRatingsForHost:input_type -> GetRatingsForHostRequest
	7, // 5: RatingServiceRPC.GetRatingsForAccommodation:input_type -> GetRatingsForAccommodationRequest
	2, // 6: RatingServiceRPC.RateHost:output_type -> RateHostResponse
	3, // 7: RatingServiceRPC.RateAccommodation:output_type -> RateAccommodationResponse
	5, // 8: RatingServiceRPC.GetRatingsForHost:output_type -> GetRatingsForHostResponse
	8, // 9: RatingServiceRPC.GetRatingsForAccommodation:output_type -> GetRatingsForAccommodationResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_review_review_service_proto_init() }
func file_review_review_service_proto_init() {
	if File_review_review_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_review_review_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateHostRequest); i {
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
		file_review_review_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateAccommodationRequest); i {
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
		file_review_review_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateHostResponse); i {
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
		file_review_review_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateAccommodationResponse); i {
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
		file_review_review_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRatingsForHostRequest); i {
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
		file_review_review_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRatingsForHostResponse); i {
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
		file_review_review_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostRating); i {
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
		file_review_review_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRatingsForAccommodationRequest); i {
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
		file_review_review_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRatingsForAccommodationResponse); i {
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
		file_review_review_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccommodationRating); i {
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
			RawDescriptor: file_review_review_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_review_review_service_proto_goTypes,
		DependencyIndexes: file_review_review_service_proto_depIdxs,
		MessageInfos:      file_review_review_service_proto_msgTypes,
	}.Build()
	File_review_review_service_proto = out.File
	file_review_review_service_proto_rawDesc = nil
	file_review_review_service_proto_goTypes = nil
	file_review_review_service_proto_depIdxs = nil
}
