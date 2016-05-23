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
	ResponseCode_RCPaymentError  ResponseCode = 8
)

var ResponseCode_name = map[int32]string{
	1: "RCSuccess",
	2: "RCInputCorrupt",
	3: "RCInputInvalid",
	4: "RCServerWarning",
	5: "RCServerError",
	6: "RCNotAuthorized",
	7: "RCClientTooOld",
	8: "RCPaymentError",
}
var ResponseCode_value = map[string]int32{
	"RCSuccess":       1,
	"RCInputCorrupt":  2,
	"RCInputInvalid":  3,
	"RCServerWarning": 4,
	"RCServerError":   5,
	"RCNotAuthorized": 6,
	"RCClientTooOld":  7,
	"RCPaymentError":  8,
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
func (ResponseCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

type DebugMessage struct {
	DebugText        []string `protobuf:"bytes,1,rep,name=debugText" json:"debugText,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DebugMessage) Reset()                    { *m = DebugMessage{} }
func (m *DebugMessage) String() string            { return proto.CompactTextString(m) }
func (*DebugMessage) ProtoMessage()               {}
func (*DebugMessage) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

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
func (*SessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

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
func (*BlitzHereAppOptions) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

type AppOptions struct {
	BlitzHereOptions *BlitzHereAppOptions `protobuf:"bytes,1,opt,name=blitzHereOptions" json:"blitzHereOptions,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *AppOptions) Reset()                    { *m = AppOptions{} }
func (m *AppOptions) String() string            { return proto.CompactTextString(m) }
func (*AppOptions) ProtoMessage()               {}
func (*AppOptions) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *AppOptions) GetBlitzHereOptions() *BlitzHereAppOptions {
	if m != nil {
		return m.BlitzHereOptions
	}
	return nil
}

type SessionResponse struct {
	UserID           *string        `protobuf:"bytes,1,opt,name=userID" json:"userID,omitempty"`
	SessionToken     *string        `protobuf:"bytes,2,opt,name=sessionToken" json:"sessionToken,omitempty"`
	ServerURL        *string        `protobuf:"bytes,3,opt,name=serverURL" json:"serverURL,omitempty"`
	UserMessages     []*UserMessage `protobuf:"bytes,4,rep,name=userMessages" json:"userMessages,omitempty"`
	UserProfile      *UserProfile   `protobuf:"bytes,5,opt,name=userProfile" json:"userProfile,omitempty"`
	ResetAllAppData  *bool          `protobuf:"varint,6,opt,name=resetAllAppData" json:"resetAllAppData,omitempty"`
	InviteRequest    *UserInvite    `protobuf:"bytes,7,opt,name=inviteRequest" json:"inviteRequest,omitempty"`
	AppOptions       *AppOptions    `protobuf:"bytes,8,opt,name=appOptions" json:"appOptions,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *SessionResponse) Reset()                    { *m = SessionResponse{} }
func (m *SessionResponse) String() string            { return proto.CompactTextString(m) }
func (*SessionResponse) ProtoMessage()               {}
func (*SessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

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

func (m *SessionResponse) GetInviteRequest() *UserInvite {
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
	UserID               *string    `protobuf:"bytes,1,opt,name=userID" json:"userID,omitempty"`
	LastMessageTimestamp *Timestamp `protobuf:"bytes,2,opt,name=lastMessageTimestamp" json:"lastMessageTimestamp,omitempty"`
	XXX_unrecognized     []byte     `json:"-"`
}

func (m *PushConnect) Reset()                    { *m = PushConnect{} }
func (m *PushConnect) String() string            { return proto.CompactTextString(m) }
func (*PushConnect) ProtoMessage()               {}
func (*PushConnect) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func (m *PushConnect) GetUserID() string {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return ""
}

func (m *PushConnect) GetLastMessageTimestamp() *Timestamp {
	if m != nil {
		return m.LastMessageTimestamp
	}
	return nil
}

type PushDisconnect struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *PushDisconnect) Reset()                    { *m = PushDisconnect{} }
func (m *PushDisconnect) String() string            { return proto.CompactTextString(m) }
func (*PushDisconnect) ProtoMessage()               {}
func (*PushDisconnect) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{6} }

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
	AcceptInviteRequest      *UserInvite               `protobuf:"bytes,10,opt,name=acceptInviteRequest" json:"acceptInviteRequest,omitempty"`
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
	FriendRequest            *FriendUpdate             `protobuf:"bytes,24,opt,name=friendRequest" json:"friendRequest,omitempty"`
	SearchCategories         *SearchCategories         `protobuf:"bytes,25,opt,name=searchCategories" json:"searchCategories,omitempty"`
	XXX_unrecognized         []byte                    `json:"-"`
}

func (m *RequestType) Reset()                    { *m = RequestType{} }
func (m *RequestType) String() string            { return proto.CompactTextString(m) }
func (*RequestType) ProtoMessage()               {}
func (*RequestType) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{7} }

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

func (m *RequestType) GetAcceptInviteRequest() *UserInvite {
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

func (m *RequestType) GetFriendRequest() *FriendUpdate {
	if m != nil {
		return m.FriendRequest
	}
	return nil
}

func (m *RequestType) GetSearchCategories() *SearchCategories {
	if m != nil {
		return m.SearchCategories
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
func (*ServerRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{8} }

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
	AcceptInviteResponse   *UserInvite             `protobuf:"bytes,9,opt,name=acceptInviteResponse" json:"acceptInviteResponse,omitempty"`
	// optional FeedPostFetchResponse  feedPostFetchResponse   = 10; Deprecated
	// optional FeedPostUpdateResponse feedPostUpdateResponse  = 11; Deprecated
	AutocompleteResponse *AutocompleteResponse `protobuf:"bytes,12,opt,name=autocompleteResponse" json:"autocompleteResponse,omitempty"`
	UserSearchResponse   *UserSearchResponse   `protobuf:"bytes,13,opt,name=userSearchResponse" json:"userSearchResponse,omitempty"`
	ConversationResponse *ConversationResponse `protobuf:"bytes,14,opt,name=conversationResponse" json:"conversationResponse,omitempty"`
	FetchConversations   *FetchConversations   `protobuf:"bytes,15,opt,name=fetchConversations" json:"fetchConversations,omitempty"`
	UserCardInfo         *UserCardInfo         `protobuf:"bytes,16,opt,name=userCardInfo" json:"userCardInfo,omitempty"`
	ChargeResponse       *Charge               `protobuf:"bytes,17,opt,name=chargeResponse" json:"chargeResponse,omitempty"`
	FriendResponse       *FriendUpdate         `protobuf:"bytes,18,opt,name=friendResponse" json:"friendResponse,omitempty"`
	SearchCategories     *SearchCategories     `protobuf:"bytes,19,opt,name=searchCategories" json:"searchCategories,omitempty"`
	FeedPostResponse     *FeedPostResponse     `protobuf:"bytes,20,opt,name=feedPostResponse" json:"feedPostResponse,omitempty"`
	XXX_unrecognized     []byte                `json:"-"`
}

func (m *ResponseType) Reset()                    { *m = ResponseType{} }
func (m *ResponseType) String() string            { return proto.CompactTextString(m) }
func (*ResponseType) ProtoMessage()               {}
func (*ResponseType) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{9} }

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

func (m *ResponseType) GetAcceptInviteResponse() *UserInvite {
	if m != nil {
		return m.AcceptInviteResponse
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

func (m *ResponseType) GetFriendResponse() *FriendUpdate {
	if m != nil {
		return m.FriendResponse
	}
	return nil
}

func (m *ResponseType) GetSearchCategories() *SearchCategories {
	if m != nil {
		return m.SearchCategories
	}
	return nil
}

func (m *ResponseType) GetFeedPostResponse() *FeedPostResponse {
	if m != nil {
		return m.FeedPostResponse
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
func (*ServerResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{10} }

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

var fileDescriptor5 = []byte{
	// 1252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xd4, 0x57, 0xdd, 0x72, 0xdb, 0x44,
	0x14, 0xc6, 0x71, 0xf3, 0xe3, 0x63, 0x5b, 0x56, 0xd6, 0x4e, 0xa2, 0xa6, 0x05, 0x5a, 0xc3, 0x30,
	0x99, 0x52, 0x52, 0x28, 0xd0, 0x61, 0x80, 0xc9, 0x90, 0xd8, 0xc9, 0xd4, 0x33, 0x81, 0x06, 0xff,
	0x0c, 0xd7, 0x8a, 0xb4, 0xb6, 0x05, 0xb2, 0x24, 0x76, 0x57, 0x81, 0xf4, 0x09, 0xb8, 0xe1, 0x29,
	0x3a, 0xdc, 0xf0, 0x2a, 0xdc, 0xf5, 0x89, 0x58, 0xad, 0x56, 0xb6, 0x56, 0x59, 0xc7, 0x19, 0xee,
	0xb8, 0x93, 0x76, 0xbf, 0xef, 0xec, 0xee, 0x39, 0xdf, 0x7e, 0x47, 0x82, 0xda, 0x00, 0x93, 0x2b,
	0x4c, 0x0e, 0x23, 0x12, 0xb2, 0x10, 0xd5, 0x4e, 0x7c, 0x8f, 0xbd, 0xfe, 0x1e, 0x53, 0x6a, 0x4f,
	0xf0, 0xfe, 0x83, 0xf0, 0xf2, 0x67, 0xec, 0x30, 0xef, 0x0a, 0x3b, 0x9f, 0xb8, 0x98, 0x3a, 0xc4,
	0x8b, 0x58, 0x28, 0xa1, 0xfb, 0xd5, 0xe1, 0x75, 0x84, 0xa9, 0x7c, 0xa9, 0x75, 0xf1, 0x95, 0xe7,
	0x60, 0xf9, 0x66, 0x9e, 0x06, 0xcc, 0x63, 0xd7, 0x43, 0x7b, 0x92, 0xcd, 0xc3, 0x19, 0xc6, 0xae,
	0x7c, 0x36, 0x2e, 0xec, 0xeb, 0x19, 0x0e, 0xd8, 0x9c, 0x3b, 0xc0, 0x36, 0x71, 0xa6, 0x19, 0x77,
	0x44, 0x31, 0x39, 0xbd, 0xca, 0xcd, 0xa3, 0x64, 0x44, 0x6e, 0x49, 0x19, 0xbb, 0x20, 0xe1, 0xd8,
	0xf3, 0xb3, 0xb1, 0xf6, 0x63, 0xe0, 0xbb, 0xb8, 0x8c, 0x27, 0x12, 0x8a, 0xb6, 0xa1, 0xe2, 0x26,
	0xef, 0x43, 0xfc, 0x3b, 0xb3, 0x4a, 0x8f, 0xca, 0x07, 0x95, 0xf6, 0x3f, 0x25, 0x30, 0x06, 0x7c,
	0xda, 0x0b, 0x83, 0x3e, 0xfe, 0x35, 0xc6, 0x94, 0xa1, 0x03, 0xd8, 0xf2, 0x43, 0xc7, 0x66, 0x7c,
	0x88, 0x83, 0x4a, 0x07, 0xd5, 0xe7, 0xbb, 0x87, 0xf9, 0x24, 0x1c, 0x9e, 0xcb, 0x59, 0xf4, 0x14,
	0xc0, 0x15, 0xa7, 0xec, 0x05, 0xe3, 0xd0, 0x5a, 0x13, 0x58, 0x4b, 0xc5, 0x76, 0xe7, 0xf3, 0xe8,
	0x09, 0x6c, 0x46, 0xe9, 0xfe, 0xac, 0xb2, 0x80, 0xde, 0x57, 0xa1, 0xb9, 0x03, 0xa0, 0x2f, 0xa1,
	0xe5, 0xdb, 0x94, 0x1d, 0x47, 0x51, 0xd7, 0x66, 0x76, 0x1f, 0x53, 0xcc, 0xf8, 0x03, 0xb6, 0xee,
	0x09, 0xe2, 0x9e, 0x4a, 0x1c, 0x7a, 0x33, 0xbe, 0x6f, 0x7b, 0x16, 0xb5, 0x77, 0xa0, 0x29, 0x66,
	0x5e, 0x62, 0x82, 0x39, 0xf7, 0x55, 0x94, 0x6c, 0x93, 0xb6, 0x7b, 0x00, 0x8b, 0x37, 0xf4, 0x0d,
	0x98, 0x97, 0x19, 0x48, 0x8e, 0xc9, 0x73, 0x3e, 0x56, 0xe3, 0xea, 0x42, 0xbd, 0x59, 0x83, 0xc6,
	0x3c, 0x5f, 0x34, 0xe2, 0x43, 0x18, 0x19, 0xb0, 0x11, 0xf3, 0xbd, 0xf7, 0xba, 0x22, 0x4c, 0x05,
	0xb5, 0xa0, 0x46, 0x53, 0xc8, 0x30, 0xfc, 0x05, 0x07, 0x22, 0x31, 0x95, 0x24, 0xf9, 0x54, 0x08,
	0x6b, 0xd4, 0x3f, 0x17, 0x09, 0xa8, 0xa0, 0x67, 0x50, 0x8b, 0x73, 0x95, 0xe4, 0xa7, 0x2b, 0xeb,
	0xd3, 0x92, 0x15, 0xf0, 0x10, 0xaa, 0xf1, 0x22, 0x4b, 0xd6, 0xfa, 0xaa, 0x34, 0xee, 0x41, 0x83,
	0x24, 0xb9, 0x3b, 0xf6, 0x7d, 0x99, 0x4a, 0x6b, 0x83, 0x73, 0xb6, 0xf8, 0xca, 0x75, 0x2f, 0xb8,
	0xf2, 0x18, 0x96, 0x45, 0xb7, 0x36, 0x75, 0xc5, 0x4b, 0x42, 0xf5, 0x04, 0x2c, 0x29, 0xb5, 0x3d,
	0xcf, 0x82, 0xb5, 0xa5, 0x43, 0xe7, 0xb2, 0x34, 0x84, 0xea, 0x45, 0x4c, 0xa7, 0x9d, 0x30, 0x08,
	0xf8, 0x6d, 0xb9, 0x91, 0x20, 0x59, 0x5d, 0x49, 0x9c, 0x97, 0x4f, 0x2a, 0x68, 0x69, 0x75, 0x4d,
	0x30, 0x92, 0xa8, 0x5d, 0x8f, 0x3a, 0x69, 0xe0, 0xf6, 0x9b, 0x2a, 0x54, 0xe5, 0x09, 0x92, 0xbb,
	0x87, 0xbe, 0x00, 0x83, 0x2a, 0x62, 0x96, 0x85, 0x7d, 0xa8, 0x86, 0x2c, 0x08, 0x9e, 0xb3, 0xe2,
	0xec, 0x8a, 0x9d, 0xd8, 0xcc, 0x99, 0xca, 0x8d, 0x3c, 0xbc, 0x99, 0x8d, 0x05, 0x06, 0x7d, 0x0d,
	0xdb, 0xb9, 0x5a, 0x8c, 0x22, 0x37, 0xd1, 0x67, 0x2a, 0xec, 0xf7, 0x97, 0x56, 0x24, 0x85, 0xa1,
	0xaf, 0xc0, 0xcc, 0x71, 0x7f, 0x8c, 0x31, 0xb9, 0x96, 0xd2, 0x7e, 0x6f, 0x29, 0x55, 0xa0, 0xd0,
	0x11, 0x34, 0xf9, 0xe1, 0xc7, 0x1e, 0x99, 0x89, 0x2b, 0x98, 0x1d, 0x73, 0x5d, 0xa7, 0xdf, 0xce,
	0x4d, 0x20, 0x17, 0x3f, 0x9a, 0xa5, 0xd3, 0x03, 0x1c, 0xb8, 0x19, 0x7d, 0x63, 0xd9, 0xb6, 0xe5,
	0xb3, 0xdc, 0xf6, 0xb7, 0xd0, 0x94, 0xe4, 0x33, 0xcc, 0x53, 0xa0, 0x6a, 0x67, 0x25, 0xfb, 0x53,
	0xa8, 0xb9, 0x39, 0x37, 0x92, 0x22, 0xda, 0x2f, 0xfa, 0x45, 0xce, 0xaf, 0xb8, 0xdc, 0xbd, 0x99,
	0x08, 0xe0, 0x87, 0xb6, 0x6b, 0x55, 0x74, 0x72, 0xef, 0x2d, 0x00, 0x5c, 0x57, 0x4d, 0xdb, 0x71,
	0x70, 0xc4, 0x7a, 0x8a, 0xb6, 0x61, 0x85, 0xb6, 0xbf, 0x83, 0xd6, 0x98, 0x9b, 0xf1, 0x45, 0x48,
	0x99, 0x72, 0xae, 0xaa, 0xe0, 0xb5, 0x55, 0xde, 0x99, 0x06, 0x89, 0x4e, 0x60, 0x27, 0x8b, 0x90,
	0x1e, 0x36, 0x0b, 0x51, 0x13, 0x21, 0x3e, 0xd0, 0x87, 0x50, 0xa0, 0x49, 0x65, 0xed, 0x98, 0x85,
	0x4e, 0x38, 0x8b, 0x7c, 0xbc, 0x88, 0x50, 0xd7, 0x55, 0xf6, 0xf8, 0x26, 0x90, 0xab, 0xb8, 0x81,
	0xb3, 0x26, 0x23, 0xd5, 0x68, 0x08, 0xee, 0x03, 0x95, 0x3b, 0xef, 0x44, 0xe7, 0x1e, 0x67, 0x49,
	0x15, 0xa7, 0x0d, 0x27, 0x5b, 0xb3, 0xb1, 0xac, 0xa0, 0x0a, 0x2c, 0x29, 0x4f, 0xb4, 0xb8, 0xe5,
	0x96, 0xa9, 0x2b, 0x4f, 0xde, 0x06, 0xf8, 0x3d, 0x8b, 0x94, 0xfb, 0x6b, 0x6d, 0xeb, 0xee, 0x99,
	0x7a, 0xc7, 0xa5, 0xe2, 0xb9, 0x6f, 0x52, 0x45, 0xf1, 0x68, 0x89, 0xe2, 0x8b, 0x40, 0x2e, 0x5a,
	0x34, 0x4e, 0x6a, 0x95, 0x9f, 0xa3, 0x56, 0x53, 0xd0, 0x1f, 0x15, 0x0b, 0x53, 0xc4, 0x25, 0xbe,
	0x97, 0xe4, 0xa7, 0xcf, 0xdb, 0x18, 0xfe, 0xcd, 0x6a, 0x2d, 0x53, 0x52, 0x3a, 0x8f, 0x5e, 0x82,
	0x15, 0x8b, 0xd4, 0xe7, 0x83, 0x0c, 0x98, 0xcd, 0x62, 0x6a, 0xed, 0x08, 0xee, 0x47, 0x05, 0xee,
	0x12, 0x74, 0x72, 0x59, 0x92, 0x75, 0x3b, 0x36, 0x71, 0x45, 0x73, 0xdd, 0xd5, 0x5d, 0x96, 0x51,
	0x0e, 0x81, 0x3e, 0x86, 0xba, 0x33, 0xb5, 0xc9, 0x64, 0xae, 0x9c, 0x3d, 0x41, 0x69, 0x15, 0x32,
	0x24, 0x20, 0xe8, 0x33, 0xa8, 0x8f, 0x89, 0x97, 0x73, 0x00, 0x4b, 0x17, 0xff, 0x4c, 0x40, 0x16,
	0x9e, 0x45, 0x45, 0xf9, 0x3b, 0xfc, 0x6d, 0x12, 0xf2, 0x29, 0x6a, 0xdd, 0xd7, 0x79, 0xd6, 0xa0,
	0x80, 0x6a, 0x8f, 0xa0, 0x9e, 0x7e, 0x52, 0x65, 0x25, 0x29, 0x36, 0xc8, 0xb4, 0x2b, 0x70, 0x39,
	0x91, 0x85, 0x97, 0x4b, 0x0f, 0x2e, 0xc8, 0x29, 0x67, 0xf6, 0xed, 0xbf, 0xb6, 0xa0, 0x96, 0xf5,
	0x60, 0xe1, 0xfe, 0x2f, 0xa0, 0x41, 0xd5, 0xd6, 0x2c, 0xed, 0xff, 0xdd, 0x25, 0xf6, 0x2f, 0xfb,
	0x77, 0x17, 0x76, 0x55, 0xff, 0x9f, 0xd3, 0xd3, 0x3d, 0x7c, 0x78, 0x5b, 0x1f, 0x98, 0x47, 0xf9,
	0x7f, 0xf6, 0x03, 0xb9, 0x6b, 0xc5, 0xa9, 0xef, 0xda, 0x0e, 0x8a, 0x86, 0xbe, 0xb9, 0xd2, 0xd0,
	0x3f, 0x07, 0x33, 0x67, 0xe8, 0x7d, 0x1c, 0xf9, 0xd7, 0xb2, 0x0d, 0xdc, 0xe2, 0xea, 0x2f, 0xa0,
	0xa5, 0xba, 0xba, 0x2c, 0x4e, 0x65, 0xb5, 0xad, 0xab, 0x86, 0x2a, 0x79, 0x35, 0x9d, 0xad, 0x1f,
	0x6b, 0x90, 0x89, 0x75, 0xe4, 0xcd, 0x51, 0xf2, 0xeb, 0x3a, 0xeb, 0x18, 0xdd, 0xc0, 0x25, 0xeb,
	0xab, 0xc6, 0x25, 0xf9, 0x86, 0x6e, 0xfd, 0x8e, 0x06, 0xb9, 0xc4, 0xba, 0x1a, 0x77, 0xb4, 0xae,
	0xa2, 0x85, 0x98, 0x2b, 0x2d, 0xe4, 0x29, 0x18, 0x99, 0x85, 0xc8, 0xbd, 0x6e, 0xdf, 0xe2, 0x21,
	0xcf, 0xc1, 0xc8, 0x3c, 0x44, 0xa2, 0xd1, 0x7f, 0x32, 0x91, 0xe6, 0x5d, 0x4c, 0x24, 0x61, 0x66,
	0x2d, 0x76, 0xbe, 0x5e, 0x4b, 0xc7, 0x3c, 0x2b, 0xa0, 0xda, 0x7f, 0x8a, 0x5f, 0x9c, 0xd4, 0x7f,
	0x64, 0x62, 0x79, 0x6a, 0x88, 0x7c, 0xee, 0x84, 0x6e, 0x6a, 0x13, 0x46, 0x71, 0xe3, 0xfd, 0x1c,
	0x42, 0x7e, 0x49, 0x8b, 0xf7, 0x4c, 0xee, 0xe9, 0x67, 0x7d, 0x2e, 0x94, 0xb0, 0xad, 0xb2, 0x2e,
	0x07, 0x79, 0x9b, 0x7a, 0xf2, 0x77, 0x69, 0xe1, 0x5b, 0x22, 0x76, 0x1d, 0x2a, 0xfd, 0xce, 0x20,
	0xe6, 0x1a, 0xa7, 0xd4, 0x2c, 0x21, 0x04, 0x46, 0xbf, 0xd3, 0x0b, 0xa2, 0x98, 0x75, 0x42, 0x42,
	0xe2, 0x88, 0x99, 0x6b, 0xb9, 0x31, 0x2e, 0x6e, 0xdb, 0xf7, 0x5c, 0xb3, 0x8c, 0x9a, 0xd0, 0xe0,
	0x34, 0x71, 0xb0, 0x9f, 0x6c, 0x12, 0x78, 0xc1, 0xc4, 0xbc, 0xc7, 0xff, 0x32, 0xea, 0xd9, 0xe0,
	0x29, 0x21, 0x21, 0x31, 0xd7, 0x53, 0xdc, 0x0f, 0x21, 0xe3, 0x12, 0x9f, 0xf2, 0x5c, 0xbe, 0xc6,
	0xae, 0xb9, 0x91, 0x06, 0xec, 0xf8, 0xbc, 0x34, 0x6c, 0x18, 0x86, 0xaf, 0x7c, 0xd7, 0xdc, 0x4c,
	0xc7, 0xe4, 0xaf, 0x68, 0x4a, 0xde, 0x3a, 0x79, 0xf6, 0xf6, 0x68, 0x0d, 0xde, 0x79, 0x7b, 0x54,
	0x46, 0xa5, 0x13, 0xfe, 0x68, 0x95, 0x60, 0x9f, 0x5f, 0x93, 0x43, 0xf1, 0x03, 0x35, 0xe5, 0xbf,
	0x46, 0xca, 0x39, 0xff, 0x28, 0x95, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x5a, 0xea, 0x8f, 0x73,
	0x3e, 0x0f, 0x00, 0x00,
}
