// Code generated by protoc-gen-go.
// source: Chat.proto
// DO NOT EDIT!

/*
Package Chat is a generated protocol buffer package.

It is generated from these files:
	Chat.proto

It has these top-level messages:
	ChatMessage
	ChatUser
	ChatRoom
	ChatConnect
	ChatEnterRoom
	ChatPresence
	ChatResponse
	ChatMessageType
*/
package Chat

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"


// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Format int32

const (
	Format_FormatUnknown  Format = 0
	Format_FormatJSON     Format = 1
	Format_FormatProtobuf Format = 2
)

var Format_name = map[int32]string{
	0: "FormatUnknown",
	1: "FormatJSON",
	2: "FormatProtobuf",
}
var Format_value = map[string]int32{
	"FormatUnknown":  0,
	"FormatJSON":     1,
	"FormatProtobuf": 2,
}

func (x Format) Enum() *Format {
	p := new(Format)
	*p = x
	return p
}
func (x Format) String() string {
	return proto.EnumName(Format_name, int32(x))
}
func (x *Format) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Format_value, data, "Format")
	if err != nil {
		return err
	}
	*x = Format(value)
	return nil
}
func (Format) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type StatusCode int32

const (
	StatusCode_StatusSuccess       StatusCode = 1
	StatusCode_StatusInputInvalid  StatusCode = 2
	StatusCode_StatusNotAuthorized StatusCode = 3
	StatusCode_StatusServerError   StatusCode = 4
)

var StatusCode_name = map[int32]string{
	1: "StatusSuccess",
	2: "StatusInputInvalid",
	3: "StatusNotAuthorized",
	4: "StatusServerError",
}
var StatusCode_value = map[string]int32{
	"StatusSuccess":       1,
	"StatusInputInvalid":  2,
	"StatusNotAuthorized": 3,
	"StatusServerError":   4,
}

