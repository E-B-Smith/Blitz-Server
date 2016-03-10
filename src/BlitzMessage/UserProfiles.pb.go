// Code generated by protoc-gen-go.
// source: UserProfiles.proto
// DO NOT EDIT!

package BlitzMessage

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"


// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ContactType int32

const (
	ContactType_CTUnknown       ContactType = 0
	ContactType_CTPhoneSMS      ContactType = 1
	ContactType_CTEmail         ContactType = 2
	ContactType_CTChat          ContactType = 3
	ContactType_CTSocialService ContactType = 4
)

var ContactType_name = map[int32]string{
	0: "CTUnknown",
	1: "CTPhoneSMS",
	2: "CTEmail",
	3: "CTChat",
	4: "CTSocialService",
}
var ContactType_value = map[string]int32{
	"CTUnknown":       0,
	"CTPhoneSMS":      1,
	"CTEmail":         2,
	"CTChat":          3,
	"CTSocialService": 4,
}

func (x ContactType) Enum() *ContactType {
	p := new(ContactType)
	*p = x
	return p
}
func (x ContactType) String() string {
	return proto.EnumName(ContactType_name, int32(x))
}
func (x *ContactType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ContactType_value, data, "ContactType")
	if err != nil {
		return err
	}
	*x = ContactType(value)
	return nil
}
func (ContactType) EnumDescriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

type UserStatus int32

const (
	UserStatus_USUnknown    UserStatus = 0
	UserStatus_USBlocked    UserStatus = 1
	UserStatus_USInvited    UserStatus = 2
	UserStatus_USActive     UserStatus = 3
	UserStatus_USConfirming UserStatus = 4
	UserStatus_USConfirmed  UserStatus = 5
)

var UserStatus_name = map[int32]string{
	0: "USUnknown",
	1: "USBlocked",
	2: "USInvited",
	3: "USActive",
	4: "USConfirming",
	5: "USConfirmed",
}
var UserStatus_value = map[string]int32{
	"USUnknown":    0,
	"USBlocked":    1,
	"USInvited":    2,
	"USActive":     3,
	"USConfirming": 4,
	"USConfirmed":  5,
}

func (x UserStatus) Enum() *UserStatus {
	p := new(UserStatus)
	*p = x
	return p
}
func (x UserStatus) String() string {
	return proto.EnumName(UserStatus_name, int32(x))
}
func (x *UserStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(UserStatus_value, data, "UserStatus")
	if err != nil {
		return err
	}
	*x = UserStatus(value)
	return nil
}
func (UserStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

type Gender int32

const (
	Gender_GUnknown Gender = 0
	Gender_GFemale  Gender = 1
	Gender_GMale    Gender = 2
	Gender_GOther   Gender = 3
)

var Gender_name = map[int32]string{
	0: "GUnknown",
	1: "GFemale",
	2: "GMale",
	3: "GOther",
}
var Gender_value = map[string]int32{
	"GUnknown": 0,
	"GFemale":  1,
	"GMale":    2,
	"GOther":   3,
}

func (x Gender) Enum() *Gender {
	p := new(Gender)
	*p = x
	return p
}
func (x Gender) String() string {
	return proto.EnumName(Gender_name, int32(x))
}
func (x *Gender) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Gender_value, data, "Gender")
	if err != nil {
		return err
	}
	*x = Gender(value)
	return nil
}
func (Gender) EnumDescriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

type ImageContent int32

const (
	ImageContent_ICUnknown        ImageContent = 0
	ImageContent_ICUserProfile    ImageContent = 1
	ImageContent_ICUserBackground ImageContent = 2
)

var ImageContent_name = map[int32]string{
	0: "ICUnknown",
	1: "ICUserProfile",
	2: "ICUserBackground",
}
var ImageContent_value = map[string]int32{
	"ICUnknown":        0,
	"ICUserProfile":    1,
	"ICUserBackground": 2,
}

func (x ImageContent) Enum() *ImageContent {
	p := new(ImageContent)
	*p = x
	return p
}
func (x ImageContent) String() string {
	return proto.EnumName(ImageContent_name, int32(x))
}
func (x *ImageContent) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ImageContent_value, data, "ImageContent")
	if err != nil {
		return err
	}
	*x = ImageContent(value)
	return nil
}
func (ImageContent) EnumDescriptor() ([]byte, []int) { return fileDescriptor8, []int{3} }

