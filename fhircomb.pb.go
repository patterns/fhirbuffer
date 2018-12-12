// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fhircomb.proto

package fhircomb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

// A search criteria to request the resource record.
type Search struct {
	// A ID is the UUID for the Patient record
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The resource type
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Search) Reset()         { *m = Search{} }
func (m *Search) String() string { return proto.CompactTextString(m) }
func (*Search) ProtoMessage()    {}
func (*Search) Descriptor() ([]byte, []int) {
	return fileDescriptor_a81805fb82f6acbf, []int{0}
}

func (m *Search) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Search.Unmarshal(m, b)
}
func (m *Search) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Search.Marshal(b, m, deterministic)
}
func (m *Search) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Search.Merge(m, src)
}
func (m *Search) XXX_Size() int {
	return xxx_messageInfo_Search.Size(m)
}
func (m *Search) XXX_DiscardUnknown() {
	xxx_messageInfo_Search.DiscardUnknown(m)
}

var xxx_messageInfo_Search proto.InternalMessageInfo

func (m *Search) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Search) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

// A resource record returned from the data store.
type Record struct {
	Resource             []byte   `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_a81805fb82f6acbf, []int{1}
}

func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetResource() []byte {
	if m != nil {
		return m.Resource
	}
	return nil
}

func init() {
	proto.RegisterType((*Search)(nil), "fhircomb.Search")
	proto.RegisterType((*Record)(nil), "fhircomb.Record")
}

func init() { proto.RegisterFile("fhircomb.proto", fileDescriptor_a81805fb82f6acbf) }

var fileDescriptor_a81805fb82f6acbf = []byte{
	// 146 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xcb, 0xc8, 0x2c,
	0x4a, 0xce, 0xcf, 0x4d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0x74,
	0xb8, 0xd8, 0x82, 0x53, 0x13, 0x8b, 0x92, 0x33, 0x84, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18,
	0x15, 0x18, 0x35, 0x38, 0x83, 0x98, 0x32, 0x53, 0x84, 0x84, 0xb8, 0x58, 0x4a, 0x2a, 0x0b, 0x52,
	0x25, 0x98, 0xc0, 0x22, 0x60, 0xb6, 0x92, 0x0a, 0x17, 0x5b, 0x50, 0x6a, 0x72, 0x7e, 0x51, 0x8a,
	0x90, 0x14, 0x17, 0x47, 0x51, 0x6a, 0x71, 0x7e, 0x69, 0x51, 0x72, 0x2a, 0x58, 0x0f, 0x4f, 0x10,
	0x9c, 0x6f, 0x64, 0xc1, 0xc5, 0xe1, 0x06, 0x35, 0x5f, 0x48, 0x87, 0x8b, 0x25, 0x28, 0x35, 0x31,
	0x45, 0x48, 0x40, 0x0f, 0xee, 0x04, 0x88, 0x7d, 0x52, 0x48, 0x22, 0x10, 0x33, 0x95, 0x18, 0x92,
	0xd8, 0xc0, 0xce, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x94, 0x3d, 0xc7, 0xb0, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FhircombClient is the client API for Fhircomb service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FhircombClient interface {
	// Obtains the resource record that matches the search criteria.
	Read(ctx context.Context, in *Search, opts ...grpc.CallOption) (*Record, error)
}

type fhircombClient struct {
	cc *grpc.ClientConn
}

func NewFhircombClient(cc *grpc.ClientConn) FhircombClient {
	return &fhircombClient{cc}
}

func (c *fhircombClient) Read(ctx context.Context, in *Search, opts ...grpc.CallOption) (*Record, error) {
	out := new(Record)
	err := c.cc.Invoke(ctx, "/fhircomb.Fhircomb/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FhircombServer is the server API for Fhircomb service.
type FhircombServer interface {
	// Obtains the resource record that matches the search criteria.
	Read(context.Context, *Search) (*Record, error)
}

func RegisterFhircombServer(s *grpc.Server, srv FhircombServer) {
	s.RegisterService(&_Fhircomb_serviceDesc, srv)
}

func _Fhircomb_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Search)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FhircombServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fhircomb.Fhircomb/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FhircombServer).Read(ctx, req.(*Search))
	}
	return interceptor(ctx, in, info, handler)
}

var _Fhircomb_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fhircomb.Fhircomb",
	HandlerType: (*FhircombServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _Fhircomb_Read_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fhircomb.proto",
}