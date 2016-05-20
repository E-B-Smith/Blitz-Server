// Code generated by protoc-gen-go.
// source: Feed.proto
// DO NOT EDIT!

package BlitzMessage

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"


// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type FeedPostType int32

const (
	FeedPostType_FPUnknown           FeedPostType = 0
	FeedPostType_FPOpenEndedQuestion FeedPostType = 1
	FeedPostType_FPOpenEndedReply    FeedPostType = 2
	FeedPostType_FPSurveyQuestion    FeedPostType = 3
	FeedPostType_FPSurveyAnswer      FeedPostType = 4
)

var FeedPostType_name = map[int32]string{
	0: "FPUnknown",
	1: "FPOpenEndedQuestion",
	2: "FPOpenEndedReply",
	3: "FPSurveyQuestion",
	4: "FPSurveyAnswer",
}
var FeedPostType_value = map[string]int32{
	"FPUnknown":           0,
	"FPOpenEndedQuestion": 1,
	"FPOpenEndedReply":    2,
	"FPSurveyQuestion":    3,
	"FPSurveyAnswer":      4,
}

func (x FeedPostType) Enum() *FeedPostType {
	p := new(FeedPostType)
	*p = x
	return p
}
func (x FeedPostType) String() string {
	return proto.EnumName(FeedPostType_name, int32(x))
}
func (x *FeedPostType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FeedPostType_value, data, "FeedPostType")
	if err != nil {
		return err
	}
	*x = FeedPostType(value)
	return nil
}
func (FeedPostType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type FeedPostScope int32

const (
	FeedPostScope_FPScopeUnknown       FeedPostScope = 0
	FeedPostScope_FPScopeLocalNetwork  FeedPostScope = 1
	FeedPostScope_FPScopeGlobalNetwork FeedPostScope = 2
)

var FeedPostScope_name = map[int32]string{
	0: "FPScopeUnknown",
	1: "FPScopeLocalNetwork",
	2: "FPScopeGlobalNetwork",
}
var FeedPostScope_value = map[string]int32{
	"FPScopeUnknown":       0,
	"FPScopeLocalNetwork":  1,
	"FPScopeGlobalNetwork": 2,
}

func (x FeedPostScope) Enum() *FeedPostScope {
	p := new(FeedPostScope)
	*p = x
	return p
}
func (x FeedPostScope) String() string {
	return proto.EnumName(FeedPostScope_name, int32(x))
}
func (x *FeedPostScope) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FeedPostScope_value, data, "FeedPostScope")
	if err != nil {
		return err
	}
	*x = FeedPostScope(value)
	return nil
}
func (FeedPostScope) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type FeedPostStatus int32

const (
	FeedPostStatus_FPSUnknown FeedPostStatus = 0
	FeedPostStatus_FPSActive  FeedPostStatus = 1
	FeedPostStatus_FPSDeleted FeedPostStatus = 2
)

var FeedPostStatus_name = map[int32]string{
	0: "FPSUnknown",
	1: "FPSActive",
	2: "FPSDeleted",
}
var FeedPostStatus_value = map[string]int32{
	"FPSUnknown": 0,
	"FPSActive":  1,
	"FPSDeleted": 2,
}

func (x FeedPostStatus) Enum() *FeedPostStatus {
	p := new(FeedPostStatus)
	*p = x
	return p
}
func (x FeedPostStatus) String() string {
	return proto.EnumName(FeedPostStatus_name, int32(x))
}
func (x *FeedPostStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FeedPostStatus_value, data, "FeedPostStatus")
	if err != nil {
		return err
	}
	*x = FeedPostStatus(value)
	return nil
}
func (FeedPostStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

type UpdateVerb int32

const (
	UpdateVerb_UVCreate UpdateVerb = 1
	UpdateVerb_UVUpdate UpdateVerb = 2
	UpdateVerb_UVDelete UpdateVerb = 3
)

var UpdateVerb_name = map[int32]string{
	1: "UVCreate",
	2: "UVUpdate",
	3: "UVDelete",
}
var UpdateVerb_value = map[string]int32{
	"UVCreate": 1,
	"UVUpdate": 2,
	"UVDelete": 3,
}

func (x UpdateVerb) Enum() *UpdateVerb {
	p := new(UpdateVerb)
	*p = x
	return p
}
func (x UpdateVerb) String() string {
	return proto.EnumName(UpdateVerb_name, int32(x))
}
func (x *UpdateVerb) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(UpdateVerb_value, data, "UpdateVerb")
	if err != nil {
		return err
	}
	*x = UpdateVerb(value)
	return nil
}
func (UpdateVerb) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

