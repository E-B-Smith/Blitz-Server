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
func (ResponseCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

type DebugMessage struct {
	DebugText        []string `protobuf:"bytes,1,rep,name=debugText" json:"debugText,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DebugMessage) Reset()                    { *m = DebugMessage{} }
func (m *DebugMessage) String() string            { return proto.CompactTextString(m) }
func (*DebugMessage) ProtoMessage()               {}
func (*DebugMessage) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

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
func (*SessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

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
func (*BlitzHereAppOptions) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

type AppOptions struct {
	BlitzHereOptions *BlitzHereAppOptions `protobuf:"bytes,1,opt,name=blitzHereOptions" json:"blitzHereOptions,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *AppOptions) Reset()                    { *m = AppOptions{} }
func (m *AppOptions) String() string            { return proto.CompactTextString(m) }
func (*AppOptions) ProtoMessage()               {}
func (*AppOptions) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

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
func (*SessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

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

type RequestType struct {
	SessionRequest      *SessionRequest       `protobuf:"bytes,1,opt,name=sessionRequest" json:"sessionRequest,omitempty"`
	UserEventBatch      *UserEventBatch       `protobuf:"bytes,2,opt,name=userEventBatch" json:"userEventBatch,omitempty"`
	UserProfileUpdate   *UserProfileUpdate    `protobuf:"bytes,3,opt,name=userProfileUpdate" json:"userProfileUpdate,omitempty"`
	UserProfileQuery    *UserProfileQuery     `protobuf:"bytes,4,opt,name=userProfileQuery" json:"userProfileQuery,omitempty"`
	ConfirmationRequest *ConfirmationRequest  `protobuf:"bytes,5,opt,name=confirmationRequest" json:"confirmationRequest,omitempty"`
	MessageSendRequest  *UserMessageUpdate    `protobuf:"bytes,6,opt,name=messageSendRequest" json:"messageSendRequest,omitempty"`
	MessageFetchRequest *UserMessageUpdate    `protobuf:"bytes,7,opt,name=messageFetchRequest" json:"messageFetchRequest,omitempty"`
	DebugMessage        *DebugMessage         `protobuf:"bytes,8,opt,name=debugMessage" json:"debugMessage,omitempty"`
	ImageUpload         *ImageUpload          `protobuf:"bytes,9,opt,name=imageUpload" json:"imageUpload,omitempty"`
	AcceptInviteRequest *AcceptInviteRequest  `protobuf:"bytes,10,opt,name=acceptInviteRequest" json:"acceptInviteRequest,omitempty"`
	FetchFeedRequest    *FeedPostFetchRequest `protobuf:"bytes,11,opt,name=fetchFeedRequest" json:"fetchFeedRequest,omitempty"`
	AutocompleteRequest *AutocompleteRequest  `protobuf:"bytes,12,opt,name=autocompleteRequest" json:"autocompleteRequest,omitempty"`
	XXX_unrecognized    []byte                `json:"-"`
}

func (m *RequestType) Reset()                    { *m = RequestType{} }
func (m *RequestType) String() string            { return proto.CompactTextString(m) }
func (*RequestType) ProtoMessage()               {}
func (*RequestType) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

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

func (m *RequestType) GetFetchFeedRequest() *FeedPostFetchRequest {
	if m != nil {
		return m.FetchFeedRequest
	}
	return nil
}

func (m *RequestType) GetAutocompleteRequest() *AutocompleteRequest {
	if m != nil {
		return m.AutocompleteRequest
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
func (*ServerRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

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
	SessionResponse      *SessionResponse        `protobuf:"bytes,1,opt,name=sessionResponse" json:"sessionResponse,omitempty"`
	UserEventResponse    *UserEventBatchResponse `protobuf:"bytes,2,opt,name=userEventResponse" json:"userEventResponse,omitempty"`
	ProfileUpdate        *UserProfileUpdate      `protobuf:"bytes,3,opt,name=profileUpdate" json:"profileUpdate,omitempty"`
	ProfileQuery         *UserProfileQuery       `protobuf:"bytes,4,opt,name=profileQuery" json:"profileQuery,omitempty"`
	ConfirmationRequest  *ConfirmationRequest    `protobuf:"bytes,5,opt,name=confirmationRequest" json:"confirmationRequest,omitempty"`
	MessageUpdate        *UserMessageUpdate      `protobuf:"bytes,6,opt,name=messageUpdate" json:"messageUpdate,omitempty"`
	DebugMessage         *DebugMessage           `protobuf:"bytes,7,opt,name=debugMessage" json:"debugMessage,omitempty"`
	ImageUploadReply     *ImageUpload            `protobuf:"bytes,8,opt,name=imageUploadReply" json:"imageUploadReply,omitempty"`
	AcceptInviteResponse *AcceptInviteResponse   `protobuf:"bytes,9,opt,name=acceptInviteResponse" json:"acceptInviteResponse,omitempty"`
	FetchFeedResponse    *FeedPostFetchResponse  `protobuf:"bytes,10,opt,name=fetchFeedResponse" json:"fetchFeedResponse,omitempty"`
	AutocompleteResponse *AutocompleteResponse   `protobuf:"bytes,11,opt,name=autocompleteResponse" json:"autocompleteResponse,omitempty"`
	XXX_unrecognized     []byte                  `json:"-"`
}

func (m *ResponseType) Reset()                    { *m = ResponseType{} }
func (m *ResponseType) String() string            { return proto.CompactTextString(m) }
func (*ResponseType) ProtoMessage()               {}
func (*ResponseType) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{7} }

func (m *ResponseType) GetSessionResponse() *SessionResponse {
	if m != nil {
		return m.SessionResponse
	}
	return nil
}

func (m *ResponseType) GetUserEventResponse() *UserEventBatchResponse {
	if m != nil {
		return m.UserEventResponse
	}
	return nil
}

func (m *ResponseType) GetProfileUpdate() *UserProfileUpdate {
	if m != nil {
		return m.ProfileUpdate
	}
	return nil
}

func (m *ResponseType) GetProfileQuery() *UserProfileQuery {
	if m != nil {
		return m.ProfileQuery
	}
	return nil
}

func (m *ResponseType) GetConfirmationRequest() *ConfirmationRequest {
	if m != nil {
		return m.ConfirmationRequest
	}
	return nil
}

func (m *ResponseType) GetMessageUpdate() *UserMessageUpdate {
	if m != nil {
		return m.MessageUpdate
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

func (m *ResponseType) GetFetchFeedResponse() *FeedPostFetchResponse {
	if m != nil {
		return m.FetchFeedResponse
	}
	return nil
}

func (m *ResponseType) GetAutocompleteResponse() *AutocompleteResponse {
	if m != nil {
		return m.AutocompleteResponse
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
func (*ServerResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{8} }

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
	proto.RegisterType((*RequestType)(nil), "BlitzMessage.RequestType")
	proto.RegisterType((*ServerRequest)(nil), "BlitzMessage.ServerRequest")
	proto.RegisterType((*ResponseType)(nil), "BlitzMessage.ResponseType")
	proto.RegisterType((*ServerResponse)(nil), "BlitzMessage.ServerResponse")
	proto.RegisterEnum("BlitzMessage.ResponseCode", ResponseCode_name, ResponseCode_value)
}

var fileDescriptor4 = []byte{
	// 954 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x56, 0xdf, 0x72, 0xdb, 0xc4,
	0x1b, 0xfd, 0x29, 0xae, 0x93, 0xf8, 0xb3, 0xec, 0x28, 0x9b, 0xfc, 0xa8, 0x08, 0xe5, 0x4f, 0x0d,
	0x17, 0x99, 0x0e, 0xb8, 0x0c, 0x94, 0x0e, 0x03, 0x9d, 0x80, 0xe3, 0x34, 0x83, 0x67, 0x0a, 0x2d,
	0xb6, 0x33, 0x5c, 0x2b, 0xd2, 0x97, 0x44, 0x20, 0x6b, 0x97, 0xdd, 0x95, 0x87, 0xf4, 0x09, 0xb8,
	0xe1, 0x96, 0xe1, 0x21, 0x78, 0x12, 0xb8, 0xeb, 0x13, 0xb1, 0x5a, 0xad, 0xec, 0x95, 0xa3, 0x24,
	0x1d, 0x2e, 0xb8, 0xd3, 0xae, 0xce, 0xf9, 0x76, 0xbf, 0xb3, 0xe7, 0xac, 0x04, 0xee, 0x04, 0xf9,
	0x1c, 0x79, 0x9f, 0x71, 0x2a, 0x29, 0x71, 0x0f, 0x93, 0x58, 0xbe, 0xfc, 0x16, 0x85, 0x08, 0xce,
	0x71, 0xef, 0x2d, 0x7a, 0xfa, 0x23, 0x86, 0x32, 0x9e, 0x63, 0xf8, 0x51, 0x84, 0x22, 0xe4, 0x31,
	0x93, 0xd4, 0x40, 0xf7, 0xda, 0xd3, 0x4b, 0x86, 0xc2, 0x0c, 0xe0, 0x18, 0x31, 0x32, 0xcf, 0xee,
	0x11, 0xce, 0xe3, 0x10, 0xcb, 0xd1, 0x04, 0x03, 0x1e, 0x5e, 0x98, 0x51, 0xe7, 0x98, 0xc7, 0x98,
	0x46, 0x25, 0xcd, 0x3b, 0x11, 0xc8, 0x9f, 0xce, 0x31, 0x95, 0xe5, 0x0c, 0xc9, 0x67, 0xcc, 0xfa,
	0x95, 0xb9, 0x17, 0x9c, 0x9e, 0xc5, 0x49, 0x39, 0xd7, 0xbb, 0x0f, 0x6a, 0x99, 0xd3, 0xec, 0xdc,
	0x40, 0xc9, 0x36, 0xb4, 0xa2, 0x7c, 0x3c, 0xc5, 0x5f, 0xa4, 0xef, 0xbc, 0xd7, 0xd8, 0x6f, 0xf5,
	0xfe, 0x76, 0xa0, 0x3b, 0x51, 0xaf, 0x63, 0x9a, 0x8e, 0xf1, 0xe7, 0x0c, 0x85, 0x24, 0xfb, 0xb0,
	0x99, 0xd0, 0x30, 0x90, 0x6a, 0x4a, 0x81, 0x9c, 0xfd, 0xf6, 0x27, 0x6f, 0xf4, 0xed, 0x8e, 0xfb,
	0xcf, 0xcc, 0x5b, 0xf2, 0x21, 0x40, 0xa4, 0xdb, 0x18, 0xa5, 0x67, 0xd4, 0x5f, 0xd3, 0x58, 0xbf,
	0x8a, 0x3d, 0x5a, 0xbc, 0x27, 0x0f, 0x60, 0x83, 0x15, 0xfb, 0xf3, 0x1b, 0x1a, 0xfa, 0x66, 0x15,
	0x6a, 0x35, 0x40, 0x3e, 0x83, 0xdd, 0x24, 0x10, 0x72, 0xc0, 0xd8, 0x51, 0x20, 0x83, 0x31, 0x0a,
	0x94, 0xea, 0x01, 0xfd, 0x3b, 0x9a, 0x78, 0xb7, 0x4a, 0x9c, 0xc6, 0x33, 0xb5, 0xef, 0x60, 0xc6,
	0x7a, 0xff, 0x87, 0x1d, 0xfd, 0xe6, 0x1b, 0xe4, 0xa8, 0xb8, 0xcf, 0x59, 0xbe, 0x4d, 0xd1, 0x1b,
	0x01, 0x2c, 0x47, 0xe4, 0x4b, 0xf0, 0x4e, 0x4b, 0x90, 0x99, 0x33, 0x7d, 0xde, 0xaf, 0xd6, 0xad,
	0x2b, 0xf5, 0xe7, 0x1a, 0x6c, 0x2d, 0xf4, 0x12, 0x4c, 0x4d, 0x21, 0xe9, 0xc2, 0x7a, 0xa6, 0xf6,
	0x3e, 0x3a, 0xd2, 0x65, 0x5a, 0x64, 0x17, 0x5c, 0x51, 0x40, 0xa6, 0xf4, 0x27, 0x4c, 0xb5, 0x30,
	0xad, 0x5c, 0x7c, 0xa1, 0x5d, 0x74, 0x32, 0x7e, 0xa6, 0x05, 0x68, 0x91, 0x87, 0xe0, 0x66, 0xd6,
	0x49, 0xaa, 0xee, 0x1a, 0xf5, 0xb2, 0x94, 0x07, 0xd8, 0x87, 0x76, 0xb6, 0x54, 0xc9, 0x6f, 0xde,
	0x26, 0xe3, 0x5d, 0xd8, 0xe2, 0xb9, 0x76, 0x83, 0x24, 0x31, 0x52, 0xfa, 0xeb, 0x8a, 0xb3, 0x49,
	0x3e, 0x87, 0x4e, 0x9c, 0xce, 0x63, 0x89, 0xe6, 0xd0, 0xfd, 0x8d, 0x3a, 0x01, 0x06, 0x61, 0x88,
	0x4c, 0x8e, 0x6c, 0x60, 0x7e, 0xe6, 0xc1, 0x42, 0x0e, 0x7f, 0xb3, 0xee, 0xcc, 0x2d, 0xb9, 0xfe,
	0x6a, 0x42, 0xdb, 0x30, 0xf3, 0x24, 0x90, 0x47, 0xd0, 0x15, 0x15, 0xb7, 0x19, 0xe5, 0xef, 0x55,
	0x2b, 0xac, 0x38, 0x52, 0xb1, 0xb2, 0x32, 0x03, 0x87, 0x81, 0x0c, 0x2f, 0x8c, 0xd7, 0xee, 0x5d,
	0xed, 0x7c, 0x89, 0x21, 0x5f, 0xc0, 0xb6, 0x25, 0xd6, 0x09, 0x8b, 0x72, 0x03, 0x15, 0xce, 0x7b,
	0xf7, 0x5a, 0xc9, 0x0a, 0x98, 0xd2, 0xc7, 0xb3, 0xb8, 0xdf, 0x67, 0xc8, 0x2f, 0x8d, 0xf7, 0xde,
	0xb9, 0x96, 0xaa, 0x51, 0xe4, 0x00, 0x76, 0x42, 0x9a, 0x9e, 0xc5, 0x7c, 0xa6, 0x33, 0x52, 0xb6,
	0xd9, 0xac, 0xd3, 0x77, 0x78, 0x15, 0xa8, 0xdc, 0x49, 0x66, 0xc5, 0xeb, 0x89, 0xba, 0x03, 0x4a,
	0xfa, 0xfa, 0x75, 0xdb, 0x36, 0xcf, 0x66, 0xdb, 0x4f, 0x60, 0xc7, 0x90, 0x8f, 0x51, 0x49, 0x50,
	0x3d, 0xdc, 0x5b, 0xd9, 0x1f, 0x83, 0x1b, 0x59, 0xd7, 0x85, 0x39, 0xdc, 0xbd, 0xd5, 0x40, 0x5b,
	0x17, 0x8a, 0xf2, 0x63, 0x3c, 0xd3, 0x05, 0x12, 0x1a, 0x44, 0x7e, 0xab, 0xce, 0x8f, 0xa3, 0x25,
	0x20, 0x17, 0x27, 0xb8, 0xea, 0x29, 0x1f, 0x5e, 0xd7, 0x7c, 0x4f, 0xc0, 0x3b, 0xcb, 0x1b, 0xcb,
	0x2f, 0xd2, 0x92, 0xdc, 0xd6, 0xe4, 0x5e, 0x95, 0x9c, 0x03, 0x5e, 0x50, 0x21, 0x6d, 0x19, 0xf4,
	0xea, 0x99, 0xa4, 0x21, 0x9d, 0xb1, 0x04, 0x97, 0xab, 0xbb, 0xb5, 0xab, 0x5f, 0x05, 0xf6, 0x4e,
	0xa0, 0x53, 0x7c, 0x07, 0xca, 0x82, 0xab, 0x41, 0x2f, 0xe2, 0xaf, 0x44, 0xe1, 0x4b, 0xcb, 0x1b,
	0xab, 0xae, 0x88, 0x62, 0x65, 0xa2, 0xf7, 0x47, 0x13, 0xdc, 0xf2, 0x2e, 0xd1, 0x21, 0x79, 0x0c,
	0x5b, 0xa2, 0x7a, 0xc5, 0x98, 0x94, 0xbc, 0x7d, 0x4d, 0x4a, 0xcc, 0x3d, 0xf4, 0x55, 0x61, 0x78,
	0x1d, 0x81, 0x05, 0xb3, 0x58, 0xfe, 0x83, 0x9b, 0x92, 0xb2, 0x28, 0xf0, 0x18, 0x3a, 0xec, 0xdf,
	0xa4, 0xe5, 0x11, 0xb8, 0xec, 0xbf, 0x4f, 0x8a, 0xda, 0xed, 0xcc, 0xf6, 0xef, 0xeb, 0x86, 0x64,
	0xd5, 0xe6, 0x1b, 0xb7, 0xda, 0xfc, 0x53, 0xf0, 0x2c, 0x9b, 0x8f, 0x91, 0x25, 0x97, 0x26, 0x1c,
	0x37, 0x78, 0xfd, 0x6b, 0xd8, 0xad, 0x7a, 0xdd, 0x1c, 0x48, 0xab, 0xce, 0xaf, 0x83, 0x1a, 0xa4,
	0x12, 0x68, 0xdb, 0x72, 0xbb, 0xa1, 0x17, 0x59, 0x79, 0xff, 0x46, 0xbb, 0x1b, 0x7e, 0xbe, 0x83,
	0x8a, 0x8d, 0x4d, 0x89, 0xda, 0xc4, 0x0c, 0x6a, 0x90, 0xbd, 0xdf, 0xf4, 0xdf, 0x41, 0x61, 0x79,
	0x53, 0x54, 0xa9, 0xc7, 0xcd, 0xf3, 0x90, 0x46, 0x85, 0x33, 0xbb, 0xab, 0xea, 0x8d, 0x2d, 0x84,
	0xf9, 0x08, 0xe9, 0x71, 0x29, 0x79, 0xf1, 0x45, 0xb4, 0x4a, 0xe9, 0xa4, 0x34, 0xea, 0x0e, 0xc2,
	0x4e, 0xc6, 0x83, 0xdf, 0x9d, 0x65, 0x54, 0x74, 0xed, 0x0e, 0xb4, 0xc6, 0xc3, 0x49, 0xa6, 0xd4,
	0x13, 0xc2, 0x73, 0x08, 0x81, 0xee, 0x78, 0x38, 0x4a, 0x59, 0x26, 0x87, 0x94, 0xf3, 0x8c, 0x49,
	0x6f, 0xcd, 0x9a, 0x53, 0xf2, 0x06, 0x49, 0x1c, 0x79, 0x0d, 0xb2, 0x03, 0x5b, 0x8a, 0xa6, 0x1b,
	0xfb, 0x21, 0xe0, 0x69, 0x9c, 0x9e, 0x7b, 0x77, 0xd4, 0x07, 0xba, 0x53, 0x4e, 0x3e, 0xe5, 0x9c,
	0x72, 0xaf, 0x59, 0xe0, 0xbe, 0xa3, 0x52, 0x89, 0x73, 0x41, 0x79, 0xfc, 0x12, 0x23, 0x6f, 0xbd,
	0x28, 0x38, 0x4c, 0xd4, 0x2f, 0x9a, 0x9c, 0x52, 0xfa, 0x3c, 0x89, 0xbc, 0x8d, 0xc3, 0x87, 0xaf,
	0x0e, 0xd6, 0xe0, 0x7f, 0xaf, 0x0e, 0x1a, 0xc4, 0x39, 0x54, 0x8f, 0xbe, 0x03, 0x7b, 0x4a, 0xcc,
	0xbe, 0xfe, 0xcf, 0xb8, 0x50, 0x7f, 0x10, 0x95, 0x9e, 0x7e, 0x75, 0x9c, 0x7f, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xfe, 0xb2, 0x56, 0x58, 0x52, 0x0a, 0x00, 0x00,
}
