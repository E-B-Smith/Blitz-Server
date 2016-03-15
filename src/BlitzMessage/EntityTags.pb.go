// Code generated by protoc-gen-go.
// source: EntityTags.proto
// DO NOT EDIT!

package BlitzMessage

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"


// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EntityType int32

const (
	EntityType_ETUnknown  EntityType = 0
	EntityType_ETUser     EntityType = 1
	EntityType_ETFeedPost EntityType = 2
)

var EntityType_name = map[int32]string{
	0: "ETUnknown",
	1: "ETUser",
	2: "ETFeedPost",
}
var EntityType_value = map[string]int32{
	"ETUnknown":  0,
	"ETUser":     1,
	"ETFeedPost": 2,
}

func (x EntityType) Enum() *EntityType {
	p := new(EntityType)
	*p = x
	return p
}
func (x EntityType) String() string {
	return proto.EnumName(EntityType_name, int32(x))
}
func (x *EntityType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EntityType_value, data, "EntityType")
	if err != nil {
		return err
	}
	*x = EntityType(value)
	return nil
}
func (EntityType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type EntityTag struct {
	EntityID         *string     `protobuf:"bytes,1,opt,name=entityID" json:"entityID,omitempty"`
	EntityType       *EntityType `protobuf:"varint,2,opt,name=entityType,enum=BlitzMessage.EntityType" json:"entityType,omitempty"`
	EntityTag        *string     `protobuf:"bytes,3,opt,name=entityTag" json:"entityTag,omitempty"`
	EntityIsTagged   *bool       `protobuf:"varint,4,opt,name=entityIsTagged" json:"entityIsTagged,omitempty"`
	EntityTagCount   *int32      `protobuf:"varint,5,opt,name=entityTagCount" json:"entityTagCount,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *EntityTag) Reset()                    { *m = EntityTag{} }
func (m *EntityTag) String() string            { return proto.CompactTextString(m) }
func (*EntityTag) ProtoMessage()               {}
func (*EntityTag) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *EntityTag) GetEntityID() string {
	if m != nil && m.EntityID != nil {
		return *m.EntityID
	}
	return ""
}

func (m *EntityTag) GetEntityType() EntityType {
	if m != nil && m.EntityType != nil {
		return *m.EntityType
	}
	return EntityType_ETUnknown
}

func (m *EntityTag) GetEntityTag() string {
	if m != nil && m.EntityTag != nil {
		return *m.EntityTag
	}
	return ""
}

func (m *EntityTag) GetEntityIsTagged() bool {
	if m != nil && m.EntityIsTagged != nil {
		return *m.EntityIsTagged
	}
	return false
}

func (m *EntityTag) GetEntityTagCount() int32 {
	if m != nil && m.EntityTagCount != nil {
		return *m.EntityTagCount
	}
	return 0
}

type EntityTagList struct {
	Tags             []*EntityTag `protobuf:"bytes,1,rep,name=tags" json:"tags,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *EntityTagList) Reset()                    { *m = EntityTagList{} }
func (m *EntityTagList) String() string            { return proto.CompactTextString(m) }
func (*EntityTagList) ProtoMessage()               {}
func (*EntityTagList) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *EntityTagList) GetTags() []*EntityTag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*EntityTag)(nil), "BlitzMessage.EntityTag")
	proto.RegisterType((*EntityTagList)(nil), "BlitzMessage.EntityTagList")
	proto.RegisterEnum("BlitzMessage.EntityType", EntityType_name, EntityType_value)
}

var fileDescriptor2 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x8f, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0xc7, 0xbb, 0x49, 0x2b, 0xcd, 0x68, 0x43, 0xdc, 0x83, 0x2e, 0xf5, 0x22, 0x05, 0xa1, 0x88,
	0x46, 0xe8, 0x41, 0x6f, 0x3d, 0x44, 0x2b, 0x08, 0x0a, 0x1e, 0xe2, 0x03, 0xa4, 0xc9, 0x10, 0xe3,
	0xc7, 0x6e, 0xd8, 0x1d, 0x95, 0xfa, 0x04, 0x9e, 0x7c, 0xb8, 0x3e, 0x91, 0x63, 0xab, 0xb1, 0x42,
	0x6f, 0x3b, 0xbf, 0x9d, 0xf9, 0x7f, 0x40, 0x34, 0xd1, 0x54, 0xd1, 0x2c, 0xcd, 0x4a, 0x17, 0xd7,
	0xd6, 0x90, 0x91, 0x5b, 0xc9, 0x53, 0x45, 0xef, 0x37, 0xe8, 0x5c, 0x56, 0x62, 0x7f, 0xcf, 0x4c,
	0x1f, 0x30, 0xa7, 0xea, 0x15, 0xf3, 0xe3, 0x02, 0x5d, 0x6e, 0xab, 0x9a, 0x8c, 0x5d, 0xae, 0x0e,
	0x3e, 0x05, 0x04, 0xcd, 0xbd, 0x8c, 0xa0, 0x8b, 0x8b, 0xe1, 0xea, 0x42, 0x89, 0x7d, 0x31, 0x0c,
	0xe4, 0x11, 0xc0, 0x92, 0xa4, 0xb3, 0x1a, 0x95, 0xc7, 0x2c, 0x1c, 0xa9, 0x78, 0x55, 0x3f, 0x9e,
	0x34, 0xff, 0x72, 0x1b, 0x02, 0xfc, 0x15, 0x53, 0xfe, 0x42, 0x60, 0x07, 0xc2, 0x1f, 0x49, 0xc7,
	0xb0, 0xc4, 0x42, 0xb5, 0x99, 0x77, 0xff, 0x38, 0xd3, 0x73, 0xf3, 0xa2, 0x49, 0x75, 0x98, 0x77,
	0x06, 0xa7, 0xd0, 0x6b, 0xf2, 0x5c, 0x57, 0x8e, 0xe4, 0x01, 0xb4, 0x89, 0xab, 0x71, 0x1e, 0x7f,
	0xb8, 0x39, 0xda, 0x5d, 0xeb, 0x9d, 0x95, 0x87, 0x67, 0x00, 0x2b, 0x41, 0x7a, 0xdc, 0x2a, 0xbd,
	0xd3, 0x8f, 0xda, 0xbc, 0xe9, 0xa8, 0x25, 0x01, 0x36, 0x78, 0x74, 0x68, 0x23, 0x21, 0x43, 0x5e,
	0x4c, 0x2f, 0x11, 0x8b, 0x5b, 0xe3, 0x28, 0xf2, 0x92, 0x93, 0xf9, 0xd8, 0x83, 0xd6, 0x7c, 0xec,
	0x4b, 0x91, 0xf0, 0x53, 0x09, 0xe8, 0xe7, 0xe6, 0x39, 0x9e, 0x7e, 0xdb, 0xdc, 0xa3, 0xc5, 0x7f,
	0x86, 0x1f, 0x42, 0x7c, 0x05, 0x00, 0x00, 0xff, 0xff, 0x23, 0x4e, 0xf1, 0xbd, 0x70, 0x01, 0x00,
	0x00,
}
