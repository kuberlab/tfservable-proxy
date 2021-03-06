// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/util/memmapped_file_system.proto

package util

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A message that describes one region of memmapped file.
type MemmappedFileSystemDirectoryElement struct {
	Offset uint64 `protobuf:"varint,1,opt,name=offset" json:"offset,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *MemmappedFileSystemDirectoryElement) Reset()         { *m = MemmappedFileSystemDirectoryElement{} }
func (m *MemmappedFileSystemDirectoryElement) String() string { return proto.CompactTextString(m) }
func (*MemmappedFileSystemDirectoryElement) ProtoMessage()    {}
func (*MemmappedFileSystemDirectoryElement) Descriptor() ([]byte, []int) {
	return fileDescriptor2, []int{0}
}

func (m *MemmappedFileSystemDirectoryElement) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *MemmappedFileSystemDirectoryElement) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// A directory of regions in a memmapped file.
type MemmappedFileSystemDirectory struct {
	Element []*MemmappedFileSystemDirectoryElement `protobuf:"bytes,1,rep,name=element" json:"element,omitempty"`
}

func (m *MemmappedFileSystemDirectory) Reset()                    { *m = MemmappedFileSystemDirectory{} }
func (m *MemmappedFileSystemDirectory) String() string            { return proto.CompactTextString(m) }
func (*MemmappedFileSystemDirectory) ProtoMessage()               {}
func (*MemmappedFileSystemDirectory) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *MemmappedFileSystemDirectory) GetElement() []*MemmappedFileSystemDirectoryElement {
	if m != nil {
		return m.Element
	}
	return nil
}

func init() {
	proto.RegisterType((*MemmappedFileSystemDirectoryElement)(nil), "tensorflow.MemmappedFileSystemDirectoryElement")
	proto.RegisterType((*MemmappedFileSystemDirectory)(nil), "tensorflow.MemmappedFileSystemDirectory")
}

func init() { proto.RegisterFile("tensorflow/core/util/memmapped_file_system.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x31, 0x4b, 0x04, 0x31,
	0x10, 0x85, 0x89, 0x1e, 0x27, 0xc6, 0x2e, 0x85, 0x6c, 0x61, 0xb1, 0x9c, 0xcd, 0x36, 0x97, 0x11,
	0xfd, 0x01, 0x82, 0x9c, 0x82, 0x85, 0x85, 0x6b, 0x67, 0x73, 0xdc, 0xed, 0x4d, 0x62, 0xb8, 0x24,
	0x13, 0x92, 0x39, 0x65, 0xff, 0xb9, 0xa5, 0xb0, 0xba, 0x6c, 0x23, 0x62, 0x37, 0x03, 0xef, 0x7d,
	0x3c, 0x3e, 0x79, 0xc5, 0x18, 0x0b, 0x65, 0xe3, 0xe9, 0x03, 0x3a, 0xca, 0x08, 0x07, 0x76, 0x1e,
	0x02, 0x86, 0xb0, 0x49, 0x09, 0x77, 0x6b, 0xe3, 0x3c, 0xae, 0x4b, 0x5f, 0x18, 0x83, 0x4e, 0x99,
	0x98, 0x94, 0x9c, 0x1a, 0x8b, 0x67, 0x79, 0xf9, 0x34, 0x46, 0x1f, 0x9c, 0xc7, 0x97, 0x21, 0xb8,
	0x72, 0x19, 0x3b, 0xa6, 0xdc, 0xdf, 0x7b, 0x0c, 0x18, 0x59, 0x9d, 0xcb, 0x39, 0x19, 0x53, 0x90,
	0x2b, 0x51, 0x8b, 0x66, 0xd6, 0xfe, 0x7c, 0x4a, 0xc9, 0x59, 0xdc, 0x04, 0xac, 0x8e, 0x6a, 0xd1,
	0x9c, 0xb6, 0xc3, 0xbd, 0x70, 0xf2, 0xe2, 0x2f, 0xa4, 0x7a, 0x94, 0x27, 0xf8, 0x8d, 0xad, 0x44,
	0x7d, 0xdc, 0x9c, 0x5d, 0x83, 0x9e, 0x06, 0xe9, 0x7f, 0xac, 0x69, 0xc7, 0xfe, 0xdd, 0xea, 0xf5,
	0xd6, 0x3a, 0x7e, 0x3b, 0x6c, 0x75, 0x47, 0x01, 0x76, 0x19, 0xfb, 0x3d, 0x4c, 0xac, 0x65, 0xc1,
	0xfc, 0xee, 0xa2, 0x5d, 0x5a, 0x82, 0xb4, 0xb7, 0xf0, 0x9b, 0xa8, 0x4f, 0x21, 0xb6, 0xf3, 0x41,
	0xcb, 0xcd, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x52, 0x21, 0xc2, 0x58, 0x4a, 0x01, 0x00, 0x00,
}
