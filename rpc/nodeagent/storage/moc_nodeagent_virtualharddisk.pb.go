// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_nodeagent_virtualharddisk.proto

package storage

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "github.com/microsoft/moc/rpc/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type VirtualHardDiskType int32

const (
	VirtualHardDiskType_OS_VIRTUALHARDDISK       VirtualHardDiskType = 0
	VirtualHardDiskType_DATADISK_VIRTUALHARDDISK VirtualHardDiskType = 1
)

var VirtualHardDiskType_name = map[int32]string{
	0: "OS_VIRTUALHARDDISK",
	1: "DATADISK_VIRTUALHARDDISK",
}

var VirtualHardDiskType_value = map[string]int32{
	"OS_VIRTUALHARDDISK":       0,
	"DATADISK_VIRTUALHARDDISK": 1,
}

func (x VirtualHardDiskType) String() string {
	return proto.EnumName(VirtualHardDiskType_name, int32(x))
}

func (VirtualHardDiskType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c1f33a566472b7b7, []int{0}
}

type VirtualHardDiskRequest struct {
	VirtualHardDiskSystems []*VirtualHardDisk `protobuf:"bytes,1,rep,name=VirtualHardDiskSystems,proto3" json:"VirtualHardDiskSystems,omitempty"`
	OperationType          common.Operation   `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}           `json:"-"`
	XXX_unrecognized       []byte             `json:"-"`
	XXX_sizecache          int32              `json:"-"`
}

func (m *VirtualHardDiskRequest) Reset()         { *m = VirtualHardDiskRequest{} }
func (m *VirtualHardDiskRequest) String() string { return proto.CompactTextString(m) }
func (*VirtualHardDiskRequest) ProtoMessage()    {}
func (*VirtualHardDiskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1f33a566472b7b7, []int{0}
}

func (m *VirtualHardDiskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualHardDiskRequest.Unmarshal(m, b)
}
func (m *VirtualHardDiskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualHardDiskRequest.Marshal(b, m, deterministic)
}
func (m *VirtualHardDiskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualHardDiskRequest.Merge(m, src)
}
func (m *VirtualHardDiskRequest) XXX_Size() int {
	return xxx_messageInfo_VirtualHardDiskRequest.Size(m)
}
func (m *VirtualHardDiskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualHardDiskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualHardDiskRequest proto.InternalMessageInfo

func (m *VirtualHardDiskRequest) GetVirtualHardDiskSystems() []*VirtualHardDisk {
	if m != nil {
		return m.VirtualHardDiskSystems
	}
	return nil
}

func (m *VirtualHardDiskRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type VirtualHardDiskResponse struct {
	VirtualHardDiskSystems []*VirtualHardDisk  `protobuf:"bytes,1,rep,name=VirtualHardDiskSystems,proto3" json:"VirtualHardDiskSystems,omitempty"`
	Result                 *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                  string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}            `json:"-"`
	XXX_unrecognized       []byte              `json:"-"`
	XXX_sizecache          int32               `json:"-"`
}

func (m *VirtualHardDiskResponse) Reset()         { *m = VirtualHardDiskResponse{} }
func (m *VirtualHardDiskResponse) String() string { return proto.CompactTextString(m) }
func (*VirtualHardDiskResponse) ProtoMessage()    {}
func (*VirtualHardDiskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1f33a566472b7b7, []int{1}
}

func (m *VirtualHardDiskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualHardDiskResponse.Unmarshal(m, b)
}
func (m *VirtualHardDiskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualHardDiskResponse.Marshal(b, m, deterministic)
}
func (m *VirtualHardDiskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualHardDiskResponse.Merge(m, src)
}
func (m *VirtualHardDiskResponse) XXX_Size() int {
	return xxx_messageInfo_VirtualHardDiskResponse.Size(m)
}
func (m *VirtualHardDiskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualHardDiskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualHardDiskResponse proto.InternalMessageInfo

func (m *VirtualHardDiskResponse) GetVirtualHardDiskSystems() []*VirtualHardDisk {
	if m != nil {
		return m.VirtualHardDiskSystems
	}
	return nil
}

func (m *VirtualHardDiskResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *VirtualHardDiskResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type SFSImageProperties struct {
	CatalogName          string   `protobuf:"bytes,1,opt,name=catalogName,proto3" json:"catalogName,omitempty"`
	Audience             string   `protobuf:"bytes,2,opt,name=audience,proto3" json:"audience,omitempty"`
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	ReleaseName          string   `protobuf:"bytes,4,opt,name=releaseName,proto3" json:"releaseName,omitempty"`
	Parts                uint32   `protobuf:"varint,5,opt,name=parts,proto3" json:"parts,omitempty"`
	DestinationDir       string   `protobuf:"bytes,6,opt,name=destinationDir,proto3" json:"destinationDir,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SFSImageProperties) Reset()         { *m = SFSImageProperties{} }
func (m *SFSImageProperties) String() string { return proto.CompactTextString(m) }
func (*SFSImageProperties) ProtoMessage()    {}
func (*SFSImageProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1f33a566472b7b7, []int{2}
}

func (m *SFSImageProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SFSImageProperties.Unmarshal(m, b)
}
func (m *SFSImageProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SFSImageProperties.Marshal(b, m, deterministic)
}
func (m *SFSImageProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SFSImageProperties.Merge(m, src)
}
func (m *SFSImageProperties) XXX_Size() int {
	return xxx_messageInfo_SFSImageProperties.Size(m)
}
func (m *SFSImageProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_SFSImageProperties.DiscardUnknown(m)
}

var xxx_messageInfo_SFSImageProperties proto.InternalMessageInfo

func (m *SFSImageProperties) GetCatalogName() string {
	if m != nil {
		return m.CatalogName
	}
	return ""
}

func (m *SFSImageProperties) GetAudience() string {
	if m != nil {
		return m.Audience
	}
	return ""
}

func (m *SFSImageProperties) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *SFSImageProperties) GetReleaseName() string {
	if m != nil {
		return m.ReleaseName
	}
	return ""
}

