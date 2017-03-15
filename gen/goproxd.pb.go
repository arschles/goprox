// Code generated by protoc-gen-go.
// source: goproxd.proto
// DO NOT EDIT!

/*
Package gen is a generated protocol buffer package.

It is generated from these files:
	goproxd.proto

It has these top-level messages:
	PackageMeta
	FullPackage
	Empty
*/
package gen

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type PackageMeta struct {
	Name    string `protobuf:"bytes,1,opt,name=Name,json=name" json:"Name,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=Version,json=version" json:"Version,omitempty"`
}

func (m *PackageMeta) Reset()                    { *m = PackageMeta{} }
func (m *PackageMeta) String() string            { return proto.CompactTextString(m) }
func (*PackageMeta) ProtoMessage()               {}
func (*PackageMeta) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type FullPackage struct {
	Metadata *PackageMeta `protobuf:"bytes,1,opt,name=Metadata,json=metadata" json:"Metadata,omitempty"`
	// the entire package tarball
	Payload []byte `protobuf:"bytes,2,opt,name=Payload,json=payload,proto3" json:"Payload,omitempty"`
}

func (m *FullPackage) Reset()                    { *m = FullPackage{} }
func (m *FullPackage) String() string            { return proto.CompactTextString(m) }
func (*FullPackage) ProtoMessage()               {}
func (*FullPackage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FullPackage) GetMetadata() *PackageMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*PackageMeta)(nil), "goproxd.PackageMeta")
	proto.RegisterType((*FullPackage)(nil), "goproxd.FullPackage")
	proto.RegisterType((*Empty)(nil), "goproxd.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for GoProxD service

type GoProxDClient interface {
	GoGet(ctx context.Context, in *PackageMeta, opts ...grpc.CallOption) (*FullPackage, error)
	UpgradePackage(ctx context.Context, in *FullPackage, opts ...grpc.CallOption) (*Empty, error)
	AddPackage(ctx context.Context, in *FullPackage, opts ...grpc.CallOption) (*Empty, error)
	PackageExists(ctx context.Context, in *PackageMeta, opts ...grpc.CallOption) (*PackageMeta, error)
}

type goProxDClient struct {
	cc *grpc.ClientConn
}

func NewGoProxDClient(cc *grpc.ClientConn) GoProxDClient {
	return &goProxDClient{cc}
}

func (c *goProxDClient) GoGet(ctx context.Context, in *PackageMeta, opts ...grpc.CallOption) (*FullPackage, error) {
	out := new(FullPackage)
	err := grpc.Invoke(ctx, "/goproxd.GoProxD/GoGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goProxDClient) UpgradePackage(ctx context.Context, in *FullPackage, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/goproxd.GoProxD/UpgradePackage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goProxDClient) AddPackage(ctx context.Context, in *FullPackage, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/goproxd.GoProxD/AddPackage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goProxDClient) PackageExists(ctx context.Context, in *PackageMeta, opts ...grpc.CallOption) (*PackageMeta, error) {
	out := new(PackageMeta)
	err := grpc.Invoke(ctx, "/goproxd.GoProxD/PackageExists", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GoProxD service

type GoProxDServer interface {
	GoGet(context.Context, *PackageMeta) (*FullPackage, error)
	UpgradePackage(context.Context, *FullPackage) (*Empty, error)
	AddPackage(context.Context, *FullPackage) (*Empty, error)
	PackageExists(context.Context, *PackageMeta) (*PackageMeta, error)
}

func RegisterGoProxDServer(s *grpc.Server, srv GoProxDServer) {
	s.RegisterService(&_GoProxD_serviceDesc, srv)
}

func _GoProxD_GoGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PackageMeta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoProxDServer).GoGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goproxd.GoProxD/GoGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoProxDServer).GoGet(ctx, req.(*PackageMeta))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoProxD_UpgradePackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FullPackage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoProxDServer).UpgradePackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goproxd.GoProxD/UpgradePackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoProxDServer).UpgradePackage(ctx, req.(*FullPackage))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoProxD_AddPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FullPackage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoProxDServer).AddPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goproxd.GoProxD/AddPackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoProxDServer).AddPackage(ctx, req.(*FullPackage))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoProxD_PackageExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PackageMeta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoProxDServer).PackageExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goproxd.GoProxD/PackageExists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoProxDServer).PackageExists(ctx, req.(*PackageMeta))
	}
	return interceptor(ctx, in, info, handler)
}

