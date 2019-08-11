// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example.proto

package example

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type AddUserRequest struct {
	Email                string          `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Metadata             *_struct.Struct `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AddUserRequest) Reset()         { *m = AddUserRequest{} }
func (m *AddUserRequest) String() string { return proto.CompactTextString(m) }
func (*AddUserRequest) ProtoMessage()    {}
func (*AddUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{0}
}

func (m *AddUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddUserRequest.Unmarshal(m, b)
}
func (m *AddUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddUserRequest.Marshal(b, m, deterministic)
}
func (m *AddUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddUserRequest.Merge(m, src)
}
func (m *AddUserRequest) XXX_Size() int {
	return xxx_messageInfo_AddUserRequest.Size(m)
}
func (m *AddUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddUserRequest proto.InternalMessageInfo

func (m *AddUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AddUserRequest) GetMetadata() *_struct.Struct {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type GetUserRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{1}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListUsersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersRequest) Reset()         { *m = ListUsersRequest{} }
func (m *ListUsersRequest) String() string { return proto.CompactTextString(m) }
func (*ListUsersRequest) ProtoMessage()    {}
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{2}
}

func (m *ListUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersRequest.Unmarshal(m, b)
}
func (m *ListUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersRequest.Marshal(b, m, deterministic)
}
func (m *ListUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersRequest.Merge(m, src)
}
func (m *ListUsersRequest) XXX_Size() int {
	return xxx_messageInfo_ListUsersRequest.Size(m)
}
func (m *ListUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersRequest proto.InternalMessageInfo

type User struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email                string               `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	Metadata             *_struct.Struct      `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{3}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *User) GetMetadata() *_struct.Struct {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*AddUserRequest)(nil), "example.AddUserRequest")
	proto.RegisterType((*GetUserRequest)(nil), "example.GetUserRequest")
	proto.RegisterType((*ListUsersRequest)(nil), "example.ListUsersRequest")
	proto.RegisterType((*User)(nil), "example.User")
}

func init() { proto.RegisterFile("example.proto", fileDescriptor_15a1dc8d40dadaa6) }

