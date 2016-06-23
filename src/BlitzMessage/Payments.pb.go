// Code generated by protoc-gen-go.
// source: Payments.proto
// DO NOT EDIT!

package BlitzMessage

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"


// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CardStatus int32

const (
	CardStatus_CSUnknown  CardStatus = 0
	CardStatus_CSStandard CardStatus = 1
	CardStatus_CSDeleted  CardStatus = 2
)

var CardStatus_name = map[int32]string{
	0: "CSUnknown",
	1: "CSStandard",
	2: "CSDeleted",
}
var CardStatus_value = map[string]int32{
	"CSUnknown":  0,
	"CSStandard": 1,
	"CSDeleted":  2,
}

func (x CardStatus) Enum() *CardStatus {
	p := new(CardStatus)
	*p = x
	return p
}
func (x CardStatus) String() string {
	return proto.EnumName(CardStatus_name, int32(x))
}
func (x *CardStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CardStatus_value, data, "CardStatus")
	if err != nil {
		return err
	}
	*x = CardStatus(value)
	return nil
}
func (CardStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

type ChargeStatus int32

const (
	ChargeStatus_CSChargeUnknown ChargeStatus = 0
	ChargeStatus_CSChargeRequest ChargeStatus = 1
	ChargeStatus_CSDeclined      ChargeStatus = 2
	ChargeStatus_CSPreauthorized ChargeStatus = 3
	ChargeStatus_CSCharged       ChargeStatus = 4
)

var ChargeStatus_name = map[int32]string{
	0: "CSChargeUnknown",
	1: "CSChargeRequest",
	2: "CSDeclined",
	3: "CSPreauthorized",
	4: "CSCharged",
}
var ChargeStatus_value = map[string]int32{
	"CSChargeUnknown": 0,
	"CSChargeRequest": 1,
	"CSDeclined":      2,
	"CSPreauthorized": 3,
	"CSCharged":       4,
}

func (x ChargeStatus) Enum() *ChargeStatus {
	p := new(ChargeStatus)
	*p = x
	return p
}
func (x ChargeStatus) String() string {
	return proto.EnumName(ChargeStatus_name, int32(x))
}
func (x *ChargeStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ChargeStatus_value, data, "ChargeStatus")
	if err != nil {
		return err
	}
	*x = ChargeStatus(value)
	return nil
}
func (ChargeStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

type ChargeTokenType int32

const (
	ChargeTokenType_CTTUnkown      ChargeTokenType = 0
	ChargeTokenType_CTTApplePay    ChargeTokenType = 1
	ChargeTokenType_CTTStripeToken ChargeTokenType = 2
	ChargeTokenType_CTTCardToken   ChargeTokenType = 3
)

var ChargeTokenType_name = map[int32]string{
	0: "CTTUnkown",
	1: "CTTApplePay",
	2: "CTTStripeToken",
	3: "CTTCardToken",
}
var ChargeTokenType_value = map[string]int32{
	"CTTUnkown":      0,
	"CTTApplePay":    1,
	"CTTStripeToken": 2,
	"CTTCardToken":   3,
}

func (x ChargeTokenType) Enum() *ChargeTokenType {
	p := new(ChargeTokenType)
	*p = x
	return p
}
func (x ChargeTokenType) String() string {
	return proto.EnumName(ChargeTokenType_name, int32(x))
}
func (x *ChargeTokenType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ChargeTokenType_value, data, "ChargeTokenType")
	if err != nil {
		return err
	}
	*x = ChargeTokenType(value)
	return nil
}
func (ChargeTokenType) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

type PurchaseType int32

const (
	PurchaseType_PurchaseTypeUnknown PurchaseType = 0
	PurchaseType_PTChatConversation  PurchaseType = 1
	PurchaseType_PTFeedPost          PurchaseType = 2
)

var PurchaseType_name = map[int32]string{
	0: "PurchaseTypeUnknown",
	1: "PTChatConversation",
	2: "PTFeedPost",
}
var PurchaseType_value = map[string]int32{
	"PurchaseTypeUnknown": 0,
	"PTChatConversation":  1,
	"PTFeedPost":          2,
}

