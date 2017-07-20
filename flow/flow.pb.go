// Code generated by protoc-gen-go.
// source: flow/flow.proto
// DO NOT EDIT!

/*
Package flow is a generated protocol buffer package.

It is generated from these files:
	flow/flow.proto
	flow/set.proto
	flow/request.proto

It has these top-level messages:
	FlowLayer
	ICMPLayer
	FlowMetric
	Flow
	FlowSet
	FlowSearchReply
*/
package flow

import proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FlowProtocol int32

const (
	FlowProtocol_ETHERNET FlowProtocol = 0
	FlowProtocol_IPV4     FlowProtocol = 1
	FlowProtocol_TCPPORT  FlowProtocol = 2
	FlowProtocol_UDPPORT  FlowProtocol = 3
	FlowProtocol_SCTPPORT FlowProtocol = 4
	FlowProtocol_IPV6     FlowProtocol = 5
	FlowProtocol_ICMPV4   FlowProtocol = 6
	FlowProtocol_ICMPV6   FlowProtocol = 7
)

var FlowProtocol_name = map[int32]string{
	0: "ETHERNET",
	1: "IPV4",
	2: "TCPPORT",
	3: "UDPPORT",
	4: "SCTPPORT",
	5: "IPV6",
	6: "ICMPV4",
	7: "ICMPV6",
}
var FlowProtocol_value = map[string]int32{
	"ETHERNET": 0,
	"IPV4":     1,
	"TCPPORT":  2,
	"UDPPORT":  3,
	"SCTPPORT": 4,
	"IPV6":     5,
	"ICMPV4":   6,
	"ICMPV6":   7,
}

func (x FlowProtocol) String() string {
	return proto.EnumName(FlowProtocol_name, int32(x))
}
func (FlowProtocol) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ICMPType int32

const (
	ICMPType_UNKNOWN                 ICMPType = 0
	ICMPType_DESTINATION_UNREACHABLE ICMPType = 1
	ICMPType_ECHO                    ICMPType = 2
	ICMPType_NEIGHBOR                ICMPType = 3
	ICMPType_ADDRESS_MASK            ICMPType = 4
	ICMPType_INFO                    ICMPType = 5
	ICMPType_PARAMETER_PROBLEM       ICMPType = 6
	ICMPType_REDIRECT                ICMPType = 7
	ICMPType_ROUTER                  ICMPType = 8
	ICMPType_SOURCE_QUENCH           ICMPType = 9
	ICMPType_TIME_EXCEEDED           ICMPType = 10
	ICMPType_TIMESTAMP               ICMPType = 11
	ICMPType_PACKET_TOO_BIG          ICMPType = 12
)

var ICMPType_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "DESTINATION_UNREACHABLE",
	2:  "ECHO",
	3:  "NEIGHBOR",
	4:  "ADDRESS_MASK",
	5:  "INFO",
	6:  "PARAMETER_PROBLEM",
	7:  "REDIRECT",
	8:  "ROUTER",
	9:  "SOURCE_QUENCH",
	10: "TIME_EXCEEDED",
	11: "TIMESTAMP",
	12: "PACKET_TOO_BIG",
}
var ICMPType_value = map[string]int32{
	"UNKNOWN":                 0,
	"DESTINATION_UNREACHABLE": 1,
	"ECHO":              2,
	"NEIGHBOR":          3,
	"ADDRESS_MASK":      4,
	"INFO":              5,
	"PARAMETER_PROBLEM": 6,
	"REDIRECT":          7,
	"ROUTER":            8,
	"SOURCE_QUENCH":     9,
	"TIME_EXCEEDED":     10,
	"TIMESTAMP":         11,
	"PACKET_TOO_BIG":    12,
}

