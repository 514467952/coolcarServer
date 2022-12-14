// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.11.0
// source: rental/rental.proto

package rentalpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type TripStatus int32

const (
	TripStatus_TS_NOT_SPECIFIED TripStatus = 0 //无规定
	TripStatus_IN_PROGRESS      TripStatus = 1 //在行程中
	TripStatus_FINISHED         TripStatus = 2 //行程结束
)

// Enum value maps for TripStatus.
var (
	TripStatus_name = map[int32]string{
		0: "TS_NOT_SPECIFIED",
		1: "IN_PROGRESS",
		2: "FINISHED",
	}
	TripStatus_value = map[string]int32{
		"TS_NOT_SPECIFIED": 0,
		"IN_PROGRESS":      1,
		"FINISHED":         2,
	}
)

func (x TripStatus) Enum() *TripStatus {
	p := new(TripStatus)
	*p = x
	return p
}

func (x TripStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TripStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_rental_rental_proto_enumTypes[0].Descriptor()
}

func (TripStatus) Type() protoreflect.EnumType {
	return &file_rental_rental_proto_enumTypes[0]
}

func (x TripStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TripStatus.Descriptor instead.
func (TripStatus) EnumDescriptor() ([]byte, []int) {
	return file_rental_rental_proto_rawDescGZIP(), []int{0}
}

//定义点的位置
type Loaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`   //纬度
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"` //经度
}

func (x *Loaction) Reset() {
	*x = Loaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rental_rental_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Loaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Loaction) ProtoMessage() {}

func (x *Loaction) ProtoReflect() protoreflect.Message {
	mi := &file_rental_rental_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Loaction.ProtoReflect.Descriptor instead.
func (*Loaction) Descriptor() ([]byte, []int) {
	return file_rental_rental_proto_rawDescGZIP(), []int{0}
}

func (x *Loaction) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Loaction) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

//定义点的状态
type LocationStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Loaction *Loaction `protobuf:"bytes,1,opt,name=Loaction,proto3" json:"Loaction,omitempty"`                   //点的位置
	FeeCent  int32     `protobuf:"varint,2,opt,name=fee_cent,json=feeCent,proto3" json:"fee_cent,omitempty"`     //点的费用
	KmDriven float64   `protobuf:"fixed64,3,opt,name=km_driven,json=kmDriven,proto3" json:"km_driven,omitempty"` //公里数
	PoiName  string    `protobuf:"bytes,4,opt,name=poi_name,json=poiName,proto3" json:"poi_name,omitempty"`      //行程的描述
}

func (x *LocationStatus) Reset() {
	*x = LocationStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rental_rental_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocationStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationStatus) ProtoMessage() {}

func (x *LocationStatus) ProtoReflect() protoreflect.Message {
	mi := &file_rental_rental_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationStatus.ProtoReflect.Descriptor instead.
func (*LocationStatus) Descriptor() ([]byte, []int) {
	return file_rental_rental_proto_rawDescGZIP(), []int{1}
}

func (x *LocationStatus) GetLoaction() *Loaction {
	if x != nil {
		return x.Loaction
	}
	return nil
}

func (x *LocationStatus) GetFeeCent() int32 {
	if x != nil {
		return x.FeeCent
	}
	return 0
}

func (x *LocationStatus) GetKmDriven() float64 {
	if x != nil {
		return x.KmDriven
	}
	return 0
}

func (x *LocationStatus) GetPoiName() string {
	if x != nil {
		return x.PoiName
	}
	return ""
}

//行程的对象
type TripEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Trip *Trip  `protobuf:"bytes,2,opt,name=trip,proto3" json:"trip,omitempty"`
}

func (x *TripEntity) Reset() {
	*x = TripEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rental_rental_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TripEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TripEntity) ProtoMessage() {}

