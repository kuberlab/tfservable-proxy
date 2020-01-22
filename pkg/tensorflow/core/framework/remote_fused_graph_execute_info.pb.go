// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/remote_fused_graph_execute_info.proto

package framework

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RemoteFusedGraphExecuteInfo_NodeType int32

const (
	RemoteFusedGraphExecuteInfo_UNUSED        RemoteFusedGraphExecuteInfo_NodeType = 0
	RemoteFusedGraphExecuteInfo_GRAPH_INPUT   RemoteFusedGraphExecuteInfo_NodeType = 1
	RemoteFusedGraphExecuteInfo_GRAPH_OUTPUT  RemoteFusedGraphExecuteInfo_NodeType = 2
	RemoteFusedGraphExecuteInfo_FUSED_NODE    RemoteFusedGraphExecuteInfo_NodeType = 3
	RemoteFusedGraphExecuteInfo_BORDER_INPUT  RemoteFusedGraphExecuteInfo_NodeType = 4
	RemoteFusedGraphExecuteInfo_BORDER_OUTPUT RemoteFusedGraphExecuteInfo_NodeType = 5
)

var RemoteFusedGraphExecuteInfo_NodeType_name = map[int32]string{
	0: "UNUSED",
	1: "GRAPH_INPUT",
	2: "GRAPH_OUTPUT",
	3: "FUSED_NODE",
	4: "BORDER_INPUT",
	5: "BORDER_OUTPUT",
}
var RemoteFusedGraphExecuteInfo_NodeType_value = map[string]int32{
	"UNUSED":        0,
	"GRAPH_INPUT":   1,
	"GRAPH_OUTPUT":  2,
	"FUSED_NODE":    3,
	"BORDER_INPUT":  4,
	"BORDER_OUTPUT": 5,
}

func (x RemoteFusedGraphExecuteInfo_NodeType) String() string {
	return proto.EnumName(RemoteFusedGraphExecuteInfo_NodeType_name, int32(x))
}
func (RemoteFusedGraphExecuteInfo_NodeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor13, []int{0, 0}
}

// Protocol buffer representing a handle to a tensorflow resource. Handles are
// not valid across executions, but can be serialized back and forth from within
// a single run.
type RemoteFusedGraphExecuteInfo struct {
	// Definition of remote graph
	RemoteGraph *GraphDef `protobuf:"bytes,1,opt,name=remote_graph,json=remoteGraph" json:"remote_graph,omitempty"`
	// Remote fused graph input node name
	GraphInputNodeName []string `protobuf:"bytes,2,rep,name=graph_input_node_name,json=graphInputNodeName" json:"graph_input_node_name,omitempty"`
	// Remote fused graph output node name
	GraphOutputNodeName []string `protobuf:"bytes,3,rep,name=graph_output_node_name,json=graphOutputNodeName" json:"graph_output_node_name,omitempty"`
	// Executor's name
	ExecutorName string `protobuf:"bytes,4,opt,name=executor_name,json=executorName" json:"executor_name,omitempty"`
	// Optional: Parameters given to the executor
	SerializedExecutorParameters []byte `protobuf:"bytes,5,opt,name=serialized_executor_parameters,json=serializedExecutorParameters,proto3" json:"serialized_executor_parameters,omitempty"`
	// Optional: Default graph input tensor shape used to allocate memory
	// before executing op
	DefaultGraphInputTensorShape []*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto `protobuf:"bytes,6,rep,name=default_graph_input_tensor_shape,json=defaultGraphInputTensorShape" json:"default_graph_input_tensor_shape,omitempty"`
	// Optional: Default graph input tensor shape used to allocate memory
	// before executing op
	// TODO(satok): Remote output tensor shape once shape information is stored
	// in NodeDef
	DefaultGraphOutputTensorShape []*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto `protobuf:"bytes,7,rep,name=default_graph_output_tensor_shape,json=defaultGraphOutputTensorShape" json:"default_graph_output_tensor_shape,omitempty"`
}

func (m *RemoteFusedGraphExecuteInfo) Reset()                    { *m = RemoteFusedGraphExecuteInfo{} }
func (m *RemoteFusedGraphExecuteInfo) String() string            { return proto.CompactTextString(m) }
func (*RemoteFusedGraphExecuteInfo) ProtoMessage()               {}
func (*RemoteFusedGraphExecuteInfo) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func (m *RemoteFusedGraphExecuteInfo) GetRemoteGraph() *GraphDef {
	if m != nil {
		return m.RemoteGraph
	}
	return nil
}

func (m *RemoteFusedGraphExecuteInfo) GetGraphInputNodeName() []string {
	if m != nil {
		return m.GraphInputNodeName
	}
	return nil
}

func (m *RemoteFusedGraphExecuteInfo) GetGraphOutputNodeName() []string {
	if m != nil {
		return m.GraphOutputNodeName
	}
	return nil
}

