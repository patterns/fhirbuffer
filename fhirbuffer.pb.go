// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fhirbuffer.proto

package fhirbuffer

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

// A search criteria to request the healthcare resource.
type Search struct {
	// A ID is the UUID of the record
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
	return fileDescriptor_ffd338a9c98fa409, []int{0}
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

// A modification to change the healthcare resource.
type Change struct {
	Resource             []byte   `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Change) Reset()         { *m = Change{} }
func (m *Change) String() string { return proto.CompactTextString(m) }
func (*Change) ProtoMessage()    {}
func (*Change) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffd338a9c98fa409, []int{1}
}

func (m *Change) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Change.Unmarshal(m, b)
}
func (m *Change) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Change.Marshal(b, m, deterministic)
}
func (m *Change) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Change.Merge(m, src)
}
func (m *Change) XXX_Size() int {
	return xxx_messageInfo_Change.Size(m)
}
func (m *Change) XXX_DiscardUnknown() {
	xxx_messageInfo_Change.DiscardUnknown(m)
}

var xxx_messageInfo_Change proto.InternalMessageInfo

func (m *Change) GetResource() []byte {
	if m != nil {
		return m.Resource
	}
	return nil
}

// A healthcare resource returned from the data store.
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
	return fileDescriptor_ffd338a9c98fa409, []int{2}
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
	proto.RegisterType((*Search)(nil), "fhirbuffer.Search")
	proto.RegisterType((*Change)(nil), "fhirbuffer.Change")
	proto.RegisterType((*Record)(nil), "fhirbuffer.Record")
}

func init() { proto.RegisterFile("fhirbuffer.proto", fileDescriptor_ffd338a9c98fa409) }

var fileDescriptor_ffd338a9c98fa409 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0xcb, 0xc8, 0x2c,
	0x4a, 0x2a, 0x4d, 0x4b, 0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x28, 0xe9, 0x70, 0xb1, 0x05, 0xa7, 0x26, 0x16, 0x25, 0x67, 0x08, 0xf1, 0x71, 0x31, 0x65, 0xa6,
	0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31, 0x65, 0xa6, 0x08, 0x09, 0x71, 0xb1, 0x94, 0x54,
	0x16, 0xa4, 0x4a, 0x30, 0x81, 0x45, 0xc0, 0x6c, 0x25, 0x15, 0x2e, 0x36, 0xe7, 0x8c, 0xc4, 0xbc,
	0xf4, 0x54, 0x21, 0x29, 0x2e, 0x8e, 0xa2, 0xd4, 0xe2, 0xfc, 0xd2, 0xa2, 0xe4, 0x54, 0xb0, 0x1e,
	0x9e, 0x20, 0x38, 0x1f, 0xa4, 0x2a, 0x28, 0x35, 0x39, 0xbf, 0x28, 0x05, 0x9f, 0x2a, 0xa3, 0x22,
	0x2e, 0x2e, 0x37, 0xb8, 0x3b, 0x84, 0x0c, 0xb8, 0x58, 0x82, 0x52, 0x13, 0x53, 0x84, 0x84, 0xf4,
	0x90, 0x9c, 0x0b, 0x71, 0x99, 0x14, 0x8a, 0x18, 0xc4, 0x64, 0x25, 0x06, 0x21, 0x23, 0x2e, 0xb6,
	0xd0, 0x82, 0x94, 0xc4, 0x92, 0x54, 0x54, 0x3d, 0x10, 0xf7, 0x61, 0xd7, 0x93, 0xc4, 0x06, 0x0e,
	0x00, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4b, 0xaa, 0x3c, 0x01, 0x14, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FhirbufferClient is the client API for Fhirbuffer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FhirbufferClient interface {
	// Obtains the healthcare resource that matches the search criteria.
	Read(ctx context.Context, in *Search, opts ...grpc.CallOption) (*Record, error)
	// Modifies the healthcare resource
	Update(ctx context.Context, in *Change, opts ...grpc.CallOption) (*Record, error)
}

type fhirbufferClient struct {
	cc *grpc.ClientConn
}

func NewFhirbufferClient(cc *grpc.ClientConn) FhirbufferClient {
	return &fhirbufferClient{cc}
}

func (c *fhirbufferClient) Read(ctx context.Context, in *Search, opts ...grpc.CallOption) (*Record, error) {
	out := new(Record)
	err := c.cc.Invoke(ctx, "/fhirbuffer.Fhirbuffer/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fhirbufferClient) Update(ctx context.Context, in *Change, opts ...grpc.CallOption) (*Record, error) {
	out := new(Record)
	err := c.cc.Invoke(ctx, "/fhirbuffer.Fhirbuffer/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FhirbufferServer is the server API for Fhirbuffer service.
type FhirbufferServer interface {
	// Obtains the healthcare resource that matches the search criteria.
	Read(context.Context, *Search) (*Record, error)
	// Modifies the healthcare resource
	Update(context.Context, *Change) (*Record, error)
}

func RegisterFhirbufferServer(s *grpc.Server, srv FhirbufferServer) {
	s.RegisterService(&_Fhirbuffer_serviceDesc, srv)
}

func _Fhirbuffer_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Search)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FhirbufferServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fhirbuffer.Fhirbuffer/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FhirbufferServer).Read(ctx, req.(*Search))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fhirbuffer_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Change)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FhirbufferServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fhirbuffer.Fhirbuffer/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FhirbufferServer).Update(ctx, req.(*Change))
	}
	return interceptor(ctx, in, info, handler)
}

var _Fhirbuffer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fhirbuffer.Fhirbuffer",
	HandlerType: (*FhirbufferServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _Fhirbuffer_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Fhirbuffer_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fhirbuffer.proto",
}
