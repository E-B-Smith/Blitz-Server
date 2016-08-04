// Code generated by protoc-gen-go.
// source: Device.proto
// DO NOT EDIT!

/*
Package BlitzMessage is a generated protocol buffer package.

It is generated from these files:
	Device.proto
	EntityTags.proto
	Feed.proto
	Globals.proto
	Payments.proto
	Search.proto
	Server.proto
	Types.proto
	UserEvents.proto
	UserMessages.proto
	UserProfiles.proto

It has these top-level messages:
	DeviceInfo
	EntityTag
	EntityTagList
	FeedPost
	FeedPostUpdateRequest
	FeedPostFetchRequest
	FeedPostResponse
	Global
	CardInfo
	UserCardInfo
	Charge
	PurchaseDescription
	FetchPurchaseDescription
	AutocompleteRequest
	AutocompleteResponse
	UserSearchRequest
	UserSearchResponse
	SearchCategory
	SearchCategories
	DebugMessage
	SessionRequest
	BlitzHereAppOptions
	AppOptions
	SessionResponse
	LoginAsAdmin
	PushConnect
	PushDisconnect
	RequestType
	ServerRequest
	ResponseType
	ServerResponse
	Timestamp
	Timespan
	Point
	Size
	Coordinate
	CoordinateRegion
	CoordinatePolygon
	Location
	Void
	KeyValue
	UserEvent
	UserEventBatch
	UserEventBatchResponse
	Conversation
	ConversationGroup
	ConversationRequest
	ConversationResponse
	FetchConversations
	FetchConversationGroups
	UpdateConversationStatus
	UserMessage
	UserMessageUpdate
	SocialIdentity
	ContactInfo
	Employment
	Education
	ImageData
	UserReview
	UserProfile
	ImageUpload
	UserProfileUpdate
	UserProfileQuery
	ConfirmationRequest
	ProfilesFromContactInfo
	FriendUpdate
	UserInvite
	UserInvites
	EditProfile
*/
package BlitzMessage

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"


// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PlatformType int32

const (
	PlatformType_PTUnknown PlatformType = 0
	PlatformType_PTiOS     PlatformType = 1
	PlatformType_PTAndroid PlatformType = 2
	PlatformType_PTWeb     PlatformType = 3
)

var PlatformType_name = map[int32]string{
	0: "PTUnknown",
	1: "PTiOS",
	2: "PTAndroid",
	3: "PTWeb",
}
var PlatformType_value = map[string]int32{
	"PTUnknown": 0,
	"PTiOS":     1,
	"PTAndroid": 2,
	"PTWeb":     3,
}

