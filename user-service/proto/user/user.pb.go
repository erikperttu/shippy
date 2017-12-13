// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

/*
Package go_micro_srv_user is a generated protocol buffer package.
It is generated from these files:
	proto/user/user.proto
It has these top-level messages:
	User
	Request
	Response
	Token
	Error
*/
package go_micro_srv_user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Company  string `protobuf:"bytes,3,opt,name=company" json:"company,omitempty"`
	Email    string `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
	Password string `protobuf:"bytes,5,opt,name=password" json:"password,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Request struct {
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Response struct {
	User   *User    `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Users  []*User  `protobuf:"bytes,2,rep,name=users" json:"users,omitempty"`
	Errors []*Error `protobuf:"bytes,3,rep,name=errors" json:"errors,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Response) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Token struct {
	Token  string   `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Valid  bool     `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
	Errors []*Error `protobuf:"bytes,3,rep,name=errors" json:"errors,omitempty"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Token) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Error struct {
	Code        int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "go.micro.srv.user.User")
	proto.RegisterType((*Request)(nil), "go.micro.srv.user.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.user.Response")
	proto.RegisterType((*Token)(nil), "go.micro.srv.user.Token")
	proto.RegisterType((*Error)(nil), "go.micro.srv.user.Error")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserServiceClient interface {
	Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error)
	ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error)
}

type userServiceClient struct {
	c           client.Client
	serviceName string
}

func NewUserServiceClient(serviceName string, c client.Client) UserServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.user"
	}
	return &userServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *userServiceClient) Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.GetAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Auth", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.ValidateToken", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	Create(context.Context, *User, *Response) error
	Get(context.Context, *User, *Response) error
	GetAll(context.Context, *Request, *Response) error
	Auth(context.Context, *User, *Token) error
	ValidateToken(context.Context, *Token, *Token) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&UserService{hdlr}, opts...))
}

type UserService struct {
	UserServiceHandler
}

func (h *UserService) Create(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Create(ctx, in, out)
}

func (h *UserService) Get(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Get(ctx, in, out)
}

func (h *UserService) GetAll(ctx context.Context, in *Request, out *Response) error {
	return h.UserServiceHandler.GetAll(ctx, in, out)
}

func (h *UserService) Auth(ctx context.Context, in *User, out *Token) error {
	return h.UserServiceHandler.Auth(ctx, in, out)
}

func (h *UserService) ValidateToken(ctx context.Context, in *Token, out *Token) error {
	return h.UserServiceHandler.ValidateToken(ctx, in, out)
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4d, 0x6b, 0xdb, 0x40,
	0x14, 0xb4, 0xf5, 0x65, 0xfb, 0x99, 0x16, 0xfa, 0x68, 0xe9, 0xe2, 0x5e, 0x8c, 0x4e, 0x85, 0x52,
	0xb5, 0xb8, 0xc7, 0x62, 0xa8, 0x31, 0xc6, 0x77, 0xb5, 0xcd, 0x5d, 0x91, 0x1e, 0xc9, 0x12, 0x49,
	0xab, 0xec, 0xae, 0x15, 0xf2, 0x4f, 0xf2, 0xcf, 0xf2, 0x77, 0xc2, 0x3e, 0xd9, 0x21, 0x10, 0xd9,
	0x01, 0x5f, 0xa4, 0x9d, 0x79, 0xa3, 0xdd, 0xd9, 0x19, 0x04, 0x9f, 0x1a, 0xad, 0xac, 0xfa, 0xb1,
	0x33, 0xa4, 0xf9, 0x91, 0x30, 0xc6, 0x0f, 0x57, 0x2a, 0xa9, 0x64, 0xae, 0x55, 0x62, 0x74, 0x9b,
	0xb8, 0x41, 0xdc, 0x42, 0xf0, 0xdf, 0x90, 0xc6, 0xf7, 0xe0, 0xc9, 0x42, 0x0c, 0xe7, 0xc3, 0xaf,
	0x93, 0xd4, 0x93, 0x05, 0x22, 0x04, 0x75, 0x56, 0x91, 0xf0, 0x98, 0xe1, 0x35, 0x0a, 0x18, 0xe5,
	0xaa, 0x6a, 0xb2, 0xfa, 0x5e, 0xf8, 0x4c, 0x1f, 0x20, 0x7e, 0x84, 0x90, 0xaa, 0x4c, 0x96, 0x22,
	0x60, 0xbe, 0x03, 0x38, 0x83, 0x71, 0x93, 0x19, 0x73, 0xa7, 0x74, 0x21, 0x42, 0x1e, 0x3c, 0xe3,
	0x78, 0x02, 0xa3, 0x94, 0x6e, 0x77, 0x64, 0x6c, 0xfc, 0x30, 0x84, 0x71, 0x4a, 0xa6, 0x51, 0xb5,
	0x21, 0xfc, 0x06, 0x81, 0xf3, 0xc5, 0x4e, 0xa6, 0x8b, 0xcf, 0xc9, 0x2b, 0xc7, 0x89, 0xb3, 0x9b,
	0xb2, 0x08, 0xbf, 0x43, 0xe8, 0xde, 0x46, 0x78, 0x73, 0xff, 0x94, 0xba, 0x53, 0xe1, 0x4f, 0x88,
	0x48, 0x6b, 0xa5, 0x8d, 0xf0, 0x59, 0x2f, 0x7a, 0xf4, 0x1b, 0x27, 0x48, 0xf7, 0xba, 0x98, 0x20,
	0xfc, 0xa7, 0x6e, 0xa8, 0x76, 0x17, 0xb4, 0x6e, 0xb1, 0x4f, 0xa8, 0x03, 0x8e, 0x6d, 0xb3, 0x52,
	0x16, 0x9c, 0xd2, 0x38, 0xed, 0xc0, 0x19, 0xc7, 0x2c, 0x21, 0x64, 0xc2, 0xa5, 0x9e, 0xab, 0x82,
	0xf8, 0x94, 0x30, 0xe5, 0x35, 0xce, 0x61, 0x5a, 0x90, 0xc9, 0xb5, 0x6c, 0xac, 0x54, 0xf5, 0xbe,
	0x90, 0x97, 0xd4, 0xe2, 0xd1, 0x83, 0xa9, 0xbb, 0xe7, 0x5f, 0xd2, 0xad, 0xcc, 0x09, 0xff, 0x40,
	0xb4, 0xd6, 0x94, 0x59, 0xc2, 0x63, 0x89, 0xcc, 0xbe, 0xf4, 0x0c, 0x0e, 0x1d, 0xc4, 0x03, 0x5c,
	0x82, 0xbf, 0x25, 0x7b, 0xf6, 0xe7, 0x6b, 0x88, 0xb6, 0x64, 0x57, 0x65, 0x89, 0xb3, 0x5e, 0x21,
	0xf7, 0xfe, 0xd6, 0x26, 0xbf, 0x21, 0x58, 0xed, 0xec, 0xf5, 0x71, 0x13, 0x7d, 0xb9, 0x72, 0x5b,
	0xf1, 0x00, 0x37, 0xf0, 0xee, 0xc2, 0x95, 0x91, 0x59, 0xea, 0x0a, 0x3c, 0x2a, 0x3e, 0xb5, 0xcd,
	0x65, 0xc4, 0xff, 0xcd, 0xaf, 0xa7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9a, 0xfd, 0x2d, 0x97, 0x50,
	0x03, 0x00, 0x00,
}