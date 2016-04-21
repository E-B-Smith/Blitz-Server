// Code generated by protoc-gen-go.
// source: UserMessages.proto
// DO NOT EDIT!

package BlitzMessage

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"


// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type UserMessageStatus int32

const (
	UserMessageStatus_MSUnknown   UserMessageStatus = 0
	UserMessageStatus_MSImportant UserMessageStatus = 1
	UserMessageStatus_MSNew       UserMessageStatus = 2
	UserMessageStatus_MSRead      UserMessageStatus = 3
	UserMessageStatus_MSArchived  UserMessageStatus = 4
)

var UserMessageStatus_name = map[int32]string{
	0: "MSUnknown",
	1: "MSImportant",
	2: "MSNew",
	3: "MSRead",
	4: "MSArchived",
}
var UserMessageStatus_value = map[string]int32{
	"MSUnknown":   0,
	"MSImportant": 1,
	"MSNew":       2,
	"MSRead":      3,
	"MSArchived":  4,
}

func (x UserMessageStatus) Enum() *UserMessageStatus {
	p := new(UserMessageStatus)
	*p = x
	return p
}
func (x UserMessageStatus) String() string {
	return proto.EnumName(UserMessageStatus_name, int32(x))
}
func (x *UserMessageStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(UserMessageStatus_value, data, "UserMessageStatus")
	if err != nil {
		return err
	}
	*x = UserMessageStatus(value)
	return nil
}
func (UserMessageStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

type UserMessageType int32

const (
	UserMessageType_MTUnknown      UserMessageType = 0
	UserMessageType_MTSystem       UserMessageType = 1
	UserMessageType_MTConversation UserMessageType = 2
	UserMessageType_MTNotification UserMessageType = 3
)

var UserMessageType_name = map[int32]string{
	0: "MTUnknown",
	1: "MTSystem",
	2: "MTConversation",
	3: "MTNotification",
}
var UserMessageType_value = map[string]int32{
	"MTUnknown":      0,
	"MTSystem":       1,
	"MTConversation": 2,
	"MTNotification": 3,
}

func (x UserMessageType) Enum() *UserMessageType {
	p := new(UserMessageType)
	*p = x
	return p
}
func (x UserMessageType) String() string {
	return proto.EnumName(UserMessageType_name, int32(x))
}
func (x *UserMessageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(UserMessageType_value, data, "UserMessageType")
	if err != nil {
		return err
	}
	*x = UserMessageType(value)
	return nil
}
func (UserMessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

type Conversation struct {
	ConversationID     *string            `protobuf:"bytes,1,opt,name=conversationID" json:"conversationID,omitempty"`
	InitiatorUserID    *string            `protobuf:"bytes,2,opt,name=initiatorUserID" json:"initiatorUserID,omitempty"`
	Status             *UserMessageStatus `protobuf:"varint,3,opt,name=status,enum=BlitzMessage.UserMessageStatus" json:"status,omitempty"`
	ParentFeedPostID   *string            `protobuf:"bytes,4,opt,name=parentFeedPostID" json:"parentFeedPostID,omitempty"`
	CreationDate       *Timestamp         `protobuf:"bytes,5,opt,name=creationDate" json:"creationDate,omitempty"`
	LastActivityDate   *Timestamp         `protobuf:"bytes,6,opt,name=lastActivityDate" json:"lastActivityDate,omitempty"`
	LastMessage        *string            `protobuf:"bytes,7,opt,name=lastMessage" json:"lastMessage,omitempty"`
	MessageCount       *int32             `protobuf:"varint,8,opt,name=messageCount" json:"messageCount,omitempty"`
	UnreadCount        *int32             `protobuf:"varint,9,opt,name=unreadCount" json:"unreadCount,omitempty"`
	MemberIDs          []string           `protobuf:"bytes,10,rep,name=memberIDs" json:"memberIDs,omitempty"`
	ClosedDate         *Timestamp         `protobuf:"bytes,11,opt,name=closedDate" json:"closedDate,omitempty"`
	HeadlineText       *string            `protobuf:"bytes,12,opt,name=headlineText" json:"headlineText,omitempty"`
	LastActivityUserID *string            `protobuf:"bytes,13,opt,name=lastActivityUserID" json:"lastActivityUserID,omitempty"`
	XXX_unrecognized   []byte             `json:"-"`
}

func (m *Conversation) Reset()                    { *m = Conversation{} }
func (m *Conversation) String() string            { return proto.CompactTextString(m) }
func (*Conversation) ProtoMessage()               {}
func (*Conversation) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *Conversation) GetConversationID() string {
	if m != nil && m.ConversationID != nil {
		return *m.ConversationID
	}
	return ""
}

func (m *Conversation) GetInitiatorUserID() string {
	if m != nil && m.InitiatorUserID != nil {
		return *m.InitiatorUserID
	}
	return ""
}

func (m *Conversation) GetStatus() UserMessageStatus {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return UserMessageStatus_MSUnknown
}

func (m *Conversation) GetParentFeedPostID() string {
	if m != nil && m.ParentFeedPostID != nil {
		return *m.ParentFeedPostID
	}
	return ""
}

func (m *Conversation) GetCreationDate() *Timestamp {
	if m != nil {
		return m.CreationDate
	}
	return nil
}

func (m *Conversation) GetLastActivityDate() *Timestamp {
	if m != nil {
		return m.LastActivityDate
	}
	return nil
}

func (m *Conversation) GetLastMessage() string {
	if m != nil && m.LastMessage != nil {
		return *m.LastMessage
	}
	return ""
}

func (m *Conversation) GetMessageCount() int32 {
	if m != nil && m.MessageCount != nil {
		return *m.MessageCount
	}
	return 0
}

func (m *Conversation) GetUnreadCount() int32 {
	if m != nil && m.UnreadCount != nil {
		return *m.UnreadCount
	}
	return 0
}

func (m *Conversation) GetMemberIDs() []string {
	if m != nil {
		return m.MemberIDs
	}
	return nil
}

func (m *Conversation) GetClosedDate() *Timestamp {
	if m != nil {
		return m.ClosedDate
	}
	return nil
}

func (m *Conversation) GetHeadlineText() string {
	if m != nil && m.HeadlineText != nil {
		return *m.HeadlineText
	}
	return ""
}

func (m *Conversation) GetLastActivityUserID() string {
	if m != nil && m.LastActivityUserID != nil {
		return *m.LastActivityUserID
	}
	return ""
}

type ConversationRequest struct {
	UserIDs          []string `protobuf:"bytes,1,rep,name=userIDs" json:"userIDs,omitempty"`
	ParentFeedPostID *string  `protobuf:"bytes,2,opt,name=parentFeedPostID" json:"parentFeedPostID,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ConversationRequest) Reset()                    { *m = ConversationRequest{} }
func (m *ConversationRequest) String() string            { return proto.CompactTextString(m) }
func (*ConversationRequest) ProtoMessage()               {}
func (*ConversationRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

func (m *ConversationRequest) GetUserIDs() []string {
	if m != nil {
		return m.UserIDs
	}
	return nil
}

func (m *ConversationRequest) GetParentFeedPostID() string {
	if m != nil && m.ParentFeedPostID != nil {
		return *m.ParentFeedPostID
	}
	return ""
}

type ConversationResponse struct {
	Conversation     *Conversation `protobuf:"bytes,1,opt,name=conversation" json:"conversation,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *ConversationResponse) Reset()                    { *m = ConversationResponse{} }
func (m *ConversationResponse) String() string            { return proto.CompactTextString(m) }
func (*ConversationResponse) ProtoMessage()               {}
func (*ConversationResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

func (m *ConversationResponse) GetConversation() *Conversation {
	if m != nil {
		return m.Conversation
	}
	return nil
}

type FetchConversations struct {
	Timespan         *Timespan       `protobuf:"bytes,1,opt,name=timespan" json:"timespan,omitempty"`
	Conversations    []*Conversation `protobuf:"bytes,2,rep,name=conversations" json:"conversations,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *FetchConversations) Reset()                    { *m = FetchConversations{} }
func (m *FetchConversations) String() string            { return proto.CompactTextString(m) }
func (*FetchConversations) ProtoMessage()               {}
func (*FetchConversations) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{3} }

func (m *FetchConversations) GetTimespan() *Timespan {
	if m != nil {
		return m.Timespan
	}
	return nil
}

func (m *FetchConversations) GetConversations() []*Conversation {
	if m != nil {
		return m.Conversations
	}
	return nil
}

type UpdateConversationStatus struct {
	ConversationID   *string            `protobuf:"bytes,1,opt,name=conversationID" json:"conversationID,omitempty"`
	Status           *UserMessageStatus `protobuf:"varint,2,opt,name=status,enum=BlitzMessage.UserMessageStatus" json:"status,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *UpdateConversationStatus) Reset()                    { *m = UpdateConversationStatus{} }
func (m *UpdateConversationStatus) String() string            { return proto.CompactTextString(m) }
func (*UpdateConversationStatus) ProtoMessage()               {}
func (*UpdateConversationStatus) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{4} }

func (m *UpdateConversationStatus) GetConversationID() string {
	if m != nil && m.ConversationID != nil {
		return *m.ConversationID
	}
	return ""
}

func (m *UpdateConversationStatus) GetStatus() UserMessageStatus {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return UserMessageStatus_MSUnknown
}

type UserMessage struct {
	MessageID        *string            `protobuf:"bytes,1,opt,name=messageID" json:"messageID,omitempty"`
	SenderID         *string            `protobuf:"bytes,2,opt,name=senderID" json:"senderID,omitempty"`
	ConversationID   *string            `protobuf:"bytes,3,opt,name=conversationID" json:"conversationID,omitempty"`
	Recipients       []string           `protobuf:"bytes,4,rep,name=recipients" json:"recipients,omitempty"`
	CreationDate     *Timestamp         `protobuf:"bytes,5,opt,name=creationDate" json:"creationDate,omitempty"`
	NotificationDate *Timestamp         `protobuf:"bytes,6,opt,name=notificationDate" json:"notificationDate,omitempty"`
	ReadDate         *Timestamp         `protobuf:"bytes,7,opt,name=readDate" json:"readDate,omitempty"`
	MessageType      *UserMessageType   `protobuf:"varint,8,opt,name=messageType,enum=BlitzMessage.UserMessageType" json:"messageType,omitempty"`
	MessageStatus    *UserMessageStatus `protobuf:"varint,9,opt,name=messageStatus,enum=BlitzMessage.UserMessageStatus" json:"messageStatus,omitempty"`
	MessageText      *string            `protobuf:"bytes,10,opt,name=messageText" json:"messageText,omitempty"`
	ActionIcon       *string            `protobuf:"bytes,11,opt,name=actionIcon" json:"actionIcon,omitempty"`
	ActionURL        *string            `protobuf:"bytes,12,opt,name=actionURL" json:"actionURL,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *UserMessage) Reset()                    { *m = UserMessage{} }
func (m *UserMessage) String() string            { return proto.CompactTextString(m) }
func (*UserMessage) ProtoMessage()               {}
func (*UserMessage) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{5} }

func (m *UserMessage) GetMessageID() string {
	if m != nil && m.MessageID != nil {
		return *m.MessageID
	}
	return ""
}

func (m *UserMessage) GetSenderID() string {
	if m != nil && m.SenderID != nil {
		return *m.SenderID
	}
	return ""
}

func (m *UserMessage) GetConversationID() string {
	if m != nil && m.ConversationID != nil {
		return *m.ConversationID
	}
	return ""
}

func (m *UserMessage) GetRecipients() []string {
	if m != nil {
		return m.Recipients
	}
	return nil
}

func (m *UserMessage) GetCreationDate() *Timestamp {
	if m != nil {
		return m.CreationDate
	}
	return nil
}

func (m *UserMessage) GetNotificationDate() *Timestamp {
	if m != nil {
		return m.NotificationDate
	}
	return nil
}

func (m *UserMessage) GetReadDate() *Timestamp {
	if m != nil {
		return m.ReadDate
	}
	return nil
}

func (m *UserMessage) GetMessageType() UserMessageType {
	if m != nil && m.MessageType != nil {
		return *m.MessageType
	}
	return UserMessageType_MTUnknown
}

func (m *UserMessage) GetMessageStatus() UserMessageStatus {
	if m != nil && m.MessageStatus != nil {
		return *m.MessageStatus
	}
	return UserMessageStatus_MSUnknown
}

func (m *UserMessage) GetMessageText() string {
	if m != nil && m.MessageText != nil {
		return *m.MessageText
	}
	return ""
}

func (m *UserMessage) GetActionIcon() string {
	if m != nil && m.ActionIcon != nil {
		return *m.ActionIcon
	}
	return ""
}

func (m *UserMessage) GetActionURL() string {
	if m != nil && m.ActionURL != nil {
		return *m.ActionURL
	}
	return ""
}

type UserMessageUpdate struct {
	Timespan         *Timespan      `protobuf:"bytes,1,opt,name=timespan" json:"timespan,omitempty"`
	Messages         []*UserMessage `protobuf:"bytes,2,rep,name=messages" json:"messages,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *UserMessageUpdate) Reset()                    { *m = UserMessageUpdate{} }
func (m *UserMessageUpdate) String() string            { return proto.CompactTextString(m) }
func (*UserMessageUpdate) ProtoMessage()               {}
func (*UserMessageUpdate) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{6} }

func (m *UserMessageUpdate) GetTimespan() *Timespan {
	if m != nil {
		return m.Timespan
	}
	return nil
}

func (m *UserMessageUpdate) GetMessages() []*UserMessage {
	if m != nil {
		return m.Messages
	}
	return nil
}

func init() {
	proto.RegisterType((*Conversation)(nil), "BlitzMessage.Conversation")
	proto.RegisterType((*ConversationRequest)(nil), "BlitzMessage.ConversationRequest")
	proto.RegisterType((*ConversationResponse)(nil), "BlitzMessage.ConversationResponse")
	proto.RegisterType((*FetchConversations)(nil), "BlitzMessage.FetchConversations")
	proto.RegisterType((*UpdateConversationStatus)(nil), "BlitzMessage.UpdateConversationStatus")
	proto.RegisterType((*UserMessage)(nil), "BlitzMessage.UserMessage")
	proto.RegisterType((*UserMessageUpdate)(nil), "BlitzMessage.UserMessageUpdate")
	proto.RegisterEnum("BlitzMessage.UserMessageStatus", UserMessageStatus_name, UserMessageStatus_value)
	proto.RegisterEnum("BlitzMessage.UserMessageType", UserMessageType_name, UserMessageType_value)
}

var fileDescriptor8 = []byte{
	// 681 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0x5b, 0x53, 0xd3, 0x40,
	0x14, 0x26, 0x0d, 0x97, 0xf6, 0xa4, 0x2d, 0x61, 0x61, 0x20, 0xd6, 0x71, 0x64, 0xfa, 0x84, 0x30,
	0x14, 0xed, 0x83, 0x8f, 0x8c, 0x5c, 0x86, 0x91, 0x19, 0xc3, 0x38, 0xb4, 0xf5, 0xc1, 0xb7, 0x90,
	0xae, 0x76, 0xb1, 0xd9, 0x0d, 0xd9, 0x2d, 0x88, 0xbf, 0xc0, 0x77, 0xdf, 0xfd, 0x51, 0xfc, 0x22,
	0xcf, 0x6e, 0x52, 0xbb, 0x55, 0x6e, 0xfa, 0x96, 0x7c, 0x7b, 0xce, 0xf7, 0x9d, 0xcb, 0xb7, 0x0b,
	0xa4, 0x27, 0x69, 0x16, 0x52, 0x29, 0xa3, 0xcf, 0x54, 0xb6, 0xd2, 0x4c, 0x28, 0x41, 0xaa, 0xfb,
	0x43, 0xa6, 0xbe, 0x15, 0x60, 0xe3, 0xa9, 0x38, 0x3b, 0xa7, 0xb1, 0x62, 0x97, 0x34, 0xde, 0xee,
	0x53, 0x19, 0x67, 0x2c, 0x55, 0x22, 0xcb, 0x43, 0x1b, 0x5e, 0xf7, 0x3a, 0x1d, 0xe7, 0x35, 0x7f,
	0xba, 0x50, 0x3d, 0x10, 0xfc, 0x92, 0x66, 0x32, 0x52, 0x4c, 0x70, 0xb2, 0x0a, 0xf5, 0xd8, 0xfa,
	0x3f, 0x3e, 0x0c, 0x9c, 0x75, 0x67, 0xa3, 0x42, 0xd6, 0x60, 0x91, 0x71, 0xa6, 0x58, 0x84, 0x44,
	0x5a, 0x1f, 0x0f, 0x4a, 0xe6, 0x60, 0x07, 0xe6, 0xa5, 0x8a, 0xd4, 0x48, 0x06, 0x2e, 0xfe, 0xd7,
	0xdb, 0xcf, 0x5b, 0x76, 0x29, 0x2d, 0xab, 0xd6, 0x8e, 0x09, 0x23, 0x01, 0xf8, 0x69, 0x94, 0x51,
	0xae, 0x8e, 0x28, 0xed, 0xbf, 0x17, 0x52, 0x21, 0xd5, 0xac, 0xa1, 0xda, 0x86, 0x6a, 0x9c, 0x51,
	0xa3, 0x7b, 0x18, 0x29, 0x1a, 0xcc, 0x21, 0xea, 0xb5, 0xd7, 0xa6, 0x09, 0xbb, 0x2c, 0xa1, 0x28,
	0x98, 0xa4, 0xe4, 0x15, 0xf8, 0xc3, 0x48, 0xaa, 0x3d, 0xdd, 0x28, 0x53, 0xd7, 0x26, 0x65, 0xfe,
	0xfe, 0x94, 0x65, 0xf0, 0x74, 0x4a, 0x71, 0x10, 0x2c, 0x18, 0xd9, 0x15, 0xa8, 0x26, 0x39, 0x70,
	0x20, 0x46, 0x5c, 0x05, 0x65, 0x44, 0xe7, 0x74, 0xe8, 0x88, 0x63, 0x35, 0xfd, 0x1c, 0xac, 0x18,
	0x70, 0x09, 0x2a, 0x09, 0x4d, 0xce, 0x74, 0xfb, 0x32, 0x80, 0x75, 0x17, 0xb3, 0xb7, 0x00, 0xe2,
	0xa1, 0x90, 0xb4, 0x6f, 0xf4, 0xbd, 0xfb, 0xf5, 0x51, 0x6a, 0x80, 0x94, 0x43, 0xc6, 0x69, 0x97,
	0x7e, 0x55, 0x41, 0xd5, 0x14, 0xd0, 0x00, 0x62, 0x37, 0x52, 0x8c, 0xb7, 0xa6, 0xcf, 0x9a, 0x6f,
	0x60, 0xd9, 0xde, 0xcf, 0x29, 0xbd, 0x18, 0x21, 0x17, 0x59, 0x84, 0x85, 0x91, 0xcc, 0xcb, 0x70,
	0x4c, 0x19, 0xb7, 0x4d, 0xd5, 0x2c, 0xa8, 0xf9, 0x16, 0x56, 0xa6, 0x19, 0x64, 0x2a, 0xb8, 0xa4,
	0xe4, 0x25, 0x4e, 0xdb, 0xc2, 0xcd, 0x9e, 0xbd, 0x76, 0x63, 0xba, 0x74, 0x3b, 0xb3, 0x79, 0x01,
	0xe4, 0x88, 0xaa, 0x78, 0x60, 0x83, 0x92, 0x6c, 0x40, 0x59, 0xe9, 0x06, 0xd3, 0x68, 0xcc, 0xb1,
	0x7a, 0x4b, 0xfb, 0x78, 0x8a, 0x0b, 0xab, 0xd9, 0x8a, 0x12, 0x0b, 0x74, 0x1f, 0x90, 0x8c, 0x21,
	0xe8, 0xa5, 0x7d, 0x9c, 0xac, 0x8d, 0x16, 0x46, 0xba, 0xcb, 0xaa, 0x13, 0x47, 0x96, 0x1e, 0xe5,
	0xc8, 0xe6, 0x0f, 0x17, 0x3c, 0x0b, 0xcd, 0xb7, 0x6c, 0x3e, 0x7f, 0x73, 0xfa, 0x50, 0x96, 0x94,
	0xf7, 0x2d, 0xdf, 0xff, 0xad, 0xee, 0x1a, 0x9c, 0x00, 0x64, 0x34, 0x66, 0x29, 0xc3, 0x5d, 0x48,
	0x34, 0xb6, 0xfb, 0x5f, 0xc6, 0xe6, 0x42, 0xb1, 0x4f, 0x2c, 0x9e, 0xa4, 0x3c, 0x60, 0xec, 0x17,
	0x50, 0xd6, 0x5e, 0x35, 0xa1, 0x0b, 0xf7, 0x87, 0xb6, 0xc1, 0x2b, 0xba, 0xd3, 0x0f, 0x81, 0x71,
	0x7b, 0xbd, 0xfd, 0xec, 0xce, 0x19, 0xe9, 0x20, 0xf2, 0x1a, 0x6a, 0x89, 0x3d, 0x32, 0x73, 0x1d,
	0x1e, 0x71, 0xd7, 0x97, 0x27, 0x5a, 0xda, 0xee, 0x30, 0x9e, 0x50, 0x14, 0x9b, 0x99, 0xe1, 0x00,
	0xcd, 0x8d, 0xa9, 0xe8, 0x91, 0xe7, 0x58, 0xef, 0xf4, 0x5d, 0x7e, 0x2b, 0x9a, 0xe7, 0xb0, 0x64,
	0x11, 0xe6, 0x2e, 0xf8, 0x07, 0xb3, 0x6d, 0x41, 0xb9, 0x90, 0x1e, 0xfb, 0xec, 0xc9, 0x9d, 0xd5,
	0x6e, 0x7e, 0x9c, 0xd2, 0x2a, 0x8a, 0xaf, 0x41, 0x25, 0xec, 0xf4, 0xf8, 0x17, 0x2e, 0xae, 0xb8,
	0x3f, 0x83, 0x57, 0xce, 0x0b, 0x3b, 0xc7, 0x49, 0x2a, 0x32, 0x15, 0x71, 0xe5, 0x3b, 0xa4, 0x02,
	0x73, 0x61, 0xe7, 0x84, 0x5e, 0xf9, 0x25, 0xec, 0x68, 0x3e, 0xec, 0x9c, 0xe2, 0x02, 0x7c, 0x97,
	0xd4, 0x01, 0xc2, 0xce, 0x5e, 0x16, 0x0f, 0xf0, 0xf9, 0xed, 0xfb, 0xb3, 0x9b, 0x1f, 0x60, 0xf1,
	0xcf, 0x71, 0x6a, 0xe6, 0xee, 0x84, 0xb9, 0x0a, 0xe5, 0xb0, 0xdb, 0xb9, 0x96, 0x8a, 0x26, 0x48,
	0x4b, 0xa0, 0x1e, 0x76, 0x6d, 0xbb, 0x23, 0xbf, 0xc1, 0x4e, 0x2c, 0x4f, 0xf8, 0xee, 0xfe, 0xce,
	0xcd, 0x6e, 0x09, 0x66, 0x6e, 0x76, 0x5d, 0xe2, 0xec, 0xe3, 0x67, 0xe0, 0x40, 0x23, 0x16, 0x49,
	0xeb, 0x4c, 0xf7, 0x38, 0xa0, 0x19, 0x9d, 0xea, 0xf6, 0xbb, 0xe3, 0xfc, 0x0a, 0x00, 0x00, 0xff,
	0xff, 0x2f, 0xae, 0x9e, 0xe3, 0x38, 0x06, 0x00, 0x00,
}
