package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SecRequest struct {
	ProductId            int64    `protobuf:"varint,1,opt,name=ProductId,proto3" json:"ProductId,omitempty"`
	Source               string   `protobuf:"bytes,2,opt,name=Source,proto3" json:"Source,omitempty"`
	AuthCode             string   `protobuf:"bytes,3,opt,name=AuthCode,proto3" json:"AuthCode,omitempty"`
	SecTime              string   `protobuf:"bytes,4,opt,name=SecTime,proto3" json:"SecTime,omitempty"`
	Nance                string   `protobuf:"bytes,5,opt,name=Nance,proto3" json:"Nance,omitempty"`
	UserId               int64    `protobuf:"varint,6,opt,name=UserId,proto3" json:"UserId,omitempty"`
	UserAuthSign         string   `protobuf:"bytes,7,opt,name=UserAuthSign,proto3" json:"UserAuthSign,omitempty"`
	AccessTime           int64    `protobuf:"varint,8,opt,name=AccessTime,proto3" json:"AccessTime,omitempty"`
	ClientAddr           string   `protobuf:"bytes,9,opt,name=ClientAddr,proto3" json:"ClientAddr,omitempty"`
	ClientRefence        string   `protobuf:"bytes,10,opt,name=ClientRefence,proto3" json:"ClientRefence,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SecRequest) Reset()         { *m = SecRequest{} }
func (m *SecRequest) String() string { return proto.CompactTextString(m) }
func (*SecRequest) ProtoMessage()    {}
func (*SecRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1202afb06d2a3a7a, []int{0}
}

func (m *SecRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecRequest.Unmarshal(m, b)
}
func (m *SecRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecRequest.Marshal(b, m, deterministic)
}
func (m *SecRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecRequest.Merge(m, src)
}
func (m *SecRequest) XXX_Size() int {
	return xxx_messageInfo_SecRequest.Size(m)
}
func (m *SecRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SecRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SecRequest proto.InternalMessageInfo

func (m *SecRequest) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *SecRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *SecRequest) GetAuthCode() string {
	if m != nil {
		return m.AuthCode
	}
	return ""
}

func (m *SecRequest) GetSecTime() string {
	if m != nil {
		return m.SecTime
	}
	return ""
}

func (m *SecRequest) GetNance() string {
	if m != nil {
		return m.Nance
	}
	return ""
}

func (m *SecRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *SecRequest) GetUserAuthSign() string {
	if m != nil {
		return m.UserAuthSign
	}
	return ""
}

func (m *SecRequest) GetAccessTime() int64 {
	if m != nil {
		return m.AccessTime
	}
	return 0
}

func (m *SecRequest) GetClientAddr() string {
	if m != nil {
		return m.ClientAddr
	}
	return ""
}

func (m *SecRequest) GetClientRefence() string {
	if m != nil {
		return m.ClientRefence
	}
	return ""
}

type SecResponse struct {
	ProductId            int64    `protobuf:"varint,1,opt,name=ProductId,proto3" json:"ProductId,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=Token,proto3" json:"Token,omitempty"`
	TokenTime            int64    `protobuf:"varint,4,opt,name=TokenTime,proto3" json:"TokenTime,omitempty"`
	Code                 int64    `protobuf:"varint,5,opt,name=Code,proto3" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SecResponse) Reset()         { *m = SecResponse{} }
func (m *SecResponse) String() string { return proto.CompactTextString(m) }
func (*SecResponse) ProtoMessage()    {}
func (*SecResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1202afb06d2a3a7a, []int{1}
}

func (m *SecResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecResponse.Unmarshal(m, b)
}
func (m *SecResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecResponse.Marshal(b, m, deterministic)
}
func (m *SecResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecResponse.Merge(m, src)
}
func (m *SecResponse) XXX_Size() int {
	return xxx_messageInfo_SecResponse.Size(m)
}
func (m *SecResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SecResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SecResponse proto.InternalMessageInfo

func (m *SecResponse) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *SecResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *SecResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *SecResponse) GetTokenTime() int64 {
	if m != nil {
		return m.TokenTime
	}
	return 0
}

func (m *SecResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*SecRequest)(nil), "pb.SecRequest")
	proto.RegisterType((*SecResponse)(nil), "pb.SecResponse")
}

