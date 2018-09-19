// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metadata/types/pilot.proto

package pbtypes

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PilotConfig struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Host                 string   `protobuf:"bytes,2,opt,name=host,proto3" json:"host"`
	ListenPort           int32    `protobuf:"varint,4,opt,name=listen_port,json=listenPort,proto3" json:"listen_port"`
	LogLevel             string   `protobuf:"bytes,5,opt,name=log_level,json=logLevel,proto3" json:"log_level"`
	TlsListenPort        int32    `protobuf:"varint,6,opt,name=tls_listen_port,json=tlsListenPort,proto3" json:"tls_listen_port"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PilotConfig) Reset()         { *m = PilotConfig{} }
func (m *PilotConfig) String() string { return proto.CompactTextString(m) }
func (*PilotConfig) ProtoMessage()    {}
func (*PilotConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b0b18e18f781d16, []int{0}
}

func (m *PilotConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PilotConfig.Unmarshal(m, b)
}
func (m *PilotConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PilotConfig.Marshal(b, m, deterministic)
}
func (m *PilotConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PilotConfig.Merge(m, src)
}
func (m *PilotConfig) XXX_Size() int {
	return xxx_messageInfo_PilotConfig.Size(m)
}
func (m *PilotConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PilotConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PilotConfig proto.InternalMessageInfo

func (m *PilotConfig) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PilotConfig) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *PilotConfig) GetListenPort() int32 {
	if m != nil {
		return m.ListenPort
	}
	return 0
}

func (m *PilotConfig) GetLogLevel() string {
	if m != nil {
		return m.LogLevel
	}
	return ""
}

func (m *PilotConfig) GetTlsListenPort() int32 {
	if m != nil {
		return m.TlsListenPort
	}
	return 0
}

type PilotTLSConfig struct {
	CaCrtData            string   `protobuf:"bytes,1,opt,name=ca_crt_data,json=caCrtData,proto3" json:"ca_crt_data"`
	ServerCrtData        string   `protobuf:"bytes,2,opt,name=server_crt_data,json=serverCrtData,proto3" json:"server_crt_data"`
	ServerKeyData        string   `protobuf:"bytes,3,opt,name=server_key_data,json=serverKeyData,proto3" json:"server_key_data"`
	ClientCrtData        string   `protobuf:"bytes,4,opt,name=client_crt_data,json=clientCrtData,proto3" json:"client_crt_data"`
	ClientKeyData        string   `protobuf:"bytes,5,opt,name=client_key_data,json=clientKeyData,proto3" json:"client_key_data"`
	PilotServerName      string   `protobuf:"bytes,6,opt,name=pilot_server_name,json=pilotServerName,proto3" json:"pilot_server_name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PilotTLSConfig) Reset()         { *m = PilotTLSConfig{} }
func (m *PilotTLSConfig) String() string { return proto.CompactTextString(m) }
func (*PilotTLSConfig) ProtoMessage()    {}
func (*PilotTLSConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b0b18e18f781d16, []int{1}
}

