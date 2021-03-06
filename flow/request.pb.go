// Code generated by protoc-gen-go.
// source: flow/request.proto
// DO NOT EDIT!

package flow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type FlowSearchReply struct {
	FlowSet *FlowSet `protobuf:"bytes,1,opt,name=FlowSet" json:"FlowSet,omitempty"`
}

func (m *FlowSearchReply) Reset()                    { *m = FlowSearchReply{} }
func (m *FlowSearchReply) String() string            { return proto.CompactTextString(m) }
func (*FlowSearchReply) ProtoMessage()               {}
func (*FlowSearchReply) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *FlowSearchReply) GetFlowSet() *FlowSet {
	if m != nil {
		return m.FlowSet
	}
	return nil
}

func init() {
	proto.RegisterType((*FlowSearchReply)(nil), "flow.FlowSearchReply")
}

func init() { proto.RegisterFile("flow/request.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 102 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0xcb, 0xc9, 0x2f,
	0xd7, 0x2f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x01, 0x89, 0x49, 0xf1, 0x81, 0x65, 0x8a, 0x53, 0xa1, 0xa2, 0x4a, 0x56, 0x5c, 0xfc, 0x6e, 0x40,
	0x91, 0xe0, 0xd4, 0xc4, 0xa2, 0xe4, 0x8c, 0xa0, 0xd4, 0x82, 0x9c, 0x4a, 0x21, 0x75, 0x2e, 0x76,
	0x88, 0x50, 0x89, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x11, 0xaf, 0x1e, 0x48, 0x93, 0x1e, 0x54,
	0x30, 0x08, 0x26, 0x9b, 0xc4, 0x06, 0x36, 0xc2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xd9,
	0xd6, 0xb8, 0x6e, 0x00, 0x00, 0x00,
}
