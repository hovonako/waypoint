// Code generated by protoc-gen-go. DO NOT EDIT.
// source: platform.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type Deploy struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Deploy) Reset()         { *m = Deploy{} }
func (m *Deploy) String() string { return proto.CompactTextString(m) }
func (*Deploy) ProtoMessage()    {}
func (*Deploy) Descriptor() ([]byte, []int) {
	return fileDescriptor_918f3d50bfb447e4, []int{0}
}

func (m *Deploy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Deploy.Unmarshal(m, b)
}
func (m *Deploy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Deploy.Marshal(b, m, deterministic)
}
func (m *Deploy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deploy.Merge(m, src)
}
func (m *Deploy) XXX_Size() int {
	return xxx_messageInfo_Deploy.Size(m)
}
func (m *Deploy) XXX_DiscardUnknown() {
	xxx_messageInfo_Deploy.DiscardUnknown(m)
}

var xxx_messageInfo_Deploy proto.InternalMessageInfo

type Deploy_Args struct {
	// args is the list of argument types. This will include some of the
	// standard types in this file (in the Args message namespace) as well
	// as custom types declared by the FuncSpec that the plugin is expected
	// to understand how to decode.
	Args                 []*any.Any `protobuf:"bytes,1,rep,name=args,proto3" json:"args,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Deploy_Args) Reset()         { *m = Deploy_Args{} }
func (m *Deploy_Args) String() string { return proto.CompactTextString(m) }
func (*Deploy_Args) ProtoMessage()    {}
func (*Deploy_Args) Descriptor() ([]byte, []int) {
	return fileDescriptor_918f3d50bfb447e4, []int{0, 0}
}

func (m *Deploy_Args) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Deploy_Args.Unmarshal(m, b)
}
func (m *Deploy_Args) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Deploy_Args.Marshal(b, m, deterministic)
}
func (m *Deploy_Args) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deploy_Args.Merge(m, src)
}
func (m *Deploy_Args) XXX_Size() int {
	return xxx_messageInfo_Deploy_Args.Size(m)
}
func (m *Deploy_Args) XXX_DiscardUnknown() {
	xxx_messageInfo_Deploy_Args.DiscardUnknown(m)
}

var xxx_messageInfo_Deploy_Args proto.InternalMessageInfo

func (m *Deploy_Args) GetArgs() []*any.Any {
	if m != nil {
		return m.Args
	}
	return nil
}

type Deploy_Resp struct {
	// result is the resulting opaque data type
	Result               *any.Any `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Deploy_Resp) Reset()         { *m = Deploy_Resp{} }
func (m *Deploy_Resp) String() string { return proto.CompactTextString(m) }
func (*Deploy_Resp) ProtoMessage()    {}
func (*Deploy_Resp) Descriptor() ([]byte, []int) {
	return fileDescriptor_918f3d50bfb447e4, []int{0, 1}
}

func (m *Deploy_Resp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Deploy_Resp.Unmarshal(m, b)
}
func (m *Deploy_Resp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Deploy_Resp.Marshal(b, m, deterministic)
}
func (m *Deploy_Resp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deploy_Resp.Merge(m, src)
}
func (m *Deploy_Resp) XXX_Size() int {
	return xxx_messageInfo_Deploy_Resp.Size(m)
}
func (m *Deploy_Resp) XXX_DiscardUnknown() {
	xxx_messageInfo_Deploy_Resp.DiscardUnknown(m)
}

var xxx_messageInfo_Deploy_Resp proto.InternalMessageInfo

func (m *Deploy_Resp) GetResult() *any.Any {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*Deploy)(nil), "proto.Deploy")
	proto.RegisterType((*Deploy_Args)(nil), "proto.Deploy.Args")
	proto.RegisterType((*Deploy_Resp)(nil), "proto.Deploy.Resp")
}

func init() {
	proto.RegisterFile("platform.proto", fileDescriptor_918f3d50bfb447e4)
}

