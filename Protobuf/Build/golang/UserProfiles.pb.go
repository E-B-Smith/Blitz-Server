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
	// 1035 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0xdd, 0x6e, 0xe3, 0xd4,
	0x13, 0x5f, 0xb7, 0x69, 0x9a, 0x8c, 0xd3, 0xd4, 0xf5, 0xf6, 0xff, 0x5f, 0xb7, 0x70, 0xb1, 0x0a,
	0x48, 0x2c, 0x81, 0x2d, 0x28, 0x48, 0x48, 0x70, 0xb1, 0x62, 0x93, 0x7e, 0x28, 0x12, 0x85, 0xa5,
	0x8e, 0xb9, 0xe0, 0xee, 0xc4, 0x9e, 0x26, 0xa7, 0x8d, 0x7d, 0xcc, 0xf1, 0x49, 0xd9, 0xf0, 0x04,
	0x5c, 0xf3, 0x08, 0x88, 0x57, 0xe1, 0x05, 0xf6, 0x39, 0x78, 0x08, 0xe6, 0x1c, 0x3b, 0xb1, 0xd3,
	0xed, 0x06, 0x71, 0xe7, 0x19, 0xcf, 0xe7, 0x6f, 0xe6, 0x37, 0x07, 0xdc, 0x20, 0x43, 0xf9, 0x4a,
	0x8a, 0x6b, 0x3e, 0xc3, 0xec, 0x24, 0x95, 0x42, 0x09, 0xb7, 0xd5, 0x9f, 0x71, 0xf5, 0xeb, 0x25,
	0x66, 0x19, 0x9b, 0xe0, 0xf1, 0x7b, 0x62, 0x7c, 0x83, 0xa1, 0xe2, 0x77, 0x18, 0x3e, 0x8f, 0x30,
	0x0b, 0x25, 0x4f, 0x95, 0x90, 0xb9, 0xe9, 0xb1, 0x3d, 0x5a, 0xa4, 0x4b, 0xbf, 0xce, 0x5f, 0x16,
	0xb4, 0x7d, 0x11, 0x72, 0x36, 0x1b, 0x46, 0x98, 0x28, 0xae, 0x16, 0xee, 0xff, 0x60, 0x2f, 0x33,
	0x1a, 0x1f, 0xe5, 0x1d, 0x0f, 0xd1, 0xb3, 0x9e, 0x6e, 0x3d, 0x6b, 0xba, 0x0e, 0x34, 0x72, 0xf5,
	0xf0, 0xd4, 0xdb, 0x7a, 0x6a, 0xe5, 0x9a, 0x39, 0x55, 0xf2, 0x1d, 0x8b, 0xd1, 0xdb, 0x36, 0x9a,
	0xc7, 0x60, 0x47, 0x3c, 0x4b, 0x67, 0x6c, 0x61, 0x94, 0x35, 0xa3, 0xdc, 0x87, 0x5d, 0x6d, 0x16,
	0x5c, 0x0d, 0xbd, 0x1d, 0xa3, 0x38, 0x80, 0x26, 0x9b, 0xab, 0xe9, 0x48, 0xdc, 0x62, 0xe2, 0xd5,
	0x8d, 0xea, 0x13, 0x00, 0xad, 0x3a, 0x7b, 0x9d, 0x72, 0x89, 0xde, 0x2e, 0xe9, 0xec, 0xde, 0x93,
	0x93, 0x6a, 0x4f, 0x27, 0x23, 0x1e, 0x63, 0xa6, 0x58, 0x9c, 0xba, 0x6e, 0x6e, 0xec, 0x63, 0x28,
	0x51, 0x79, 0x0d, 0x1d, 0xa0, 0x33, 0x06, 0x7b, 0x20, 0x12, 0xc5, 0x42, 0x35, 0x4c, 0xae, 0x85,
	0x7b, 0x02, 0x76, 0x98, 0x8b, 0xba, 0x59, 0xd3, 0x41, 0xbb, 0x77, 0xb4, 0x1e, 0x70, 0x50, 0x1a,
	0xe8, 0x1a, 0x0b, 0x7b, 0xea, 0x4d, 0x77, 0x4b, 0x39, 0x78, 0xf6, 0x23, 0x4a, 0x7e, 0xcd, 0x31,
	0x32, 0xdd, 0x35, 0x3a, 0xbf, 0x5b, 0x00, 0x67, 0x71, 0x3a, 0x13, 0x8b, 0x98, 0xa0, 0xd2, 0xed,
	0xdf, 0x88, 0xf1, 0x88, 0xab, 0x99, 0x4e, 0x50, 0xb4, 0x1f, 0x8a, 0x38, 0x65, 0x49, 0xde, 0xfe,
	0x0a, 0xa5, 0x99, 0x08, 0x99, 0xe2, 0x22, 0x29, 0x50, 0x22, 0x0d, 0x4f, 0xa2, 0x79, 0xa6, 0xe4,
	0xa2, 0x80, 0xe8, 0x19, 0x34, 0x94, 0x6e, 0x8f, 0x5c, 0x0d, 0x46, 0x76, 0xef, 0xff, 0x0f, 0x34,
	0x4f, 0x7f, 0x75, 0xa1, 0xd9, 0x3c, 0x8e, 0x19, 0xb9, 0x1a, 0xe4, 0x3a, 0xb7, 0xd0, 0x3c, 0x8b,
	0xe6, 0x79, 0x7c, 0x5d, 0x75, 0x16, 0x4e, 0x85, 0x98, 0x99, 0xfc, 0x79, 0x51, 0x6d, 0xa8, 0x47,
	0x38, 0x91, 0x58, 0xa9, 0x07, 0xe3, 0x74, 0xca, 0x32, 0x9e, 0x15, 0xf5, 0x54, 0xb3, 0xd7, 0x36,
	0x65, 0xef, 0xfc, 0x69, 0x41, 0x73, 0x18, 0x93, 0xea, 0x94, 0x29, 0xe6, 0x7e, 0x0e, 0x2d, 0xae,
	0x05, 0x0d, 0x24, 0x01, 0x62, 0xf2, 0xb5, 0x7b, 0xc7, 0xeb, 0xbe, 0xc3, 0x8a, 0x85, 0x41, 0x55,
	0xcb, 0xfd, 0x85, 0xc2, 0xcc, 0xd4, 0xd3, 0xca, 0x41, 0x33, 0xbf, 0xcd, 0xa8, 0x4a, 0x88, 0xb4,
	0x61, 0x70, 0xf5, 0x6d, 0x01, 0x51, 0x17, 0x9a, 0x11, 0x53, 0xf8, 0x32, 0x8a, 0x68, 0x1e, 0x3b,
	0x1b, 0x17, 0xa4, 0xf3, 0x77, 0x0d, 0xec, 0x0a, 0x47, 0x34, 0x04, 0x7a, 0x03, 0x69, 0x71, 0x73,
	0x48, 0x3e, 0x05, 0xd0, 0xb2, 0xaf, 0x98, 0x9a, 0xe7, 0x65, 0xb4, 0x7b, 0xde, 0x7a, 0xb0, 0x60,
	0xf5, 0xdf, 0x7d, 0x0e, 0x2d, 0x5a, 0x34, 0x03, 0x30, 0xb5, 0x9d, 0x57, 0xb8, 0x61, 0x3b, 0x3f,
	0xa6, 0x79, 0xb3, 0x4c, 0xf9, 0x88, 0x4b, 0x34, 0xdf, 0x69, 0xda, 0x82, 0x5a, 0xa2, 0x07, 0x95,
	0xd3, 0xe2, 0x43, 0xa8, 0x4f, 0x30, 0x89, 0x50, 0x9a, 0xc9, 0xb6, 0x7b, 0x87, 0xeb, 0x6e, 0x17,
	0xe6, 0x9f, 0x0e, 0x3f, 0xe6, 0x52, 0x4d, 0x23, 0xb6, 0xf8, 0x37, 0x9e, 0x7c, 0x04, 0x75, 0x03,
	0x62, 0x46, 0x1c, 0xd9, 0x7e, 0xdb, 0xb0, 0x1c, 0xe4, 0x97, 0xe0, 0x64, 0xd5, 0x1b, 0xc0, 0xc9,
	0xa5, 0x69, 0x5c, 0xde, 0x5f, 0x77, 0xb9, 0x77, 0x29, 0x4a, 0x96, 0x69, 0xd2, 0x79, 0x60, 0x5c,
	0x1e, 0x66, 0x99, 0x61, 0xe5, 0x17, 0x70, 0x10, 0xce, 0xa5, 0x24, 0xef, 0x92, 0x46, 0x9e, 0x6d,
	0x9a, 0xb8, 0x07, 0x7f, 0x85, 0x66, 0x34, 0x2c, 0x2c, 0xad, 0x5b, 0x26, 0xc7, 0xbb, 0xad, 0x69,
	0x4d, 0x70, 0x49, 0x07, 0x6f, 0xef, 0xa1, 0xb6, 0x4b, 0xb6, 0xd0, 0xa1, 0xc3, 0xd7, 0x29, 0x4a,
	0xc5, 0x33, 0x1c, 0xb1, 0x49, 0xe6, 0xb5, 0xc9, 0xbe, 0xe9, 0x1e, 0xd2, 0x5a, 0xd3, 0x3e, 0x4a,
	0x42, 0xd1, 0x68, 0xf7, 0x8d, 0xf6, 0x08, 0x0e, 0xc6, 0x2c, 0xbc, 0x9d, 0x48, 0x31, 0x4f, 0x22,
	0xbf, 0xa0, 0xa0, 0x63, 0x28, 0xf8, 0x15, 0xd8, 0x06, 0xcb, 0x80, 0xaa, 0x60, 0x91, 0x2e, 0x81,
	0x2f, 0xa1, 0xa5, 0x85, 0xdb, 0x84, 0x7c, 0xe7, 0x1b, 0x38, 0xa8, 0x2c, 0x6a, 0x90, 0xea, 0x15,
	0xa7, 0x63, 0xd8, 0x48, 0x8b, 0xeb, 0x5e, 0xf8, 0x1f, 0xbd, 0xbd, 0x9c, 0x85, 0x4b, 0xe7, 0x03,
	0x70, 0x2a, 0xe2, 0x0f, 0x73, 0x94, 0x8b, 0xe5, 0xc5, 0x1d, 0x9e, 0xe6, 0xfe, 0xcd, 0xce, 0x1f,
	0x16, 0x3c, 0xa6, 0x41, 0x5c, 0x73, 0x19, 0x9b, 0xd6, 0xaf, 0xf0, 0xe7, 0x39, 0xb5, 0x77, 0x7f,
	0x80, 0x96, 0x19, 0xc5, 0x86, 0x01, 0x76, 0x61, 0xb7, 0xa8, 0xcc, 0xb0, 0x66, 0x53, 0x61, 0xae,
	0x07, 0x4e, 0x58, 0x49, 0x39, 0x10, 0xd1, 0x92, 0xdc, 0x84, 0x3b, 0x4f, 0xee, 0x38, 0x41, 0x1c,
	0xe4, 0xac, 0x34, 0x0c, 0xef, 0x9c, 0xc3, 0x93, 0xe5, 0xa3, 0x76, 0x2e, 0x45, 0x5c, 0xcd, 0xfb,
	0x5f, 0x10, 0xe9, 0xfe, 0xb4, 0x7a, 0x0a, 0xcc, 0x69, 0xdf, 0x83, 0xe6, 0x60, 0x14, 0x24, 0xb7,
	0x89, 0xf8, 0x25, 0x71, 0x1e, 0xd1, 0x2d, 0x80, 0xc1, 0xe8, 0xd5, 0x54, 0x24, 0xe8, 0x5f, 0xfa,
	0x8e, 0xe5, 0xda, 0xb0, 0x3b, 0x18, 0x9d, 0xc5, 0x8c, 0xcf, 0x9c, 0x2d, 0x3a, 0x4f, 0xf5, 0xc1,
	0x68, 0x30, 0x65, 0xca, 0xd9, 0xa6, 0xbb, 0xb4, 0x3f, 0x18, 0xf9, 0xd5, 0x87, 0xd0, 0xa9, 0x75,
	0x6f, 0x00, 0x2a, 0x97, 0x81, 0x42, 0x07, 0x7e, 0x19, 0xda, 0x88, 0x7d, 0x3a, 0xf6, 0xb7, 0x18,
	0x51, 0x64, 0x23, 0x0e, 0x4d, 0xa3, 0x11, 0xc5, 0x6e, 0x41, 0x23, 0xf0, 0x5f, 0x9a, 0x47, 0x99,
	0xa2, 0x3b, 0xd0, 0x0a, 0xfc, 0x62, 0x24, 0x3c, 0x99, 0x38, 0x35, 0x1a, 0x9a, 0xbd, 0xd2, 0x90,
	0xc3, 0x4e, 0xf7, 0x6b, 0xa8, 0x17, 0x9c, 0x27, 0xd7, 0x8b, 0x32, 0x0d, 0x55, 0x7c, 0x71, 0x8e,
	0x31, 0x9b, 0x21, 0x25, 0x69, 0xc2, 0xce, 0xc5, 0xa5, 0xfe, 0x34, 0xc5, 0x5f, 0x7c, 0xaf, 0xa6,
	0x28, 0x9d, 0xed, 0xee, 0x39, 0xb4, 0xd6, 0x0e, 0x2f, 0xd5, 0x32, 0x1c, 0x94, 0x21, 0x0e, 0x60,
	0x8f, 0xc4, 0x12, 0x33, 0x0a, 0x74, 0x08, 0x4e, 0xae, 0xea, 0xaf, 0xb6, 0xdc, 0xd9, 0xea, 0x7f,
	0xf6, 0xe6, 0xc5, 0x16, 0x3c, 0x7a, 0xf3, 0x62, 0xdb, 0xb5, 0xfa, 0xf4, 0xe9, 0x59, 0x70, 0x4c,
	0x2f, 0xdc, 0xc9, 0x58, 0x63, 0x4f, 0x99, 0x70, 0x6d, 0x0a, 0xbf, 0x59, 0xd6, 0x3f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x0d, 0x7b, 0x42, 0x4a, 0x9c, 0x08, 0x00, 0x00,
}
