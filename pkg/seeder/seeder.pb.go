// Copyright 2020 Lennart Espe
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: seeder.proto

package seeder

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type HasRequest struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HasRequest) Reset()         { *m = HasRequest{} }
func (m *HasRequest) String() string { return proto.CompactTextString(m) }
func (*HasRequest) ProtoMessage()    {}
func (*HasRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62eb1d2ab518664f, []int{0}
}

func (m *HasRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HasRequest.Unmarshal(m, b)
}
func (m *HasRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HasRequest.Marshal(b, m, deterministic)
}
func (m *HasRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HasRequest.Merge(m, src)
}
func (m *HasRequest) XXX_Size() int {
	return xxx_messageInfo_HasRequest.Size(m)
}
func (m *HasRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HasRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HasRequest proto.InternalMessageInfo

func (m *HasRequest) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type HasResponse struct {
	Chunks               []uint64 `protobuf:"varint,1,rep,packed,name=chunks,proto3" json:"chunks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HasResponse) Reset()         { *m = HasResponse{} }
func (m *HasResponse) String() string { return proto.CompactTextString(m) }
func (*HasResponse) ProtoMessage()    {}
func (*HasResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_62eb1d2ab518664f, []int{1}
}

func (m *HasResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HasResponse.Unmarshal(m, b)
}
func (m *HasResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HasResponse.Marshal(b, m, deterministic)
}
func (m *HasResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HasResponse.Merge(m, src)
}
func (m *HasResponse) XXX_Size() int {
	return xxx_messageInfo_HasResponse.Size(m)
}
func (m *HasResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HasResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HasResponse proto.InternalMessageInfo

func (m *HasResponse) GetChunks() []uint64 {
	if m != nil {
		return m.Chunks
	}
	return nil
}

type FetchRequest struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Chunk                int64    `protobuf:"varint,2,opt,name=chunk,proto3" json:"chunk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchRequest) Reset()         { *m = FetchRequest{} }
func (m *FetchRequest) String() string { return proto.CompactTextString(m) }
func (*FetchRequest) ProtoMessage()    {}
func (*FetchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62eb1d2ab518664f, []int{2}
}

func (m *FetchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchRequest.Unmarshal(m, b)
}
func (m *FetchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchRequest.Marshal(b, m, deterministic)
}
func (m *FetchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchRequest.Merge(m, src)
}
func (m *FetchRequest) XXX_Size() int {
	return xxx_messageInfo_FetchRequest.Size(m)
}
func (m *FetchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchRequest proto.InternalMessageInfo

func (m *FetchRequest) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *FetchRequest) GetChunk() int64 {
	if m != nil {
		return m.Chunk
	}
	return 0
}

type FetchResponse struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchResponse) Reset()         { *m = FetchResponse{} }
func (m *FetchResponse) String() string { return proto.CompactTextString(m) }
func (*FetchResponse) ProtoMessage()    {}
func (*FetchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_62eb1d2ab518664f, []int{3}
}

func (m *FetchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchResponse.Unmarshal(m, b)
}
func (m *FetchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchResponse.Marshal(b, m, deterministic)
}
func (m *FetchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchResponse.Merge(m, src)
}
func (m *FetchResponse) XXX_Size() int {
	return xxx_messageInfo_FetchResponse.Size(m)
}
func (m *FetchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FetchResponse proto.InternalMessageInfo

func (m *FetchResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*HasRequest)(nil), "seeder.HasRequest")
	proto.RegisterType((*HasResponse)(nil), "seeder.HasResponse")
	proto.RegisterType((*FetchRequest)(nil), "seeder.FetchRequest")
	proto.RegisterType((*FetchResponse)(nil), "seeder.FetchResponse")
}

func init() { proto.RegisterFile("seeder.proto", fileDescriptor_62eb1d2ab518664f) }

