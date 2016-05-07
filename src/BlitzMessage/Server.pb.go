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
	ChargeRequest            *Charge                   `protobuf:"bytes,23,opt,name=chargeRequest" json:"chargeRequest,omitempty"`
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

func (m *RequestType) GetChargeRequest() *Charge {
	if m != nil {
		return m.ChargeRequest
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
	ChargeResponse         *Charge                 `protobuf:"bytes,17,opt,name=chargeResponse" json:"chargeResponse,omitempty"`
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

func (m *ResponseType) GetChargeResponse() *Charge {
	if m != nil {
		return m.ChargeResponse
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
	// 1209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xd4, 0x57, 0xcb, 0x72, 0xe4, 0x34,
	0x14, 0xa5, 0xd3, 0x93, 0x47, 0xdf, 0x7e, 0x39, 0xca, 0xcb, 0x64, 0x66, 0x60, 0xa6, 0xa1, 0xa8,
	0xd4, 0x30, 0x64, 0x28, 0x5e, 0x45, 0x01, 0x95, 0x22, 0xe9, 0x4e, 0x6a, 0xba, 0x2a, 0x30, 0xa1,
	0x1f, 0xc5, 0xda, 0xb1, 0x95, 0xb4, 0xc1, 0x6d, 0x19, 0x49, 0x0e, 0x64, 0xbe, 0x80, 0x0d, 0x0b,
	0x36, 0x7c, 0x05, 0xff, 0xc0, 0x07, 0xb0, 0x9b, 0x2f, 0x42, 0x96, 0xe5, 0x6e, 0xc9, 0x51, 0x27,
	0xd9, 0xb2, 0xb3, 0xa5, 0x73, 0xaf, 0xee, 0xe3, 0xdc, 0x23, 0x1b, 0x1a, 0x43, 0x4c, 0xaf, 0x30,
	0xdd, 0x4f, 0x28, 0xe1, 0x04, 0x35, 0x8e, 0xa2, 0x90, 0xbf, 0xfe, 0x0e, 0x33, 0xe6, 0x5d, 0xe2,
	0xdd, 0x87, 0xe4, 0xfc, 0x27, 0xec, 0xf3, 0xf0, 0x0a, 0xfb, 0x1f, 0x05, 0x98, 0xf9, 0x34, 0x4c,
	0x38, 0x51, 0xd0, 0xdd, 0xfa, 0xe8, 0x3a, 0xc1, 0x4c, 0xbd, 0x34, 0x7a, 0xf8, 0x2a, 0xf4, 0xb1,
	0x7a, 0x73, 0x8e, 0x63, 0x1e, 0xf2, 0xeb, 0x91, 0x77, 0x59, 0xec, 0xc3, 0x09, 0xc6, 0x81, 0x7a,
	0x6e, 0x9e, 0xd0, 0x10, 0xc7, 0x41, 0xb1, 0xd5, 0x3a, 0xf3, 0xae, 0xa7, 0x38, 0xe6, 0x33, 0x57,
	0x43, 0xec, 0x51, 0x7f, 0x52, 0xb8, 0x1a, 0x33, 0x4c, 0x8f, 0xaf, 0xb4, 0x7d, 0x94, 0xad, 0xa8,
	0x08, 0x8d, 0xb5, 0x33, 0x4a, 0x2e, 0xc2, 0xa8, 0x58, 0xeb, 0x3c, 0x05, 0x11, 0xd4, 0x79, 0x7a,
	0xa9, 0xa0, 0x68, 0x1d, 0x6a, 0x41, 0xf6, 0x3e, 0xc2, 0xbf, 0x71, 0xb7, 0xf2, 0xa4, 0xba, 0x57,
	0xeb, 0xfc, 0x5b, 0x81, 0xd6, 0x50, 0x6c, 0x87, 0x24, 0x1e, 0xe0, 0x5f, 0x52, 0xcc, 0x38, 0xda,
	0x83, 0xb5, 0x88, 0xf8, 0x1e, 0x17, 0x4b, 0x02, 0x54, 0xd9, 0xab, 0x7f, 0xb2, 0xbd, 0xaf, 0xd7,
	0x64, 0xff, 0x54, 0xed, 0xa2, 0xe7, 0x00, 0x81, 0x4c, 0xba, 0x1f, 0x5f, 0x10, 0x77, 0x49, 0x62,
	0x5d, 0x13, 0xdb, 0x9b, 0xed, 0xa3, 0x67, 0xb0, 0x9a, 0xe4, 0xf1, 0xb9, 0x55, 0x09, 0x7d, 0xdb,
	0x84, 0x6a, 0x09, 0xa0, 0xcf, 0x61, 0x33, 0xf2, 0x18, 0x3f, 0x4c, 0x92, 0x9e, 0xc7, 0xbd, 0x01,
	0x66, 0x98, 0x8b, 0x07, 0xec, 0x3e, 0x90, 0x86, 0x3b, 0xa6, 0xe1, 0x28, 0x9c, 0x8a, 0xb8, 0xbd,
	0x69, 0xd2, 0xd9, 0x82, 0x0d, 0xb9, 0xf3, 0x12, 0x53, 0x2c, 0x6c, 0x5f, 0x25, 0x59, 0x98, 0xac,
	0xd3, 0x07, 0x98, 0xbf, 0xa1, 0xaf, 0xc1, 0x39, 0x2f, 0x40, 0x6a, 0x4d, 0xe5, 0xf9, 0xd4, 0xf4,
	0x6b, 0x73, 0xf5, 0xf7, 0x12, 0xb4, 0x67, 0xf5, 0x62, 0x89, 0x58, 0xc2, 0xa8, 0x05, 0x2b, 0xa9,
	0x88, 0xbd, 0xdf, 0x93, 0x6e, 0x6a, 0x68, 0x13, 0x1a, 0x2c, 0x87, 0x8c, 0xc8, 0xcf, 0x38, 0x96,
	0x85, 0xa9, 0x65, 0xc5, 0x67, 0x92, 0x67, 0xe3, 0xc1, 0xa9, 0x2c, 0x40, 0x0d, 0xbd, 0x80, 0x46,
	0xaa, 0x75, 0x52, 0x64, 0x57, 0xb5, 0x97, 0xa5, 0x68, 0xe0, 0x3e, 0xd4, 0xd3, 0x79, 0x95, 0xdc,
	0xe5, 0xbb, 0xca, 0xb8, 0x03, 0x6d, 0x9a, 0xd5, 0xee, 0x30, 0x8a, 0x54, 0x29, 0xdd, 0x15, 0x61,
	0xb3, 0x86, 0xbe, 0x84, 0x66, 0x18, 0x5f, 0x85, 0x1c, 0xab, 0xa6, 0xbb, 0xab, 0xb6, 0x02, 0x1c,
	0xfa, 0x3e, 0x4e, 0x78, 0x5f, 0x07, 0x66, 0x3d, 0xf7, 0x66, 0xe5, 0x70, 0xd7, 0x6c, 0x3d, 0xd7,
	0xca, 0xf5, 0x18, 0xea, 0x67, 0x29, 0x9b, 0x74, 0x49, 0x1c, 0x8b, 0x29, 0x2a, 0x57, 0xaa, 0xe3,
	0x40, 0x2b, 0xdb, 0xee, 0x85, 0xcc, 0xcf, 0x11, 0x9d, 0x7f, 0x00, 0xea, 0xea, 0xa8, 0x6c, 0xb8,
	0xd0, 0x67, 0xd0, 0x62, 0x06, 0x3d, 0x55, 0xab, 0x1e, 0x99, 0x47, 0x96, 0x28, 0x2c, 0xac, 0xd2,
	0x62, 0x68, 0x8e, 0x3c, 0xee, 0x4f, 0x14, 0x39, 0x1f, 0xdd, 0x2c, 0xd5, 0x1c, 0x83, 0xbe, 0x82,
	0x75, 0xad, 0xba, 0xe3, 0x24, 0xc8, 0x18, 0x97, 0x53, 0xf5, 0xdd, 0x85, 0x35, 0xce, 0x61, 0xa2,
	0xa0, 0x8e, 0x66, 0xfb, 0x43, 0x8a, 0xe9, 0xb5, 0x22, 0xeb, 0x3b, 0x0b, 0x4d, 0x25, 0x0a, 0x1d,
	0xc0, 0x86, 0x48, 0xfe, 0x22, 0xa4, 0x53, 0x39, 0x54, 0x45, 0x9a, 0xcb, 0xb6, 0x86, 0x74, 0x6f,
	0x02, 0x05, 0x9d, 0xd1, 0x34, 0xdf, 0x1e, 0x0a, 0x49, 0x29, 0xcc, 0x57, 0x16, 0x85, 0xad, 0x9e,
	0x55, 0xd8, 0xdf, 0xc0, 0x86, 0x32, 0x3e, 0xc1, 0xa2, 0x04, 0x26, 0x1b, 0xee, 0xb4, 0xfe, 0x18,
	0x1a, 0x81, 0xa6, 0x2f, 0x8a, 0x0d, 0xbb, 0x65, 0x05, 0xd0, 0x14, 0x48, 0x10, 0x38, 0x9c, 0x4a,
	0x07, 0x11, 0xf1, 0x02, 0xb7, 0x66, 0x23, 0x70, 0x7f, 0x0e, 0xc8, 0x8a, 0xe3, 0xdd, 0x24, 0xa1,
	0x0b, 0xf7, 0x65, 0xeb, 0xb7, 0xb0, 0x79, 0x21, 0x64, 0xf7, 0x8c, 0x30, 0x6e, 0x24, 0x58, 0x97,
	0x0e, 0x3a, 0xa6, 0x83, 0x13, 0x0b, 0x12, 0x1d, 0xc1, 0x56, 0xe1, 0x21, 0xcf, 0xba, 0x70, 0xd1,
	0x90, 0x2e, 0xde, 0xb3, 0xbb, 0x30, 0xa0, 0x32, 0x8b, 0x94, 0x13, 0x9f, 0x4c, 0x93, 0x08, 0xcf,
	0x3d, 0x34, 0xad, 0x59, 0xdc, 0x04, 0x0a, 0x3a, 0xb7, 0x71, 0x71, 0x9d, 0x28, 0x5a, 0xb6, 0xa4,
	0xed, 0x43, 0xd3, 0x76, 0x76, 0xe7, 0x9c, 0x86, 0xc2, 0x4a, 0xd1, 0x39, 0xbf, 0x4b, 0x8a, 0x33,
	0xdb, 0x8b, 0x3a, 0x6b, 0xc0, 0xb2, 0x3e, 0x25, 0xf3, 0xb9, 0x75, 0x1d, 0x5b, 0x9f, 0xf4, 0xc1,
	0x16, 0x03, 0x97, 0x18, 0x83, 0xec, 0xae, 0xdb, 0x06, 0xce, 0x1c, 0x76, 0x45, 0x7d, 0x21, 0x89,
	0xcc, 0xa0, 0x3e, 0x5a, 0x40, 0xfd, 0x32, 0x50, 0xb0, 0x17, 0x5d, 0x64, 0xbd, 0xd2, 0xf7, 0x98,
	0xbb, 0x21, 0xcd, 0x9f, 0x94, 0x1b, 0x53, 0xc6, 0x65, 0x4a, 0x96, 0xd5, 0x67, 0x20, 0x6e, 0x28,
	0xfc, 0xab, 0xbb, 0x69, 0x53, 0xb2, 0xf1, 0x6c, 0x1f, 0xbd, 0x04, 0x37, 0x95, 0xa5, 0xd7, 0x9d,
	0x0c, 0xb9, 0xc7, 0x53, 0xe6, 0x6e, 0x49, 0xdb, 0x0f, 0x4a, 0xb6, 0x0b, 0xd0, 0xd9, 0xd4, 0x64,
	0xe7, 0x76, 0x3d, 0x1a, 0xc8, 0x7b, 0x73, 0xdb, 0x36, 0x35, 0x63, 0x0d, 0x81, 0x3e, 0x84, 0xa6,
	0x3f, 0xf1, 0xe8, 0xe5, 0x8c, 0x39, 0x3b, 0xd2, 0x64, 0xb3, 0x54, 0x21, 0x09, 0xe9, 0x8c, 0xa1,
	0x99, 0x7f, 0xcf, 0x14, 0x55, 0x2a, 0x5f, 0x47, 0xf9, 0x25, 0x25, 0x3a, 0x4c, 0xe7, 0x3a, 0xab,
	0xf4, 0xb1, 0xd4, 0x61, 0x4d, 0x88, 0x3b, 0x7f, 0xae, 0x41, 0xa3, 0xb8, 0xf1, 0xa4, 0x32, 0x7f,
	0x01, 0x6d, 0x66, 0x5e, 0x84, 0x4a, 0x9a, 0x1f, 0x2f, 0x90, 0x66, 0x75, 0x5b, 0xf6, 0x60, 0xdb,
	0xd4, 0xe6, 0x99, 0x79, 0x1e, 0xc3, 0xfb, 0xb7, 0x69, 0xf4, 0xcc, 0xcb, 0xff, 0x53, 0xab, 0x55,
	0xd4, 0x86, 0x8a, 0xde, 0x57, 0xaa, 0xcb, 0x62, 0xbb, 0x7a, 0xa7, 0xd8, 0x7e, 0x0a, 0x8e, 0x26,
	0xb6, 0x03, 0x9c, 0x44, 0xd7, 0x4a, 0xa2, 0x6f, 0x51, 0x5c, 0xa1, 0x98, 0xa6, 0xe2, 0xaa, 0xe6,
	0xd4, 0x6c, 0x8a, 0x79, 0x68, 0x41, 0xea, 0x8a, 0xa9, 0x94, 0x54, 0xb9, 0x80, 0xdb, 0x14, 0xd3,
	0x80, 0x66, 0x24, 0x29, 0xab, 0xae, 0x72, 0x52, 0xb7, 0x91, 0xe4, 0xc4, 0x8a, 0x95, 0xb9, 0x18,
	0x72, 0xaa, 0x7c, 0x34, 0xac, 0xb9, 0x58, 0x90, 0x99, 0xc2, 0xe8, 0x1a, 0xaa, 0xec, 0x9b, 0x36,
	0x85, 0x19, 0xdf, 0xc0, 0x65, 0xe7, 0x9b, 0xfa, 0xa6, 0xec, 0x5b, 0xb6, 0xf3, 0xbb, 0x16, 0xe4,
	0x02, 0x85, 0x6b, 0xdf, 0x53, 0xe1, 0xca, 0x4a, 0xe3, 0xdc, 0xa9, 0x34, 0xcf, 0xa1, 0x55, 0x28,
	0x8d, 0x8a, 0x75, 0xfd, 0x16, 0xa9, 0xf9, 0x43, 0xfe, 0x3c, 0xe4, 0x5a, 0xa3, 0x02, 0x16, 0x47,
	0x52, 0xf5, 0xdc, 0x25, 0x41, 0x2e, 0x09, 0xad, 0xf2, 0x91, 0x03, 0x0d, 0xa1, 0xbe, 0x51, 0xe5,
	0x7b, 0x41, 0xed, 0xfc, 0x83, 0x59, 0x73, 0x25, 0x25, 0xaa, 0x6a, 0x8b, 0x5e, 0x97, 0xa4, 0x67,
	0x7f, 0x55, 0xe6, 0x1a, 0x25, 0x7d, 0x37, 0xa1, 0x36, 0xe8, 0x0e, 0x53, 0xc1, 0x52, 0xc6, 0x9c,
	0x0a, 0x42, 0xd0, 0x1a, 0x74, 0xfb, 0x71, 0x92, 0xf2, 0x2e, 0xa1, 0x34, 0x4d, 0xb8, 0xb3, 0xa4,
	0xad, 0x09, 0x1a, 0x7b, 0x51, 0x18, 0x38, 0x55, 0xb4, 0x01, 0x6d, 0x61, 0x26, 0x13, 0xfb, 0xd1,
	0xa3, 0x71, 0x18, 0x5f, 0x3a, 0x0f, 0xc4, 0xf7, 0x7b, 0xb3, 0x58, 0x3c, 0xa6, 0x94, 0x50, 0x67,
	0x39, 0xc7, 0x7d, 0x4f, 0xb8, 0xa0, 0xce, 0x84, 0xd0, 0xf0, 0x35, 0x0e, 0x9c, 0x95, 0xdc, 0x61,
	0x37, 0x12, 0xff, 0x77, 0x7c, 0x44, 0xc8, 0xab, 0x28, 0x70, 0x56, 0x8f, 0x5e, 0xbc, 0x39, 0x58,
	0x82, 0xb7, 0xde, 0x1c, 0x54, 0x51, 0xe5, 0x48, 0x3c, 0xba, 0x15, 0xd8, 0x15, 0x54, 0xdb, 0x97,
	0xbf, 0x21, 0x13, 0xf1, 0x83, 0x61, 0xe4, 0xf4, 0x7b, 0xa5, 0xf2, 0x5f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x3f, 0x62, 0x44, 0x42, 0x93, 0x0e, 0x00, 0x00,
}