func (x *TripEntity) ProtoReflect() protoreflect.Message {
	mi := &file_rental_rental_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TripEntity.ProtoReflect.Descriptor instead.
func (*TripEntity) Descriptor() ([]byte, []int) {
	return file_rental_rental_proto_rawDescGZIP(), []int{2}
}

func (x *TripEntity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TripEntity) GetTrip() *Trip {
	if x != nil {
		return x.Trip
	}
	return nil
}

//行程的数据
type Trip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID string          `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`                      //开车的人
	CarID     string          `protobuf:"bytes,2,opt,name=carID,proto3" json:"carID,omitempty"`                              //开的哪辆车
	Start     *LocationStatus `protobuf:"bytes,3,opt,name=start,proto3" json:"start,omitempty"`                              //起点
	Current   *LocationStatus `protobuf:"bytes,4,opt,name=current,proto3" json:"current,omitempty"`                          //当前点
	End       *LocationStatus `protobuf:"bytes,5,opt,name=end,proto3" json:"end,omitempty"`                                  //结束点
	Status    TripStatus      `protobuf:"varint,6,opt,name=status,proto3,enum=rental.v1.TripStatus" json:"status,omitempty"` //行程的状态
}

func (x *Trip) Reset() {
	*x = Trip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rental_rental_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trip) ProtoMessage() {}

func (x *Trip) ProtoReflect() protoreflect.Message {
	mi := &file_rental_rental_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trip.ProtoReflect.Descriptor instead.
func (*Trip) Descriptor() ([]byte, []int) {
	return file_rental_rental_proto_rawDescGZIP(), []int{3}
}

func (x *Trip) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *Trip) GetCarID() string {
	if x != nil {
		return x.CarID
	}
	return ""
}

func (x *Trip) GetStart() *LocationStatus {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *Trip) GetCurrent() *LocationStatus {
	if x != nil {
		return x.Current
	}
	return nil
}

func (x *Trip) GetEnd() *LocationStatus {
	if x != nil {
		return x.End
	}
	return nil
}

func (x *Trip) GetStatus() TripStatus {
	if x != nil {
		return x.Status
	}
	return TripStatus_TS_NOT_SPECIFIED
}

type CreateTripRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start *Loaction `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	CarId string    `protobuf:"bytes,2,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
}

func (x *CreateTripRequest) Reset() {
	*x = CreateTripRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rental_rental_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTripRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTripRequest) ProtoMessage() {}

