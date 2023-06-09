// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.0
// source: notification/notification-service.proto

package notification

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

type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notification string `protobuf:"bytes,1,opt,name=notification,proto3" json:"notification,omitempty"`
}

func (x *Notification) Reset() {
	*x = Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_notification_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_notification_notification_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_notification_notification_service_proto_rawDescGZIP(), []int{0}
}

func (x *Notification) GetNotification() string {
	if x != nil {
		return x.Notification
	}
	return ""
}

type GetNotificationsForUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetNotificationsForUserRequest) Reset() {
	*x = GetNotificationsForUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_notification_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotificationsForUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotificationsForUserRequest) ProtoMessage() {}

func (x *GetNotificationsForUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notification_notification_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotificationsForUserRequest.ProtoReflect.Descriptor instead.
func (*GetNotificationsForUserRequest) Descriptor() ([]byte, []int) {
	return file_notification_notification_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetNotificationsForUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetNotificationsForUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notifications []*Notification `protobuf:"bytes,1,rep,name=notifications,proto3" json:"notifications,omitempty"`
}

func (x *GetNotificationsForUserResponse) Reset() {
	*x = GetNotificationsForUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_notification_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotificationsForUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotificationsForUserResponse) ProtoMessage() {}

func (x *GetNotificationsForUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notification_notification_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotificationsForUserResponse.ProtoReflect.Descriptor instead.
func (*GetNotificationsForUserResponse) Descriptor() ([]byte, []int) {
	return file_notification_notification_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetNotificationsForUserResponse) GetNotifications() []*Notification {
	if x != nil {
		return x.Notifications
	}
	return nil
}

type NotificationType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RESERVATION_REQUEST bool `protobuf:"varint,1,opt,name=RESERVATION_REQUEST,json=RESERVATIONREQUEST,proto3" json:"RESERVATION_REQUEST,omitempty"`
	RESERVATION_CANCEL  bool `protobuf:"varint,2,opt,name=RESERVATION_CANCEL,json=RESERVATIONCANCEL,proto3" json:"RESERVATION_CANCEL,omitempty"`
	HOST_RATE           bool `protobuf:"varint,3,opt,name=HOST_RATE,json=HOSTRATE,proto3" json:"HOST_RATE,omitempty"`
	ACCOM_RATE          bool `protobuf:"varint,4,opt,name=ACCOM_RATE,json=ACCOMRATE,proto3" json:"ACCOM_RATE,omitempty"`
	GOT_FEATURED_HOST   bool `protobuf:"varint,5,opt,name=GOT_FEATURED_HOST,json=GOTFEATUREDHOST,proto3" json:"GOT_FEATURED_HOST,omitempty"`
	LOST_FEATURED_HOST  bool `protobuf:"varint,6,opt,name=LOST_FEATURED_HOST,json=LOSTFEATUREDHOST,proto3" json:"LOST_FEATURED_HOST,omitempty"`
	GUEST_ACCEPTED      bool `protobuf:"varint,7,opt,name=GUEST_ACCEPTED,json=GUESTACCEPTED,proto3" json:"GUEST_ACCEPTED,omitempty"`
}

func (x *NotificationType) Reset() {
	*x = NotificationType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_notification_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationType) ProtoMessage() {}

func (x *NotificationType) ProtoReflect() protoreflect.Message {
	mi := &file_notification_notification_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationType.ProtoReflect.Descriptor instead.
func (*NotificationType) Descriptor() ([]byte, []int) {
	return file_notification_notification_service_proto_rawDescGZIP(), []int{3}
}

func (x *NotificationType) GetRESERVATION_REQUEST() bool {
	if x != nil {
		return x.RESERVATION_REQUEST
	}
	return false
}

func (x *NotificationType) GetRESERVATION_CANCEL() bool {
	if x != nil {
		return x.RESERVATION_CANCEL
	}
	return false
}

func (x *NotificationType) GetHOST_RATE() bool {
	if x != nil {
		return x.HOST_RATE
	}
	return false
}

func (x *NotificationType) GetACCOM_RATE() bool {
	if x != nil {
		return x.ACCOM_RATE
	}
	return false
}

func (x *NotificationType) GetGOT_FEATURED_HOST() bool {
	if x != nil {
		return x.GOT_FEATURED_HOST
	}
	return false
}

func (x *NotificationType) GetLOST_FEATURED_HOST() bool {
	if x != nil {
		return x.LOST_FEATURED_HOST
	}
	return false
}

func (x *NotificationType) GetGUEST_ACCEPTED() bool {
	if x != nil {
		return x.GUEST_ACCEPTED
	}
	return false
}

type AddNotificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notification *Notification `protobuf:"bytes,1,opt,name=notification,proto3" json:"notification,omitempty"`
}

func (x *AddNotificationRequest) Reset() {
	*x = AddNotificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_notification_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNotificationRequest) ProtoMessage() {}

func (x *AddNotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notification_notification_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNotificationRequest.ProtoReflect.Descriptor instead.
func (*AddNotificationRequest) Descriptor() ([]byte, []int) {
	return file_notification_notification_service_proto_rawDescGZIP(), []int{4}
}

func (x *AddNotificationRequest) GetNotification() *Notification {
	if x != nil {
		return x.Notification
	}
	return nil
}

type AddNotificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *AddNotificationResponse) Reset() {
	*x = AddNotificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_notification_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNotificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNotificationResponse) ProtoMessage() {}

func (x *AddNotificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notification_notification_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNotificationResponse.ProtoReflect.Descriptor instead.
func (*AddNotificationResponse) Descriptor() ([]byte, []int) {
	return file_notification_notification_service_proto_rawDescGZIP(), []int{5}
}

func (x *AddNotificationResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_notification_notification_service_proto protoreflect.FileDescriptor

var file_notification_notification_service_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x38, 0x0a, 0x1e,
	0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x56, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0d, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xaf,
	0x02, 0x0a, 0x10, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x52, 0x45, 0x53, 0x45, 0x52, 0x56, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x12, 0x52, 0x45, 0x53, 0x45, 0x52, 0x56, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x52, 0x45, 0x51,
	0x55, 0x45, 0x53, 0x54, 0x12, 0x2d, 0x0a, 0x12, 0x52, 0x45, 0x53, 0x45, 0x52, 0x56, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x11, 0x52, 0x45, 0x53, 0x45, 0x52, 0x56, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x43, 0x41, 0x4e,
	0x43, 0x45, 0x4c, 0x12, 0x1b, 0x0a, 0x09, 0x48, 0x4f, 0x53, 0x54, 0x5f, 0x52, 0x41, 0x54, 0x45,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x48, 0x4f, 0x53, 0x54, 0x52, 0x41, 0x54, 0x45,
	0x12, 0x1d, 0x0a, 0x0a, 0x41, 0x43, 0x43, 0x4f, 0x4d, 0x5f, 0x52, 0x41, 0x54, 0x45, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x41, 0x43, 0x43, 0x4f, 0x4d, 0x52, 0x41, 0x54, 0x45, 0x12,
	0x2a, 0x0a, 0x11, 0x47, 0x4f, 0x54, 0x5f, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x44, 0x5f,
	0x48, 0x4f, 0x53, 0x54, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x47, 0x4f, 0x54, 0x46,
	0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x44, 0x48, 0x4f, 0x53, 0x54, 0x12, 0x2c, 0x0a, 0x12, 0x4c,
	0x4f, 0x53, 0x54, 0x5f, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x44, 0x5f, 0x48, 0x4f, 0x53,
	0x54, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x4c, 0x4f, 0x53, 0x54, 0x46, 0x45, 0x41,
	0x54, 0x55, 0x52, 0x45, 0x44, 0x48, 0x4f, 0x53, 0x54, 0x12, 0x25, 0x0a, 0x0e, 0x47, 0x55, 0x45,
	0x53, 0x54, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0d, 0x47, 0x55, 0x45, 0x53, 0x54, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44,
	0x22, 0x4b, 0x0a, 0x16, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x0c, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x35, 0x0a,
	0x17, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0xf7, 0x01, 0x0a, 0x16, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x50, 0x43, 0x12,
	0x7c, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x47, 0x65, 0x74,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x46, 0x6f, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x47, 0x65,
	0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x46, 0x6f,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x7d, 0x12, 0x5f, 0x0a,
	0x0f, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x17, 0x2e, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x41, 0x64, 0x64, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x11, 0x2f, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x64, 0x64, 0x42, 0x14,
	0x5a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notification_notification_service_proto_rawDescOnce sync.Once
	file_notification_notification_service_proto_rawDescData = file_notification_notification_service_proto_rawDesc
)

func file_notification_notification_service_proto_rawDescGZIP() []byte {
	file_notification_notification_service_proto_rawDescOnce.Do(func() {
		file_notification_notification_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_notification_notification_service_proto_rawDescData)
	})
	return file_notification_notification_service_proto_rawDescData
}

var file_notification_notification_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_notification_notification_service_proto_goTypes = []interface{}{
	(*Notification)(nil),                    // 0: Notification
	(*GetNotificationsForUserRequest)(nil),  // 1: GetNotificationsForUserRequest
	(*GetNotificationsForUserResponse)(nil), // 2: GetNotificationsForUserResponse
	(*NotificationType)(nil),                // 3: NotificationType
	(*AddNotificationRequest)(nil),          // 4: AddNotificationRequest
	(*AddNotificationResponse)(nil),         // 5: AddNotificationResponse
}
var file_notification_notification_service_proto_depIdxs = []int32{
	0, // 0: GetNotificationsForUserResponse.notifications:type_name -> Notification
	0, // 1: AddNotificationRequest.notification:type_name -> Notification
	1, // 2: NotificationServiceRPC.GetNotificationsForUser:input_type -> GetNotificationsForUserRequest
	4, // 3: NotificationServiceRPC.AddNotification:input_type -> AddNotificationRequest
	2, // 4: NotificationServiceRPC.GetNotificationsForUser:output_type -> GetNotificationsForUserResponse
	5, // 5: NotificationServiceRPC.AddNotification:output_type -> AddNotificationResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_notification_notification_service_proto_init() }
func file_notification_notification_service_proto_init() {
	if File_notification_notification_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notification_notification_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notification); i {
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
		file_notification_notification_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotificationsForUserRequest); i {
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
		file_notification_notification_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotificationsForUserResponse); i {
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
		file_notification_notification_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationType); i {
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
		file_notification_notification_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddNotificationRequest); i {
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
		file_notification_notification_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddNotificationResponse); i {
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
			RawDescriptor: file_notification_notification_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notification_notification_service_proto_goTypes,
		DependencyIndexes: file_notification_notification_service_proto_depIdxs,
		MessageInfos:      file_notification_notification_service_proto_msgTypes,
	}.Build()
	File_notification_notification_service_proto = out.File
	file_notification_notification_service_proto_rawDesc = nil
	file_notification_notification_service_proto_goTypes = nil
	file_notification_notification_service_proto_depIdxs = nil
}
