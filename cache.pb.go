// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cache.proto

package simplegocache

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type NewCacheMsg struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Ttl                  int64    `protobuf:"varint,2,opt,name=ttl,proto3" json:"ttl,omitempty"`
	MaxElements          int32    `protobuf:"varint,3,opt,name=maxElements,proto3" json:"maxElements,omitempty"`
	MaxElementSize       int32    `protobuf:"varint,4,opt,name=maxElementSize,proto3" json:"maxElementSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewCacheMsg) Reset()         { *m = NewCacheMsg{} }
func (m *NewCacheMsg) String() string { return proto.CompactTextString(m) }
func (*NewCacheMsg) ProtoMessage()    {}
func (*NewCacheMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fca3b110c9bbf3a, []int{0}
}

func (m *NewCacheMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewCacheMsg.Unmarshal(m, b)
}
func (m *NewCacheMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewCacheMsg.Marshal(b, m, deterministic)
}
func (m *NewCacheMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewCacheMsg.Merge(m, src)
}
func (m *NewCacheMsg) XXX_Size() int {
	return xxx_messageInfo_NewCacheMsg.Size(m)
}
func (m *NewCacheMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_NewCacheMsg.DiscardUnknown(m)
}

var xxx_messageInfo_NewCacheMsg proto.InternalMessageInfo

func (m *NewCacheMsg) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *NewCacheMsg) GetTtl() int64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func (m *NewCacheMsg) GetMaxElements() int32 {
	if m != nil {
		return m.MaxElements
	}
	return 0
}

func (m *NewCacheMsg) GetMaxElementSize() int32 {
	if m != nil {
		return m.MaxElementSize
	}
	return 0
}

type EntryMsg struct {
	Content              []byte   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	CacheKey             string   `protobuf:"bytes,3,opt,name=cacheKey,proto3" json:"cacheKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntryMsg) Reset()         { *m = EntryMsg{} }
func (m *EntryMsg) String() string { return proto.CompactTextString(m) }
func (*EntryMsg) ProtoMessage()    {}
func (*EntryMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fca3b110c9bbf3a, []int{1}
}

func (m *EntryMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntryMsg.Unmarshal(m, b)
}
func (m *EntryMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntryMsg.Marshal(b, m, deterministic)
}
func (m *EntryMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntryMsg.Merge(m, src)
}
func (m *EntryMsg) XXX_Size() int {
	return xxx_messageInfo_EntryMsg.Size(m)
}
func (m *EntryMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_EntryMsg.DiscardUnknown(m)
}

var xxx_messageInfo_EntryMsg proto.InternalMessageInfo

func (m *EntryMsg) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *EntryMsg) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *EntryMsg) GetCacheKey() string {
	if m != nil {
		return m.CacheKey
	}
	return ""
}

type CacheMsg struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CacheMsg) Reset()         { *m = CacheMsg{} }
func (m *CacheMsg) String() string { return proto.CompactTextString(m) }
func (*CacheMsg) ProtoMessage()    {}
func (*CacheMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fca3b110c9bbf3a, []int{2}
}

func (m *CacheMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CacheMsg.Unmarshal(m, b)
}
func (m *CacheMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CacheMsg.Marshal(b, m, deterministic)
}
func (m *CacheMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CacheMsg.Merge(m, src)
}
func (m *CacheMsg) XXX_Size() int {
	return xxx_messageInfo_CacheMsg.Size(m)
}
func (m *CacheMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_CacheMsg.DiscardUnknown(m)
}

var xxx_messageInfo_CacheMsg proto.InternalMessageInfo