func (x ICMPType) String() string {
	return proto.EnumName(ICMPType_name, int32(x))
}
func (ICMPType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type FlowLayer struct {
	Protocol FlowProtocol `protobuf:"varint,1,opt,name=Protocol,enum=flow.FlowProtocol" json:"Protocol,omitempty"`
	A        string       `protobuf:"bytes,3,opt,name=A" json:"A,omitempty"`
	B        string       `protobuf:"bytes,4,opt,name=B" json:"B,omitempty"`
	ID       int64        `protobuf:"varint,5,opt,name=ID" json:"ID"`
}

func (m *FlowLayer) Reset()                    { *m = FlowLayer{} }
func (m *FlowLayer) String() string            { return proto.CompactTextString(m) }
func (*FlowLayer) ProtoMessage()               {}
func (*FlowLayer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FlowLayer) GetProtocol() FlowProtocol {
	if m != nil {
		return m.Protocol
	}
	return FlowProtocol_ETHERNET
}

func (m *FlowLayer) GetA() string {
	if m != nil {
		return m.A
	}
	return ""
}

func (m *FlowLayer) GetB() string {
	if m != nil {
		return m.B
	}
	return ""
}

func (m *FlowLayer) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

type ICMPLayer struct {
	Type ICMPType `protobuf:"varint,1,opt,name=Type,enum=flow.ICMPType" json:"Type,omitempty"`
	Code uint32   `protobuf:"varint,2,opt,name=Code" json:"Code,omitempty"`
	ID   uint32   `protobuf:"varint,3,opt,name=ID" json:"ID,omitempty"`
}

func (m *ICMPLayer) Reset()                    { *m = ICMPLayer{} }
func (m *ICMPLayer) String() string            { return proto.CompactTextString(m) }
func (*ICMPLayer) ProtoMessage()               {}
func (*ICMPLayer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ICMPLayer) GetType() ICMPType {
	if m != nil {
		return m.Type
	}
	return ICMPType_UNKNOWN
}

func (m *ICMPLayer) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ICMPLayer) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

type FlowMetric struct {
	ABPackets int64 `protobuf:"varint,2,opt,name=ABPackets" json:"ABPackets"`
	ABBytes   int64 `protobuf:"varint,3,opt,name=ABBytes" json:"ABBytes"`
	BAPackets int64 `protobuf:"varint,4,opt,name=BAPackets" json:"BAPackets"`
	BABytes   int64 `protobuf:"varint,5,opt,name=BABytes" json:"BABytes"`
}

func (m *FlowMetric) Reset()                    { *m = FlowMetric{} }
func (m *FlowMetric) String() string            { return proto.CompactTextString(m) }
func (*FlowMetric) ProtoMessage()               {}
func (*FlowMetric) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FlowMetric) GetABPackets() int64 {
	if m != nil {
		return m.ABPackets
	}
	return 0
}

func (m *FlowMetric) GetABBytes() int64 {
	if m != nil {
		return m.ABBytes
	}
	return 0
}

func (m *FlowMetric) GetBAPackets() int64 {
	if m != nil {
		return m.BAPackets
	}
	return 0
}

func (m *FlowMetric) GetBABytes() int64 {
	if m != nil {
		return m.BABytes
	}
	return 0
}

type Flow struct {
	// Flow Universally Unique IDentifier
	// flow.UUID is unique in the universe, as it should be used as a key of an
	// hashtable. By design 2 different flows, their UUID are always different.
	// flow.UUID can be used as Database Index.
	UUID       string `protobuf:"bytes,1,opt,name=UUID" json:"UUID,omitempty"`
	LayersPath string `protobuf:"bytes,2,opt,name=LayersPath" json:"LayersPath,omitempty"`
	// Application is the last layer which is not a payload.
	Application string `protobuf:"bytes,3,opt,name=Application" json:"Application,omitempty"`
	// Data Flow info
	Link      *FlowLayer `protobuf:"bytes,20,opt,name=Link" json:"Link,omitempty"`
	Network   *FlowLayer `protobuf:"bytes,21,opt,name=Network" json:"Network,omitempty"`
	Transport *FlowLayer `protobuf:"bytes,22,opt,name=Transport" json:"Transport,omitempty"`
	ICMP      *ICMPLayer `protobuf:"bytes,23,opt,name=ICMP" json:"ICMP,omitempty"`
	// Data Flow Metric info from the 1st layer
	// amount of data between two updates
	LastUpdateMetric *FlowMetric `protobuf:"bytes,31,opt,name=LastUpdateMetric" json:"LastUpdateMetric,omitempty"`
	// Total amount of data for the whole flow duration
	Metric          *FlowMetric `protobuf:"bytes,32,opt,name=Metric" json:"Metric,omitempty"`
	Start           int64       `protobuf:"varint,10,opt,name=Start" json:"Start"`
	Last            int64       `protobuf:"varint,11,opt,name=Last" json:"Last"`
	LastUpdateStart int64       `protobuf:"varint,12,opt,name=LastUpdateStart" json:"LastUpdateStart"`
	LastUpdateLast  int64       `protobuf:"varint,13,opt,name=LastUpdateLast" json:"LastUpdateLast"`
	// Flow Tracking IDentifier, from 1st packet bytes
	// flow.TrackingID could be used to identify an unique flow whatever it has
	// been captured on the infrastructure. flow.TrackingID is calculated from
	// the bytes of the first packet of his session.
	// flow.TrackingID can be used as a Tag.
	TrackingID   string `protobuf:"bytes,50,opt,name=TrackingID" json:"TrackingID,omitempty"`
	L3TrackingID string `protobuf:"bytes,51,opt,name=L3TrackingID" json:"L3TrackingID,omitempty"`
	// Flow Parent UUID is used as reference to the parent flow
	// Flow.ParentUUID is the same value that point to his parent flow.UUID
	ParentUUID string `protobuf:"bytes,6,opt,name=ParentUUID" json:"ParentUUID"`
	// Topology info
	NodeTID  string `protobuf:"bytes,33,opt,name=NodeTID" json:"NodeTID,omitempty"`
	ANodeTID string `protobuf:"bytes,34,opt,name=ANodeTID" json:"ANodeTID,omitempty"`
	BNodeTID string `protobuf:"bytes,35,opt,name=BNodeTID" json:"BNodeTID,omitempty"`
}