func (x PlatformType) Enum() *PlatformType {
	p := new(PlatformType)
	*p = x
	return p
}
func (x PlatformType) String() string {
	return proto.EnumName(PlatformType_name, int32(x))
}
func (x *PlatformType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PlatformType_value, data, "PlatformType")
	if err != nil {
		return err
	}
	*x = PlatformType(value)
	return nil
}
func (PlatformType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type DeviceInfo struct {
	VendorUID           *string       `protobuf:"bytes,1,opt,name=vendorUID" json:"vendorUID,omitempty"`
	AdvertisingUID      *string       `protobuf:"bytes,2,opt,name=advertisingUID" json:"advertisingUID,omitempty"`
	PlatformType        *PlatformType `protobuf:"varint,3,opt,name=platformType,enum=BlitzMessage.PlatformType" json:"platformType,omitempty"`
	ModelName           *string       `protobuf:"bytes,4,opt,name=modelName" json:"modelName,omitempty"`
	SystemVersion       *string       `protobuf:"bytes,5,opt,name=systemVersion" json:"systemVersion,omitempty"`
	Language            *string       `protobuf:"bytes,6,opt,name=language" json:"language,omitempty"`
	Timezone            *string       `protobuf:"bytes,7,opt,name=timezone" json:"timezone,omitempty"`
	PhoneCountryCode    *string       `protobuf:"bytes,8,opt,name=phoneCountryCode" json:"phoneCountryCode,omitempty"`
	ScreenSize          *Size         `protobuf:"bytes,9,opt,name=screenSize" json:"screenSize,omitempty"`
	ScreenScale         *float32      `protobuf:"fixed32,10,opt,name=screenScale,def=1" json:"screenScale,omitempty"`
	AppID               *string       `protobuf:"bytes,11,opt,name=appID" json:"appID,omitempty"`
	AppVersion          *string       `protobuf:"bytes,12,opt,name=appVersion" json:"appVersion,omitempty"`
	NotificationToken   *string       `protobuf:"bytes,13,opt,name=notificationToken" json:"notificationToken,omitempty"`
	AppIsReleaseVersion *bool         `protobuf:"varint,14,opt,name=appIsReleaseVersion" json:"appIsReleaseVersion,omitempty"`
	UserTags            []string      `protobuf:"bytes,15,rep,name=userTags" json:"userTags,omitempty"`
	DeviceUDID          *string       `protobuf:"bytes,16,opt,name=deviceUDID" json:"deviceUDID,omitempty"`
	ColorDepth          *float32      `protobuf:"fixed32,17,opt,name=colorDepth" json:"colorDepth,omitempty"`
	IPAddress           *string       `protobuf:"bytes,18,opt,name=IPAddress" json:"IPAddress,omitempty"`
	SystemBuildVersion  *string       `protobuf:"bytes,19,opt,name=systemBuildVersion" json:"systemBuildVersion,omitempty"`
	LocalIPAddress      *string       `protobuf:"bytes,20,opt,name=localIPAddress" json:"localIPAddress,omitempty"`
	XXX_unrecognized    []byte        `json:"-"`
}

func (m *DeviceInfo) Reset()                    { *m = DeviceInfo{} }
func (m *DeviceInfo) String() string            { return proto.CompactTextString(m) }
func (*DeviceInfo) ProtoMessage()               {}
func (*DeviceInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

const Default_DeviceInfo_ScreenScale float32 = 1

func (m *DeviceInfo) GetVendorUID() string {
	if m != nil && m.VendorUID != nil {
		return *m.VendorUID
	}
	return ""
}

func (m *DeviceInfo) GetAdvertisingUID() string {
	if m != nil && m.AdvertisingUID != nil {
		return *m.AdvertisingUID
	}
	return ""
}

func (m *DeviceInfo) GetPlatformType() PlatformType {
	if m != nil && m.PlatformType != nil {
		return *m.PlatformType
	}
	return PlatformType_PTUnknown
}

func (m *DeviceInfo) GetModelName() string {
	if m != nil && m.ModelName != nil {
		return *m.ModelName
	}
	return ""
}

func (m *DeviceInfo) GetSystemVersion() string {
	if m != nil && m.SystemVersion != nil {
		return *m.SystemVersion
	}
	return ""
}

func (m *DeviceInfo) GetLanguage() string {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return ""
}

func (m *DeviceInfo) GetTimezone() string {
	if m != nil && m.Timezone != nil {
		return *m.Timezone
	}
	return ""
}

func (m *DeviceInfo) GetPhoneCountryCode() string {
	if m != nil && m.PhoneCountryCode != nil {
		return *m.PhoneCountryCode
	}
	return ""
}

func (m *DeviceInfo) GetScreenSize() *Size {
	if m != nil {
		return m.ScreenSize
	}
	return nil
}

func (m *DeviceInfo) GetScreenScale() float32 {
	if m != nil && m.ScreenScale != nil {
		return *m.ScreenScale
	}
	return Default_DeviceInfo_ScreenScale
}

func (m *DeviceInfo) GetAppID() string {
	if m != nil && m.AppID != nil {
		return *m.AppID
	}
	return ""
}

func (m *DeviceInfo) GetAppVersion() string {
	if m != nil && m.AppVersion != nil {
		return *m.AppVersion
	}
	return ""
}

func (m *DeviceInfo) GetNotificationToken() string {
	if m != nil && m.NotificationToken != nil {
		return *m.NotificationToken
	}
	return ""
}

func (m *DeviceInfo) GetAppIsReleaseVersion() bool {
	if m != nil && m.AppIsReleaseVersion != nil {
		return *m.AppIsReleaseVersion
	}
	return false
}

func (m *DeviceInfo) GetUserTags() []string {
	if m != nil {
		return m.UserTags
	}
	return nil
}

func (m *DeviceInfo) GetDeviceUDID() string {
	if m != nil && m.DeviceUDID != nil {
		return *m.DeviceUDID
	}
	return ""
}

func (m *DeviceInfo) GetColorDepth() float32 {
	if m != nil && m.ColorDepth != nil {
		return *m.ColorDepth
	}
	return 0
}

func (m *DeviceInfo) GetIPAddress() string {
	if m != nil && m.IPAddress != nil {
		return *m.IPAddress
	}
	return ""
}

func (m *DeviceInfo) GetSystemBuildVersion() string {
	if m != nil && m.SystemBuildVersion != nil {
		return *m.SystemBuildVersion
	}
	return ""
}

func (m *DeviceInfo) GetLocalIPAddress() string {
	if m != nil && m.LocalIPAddress != nil {
		return *m.LocalIPAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*DeviceInfo)(nil), "BlitzMessage.DeviceInfo")
	proto.RegisterEnum("BlitzMessage.PlatformType", PlatformType_name, PlatformType_value)
}

func init() { proto.RegisterFile("Device.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x51, 0xc1, 0x72, 0xd3, 0x30,
	0x10, 0xad, 0x13, 0x0a, 0xc9, 0xc6, 0x09, 0x8e, 0x0a, 0x8c, 0x70, 0x2f, 0x1d, 0x0e, 0x0c, 0xc3,
	0x0c, 0x01, 0x7a, 0xe4, 0xd0, 0x99, 0xba, 0xb9, 0xe4, 0x00, 0x64, 0xa8, 0x03, 0x67, 0xc5, 0xda,
	0x24, 0xa2, 0xb6, 0xe4, 0x91, 0x94, 0x30, 0xc9, 0x17, 0xf0, 0x7d, 0xfd, 0x0f, 0xfe, 0x01, 0x49,
	0xc6, 0x43, 0x73, 0xd3, 0xbe, 0xb7, 0xfb, 0x76, 0xdf, 0x13, 0xc4, 0x53, 0xdc, 0x89, 0x02, 0x27,
	0xb5, 0x56, 0x56, 0x91, 0x38, 0x2b, 0x85, 0x3d, 0x7c, 0x46, 0x63, 0xd8, 0x1a, 0xd3, 0x73, 0xb5,
	0xfc, 0x89, 0x85, 0x15, 0x3b, 0x2c, 0xde, 0x71, 0x34, 0x85, 0x16, 0xb5, 0x55, 0xba, 0x69, 0x4d,
	0x07, 0xf9, 0xbe, 0x46, 0xd3, 0x14, 0xaf, 0xfe, 0x74, 0x01, 0x1a, 0xa1, 0x99, 0x5c, 0x29, 0x32,
	0x86, 0xfe, 0x0e, 0x25, 0x57, 0x7a, 0x31, 0x9b, 0xd2, 0xe8, 0x22, 0x7a, 0xd3, 0x27, 0x2f, 0x60,
	0xc4, 0xf8, 0x0e, 0xb5, 0x15, 0x46, 0xc8, 0xb5, 0xc7, 0x3b, 0x01, 0xff, 0x00, 0x71, 0x5d, 0x32,
	0xbb, 0x52, 0xba, 0xf2, 0x82, 0xb4, 0xeb, 0xd0, 0xd1, 0x65, 0x3a, 0x79, 0x78, 0xc8, 0x64, 0xfe,
	0xa0, 0xc3, 0x8b, 0x57, 0x8a, 0x63, 0xf9, 0x85, 0x55, 0x48, 0x1f, 0x05, 0x91, 0xe7, 0x30, 0x34,
	0x7b, 0x63, 0xb1, 0xfa, 0x8e, 0xda, 0x08, 0x25, 0xe9, 0x69, 0x80, 0x13, 0xe8, 0x95, 0x4c, 0xae,
	0xb7, 0x4e, 0x82, 0x3e, 0x6e, 0x11, 0x2b, 0x2a, 0x3c, 0x28, 0x89, 0xf4, 0x49, 0x40, 0x28, 0x24,
	0xf5, 0xc6, 0x95, 0x37, 0x6a, 0x2b, 0xad, 0xde, 0xdf, 0x38, 0x65, 0xda, 0x0b, 0xcc, 0x6b, 0x00,
	0x67, 0x19, 0x51, 0xde, 0x8a, 0x03, 0xd2, 0xbe, 0xc3, 0x06, 0x97, 0xe4, 0xf8, 0x2e, 0xcf, 0x38,
	0x67, 0x83, 0x7f, 0x7d, 0x05, 0x2b, 0x91, 0x82, 0x6b, 0xec, 0x7c, 0x8a, 0x3e, 0x92, 0x21, 0x9c,
	0xb2, 0xba, 0x76, 0x46, 0x07, 0x41, 0x8e, 0x00, 0xb8, 0xb2, 0x3d, 0x30, 0x0e, 0xd8, 0x4b, 0x18,
	0x4b, 0x65, 0xc5, 0x4a, 0x14, 0xcc, 0x3a, 0x34, 0x57, 0x77, 0x28, 0xe9, 0x30, 0x50, 0xe7, 0x70,
	0xe6, 0xa7, 0xcd, 0x37, 0x2c, 0x91, 0x19, 0x6c, 0xe7, 0x46, 0x8e, 0xec, 0x79, 0x1b, 0x5b, 0x83,
	0x3a, 0x67, 0x6b, 0x43, 0x9f, 0x5e, 0x74, 0x1b, 0x75, 0x1e, 0xf2, 0x5f, 0x4c, 0xdd, 0xc6, 0xa4,
	0xdd, 0x58, 0xa8, 0x52, 0xe9, 0x29, 0xd6, 0x76, 0x43, 0xc7, 0xfe, 0x2e, 0x1f, 0xde, 0x6c, 0x7e,
	0xcd, 0xb9, 0x76, 0x0e, 0x28, 0x09, 0x6d, 0x29, 0x90, 0x26, 0xbc, 0x6c, 0x2b, 0x4a, 0xde, 0x2e,
	0x3a, 0x6b, 0x7f, 0xad, 0x54, 0xce, 0xd6, 0xff, 0x99, 0x67, 0x1e, 0x7f, 0x9b, 0x41, 0x7c, 0xf4,
	0x27, 0x43, 0xe8, 0xcf, 0xf3, 0x85, 0xbc, 0x93, 0xea, 0x97, 0x4c, 0x4e, 0x48, 0x1f, 0x4e, 0xe7,
	0xb9, 0xf8, 0x7a, 0x9b, 0x44, 0x0d, 0x73, 0x2d, 0xb9, 0x56, 0x82, 0x27, 0x9d, 0x86, 0xf9, 0x81,
	0xcb, 0xa4, 0x9b, 0xbd, 0xbf, 0xbf, 0xea, 0xc0, 0xc9, 0xfd, 0x55, 0x97, 0x44, 0x99, 0x7b, 0xd2,
	0x08, 0xd2, 0x42, 0x55, 0x93, 0xa5, 0x0f, 0x78, 0x83, 0x1a, 0x8f, 0xa2, 0xfe, 0x1d, 0x45, 0x7f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x41, 0xda, 0x9b, 0xec, 0xab, 0x02, 0x00, 0x00,
}
