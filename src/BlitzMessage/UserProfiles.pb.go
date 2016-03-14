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
func (ContactType) EnumDescriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

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
func (UserStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

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
func (Gender) EnumDescriptor() ([]byte, []int) { return fileDescriptor10, []int{2} }

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
func (ImageContent) EnumDescriptor() ([]byte, []int) { return fileDescriptor10, []int{3} }

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
func (*SocialIdentity) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

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
func (*ContactInfo) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

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
func (*Employment) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{2} }

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
func (*Education) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{3} }

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
	Crc32            *int64        `protobuf:"varint,6,opt,name=crc32" json:"crc32,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *ImageData) Reset()                    { *m = ImageData{} }
func (m *ImageData) String() string            { return proto.CompactTextString(m) }
func (*ImageData) ProtoMessage()               {}
func (*ImageData) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{4} }

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

func (m *ImageData) GetCrc32() int64 {
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
	ExpertiseTags     []*EntityTag      `protobuf:"bytes,14,rep,name=expertiseTags" json:"expertiseTags,omitempty"`
	InterestTags      []string          `protobuf:"bytes,15,rep,name=interestTags" json:"interestTags,omitempty"`
	BackgroundSummary *string           `protobuf:"bytes,16,opt,name=backgroundSummary" json:"backgroundSummary,omitempty"`
	XXX_unrecognized  []byte            `json:"-"`
}

func (m *UserProfile) Reset()                    { *m = UserProfile{} }
func (m *UserProfile) String() string            { return proto.CompactTextString(m) }
func (*UserProfile) ProtoMessage()               {}
func (*UserProfile) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{5} }

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

func (m *UserProfile) GetExpertiseTags() []*EntityTag {
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
func (*ImageUpload) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{6} }

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
func (*UserProfileUpdate) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{7} }

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
func (*UserProfileQuery) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{8} }

func (m *UserProfileQuery) GetUserIDs() []string {
	if m != nil {
		return m.UserIDs
	}
	return nil
}

type ConfirmationRequest struct {
	ContactInfo      *ContactInfo `protobuf:"bytes,1,opt,name=contactInfo" json:"contactInfo,omitempty"`
	UserProfile      *UserProfile `protobuf:"bytes,2,opt,name=userProfile" json:"userProfile,omitempty"`
	ConfirmationCode *string      `protobuf:"bytes,3,opt,name=confirmationCode" json:"confirmationCode,omitempty"`
	InviterUserID    *string      `protobuf:"bytes,4,opt,name=inviterUserID" json:"inviterUserID,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ConfirmationRequest) Reset()                    { *m = ConfirmationRequest{} }
func (m *ConfirmationRequest) String() string            { return proto.CompactTextString(m) }
func (*ConfirmationRequest) ProtoMessage()               {}
func (*ConfirmationRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{9} }

func (m *ConfirmationRequest) GetContactInfo() *ContactInfo {
	if m != nil {
		return m.ContactInfo
	}
	return nil
}