type FeedPost struct {
	PostID                   *string        `protobuf:"bytes,1,opt,name=postID" json:"postID,omitempty"`
	ParentID                 *string        `protobuf:"bytes,2,opt,name=parentID" json:"parentID,omitempty"`
	PostType                 *FeedPostType  `protobuf:"varint,3,opt,name=postType,enum=BlitzMessage.FeedPostType" json:"postType,omitempty"`
	PostScope                *FeedPostScope `protobuf:"varint,4,opt,name=postScope,enum=BlitzMessage.FeedPostScope" json:"postScope,omitempty"`
	UserID                   *string        `protobuf:"bytes,5,opt,name=userID" json:"userID,omitempty"`
	AnonymousPost            *bool          `protobuf:"varint,6,opt,name=anonymousPost,def=0" json:"anonymousPost,omitempty"`
	Timestamp                *Timestamp     `protobuf:"bytes,7,opt,name=timestamp" json:"timestamp,omitempty"`
	TimespanActive           *Timespan      `protobuf:"bytes,8,opt,name=timespanActive" json:"timespanActive,omitempty"`
	HeadlineText             *string        `protobuf:"bytes,9,opt,name=headlineText" json:"headlineText,omitempty"`
	BodyText                 *string        `protobuf:"bytes,10,opt,name=bodyText" json:"bodyText,omitempty"`
	PostTags                 []*EntityTag   `protobuf:"bytes,12,rep,name=postTags" json:"postTags,omitempty"`
	RepliesDeprecated        []*FeedPost    `protobuf:"bytes,13,rep,name=replies_deprecated" json:"replies_deprecated,omitempty"`
	MayAddReply              *bool          `protobuf:"varint,14,opt,name=mayAddReply" json:"mayAddReply,omitempty"`
	MayChooseMulitpleReplies *bool          `protobuf:"varint,15,opt,name=mayChooseMulitpleReplies" json:"mayChooseMulitpleReplies,omitempty"`
	SurveyAnswerSequence     *int32         `protobuf:"varint,16,opt,name=surveyAnswerSequence" json:"surveyAnswerSequence,omitempty"`
	AreMoreReplies           *bool          `protobuf:"varint,17,opt,name=areMoreReplies" json:"areMoreReplies,omitempty"`
	TotalVoteCount           *int32         `protobuf:"varint,18,opt,name=totalVoteCount" json:"totalVoteCount,omitempty"`
	XXX_unrecognized         []byte         `json:"-"`
}

func (m *FeedPost) Reset()                    { *m = FeedPost{} }
func (m *FeedPost) String() string            { return proto.CompactTextString(m) }
func (*FeedPost) ProtoMessage()               {}
func (*FeedPost) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

const Default_FeedPost_AnonymousPost bool = false

func (m *FeedPost) GetPostID() string {
	if m != nil && m.PostID != nil {
		return *m.PostID
	}
	return ""
}

func (m *FeedPost) GetParentID() string {
	if m != nil && m.ParentID != nil {
		return *m.ParentID
	}
	return ""
}

func (m *FeedPost) GetPostType() FeedPostType {
	if m != nil && m.PostType != nil {
		return *m.PostType
	}
	return FeedPostType_FPUnknown
}

func (m *FeedPost) GetPostScope() FeedPostScope {
	if m != nil && m.PostScope != nil {
		return *m.PostScope
	}
	return FeedPostScope_FPScopeUnknown
}

func (m *FeedPost) GetUserID() string {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return ""
}

func (m *FeedPost) GetAnonymousPost() bool {
	if m != nil && m.AnonymousPost != nil {
		return *m.AnonymousPost
	}
	return Default_FeedPost_AnonymousPost
}

