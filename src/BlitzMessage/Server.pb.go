// Code generated by protoc-gen-go.
// source: Server.proto
// DO NOT EDIT!

package BlitzMessage

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"


// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ResponseCode int32

const (
	ResponseCode_RCSuccess       ResponseCode = 1
	ResponseCode_RCInputCorrupt  ResponseCode = 2
	ResponseCode_RCInputInvalid  ResponseCode = 3
	ResponseCode_RCServerWarning ResponseCode = 4
	ResponseCode_RCServerError   ResponseCode = 5
	ResponseCode_RCNotAuthorized ResponseCode = 6
	ResponseCode_RCClientTooOld  ResponseCode = 7
)

var ResponseCode_name = map[int32]string{
	1: "RCSuccess",
	2: "RCInputCorrupt",
	3: "RCInputInvalid",
	4: "RCServerWarning",
	5: "RCServerError",
	6: "RCNotAuthorized",
	7: "RCClientTooOld",
}
var ResponseCode_value = map[string]int32{
	"RCSuccess":       1,
	"RCInputCorrupt":  2,
	"RCInputInvalid":  3,
	"RCServerWarning": 4,
	"RCServerError":   5,
	"RCNotAuthorized": 6,
	"RCClientTooOld":  7,
}

func (x ResponseCode) Enum() *ResponseCode {
	p := new(ResponseCode)
	*p = x
	return p
}
func (x ResponseCode) String() string {
	return proto.EnumName(ResponseCode_name, int32(x))
}
func (x *ResponseCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ResponseCode_value, data, "ResponseCode")
	if err != nil {
		return err
	}
	*x = ResponseCode(value)
	return nil
}
func (ResponseCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

type DebugMessage struct {
	DebugText        []string `protobuf:"bytes,1,rep,name=debugText" json:"debugText,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DebugMessage) Reset()                    { *m = DebugMessage{} }
func (m *DebugMessage) String() string            { return proto.CompactTextString(m) }
func (*DebugMessage) ProtoMessage()               {}
func (*DebugMessage) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *DebugMessage) GetDebugText() []string {
	if m != nil {
		return m.DebugText
	}
	return nil
}

type SessionRequest struct {
	Location             *Location    `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`
	DeviceInfo           *DeviceInfo  `protobuf:"bytes,2,opt,name=deviceInfo" json:"deviceInfo,omitempty"`
	Profile              *UserProfile `protobuf:"bytes,3,opt,name=profile" json:"profile,omitempty"`
	LastAppDataResetDate *Timestamp   `protobuf:"bytes,4,opt,name=lastAppDataResetDate" json:"lastAppDataResetDate,omitempty"`
	XXX_unrecognized     []byte       `json:"-"`
}

func (m *SessionRequest) Reset()                    { *m = SessionRequest{} }
func (m *SessionRequest) String() string            { return proto.CompactTextString(m) }
func (*SessionRequest) ProtoMessage()               {}
func (*SessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *SessionRequest) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *SessionRequest) GetDeviceInfo() *DeviceInfo {
	if m != nil {
		return m.DeviceInfo
	}
	return nil
}

func (m *SessionRequest) GetProfile() *UserProfile {
	if m != nil {
		return m.Profile
	}
	return nil
}

func (m *SessionRequest) GetLastAppDataResetDate() *Timestamp {
	if m != nil {
		return m.LastAppDataResetDate
	}
	return nil
}

type BlitzHereAppOptions struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *BlitzHereAppOptions) Reset()                    { *m = BlitzHereAppOptions{} }
func (m *BlitzHereAppOptions) String() string            { return proto.CompactTextString(m) }
func (*BlitzHereAppOptions) ProtoMessage()               {}
func (*BlitzHereAppOptions) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

type AppOptions struct {
	BlitzHereOptions *BlitzHereAppOptions `protobuf:"bytes,1,opt,name=blitzHereOptions" json:"blitzHereOptions,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *AppOptions) Reset()                    { *m = AppOptions{} }
func (m *AppOptions) String() string            { return proto.CompactTextString(m) }
func (*AppOptions) ProtoMessage()               {}
func (*AppOptions) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{3} }

func (m *AppOptions) GetBlitzHereOptions() *BlitzHereAppOptions {
	if m != nil {
		return m.BlitzHereOptions
	}
	return nil
}

type SessionResponse struct {
	UserID           *string              `protobuf:"bytes,1,opt,name=userID" json:"userID,omitempty"`
	SessionToken     *string              `protobuf:"bytes,2,opt,name=sessionToken" json:"sessionToken,omitempty"`
	ServerURL        *string              `protobuf:"bytes,3,opt,name=serverURL" json:"serverURL,omitempty"`
	UserMessages     []*UserMessage       `protobuf:"bytes,4,rep,name=userMessages" json:"userMessages,omitempty"`
	UserProfile      *UserProfile         `protobuf:"bytes,5,opt,name=userProfile" json:"userProfile,omitempty"`
	ResetAllAppData  *bool                `protobuf:"varint,6,opt,name=resetAllAppData" json:"resetAllAppData,omitempty"`
	InviteRequest    *AcceptInviteRequest `protobuf:"bytes,7,opt,name=inviteRequest" json:"inviteRequest,omitempty"`
	AppOptions       *AppOptions          `protobuf:"bytes,8,opt,name=appOptions" json:"appOptions,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *SessionResponse) Reset()                    { *m = SessionResponse{} }
func (m *SessionResponse) String() string            { return proto.CompactTextString(m) }
func (*SessionResponse) ProtoMessage()               {}
func (*SessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{4} }

func (m *SessionResponse) GetUserID() string {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return ""
}

func (m *SessionResponse) GetSessionToken() string {
	if m != nil && m.SessionToken != nil {
		return *m.SessionToken
	}
	return ""
}

func (m *SessionResponse) GetServerURL() string {
	if m != nil && m.ServerURL != nil {
		return *m.ServerURL
	}
	return ""
}

func (m *SessionResponse) GetUserMessages() []*UserMessage {
	if m != nil {
		return m.UserMessages
	}
	return nil
}

func (m *SessionResponse) GetUserProfile() *UserProfile {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

func (m *SessionResponse) GetResetAllAppData() bool {
	if m != nil && m.ResetAllAppData != nil {
		return *m.ResetAllAppData
	}
	return false
}

func (m *SessionResponse) GetInviteRequest() *AcceptInviteRequest {
	if m != nil {
		return m.InviteRequest
	}
	return nil
}

func (m *SessionResponse) GetAppOptions() *AppOptions {
	if m != nil {
		return m.AppOptions
	}
	return nil
}

type PushConnect struct {
	UserID           *string `protobuf:"bytes,1,opt,name=userID" json:"userID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PushConnect) Reset()                    { *m = PushConnect{} }
func (m *PushConnect) String() string            { return proto.CompactTextString(m) }
func (*PushConnect) ProtoMessage()               {}
func (*PushConnect) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{5} }

func (m *PushConnect) GetUserID() string {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return ""
}

type PushDisconnect struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *PushDisconnect) Reset()                    { *m = PushDisconnect{} }
func (m *PushDisconnect) String() string            { return proto.CompactTextString(m) }
func (*PushDisconnect) ProtoMessage()               {}
func (*PushDisconnect) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{6} }

