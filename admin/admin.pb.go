// Code generated by protoc-gen-go.
// source: admin.proto
// DO NOT EDIT!

/*
Package admin is a generated protocol buffer package.

It is generated from these files:
	admin.proto

It has these top-level messages:
	FullPackage
	PackageList
	Package
	Empty
*/
package admin

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

type FullPackage struct {
	Name string `protobuf:"bytes,1,opt,name=Name,json=name" json:"Name,omitempty"`
	SHA  string `protobuf:"bytes,2,opt,name=SHA,json=sHA" json:"SHA,omitempty"`
}

func (m *FullPackage) Reset()                    { *m = FullPackage{} }
func (m *FullPackage) String() string            { return proto.CompactTextString(m) }
func (*FullPackage) ProtoMessage()               {}
func (*FullPackage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type PackageList struct {
	Packages []*Package `protobuf:"bytes,1,rep,name=packages" json:"packages,omitempty"`
}

func (m *PackageList) Reset()                    { *m = PackageList{} }
func (m *PackageList) String() string            { return proto.CompactTextString(m) }
func (*PackageList) ProtoMessage()               {}
func (*PackageList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PackageList) GetPackages() []*Package {
	if m != nil {
		return m.Packages
	}
	return nil
}

type Package struct {
	Name string `protobuf:"bytes,1,opt,name=Name,json=name" json:"Name,omitempty"`
}

func (m *Package) Reset()                    { *m = Package{} }
func (m *Package) String() string            { return proto.CompactTextString(m) }
func (*Package) ProtoMessage()               {}
func (*Package) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*FullPackage)(nil), "admin.FullPackage")
	proto.RegisterType((*PackageList)(nil), "admin.PackageList")
	proto.RegisterType((*Package)(nil), "admin.Package")
	proto.RegisterType((*Empty)(nil), "admin.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Admin service

type AdminClient interface {
	GetPackages(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PackageList, error)
	UpgradePackage(ctx context.Context, in *FullPackage, opts ...grpc.CallOption) (*FullPackage, error)
	AddPackage(ctx context.Context, in *FullPackage, opts ...grpc.CallOption) (*FullPackage, error)
}

type adminClient struct {
	cc *grpc.ClientConn
}

func NewAdminClient(cc *grpc.ClientConn) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) GetPackages(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PackageList, error) {
	out := new(PackageList)
	err := grpc.Invoke(ctx, "/admin.Admin/GetPackages", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) UpgradePackage(ctx context.Context, in *FullPackage, opts ...grpc.CallOption) (*FullPackage, error) {
	out := new(FullPackage)
	err := grpc.Invoke(ctx, "/admin.Admin/UpgradePackage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) AddPackage(ctx context.Context, in *FullPackage, opts ...grpc.CallOption) (*FullPackage, error) {
	out := new(FullPackage)
	err := grpc.Invoke(ctx, "/admin.Admin/AddPackage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Admin service

type AdminServer interface {
	GetPackages(context.Context, *Empty) (*PackageList, error)
	UpgradePackage(context.Context, *FullPackage) (*FullPackage, error)
	AddPackage(context.Context, *FullPackage) (*FullPackage, error)
}

func RegisterAdminServer(s *grpc.Server, srv AdminServer) {
	s.RegisterService(&_Admin_serviceDesc, srv)
}

func _Admin_GetPackages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetPackages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.Admin/GetPackages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetPackages(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_UpgradePackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FullPackage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).UpgradePackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.Admin/UpgradePackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).UpgradePackage(ctx, req.(*FullPackage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_AddPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FullPackage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).AddPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.Admin/AddPackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).AddPackage(ctx, req.(*FullPackage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Admin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "admin.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPackages",
			Handler:    _Admin_GetPackages_Handler,
		},
		{
			MethodName: "UpgradePackage",
			Handler:    _Admin_UpgradePackage_Handler,
		},
		{
			MethodName: "AddPackage",
			Handler:    _Admin_AddPackage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x4c, 0xc9, 0xcd,
	0xcc, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0x8c, 0xb9, 0xb8, 0xdd,
	0x4a, 0x73, 0x72, 0x02, 0x12, 0x93, 0xb3, 0x13, 0xd3, 0x53, 0x85, 0x84, 0xb8, 0x58, 0xfc, 0x12,
	0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x58, 0xf2, 0x80, 0x6c, 0x21, 0x01, 0x2e,
	0xe6, 0x60, 0x0f, 0x47, 0x09, 0x26, 0xb0, 0x10, 0x73, 0xb1, 0x87, 0xa3, 0x92, 0x25, 0x17, 0x37,
	0x54, 0x83, 0x4f, 0x66, 0x71, 0x89, 0x90, 0x16, 0x17, 0x47, 0x01, 0x84, 0x5b, 0x0c, 0xd4, 0xc8,
	0xac, 0xc1, 0x6d, 0xc4, 0xa7, 0x07, 0xb1, 0x0a, 0xaa, 0x2a, 0x08, 0x2e, 0xaf, 0x24, 0xcb, 0xc5,
	0x8e, 0xc7, 0x2e, 0x25, 0x76, 0x2e, 0x56, 0xd7, 0xdc, 0x82, 0x92, 0x4a, 0xa3, 0x75, 0x8c, 0x5c,
	0xac, 0x8e, 0x20, 0x33, 0x84, 0x0c, 0xb9, 0xb8, 0xdd, 0x53, 0x4b, 0xa0, 0x9a, 0x8a, 0x85, 0x78,
	0xa0, 0x46, 0x83, 0x95, 0x49, 0x09, 0xa1, 0x5a, 0x04, 0x72, 0x8e, 0x12, 0x83, 0x90, 0x15, 0x17,
	0x5f, 0x68, 0x41, 0x7a, 0x51, 0x62, 0x4a, 0x2a, 0xdc, 0x2e, 0xa8, 0x3a, 0x24, 0xbf, 0x4a, 0x61,
	0x11, 0x03, 0xea, 0x35, 0xe3, 0xe2, 0x72, 0x4c, 0x49, 0x21, 0x59, 0x5f, 0x12, 0x1b, 0x38, 0x58,
	0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x4e, 0xf7, 0xd9, 0x65, 0x01, 0x00, 0x00,
}