var fileDescriptor_918f3d50bfb447e4 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0x4d, 0x4b, 0x72, 0x41,
	0x14, 0x46, 0xf4, 0xf5, 0xad, 0xa3, 0x45, 0x1c, 0x2c, 0x6c, 0x82, 0x90, 0x56, 0x2e, 0x64, 0x14,
	0x0b, 0xdc, 0x44, 0x24, 0x5a, 0x20, 0xb4, 0x08, 0xfd, 0x05, 0x2a, 0xc7, 0x41, 0xb8, 0x77, 0x66,
	0x9a, 0x8f, 0xc5, 0xfd, 0xb9, 0xfd, 0x93, 0x70, 0x66, 0x14, 0x52, 0xee, 0xa2, 0xd5, 0xe5, 0x9c,
	0xe7, 0x63, 0x9e, 0xf3, 0x5c, 0xb8, 0xd4, 0xd9, 0xd2, 0x6d, 0x94, 0xc9, 0xb9, 0x36, 0xca, 0x29,
	0xfc, 0x17, 0x3e, 0xec, 0x56, 0x28, 0x25, 0x32, 0xea, 0x87, 0x69, 0xe5, 0x37, 0xfd, 0xa5, 0x2c,
	0x22, 0x83, 0xdd, 0x1d, 0x43, 0x94, 0x6b, 0xb7, 0x07, 0x9b, 0x3a, 0xf3, 0x62, 0x2b, 0xe3, 0xf4,
	0xa0, 0xa1, 0x3e, 0x25, 0x9d, 0xa9, 0x82, 0x0d, 0xa0, 0x36, 0x36, 0xc2, 0x62, 0x17, 0x6a, 0x4b,
	0x23, 0x6c, 0xbb, 0xd2, 0xa9, 0x76, 0x1b, 0xc3, 0x16, 0x8f, 0x5e, 0x7c, 0xef, 0xc5, 0xc7, 0xb2,
	0x98, 0x07, 0x06, 0x7b, 0x82, 0xda, 0x9c, 0xac, 0xc6, 0x1e, 0xd4, 0x0d, 0x59, 0x9f, 0xb9, 0x76,
	0xa5, 0x53, 0x29, 0xd5, 0x24, 0xce, 0xf0, 0xbb, 0x0a, 0x67, 0x9f, 0xe9, 0x22, 0x7c, 0x85, 0xe6,
	0x44, 0xc9, 0xcd, 0x56, 0x2c, 0x9c, 0xf1, 0x6b, 0x87, 0x37, 0x27, 0xd2, 0xb7, 0x5d, 0x74, 0xd6,
	0x8e, 0x0b, 0x1e, 0xc9, 0x3c, 0xb2, 0xc3, 0xe3, 0x13, 0x38, 0x8f, 0x4b, 0x6f, 0x08, 0xef, 0x7f,
	0xd3, 0x0e, 0xc0, 0x9c, 0xbe, 0x3c, 0x59, 0xc7, 0x4a, 0xec, 0xb1, 0x07, 0x10, 0x5b, 0x58, 0x68,
	0x5a, 0x63, 0x33, 0xb9, 0xc4, 0xa7, 0xaf, 0xd2, 0xf4, 0xee, 0xe5, 0x7a, 0x07, 0x0f, 0x71, 0xb0,
	0xef, 0x0c, 0x31, 0x61, 0x71, 0xe4, 0xbb, 0xfe, 0xd8, 0xd1, 0x2e, 0x84, 0x7c, 0x86, 0xc6, 0xcc,
	0x4e, 0xc9, 0x3a, 0xa3, 0x0a, 0x32, 0xa5, 0x57, 0x5e, 0x27, 0xe9, 0x2c, 0xd7, 0x19, 0xe5, 0x24,
	0x9d, 0x0d, 0xea, 0x11, 0x34, 0x92, 0x36, 0xc4, 0x2b, 0x53, 0x9f, 0x06, 0x1d, 0xc1, 0xff, 0x24,
	0xc4, 0xd6, 0x11, 0x18, 0xb3, 0x96, 0xf5, 0xf1, 0x02, 0x17, 0x33, 0xfb, 0xa1, 0xc4, 0xe1, 0x3f,
	0xfd, 0x2d, 0xf1, 0xaa, 0x1e, 0xb6, 0x8f, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8a, 0xa4, 0x29,
	0x83, 0xbb, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PlatformClient is the client API for Platform service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlatformClient interface {
	ConfigStruct(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Config_StructResp, error)
	Configure(ctx context.Context, in *Config_ConfigureRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	DeploySpec(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*FuncSpec2, error)
	Deploy(ctx context.Context, in *Deploy_Args, opts ...grpc.CallOption) (*Deploy_Resp, error)
	// component.Destroyer optional implementation
	IsDestroyer(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ImplementsResp, error)
	DestroySpec(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*FuncSpec2, error)
	Destroy(ctx context.Context, in *FuncSpec_Args, opts ...grpc.CallOption) (*empty.Empty, error)
	// IsLogPlatform returns true if this platform also implements LogPlatform.
	IsLogPlatform(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ImplementsResp, error)
}

type platformClient struct {
	cc grpc.ClientConnInterface
}

func NewPlatformClient(cc grpc.ClientConnInterface) PlatformClient {
	return &platformClient{cc}
}

func (c *platformClient) ConfigStruct(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Config_StructResp, error) {
	out := new(Config_StructResp)
	err := c.cc.Invoke(ctx, "/proto.Platform/ConfigStruct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) Configure(ctx context.Context, in *Config_ConfigureRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.Platform/Configure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) DeploySpec(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*FuncSpec2, error) {
	out := new(FuncSpec2)
	err := c.cc.Invoke(ctx, "/proto.Platform/DeploySpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) Deploy(ctx context.Context, in *Deploy_Args, opts ...grpc.CallOption) (*Deploy_Resp, error) {
	out := new(Deploy_Resp)
	err := c.cc.Invoke(ctx, "/proto.Platform/Deploy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) IsDestroyer(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ImplementsResp, error) {
	out := new(ImplementsResp)
	err := c.cc.Invoke(ctx, "/proto.Platform/IsDestroyer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) DestroySpec(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*FuncSpec2, error) {
	out := new(FuncSpec2)
	err := c.cc.Invoke(ctx, "/proto.Platform/DestroySpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) Destroy(ctx context.Context, in *FuncSpec_Args, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.Platform/Destroy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) IsLogPlatform(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ImplementsResp, error) {
	out := new(ImplementsResp)
	err := c.cc.Invoke(ctx, "/proto.Platform/IsLogPlatform", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlatformServer is the server API for Platform service.
type PlatformServer interface {
	ConfigStruct(context.Context, *empty.Empty) (*Config_StructResp, error)
	Configure(context.Context, *Config_ConfigureRequest) (*empty.Empty, error)
	DeploySpec(context.Context, *Empty) (*FuncSpec2, error)
	Deploy(context.Context, *Deploy_Args) (*Deploy_Resp, error)
	// component.Destroyer optional implementation
	IsDestroyer(context.Context, *empty.Empty) (*ImplementsResp, error)
	DestroySpec(context.Context, *empty.Empty) (*FuncSpec2, error)
	Destroy(context.Context, *FuncSpec_Args) (*empty.Empty, error)
	// IsLogPlatform returns true if this platform also implements LogPlatform.
	IsLogPlatform(context.Context, *empty.Empty) (*ImplementsResp, error)
}

// UnimplementedPlatformServer can be embedded to have forward compatible implementations.
type UnimplementedPlatformServer struct {
}

func (*UnimplementedPlatformServer) ConfigStruct(ctx context.Context, req *empty.Empty) (*Config_StructResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigStruct not implemented")
}
func (*UnimplementedPlatformServer) Configure(ctx context.Context, req *Config_ConfigureRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Configure not implemented")
}
func (*UnimplementedPlatformServer) DeploySpec(ctx context.Context, req *Empty) (*FuncSpec2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeploySpec not implemented")
}
func (*UnimplementedPlatformServer) Deploy(ctx context.Context, req *Deploy_Args) (*Deploy_Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deploy not implemented")
}
func (*UnimplementedPlatformServer) IsDestroyer(ctx context.Context, req *empty.Empty) (*ImplementsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsDestroyer not implemented")
}
func (*UnimplementedPlatformServer) DestroySpec(ctx context.Context, req *empty.Empty) (*FuncSpec2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DestroySpec not implemented")
}
func (*UnimplementedPlatformServer) Destroy(ctx context.Context, req *FuncSpec_Args) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Destroy not implemented")
}
func (*UnimplementedPlatformServer) IsLogPlatform(ctx context.Context, req *empty.Empty) (*ImplementsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsLogPlatform not implemented")
}

func RegisterPlatformServer(s *grpc.Server, srv PlatformServer) {
	s.RegisterService(&_Platform_serviceDesc, srv)
}

func _Platform_ConfigStruct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).ConfigStruct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Platform/ConfigStruct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).ConfigStruct(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Config_ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Platform/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).Configure(ctx, req.(*Config_ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_DeploySpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).DeploySpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Platform/DeploySpec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).DeploySpec(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_Deploy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Deploy_Args)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).Deploy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Platform/Deploy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).Deploy(ctx, req.(*Deploy_Args))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_IsDestroyer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).IsDestroyer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Platform/IsDestroyer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).IsDestroyer(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_DestroySpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).DestroySpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Platform/DestroySpec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).DestroySpec(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_Destroy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FuncSpec_Args)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).Destroy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Platform/Destroy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).Destroy(ctx, req.(*FuncSpec_Args))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_IsLogPlatform_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).IsLogPlatform(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Platform/IsLogPlatform",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).IsLogPlatform(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Platform_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Platform",
	HandlerType: (*PlatformServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConfigStruct",
			Handler:    _Platform_ConfigStruct_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _Platform_Configure_Handler,
		},
		{
			MethodName: "DeploySpec",
			Handler:    _Platform_DeploySpec_Handler,
		},
		{
			MethodName: "Deploy",
			Handler:    _Platform_Deploy_Handler,
		},
		{
			MethodName: "IsDestroyer",
			Handler:    _Platform_IsDestroyer_Handler,
		},
		{
			MethodName: "DestroySpec",
			Handler:    _Platform_DestroySpec_Handler,
		},
		{
			MethodName: "Destroy",
			Handler:    _Platform_Destroy_Handler,
		},
		{
			MethodName: "IsLogPlatform",
			Handler:    _Platform_IsLogPlatform_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "platform.proto",
}
