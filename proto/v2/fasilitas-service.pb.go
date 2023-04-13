// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.20.1
// source: proto/v2/fasilitas-service.proto

package v2

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetMasterRegionByParamReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionId    int32  `protobuf:"varint,1,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	Propinsi    string `protobuf:"bytes,2,opt,name=propinsi,proto3" json:"propinsi,omitempty"`
	KabKota     string `protobuf:"bytes,3,opt,name=kab_kota,json=kabKota,proto3" json:"kab_kota,omitempty"`
	Kecamatan   string `protobuf:"bytes,4,opt,name=kecamatan,proto3" json:"kecamatan,omitempty"`
	Kelurahan   string `protobuf:"bytes,5,opt,name=kelurahan,proto3" json:"kelurahan,omitempty"`
	FlagIbukota string `protobuf:"bytes,6,opt,name=flag_ibukota,json=flagIbukota,proto3" json:"flag_ibukota,omitempty"`
	Keterangan  string `protobuf:"bytes,7,opt,name=keterangan,proto3" json:"keterangan,omitempty"`
	ParentId    int32  `protobuf:"varint,8,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Nama        string `protobuf:"bytes,9,opt,name=nama,proto3" json:"nama,omitempty"`
	Level       string `protobuf:"bytes,10,opt,name=level,proto3" json:"level,omitempty"`
	ID          int32  `protobuf:"varint,11,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetMasterRegionByParamReq) Reset() {
	*x = GetMasterRegionByParamReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v2_fasilitas_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMasterRegionByParamReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMasterRegionByParamReq) ProtoMessage() {}

func (x *GetMasterRegionByParamReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v2_fasilitas_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMasterRegionByParamReq.ProtoReflect.Descriptor instead.
func (*GetMasterRegionByParamReq) Descriptor() ([]byte, []int) {
	return file_proto_v2_fasilitas_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetMasterRegionByParamReq) GetRegionId() int32 {
	if x != nil {
		return x.RegionId
	}
	return 0
}

func (x *GetMasterRegionByParamReq) GetPropinsi() string {
	if x != nil {
		return x.Propinsi
	}
	return ""
}

func (x *GetMasterRegionByParamReq) GetKabKota() string {
	if x != nil {
		return x.KabKota
	}
	return ""
}

func (x *GetMasterRegionByParamReq) GetKecamatan() string {
	if x != nil {
		return x.Kecamatan
	}
	return ""
}

func (x *GetMasterRegionByParamReq) GetKelurahan() string {
	if x != nil {
		return x.Kelurahan
	}
	return ""
}

func (x *GetMasterRegionByParamReq) GetFlagIbukota() string {
	if x != nil {
		return x.FlagIbukota
	}
	return ""
}

func (x *GetMasterRegionByParamReq) GetKeterangan() string {
	if x != nil {
		return x.Keterangan
	}
	return ""
}

func (x *GetMasterRegionByParamReq) GetParentId() int32 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *GetMasterRegionByParamReq) GetNama() string {
	if x != nil {
		return x.Nama
	}
	return ""
}

func (x *GetMasterRegionByParamReq) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *GetMasterRegionByParamReq) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type MasterRegion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionId    int32  `protobuf:"varint,1,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	Propinsi    string `protobuf:"bytes,2,opt,name=propinsi,proto3" json:"propinsi,omitempty"`
	KabKota     string `protobuf:"bytes,3,opt,name=kab_kota,json=kabKota,proto3" json:"kab_kota,omitempty"`
	Kecamatan   string `protobuf:"bytes,4,opt,name=kecamatan,proto3" json:"kecamatan,omitempty"`
	Kelurahan   string `protobuf:"bytes,5,opt,name=kelurahan,proto3" json:"kelurahan,omitempty"`
	FlagIbukota string `protobuf:"bytes,6,opt,name=flag_ibukota,json=flagIbukota,proto3" json:"flag_ibukota,omitempty"`
	Keterangan  string `protobuf:"bytes,7,opt,name=keterangan,proto3" json:"keterangan,omitempty"`
	ParentId    int32  `protobuf:"varint,8,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Nama        string `protobuf:"bytes,9,opt,name=nama,proto3" json:"nama,omitempty"`
	Level       string `protobuf:"bytes,10,opt,name=level,proto3" json:"level,omitempty"`
	ID          int32  `protobuf:"varint,11,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *MasterRegion) Reset() {
	*x = MasterRegion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v2_fasilitas_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MasterRegion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MasterRegion) ProtoMessage() {}

