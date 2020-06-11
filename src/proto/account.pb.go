// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	account.proto
	session.proto

It has these top-level messages:
	LoginRequest
	LoginResponse
	LogoutRequest
	LogoutResponse
	CreateRequest
	CreateResponse
	UpdateRequest
	UpdateResponse
	ShowRequest
	ShowResponse
	TokenString
	AccountID
	AccountType
	AccountInfo
	Status
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type LoginRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto1.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto1.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type LogoutRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *LogoutRequest) Reset()                    { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string            { return proto1.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()               {}
func (*LogoutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LogoutRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type LogoutResponse struct {
	IsLoggedOut bool `protobuf:"varint,1,opt,name=isLoggedOut" json:"isLoggedOut,omitempty"`
}

func (m *LogoutResponse) Reset()                    { *m = LogoutResponse{} }
func (m *LogoutResponse) String() string            { return proto1.CompactTextString(m) }
func (*LogoutResponse) ProtoMessage()               {}
func (*LogoutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LogoutResponse) GetIsLoggedOut() bool {
	if m != nil {
		return m.IsLoggedOut
	}
	return false
}

type CreateRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Email    string `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto1.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CreateRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CreateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type CreateResponse struct {
	IsCreated bool `protobuf:"varint,1,opt,name=isCreated" json:"isCreated,omitempty"`
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto1.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CreateResponse) GetIsCreated() bool {
	if m != nil {
		return m.IsCreated
	}
	return false
}

type UpdateRequest struct {
	Token    string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Avatar   string `protobuf:"bytes,4,opt,name=avatar" json:"avatar,omitempty"`
	Password string `protobuf:"bytes,5,opt,name=password" json:"password,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto1.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UpdateRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *UpdateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UpdateRequest) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *UpdateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UpdateResponse struct {
	IsUpdated bool `protobuf:"varint,1,opt,name=isUpdated" json:"isUpdated,omitempty"`
}

func (m *UpdateResponse) Reset()                    { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string            { return proto1.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()               {}
func (*UpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *UpdateResponse) GetIsUpdated() bool {
	if m != nil {
		return m.IsUpdated
	}
	return false
}

type ShowRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *ShowRequest) Reset()                    { *m = ShowRequest{} }
func (m *ShowRequest) String() string            { return proto1.CompactTextString(m) }
func (*ShowRequest) ProtoMessage()               {}
func (*ShowRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ShowRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ShowResponse struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Email   string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Avatar  string `protobuf:"bytes,3,opt,name=avatar" json:"avatar,omitempty"`
	Diamond int64  `protobuf:"varint,4,opt,name=diamond" json:"diamond,omitempty"`
}

func (m *ShowResponse) Reset()                    { *m = ShowResponse{} }
func (m *ShowResponse) String() string            { return proto1.CompactTextString(m) }
func (*ShowResponse) ProtoMessage()               {}
func (*ShowResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ShowResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ShowResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ShowResponse) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *ShowResponse) GetDiamond() int64 {
	if m != nil {
		return m.Diamond
	}
	return 0
}

func init() {
	proto1.RegisterType((*LoginRequest)(nil), "proto.LoginRequest")
	proto1.RegisterType((*LoginResponse)(nil), "proto.LoginResponse")
	proto1.RegisterType((*LogoutRequest)(nil), "proto.LogoutRequest")
	proto1.RegisterType((*LogoutResponse)(nil), "proto.LogoutResponse")
	proto1.RegisterType((*CreateRequest)(nil), "proto.CreateRequest")
	proto1.RegisterType((*CreateResponse)(nil), "proto.CreateResponse")
	proto1.RegisterType((*UpdateRequest)(nil), "proto.UpdateRequest")
	proto1.RegisterType((*UpdateResponse)(nil), "proto.UpdateResponse")
	proto1.RegisterType((*ShowRequest)(nil), "proto.ShowRequest")
	proto1.RegisterType((*ShowResponse)(nil), "proto.ShowResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AccountService service

type AccountServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Show(ctx context.Context, in *ShowRequest, opts ...grpc.CallOption) (*ShowResponse, error)
}

type accountServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccountServiceClient(cc *grpc.ClientConn) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := grpc.Invoke(ctx, "/proto.AccountService/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := grpc.Invoke(ctx, "/proto.AccountService/Logout", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := grpc.Invoke(ctx, "/proto.AccountService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := grpc.Invoke(ctx, "/proto.AccountService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) Show(ctx context.Context, in *ShowRequest, opts ...grpc.CallOption) (*ShowResponse, error) {
	out := new(ShowResponse)
	err := grpc.Invoke(ctx, "/proto.AccountService/Show", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AccountService service

type AccountServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Show(context.Context, *ShowRequest) (*ShowResponse, error)
}

func RegisterAccountServiceServer(s *grpc.Server, srv AccountServiceServer) {
	s.RegisterService(&_AccountService_serviceDesc, srv)
}

func _AccountService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AccountService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AccountService/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AccountService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AccountService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_Show_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Show(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AccountService/Show",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Show(ctx, req.(*ShowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AccountService_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _AccountService_Logout_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _AccountService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AccountService_Update_Handler,
		},
		{
			MethodName: "Show",
			Handler:    _AccountService_Show_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account.proto",
}

func init() { proto1.RegisterFile("account.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xcd, 0x4e, 0x83, 0x40,
	0x10, 0x0e, 0x50, 0xb0, 0x9d, 0x16, 0x0e, 0xdb, 0x6a, 0x08, 0xf1, 0xd0, 0x60, 0x4c, 0x3c, 0x61,
	0x52, 0xd3, 0x07, 0x30, 0x26, 0x9e, 0x9a, 0x98, 0xd0, 0xf8, 0x00, 0x6b, 0xd9, 0x54, 0xd4, 0xb2,
	0x94, 0x5d, 0xda, 0x07, 0xf0, 0x2d, 0x7c, 0x5a, 0xd3, 0xfd, 0xa1, 0x0b, 0x21, 0x5e, 0x3c, 0xb5,
	0xf3, 0xcd, 0x7c, 0xc3, 0xf7, 0xed, 0x7c, 0xe0, 0xe3, 0xcd, 0x86, 0xd6, 0x05, 0x4f, 0xca, 0x8a,
	0x72, 0x8a, 0x5c, 0xf1, 0x13, 0x3f, 0xc3, 0x64, 0x45, 0xb7, 0x79, 0x91, 0x92, 0x7d, 0x4d, 0x18,
	0x47, 0x11, 0x0c, 0x6b, 0x46, 0xaa, 0x02, 0xef, 0x48, 0x68, 0xcd, 0xad, 0xbb, 0x51, 0xda, 0xd4,
	0xa7, 0x5e, 0x89, 0x19, 0x3b, 0xd2, 0x2a, 0x0b, 0x6d, 0xd9, 0xd3, 0x75, 0x7c, 0x0b, 0xbe, 0xda,
	0xc3, 0x4a, 0x5a, 0x30, 0x82, 0x66, 0xe0, 0x72, 0xfa, 0x49, 0x0a, 0xb5, 0x45, 0x16, 0x6a, 0x8c,
	0xd6, 0x5c, 0x7f, 0xaf, 0x7f, 0x6c, 0x01, 0x81, 0x1e, 0x53, 0xeb, 0xe6, 0x30, 0xce, 0xd9, 0x8a,
	0x6e, 0xb7, 0x24, 0x7b, 0xa9, 0xb9, 0x98, 0x1e, 0xa6, 0x26, 0x14, 0xef, 0xc1, 0x7f, 0xaa, 0x08,
	0xe6, 0xe4, 0x9f, 0x56, 0x10, 0x82, 0x81, 0xe0, 0x38, 0x02, 0x17, 0xff, 0x4f, 0x32, 0xc9, 0x0e,
	0xe7, 0x5f, 0xe1, 0x40, 0xca, 0x14, 0x45, 0x9c, 0x40, 0xa0, 0x3f, 0xa9, 0x64, 0x5e, 0xc3, 0x28,
	0x67, 0x12, 0xcb, 0x94, 0xc8, 0x33, 0x10, 0x7f, 0x5b, 0xe0, 0xbf, 0x96, 0x99, 0xa1, 0xb1, 0xd7,
	0x7e, 0xa3, 0xc0, 0xee, 0x53, 0xe0, 0x18, 0x0a, 0xd0, 0x15, 0x78, 0xf8, 0x80, 0x39, 0xae, 0x94,
	0x30, 0x55, 0xb5, 0xfc, 0xb9, 0x9d, 0x53, 0x25, 0x10, 0x68, 0x11, 0xa6, 0x6a, 0x89, 0x19, 0xaa,
	0x15, 0x10, 0xdf, 0xc0, 0x78, 0xfd, 0x4e, 0x8f, 0x7f, 0x5f, 0xec, 0x03, 0x26, 0x72, 0x48, 0xad,
	0xd4, 0x16, 0xac, 0x3e, 0x0b, 0x76, 0xbf, 0x05, 0xa7, 0x65, 0x21, 0x84, 0x8b, 0x2c, 0xc7, 0x3b,
	0x5a, 0x64, 0xc2, 0x9b, 0x93, 0xea, 0x72, 0xf1, 0x63, 0x43, 0xf0, 0x28, 0xc3, 0xbc, 0x26, 0xd5,
	0x21, 0xdf, 0x10, 0xb4, 0x00, 0x57, 0xc4, 0x0f, 0x4d, 0x65, 0xbc, 0x13, 0x33, 0xd4, 0xd1, 0xac,
	0x0d, 0x2a, 0x89, 0x4b, 0xf0, 0x64, 0xc8, 0x90, 0xd1, 0x3f, 0x47, 0x33, 0xba, 0xec, 0xa0, 0x67,
	0x9a, 0xbc, 0x67, 0x43, 0x6b, 0xc5, 0xae, 0xa1, 0x75, 0x92, 0xb1, 0x04, 0x4f, 0x3e, 0x68, 0x43,
	0x6b, 0x25, 0xa1, 0xa1, 0x75, 0x4e, 0x73, 0x0f, 0x83, 0xd3, 0xbb, 0x22, 0xa4, 0xda, 0xc6, 0x25,
	0xa2, 0x69, 0x0b, 0x93, 0x84, 0x37, 0x4f, 0x60, 0x0f, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0b,
	0xd0, 0x44, 0xda, 0xef, 0x03, 0x00, 0x00,
}