var _GoProxD_serviceDesc = grpc.ServiceDesc{
	ServiceName: "goproxd.GoProxD",
	HandlerType: (*GoProxDServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GoGet",
			Handler:    _GoProxD_GoGet_Handler,
		},
		{
			MethodName: "UpgradePackage",
			Handler:    _GoProxD_UpgradePackage_Handler,
		},
		{
			MethodName: "AddPackage",
			Handler:    _GoProxD_AddPackage_Handler,
		},
		{
			MethodName: "PackageExists",
			Handler:    _GoProxD_PackageExists_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xcf, 0x2f, 0x28,
	0xca, 0xaf, 0x48, 0xd1, 0x03, 0x92, 0x25, 0xf9, 0x42, 0xec, 0x50, 0xae, 0x92, 0x35, 0x17, 0x77,
	0x40, 0x62, 0x72, 0x76, 0x62, 0x7a, 0xaa, 0x6f, 0x6a, 0x49, 0xa2, 0x90, 0x10, 0x17, 0x8b, 0x5f,
	0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x4b, 0x1e, 0x90, 0x2d, 0x24, 0xc1,
	0xc5, 0x1e, 0x96, 0x5a, 0x54, 0x9c, 0x99, 0x9f, 0x27, 0xc1, 0x04, 0x16, 0x66, 0x2f, 0x83, 0x70,
	0x95, 0x22, 0xb9, 0xb8, 0xdd, 0x4a, 0x73, 0x72, 0xa0, 0x06, 0x08, 0x19, 0x70, 0x71, 0x80, 0x0c,
	0x49, 0x49, 0x2c, 0x49, 0x04, 0x1b, 0xc0, 0x6d, 0x24, 0xa2, 0x07, 0xb3, 0x16, 0xc9, 0x92, 0x20,
	0x8e, 0x5c, 0xa8, 0x2a, 0x90, 0xd1, 0x01, 0x89, 0x95, 0x39, 0xf9, 0x89, 0x29, 0x60, 0xa3, 0x79,
	0x82, 0xd8, 0x0b, 0x20, 0x5c, 0x25, 0x76, 0x2e, 0x56, 0xd7, 0xdc, 0x82, 0x92, 0x4a, 0xa3, 0xf7,
	0x8c, 0x5c, 0xec, 0xee, 0xf9, 0x01, 0x40, 0x43, 0x5c, 0x84, 0x4c, 0xb9, 0x58, 0xdd, 0xf3, 0xdd,
	0x53, 0x4b, 0x84, 0xb0, 0x9a, 0x2b, 0x85, 0x10, 0x45, 0x72, 0x95, 0x12, 0x83, 0x90, 0x05, 0x17,
	0x5f, 0x68, 0x41, 0x7a, 0x51, 0x62, 0x4a, 0x2a, 0xcc, 0xa5, 0x58, 0x55, 0x4a, 0xf1, 0xc1, 0x45,
	0xc1, 0x56, 0x03, 0x75, 0x9a, 0x70, 0x71, 0x39, 0xa6, 0xa4, 0x90, 0xaa, 0xcb, 0x96, 0x8b, 0x17,
	0x2a, 0xe9, 0x5a, 0x91, 0x59, 0x5c, 0x52, 0x4c, 0xd0, 0xb9, 0x48, 0xa2, 0x4a, 0x0c, 0x4e, 0xac,
	0x51, 0xcc, 0xe9, 0xa9, 0x79, 0x49, 0x6c, 0xe0, 0x98, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x5f, 0x1c, 0x8f, 0x01, 0xba, 0x01, 0x00, 0x00,
}