func (m *ConfirmationRequest) GetUserProfile() *UserProfile {
	if m != nil {
		return m.UserProfile
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
func (*ProfilesFromContactInfo) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{10} }

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

var fileDescriptor10 = []byte{
	// 1056 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0xdd, 0x72, 0xdb, 0x44,
	0x14, 0xae, 0xe2, 0xd8, 0xb5, 0x8f, 0x6c, 0x47, 0x51, 0x03, 0x55, 0x02, 0x17, 0x1d, 0xc3, 0x0c,
	0x25, 0x50, 0xc3, 0xb8, 0x33, 0xcc, 0xc0, 0x45, 0x87, 0xda, 0xf9, 0x19, 0xcf, 0x10, 0x28, 0x91,
	0xc4, 0x05, 0x77, 0x6b, 0x69, 0x63, 0x6f, 0x22, 0x69, 0xc5, 0x6a, 0x15, 0x6a, 0xee, 0x99, 0xe1,
	0x9a, 0x87, 0xe0, 0x05, 0x78, 0x06, 0x5e, 0xa0, 0x4f, 0xc4, 0xd9, 0x95, 0x6c, 0xc9, 0x69, 0x6a,
	0x86, 0x3b, 0xed, 0xd1, 0xf9, 0xfd, 0xce, 0xf9, 0xce, 0x01, 0xdb, 0xcf, 0xa8, 0x78, 0x25, 0xf8,
	0x15, 0x8b, 0x68, 0x36, 0x4c, 0x05, 0x97, 0xdc, 0xee, 0x8e, 0x23, 0x26, 0x7f, 0xbb, 0xa0, 0x59,
	0x46, 0xe6, 0xf4, 0xe8, 0x03, 0x3e, 0xbb, 0xa6, 0x81, 0x64, 0xb7, 0x34, 0x78, 0x16, 0xd2, 0x2c,
	0x10, 0x2c, 0x95, 0x5c, 0x14, 0xaa, 0x47, 0xa6, 0xb7, 0x4c, 0x57, 0x76, 0x47, 0xd6, 0x69, 0x22,
	0x99, 0x5c, 0x7a, 0x64, 0x5e, 0x4a, 0x06, 0xff, 0x18, 0xd0, 0x77, 0x79, 0xc0, 0x48, 0x34, 0x0d,
	0xa9, 0xfe, 0x69, 0xbf, 0x07, 0xbd, 0x4c, 0x4b, 0x5c, 0x2a, 0x6e, 0x59, 0x40, 0x1d, 0xe3, 0xc9,
	0xce, 0xd3, 0x8e, 0x6d, 0x41, 0xbb, 0x10, 0x4f, 0x4f, 0x9c, 0x9d, 0x27, 0x46, 0x21, 0xc9, 0x31,
	0xb7, 0xef, 0x49, 0x4c, 0x9d, 0x86, 0x96, 0x3c, 0x02, 0x33, 0x64, 0x59, 0x1a, 0x91, 0xa5, 0x16,
	0xee, 0x6a, 0xe1, 0x1e, 0x3c, 0x54, 0x6a, 0xfe, 0xe5, 0xd4, 0x69, 0x6a, 0xc1, 0x3e, 0x74, 0x48,
	0x2e, 0x17, 0x1e, 0xbf, 0xa1, 0x89, 0xd3, 0xd2, 0xa2, 0xcf, 0x00, 0x94, 0xe8, 0xf4, 0x75, 0xca,
	0x04, 0x75, 0x1e, 0xa2, 0xcc, 0x1c, 0x3d, 0x1e, 0xd6, 0xab, 0x1c, 0x7a, 0x2c, 0xa6, 0x99, 0x24,
	0x71, 0x6a, 0xdb, 0x85, 0xb2, 0x4b, 0x03, 0x41, 0xa5, 0xd3, 0x56, 0x0e, 0x06, 0x33, 0x30, 0x27,
	0x3c, 0x91, 0x24, 0x90, 0xd3, 0xe4, 0x8a, 0xdb, 0x43, 0x30, 0x83, 0xe2, 0xa9, 0xca, 0xd7, 0x15,
	0xf4, 0x47, 0x87, 0x9b, 0x0e, 0x27, 0x95, 0x82, 0xca, 0xb1, 0xd4, 0xc7, 0xda, 0x54, 0xb5, 0x18,
	0x83, 0x65, 0x3f, 0x51, 0xc1, 0xae, 0x18, 0x0d, 0x75, 0x75, 0xed, 0xc1, 0x9f, 0x06, 0xc0, 0x69,
	0x9c, 0x46, 0x7c, 0x19, 0x23, 0x54, 0xaa, 0xfc, 0x6b, 0x3e, 0xf3, 0x98, 0x8c, 0x54, 0x80, 0xb2,
	0xfc, 0x80, 0xc7, 0x29, 0x49, 0x8a, 0xf2, 0xd7, 0x28, 0x45, 0x3c, 0x20, 0x92, 0xf1, 0xa4, 0x44,
	0x09, 0x25, 0x2c, 0x09, 0xf3, 0x4c, 0x8a, 0x65, 0x09, 0xd1, 0x53, 0x68, 0x4b, 0x55, 0x1e, 0x9a,
	0x6a, 0x8c, 0xcc, 0xd1, 0xfb, 0xf7, 0x14, 0x8f, 0x7f, 0x55, 0xa2, 0x59, 0x1e, 0xc7, 0x04, 0x4d,
	0x35, 0x72, 0x83, 0x1b, 0xe8, 0x9c, 0x86, 0x79, 0xe1, 0x5f, 0x65, 0x9d, 0x05, 0x0b, 0xce, 0x23,
	0x1d, 0xbf, 0x48, 0xaa, 0x0f, 0xad, 0x90, 0xce, 0x05, 0xad, 0xe5, 0x43, 0xe3, 0x74, 0x41, 0x32,
	0x96, 0x95, 0xf9, 0xd4, 0xa3, 0xef, 0x6e, 0x8b, 0x3e, 0xf8, 0xdb, 0x80, 0xce, 0x34, 0x46, 0xd1,
	0x09, 0x91, 0xc4, 0xfe, 0x12, 0xba, 0x4c, 0x3d, 0x14, 0x90, 0x08, 0x88, 0x8e, 0xd7, 0x1f, 0x1d,
	0x6d, 0xda, 0x4e, 0x6b, 0x1a, 0x1a, 0x55, 0xf5, 0x1e, 0x2f, 0x25, 0xcd, 0x74, 0x3e, 0xdd, 0x02,
	0x34, 0xfd, 0x5b, 0xb7, 0xaa, 0x82, 0x48, 0x29, 0xfa, 0x97, 0xdf, 0x95, 0x10, 0x1d, 0x43, 0x27,
	0x24, 0x92, 0xbe, 0x0c, 0x43, 0xec, 0x47, 0x73, 0xfb, 0x80, 0xf4, 0xa0, 0x19, 0x88, 0xe0, 0xf9,
	0x48, 0x43, 0xd4, 0x18, 0xfc, 0xde, 0x04, 0xb3, 0x46, 0x22, 0x85, 0x88, 0x1a, 0x48, 0x9c, 0xe3,
	0x02, 0xa1, 0xcf, 0x01, 0xd4, 0xdb, 0x95, 0x44, 0xe6, 0x45, 0x56, 0xfd, 0x91, 0xb3, 0xe9, 0xdb,
	0x5f, 0xff, 0xb7, 0x9f, 0x41, 0x17, 0xe7, 0x4e, 0xe3, 0x8d, 0x28, 0x14, 0x09, 0x6f, 0xc9, 0xe5,
	0x53, 0x6c, 0x3f, 0xc9, 0xa4, 0x4b, 0xe9, 0x0a, 0xdc, 0x77, 0xaa, 0x76, 0x61, 0x37, 0x51, 0x7d,
	0x2b, 0x58, 0xf2, 0x31, 0xb4, 0xe6, 0x34, 0x09, 0xa9, 0xd0, 0x55, 0xf4, 0x47, 0x07, 0x9b, 0x66,
	0xe7, 0xfa, 0x9f, 0x72, 0x3f, 0x63, 0x42, 0x2e, 0x42, 0xb2, 0xfc, 0x2f, 0xda, 0x7c, 0x02, 0x2d,
	0x8d, 0x69, 0x86, 0x94, 0x69, 0xbc, 0xad, 0x58, 0xf5, 0xf5, 0x2b, 0xb0, 0xb2, 0xfa, 0x4a, 0x60,
	0x68, 0xd2, 0xd1, 0x26, 0x1f, 0x6e, 0x9a, 0xdc, 0x59, 0x1c, 0x15, 0xe9, 0x14, 0x07, 0x1d, 0xd0,
	0x26, 0xf7, 0x93, 0x4e, 0x93, 0xf4, 0x39, 0xec, 0x07, 0xb9, 0x10, 0x68, 0x5d, 0xb1, 0xca, 0x31,
	0x75, 0x11, 0x77, 0xe0, 0xaf, 0xb1, 0x0e, 0x9b, 0x45, 0x2b, 0xed, 0xae, 0x8e, 0xf1, 0x6e, 0x6d,
	0x9c, 0x1a, 0xba, 0x62, 0x87, 0xd3, 0xbb, 0xaf, 0xec, 0x8a, 0x3c, 0x43, 0xe8, 0xd1, 0xd7, 0x29,
	0x15, 0x92, 0x65, 0x54, 0x6d, 0x48, 0xa7, 0x7f, 0xaf, 0xfe, 0x6a, 0x83, 0xda, 0x07, 0x38, 0xfe,
	0x38, 0xb7, 0x02, 0xe1, 0xd5, 0xea, 0x7b, 0xa8, 0xde, 0xb1, 0x0f, 0x61, 0x7f, 0x46, 0x82, 0x9b,
	0xb9, 0xe0, 0x79, 0x12, 0xba, 0x25, 0x55, 0x2d, 0x4d, 0xd5, 0xaf, 0xc1, 0xd4, 0x20, 0xfb, 0x98,
	0x1e, 0x09, 0x55, 0x6e, 0x6c, 0x85, 0x39, 0x4e, 0xe2, 0xb6, 0x96, 0x0c, 0xbe, 0x85, 0xfd, 0xda,
	0x04, 0xfb, 0xa9, 0xa2, 0x02, 0x2e, 0xcd, 0x76, 0x5a, 0xde, 0x85, 0xd2, 0xfe, 0xf0, 0xed, 0xa9,
	0x2d, 0x4d, 0x06, 0x1f, 0x81, 0x55, 0x7b, 0xfe, 0x98, 0x53, 0xb1, 0x5c, 0x6d, 0xe6, 0xe9, 0x49,
	0x61, 0xdf, 0x19, 0xfc, 0x65, 0xc0, 0x23, 0xec, 0xd0, 0x15, 0x13, 0xb1, 0xc6, 0xe4, 0x92, 0xfe,
	0x92, 0x63, 0x79, 0x77, 0x3b, 0x6b, 0xe8, 0x1e, 0x6d, 0xe9, 0x2c, 0xea, 0xe7, 0x55, 0x30, 0x4d,
	0xa9, 0x6d, 0xc9, 0xd9, 0x0e, 0x58, 0x41, 0x2d, 0xec, 0x84, 0x87, 0xab, 0x45, 0x80, 0xc7, 0x88,
	0x25, 0xb7, 0x0c, 0x61, 0xf6, 0x0b, 0xca, 0xea, 0x6d, 0x30, 0x38, 0x83, 0xc7, 0xab, 0x93, 0x78,
	0x26, 0x78, 0x5c, 0x8f, 0xfd, 0x7f, 0x50, 0x39, 0xfe, 0x79, 0x7d, 0x36, 0xf4, 0x19, 0xe8, 0x41,
	0x67, 0xe2, 0xf9, 0xc9, 0x4d, 0xc2, 0x7f, 0x4d, 0xac, 0x07, 0xb8, 0x28, 0x60, 0xe2, 0xbd, 0x5a,
	0xf0, 0x84, 0xba, 0x17, 0xae, 0x65, 0xd8, 0x26, 0x3c, 0x9c, 0x78, 0xa7, 0x31, 0x61, 0x91, 0xb5,
	0x83, 0xab, 0xac, 0x35, 0xf1, 0x26, 0x0b, 0x22, 0xad, 0x06, 0xee, 0xb0, 0xbd, 0x89, 0xe7, 0xd6,
	0x8f, 0xa6, 0xb5, 0x7b, 0x7c, 0x0d, 0x50, 0x5b, 0x1b, 0xe8, 0xda, 0x77, 0x2b, 0xd7, 0xfa, 0x39,
	0xc6, 0xc3, 0x70, 0x43, 0x43, 0xf4, 0xac, 0x9f, 0x53, 0x5d, 0x68, 0x88, 0xbe, 0xbb, 0xd0, 0xf6,
	0xdd, 0x97, 0xfa, 0xa4, 0xa3, 0x77, 0x0b, 0xba, 0xbe, 0x5b, 0xb6, 0x85, 0x25, 0x73, 0x6b, 0x17,
	0x1b, 0x67, 0xae, 0x25, 0x68, 0xd0, 0x3c, 0xfe, 0x06, 0x5a, 0xe5, 0x42, 0x40, 0xd3, 0xf3, 0x2a,
	0x0c, 0x66, 0x7c, 0x7e, 0x46, 0x63, 0x12, 0x51, 0x0c, 0xd2, 0x81, 0xe6, 0xf9, 0x85, 0xfa, 0xd4,
	0xc9, 0x9f, 0xff, 0x20, 0x17, 0x54, 0x58, 0x8d, 0xe3, 0x33, 0xe8, 0x6e, 0x2c, 0x69, 0xcc, 0x65,
	0x3a, 0xa9, 0x5c, 0xec, 0x43, 0x0f, 0x9f, 0x15, 0x66, 0xe8, 0xe8, 0x00, 0xac, 0x42, 0x34, 0x5e,
	0x4f, 0xba, 0xb5, 0x33, 0xfe, 0xe2, 0xcd, 0x8b, 0x1d, 0x78, 0xf0, 0xe6, 0x45, 0xc3, 0x36, 0xc6,
	0xf8, 0xe9, 0x18, 0x70, 0x84, 0xd7, 0x70, 0x38, 0x53, 0xd8, 0x63, 0x24, 0xba, 0xd1, 0x85, 0x3f,
	0x0c, 0xe3, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x26, 0x7d, 0x49, 0x36, 0xda, 0x08, 0x00, 0x00,
}