func (x *MasterRegion) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v2_fasilitas_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MasterRegion.ProtoReflect.Descriptor instead.
func (*MasterRegion) Descriptor() ([]byte, []int) {
	return file_proto_v2_fasilitas_service_proto_rawDescGZIP(), []int{1}
}

func (x *MasterRegion) GetRegionId() int32 {
	if x != nil {
		return x.RegionId
	}
	return 0
}

func (x *MasterRegion) GetPropinsi() string {
	if x != nil {
		return x.Propinsi
	}
	return ""
}

func (x *MasterRegion) GetKabKota() string {
	if x != nil {
		return x.KabKota
	}
	return ""
}

func (x *MasterRegion) GetKecamatan() string {
	if x != nil {
		return x.Kecamatan
	}
	return ""
}

func (x *MasterRegion) GetKelurahan() string {
	if x != nil {
		return x.Kelurahan
	}
	return ""
}

func (x *MasterRegion) GetFlagIbukota() string {
	if x != nil {
		return x.FlagIbukota
	}
	return ""
}

func (x *MasterRegion) GetKeterangan() string {
	if x != nil {
		return x.Keterangan
	}
	return ""
}

func (x *MasterRegion) GetParentId() int32 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *MasterRegion) GetNama() string {
	if x != nil {
		return x.Nama
	}
	return ""
}

func (x *MasterRegion) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *MasterRegion) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

var File_proto_v2_fasilitas_service_proto protoreflect.FileDescriptor

var file_proto_v2_fasilitas_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x61, 0x73, 0x69, 0x6c,
	0x69, 0x74, 0x61, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x21, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x66, 0x61, 0x73, 0x69,
	0x6c, 0x69, 0x74, 0x61, 0x73, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x76, 0x32, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x32, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc5, 0x02, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x12,
	0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x70, 0x69, 0x6e, 0x73, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x70, 0x69, 0x6e, 0x73, 0x69, 0x12, 0x19, 0x0a, 0x08, 0x6b, 0x61, 0x62, 0x5f,
	0x6b, 0x6f, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x61, 0x62, 0x4b,
	0x6f, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x6b, 0x65, 0x63, 0x61, 0x6d, 0x61, 0x74, 0x61, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6b, 0x65, 0x63, 0x61, 0x6d, 0x61, 0x74, 0x61,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6b, 0x65, 0x6c, 0x75, 0x72, 0x61, 0x68, 0x61, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6b, 0x65, 0x6c, 0x75, 0x72, 0x61, 0x68, 0x61, 0x6e, 0x12,
	0x21, 0x0a, 0x0c, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x69, 0x62, 0x75, 0x6b, 0x6f, 0x74, 0x61, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x6c, 0x61, 0x67, 0x49, 0x62, 0x75, 0x6b, 0x6f,
	0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x6b, 0x65, 0x74, 0x65, 0x72, 0x61, 0x6e, 0x67, 0x61, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6b, 0x65, 0x74, 0x65, 0x72, 0x61, 0x6e, 0x67,
	0x61, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x22, 0xb8, 0x02, 0x0a, 0x0c, 0x4d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x69,
	0x6e, 0x73, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x69,
	0x6e, 0x73, 0x69, 0x12, 0x19, 0x0a, 0x08, 0x6b, 0x61, 0x62, 0x5f, 0x6b, 0x6f, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x61, 0x62, 0x4b, 0x6f, 0x74, 0x61, 0x12, 0x1c,
	0x0a, 0x09, 0x6b, 0x65, 0x63, 0x61, 0x6d, 0x61, 0x74, 0x61, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6b, 0x65, 0x63, 0x61, 0x6d, 0x61, 0x74, 0x61, 0x6e, 0x12, 0x1c, 0x0a, 0x09,
	0x6b, 0x65, 0x6c, 0x75, 0x72, 0x61, 0x68, 0x61, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6b, 0x65, 0x6c, 0x75, 0x72, 0x61, 0x68, 0x61, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6c,
	0x61, 0x67, 0x5f, 0x69, 0x62, 0x75, 0x6b, 0x6f, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x66, 0x6c, 0x61, 0x67, 0x49, 0x62, 0x75, 0x6b, 0x6f, 0x74, 0x61, 0x12, 0x1e, 0x0a,
	0x0a, 0x6b, 0x65, 0x74, 0x65, 0x72, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6b, 0x65, 0x74, 0x65, 0x72, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x61, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x49, 0x44, 0x32, 0xdc, 0x02, 0x0a, 0x11, 0x46, 0x61, 0x73, 0x69, 0x6c, 0x69, 0x74,
	0x61, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9e, 0x01, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x42, 0x79,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x3c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x66, 0x61, 0x73, 0x69, 0x6c, 0x69, 0x74, 0x61, 0x73, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x52, 0x65, 0x71, 0x1a, 0x2f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x66, 0x61,
	0x73, 0x69, 0x6c, 0x69, 0x74, 0x61, 0x73, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x32, 0x2e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x6d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0xa5, 0x01, 0x0a, 0x18,
	0x47, 0x65, 0x74, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x42,
	0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x3c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x66, 0x61, 0x73, 0x69, 0x6c, 0x69, 0x74, 0x61, 0x73, 0x2e, 0x6d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74,
	0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x2f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x66, 0x61, 0x73, 0x69, 0x6c, 0x69, 0x74, 0x61, 0x73, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x32, 0x2e, 0x4d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12,
	0x12, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2f, 0x7b,
	0x49, 0x44, 0x7d, 0x42, 0x1c, 0x5a, 0x1a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x66,
	0x61, 0x73, 0x69, 0x6c, 0x69, 0x74, 0x61, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76,
	0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_v2_fasilitas_service_proto_rawDescOnce sync.Once
	file_proto_v2_fasilitas_service_proto_rawDescData = file_proto_v2_fasilitas_service_proto_rawDesc
)

