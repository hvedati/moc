// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_cloudagent_networkinterface.proto

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

type NetworkInterface_NetworkInterfaceType int32

const (
	NetworkInterface_Local  NetworkInterface_NetworkInterfaceType = 0
	NetworkInterface_Remote NetworkInterface_NetworkInterfaceType = 1
)

var NetworkInterface_NetworkInterfaceType_name = map[int32]string{
	0: "Local",
	1: "Remote",
}

var NetworkInterface_NetworkInterfaceType_value = map[string]int32{
	"Local":  0,
	"Remote": 1,
}

func (x NetworkInterface_NetworkInterfaceType) String() string {
	return proto.EnumName(NetworkInterface_NetworkInterfaceType_name, int32(x))
}

func (NetworkInterface_NetworkInterfaceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_641284ba8360303c, []int{6, 0}
}

type NetworkInterfaceRequest struct {
	NetworkInterfaces    []*NetworkInterface `protobuf:"bytes,1,rep,name=NetworkInterfaces,proto3" json:"NetworkInterfaces,omitempty"`
	OperationType        common.Operation    `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *NetworkInterfaceRequest) Reset()         { *m = NetworkInterfaceRequest{} }
func (m *NetworkInterfaceRequest) String() string { return proto.CompactTextString(m) }
func (*NetworkInterfaceRequest) ProtoMessage()    {}
func (*NetworkInterfaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_641284ba8360303c, []int{0}
}

func (m *NetworkInterfaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkInterfaceRequest.Unmarshal(m, b)
}
func (m *NetworkInterfaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkInterfaceRequest.Marshal(b, m, deterministic)
}
func (m *NetworkInterfaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkInterfaceRequest.Merge(m, src)
}
func (m *NetworkInterfaceRequest) XXX_Size() int {
	return xxx_messageInfo_NetworkInterfaceRequest.Size(m)
}
func (m *NetworkInterfaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkInterfaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkInterfaceRequest proto.InternalMessageInfo

func (m *NetworkInterfaceRequest) GetNetworkInterfaces() []*NetworkInterface {
	if m != nil {
		return m.NetworkInterfaces
	}
	return nil
}

func (m *NetworkInterfaceRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type NetworkInterfaceResponse struct {
	NetworkInterfaces    []*NetworkInterface `protobuf:"bytes,1,rep,name=NetworkInterfaces,proto3" json:"NetworkInterfaces,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *NetworkInterfaceResponse) Reset()         { *m = NetworkInterfaceResponse{} }
func (m *NetworkInterfaceResponse) String() string { return proto.CompactTextString(m) }
func (*NetworkInterfaceResponse) ProtoMessage()    {}
func (*NetworkInterfaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_641284ba8360303c, []int{1}
}

func (m *NetworkInterfaceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkInterfaceResponse.Unmarshal(m, b)
}
func (m *NetworkInterfaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkInterfaceResponse.Marshal(b, m, deterministic)
}
func (m *NetworkInterfaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkInterfaceResponse.Merge(m, src)
}
func (m *NetworkInterfaceResponse) XXX_Size() int {
	return xxx_messageInfo_NetworkInterfaceResponse.Size(m)
}
func (m *NetworkInterfaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkInterfaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkInterfaceResponse proto.InternalMessageInfo

func (m *NetworkInterfaceResponse) GetNetworkInterfaces() []*NetworkInterface {
	if m != nil {
		return m.NetworkInterfaces
	}
	return nil
}

func (m *NetworkInterfaceResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *NetworkInterfaceResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type NetworkInterfacePrecheckRequest struct {
	NetworkInterfaces    []*NetworkInterface `protobuf:"bytes,1,rep,name=NetworkInterfaces,proto3" json:"NetworkInterfaces,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *NetworkInterfacePrecheckRequest) Reset()         { *m = NetworkInterfacePrecheckRequest{} }
func (m *NetworkInterfacePrecheckRequest) String() string { return proto.CompactTextString(m) }
func (*NetworkInterfacePrecheckRequest) ProtoMessage()    {}
func (*NetworkInterfacePrecheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_641284ba8360303c, []int{2}
}

func (m *NetworkInterfacePrecheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkInterfacePrecheckRequest.Unmarshal(m, b)
}
func (m *NetworkInterfacePrecheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkInterfacePrecheckRequest.Marshal(b, m, deterministic)
}
func (m *NetworkInterfacePrecheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkInterfacePrecheckRequest.Merge(m, src)
}
func (m *NetworkInterfacePrecheckRequest) XXX_Size() int {
	return xxx_messageInfo_NetworkInterfacePrecheckRequest.Size(m)
}
func (m *NetworkInterfacePrecheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkInterfacePrecheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkInterfacePrecheckRequest proto.InternalMessageInfo

func (m *NetworkInterfacePrecheckRequest) GetNetworkInterfaces() []*NetworkInterface {
	if m != nil {
		return m.NetworkInterfaces
	}
	return nil
}

type NetworkInterfacePrecheckResponse struct {
	// The precheck result: true if the precheck criteria is passed; otherwise, false
	Result *wrappers.BoolValue `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
	// The error message if the precheck is not passed; otherwise, empty string
	Error                string   `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkInterfacePrecheckResponse) Reset()         { *m = NetworkInterfacePrecheckResponse{} }
func (m *NetworkInterfacePrecheckResponse) String() string { return proto.CompactTextString(m) }
func (*NetworkInterfacePrecheckResponse) ProtoMessage()    {}
func (*NetworkInterfacePrecheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_641284ba8360303c, []int{3}
}

func (m *NetworkInterfacePrecheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkInterfacePrecheckResponse.Unmarshal(m, b)
}
func (m *NetworkInterfacePrecheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkInterfacePrecheckResponse.Marshal(b, m, deterministic)
}
func (m *NetworkInterfacePrecheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkInterfacePrecheckResponse.Merge(m, src)
}
func (m *NetworkInterfacePrecheckResponse) XXX_Size() int {
	return xxx_messageInfo_NetworkInterfacePrecheckResponse.Size(m)
}
func (m *NetworkInterfacePrecheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkInterfacePrecheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkInterfacePrecheckResponse proto.InternalMessageInfo

func (m *NetworkInterfacePrecheckResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *NetworkInterfacePrecheckResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type InboundNatRule struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InboundNatRule) Reset()         { *m = InboundNatRule{} }
func (m *InboundNatRule) String() string { return proto.CompactTextString(m) }
func (*InboundNatRule) ProtoMessage()    {}
func (*InboundNatRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_641284ba8360303c, []int{4}
}

func (m *InboundNatRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InboundNatRule.Unmarshal(m, b)
}
func (m *InboundNatRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InboundNatRule.Marshal(b, m, deterministic)
}
func (m *InboundNatRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InboundNatRule.Merge(m, src)
}
func (m *InboundNatRule) XXX_Size() int {
	return xxx_messageInfo_InboundNatRule.Size(m)
}
func (m *InboundNatRule) XXX_DiscardUnknown() {
	xxx_messageInfo_InboundNatRule.DiscardUnknown(m)
}

var xxx_messageInfo_InboundNatRule proto.InternalMessageInfo

func (m *InboundNatRule) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type IpConfiguration struct {
	Ipaddress               string                                `protobuf:"bytes,1,opt,name=ipaddress,proto3" json:"ipaddress,omitempty"`
	Prefixlength            string                                `protobuf:"bytes,2,opt,name=prefixlength,proto3" json:"prefixlength,omitempty"`
	Subnetid                string                                `protobuf:"bytes,3,opt,name=subnetid,proto3" json:"subnetid,omitempty"`
	Primary                 bool                                  `protobuf:"varint,4,opt,name=primary,proto3" json:"primary,omitempty"`
	Loadbalanceraddresspool []string                              `protobuf:"bytes,5,rep,name=loadbalanceraddresspool,proto3" json:"loadbalanceraddresspool,omitempty"`
	Allocation              common.IPAllocationMethod             `protobuf:"varint,6,opt,name=allocation,proto3,enum=moc.IPAllocationMethod" json:"allocation,omitempty"`
	Gateway                 string                                `protobuf:"bytes,7,opt,name=gateway,proto3" json:"gateway,omitempty"`
	InboundNatRules         []*InboundNatRule                     `protobuf:"bytes,8,rep,name=inboundNatRules,proto3" json:"inboundNatRules,omitempty"`
	Tags                    *common.Tags                          `protobuf:"bytes,9,opt,name=tags,proto3" json:"tags,omitempty"`
	NetworkType             common.NetworkType                    `protobuf:"varint,10,opt,name=networkType,proto3,enum=moc.NetworkType" json:"networkType,omitempty"`
	NetworkSecurityGroupRef *common.NetworkSecurityGroupReference `protobuf:"bytes,12,opt,name=networkSecurityGroupRef,proto3" json:"networkSecurityGroupRef,omitempty"`
	Subnet                  *common.SubnetReference               `protobuf:"bytes,13,opt,name=subnet,proto3" json:"subnet,omitempty"`
	PublicIPAddressRef      *common.PublicIPAddressReference      `protobuf:"bytes,14,opt,name=publicIPAddressRef,proto3" json:"publicIPAddressRef,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                              `json:"-"`
	XXX_unrecognized        []byte                                `json:"-"`
	XXX_sizecache           int32                                 `json:"-"`
}

func (m *IpConfiguration) Reset()         { *m = IpConfiguration{} }
func (m *IpConfiguration) String() string { return proto.CompactTextString(m) }
func (*IpConfiguration) ProtoMessage()    {}
func (*IpConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_641284ba8360303c, []int{5}
}

func (m *IpConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpConfiguration.Unmarshal(m, b)
}
func (m *IpConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpConfiguration.Marshal(b, m, deterministic)
}
func (m *IpConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpConfiguration.Merge(m, src)
}
func (m *IpConfiguration) XXX_Size() int {
	return xxx_messageInfo_IpConfiguration.Size(m)
}
func (m *IpConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_IpConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_IpConfiguration proto.InternalMessageInfo

func (m *IpConfiguration) GetIpaddress() string {
	if m != nil {
		return m.Ipaddress
	}
	return ""
}

func (m *IpConfiguration) GetPrefixlength() string {
	if m != nil {
		return m.Prefixlength
	}
	return ""
}

func (m *IpConfiguration) GetSubnetid() string {
	if m != nil {
		return m.Subnetid
	}
	return ""
}

func (m *IpConfiguration) GetPrimary() bool {
	if m != nil {
		return m.Primary
	}
	return false
}

func (m *IpConfiguration) GetLoadbalanceraddresspool() []string {
	if m != nil {
		return m.Loadbalanceraddresspool
	}
	return nil
}

func (m *IpConfiguration) GetAllocation() common.IPAllocationMethod {
	if m != nil {
		return m.Allocation
	}
	return common.IPAllocationMethod_Invalid
}

func (m *IpConfiguration) GetGateway() string {
	if m != nil {
		return m.Gateway
	}
	return ""
}

func (m *IpConfiguration) GetInboundNatRules() []*InboundNatRule {
	if m != nil {
		return m.InboundNatRules
	}
	return nil
}

func (m *IpConfiguration) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *IpConfiguration) GetNetworkType() common.NetworkType {
	if m != nil {
		return m.NetworkType
	}
	return common.NetworkType_VIRTUAL_NETWORK
}

func (m *IpConfiguration) GetNetworkSecurityGroupRef() *common.NetworkSecurityGroupReference {
	if m != nil {
		return m.NetworkSecurityGroupRef
	}
	return nil
}

func (m *IpConfiguration) GetSubnet() *common.SubnetReference {
	if m != nil {
		return m.Subnet
	}
	return nil
}

func (m *IpConfiguration) GetPublicIPAddressRef() *common.PublicIPAddressReference {
	if m != nil {
		return m.PublicIPAddressRef
	}
	return nil
}

type NetworkInterface struct {
	Name                 string                                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string                                `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Type                 NetworkInterface_NetworkInterfaceType `protobuf:"varint,3,opt,name=type,proto3,enum=moc.cloudagent.network.NetworkInterface_NetworkInterfaceType" json:"type,omitempty"`
	IpConfigurations     []*IpConfiguration                    `protobuf:"bytes,4,rep,name=ipConfigurations,proto3" json:"ipConfigurations,omitempty"`
	Macaddress           string                                `protobuf:"bytes,5,opt,name=macaddress,proto3" json:"macaddress,omitempty"`
	Dns                  *common.Dns                           `protobuf:"bytes,6,opt,name=dns,proto3" json:"dns,omitempty"`
	NodeName             string                                `protobuf:"bytes,7,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
	GroupName            string                                `protobuf:"bytes,8,opt,name=groupName,proto3" json:"groupName,omitempty"`
	LocationName         string                                `protobuf:"bytes,9,opt,name=locationName,proto3" json:"locationName,omitempty"`
	Status               *common.Status                        `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	VirtualMachineName   string                                `protobuf:"bytes,11,opt,name=virtualMachineName,proto3" json:"virtualMachineName,omitempty"`
	IovWeight            uint32                                `protobuf:"varint,12,opt,name=iovWeight,proto3" json:"iovWeight,omitempty"`
	Tags                 *common.Tags                          `protobuf:"bytes,13,opt,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *NetworkInterface) Reset()         { *m = NetworkInterface{} }
func (m *NetworkInterface) String() string { return proto.CompactTextString(m) }
func (*NetworkInterface) ProtoMessage()    {}
func (*NetworkInterface) Descriptor() ([]byte, []int) {
	return fileDescriptor_641284ba8360303c, []int{6}
}

func (m *NetworkInterface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkInterface.Unmarshal(m, b)
}
func (m *NetworkInterface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkInterface.Marshal(b, m, deterministic)
}
func (m *NetworkInterface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkInterface.Merge(m, src)
}
func (m *NetworkInterface) XXX_Size() int {
	return xxx_messageInfo_NetworkInterface.Size(m)
}
func (m *NetworkInterface) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkInterface.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkInterface proto.InternalMessageInfo

func (m *NetworkInterface) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkInterface) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NetworkInterface) GetType() NetworkInterface_NetworkInterfaceType {
	if m != nil {
		return m.Type
	}
	return NetworkInterface_Local
}

func (m *NetworkInterface) GetIpConfigurations() []*IpConfiguration {
	if m != nil {
		return m.IpConfigurations
	}
	return nil
}

func (m *NetworkInterface) GetMacaddress() string {
	if m != nil {
		return m.Macaddress
	}
	return ""
}

func (m *NetworkInterface) GetDns() *common.Dns {
	if m != nil {
		return m.Dns
	}
	return nil
}

func (m *NetworkInterface) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *NetworkInterface) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *NetworkInterface) GetLocationName() string {
	if m != nil {
		return m.LocationName
	}
	return ""
}

