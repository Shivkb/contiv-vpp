// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nat.proto

/*
Package nat is a generated protocol buffer package.

It is generated from these files:
	nat.proto

It has these top-level messages:
	Nat44Global
	Nat44SNat
	Nat44DNat
*/
package nat

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Protocol int32

const (
	Protocol_TCP  Protocol = 0
	Protocol_UDP  Protocol = 1
	Protocol_ICMP Protocol = 2
)

var Protocol_name = map[int32]string{
	0: "TCP",
	1: "UDP",
	2: "ICMP",
}
var Protocol_value = map[string]int32{
	"TCP":  0,
	"UDP":  1,
	"ICMP": 2,
}

func (x Protocol) String() string {
	return proto.EnumName(Protocol_name, int32(x))
}
func (Protocol) EnumDescriptor() ([]byte, []int) { return fileDescriptorNat, []int{0} }

type TwiceNatMode int32

const (
	TwiceNatMode_DISABLED TwiceNatMode = 0
	TwiceNatMode_ENABLED  TwiceNatMode = 1
	TwiceNatMode_SELF     TwiceNatMode = 2
)

var TwiceNatMode_name = map[int32]string{
	0: "DISABLED",
	1: "ENABLED",
	2: "SELF",
}
var TwiceNatMode_value = map[string]int32{
	"DISABLED": 0,
	"ENABLED":  1,
	"SELF":     2,
}

func (x TwiceNatMode) String() string {
	return proto.EnumName(TwiceNatMode_name, int32(x))
}
func (TwiceNatMode) EnumDescriptor() ([]byte, []int) { return fileDescriptorNat, []int{1} }

// NAT44 global config
type Nat44Global struct {
	Forwarding    bool                        `protobuf:"varint,2,opt,name=forwarding,proto3" json:"forwarding,omitempty"`
	NatInterfaces []*Nat44Global_NatInterface `protobuf:"bytes,3,rep,name=nat_interfaces,json=natInterfaces" json:"nat_interfaces,omitempty"`
	AddressPools  []*Nat44Global_AddressPool  `protobuf:"bytes,5,rep,name=address_pools,json=addressPools" json:"address_pools,omitempty"`
}

func (m *Nat44Global) Reset()                    { *m = Nat44Global{} }
func (m *Nat44Global) String() string            { return proto.CompactTextString(m) }
func (*Nat44Global) ProtoMessage()               {}
func (*Nat44Global) Descriptor() ([]byte, []int) { return fileDescriptorNat, []int{0} }

func (m *Nat44Global) GetForwarding() bool {
	if m != nil {
		return m.Forwarding
	}
	return false
}

func (m *Nat44Global) GetNatInterfaces() []*Nat44Global_NatInterface {
	if m != nil {
		return m.NatInterfaces
	}
	return nil
}

func (m *Nat44Global) GetAddressPools() []*Nat44Global_AddressPool {
	if m != nil {
		return m.AddressPools
	}
	return nil
}

type Nat44Global_NatInterface struct {
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	IsInside      bool   `protobuf:"varint,2,opt,name=is_inside,json=isInside,proto3" json:"is_inside,omitempty"`
	OutputFeature bool   `protobuf:"varint,3,opt,name=output_feature,json=outputFeature,proto3" json:"output_feature,omitempty"`
}

func (m *Nat44Global_NatInterface) Reset()                    { *m = Nat44Global_NatInterface{} }
func (m *Nat44Global_NatInterface) String() string            { return proto.CompactTextString(m) }
func (*Nat44Global_NatInterface) ProtoMessage()               {}
func (*Nat44Global_NatInterface) Descriptor() ([]byte, []int) { return fileDescriptorNat, []int{0, 0} }

func (m *Nat44Global_NatInterface) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Nat44Global_NatInterface) GetIsInside() bool {
	if m != nil {
		return m.IsInside
	}
	return false
}

func (m *Nat44Global_NatInterface) GetOutputFeature() bool {
	if m != nil {
		return m.OutputFeature
	}
	return false
}