func (x PurchaseType) Enum() *PurchaseType {
	p := new(PurchaseType)
	*p = x
	return p
}
func (x PurchaseType) String() string {
	return proto.EnumName(PurchaseType_name, int32(x))
}
func (x *PurchaseType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PurchaseType_value, data, "PurchaseType")
	if err != nil {
		return err
	}
	*x = PurchaseType(value)
	return nil
}
func (PurchaseType) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

type CardInfo struct {
	CardStatus       *CardStatus `protobuf:"varint,1,opt,name=cardStatus,enum=BlitzMessage.CardStatus" json:"cardStatus,omitempty"`
	CardHolderName   *string     `protobuf:"bytes,2,opt,name=cardHolderName" json:"cardHolderName,omitempty"`
	MemoText         *string     `protobuf:"bytes,3,opt,name=memoText" json:"memoText,omitempty"`
	Brand            *string     `protobuf:"bytes,4,opt,name=brand" json:"brand,omitempty"`
	Last4            *string     `protobuf:"bytes,5,opt,name=last4" json:"last4,omitempty"`
	ExpireMonth      *int32      `protobuf:"varint,6,opt,name=expireMonth" json:"expireMonth,omitempty"`
	ExpireYear       *int32      `protobuf:"varint,7,opt,name=expireYear" json:"expireYear,omitempty"`
	Token            *string     `protobuf:"bytes,8,opt,name=token" json:"token,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *CardInfo) Reset()                    { *m = CardInfo{} }
func (m *CardInfo) String() string            { return proto.CompactTextString(m) }
func (*CardInfo) ProtoMessage()               {}
func (*CardInfo) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *CardInfo) GetCardStatus() CardStatus {
	if m != nil && m.CardStatus != nil {
		return *m.CardStatus
	}
	return CardStatus_CSUnknown
}

func (m *CardInfo) GetCardHolderName() string {
	if m != nil && m.CardHolderName != nil {
		return *m.CardHolderName
	}
	return ""
}

func (m *CardInfo) GetMemoText() string {
	if m != nil && m.MemoText != nil {
		return *m.MemoText
	}
	return ""
}

func (m *CardInfo) GetBrand() string {
	if m != nil && m.Brand != nil {
		return *m.Brand
	}
	return ""
}

func (m *CardInfo) GetLast4() string {
	if m != nil && m.Last4 != nil {
		return *m.Last4
	}
	return ""
}

func (m *CardInfo) GetExpireMonth() int32 {
	if m != nil && m.ExpireMonth != nil {
		return *m.ExpireMonth
	}
	return 0
}

func (m *CardInfo) GetExpireYear() int32 {
	if m != nil && m.ExpireYear != nil {
		return *m.ExpireYear
	}
	return 0
}

func (m *CardInfo) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

type UserCardInfo struct {
	UserID           *string     `protobuf:"bytes,1,opt,name=userID" json:"userID,omitempty"`
	CardInfo         []*CardInfo `protobuf:"bytes,2,rep,name=cardInfo" json:"cardInfo,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *UserCardInfo) Reset()                    { *m = UserCardInfo{} }
func (m *UserCardInfo) String() string            { return proto.CompactTextString(m) }
func (*UserCardInfo) ProtoMessage()               {}
func (*UserCardInfo) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *UserCardInfo) GetUserID() string {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return ""
}

func (m *UserCardInfo) GetCardInfo() []*CardInfo {
	if m != nil {
		return m.CardInfo
	}
	return nil
}

