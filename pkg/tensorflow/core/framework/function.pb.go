// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/function.proto

package framework

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A library is a set of named functions.
type FunctionDefLibrary struct {
	Function []*FunctionDef `protobuf:"bytes,1,rep,name=function" json:"function,omitempty"`
	Gradient []*GradientDef `protobuf:"bytes,2,rep,name=gradient" json:"gradient,omitempty"`
}

func (m *FunctionDefLibrary) Reset()                    { *m = FunctionDefLibrary{} }
func (m *FunctionDefLibrary) String() string            { return proto.CompactTextString(m) }
func (*FunctionDefLibrary) ProtoMessage()               {}
func (*FunctionDefLibrary) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *FunctionDefLibrary) GetFunction() []*FunctionDef {
	if m != nil {
		return m.Function
	}
	return nil
}

func (m *FunctionDefLibrary) GetGradient() []*GradientDef {
	if m != nil {
		return m.Gradient
	}
	return nil
}

// A function can be instantiated when the runtime can bind every attr
// with a value. When a GraphDef has a call to a function, it must
// have binding for every attr defined in the signature.
//
// TODO(zhifengc):
//   * device spec, etc.
type FunctionDef struct {
	// The definition of the function's name, arguments, return values,
	// attrs etc.
	Signature *OpDef `protobuf:"bytes,1,opt,name=signature" json:"signature,omitempty"`
	// Attributes specific to this function definition.
	Attr map[string]*AttrValue `protobuf:"bytes,5,rep,name=attr" json:"attr,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// By convention, "op" in node_def is resolved by consulting with a
	// user-defined library first. If not resolved, "func" is assumed to
	// be a builtin op.
	NodeDef []*NodeDef `protobuf:"bytes,3,rep,name=node_def,json=nodeDef" json:"node_def,omitempty"`
	// A mapping from the output arg names from `signature` to the
	// outputs from `node_def` that should be returned by the function.
	Ret map[string]string `protobuf:"bytes,4,rep,name=ret" json:"ret,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *FunctionDef) Reset()                    { *m = FunctionDef{} }
func (m *FunctionDef) String() string            { return proto.CompactTextString(m) }
func (*FunctionDef) ProtoMessage()               {}
func (*FunctionDef) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *FunctionDef) GetSignature() *OpDef {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *FunctionDef) GetAttr() map[string]*AttrValue {
	if m != nil {
		return m.Attr
	}
	return nil
}

func (m *FunctionDef) GetNodeDef() []*NodeDef {
	if m != nil {
		return m.NodeDef
	}
	return nil
}

func (m *FunctionDef) GetRet() map[string]string {
	if m != nil {
		return m.Ret
	}
	return nil
}

// GradientDef defines the gradient function of a function defined in
// a function library.
//
// A gradient function g (specified by gradient_func) for a function f
// (specified by function_name) must follow the following:
//
// The function 'f' must be a numerical function which takes N inputs
// and produces M outputs. Its gradient function 'g', which is a
// function taking N + M inputs and produces N outputs.
//
// I.e. if we have
//    (y1, y2, ..., y_M) = f(x1, x2, ..., x_N),
// then, g is
//    (dL/dx1, dL/dx2, ..., dL/dx_N) = g(x1, x2, ..., x_N,
//                                      dL/dy1, dL/dy2, ..., dL/dy_M),
// where L is a scalar-value function of (x1, x2, ..., xN) (e.g., the
// loss function). dL/dx_i is the partial derivative of L with respect
// to x_i.
type GradientDef struct {
	FunctionName string `protobuf:"bytes,1,opt,name=function_name,json=functionName" json:"function_name,omitempty"`
	GradientFunc string `protobuf:"bytes,2,opt,name=gradient_func,json=gradientFunc" json:"gradient_func,omitempty"`
}

func (m *GradientDef) Reset()                    { *m = GradientDef{} }
func (m *GradientDef) String() string            { return proto.CompactTextString(m) }
func (*GradientDef) ProtoMessage()               {}
func (*GradientDef) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *GradientDef) GetFunctionName() string {
	if m != nil {
		return m.FunctionName
	}
	return ""
}

func (m *GradientDef) GetGradientFunc() string {
	if m != nil {
		return m.GradientFunc
	}
	return ""
}

func init() {
	proto.RegisterType((*FunctionDefLibrary)(nil), "tensorflow.FunctionDefLibrary")
	proto.RegisterType((*FunctionDef)(nil), "tensorflow.FunctionDef")
	proto.RegisterType((*GradientDef)(nil), "tensorflow.GradientDef")
}

