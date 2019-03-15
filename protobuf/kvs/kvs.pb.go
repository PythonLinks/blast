// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/kvs/kvs.proto

package kvs

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	raft "github.com/mosuka/blast/protobuf/raft"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetRequest struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e9909cfc2f34163, []int{0}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type GetResponse struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e9909cfc2f34163, []int{1}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type PutRequest struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutRequest) Reset()         { *m = PutRequest{} }
func (m *PutRequest) String() string { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()    {}
func (*PutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e9909cfc2f34163, []int{2}
}

func (m *PutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutRequest.Unmarshal(m, b)
}
func (m *PutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutRequest.Marshal(b, m, deterministic)
}
func (m *PutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutRequest.Merge(m, src)
}
func (m *PutRequest) XXX_Size() int {
	return xxx_messageInfo_PutRequest.Size(m)
}
func (m *PutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutRequest proto.InternalMessageInfo

func (m *PutRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *PutRequest) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type DeleteRequest struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e9909cfc2f34163, []int{3}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type KeyValuePair struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyValuePair) Reset()         { *m = KeyValuePair{} }
func (m *KeyValuePair) String() string { return proto.CompactTextString(m) }
func (*KeyValuePair) ProtoMessage()    {}
func (*KeyValuePair) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e9909cfc2f34163, []int{4}
}

func (m *KeyValuePair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyValuePair.Unmarshal(m, b)
}
func (m *KeyValuePair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyValuePair.Marshal(b, m, deterministic)
}
func (m *KeyValuePair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValuePair.Merge(m, src)
}
func (m *KeyValuePair) XXX_Size() int {
	return xxx_messageInfo_KeyValuePair.Size(m)
}
func (m *KeyValuePair) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValuePair.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValuePair proto.InternalMessageInfo

func (m *KeyValuePair) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KeyValuePair) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type KVSCommand struct {
	Op                   string   `protobuf:"bytes,1,opt,name=op,proto3" json:"op,omitempty"`
	Key                  []byte   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KVSCommand) Reset()         { *m = KVSCommand{} }
func (m *KVSCommand) String() string { return proto.CompactTextString(m) }
func (*KVSCommand) ProtoMessage()    {}
func (*KVSCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e9909cfc2f34163, []int{5}
}

func (m *KVSCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KVSCommand.Unmarshal(m, b)
}
func (m *KVSCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KVSCommand.Marshal(b, m, deterministic)
}
func (m *KVSCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVSCommand.Merge(m, src)
}
func (m *KVSCommand) XXX_Size() int {
	return xxx_messageInfo_KVSCommand.Size(m)
}
func (m *KVSCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_KVSCommand.DiscardUnknown(m)
}

var xxx_messageInfo_KVSCommand proto.InternalMessageInfo

func (m *KVSCommand) GetOp() string {
	if m != nil {
		return m.Op
	}
	return ""
}

func (m *KVSCommand) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KVSCommand) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "kvs.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "kvs.GetResponse")
	proto.RegisterType((*PutRequest)(nil), "kvs.PutRequest")
	proto.RegisterType((*DeleteRequest)(nil), "kvs.DeleteRequest")
	proto.RegisterType((*KeyValuePair)(nil), "kvs.KeyValuePair")
	proto.RegisterType((*KVSCommand)(nil), "kvs.KVSCommand")
}

func init() { proto.RegisterFile("protobuf/kvs/kvs.proto", fileDescriptor_6e9909cfc2f34163) }