func (m *CacheMsg) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type Bool struct {
	InCache              bool     `protobuf:"varint,1,opt,name=in_cache,json=inCache,proto3" json:"in_cache,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bool) Reset()         { *m = Bool{} }
func (m *Bool) String() string { return proto.CompactTextString(m) }
func (*Bool) ProtoMessage()    {}
func (*Bool) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fca3b110c9bbf3a, []int{3}
}

func (m *Bool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bool.Unmarshal(m, b)
}
func (m *Bool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bool.Marshal(b, m, deterministic)
}
func (m *Bool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bool.Merge(m, src)
}
func (m *Bool) XXX_Size() int {
	return xxx_messageInfo_Bool.Size(m)
}
func (m *Bool) XXX_DiscardUnknown() {
	xxx_messageInfo_Bool.DiscardUnknown(m)
}

var xxx_messageInfo_Bool proto.InternalMessageInfo

func (m *Bool) GetInCache() bool {
	if m != nil {
		return m.InCache
	}
	return false
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fca3b110c9bbf3a, []int{4}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*NewCacheMsg)(nil), "simplegocache.NewCacheMsg")
	proto.RegisterType((*EntryMsg)(nil), "simplegocache.EntryMsg")
	proto.RegisterType((*CacheMsg)(nil), "simplegocache.CacheMsg")
	proto.RegisterType((*Bool)(nil), "simplegocache.Bool")
	proto.RegisterType((*Empty)(nil), "simplegocache.Empty")
}

func init() { proto.RegisterFile("cache.proto", fileDescriptor_5fca3b110c9bbf3a) }

var fileDescriptor_5fca3b110c9bbf3a = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4f, 0x4b, 0xfb, 0x40,
	0x14, 0x24, 0x4d, 0xd3, 0xa4, 0xaf, 0xbf, 0x9f, 0xc8, 0x2a, 0x18, 0x8b, 0x87, 0x98, 0x83, 0xf4,
	0xd4, 0x83, 0x3d, 0x54, 0xbc, 0x88, 0xd6, 0x1e, 0x44, 0x14, 0x59, 0xf1, 0x2c, 0xb1, 0x7d, 0xc4,
	0xe0, 0xfe, 0x09, 0xdd, 0x15, 0x8d, 0x7e, 0x4c, 0xbf, 0x90, 0xec, 0x4a, 0xda, 0x18, 0x4c, 0xa1,
	0xb7, 0xf7, 0x66, 0x32, 0xbc, 0x99, 0x09, 0x0b, 0xbd, 0x59, 0x32, 0x7b, 0xc6, 0x61, 0xbe, 0x90,
	0x5a, 0x92, 0xff, 0x2a, 0xe3, 0x39, 0xc3, 0x54, 0x5a, 0x30, 0xfe, 0x84, 0xde, 0x2d, 0xbe, 0x4d,
	0xcc, 0x7c, 0xa3, 0x52, 0xb2, 0x0d, 0xee, 0x0b, 0x16, 0xa1, 0x13, 0x39, 0x83, 0x2e, 0x35, 0xa3,
	0x41, 0xb4, 0x66, 0x61, 0x2b, 0x72, 0x06, 0x2e, 0x35, 0x23, 0x89, 0xa0, 0xc7, 0x93, 0xf7, 0x29,
	0x43, 0x8e, 0x42, 0xab, 0xd0, 0x8d, 0x9c, 0x81, 0x47, 0xab, 0x10, 0x39, 0x82, 0xad, 0xd5, 0x7a,
	0x9f, 0x7d, 0x60, 0xd8, 0xb6, 0x1f, 0xd5, 0xd0, 0x98, 0x42, 0x30, 0x15, 0x7a, 0x51, 0x98, 0xcb,
	0x21, 0xf8, 0x33, 0x29, 0x34, 0x0a, 0x6d, 0xaf, 0xff, 0xa3, 0xe5, 0x5a, 0x7a, 0x6a, 0xad, 0x3c,
	0xf5, 0x21, 0xb0, 0xee, 0xaf, 0xb1, 0xb0, 0xe7, 0xbb, 0x74, 0xb9, 0xc7, 0x07, 0x10, 0x54, 0xd3,
	0x70, 0x95, 0x96, 0x69, 0xb8, 0x4a, 0xe3, 0x43, 0x68, 0x5f, 0x48, 0xc9, 0xc8, 0x3e, 0x04, 0x99,
	0x78, 0xb4, 0x22, 0x4b, 0x07, 0xd4, 0xcf, 0x84, 0xd5, 0xc5, 0x3e, 0x78, 0x53, 0x9e, 0xeb, 0xe2,
	0xf8, 0xcb, 0x05, 0xcf, 0x42, 0xe4, 0x0c, 0x82, 0xb2, 0x24, 0xd2, 0x1f, 0xfe, 0x2a, 0x70, 0x58,
	0x69, 0xaf, 0xbf, 0x57, 0xe3, 0x96, 0x46, 0xc6, 0xe0, 0x9e, 0xcf, 0xe7, 0xa4, 0xce, 0x97, 0xe1,
	0xd7, 0x09, 0x3b, 0x97, 0xc8, 0x50, 0x63, 0xb3, 0x76, 0xb7, 0x4e, 0x18, 0xf3, 0xe4, 0x04, 0xda,
	0x14, 0x93, 0x0d, 0x4e, 0x2e, 0x7f, 0xc4, 0x18, 0xfc, 0xab, 0x9f, 0x2a, 0x9a, 0xc5, 0x3b, 0x35,
	0xc2, 0x76, 0x7a, 0x0a, 0x9d, 0x87, 0x7c, 0x9e, 0xac, 0xf3, 0xda, 0x98, 0x73, 0x04, 0xde, 0xdd,
	0xe2, 0x55, 0x20, 0xf9, 0x33, 0x4d, 0x43, 0xc6, 0x11, 0x78, 0x13, 0x26, 0xd5, 0x46, 0xa2, 0xa7,
	0x8e, 0x7d, 0x06, 0xa3, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x6d, 0xb2, 0xc3, 0x15, 0x03,
	0x00, 0x00,
}
