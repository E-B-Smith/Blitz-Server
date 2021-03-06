// Code generated by protoc-gen-go.
// source: Search.proto
// DO NOT EDIT!

package BlitzMessage

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"


// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SearchType int32

const (
	SearchType_STSearchAll SearchType = 0
	SearchType_STUsers     SearchType = 1
	SearchType_STTopics    SearchType = 2
)

var SearchType_name = map[int32]string{
	0: "STSearchAll",
	1: "STUsers",
	2: "STTopics",
}
var SearchType_value = map[string]int32{
	"STSearchAll": 0,
	"STUsers":     1,
	"STTopics":    2,
}

func (x SearchType) Enum() *SearchType {
	p := new(SearchType)
	*p = x
	return p
}
func (x SearchType) String() string {
	return proto.EnumName(SearchType_name, int32(x))
}
func (x *SearchType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SearchType_value, data, "SearchType")
	if err != nil {
		return err
	}
	*x = SearchType(value)
	return nil
}
func (SearchType) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

type AutocompleteRequest struct {
	Query            *string     `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
	SearchType       *SearchType `protobuf:"varint,2,opt,name=searchType,enum=BlitzMessage.SearchType" json:"searchType,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *AutocompleteRequest) Reset()                    { *m = AutocompleteRequest{} }
func (m *AutocompleteRequest) String() string            { return proto.CompactTextString(m) }
func (*AutocompleteRequest) ProtoMessage()               {}
func (*AutocompleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *AutocompleteRequest) GetQuery() string {
	if m != nil && m.Query != nil {
		return *m.Query
	}
	return ""
}

func (m *AutocompleteRequest) GetSearchType() SearchType {
	if m != nil && m.SearchType != nil {
		return *m.SearchType
	}
	return SearchType_STSearchAll
}

type AutocompleteResponse struct {
	Query            *string  `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
	Suggestions      []string `protobuf:"bytes,2,rep,name=suggestions" json:"suggestions,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *AutocompleteResponse) Reset()                    { *m = AutocompleteResponse{} }
func (m *AutocompleteResponse) String() string            { return proto.CompactTextString(m) }
func (*AutocompleteResponse) ProtoMessage()               {}
func (*AutocompleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *AutocompleteResponse) GetQuery() string {
	if m != nil && m.Query != nil {
		return *m.Query
	}
	return ""
}

func (m *AutocompleteResponse) GetSuggestions() []string {
	if m != nil {
		return m.Suggestions
	}
	return nil
}

type UserSearchRequest struct {
	Query            *string `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UserSearchRequest) Reset()                    { *m = UserSearchRequest{} }
func (m *UserSearchRequest) String() string            { return proto.CompactTextString(m) }
func (*UserSearchRequest) ProtoMessage()               {}
func (*UserSearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *UserSearchRequest) GetQuery() string {
	if m != nil && m.Query != nil {
		return *m.Query
	}
	return ""
}

type UserSearchResponse struct {
	Query            *string        `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
	Profiles         []*UserProfile `protobuf:"bytes,2,rep,name=profiles" json:"profiles,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *UserSearchResponse) Reset()                    { *m = UserSearchResponse{} }
func (m *UserSearchResponse) String() string            { return proto.CompactTextString(m) }
func (*UserSearchResponse) ProtoMessage()               {}
func (*UserSearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *UserSearchResponse) GetQuery() string {
	if m != nil && m.Query != nil {
		return *m.Query
	}
	return ""
}

func (m *UserSearchResponse) GetProfiles() []*UserProfile {
	if m != nil {
		return m.Profiles
	}
	return nil
}

type SearchCategory struct {
	Item             *string `protobuf:"bytes,1,opt,name=item" json:"item,omitempty"`
	Parent           *string `protobuf:"bytes,2,opt,name=parent" json:"parent,omitempty"`
	IsLeaf           *bool   `protobuf:"varint,3,opt,name=isLeaf" json:"isLeaf,omitempty"`
	DescriptionText  *string `protobuf:"bytes,4,opt,name=descriptionText" json:"descriptionText,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SearchCategory) Reset()                    { *m = SearchCategory{} }
func (m *SearchCategory) String() string            { return proto.CompactTextString(m) }
func (*SearchCategory) ProtoMessage()               {}
func (*SearchCategory) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *SearchCategory) GetItem() string {
	if m != nil && m.Item != nil {
		return *m.Item
	}
	return ""
}

func (m *SearchCategory) GetParent() string {
	if m != nil && m.Parent != nil {
		return *m.Parent
	}
	return ""
}

func (m *SearchCategory) GetIsLeaf() bool {
	if m != nil && m.IsLeaf != nil {
		return *m.IsLeaf
	}
	return false
}

func (m *SearchCategory) GetDescriptionText() string {
	if m != nil && m.DescriptionText != nil {
		return *m.DescriptionText
	}
	return ""
}

type SearchCategories struct {
	Categories       []*SearchCategory `protobuf:"bytes,1,rep,name=categories" json:"categories,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *SearchCategories) Reset()                    { *m = SearchCategories{} }
func (m *SearchCategories) String() string            { return proto.CompactTextString(m) }
func (*SearchCategories) ProtoMessage()               {}
func (*SearchCategories) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func (m *SearchCategories) GetCategories() []*SearchCategory {
	if m != nil {
		return m.Categories
	}
	return nil
}

func init() {
	proto.RegisterType((*AutocompleteRequest)(nil), "BlitzMessage.AutocompleteRequest")
	proto.RegisterType((*AutocompleteResponse)(nil), "BlitzMessage.AutocompleteResponse")
	proto.RegisterType((*UserSearchRequest)(nil), "BlitzMessage.UserSearchRequest")
	proto.RegisterType((*UserSearchResponse)(nil), "BlitzMessage.UserSearchResponse")
	proto.RegisterType((*SearchCategory)(nil), "BlitzMessage.SearchCategory")
	proto.RegisterType((*SearchCategories)(nil), "BlitzMessage.SearchCategories")
	proto.RegisterEnum("BlitzMessage.SearchType", SearchType_name, SearchType_value)
}

func init() { proto.RegisterFile("Search.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0xd7, 0xe9, 0x02, 0xdd, 0x49, 0xe8, 0x06, 0x2f, 0x12, 0xa6, 0x70, 0x88, 0x72, 0x8a,
	0xf8, 0x13, 0xd0, 0x9e, 0x10, 0x87, 0x4a, 0x5b, 0x38, 0x82, 0xb4, 0xda, 0x04, 0x71, 0x4e, 0xcd,
	0x34, 0x35, 0x4a, 0x63, 0xd7, 0xe3, 0x20, 0xca, 0x13, 0xf0, 0x7c, 0x7d, 0x22, 0xd4, 0xa4, 0x2d,
	0xcd, 0xaa, 0xc7, 0xb1, 0x47, 0xbf, 0xef, 0x37, 0x1f, 0x04, 0x19, 0x16, 0x56, 0x2e, 0x52, 0x63,
	0xb5, 0xd3, 0x3c, 0x98, 0x56, 0xca, 0xfd, 0xf9, 0x8a, 0x44, 0x45, 0x89, 0xe3, 0x17, 0x7a, 0xf6,
	0x13, 0xa5, 0x53, 0xbf, 0x50, 0xbe, 0xfd, 0x81, 0x24, 0xad, 0x32, 0x4e, 0xdb, 0x6e, 0x75, 0xcc,
	0xbf, 0x11, 0xda, 0x5b, 0xab, 0xe7, 0xaa, 0x42, 0xea, 0xde, 0xe2, 0x3b, 0xb8, 0xba, 0x69, 0x9c,
	0x96, 0x7a, 0x69, 0x2a, 0x74, 0x78, 0x87, 0xab, 0x06, 0xc9, 0xf1, 0xc7, 0xf0, 0x60, 0xd5, 0xa0,
	0x5d, 0x0b, 0x16, 0xb1, 0xe4, 0x82, 0xbf, 0x01, 0xa0, 0x36, 0x34, 0x5f, 0x1b, 0x14, 0x5e, 0xc4,
	0x92, 0xd1, 0xb5, 0x48, 0x8f, 0x93, 0xd3, 0xec, 0xf0, 0x1f, 0x7f, 0x84, 0xa7, 0x7d, 0x26, 0x19,
	0x5d, 0x13, 0xde, 0x87, 0x5e, 0x81, 0x4f, 0x4d, 0x59, 0x22, 0x39, 0xa5, 0x6b, 0x12, 0x5e, 0x34,
	0x48, 0x2e, 0xe2, 0x18, 0x9e, 0x6c, 0x2d, 0x3b, 0xda, 0x69, 0x9b, 0xf8, 0x16, 0xf8, 0xf1, 0xce,
	0x69, 0xfa, 0x6b, 0x18, 0x9a, 0xdd, 0xa9, 0x2d, 0xda, 0xbf, 0x7e, 0xde, 0x17, 0x3e, 0x2a, 0x23,
	0xfe, 0x0e, 0xa3, 0x8e, 0xf6, 0xa9, 0x70, 0x58, 0x6a, 0xbb, 0xe6, 0x01, 0x9c, 0x2b, 0x87, 0xcb,
	0x1d, 0x6c, 0x04, 0x0f, 0x4d, 0x61, 0xb1, 0x76, 0xed, 0xed, 0xed, 0xac, 0xe8, 0x0b, 0x16, 0x73,
	0x31, 0x88, 0x58, 0x32, 0xe4, 0xcf, 0xe0, 0x72, 0xdf, 0xb6, 0xd2, 0x75, 0x8e, 0xbf, 0x9d, 0x38,
	0x6f, 0x55, 0x3f, 0x43, 0xd8, 0x03, 0x2b, 0x24, 0xfe, 0x1e, 0x40, 0x1e, 0x26, 0xc1, 0x5a, 0xb7,
	0x97, 0xa7, 0xca, 0xdc, 0xcb, 0xbc, 0xfa, 0x00, 0xf0, 0xbf, 0x5e, 0x7e, 0x09, 0x7e, 0x96, 0x77,
	0xf3, 0x4d, 0x55, 0x85, 0x67, 0xdc, 0x87, 0x47, 0x59, 0xbe, 0x3d, 0x87, 0x42, 0xc6, 0x03, 0x18,
	0x66, 0x79, 0xae, 0x8d, 0x92, 0x14, 0x7a, 0xd3, 0x77, 0x9b, 0x89, 0x07, 0x67, 0x9b, 0xc9, 0x80,
	0xb3, 0xe9, 0x66, 0xe2, 0x09, 0x06, 0x63, 0xa9, 0x97, 0xe9, 0x6c, 0x1b, 0xb6, 0x40, 0x8b, 0xbd,
	0xd8, 0xbf, 0x8c, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x12, 0xff, 0xb6, 0xa2, 0x5d, 0x02, 0x00,
	0x00,
}