var fileDescriptor_6e9909cfc2f34163 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0xdb, 0xe4, 0x6d, 0x79, 0xdf, 0x79, 0xab, 0xd6, 0x41, 0x4a, 0x89, 0x20, 0x1a, 0x45,
	0xc4, 0x43, 0x22, 0x56, 0x7b, 0xf2, 0xa4, 0x95, 0x82, 0xf5, 0x50, 0x5a, 0xe8, 0xc1, 0xdb, 0x46,
	0xa7, 0x7f, 0xc8, 0x9f, 0x8d, 0xdd, 0x4d, 0xa0, 0x1f, 0xd0, 0xef, 0x25, 0xd9, 0x6d, 0x9b, 0x0a,
	0xa6, 0x7a, 0x48, 0xd8, 0x99, 0x7d, 0x7e, 0x4f, 0xc2, 0x3c, 0x03, 0x8d, 0x78, 0xce, 0x25, 0xf7,
	0x92, 0xb1, 0xeb, 0xa7, 0x22, 0x7b, 0x1c, 0xd5, 0x40, 0xd3, 0x4f, 0x85, 0x75, 0x38, 0xe1, 0x7c,
	0x12, 0x90, 0xbb, 0xd6, 0x50, 0x18, 0xcb, 0x85, 0x56, 0x58, 0xcd, 0x75, 0x77, 0xce, 0xc6, 0x52,
	0xbd, 0xf4, 0x8d, 0x7d, 0x04, 0xd0, 0x25, 0x39, 0xa0, 0xf7, 0x84, 0x84, 0xc4, 0x3a, 0x98, 0x3e,
	0x2d, 0x9a, 0xe5, 0xe3, 0xf2, 0x45, 0x6d, 0x90, 0x1d, 0xed, 0x53, 0xf8, 0xaf, 0xee, 0x45, 0xcc,
	0x23, 0x41, 0x78, 0x00, 0x95, 0x94, 0x05, 0x09, 0x2d, 0x25, 0xba, 0xb0, 0x6f, 0x00, 0xfa, 0x49,
	0xb1, 0x49, 0x4e, 0x19, 0x9b, 0xd4, 0x09, 0xec, 0x74, 0x28, 0x20, 0x49, 0xc5, 0x5f, 0x6f, 0x43,
	0xad, 0x47, 0x8b, 0x51, 0x26, 0xef, 0xb3, 0xd9, 0xfc, 0xd7, 0xd6, 0x1d, 0x80, 0xde, 0x68, 0xf8,
	0xc0, 0xc3, 0x90, 0x45, 0x6f, 0xb8, 0x0b, 0x06, 0x8f, 0x15, 0xf4, 0x6f, 0x60, 0xf0, 0x78, 0xe5,
	0x62, 0x7c, 0xe3, 0x62, 0x6e, 0xb8, 0x5c, 0x7f, 0x18, 0x60, 0xf6, 0x46, 0x43, 0x6c, 0xc1, 0x9f,
	0x27, 0x3e, 0x8b, 0x70, 0xdf, 0x51, 0x83, 0xcb, 0xce, 0xcb, 0x5f, 0xb6, 0x1a, 0x8e, 0x1e, 0xbb,
	0xb3, 0x1a, 0xb0, 0xf3, 0x98, 0x8d, 0xdd, 0x2e, 0xe1, 0x2d, 0x54, 0x9e, 0x89, 0xa5, 0x84, 0xa8,
	0x29, 0x55, 0xfc, 0x8c, 0xdd, 0xc1, 0xdf, 0x61, 0xc4, 0x62, 0x31, 0xe5, 0x12, 0x0b, 0x54, 0x5b,
	0xe8, 0x4b, 0x30, 0xbb, 0x24, 0x71, 0xcf, 0xc9, 0x96, 0x23, 0xcf, 0xd5, 0xaa, 0xe7, 0x0d, 0x1d,
	0xa4, 0x5d, 0xc2, 0x2b, 0x30, 0xfb, 0xc9, 0x4a, 0x9b, 0xc7, 0xb7, 0xc5, 0xbd, 0x0d, 0x55, 0x1d,
	0x18, 0xa2, 0x82, 0xbe, 0xa4, 0x57, 0xcc, 0xdd, 0x9f, 0xbf, 0x9c, 0x4d, 0x66, 0x72, 0x9a, 0x78,
	0xce, 0x2b, 0x0f, 0xdd, 0x90, 0x8b, 0xc4, 0x67, 0xae, 0x17, 0x30, 0x21, 0xdd, 0xcd, 0x8d, 0xf6,
	0xaa, 0xaa, 0x6a, 0x7d, 0x06, 0x00, 0x00, 0xff, 0xff, 0x20, 0xc5, 0x69, 0x1c, 0xe8, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KVSClient is the client API for KVS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KVSClient interface {
	Join(ctx context.Context, in *raft.JoinRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Leave(ctx context.Context, in *raft.LeaveRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Snapshot(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type kVSClient struct {
	cc *grpc.ClientConn
}

func NewKVSClient(cc *grpc.ClientConn) KVSClient {
	return &kVSClient{cc}
}

func (c *kVSClient) Join(ctx context.Context, in *raft.JoinRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/kvs.KVS/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVSClient) Leave(ctx context.Context, in *raft.LeaveRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/kvs.KVS/Leave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVSClient) Snapshot(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/kvs.KVS/Snapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVSClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/kvs.KVS/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVSClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/kvs.KVS/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVSClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/kvs.KVS/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KVSServer is the server API for KVS service.
type KVSServer interface {
	Join(context.Context, *raft.JoinRequest) (*empty.Empty, error)
	Leave(context.Context, *raft.LeaveRequest) (*empty.Empty, error)
	Snapshot(context.Context, *empty.Empty) (*empty.Empty, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Put(context.Context, *PutRequest) (*empty.Empty, error)
	Delete(context.Context, *DeleteRequest) (*empty.Empty, error)
}

func RegisterKVSServer(s *grpc.Server, srv KVSServer) {
	s.RegisterService(&_KVS_serviceDesc, srv)
}

func _KVS_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(raft.JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVSServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvs.KVS/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVSServer).Join(ctx, req.(*raft.JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVS_Leave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(raft.LeaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVSServer).Leave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvs.KVS/Leave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVSServer).Leave(ctx, req.(*raft.LeaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVS_Snapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVSServer).Snapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvs.KVS/Snapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVSServer).Snapshot(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVS_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVSServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvs.KVS/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVSServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVS_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVSServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvs.KVS/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVSServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVS_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVSServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvs.KVS/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVSServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KVS_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kvs.KVS",
	HandlerType: (*KVSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Join",
			Handler:    _KVS_Join_Handler,
		},
		{
			MethodName: "Leave",
			Handler:    _KVS_Leave_Handler,
		},
		{
			MethodName: "Snapshot",
			Handler:    _KVS_Snapshot_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _KVS_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _KVS_Put_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _KVS_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/kvs/kvs.proto",
}