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

type AutocompleteRequest struct {
	Query            *string `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
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

func init() {
	proto.RegisterType((*AutocompleteRequest)(nil), "BlitzMessage.AutocompleteRequest")
	proto.RegisterType((*AutocompleteResponse)(nil), "BlitzMessage.AutocompleteResponse")
	proto.RegisterType((*UserSearchRequest)(nil), "BlitzMessage.UserSearchRequest")
	proto.RegisterType((*UserSearchResponse)(nil), "BlitzMessage.UserSearchResponse")
}

var fileDescriptor5 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x09, 0x4e, 0x4d, 0x2c,
	0x4a, 0xce, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x71, 0xca, 0xc9, 0x2c, 0xa9, 0xf2,
	0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x95, 0x92, 0xce, 0x4f, 0xca, 0x4a, 0x4d, 0x2e, 0xc9, 0x2c,
	0x4b, 0x4d, 0xd6, 0x4d, 0x49, 0x2d, 0x4e, 0x2e, 0xca, 0x2c, 0x28, 0xc9, 0x2f, 0x82, 0x28, 0x95,
	0x12, 0x0a, 0x2d, 0x4e, 0x2d, 0x0a, 0x28, 0xca, 0x4f, 0xcb, 0xcc, 0x49, 0x2d, 0x86, 0x88, 0x29,
	0xa9, 0x70, 0x09, 0x3b, 0x96, 0x96, 0xe4, 0x27, 0xe7, 0xe7, 0x16, 0xe4, 0xa4, 0x96, 0xa4, 0x06,
	0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0xf1, 0x72, 0xb1, 0x02, 0x19, 0x45, 0x95, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0x9c, 0x4a, 0x56, 0x5c, 0x22, 0xa8, 0xaa, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53,
	0xd1, 0x94, 0x09, 0x09, 0x73, 0x71, 0x17, 0x97, 0xa6, 0xa7, 0x03, 0x0d, 0xc8, 0x04, 0x4a, 0x4b,
	0x30, 0x29, 0x30, 0x03, 0xf5, 0x2a, 0x71, 0x09, 0x82, 0xec, 0x85, 0x38, 0x1a, 0x87, 0xf9, 0x01,
	0x5c, 0x42, 0xc8, 0x6a, 0xb0, 0x9b, 0xae, 0xcd, 0xc5, 0x51, 0x00, 0x75, 0x3c, 0xd8, 0x68, 0x6e,
	0x23, 0x49, 0x3d, 0x64, 0xcf, 0xeb, 0x21, 0x79, 0xcf, 0x49, 0xff, 0x92, 0x1d, 0x13, 0x17, 0xc3,
	0x25, 0x3b, 0x66, 0x21, 0x46, 0x27, 0x20, 0x53, 0x82, 0x91, 0x4b, 0x0a, 0xe8, 0x7a, 0xbd, 0x24,
	0x90, 0xfa, 0x8c, 0xd4, 0xa2, 0x54, 0x14, 0x9d, 0x1d, 0x8c, 0x8c, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xcf, 0x96, 0x43, 0x14, 0x56, 0x01, 0x00, 0x00,
}