func (m *NetworkInterface) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *NetworkInterface) GetVirtualMachineName() string {
	if m != nil {
		return m.VirtualMachineName
	}
	return ""
}

func (m *NetworkInterface) GetIovWeight() uint32 {
	if m != nil {
		return m.IovWeight
	}
	return 0
}

func (m *NetworkInterface) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterEnum("moc.cloudagent.network.NetworkInterface_NetworkInterfaceType", NetworkInterface_NetworkInterfaceType_name, NetworkInterface_NetworkInterfaceType_value)
	proto.RegisterType((*NetworkInterfaceRequest)(nil), "moc.cloudagent.network.NetworkInterfaceRequest")
	proto.RegisterType((*NetworkInterfaceResponse)(nil), "moc.cloudagent.network.NetworkInterfaceResponse")
	proto.RegisterType((*NetworkInterfacePrecheckRequest)(nil), "moc.cloudagent.network.NetworkInterfacePrecheckRequest")
	proto.RegisterType((*NetworkInterfacePrecheckResponse)(nil), "moc.cloudagent.network.NetworkInterfacePrecheckResponse")
	proto.RegisterType((*InboundNatRule)(nil), "moc.cloudagent.network.InboundNatRule")
	proto.RegisterType((*IpConfiguration)(nil), "moc.cloudagent.network.IpConfiguration")
	proto.RegisterType((*NetworkInterface)(nil), "moc.cloudagent.network.NetworkInterface")
}

