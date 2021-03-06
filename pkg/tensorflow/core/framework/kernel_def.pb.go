// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/kernel_def.proto

package framework

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type KernelDef struct {
	// Must match the name of an Op.
	Op string `protobuf:"bytes,1,opt,name=op" json:"op,omitempty"`
	// Type of device this kernel runs on.
	DeviceType string                      `protobuf:"bytes,2,opt,name=device_type,json=deviceType" json:"device_type,omitempty"`
	Constraint []*KernelDef_AttrConstraint `protobuf:"bytes,3,rep,name=constraint" json:"constraint,omitempty"`
	// Names of the Op's input_/output_args that reside in host memory
	// instead of device memory.
	HostMemoryArg []string `protobuf:"bytes,4,rep,name=host_memory_arg,json=hostMemoryArg" json:"host_memory_arg,omitempty"`
	// This allows experimental kernels to be registered for an op that
	// won't be used unless the user specifies a "_kernel" attr with
	// value matching this.
	Label string `protobuf:"bytes,5,opt,name=label" json:"label,omitempty"`
}

func (m *KernelDef) Reset()                    { *m = KernelDef{} }
func (m *KernelDef) String() string            { return proto.CompactTextString(m) }
func (*KernelDef) ProtoMessage()               {}
func (*KernelDef) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *KernelDef) GetOp() string {
	if m != nil {
		return m.Op
	}
	return ""
}

func (m *KernelDef) GetDeviceType() string {
	if m != nil {
		return m.DeviceType
	}
	return ""
}

func (m *KernelDef) GetConstraint() []*KernelDef_AttrConstraint {
	if m != nil {
		return m.Constraint
	}
	return nil
}

func (m *KernelDef) GetHostMemoryArg() []string {
	if m != nil {
		return m.HostMemoryArg
	}
	return nil
}

func (m *KernelDef) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

type KernelDef_AttrConstraint struct {
	// Name of an attr from the Op.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// A list of values that this kernel supports for this attr.
	// Like OpDef.AttrDef.allowed_values, except for kernels instead of Ops.
	AllowedValues *AttrValue `protobuf:"bytes,2,opt,name=allowed_values,json=allowedValues" json:"allowed_values,omitempty"`
}

func (m *KernelDef_AttrConstraint) Reset()                    { *m = KernelDef_AttrConstraint{} }
func (m *KernelDef_AttrConstraint) String() string            { return proto.CompactTextString(m) }
func (*KernelDef_AttrConstraint) ProtoMessage()               {}
func (*KernelDef_AttrConstraint) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0, 0} }

func (m *KernelDef_AttrConstraint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KernelDef_AttrConstraint) GetAllowedValues() *AttrValue {
	if m != nil {
		return m.AllowedValues
	}
	return nil
}

func init() {
	proto.RegisterType((*KernelDef)(nil), "tensorflow.KernelDef")
	proto.RegisterType((*KernelDef_AttrConstraint)(nil), "tensorflow.KernelDef.AttrConstraint")
}

func init() { proto.RegisterFile("tensorflow/core/framework/kernel_def.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x4d, 0x6b, 0xc2, 0x40,
	0x10, 0x25, 0x51, 0x0b, 0xae, 0xa8, 0xb0, 0xb4, 0x10, 0xbc, 0x54, 0x4a, 0x29, 0x52, 0x30, 0x01,
	0x7b, 0xed, 0x45, 0xeb, 0xad, 0x14, 0x24, 0x94, 0x1e, 0x7a, 0x09, 0x9b, 0x64, 0x12, 0x43, 0x3e,
	0x26, 0xcc, 0xae, 0x91, 0xfc, 0x9f, 0xfe, 0xc8, 0x1e, 0x4b, 0x36, 0x12, 0xed, 0xa1, 0xbd, 0xed,
	0xbc, 0x79, 0xef, 0xed, 0xbc, 0x19, 0xf6, 0xa8, 0xa0, 0x90, 0x48, 0x51, 0x86, 0x47, 0x27, 0x40,
	0x02, 0x27, 0x22, 0x91, 0xc3, 0x11, 0x29, 0x75, 0x52, 0xa0, 0x02, 0x32, 0x2f, 0x84, 0xc8, 0x2e,
	0x09, 0x15, 0x72, 0x76, 0xe6, 0xce, 0xfe, 0xd1, 0x09, 0xa5, 0xc8, 0xab, 0x44, 0x76, 0x80, 0x56,
	0x77, 0xf7, 0x65, 0xb2, 0xe1, 0xab, 0x36, 0xdb, 0x42, 0xc4, 0x27, 0xcc, 0xc4, 0xd2, 0x32, 0xe6,
	0xc6, 0x62, 0xe8, 0x9a, 0x58, 0xf2, 0x5b, 0x36, 0x0a, 0xa1, 0x4a, 0x02, 0xf0, 0x54, 0x5d, 0x82,
	0x65, 0xea, 0x06, 0x6b, 0xa1, 0xf7, 0xba, 0x04, 0xbe, 0x65, 0x2c, 0xc0, 0x42, 0x2a, 0x12, 0x49,
	0xa1, 0xac, 0xde, 0xbc, 0xb7, 0x18, 0xad, 0xee, 0xed, 0xf3, 0xff, 0x76, 0xe7, 0x6d, 0xaf, 0x95,
	0xa2, 0x97, 0x8e, 0xeb, 0x5e, 0xe8, 0xf8, 0x03, 0x9b, 0xee, 0x51, 0x2a, 0x2f, 0x87, 0x1c, 0xa9,
	0xf6, 0x04, 0xc5, 0x56, 0x7f, 0xde, 0x5b, 0x0c, 0xdd, 0x71, 0x03, 0xbf, 0x69, 0x74, 0x4d, 0x31,
	0xbf, 0x66, 0x83, 0x4c, 0xf8, 0x90, 0x59, 0x03, 0x3d, 0x48, 0x5b, 0xcc, 0x7c, 0x36, 0xf9, 0xed,
	0xcd, 0x39, 0xeb, 0x17, 0x22, 0x87, 0x53, 0x10, 0xfd, 0xe6, 0xcf, 0x6c, 0x22, 0xb2, 0x0c, 0x8f,
	0x10, 0xb6, 0xf9, 0xa5, 0x4e, 0x33, 0x5a, 0xdd, 0x5c, 0x4e, 0xdb, 0xf8, 0x7c, 0x34, 0x5d, 0x77,
	0x7c, 0x22, 0xeb, 0x4a, 0x6e, 0x2a, 0x66, 0x21, 0xc5, 0x97, 0xd4, 0x6e, 0xa7, 0x9b, 0x69, 0x97,
	0x71, 0xd7, 0xac, 0x54, 0xee, 0x8c, 0xcf, 0x6d, 0x9c, 0xa8, 0xfd, 0xc1, 0xb7, 0x03, 0xcc, 0x9d,
	0x90, 0xa0, 0x4e, 0x9d, 0xb3, 0x72, 0x29, 0x81, 0xaa, 0xa4, 0x88, 0x97, 0x31, 0x3a, 0x65, 0x1a,
	0x3b, 0x7f, 0x1e, 0xeb, 0xdb, 0x30, 0xfc, 0x2b, 0x7d, 0xa5, 0xa7, 0x9f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x40, 0x5b, 0x28, 0x46, 0x0b, 0x02, 0x00, 0x00,
}