type RequestType struct {
	SessionRequest           *SessionRequest           `protobuf:"bytes,1,opt,name=sessionRequest" json:"sessionRequest,omitempty"`
	UserEventBatch           *UserEventBatch           `protobuf:"bytes,2,opt,name=userEventBatch" json:"userEventBatch,omitempty"`
	UserProfileUpdate        *UserProfileUpdate        `protobuf:"bytes,3,opt,name=userProfileUpdate" json:"userProfileUpdate,omitempty"`
	UserProfileQuery         *UserProfileQuery         `protobuf:"bytes,4,opt,name=userProfileQuery" json:"userProfileQuery,omitempty"`
	ConfirmationRequest      *ConfirmationRequest      `protobuf:"bytes,5,opt,name=confirmationRequest" json:"confirmationRequest,omitempty"`
	MessageSendRequest       *UserMessageUpdate        `protobuf:"bytes,6,opt,name=messageSendRequest" json:"messageSendRequest,omitempty"`
	MessageFetchRequest      *UserMessageUpdate        `protobuf:"bytes,7,opt,name=messageFetchRequest" json:"messageFetchRequest,omitempty"`
	DebugMessage             *DebugMessage             `protobuf:"bytes,8,opt,name=debugMessage" json:"debugMessage,omitempty"`
	ImageUpload              *ImageUpload              `protobuf:"bytes,9,opt,name=imageUpload" json:"imageUpload,omitempty"`
	AcceptInviteRequest      *AcceptInviteRequest      `protobuf:"bytes,10,opt,name=acceptInviteRequest" json:"acceptInviteRequest,omitempty"`
	FeedPostFetchRequest     *FeedPostFetchRequest     `protobuf:"bytes,11,opt,name=feedPostFetchRequest" json:"feedPostFetchRequest,omitempty"`
	FeedPostUpdateRequest    *FeedPostUpdateRequest    `protobuf:"bytes,12,opt,name=feedPostUpdateRequest" json:"feedPostUpdateRequest,omitempty"`
	AutocompleteRequest      *AutocompleteRequest      `protobuf:"bytes,13,opt,name=autocompleteRequest" json:"autocompleteRequest,omitempty"`
	EntityTagUpdate          *EntityTagList            `protobuf:"bytes,14,opt,name=entityTagUpdate" json:"entityTagUpdate,omitempty"`
	UserSearchRequest        *UserSearchRequest        `protobuf:"bytes,15,opt,name=userSearchRequest" json:"userSearchRequest,omitempty"`
	PushConnect              *PushConnect              `protobuf:"bytes,16,opt,name=pushConnect" json:"pushConnect,omitempty"`
	PushDisconnect           *PushDisconnect           `protobuf:"bytes,17,opt,name=pushDisconnect" json:"pushDisconnect,omitempty"`
	ConversationRequest      *ConversationRequest      `protobuf:"bytes,18,opt,name=conversationRequest" json:"conversationRequest,omitempty"`
	FetchConversations       *FetchConversations       `protobuf:"bytes,19,opt,name=fetchConversations" json:"fetchConversations,omitempty"`
	UserReview               *UserReview               `protobuf:"bytes,20,opt,name=userReview" json:"userReview,omitempty"`
	UpdateConversationStatus *UpdateConversationStatus `protobuf:"bytes,21,opt,name=updateConversationStatus" json:"updateConversationStatus,omitempty"`
	UserCardInfo             *UserCardInfo             `protobuf:"bytes,22,opt,name=userCardInfo" json:"userCardInfo,omitempty"`
	XXX_unrecognized         []byte                    `json:"-"`
}