func (x *CreateTripRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rental_rental_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTripRequest.ProtoReflect.Descriptor instead.
func (*CreateTripRequest) Descriptor() ([]byte, []int) {
	return file_rental_rental_proto_rawDescGZIP(), []int{4}
}

func (x *CreateTripRequest) GetStart() *Loaction {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *CreateTripRequest) GetCarId() string {
	if x != nil {
		return x.CarId
	}
	return ""
}

type GetTripRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` //根据id获取行程的信息
}

func (x *GetTripRequest) Reset() {
	*x = GetTripRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rental_rental_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTripRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTripRequest) ProtoMessage() {}

func (x *GetTripRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rental_rental_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTripRequest.ProtoReflect.Descriptor instead.
func (*GetTripRequest) Descriptor() ([]byte, []int) {
	return file_rental_rental_proto_rawDescGZIP(), []int{5}
}

func (x *GetTripRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetTripsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status TripStatus `protobuf:"varint,1,opt,name=status,proto3,enum=rental.v1.TripStatus" json:"status,omitempty"` //根据行程的状态查找
}

func (x *GetTripsRequest) Reset() {
	*x = GetTripsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rental_rental_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTripsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTripsRequest) ProtoMessage() {}

func (x *GetTripsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rental_rental_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTripsRequest.ProtoReflect.Descriptor instead.
func (*GetTripsRequest) Descriptor() ([]byte, []int) {
	return file_rental_rental_proto_rawDescGZIP(), []int{6}
}

func (x *GetTripsRequest) GetStatus() TripStatus {
	if x != nil {
		return x.Status
	}
	return TripStatus_TS_NOT_SPECIFIED
}

type GetTripsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Trips []*TripEntity `protobuf:"bytes,1,rep,name=trips,proto3" json:"trips,omitempty"`
}

func (x *GetTripsResponse) Reset() {
	*x = GetTripsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rental_rental_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTripsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTripsResponse) ProtoMessage() {}

func (x *GetTripsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rental_rental_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTripsResponse.ProtoReflect.Descriptor instead.
func (*GetTripsResponse) Descriptor() ([]byte, []int) {
	return file_rental_rental_proto_rawDescGZIP(), []int{7}
}

func (x *GetTripsResponse) GetTrips() []*TripEntity {
	if x != nil {
		return x.Trips
	}
	return nil
}

type UpdateTripRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                           //行程id
	Current *Loaction `protobuf:"bytes,2,opt,name=current,proto3" json:"current,omitempty"`                 //当前行程的位置
	EndTrip bool      `protobuf:"varint,3,opt,name=end_trip,json=endTrip,proto3" json:"end_trip,omitempty"` //是否结束行程
}

func (x *UpdateTripRequest) Reset() {
	*x = UpdateTripRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rental_rental_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTripRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTripRequest) ProtoMessage() {}

func (x *UpdateTripRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rental_rental_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTripRequest.ProtoReflect.Descriptor instead.
func (*UpdateTripRequest) Descriptor() ([]byte, []int) {
	return file_rental_rental_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateTripRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateTripRequest) GetCurrent() *Loaction {
	if x != nil {
		return x.Current
	}
	return nil
}

func (x *UpdateTripRequest) GetEndTrip() bool {
	if x != nil {
		return x.EndTrip
	}
	return false
}

var File_rental_rental_proto protoreflect.FileDescriptor

var file_rental_rental_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2f, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31,
	0x22, 0x44, 0x0a, 0x08, 0x4c, 0x6f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08,
	0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e,
	0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x94, 0x01, 0x0a, 0x0e, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x08, 0x4c, 0x6f, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x65,
	0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x08, 0x4c, 0x6f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x65,
	0x65, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x66, 0x65,
	0x65, 0x43, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6b, 0x6d, 0x5f, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6b, 0x6d, 0x44, 0x72, 0x69, 0x76,
	0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6f, 0x69, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6f, 0x69, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x41, 0x0a,
	0x0a, 0x54, 0x72, 0x69, 0x70, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x04, 0x74,
	0x72, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x69, 0x70, 0x52, 0x04, 0x74, 0x72, 0x69, 0x70,
	0x22, 0xfc, 0x01, 0x0a, 0x04, 0x54, 0x72, 0x69, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x61, 0x72, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x61, 0x72, 0x49, 0x44, 0x12, 0x2f, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72,
	0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x33,
	0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x03, 0x65, 0x6e, 0x64,
	0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x15, 0x2e, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x69,
	0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x55, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x6f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12,
	0x15, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x63, 0x61, 0x72, 0x49, 0x64, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x40, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54,
	0x72, 0x69, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x72, 0x65,
	0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x69, 0x70, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3f, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x54, 0x72, 0x69, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b,
	0x0a, 0x05, 0x74, 0x72, 0x69, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x69, 0x70, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x52, 0x05, 0x74, 0x72, 0x69, 0x70, 0x73, 0x22, 0x6d, 0x0a, 0x11, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x2d, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x72, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x72, 0x69, 0x70, 0x2a, 0x41, 0x0a, 0x0a, 0x54, 0x72,
	0x69, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x53, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f,
	0x0a, 0x0b, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12,
	0x0c, 0x0a, 0x08, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x02, 0x32, 0x89, 0x02,
	0x0a, 0x0b, 0x54, 0x72, 0x69, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x70, 0x12, 0x1c, 0x2e, 0x72, 0x65,
	0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72,
	0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x72, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x69, 0x70, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x35, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x12, 0x19, 0x2e, 0x72, 0x65,
	0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x72, 0x69, 0x70, 0x12, 0x43, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x72,
	0x69, 0x70, 0x73, 0x12, 0x1a, 0x2e, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x72, 0x69, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x70, 0x12, 0x1c, 0x2e, 0x72, 0x65, 0x6e,
	0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x72, 0x65, 0x6e, 0x74, 0x61,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x69, 0x70, 0x42, 0x24, 0x5a, 0x22, 0x63, 0x6f, 0x6f,
	0x6c, 0x63, 0x61, 0x72, 0x2f, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rental_rental_proto_rawDescOnce sync.Once
	file_rental_rental_proto_rawDescData = file_rental_rental_proto_rawDesc
)

func file_rental_rental_proto_rawDescGZIP() []byte {
	file_rental_rental_proto_rawDescOnce.Do(func() {
		file_rental_rental_proto_rawDescData = protoimpl.X.CompressGZIP(file_rental_rental_proto_rawDescData)
	})
	return file_rental_rental_proto_rawDescData
}

var file_rental_rental_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rental_rental_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_rental_rental_proto_goTypes = []interface{}{
	(TripStatus)(0),           // 0: rental.v1.TripStatus
	(*Loaction)(nil),          // 1: rental.v1.Loaction
	(*LocationStatus)(nil),    // 2: rental.v1.LocationStatus
	(*TripEntity)(nil),        // 3: rental.v1.TripEntity
	(*Trip)(nil),              // 4: rental.v1.Trip
	(*CreateTripRequest)(nil), // 5: rental.v1.CreateTripRequest
	(*GetTripRequest)(nil),    // 6: rental.v1.GetTripRequest
	(*GetTripsRequest)(nil),   // 7: rental.v1.GetTripsRequest
	(*GetTripsResponse)(nil),  // 8: rental.v1.GetTripsResponse
	(*UpdateTripRequest)(nil), // 9: rental.v1.UpdateTripRequest
}
var file_rental_rental_proto_depIdxs = []int32{
	1,  // 0: rental.v1.LocationStatus.Loaction:type_name -> rental.v1.Loaction
	4,  // 1: rental.v1.TripEntity.trip:type_name -> rental.v1.Trip
	2,  // 2: rental.v1.Trip.start:type_name -> rental.v1.LocationStatus
	2,  // 3: rental.v1.Trip.current:type_name -> rental.v1.LocationStatus
	2,  // 4: rental.v1.Trip.end:type_name -> rental.v1.LocationStatus
	0,  // 5: rental.v1.Trip.status:type_name -> rental.v1.TripStatus
	1,  // 6: rental.v1.CreateTripRequest.start:type_name -> rental.v1.Loaction
	0,  // 7: rental.v1.GetTripsRequest.status:type_name -> rental.v1.TripStatus
	3,  // 8: rental.v1.GetTripsResponse.trips:type_name -> rental.v1.TripEntity
	1,  // 9: rental.v1.UpdateTripRequest.current:type_name -> rental.v1.Loaction
	5,  // 10: rental.v1.TripService.CreateTrip:input_type -> rental.v1.CreateTripRequest
	6,  // 11: rental.v1.TripService.GetTrip:input_type -> rental.v1.GetTripRequest
	7,  // 12: rental.v1.TripService.GetTrips:input_type -> rental.v1.GetTripsRequest
	9,  // 13: rental.v1.TripService.UpdateTrip:input_type -> rental.v1.UpdateTripRequest
	3,  // 14: rental.v1.TripService.CreateTrip:output_type -> rental.v1.TripEntity
	4,  // 15: rental.v1.TripService.GetTrip:output_type -> rental.v1.Trip
	8,  // 16: rental.v1.TripService.GetTrips:output_type -> rental.v1.GetTripsResponse
	4,  // 17: rental.v1.TripService.UpdateTrip:output_type -> rental.v1.Trip
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_rental_rental_proto_init() }
func file_rental_rental_proto_init() {
	if File_rental_rental_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rental_rental_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Loaction); i {
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
		file_rental_rental_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocationStatus); i {
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
		file_rental_rental_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TripEntity); i {
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
		file_rental_rental_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trip); i {
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
		file_rental_rental_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTripRequest); i {
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
		file_rental_rental_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTripRequest); i {
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
		file_rental_rental_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTripsRequest); i {
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
		file_rental_rental_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTripsResponse); i {
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
		file_rental_rental_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTripRequest); i {
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
			RawDescriptor: file_rental_rental_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rental_rental_proto_goTypes,
		DependencyIndexes: file_rental_rental_proto_depIdxs,
		EnumInfos:         file_rental_rental_proto_enumTypes,
		MessageInfos:      file_rental_rental_proto_msgTypes,
	}.Build()
	File_rental_rental_proto = out.File
	file_rental_rental_proto_rawDesc = nil
	file_rental_rental_proto_goTypes = nil
	file_rental_rental_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TripServiceClient is the client API for TripService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TripServiceClient interface {
	CreateTrip(ctx context.Context, in *CreateTripRequest, opts ...grpc.CallOption) (*TripEntity, error)
	GetTrip(ctx context.Context, in *GetTripRequest, opts ...grpc.CallOption) (*Trip, error)
	GetTrips(ctx context.Context, in *GetTripsRequest, opts ...grpc.CallOption) (*GetTripsResponse, error)
	UpdateTrip(ctx context.Context, in *UpdateTripRequest, opts ...grpc.CallOption) (*Trip, error)
}

type tripServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTripServiceClient(cc grpc.ClientConnInterface) TripServiceClient {
	return &tripServiceClient{cc}
}

func (c *tripServiceClient) CreateTrip(ctx context.Context, in *CreateTripRequest, opts ...grpc.CallOption) (*TripEntity, error) {
	out := new(TripEntity)
	err := c.cc.Invoke(ctx, "/rental.v1.TripService/CreateTrip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tripServiceClient) GetTrip(ctx context.Context, in *GetTripRequest, opts ...grpc.CallOption) (*Trip, error) {
	out := new(Trip)
	err := c.cc.Invoke(ctx, "/rental.v1.TripService/GetTrip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tripServiceClient) GetTrips(ctx context.Context, in *GetTripsRequest, opts ...grpc.CallOption) (*GetTripsResponse, error) {
	out := new(GetTripsResponse)
	err := c.cc.Invoke(ctx, "/rental.v1.TripService/GetTrips", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tripServiceClient) UpdateTrip(ctx context.Context, in *UpdateTripRequest, opts ...grpc.CallOption) (*Trip, error) {
	out := new(Trip)
	err := c.cc.Invoke(ctx, "/rental.v1.TripService/UpdateTrip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TripServiceServer is the server API for TripService service.
type TripServiceServer interface {
	CreateTrip(context.Context, *CreateTripRequest) (*TripEntity, error)
	GetTrip(context.Context, *GetTripRequest) (*Trip, error)
	GetTrips(context.Context, *GetTripsRequest) (*GetTripsResponse, error)
	UpdateTrip(context.Context, *UpdateTripRequest) (*Trip, error)
}

// UnimplementedTripServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTripServiceServer struct {
}

func (*UnimplementedTripServiceServer) CreateTrip(context.Context, *CreateTripRequest) (*TripEntity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTrip not implemented")
}
func (*UnimplementedTripServiceServer) GetTrip(context.Context, *GetTripRequest) (*Trip, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrip not implemented")
}
func (*UnimplementedTripServiceServer) GetTrips(context.Context, *GetTripsRequest) (*GetTripsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrips not implemented")
}
func (*UnimplementedTripServiceServer) UpdateTrip(context.Context, *UpdateTripRequest) (*Trip, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTrip not implemented")
}

func RegisterTripServiceServer(s *grpc.Server, srv TripServiceServer) {
	s.RegisterService(&_TripService_serviceDesc, srv)
}

func _TripService_CreateTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripServiceServer).CreateTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rental.v1.TripService/CreateTrip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripServiceServer).CreateTrip(ctx, req.(*CreateTripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TripService_GetTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripServiceServer).GetTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rental.v1.TripService/GetTrip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripServiceServer).GetTrip(ctx, req.(*GetTripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TripService_GetTrips_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTripsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripServiceServer).GetTrips(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rental.v1.TripService/GetTrips",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripServiceServer).GetTrips(ctx, req.(*GetTripsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TripService_UpdateTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripServiceServer).UpdateTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rental.v1.TripService/UpdateTrip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripServiceServer).UpdateTrip(ctx, req.(*UpdateTripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TripService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rental.v1.TripService",
	HandlerType: (*TripServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTrip",
			Handler:    _TripService_CreateTrip_Handler,
		},
		{
			MethodName: "GetTrip",
			Handler:    _TripService_GetTrip_Handler,
		},
		{
			MethodName: "GetTrips",
			Handler:    _TripService_GetTrips_Handler,
		},
		{
			MethodName: "UpdateTrip",
			Handler:    _TripService_UpdateTrip_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rental/rental.proto",
}