type Nat44Global_AddressPool struct {
	FirstSrcAddress string `protobuf:"bytes,1,opt,name=first_src_address,json=firstSrcAddress,proto3" json:"first_src_address,omitempty"`
	LastSrcAddress  string `protobuf:"bytes,2,opt,name=last_src_address,json=lastSrcAddress,proto3" json:"last_src_address,omitempty"`
	VrfId           uint32 `protobuf:"varint,3,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	TwiceNat        bool   `protobuf:"varint,4,opt,name=twice_nat,json=twiceNat,proto3" json:"twice_nat,omitempty"`
}

func (m *Nat44Global_AddressPool) Reset()                    { *m = Nat44Global_AddressPool{} }
func (m *Nat44Global_AddressPool) String() string            { return proto.CompactTextString(m) }
func (*Nat44Global_AddressPool) ProtoMessage()               {}
func (*Nat44Global_AddressPool) Descriptor() ([]byte, []int) { return fileDescriptorNat, []int{0, 1} }

func (m *Nat44Global_AddressPool) GetFirstSrcAddress() string {
	if m != nil {
		return m.FirstSrcAddress
	}
	return ""
}

func (m *Nat44Global_AddressPool) GetLastSrcAddress() string {
	if m != nil {
		return m.LastSrcAddress
	}
	return ""
}

func (m *Nat44Global_AddressPool) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *Nat44Global_AddressPool) GetTwiceNat() bool {
	if m != nil {
		return m.TwiceNat
	}
	return false
}

// Many-to-one (SNAT) setup
type Nat44SNat struct {
	SnatConfigs []*Nat44SNat_SNatConfig `protobuf:"bytes,1,rep,name=snat_configs,json=snatConfigs" json:"snat_configs,omitempty"`
}

func (m *Nat44SNat) Reset()                    { *m = Nat44SNat{} }
func (m *Nat44SNat) String() string            { return proto.CompactTextString(m) }
func (*Nat44SNat) ProtoMessage()               {}
func (*Nat44SNat) Descriptor() ([]byte, []int) { return fileDescriptorNat, []int{1} }

func (m *Nat44SNat) GetSnatConfigs() []*Nat44SNat_SNatConfig {
	if m != nil {
		return m.SnatConfigs
	}
	return nil
}

type Nat44SNat_SNatConfig struct {
	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
}

func (m *Nat44SNat_SNatConfig) Reset()                    { *m = Nat44SNat_SNatConfig{} }
func (m *Nat44SNat_SNatConfig) String() string            { return proto.CompactTextString(m) }
func (*Nat44SNat_SNatConfig) ProtoMessage()               {}
func (*Nat44SNat_SNatConfig) Descriptor() ([]byte, []int) { return fileDescriptorNat, []int{1, 0} }

func (m *Nat44SNat_SNatConfig) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

// One-to-many (DNAT) setup
type Nat44DNat struct {
	DnatConfigs []*Nat44DNat_DNatConfig `protobuf:"bytes,1,rep,name=dnat_configs,json=dnatConfigs" json:"dnat_configs,omitempty"`
}

func (m *Nat44DNat) Reset()                    { *m = Nat44DNat{} }
func (m *Nat44DNat) String() string            { return proto.CompactTextString(m) }
func (*Nat44DNat) ProtoMessage()               {}
func (*Nat44DNat) Descriptor() ([]byte, []int) { return fileDescriptorNat, []int{2} }

func (m *Nat44DNat) GetDnatConfigs() []*Nat44DNat_DNatConfig {
	if m != nil {
		return m.DnatConfigs
	}
	return nil
}

type Nat44DNat_DNatConfig struct {
	Label      string                                  `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	StMappings []*Nat44DNat_DNatConfig_StaticMapping   `protobuf:"bytes,4,rep,name=st_mappings,json=stMappings" json:"st_mappings,omitempty"`
	IdMappings []*Nat44DNat_DNatConfig_IdentityMapping `protobuf:"bytes,6,rep,name=id_mappings,json=idMappings" json:"id_mappings,omitempty"`
}

func (m *Nat44DNat_DNatConfig) Reset()                    { *m = Nat44DNat_DNatConfig{} }
func (m *Nat44DNat_DNatConfig) String() string            { return proto.CompactTextString(m) }
func (*Nat44DNat_DNatConfig) ProtoMessage()               {}
func (*Nat44DNat_DNatConfig) Descriptor() ([]byte, []int) { return fileDescriptorNat, []int{2, 0} }

func (m *Nat44DNat_DNatConfig) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Nat44DNat_DNatConfig) GetStMappings() []*Nat44DNat_DNatConfig_StaticMapping {
	if m != nil {
		return m.StMappings
	}
	return nil
}

func (m *Nat44DNat_DNatConfig) GetIdMappings() []*Nat44DNat_DNatConfig_IdentityMapping {
	if m != nil {
		return m.IdMappings
	}
	return nil
}