func init() {
	proto.RegisterFile("moc_cloudagent_networkinterface.proto", fileDescriptor_641284ba8360303c)
}

var fileDescriptor_641284ba8360303c = []byte{
	// 901 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x4d, 0x6f, 0xe3, 0x44,
	0x18, 0xde, 0x7c, 0x36, 0x79, 0xd3, 0x66, 0xc3, 0xa8, 0x10, 0x2b, 0x62, 0x4b, 0x64, 0x16, 0xc8,
	0x01, 0x1c, 0x08, 0x48, 0xbb, 0x17, 0x0e, 0x2d, 0x20, 0x14, 0x44, 0x4b, 0x98, 0xae, 0x16, 0x09,
	0x21, 0xad, 0x26, 0xf6, 0xc4, 0x19, 0xd5, 0x9e, 0x31, 0xe3, 0x71, 0x4b, 0x4e, 0x1c, 0xe1, 0x9f,
	0xc0, 0x1f, 0xe0, 0xc6, 0x95, 0xff, 0x85, 0xfc, 0xda, 0x6e, 0x9c, 0x2f, 0xa9, 0x7b, 0xd8, 0x53,
	0x32, 0xef, 0xf3, 0xbc, 0xcf, 0xbc, 0x9f, 0x63, 0xf8, 0x20, 0x54, 0xee, 0x2b, 0x37, 0x50, 0x89,
	0xc7, 0x7c, 0x2e, 0xcd, 0x2b, 0xc9, 0xcd, 0x9d, 0xd2, 0x37, 0x42, 0x1a, 0xae, 0x17, 0xcc, 0xe5,
	0x4e, 0xa4, 0x95, 0x51, 0xe4, 0x9d, 0x50, 0xb9, 0xce, 0x9a, 0xe6, 0xe4, 0xb4, 0xc1, 0x99, 0xaf,
	0x94, 0x1f, 0xf0, 0x31, 0xb2, 0xe6, 0xc9, 0x62, 0x7c, 0xa7, 0x59, 0x14, 0x71, 0x1d, 0x67, 0x7e,
	0x83, 0x3e, 0xca, 0xab, 0x30, 0x54, 0x32, 0xff, 0xc9, 0x81, 0xb3, 0x12, 0x90, 0x8b, 0x95, 0x71,
	0xfb, 0xaf, 0x0a, 0xf4, 0xaf, 0x32, 0xfb, 0xb4, 0x88, 0x85, 0xf2, 0x5f, 0x13, 0x1e, 0x1b, 0xf2,
	0x12, 0xde, 0xda, 0x86, 0x62, 0xab, 0x32, 0xac, 0x8d, 0x3a, 0x93, 0x91, 0xb3, 0x3f, 0x50, 0x67,
	0x47, 0x6b, 0x57, 0x82, 0x7c, 0x01, 0x27, 0x3f, 0x44, 0x5c, 0x33, 0x23, 0x94, 0x7c, 0xb1, 0x8a,
	0xb8, 0x55, 0x1d, 0x56, 0x46, 0xdd, 0x49, 0x17, 0x35, 0xef, 0x11, 0xba, 0x49, 0xb2, 0xff, 0xad,
	0x80, 0xb5, 0x1b, 0x69, 0x1c, 0x29, 0x19, 0xf3, 0x37, 0x16, 0xea, 0x04, 0x9a, 0x94, 0xc7, 0x49,
	0x60, 0x30, 0xc6, 0xce, 0x64, 0xe0, 0x64, 0x8d, 0x70, 0x8a, 0x46, 0x38, 0x17, 0x4a, 0x05, 0x2f,
	0x59, 0x90, 0x70, 0x9a, 0x33, 0xc9, 0x29, 0x34, 0xbe, 0xd1, 0x5a, 0x69, 0xab, 0x36, 0xac, 0x8c,
	0xda, 0x34, 0x3b, 0xd8, 0x2b, 0x78, 0x6f, 0x5b, 0x7e, 0xa6, 0xb9, 0xbb, 0xe4, 0xee, 0xcd, 0x1b,
	0xae, 0xb7, 0x1d, 0xc0, 0xf0, 0xf0, 0xd5, 0x79, 0x01, 0xd7, 0x89, 0x56, 0x5e, 0x3f, 0xd1, 0x6a,
	0x39, 0xd1, 0xa7, 0xd0, 0x9d, 0xca, 0xb9, 0x4a, 0xa4, 0x77, 0xc5, 0x0c, 0x4d, 0x02, 0x4e, 0x08,
	0xd4, 0x25, 0x0b, 0x39, 0x2a, 0xb7, 0x29, 0xfe, 0xb7, 0xff, 0x6e, 0xc0, 0xe3, 0x69, 0xf4, 0x95,
	0x92, 0x0b, 0xe1, 0x27, 0x59, 0x97, 0x89, 0x0d, 0x6d, 0x11, 0x31, 0xcf, 0xd3, 0x3c, 0x8e, 0x33,
	0xf2, 0x45, 0xfd, 0xcf, 0x7f, 0xac, 0x0a, 0x5d, 0x9b, 0x89, 0x0d, 0xc7, 0x91, 0xe6, 0x0b, 0xf1,
	0x5b, 0xc0, 0xa5, 0x6f, 0x96, 0xf9, 0xd5, 0x1b, 0x36, 0x32, 0x80, 0x56, 0x9c, 0xcc, 0x25, 0x37,
	0xc2, 0xcb, 0x7b, 0x70, 0x7f, 0x26, 0x16, 0x1c, 0x45, 0x5a, 0x84, 0x4c, 0xaf, 0xac, 0xfa, 0xb0,
	0x32, 0x6a, 0xd1, 0xe2, 0x48, 0x9e, 0x43, 0x3f, 0x50, 0xcc, 0x9b, 0xb3, 0x80, 0x49, 0x97, 0xeb,
	0xfc, 0xc2, 0x48, 0xa9, 0xc0, 0x6a, 0x0c, 0x6b, 0xa3, 0x36, 0x3d, 0x04, 0x93, 0x67, 0x00, 0x2c,
	0x08, 0x94, 0x8b, 0x59, 0x58, 0x4d, 0x1c, 0xe6, 0x3e, 0x36, 0x6c, 0x3a, 0x3b, 0xbf, 0x07, 0x2e,
	0xb9, 0x59, 0x2a, 0x8f, 0x96, 0xa8, 0xe4, 0x0c, 0x8e, 0x7c, 0x66, 0xf8, 0x1d, 0x5b, 0x59, 0x47,
	0xa5, 0x74, 0x0b, 0x23, 0x99, 0xc1, 0x63, 0xb1, 0x51, 0xca, 0xd8, 0x6a, 0xe1, 0x38, 0x7c, 0x78,
	0x68, 0x1c, 0x36, 0x2b, 0x4f, 0xb7, 0xdd, 0xc9, 0x13, 0xa8, 0x1b, 0xe6, 0xc7, 0x56, 0x1b, 0x9b,
	0xdc, 0x46, 0x99, 0x17, 0xcc, 0x8f, 0x29, 0x9a, 0xc9, 0x04, 0x3a, 0xb9, 0x12, 0xee, 0x25, 0x60,
	0x2a, 0x3d, 0x64, 0x5d, 0xad, 0xed, 0xb4, 0x4c, 0x22, 0xbf, 0x40, 0x3f, 0x3f, 0x5e, 0x73, 0x37,
	0xd1, 0xc2, 0xac, 0xbe, 0xd5, 0x2a, 0x89, 0x28, 0x5f, 0x58, 0xc7, 0x78, 0x8b, 0x5d, 0xf6, 0xdf,
	0xe6, 0x70, 0xcd, 0xa5, 0xcb, 0xe9, 0x21, 0x09, 0xf2, 0x31, 0x34, 0xb3, 0xde, 0x59, 0x27, 0x28,
	0x76, 0x8a, 0x62, 0xd7, 0x68, 0x5a, 0xbb, 0xe7, 0x1c, 0x72, 0x09, 0x24, 0x4a, 0xe6, 0x81, 0x70,
	0xa7, 0xb3, 0xf3, 0xac, 0x41, 0x69, 0x18, 0x5d, 0xf4, 0x7c, 0x82, 0x9e, 0xb3, 0x1d, 0x38, 0x97,
	0xd8, 0xe3, 0xf8, 0x5d, 0xbd, 0xd5, 0xe9, 0x1d, 0xdb, 0xff, 0xd5, 0xa1, 0xb7, 0xbd, 0x3f, 0xfb,
	0x66, 0x9a, 0x74, 0xa1, 0x2a, 0xbc, 0x7c, 0x22, 0xab, 0xc2, 0x23, 0x3f, 0x42, 0xdd, 0xa4, 0x65,
	0xac, 0x61, 0x19, 0xbf, 0x7c, 0xe8, 0x0a, 0xef, 0x18, 0xb0, 0xe6, 0x28, 0x45, 0xae, 0xa1, 0x27,
	0x36, 0xb7, 0x26, 0xb6, 0xea, 0x38, 0x12, 0x1f, 0x1d, 0x1c, 0x89, 0x4d, 0x3e, 0xdd, 0x11, 0x20,
	0x4f, 0x01, 0x42, 0xe6, 0x16, 0x8b, 0xd7, 0x28, 0x4d, 0x62, 0xc9, 0x4e, 0x86, 0x50, 0xf3, 0x64,
	0x8c, 0xe3, 0xdd, 0x99, 0xb4, 0xf0, 0xb6, 0xaf, 0x65, 0x9c, 0x13, 0x53, 0x28, 0xdd, 0x3b, 0xa9,
	0x3c, 0x7e, 0x95, 0xd6, 0xe5, 0x28, 0xdb, 0xbb, 0xe2, 0x4c, 0xde, 0x85, 0xb6, 0x9f, 0xf6, 0x14,
	0xc1, 0x16, 0x82, 0x6b, 0x43, 0xba, 0xd5, 0xc5, 0x52, 0x20, 0xa1, 0x9d, 0x6d, 0x75, 0xd9, 0x46,
	0xde, 0x87, 0x66, 0x6c, 0x98, 0x49, 0x62, 0x1c, 0xcb, 0xce, 0xa4, 0x93, 0x4d, 0x02, 0x9a, 0x68,
	0x0e, 0x11, 0x07, 0xc8, 0xad, 0xd0, 0x26, 0x61, 0xc1, 0x25, 0x73, 0x97, 0x42, 0x66, 0xc1, 0x74,
	0x50, 0x6e, 0x0f, 0x92, 0x86, 0x25, 0xd4, 0xed, 0x4f, 0x5c, 0xf8, 0x4b, 0x83, 0xe3, 0x7a, 0x42,
	0xd7, 0x86, 0xfb, 0x6d, 0x39, 0xd9, 0xbb, 0x2d, 0xf6, 0x27, 0x70, 0xba, 0xaf, 0x55, 0xa4, 0x0d,
	0x8d, 0xef, 0x95, 0xcb, 0x82, 0xde, 0x23, 0x02, 0xe9, 0xb3, 0x1a, 0x2a, 0xc3, 0x7b, 0x95, 0xc9,
	0x1f, 0x55, 0x78, 0x7b, 0x9b, 0x7f, 0x9e, 0xb6, 0x8a, 0x84, 0xd0, 0x9c, 0xca, 0x5b, 0x75, 0xc3,
	0xc9, 0xf8, 0xc1, 0xef, 0x7c, 0xf6, 0xcd, 0x18, 0x7c, 0xfa, 0x70, 0x87, 0xec, 0xa5, 0xb7, 0x1f,
	0x91, 0xdf, 0xa1, 0x55, 0xbc, 0xff, 0xe4, 0xd9, 0x43, 0xfd, 0xb7, 0x3e, 0x56, 0x83, 0xe7, 0xaf,
	0xef, 0x58, 0x04, 0x70, 0xf1, 0xd9, 0xcf, 0x63, 0x5f, 0x98, 0x65, 0x32, 0x77, 0x5c, 0x15, 0x8e,
	0x43, 0xe1, 0x6a, 0x15, 0xab, 0x85, 0x19, 0x87, 0xca, 0x1d, 0xeb, 0xc8, 0x1d, 0xaf, 0x55, 0xc7,
	0xb9, 0xea, 0xbc, 0x89, 0xdf, 0xa1, 0xcf, 0xff, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x79, 0xdc,
	0x25, 0x48, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NetworkInterfaceAgentClient is the client API for NetworkInterfaceAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkInterfaceAgentClient interface {
	Invoke(ctx context.Context, in *NetworkInterfaceRequest, opts ...grpc.CallOption) (*NetworkInterfaceResponse, error)
	// Prechecks whether the system is able to create specified network interfaces (but does not actually create them).
	Precheck(ctx context.Context, in *NetworkInterfacePrecheckRequest, opts ...grpc.CallOption) (*NetworkInterfacePrecheckResponse, error)
}

type networkInterfaceAgentClient struct {
	cc *grpc.ClientConn
}

func NewNetworkInterfaceAgentClient(cc *grpc.ClientConn) NetworkInterfaceAgentClient {
	return &networkInterfaceAgentClient{cc}
}

func (c *networkInterfaceAgentClient) Invoke(ctx context.Context, in *NetworkInterfaceRequest, opts ...grpc.CallOption) (*NetworkInterfaceResponse, error) {
	out := new(NetworkInterfaceResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.network.NetworkInterfaceAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkInterfaceAgentClient) Precheck(ctx context.Context, in *NetworkInterfacePrecheckRequest, opts ...grpc.CallOption) (*NetworkInterfacePrecheckResponse, error) {
	out := new(NetworkInterfacePrecheckResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.network.NetworkInterfaceAgent/Precheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkInterfaceAgentServer is the server API for NetworkInterfaceAgent service.
type NetworkInterfaceAgentServer interface {
	Invoke(context.Context, *NetworkInterfaceRequest) (*NetworkInterfaceResponse, error)
	// Prechecks whether the system is able to create specified network interfaces (but does not actually create them).
	Precheck(context.Context, *NetworkInterfacePrecheckRequest) (*NetworkInterfacePrecheckResponse, error)
}

// UnimplementedNetworkInterfaceAgentServer can be embedded to have forward compatible implementations.
type UnimplementedNetworkInterfaceAgentServer struct {
}

func (*UnimplementedNetworkInterfaceAgentServer) Invoke(ctx context.Context, req *NetworkInterfaceRequest) (*NetworkInterfaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}
func (*UnimplementedNetworkInterfaceAgentServer) Precheck(ctx context.Context, req *NetworkInterfacePrecheckRequest) (*NetworkInterfacePrecheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Precheck not implemented")
}

func RegisterNetworkInterfaceAgentServer(s *grpc.Server, srv NetworkInterfaceAgentServer) {
	s.RegisterService(&_NetworkInterfaceAgent_serviceDesc, srv)
}

func _NetworkInterfaceAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkInterfaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkInterfaceAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.network.NetworkInterfaceAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkInterfaceAgentServer).Invoke(ctx, req.(*NetworkInterfaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkInterfaceAgent_Precheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkInterfacePrecheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkInterfaceAgentServer).Precheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.network.NetworkInterfaceAgent/Precheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkInterfaceAgentServer).Precheck(ctx, req.(*NetworkInterfacePrecheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkInterfaceAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.cloudagent.network.NetworkInterfaceAgent",
	HandlerType: (*NetworkInterfaceAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _NetworkInterfaceAgent_Invoke_Handler,
		},
		{
			MethodName: "Precheck",
			Handler:    _NetworkInterfaceAgent_Precheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moc_cloudagent_networkinterface.proto",
}
