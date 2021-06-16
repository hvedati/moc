// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_cloudagent_virtualnetwork.proto

package network

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

type VirtualNetworkType int32

const (
	VirtualNetworkType_NAT         VirtualNetworkType = 0
	VirtualNetworkType_Transparent VirtualNetworkType = 1
	VirtualNetworkType_L2Bridge    VirtualNetworkType = 2
	VirtualNetworkType_L2Tunnel    VirtualNetworkType = 3
	VirtualNetworkType_ICS         VirtualNetworkType = 4
	VirtualNetworkType_Private     VirtualNetworkType = 5
	VirtualNetworkType_Overlay     VirtualNetworkType = 6
	VirtualNetworkType_Internal    VirtualNetworkType = 7
	VirtualNetworkType_Mirrored    VirtualNetworkType = 8
)

var VirtualNetworkType_name = map[int32]string{
	0: "NAT",
	1: "Transparent",
	2: "L2Bridge",
	3: "L2Tunnel",
	4: "ICS",
	5: "Private",
	6: "Overlay",
	7: "Internal",
	8: "Mirrored",
}

var VirtualNetworkType_value = map[string]int32{
	"NAT":         0,
	"Transparent": 1,
	"L2Bridge":    2,
	"L2Tunnel":    3,
	"ICS":         4,
	"Private":     5,
	"Overlay":     6,
	"Internal":    7,
	"Mirrored":    8,
}

func (x VirtualNetworkType) String() string {
	return proto.EnumName(VirtualNetworkType_name, int32(x))
}

func (VirtualNetworkType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_65af3e74534f0214, []int{0}
}

type IPPoolType int32

const (
	IPPoolType_VM      IPPoolType = 0
	IPPoolType_VIPPool IPPoolType = 1
)

var IPPoolType_name = map[int32]string{
	0: "VM",
	1: "VIPPool",
}

var IPPoolType_value = map[string]int32{
	"VM":      0,
	"VIPPool": 1,
}

func (x IPPoolType) String() string {
	return proto.EnumName(IPPoolType_name, int32(x))
}

func (IPPoolType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_65af3e74534f0214, []int{1}
}

