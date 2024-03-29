package pb

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CheckTokenRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckTokenRequest) Reset()         { *m = CheckTokenRequest{} }
func (m *CheckTokenRequest) String() string { return proto.CompactTextString(m) }
func (*CheckTokenRequest) ProtoMessage()    {}
func (*CheckTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ce0b12f599e9f07, []int{0}
}

func (m *CheckTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckTokenRequest.Unmarshal(m, b)
}
func (m *CheckTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckTokenRequest.Marshal(b, m, deterministic)
}
func (m *CheckTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckTokenRequest.Merge(m, src)
}
func (m *CheckTokenRequest) XXX_Size() int {
	return xxx_messageInfo_CheckTokenRequest.Size(m)
}
func (m *CheckTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckTokenRequest proto.InternalMessageInfo

func (m *CheckTokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ClientDetails struct {
	ClientId                    string   `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	AccessTokenValiditySeconds  int32    `protobuf:"varint,2,opt,name=accessTokenValiditySeconds,proto3" json:"accessTokenValiditySeconds,omitempty"`
	RefreshTokenValiditySeconds int32    `protobuf:"varint,3,opt,name=refreshTokenValiditySeconds,proto3" json:"refreshTokenValiditySeconds,omitempty"`
	AuthorizedGrantTypes        []string `protobuf:"bytes,4,rep,name=authorizedGrantTypes,proto3" json:"authorizedGrantTypes,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *ClientDetails) Reset()         { *m = ClientDetails{} }
func (m *ClientDetails) String() string { return proto.CompactTextString(m) }
func (*ClientDetails) ProtoMessage()    {}
func (*ClientDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ce0b12f599e9f07, []int{1}
}

func (m *ClientDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientDetails.Unmarshal(m, b)
}
func (m *ClientDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientDetails.Marshal(b, m, deterministic)
}
func (m *ClientDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientDetails.Merge(m, src)
}
func (m *ClientDetails) XXX_Size() int {
	return xxx_messageInfo_ClientDetails.Size(m)
}
func (m *ClientDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientDetails.DiscardUnknown(m)
}

var xxx_messageInfo_ClientDetails proto.InternalMessageInfo

func (m *ClientDetails) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *ClientDetails) GetAccessTokenValiditySeconds() int32 {
	if m != nil {
		return m.AccessTokenValiditySeconds
	}
	return 0
}

func (m *ClientDetails) GetRefreshTokenValiditySeconds() int32 {
	if m != nil {
		return m.RefreshTokenValiditySeconds
	}
	return 0
}

func (m *ClientDetails) GetAuthorizedGrantTypes() []string {
	if m != nil {
		return m.AuthorizedGrantTypes
	}
	return nil
}

type UserDetails struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Authorities          []string `protobuf:"bytes,3,rep,name=authorities,proto3" json:"authorities,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserDetails) Reset()         { *m = UserDetails{} }
func (m *UserDetails) String() string { return proto.CompactTextString(m) }
func (*UserDetails) ProtoMessage()    {}
func (*UserDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ce0b12f599e9f07, []int{2}
}

func (m *UserDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDetails.Unmarshal(m, b)
}
func (m *UserDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDetails.Marshal(b, m, deterministic)
}
func (m *UserDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDetails.Merge(m, src)
}
func (m *UserDetails) XXX_Size() int {
	return xxx_messageInfo_UserDetails.Size(m)
}
func (m *UserDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDetails.DiscardUnknown(m)
}

var xxx_messageInfo_UserDetails proto.InternalMessageInfo

func (m *UserDetails) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserDetails) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserDetails) GetAuthorities() []string {
	if m != nil {
		return m.Authorities
	}
	return nil
}

type CheckTokenResponse struct {
	UserDetails          *UserDetails   `protobuf:"bytes,1,opt,name=userDetails,proto3" json:"userDetails,omitempty"`
	ClientDetails        *ClientDetails `protobuf:"bytes,2,opt,name=clientDetails,proto3" json:"clientDetails,omitempty"`
	IsValidToken         bool           `protobuf:"varint,3,opt,name=isValidToken,proto3" json:"isValidToken,omitempty"`
	Err                  string         `protobuf:"bytes,4,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CheckTokenResponse) Reset()         { *m = CheckTokenResponse{} }
func (m *CheckTokenResponse) String() string { return proto.CompactTextString(m) }
func (*CheckTokenResponse) ProtoMessage()    {}
func (*CheckTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ce0b12f599e9f07, []int{3}
}

func (m *CheckTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckTokenResponse.Unmarshal(m, b)
}
func (m *CheckTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckTokenResponse.Marshal(b, m, deterministic)
}
func (m *CheckTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckTokenResponse.Merge(m, src)
}
func (m *CheckTokenResponse) XXX_Size() int {
	return xxx_messageInfo_CheckTokenResponse.Size(m)
}
func (m *CheckTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckTokenResponse proto.InternalMessageInfo

func (m *CheckTokenResponse) GetUserDetails() *UserDetails {
	if m != nil {
		return m.UserDetails
	}
	return nil
}

func (m *CheckTokenResponse) GetClientDetails() *ClientDetails {
	if m != nil {
		return m.ClientDetails
	}
	return nil
}

func (m *CheckTokenResponse) GetIsValidToken() bool {
	if m != nil {
		return m.IsValidToken
	}
	return false
}

func (m *CheckTokenResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*CheckTokenRequest)(nil), "pb.CheckTokenRequest")
	proto.RegisterType((*ClientDetails)(nil), "pb.ClientDetails")
	proto.RegisterType((*UserDetails)(nil), "pb.UserDetails")
	proto.RegisterType((*CheckTokenResponse)(nil), "pb.CheckTokenResponse")
}

func init() { proto.RegisterFile("oauth.proto", fileDescriptor_7ce0b12f599e9f07) }

var fileDescriptor_7ce0b12f599e9f07 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x4f, 0xc2, 0x30,
	0x14, 0xc7, 0x33, 0x06, 0x04, 0xde, 0x20, 0x4a, 0x83, 0x64, 0xc1, 0xcb, 0xb2, 0x13, 0x5e, 0x48,
	0x9c, 0x07, 0x0f, 0x26, 0x46, 0x83, 0x89, 0x31, 0x1e, 0x4c, 0x0a, 0x7a, 0x1f, 0xdd, 0x33, 0x6b,
	0xc0, 0x75, 0xb6, 0x9d, 0x09, 0x7e, 0x30, 0xbf, 0x8f, 0xdf, 0xc4, 0xb4, 0x4c, 0x18, 0x84, 0x70,
	0xeb, 0xff, 0xed, 0xbf, 0xf7, 0xfe, 0xbf, 0xd7, 0x82, 0x27, 0xe2, 0x42, 0xa7, 0xe3, 0x5c, 0x0a,
	0x2d, 0x48, 0x2d, 0x9f, 0x87, 0x17, 0xd0, 0x9b, 0xa4, 0xc8, 0x16, 0x33, 0xb1, 0xc0, 0x8c, 0xe2,
	0x67, 0x81, 0x4a, 0x93, 0x3e, 0x34, 0xb4, 0xd1, 0xbe, 0x13, 0x38, 0xa3, 0x36, 0x5d, 0x8b, 0xf0,
	0xd7, 0x81, 0xee, 0x64, 0xc9, 0x31, 0xd3, 0x0f, 0xa8, 0x63, 0xbe, 0x54, 0x64, 0x08, 0x2d, 0x66,
	0x0b, 0x4f, 0x49, 0x69, 0xdd, 0x68, 0x72, 0x0b, 0xc3, 0x98, 0x31, 0x54, 0xca, 0x76, 0x7e, 0x8b,
	0x97, 0x3c, 0xe1, 0x7a, 0x35, 0x45, 0x26, 0xb2, 0x44, 0xf9, 0xb5, 0xc0, 0x19, 0x35, 0xe8, 0x11,
	0x07, 0xb9, 0x83, 0x73, 0x89, 0xef, 0x12, 0x55, 0x7a, 0xb0, 0x81, 0x6b, 0x1b, 0x1c, 0xb3, 0x90,
	0x08, 0xfa, 0x06, 0x56, 0x48, 0xfe, 0x8d, 0xc9, 0xa3, 0x8c, 0x33, 0x3d, 0x5b, 0xe5, 0xa8, 0xfc,
	0x7a, 0xe0, 0x8e, 0xda, 0xf4, 0xe0, 0xb7, 0x90, 0x81, 0xf7, 0xaa, 0x50, 0xfe, 0x03, 0x0e, 0xa0,
	0x59, 0x28, 0x94, 0x25, 0x9e, 0x4b, 0x4b, 0x65, 0xc0, 0xcd, 0x29, 0x8b, 0x3f, 0xd0, 0xa2, 0xb4,
	0xe9, 0x46, 0x93, 0x00, 0xbc, 0xb2, 0xb5, 0xe6, 0x68, 0x82, 0x9a, 0x69, 0xd5, 0x52, 0xf8, 0xe3,
	0x00, 0xa9, 0x2e, 0x5d, 0xe5, 0x22, 0x53, 0x48, 0x2e, 0xc1, 0x2b, 0xb6, 0xb3, 0xed, 0x44, 0x2f,
	0x3a, 0x19, 0xe7, 0xf3, 0x71, 0x25, 0x12, 0xad, 0x7a, 0xc8, 0x35, 0x74, 0x59, 0xf5, 0x46, 0x6c,
	0x18, 0x2f, 0xea, 0x99, 0x9f, 0x76, 0xae, 0x8a, 0xee, 0xfa, 0x48, 0x08, 0x1d, 0xae, 0xec, 0xc2,
	0x6c, 0x06, 0xbb, 0xce, 0x16, 0xdd, 0xa9, 0x91, 0x53, 0x70, 0x51, 0x4a, 0xbf, 0x6e, 0xf9, 0xcc,
	0x31, 0x7a, 0x86, 0xce, 0xcb, 0x7d, 0xa1, 0xd3, 0x29, 0xca, 0x2f, 0xce, 0x90, 0xdc, 0x00, 0x6c,
	0x39, 0xc8, 0x99, 0x9d, 0xba, 0xff, 0x98, 0x86, 0x83, 0xfd, 0xf2, 0x1a, 0x77, 0xde, 0xb4, 0x8f,
	0xf0, 0xea, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xd3, 0xa2, 0x2a, 0x89, 0x93, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OAuthServiceClient is the client API for OAuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OAuthServiceClient interface {
	// token 校验
	CheckToken(ctx context.Context, in *CheckTokenRequest, opts ...grpc.CallOption) (*CheckTokenResponse, error)
}

type oAuthServiceClient struct {
	cc *grpc.ClientConn
}

func NewOAuthServiceClient(cc *grpc.ClientConn) OAuthServiceClient {
	return &oAuthServiceClient{cc}
}

func (c *oAuthServiceClient) CheckToken(ctx context.Context, in *CheckTokenRequest, opts ...grpc.CallOption) (*CheckTokenResponse, error) {
	out := new(CheckTokenResponse)
	err := c.cc.Invoke(ctx, "/pb.OAuthService/CheckToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OAuthServiceServer is the server API for OAuthService service.
type OAuthServiceServer interface {
	// token 校验
	CheckToken(context.Context, *CheckTokenRequest) (*CheckTokenResponse, error)
}

// UnimplementedOAuthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOAuthServiceServer struct {
}

func (*UnimplementedOAuthServiceServer) CheckToken(ctx context.Context, req *CheckTokenRequest) (*CheckTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckToken not implemented")
}

func RegisterOAuthServiceServer(s *grpc.Server, srv OAuthServiceServer) {
	s.RegisterService(&_OAuthService_serviceDesc, srv)
}

func _OAuthService_CheckToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuthServiceServer).CheckToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OAuthService/CheckToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuthServiceServer).CheckToken(ctx, req.(*CheckTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OAuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.OAuthService",
	HandlerType: (*OAuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckToken",
			Handler:    _OAuthService_CheckToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "oauth.proto",
}
