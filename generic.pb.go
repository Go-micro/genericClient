// Code generated by protoc-gen-go.
// source: generic.proto
// DO NOT EDIT!

/*
Package GenericClient is a generated protocol buffer package.

It is generated from these files:
	generic.proto

It has these top-level messages:
	Request
	Response
*/
package GenericClient

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

type Request struct {
	Body []byte `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type Response struct {
	Body []byte `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "GenericClient.Request")
	proto.RegisterType((*Response)(nil), "GenericClient.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Generic service

type GenericClient interface {
	Printify(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type genericClient struct {
	cc *grpc.ClientConn
}

func NewGenericClient(cc *grpc.ClientConn) GenericClient {
	return &genericClient{cc}
}

func (c *genericClient) Printify(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/GenericClient.Generic/Printify", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Generic service

type GenericServer interface {
	Printify(context.Context, *Request) (*Response, error)
}

func RegisterGenericServer(s *grpc.Server, srv GenericServer) {
	s.RegisterService(&_Generic_serviceDesc, srv)
}

func _Generic_Printify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenericServer).Printify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GenericClient.Generic/Printify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenericServer).Printify(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Generic_serviceDesc = grpc.ServiceDesc{
	ServiceName: "GenericClient.Generic",
	HandlerType: (*GenericServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Printify",
			Handler:    _Generic_Printify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "generic.proto",
}

func init() { proto.RegisterFile("generic.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 132 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4f, 0xcd, 0x4b,
	0x2d, 0xca, 0x4c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x75, 0x87, 0x70, 0x9d, 0x73,
	0x32, 0x53, 0xf3, 0x4a, 0x94, 0x64, 0xb9, 0xd8, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84,
	0x84, 0xb8, 0x58, 0x92, 0xf2, 0x53, 0x2a, 0x25, 0x18, 0x15, 0x18, 0x35, 0x78, 0x82, 0xc0, 0x6c,
	0x25, 0x39, 0x2e, 0x8e, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x6c, 0xf2, 0x46, 0x1e,
	0x5c, 0xec, 0x50, 0xf3, 0x84, 0x6c, 0xb9, 0x38, 0x02, 0x8a, 0x32, 0xf3, 0x4a, 0x32, 0xd3, 0x2a,
	0x85, 0xc4, 0xf4, 0x50, 0x6c, 0xd1, 0x83, 0x5a, 0x21, 0x25, 0x8e, 0x21, 0x0e, 0x31, 0x5b, 0x89,
	0x21, 0x89, 0x0d, 0xec, 0x3c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdf, 0xe5, 0xd5, 0x42,
	0xaf, 0x00, 0x00, 0x00,
}