var fileDescriptor_15a1dc8d40dadaa6 = []byte{
	// 547 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x65, 0xb7, 0xa5, 0x74, 0xa3, 0x04, 0xb4, 0xaa, 0x68, 0x6a, 0x15, 0xb1, 0x32, 0x97,
	0xaa, 0x22, 0x76, 0x9b, 0xde, 0x1a, 0x71, 0x30, 0x20, 0x45, 0x45, 0x1c, 0x50, 0x0a, 0x97, 0x72,
	0x40, 0x6b, 0x7b, 0xea, 0x2c, 0xb2, 0xbd, 0x66, 0x77, 0x92, 0x12, 0x01, 0x17, 0x38, 0x73, 0x09,
	0x77, 0x24, 0xde, 0x81, 0x37, 0xe1, 0x15, 0x78, 0x10, 0xe4, 0x7f, 0x69, 0xe2, 0x5c, 0x10, 0xa7,
	0xec, 0xce, 0x7c, 0x99, 0xef, 0xb7, 0xdf, 0xc8, 0xa4, 0x0d, 0x1f, 0x78, 0x92, 0xc5, 0xe0, 0x64,
	0x4a, 0xa2, 0xa4, 0xdb, 0xd5, 0xd5, 0x7a, 0x10, 0x49, 0x19, 0xc5, 0xe0, 0x16, 0x65, 0x7f, 0x72,
	0xe5, 0xa2, 0x48, 0x40, 0x23, 0x4f, 0xb2, 0x52, 0x69, 0x1d, 0x34, 0x05, 0x1a, 0xd5, 0x24, 0xc0,
	0x46, 0x97, 0x67, 0xc2, 0xe5, 0x69, 0x2a, 0x91, 0xa3, 0x90, 0xa9, 0xae, 0xba, 0x8f, 0x8a, 0x9f,
	0xa0, 0x17, 0x41, 0xda, 0xd3, 0xd7, 0x3c, 0x8a, 0x40, 0xb9, 0x32, 0x2b, 0x14, 0xeb, 0x6a, 0xfb,
	0x0d, 0xe9, 0x78, 0x61, 0xf8, 0x5a, 0x83, 0x1a, 0xc1, 0xfb, 0x09, 0x68, 0xa4, 0xbb, 0x64, 0x0b,
	0x12, 0x2e, 0xe2, 0xae, 0xc1, 0x8c, 0xc3, 0x9d, 0x51, 0x79, 0xa1, 0xa7, 0xe4, 0x76, 0x02, 0xc8,
	0x43, 0x8e, 0xbc, 0x6b, 0x32, 0xe3, 0xb0, 0xd5, 0xdf, 0x73, 0x4a, 0x0c, 0xa7, 0x86, 0x74, 0x2e,
	0x0a, 0xc8, 0xd1, 0x42, 0x68, 0x33, 0xd2, 0x19, 0x02, 0x2e, 0x0f, 0xef, 0x10, 0x53, 0x84, 0xd5,
	0x64, 0x53, 0x84, 0x36, 0x25, 0x77, 0x5f, 0x08, 0x5d, 0x48, 0x74, 0xa5, 0xb1, 0x7f, 0x18, 0x64,
	0x33, 0x2f, 0x34, 0xc5, 0x37, 0x64, 0xe6, 0x32, 0xd9, 0x80, 0xb4, 0x02, 0x05, 0x1c, 0xe1, 0x6d,
	0x9e, 0x62, 0x77, 0xa3, 0x80, 0xb3, 0xd6, 0xe0, 0x5e, 0xd5, 0x11, 0x8f, 0x48, 0x29, 0xcf, 0x0b,
	0x2b, 0xcf, 0xda, 0xfc, 0xc7, 0x67, 0xf5, 0x7f, 0x6e, 0x90, 0x56, 0x0e, 0x78, 0x01, 0x6a, 0x2a,
	0x02, 0xa0, 0x9f, 0xc8, 0x76, 0x95, 0x21, 0xdd, 0x73, 0xea, 0x95, 0xaf, 0xa6, 0x6a, 0xb5, 0x17,
	0x8d, 0xbc, 0x6a, 0x9f, 0xcf, 0x3d, 0x87, 0x6c, 0x15, 0xcf, 0xa6, 0xc4, 0x0b, 0x43, 0xc6, 0xd9,
	0x44, 0x83, 0xb2, 0xf6, 0x6f, 0xce, 0x0c, 0x25, 0xc3, 0x31, 0x30, 0x0d, 0x6a, 0x0a, 0xca, 0xf9,
	0xf2, 0xfb, 0xcf, 0x77, 0x93, 0xda, 0xed, 0x62, 0xeb, 0xd3, 0x13, 0x37, 0x97, 0xe8, 0x33, 0xe3,
	0x88, 0x7e, 0x33, 0xc8, 0x76, 0x95, 0xf2, 0x92, 0xfd, 0x6a, 0xee, 0x4d, 0xfb, 0xcb, 0xb9, 0xf7,
	0xb8, 0xb6, 0xbf, 0x37, 0x04, 0xac, 0x2d, 0xfd, 0x59, 0x6e, 0x29, 0x14, 0x3b, 0x7f, 0x66, 0x1d,
	0x0c, 0x01, 0x75, 0xdd, 0xb8, 0x52, 0x32, 0x59, 0xa3, 0xd9, 0xa5, 0x74, 0x85, 0xc6, 0xfd, 0x28,
	0xc2, 0xcf, 0xf4, 0xab, 0x41, 0x76, 0x16, 0x3b, 0xa5, 0xfb, 0x0b, 0xe3, 0xe6, 0x9e, 0x9b, 0x4c,
	0xcf, 0xe7, 0x5e, 0x7f, 0x11, 0x49, 0x2e, 0x2e, 0x9c, 0xb5, 0x75, 0xbf, 0x38, 0xf3, 0x38, 0x2e,
	0xef, 0x4c, 0xa6, 0x6b, 0x20, 0x77, 0xe8, 0x6a, 0x2c, 0xc7, 0xc6, 0x93, 0x5f, 0xc6, 0xdc, 0x4b,
	0xe9, 0x56, 0x7f, 0xe3, 0xc4, 0x39, 0x3e, 0x32, 0x4c, 0xe5, 0x93, 0x87, 0xd1, 0xe8, 0xe5, 0xd3,
	0x5e, 0xc4, 0x11, 0xae, 0xf9, 0x8c, 0xf9, 0x52, 0xc4, 0xa0, 0xb2, 0x98, 0x23, 0x30, 0x05, 0x99,
	0xd4, 0x02, 0xa5, 0x9a, 0xd1, 0xc1, 0x18, 0x31, 0xd3, 0x67, 0xae, 0x1b, 0x09, 0x1c, 0x4f, 0x7c,
	0x27, 0x90, 0x89, 0xfb, 0x4e, 0x8e, 0x79, 0xea, 0x2b, 0x9e, 0x86, 0x63, 0xa9, 0x34, 0xba, 0x91,
	0xca, 0x82, 0x7a, 0x4e, 0x6f, 0x69, 0xce, 0xa5, 0xf7, 0x1f, 0x7f, 0x2a, 0x3f, 0xf1, 0x41, 0x95,
	0x84, 0x7f, 0xab, 0xb8, 0x9e, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xc8, 0xd6, 0xcf, 0x3a,
	0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*User, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (UserService_ListUsersClient, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/example.UserService/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/example.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (UserService_ListUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UserService_serviceDesc.Streams[0], "/example.UserService/ListUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceListUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserService_ListUsersClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type userServiceListUsersClient struct {
	grpc.ClientStream
}

func (x *userServiceListUsersClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	AddUser(context.Context, *AddUserRequest) (*User, error)
	GetUser(context.Context, *GetUserRequest) (*User, error)
	ListUsers(*ListUsersRequest, UserService_ListUsersServer) error
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) AddUser(ctx context.Context, req *AddUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (*UnimplementedUserServiceServer) GetUser(ctx context.Context, req *GetUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedUserServiceServer) ListUsers(req *ListUsersRequest, srv UserService_ListUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.UserService/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddUser(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListUsersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServiceServer).ListUsers(m, &userServiceListUsersServer{stream})
}

type UserService_ListUsersServer interface {
	Send(*User) error
	grpc.ServerStream
}

type userServiceListUsersServer struct {
	grpc.ServerStream
}

func (x *userServiceListUsersServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _UserService_AddUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListUsers",
			Handler:       _UserService_ListUsers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "example.proto",
}