func file_proto_v2_fasilitas_service_proto_rawDescGZIP() []byte {
	file_proto_v2_fasilitas_service_proto_rawDescOnce.Do(func() {
		file_proto_v2_fasilitas_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v2_fasilitas_service_proto_rawDescData)
	})
	return file_proto_v2_fasilitas_service_proto_rawDescData
}

var file_proto_v2_fasilitas_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_v2_fasilitas_service_proto_goTypes = []interface{}{
	(*GetMasterRegionByParamReq)(nil), // 0: service.fasilitas.master.proto.v2.GetMasterRegionByParamReq
	(*MasterRegion)(nil),              // 1: service.fasilitas.master.proto.v2.MasterRegion
}
var file_proto_v2_fasilitas_service_proto_depIdxs = []int32{
	0, // 0: service.fasilitas.master.proto.v2.Fasilitas_service.GetMasterRegionByParam:input_type -> service.fasilitas.master.proto.v2.GetMasterRegionByParamReq
	0, // 1: service.fasilitas.master.proto.v2.Fasilitas_service.GetMasterRegionByParamId:input_type -> service.fasilitas.master.proto.v2.GetMasterRegionByParamReq
	1, // 2: service.fasilitas.master.proto.v2.Fasilitas_service.GetMasterRegionByParam:output_type -> service.fasilitas.master.proto.v2.MasterRegion
	1, // 3: service.fasilitas.master.proto.v2.Fasilitas_service.GetMasterRegionByParamId:output_type -> service.fasilitas.master.proto.v2.MasterRegion
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_v2_fasilitas_service_proto_init() }
func file_proto_v2_fasilitas_service_proto_init() {
	if File_proto_v2_fasilitas_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_v2_fasilitas_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMasterRegionByParamReq); i {
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
		file_proto_v2_fasilitas_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MasterRegion); i {
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
			RawDescriptor: file_proto_v2_fasilitas_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_v2_fasilitas_service_proto_goTypes,
		DependencyIndexes: file_proto_v2_fasilitas_service_proto_depIdxs,
		MessageInfos:      file_proto_v2_fasilitas_service_proto_msgTypes,
	}.Build()
	File_proto_v2_fasilitas_service_proto = out.File
	file_proto_v2_fasilitas_service_proto_rawDesc = nil
	file_proto_v2_fasilitas_service_proto_goTypes = nil
	file_proto_v2_fasilitas_service_proto_depIdxs = nil
}