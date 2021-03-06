// Code generated by protoc-gen-go.
// source: UserEvents.proto
// DO NOT EDIT!

package BlitzMessage

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"


// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type UserEvent struct {
	Timestamp        *Timestamp `protobuf:"bytes,1,req,name=timestamp" json:"timestamp,omitempty"`
	Location         *Location  `protobuf:"bytes,2,opt,name=location" json:"location,omitempty"`
	Event            *string    `protobuf:"bytes,3,req,name=event" json:"event,omitempty"`
	EventData        []string   `protobuf:"bytes,4,rep,name=eventData" json:"eventData,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *UserEvent) Reset()                    { *m = UserEvent{} }
func (m *UserEvent) String() string            { return proto.CompactTextString(m) }
func (*UserEvent) ProtoMessage()               {}
func (*UserEvent) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *UserEvent) GetTimestamp() *Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *UserEvent) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *UserEvent) GetEvent() string {
	if m != nil && m.Event != nil {
		return *m.Event
	}
	return ""
}

func (m *UserEvent) GetEventData() []string {
	if m != nil {
		return m.EventData
	}
	return nil
}

type UserEventBatch struct {
	UserEvents       []*UserEvent `protobuf:"bytes,1,rep,name=userEvents" json:"userEvents,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *UserEventBatch) Reset()                    { *m = UserEventBatch{} }
func (m *UserEventBatch) String() string            { return proto.CompactTextString(m) }
func (*UserEventBatch) ProtoMessage()               {}
func (*UserEventBatch) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

func (m *UserEventBatch) GetUserEvents() []*UserEvent {
	if m != nil {
		return m.UserEvents
	}
	return nil
}

type UserEventBatchResponse struct {
	LatestEventUpdate *Timestamp `protobuf:"bytes,1,opt,name=latestEventUpdate" json:"latestEventUpdate,omitempty"`
	XXX_unrecognized  []byte     `json:"-"`
}

func (m *UserEventBatchResponse) Reset()                    { *m = UserEventBatchResponse{} }
func (m *UserEventBatchResponse) String() string            { return proto.CompactTextString(m) }
func (*UserEventBatchResponse) ProtoMessage()               {}
func (*UserEventBatchResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

func (m *UserEventBatchResponse) GetLatestEventUpdate() *Timestamp {
	if m != nil {
		return m.LatestEventUpdate
	}
	return nil
}

func init() {
	proto.RegisterType((*UserEvent)(nil), "BlitzMessage.UserEvent")
	proto.RegisterType((*UserEventBatch)(nil), "BlitzMessage.UserEventBatch")
	proto.RegisterType((*UserEventBatchResponse)(nil), "BlitzMessage.UserEventBatchResponse")
}

func init() { proto.RegisterFile("UserEvents.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x8f, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xdd, 0x44, 0xc1, 0x4c, 0x55, 0xec, 0x1e, 0xea, 0x12, 0x2f, 0x21, 0xa7, 0xa0, 0x18,
	0x21, 0x77, 0x73, 0x08, 0x7a, 0xab, 0x17, 0x69, 0x1f, 0x60, 0xbb, 0x1d, 0xec, 0x4a, 0x92, 0x5d,
	0x76, 0xc7, 0x82, 0x5e, 0xbd, 0xf8, 0x7c, 0x79, 0x22, 0x69, 0xac, 0xc1, 0x20, 0x1e, 0x87, 0xff,
	0xff, 0x7e, 0xbe, 0x81, 0xf3, 0xa5, 0x47, 0xf7, 0xb0, 0xc5, 0x96, 0x7c, 0x6e, 0x9d, 0x21, 0xc3,
	0x4f, 0xaa, 0x5a, 0xd3, 0xfb, 0x23, 0x7a, 0x2f, 0x9f, 0x31, 0xbe, 0x34, 0xab, 0x17, 0x54, 0xa4,
	0xb7, 0xa8, 0x6e, 0xd6, 0xe8, 0x95, 0xd3, 0x96, 0x8c, 0xfb, 0xae, 0xc6, 0x93, 0xc5, 0x9b, 0xc5,
	0x3d, 0x97, 0x7e, 0x30, 0x88, 0x86, 0x31, 0x7e, 0x05, 0x11, 0xe9, 0x06, 0x3d, 0xc9, 0xc6, 0x0a,
	0x96, 0x04, 0xd9, 0xa4, 0xb8, 0xc8, 0x7f, 0x2f, 0xe7, 0x8b, 0x9f, 0x98, 0x67, 0x70, 0x5c, 0x1b,
	0x25, 0x49, 0x9b, 0x56, 0x04, 0x09, 0xcb, 0x26, 0xc5, 0x6c, 0x5c, 0x9d, 0xef, 0x53, 0x7e, 0x0a,
	0x47, 0xb8, 0x9b, 0x17, 0x61, 0x12, 0x64, 0x11, 0x9f, 0x42, 0xd4, 0x9f, 0xf7, 0x92, 0xa4, 0x38,
	0x4c, 0xc2, 0x2c, 0x4a, 0xef, 0xe0, 0x6c, 0x90, 0xa8, 0x24, 0xa9, 0x0d, 0xbf, 0x06, 0x78, 0x1d,
	0x7e, 0x14, 0x2c, 0x09, 0xff, 0xaa, 0x0c, 0x44, 0x3a, 0x87, 0xd9, 0x18, 0x7f, 0x42, 0x6f, 0x4d,
	0xeb, 0x91, 0x17, 0x30, 0xad, 0x25, 0xa1, 0xa7, 0x3e, 0x5b, 0xda, 0xb5, 0x24, 0x14, 0xac, 0xb7,
	0xfd, 0xef, 0xb1, 0xea, 0xb6, 0x2b, 0x03, 0x38, 0xe8, 0xca, 0x90, 0xb3, 0xaa, 0x2b, 0x03, 0xc1,
	0x20, 0x56, 0xa6, 0xc9, 0x57, 0xbb, 0xf6, 0x06, 0x1d, 0x8e, 0xb8, 0x4f, 0xc6, 0xbe, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x3f, 0xe3, 0x42, 0x2e, 0x8e, 0x01, 0x00, 0x00,
}