type SocialIdentity struct {
	SocialService    *string    `protobuf:"bytes,1,req,name=socialService" json:"socialService,omitempty"`
	SocialID         *string    `protobuf:"bytes,2,opt,name=socialID" json:"socialID,omitempty"`
	UserName         *string    `protobuf:"bytes,3,opt,name=userName" json:"userName,omitempty"`
	DisplayName      *string    `protobuf:"bytes,4,opt,name=displayName" json:"displayName,omitempty"`
	UserURI          *string    `protobuf:"bytes,5,opt,name=userURI" json:"userURI,omitempty"`
	AuthToken        *string    `protobuf:"bytes,6,opt,name=authToken" json:"authToken,omitempty"`
	AuthExpire       *Timestamp `protobuf:"bytes,7,opt,name=authExpire" json:"authExpire,omitempty"`
	AuthSecret       *string    `protobuf:"bytes,8,opt,name=authSecret" json:"authSecret,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *SocialIdentity) Reset()                    { *m = SocialIdentity{} }
func (m *SocialIdentity) String() string            { return proto.CompactTextString(m) }
func (*SocialIdentity) ProtoMessage()               {}
func (*SocialIdentity) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *SocialIdentity) GetSocialService() string {
	if m != nil && m.SocialService != nil {
		return *m.SocialService
	}
	return ""
}

func (m *SocialIdentity) GetSocialID() string {
	if m != nil && m.SocialID != nil {
		return *m.SocialID
	}
	return ""
}

func (m *SocialIdentity) GetUserName() string {
	if m != nil && m.UserName != nil {
		return *m.UserName
	}
	return ""
}

func (m *SocialIdentity) GetDisplayName() string {
	if m != nil && m.DisplayName != nil {
		return *m.DisplayName
	}
	return ""
}

func (m *SocialIdentity) GetUserURI() string {
	if m != nil && m.UserURI != nil {
		return *m.UserURI
	}
	return ""
}

func (m *SocialIdentity) GetAuthToken() string {
	if m != nil && m.AuthToken != nil {
		return *m.AuthToken
	}
	return ""
}

func (m *SocialIdentity) GetAuthExpire() *Timestamp {
	if m != nil {
		return m.AuthExpire
	}
	return nil
}

func (m *SocialIdentity) GetAuthSecret() string {
	if m != nil && m.AuthSecret != nil {
		return *m.AuthSecret
	}
	return ""
}

type ContactInfo struct {
	ContactType      *ContactType `protobuf:"varint,1,req,name=contactType,enum=BlitzMessage.ContactType" json:"contactType,omitempty"`
	Contact          *string      `protobuf:"bytes,2,req,name=contact" json:"contact,omitempty"`
	IsVerified       *bool        `protobuf:"varint,3,opt,name=isVerified" json:"isVerified,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ContactInfo) Reset()                    { *m = ContactInfo{} }
func (m *ContactInfo) String() string            { return proto.CompactTextString(m) }
func (*ContactInfo) ProtoMessage()               {}
func (*ContactInfo) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

func (m *ContactInfo) GetContactType() ContactType {
	if m != nil && m.ContactType != nil {
		return *m.ContactType
	}
	return ContactType_CTUnknown
}

func (m *ContactInfo) GetContact() string {
	if m != nil && m.Contact != nil {
		return *m.Contact
	}
	return ""
}

func (m *ContactInfo) GetIsVerified() bool {
	if m != nil && m.IsVerified != nil {
		return *m.IsVerified
	}
	return false
}

type Employment struct {
	JobTitle         *string   `protobuf:"bytes,1,opt,name=jobTitle" json:"jobTitle,omitempty"`
	CompanyName      *string   `protobuf:"bytes,2,opt,name=companyName" json:"companyName,omitempty"`
	Location         *string   `protobuf:"bytes,3,opt,name=location" json:"location,omitempty"`
	Industry         *string   `protobuf:"bytes,4,opt,name=industry" json:"industry,omitempty"`
	Timespan         *Timespan `protobuf:"bytes,5,opt,name=timespan" json:"timespan,omitempty"`
	Summary          *string   `protobuf:"bytes,6,opt,name=summary" json:"summary,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Employment) Reset()                    { *m = Employment{} }
func (m *Employment) String() string            { return proto.CompactTextString(m) }
func (*Employment) ProtoMessage()               {}
func (*Employment) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

func (m *Employment) GetJobTitle() string {
	if m != nil && m.JobTitle != nil {
		return *m.JobTitle
	}
	return ""
}

func (m *Employment) GetCompanyName() string {
	if m != nil && m.CompanyName != nil {
		return *m.CompanyName
	}
	return ""
}

func (m *Employment) GetLocation() string {
	if m != nil && m.Location != nil {
		return *m.Location
	}
	return ""
}

func (m *Employment) GetIndustry() string {
	if m != nil && m.Industry != nil {
		return *m.Industry
	}
	return ""
}

func (m *Employment) GetTimespan() *Timespan {
	if m != nil {
		return m.Timespan
	}
	return nil
}

func (m *Employment) GetSummary() string {
	if m != nil && m.Summary != nil {
		return *m.Summary
	}
	return ""
}

type Education struct {
	SchoolName       *string   `protobuf:"bytes,1,opt,name=schoolName" json:"schoolName,omitempty"`
	Degree           *string   `protobuf:"bytes,2,opt,name=degree" json:"degree,omitempty"`
	Emphasis         *string   `protobuf:"bytes,3,opt,name=emphasis" json:"emphasis,omitempty"`
	Timespan         *Timespan `protobuf:"bytes,4,opt,name=timespan" json:"timespan,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Education) Reset()                    { *m = Education{} }
func (m *Education) String() string            { return proto.CompactTextString(m) }
func (*Education) ProtoMessage()               {}
func (*Education) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{3} }

func (m *Education) GetSchoolName() string {
	if m != nil && m.SchoolName != nil {
		return *m.SchoolName
	}
	return ""
}

func (m *Education) GetDegree() string {
	if m != nil && m.Degree != nil {
		return *m.Degree
	}
	return ""
}

func (m *Education) GetEmphasis() string {
	if m != nil && m.Emphasis != nil {
		return *m.Emphasis
	}
	return ""
}

func (m *Education) GetTimespan() *Timespan {
	if m != nil {
		return m.Timespan
	}
	return nil
}

type ImageData struct {
	ImageContent     *ImageContent `protobuf:"varint,1,opt,name=imageContent,enum=BlitzMessage.ImageContent" json:"imageContent,omitempty"`
	ImageBytes       []byte        `protobuf:"bytes,2,opt,name=imageBytes" json:"imageBytes,omitempty"`
	ContentType      *string       `protobuf:"bytes,3,opt,name=contentType" json:"contentType,omitempty"`
	ImageURL         *string       `protobuf:"bytes,4,opt,name=imageURL" json:"imageURL,omitempty"`
	DateAdded        *Timestamp    `protobuf:"bytes,5,opt,name=dateAdded" json:"dateAdded,omitempty"`
	Crc32            *uint32       `protobuf:"varint,6,opt,name=crc32" json:"crc32,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *ImageData) Reset()                    { *m = ImageData{} }
func (m *ImageData) String() string            { return proto.CompactTextString(m) }
func (*ImageData) ProtoMessage()               {}
func (*ImageData) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{4} }

func (m *ImageData) GetImageContent() ImageContent {
	if m != nil && m.ImageContent != nil {
		return *m.ImageContent
	}
	return ImageContent_ICUnknown
}

func (m *ImageData) GetImageBytes() []byte {
	if m != nil {
		return m.ImageBytes
	}
	return nil
}

func (m *ImageData) GetContentType() string {
	if m != nil && m.ContentType != nil {
		return *m.ContentType
	}
	return ""
}

func (m *ImageData) GetImageURL() string {
	if m != nil && m.ImageURL != nil {
		return *m.ImageURL
	}
	return ""
}

func (m *ImageData) GetDateAdded() *Timestamp {
	if m != nil {
		return m.DateAdded
	}
	return nil
}

func (m *ImageData) GetCrc32() uint32 {
	if m != nil && m.Crc32 != nil {
		return *m.Crc32
	}
	return 0
}

type UserProfile struct {
	UserID            *string           `protobuf:"bytes,1,opt,name=userID" json:"userID,omitempty"`
	UserStatus        *UserStatus       `protobuf:"varint,2,opt,name=userStatus,enum=BlitzMessage.UserStatus" json:"userStatus,omitempty"`
	CreationDate      *Timestamp        `protobuf:"bytes,3,opt,name=creationDate" json:"creationDate,omitempty"`
	LastSeen          *Timestamp        `protobuf:"bytes,4,opt,name=lastSeen" json:"lastSeen,omitempty"`
	Name              *string           `protobuf:"bytes,5,opt,name=name" json:"name,omitempty"`
	Gender            *Gender           `protobuf:"varint,6,opt,name=gender,enum=BlitzMessage.Gender" json:"gender,omitempty"`
	Birthday          *Timestamp        `protobuf:"bytes,7,opt,name=birthday" json:"birthday,omitempty"`
	Images            []*ImageData      `protobuf:"bytes,8,rep,name=images" json:"images,omitempty"`
	SocialIdentities  []*SocialIdentity `protobuf:"bytes,9,rep,name=socialIdentities" json:"socialIdentities,omitempty"`
	ContactInfo       []*ContactInfo    `protobuf:"bytes,10,rep,name=contactInfo" json:"contactInfo,omitempty"`
	CurrentEmployment *Employment       `protobuf:"bytes,11,opt,name=currentEmployment" json:"currentEmployment,omitempty"`
	Employment        []*Employment     `protobuf:"bytes,12,rep,name=employment" json:"employment,omitempty"`
	Education         []*Education      `protobuf:"bytes,13,rep,name=education" json:"education,omitempty"`
	ExpertiseTags     []string          `protobuf:"bytes,14,rep,name=expertiseTags" json:"expertiseTags,omitempty"`
	InterestTags      []string          `protobuf:"bytes,15,rep,name=interestTags" json:"interestTags,omitempty"`
	BackgroundSummary *string           `protobuf:"bytes,16,opt,name=backgroundSummary" json:"backgroundSummary,omitempty"`
	XXX_unrecognized  []byte            `json:"-"`
}

func (m *UserProfile) Reset()                    { *m = UserProfile{} }
func (m *UserProfile) String() string            { return proto.CompactTextString(m) }
func (*UserProfile) ProtoMessage()               {}
func (*UserProfile) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{5} }

func (m *UserProfile) GetUserID() string {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return ""
}

func (m *UserProfile) GetUserStatus() UserStatus {
	if m != nil && m.UserStatus != nil {
		return *m.UserStatus
	}
	return UserStatus_USUnknown
}

func (m *UserProfile) GetCreationDate() *Timestamp {
	if m != nil {
		return m.CreationDate
	}
	return nil
}

func (m *UserProfile) GetLastSeen() *Timestamp {
	if m != nil {
		return m.LastSeen
	}
	return nil
}

func (m *UserProfile) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *UserProfile) GetGender() Gender {
	if m != nil && m.Gender != nil {
		return *m.Gender
	}
	return Gender_GUnknown
}

func (m *UserProfile) GetBirthday() *Timestamp {
	if m != nil {
		return m.Birthday
	}
	return nil
}

func (m *UserProfile) GetImages() []*ImageData {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *UserProfile) GetSocialIdentities() []*SocialIdentity {
	if m != nil {
		return m.SocialIdentities
	}
	return nil
}

func (m *UserProfile) GetContactInfo() []*ContactInfo {
	if m != nil {
		return m.ContactInfo
	}
	return nil
}

func (m *UserProfile) GetCurrentEmployment() *Employment {
	if m != nil {
		return m.CurrentEmployment
	}
	return nil
}

func (m *UserProfile) GetEmployment() []*Employment {
	if m != nil {
		return m.Employment
	}
	return nil
}

func (m *UserProfile) GetEducation() []*Education {
	if m != nil {
		return m.Education
	}
	return nil
}

func (m *UserProfile) GetExpertiseTags() []string {
	if m != nil {
		return m.ExpertiseTags
	}
	return nil
}

func (m *UserProfile) GetInterestTags() []string {
	if m != nil {
		return m.InterestTags
	}
	return nil
}

func (m *UserProfile) GetBackgroundSummary() string {
	if m != nil && m.BackgroundSummary != nil {
		return *m.BackgroundSummary
	}
	return ""
}

type ImageUpload struct {
	ImageData        []*ImageData `protobuf:"bytes,1,rep,name=imageData" json:"imageData,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ImageUpload) Reset()                    { *m = ImageUpload{} }
func (m *ImageUpload) String() string            { return proto.CompactTextString(m) }
func (*ImageUpload) ProtoMessage()               {}
func (*ImageUpload) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{6} }

func (m *ImageUpload) GetImageData() []*ImageData {
	if m != nil {
		return m.ImageData
	}
	return nil
}

type UserProfileUpdate struct {
	Profiles         []*UserProfile `protobuf:"bytes,1,rep,name=profiles" json:"profiles,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *UserProfileUpdate) Reset()                    { *m = UserProfileUpdate{} }
func (m *UserProfileUpdate) String() string            { return proto.CompactTextString(m) }
func (*UserProfileUpdate) ProtoMessage()               {}
func (*UserProfileUpdate) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{7} }

func (m *UserProfileUpdate) GetProfiles() []*UserProfile {
	if m != nil {
		return m.Profiles
	}
	return nil
}

type UserProfileQuery struct {
	UserIDs          []string `protobuf:"bytes,1,rep,name=userIDs" json:"userIDs,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *UserProfileQuery) Reset()                    { *m = UserProfileQuery{} }
func (m *UserProfileQuery) String() string            { return proto.CompactTextString(m) }
func (*UserProfileQuery) ProtoMessage()               {}
func (*UserProfileQuery) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{8} }

func (m *UserProfileQuery) GetUserIDs() []string {
	if m != nil {
		return m.UserIDs
	}
	return nil
}

type ConfirmationRequest struct {
	ContactInfo      *ContactInfo `protobuf:"bytes,1,opt,name=contactInfo" json:"contactInfo,omitempty"`
	Profile          *UserProfile `protobuf:"bytes,2,opt,name=profile" json:"profile,omitempty"`
	ConfirmationCode *string      `protobuf:"bytes,3,opt,name=confirmationCode" json:"confirmationCode,omitempty"`
	InviterUserID    *string      `protobuf:"bytes,4,opt,name=inviterUserID" json:"inviterUserID,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ConfirmationRequest) Reset()                    { *m = ConfirmationRequest{} }
func (m *ConfirmationRequest) String() string            { return proto.CompactTextString(m) }
func (*ConfirmationRequest) ProtoMessage()               {}
func (*ConfirmationRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{9} }

func (m *ConfirmationRequest) GetContactInfo() *ContactInfo {
	if m != nil {
		return m.ContactInfo
	}
	return nil
}

func (m *ConfirmationRequest) GetProfile() *UserProfile {
	if m != nil {
		return m.Profile
	}
	return nil
}

func (m *ConfirmationRequest) GetConfirmationCode() string {
	if m != nil && m.ConfirmationCode != nil {
		return *m.ConfirmationCode
	}
	return ""
}

func (m *ConfirmationRequest) GetInviterUserID() string {
	if m != nil && m.InviterUserID != nil {
		return *m.InviterUserID
	}
	return ""
}

type ProfilesFromContactInfo struct {
	Profiles         []*UserProfile `protobuf:"bytes,1,rep,name=profiles" json:"profiles,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *ProfilesFromContactInfo) Reset()                    { *m = ProfilesFromContactInfo{} }
func (m *ProfilesFromContactInfo) String() string            { return proto.CompactTextString(m) }
func (*ProfilesFromContactInfo) ProtoMessage()               {}
func (*ProfilesFromContactInfo) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{10} }

func (m *ProfilesFromContactInfo) GetProfiles() []*UserProfile {
	if m != nil {
		return m.Profiles
	}
	return nil
}

func init() {
	proto.RegisterType((*SocialIdentity)(nil), "BlitzMessage.SocialIdentity")
	proto.RegisterType((*ContactInfo)(nil), "BlitzMessage.ContactInfo")
	proto.RegisterType((*Employment)(nil), "BlitzMessage.Employment")
	proto.RegisterType((*Education)(nil), "BlitzMessage.Education")
	proto.RegisterType((*ImageData)(nil), "BlitzMessage.ImageData")
	proto.RegisterType((*UserProfile)(nil), "BlitzMessage.UserProfile")
	proto.RegisterType((*ImageUpload)(nil), "BlitzMessage.ImageUpload")
	proto.RegisterType((*UserProfileUpdate)(nil), "BlitzMessage.UserProfileUpdate")
	proto.RegisterType((*UserProfileQuery)(nil), "BlitzMessage.UserProfileQuery")
	proto.RegisterType((*ConfirmationRequest)(nil), "BlitzMessage.ConfirmationRequest")
	proto.RegisterType((*ProfilesFromContactInfo)(nil), "BlitzMessage.ProfilesFromContactInfo")
	proto.RegisterEnum("BlitzMessage.ContactType", ContactType_name, ContactType_value)
	proto.RegisterEnum("BlitzMessage.UserStatus", UserStatus_name, UserStatus_value)
	proto.RegisterEnum("BlitzMessage.Gender", Gender_name, Gender_value)
	proto.RegisterEnum("BlitzMessage.ImageContent", ImageContent_name, ImageContent_value)
}

var fileDescriptor8 = []byte{
	// 1044 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0xdd, 0x72, 0xdb, 0x54,
	0x10, 0xae, 0x12, 0xc7, 0xb1, 0x57, 0xb6, 0xab, 0xa8, 0x81, 0x2a, 0x81, 0x8b, 0x8e, 0x61, 0x86,
	0x62, 0x68, 0x60, 0xdc, 0x19, 0x66, 0xe0, 0xa2, 0x43, 0xed, 0xfc, 0x8c, 0x67, 0x08, 0x94, 0xc8,
	0xe2, 0x82, 0xbb, 0x63, 0xe9, 0xc4, 0x3e, 0x89, 0xa5, 0x23, 0x8e, 0x8e, 0x43, 0xcd, 0x13, 0x70,
	0xcd, 0x23, 0xf0, 0x0a, 0x3c, 0x03, 0x2f, 0xd0, 0xe7, 0xe0, 0x21, 0xd8, 0xb3, 0x92, 0x23, 0x39,
	0x4d, 0xcd, 0x70, 0xa7, 0x5d, 0xed, 0xef, 0xb7, 0xfb, 0xed, 0x01, 0x37, 0xc8, 0xb8, 0x7a, 0xa5,
	0xe4, 0xa5, 0x98, 0xf3, 0xec, 0x28, 0x55, 0x52, 0x4b, 0xb7, 0x35, 0x98, 0x0b, 0xfd, 0xdb, 0x39,
	0xcf, 0x32, 0x36, 0xe5, 0x87, 0x1f, 0xc8, 0xc9, 0x15, 0x0f, 0xb5, 0xb8, 0xe1, 0xe1, 0xb3, 0x88,
	0x67, 0xa1, 0x12, 0xa9, 0x96, 0x2a, 0x37, 0x3d, 0xb4, 0xc7, 0xcb, 0x74, 0xe5, 0xd7, 0xfd, 0xdb,
	0x82, 0x8e, 0x2f, 0x43, 0xc1, 0xe6, 0xa3, 0x88, 0x27, 0x5a, 0xe8, 0xa5, 0xfb, 0x1e, 0xb4, 0x33,
	0xd2, 0xf8, 0x5c, 0xdd, 0x88, 0x90, 0x7b, 0xd6, 0x93, 0xad, 0xa7, 0x4d, 0xd7, 0x81, 0x46, 0xae,
	0x1e, 0x1d, 0x7b, 0x5b, 0x4f, 0xac, 0x5c, 0xb3, 0xc0, 0x4a, 0xbe, 0x67, 0x31, 0xf7, 0xb6, 0x49,
	0xf3, 0x08, 0xec, 0x48, 0x64, 0xe9, 0x9c, 0x2d, 0x49, 0x59, 0x23, 0xe5, 0x43, 0xd8, 0x35, 0x66,
	0xc1, 0xc5, 0xc8, 0xdb, 0x21, 0xc5, 0x1e, 0x34, 0xd9, 0x42, 0xcf, 0xc6, 0xf2, 0x9a, 0x27, 0x5e,
	0x9d, 0x54, 0x9f, 0x01, 0x18, 0xd5, 0xc9, 0xeb, 0x54, 0x28, 0xee, 0xed, 0xa2, 0xce, 0xee, 0x3f,
	0x3e, 0xaa, 0xf6, 0x74, 0x34, 0x16, 0x31, 0xcf, 0x34, 0x8b, 0x53, 0xd7, 0xcd, 0x8d, 0x7d, 0x1e,
	0x2a, 0xae, 0xbd, 0x86, 0x09, 0xd0, 0x9d, 0x80, 0x3d, 0x94, 0x89, 0x66, 0xa1, 0x1e, 0x25, 0x97,
	0xd2, 0x3d, 0x02, 0x3b, 0xcc, 0x45, 0xd3, 0x2c, 0x75, 0xd0, 0xe9, 0x1f, 0xac, 0x07, 0x1c, 0x96,
	0x06, 0xa6, 0xc6, 0xc2, 0x1e, 0x7b, 0x33, 0xdd, 0x62, 0x0e, 0x91, 0xfd, 0xc4, 0x95, 0xb8, 0x14,
	0x3c, 0xa2, 0xee, 0x1a, 0xdd, 0x3f, 0x2c, 0x80, 0x93, 0x38, 0x9d, 0xcb, 0x65, 0x8c, 0x50, 0x99,
	0xf6, 0xaf, 0xe4, 0x64, 0x2c, 0xf4, 0xdc, 0x24, 0x28, 0xda, 0x0f, 0x65, 0x9c, 0xb2, 0x24, 0x6f,
	0xff, 0x16, 0xa5, 0xb9, 0x0c, 0x99, 0x16, 0x32, 0x29, 0x50, 0x42, 0x8d, 0x48, 0xa2, 0x45, 0xa6,
	0xd5, 0xb2, 0x80, 0xe8, 0x29, 0x34, 0xb4, 0x69, 0x0f, 0x5d, 0x09, 0x23, 0xbb, 0xff, 0xfe, 0x3d,
	0xcd, 0xe3, 0x5f, 0x53, 0x68, 0xb6, 0x88, 0x63, 0x86, 0xae, 0x84, 0x5c, 0xf7, 0x1a, 0x9a, 0x27,
	0xd1, 0x22, 0x8f, 0x6f, 0xaa, 0xce, 0xc2, 0x99, 0x94, 0x73, 0xca, 0x9f, 0x17, 0xd5, 0x81, 0x7a,
	0xc4, 0xa7, 0x8a, 0x57, 0xea, 0xe1, 0x71, 0x3a, 0x63, 0x99, 0xc8, 0x8a, 0x7a, 0xaa, 0xd9, 0x6b,
	0x9b, 0xb2, 0x77, 0xff, 0xb2, 0xa0, 0x39, 0x8a, 0x51, 0x75, 0xcc, 0x34, 0x73, 0xbf, 0x84, 0x96,
	0x30, 0x82, 0x01, 0x12, 0x01, 0xa1, 0x7c, 0x9d, 0xfe, 0xe1, 0xba, 0xef, 0xa8, 0x62, 0x41, 0xa8,
	0x1a, 0x79, 0xb0, 0xd4, 0x3c, 0xa3, 0x7a, 0x5a, 0x39, 0x68, 0xf4, 0x9b, 0x46, 0x55, 0x42, 0x64,
	0x0c, 0x83, 0x8b, 0xef, 0x0a, 0x88, 0x7a, 0xd0, 0x8c, 0x98, 0xe6, 0x2f, 0xa3, 0x08, 0xe7, 0xb1,
	0xb3, 0x79, 0x41, 0xda, 0xb0, 0x13, 0xaa, 0xf0, 0x79, 0x9f, 0x20, 0x6a, 0x77, 0xff, 0xa9, 0x81,
	0x5d, 0xa1, 0x8c, 0x41, 0xc4, 0x2c, 0x24, 0xee, 0x71, 0x8e, 0xd0, 0xe7, 0x00, 0x46, 0xf6, 0x35,
	0xd3, 0x8b, 0xbc, 0xaa, 0x4e, 0xdf, 0x5b, 0x8f, 0x1d, 0xdc, 0xfe, 0x77, 0x9f, 0x41, 0x0b, 0xf7,
	0x8e, 0xf0, 0x46, 0x14, 0xf2, 0x82, 0x37, 0xd4, 0xf2, 0x29, 0x8e, 0x9f, 0x65, 0xda, 0xe7, 0x7c,
	0x05, 0xee, 0x3b, 0x4d, 0x5b, 0x50, 0x4b, 0xcc, 0xdc, 0x72, 0x96, 0x7c, 0x0c, 0xf5, 0x29, 0x4f,
	0x22, 0xae, 0xa8, 0x8b, 0x4e, 0x7f, 0x7f, 0xdd, 0xed, 0x8c, 0xfe, 0x99, 0xf0, 0x13, 0xa1, 0xf4,
	0x2c, 0x62, 0xcb, 0xff, 0xa2, 0xcd, 0x27, 0x50, 0x27, 0x4c, 0x33, 0xa4, 0xcc, 0xf6, 0xdb, 0x86,
	0xe5, 0x5c, 0xbf, 0x02, 0x27, 0xab, 0x9e, 0x04, 0x81, 0x2e, 0x4d, 0x72, 0xf9, 0x70, 0xdd, 0xe5,
	0xce, 0xe1, 0x28, 0x49, 0x67, 0x38, 0xe8, 0x01, 0xb9, 0xdc, 0x4f, 0x3a, 0x22, 0xe9, 0x73, 0xd8,
	0x0b, 0x17, 0x4a, 0xa1, 0x77, 0xc9, 0x2a, 0xcf, 0xa6, 0x26, 0xee, 0xc0, 0x5f, 0x61, 0x1d, 0x0e,
	0x8b, 0x97, 0xd6, 0x2d, 0xca, 0xf1, 0x6e, 0x6b, 0xdc, 0x1a, 0xbe, 0x62, 0x87, 0xd7, 0xbe, 0xaf,
	0xed, 0x92, 0x3c, 0x78, 0xf7, 0xf8, 0xeb, 0x94, 0x2b, 0x2d, 0x32, 0x3e, 0x66, 0xd3, 0xcc, 0xeb,
	0xa0, 0x7d, 0xd3, 0xdd, 0xc7, 0x2d, 0xc7, 0xf5, 0x54, 0x88, 0x22, 0x69, 0x1f, 0x92, 0xf6, 0x00,
	0xf6, 0x26, 0x2c, 0xbc, 0x9e, 0x2a, 0xb9, 0x48, 0x22, 0xbf, 0x60, 0xa4, 0x43, 0x8c, 0xfc, 0x1a,
	0x6c, 0xc2, 0x32, 0xc0, 0x2a, 0x58, 0x64, 0x4a, 0x10, 0x2b, 0x68, 0x71, 0xe1, 0x36, 0x21, 0xdf,
	0xfd, 0x16, 0xf6, 0x2a, 0x8b, 0x1a, 0xa4, 0x66, 0xe3, 0xf1, 0x36, 0x36, 0xd2, 0xe2, 0xd8, 0x17,
	0xfe, 0x07, 0x6f, 0x2f, 0x67, 0xe1, 0xd2, 0xfd, 0x08, 0x9c, 0x8a, 0xf8, 0xe3, 0x82, 0xab, 0xe5,
	0xea, 0x00, 0x8f, 0x8e, 0x73, 0xff, 0x66, 0xf7, 0x4f, 0x0b, 0x1e, 0xe1, 0x20, 0x2e, 0x85, 0x8a,
	0xa9, 0xf5, 0x0b, 0xfe, 0xcb, 0x02, 0xdb, 0xbb, 0x3b, 0x40, 0x8b, 0x46, 0xb1, 0x61, 0x80, 0x3d,
	0xd8, 0x2d, 0x2a, 0x23, 0xd6, 0x6c, 0x2a, 0xcc, 0xf5, 0xc0, 0x09, 0x2b, 0x29, 0x87, 0x32, 0x5a,
	0x71, 0x1d, 0x71, 0x17, 0xc9, 0x8d, 0x40, 0x88, 0x83, 0x9c, 0x95, 0x44, 0xf8, 0xee, 0x29, 0x3c,
	0x5e, 0xbd, 0x71, 0xa7, 0x4a, 0xc6, 0xd5, 0xbc, 0xff, 0x07, 0x91, 0xde, 0xcf, 0xb7, 0x2f, 0x03,
	0x5d, 0xfa, 0x36, 0x34, 0x87, 0xe3, 0x20, 0xb9, 0x4e, 0xe4, 0xaf, 0x89, 0xf3, 0x00, 0x6f, 0x01,
	0x0c, 0xc7, 0xaf, 0x66, 0x32, 0xe1, 0xfe, 0xb9, 0xef, 0x58, 0xae, 0x0d, 0xbb, 0xc3, 0xf1, 0x49,
	0xcc, 0xc4, 0xdc, 0xd9, 0xc2, 0x6b, 0x55, 0x1f, 0x8e, 0x87, 0x33, 0xa6, 0x9d, 0x6d, 0x3c, 0x53,
	0x0f, 0x87, 0x63, 0xbf, 0xfa, 0x2e, 0x3a, 0xb5, 0xde, 0x15, 0x40, 0xe5, 0x32, 0x60, 0xe8, 0xc0,
	0x2f, 0x43, 0x93, 0x38, 0xc0, 0xdb, 0x7f, 0xcd, 0x23, 0x8c, 0x4c, 0xe2, 0x88, 0x1a, 0x8d, 0x30,
	0x76, 0x0b, 0x1a, 0x81, 0xff, 0x92, 0xde, 0x68, 0x8c, 0xee, 0x40, 0x2b, 0xf0, 0x8b, 0x91, 0x88,
	0x64, 0xea, 0xd4, 0x70, 0x68, 0xf6, 0xad, 0x06, 0x1d, 0x76, 0x7a, 0xdf, 0x40, 0xbd, 0xe0, 0x3c,
	0xba, 0x9e, 0x95, 0x69, 0xb0, 0xe2, 0xb3, 0x53, 0x1e, 0xb3, 0x39, 0xc7, 0x24, 0x4d, 0xd8, 0x39,
	0x3b, 0x37, 0x9f, 0x54, 0xfc, 0xd9, 0x0f, 0x7a, 0xc6, 0x95, 0xb3, 0xdd, 0x3b, 0x85, 0xd6, 0xda,
	0x1d, 0xc6, 0x5a, 0x46, 0xc3, 0x32, 0xc4, 0x1e, 0xb4, 0x51, 0x2c, 0x31, 0xc3, 0x40, 0xfb, 0xe0,
	0xe4, 0xaa, 0xc1, 0xed, 0x96, 0x3b, 0x5b, 0x83, 0x2f, 0xde, 0xbc, 0xd8, 0x82, 0x07, 0x6f, 0x5e,
	0x6c, 0xbb, 0xd6, 0x00, 0x3f, 0x3d, 0x0b, 0x0e, 0xf1, 0xc1, 0x3b, 0x9a, 0x18, 0xec, 0x31, 0x13,
	0x5f, 0x9b, 0xc2, 0xef, 0x96, 0xf5, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5e, 0xb7, 0x66, 0xd0,
	0xab, 0x08, 0x00, 0x00,
}