func init() { proto.RegisterFile("tensorflow/core/framework/function.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x8b, 0x13, 0x31,
	0x14, 0xc6, 0x99, 0x76, 0x57, 0xdb, 0xd7, 0x55, 0x34, 0x2a, 0x86, 0x9e, 0x6a, 0x05, 0x29, 0xca,
	0x4e, 0xa0, 0x8b, 0x22, 0xde, 0x5c, 0x56, 0xbd, 0x48, 0x5d, 0xe6, 0xa0, 0xe0, 0xa5, 0xa4, 0x9d,
	0x37, 0xe3, 0xd0, 0x6d, 0x52, 0x5e, 0x33, 0x5b, 0xe6, 0xe2, 0xff, 0xea, 0x7f, 0xe1, 0x51, 0x92,
	0x99, 0x74, 0x02, 0x3a, 0xde, 0x42, 0xf2, 0xfd, 0xbe, 0xef, 0xe5, 0xbd, 0x07, 0x33, 0x83, 0x6a,
	0xaf, 0x29, 0xbb, 0xd1, 0x07, 0xb1, 0xd6, 0x84, 0x22, 0x23, 0xb9, 0xc5, 0x83, 0xa6, 0x8d, 0xc8,
	0x4a, 0xb5, 0x36, 0x85, 0x56, 0xf1, 0x8e, 0xb4, 0xd1, 0x0c, 0x5a, 0xe5, 0xf8, 0x65, 0x37, 0x25,
	0x8d, 0xa1, 0xe5, 0xad, 0xbc, 0x29, 0xb1, 0xe6, 0xc6, 0xff, 0x49, 0x50, 0x3a, 0xc5, 0x65, 0x8a,
	0x59, 0xa3, 0x7c, 0xd1, 0xad, 0xd4, 0xbb, 0x56, 0x37, 0xfd, 0x09, 0xec, 0x63, 0x53, 0xdb, 0x15,
	0x66, 0x9f, 0x8b, 0x15, 0x49, 0xaa, 0xd8, 0x05, 0x0c, 0x7c, 0xc5, 0x3c, 0x9a, 0xf4, 0x67, 0xa3,
	0xf9, 0xd3, 0xb8, 0x35, 0x8c, 0x03, 0x22, 0x39, 0x0a, 0x2d, 0x94, 0x93, 0x4c, 0x0b, 0x54, 0x86,
	0xf7, 0xfe, 0x86, 0x3e, 0x35, 0x6f, 0x0e, 0xf2, 0xc2, 0xe9, 0xaf, 0x1e, 0x8c, 0x02, 0x3b, 0x26,
	0x60, 0xb8, 0x2f, 0x72, 0x25, 0x4d, 0x49, 0xc8, 0xa3, 0x49, 0x34, 0x1b, 0xcd, 0x1f, 0x86, 0x2e,
	0x5f, 0x76, 0x96, 0x6f, 0x35, 0xec, 0x35, 0x9c, 0xd8, 0x36, 0xf1, 0x53, 0x97, 0xf8, 0xac, 0xa3,
	0xcc, 0xf8, 0xbd, 0x31, 0xf4, 0x41, 0x19, 0xaa, 0x12, 0x27, 0x67, 0x31, 0x0c, 0x7c, 0xc7, 0x78,
	0xdf, 0xa1, 0x8f, 0x42, 0x74, 0xa1, 0x53, 0xb4, 0x41, 0x77, 0x55, 0x7d, 0x60, 0x73, 0xe8, 0x13,
	0x1a, 0x7e, 0xe2, 0xa4, 0x93, 0xae, 0x94, 0x04, 0x4d, 0x1d, 0x62, 0xc5, 0xe3, 0x05, 0x0c, 0x8f,
	0xb1, 0xec, 0x01, 0xf4, 0x37, 0x58, 0xb9, 0x2f, 0x0d, 0x13, 0x7b, 0x64, 0xaf, 0xe0, 0xd4, 0xcd,
	0x96, 0xf7, 0xdc, 0x37, 0x9f, 0x84, 0xa6, 0x96, 0xfb, 0x6a, 0x1f, 0x93, 0x5a, 0xf3, 0xae, 0xf7,
	0x36, 0x1a, 0xbf, 0x81, 0x81, 0x0f, 0xf8, 0x87, 0xdd, 0xe3, 0xd0, 0x6e, 0x18, 0x70, 0xd3, 0x6f,
	0x30, 0x0a, 0x9a, 0xcf, 0x9e, 0xc3, 0x3d, 0x3f, 0xb3, 0xa5, 0x92, 0x5b, 0x6c, 0x4c, 0xce, 0xfc,
	0xe5, 0x42, 0x6e, 0xd1, 0x8a, 0xfc, 0x8c, 0x96, 0xf6, 0xa1, 0x71, 0x3d, 0xf3, 0x97, 0xf6, 0xd7,
	0x97, 0x25, 0x70, 0x4d, 0x79, 0x58, 0xf7, 0x71, 0xcb, 0x2e, 0xef, 0xfb, 0xbe, 0x5c, 0xdb, 0x3d,
	0xdb, 0x5f, 0x47, 0xdf, 0xaf, 0xf2, 0xc2, 0xfc, 0x28, 0x57, 0xf1, 0x5a, 0x6f, 0x45, 0x4a, 0x58,
	0x6d, 0x44, 0x0b, 0x9e, 0xef, 0x91, 0x6e, 0x0b, 0x95, 0x9f, 0xe7, 0x5a, 0xec, 0x36, 0xb9, 0xe8,
	0xdc, 0xde, 0xdf, 0x51, 0xb4, 0xba, 0xe3, 0x56, 0xf7, 0xe2, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x11, 0x05, 0xaf, 0x3d, 0x70, 0x03, 0x00, 0x00,
}