func init() { proto.RegisterFile("seckill.proto", fileDescriptor_1202afb06d2a3a7a) }

var fileDescriptor_1202afb06d2a3a7a = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x4d, 0xda, 0xf4, 0xcf, 0x68, 0x2b, 0x0c, 0x45, 0x96, 0x22, 0x52, 0x82, 0x87, 0x1e,
	0xa4, 0x07, 0xbd, 0x0b, 0xa5, 0xa7, 0x22, 0x88, 0x24, 0xf5, 0x01, 0xec, 0xee, 0xa8, 0x4b, 0x63,
	0x36, 0xee, 0x26, 0xbe, 0x84, 0x2f, 0xe3, 0x23, 0xca, 0x4e, 0x62, 0xdb, 0x9c, 0xbc, 0xcd, 0xf7,
	0xfb, 0xb2, 0x33, 0x3b, 0x5f, 0x16, 0x46, 0x8e, 0xe4, 0x4e, 0x67, 0xd9, 0xa2, 0xb0, 0xa6, 0x34,
	0x18, 0x16, 0xdb, 0xf8, 0x27, 0x04, 0x48, 0x49, 0x26, 0xf4, 0x59, 0x91, 0x2b, 0xf1, 0x12, 0x86,
	0x4f, 0xd6, 0xa8, 0x4a, 0x96, 0x6b, 0x25, 0x82, 0x59, 0x30, 0xef, 0x24, 0x07, 0x80, 0x17, 0xd0,
	0x4b, 0x4d, 0x65, 0x25, 0x89, 0x70, 0x16, 0xcc, 0x87, 0x49, 0xa3, 0x70, 0x0a, 0x83, 0x65, 0x55,
	0xbe, 0xaf, 0x8c, 0x22, 0xd1, 0x61, 0x67, 0xaf, 0x51, 0x40, 0x3f, 0x25, 0xb9, 0xd1, 0x1f, 0x24,
	0xba, 0x6c, 0xfd, 0x49, 0x9c, 0x40, 0xf4, 0xf8, 0x92, 0x4b, 0x12, 0x11, 0xf3, 0x5a, 0xf8, 0x19,
	0xcf, 0x8e, 0xec, 0x5a, 0x89, 0x1e, 0x8f, 0x6f, 0x14, 0xc6, 0x70, 0xe6, 0x2b, 0xdf, 0x37, 0xd5,
	0x6f, 0xb9, 0xe8, 0xf3, 0xa1, 0x16, 0xc3, 0x2b, 0x80, 0xa5, 0x94, 0xe4, 0x1c, 0x8f, 0x1b, 0xf0,
	0xf9, 0x23, 0xe2, 0xfd, 0x55, 0xa6, 0x29, 0x2f, 0x97, 0x4a, 0x59, 0x31, 0xe4, 0x0e, 0x47, 0x04,
	0xaf, 0x61, 0x54, 0xab, 0x84, 0x5e, 0xc9, 0xdf, 0x0c, 0xf8, 0x93, 0x36, 0x8c, 0xbf, 0x03, 0x38,
	0xe5, 0xc8, 0x5c, 0x61, 0x72, 0x47, 0xff, 0x67, 0xd6, 0xec, 0x13, 0xb6, 0xf6, 0x99, 0x40, 0xb4,
	0x31, 0x3b, 0xca, 0x9b, 0xc0, 0x6a, 0xe1, 0x7b, 0x71, 0xb1, 0xcf, 0xab, 0x93, 0x1c, 0x00, 0x22,
	0x74, 0x39, 0xe3, 0x88, 0x0d, 0xae, 0x6f, 0xef, 0x61, 0x9c, 0x92, 0x7c, 0xd0, 0x59, 0x96, 0x92,
	0xfd, 0xd2, 0x92, 0xf0, 0x06, 0xfa, 0xae, 0x26, 0x38, 0x5e, 0x14, 0xdb, 0xc5, 0xe1, 0xf7, 0x4e,
	0xcf, 0xf7, 0xba, 0xbe, 0x7b, 0x7c, 0xb2, 0xed, 0xf1, 0x5b, 0xb8, 0xfb, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0xce, 0x46, 0x57, 0x89, 0x1c, 0x02, 0x00, 0x00,
}