func (m *PilotTLSConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PilotTLSConfig.Unmarshal(m, b)
}
func (m *PilotTLSConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PilotTLSConfig.Marshal(b, m, deterministic)
}
func (m *PilotTLSConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PilotTLSConfig.Merge(m, src)
}
func (m *PilotTLSConfig) XXX_Size() int {
	return xxx_messageInfo_PilotTLSConfig.Size(m)
}
func (m *PilotTLSConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PilotTLSConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PilotTLSConfig proto.InternalMessageInfo

func (m *PilotTLSConfig) GetCaCrtData() string {
	if m != nil {
		return m.CaCrtData
	}
	return ""
}

func (m *PilotTLSConfig) GetServerCrtData() string {
	if m != nil {
		return m.ServerCrtData
	}
	return ""
}

func (m *PilotTLSConfig) GetServerKeyData() string {
	if m != nil {
		return m.ServerKeyData
	}
	return ""
}

func (m *PilotTLSConfig) GetClientCrtData() string {
	if m != nil {
		return m.ClientCrtData
	}
	return ""
}

func (m *PilotTLSConfig) GetClientKeyData() string {
	if m != nil {
		return m.ClientKeyData
	}
	return ""
}

func (m *PilotTLSConfig) GetPilotServerName() string {
	if m != nil {
		return m.PilotServerName
	}
	return ""
}

type PilotClientTLSConfig struct {
	CaCrtData            string   `protobuf:"bytes,1,opt,name=ca_crt_data,json=caCrtData,proto3" json:"ca_crt_data"`
	ClientCrtData        string   `protobuf:"bytes,2,opt,name=client_crt_data,json=clientCrtData,proto3" json:"client_crt_data"`
	ClientKeyData        string   `protobuf:"bytes,3,opt,name=client_key_data,json=clientKeyData,proto3" json:"client_key_data"`
	PilotServerName      string   `protobuf:"bytes,4,opt,name=pilot_server_name,json=pilotServerName,proto3" json:"pilot_server_name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PilotClientTLSConfig) Reset()         { *m = PilotClientTLSConfig{} }
func (m *PilotClientTLSConfig) String() string { return proto.CompactTextString(m) }
func (*PilotClientTLSConfig) ProtoMessage()    {}
func (*PilotClientTLSConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b0b18e18f781d16, []int{2}
}

func (m *PilotClientTLSConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PilotClientTLSConfig.Unmarshal(m, b)
}
func (m *PilotClientTLSConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PilotClientTLSConfig.Marshal(b, m, deterministic)
}
func (m *PilotClientTLSConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PilotClientTLSConfig.Merge(m, src)
}
func (m *PilotClientTLSConfig) XXX_Size() int {
	return xxx_messageInfo_PilotClientTLSConfig.Size(m)
}
func (m *PilotClientTLSConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PilotClientTLSConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PilotClientTLSConfig proto.InternalMessageInfo

func (m *PilotClientTLSConfig) GetCaCrtData() string {
	if m != nil {
		return m.CaCrtData
	}
	return ""
}

func (m *PilotClientTLSConfig) GetClientCrtData() string {
	if m != nil {
		return m.ClientCrtData
	}
	return ""
}

func (m *PilotClientTLSConfig) GetClientKeyData() string {
	if m != nil {
		return m.ClientKeyData
	}
	return ""
}

func (m *PilotClientTLSConfig) GetPilotServerName() string {
	if m != nil {
		return m.PilotServerName
	}
	return ""
}

type PilotEndpoint struct {
	PilotId              string   `protobuf:"bytes,1,opt,name=pilot_id,json=pilotId,proto3" json:"pilot_id"`
	PilotHost            string   `protobuf:"bytes,2,opt,name=pilot_host,json=pilotHost,proto3" json:"pilot_host"`
	PilotPort            int32    `protobuf:"varint,3,opt,name=pilot_port,json=pilotPort,proto3" json:"pilot_port"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PilotEndpoint) Reset()         { *m = PilotEndpoint{} }
func (m *PilotEndpoint) String() string { return proto.CompactTextString(m) }
func (*PilotEndpoint) ProtoMessage()    {}
func (*PilotEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b0b18e18f781d16, []int{3}
}

func (m *PilotEndpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PilotEndpoint.Unmarshal(m, b)
}
func (m *PilotEndpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PilotEndpoint.Marshal(b, m, deterministic)
}
func (m *PilotEndpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PilotEndpoint.Merge(m, src)
}
func (m *PilotEndpoint) XXX_Size() int {
	return xxx_messageInfo_PilotEndpoint.Size(m)
}
func (m *PilotEndpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_PilotEndpoint.DiscardUnknown(m)
}

var xxx_messageInfo_PilotEndpoint proto.InternalMessageInfo

func (m *PilotEndpoint) GetPilotId() string {
	if m != nil {
		return m.PilotId
	}
	return ""
}

func (m *PilotEndpoint) GetPilotHost() string {
	if m != nil {
		return m.PilotHost
	}
	return ""
}

func (m *PilotEndpoint) GetPilotPort() int32 {
	if m != nil {
		return m.PilotPort
	}
	return 0
}

func init() {
	proto.RegisterType((*PilotConfig)(nil), "metadata.types.PilotConfig")
	proto.RegisterType((*PilotTLSConfig)(nil), "metadata.types.PilotTLSConfig")
	proto.RegisterType((*PilotClientTLSConfig)(nil), "metadata.types.PilotClientTLSConfig")
	proto.RegisterType((*PilotEndpoint)(nil), "metadata.types.PilotEndpoint")
}

func init() { proto.RegisterFile("metadata/types/pilot.proto", fileDescriptor_9b0b18e18f781d16) }

var fileDescriptor_9b0b18e18f781d16 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4f, 0xef, 0xd2, 0x40,
	0x14, 0x4c, 0x4b, 0x41, 0xfa, 0x08, 0x10, 0x37, 0x1e, 0xaa, 0x46, 0x25, 0x1c, 0x0c, 0xf1, 0x40,
	0x0f, 0x26, 0xc6, 0xc4, 0x9b, 0x68, 0xa2, 0x91, 0x18, 0x02, 0x9e, 0xbc, 0x34, 0x4b, 0xbb, 0x96,
	0x0d, 0x4b, 0x77, 0xb3, 0x7d, 0x21, 0xf2, 0x49, 0xfc, 0x1e, 0x7e, 0x3e, 0x0f, 0xa6, 0x6f, 0x4b,
	0xe9, 0xef, 0xcf, 0x81, 0xdf, 0xa9, 0x7d, 0x33, 0xb3, 0xd3, 0xd9, 0xee, 0x2c, 0x3c, 0x3b, 0x08,
	0xe4, 0x19, 0x47, 0x1e, 0xe3, 0xc9, 0x88, 0x32, 0x36, 0x52, 0x69, 0x9c, 0x1b, 0xab, 0x51, 0xb3,
	0xd1, 0x99, 0x9b, 0x13, 0x37, 0xfd, 0xe3, 0xc1, 0x60, 0x55, 0xf1, 0x0b, 0x5d, 0xfc, 0x92, 0x39,
	0x1b, 0x81, 0x2f, 0xb3, 0xc8, 0x9b, 0x78, 0xb3, 0x70, 0xed, 0xcb, 0x8c, 0x31, 0x08, 0x76, 0xba,
	0xc4, 0xc8, 0x27, 0x84, 0xde, 0xd9, 0x2b, 0x18, 0x28, 0x59, 0xa2, 0x28, 0x12, 0xa3, 0x2d, 0x46,
	0xc1, 0xc4, 0x9b, 0x75, 0xd7, 0xe0, 0xa0, 0x95, 0xb6, 0xc8, 0x9e, 0x43, 0xa8, 0x74, 0x9e, 0x28,
	0x71, 0x14, 0x2a, 0xea, 0xd2, 0xca, 0xbe, 0xd2, 0xf9, 0xb2, 0x9a, 0xd9, 0x6b, 0x18, 0xa3, 0x2a,
	0x93, 0xb6, 0x43, 0x8f, 0x1c, 0x86, 0xa8, 0xca, 0x65, 0x63, 0x32, 0xfd, 0xe7, 0xc1, 0x88, 0x92,
	0xfd, 0x58, 0x6e, 0xea, 0x70, 0x2f, 0x61, 0x90, 0xf2, 0x24, 0xb5, 0x98, 0x54, 0x3b, 0xa8, 0x53,
	0x86, 0x29, 0x5f, 0x58, 0xfc, 0xc4, 0x91, 0x57, 0xd6, 0xa5, 0xb0, 0x47, 0x61, 0x2f, 0x1a, 0x97,
	0x7b, 0xe8, 0xe0, 0xbb, 0xba, 0xbd, 0x38, 0x39, 0x5d, 0xa7, 0xad, 0xfb, 0x26, 0x4e, 0x67, 0x5d,
	0xaa, 0xa4, 0x28, 0xf0, 0xe2, 0x17, 0x38, 0x9d, 0x83, 0x5b, 0x7e, 0xb5, 0xae, 0xf1, 0xeb, 0xb6,
	0x75, 0x67, 0xbf, 0x37, 0xf0, 0x98, 0xce, 0x22, 0xa9, 0xbf, 0x5e, 0xf0, 0x83, 0xa0, 0xcd, 0x87,
	0xeb, 0x31, 0x11, 0x1b, 0xc2, 0xbf, 0xf3, 0x83, 0x98, 0xfe, 0xf5, 0xe0, 0x89, 0x3b, 0x18, 0xb2,
	0x78, 0xd0, 0x4f, 0xb8, 0x1d, 0xda, 0xbf, 0x32, 0x74, 0xe7, 0xea, 0xd0, 0xc1, 0xfd, 0xa1, 0x77,
	0x30, 0xa4, 0xcc, 0x9f, 0x8b, 0xcc, 0x68, 0x59, 0x20, 0x7b, 0x0a, 0x7d, 0xb7, 0xb8, 0x29, 0xd5,
	0x23, 0x9a, 0xbf, 0x66, 0xec, 0x05, 0x80, 0xa3, 0x5a, 0xfd, 0x0a, 0x09, 0xf9, 0x52, 0x95, 0xac,
	0xa1, 0xa9, 0x21, 0x1d, 0x6a, 0x88, 0xa3, 0xab, 0x76, 0x7c, 0x7c, 0xff, 0xf3, 0x9d, 0x36, 0xa2,
	0x30, 0x12, 0xad, 0xfc, 0x3d, 0x97, 0x3a, 0xbe, 0x4c, 0xb1, 0xd9, 0xe7, 0xb1, 0xd9, 0xc6, 0x37,
	0x6f, 0xc1, 0x07, 0xb3, 0xa5, 0xe7, 0xb6, 0x47, 0x17, 0xe1, 0xed, 0xff, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xc6, 0x22, 0x40, 0x2c, 0x26, 0x03, 0x00, 0x00,
}