type Nat44DNat_DNatConfig_StaticMapping struct {
	VrfId             uint32                                        `protobuf:"varint,1,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	ExternalInterface string                                        `protobuf:"bytes,2,opt,name=external_interface,json=externalInterface,proto3" json:"external_interface,omitempty"`
	ExternalIp        string                                        `protobuf:"bytes,3,opt,name=external_ip,json=externalIp,proto3" json:"external_ip,omitempty"`
	ExternalPort      uint32                                        `protobuf:"varint,4,opt,name=external_port,json=externalPort,proto3" json:"external_port,omitempty"`
	LocalIps          []*Nat44DNat_DNatConfig_StaticMapping_LocalIP `protobuf:"bytes,5,rep,name=local_ips,json=localIps" json:"local_ips,omitempty"`
	Protocol          Protocol                                      `protobuf:"varint,6,opt,name=protocol,proto3,enum=nat.Protocol" json:"protocol,omitempty"`
	TwiceNat          TwiceNatMode                                  `protobuf:"varint,7,opt,name=twice_nat,json=twiceNat,proto3,enum=nat.TwiceNatMode" json:"twice_nat,omitempty"`
}

func (m *Nat44DNat_DNatConfig_StaticMapping) Reset()         { *m = Nat44DNat_DNatConfig_StaticMapping{} }
func (m *Nat44DNat_DNatConfig_StaticMapping) String() string { return proto.CompactTextString(m) }
func (*Nat44DNat_DNatConfig_StaticMapping) ProtoMessage()    {}
func (*Nat44DNat_DNatConfig_StaticMapping) Descriptor() ([]byte, []int) {
	return fileDescriptorNat, []int{2, 0, 0}
}

func (m *Nat44DNat_DNatConfig_StaticMapping) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *Nat44DNat_DNatConfig_StaticMapping) GetExternalInterface() string {
	if m != nil {
		return m.ExternalInterface
	}
	return ""
}

func (m *Nat44DNat_DNatConfig_StaticMapping) GetExternalIp() string {
	if m != nil {
		return m.ExternalIp
	}
	return ""
}

func (m *Nat44DNat_DNatConfig_StaticMapping) GetExternalPort() uint32 {
	if m != nil {
		return m.ExternalPort
	}
	return 0
}

func (m *Nat44DNat_DNatConfig_StaticMapping) GetLocalIps() []*Nat44DNat_DNatConfig_StaticMapping_LocalIP {
	if m != nil {
		return m.LocalIps
	}
	return nil
}

func (m *Nat44DNat_DNatConfig_StaticMapping) GetProtocol() Protocol {
	if m != nil {
		return m.Protocol
	}
	return Protocol_TCP
}

func (m *Nat44DNat_DNatConfig_StaticMapping) GetTwiceNat() TwiceNatMode {
	if m != nil {
		return m.TwiceNat
	}
	return TwiceNatMode_DISABLED
}

type Nat44DNat_DNatConfig_StaticMapping_LocalIP struct {
	LocalIp     string `protobuf:"bytes,1,opt,name=local_ip,json=localIp,proto3" json:"local_ip,omitempty"`
	LocalPort   uint32 `protobuf:"varint,3,opt,name=local_port,json=localPort,proto3" json:"local_port,omitempty"`
	Probability uint32 `protobuf:"varint,2,opt,name=probability,proto3" json:"probability,omitempty"`
}

func (m *Nat44DNat_DNatConfig_StaticMapping_LocalIP) Reset() {
	*m = Nat44DNat_DNatConfig_StaticMapping_LocalIP{}
}
func (m *Nat44DNat_DNatConfig_StaticMapping_LocalIP) String() string {
	return proto.CompactTextString(m)
}
func (*Nat44DNat_DNatConfig_StaticMapping_LocalIP) ProtoMessage() {}
func (*Nat44DNat_DNatConfig_StaticMapping_LocalIP) Descriptor() ([]byte, []int) {
	return fileDescriptorNat, []int{2, 0, 0, 0}
}

func (m *Nat44DNat_DNatConfig_StaticMapping_LocalIP) GetLocalIp() string {
	if m != nil {
		return m.LocalIp
	}
	return ""
}

func (m *Nat44DNat_DNatConfig_StaticMapping_LocalIP) GetLocalPort() uint32 {
	if m != nil {
		return m.LocalPort
	}
	return 0
}

func (m *Nat44DNat_DNatConfig_StaticMapping_LocalIP) GetProbability() uint32 {
	if m != nil {
		return m.Probability
	}
	return 0
}

type Nat44DNat_DNatConfig_IdentityMapping struct {
	VrfId              uint32   `protobuf:"varint,1,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	AddressedInterface string   `protobuf:"bytes,2,opt,name=addressed_interface,json=addressedInterface,proto3" json:"addressed_interface,omitempty"`
	IpAddress          string   `protobuf:"bytes,3,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	Port               uint32   `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	Protocol           Protocol `protobuf:"varint,5,opt,name=protocol,proto3,enum=nat.Protocol" json:"protocol,omitempty"`
}