type VirtualNetworkRequest struct {
	VirtualNetworks      []*VirtualNetwork `protobuf:"bytes,1,rep,name=VirtualNetworks,proto3" json:"VirtualNetworks,omitempty"`
	OperationType        common.Operation  `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *VirtualNetworkRequest) Reset()         { *m = VirtualNetworkRequest{} }
func (m *VirtualNetworkRequest) String() string { return proto.CompactTextString(m) }
func (*VirtualNetworkRequest) ProtoMessage()    {}
func (*VirtualNetworkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_65af3e74534f0214, []int{0}
}

func (m *VirtualNetworkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualNetworkRequest.Unmarshal(m, b)
}
func (m *VirtualNetworkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualNetworkRequest.Marshal(b, m, deterministic)
}
func (m *VirtualNetworkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualNetworkRequest.Merge(m, src)
}
func (m *VirtualNetworkRequest) XXX_Size() int {
	return xxx_messageInfo_VirtualNetworkRequest.Size(m)
}
func (m *VirtualNetworkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualNetworkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualNetworkRequest proto.InternalMessageInfo

func (m *VirtualNetworkRequest) GetVirtualNetworks() []*VirtualNetwork {
	if m != nil {
		return m.VirtualNetworks
	}
	return nil
}

func (m *VirtualNetworkRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type VirtualNetworkResponse struct {
	VirtualNetworks      []*VirtualNetwork   `protobuf:"bytes,1,rep,name=VirtualNetworks,proto3" json:"VirtualNetworks,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *VirtualNetworkResponse) Reset()         { *m = VirtualNetworkResponse{} }
func (m *VirtualNetworkResponse) String() string { return proto.CompactTextString(m) }
func (*VirtualNetworkResponse) ProtoMessage()    {}
func (*VirtualNetworkResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_65af3e74534f0214, []int{1}
}

func (m *VirtualNetworkResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualNetworkResponse.Unmarshal(m, b)
}
func (m *VirtualNetworkResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualNetworkResponse.Marshal(b, m, deterministic)
}
func (m *VirtualNetworkResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualNetworkResponse.Merge(m, src)
}
func (m *VirtualNetworkResponse) XXX_Size() int {
	return xxx_messageInfo_VirtualNetworkResponse.Size(m)
}
func (m *VirtualNetworkResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualNetworkResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualNetworkResponse proto.InternalMessageInfo

func (m *VirtualNetworkResponse) GetVirtualNetworks() []*VirtualNetwork {
	if m != nil {
		return m.VirtualNetworks
	}
	return nil
}

func (m *VirtualNetworkResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *VirtualNetworkResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type VirtualNetwork struct {
	Name                 string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string             `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Subnets              []*Subnet          `protobuf:"bytes,3,rep,name=subnets,proto3" json:"subnets,omitempty"`
	Dns                  *common.Dns        `protobuf:"bytes,4,opt,name=dns,proto3" json:"dns,omitempty"`
	Type                 VirtualNetworkType `protobuf:"varint,5,opt,name=type,proto3,enum=moc.cloudagent.network.VirtualNetworkType" json:"type,omitempty"`
	Nodefqdn             string             `protobuf:"bytes,6,opt,name=nodefqdn,proto3" json:"nodefqdn,omitempty"`
	GroupName            string             `protobuf:"bytes,7,opt,name=groupName,proto3" json:"groupName,omitempty"`
	Status               *common.Status     `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	LocationName         string             `protobuf:"bytes,9,opt,name=locationName,proto3" json:"locationName,omitempty"`
	MacPoolName          string             `protobuf:"bytes,10,opt,name=macPoolName,proto3" json:"macPoolName,omitempty"`
	Tags                 *common.Tags       `protobuf:"bytes,12,opt,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *VirtualNetwork) Reset()         { *m = VirtualNetwork{} }
func (m *VirtualNetwork) String() string { return proto.CompactTextString(m) }
func (*VirtualNetwork) ProtoMessage()    {}
func (*VirtualNetwork) Descriptor() ([]byte, []int) {
	return fileDescriptor_65af3e74534f0214, []int{2}
}

func (m *VirtualNetwork) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualNetwork.Unmarshal(m, b)
}
func (m *VirtualNetwork) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualNetwork.Marshal(b, m, deterministic)
}
func (m *VirtualNetwork) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualNetwork.Merge(m, src)
}
func (m *VirtualNetwork) XXX_Size() int {
	return xxx_messageInfo_VirtualNetwork.Size(m)
}
func (m *VirtualNetwork) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualNetwork.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualNetwork proto.InternalMessageInfo

func (m *VirtualNetwork) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VirtualNetwork) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VirtualNetwork) GetSubnets() []*Subnet {
	if m != nil {
		return m.Subnets
	}
	return nil
}

func (m *VirtualNetwork) GetDns() *common.Dns {
	if m != nil {
		return m.Dns
	}
	return nil
}

func (m *VirtualNetwork) GetType() VirtualNetworkType {
	if m != nil {
		return m.Type
	}
	return VirtualNetworkType_NAT
}

func (m *VirtualNetwork) GetNodefqdn() string {
	if m != nil {
		return m.Nodefqdn
	}
	return ""
}

func (m *VirtualNetwork) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *VirtualNetwork) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *VirtualNetwork) GetLocationName() string {
	if m != nil {
		return m.LocationName
	}
	return ""
}

func (m *VirtualNetwork) GetMacPoolName() string {
	if m != nil {
		return m.MacPoolName
	}
	return ""
}

func (m *VirtualNetwork) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

type Subnet struct {
	Name                 string                    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string                    `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Cidr                 string                    `protobuf:"bytes,3,opt,name=cidr,proto3" json:"cidr,omitempty"`
	Routes               []*Route                  `protobuf:"bytes,4,rep,name=routes,proto3" json:"routes,omitempty"`
	Allocation           common.IPAllocationMethod `protobuf:"varint,5,opt,name=allocation,proto3,enum=moc.IPAllocationMethod" json:"allocation,omitempty"`
	Vlan                 uint32                    `protobuf:"varint,6,opt,name=vlan,proto3" json:"vlan,omitempty"`
	Ippools              []*IPPool                 `protobuf:"bytes,7,rep,name=ippools,proto3" json:"ippools,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Subnet) Reset()         { *m = Subnet{} }
func (m *Subnet) String() string { return proto.CompactTextString(m) }
func (*Subnet) ProtoMessage()    {}
func (*Subnet) Descriptor() ([]byte, []int) {
	return fileDescriptor_65af3e74534f0214, []int{3}
}

func (m *Subnet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subnet.Unmarshal(m, b)
}
func (m *Subnet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subnet.Marshal(b, m, deterministic)
}
func (m *Subnet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subnet.Merge(m, src)
}
func (m *Subnet) XXX_Size() int {
	return xxx_messageInfo_Subnet.Size(m)
}
func (m *Subnet) XXX_DiscardUnknown() {
	xxx_messageInfo_Subnet.DiscardUnknown(m)
}

var xxx_messageInfo_Subnet proto.InternalMessageInfo

func (m *Subnet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Subnet) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Subnet) GetCidr() string {
	if m != nil {
		return m.Cidr
	}
	return ""
}

func (m *Subnet) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

func (m *Subnet) GetAllocation() common.IPAllocationMethod {
	if m != nil {
		return m.Allocation
	}
	return common.IPAllocationMethod_Invalid
}

func (m *Subnet) GetVlan() uint32 {
	if m != nil {
		return m.Vlan
	}
	return 0
}

func (m *Subnet) GetIppools() []*IPPool {
	if m != nil {
		return m.Ippools
	}
	return nil
}

type IPPoolInfo struct {
	Used                 string   `protobuf:"bytes,1,opt,name=used,proto3" json:"used,omitempty"`
	Available            string   `protobuf:"bytes,2,opt,name=available,proto3" json:"available,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IPPoolInfo) Reset()         { *m = IPPoolInfo{} }
func (m *IPPoolInfo) String() string { return proto.CompactTextString(m) }
func (*IPPoolInfo) ProtoMessage()    {}
func (*IPPoolInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_65af3e74534f0214, []int{4}
}

func (m *IPPoolInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPPoolInfo.Unmarshal(m, b)
}
func (m *IPPoolInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPPoolInfo.Marshal(b, m, deterministic)
}
func (m *IPPoolInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPPoolInfo.Merge(m, src)
}
func (m *IPPoolInfo) XXX_Size() int {
	return xxx_messageInfo_IPPoolInfo.Size(m)
}
func (m *IPPoolInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_IPPoolInfo.DiscardUnknown(m)
}

var xxx_messageInfo_IPPoolInfo proto.InternalMessageInfo

func (m *IPPoolInfo) GetUsed() string {
	if m != nil {
		return m.Used
	}
	return ""
}

func (m *IPPoolInfo) GetAvailable() string {
	if m != nil {
		return m.Available
	}
	return ""
}

type IPPool struct {
	Name                 string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 IPPoolType  `protobuf:"varint,2,opt,name=type,proto3,enum=moc.cloudagent.network.IPPoolType" json:"type,omitempty"`
	Start                string      `protobuf:"bytes,3,opt,name=start,proto3" json:"start,omitempty"`
	End                  string      `protobuf:"bytes,4,opt,name=end,proto3" json:"end,omitempty"`
	Info                 *IPPoolInfo `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *IPPool) Reset()         { *m = IPPool{} }
func (m *IPPool) String() string { return proto.CompactTextString(m) }
func (*IPPool) ProtoMessage()    {}
func (*IPPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_65af3e74534f0214, []int{5}
}

func (m *IPPool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPPool.Unmarshal(m, b)
}
func (m *IPPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPPool.Marshal(b, m, deterministic)
}
func (m *IPPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPPool.Merge(m, src)
}
func (m *IPPool) XXX_Size() int {
	return xxx_messageInfo_IPPool.Size(m)
}
func (m *IPPool) XXX_DiscardUnknown() {
	xxx_messageInfo_IPPool.DiscardUnknown(m)
}

var xxx_messageInfo_IPPool proto.InternalMessageInfo

func (m *IPPool) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *IPPool) GetType() IPPoolType {
	if m != nil {
		return m.Type
	}
	return IPPoolType_VM
}

func (m *IPPool) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *IPPool) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

func (m *IPPool) GetInfo() *IPPoolInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type Ipam struct {
	Type                 string    `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Subnets              []*Subnet `protobuf:"bytes,2,rep,name=subnets,proto3" json:"subnets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Ipam) Reset()         { *m = Ipam{} }
func (m *Ipam) String() string { return proto.CompactTextString(m) }
func (*Ipam) ProtoMessage()    {}
func (*Ipam) Descriptor() ([]byte, []int) {
	return fileDescriptor_65af3e74534f0214, []int{6}
}

func (m *Ipam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipam.Unmarshal(m, b)
}
func (m *Ipam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipam.Marshal(b, m, deterministic)
}
func (m *Ipam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipam.Merge(m, src)
}
func (m *Ipam) XXX_Size() int {
	return xxx_messageInfo_Ipam.Size(m)
}
func (m *Ipam) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipam.DiscardUnknown(m)
}

var xxx_messageInfo_Ipam proto.InternalMessageInfo

func (m *Ipam) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Ipam) GetSubnets() []*Subnet {
	if m != nil {
		return m.Subnets
	}
	return nil
}

type Route struct {
	Nexthop              string   `protobuf:"bytes,1,opt,name=nexthop,proto3" json:"nexthop,omitempty"`
	Destinationprefix    string   `protobuf:"bytes,2,opt,name=destinationprefix,proto3" json:"destinationprefix,omitempty"`
	Metric               uint32   `protobuf:"varint,3,opt,name=metric,proto3" json:"metric,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_65af3e74534f0214, []int{7}
}

func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetNexthop() string {
	if m != nil {
		return m.Nexthop
	}
	return ""
}

func (m *Route) GetDestinationprefix() string {
	if m != nil {
		return m.Destinationprefix
	}
	return ""
}

func (m *Route) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func init() {
	proto.RegisterEnum("moc.cloudagent.network.VirtualNetworkType", VirtualNetworkType_name, VirtualNetworkType_value)
	proto.RegisterEnum("moc.cloudagent.network.IPPoolType", IPPoolType_name, IPPoolType_value)
	proto.RegisterType((*VirtualNetworkRequest)(nil), "moc.cloudagent.network.VirtualNetworkRequest")
	proto.RegisterType((*VirtualNetworkResponse)(nil), "moc.cloudagent.network.VirtualNetworkResponse")
	proto.RegisterType((*VirtualNetwork)(nil), "moc.cloudagent.network.VirtualNetwork")
	proto.RegisterType((*Subnet)(nil), "moc.cloudagent.network.Subnet")
	proto.RegisterType((*IPPoolInfo)(nil), "moc.cloudagent.network.IPPoolInfo")
	proto.RegisterType((*IPPool)(nil), "moc.cloudagent.network.IPPool")
	proto.RegisterType((*Ipam)(nil), "moc.cloudagent.network.Ipam")
	proto.RegisterType((*Route)(nil), "moc.cloudagent.network.Route")
}

func init() {
	proto.RegisterFile("moc_cloudagent_virtualnetwork.proto", fileDescriptor_65af3e74534f0214)
}

var fileDescriptor_65af3e74534f0214 = []byte{
	// 863 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xcf, 0x6f, 0xdb, 0x36,
	0x14, 0x8e, 0xfc, 0x43, 0xb6, 0x9f, 0x93, 0x54, 0xe3, 0xba, 0x54, 0x30, 0xd6, 0x20, 0x53, 0x81,
	0x21, 0x08, 0x36, 0x19, 0xf3, 0x7e, 0x9e, 0x0a, 0x24, 0xdb, 0x0e, 0x1e, 0x96, 0xd4, 0x60, 0x8c,
	0x1c, 0x76, 0x29, 0x68, 0x89, 0x56, 0x84, 0x4a, 0xa4, 0x4a, 0x52, 0x6e, 0x7b, 0xdb, 0x75, 0xf7,
	0xdd, 0xf7, 0x4f, 0xec, 0xef, 0x1b, 0x06, 0x3e, 0x49, 0x89, 0x9d, 0xb6, 0x99, 0x77, 0xe8, 0xc9,
	0x7c, 0xef, 0x7d, 0xef, 0xe3, 0xc7, 0xef, 0x91, 0x32, 0x3c, 0xc9, 0x65, 0xf4, 0x3c, 0xca, 0x64,
	0x19, 0xb3, 0x84, 0x0b, 0xf3, 0x7c, 0x95, 0x2a, 0x53, 0xb2, 0x4c, 0x70, 0xf3, 0x4a, 0xaa, 0x17,
	0x61, 0xa1, 0xa4, 0x91, 0xe4, 0x20, 0x97, 0x51, 0x78, 0x0b, 0x0a, 0xeb, 0xea, 0xe8, 0x30, 0x91,
	0x32, 0xc9, 0xf8, 0x18, 0x51, 0x8b, 0x72, 0x39, 0x7e, 0xa5, 0x58, 0x51, 0x70, 0xa5, 0xab, 0xbe,
	0xd1, 0x23, 0x24, 0x97, 0x79, 0x2e, 0x45, 0xfd, 0x53, 0x17, 0x0e, 0xd7, 0x0a, 0x35, 0xd9, 0x7a,
	0x3d, 0xf8, 0xcb, 0x81, 0x4f, 0xae, 0x2a, 0x25, 0x17, 0x55, 0x99, 0xf2, 0x97, 0x25, 0xd7, 0x86,
	0xcc, 0xe0, 0xc1, 0x66, 0x41, 0xfb, 0xce, 0x51, 0xfb, 0x78, 0x38, 0xf9, 0x3c, 0x7c, 0xb7, 0xc8,
	0xf0, 0x0e, 0xcf, 0xdd, 0x76, 0xf2, 0x0d, 0xec, 0x3d, 0x2b, 0xb8, 0x62, 0x26, 0x95, 0x62, 0xfe,
	0xa6, 0xe0, 0x7e, 0xeb, 0xc8, 0x39, 0xde, 0x9f, 0xec, 0x23, 0xdf, 0x4d, 0x85, 0x6e, 0x82, 0x82,
	0xbf, 0x1d, 0x38, 0xb8, 0xab, 0x50, 0x17, 0x52, 0x68, 0xfe, 0x01, 0x24, 0x4e, 0xc0, 0xa5, 0x5c,
	0x97, 0x99, 0x41, 0x6d, 0xc3, 0xc9, 0x28, 0xac, 0x8c, 0x0f, 0x1b, 0xe3, 0xc3, 0x33, 0x29, 0xb3,
	0x2b, 0x96, 0x95, 0x9c, 0xd6, 0x48, 0xf2, 0x10, 0xba, 0x3f, 0x2b, 0x25, 0x95, 0xdf, 0x3e, 0x72,
	0x8e, 0x07, 0xb4, 0x0a, 0x82, 0x3f, 0xdb, 0xb0, 0xbf, 0xc9, 0x4e, 0x08, 0x74, 0x04, 0xcb, 0xb9,
	0xef, 0x20, 0x0e, 0xd7, 0x64, 0x1f, 0x5a, 0x69, 0x8c, 0x9b, 0x0d, 0x68, 0x2b, 0x8d, 0xc9, 0x0f,
	0xd0, 0xd3, 0xe5, 0x42, 0x70, 0xa3, 0xfd, 0x36, 0x1e, 0xe5, 0xf0, 0x7d, 0x47, 0xb9, 0x44, 0x18,
	0x6d, 0xe0, 0x64, 0x04, 0xed, 0x58, 0x68, 0xbf, 0x83, 0xba, 0xfb, 0xd8, 0xf5, 0x93, 0xd0, 0xd4,
	0x26, 0xc9, 0x53, 0xe8, 0x18, 0x6b, 0x78, 0x17, 0x0d, 0x3f, 0xd9, 0xce, 0x1d, 0xeb, 0x3e, 0xc5,
	0x3e, 0x32, 0x82, 0xbe, 0x90, 0x31, 0x5f, 0xbe, 0x8c, 0x85, 0xef, 0xa2, 0xd6, 0x9b, 0x98, 0x7c,
	0x0a, 0x83, 0x44, 0xc9, 0xb2, 0xb8, 0xb0, 0x47, 0xeb, 0x61, 0xf1, 0x36, 0x41, 0x9e, 0x80, 0xab,
	0x0d, 0x33, 0xa5, 0xf6, 0xfb, 0x28, 0x6c, 0x88, 0x7b, 0x5f, 0x62, 0x8a, 0xd6, 0x25, 0x12, 0xc0,
	0x6e, 0x26, 0x23, 0x1c, 0x39, 0xb2, 0x0c, 0x90, 0x65, 0x23, 0x47, 0x8e, 0x60, 0x98, 0xb3, 0x68,
	0x26, 0x65, 0x86, 0x10, 0x40, 0xc8, 0x7a, 0x8a, 0x3c, 0x86, 0x8e, 0x61, 0x89, 0xf6, 0x77, 0x71,
	0xa3, 0x01, 0x6e, 0x34, 0x67, 0x89, 0xa6, 0x98, 0xfe, 0xa5, 0xd3, 0x1f, 0x7a, 0xbb, 0xc1, 0x3f,
	0x0e, 0xb8, 0x95, 0x73, 0x5b, 0x8d, 0x83, 0x40, 0x27, 0x4a, 0xe3, 0x66, 0xb4, 0xb8, 0x26, 0xdf,
	0x82, 0xab, 0x64, 0x69, 0xb8, 0xf5, 0xda, 0x4e, 0xe8, 0xf1, 0xfb, 0xec, 0xa4, 0x16, 0x45, 0x6b,
	0x30, 0xf9, 0x1e, 0x80, 0x65, 0xcd, 0x91, 0xea, 0x49, 0x3c, 0xc2, 0xd6, 0xe9, 0xec, 0xf4, 0xa6,
	0x70, 0xce, 0xcd, 0xb5, 0x8c, 0xe9, 0x1a, 0xd4, 0x6a, 0x58, 0x65, 0xac, 0x32, 0x7e, 0x8f, 0xe2,
	0xda, 0x5e, 0x93, 0xb4, 0x28, 0xa4, 0xcc, 0xb4, 0xdf, 0xbb, 0xff, 0x9a, 0x4c, 0x67, 0xd6, 0x20,
	0xda, 0xc0, 0x83, 0xa7, 0x00, 0x55, 0x6a, 0x2a, 0x96, 0xd2, 0x72, 0x97, 0x9a, 0xc7, 0x8d, 0x07,
	0x76, 0x6d, 0x07, 0xca, 0x56, 0x2c, 0xcd, 0xd8, 0x22, 0xe3, 0xb5, 0x15, 0xb7, 0x09, 0xfb, 0x1c,
	0xdd, 0x8a, 0xe0, 0x9d, 0x06, 0x7e, 0x57, 0xdf, 0xb4, 0xea, 0x69, 0x07, 0xf7, 0xab, 0x5a, 0xbb,
	0x61, 0x0f, 0xa1, 0xab, 0x0d, 0x53, 0xa6, 0x79, 0x44, 0x18, 0x10, 0x0f, 0xda, 0x5c, 0xc4, 0x78,
	0xa7, 0x07, 0xd4, 0x2e, 0x2d, 0x7f, 0x2a, 0x96, 0x12, 0xfd, 0x1b, 0xfe, 0x17, 0xbf, 0x3d, 0x22,
	0x45, 0x7c, 0x30, 0x87, 0xce, 0xb4, 0x60, 0xb9, 0xd5, 0x8c, 0xfa, 0x6a, 0xcd, 0xb8, 0xf7, 0xda,
	0x9b, 0x6b, 0xfd, 0xaf, 0x37, 0x17, 0x24, 0xd0, 0xc5, 0x21, 0x13, 0x1f, 0x7a, 0x82, 0xbf, 0x36,
	0xd7, 0xb2, 0xa8, 0x99, 0x9b, 0x90, 0x7c, 0x01, 0x1f, 0xc5, 0x5c, 0x9b, 0x54, 0xe0, 0x30, 0x0b,
	0xc5, 0x97, 0xe9, 0xeb, 0xda, 0xd5, 0xb7, 0x0b, 0xe4, 0x00, 0xdc, 0x9c, 0x1b, 0x95, 0x46, 0xe8,
	0xc3, 0x1e, 0xad, 0xa3, 0x93, 0x3f, 0x1c, 0x20, 0x6f, 0xbf, 0x4e, 0xd2, 0x83, 0xf6, 0xc5, 0xe9,
	0xdc, 0xdb, 0x21, 0x0f, 0x60, 0x38, 0x57, 0x4c, 0xe8, 0x82, 0x29, 0x2e, 0x8c, 0xe7, 0x90, 0x5d,
	0xe8, 0xff, 0x3a, 0x39, 0x53, 0x69, 0x9c, 0x70, 0xaf, 0x55, 0x45, 0xf3, 0x52, 0x08, 0x9e, 0x79,
	0x6d, 0xdb, 0x35, 0xfd, 0xf1, 0xd2, 0xeb, 0x90, 0x21, 0xf4, 0x66, 0x2a, 0x5d, 0x31, 0xc3, 0xbd,
	0xae, 0x0d, 0x9e, 0xad, 0xb8, 0xca, 0xd8, 0x1b, 0xcf, 0xb5, 0x0d, 0x53, 0x61, 0xb8, 0x12, 0x2c,
	0xf3, 0x7a, 0x36, 0x3a, 0x4f, 0xed, 0x57, 0x8d, 0xc7, 0x5e, 0xff, 0xe4, 0xb3, 0xe6, 0x06, 0xa1,
	0x04, 0x17, 0x5a, 0x57, 0xe7, 0xde, 0x8e, 0x6d, 0xbf, 0xaa, 0xd2, 0x9e, 0x33, 0xf9, 0xdd, 0x81,
	0x8f, 0x37, 0xe5, 0x9e, 0x5a, 0x1f, 0x49, 0x0a, 0xee, 0x54, 0xac, 0xe4, 0x0b, 0x4e, 0xbe, 0xdc,
	0xf2, 0x0b, 0x5d, 0xfd, 0x19, 0x8d, 0xc2, 0x6d, 0xe1, 0xd5, 0x3f, 0x43, 0xb0, 0x73, 0xf6, 0xd5,
	0x6f, 0xe3, 0x24, 0x35, 0xd7, 0xe5, 0x22, 0x8c, 0x64, 0x3e, 0xce, 0xd3, 0x48, 0x49, 0x2d, 0x97,
	0x66, 0x9c, 0xcb, 0x68, 0xac, 0x8a, 0x68, 0x7c, 0xcb, 0x35, 0xae, 0xb9, 0x16, 0x2e, 0x7e, 0xe4,
	0xbf, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0xc0, 0xc9, 0x63, 0x62, 0xaa, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VirtualNetworkAgentClient is the client API for VirtualNetworkAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VirtualNetworkAgentClient interface {
	Invoke(ctx context.Context, in *VirtualNetworkRequest, opts ...grpc.CallOption) (*VirtualNetworkResponse, error)
}

type virtualNetworkAgentClient struct {
	cc *grpc.ClientConn
}

func NewVirtualNetworkAgentClient(cc *grpc.ClientConn) VirtualNetworkAgentClient {
	return &virtualNetworkAgentClient{cc}
}

func (c *virtualNetworkAgentClient) Invoke(ctx context.Context, in *VirtualNetworkRequest, opts ...grpc.CallOption) (*VirtualNetworkResponse, error) {
	out := new(VirtualNetworkResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.network.VirtualNetworkAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VirtualNetworkAgentServer is the server API for VirtualNetworkAgent service.
type VirtualNetworkAgentServer interface {
	Invoke(context.Context, *VirtualNetworkRequest) (*VirtualNetworkResponse, error)
}

// UnimplementedVirtualNetworkAgentServer can be embedded to have forward compatible implementations.
type UnimplementedVirtualNetworkAgentServer struct {
}

func (*UnimplementedVirtualNetworkAgentServer) Invoke(ctx context.Context, req *VirtualNetworkRequest) (*VirtualNetworkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}

func RegisterVirtualNetworkAgentServer(s *grpc.Server, srv VirtualNetworkAgentServer) {
	s.RegisterService(&_VirtualNetworkAgent_serviceDesc, srv)
}

func _VirtualNetworkAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VirtualNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualNetworkAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.network.VirtualNetworkAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualNetworkAgentServer).Invoke(ctx, req.(*VirtualNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VirtualNetworkAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.cloudagent.network.VirtualNetworkAgent",
	HandlerType: (*VirtualNetworkAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _VirtualNetworkAgent_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moc_cloudagent_virtualnetwork.proto",
}