func (m *Flow) Reset()                    { *m = Flow{} }
func (m *Flow) String() string            { return proto.CompactTextString(m) }
func (*Flow) ProtoMessage()               {}
func (*Flow) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Flow) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *Flow) GetLayersPath() string {
	if m != nil {
		return m.LayersPath
	}
	return ""
}

func (m *Flow) GetApplication() string {
	if m != nil {
		return m.Application
	}
	return ""
}

func (m *Flow) GetLink() *FlowLayer {
	if m != nil {
		return m.Link
	}
	return nil
}

func (m *Flow) GetNetwork() *FlowLayer {
	if m != nil {
		return m.Network
	}
	return nil
}

func (m *Flow) GetTransport() *FlowLayer {
	if m != nil {
		return m.Transport
	}
	return nil
}

func (m *Flow) GetICMP() *ICMPLayer {
	if m != nil {
		return m.ICMP
	}
	return nil
}

func (m *Flow) GetLastUpdateMetric() *FlowMetric {
	if m != nil {
		return m.LastUpdateMetric
	}
	return nil
}

func (m *Flow) GetMetric() *FlowMetric {
	if m != nil {
		return m.Metric
	}
	return nil
}

func (m *Flow) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Flow) GetLast() int64 {
	if m != nil {
		return m.Last
	}
	return 0
}

func (m *Flow) GetLastUpdateStart() int64 {
	if m != nil {
		return m.LastUpdateStart
	}
	return 0
}

func (m *Flow) GetLastUpdateLast() int64 {
	if m != nil {
		return m.LastUpdateLast
	}
	return 0
}

func (m *Flow) GetTrackingID() string {
	if m != nil {
		return m.TrackingID
	}
	return ""
}

func (m *Flow) GetL3TrackingID() string {
	if m != nil {
		return m.L3TrackingID
	}
	return ""
}

func (m *Flow) GetParentUUID() string {
	if m != nil {
		return m.ParentUUID
	}
	return ""
}

func (m *Flow) GetNodeTID() string {
	if m != nil {
		return m.NodeTID
	}
	return ""
}

func (m *Flow) GetANodeTID() string {
	if m != nil {
		return m.ANodeTID
	}
	return ""
}

func (m *Flow) GetBNodeTID() string {
	if m != nil {
		return m.BNodeTID
	}
	return ""
}

func init() {
	proto.RegisterType((*FlowLayer)(nil), "flow.FlowLayer")
	proto.RegisterType((*ICMPLayer)(nil), "flow.ICMPLayer")
	proto.RegisterType((*FlowMetric)(nil), "flow.FlowMetric")
	proto.RegisterType((*Flow)(nil), "flow.Flow")
	proto.RegisterEnum("flow.FlowProtocol", FlowProtocol_name, FlowProtocol_value)
	proto.RegisterEnum("flow.ICMPType", ICMPType_name, ICMPType_value)
}

