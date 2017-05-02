// Code generated by protoc-gen-go.
// source: gitbin/v1alpha1/gitbin.proto
// DO NOT EDIT!

/*
Package gitbin is a generated protocol buffer package.

It is generated from these files:
	gitbin/v1alpha1/gitbin.proto

It has these top-level messages:
	BinarySearch
	BinaryDownload
*/
package gitbin

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BinaryDownload_Format int32

const (
	BinaryDownload_BINARY BinaryDownload_Format = 0
	BinaryDownload_TARGZ  BinaryDownload_Format = 1
	BinaryDownload_ZIP    BinaryDownload_Format = 2
)

var BinaryDownload_Format_name = map[int32]string{
	0: "BINARY",
	1: "TARGZ",
	2: "ZIP",
}
var BinaryDownload_Format_value = map[string]int32{
	"BINARY": 0,
	"TARGZ":  1,
	"ZIP":    2,
}

func (x BinaryDownload_Format) String() string {
	return proto.EnumName(BinaryDownload_Format_name, int32(x))
}
func (BinaryDownload_Format) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

// Search criteria for a binary download
// First two are mandatory
// If version is empty, aim for the latest
// When os or arch is empty, it probably does not matter
type BinarySearch struct {
	Owner      string `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Repository string `protobuf:"bytes,2,opt,name=repository" json:"repository,omitempty"`
	Version    string `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
	Os         string `protobuf:"bytes,4,opt,name=os" json:"os,omitempty"`
	Arch       string `protobuf:"bytes,5,opt,name=arch" json:"arch,omitempty"`
}

func (m *BinarySearch) Reset()                    { *m = BinarySearch{} }
func (m *BinarySearch) String() string            { return proto.CompactTextString(m) }
func (*BinarySearch) ProtoMessage()               {}
func (*BinarySearch) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BinarySearch) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *BinarySearch) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

func (m *BinarySearch) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *BinarySearch) GetOs() string {
	if m != nil {
		return m.Os
	}
	return ""
}

func (m *BinarySearch) GetArch() string {
	if m != nil {
		return m.Arch
	}
	return ""
}