func (m *SFSImageProperties) GetParts() uint32 {
	if m != nil {
		return m.Parts
	}
	return 0
}

func (m *SFSImageProperties) GetDestinationDir() string {
	if m != nil {
		return m.DestinationDir
	}
	return ""
}

type LocalImageProperties struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocalImageProperties) Reset()         { *m = LocalImageProperties{} }
func (m *LocalImageProperties) String() string { return proto.CompactTextString(m) }
func (*LocalImageProperties) ProtoMessage()    {}
func (*LocalImageProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1f33a566472b7b7, []int{3}
}

func (m *LocalImageProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalImageProperties.Unmarshal(m, b)
}
func (m *LocalImageProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalImageProperties.Marshal(b, m, deterministic)
}
func (m *LocalImageProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalImageProperties.Merge(m, src)
}
func (m *LocalImageProperties) XXX_Size() int {
	return xxx_messageInfo_LocalImageProperties.Size(m)
}
func (m *LocalImageProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalImageProperties.DiscardUnknown(m)
}

var xxx_messageInfo_LocalImageProperties proto.InternalMessageInfo

func (m *LocalImageProperties) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type CloneImageProperties struct {
	CloneSource          string   `protobuf:"bytes,1,opt,name=cloneSource,proto3" json:"cloneSource,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloneImageProperties) Reset()         { *m = CloneImageProperties{} }
func (m *CloneImageProperties) String() string { return proto.CompactTextString(m) }
func (*CloneImageProperties) ProtoMessage()    {}
func (*CloneImageProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1f33a566472b7b7, []int{4}
}

func (m *CloneImageProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloneImageProperties.Unmarshal(m, b)
}
func (m *CloneImageProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloneImageProperties.Marshal(b, m, deterministic)
}
func (m *CloneImageProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloneImageProperties.Merge(m, src)
}
func (m *CloneImageProperties) XXX_Size() int {
	return xxx_messageInfo_CloneImageProperties.Size(m)
}
func (m *CloneImageProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_CloneImageProperties.DiscardUnknown(m)
}

var xxx_messageInfo_CloneImageProperties proto.InternalMessageInfo

func (m *CloneImageProperties) GetCloneSource() string {
	if m != nil {
		return m.CloneSource
	}
	return ""
}

type HttpImageProperties struct {
	HttpURL              string   `protobuf:"bytes,1,opt,name=httpURL,proto3" json:"httpURL,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HttpImageProperties) Reset()         { *m = HttpImageProperties{} }
func (m *HttpImageProperties) String() string { return proto.CompactTextString(m) }
func (*HttpImageProperties) ProtoMessage()    {}
func (*HttpImageProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1f33a566472b7b7, []int{5}
}

func (m *HttpImageProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpImageProperties.Unmarshal(m, b)
}
func (m *HttpImageProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpImageProperties.Marshal(b, m, deterministic)
}
func (m *HttpImageProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpImageProperties.Merge(m, src)
}
func (m *HttpImageProperties) XXX_Size() int {
	return xxx_messageInfo_HttpImageProperties.Size(m)
}
func (m *HttpImageProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpImageProperties.DiscardUnknown(m)
}

var xxx_messageInfo_HttpImageProperties proto.InternalMessageInfo

func (m *HttpImageProperties) GetHttpURL() string {
	if m != nil {
		return m.HttpURL
	}
	return ""
}

type AzureGalleryImageProperties struct {
	SasURI               string   `protobuf:"bytes,1,opt,name=sasURI,proto3" json:"sasURI,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AzureGalleryImageProperties) Reset()         { *m = AzureGalleryImageProperties{} }
func (m *AzureGalleryImageProperties) String() string { return proto.CompactTextString(m) }
func (*AzureGalleryImageProperties) ProtoMessage()    {}
func (*AzureGalleryImageProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1f33a566472b7b7, []int{6}
}

func (m *AzureGalleryImageProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AzureGalleryImageProperties.Unmarshal(m, b)
}
func (m *AzureGalleryImageProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AzureGalleryImageProperties.Marshal(b, m, deterministic)
}
func (m *AzureGalleryImageProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AzureGalleryImageProperties.Merge(m, src)
}
func (m *AzureGalleryImageProperties) XXX_Size() int {
	return xxx_messageInfo_AzureGalleryImageProperties.Size(m)
}
func (m *AzureGalleryImageProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_AzureGalleryImageProperties.DiscardUnknown(m)
}

var xxx_messageInfo_AzureGalleryImageProperties proto.InternalMessageInfo

func (m *AzureGalleryImageProperties) GetSasURI() string {
	if m != nil {
		return m.SasURI
	}
	return ""
}

func (m *AzureGalleryImageProperties) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type VirtualHardDisk struct {
	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id     string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Source string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Path   string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	// Storage container name to hold this vhd
	ContainerName        string                     `protobuf:"bytes,5,opt,name=containerName,proto3" json:"containerName,omitempty"`
	Status               *common.Status             `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	Size                 int64                      `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty"`
	Dynamic              bool                       `protobuf:"varint,8,opt,name=dynamic,proto3" json:"dynamic,omitempty"`
	Blocksizebytes       int32                      `protobuf:"varint,9,opt,name=blocksizebytes,proto3" json:"blocksizebytes,omitempty"`
	Logicalsectorbytes   int32                      `protobuf:"varint,10,opt,name=logicalsectorbytes,proto3" json:"logicalsectorbytes,omitempty"`
	Physicalsectorbytes  int32                      `protobuf:"varint,11,opt,name=physicalsectorbytes,proto3" json:"physicalsectorbytes,omitempty"`
	Controllernumber     int64                      `protobuf:"varint,12,opt,name=controllernumber,proto3" json:"controllernumber,omitempty"`
	Controllerlocation   int64                      `protobuf:"varint,13,opt,name=controllerlocation,proto3" json:"controllerlocation,omitempty"`
	Disknumber           int64                      `protobuf:"varint,14,opt,name=disknumber,proto3" json:"disknumber,omitempty"`
	VirtualmachineName   string                     `protobuf:"bytes,15,opt,name=virtualmachineName,proto3" json:"virtualmachineName,omitempty"`
	Scsipath             string                     `protobuf:"bytes,16,opt,name=scsipath,proto3" json:"scsipath,omitempty"`
	Virtualharddisktype  VirtualHardDiskType        `protobuf:"varint,17,opt,name=virtualharddisktype,proto3,enum=moc.nodeagent.storage.VirtualHardDiskType" json:"virtualharddisktype,omitempty"`
	Entity               *common.Entity             `protobuf:"bytes,18,opt,name=entity,proto3" json:"entity,omitempty"`
	Tags                 *common.Tags               `protobuf:"bytes,19,opt,name=tags,proto3" json:"tags,omitempty"`
	SourceType           common.ImageSource         `protobuf:"varint,20,opt,name=sourceType,proto3,enum=moc.ImageSource" json:"sourceType,omitempty"`
	HyperVGeneration     common.HyperVGeneration    `protobuf:"varint,21,opt,name=hyperVGeneration,proto3,enum=moc.HyperVGeneration" json:"hyperVGeneration,omitempty"`
	CloudInitDataSource  common.CloudInitDataSource `protobuf:"varint,22,opt,name=cloudInitDataSource,proto3,enum=moc.CloudInitDataSource" json:"cloudInitDataSource,omitempty"`
	DiskFileFormat       common.DiskFileFormat      `protobuf:"varint,23,opt,name=diskFileFormat,proto3,enum=moc.DiskFileFormat" json:"diskFileFormat,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *VirtualHardDisk) Reset()         { *m = VirtualHardDisk{} }
func (m *VirtualHardDisk) String() string { return proto.CompactTextString(m) }
func (*VirtualHardDisk) ProtoMessage()    {}
func (*VirtualHardDisk) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1f33a566472b7b7, []int{7}
}

func (m *VirtualHardDisk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualHardDisk.Unmarshal(m, b)
}
func (m *VirtualHardDisk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualHardDisk.Marshal(b, m, deterministic)
}
func (m *VirtualHardDisk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualHardDisk.Merge(m, src)
}
func (m *VirtualHardDisk) XXX_Size() int {
	return xxx_messageInfo_VirtualHardDisk.Size(m)
}
func (m *VirtualHardDisk) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualHardDisk.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualHardDisk proto.InternalMessageInfo

func (m *VirtualHardDisk) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VirtualHardDisk) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VirtualHardDisk) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *VirtualHardDisk) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *VirtualHardDisk) GetContainerName() string {
	if m != nil {
		return m.ContainerName
	}
	return ""
}

func (m *VirtualHardDisk) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *VirtualHardDisk) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *VirtualHardDisk) GetDynamic() bool {
	if m != nil {
		return m.Dynamic
	}
	return false
}

func (m *VirtualHardDisk) GetBlocksizebytes() int32 {
	if m != nil {
		return m.Blocksizebytes
	}
	return 0
}

func (m *VirtualHardDisk) GetLogicalsectorbytes() int32 {
	if m != nil {
		return m.Logicalsectorbytes
	}
	return 0
}

func (m *VirtualHardDisk) GetPhysicalsectorbytes() int32 {
	if m != nil {
		return m.Physicalsectorbytes
	}
	return 0
}

func (m *VirtualHardDisk) GetControllernumber() int64 {
	if m != nil {
		return m.Controllernumber
	}
	return 0
}

func (m *VirtualHardDisk) GetControllerlocation() int64 {
	if m != nil {
		return m.Controllerlocation
	}
	return 0
}

func (m *VirtualHardDisk) GetDisknumber() int64 {
	if m != nil {
		return m.Disknumber
	}
	return 0
}

func (m *VirtualHardDisk) GetVirtualmachineName() string {
	if m != nil {
		return m.VirtualmachineName
	}
	return ""
}

func (m *VirtualHardDisk) GetScsipath() string {
	if m != nil {
		return m.Scsipath
	}
	return ""
}

func (m *VirtualHardDisk) GetVirtualharddisktype() VirtualHardDiskType {
	if m != nil {
		return m.Virtualharddisktype
	}
	return VirtualHardDiskType_OS_VIRTUALHARDDISK
}

func (m *VirtualHardDisk) GetEntity() *common.Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *VirtualHardDisk) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *VirtualHardDisk) GetSourceType() common.ImageSource {
	if m != nil {
		return m.SourceType
	}
	return common.ImageSource_LOCAL_SOURCE
}

func (m *VirtualHardDisk) GetHyperVGeneration() common.HyperVGeneration {
	if m != nil {
		return m.HyperVGeneration
	}
	return common.HyperVGeneration_HyperVGenerationV1
}

func (m *VirtualHardDisk) GetCloudInitDataSource() common.CloudInitDataSource {
	if m != nil {
		return m.CloudInitDataSource
	}
	return common.CloudInitDataSource_NoCloud
}

func (m *VirtualHardDisk) GetDiskFileFormat() common.DiskFileFormat {
	if m != nil {
		return m.DiskFileFormat
	}
	return common.DiskFileFormat_DiskFileFormatVHD
}

func init() {
	proto.RegisterEnum("moc.nodeagent.storage.VirtualHardDiskType", VirtualHardDiskType_name, VirtualHardDiskType_value)
	proto.RegisterType((*VirtualHardDiskRequest)(nil), "moc.nodeagent.storage.VirtualHardDiskRequest")
	proto.RegisterType((*VirtualHardDiskResponse)(nil), "moc.nodeagent.storage.VirtualHardDiskResponse")
	proto.RegisterType((*SFSImageProperties)(nil), "moc.nodeagent.storage.SFSImageProperties")
	proto.RegisterType((*LocalImageProperties)(nil), "moc.nodeagent.storage.LocalImageProperties")
	proto.RegisterType((*CloneImageProperties)(nil), "moc.nodeagent.storage.CloneImageProperties")
	proto.RegisterType((*HttpImageProperties)(nil), "moc.nodeagent.storage.HttpImageProperties")
	proto.RegisterType((*AzureGalleryImageProperties)(nil), "moc.nodeagent.storage.AzureGalleryImageProperties")
	proto.RegisterType((*VirtualHardDisk)(nil), "moc.nodeagent.storage.VirtualHardDisk")
}

func init() {
	proto.RegisterFile("moc_nodeagent_virtualharddisk.proto", fileDescriptor_c1f33a566472b7b7)
}

var fileDescriptor_c1f33a566472b7b7 = []byte{
	// 984 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0x5b, 0x6f, 0xdb, 0x36,
	0x14, 0xae, 0x72, 0x71, 0x9a, 0xe3, 0xc5, 0x4d, 0xe9, 0x5c, 0x34, 0x77, 0x2d, 0x0c, 0x77, 0x28,
	0x8c, 0x00, 0x93, 0x03, 0x6f, 0x0f, 0x03, 0xf6, 0xe4, 0xc6, 0x49, 0xe3, 0x35, 0x68, 0x06, 0x3a,
	0xc9, 0xc3, 0x30, 0x2c, 0xa0, 0x69, 0xc6, 0x26, 0x22, 0x89, 0x1a, 0x49, 0x65, 0x70, 0x7f, 0xd4,
	0xfe, 0xc4, 0x1e, 0xd7, 0x1f, 0x35, 0xf0, 0x48, 0xbe, 0x44, 0xf6, 0x80, 0xbc, 0xec, 0xc9, 0x3a,
	0xe7, 0xfb, 0xce, 0xa7, 0x73, 0x11, 0x0f, 0x0d, 0x6f, 0x23, 0xc5, 0x6f, 0x63, 0x35, 0x14, 0x6c,
	0x24, 0x62, 0x7b, 0xfb, 0x20, 0xb5, 0x4d, 0x59, 0x38, 0x66, 0x7a, 0x38, 0x94, 0xe6, 0x3e, 0x48,
	0xb4, 0xb2, 0x8a, 0xec, 0x47, 0x8a, 0x07, 0x33, 0x52, 0x60, 0xac, 0xd2, 0x6c, 0x24, 0x6a, 0x6f,
	0x46, 0x4a, 0x8d, 0x42, 0xd1, 0x42, 0xd2, 0x20, 0xbd, 0x6b, 0xfd, 0xa9, 0x59, 0x92, 0x08, 0x6d,
	0xb2, 0xb0, 0xda, 0xa1, 0xd3, 0xe6, 0x2a, 0x8a, 0x54, 0x9c, 0xff, 0xe4, 0xc0, 0xeb, 0x05, 0x20,
	0x56, 0x56, 0xde, 0x49, 0xce, 0xac, 0x9c, 0xc1, 0xaf, 0x8a, 0xba, 0x22, 0x4a, 0xec, 0x24, 0x03,
	0x1b, 0x7f, 0x79, 0x70, 0x70, 0x93, 0x65, 0x79, 0xce, 0xf4, 0xb0, 0x2b, 0xcd, 0x3d, 0x15, 0x7f,
	0xa4, 0xc2, 0x58, 0xf2, 0xfb, 0x12, 0xd2, 0x9f, 0x18, 0x2b, 0x22, 0xe3, 0x7b, 0xf5, 0xf5, 0x66,
	0xb9, 0xfd, 0x2e, 0x58, 0x59, 0x47, 0x50, 0x94, 0xfb, 0x0f, 0x15, 0xf2, 0x03, 0xec, 0x5c, 0x26,
	0x42, 0x63, 0xaa, 0x57, 0x93, 0x44, 0xf8, 0x6b, 0x75, 0xaf, 0x59, 0x69, 0x57, 0x50, 0x76, 0x86,
	0xd0, 0xc7, 0xa4, 0xc6, 0xdf, 0x1e, 0x1c, 0x2e, 0x25, 0x6c, 0x12, 0x15, 0x1b, 0xf1, 0xbf, 0x67,
	0xdc, 0x86, 0x12, 0x15, 0x26, 0x0d, 0x2d, 0xa6, 0x5a, 0x6e, 0xd7, 0x82, 0xac, 0xb5, 0xc1, 0xb4,
	0xb5, 0xc1, 0x7b, 0xa5, 0xc2, 0x1b, 0x16, 0xa6, 0x82, 0xe6, 0x4c, 0xb2, 0x07, 0x9b, 0xa7, 0x5a,
	0x2b, 0xed, 0xaf, 0xd7, 0xbd, 0xe6, 0x36, 0xcd, 0x8c, 0xc6, 0x17, 0x0f, 0x48, 0xff, 0xac, 0xdf,
	0x8b, 0xd8, 0x48, 0xfc, 0xa2, 0x55, 0x22, 0xb4, 0x95, 0xc2, 0x90, 0x3a, 0x94, 0x39, 0xb3, 0x2c,
	0x54, 0xa3, 0x4f, 0x2c, 0x12, 0xbe, 0x87, 0x21, 0x8b, 0x2e, 0x52, 0x83, 0xe7, 0x2c, 0x1d, 0x4a,
	0x11, 0xf3, 0xac, 0x5f, 0xdb, 0x74, 0x66, 0x13, 0x1f, 0xb6, 0x1e, 0x84, 0x36, 0x52, 0xc5, 0xf9,
	0xcb, 0xa6, 0xa6, 0xd3, 0xd5, 0x22, 0x14, 0xcc, 0x08, 0xd4, 0xdd, 0xc8, 0x74, 0x17, 0x5c, 0x2e,
	0xcd, 0x84, 0x69, 0x6b, 0xfc, 0xcd, 0xba, 0xd7, 0xdc, 0xa1, 0x99, 0x41, 0xde, 0x41, 0x65, 0x28,
	0x8c, 0x95, 0x31, 0xf6, 0xbf, 0x2b, 0xb5, 0x5f, 0xc2, 0xd0, 0x82, 0xb7, 0x71, 0x04, 0x7b, 0x17,
	0x8a, 0xb3, 0xb0, 0x58, 0x0f, 0x81, 0x8d, 0x84, 0xd9, 0x71, 0x5e, 0x08, 0x3e, 0x37, 0x7e, 0x84,
	0xbd, 0x93, 0x50, 0xc5, 0x62, 0x55, 0xed, 0xce, 0xdf, 0x57, 0xa9, 0xe6, 0xf3, 0xda, 0xe7, 0xae,
	0x46, 0x0b, 0xaa, 0xe7, 0xd6, 0x26, 0xc5, 0x40, 0x1f, 0xb6, 0xc6, 0xd6, 0x26, 0xd7, 0xf4, 0x22,
	0x0f, 0x9a, 0x9a, 0x8d, 0x4b, 0x78, 0xd5, 0xf9, 0x9c, 0x6a, 0xf1, 0x81, 0x85, 0xa1, 0xd0, 0x93,
	0x62, 0xe0, 0x01, 0x94, 0x0c, 0x33, 0xd7, 0xb4, 0x97, 0xc7, 0xe5, 0xd6, 0x62, 0x1f, 0xd7, 0x1e,
	0xf5, 0xb1, 0xf1, 0x65, 0x0b, 0x5e, 0x14, 0xbe, 0x0d, 0x57, 0x63, 0x3c, 0x1f, 0x16, 0x3e, 0x93,
	0x0a, 0xac, 0xc9, 0x61, 0x1e, 0xbc, 0x26, 0x87, 0xf8, 0xa6, 0xac, 0xac, 0xf5, 0xfc, 0x4d, 0x68,
	0xcd, 0xfa, 0xb3, 0x31, 0xef, 0x0f, 0xf9, 0x16, 0x76, 0xb8, 0x8a, 0x2d, 0x93, 0xb1, 0xd0, 0x38,
	0xad, 0x4d, 0x04, 0x1f, 0x3b, 0xc9, 0x5b, 0x28, 0x19, 0xcb, 0x6c, 0x6a, 0x70, 0x22, 0xe5, 0x76,
	0x19, 0x3f, 0xed, 0x3e, 0xba, 0x68, 0x0e, 0x39, 0x79, 0x23, 0x3f, 0x0b, 0x7f, 0xab, 0xee, 0x35,
	0xd7, 0x29, 0x3e, 0xbb, 0xe2, 0x86, 0x93, 0x98, 0x45, 0x92, 0xfb, 0xcf, 0xeb, 0x5e, 0xf3, 0x39,
	0x9d, 0x9a, 0x6e, 0xd8, 0x83, 0x50, 0xf1, 0x7b, 0x47, 0x1b, 0x4c, 0xac, 0x30, 0xfe, 0x76, 0xdd,
	0x6b, 0x6e, 0xd2, 0x82, 0x97, 0x04, 0x40, 0x42, 0x35, 0x92, 0x9c, 0x85, 0x46, 0x70, 0xab, 0x74,
	0xc6, 0x05, 0xe4, 0xae, 0x40, 0xc8, 0x31, 0x54, 0x93, 0xf1, 0xc4, 0x14, 0x03, 0xca, 0x18, 0xb0,
	0x0a, 0x22, 0x47, 0xb0, 0xeb, 0xaa, 0xd5, 0xca, 0x8d, 0x2d, 0x4e, 0xa3, 0x81, 0xd0, 0xfe, 0x57,
	0x58, 0xc3, 0x92, 0xdf, 0x65, 0x33, 0xf7, 0x85, 0x2a, 0xdb, 0x7c, 0xfe, 0x0e, 0xb2, 0x57, 0x20,
	0xe4, 0x0d, 0x80, 0x5b, 0xc5, 0xb9, 0x6a, 0x05, 0x79, 0x0b, 0x1e, 0xa7, 0x97, 0x6f, 0xed, 0x88,
	0xf1, 0xb1, 0x8c, 0xb3, 0x13, 0xf3, 0x02, 0x67, 0xb0, 0x02, 0x71, 0x07, 0xd2, 0x70, 0x23, 0x71,
	0x8c, 0xbb, 0xd9, 0x81, 0x9c, 0xda, 0xe4, 0x37, 0xa8, 0x16, 0x6e, 0x00, 0xeb, 0xf6, 0xdc, 0x4b,
	0xdc, 0x73, 0x47, 0x4f, 0x5b, 0x46, 0x6e, 0xe9, 0xd1, 0x55, 0x32, 0xee, 0x13, 0x10, 0xb1, 0x95,
	0x76, 0xe2, 0x93, 0x85, 0x4f, 0xe0, 0x14, 0x5d, 0x34, 0x87, 0xc8, 0x6b, 0xd8, 0xb0, 0x6c, 0x64,
	0xfc, 0x2a, 0x52, 0xb6, 0x91, 0x72, 0xc5, 0x46, 0x86, 0xa2, 0x9b, 0x1c, 0x03, 0x64, 0x9f, 0x22,
	0x2e, 0xe0, 0x3d, 0x4c, 0x6c, 0x17, 0x49, 0x78, 0x58, 0xb2, 0x83, 0x47, 0x17, 0x38, 0xa4, 0x03,
	0xbb, 0xe3, 0x49, 0x22, 0xf4, 0xcd, 0x07, 0x11, 0xe7, 0x7b, 0xd9, 0xdf, 0xc7, 0xb8, 0x7d, 0x8c,
	0x3b, 0x2f, 0x80, 0x74, 0x89, 0x4e, 0x7e, 0x86, 0x2a, 0x0f, 0x55, 0x3a, 0xec, 0xc5, 0xd2, 0x76,
	0x99, 0x65, 0xf9, 0x89, 0x3f, 0x40, 0x15, 0x1f, 0x55, 0x4e, 0x96, 0x71, 0xba, 0x2a, 0x88, 0xfc,
	0x04, 0x15, 0xd7, 0x90, 0x33, 0x19, 0x8a, 0x33, 0xa5, 0x23, 0x66, 0xfd, 0x43, 0x94, 0xa9, 0xa2,
	0x4c, 0xf7, 0x11, 0x44, 0x0b, 0xd4, 0xa3, 0x8f, 0x50, 0x5d, 0xd1, 0x6d, 0x72, 0x00, 0xe4, 0xb2,
	0x7f, 0x7b, 0xd3, 0xa3, 0x57, 0xd7, 0x9d, 0x8b, 0xf3, 0x0e, 0xed, 0x76, 0x7b, 0xfd, 0x8f, 0xbb,
	0xcf, 0xc8, 0x37, 0xe0, 0x77, 0x3b, 0x57, 0x1d, 0x67, 0x2d, 0xa1, 0x5e, 0xfb, 0x1f, 0x0f, 0xf6,
	0x0a, 0x6a, 0x1d, 0x37, 0x58, 0x22, 0xa1, 0xd4, 0x8b, 0x1f, 0xd4, 0xbd, 0x20, 0xdf, 0x3d, 0xf1,
	0xfe, 0xc9, 0x2e, 0xe0, 0x5a, 0xf0, 0x54, 0x7a, 0x76, 0xfd, 0x35, 0x9e, 0x91, 0x73, 0x78, 0x79,
	0x32, 0x16, 0xfc, 0xfe, 0xd3, 0xc2, 0xbf, 0x00, 0x72, 0xb0, 0x74, 0x4b, 0x9d, 0xba, 0x3f, 0x00,
	0xb5, 0xaf, 0x51, 0x7e, 0x91, 0x3a, 0x57, 0x7a, 0x7f, 0xfc, 0x6b, 0x30, 0x92, 0x76, 0x9c, 0x0e,
	0x02, 0xae, 0xa2, 0x56, 0x24, 0xb9, 0x56, 0x46, 0xdd, 0xd9, 0x56, 0xa4, 0x78, 0x4b, 0x27, 0xbc,
	0x35, 0xcb, 0xaa, 0x95, 0x67, 0x35, 0x28, 0xa1, 0xfc, 0xf7, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x6b, 0x3e, 0x1d, 0x68, 0x03, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VirtualHardDiskAgentClient is the client API for VirtualHardDiskAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VirtualHardDiskAgentClient interface {
	Invoke(ctx context.Context, in *VirtualHardDiskRequest, opts ...grpc.CallOption) (*VirtualHardDiskResponse, error)
	CheckNotification(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*common.NotificationResponse, error)
}

type virtualHardDiskAgentClient struct {
	cc *grpc.ClientConn
}

func NewVirtualHardDiskAgentClient(cc *grpc.ClientConn) VirtualHardDiskAgentClient {
	return &virtualHardDiskAgentClient{cc}
}

func (c *virtualHardDiskAgentClient) Invoke(ctx context.Context, in *VirtualHardDiskRequest, opts ...grpc.CallOption) (*VirtualHardDiskResponse, error) {
	out := new(VirtualHardDiskResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.storage.VirtualHardDiskAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualHardDiskAgentClient) CheckNotification(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*common.NotificationResponse, error) {
	out := new(common.NotificationResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.storage.VirtualHardDiskAgent/CheckNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VirtualHardDiskAgentServer is the server API for VirtualHardDiskAgent service.
type VirtualHardDiskAgentServer interface {
	Invoke(context.Context, *VirtualHardDiskRequest) (*VirtualHardDiskResponse, error)
	CheckNotification(context.Context, *empty.Empty) (*common.NotificationResponse, error)
}

// UnimplementedVirtualHardDiskAgentServer can be embedded to have forward compatible implementations.
type UnimplementedVirtualHardDiskAgentServer struct {
}

func (*UnimplementedVirtualHardDiskAgentServer) Invoke(ctx context.Context, req *VirtualHardDiskRequest) (*VirtualHardDiskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}
func (*UnimplementedVirtualHardDiskAgentServer) CheckNotification(ctx context.Context, req *empty.Empty) (*common.NotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckNotification not implemented")
}

func RegisterVirtualHardDiskAgentServer(s *grpc.Server, srv VirtualHardDiskAgentServer) {
	s.RegisterService(&_VirtualHardDiskAgent_serviceDesc, srv)
}

func _VirtualHardDiskAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VirtualHardDiskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualHardDiskAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.nodeagent.storage.VirtualHardDiskAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualHardDiskAgentServer).Invoke(ctx, req.(*VirtualHardDiskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualHardDiskAgent_CheckNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualHardDiskAgentServer).CheckNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.nodeagent.storage.VirtualHardDiskAgent/CheckNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualHardDiskAgentServer).CheckNotification(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _VirtualHardDiskAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.nodeagent.storage.VirtualHardDiskAgent",
	HandlerType: (*VirtualHardDiskAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _VirtualHardDiskAgent_Invoke_Handler,
		},
		{
			MethodName: "CheckNotification",
			Handler:    _VirtualHardDiskAgent_CheckNotification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moc_nodeagent_virtualharddisk.proto",
}