func (m *RemoteFusedGraphExecuteInfo) GetExecutorName() string {
	if m != nil {
		return m.ExecutorName
	}
	return ""
}

func (m *RemoteFusedGraphExecuteInfo) GetSerializedExecutorParameters() []byte {
	if m != nil {
		return m.SerializedExecutorParameters
	}
	return nil
}

func (m *RemoteFusedGraphExecuteInfo) GetDefaultGraphInputTensorShape() []*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto {
	if m != nil {
		return m.DefaultGraphInputTensorShape
	}
	return nil
}

func (m *RemoteFusedGraphExecuteInfo) GetDefaultGraphOutputTensorShape() []*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto {
	if m != nil {
		return m.DefaultGraphOutputTensorShape
	}
	return nil
}

type RemoteFusedGraphExecuteInfo_TensorShapeTypeProto struct {
	Dtype DataType          `protobuf:"varint,1,opt,name=dtype,enum=tensorflow.DataType" json:"dtype,omitempty"`
	Shape *TensorShapeProto `protobuf:"bytes,2,opt,name=shape" json:"shape,omitempty"`
}

func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) Reset() {
	*m = RemoteFusedGraphExecuteInfo_TensorShapeTypeProto{}
}
func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) String() string {
	return proto.CompactTextString(m)
}
func (*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) ProtoMessage() {}
func (*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) Descriptor() ([]byte, []int) {
	return fileDescriptor13, []int{0, 0}
}

func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) GetDtype() DataType {
	if m != nil {
		return m.Dtype
	}
	return DataType_DT_INVALID
}

func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) GetShape() *TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func init() {
	proto.RegisterType((*RemoteFusedGraphExecuteInfo)(nil), "tensorflow.RemoteFusedGraphExecuteInfo")
	proto.RegisterType((*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto)(nil), "tensorflow.RemoteFusedGraphExecuteInfo.TensorShapeTypeProto")
	proto.RegisterEnum("tensorflow.RemoteFusedGraphExecuteInfo_NodeType", RemoteFusedGraphExecuteInfo_NodeType_name, RemoteFusedGraphExecuteInfo_NodeType_value)
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/remote_fused_graph_execute_info.proto", fileDescriptor13)
}