type Charge struct {
	ChargeID          *string          `protobuf:"bytes,1,opt,name=chargeID" json:"chargeID,omitempty"`
	Timestamp         *Timestamp       `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	ChargeStatus      *ChargeStatus    `protobuf:"varint,3,opt,name=chargeStatus,enum=BlitzMessage.ChargeStatus" json:"chargeStatus,omitempty"`
	PayerID           *string          `protobuf:"bytes,4,opt,name=payerID" json:"payerID,omitempty"`
	PayeeIDDeprecated *string          `protobuf:"bytes,5,opt,name=payeeID_deprecated" json:"payeeID_deprecated,omitempty"`
	PurchaseType      *PurchaseType    `protobuf:"varint,6,opt,name=purchaseType,enum=BlitzMessage.PurchaseType" json:"purchaseType,omitempty"`
	PurchaseTypeID    *string          `protobuf:"bytes,7,opt,name=purchaseTypeID" json:"purchaseTypeID,omitempty"`
	MemoText          *string          `protobuf:"bytes,8,opt,name=memoText" json:"memoText,omitempty"`
	Amount            *string          `protobuf:"bytes,9,opt,name=amount" json:"amount,omitempty"`
	Currency          *string          `protobuf:"bytes,10,opt,name=currency" json:"currency,omitempty"`
	TokenType         *ChargeTokenType `protobuf:"varint,11,opt,name=tokenType,enum=BlitzMessage.ChargeTokenType" json:"tokenType,omitempty"`
	ChargeToken       *string          `protobuf:"bytes,12,opt,name=chargeToken" json:"chargeToken,omitempty"`
	ProcessorReason   *string          `protobuf:"bytes,13,opt,name=processorReason" json:"processorReason,omitempty"`
	XXX_unrecognized  []byte           `json:"-"`
}

func (m *Charge) Reset()                    { *m = Charge{} }
func (m *Charge) String() string            { return proto.CompactTextString(m) }
func (*Charge) ProtoMessage()               {}
func (*Charge) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *Charge) GetChargeID() string {
	if m != nil && m.ChargeID != nil {
		return *m.ChargeID
	}
	return ""
}

func (m *Charge) GetTimestamp() *Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Charge) GetChargeStatus() ChargeStatus {
	if m != nil && m.ChargeStatus != nil {
		return *m.ChargeStatus
	}
	return ChargeStatus_CSChargeUnknown
}

func (m *Charge) GetPayerID() string {
	if m != nil && m.PayerID != nil {
		return *m.PayerID
	}
	return ""
}

func (m *Charge) GetPayeeIDDeprecated() string {
	if m != nil && m.PayeeIDDeprecated != nil {
		return *m.PayeeIDDeprecated
	}
	return ""
}

func (m *Charge) GetPurchaseType() PurchaseType {
	if m != nil && m.PurchaseType != nil {
		return *m.PurchaseType
	}
	return PurchaseType_PurchaseTypeUnknown
}

func (m *Charge) GetPurchaseTypeID() string {
	if m != nil && m.PurchaseTypeID != nil {
		return *m.PurchaseTypeID
	}
	return ""
}

func (m *Charge) GetMemoText() string {
	if m != nil && m.MemoText != nil {
		return *m.MemoText
	}
	return ""
}

func (m *Charge) GetAmount() string {
	if m != nil && m.Amount != nil {
		return *m.Amount
	}
	return ""
}

func (m *Charge) GetCurrency() string {
	if m != nil && m.Currency != nil {
		return *m.Currency
	}
	return ""
}

func (m *Charge) GetTokenType() ChargeTokenType {
	if m != nil && m.TokenType != nil {
		return *m.TokenType
	}
	return ChargeTokenType_CTTUnkown
}

func (m *Charge) GetChargeToken() string {
	if m != nil && m.ChargeToken != nil {
		return *m.ChargeToken
	}
	return ""
}

func (m *Charge) GetProcessorReason() string {
	if m != nil && m.ProcessorReason != nil {
		return *m.ProcessorReason
	}
	return ""
}

type PurchaseDescription struct {
	//  Sent from client for purchase request:
	PurchaseType   *PurchaseType `protobuf:"varint,1,opt,name=purchaseType,enum=BlitzMessage.PurchaseType" json:"purchaseType,omitempty"`
	PurchaseTypeID *string       `protobuf:"bytes,2,opt,name=purchaseTypeID" json:"purchaseTypeID,omitempty"`
	//  Sent to client as purchase request response:
	PurchaseIDWha    *string `protobuf:"bytes,3,opt,name=purchaseID_wha" json:"purchaseID_wha,omitempty"`
	MemoText         *string `protobuf:"bytes,4,opt,name=memoText" json:"memoText,omitempty"`
	Amount           *string `protobuf:"bytes,5,opt,name=amount" json:"amount,omitempty"`
	Currency         *string `protobuf:"bytes,6,opt,name=currency" json:"currency,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PurchaseDescription) Reset()                    { *m = PurchaseDescription{} }
func (m *PurchaseDescription) String() string            { return proto.CompactTextString(m) }
func (*PurchaseDescription) ProtoMessage()               {}
func (*PurchaseDescription) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *PurchaseDescription) GetPurchaseType() PurchaseType {
	if m != nil && m.PurchaseType != nil {
		return *m.PurchaseType
	}
	return PurchaseType_PurchaseTypeUnknown
}