func (m *RequestType) Reset()                    { *m = RequestType{} }
func (m *RequestType) String() string            { return proto.CompactTextString(m) }
func (*RequestType) ProtoMessage()               {}
func (*RequestType) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{7} }

func (m *RequestType) GetSessionRequest() *SessionRequest {
	if m != nil {
		return m.SessionRequest
	}
	return nil
}

func (m *RequestType) GetUserEventBatch() *UserEventBatch {
	if m != nil {
		return m.UserEventBatch
	}
	return nil
}

func (m *RequestType) GetUserProfileUpdate() *UserProfileUpdate {
	if m != nil {
		return m.UserProfileUpdate
	}
	return nil
}

func (m *RequestType) GetUserProfileQuery() *UserProfileQuery {
	if m != nil {
		return m.UserProfileQuery
	}
	return nil
}

func (m *RequestType) GetConfirmationRequest() *ConfirmationRequest {
	if m != nil {
		return m.ConfirmationRequest
	}
	return nil
}

func (m *RequestType) GetMessageSendRequest() *UserMessageUpdate {
	if m != nil {
		return m.MessageSendRequest
	}
	return nil
}

func (m *RequestType) GetMessageFetchRequest() *UserMessageUpdate {
	if m != nil {
		return m.MessageFetchRequest
	}
	return nil
}

func (m *RequestType) GetDebugMessage() *DebugMessage {
	if m != nil {
		return m.DebugMessage
	}
	return nil
}

func (m *RequestType) GetImageUpload() *ImageUpload {
	if m != nil {
		return m.ImageUpload
	}
	return nil
}

func (m *RequestType) GetAcceptInviteRequest() *AcceptInviteRequest {
	if m != nil {
		return m.AcceptInviteRequest
	}
	return nil
}

func (m *RequestType) GetFeedPostFetchRequest() *FeedPostFetchRequest {
	if m != nil {
		return m.FeedPostFetchRequest
	}
	return nil
}

func (m *RequestType) GetFeedPostUpdateRequest() *FeedPostUpdateRequest {
	if m != nil {
		return m.FeedPostUpdateRequest
	}
	return nil
}

func (m *RequestType) GetAutocompleteRequest() *AutocompleteRequest {
	if m != nil {
		return m.AutocompleteRequest
	}
	return nil
}

func (m *RequestType) GetEntityTagUpdate() *EntityTagList {
	if m != nil {
		return m.EntityTagUpdate
	}
	return nil
}

func (m *RequestType) GetUserSearchRequest() *UserSearchRequest {
	if m != nil {
		return m.UserSearchRequest
	}
	return nil
}

func (m *RequestType) GetPushConnect() *PushConnect {
	if m != nil {
		return m.PushConnect
	}
	return nil
}

func (m *RequestType) GetPushDisconnect() *PushDisconnect {
	if m != nil {
		return m.PushDisconnect
	}
	return nil
}

func (m *RequestType) GetConversationRequest() *ConversationRequest {
	if m != nil {
		return m.ConversationRequest
	}
	return nil
}

func (m *RequestType) GetFetchConversations() *FetchConversations {
	if m != nil {
		return m.FetchConversations
	}
	return nil
}

func (m *RequestType) GetUserReview() *UserReview {
	if m != nil {
		return m.UserReview
	}
	return nil
}

func (m *RequestType) GetUpdateConversationStatus() *UpdateConversationStatus {
	if m != nil {
		return m.UpdateConversationStatus
	}
	return nil
}

func (m *RequestType) GetUserCardInfo() *UserCardInfo {
	if m != nil {
		return m.UserCardInfo
	}
	return nil
}

