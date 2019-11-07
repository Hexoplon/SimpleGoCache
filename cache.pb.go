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

type Entry struct {
	Content              []byte   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	CacheKey             string   `protobuf:"bytes,3,opt,name=cacheKey,proto3" json:"cacheKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Entry) Reset()         { *m = Entry{} }
func (m *Entry) String() string { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()    {}
func (*Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fca3b110c9bbf3a, []int{1}
}

func (m *Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entry.Unmarshal(m, b)
}
func (m *Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entry.Marshal(b, m, deterministic)
}
func (m *Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entry.Merge(m, src)
}
func (m *Entry) XXX_Size() int {
	return xxx_messageInfo_Entry.Size(m)
}
func (m *Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Entry proto.InternalMessageInfo

func (m *Entry) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Entry) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Entry) GetCacheKey() string {
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
	proto.RegisterType((*Entry)(nil), "simplegocache.Entry")
	proto.RegisterType((*CacheMsg)(nil), "simplegocache.CacheMsg")
	proto.RegisterType((*Bool)(nil), "simplegocache.Bool")
	proto.RegisterType((*Empty)(nil), "simplegocache.Empty")
}

func init() { proto.RegisterFile("cache.proto", fileDescriptor_5fca3b110c9bbf3a) }

var fileDescriptor_5fca3b110c9bbf3a = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x4f, 0x83, 0x40,
	0x10, 0x0d, 0xa5, 0x14, 0x3a, 0x55, 0x63, 0x56, 0x13, 0x91, 0x78, 0x40, 0x0e, 0xa6, 0xa7, 0x1e,
	0xda, 0x26, 0x5e, 0xbc, 0x68, 0xed, 0xc1, 0x18, 0x3f, 0xb2, 0xc6, 0xb3, 0xc1, 0x32, 0x41, 0x22,
	0xec, 0x92, 0xb2, 0x46, 0xd1, 0xff, 0xe6, 0x6f, 0x33, 0x3b, 0x86, 0xb6, 0x62, 0x49, 0xf4, 0x36,
	0xf3, 0x86, 0xc7, 0xfb, 0xc8, 0x42, 0x6f, 0x16, 0xce, 0x9e, 0x70, 0x90, 0xcf, 0xa5, 0x92, 0x6c,
	0xb3, 0x48, 0xb2, 0x3c, 0xc5, 0x58, 0x12, 0x18, 0x7c, 0x40, 0xef, 0x1a, 0x5f, 0x27, 0x7a, 0xbe,
	0x2a, 0x62, 0xb6, 0x0d, 0xe6, 0x33, 0x96, 0xae, 0xe1, 0x1b, 0xfd, 0x2e, 0xd7, 0xa3, 0x46, 0x94,
	0x4a, 0xdd, 0x96, 0x6f, 0xf4, 0x4d, 0xae, 0x47, 0xe6, 0x43, 0x2f, 0x0b, 0xdf, 0xa6, 0x29, 0x66,
	0x28, 0x54, 0xe1, 0x9a, 0xbe, 0xd1, 0xb7, 0xf8, 0x2a, 0xc4, 0x8e, 0x60, 0x6b, 0xb9, 0xde, 0x25,
	0xef, 0xe8, 0xb6, 0xe9, 0xa3, 0x1a, 0x1a, 0xdc, 0x80, 0x35, 0x15, 0x6a, 0x5e, 0x32, 0x17, 0xec,
	0x99, 0x14, 0x0a, 0x85, 0x22, 0xe9, 0x0d, 0x5e, 0xad, 0x95, 0xa1, 0xd6, 0xd2, 0x90, 0x07, 0x0e,
	0x59, 0xbf, 0xc4, 0x92, 0xb4, 0xbb, 0x7c, 0xb1, 0x07, 0x07, 0xe0, 0xac, 0x46, 0xc9, 0x8a, 0xb8,
	0x8a, 0x92, 0x15, 0x71, 0x70, 0x08, 0xed, 0x33, 0x29, 0x53, 0xb6, 0x0f, 0x4e, 0x22, 0x1e, 0x88,
	0x44, 0x67, 0x87, 0xdb, 0x89, 0x20, 0x5e, 0x60, 0x83, 0x35, 0xcd, 0x72, 0x55, 0x0e, 0x3f, 0x4d,
	0xb0, 0x08, 0x62, 0x27, 0xe0, 0x54, 0x0d, 0x31, 0x6f, 0xf0, 0xa3, 0xbd, 0xc1, 0x4a, 0x75, 0xde,
	0x6e, 0xed, 0x46, 0xff, 0x61, 0x63, 0x30, 0x4f, 0xa3, 0x88, 0xfd, 0x3a, 0xea, 0xd8, 0xde, 0x5e,
	0x0d, 0x5d, 0x78, 0x1f, 0x43, 0xe7, 0x1c, 0x53, 0x54, 0xd8, 0x40, 0x5c, 0xaf, 0x35, 0x84, 0x36,
	0xc7, 0x30, 0xfa, 0x2b, 0x87, 0x9a, 0x1f, 0x83, 0x7d, 0xf1, 0x9d, 0xbd, 0x81, 0xb6, 0x53, 0x43,
	0xa9, 0xc1, 0x63, 0xe8, 0xdc, 0xe7, 0x51, 0xd8, 0xe8, 0xaf, 0x31, 0xd8, 0x08, 0xac, 0xdb, 0xf9,
	0x8b, 0x58, 0xc3, 0xd3, 0x09, 0x1a, 0x72, 0x8d, 0xc0, 0x9a, 0xa4, 0xb2, 0xf8, 0x17, 0xe9, 0xb1,
	0x43, 0xcf, 0x7d, 0xf4, 0x15, 0x00, 0x00, 0xff, 0xff, 0xfe, 0xfd, 0x6f, 0x29, 0xfd, 0x02, 0x00,
	0x00,
}