// Binary download information
type BinaryDownload struct {
	Url    string                `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Format BinaryDownload_Format `protobuf:"varint,2,opt,name=format,enum=binhq.gitbin.v1alpha1.BinaryDownload_Format" json:"format,omitempty"`
	Path   string                `protobuf:"bytes,3,opt,name=path" json:"path,omitempty"`
}

func (m *BinaryDownload) Reset()                    { *m = BinaryDownload{} }
func (m *BinaryDownload) String() string            { return proto.CompactTextString(m) }
func (*BinaryDownload) ProtoMessage()               {}
func (*BinaryDownload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BinaryDownload) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *BinaryDownload) GetFormat() BinaryDownload_Format {
	if m != nil {
		return m.Format
	}
	return BinaryDownload_BINARY
}

func (m *BinaryDownload) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*BinarySearch)(nil), "binhq.gitbin.v1alpha1.BinarySearch")
	proto.RegisterType((*BinaryDownload)(nil), "binhq.gitbin.v1alpha1.BinaryDownload")
	proto.RegisterEnum("binhq.gitbin.v1alpha1.BinaryDownload_Format", BinaryDownload_Format_name, BinaryDownload_Format_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Gitbin service

type GitbinClient interface {
	// Look for binary download information
	FindBinary(ctx context.Context, in *BinarySearch, opts ...grpc.CallOption) (*BinaryDownload, error)
}

type gitbinClient struct {
	cc *grpc.ClientConn
}

func NewGitbinClient(cc *grpc.ClientConn) GitbinClient {
	return &gitbinClient{cc}
}

func (c *gitbinClient) FindBinary(ctx context.Context, in *BinarySearch, opts ...grpc.CallOption) (*BinaryDownload, error) {
	out := new(BinaryDownload)
	err := grpc.Invoke(ctx, "/binhq.gitbin.v1alpha1.Gitbin/FindBinary", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Gitbin service

type GitbinServer interface {
	// Look for binary download information
	FindBinary(context.Context, *BinarySearch) (*BinaryDownload, error)
}

func RegisterGitbinServer(s *grpc.Server, srv GitbinServer) {
	s.RegisterService(&_Gitbin_serviceDesc, srv)
}

func _Gitbin_FindBinary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BinarySearch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitbinServer).FindBinary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/binhq.gitbin.v1alpha1.Gitbin/FindBinary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitbinServer).FindBinary(ctx, req.(*BinarySearch))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gitbin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "binhq.gitbin.v1alpha1.Gitbin",
	HandlerType: (*GitbinServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindBinary",
			Handler:    _Gitbin_FindBinary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gitbin/v1alpha1/gitbin.proto",
}

func init() { proto.RegisterFile("gitbin/v1alpha1/gitbin.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x51, 0x51, 0x4b, 0xf3, 0x30,
	0x14, 0xfd, 0xda, 0xad, 0xd9, 0xb7, 0x8b, 0x94, 0x72, 0x51, 0x08, 0x22, 0x22, 0x15, 0x61, 0x0f,
	0xd2, 0xb1, 0xf9, 0x0b, 0x56, 0xc6, 0xc6, 0x5e, 0x44, 0xaa, 0x0f, 0xda, 0xb7, 0xd4, 0x55, 0x1b,
	0xa8, 0xb9, 0x35, 0xad, 0x1b, 0x7b, 0xf5, 0xc7, 0xf8, 0x3b, 0xa5, 0x49, 0x07, 0x0a, 0x22, 0xbe,
	0x9d, 0x73, 0xee, 0xb9, 0x39, 0x27, 0x09, 0x9c, 0x3c, 0xcb, 0x26, 0x93, 0x6a, 0xbc, 0x99, 0x88,
	0xb2, 0x2a, 0xc4, 0x64, 0x6c, 0x79, 0x54, 0x69, 0x6a, 0x08, 0x8f, 0x32, 0xa9, 0x8a, 0xd7, 0xa8,
	0xd3, 0xf6, 0x9e, 0xf0, 0xdd, 0x81, 0x83, 0x58, 0x2a, 0xa1, 0x77, 0xb7, 0xb9, 0xd0, 0x8f, 0x05,
	0x1e, 0x82, 0x47, 0x5b, 0x95, 0x6b, 0xee, 0x9c, 0x39, 0xa3, 0x61, 0x62, 0x09, 0x9e, 0x02, 0xe8,
	0xbc, 0xa2, 0x5a, 0x36, 0xa4, 0x77, 0xdc, 0x35, 0xa3, 0x2f, 0x0a, 0x72, 0x18, 0x6c, 0x72, 0x5d,
	0x4b, 0x52, 0xbc, 0x67, 0x86, 0x7b, 0x8a, 0x3e, 0xb8, 0x54, 0xf3, 0xbe, 0x11, 0x5d, 0xaa, 0x11,
	0xa1, 0xdf, 0xe6, 0x70, 0xcf, 0x28, 0x06, 0x87, 0x1f, 0x0e, 0xf8, 0xb6, 0xc4, 0x9c, 0xb6, 0xaa,
	0x24, 0xb1, 0xc6, 0x00, 0x7a, 0x6f, 0xba, 0xec, 0x4a, 0xb4, 0x10, 0xe7, 0xc0, 0x9e, 0x48, 0xbf,
	0x88, 0xc6, 0xc4, 0xfb, 0xd3, 0xcb, 0xe8, 0xc7, 0x1b, 0x45, 0xdf, 0x0f, 0x8a, 0x16, 0x66, 0x27,
	0xe9, 0x76, 0xdb, 0xf8, 0x4a, 0x34, 0x45, 0xd7, 0xd2, 0xe0, 0x70, 0x04, 0xcc, 0xba, 0x10, 0x80,
	0xc5, 0xab, 0xeb, 0x59, 0xf2, 0x10, 0xfc, 0xc3, 0x21, 0x78, 0x77, 0xb3, 0x64, 0x99, 0x06, 0x0e,
	0x0e, 0xa0, 0x97, 0xae, 0x6e, 0x02, 0x77, 0x9a, 0x01, 0x5b, 0x9a, 0x38, 0xbc, 0x07, 0x58, 0x48,
	0xb5, 0xb6, 0x61, 0x78, 0xfe, 0x6b, 0x17, 0xfb, 0xb2, 0xc7, 0x17, 0x7f, 0x2a, 0x1c, 0xff, 0x4f,
	0x99, 0x75, 0x64, 0xcc, 0xfc, 0xdc, 0xd5, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0c, 0x45, 0x8f,
	0x69, 0xd9, 0x01, 0x00, 0x00,
}