func init() { proto.RegisterFile("flow/flow.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 726 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x54, 0xdd, 0x4e, 0xdb, 0x48,
	0x14, 0xde, 0x10, 0xe7, 0xef, 0xe4, 0x07, 0x33, 0x82, 0xc5, 0xda, 0x5d, 0xed, 0xb2, 0x41, 0xaa,
	0x28, 0x52, 0xa9, 0x04, 0x15, 0x57, 0xbd, 0x19, 0xdb, 0x03, 0xb1, 0x48, 0x6c, 0x77, 0x3c, 0x6e,
	0x7b, 0x17, 0xb9, 0xc1, 0xa5, 0x51, 0x50, 0x1c, 0x25, 0x96, 0x10, 0x52, 0x9f, 0xa4, 0x8f, 0xd7,
	0x27, 0xe9, 0x99, 0x99, 0x38, 0x4e, 0x29, 0xbd, 0x81, 0x39, 0xdf, 0xdf, 0x39, 0xf6, 0x19, 0x07,
	0x76, 0x3f, 0xdf, 0x67, 0x0f, 0xaf, 0xe5, 0x9f, 0xb3, 0xc5, 0x32, 0xcb, 0x33, 0x62, 0xc8, 0x73,
	0xff, 0x0e, 0x5a, 0x57, 0xf8, 0x7f, 0x98, 0x3c, 0xa6, 0x4b, 0x72, 0x06, 0xcd, 0x50, 0x72, 0x93,
	0xec, 0xde, 0xaa, 0x1c, 0x55, 0x4e, 0x7a, 0xe7, 0xe4, 0x4c, 0x39, 0xa4, 0xa4, 0x60, 0xf8, 0x46,
	0x43, 0x3a, 0x50, 0xa1, 0x56, 0x15, 0x85, 0x2d, 0x5e, 0xa1, 0xb2, 0xb2, 0x2d, 0x43, 0x57, 0x36,
	0xe9, 0xc1, 0x8e, 0xe7, 0x5a, 0x35, 0x2c, 0xab, 0x1c, 0x4f, 0xfd, 0x08, 0x5a, 0x9e, 0x33, 0x0a,
	0x75, 0xa3, 0x3e, 0x18, 0xe2, 0x71, 0x91, 0xae, 0x9b, 0xf4, 0x74, 0x13, 0x49, 0x4b, 0x94, 0x2b,
	0x8e, 0x10, 0x30, 0x9c, 0xec, 0x36, 0xb5, 0x76, 0x50, 0xd3, 0xe5, 0xea, 0xbc, 0x0e, 0xad, 0x2a,
	0x44, 0x86, 0x7e, 0x05, 0x90, 0xa3, 0x8d, 0xd2, 0x7c, 0x39, 0x9d, 0x90, 0x7f, 0xa0, 0x45, 0xed,
	0x30, 0x99, 0xcc, 0xd2, 0x7c, 0xa5, 0x6c, 0x55, 0x5e, 0x02, 0xc4, 0x82, 0x06, 0xb5, 0xed, 0xc7,
	0x3c, 0x5d, 0xa9, 0x80, 0x2a, 0x2f, 0x4a, 0xe9, 0xb3, 0x69, 0xe1, 0x33, 0xb4, 0x6f, 0x03, 0x48,
	0x9f, 0x4d, 0xb5, 0x4f, 0x3f, 0x4d, 0x51, 0xf6, 0xbf, 0xd5, 0xc0, 0x90, 0xed, 0xe5, 0xa8, 0x71,
	0x8c, 0x83, 0x55, 0xd4, 0xc3, 0xab, 0x33, 0xf9, 0x17, 0x40, 0x3d, 0xeb, 0x2a, 0x4c, 0xf2, 0x2f,
	0x6a, 0x9a, 0x16, 0xdf, 0x42, 0xc8, 0x11, 0xb4, 0xe9, 0x62, 0x71, 0x3f, 0x9d, 0x24, 0xf9, 0x34,
	0x9b, 0xaf, 0xdf, 0xe2, 0x36, 0x44, 0x8e, 0xc1, 0x18, 0x4e, 0xe7, 0x33, 0x6b, 0x1f, 0xa9, 0xf6,
	0xf9, 0x6e, 0xb9, 0x09, 0x95, 0xc2, 0x15, 0x49, 0x5e, 0x42, 0xc3, 0x4f, 0xf3, 0x87, 0x6c, 0x39,
	0xb3, 0x0e, 0x9e, 0xd7, 0x15, 0x3c, 0x79, 0x05, 0x2d, 0xb1, 0x4c, 0xe6, 0xab, 0x45, 0xb6, 0xcc,
	0xad, 0x3f, 0x9f, 0x17, 0x97, 0x0a, 0xd9, 0x5e, 0x6e, 0xc4, 0x3a, 0xdc, 0x56, 0x6e, 0x56, 0xc8,
	0x15, 0x49, 0xde, 0x82, 0x39, 0x4c, 0x56, 0x79, 0xbc, 0xb8, 0x4d, 0xf2, 0x54, 0xaf, 0xc1, 0xfa,
	0x4f, 0x19, 0xcc, 0x32, 0x5a, 0xe3, 0xfc, 0x17, 0x25, 0x39, 0x81, 0xfa, 0xda, 0x73, 0xf4, 0x1b,
	0xcf, 0x9a, 0x27, 0xfb, 0x50, 0x8b, 0xf2, 0x04, 0xe7, 0x06, 0xb5, 0x02, 0x5d, 0xc8, 0xf7, 0x2e,
	0x33, 0xad, 0xb6, 0x02, 0xd5, 0x19, 0x33, 0x77, 0xcb, 0x3e, 0xda, 0xd3, 0x51, 0xf4, 0x53, 0x98,
	0xbc, 0x80, 0x5e, 0x09, 0xa9, 0x9c, 0xae, 0x12, 0x3e, 0x41, 0xe5, 0x26, 0xf1, 0xad, 0x4c, 0x66,
	0xd3, 0xf9, 0x1d, 0xee, 0xf8, 0x5c, 0x6f, 0xb2, 0x44, 0xf0, 0x32, 0x77, 0x86, 0x17, 0x5b, 0x8a,
	0x0b, 0xa5, 0xf8, 0x09, 0x93, 0x19, 0x61, 0xb2, 0x4c, 0xe7, 0xb9, 0xba, 0x27, 0x75, 0x9d, 0x51,
	0x22, 0xf2, 0x92, 0xf9, 0x78, 0xc1, 0x05, 0x92, 0xff, 0x2b, 0xb2, 0x28, 0xc9, 0x5f, 0xd0, 0xa4,
	0x05, 0xd5, 0x57, 0xd4, 0xa6, 0x96, 0x9c, 0x5d, 0x70, 0xc7, 0x9a, 0x2b, 0xea, 0xd3, 0x05, 0x74,
	0xb6, 0xbf, 0x5a, 0xfc, 0x3a, 0x9b, 0x4c, 0x0c, 0x18, 0xf7, 0x99, 0x30, 0xff, 0x20, 0x4d, 0x5c,
	0x6e, 0xf8, 0xfe, 0x8d, 0x59, 0x21, 0x6d, 0x68, 0x08, 0x27, 0x0c, 0x03, 0x2e, 0xcc, 0x1d, 0x59,
	0xc4, 0xae, 0x2e, 0xaa, 0xd2, 0x11, 0x39, 0x42, 0x57, 0xc6, 0xda, 0x71, 0x69, 0xd6, 0x08, 0x40,
	0x5d, 0xee, 0x1e, 0xdd, 0xf5, 0xcd, 0xf9, 0xd2, 0x6c, 0x9c, 0x7e, 0xaf, 0x40, 0xb3, 0xf8, 0x86,
	0x55, 0x92, 0x7f, 0xe3, 0x07, 0x1f, 0x7c, 0xec, 0xf6, 0x37, 0x1c, 0xba, 0x2c, 0x12, 0x9e, 0x4f,
	0x85, 0x17, 0xf8, 0xe3, 0xd8, 0xe7, 0x8c, 0x3a, 0x03, 0x6a, 0x0f, 0x19, 0x0e, 0x80, 0xc1, 0xcc,
	0x19, 0x04, 0xd8, 0x1d, 0x1b, 0xfa, 0xcc, 0xbb, 0x1e, 0xd8, 0x01, 0xc7, 0xf6, 0x26, 0x74, 0xa8,
	0xeb, 0x72, 0x16, 0x45, 0xe3, 0x11, 0x8d, 0x6e, 0xd6, 0x23, 0xf8, 0x57, 0x01, 0x8e, 0x70, 0x00,
	0x7b, 0x21, 0xe5, 0x74, 0xc4, 0x04, 0xe3, 0xe3, 0x90, 0x07, 0x98, 0x34, 0xc2, 0x69, 0x30, 0x80,
	0x33, 0xd7, 0xe3, 0xcc, 0x11, 0x66, 0x43, 0xce, 0xc6, 0x83, 0x18, 0x15, 0x66, 0x93, 0xec, 0x41,
	0x37, 0x0a, 0x62, 0xee, 0xb0, 0xf1, 0xbb, 0x98, 0xf9, 0xce, 0xc0, 0x6c, 0x49, 0x48, 0x78, 0x23,
	0x36, 0x66, 0x1f, 0x1d, 0xc6, 0x5c, 0xe6, 0x9a, 0x40, 0xba, 0xf8, 0x85, 0x20, 0x14, 0x09, 0x3a,
	0x0a, 0xcd, 0x36, 0x5e, 0xaf, 0x5e, 0x48, 0x9d, 0x1b, 0x26, 0xc6, 0x22, 0x08, 0xc6, 0xb6, 0x77,
	0x6d, 0x76, 0x3e, 0xd5, 0xd5, 0x8f, 0xe7, 0xc5, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x12, 0x6e,
	0xfa, 0x64, 0x4f, 0x05, 0x00, 0x00,
}
