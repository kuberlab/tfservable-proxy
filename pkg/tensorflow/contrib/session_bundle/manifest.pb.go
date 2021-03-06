// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/contrib/session_bundle/manifest.proto

/*
Package session_bundle is a generated protocol buffer package.

It is generated from these files:
	tensorflow/contrib/session_bundle/manifest.proto

It has these top-level messages:
	Signatures
	TensorBinding
	AssetFile
	Signature
	RegressionSignature
	ClassificationSignature
	GenericSignature
*/
package session_bundle

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

// Signatures of model export.
type Signatures struct {
	// Default signature of the graph.
	// WARNING(break-tutorial-inline-code): The following code snippet is
	// in-lined in tutorials, please update tutorial documents accordingly
	// whenever code changes.
	DefaultSignature *Signature `protobuf:"bytes,1,opt,name=default_signature,json=defaultSignature" json:"default_signature,omitempty"`
	// Named signatures of the graph.
	NamedSignatures map[string]*Signature `protobuf:"bytes,2,rep,name=named_signatures,json=namedSignatures" json:"named_signatures,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Signatures) Reset()                    { *m = Signatures{} }
func (m *Signatures) String() string            { return proto.CompactTextString(m) }
func (*Signatures) ProtoMessage()               {}
func (*Signatures) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Signatures) GetDefaultSignature() *Signature {
	if m != nil {
		return m.DefaultSignature
	}
	return nil
}

func (m *Signatures) GetNamedSignatures() map[string]*Signature {
	if m != nil {
		return m.NamedSignatures
	}
	return nil
}

// A binding to a tensor including the name and, possibly in the future, type
// or other metadata. For example, this may specify whether a tensor supports
// batch vs single inference.
type TensorBinding struct {
	// The name of the tensor to bind to.
	TensorName string `protobuf:"bytes,1,opt,name=tensor_name,json=tensorName" json:"tensor_name,omitempty"`
}

func (m *TensorBinding) Reset()                    { *m = TensorBinding{} }
func (m *TensorBinding) String() string            { return proto.CompactTextString(m) }
func (*TensorBinding) ProtoMessage()               {}
func (*TensorBinding) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TensorBinding) GetTensorName() string {
	if m != nil {
		return m.TensorName
	}
	return ""
}

// An asset file or set of sharded files with the same name that will be bound
// to a tensor at init / session_bundle load time.
type AssetFile struct {
	// The tensor to bind the asset filename to.
	TensorBinding *TensorBinding `protobuf:"bytes,1,opt,name=tensor_binding,json=tensorBinding" json:"tensor_binding,omitempty"`
	// The filename within the assets directory. Note: does not include the base
	// path or asset directory prefix. Base paths can and will change when models
	// are deployed for serving.
	Filename string `protobuf:"bytes,2,opt,name=filename" json:"filename,omitempty"`
}

func (m *AssetFile) Reset()                    { *m = AssetFile{} }
func (m *AssetFile) String() string            { return proto.CompactTextString(m) }
func (*AssetFile) ProtoMessage()               {}
func (*AssetFile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AssetFile) GetTensorBinding() *TensorBinding {
	if m != nil {
		return m.TensorBinding
	}
	return nil
}

func (m *AssetFile) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

// A Signature specifies the inputs and outputs of commonly used graphs.
type Signature struct {
	// Types that are valid to be assigned to Type:
	//	*Signature_RegressionSignature
	//	*Signature_ClassificationSignature
	//	*Signature_GenericSignature
	Type isSignature_Type `protobuf_oneof:"type"`
}

func (m *Signature) Reset()                    { *m = Signature{} }
func (m *Signature) String() string            { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()               {}
func (*Signature) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type isSignature_Type interface {
	isSignature_Type()
}

type Signature_RegressionSignature struct {
	RegressionSignature *RegressionSignature `protobuf:"bytes,1,opt,name=regression_signature,json=regressionSignature,oneof"`
}
type Signature_ClassificationSignature struct {
	ClassificationSignature *ClassificationSignature `protobuf:"bytes,2,opt,name=classification_signature,json=classificationSignature,oneof"`
}
type Signature_GenericSignature struct {
	GenericSignature *GenericSignature `protobuf:"bytes,3,opt,name=generic_signature,json=genericSignature,oneof"`
}

func (*Signature_RegressionSignature) isSignature_Type()     {}
func (*Signature_ClassificationSignature) isSignature_Type() {}
func (*Signature_GenericSignature) isSignature_Type()        {}

func (m *Signature) GetType() isSignature_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Signature) GetRegressionSignature() *RegressionSignature {
	if x, ok := m.GetType().(*Signature_RegressionSignature); ok {
		return x.RegressionSignature
	}
	return nil
}

func (m *Signature) GetClassificationSignature() *ClassificationSignature {
	if x, ok := m.GetType().(*Signature_ClassificationSignature); ok {
		return x.ClassificationSignature
	}
	return nil
}

func (m *Signature) GetGenericSignature() *GenericSignature {
	if x, ok := m.GetType().(*Signature_GenericSignature); ok {
		return x.GenericSignature
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Signature) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Signature_OneofMarshaler, _Signature_OneofUnmarshaler, _Signature_OneofSizer, []interface{}{
		(*Signature_RegressionSignature)(nil),
		(*Signature_ClassificationSignature)(nil),
		(*Signature_GenericSignature)(nil),
	}
}

func _Signature_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Signature)
	// type
	switch x := m.Type.(type) {
	case *Signature_RegressionSignature:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RegressionSignature); err != nil {
			return err
		}
	case *Signature_ClassificationSignature:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClassificationSignature); err != nil {
			return err
		}
	case *Signature_GenericSignature:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GenericSignature); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Signature.Type has unexpected type %T", x)
	}
	return nil
}

func _Signature_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Signature)
	switch tag {
	case 1: // type.regression_signature
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RegressionSignature)
		err := b.DecodeMessage(msg)
		m.Type = &Signature_RegressionSignature{msg}
		return true, err
	case 2: // type.classification_signature
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ClassificationSignature)
		err := b.DecodeMessage(msg)
		m.Type = &Signature_ClassificationSignature{msg}
		return true, err
	case 3: // type.generic_signature
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GenericSignature)
		err := b.DecodeMessage(msg)
		m.Type = &Signature_GenericSignature{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Signature_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Signature)
	// type
	switch x := m.Type.(type) {
	case *Signature_RegressionSignature:
		s := proto.Size(x.RegressionSignature)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Signature_ClassificationSignature:
		s := proto.Size(x.ClassificationSignature)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Signature_GenericSignature:
		s := proto.Size(x.GenericSignature)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// RegressionSignature specifies a graph that takes an input and returns an
// output.
type RegressionSignature struct {
	Input  *TensorBinding `protobuf:"bytes,1,opt,name=input" json:"input,omitempty"`
	Output *TensorBinding `protobuf:"bytes,2,opt,name=output" json:"output,omitempty"`
}

func (m *RegressionSignature) Reset()                    { *m = RegressionSignature{} }
func (m *RegressionSignature) String() string            { return proto.CompactTextString(m) }
func (*RegressionSignature) ProtoMessage()               {}
func (*RegressionSignature) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RegressionSignature) GetInput() *TensorBinding {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *RegressionSignature) GetOutput() *TensorBinding {
	if m != nil {
		return m.Output
	}
	return nil
}

// ClassificationSignature specifies a graph that takes an input and returns
// classes and their scores.
// WARNING(break-tutorial-inline-code): The following code snippet is
// in-lined in tutorials, please update tutorial documents accordingly
// whenever code changes.
type ClassificationSignature struct {
	Input   *TensorBinding `protobuf:"bytes,1,opt,name=input" json:"input,omitempty"`
	Classes *TensorBinding `protobuf:"bytes,2,opt,name=classes" json:"classes,omitempty"`
	Scores  *TensorBinding `protobuf:"bytes,3,opt,name=scores" json:"scores,omitempty"`
}

func (m *ClassificationSignature) Reset()                    { *m = ClassificationSignature{} }
func (m *ClassificationSignature) String() string            { return proto.CompactTextString(m) }
func (*ClassificationSignature) ProtoMessage()               {}
func (*ClassificationSignature) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ClassificationSignature) GetInput() *TensorBinding {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *ClassificationSignature) GetClasses() *TensorBinding {
	if m != nil {
		return m.Classes
	}
	return nil
}

func (m *ClassificationSignature) GetScores() *TensorBinding {
	if m != nil {
		return m.Scores
	}
	return nil
}

// GenericSignature specifies a map from logical name to Tensor name.
// Typical application of GenericSignature is to use a single GenericSignature
// that includes all of the Tensor nodes and target names that may be useful at
// serving, analysis or debugging time. The recommended name for this signature
// in the ModelManifest is "generic_bindings".
type GenericSignature struct {
	Map map[string]*TensorBinding `protobuf:"bytes,1,rep,name=map" json:"map,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *GenericSignature) Reset()                    { *m = GenericSignature{} }
func (m *GenericSignature) String() string            { return proto.CompactTextString(m) }
func (*GenericSignature) ProtoMessage()               {}
func (*GenericSignature) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GenericSignature) GetMap() map[string]*TensorBinding {
	if m != nil {
		return m.Map
	}
	return nil
}

