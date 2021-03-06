// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pixeldothorse.proto

/*
Package pixeldothorse is a generated protocol buffer package.

pixeldothorse is the public, user-facing API for the Pixel.Horse game.

It is generated from these files:
	pixeldothorse.proto

It has these top-level messages:
	Nil
*/
package pixeldothorse

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

type Nil struct {
}

func (m *Nil) Reset()                    { *m = Nil{} }
func (m *Nil) String() string            { return proto.CompactTextString(m) }
func (*Nil) ProtoMessage()               {}
func (*Nil) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Nil)(nil), "io.pixeldothorse.Nil")
}

func init() { proto.RegisterFile("pixeldothorse.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 101 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0xc8, 0xac, 0x48,
	0xcd, 0x49, 0xc9, 0x2f, 0xc9, 0xc8, 0x2f, 0x2a, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0xc8, 0xcc, 0xd7, 0x43, 0x11, 0x57, 0x62, 0xe5, 0x62, 0xf6, 0xcb, 0xcc, 0x31, 0xb2, 0xe7,
	0x62, 0x09, 0xc8, 0xcc, 0x4b, 0x17, 0x32, 0xe7, 0x62, 0xf7, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f,
	0x15, 0x12, 0xd5, 0x43, 0x57, 0xac, 0xe7, 0x97, 0x99, 0x23, 0x85, 0x5d, 0xd8, 0x89, 0x3f, 0x8a,
	0x17, 0x45, 0x30, 0x89, 0x0d, 0x6c, 0xa3, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xdf, 0xb4,
	0x43, 0x88, 0x00, 0x00, 0x00,
}