type ServerRequest struct {
	SessionToken     *string      `protobuf:"bytes,1,opt,name=sessionToken" json:"sessionToken,omitempty"`
	RequestType      *RequestType `protobuf:"bytes,2,opt,name=requestType" json:"requestType,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ServerRequest) Reset()                    { *m = ServerRequest{} }
func (m *ServerRequest) String() string            { return proto.CompactTextString(m) }
func (*ServerRequest) ProtoMessage()               {}
func (*ServerRequest) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{8} }

func (m *ServerRequest) GetSessionToken() string {
	if m != nil && m.SessionToken != nil {
		return *m.SessionToken
	}
	return ""
}

func (m *ServerRequest) GetRequestType() *RequestType {
	if m != nil {
		return m.RequestType
	}
	return nil
}

type ResponseType struct {
	SessionResponse        *SessionResponse        `protobuf:"bytes,1,opt,name=sessionResponse" json:"sessionResponse,omitempty"`
	UserEventBatchResponse *UserEventBatchResponse `protobuf:"bytes,2,opt,name=userEventBatchResponse" json:"userEventBatchResponse,omitempty"`
	UserProfileUpdate      *UserProfileUpdate      `protobuf:"bytes,3,opt,name=userProfileUpdate" json:"userProfileUpdate,omitempty"`
	UserProfileQuery       *UserProfileQuery       `protobuf:"bytes,4,opt,name=userProfileQuery" json:"userProfileQuery,omitempty"`
	ConfirmationRequest    *ConfirmationRequest    `protobuf:"bytes,5,opt,name=confirmationRequest" json:"confirmationRequest,omitempty"`
	UserMessageUpdate      *UserMessageUpdate      `protobuf:"bytes,6,opt,name=userMessageUpdate" json:"userMessageUpdate,omitempty"`
	DebugMessage           *DebugMessage           `protobuf:"bytes,7,opt,name=debugMessage" json:"debugMessage,omitempty"`
	ImageUploadReply       *ImageUpload            `protobuf:"bytes,8,opt,name=imageUploadReply" json:"imageUploadReply,omitempty"`
	AcceptInviteResponse   *AcceptInviteResponse   `protobuf:"bytes,9,opt,name=acceptInviteResponse" json:"acceptInviteResponse,omitempty"`
	FeedPostFetchResponse  *FeedPostFetchResponse  `protobuf:"bytes,10,opt,name=feedPostFetchResponse" json:"feedPostFetchResponse,omitempty"`
	FeedPostUpdateResponse *FeedPostUpdateResponse `protobuf:"bytes,11,opt,name=feedPostUpdateResponse" json:"feedPostUpdateResponse,omitempty"`
	AutocompleteResponse   *AutocompleteResponse   `protobuf:"bytes,12,opt,name=autocompleteResponse" json:"autocompleteResponse,omitempty"`
	UserSearchResponse     *UserSearchResponse     `protobuf:"bytes,13,opt,name=userSearchResponse" json:"userSearchResponse,omitempty"`
	ConversationResponse   *ConversationResponse   `protobuf:"bytes,14,opt,name=conversationResponse" json:"conversationResponse,omitempty"`
	FetchConversations     *FetchConversations     `protobuf:"bytes,15,opt,name=fetchConversations" json:"fetchConversations,omitempty"`
	UserCardInfo           *UserCardInfo           `protobuf:"bytes,16,opt,name=userCardInfo" json:"userCardInfo,omitempty"`
	XXX_unrecognized       []byte                  `json:"-"`
}

func (m *ResponseType) Reset()                    { *m = ResponseType{} }
func (m *ResponseType) String() string            { return proto.CompactTextString(m) }
func (*ResponseType) ProtoMessage()               {}
func (*ResponseType) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{9} }

func (m *ResponseType) GetSessionResponse() *SessionResponse {
	if m != nil {
		return m.SessionResponse
	}
	return nil
}

func (m *ResponseType) GetUserEventBatchResponse() *UserEventBatchResponse {
	if m != nil {
		return m.UserEventBatchResponse
	}
	return nil
}

func (m *ResponseType) GetUserProfileUpdate() *UserProfileUpdate {
	if m != nil {
		return m.UserProfileUpdate
	}
	return nil
}

func (m *ResponseType) GetUserProfileQuery() *UserProfileQuery {
	if m != nil {
		return m.UserProfileQuery
	}
	return nil
}

func (m *ResponseType) GetConfirmationRequest() *ConfirmationRequest {
	if m != nil {
		return m.ConfirmationRequest
	}
	return nil
}

func (m *ResponseType) GetUserMessageUpdate() *UserMessageUpdate {
	if m != nil {
		return m.UserMessageUpdate
	}
	return nil
}

func (m *ResponseType) GetDebugMessage() *DebugMessage {
	if m != nil {
		return m.DebugMessage
	}
	return nil
}

func (m *ResponseType) GetImageUploadReply() *ImageUpload {
	if m != nil {
		return m.ImageUploadReply
	}
	return nil
}

func (m *ResponseType) GetAcceptInviteResponse() *AcceptInviteResponse {
	if m != nil {
		return m.AcceptInviteResponse
	}
	return nil
}

func (m *ResponseType) GetFeedPostFetchResponse() *FeedPostFetchResponse {
	if m != nil {
		return m.FeedPostFetchResponse
	}
	return nil
}

func (m *ResponseType) GetFeedPostUpdateResponse() *FeedPostUpdateResponse {
	if m != nil {
		return m.FeedPostUpdateResponse
	}
	return nil
}

func (m *ResponseType) GetAutocompleteResponse() *AutocompleteResponse {
	if m != nil {
		return m.AutocompleteResponse
	}
	return nil
}

func (m *ResponseType) GetUserSearchResponse() *UserSearchResponse {
	if m != nil {
		return m.UserSearchResponse
	}
	return nil
}

func (m *ResponseType) GetConversationResponse() *ConversationResponse {
	if m != nil {
		return m.ConversationResponse
	}
	return nil
}

func (m *ResponseType) GetFetchConversations() *FetchConversations {
	if m != nil {
		return m.FetchConversations
	}
	return nil
}

func (m *ResponseType) GetUserCardInfo() *UserCardInfo {
	if m != nil {
		return m.UserCardInfo
	}
	return nil
}

type ServerResponse struct {
	ResponseCode     *ResponseCode `protobuf:"varint,1,opt,name=responseCode,enum=BlitzMessage.ResponseCode" json:"responseCode,omitempty"`
	ResponseMessage  *string       `protobuf:"bytes,2,opt,name=responseMessage" json:"responseMessage,omitempty"`
	ResponseType     *ResponseType `protobuf:"bytes,3,opt,name=responseType" json:"responseType,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *ServerResponse) Reset()                    { *m = ServerResponse{} }
func (m *ServerResponse) String() string            { return proto.CompactTextString(m) }
func (*ServerResponse) ProtoMessage()               {}
func (*ServerResponse) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{10} }

