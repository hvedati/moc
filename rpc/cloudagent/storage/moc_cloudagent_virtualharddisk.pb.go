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

type HyperVGenerationType int32

const (
	HyperVGenerationType_HyperVGenerationTypeV1 HyperVGenerationType = 0
	HyperVGenerationType_HyperVGenerationTypeV2 HyperVGenerationType = 1
)


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

type VirtualHardDisk struct {
	Name                       string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                         string         `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	ContainerName              string         `protobuf:"bytes,3,opt,name=containerName,proto3" json:"containerName,omitempty"`
	Path                       string         `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Status                     *common.Status `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Size                       int64          `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`
	Dynamic                    bool           `protobuf:"varint,7,opt,name=dynamic,proto3" json:"dynamic,omitempty"`
	Blocksizebytes             int32          `protobuf:"varint,8,opt,name=blocksizebytes,proto3" json:"blocksizebytes,omitempty"`
	Logicalsectorbytes         int32          `protobuf:"varint,9,opt,name=logicalsectorbytes,proto3" json:"logicalsectorbytes,omitempty"`
	Physicalsectorbytes        int32          `protobuf:"varint,10,opt,name=physicalsectorbytes,proto3" json:"physicalsectorbytes,omitempty"`
	Controllernumber           int64          `protobuf:"varint,11,opt,name=controllernumber,proto3" json:"controllernumber,omitempty"`
	Controllerlocation         int64          `protobuf:"varint,12,opt,name=controllerlocation,proto3" json:"controllerlocation,omitempty"`
	Disknumber                 int64          `protobuf:"varint,13,opt,name=disknumber,proto3" json:"disknumber,omitempty"`
	VirtualmachineName         string         `protobuf:"bytes,14,opt,name=virtualmachineName,proto3" json:"virtualmachineName,omitempty"`
	Scsipath                   string         `protobuf:"bytes,15,opt,name=scsipath,proto3" json:"scsipath,omitempty"`
	AttachedVirtualmachineName string         `protobuf:"bytes,16,opt,name=attachedVirtualmachineName,proto3" json:"attachedVirtualmachineName,omitempty"`
	AttachedNodeName           string         `protobuf:"bytes,17,opt,name=attachedNodeName,proto3" json:"attachedNodeName,omitempty"`
	GroupName                  string         `protobuf:"bytes,18,opt,name=groupName,proto3" json:"groupName,omitempty"`
	LocationName               string         `protobuf:"bytes,19,opt,name=locationName,proto3" json:"locationName,omitempty"`
	Hypervgenerationtype       HyperVGenerationType `protobuf:"varint,20,opt,name=hypervgenerationtype,proto3,enum=moc.cloudagent.storage.HyperVGenerationType" json:"hypervgenerationtype,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}       `json:"-"`
	XXX_unrecognized           []byte         `json:"-"`
	XXX_sizecache              int32          `json:"-"`
}

func (m *VirtualHardDisk) Reset()         { *m = VirtualHardDisk{} }
func (m *VirtualHardDisk) String() string { return proto.CompactTextString(m) }
func (*VirtualHardDisk) ProtoMessage()    {}
func (*VirtualHardDisk) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3ff93a45373b349, []int{2}
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

func init() {
	proto.RegisterType((*VirtualHardDiskRequest)(nil), "moc.cloudagent.storage.VirtualHardDiskRequest")
	proto.RegisterType((*VirtualHardDiskResponse)(nil), "moc.cloudagent.storage.VirtualHardDiskResponse")
	proto.RegisterType((*VirtualHardDisk)(nil), "moc.cloudagent.storage.VirtualHardDisk")
}

func init() {
	proto.RegisterFile("moc_cloudagent_virtualharddisk.proto", fileDescriptor_e3ff93a45373b349)
}

