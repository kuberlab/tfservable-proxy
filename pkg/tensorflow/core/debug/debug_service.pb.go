// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/debug/debug_service.proto

/*
Package debug is a generated protocol buffer package.

It is generated from these files:
	tensorflow/core/debug/debug_service.proto
	tensorflow/core/debug/debugger_event_metadata.proto

It has these top-level messages:
	EventReply
	DebuggerEventMetadata
*/
package debug

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import tensorflow5 "github.com/kuberlab/tfservable-proxy/pkg/tensorflow/core/util"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EventReply_DebugOpStateChange_Change int32

const (
	EventReply_DebugOpStateChange_DISABLE EventReply_DebugOpStateChange_Change = 0
	EventReply_DebugOpStateChange_ENABLE  EventReply_DebugOpStateChange_Change = 1
)

var EventReply_DebugOpStateChange_Change_name = map[int32]string{
	0: "DISABLE",
	1: "ENABLE",
}
var EventReply_DebugOpStateChange_Change_value = map[string]int32{
	"DISABLE": 0,
	"ENABLE":  1,
}

func (x EventReply_DebugOpStateChange_Change) String() string {
	return proto.EnumName(EventReply_DebugOpStateChange_Change_name, int32(x))
}
func (EventReply_DebugOpStateChange_Change) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 0}
}

// Reply message from EventListener to the client, i.e., to the source of the
// Event protocol buffers, e.g., debug ops inserted by a debugged runtime to a
// TensorFlow graph being executed.
type EventReply struct {
	DebugOpStateChanges []*EventReply_DebugOpStateChange `protobuf:"bytes,1,rep,name=debug_op_state_changes,json=debugOpStateChanges" json:"debug_op_state_changes,omitempty"`
}

func (m *EventReply) Reset()                    { *m = EventReply{} }
func (m *EventReply) String() string            { return proto.CompactTextString(m) }
func (*EventReply) ProtoMessage()               {}
func (*EventReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *EventReply) GetDebugOpStateChanges() []*EventReply_DebugOpStateChange {
	if m != nil {
		return m.DebugOpStateChanges
	}
	return nil
}

type EventReply_DebugOpStateChange struct {
	Change     EventReply_DebugOpStateChange_Change `protobuf:"varint,1,opt,name=change,enum=tensorflow.EventReply_DebugOpStateChange_Change" json:"change,omitempty"`
	NodeName   string                               `protobuf:"bytes,2,opt,name=node_name,json=nodeName" json:"node_name,omitempty"`
	OutputSlot int32                                `protobuf:"varint,3,opt,name=output_slot,json=outputSlot" json:"output_slot,omitempty"`
	DebugOp    string                               `protobuf:"bytes,4,opt,name=debug_op,json=debugOp" json:"debug_op,omitempty"`
}

func (m *EventReply_DebugOpStateChange) Reset()         { *m = EventReply_DebugOpStateChange{} }
func (m *EventReply_DebugOpStateChange) String() string { return proto.CompactTextString(m) }
func (*EventReply_DebugOpStateChange) ProtoMessage()    {}
func (*EventReply_DebugOpStateChange) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

func (m *EventReply_DebugOpStateChange) GetChange() EventReply_DebugOpStateChange_Change {
	if m != nil {
		return m.Change
	}
	return EventReply_DebugOpStateChange_DISABLE
}

func (m *EventReply_DebugOpStateChange) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *EventReply_DebugOpStateChange) GetOutputSlot() int32 {
	if m != nil {
		return m.OutputSlot
	}
	return 0
}

func (m *EventReply_DebugOpStateChange) GetDebugOp() string {
	if m != nil {
		return m.DebugOp
	}
	return ""
}

