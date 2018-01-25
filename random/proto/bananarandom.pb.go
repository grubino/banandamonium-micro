// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bananarandom.proto

/*
Package bananarandom is a generated protocol buffer package.

It is generated from these files:
	bananarandom.proto

It has these top-level messages:
	RollRequest
	RollResponse
*/
package bananarandom

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RollRequest struct {
	GameId   int32 `protobuf:"varint,1,opt,name=gameId" json:"gameId,omitempty"`
	PlayerId int32 `protobuf:"varint,2,opt,name=playerId" json:"playerId,omitempty"`
}

func (m *RollRequest) Reset()                    { *m = RollRequest{} }
func (m *RollRequest) String() string            { return proto.CompactTextString(m) }
func (*RollRequest) ProtoMessage()               {}
func (*RollRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RollRequest) GetGameId() int32 {
	if m != nil {
		return m.GameId
	}
	return 0
}

func (m *RollRequest) GetPlayerId() int32 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

type RollResponse struct {
	Results []int32 `protobuf:"varint,1,rep,packed,name=results" json:"results,omitempty"`
}

func (m *RollResponse) Reset()                    { *m = RollResponse{} }
func (m *RollResponse) String() string            { return proto.CompactTextString(m) }
func (*RollResponse) ProtoMessage()               {}
func (*RollResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RollResponse) GetResults() []int32 {
	if m != nil {
		return m.Results
	}
	return nil
}

func init() {
	proto.RegisterType((*RollRequest)(nil), "RollRequest")
	proto.RegisterType((*RollResponse)(nil), "RollResponse")
}

func init() { proto.RegisterFile("bananarandom.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0x4a, 0xcc, 0x4b,
	0xcc, 0x4b, 0x2c, 0x4a, 0xcc, 0x4b, 0xc9, 0xcf, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x72,
	0xe4, 0xe2, 0x0e, 0xca, 0xcf, 0xc9, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe3,
	0x62, 0x4b, 0x4f, 0xcc, 0x4d, 0xf5, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x82, 0xf2,
	0x84, 0xa4, 0xb8, 0x38, 0x0a, 0x72, 0x12, 0x2b, 0x53, 0x8b, 0x3c, 0x53, 0x24, 0x98, 0xc0, 0x32,
	0x70, 0xbe, 0x92, 0x06, 0x17, 0x0f, 0xc4, 0x88, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x09,
	0x2e, 0xf6, 0xa2, 0xd4, 0xe2, 0xd2, 0x9c, 0x92, 0x62, 0x09, 0x46, 0x05, 0x66, 0x0d, 0xd6, 0x20,
	0x18, 0xd7, 0xc8, 0x92, 0x8b, 0xc7, 0x09, 0xec, 0x84, 0x20, 0xb0, 0x13, 0x84, 0x34, 0xb9, 0x38,
	0x40, 0x3a, 0x5d, 0x32, 0x93, 0x53, 0x85, 0x78, 0xf4, 0x90, 0xdc, 0x21, 0xc5, 0xab, 0x87, 0x6c,
	0xa4, 0x12, 0x43, 0x12, 0x1b, 0xd8, 0xb9, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x12,
	0xd8, 0xd5, 0xc4, 0x00, 0x00, 0x00,
}