func (m *ServerResponse) GetResponseCode() ResponseCode {
	if m != nil && m.ResponseCode != nil {
		return *m.ResponseCode
	}
	return ResponseCode_RCSuccess
}

func (m *ServerResponse) GetResponseMessage() string {
	if m != nil && m.ResponseMessage != nil {
		return *m.ResponseMessage
	}
	return ""
}

func (m *ServerResponse) GetResponseType() *ResponseType {
	if m != nil {
		return m.ResponseType
	}
	return nil
}

func init() {
	proto.RegisterType((*DebugMessage)(nil), "BlitzMessage.DebugMessage")
	proto.RegisterType((*SessionRequest)(nil), "BlitzMessage.SessionRequest")
	proto.RegisterType((*BlitzHereAppOptions)(nil), "BlitzMessage.BlitzHereAppOptions")
	proto.RegisterType((*AppOptions)(nil), "BlitzMessage.AppOptions")
	proto.RegisterType((*SessionResponse)(nil), "BlitzMessage.SessionResponse")
	proto.RegisterType((*PushConnect)(nil), "BlitzMessage.PushConnect")
	proto.RegisterType((*PushDisconnect)(nil), "BlitzMessage.PushDisconnect")
	proto.RegisterType((*RequestType)(nil), "BlitzMessage.RequestType")
	proto.RegisterType((*ServerRequest)(nil), "BlitzMessage.ServerRequest")
	proto.RegisterType((*ResponseType)(nil), "BlitzMessage.ResponseType")
	proto.RegisterType((*ServerResponse)(nil), "BlitzMessage.ServerResponse")
	proto.RegisterEnum("BlitzMessage.ResponseCode", ResponseCode_name, ResponseCode_value)
}