func init() {
	proto.RegisterType((*Signatures)(nil), "tensorflow.serving.Signatures")
	proto.RegisterType((*TensorBinding)(nil), "tensorflow.serving.TensorBinding")
	proto.RegisterType((*AssetFile)(nil), "tensorflow.serving.AssetFile")
	proto.RegisterType((*Signature)(nil), "tensorflow.serving.Signature")
	proto.RegisterType((*RegressionSignature)(nil), "tensorflow.serving.RegressionSignature")
	proto.RegisterType((*ClassificationSignature)(nil), "tensorflow.serving.ClassificationSignature")
	proto.RegisterType((*GenericSignature)(nil), "tensorflow.serving.GenericSignature")
}

func init() { proto.RegisterFile("tensorflow/contrib/session_bundle/manifest.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x5b, 0x6e, 0xd3, 0x4c,
	0x14, 0xc7, 0x3f, 0x3b, 0x5f, 0x43, 0x73, 0xa2, 0x82, 0x3b, 0xad, 0x54, 0x2b, 0x12, 0xa2, 0x58,
	0x48, 0x54, 0x42, 0xb1, 0xab, 0xe6, 0xa1, 0x5c, 0x1e, 0x10, 0x41, 0x40, 0x85, 0xa0, 0x0f, 0x2e,
	0x2f, 0x20, 0x44, 0x34, 0xb6, 0xc7, 0xee, 0x28, 0xce, 0x8c, 0x99, 0x19, 0x17, 0x65, 0x09, 0xac,
	0x81, 0x3d, 0x20, 0xd6, 0xc1, 0xaa, 0x90, 0x2f, 0xb1, 0xe3, 0xd4, 0x80, 0xc5, 0x5b, 0x7c, 0x72,
	0x7e, 0xe7, 0x7f, 0x2e, 0xfe, 0x1b, 0x8e, 0x15, 0x61, 0x92, 0x8b, 0x30, 0xe6, 0x5f, 0x1c, 0x9f,
	0x33, 0x25, 0xa8, 0xe7, 0x48, 0x22, 0x25, 0xe5, 0x6c, 0xe6, 0xa5, 0x2c, 0x88, 0x89, 0xb3, 0xc0,
	0x8c, 0x86, 0x44, 0x2a, 0x3b, 0x11, 0x5c, 0x71, 0x84, 0x6a, 0xc2, 0x96, 0x44, 0x5c, 0x51, 0x16,
	0x59, 0xdf, 0x74, 0x80, 0x0b, 0x1a, 0x31, 0xac, 0x52, 0x41, 0x24, 0x7a, 0x0d, 0xbb, 0x01, 0x09,
	0x71, 0x1a, 0xab, 0x99, 0x5c, 0x45, 0x4d, 0xed, 0x50, 0x3b, 0x1a, 0x9e, 0xdc, 0xb6, 0xaf, 0xe3,
	0x76, 0x85, 0xba, 0x46, 0xc9, 0x55, 0x11, 0xf4, 0x09, 0x0c, 0x86, 0x17, 0x24, 0xa8, 0x2b, 0x49,
	0x53, 0x3f, 0xec, 0x1d, 0x0d, 0x4f, 0x26, 0x7f, 0x2c, 0x25, 0xed, 0xf3, 0x0c, 0xab, 0x9f, 0x5f,
	0x30, 0x25, 0x96, 0xee, 0x2d, 0xd6, 0x8c, 0x8e, 0x30, 0xec, 0xb7, 0x25, 0x22, 0x03, 0x7a, 0x73,
	0xb2, 0xcc, 0xbb, 0x1e, 0xb8, 0xd9, 0x4f, 0x34, 0x81, 0xad, 0x2b, 0x1c, 0xa7, 0xc4, 0xd4, 0xbb,
	0x4c, 0x52, 0xe4, 0x3e, 0xd6, 0x1f, 0x6a, 0xd6, 0x31, 0xec, 0xbc, 0xcb, 0x53, 0xa7, 0x94, 0x05,
	0x94, 0x45, 0xe8, 0x0e, 0x0c, 0x0b, 0x76, 0x96, 0x75, 0x53, 0x6a, 0x40, 0x11, 0xca, 0x9a, 0xb1,
	0x3e, 0xc3, 0xe0, 0x99, 0x94, 0x44, 0xbd, 0xa4, 0x31, 0x41, 0x67, 0x70, 0xb3, 0xcc, 0xf6, 0x0a,
	0xbe, 0x5c, 0xe5, 0xdd, 0xb6, 0x06, 0x1a, 0x42, 0xee, 0x8e, 0x6a, 0xe8, 0x8e, 0x60, 0x3b, 0xa4,
	0x31, 0xc9, 0x45, 0xf5, 0x5c, 0xb4, 0x7a, 0xb6, 0x7e, 0xe8, 0x30, 0xa8, 0xb7, 0xfe, 0x11, 0xf6,
	0x05, 0x89, 0x44, 0xf9, 0x22, 0x6c, 0x1e, 0xf1, 0x7e, 0x9b, 0xb2, 0x5b, 0xe5, 0x57, 0x65, 0xce,
	0xfe, 0x73, 0xf7, 0xc4, 0xf5, 0x30, 0xba, 0x04, 0xd3, 0x8f, 0xb1, 0x94, 0x34, 0xa4, 0x3e, 0x56,
	0x4d, 0x85, 0x62, 0xb9, 0x0f, 0xda, 0x14, 0x9e, 0x37, 0x98, 0x75, 0x95, 0x03, 0xbf, 0xfd, 0x2f,
	0x74, 0x01, 0xbb, 0x11, 0x61, 0x44, 0x50, 0x7f, 0x4d, 0xa2, 0x97, 0x4b, 0xdc, 0x6b, 0x93, 0x78,
	0x55, 0x24, 0xaf, 0xd7, 0x36, 0xa2, 0x8d, 0xd8, 0xb4, 0x0f, 0xff, 0xab, 0x65, 0x42, 0xac, 0xaf,
	0x1a, 0xec, 0xb5, 0x4c, 0x8d, 0x4e, 0x61, 0x8b, 0xb2, 0x24, 0x55, 0xdd, 0xef, 0x54, 0xe4, 0xa3,
	0x47, 0xd0, 0xe7, 0xa9, 0xca, 0x48, 0xbd, 0x2b, 0x59, 0x02, 0xd6, 0x4f, 0x0d, 0x0e, 0x7e, 0xb3,
	0x9f, 0x7f, 0xef, 0xe7, 0x09, 0xdc, 0xc8, 0x17, 0x9b, 0x5b, 0xae, 0x23, 0xba, 0x22, 0xb2, 0x61,
	0xa4, 0xcf, 0x33, 0xbb, 0xf6, 0x3a, 0x0f, 0x53, 0x00, 0xd6, 0x77, 0x0d, 0x8c, 0xcd, 0x4b, 0xa0,
	0xa7, 0xd0, 0x5b, 0xe0, 0xc4, 0xd4, 0x72, 0xef, 0x8f, 0xbb, 0x1c, 0xcf, 0x7e, 0x8b, 0x93, 0xc2,
	0xf5, 0x19, 0x39, 0x7a, 0x0f, 0xdb, 0xab, 0x40, 0x8b, 0xbb, 0x4f, 0x9b, 0xee, 0xee, 0xb2, 0xa4,
	0xca, 0xe1, 0xd3, 0xf3, 0x0f, 0x6f, 0x22, 0xaa, 0x2e, 0x53, 0xcf, 0xf6, 0xf9, 0xc2, 0x09, 0x04,
	0x59, 0xce, 0x9d, 0x9a, 0x1f, 0x97, 0xfc, 0x38, 0xe2, 0x4e, 0x32, 0x8f, 0x9c, 0xbf, 0x7e, 0x72,
	0xbd, 0x7e, 0xfe, 0xa9, 0x9d, 0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x58, 0x07, 0x35, 0x35, 0x9e,
	0x05, 0x00, 0x00,
}
