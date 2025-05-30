// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_cloudagent_virtualharddisk.proto

package storage

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type VirtualHardDiskRequest struct {
	VirtualHardDisks     []*VirtualHardDisk `protobuf:"bytes,1,rep,name=VirtualHardDisks,proto3" json:"VirtualHardDisks,omitempty"`
	OperationType        common.Operation   `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *VirtualHardDiskRequest) Reset()         { *m = VirtualHardDiskRequest{} }
func (m *VirtualHardDiskRequest) String() string { return proto.CompactTextString(m) }
func (*VirtualHardDiskRequest) ProtoMessage()    {}
func (*VirtualHardDiskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3ff93a45373b349, []int{0}
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

func (m *VirtualHardDiskRequest) GetVirtualHardDisks() []*VirtualHardDisk {
	if m != nil {
		return m.VirtualHardDisks
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
	VirtualHardDisks     []*VirtualHardDisk  `protobuf:"bytes,1,rep,name=VirtualHardDisks,proto3" json:"VirtualHardDisks,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *VirtualHardDiskResponse) Reset()         { *m = VirtualHardDiskResponse{} }
func (m *VirtualHardDiskResponse) String() string { return proto.CompactTextString(m) }
func (*VirtualHardDiskResponse) ProtoMessage()    {}
func (*VirtualHardDiskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3ff93a45373b349, []int{1}
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

func (m *VirtualHardDiskResponse) GetVirtualHardDisks() []*VirtualHardDisk {
	if m != nil {
		return m.VirtualHardDisks
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

type VirtualHardDiskOperationRequest struct {
	VirtualHardDisks     []*VirtualHardDisk             `protobuf:"bytes,1,rep,name=VirtualHardDisks,proto3" json:"VirtualHardDisks,omitempty"`
	OperationType        common.ProviderAccessOperation `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.ProviderAccessOperation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *VirtualHardDiskOperationRequest) Reset()         { *m = VirtualHardDiskOperationRequest{} }
func (m *VirtualHardDiskOperationRequest) String() string { return proto.CompactTextString(m) }
func (*VirtualHardDiskOperationRequest) ProtoMessage()    {}
func (*VirtualHardDiskOperationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3ff93a45373b349, []int{2}
}

func (m *VirtualHardDiskOperationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualHardDiskOperationRequest.Unmarshal(m, b)
}
func (m *VirtualHardDiskOperationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualHardDiskOperationRequest.Marshal(b, m, deterministic)
}
func (m *VirtualHardDiskOperationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualHardDiskOperationRequest.Merge(m, src)
}
func (m *VirtualHardDiskOperationRequest) XXX_Size() int {
	return xxx_messageInfo_VirtualHardDiskOperationRequest.Size(m)
}
func (m *VirtualHardDiskOperationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualHardDiskOperationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualHardDiskOperationRequest proto.InternalMessageInfo

func (m *VirtualHardDiskOperationRequest) GetVirtualHardDisks() []*VirtualHardDisk {
	if m != nil {
		return m.VirtualHardDisks
	}
	return nil
}

func (m *VirtualHardDiskOperationRequest) GetOperationType() common.ProviderAccessOperation {
	if m != nil {
		return m.OperationType
	}
	return common.ProviderAccessOperation_Unspecified
}

type VirtualHardDiskOperationResponse struct {
	VirtualHardDisks     []*VirtualHardDisk  `protobuf:"bytes,1,rep,name=VirtualHardDisks,proto3" json:"VirtualHardDisks,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *VirtualHardDiskOperationResponse) Reset()         { *m = VirtualHardDiskOperationResponse{} }
func (m *VirtualHardDiskOperationResponse) String() string { return proto.CompactTextString(m) }
func (*VirtualHardDiskOperationResponse) ProtoMessage()    {}
func (*VirtualHardDiskOperationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3ff93a45373b349, []int{3}
}

func (m *VirtualHardDiskOperationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualHardDiskOperationResponse.Unmarshal(m, b)
}
func (m *VirtualHardDiskOperationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualHardDiskOperationResponse.Marshal(b, m, deterministic)
}
func (m *VirtualHardDiskOperationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualHardDiskOperationResponse.Merge(m, src)
}
func (m *VirtualHardDiskOperationResponse) XXX_Size() int {
	return xxx_messageInfo_VirtualHardDiskOperationResponse.Size(m)
}
func (m *VirtualHardDiskOperationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualHardDiskOperationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualHardDiskOperationResponse proto.InternalMessageInfo

func (m *VirtualHardDiskOperationResponse) GetVirtualHardDisks() []*VirtualHardDisk {
	if m != nil {
		return m.VirtualHardDisks
	}
	return nil
}

func (m *VirtualHardDiskOperationResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *VirtualHardDiskOperationResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type VirtualHardDisk struct {
	Name                       string                     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                         string                     `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	ContainerName              string                     `protobuf:"bytes,3,opt,name=containerName,proto3" json:"containerName,omitempty"`
	Path                       string                     `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Status                     *common.Status             `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Size                       int64                      `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`
	Dynamic                    bool                       `protobuf:"varint,7,opt,name=dynamic,proto3" json:"dynamic,omitempty"`
	Blocksizebytes             int32                      `protobuf:"varint,8,opt,name=blocksizebytes,proto3" json:"blocksizebytes,omitempty"`
	Logicalsectorbytes         int32                      `protobuf:"varint,9,opt,name=logicalsectorbytes,proto3" json:"logicalsectorbytes,omitempty"`
	Physicalsectorbytes        int32                      `protobuf:"varint,10,opt,name=physicalsectorbytes,proto3" json:"physicalsectorbytes,omitempty"`
	Controllernumber           int64                      `protobuf:"varint,11,opt,name=controllernumber,proto3" json:"controllernumber,omitempty"`
	Controllerlocation         int64                      `protobuf:"varint,12,opt,name=controllerlocation,proto3" json:"controllerlocation,omitempty"`
	Disknumber                 int64                      `protobuf:"varint,13,opt,name=disknumber,proto3" json:"disknumber,omitempty"`
	VirtualmachineName         string                     `protobuf:"bytes,14,opt,name=virtualmachineName,proto3" json:"virtualmachineName,omitempty"`
	Scsipath                   string                     `protobuf:"bytes,15,opt,name=scsipath,proto3" json:"scsipath,omitempty"`
	AttachedVirtualmachineName string                     `protobuf:"bytes,16,opt,name=attachedVirtualmachineName,proto3" json:"attachedVirtualmachineName,omitempty"`
	AttachedNodeName           string                     `protobuf:"bytes,17,opt,name=attachedNodeName,proto3" json:"attachedNodeName,omitempty"`
	GroupName                  string                     `protobuf:"bytes,18,opt,name=groupName,proto3" json:"groupName,omitempty"`
	LocationName               string                     `protobuf:"bytes,19,opt,name=locationName,proto3" json:"locationName,omitempty"`
	HyperVGeneration           common.HyperVGeneration    `protobuf:"varint,20,opt,name=hyperVGeneration,proto3,enum=moc.HyperVGeneration" json:"hyperVGeneration,omitempty"`
	DiskFileFormat             common.DiskFileFormat      `protobuf:"varint,21,opt,name=diskFileFormat,proto3,enum=moc.DiskFileFormat" json:"diskFileFormat,omitempty"`
	CloudInitDataSource        common.CloudInitDataSource `protobuf:"varint,22,opt,name=cloudInitDataSource,proto3,enum=moc.CloudInitDataSource" json:"cloudInitDataSource,omitempty"`
	Tags                       *common.Tags               `protobuf:"bytes,23,opt,name=tags,proto3" json:"tags,omitempty"`
	SourcePath                 string                     `protobuf:"bytes,24,opt,name=sourcePath,proto3" json:"sourcePath,omitempty"`
	SourceType                 common.ImageSource         `protobuf:"varint,25,opt,name=sourceType,proto3,enum=moc.ImageSource" json:"sourceType,omitempty"`
	TargetUrl                  string                     `protobuf:"bytes,26,opt,name=targetUrl,proto3" json:"targetUrl,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}                   `json:"-"`
	XXX_unrecognized           []byte                     `json:"-"`
	XXX_sizecache              int32                      `json:"-"`
}

func (m *VirtualHardDisk) Reset()         { *m = VirtualHardDisk{} }
func (m *VirtualHardDisk) String() string { return proto.CompactTextString(m) }
func (*VirtualHardDisk) ProtoMessage()    {}
func (*VirtualHardDisk) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3ff93a45373b349, []int{4}
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

func (m *VirtualHardDisk) GetContainerName() string {
	if m != nil {
		return m.ContainerName
	}
	return ""
}

func (m *VirtualHardDisk) GetPath() string {
	if m != nil {
		return m.Path
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

func (m *VirtualHardDisk) GetAttachedVirtualmachineName() string {
	if m != nil {
		return m.AttachedVirtualmachineName
	}
	return ""
}

func (m *VirtualHardDisk) GetAttachedNodeName() string {
	if m != nil {
		return m.AttachedNodeName
	}
	return ""
}

func (m *VirtualHardDisk) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *VirtualHardDisk) GetLocationName() string {
	if m != nil {
		return m.LocationName
	}
	return ""
}

func (m *VirtualHardDisk) GetHyperVGeneration() common.HyperVGeneration {
	if m != nil {
		return m.HyperVGeneration
	}
	return common.HyperVGeneration_HyperVGenerationV2
}

func (m *VirtualHardDisk) GetDiskFileFormat() common.DiskFileFormat {
	if m != nil {
		return m.DiskFileFormat
	}
	return common.DiskFileFormat_DiskFileFormatVHDX
}

func (m *VirtualHardDisk) GetCloudInitDataSource() common.CloudInitDataSource {
	if m != nil {
		return m.CloudInitDataSource
	}
	return common.CloudInitDataSource_NoCloud
}

func (m *VirtualHardDisk) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *VirtualHardDisk) GetSourcePath() string {
	if m != nil {
		return m.SourcePath
	}
	return ""
}

func (m *VirtualHardDisk) GetSourceType() common.ImageSource {
	if m != nil {
		return m.SourceType
	}
	return common.ImageSource_LOCAL_SOURCE
}

func (m *VirtualHardDisk) GetTargetUrl() string {
	if m != nil {
		return m.TargetUrl
	}
	return ""
}

type VirtualHardDiskPrecheckRequest struct {
	VirtualHardDisks     []*VirtualHardDisk `protobuf:"bytes,1,rep,name=VirtualHardDisks,proto3" json:"VirtualHardDisks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *VirtualHardDiskPrecheckRequest) Reset()         { *m = VirtualHardDiskPrecheckRequest{} }
func (m *VirtualHardDiskPrecheckRequest) String() string { return proto.CompactTextString(m) }
func (*VirtualHardDiskPrecheckRequest) ProtoMessage()    {}
func (*VirtualHardDiskPrecheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3ff93a45373b349, []int{5}
}

func (m *VirtualHardDiskPrecheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualHardDiskPrecheckRequest.Unmarshal(m, b)
}
func (m *VirtualHardDiskPrecheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualHardDiskPrecheckRequest.Marshal(b, m, deterministic)
}
func (m *VirtualHardDiskPrecheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualHardDiskPrecheckRequest.Merge(m, src)
}
func (m *VirtualHardDiskPrecheckRequest) XXX_Size() int {
	return xxx_messageInfo_VirtualHardDiskPrecheckRequest.Size(m)
}
func (m *VirtualHardDiskPrecheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualHardDiskPrecheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualHardDiskPrecheckRequest proto.InternalMessageInfo

func (m *VirtualHardDiskPrecheckRequest) GetVirtualHardDisks() []*VirtualHardDisk {
	if m != nil {
		return m.VirtualHardDisks
	}
	return nil
}

type VirtualHardDiskPrecheckResponse struct {
	// The precheck result: true if the precheck criteria is passed; otherwise, false
	Result *wrappers.BoolValue `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
	// The error message if the precheck is not passed; otherwise, empty string
	Error                string   `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VirtualHardDiskPrecheckResponse) Reset()         { *m = VirtualHardDiskPrecheckResponse{} }
func (m *VirtualHardDiskPrecheckResponse) String() string { return proto.CompactTextString(m) }
func (*VirtualHardDiskPrecheckResponse) ProtoMessage()    {}
func (*VirtualHardDiskPrecheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3ff93a45373b349, []int{6}
}

func (m *VirtualHardDiskPrecheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualHardDiskPrecheckResponse.Unmarshal(m, b)
}
func (m *VirtualHardDiskPrecheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualHardDiskPrecheckResponse.Marshal(b, m, deterministic)
}
func (m *VirtualHardDiskPrecheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualHardDiskPrecheckResponse.Merge(m, src)
}
func (m *VirtualHardDiskPrecheckResponse) XXX_Size() int {
	return xxx_messageInfo_VirtualHardDiskPrecheckResponse.Size(m)
}
func (m *VirtualHardDiskPrecheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualHardDiskPrecheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualHardDiskPrecheckResponse proto.InternalMessageInfo

func (m *VirtualHardDiskPrecheckResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *VirtualHardDiskPrecheckResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*VirtualHardDiskRequest)(nil), "moc.cloudagent.storage.VirtualHardDiskRequest")
	proto.RegisterType((*VirtualHardDiskResponse)(nil), "moc.cloudagent.storage.VirtualHardDiskResponse")
	proto.RegisterType((*VirtualHardDiskOperationRequest)(nil), "moc.cloudagent.storage.VirtualHardDiskOperationRequest")
	proto.RegisterType((*VirtualHardDiskOperationResponse)(nil), "moc.cloudagent.storage.VirtualHardDiskOperationResponse")
	proto.RegisterType((*VirtualHardDisk)(nil), "moc.cloudagent.storage.VirtualHardDisk")
	proto.RegisterType((*VirtualHardDiskPrecheckRequest)(nil), "moc.cloudagent.storage.VirtualHardDiskPrecheckRequest")
	proto.RegisterType((*VirtualHardDiskPrecheckResponse)(nil), "moc.cloudagent.storage.VirtualHardDiskPrecheckResponse")
}

func init() {
	proto.RegisterFile("moc_cloudagent_virtualharddisk.proto", fileDescriptor_e3ff93a45373b349)
}

var fileDescriptor_e3ff93a45373b349 = []byte{
	// 881 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x55, 0x41, 0x6f, 0xe3, 0x44,
	0x14, 0xc6, 0x6d, 0x9a, 0x26, 0xaf, 0xdb, 0x6c, 0x98, 0x76, 0xdb, 0xc1, 0x5a, 0x4a, 0x14, 0x56,
	0x10, 0x71, 0xb0, 0x97, 0x80, 0x00, 0x09, 0x09, 0xa9, 0xa5, 0x2c, 0x5b, 0x0e, 0x4b, 0xe5, 0x2e,
	0x3d, 0x70, 0x59, 0x4d, 0xc6, 0xb3, 0xce, 0x28, 0xb6, 0xc7, 0xcc, 0x8c, 0x8b, 0x82, 0x04, 0x07,
	0x4e, 0xfc, 0x0f, 0x4e, 0xfc, 0x00, 0x24, 0xc4, 0x8d, 0x7f, 0x86, 0xfc, 0xec, 0x34, 0xa9, 0x93,
	0xa2, 0xf4, 0xd0, 0x03, 0xa7, 0xc4, 0xdf, 0xf7, 0xbd, 0xef, 0xbd, 0x79, 0xcf, 0xf3, 0x0c, 0x4f,
	0x12, 0xc5, 0x5f, 0xf1, 0x58, 0xe5, 0x21, 0x8b, 0x44, 0x6a, 0x5f, 0x5d, 0x49, 0x6d, 0x73, 0x16,
	0x8f, 0x99, 0x0e, 0x43, 0x69, 0x26, 0x5e, 0xa6, 0x95, 0x55, 0xe4, 0x20, 0x51, 0xdc, 0x9b, 0xab,
	0x3c, 0x63, 0x95, 0x66, 0x91, 0x70, 0x8f, 0x22, 0xa5, 0xa2, 0x58, 0xf8, 0xa8, 0x1a, 0xe5, 0xaf,
	0xfd, 0x1f, 0x35, 0xcb, 0x32, 0xa1, 0x4d, 0x19, 0xe7, 0x1e, 0xa2, 0xbb, 0x4a, 0x12, 0x95, 0x56,
	0x3f, 0x25, 0xd1, 0xff, 0xdd, 0x81, 0x83, 0xcb, 0x32, 0xd5, 0x73, 0xa6, 0xc3, 0x53, 0x69, 0x26,
	0x81, 0xf8, 0x21, 0x17, 0xc6, 0x92, 0x0b, 0xe8, 0xd6, 0x18, 0x43, 0x9d, 0xde, 0xe6, 0x60, 0x67,
	0xf8, 0xbe, 0xb7, 0xba, 0x0c, 0xaf, 0xee, 0xb4, 0x64, 0x40, 0x3e, 0x86, 0xdd, 0x6f, 0x33, 0xa1,
	0x99, 0x95, 0x2a, 0x7d, 0x39, 0xcd, 0x04, 0xdd, 0xe8, 0x39, 0x83, 0xce, 0xb0, 0x83, 0x8e, 0xd7,
	0x4c, 0x70, 0x53, 0xd4, 0xff, 0xcb, 0x81, 0xc3, 0xa5, 0x2a, 0x4d, 0xa6, 0x52, 0x23, 0xee, 0xa7,
	0xcc, 0x21, 0x34, 0x03, 0x61, 0xf2, 0xd8, 0x62, 0x7d, 0x3b, 0x43, 0xd7, 0x2b, 0x1b, 0xec, 0xcd,
	0x1a, 0xec, 0x9d, 0x28, 0x15, 0x5f, 0xb2, 0x38, 0x17, 0x41, 0xa5, 0x24, 0xfb, 0xb0, 0xf5, 0x95,
	0xd6, 0x4a, 0xd3, 0xcd, 0x9e, 0x33, 0x68, 0x07, 0xe5, 0x43, 0xff, 0x6f, 0x07, 0xde, 0xa9, 0xd9,
	0xcf, 0x8f, 0x79, 0x9f, 0x9d, 0x3e, 0x59, 0xdd, 0xe9, 0xc7, 0xe8, 0x78, 0xae, 0xd5, 0x95, 0x0c,
	0x85, 0x3e, 0xe6, 0x5c, 0x18, 0x73, 0x6b, 0xdf, 0xff, 0x71, 0xa0, 0x77, 0x7b, 0xf1, 0xff, 0x8f,
	0x01, 0xfc, 0xd1, 0x82, 0x87, 0x35, 0x7b, 0x42, 0xa0, 0x91, 0xb2, 0x44, 0x50, 0x07, 0x85, 0xf8,
	0x9f, 0x74, 0x60, 0x43, 0x86, 0x98, 0xad, 0x1d, 0x6c, 0xc8, 0x90, 0x3c, 0x81, 0x5d, 0xae, 0x52,
	0xcb, 0x64, 0x2a, 0xf4, 0x8b, 0x42, 0x5c, 0xba, 0xde, 0x04, 0x09, 0x85, 0x46, 0xc6, 0xec, 0x98,
	0x36, 0x0a, 0xf2, 0xa4, 0xf1, 0xdb, 0x9f, 0xd4, 0x09, 0x10, 0x21, 0xef, 0x42, 0xd3, 0x58, 0x66,
	0x73, 0x43, 0xb7, 0xf0, 0x04, 0x3b, 0xd8, 0x8c, 0x0b, 0x84, 0x82, 0x8a, 0x2a, 0x0a, 0x31, 0xf2,
	0x27, 0x41, 0x9b, 0x3d, 0x67, 0xb0, 0x19, 0xe0, 0x7f, 0x42, 0x61, 0x3b, 0x9c, 0xa6, 0x2c, 0x91,
	0x9c, 0x6e, 0xf7, 0x9c, 0x41, 0x2b, 0x98, 0x3d, 0x92, 0xf7, 0xa0, 0x33, 0x8a, 0x15, 0x9f, 0x14,
	0xb2, 0xd1, 0xd4, 0x0a, 0x43, 0x5b, 0x3d, 0x67, 0xb0, 0x15, 0xd4, 0x50, 0xe2, 0x01, 0x89, 0x55,
	0x24, 0x39, 0x8b, 0x8d, 0xe0, 0x56, 0xe9, 0x52, 0xdb, 0x46, 0xed, 0x0a, 0x86, 0x3c, 0x85, 0xbd,
	0x6c, 0x3c, 0x35, 0xf5, 0x00, 0xc0, 0x80, 0x55, 0x14, 0xf9, 0x00, 0xba, 0x45, 0x1f, 0xb4, 0x8a,
	0x63, 0xa1, 0xd3, 0x3c, 0x19, 0x09, 0x4d, 0x77, 0xf0, 0x0c, 0x4b, 0x78, 0x51, 0xcd, 0x1c, 0x8b,
	0x15, 0xc7, 0xb7, 0x87, 0x3e, 0x40, 0xf5, 0x0a, 0x86, 0x1c, 0x01, 0x14, 0x1b, 0xaf, 0x72, 0xdd,
	0x45, 0xdd, 0x02, 0x52, 0xf8, 0x55, 0xcb, 0x31, 0x61, 0x7c, 0x2c, 0x53, 0x81, 0xd3, 0xe9, 0xe0,
	0x74, 0x56, 0x30, 0xc4, 0x85, 0x96, 0xe1, 0x46, 0xe2, 0x98, 0x1e, 0xa2, 0xea, 0xfa, 0x99, 0x7c,
	0x01, 0x2e, 0xb3, 0x96, 0xf1, 0xb1, 0x08, 0x2f, 0x97, 0x3d, 0xbb, 0xa8, 0xfe, 0x0f, 0x45, 0xd1,
	0x87, 0x19, 0xfb, 0x42, 0x85, 0x65, 0xd4, 0x9b, 0x18, 0xb5, 0x84, 0x93, 0xc7, 0xd0, 0x8e, 0xb4,
	0xca, 0x33, 0x14, 0x11, 0x14, 0xcd, 0x01, 0xd2, 0x87, 0x07, 0xb3, 0x0e, 0xa0, 0x60, 0x0f, 0x05,
	0x37, 0x30, 0x72, 0x0c, 0xdd, 0xf1, 0x34, 0x13, 0xfa, 0xf2, 0x6b, 0x91, 0x56, 0xb7, 0x90, 0xee,
	0xe3, 0xad, 0x7e, 0x84, 0x2f, 0xd7, 0xf3, 0x1a, 0x19, 0x2c, 0xc9, 0xc9, 0xe7, 0xd0, 0x29, 0x5a,
	0xf9, 0x4c, 0xc6, 0xe2, 0x99, 0xd2, 0x09, 0xb3, 0xf4, 0x11, 0x1a, 0xec, 0xa1, 0xc1, 0xe9, 0x0d,
	0x2a, 0xa8, 0x49, 0xc9, 0x37, 0xb0, 0x87, 0x97, 0xf9, 0x2c, 0x95, 0xf6, 0x94, 0x59, 0x76, 0xa1,
	0x72, 0xcd, 0x05, 0x3d, 0x40, 0x07, 0x8a, 0x0e, 0x5f, 0x2e, 0xf3, 0xc1, 0xaa, 0x20, 0xf2, 0x36,
	0x34, 0x2c, 0x8b, 0x0c, 0x3d, 0xc4, 0xcb, 0xd1, 0xc6, 0xe0, 0x97, 0x2c, 0x32, 0x01, 0xc2, 0xc5,
	0x4b, 0x60, 0x50, 0x78, 0x5e, 0x8c, 0x8d, 0x62, 0x33, 0x16, 0x10, 0xf2, 0x74, 0xc6, 0xe3, 0x6a,
	0x7b, 0x0b, 0x2b, 0xe8, 0xa2, 0xc9, 0x59, 0xc2, 0x22, 0x51, 0x65, 0x5e, 0xd0, 0x90, 0x3e, 0xb4,
	0x2d, 0xd3, 0x91, 0xb0, 0xdf, 0xe9, 0x98, 0xba, 0x0b, 0xd7, 0x75, 0x0e, 0xf7, 0x73, 0x38, 0xaa,
	0xad, 0x8a, 0x73, 0x2d, 0xf8, 0x58, 0xf0, 0x7b, 0xfd, 0x28, 0xf6, 0x27, 0x4b, 0x9f, 0x88, 0x79,
	0xda, 0x6a, 0xc9, 0xce, 0xf7, 0xa1, 0x73, 0xf7, 0x7d, 0xb8, 0xb1, 0xb0, 0x0f, 0x87, 0xbf, 0x6e,
	0xc2, 0x7e, 0x2d, 0xdb, 0x71, 0x51, 0x2f, 0x99, 0x40, 0xf3, 0x2c, 0xbd, 0x52, 0x13, 0x41, 0xbc,
	0x75, 0x8f, 0x52, 0x36, 0xc5, 0xf5, 0xd7, 0xd6, 0x97, 0xa7, 0xe9, 0xbf, 0x41, 0x7e, 0x86, 0xd6,
	0xec, 0x8c, 0xe4, 0x93, 0x35, 0xc3, 0x6b, 0xb3, 0x70, 0x3f, 0xbd, 0x73, 0xdc, 0x75, 0xfa, 0x5f,
	0x60, 0xbb, 0xfc, 0x90, 0x09, 0xb2, 0xae, 0x4b, 0xfd, 0xab, 0xed, 0x7e, 0x76, 0xf7, 0xc0, 0x59,
	0xfe, 0x93, 0x0f, 0xbf, 0xf7, 0x23, 0x69, 0xc7, 0xf9, 0xc8, 0xe3, 0x2a, 0xf1, 0x13, 0xc9, 0xb5,
	0x32, 0xea, 0xb5, 0xf5, 0x13, 0xc5, 0x7d, 0x9d, 0x71, 0x7f, 0xee, 0xea, 0x57, 0xae, 0xa3, 0x26,
	0x4e, 0xfa, 0xa3, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x2c, 0x01, 0x21, 0x29, 0x0a, 0x00,
	0x00,
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
	// Prechecks whether the system is able to create specified virtual hard disks (but does not actually create them).
	Precheck(ctx context.Context, in *VirtualHardDiskPrecheckRequest, opts ...grpc.CallOption) (*VirtualHardDiskPrecheckResponse, error)
	Operate(ctx context.Context, in *VirtualHardDiskOperationRequest, opts ...grpc.CallOption) (*VirtualHardDiskOperationResponse, error)
}

type virtualHardDiskAgentClient struct {
	cc *grpc.ClientConn
}

func NewVirtualHardDiskAgentClient(cc *grpc.ClientConn) VirtualHardDiskAgentClient {
	return &virtualHardDiskAgentClient{cc}
}

func (c *virtualHardDiskAgentClient) Invoke(ctx context.Context, in *VirtualHardDiskRequest, opts ...grpc.CallOption) (*VirtualHardDiskResponse, error) {
	out := new(VirtualHardDiskResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.storage.VirtualHardDiskAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualHardDiskAgentClient) Precheck(ctx context.Context, in *VirtualHardDiskPrecheckRequest, opts ...grpc.CallOption) (*VirtualHardDiskPrecheckResponse, error) {
	out := new(VirtualHardDiskPrecheckResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.storage.VirtualHardDiskAgent/Precheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualHardDiskAgentClient) Operate(ctx context.Context, in *VirtualHardDiskOperationRequest, opts ...grpc.CallOption) (*VirtualHardDiskOperationResponse, error) {
	out := new(VirtualHardDiskOperationResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.storage.VirtualHardDiskAgent/Operate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VirtualHardDiskAgentServer is the server API for VirtualHardDiskAgent service.
type VirtualHardDiskAgentServer interface {
	Invoke(context.Context, *VirtualHardDiskRequest) (*VirtualHardDiskResponse, error)
	// Prechecks whether the system is able to create specified virtual hard disks (but does not actually create them).
	Precheck(context.Context, *VirtualHardDiskPrecheckRequest) (*VirtualHardDiskPrecheckResponse, error)
	Operate(context.Context, *VirtualHardDiskOperationRequest) (*VirtualHardDiskOperationResponse, error)
}

// UnimplementedVirtualHardDiskAgentServer can be embedded to have forward compatible implementations.
type UnimplementedVirtualHardDiskAgentServer struct {
}

func (*UnimplementedVirtualHardDiskAgentServer) Invoke(ctx context.Context, req *VirtualHardDiskRequest) (*VirtualHardDiskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}
func (*UnimplementedVirtualHardDiskAgentServer) Precheck(ctx context.Context, req *VirtualHardDiskPrecheckRequest) (*VirtualHardDiskPrecheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Precheck not implemented")
}
func (*UnimplementedVirtualHardDiskAgentServer) Operate(ctx context.Context, req *VirtualHardDiskOperationRequest) (*VirtualHardDiskOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Operate not implemented")
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
		FullMethod: "/moc.cloudagent.storage.VirtualHardDiskAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualHardDiskAgentServer).Invoke(ctx, req.(*VirtualHardDiskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualHardDiskAgent_Precheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VirtualHardDiskPrecheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualHardDiskAgentServer).Precheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.storage.VirtualHardDiskAgent/Precheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualHardDiskAgentServer).Precheck(ctx, req.(*VirtualHardDiskPrecheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualHardDiskAgent_Operate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VirtualHardDiskOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualHardDiskAgentServer).Operate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.storage.VirtualHardDiskAgent/Operate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualHardDiskAgentServer).Operate(ctx, req.(*VirtualHardDiskOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VirtualHardDiskAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.cloudagent.storage.VirtualHardDiskAgent",
	HandlerType: (*VirtualHardDiskAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _VirtualHardDiskAgent_Invoke_Handler,
		},
		{
			MethodName: "Precheck",
			Handler:    _VirtualHardDiskAgent_Precheck_Handler,
		},
		{
			MethodName: "Operate",
			Handler:    _VirtualHardDiskAgent_Operate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moc_cloudagent_virtualharddisk.proto",
}