var fileDescriptor13 = []byte{
	// 542 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xc1, 0x8f, 0xd2, 0x40,
	0x14, 0xc6, 0x2d, 0x2c, 0xe8, 0x3e, 0xd8, 0x15, 0xc7, 0xd5, 0x10, 0x44, 0x53, 0xd7, 0x98, 0x10,
	0xe3, 0xd2, 0xc8, 0x1e, 0xbc, 0x98, 0x18, 0x49, 0x59, 0xe4, 0x02, 0xa4, 0x0b, 0x17, 0x2f, 0x93,
	0x2e, 0x7d, 0x2d, 0xcd, 0x42, 0xa7, 0x99, 0x4e, 0x77, 0xc5, 0xb3, 0xd1, 0xbf, 0xc3, 0xff, 0xd2,
	0xa3, 0x99, 0x99, 0x2e, 0x94, 0x44, 0x38, 0xed, 0xb1, 0xf3, 0x7e, 0xdf, 0xeb, 0x37, 0xdf, 0xbc,
	0x07, 0x9f, 0x05, 0x46, 0x09, 0xe3, 0xfe, 0x82, 0xdd, 0x5a, 0x33, 0xc6, 0xd1, 0xf2, 0xb9, 0xbb,
	0xc4, 0x5b, 0xc6, 0xaf, 0x2d, 0x8e, 0x4b, 0x26, 0x90, 0xfa, 0x69, 0x82, 0x1e, 0x0d, 0xb8, 0x1b,
	0xcf, 0x29, 0x7e, 0xc7, 0x59, 0x2a, 0x90, 0x86, 0x91, 0xcf, 0xda, 0x31, 0x67, 0x82, 0x11, 0xd8,
	0x34, 0x68, 0xbc, 0xdd, 0xdd, 0x4c, 0xe9, 0xb5, 0xa4, 0xf1, 0x7e, 0x37, 0xa6, 0x2b, 0x34, 0x99,
	0xbb, 0x31, 0x66, 0xf4, 0x9e, 0xa6, 0x62, 0x15, 0x63, 0xa2, 0xb1, 0xd3, 0x3f, 0x65, 0x78, 0xe1,
	0x28, 0xc7, 0x17, 0xd2, 0x70, 0x5f, 0xfe, 0xaf, 0xa7, 0xed, 0x0e, 0x22, 0x9f, 0x91, 0x8f, 0x50,
	0xcd, 0x2e, 0xa4, 0xac, 0xd4, 0x0d, 0xd3, 0x68, 0x55, 0x3a, 0x27, 0xed, 0x4d, 0xf7, 0xb6, 0xd2,
	0xd8, 0xe8, 0x3b, 0x15, 0x4d, 0xaa, 0x6f, 0xf2, 0x01, 0x9e, 0xe9, 0xcb, 0x87, 0x51, 0x9c, 0x0a,
	0x1a, 0x31, 0x0f, 0x69, 0xe4, 0x2e, 0xb1, 0x5e, 0x30, 0x8b, 0xad, 0x43, 0x87, 0xa8, 0xe2, 0x40,
	0xd6, 0x86, 0xcc, 0xc3, 0xa1, 0xbb, 0x44, 0x72, 0x0e, 0xcf, 0xb5, 0x84, 0xa5, 0x62, 0x5b, 0x53,
	0x54, 0x9a, 0xa7, 0xaa, 0x3a, 0x52, 0xc5, 0xb5, 0xe8, 0x0d, 0x1c, 0xe9, 0x78, 0x19, 0xd7, 0xec,
	0x81, 0x69, 0xb4, 0x0e, 0x9d, 0xea, 0xdd, 0xa1, 0x82, 0x6c, 0x78, 0x95, 0x20, 0x0f, 0xdd, 0x45,
	0xf8, 0x03, 0x3d, 0xba, 0xe6, 0x63, 0x57, 0x66, 0x22, 0x90, 0x27, 0xf5, 0x92, 0x69, 0xb4, 0xaa,
	0x4e, 0x73, 0x43, 0xf5, 0x32, 0x68, 0xbc, 0x66, 0xc8, 0x4f, 0x03, 0x4c, 0x0f, 0x7d, 0x37, 0x5d,
	0x08, 0x9a, 0xbf, 0x5b, 0x3e, 0xfd, 0x7a, 0xd9, 0x2c, 0xb6, 0x2a, 0x9d, 0x4f, 0xf9, 0x80, 0xf6,
	0xe4, 0xdb, 0x9e, 0x28, 0xec, 0x52, 0x4a, 0x27, 0xab, 0x18, 0xc7, 0xf2, 0x51, 0x9c, 0x66, 0xf6,
	0x97, 0xfe, 0x3a, 0xa3, 0x1c, 0x46, 0x7e, 0x19, 0xf0, 0x7a, 0xdb, 0x46, 0x96, 0xd7, 0x96, 0x8f,
	0x87, 0xf7, 0xe0, 0xe3, 0x65, 0xde, 0x87, 0xce, 0x3d, 0xc7, 0x35, 0x6e, 0xe0, 0xe4, 0x7f, 0x32,
	0xf2, 0x0e, 0x4a, 0x9e, 0x9c, 0x31, 0x35, 0x2c, 0xc7, 0xdb, 0xc3, 0x62, 0xbb, 0xc2, 0x95, 0xa4,
	0xa3, 0x11, 0xd2, 0x81, 0x92, 0xf6, 0x5b, 0x50, 0x83, 0xd5, 0xcc, 0xb3, 0xb9, 0xe6, 0xda, 0x8f,
	0x46, 0x4f, 0x23, 0x78, 0x24, 0x9f, 0x5f, 0xb6, 0x21, 0x00, 0xe5, 0xe9, 0x70, 0x7a, 0xd9, 0xb3,
	0x6b, 0x0f, 0xc8, 0x63, 0xa8, 0xf4, 0x9d, 0x2f, 0xe3, 0xaf, 0x74, 0x30, 0x1c, 0x4f, 0x27, 0x35,
	0x83, 0xd4, 0xa0, 0xaa, 0x0f, 0x46, 0xd3, 0x89, 0x3c, 0x29, 0x90, 0x63, 0x80, 0x0b, 0x49, 0xd3,
	0xe1, 0xc8, 0xee, 0xd5, 0x8a, 0x92, 0xe8, 0x8e, 0x1c, 0xbb, 0xe7, 0x64, 0x9a, 0x03, 0xf2, 0x04,
	0x8e, 0xb2, 0x93, 0x4c, 0x54, 0xea, 0xfe, 0x36, 0xa0, 0xce, 0x78, 0x90, 0xb7, 0xb6, 0x5e, 0xa6,
	0xae, 0xb9, 0x27, 0x55, 0xe5, 0x7a, 0x6c, 0x7c, 0xb3, 0x83, 0x50, 0xcc, 0xd3, 0xab, 0xf6, 0x8c,
	0x2d, 0x2d, 0x8f, 0xe3, 0xea, 0x6e, 0x61, 0x65, 0xab, 0xb3, 0x04, 0xf9, 0x4d, 0x18, 0x05, 0x67,
	0x01, 0xb3, 0xe2, 0xeb, 0xc0, 0xda, 0xb9, 0xb6, 0x7f, 0x0d, 0xe3, 0xaa, 0xac, 0x96, 0xf6, 0xfc,
	0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x10, 0xc3, 0x2a, 0x07, 0x7f, 0x04, 0x00, 0x00,
}