func (m *PurchaseDescription) GetPurchaseTypeID() string {
	if m != nil && m.PurchaseTypeID != nil {
		return *m.PurchaseTypeID
	}
	return ""
}

func (m *PurchaseDescription) GetPurchaseIDWha() string {
	if m != nil && m.PurchaseIDWha != nil {
		return *m.PurchaseIDWha
	}
	return ""
}

func (m *PurchaseDescription) GetMemoText() string {
	if m != nil && m.MemoText != nil {
		return *m.MemoText
	}
	return ""
}

func (m *PurchaseDescription) GetAmount() string {
	if m != nil && m.Amount != nil {
		return *m.Amount
	}
	return ""
}

func (m *PurchaseDescription) GetCurrency() string {
	if m != nil && m.Currency != nil {
		return *m.Currency
	}
	return ""
}

func init() {
	proto.RegisterType((*CardInfo)(nil), "BlitzMessage.CardInfo")
	proto.RegisterType((*UserCardInfo)(nil), "BlitzMessage.UserCardInfo")
	proto.RegisterType((*Charge)(nil), "BlitzMessage.Charge")
	proto.RegisterType((*PurchaseDescription)(nil), "BlitzMessage.PurchaseDescription")
	proto.RegisterEnum("BlitzMessage.CardStatus", CardStatus_name, CardStatus_value)
	proto.RegisterEnum("BlitzMessage.ChargeStatus", ChargeStatus_name, ChargeStatus_value)
	proto.RegisterEnum("BlitzMessage.ChargeTokenType", ChargeTokenType_name, ChargeTokenType_value)
	proto.RegisterEnum("BlitzMessage.PurchaseType", PurchaseType_name, PurchaseType_value)
}

