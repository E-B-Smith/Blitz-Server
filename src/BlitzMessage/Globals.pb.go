// Code generated by protoc-gen-go.
// source: Globals.proto
// DO NOT EDIT!

package BlitzMessage

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"


// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Global struct {
	//  Global users:
	SystemUserID     *string `protobuf:"bytes,1,opt,name=SystemUserID,def=e4227250-c8e5-4da9-8177-f084020910b8" json:"SystemUserID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Global) Reset()                    { *m = Global{} }
func (m *Global) String() string            { return proto.CompactTextString(m) }
func (*Global) ProtoMessage()               {}
func (*Global) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

const Default_Global_SystemUserID string = "e4227250-c8e5-4da9-8177-f084020910b8"

func (m *Global) GetSystemUserID() string {
	if m != nil && m.SystemUserID != nil {
		return *m.SystemUserID
	}
	return Default_Global_SystemUserID
}

func init() {
	proto.RegisterType((*Global)(nil), "BlitzMessage.Global")
}

var fileDescriptor3 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x75, 0xcf, 0xc9, 0x4f,
	0x4a, 0xcc, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x71, 0xca, 0xc9, 0x2c, 0xa9,
	0xf2, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x95, 0x92, 0xce, 0x4f, 0xca, 0x4a, 0x4d, 0x2e, 0xc9,
	0x2c, 0x4b, 0x4d, 0xd6, 0x4d, 0x49, 0x2d, 0x4e, 0x2e, 0xca, 0x2c, 0x28, 0xc9, 0x2f, 0x82, 0x28,
	0x55, 0x72, 0xe1, 0x62, 0x83, 0xe8, 0x15, 0xb2, 0xe2, 0xe2, 0x09, 0xae, 0x2c, 0x2e, 0x49, 0xcd,
	0x0d, 0x2d, 0x4e, 0x2d, 0xf2, 0x74, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0xb4, 0x52, 0x49, 0x35,
	0x31, 0x32, 0x32, 0x37, 0x32, 0x35, 0xd0, 0x4d, 0xb6, 0x48, 0x35, 0xd5, 0x35, 0x49, 0x49, 0xb4,
	0xd4, 0xb5, 0x30, 0x34, 0x37, 0xd7, 0x4d, 0x33, 0xb0, 0x30, 0x31, 0x30, 0x32, 0xb0, 0x34, 0x34,
	0x48, 0xb2, 0x70, 0xd2, 0xbf, 0x64, 0xc7, 0xc4, 0xc5, 0x70, 0xc9, 0x8e, 0x59, 0x88, 0xd1, 0x09,
	0xc8, 0x94, 0x60, 0xe4, 0x92, 0x4a, 0xce, 0xcf, 0xd5, 0x4b, 0x02, 0x39, 0x23, 0x23, 0xb5, 0x28,
	0x55, 0x0f, 0xd9, 0x41, 0x1d, 0x8c, 0x8c, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x04, 0x98,
	0x01, 0xb1, 0x00, 0x00, 0x00,
}