func (m *FeedPost) GetTimestamp() *Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *FeedPost) GetTimespanActive() *Timespan {
	if m != nil {
		return m.TimespanActive
	}
	return nil
}

func (m *FeedPost) GetHeadlineText() string {
	if m != nil && m.HeadlineText != nil {
		return *m.HeadlineText
	}
	return ""
}

func (m *FeedPost) GetBodyText() string {
	if m != nil && m.BodyText != nil {
		return *m.BodyText
	}
	return ""
}

func (m *FeedPost) GetPostTags() []*EntityTag {
	if m != nil {
		return m.PostTags
	}
	return nil
}

func (m *FeedPost) GetRepliesDeprecated() []*FeedPost {
	if m != nil {
		return m.RepliesDeprecated
	}
	return nil
}

func (m *FeedPost) GetMayAddReply() bool {
	if m != nil && m.MayAddReply != nil {
		return *m.MayAddReply
	}
	return false
}

func (m *FeedPost) GetMayChooseMulitpleReplies() bool {
	if m != nil && m.MayChooseMulitpleReplies != nil {
		return *m.MayChooseMulitpleReplies
	}
	return false
}

func (m *FeedPost) GetSurveyAnswerSequence() int32 {
	if m != nil && m.SurveyAnswerSequence != nil {
		return *m.SurveyAnswerSequence
	}
	return 0
}

func (m *FeedPost) GetAreMoreReplies() bool {
	if m != nil && m.AreMoreReplies != nil {
		return *m.AreMoreReplies
	}
	return false
}

func (m *FeedPost) GetTotalVoteCount() int32 {
	if m != nil && m.TotalVoteCount != nil {
		return *m.TotalVoteCount
	}
	return 0
}

type FeedPostUpdateRequest struct {
	UpdateVerb         *UpdateVerb `protobuf:"varint,1,opt,name=updateVerb,enum=BlitzMessage.UpdateVerb" json:"updateVerb,omitempty"`
	FeedPostDeprecated *FeedPost   `protobuf:"bytes,2,opt,name=feedPost_deprecated" json:"feedPost_deprecated,omitempty"`
	FeedPosts          []*FeedPost `protobuf:"bytes,3,rep,name=feedPosts" json:"feedPosts,omitempty"`
	XXX_unrecognized   []byte      `json:"-"`
}

func (m *FeedPostUpdateRequest) Reset()                    { *m = FeedPostUpdateRequest{} }
func (m *FeedPostUpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*FeedPostUpdateRequest) ProtoMessage()               {}
func (*FeedPostUpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *FeedPostUpdateRequest) GetUpdateVerb() UpdateVerb {
	if m != nil && m.UpdateVerb != nil {
		return *m.UpdateVerb
	}
	return UpdateVerb_UVCreate
}

func (m *FeedPostUpdateRequest) GetFeedPostDeprecated() *FeedPost {
	if m != nil {
		return m.FeedPostDeprecated
	}
	return nil
}

func (m *FeedPostUpdateRequest) GetFeedPosts() []*FeedPost {
	if m != nil {
		return m.FeedPosts
	}
	return nil
}

