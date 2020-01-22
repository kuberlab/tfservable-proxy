// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/python/training/checkpoint_state.proto

/*
Package training is a generated protocol buffer package.

It is generated from these files:
	tensorflow/python/training/checkpoint_state.proto

It has these top-level messages:
	CheckpointState
*/
package training

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

// Protocol buffer representing the checkpoint state.
//
// TODO(touts): Add other attributes as needed.
type CheckpointState struct {
	// Path to the most-recent model checkpoint.
	ModelCheckpointPath string `protobuf:"bytes,1,opt,name=model_checkpoint_path,json=modelCheckpointPath" json:"model_checkpoint_path,omitempty"`
	// Paths to all not-yet-deleted model checkpoints, sorted from oldest to
	// newest.
	// Note that the value of model_checkpoint_path should be the last item in
	// this list.
	AllModelCheckpointPaths []string `protobuf:"bytes,2,rep,name=all_model_checkpoint_paths,json=allModelCheckpointPaths" json:"all_model_checkpoint_paths,omitempty"`
}

func (m *CheckpointState) Reset()                    { *m = CheckpointState{} }
func (m *CheckpointState) String() string            { return proto.CompactTextString(m) }
func (*CheckpointState) ProtoMessage()               {}
func (*CheckpointState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CheckpointState) GetModelCheckpointPath() string {
	if m != nil {
		return m.ModelCheckpointPath
	}
	return ""
}

func (m *CheckpointState) GetAllModelCheckpointPaths() []string {
	if m != nil {
		return m.AllModelCheckpointPaths
	}
	return nil
}

func init() {
	proto.RegisterType((*CheckpointState)(nil), "tensorflow.CheckpointState")
}

func init() { proto.RegisterFile("tensorflow/python/training/checkpoint_state.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x2c, 0x49, 0xcd, 0x2b,
	0xce, 0x2f, 0x4a, 0xcb, 0xc9, 0x2f, 0xd7, 0x2f, 0xa8, 0x2c, 0xc9, 0xc8, 0xcf, 0xd3, 0x2f, 0x29,
	0x4a, 0xcc, 0xcc, 0xcb, 0xcc, 0x4b, 0xd7, 0x4f, 0xce, 0x48, 0x4d, 0xce, 0x2e, 0xc8, 0xcf, 0xcc,
	0x2b, 0x89, 0x2f, 0x2e, 0x49, 0x2c, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42,
	0x68, 0x51, 0x6a, 0x62, 0xe4, 0xe2, 0x77, 0x86, 0x2b, 0x0b, 0x06, 0xa9, 0x12, 0x32, 0xe2, 0x12,
	0xcd, 0xcd, 0x4f, 0x49, 0xcd, 0x89, 0x47, 0xd2, 0x5f, 0x90, 0x58, 0x92, 0x21, 0xc1, 0xa8, 0xc0,
	0xa8, 0xc1, 0x19, 0x24, 0x0c, 0x96, 0x44, 0x68, 0x0a, 0x48, 0x2c, 0xc9, 0x10, 0xb2, 0xe6, 0x92,
	0x4a, 0xcc, 0xc9, 0x89, 0xc7, 0xaa, 0xaf, 0x58, 0x82, 0x49, 0x81, 0x59, 0x83, 0x33, 0x48, 0x3c,
	0x31, 0x27, 0xc7, 0x17, 0x53, 0x6f, 0xb1, 0x93, 0x57, 0x94, 0x6b, 0x7a, 0x66, 0x49, 0x46, 0x69,
	0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x4a, 0x51, 0x6a, 0x65, 0xb6, 0x3e, 0xc2, 0x8d, 0xba, 0xc5,
	0xa9, 0x45, 0x65, 0x99, 0x79, 0xe9, 0xba, 0xe9, 0xf9, 0xfa, 0x05, 0xd9, 0xe9, 0xfa, 0xb8, 0x3d,
	0xfc, 0x83, 0x91, 0x31, 0x89, 0x0d, 0xec, 0x47, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa1,
	0xc9, 0xa4, 0x47, 0x18, 0x01, 0x00, 0x00,
}