func init() { proto.RegisterFile("Payments.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 667 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x4e, 0xdb, 0x4a,
	0x14, 0xc6, 0x09, 0x81, 0xe4, 0x24, 0x24, 0x91, 0x91, 0x60, 0x94, 0xab, 0x2b, 0x5d, 0xb1, 0x42,
	0xd1, 0x6d, 0x5a, 0xa1, 0xae, 0xba, 0x40, 0x2a, 0x89, 0x2a, 0x58, 0xd0, 0x46, 0xc4, 0xa8, 0xea,
	0xaa, 0x1a, 0xec, 0x53, 0xe2, 0x12, 0xcf, 0xb8, 0xe3, 0x31, 0x10, 0xb6, 0xdd, 0xf4, 0x5d, 0xfa,
	0x02, 0x7d, 0x0e, 0x9e, 0xa8, 0x67, 0x66, 0x1c, 0xe2, 0x90, 0xae, 0xba, 0xf3, 0x7c, 0xe7, 0xe7,
	0x3b, 0xe7, 0xfb, 0x4e, 0x02, 0xed, 0x31, 0x9f, 0x27, 0x28, 0x74, 0x36, 0x48, 0x95, 0xd4, 0xd2,
	0x6f, 0x9d, 0xcc, 0x62, 0xfd, 0x70, 0x8e, 0x59, 0xc6, 0xaf, 0xb1, 0xf7, 0x8f, 0xbc, 0xfa, 0x8a,
	0xa1, 0x8e, 0x6f, 0x31, 0x7c, 0x11, 0x61, 0x16, 0xaa, 0x38, 0xd5, 0x52, 0xb9, 0xd4, 0x5e, 0x33,
	0x98, 0xa7, 0x58, 0xd4, 0x1d, 0xfc, 0xf2, 0xa0, 0x3e, 0xe4, 0x2a, 0x3a, 0x13, 0x5f, 0xa4, 0xff,
	0x3f, 0x40, 0x48, 0xdf, 0x13, 0xcd, 0x75, 0x9e, 0x31, 0xef, 0x3f, 0xef, 0xb0, 0x7d, 0xc4, 0x06,
	0xe5, 0xce, 0x83, 0xe1, 0x53, 0xdc, 0xdf, 0x83, 0xb6, 0xc9, 0x3e, 0x95, 0xb3, 0x08, 0xd5, 0x7b,
	0x9e, 0x20, 0xab, 0x50, 0x45, 0xc3, 0xef, 0x42, 0x3d, 0xc1, 0x44, 0x06, 0x78, 0xaf, 0x59, 0xd5,
	0x22, 0x3b, 0x50, 0xbb, 0x52, 0x5c, 0x44, 0x6c, 0x73, 0xf1, 0x9c, 0xf1, 0x4c, 0xbf, 0x66, 0x35,
	0xfb, 0xdc, 0x85, 0x26, 0xde, 0xa7, 0xb1, 0xc2, 0x73, 0x29, 0xf4, 0x94, 0x6d, 0x11, 0x58, 0xf3,
	0x7d, 0x00, 0x07, 0x7e, 0x42, 0xae, 0xd8, 0xb6, 0xc5, 0xa8, 0x4e, 0xcb, 0x1b, 0x14, 0xac, 0x6e,
	0xea, 0x0e, 0x4e, 0xa1, 0x75, 0x99, 0xa1, 0x7a, 0x9a, 0xbe, 0x0d, 0x5b, 0x39, 0xbd, 0xcf, 0x46,
	0x76, 0xf2, 0x86, 0x7f, 0x08, 0xf5, 0xb0, 0x88, 0xd1, 0x64, 0xd5, 0xc3, 0xe6, 0xd1, 0xde, 0xfa,
	0x2e, 0x26, 0x7a, 0xf0, 0xbd, 0x0a, 0x5b, 0xc3, 0x29, 0x57, 0xd7, 0x68, 0x86, 0x0f, 0xed, 0xd7,
	0x53, 0x9b, 0x3e, 0x34, 0x74, 0x9c, 0x60, 0xa6, 0x79, 0x92, 0xda, 0x0d, 0x9b, 0x47, 0xfb, 0xab,
	0x7d, 0x82, 0x45, 0xd8, 0x7f, 0x05, 0x2d, 0x57, 0x5d, 0x48, 0x58, 0xb5, 0x12, 0xf6, 0x9e, 0xd1,
	0x96, 0x32, 0xfc, 0x0e, 0x6c, 0xa7, 0x7c, 0x6e, 0xa7, 0x76, 0xe2, 0xf4, 0xc0, 0x37, 0x00, 0xf1,
	0x7f, 0x8e, 0x30, 0x55, 0x18, 0x72, 0x8d, 0x51, 0xa1, 0x14, 0xb5, 0x4f, 0x73, 0x45, 0x0c, 0x19,
	0x1a, 0x0f, 0xad, 0x54, 0x6b, 0xed, 0xc7, 0xa5, 0x0c, 0xe3, 0x51, 0xb9, 0x82, 0x58, 0xb6, 0xd7,
	0x3c, 0xb2, 0x6a, 0x1a, 0xf5, 0x78, 0x22, 0x73, 0xa1, 0x59, 0x63, 0x91, 0x11, 0xe6, 0x4a, 0xa1,
	0x08, 0xe7, 0x0c, 0x0a, 0xf6, 0x86, 0x95, 0xdf, 0x52, 0x37, 0x2d, 0xf5, 0xbf, 0x7f, 0xda, 0x2c,
	0x58, 0x24, 0x19, 0x67, 0xc3, 0x25, 0xc4, 0x5a, 0xb6, 0xcd, 0x3e, 0x74, 0xe8, 0xf4, 0x42, 0x2a,
	0x91, 0xea, 0x02, 0x79, 0x26, 0x05, 0xdb, 0xb1, 0x7e, 0xfe, 0xf4, 0x60, 0x77, 0x31, 0xfc, 0xa8,
	0x38, 0xda, 0x58, 0x8a, 0xb5, 0xad, 0xbd, 0xbf, 0xd8, 0xda, 0x5d, 0x66, 0x09, 0x27, 0x79, 0xef,
	0xa6, 0xbc, 0xb8, 0xcf, 0xb2, 0x1a, 0x9b, 0xcf, 0xd4, 0xa8, 0xad, 0xa9, 0x61, 0x54, 0x6f, 0xf4,
	0xdf, 0x00, 0x94, 0x7e, 0x0b, 0x3b, 0xd0, 0x18, 0x4e, 0x2e, 0xc5, 0x8d, 0x90, 0x77, 0xa2, 0xbb,
	0x41, 0xe5, 0x30, 0x9c, 0x50, 0x48, 0x44, 0x94, 0xd2, 0xf5, 0x5c, 0x78, 0x84, 0x33, 0x24, 0x2f,
	0xbb, 0x95, 0xfe, 0x0c, 0x5a, 0x2b, 0x47, 0xb0, 0x0b, 0x9d, 0xe1, 0xc4, 0x21, 0xcb, 0x1e, 0x25,
	0xf0, 0x02, 0xbf, 0xe5, 0x74, 0x61, 0xd4, 0xc8, 0x36, 0x1e, 0x61, 0x38, 0x8b, 0x85, 0xe9, 0xe4,
	0x92, 0xc6, 0x0a, 0x79, 0xae, 0xa7, 0x52, 0xc5, 0x0f, 0x04, 0x56, 0x1d, 0x9b, 0xab, 0x8c, 0xba,
	0x9b, 0xfd, 0x8f, 0x94, 0xf3, 0xcc, 0x18, 0x93, 0x11, 0x04, 0xc4, 0xe5, 0xa8, 0x3a, 0xd0, 0xa4,
	0xe7, 0xdb, 0x34, 0x9d, 0x21, 0xfd, 0xad, 0x10, 0x8d, 0x0f, 0x6d, 0x02, 0x26, 0x9a, 0x2c, 0x70,
	0x55, 0x44, 0xd5, 0xa5, 0xa1, 0x83, 0xc0, 0xec, 0xec, 0x90, 0x6a, 0xff, 0x03, 0xb4, 0x56, 0x64,
	0xdf, 0x5f, 0xfa, 0x67, 0xde, 0xcb, 0x55, 0xf6, 0xc0, 0x1f, 0x07, 0x34, 0x83, 0x1e, 0x4a, 0x71,
	0x8b, 0x2a, 0xe3, 0xc6, 0x57, 0xb7, 0xcd, 0x38, 0x78, 0x87, 0x18, 0x8d, 0x25, 0x6d, 0x57, 0x39,
	0x79, 0xf9, 0x78, 0x5c, 0x81, 0x8d, 0xc7, 0xe3, 0xaa, 0xef, 0x9d, 0xd0, 0x27, 0xf3, 0xa0, 0x17,
	0xca, 0x64, 0x70, 0x65, 0x8c, 0x9e, 0xa2, 0xc2, 0x15, 0xcb, 0x7f, 0x78, 0xde, 0xef, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xf6, 0xc6, 0x28, 0xe9, 0x06, 0x05, 0x00, 0x00,
}