var fileDescriptor_e3ff93a45373b349 = []byte{
	// 611 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xc1, 0x6e, 0xd3, 0x4e,
	0x10, 0xc6, 0xff, 0x6e, 0x9b, 0xb4, 0x99, 0xb4, 0x69, 0xff, 0xdb, 0xaa, 0x5d, 0x45, 0xa8, 0x8a,
	0x42, 0x05, 0x11, 0x07, 0x1b, 0x02, 0x67, 0x24, 0x2a, 0x90, 0xe0, 0x52, 0x24, 0x17, 0xf5, 0xc0,
	0xa5, 0xda, 0xac, 0xb7, 0xce, 0x2a, 0x6b, 0x8f, 0xd9, 0x5d, 0x17, 0x85, 0x2b, 0x8f, 0xc2, 0x4b,
	0x70, 0xe6, 0xc9, 0x90, 0xc7, 0x49, 0x43, 0x93, 0x08, 0xf5, 0xc2, 0xa9, 0xbb, 0xdf, 0xf7, 0x9b,
	0xcf, 0xd3, 0xb1, 0x27, 0x70, 0x96, 0xa1, 0xbc, 0x96, 0x06, 0xcb, 0x44, 0xa4, 0x2a, 0xf7, 0xd7,
	0xb7, 0xda, 0xfa, 0x52, 0x98, 0xb1, 0xb0, 0x49, 0xa2, 0xdd, 0x24, 0x2c, 0x2c, 0x7a, 0x64, 0xc7,
	0x19, 0xca, 0x70, 0x41, 0x85, 0xce, 0xa3, 0x15, 0xa9, 0xea, 0x9e, 0xa6, 0x88, 0xa9, 0x51, 0x11,
	0x51, 0xa3, 0xf2, 0x26, 0xfa, 0x6a, 0x45, 0x51, 0x28, 0xeb, 0xea, 0xba, 0xee, 0x09, 0xa5, 0x63,
	0x96, 0x61, 0x3e, 0xfb, 0x53, 0x1b, 0xfd, 0x1f, 0x01, 0x1c, 0x5f, 0xd5, 0x8f, 0x7a, 0x2f, 0x6c,
	0xf2, 0x56, 0xbb, 0x49, 0xac, 0xbe, 0x94, 0xca, 0x79, 0x76, 0x09, 0x07, 0x4b, 0x8e, 0xe3, 0x41,
	0x6f, 0x73, 0xd0, 0x1e, 0x3e, 0x0d, 0xd7, 0xb7, 0x11, 0x2e, 0x27, 0xad, 0x04, 0xb0, 0x57, 0xb0,
	0xf7, 0xb1, 0x50, 0x56, 0x78, 0x8d, 0xf9, 0xa7, 0x69, 0xa1, 0xf8, 0x46, 0x2f, 0x18, 0x74, 0x86,
	0x1d, 0x4a, 0xbc, 0x73, 0xe2, 0xfb, 0x50, 0xff, 0x67, 0x00, 0x27, 0x2b, 0x5d, 0xba, 0x02, 0x73,
	0xa7, 0xfe, 0x4d, 0x9b, 0x43, 0x68, 0xc6, 0xca, 0x95, 0xc6, 0x53, 0x7f, 0xed, 0x61, 0x37, 0xac,
	0x07, 0x1c, 0xce, 0x07, 0x1c, 0x9e, 0x23, 0x9a, 0x2b, 0x61, 0x4a, 0x15, 0xcf, 0x48, 0x76, 0x04,
	0x8d, 0x77, 0xd6, 0xa2, 0xe5, 0x9b, 0xbd, 0x60, 0xd0, 0x8a, 0xeb, 0x4b, 0xff, 0x57, 0x03, 0xf6,
	0x97, 0xe2, 0x19, 0x83, 0xad, 0x5c, 0x64, 0x8a, 0x07, 0x04, 0xd2, 0x99, 0x75, 0x60, 0x43, 0x27,
	0xf4, 0xb4, 0x56, 0xbc, 0xa1, 0x13, 0x76, 0x06, 0x7b, 0x12, 0x73, 0x2f, 0x74, 0xae, 0xec, 0x45,
	0x05, 0xd7, 0xa9, 0xf7, 0xc5, 0x2a, 0xa9, 0x10, 0x7e, 0xcc, 0xb7, 0xea, 0xa4, 0xea, 0xcc, 0x1e,
	0x43, 0xd3, 0x79, 0xe1, 0x4b, 0xc7, 0x1b, 0xd4, 0x7b, 0x9b, 0xc6, 0x70, 0x49, 0x52, 0x3c, 0xb3,
	0xaa, 0x42, 0xa7, 0xbf, 0x29, 0xde, 0xec, 0x05, 0x83, 0xcd, 0x98, 0xce, 0x8c, 0xc3, 0x76, 0x32,
	0xcd, 0x45, 0xa6, 0x25, 0xdf, 0xee, 0x05, 0x83, 0x9d, 0x78, 0x7e, 0x65, 0x4f, 0xa0, 0x33, 0x32,
	0x28, 0x27, 0x15, 0x36, 0x9a, 0x7a, 0xe5, 0xf8, 0x4e, 0x2f, 0x18, 0x34, 0xe2, 0x25, 0x95, 0x85,
	0xc0, 0x0c, 0xa6, 0x5a, 0x0a, 0xe3, 0x94, 0xf4, 0x68, 0x6b, 0xb6, 0x45, 0xec, 0x1a, 0x87, 0x3d,
	0x87, 0xc3, 0x62, 0x3c, 0x75, 0xcb, 0x05, 0x40, 0x05, 0xeb, 0x2c, 0xf6, 0x0c, 0x0e, 0xaa, 0x09,
	0x58, 0x34, 0x46, 0xd9, 0xbc, 0xcc, 0x46, 0xca, 0xf2, 0x36, 0xfd, 0x0f, 0x2b, 0x7a, 0xd5, 0xcd,
	0x42, 0x33, 0x28, 0xe9, 0x7b, 0xe2, 0xbb, 0x44, 0xaf, 0x71, 0xd8, 0x29, 0x40, 0xb5, 0x6a, 0xb3,
	0xd4, 0x3d, 0xe2, 0xfe, 0x50, 0xaa, 0xbc, 0xd9, 0x56, 0x66, 0x42, 0x8e, 0x75, 0xae, 0xe8, 0xbd,
	0x74, 0x68, 0xf4, 0x6b, 0x1c, 0xd6, 0x85, 0x1d, 0x27, 0x9d, 0xa6, 0x17, 0xb4, 0x4f, 0xd4, 0xdd,
	0x9d, 0xbd, 0x86, 0xae, 0xf0, 0x5e, 0xc8, 0xb1, 0x4a, 0xae, 0x56, 0x33, 0x0f, 0x88, 0xfe, 0x0b,
	0x51, 0xcd, 0x61, 0xee, 0x5e, 0x60, 0x52, 0x57, 0xfd, 0x4f, 0x55, 0x2b, 0x3a, 0x7b, 0x04, 0xad,
	0xd4, 0x62, 0x59, 0x10, 0xc4, 0x08, 0x5a, 0x08, 0xac, 0x0f, 0xbb, 0xf3, 0x09, 0x10, 0x70, 0x48,
	0xc0, 0x3d, 0x6d, 0xf8, 0x3d, 0x80, 0xa3, 0xa5, 0x8f, 0xf8, 0x4d, 0xb5, 0x51, 0x6c, 0x02, 0xcd,
	0x0f, 0xf9, 0x2d, 0x4e, 0x14, 0x0b, 0x1f, 0xba, 0x6c, 0xf5, 0xaf, 0x4b, 0x37, 0x7a, 0x30, 0x5f,
	0xef, 0x79, 0xff, 0xbf, 0xf3, 0x17, 0x9f, 0xa3, 0x54, 0xfb, 0x71, 0x39, 0x0a, 0x25, 0x66, 0x51,
	0xa6, 0xa5, 0x45, 0x87, 0x37, 0x3e, 0xca, 0x50, 0x46, 0xb6, 0x90, 0xd1, 0x22, 0x2c, 0x9a, 0x85,
	0x8d, 0x9a, 0xb4, 0xaf, 0x2f, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x8f, 0xba, 0xce, 0x5e,
	0x05, 0x00, 0x00,
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

// VirtualHardDiskAgentServer is the server API for VirtualHardDiskAgent service.
type VirtualHardDiskAgentServer interface {
	Invoke(context.Context, *VirtualHardDiskRequest) (*VirtualHardDiskResponse, error)
}

// UnimplementedVirtualHardDiskAgentServer can be embedded to have forward compatible implementations.
type UnimplementedVirtualHardDiskAgentServer struct {
}

func (*UnimplementedVirtualHardDiskAgentServer) Invoke(ctx context.Context, req *VirtualHardDiskRequest) (*VirtualHardDiskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
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

var _VirtualHardDiskAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.cloudagent.storage.VirtualHardDiskAgent",
	HandlerType: (*VirtualHardDiskAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _VirtualHardDiskAgent_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moc_cloudagent_virtualharddisk.proto",
}