var fileDescriptor6 = []byte{
	// 1181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xd4, 0x57, 0xdb, 0x72, 0xe3, 0x34,
	0x18, 0x26, 0xcd, 0xf6, 0x90, 0x3f, 0x27, 0x57, 0x3d, 0xac, 0xe9, 0xee, 0xc2, 0x6e, 0x60, 0x98,
	0xce, 0x0e, 0x74, 0x19, 0x4e, 0xc3, 0x00, 0xd3, 0xa1, 0x4d, 0xda, 0xd9, 0xcc, 0x14, 0xb6, 0xe4,
	0x30, 0x5c, 0xbb, 0xb6, 0xda, 0x1a, 0x1c, 0xcb, 0x48, 0x72, 0x21, 0xfb, 0x04, 0xdc, 0x70, 0xc3,
	0x05, 0x4f, 0xc1, 0x9b, 0x70, 0xb7, 0xaf, 0xc0, 0x8b, 0x20, 0xcb, 0x72, 0x22, 0x39, 0x4a, 0xdb,
	0x5b, 0xee, 0x6c, 0xe9, 0xfb, 0x7f, 0xfd, 0x87, 0xef, 0xff, 0x64, 0x43, 0x63, 0x88, 0xe9, 0x0d,
	0xa6, 0x07, 0x09, 0x25, 0x9c, 0xa0, 0xc6, 0x71, 0x14, 0xf2, 0xd7, 0xdf, 0x61, 0xc6, 0xbc, 0x2b,
	0xbc, 0xf7, 0x88, 0x5c, 0xfc, 0x84, 0x7d, 0x1e, 0xde, 0x60, 0xff, 0xa3, 0x00, 0x33, 0x9f, 0x86,
	0x09, 0x27, 0x0a, 0xba, 0x57, 0x1f, 0x4d, 0x13, 0xcc, 0xd4, 0x4b, 0xa3, 0x87, 0x6f, 0x42, 0x1f,
	0xab, 0x37, 0xe7, 0x24, 0xe6, 0x21, 0x9f, 0x8e, 0xbc, 0xab, 0x62, 0x1f, 0x4e, 0x31, 0x0e, 0xd4,
	0x73, 0xf3, 0x94, 0x86, 0x38, 0x0e, 0x8a, 0xad, 0xd6, 0xb9, 0x37, 0x9d, 0xe0, 0x98, 0xcf, 0x5c,
	0x0d, 0xb1, 0x47, 0xfd, 0xeb, 0xc2, 0xd5, 0x98, 0x61, 0x7a, 0x72, 0xa3, 0xed, 0xa3, 0x6c, 0x45,
	0x45, 0x68, 0xac, 0x9d, 0x53, 0x72, 0x19, 0x46, 0xc5, 0x5a, 0xe7, 0x19, 0x88, 0xa0, 0x2e, 0xd2,
	0x2b, 0x05, 0x45, 0x9b, 0x50, 0x0b, 0xb2, 0xf7, 0x11, 0xfe, 0x8d, 0xbb, 0x95, 0xa7, 0xd5, 0xfd,
	0x5a, 0xe7, 0x9f, 0x0a, 0xb4, 0x86, 0x62, 0x3b, 0x24, 0xf1, 0x00, 0xff, 0x92, 0x62, 0xc6, 0xd1,
	0x3e, 0x6c, 0x44, 0xc4, 0xf7, 0xb8, 0x58, 0x12, 0xa0, 0xca, 0x7e, 0xfd, 0x93, 0xdd, 0x03, 0xbd,
	0x26, 0x07, 0x67, 0x6a, 0x17, 0x7d, 0x08, 0x10, 0xc8, 0xa4, 0xfb, 0xf1, 0x25, 0x71, 0x57, 0x24,
	0xd6, 0x35, 0xb1, 0xbd, 0xd9, 0x3e, 0x7a, 0x0e, 0xeb, 0x49, 0x1e, 0x9f, 0x5b, 0x95, 0xd0, 0xb7,
	0x4d, 0xa8, 0x96, 0x00, 0xfa, 0x1c, 0xb6, 0x23, 0x8f, 0xf1, 0xa3, 0x24, 0xe9, 0x79, 0xdc, 0x1b,
	0x60, 0x86, 0xb9, 0x78, 0xc0, 0xee, 0x03, 0x69, 0xf8, 0xd0, 0x34, 0x1c, 0x85, 0x13, 0x11, 0xb7,
	0x37, 0x49, 0x3a, 0x3b, 0xb0, 0x25, 0x77, 0x5e, 0x62, 0x8a, 0x85, 0xed, 0xab, 0x24, 0x0b, 0x93,
	0x75, 0xfa, 0x00, 0xf3, 0x37, 0xf4, 0x35, 0x38, 0x17, 0x05, 0x48, 0xad, 0xa9, 0x3c, 0x9f, 0x99,
	0x7e, 0x6d, 0xae, 0xfe, 0x5e, 0x81, 0xf6, 0xac, 0x5e, 0x2c, 0x11, 0x4b, 0x18, 0xb5, 0x60, 0x2d,
	0x15, 0xb1, 0xf7, 0x7b, 0xd2, 0x4d, 0x0d, 0x6d, 0x43, 0x83, 0xe5, 0x90, 0x11, 0xf9, 0x19, 0xc7,
	0xb2, 0x30, 0xb5, 0xac, 0xf8, 0x4c, 0xf2, 0x6c, 0x3c, 0x38, 0x93, 0x05, 0xa8, 0xa1, 0x17, 0xd0,
	0x48, 0xb5, 0x4e, 0x8a, 0xec, 0xaa, 0xf6, 0xb2, 0x14, 0x0d, 0x3c, 0x80, 0x7a, 0x3a, 0xaf, 0x92,
	0xbb, 0x7a, 0x57, 0x19, 0x1f, 0x42, 0x9b, 0x66, 0xb5, 0x3b, 0x8a, 0x22, 0x55, 0x4a, 0x77, 0x4d,
	0xd8, 0x6c, 0xa0, 0x2f, 0xa1, 0x19, 0xc6, 0x37, 0x21, 0xc7, 0xaa, 0xe9, 0xee, 0xba, 0xad, 0x00,
	0x47, 0xbe, 0x8f, 0x13, 0xde, 0xd7, 0x81, 0x59, 0xcf, 0xbd, 0x59, 0x39, 0xdc, 0x0d, 0x5b, 0xcf,
	0xb5, 0x72, 0x3d, 0x81, 0xfa, 0x79, 0xca, 0xae, 0xbb, 0x24, 0x8e, 0xc5, 0x14, 0x95, 0x2b, 0xd5,
	0x71, 0xa0, 0x95, 0x6d, 0xf7, 0x42, 0xe6, 0xe7, 0x88, 0xce, 0x9f, 0x00, 0x75, 0x75, 0x54, 0x36,
	0x5c, 0xe8, 0x33, 0x68, 0x31, 0x83, 0x9e, 0xaa, 0x55, 0x8f, 0xcd, 0x23, 0x4b, 0x14, 0x16, 0x56,
	0x69, 0x31, 0x34, 0xc7, 0x1e, 0xf7, 0xaf, 0x15, 0x39, 0x1f, 0x2f, 0x96, 0x6a, 0x8e, 0x41, 0x5f,
	0xc1, 0xa6, 0x56, 0xdd, 0x71, 0x12, 0x64, 0x8c, 0xcb, 0xa9, 0xfa, 0xee, 0xd2, 0x1a, 0xe7, 0x30,
	0x51, 0x50, 0x47, 0xb3, 0xfd, 0x21, 0xc5, 0x74, 0xaa, 0xc8, 0xfa, 0xce, 0x52, 0x53, 0x89, 0x42,
	0x87, 0xb0, 0x25, 0x92, 0xbf, 0x0c, 0xe9, 0x44, 0x0e, 0x55, 0x91, 0xe6, 0xaa, 0xad, 0x21, 0xdd,
	0x45, 0xa0, 0xa0, 0x33, 0x9a, 0xe4, 0xdb, 0x43, 0x21, 0x29, 0x85, 0xf9, 0xda, 0xb2, 0xb0, 0xd5,
	0xb3, 0x0a, 0xfb, 0x1b, 0xd8, 0x52, 0xc6, 0xa7, 0x58, 0x94, 0xc0, 0x64, 0xc3, 0x9d, 0xd6, 0x1f,
	0x43, 0x23, 0xd0, 0xf4, 0x45, 0xb1, 0x61, 0xaf, 0xac, 0x00, 0x9a, 0x02, 0x09, 0x02, 0x87, 0x13,
	0xe9, 0x20, 0x22, 0x5e, 0xe0, 0xd6, 0x6c, 0x04, 0xee, 0xcf, 0x01, 0x59, 0x71, 0xbc, 0x45, 0x12,
	0xba, 0x70, 0x5f, 0xb6, 0x7e, 0x0b, 0xdb, 0x97, 0x42, 0x76, 0xcf, 0x09, 0xe3, 0x46, 0x82, 0x75,
	0xe9, 0xa0, 0x63, 0x3a, 0x38, 0xb5, 0x20, 0xd1, 0x31, 0xec, 0x14, 0x1e, 0xf2, 0xac, 0x0b, 0x17,
	0x0d, 0xe9, 0xe2, 0x3d, 0xbb, 0x0b, 0x03, 0x2a, 0xb3, 0x48, 0x39, 0xf1, 0xc9, 0x24, 0x89, 0xf0,
	0xdc, 0x43, 0xd3, 0x9a, 0xc5, 0x22, 0x50, 0xd0, 0xb9, 0x8d, 0x8b, 0xeb, 0x44, 0xd1, 0xb2, 0x25,
	0x6d, 0x1f, 0x99, 0xb6, 0xb3, 0x3b, 0xe7, 0x2c, 0x14, 0x56, 0x8a, 0xce, 0xf9, 0x5d, 0x52, 0x9c,
	0xd9, 0x5e, 0xd6, 0x59, 0x03, 0x96, 0xf5, 0x29, 0x99, 0xcf, 0xad, 0xeb, 0xd8, 0xfa, 0xa4, 0x0f,
	0xb6, 0x18, 0xb8, 0xc4, 0x18, 0x64, 0x77, 0xd3, 0x36, 0x70, 0xe6, 0xb0, 0x2b, 0xea, 0x0b, 0x49,
	0x64, 0x06, 0xf5, 0xd1, 0x12, 0xea, 0x97, 0x81, 0x82, 0xbd, 0xe8, 0x32, 0xeb, 0x95, 0xbe, 0xc7,
	0xdc, 0x2d, 0x69, 0xfe, 0xb4, 0xdc, 0x98, 0x32, 0x2e, 0x53, 0xb2, 0xac, 0x3e, 0x03, 0x71, 0x43,
	0xe1, 0x5f, 0xdd, 0x6d, 0x9b, 0x92, 0x8d, 0x67, 0xfb, 0xe8, 0x25, 0xb8, 0xa9, 0x2c, 0xbd, 0xee,
	0x64, 0xc8, 0x3d, 0x9e, 0x32, 0x77, 0x47, 0xda, 0x7e, 0x50, 0xb2, 0x5d, 0x82, 0xce, 0xa6, 0x26,
	0x3b, 0xb7, 0xeb, 0xd1, 0x40, 0xde, 0x9b, 0xbb, 0xb6, 0xa9, 0x19, 0x6b, 0x88, 0xce, 0x18, 0x9a,
	0xf9, 0x27, 0x4a, 0x91, 0x78, 0xf9, 0x86, 0xc9, 0xef, 0x1d, 0xd1, 0x34, 0x3a, 0x97, 0x4e, 0x25,
	0x79, 0xa5, 0xa6, 0x69, 0xda, 0xda, 0xf9, 0x77, 0x1d, 0x1a, 0xc5, 0x25, 0x26, 0xc5, 0xf6, 0x0b,
	0x68, 0x33, 0xf3, 0x6e, 0x53, 0x6a, 0xfb, 0x64, 0x89, 0xda, 0xaa, 0x0b, 0xb0, 0x07, 0xbb, 0xa6,
	0xdc, 0xce, 0xcc, 0xf3, 0x18, 0xde, 0xbf, 0x4d, 0x76, 0x67, 0x5e, 0xfe, 0x9f, 0xf2, 0xab, 0xa2,
	0x36, 0x84, 0xf1, 0xbe, 0xea, 0x5b, 0xd6, 0xcf, 0xf5, 0x3b, 0xf5, 0xf3, 0x53, 0x70, 0x34, 0xfd,
	0x1c, 0xe0, 0x24, 0x9a, 0x2a, 0xd5, 0xbd, 0x45, 0x44, 0x85, 0x08, 0x9a, 0x22, 0xaa, 0x9a, 0x53,
	0xb3, 0x89, 0xe0, 0x91, 0x05, 0xa9, 0x8b, 0xa0, 0x12, 0x47, 0xe5, 0x02, 0x6e, 0x13, 0x41, 0x03,
	0x9a, 0x91, 0xa4, 0x2c, 0xa4, 0xca, 0x49, 0xdd, 0x46, 0x92, 0x53, 0x2b, 0x56, 0xe6, 0x62, 0x28,
	0xa4, 0xf2, 0xd1, 0xb0, 0xe6, 0x62, 0x41, 0x66, 0xa2, 0xa1, 0xcb, 0xa2, 0xb2, 0x6f, 0xda, 0x44,
	0x63, 0xbc, 0x80, 0xcb, 0xce, 0x37, 0x25, 0x4b, 0xd9, 0xb7, 0x6c, 0xe7, 0x77, 0x2d, 0xc8, 0x25,
	0xa2, 0xd5, 0xbe, 0xa7, 0x68, 0x95, 0xc5, 0xc3, 0xb9, 0x53, 0x3c, 0xfe, 0x90, 0x5f, 0xf8, 0xb9,
	0x7a, 0xa8, 0x10, 0x84, 0x13, 0xaa, 0x9e, 0xbb, 0x24, 0xc8, 0x87, 0xbc, 0x55, 0x76, 0x32, 0xd0,
	0x10, 0xea, 0x43, 0x52, 0xbe, 0x17, 0x64, 0xcd, 0xbf, 0x6a, 0x35, 0x57, 0x52, 0x74, 0xaa, 0xb6,
	0x78, 0x74, 0x91, 0x79, 0xfe, 0x57, 0x65, 0xae, 0x3a, 0xd2, 0x77, 0x13, 0x6a, 0x83, 0xee, 0x30,
	0x15, 0xbc, 0x63, 0xcc, 0xa9, 0x20, 0x04, 0xad, 0x41, 0xb7, 0x1f, 0x27, 0x29, 0xef, 0x12, 0x4a,
	0xd3, 0x84, 0x3b, 0x2b, 0xda, 0x9a, 0x20, 0xa6, 0x17, 0x85, 0x81, 0x53, 0x45, 0x5b, 0xd0, 0x16,
	0x66, 0x32, 0xb1, 0x1f, 0x3d, 0x1a, 0x87, 0xf1, 0x95, 0xf3, 0x40, 0x7c, 0x64, 0x37, 0x8b, 0xc5,
	0x13, 0x4a, 0x09, 0x75, 0x56, 0x73, 0xdc, 0xf7, 0x84, 0x0b, 0x32, 0x5c, 0x13, 0x1a, 0xbe, 0xc6,
	0x81, 0xb3, 0x96, 0x3b, 0xec, 0x46, 0xe2, 0x27, 0x8c, 0x8f, 0x08, 0x79, 0x15, 0x05, 0xce, 0xfa,
	0xf1, 0x8b, 0x37, 0x87, 0x2b, 0xf0, 0xd6, 0x9b, 0xc3, 0x2a, 0xaa, 0x1c, 0x8b, 0x47, 0xb7, 0x02,
	0x7b, 0x82, 0x3c, 0x07, 0xf2, 0x5f, 0xe1, 0x5a, 0xfc, 0x05, 0x18, 0x39, 0xfd, 0x5e, 0xa9, 0xfc,
	0x17, 0x00, 0x00, 0xff, 0xff, 0x0e, 0xd5, 0x19, 0xe1, 0x38, 0x0e, 0x00, 0x00,
}