func (x StatusCode) Enum() *StatusCode {
	p := new(StatusCode)
	*p = x
	return p
}
func (x StatusCode) String() string {
	return proto.EnumName(StatusCode_name, int32(x))
}
func (x *StatusCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(StatusCode_value, data, "StatusCode")
	if err != nil {
		return err
	}
	*x = StatusCode(value)
	return nil
}
func (StatusCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ChatMessage struct {
	SenderID         *string  `protobuf:"bytes,1,opt,name=senderID" json:"senderID,omitempty"`
	RoomID           *string  `protobuf:"bytes,2,opt,name=roomID" json:"roomID,omitempty"`
	Timestamp        *float64 `protobuf:"fixed64,3,opt,name=timestamp" json:"timestamp,omitempty"`
	Message          *string  `protobuf:"bytes,4,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ChatMessage) Reset()                    { *m = ChatMessage{} }
func (m *ChatMessage) String() string            { return proto.CompactTextString(m) }
func (*ChatMessage) ProtoMessage()               {}
func (*ChatMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ChatMessage) GetSenderID() string {
	if m != nil && m.SenderID != nil {
		return *m.SenderID
	}
	return ""
}

func (m *ChatMessage) GetRoomID() string {
	if m != nil && m.RoomID != nil {
		return *m.RoomID
	}
	return ""
}

func (m *ChatMessage) GetTimestamp() float64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *ChatMessage) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type ChatUser struct {
	UserID           *string `protobuf:"bytes,1,opt,name=userID" json:"userID,omitempty"`
	Nickname         *string `protobuf:"bytes,2,opt,name=nickname" json:"nickname,omitempty"`
	Format           *Format `protobuf:"varint,3,opt,name=format,enum=Chat.Format" json:"format,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ChatUser) Reset()                    { *m = ChatUser{} }
func (m *ChatUser) String() string            { return proto.CompactTextString(m) }
func (*ChatUser) ProtoMessage()               {}
func (*ChatUser) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ChatUser) GetUserID() string {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return ""
}

func (m *ChatUser) GetNickname() string {
	if m != nil && m.Nickname != nil {
		return *m.Nickname
	}
	return ""
}

func (m *ChatUser) GetFormat() Format {
	if m != nil && m.Format != nil {
		return *m.Format
	}
	return Format_FormatUnknown
}

type ChatRoom struct {
	RoomID           *string `protobuf:"bytes,1,opt,name=roomID" json:"roomID,omitempty"`
	RoomName         *string `protobuf:"bytes,2,opt,name=roomName" json:"roomName,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ChatRoom) Reset()                    { *m = ChatRoom{} }
func (m *ChatRoom) String() string            { return proto.CompactTextString(m) }
func (*ChatRoom) ProtoMessage()               {}
func (*ChatRoom) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ChatRoom) GetRoomID() string {
	if m != nil && m.RoomID != nil {
		return *m.RoomID
	}
	return ""
}

func (m *ChatRoom) GetRoomName() string {
	if m != nil && m.RoomName != nil {
		return *m.RoomName
	}
	return ""
}

type ChatConnect struct {
	IsConnecting     *bool       `protobuf:"varint,1,opt,name=isConnecting" json:"isConnecting,omitempty"`
	User             *ChatUser   `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	Rooms            []*ChatRoom `protobuf:"bytes,3,rep,name=rooms" json:"rooms,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *ChatConnect) Reset()                    { *m = ChatConnect{} }
func (m *ChatConnect) String() string            { return proto.CompactTextString(m) }
func (*ChatConnect) ProtoMessage()               {}
func (*ChatConnect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ChatConnect) GetIsConnecting() bool {
	if m != nil && m.IsConnecting != nil {
		return *m.IsConnecting
	}
	return false
}

func (m *ChatConnect) GetUser() *ChatUser {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *ChatConnect) GetRooms() []*ChatRoom {
	if m != nil {
		return m.Rooms
	}
	return nil
}

type ChatEnterRoom struct {
	User             *ChatUser `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Room             *ChatRoom `protobuf:"bytes,2,opt,name=room" json:"room,omitempty"`
	UserIsEntering   *bool     `protobuf:"varint,3,opt,name=userIsEntering" json:"userIsEntering,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *ChatEnterRoom) Reset()                    { *m = ChatEnterRoom{} }
func (m *ChatEnterRoom) String() string            { return proto.CompactTextString(m) }
func (*ChatEnterRoom) ProtoMessage()               {}
func (*ChatEnterRoom) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ChatEnterRoom) GetUser() *ChatUser {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *ChatEnterRoom) GetRoom() *ChatRoom {
	if m != nil {
		return m.Room
	}
	return nil
}

func (m *ChatEnterRoom) GetUserIsEntering() bool {
	if m != nil && m.UserIsEntering != nil {
		return *m.UserIsEntering
	}
	return false
}

type ChatPresence struct {
	Room             *ChatRoom   `protobuf:"bytes,1,opt,name=room" json:"room,omitempty"`
	Users            []*ChatUser `protobuf:"bytes,2,rep,name=users" json:"users,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *ChatPresence) Reset()                    { *m = ChatPresence{} }
func (m *ChatPresence) String() string            { return proto.CompactTextString(m) }
func (*ChatPresence) ProtoMessage()               {}
func (*ChatPresence) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ChatPresence) GetRoom() *ChatRoom {
	if m != nil {
		return m.Room
	}
	return nil
}

func (m *ChatPresence) GetUsers() []*ChatUser {
	if m != nil {
		return m.Users
	}
	return nil
}

type ChatResponse struct {
	Code             *StatusCode `protobuf:"varint,1,opt,name=code,enum=Chat.StatusCode" json:"code,omitempty"`
	Message          *string     `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *ChatResponse) Reset()                    { *m = ChatResponse{} }
func (m *ChatResponse) String() string            { return proto.CompactTextString(m) }
func (*ChatResponse) ProtoMessage()               {}
func (*ChatResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ChatResponse) GetCode() StatusCode {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return StatusCode_StatusSuccess
}

func (m *ChatResponse) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type ChatMessageType struct {
	ChatMessage      *ChatMessage   `protobuf:"bytes,1,opt,name=chatMessage" json:"chatMessage,omitempty"`
	ChatConnect      *ChatConnect   `protobuf:"bytes,2,opt,name=chatConnect" json:"chatConnect,omitempty"`
	ChatEnterRoom    *ChatEnterRoom `protobuf:"bytes,3,opt,name=chatEnterRoom" json:"chatEnterRoom,omitempty"`
	ChatPresence     *ChatPresence  `protobuf:"bytes,4,opt,name=chatPresence" json:"chatPresence,omitempty"`
	ChatResponse     *ChatResponse  `protobuf:"bytes,5,opt,name=chatResponse" json:"chatResponse,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *ChatMessageType) Reset()                    { *m = ChatMessageType{} }
func (m *ChatMessageType) String() string            { return proto.CompactTextString(m) }
func (*ChatMessageType) ProtoMessage()               {}
func (*ChatMessageType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ChatMessageType) GetChatMessage() *ChatMessage {
	if m != nil {
		return m.ChatMessage
	}
	return nil
}

func (m *ChatMessageType) GetChatConnect() *ChatConnect {
	if m != nil {
		return m.ChatConnect
	}
	return nil
}

func (m *ChatMessageType) GetChatEnterRoom() *ChatEnterRoom {
	if m != nil {
		return m.ChatEnterRoom
	}
	return nil
}

func (m *ChatMessageType) GetChatPresence() *ChatPresence {
	if m != nil {
		return m.ChatPresence
	}
	return nil
}

func (m *ChatMessageType) GetChatResponse() *ChatResponse {
	if m != nil {
		return m.ChatResponse
	}
	return nil
}

func init() {
	proto.RegisterType((*ChatMessage)(nil), "Chat.ChatMessage")
	proto.RegisterType((*ChatUser)(nil), "Chat.ChatUser")
	proto.RegisterType((*ChatRoom)(nil), "Chat.ChatRoom")
	proto.RegisterType((*ChatConnect)(nil), "Chat.ChatConnect")
	proto.RegisterType((*ChatEnterRoom)(nil), "Chat.ChatEnterRoom")
	proto.RegisterType((*ChatPresence)(nil), "Chat.ChatPresence")
	proto.RegisterType((*ChatResponse)(nil), "Chat.ChatResponse")
	proto.RegisterType((*ChatMessageType)(nil), "Chat.ChatMessageType")
	proto.RegisterEnum("Chat.Format", Format_name, Format_value)
	proto.RegisterEnum("Chat.StatusCode", StatusCode_name, StatusCode_value)
}

var fileDescriptor0 = []byte{
	// 551 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x93, 0x5d, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xeb, 0xc4, 0x0d, 0xe9, 0x24, 0x71, 0x9d, 0x2d, 0x94, 0xf0, 0x29, 0xe4, 0x07, 0x54,
	0x45, 0x10, 0xa4, 0x1c, 0xa0, 0x15, 0x84, 0x22, 0xa5, 0xa8, 0xa1, 0x6a, 0xe8, 0x3b, 0xce, 0x7a,
	0x9a, 0x98, 0xd6, 0xbb, 0xd6, 0xee, 0x3a, 0x88, 0x9e, 0x80, 0xf3, 0xf5, 0x06, 0xdc, 0x84, 0xfd,
	0x70, 0x12, 0xd3, 0xf6, 0x6d, 0x67, 0x3c, 0xf3, 0x9b, 0xff, 0xfe, 0x77, 0x0c, 0x30, 0x5a, 0xc4,
	0x6a, 0x90, 0x0b, 0xae, 0x38, 0xf1, 0xcd, 0xf9, 0xf9, 0x0b, 0x3e, 0xfb, 0x89, 0x54, 0xa5, 0x4b,
	0xa4, 0xef, 0x13, 0x94, 0x54, 0xa4, 0xb9, 0xe2, 0xc2, 0x95, 0x44, 0x53, 0x68, 0x99, 0xa2, 0x53,
	0x94, 0x32, 0x9e, 0x23, 0x09, 0xa1, 0x29, 0x91, 0x25, 0x28, 0xc6, 0x9f, 0x7b, 0xde, 0x1b, 0xef,
	0x60, 0x87, 0x04, 0xd0, 0x10, 0x9c, 0x67, 0x3a, 0xae, 0xd9, 0xb8, 0x0b, 0x3b, 0x2a, 0xcd, 0x50,
	0xaa, 0x38, 0xcb, 0x7b, 0x75, 0x9d, 0xf2, 0xc8, 0x2e, 0x3c, 0xca, 0x5c, 0x7f, 0xcf, 0x37, 0x35,
	0xd1, 0x09, 0x34, 0x0d, 0xf4, 0x42, 0xa2, 0x30, 0xfd, 0x85, 0xac, 0xf0, 0xf4, 0x04, 0x96, 0xd2,
	0x2b, 0x16, 0x67, 0x58, 0x12, 0x5f, 0x42, 0xe3, 0x92, 0x8b, 0x2c, 0x56, 0x16, 0x17, 0x0c, 0xdb,
	0x03, 0x7b, 0x85, 0x2f, 0x36, 0x17, 0xbd, 0x73, 0xac, 0x73, 0xad, 0xa1, 0xa2, 0x65, 0xcd, 0x32,
	0xf1, 0x64, 0xcd, 0x8a, 0x7e, 0xb8, 0xeb, 0x8c, 0x38, 0x63, 0xfa, 0xca, 0xe4, 0x31, 0xb4, 0x53,
	0x59, 0x06, 0x29, 0x9b, 0xdb, 0xb6, 0xa6, 0x1e, 0xe8, 0x1b, 0x49, 0xb6, 0xa5, 0x35, 0x0c, 0xdc,
	0xb8, 0xb5, 0xe0, 0x57, 0xb0, 0x6d, 0xa0, 0x52, 0xab, 0xa9, 0xff, 0xff, 0xd9, 0x68, 0x88, 0x28,
	0x74, 0xcc, 0xf9, 0x98, 0x29, 0x14, 0x56, 0xd4, 0x8a, 0xe6, 0x3d, 0x48, 0xd3, 0x5f, 0x0d, 0xed,
	0xfe, 0x2c, 0xdb, 0xbb, 0x0f, 0x81, 0x35, 0x47, 0x5a, 0x9c, 0x51, 0x68, 0x2c, 0x68, 0x46, 0x5f,
	0xa1, 0x6d, 0x6a, 0xce, 0x04, 0xea, 0xd7, 0xa0, 0xb8, 0xa6, 0x78, 0x0f, 0x52, 0xb4, 0x62, 0x43,
	0x91, 0x7a, 0x48, 0xfd, 0xbe, 0x84, 0xe8, 0xc8, 0xc1, 0xce, 0x51, 0xe6, 0x9c, 0x49, 0x24, 0xaf,
	0xc1, 0xa7, 0x3c, 0x41, 0x0b, 0x0b, 0x86, 0xa1, 0xab, 0x9e, 0xaa, 0x58, 0x15, 0xda, 0xaa, 0x04,
	0xab, 0xcf, 0xe9, 0x4c, 0xfd, 0xeb, 0xc1, 0x6e, 0x65, 0x49, 0xbe, 0xff, 0xce, 0x91, 0xbc, 0x85,
	0x16, 0xdd, 0xa4, 0x4a, 0x61, 0xdd, 0xcd, 0xe4, 0xd5, 0x42, 0x95, 0x75, 0xe5, 0x1b, 0x94, 0x36,
	0x54, 0xea, 0x56, 0x2f, 0xd5, 0x87, 0x0e, 0xad, 0xda, 0x6a, 0x8d, 0x68, 0x0d, 0xf7, 0x36, 0x95,
	0x1b, 0xc7, 0x0f, 0xa0, 0x4d, 0x2b, 0xee, 0xd8, 0xa5, 0x6b, 0x0d, 0xc9, 0xa6, 0x74, 0xed, 0x5b,
	0x59, 0xb9, 0xba, 0x7a, 0x6f, 0xfb, 0x6e, 0xe5, 0xea, 0x4b, 0xff, 0x08, 0x1a, 0x6e, 0xe1, 0xf4,
	0x82, 0x77, 0xdc, 0xe9, 0x82, 0x5d, 0x31, 0xfe, 0x8b, 0x85, 0x5b, 0x7a, 0xef, 0xc0, 0xa5, 0x4e,
	0xa6, 0xdf, 0x26, 0xa1, 0x47, 0x08, 0x04, 0x2e, 0x3e, 0x33, 0xff, 0xd0, 0xac, 0xb8, 0x0c, 0x6b,
	0xfd, 0x39, 0x40, 0xc5, 0x43, 0x0d, 0x71, 0xd1, 0xb4, 0xa0, 0x54, 0x5b, 0xa1, 0x9b, 0xf6, 0x81,
	0xb8, 0xd4, 0x98, 0xe5, 0x85, 0x1a, 0xb3, 0x65, 0x7c, 0x9d, 0x26, 0x61, 0x8d, 0x3c, 0x85, 0x3d,
	0x97, 0x9f, 0x70, 0xf5, 0xb1, 0x50, 0x0b, 0x2e, 0xd2, 0x1b, 0x4c, 0xc2, 0x3a, 0x79, 0x02, 0xdd,
	0x92, 0x81, 0x62, 0x89, 0xe2, 0x58, 0x08, 0x2e, 0x42, 0xff, 0xd3, 0x87, 0xdb, 0xc3, 0x1a, 0x6c,
	0xdd, 0x1e, 0xfa, 0xa4, 0x36, 0x3a, 0xd5, 0xe7, 0x9e, 0x07, 0xcf, 0x28, 0xcf, 0x06, 0xb3, 0xeb,
	0x54, 0xdd, 0x2c, 0x50, 0x60, 0xd5, 0xff, 0x3f, 0x9e, 0xf7, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x01,
	0x57, 0xab, 0x63, 0x12, 0x04, 0x00, 0x00,
}