func (m *Nat44DNat_DNatConfig_IdentityMapping) Reset()         { *m = Nat44DNat_DNatConfig_IdentityMapping{} }
func (m *Nat44DNat_DNatConfig_IdentityMapping) String() string { return proto.CompactTextString(m) }
func (*Nat44DNat_DNatConfig_IdentityMapping) ProtoMessage()    {}
func (*Nat44DNat_DNatConfig_IdentityMapping) Descriptor() ([]byte, []int) {
	return fileDescriptorNat, []int{2, 0, 1}
}

func (m *Nat44DNat_DNatConfig_IdentityMapping) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *Nat44DNat_DNatConfig_IdentityMapping) GetAddressedInterface() string {
	if m != nil {
		return m.AddressedInterface
	}
	return ""
}

func (m *Nat44DNat_DNatConfig_IdentityMapping) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *Nat44DNat_DNatConfig_IdentityMapping) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Nat44DNat_DNatConfig_IdentityMapping) GetProtocol() Protocol {
	if m != nil {
		return m.Protocol
	}
	return Protocol_TCP
}

func init() {
	proto.RegisterType((*Nat44Global)(nil), "nat.Nat44Global")
	proto.RegisterType((*Nat44Global_NatInterface)(nil), "nat.Nat44Global.NatInterface")
	proto.RegisterType((*Nat44Global_AddressPool)(nil), "nat.Nat44Global.AddressPool")
	proto.RegisterType((*Nat44SNat)(nil), "nat.Nat44SNat")
	proto.RegisterType((*Nat44SNat_SNatConfig)(nil), "nat.Nat44SNat.SNatConfig")
	proto.RegisterType((*Nat44DNat)(nil), "nat.Nat44DNat")
	proto.RegisterType((*Nat44DNat_DNatConfig)(nil), "nat.Nat44DNat.DNatConfig")
	proto.RegisterType((*Nat44DNat_DNatConfig_StaticMapping)(nil), "nat.Nat44DNat.DNatConfig.StaticMapping")
	proto.RegisterType((*Nat44DNat_DNatConfig_StaticMapping_LocalIP)(nil), "nat.Nat44DNat.DNatConfig.StaticMapping.LocalIP")
	proto.RegisterType((*Nat44DNat_DNatConfig_IdentityMapping)(nil), "nat.Nat44DNat.DNatConfig.IdentityMapping")
	proto.RegisterEnum("nat.Protocol", Protocol_name, Protocol_value)
	proto.RegisterEnum("nat.TwiceNatMode", TwiceNatMode_name, TwiceNatMode_value)
}

func init() { proto.RegisterFile("nat.proto", fileDescriptorNat) }