var fileDescriptor_62eb1d2ab518664f = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcd, 0x8a, 0x83, 0x30,
	0x10, 0xde, 0xac, 0x9a, 0xc3, 0xac, 0x7b, 0x99, 0x75, 0x17, 0xf1, 0x14, 0xb2, 0x14, 0x3c, 0x49,
	0x69, 0xa1, 0xf4, 0x09, 0x8a, 0xe7, 0xf4, 0x09, 0x52, 0x0d, 0x04, 0x0a, 0x6a, 0x9d, 0xf8, 0xfe,
	0xa5, 0x31, 0xd2, 0x9f, 0x43, 0x6f, 0xf3, 0x4d, 0xbe, 0xbf, 0x0c, 0xa4, 0x64, 0x4c, 0x6b, 0xc6,
	0x6a, 0x18, 0x7b, 0xd7, 0x23, 0x9f, 0x91, 0x14, 0x00, 0xb5, 0x26, 0x65, 0x2e, 0x93, 0x21, 0x87,
	0x08, 0xb1, 0xd5, 0x64, 0x73, 0x26, 0x58, 0x99, 0x2a, 0x3f, 0xcb, 0x15, 0x7c, 0x79, 0x06, 0x0d,
	0x7d, 0x47, 0x06, 0xff, 0x80, 0x37, 0x76, 0xea, 0xce, 0x94, 0x33, 0x11, 0x95, 0xb1, 0x0a, 0x48,
	0xee, 0x21, 0x3d, 0x18, 0xd7, 0xd8, 0x37, 0x56, 0x98, 0x41, 0xe2, 0xd9, 0xf9, 0xa7, 0x60, 0x65,
	0xa4, 0x66, 0x20, 0xff, 0xe1, 0x3b, 0x28, 0x43, 0x04, 0x42, 0xdc, 0x6a, 0xa7, 0x17, 0xe9, 0x6d,
	0xde, 0x8c, 0xc0, 0x8f, 0xbe, 0x31, 0xae, 0x21, 0xaa, 0x35, 0x21, 0x56, 0xe1, 0x3f, 0xf7, 0xfa,
	0xc5, 0xcf, 0xd3, 0x6e, 0x76, 0x93, 0x1f, 0xb8, 0x83, 0xc4, 0x07, 0x60, 0xb6, 0xbc, 0x3f, 0x36,
	0x2d, 0x7e, 0x5f, 0xb6, 0x8b, 0xee, 0xc4, 0xfd, 0xa9, 0xb6, 0xd7, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xbe, 0xbb, 0x05, 0x74, 0x3a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SeederClient is the client API for Seeder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SeederClient interface {
	Has(ctx context.Context, in *HasRequest, opts ...grpc.CallOption) (*HasResponse, error)
	Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error)
}

type seederClient struct {
	cc *grpc.ClientConn
}

func NewSeederClient(cc *grpc.ClientConn) SeederClient {
	return &seederClient{cc}
}

func (c *seederClient) Has(ctx context.Context, in *HasRequest, opts ...grpc.CallOption) (*HasResponse, error) {
	out := new(HasResponse)
	err := c.cc.Invoke(ctx, "/seeder.Seeder/Has", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seederClient) Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error) {
	out := new(FetchResponse)
	err := c.cc.Invoke(ctx, "/seeder.Seeder/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SeederServer is the server API for Seeder service.
type SeederServer interface {
	Has(context.Context, *HasRequest) (*HasResponse, error)
	Fetch(context.Context, *FetchRequest) (*FetchResponse, error)
}

// UnimplementedSeederServer can be embedded to have forward compatible implementations.
type UnimplementedSeederServer struct {
}

func (*UnimplementedSeederServer) Has(ctx context.Context, req *HasRequest) (*HasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Has not implemented")
}
func (*UnimplementedSeederServer) Fetch(ctx context.Context, req *FetchRequest) (*FetchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}

func RegisterSeederServer(s *grpc.Server, srv SeederServer) {
	s.RegisterService(&_Seeder_serviceDesc, srv)
}

func _Seeder_Has_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeederServer).Has(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seeder.Seeder/Has",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeederServer).Has(ctx, req.(*HasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Seeder_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeederServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seeder.Seeder/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeederServer).Fetch(ctx, req.(*FetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Seeder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "seeder.Seeder",
	HandlerType: (*SeederServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Has",
			Handler:    _Seeder_Has_Handler,
		},
		{
			MethodName: "Fetch",
			Handler:    _Seeder_Fetch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "seeder.proto",
}
