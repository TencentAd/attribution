// Code generated by protoc-gen-go. DO NOT EDIT.
// source: attribution/proto/conv/conv.proto

package conv

import (
	click "attribution/proto/click"
	user "attribution/proto/user"
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

type ConversionLog struct {
	UserData             *user.UserData  `protobuf:"bytes,1,opt,name=user_data,json=userData,proto3" json:"user_data,omitempty"`
	EventTime            int64           `protobuf:"varint,2,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	AppId                string          `protobuf:"bytes,3,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	ConvId               string          `protobuf:"bytes,4,opt,name=conv_id,json=convId,proto3" json:"conv_id,omitempty"`
	MatchClick           *click.ClickLog `protobuf:"bytes,1000,opt,name=match_click,json=matchClick,proto3" json:"match_click,omitempty"`
	OriginalContent      string          `protobuf:"bytes,1001,opt,name=original_content,json=originalContent,proto3" json:"original_content,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ConversionLog) Reset()         { *m = ConversionLog{} }
func (m *ConversionLog) String() string { return proto.CompactTextString(m) }
func (*ConversionLog) ProtoMessage()    {}
func (*ConversionLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_e165194b454c0299, []int{0}
}

func (m *ConversionLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversionLog.Unmarshal(m, b)
}
func (m *ConversionLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversionLog.Marshal(b, m, deterministic)
}
func (m *ConversionLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversionLog.Merge(m, src)
}
func (m *ConversionLog) XXX_Size() int {
	return xxx_messageInfo_ConversionLog.Size(m)
}
func (m *ConversionLog) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversionLog.DiscardUnknown(m)
}

var xxx_messageInfo_ConversionLog proto.InternalMessageInfo

func (m *ConversionLog) GetUserData() *user.UserData {
	if m != nil {
		return m.UserData
	}
	return nil
}

func (m *ConversionLog) GetEventTime() int64 {
	if m != nil {
		return m.EventTime
	}
	return 0
}

func (m *ConversionLog) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *ConversionLog) GetConvId() string {
	if m != nil {
		return m.ConvId
	}
	return ""
}

func (m *ConversionLog) GetMatchClick() *click.ClickLog {
	if m != nil {
		return m.MatchClick
	}
	return nil
}

func (m *ConversionLog) GetOriginalContent() string {
	if m != nil {
		return m.OriginalContent
	}
	return ""
}

func init() {
	proto.RegisterType((*ConversionLog)(nil), "conv.ConversionLog")
}

func init() { proto.RegisterFile("attribution/proto/conv/conv.proto", fileDescriptor_e165194b454c0299) }

var fileDescriptor_e165194b454c0299 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0xc1, 0x4a, 0xc3, 0x40,
	0x14, 0x24, 0xb6, 0x36, 0x66, 0x8b, 0x56, 0x16, 0xc4, 0x50, 0x10, 0x62, 0xf5, 0x10, 0x14, 0xb6,
	0xa8, 0x7f, 0x60, 0xbc, 0x14, 0x7a, 0x0a, 0x7a, 0x0e, 0xdb, 0xcd, 0x12, 0x17, 0x9b, 0x7d, 0x61,
	0xf3, 0x92, 0x6f, 0x56, 0xbf, 0x42, 0xf6, 0xad, 0x82, 0x87, 0x5e, 0x86, 0x99, 0xc9, 0x30, 0x99,
	0x7d, 0xec, 0x5a, 0x22, 0x3a, 0xb3, 0x1b, 0xd0, 0x80, 0x5d, 0x77, 0x0e, 0x10, 0xd6, 0x0a, 0xec,
	0x48, 0x20, 0x48, 0xf3, 0xa9, 0xe7, 0xcb, 0x03, 0xc1, 0xa1, 0xd7, 0x8e, 0x20, 0x04, 0x97, 0x37,
	0x07, 0xba, 0xf6, 0x46, 0x7d, 0x04, 0x0c, 0xa1, 0xd5, 0x77, 0xc4, 0x4e, 0x0b, 0xb0, 0xa3, 0x76,
	0xbd, 0x01, 0xbb, 0x85, 0x86, 0xdf, 0xb3, 0xc4, 0x97, 0x54, 0xb5, 0x44, 0x99, 0x46, 0x59, 0x94,
	0xcf, 0x1f, 0xcf, 0x04, 0xd5, 0xbe, 0xf5, 0xda, 0xbd, 0x48, 0x94, 0xe5, 0xc9, 0xf0, 0xcb, 0xf8,
	0x15, 0x63, 0x7a, 0xd4, 0x16, 0x2b, 0x34, 0xad, 0x4e, 0x8f, 0xb2, 0x28, 0x9f, 0x94, 0x09, 0x39,
	0xaf, 0xa6, 0xd5, 0xfc, 0x82, 0xcd, 0x64, 0xd7, 0x55, 0xa6, 0x4e, 0x27, 0x59, 0x94, 0x27, 0xe5,
	0xb1, 0xec, 0xba, 0x4d, 0xcd, 0x2f, 0x59, 0xec, 0x1f, 0xe1, 0xfd, 0x29, 0xf9, 0x33, 0x2f, 0x37,
	0x35, 0x7f, 0x60, 0xf3, 0x56, 0xa2, 0x7a, 0xaf, 0x68, 0x62, 0xfa, 0x19, 0xd3, 0xef, 0x17, 0x22,
	0x2c, 0x2e, 0x3c, 0x6e, 0xa1, 0x29, 0x19, 0x85, 0x48, 0xf2, 0x3b, 0x76, 0x0e, 0xce, 0x34, 0xc6,
	0xca, 0x7d, 0xa5, 0xc0, 0xa2, 0xb6, 0x98, 0x7e, 0xc5, 0xd4, 0xba, 0xf8, 0xfb, 0x50, 0x04, 0xff,
	0xf9, 0x96, 0xad, 0x14, 0xb4, 0x02, 0xb5, 0x55, 0xda, 0xa2, 0xf8, 0x77, 0x9f, 0x70, 0x0d, 0xe1,
	0x67, 0xec, 0x66, 0xc4, 0x9f, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xfa, 0xc0, 0x2b, 0x38, 0x8c,
	0x01, 0x00, 0x00,
}