var fileDescriptorNat = []byte{
	// 720 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x4e, 0xdb, 0x48,
	0x14, 0xc6, 0xf9, 0x75, 0x8e, 0xe3, 0x10, 0x66, 0x77, 0x25, 0x93, 0x5d, 0x76, 0xa3, 0xac, 0xda,
	0x06, 0xa4, 0x06, 0x09, 0xb8, 0xec, 0x0d, 0x25, 0xd0, 0xba, 0x0a, 0x51, 0xe4, 0xd0, 0x6b, 0x6b,
	0x62, 0x8f, 0xd1, 0x48, 0xc6, 0x63, 0xcd, 0x0c, 0x50, 0x9e, 0xa4, 0xb7, 0xbd, 0xef, 0x03, 0xb4,
	0x0f, 0xd5, 0x87, 0xa8, 0x3c, 0x1e, 0xc7, 0x06, 0x04, 0xed, 0x4d, 0x74, 0xe6, 0x9c, 0xcf, 0xdf,
	0x77, 0x7e, 0x03, 0x9d, 0x04, 0xcb, 0x49, 0xca, 0x99, 0x64, 0xa8, 0x9e, 0x60, 0x39, 0xfa, 0x5e,
	0x07, 0x6b, 0x8e, 0xe5, 0xd1, 0xd1, 0xbb, 0x98, 0xad, 0x70, 0x8c, 0xfe, 0x05, 0x88, 0x18, 0xbf,
	0xc5, 0x3c, 0xa4, 0xc9, 0xa5, 0x53, 0x1b, 0x1a, 0x63, 0xd3, 0xab, 0x78, 0xd0, 0x14, 0x7a, 0x09,
	0x96, 0x3e, 0x4d, 0x24, 0xe1, 0x11, 0x0e, 0x88, 0x70, 0xea, 0xc3, 0xfa, 0xd8, 0x3a, 0xd8, 0x99,
	0x64, 0xc4, 0x15, 0xa6, 0xcc, 0x76, 0x0b, 0x94, 0x67, 0x27, 0x95, 0x97, 0x40, 0xc7, 0x60, 0xe3,
	0x30, 0xe4, 0x44, 0x08, 0x3f, 0x65, 0x2c, 0x16, 0x4e, 0x53, 0x91, 0xfc, 0xf3, 0x88, 0xe4, 0x38,
	0x47, 0x2d, 0x18, 0x8b, 0xbd, 0x2e, 0x2e, 0x1f, 0x62, 0x10, 0x41, 0xb7, 0xaa, 0x80, 0x10, 0x34,
	0x12, 0x7c, 0x45, 0x1c, 0x63, 0x68, 0x8c, 0x3b, 0x9e, 0xb2, 0xd1, 0xdf, 0xd0, 0xa1, 0xc2, 0xa7,
	0x89, 0xa0, 0x21, 0xd1, 0xb5, 0x98, 0x54, 0xb8, 0xea, 0x8d, 0x5e, 0x40, 0x8f, 0x5d, 0xcb, 0xf4,
	0x5a, 0xfa, 0x11, 0xc1, 0xf2, 0x9a, 0x13, 0xa7, 0xae, 0x10, 0x76, 0xee, 0x3d, 0xcb, 0x9d, 0x83,
	0xcf, 0x06, 0x58, 0x95, 0x2c, 0xd0, 0x1e, 0x6c, 0x45, 0x94, 0x0b, 0xe9, 0x0b, 0x1e, 0xf8, 0x3a,
	0x23, 0x2d, 0xba, 0xa9, 0x02, 0x4b, 0x1e, 0x68, 0x3c, 0x1a, 0x43, 0x3f, 0xc6, 0x0f, 0xa0, 0x35,
	0x05, 0xed, 0x65, 0xfe, 0x0a, 0xf2, 0x2f, 0x68, 0xdd, 0xf0, 0xc8, 0xa7, 0xa1, 0x4a, 0xc2, 0xf6,
	0x9a, 0x37, 0x3c, 0x72, 0xc3, 0xac, 0x00, 0x79, 0x4b, 0x03, 0xe2, 0x27, 0x58, 0x3a, 0x8d, 0xbc,
	0x00, 0xe5, 0x98, 0x63, 0x39, 0xba, 0x82, 0x8e, 0x6a, 0xd5, 0x72, 0x8e, 0x25, 0x7a, 0x03, 0x5d,
	0x91, 0x0d, 0x26, 0x60, 0x49, 0x44, 0x2f, 0xb3, 0x8c, 0xb2, 0x86, 0x6e, 0x97, 0x0d, 0xcd, 0x50,
	0x93, 0xec, 0xe7, 0x44, 0x21, 0x3c, 0x2b, 0x83, 0xe7, 0xb6, 0x18, 0x8c, 0x00, 0xca, 0x10, 0xfa,
	0x13, 0x9a, 0x31, 0x5e, 0x91, 0x58, 0x97, 0x95, 0x3f, 0x46, 0x3f, 0x5a, 0x5a, 0x6f, 0xaa, 0xf5,
	0xc2, 0x67, 0xf5, 0x32, 0xd4, 0x64, 0x5a, 0xd1, 0x0b, 0x2b, 0x7a, 0x5f, 0x5a, 0x00, 0xd3, 0x5f,
	0x08, 0xa2, 0xf7, 0x60, 0x09, 0xe9, 0x5f, 0xe1, 0x34, 0xa5, 0xc9, 0xa5, 0x70, 0x1a, 0x4a, 0xe1,
	0xd5, 0x93, 0x0a, 0x93, 0xa5, 0xc4, 0x92, 0x06, 0xe7, 0x39, 0xde, 0x03, 0x21, 0xb5, 0x29, 0xd0,
	0x07, 0xb0, 0x68, 0x58, 0x32, 0xb5, 0x14, 0xd3, 0xee, 0xd3, 0x4c, 0x6e, 0x48, 0x12, 0x49, 0xe5,
	0xdd, 0x9a, 0x8b, 0x86, 0x05, 0xd7, 0xe0, 0x6b, 0x1d, 0xec, 0x7b, 0x4a, 0x95, 0xd9, 0x19, 0xd5,
	0xd9, 0xbd, 0x06, 0x44, 0x3e, 0x49, 0xc2, 0x13, 0x1c, 0x97, 0xe7, 0xa2, 0xc7, 0xbf, 0x55, 0x44,
	0xca, 0xfd, 0xfd, 0x0f, 0xac, 0x12, 0x9e, 0xaa, 0x35, 0xe8, 0x78, 0xb0, 0xc6, 0xa5, 0xe8, 0x7f,
	0xb0, 0xd7, 0x80, 0x94, 0xf1, 0x7c, 0x1f, 0x6c, 0xaf, 0x5b, 0x38, 0x17, 0x8c, 0x4b, 0x34, 0x83,
	0x4e, 0xcc, 0x02, 0x45, 0x51, 0x1c, 0xd5, 0xfe, 0x6f, 0x76, 0x6c, 0x32, 0xcb, 0x3e, 0x74, 0x17,
	0x9e, 0xa9, 0x18, 0xdc, 0x54, 0xa0, 0x5d, 0x30, 0xd5, 0x5f, 0x45, 0xc0, 0x62, 0xa7, 0x35, 0x34,
	0xc6, 0xbd, 0x03, 0x5b, 0x91, 0x2d, 0xb4, 0xd3, 0x5b, 0x87, 0xd1, 0xa4, 0xba, 0xa9, 0x6d, 0x85,
	0xdd, 0x52, 0xd8, 0x0b, 0xbd, 0xae, 0xe7, 0x2c, 0x24, 0xe5, 0xf2, 0x0e, 0x08, 0xb4, 0xb5, 0x1e,
	0xda, 0x06, 0xb3, 0xc8, 0x59, 0x2f, 0x40, 0x5b, 0x67, 0x80, 0x76, 0x00, 0xf2, 0x90, 0x2a, 0x38,
	0x3f, 0x8d, 0xbc, 0x40, 0x55, 0xed, 0x10, 0xac, 0x94, 0xb3, 0x15, 0x5e, 0xd1, 0x98, 0xca, 0x3b,
	0xd5, 0x5b, 0xdb, 0xab, 0xba, 0x06, 0xdf, 0x0c, 0xd8, 0x7c, 0x30, 0xcd, 0xa7, 0xe6, 0xb5, 0x0f,
	0x7f, 0xe8, 0x1b, 0x25, 0xe1, 0xa3, 0x81, 0xa1, 0x75, 0xa8, 0x9c, 0xd8, 0x0e, 0x00, 0x4d, 0xd7,
	0x77, 0x9d, 0x0f, 0xac, 0x43, 0xd3, 0xe2, 0xa4, 0x11, 0x34, 0x2a, 0x63, 0x52, 0xf6, 0xbd, 0x86,
	0x36, 0x9f, 0x6d, 0xe8, 0xde, 0x4b, 0x30, 0x0b, 0x2f, 0x6a, 0x43, 0xfd, 0xe2, 0x64, 0xd1, 0xdf,
	0xc8, 0x8c, 0x8f, 0xd3, 0x45, 0xdf, 0x40, 0x26, 0x34, 0xdc, 0x93, 0xf3, 0x45, 0xbf, 0xb6, 0x77,
	0x08, 0xdd, 0x6a, 0x8b, 0x51, 0x17, 0xcc, 0xa9, 0xbb, 0x3c, 0x7e, 0x3b, 0x3b, 0x9d, 0xf6, 0x37,
	0x90, 0x05, 0xed, 0xd3, 0x79, 0xfe, 0x50, 0x1f, 0x2d, 0x4f, 0x67, 0x67, 0xfd, 0xda, 0xaa, 0xa5,
	0x64, 0x0e, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0x46, 0x7f, 0x74, 0xc4, 0x0e, 0x06, 0x00, 0x00,
}