func init() {
	proto.RegisterType((*EventReply)(nil), "tensorflow.EventReply")
	proto.RegisterType((*EventReply_DebugOpStateChange)(nil), "tensorflow.EventReply.DebugOpStateChange")
	proto.RegisterEnum("tensorflow.EventReply_DebugOpStateChange_Change", EventReply_DebugOpStateChange_Change_name, EventReply_DebugOpStateChange_Change_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EventListener service

type EventListenerClient interface {
	// Client(s) can use this RPC method to send the EventListener Event protos.
	// The Event protos can hold information such as:
	//   1) intermediate tensors from a debugged graph being executed, which can
	//      be sent from DebugIdentity ops configured with grpc URLs.
	//   2) GraphDefs of partition graphs, which can be sent from special debug
	//      ops that get executed immediately after the beginning of the graph
	//      execution.
	SendEvents(ctx context.Context, opts ...grpc.CallOption) (EventListener_SendEventsClient, error)
}

type eventListenerClient struct {
	cc *grpc.ClientConn
}

func NewEventListenerClient(cc *grpc.ClientConn) EventListenerClient {
	return &eventListenerClient{cc}
}

func (c *eventListenerClient) SendEvents(ctx context.Context, opts ...grpc.CallOption) (EventListener_SendEventsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_EventListener_serviceDesc.Streams[0], c.cc, "/tensorflow.EventListener/SendEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventListenerSendEventsClient{stream}
	return x, nil
}

type EventListener_SendEventsClient interface {
	Send(*tensorflow5.Event) error
	Recv() (*EventReply, error)
	grpc.ClientStream
}

type eventListenerSendEventsClient struct {
	grpc.ClientStream
}

func (x *eventListenerSendEventsClient) Send(m *tensorflow5.Event) error {
	return x.ClientStream.SendMsg(m)
}

func (x *eventListenerSendEventsClient) Recv() (*EventReply, error) {
	m := new(EventReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for EventListener service

type EventListenerServer interface {
	// Client(s) can use this RPC method to send the EventListener Event protos.
	// The Event protos can hold information such as:
	//   1) intermediate tensors from a debugged graph being executed, which can
	//      be sent from DebugIdentity ops configured with grpc URLs.
	//   2) GraphDefs of partition graphs, which can be sent from special debug
	//      ops that get executed immediately after the beginning of the graph
	//      execution.
	SendEvents(EventListener_SendEventsServer) error
}

func RegisterEventListenerServer(s *grpc.Server, srv EventListenerServer) {
	s.RegisterService(&_EventListener_serviceDesc, srv)
}

func _EventListener_SendEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EventListenerServer).SendEvents(&eventListenerSendEventsServer{stream})
}

type EventListener_SendEventsServer interface {
	Send(*EventReply) error
	Recv() (*tensorflow5.Event, error)
	grpc.ServerStream
}

type eventListenerSendEventsServer struct {
	grpc.ServerStream
}

func (x *eventListenerSendEventsServer) Send(m *EventReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *eventListenerSendEventsServer) Recv() (*tensorflow5.Event, error) {
	m := new(tensorflow5.Event)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _EventListener_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tensorflow.EventListener",
	HandlerType: (*EventListenerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendEvents",
			Handler:       _EventListener_SendEvents_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "tensorflow/core/debug/debug_service.proto",
}

func init() { proto.RegisterFile("tensorflow/core/debug/debug_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x86, 0xef, 0xc0, 0xbd, 0x05, 0x0e, 0xb9, 0x46, 0xc7, 0x84, 0x54, 0x5c, 0x58, 0x59, 0x95,
	0x05, 0x2d, 0xc1, 0xa5, 0x1b, 0x45, 0x48, 0x34, 0x21, 0x98, 0xb4, 0x3b, 0x17, 0x36, 0xa5, 0x3d,
	0x96, 0x86, 0x32, 0xd3, 0x74, 0xa6, 0x18, 0x1e, 0xd1, 0x57, 0xf0, 0x69, 0x4c, 0xa7, 0x35, 0x4d,
	0x84, 0x85, 0x9b, 0xc9, 0xcc, 0x9f, 0xff, 0x3b, 0xe7, 0xfc, 0x39, 0x03, 0x43, 0x89, 0x4c, 0xf0,
	0xec, 0x2d, 0xe1, 0xef, 0x76, 0xc0, 0x33, 0xb4, 0x43, 0x5c, 0xe5, 0x51, 0x79, 0x7a, 0x02, 0xb3,
	0x5d, 0x1c, 0xa0, 0x95, 0x66, 0x5c, 0x72, 0x0a, 0xb5, 0xb5, 0x6f, 0xfc, 0xc4, 0x72, 0x19, 0x27,
	0x36, 0xee, 0x90, 0xc9, 0xd2, 0x3d, 0xf8, 0x68, 0x00, 0xcc, 0x8b, 0xb7, 0x83, 0x69, 0xb2, 0xa7,
	0xaf, 0xd0, 0x2b, 0x6b, 0xf2, 0xd4, 0x13, 0xd2, 0x97, 0xe8, 0x05, 0x6b, 0x9f, 0x45, 0x28, 0x74,
	0x62, 0x34, 0xcd, 0xee, 0x64, 0x68, 0xd5, 0x15, 0xad, 0x9a, 0xb3, 0x66, 0x05, 0xf4, 0x9c, 0xba,
	0x05, 0xf2, 0xa0, 0x08, 0xe7, 0x3c, 0x3c, 0xd0, 0x44, 0xff, 0x93, 0x00, 0x3d, 0xf4, 0xd2, 0x47,
	0xd0, 0xca, 0x3e, 0x3a, 0x31, 0x88, 0x79, 0x32, 0x19, 0xff, 0xba, 0x8d, 0x55, 0x75, 0xab, 0x78,
	0x7a, 0x09, 0x1d, 0xc6, 0x43, 0xf4, 0x98, 0xbf, 0x45, 0xbd, 0x61, 0x10, 0xb3, 0xe3, 0xb4, 0x0b,
	0x61, 0xe9, 0x6f, 0x91, 0x5e, 0x41, 0x97, 0xe7, 0x32, 0xcd, 0xa5, 0x27, 0x12, 0x2e, 0xf5, 0xa6,
	0x41, 0xcc, 0x7f, 0x0e, 0x94, 0x92, 0x9b, 0x70, 0x49, 0x2f, 0xa0, 0xfd, 0x1d, 0x5f, 0xff, 0xab,
	0xe0, 0x56, 0x95, 0x62, 0x70, 0x0d, 0x5a, 0x35, 0x6c, 0x17, 0x5a, 0xb3, 0x27, 0xf7, 0x7e, 0xba,
	0x98, 0x9f, 0xfe, 0xa1, 0x00, 0xda, 0x7c, 0xa9, 0xee, 0x64, 0xb2, 0x80, 0xff, 0x6a, 0xd6, 0x45,
	0x2c, 0x24, 0x32, 0xcc, 0xe8, 0x2d, 0x80, 0x8b, 0x2c, 0x54, 0xa2, 0xa0, 0x67, 0x07, 0xa1, 0xfa,
	0xbd, 0xe3, 0x39, 0x4d, 0x32, 0x26, 0xd3, 0xe9, 0xcb, 0x5d, 0x14, 0xcb, 0x75, 0xbe, 0xb2, 0x02,
	0xbe, 0xb5, 0xc3, 0x0c, 0xf7, 0x1b, 0xbb, 0x76, 0x8f, 0xd4, 0xd2, 0x59, 0x34, 0x8a, 0xb8, 0x9d,
	0x6e, 0x22, 0xfb, 0xe8, 0xff, 0x58, 0x69, 0x6a, 0xc9, 0x37, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x7e, 0xc6, 0x48, 0x10, 0x3f, 0x02, 0x00, 0x00,
}