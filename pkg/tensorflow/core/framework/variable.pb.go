// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/variable.proto

package framework

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Protocol buffer representing a Variable.
type VariableDef struct {
	// Name of the variable tensor.
	VariableName string `protobuf:"bytes,1,opt,name=variable_name,json=variableName" json:"variable_name,omitempty"`
	// Name of the initializer op.
	InitializerName string `protobuf:"bytes,2,opt,name=initializer_name,json=initializerName" json:"initializer_name,omitempty"`
	// Name of the snapshot tensor.
	SnapshotName string `protobuf:"bytes,3,opt,name=snapshot_name,json=snapshotName" json:"snapshot_name,omitempty"`
	// Support for saving variables as slices of a larger variable.
	SaveSliceInfoDef *SaveSliceInfoDef `protobuf:"bytes,4,opt,name=save_slice_info_def,json=saveSliceInfoDef" json:"save_slice_info_def,omitempty"`
	// Whether to represent this as a ResourceVariable.
	IsResource bool `protobuf:"varint,5,opt,name=is_resource,json=isResource" json:"is_resource,omitempty"`
}

func (m *VariableDef) Reset()                    { *m = VariableDef{} }
func (m *VariableDef) String() string            { return proto.CompactTextString(m) }
func (*VariableDef) ProtoMessage()               {}
func (*VariableDef) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{0} }

func (m *VariableDef) GetVariableName() string {
	if m != nil {
		return m.VariableName
	}
	return ""
}

func (m *VariableDef) GetInitializerName() string {
	if m != nil {
		return m.InitializerName
	}
	return ""
}

func (m *VariableDef) GetSnapshotName() string {
	if m != nil {
		return m.SnapshotName
	}
	return ""
}

func (m *VariableDef) GetSaveSliceInfoDef() *SaveSliceInfoDef {
	if m != nil {
		return m.SaveSliceInfoDef
	}
	return nil
}

func (m *VariableDef) GetIsResource() bool {
	if m != nil {
		return m.IsResource
	}
	return false
}

type SaveSliceInfoDef struct {
	// Name of the full variable of which this is a slice.
	FullName string `protobuf:"bytes,1,opt,name=full_name,json=fullName" json:"full_name,omitempty"`
	// Shape of the full variable.
	FullShape []int64 `protobuf:"varint,2,rep,packed,name=full_shape,json=fullShape" json:"full_shape,omitempty"`
	// Offset of this variable into the full variable.
	VarOffset []int64 `protobuf:"varint,3,rep,packed,name=var_offset,json=varOffset" json:"var_offset,omitempty"`
	// Shape of this variable.
	VarShape []int64 `protobuf:"varint,4,rep,packed,name=var_shape,json=varShape" json:"var_shape,omitempty"`
}

func (m *SaveSliceInfoDef) Reset()                    { *m = SaveSliceInfoDef{} }
func (m *SaveSliceInfoDef) String() string            { return proto.CompactTextString(m) }
func (*SaveSliceInfoDef) ProtoMessage()               {}
func (*SaveSliceInfoDef) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{1} }

func (m *SaveSliceInfoDef) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *SaveSliceInfoDef) GetFullShape() []int64 {
	if m != nil {
		return m.FullShape
	}
	return nil
}

func (m *SaveSliceInfoDef) GetVarOffset() []int64 {
	if m != nil {
		return m.VarOffset
	}
	return nil
}

func (m *SaveSliceInfoDef) GetVarShape() []int64 {
	if m != nil {
		return m.VarShape
	}
	return nil
}

func init() {
	proto.RegisterType((*VariableDef)(nil), "tensorflow.VariableDef")
	proto.RegisterType((*SaveSliceInfoDef)(nil), "tensorflow.SaveSliceInfoDef")
}

func init() { proto.RegisterFile("tensorflow/core/framework/variable.proto", fileDescriptor22) }