type FeedPostUpdateResponse struct {
	FeedPost         *FeedPost `protobuf:"bytes,1,opt,name=feedPost" json:"feedPost,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *FeedPostUpdateResponse) Reset()                    { *m = FeedPostUpdateResponse{} }
func (m *FeedPostUpdateResponse) String() string            { return proto.CompactTextString(m) }
func (*FeedPostUpdateResponse) ProtoMessage()               {}
func (*FeedPostUpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *FeedPostUpdateResponse) GetFeedPost() *FeedPost {
	if m != nil {
		return m.FeedPost
	}
	return nil
}

type FeedPostFetchRequest struct {
	Timespan         *Timespan      `protobuf:"bytes,1,opt,name=timespan" json:"timespan,omitempty"`
	FeedScope        *FeedPostScope `protobuf:"varint,2,opt,name=feedScope,enum=BlitzMessage.FeedPostScope" json:"feedScope,omitempty"`
	ParentID         *string        `protobuf:"bytes,3,opt,name=parentID" json:"parentID,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *FeedPostFetchRequest) Reset()                    { *m = FeedPostFetchRequest{} }
func (m *FeedPostFetchRequest) String() string            { return proto.CompactTextString(m) }
func (*FeedPostFetchRequest) ProtoMessage()               {}
func (*FeedPostFetchRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *FeedPostFetchRequest) GetTimespan() *Timespan {
	if m != nil {
		return m.Timespan
	}
	return nil
}

func (m *FeedPostFetchRequest) GetFeedScope() FeedPostScope {
	if m != nil && m.FeedScope != nil {
		return *m.FeedScope
	}
	return FeedPostScope_FPScopeUnknown
}

func (m *FeedPostFetchRequest) GetParentID() string {
	if m != nil && m.ParentID != nil {
		return *m.ParentID
	}
	return ""
}

type FeedPostFetchResponse struct {
	FeedPosts        []*FeedPost `protobuf:"bytes,1,rep,name=feedPosts" json:"feedPosts,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *FeedPostFetchResponse) Reset()                    { *m = FeedPostFetchResponse{} }
func (m *FeedPostFetchResponse) String() string            { return proto.CompactTextString(m) }
func (*FeedPostFetchResponse) ProtoMessage()               {}
func (*FeedPostFetchResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *FeedPostFetchResponse) GetFeedPosts() []*FeedPost {
	if m != nil {
		return m.FeedPosts
	}
	return nil
}

func init() {
	proto.RegisterType((*FeedPost)(nil), "BlitzMessage.FeedPost")
	proto.RegisterType((*FeedPostUpdateRequest)(nil), "BlitzMessage.FeedPostUpdateRequest")
	proto.RegisterType((*FeedPostUpdateResponse)(nil), "BlitzMessage.FeedPostUpdateResponse")
	proto.RegisterType((*FeedPostFetchRequest)(nil), "BlitzMessage.FeedPostFetchRequest")
	proto.RegisterType((*FeedPostFetchResponse)(nil), "BlitzMessage.FeedPostFetchResponse")
	proto.RegisterEnum("BlitzMessage.FeedPostType", FeedPostType_name, FeedPostType_value)
	proto.RegisterEnum("BlitzMessage.FeedPostScope", FeedPostScope_name, FeedPostScope_value)
	proto.RegisterEnum("BlitzMessage.FeedPostStatus", FeedPostStatus_name, FeedPostStatus_value)
	proto.RegisterEnum("BlitzMessage.UpdateVerb", UpdateVerb_name, UpdateVerb_value)
}

var fileDescriptor2 = []byte{
	// 713 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x54, 0x4d, 0x4f, 0xdb, 0x4a,
	0x14, 0xc5, 0x09, 0xf0, 0x92, 0x4b, 0x92, 0xe7, 0x37, 0x40, 0x18, 0x05, 0x16, 0x88, 0x15, 0x44,
	0xbc, 0x54, 0x4a, 0xa5, 0x2e, 0xba, 0xa0, 0x22, 0x40, 0xaa, 0x4a, 0xa5, 0xa5, 0xe4, 0x63, 0x5b,
	0x39, 0xf6, 0x85, 0xb8, 0x38, 0x33, 0xae, 0x67, 0x0c, 0x4d, 0x97, 0x5d, 0xf5, 0x67, 0x74, 0xd3,
	0x3f, 0xc4, 0x2f, 0xea, 0xf5, 0x38, 0x0e, 0x09, 0x4d, 0xe9, 0x8a, 0xcc, 0xb9, 0xe7, 0x9e, 0xb9,
	0xf7, 0xf8, 0x0c, 0x00, 0x6d, 0x44, 0xaf, 0x11, 0x46, 0x52, 0x4b, 0x56, 0x6a, 0x05, 0xbe, 0xfe,
	0x7a, 0x8e, 0x4a, 0x39, 0xd7, 0x58, 0xdb, 0x96, 0x83, 0x4f, 0xe8, 0x6a, 0xff, 0x16, 0xdd, 0xff,
	0x3d, 0x54, 0x6e, 0xe4, 0x87, 0x5a, 0x46, 0x29, 0xb5, 0xb6, 0xd6, 0x1d, 0x87, 0xa8, 0x26, 0x07,
	0xfb, 0x4c, 0x68, 0x5f, 0x8f, 0xbb, 0xce, 0xf5, 0x04, 0xd9, 0xfb, 0xb1, 0x0c, 0x85, 0x44, 0xf8,
	0x42, 0x2a, 0xcd, 0x2a, 0xb0, 0x1a, 0xd2, 0xdf, 0x37, 0xa7, 0xdc, 0xda, 0xb5, 0xf6, 0x8b, 0xcc,
	0x86, 0x42, 0xe8, 0x44, 0x28, 0x12, 0x24, 0x67, 0x90, 0x43, 0x42, 0x88, 0x91, 0x68, 0xf2, 0x3c,
	0x21, 0x95, 0x66, 0xad, 0x31, 0x3b, 0x4b, 0x23, 0xd3, 0x4a, 0x18, 0xac, 0x01, 0xc5, 0x84, 0xdd,
	0x71, 0x25, 0xd1, 0x97, 0x0d, 0x7d, 0x7b, 0x31, 0xdd, 0x50, 0x92, 0xfb, 0x63, 0x85, 0x11, 0xdd,
	0xb6, 0x62, 0x6e, 0xdb, 0x81, 0xb2, 0x23, 0xa4, 0x18, 0x8f, 0x64, 0xac, 0x12, 0x16, 0x5f, 0x25,
	0xb8, 0xf0, 0x72, 0xe5, 0xca, 0x09, 0x14, 0xb2, 0x3a, 0x14, 0xb5, 0x3f, 0x42, 0xa5, 0x9d, 0x51,
	0xc8, 0xff, 0xa1, 0xca, 0x5a, 0x73, 0x6b, 0x5e, 0xbd, 0x9b, 0x95, 0x69, 0x92, 0x8a, 0xe1, 0x86,
	0x8e, 0x38, 0x36, 0x4e, 0xf1, 0x82, 0x69, 0xa8, 0x2e, 0x68, 0x20, 0x0e, 0xdb, 0x80, 0xd2, 0x10,
	0x1d, 0x2f, 0xf0, 0x05, 0x76, 0xf1, 0x8b, 0xe6, 0xc5, 0xcc, 0x8f, 0x81, 0xf4, 0xc6, 0x06, 0x01,
	0x83, 0x1c, 0x4c, 0xfc, 0x20, 0x43, 0x79, 0x69, 0x37, 0xff, 0xfb, 0x08, 0x53, 0xc3, 0x59, 0x13,
	0x58, 0x84, 0x61, 0xe0, 0xa3, 0xfa, 0xe8, 0x61, 0x18, 0xa1, 0xeb, 0x68, 0xf4, 0x78, 0xd9, 0x34,
	0x55, 0x17, 0xbb, 0xc2, 0xd6, 0x61, 0x6d, 0xe4, 0x8c, 0x8f, 0x3d, 0xef, 0x92, 0x3a, 0xc7, 0xbc,
	0x92, 0xac, 0xcf, 0x76, 0x81, 0x13, 0x78, 0x32, 0x94, 0x52, 0xe1, 0x79, 0x4c, 0x7d, 0x61, 0x80,
	0x97, 0xa9, 0x32, 0xff, 0xd7, 0x30, 0x76, 0x60, 0x43, 0xc5, 0xd1, 0x2d, 0x8e, 0x8f, 0x85, 0xba,
	0xc3, 0xa8, 0x83, 0x9f, 0x63, 0x14, 0x2e, 0x72, 0x9b, 0xaa, 0x2b, 0xac, 0x0a, 0x15, 0xfa, 0xa8,
	0xe7, 0x32, 0x9a, 0x76, 0xfd, 0x67, 0xba, 0x08, 0xd7, 0x52, 0x3b, 0x41, 0x5f, 0x6a, 0x3c, 0x91,
	0xb1, 0xd0, 0x9c, 0x25, 0xfc, 0xbd, 0x9f, 0x16, 0x6c, 0x66, 0x13, 0xf5, 0x42, 0x8f, 0xc6, 0xbe,
	0x4c, 0x04, 0x69, 0xbc, 0x43, 0x80, 0xd8, 0x00, 0x7d, 0x8c, 0x06, 0x26, 0x33, 0x95, 0x26, 0x9f,
	0x5f, 0xa5, 0x37, 0xad, 0xb3, 0xe7, 0xb0, 0x7e, 0x35, 0x91, 0x99, 0x75, 0x20, 0xb7, 0xe8, 0x43,
	0x4c, 0x1d, 0x38, 0x80, 0x62, 0xd6, 0xa4, 0x28, 0x71, 0x4f, 0x98, 0xb5, 0xd7, 0x82, 0xea, 0xe3,
	0x31, 0x55, 0x28, 0x05, 0x25, 0x65, 0x1f, 0x0a, 0x99, 0x88, 0x99, 0xf2, 0xcf, 0x1a, 0xdf, 0x2c,
	0xd8, 0xc8, 0x0e, 0x6d, 0xd4, 0xee, 0x30, 0x5b, 0x95, 0x24, 0xb2, 0x00, 0x2d, 0x96, 0x98, 0x46,
	0xa7, 0x91, 0x4e, 0x9c, 0x86, 0x3e, 0xf7, 0xf7, 0xd0, 0xcf, 0x3e, 0xb2, 0xe4, 0x49, 0x15, 0x69,
	0x91, 0xcd, 0x47, 0x33, 0x4c, 0xf6, 0x98, 0x33, 0xc3, 0x7a, 0xca, 0x8c, 0xfa, 0x2d, 0x94, 0xe6,
	0x9e, 0x62, 0x19, 0x8a, 0xed, 0x8b, 0x9e, 0xb8, 0x11, 0xf2, 0x4e, 0xd8, 0x4b, 0x6c, 0x0b, 0xd6,
	0xdb, 0x17, 0xef, 0x43, 0x14, 0x67, 0xc2, 0x43, 0xef, 0x43, 0xb2, 0xa2, 0x2f, 0x85, 0x6d, 0x51,
	0xf0, 0xed, 0x99, 0x82, 0x89, 0x9d, 0x9d, 0x4b, 0xd1, 0x8e, 0x89, 0xd4, 0x94, 0x9b, 0x67, 0x0c,
	0x2a, 0x19, 0x9a, 0x06, 0xcd, 0x5e, 0xae, 0xf7, 0xa1, 0x3c, 0xbf, 0x5e, 0x4a, 0x4a, 0x7e, 0x3e,
	0xba, 0xdd, 0x60, 0x6f, 0xa5, 0xeb, 0x04, 0xef, 0x50, 0xdf, 0xc9, 0xe8, 0x86, 0x6e, 0xe7, 0xe4,
	0x7e, 0x5a, 0x78, 0x1d, 0xc8, 0xc1, 0x43, 0x25, 0x57, 0x7f, 0x45, 0x32, 0x99, 0xae, 0x76, 0x74,
	0xac, 0xe8, 0x9f, 0x05, 0x10, 0xf7, 0x41, 0xd4, 0x6c, 0xd8, 0x49, 0x5f, 0x37, 0x49, 0xa5, 0xe5,
	0x53, 0x0c, 0x90, 0x42, 0x46, 0x02, 0x2f, 0x00, 0x66, 0xb2, 0x58, 0x82, 0x42, 0xaf, 0x7f, 0x12,
	0x21, 0x9d, 0x89, 0x6b, 0x4e, 0x69, 0x95, 0x96, 0x35, 0xa7, 0xb4, 0xd1, 0xce, 0xb7, 0x9e, 0xdd,
	0x1f, 0xe5, 0x60, 0xe9, 0xfe, 0x28, 0xcf, 0xac, 0x16, 0xfd, 0xe4, 0x16, 0xd4, 0x5c, 0x39, 0x6a,
	0x0c, 0x12, 0xd7, 0x87, 0x18, 0xe1, 0x9c, 0xff, 0xdf, 0x2d, 0xeb, 0x57, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x9e, 0x4a, 0xc0, 0x82, 0xa8, 0x05, 0x00, 0x00,
}