var fileDescriptor22 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x41, 0x4b, 0xeb, 0x40,
	0x10, 0xc7, 0xd9, 0x97, 0xbe, 0x47, 0xbb, 0x7d, 0xd5, 0x12, 0x2f, 0x01, 0x15, 0x4b, 0x7b, 0x89,
	0x87, 0x26, 0xa0, 0xdf, 0xa0, 0xf4, 0x22, 0x82, 0x96, 0x14, 0x3c, 0x78, 0x09, 0xdb, 0x74, 0x36,
	0x5d, 0x9a, 0x64, 0xc3, 0x6c, 0xb2, 0x45, 0x3f, 0x82, 0x9f, 0xd6, 0xa3, 0x47, 0xd9, 0x4d, 0x63,
	0x4b, 0xc1, 0xeb, 0xef, 0xff, 0x9b, 0x61, 0x86, 0x19, 0xea, 0x57, 0x50, 0x28, 0x89, 0x3c, 0x93,
	0xbb, 0x30, 0x91, 0x08, 0x21, 0x47, 0x96, 0xc3, 0x4e, 0xe2, 0x36, 0xd4, 0x0c, 0x05, 0x5b, 0x65,
	0x10, 0x94, 0x28, 0x2b, 0xe9, 0xd2, 0x83, 0x39, 0xfe, 0x24, 0xb4, 0xff, 0xb2, 0x8f, 0xe7, 0xc0,
	0xdd, 0x09, 0x1d, 0xb4, 0x76, 0x5c, 0xb0, 0x1c, 0x3c, 0x32, 0x22, 0x7e, 0x2f, 0xfa, 0xdf, 0xc2,
	0x27, 0x96, 0x83, 0x7b, 0x4b, 0x87, 0xa2, 0x10, 0x95, 0x60, 0x99, 0x78, 0x07, 0x6c, 0xbc, 0x3f,
	0xd6, 0x3b, 0x3f, 0xe2, 0x56, 0x9d, 0xd0, 0x81, 0x2a, 0x58, 0xa9, 0x36, 0xb2, 0x6a, 0x3c, 0xa7,
	0xe9, 0xd7, 0x42, 0x2b, 0x3d, 0xd2, 0x0b, 0xc5, 0x34, 0xc4, 0x2a, 0x13, 0x09, 0xc4, 0xa2, 0xe0,
	0x32, 0x5e, 0x03, 0xf7, 0x3a, 0x23, 0xe2, 0xf7, 0xef, 0xae, 0x82, 0xc3, 0xb8, 0xc1, 0x92, 0x69,
	0x58, 0x1a, 0xeb, 0xa1, 0xe0, 0x72, 0x0e, 0x3c, 0x1a, 0xaa, 0x13, 0xe2, 0xde, 0xd0, 0xbe, 0x50,
	0x31, 0x82, 0x92, 0x35, 0x26, 0xe0, 0xfd, 0x1d, 0x11, 0xbf, 0x1b, 0x51, 0xa1, 0xa2, 0x3d, 0x19,
	0x7f, 0x10, 0x3a, 0x3c, 0xed, 0xe3, 0x5e, 0xd2, 0x1e, 0xaf, 0xb3, 0xec, 0x78, 0xe7, 0xae, 0x01,
	0x76, 0xbe, 0x6b, 0x4a, 0x6d, 0xa8, 0x36, 0xac, 0x34, 0x9b, 0x3a, 0xbe, 0x13, 0x59, 0x7d, 0x69,
	0x80, 0x89, 0x35, 0xc3, 0x58, 0x72, 0xae, 0xa0, 0xf2, 0x9c, 0x26, 0xd6, 0x0c, 0x9f, 0x2d, 0x30,
	0xad, 0x4d, 0xdc, 0x14, 0x77, 0x6c, 0xda, 0xd5, 0x0c, 0x6d, 0xed, 0xac, 0xa6, 0x9e, 0xc4, 0xf4,
	0x78, 0xc5, 0x9f, 0xb3, 0xcd, 0xce, 0xda, 0xc3, 0x2c, 0xcc, 0xd9, 0xd4, 0x82, 0xbc, 0xce, 0x53,
	0x51, 0x6d, 0xea, 0x55, 0x90, 0xc8, 0x3c, 0x5c, 0x23, 0xbc, 0x6d, 0xc3, 0x43, 0xe1, 0x54, 0x01,
	0x6a, 0x51, 0xa4, 0xd3, 0x54, 0x86, 0xe5, 0x36, 0x0d, 0x7f, 0x7d, 0x87, 0x2f, 0x42, 0x56, 0xff,
	0xec, 0x27, 0xdc, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x15, 0x82, 0xa6, 0x49, 0x35, 0x02, 0x00,
	0x00,
}