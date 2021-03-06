// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/log_memory.proto

package framework

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MemoryLogStep struct {
	// Process-unique step id.
	StepId int64 `protobuf:"varint,1,opt,name=step_id,json=stepId" json:"step_id,omitempty"`
	// Handle describing the feeds and fetches of the step.
	Handle string `protobuf:"bytes,2,opt,name=handle" json:"handle,omitempty"`
}

func (m *MemoryLogStep) Reset()                    { *m = MemoryLogStep{} }
func (m *MemoryLogStep) String() string            { return proto.CompactTextString(m) }
func (*MemoryLogStep) ProtoMessage()               {}
func (*MemoryLogStep) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *MemoryLogStep) GetStepId() int64 {
	if m != nil {
		return m.StepId
	}
	return 0
}

func (m *MemoryLogStep) GetHandle() string {
	if m != nil {
		return m.Handle
	}
	return ""
}

type MemoryLogTensorAllocation struct {
	// Process-unique step id.
	StepId int64 `protobuf:"varint,1,opt,name=step_id,json=stepId" json:"step_id,omitempty"`
	// Name of the kernel making the allocation as set in GraphDef,
	// e.g., "affine2/weights/Assign".
	KernelName string `protobuf:"bytes,2,opt,name=kernel_name,json=kernelName" json:"kernel_name,omitempty"`
	// Allocated tensor details.
	Tensor *TensorDescription `protobuf:"bytes,3,opt,name=tensor" json:"tensor,omitempty"`
}

func (m *MemoryLogTensorAllocation) Reset()                    { *m = MemoryLogTensorAllocation{} }
func (m *MemoryLogTensorAllocation) String() string            { return proto.CompactTextString(m) }
func (*MemoryLogTensorAllocation) ProtoMessage()               {}
func (*MemoryLogTensorAllocation) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

func (m *MemoryLogTensorAllocation) GetStepId() int64 {
	if m != nil {
		return m.StepId
	}
	return 0
}

func (m *MemoryLogTensorAllocation) GetKernelName() string {
	if m != nil {
		return m.KernelName
	}
	return ""
}

func (m *MemoryLogTensorAllocation) GetTensor() *TensorDescription {
	if m != nil {
		return m.Tensor
	}
	return nil
}

type MemoryLogTensorDeallocation struct {
	// Id of the tensor buffer being deallocated, used to match to a
	// corresponding allocation.
	AllocationId int64 `protobuf:"varint,1,opt,name=allocation_id,json=allocationId" json:"allocation_id,omitempty"`
	// Name of the allocator used.
	AllocatorName string `protobuf:"bytes,2,opt,name=allocator_name,json=allocatorName" json:"allocator_name,omitempty"`
}

func (m *MemoryLogTensorDeallocation) Reset()                    { *m = MemoryLogTensorDeallocation{} }
func (m *MemoryLogTensorDeallocation) String() string            { return proto.CompactTextString(m) }
func (*MemoryLogTensorDeallocation) ProtoMessage()               {}
func (*MemoryLogTensorDeallocation) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

func (m *MemoryLogTensorDeallocation) GetAllocationId() int64 {
	if m != nil {
		return m.AllocationId
	}
	return 0
}

func (m *MemoryLogTensorDeallocation) GetAllocatorName() string {
	if m != nil {
		return m.AllocatorName
	}
	return ""
}

type MemoryLogTensorOutput struct {
	// Process-unique step id.
	StepId int64 `protobuf:"varint,1,opt,name=step_id,json=stepId" json:"step_id,omitempty"`
	// Name of the kernel producing an output as set in GraphDef, e.g.,
	// "affine2/weights/Assign".
	KernelName string `protobuf:"bytes,2,opt,name=kernel_name,json=kernelName" json:"kernel_name,omitempty"`
	// Index of the output being set.
	Index int32 `protobuf:"varint,3,opt,name=index" json:"index,omitempty"`
	// Output tensor details.
	Tensor *TensorDescription `protobuf:"bytes,4,opt,name=tensor" json:"tensor,omitempty"`
}

func (m *MemoryLogTensorOutput) Reset()                    { *m = MemoryLogTensorOutput{} }
func (m *MemoryLogTensorOutput) String() string            { return proto.CompactTextString(m) }
func (*MemoryLogTensorOutput) ProtoMessage()               {}
func (*MemoryLogTensorOutput) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{3} }

func (m *MemoryLogTensorOutput) GetStepId() int64 {
	if m != nil {
		return m.StepId
	}
	return 0
}

func (m *MemoryLogTensorOutput) GetKernelName() string {
	if m != nil {
		return m.KernelName
	}
	return ""
}

func (m *MemoryLogTensorOutput) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *MemoryLogTensorOutput) GetTensor() *TensorDescription {
	if m != nil {
		return m.Tensor
	}
	return nil
}

type MemoryLogRawAllocation struct {
	// Process-unique step id.
	StepId int64 `protobuf:"varint,1,opt,name=step_id,json=stepId" json:"step_id,omitempty"`
	// Name of the operation making the allocation.
	Operation string `protobuf:"bytes,2,opt,name=operation" json:"operation,omitempty"`
	// Number of bytes in the allocation.
	NumBytes int64 `protobuf:"varint,3,opt,name=num_bytes,json=numBytes" json:"num_bytes,omitempty"`
	// Address of the allocation.
	Ptr uint64 `protobuf:"varint,4,opt,name=ptr" json:"ptr,omitempty"`
	// Id of the tensor buffer being allocated, used to match to a
	// corresponding deallocation.
	AllocationId int64 `protobuf:"varint,5,opt,name=allocation_id,json=allocationId" json:"allocation_id,omitempty"`
	// Name of the allocator used.
	AllocatorName string `protobuf:"bytes,6,opt,name=allocator_name,json=allocatorName" json:"allocator_name,omitempty"`
}

func (m *MemoryLogRawAllocation) Reset()                    { *m = MemoryLogRawAllocation{} }
func (m *MemoryLogRawAllocation) String() string            { return proto.CompactTextString(m) }
func (*MemoryLogRawAllocation) ProtoMessage()               {}
func (*MemoryLogRawAllocation) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{4} }

func (m *MemoryLogRawAllocation) GetStepId() int64 {
	if m != nil {
		return m.StepId
	}
	return 0
}

func (m *MemoryLogRawAllocation) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func (m *MemoryLogRawAllocation) GetNumBytes() int64 {
	if m != nil {
		return m.NumBytes
	}
	return 0
}

func (m *MemoryLogRawAllocation) GetPtr() uint64 {
	if m != nil {
		return m.Ptr
	}
	return 0
}

func (m *MemoryLogRawAllocation) GetAllocationId() int64 {
	if m != nil {
		return m.AllocationId
	}
	return 0
}

func (m *MemoryLogRawAllocation) GetAllocatorName() string {
	if m != nil {
		return m.AllocatorName
	}
	return ""
}

type MemoryLogRawDeallocation struct {
	// Process-unique step id.
	StepId int64 `protobuf:"varint,1,opt,name=step_id,json=stepId" json:"step_id,omitempty"`
	// Name of the operation making the deallocation.
	Operation string `protobuf:"bytes,2,opt,name=operation" json:"operation,omitempty"`
	// Id of the tensor buffer being deallocated, used to match to a
	// corresponding allocation.
	AllocationId int64 `protobuf:"varint,3,opt,name=allocation_id,json=allocationId" json:"allocation_id,omitempty"`
	// Name of the allocator used.
	AllocatorName string `protobuf:"bytes,4,opt,name=allocator_name,json=allocatorName" json:"allocator_name,omitempty"`
	// True if the deallocation is queued and will be performed later,
	// e.g. for GPU lazy freeing of buffers.
	Deferred bool `protobuf:"varint,5,opt,name=deferred" json:"deferred,omitempty"`
}

func (m *MemoryLogRawDeallocation) Reset()                    { *m = MemoryLogRawDeallocation{} }
func (m *MemoryLogRawDeallocation) String() string            { return proto.CompactTextString(m) }
func (*MemoryLogRawDeallocation) ProtoMessage()               {}
func (*MemoryLogRawDeallocation) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{5} }

func (m *MemoryLogRawDeallocation) GetStepId() int64 {
	if m != nil {
		return m.StepId
	}
	return 0
}

func (m *MemoryLogRawDeallocation) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func (m *MemoryLogRawDeallocation) GetAllocationId() int64 {
	if m != nil {
		return m.AllocationId
	}
	return 0
}

func (m *MemoryLogRawDeallocation) GetAllocatorName() string {
	if m != nil {
		return m.AllocatorName
	}
	return ""
}

func (m *MemoryLogRawDeallocation) GetDeferred() bool {
	if m != nil {
		return m.Deferred
	}
	return false
}

func init() {
	proto.RegisterType((*MemoryLogStep)(nil), "tensorflow.MemoryLogStep")
	proto.RegisterType((*MemoryLogTensorAllocation)(nil), "tensorflow.MemoryLogTensorAllocation")
	proto.RegisterType((*MemoryLogTensorDeallocation)(nil), "tensorflow.MemoryLogTensorDeallocation")
	proto.RegisterType((*MemoryLogTensorOutput)(nil), "tensorflow.MemoryLogTensorOutput")
	proto.RegisterType((*MemoryLogRawAllocation)(nil), "tensorflow.MemoryLogRawAllocation")
	proto.RegisterType((*MemoryLogRawDeallocation)(nil), "tensorflow.MemoryLogRawDeallocation")
}

func init() { proto.RegisterFile("tensorflow/core/framework/log_memory.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xb5, 0x38, 0x31, 0xc9, 0x94, 0x00, 0xb2, 0xa0, 0x98, 0x16, 0x44, 0x14, 0x84, 0x14,
	0x21, 0x35, 0x96, 0x8a, 0xb8, 0x43, 0x94, 0x4b, 0xa5, 0x02, 0x95, 0xe1, 0xc4, 0xc5, 0x72, 0xe2,
	0x89, 0x6b, 0xc5, 0xbb, 0xb3, 0x5a, 0xaf, 0x1b, 0xf2, 0x0e, 0x3c, 0x03, 0xef, 0xc1, 0x2b, 0xf0,
	0x44, 0x1c, 0x91, 0xbd, 0xa9, 0x37, 0xb4, 0x8d, 0x14, 0xe8, 0xcd, 0xff, 0x78, 0x67, 0xe6, 0xfb,
	0x67, 0x34, 0xf0, 0x5a, 0xa3, 0x28, 0x48, 0xcd, 0x73, 0x5a, 0x06, 0x33, 0x52, 0x18, 0xcc, 0x55,
	0xcc, 0x71, 0x49, 0x6a, 0x11, 0xe4, 0x94, 0x46, 0x1c, 0x39, 0xa9, 0xd5, 0x48, 0x2a, 0xd2, 0xe4,
	0x81, 0x7d, 0x7b, 0x70, 0xbc, 0x3d, 0xcf, 0xfc, 0x89, 0x12, 0x2c, 0x66, 0x2a, 0x93, 0x3a, 0x23,
	0x61, 0xf2, 0x07, 0xef, 0xa0, 0xf7, 0xa1, 0xae, 0x77, 0x4a, 0xe9, 0x67, 0x8d, 0xd2, 0x7b, 0x02,
	0x77, 0x0b, 0x8d, 0x32, 0xca, 0x12, 0x9f, 0xf5, 0xd9, 0xd0, 0x09, 0xdd, 0x4a, 0x9e, 0x24, 0xde,
	0x3e, 0xb8, 0xe7, 0xb1, 0x48, 0x72, 0xf4, 0xef, 0xf4, 0xd9, 0xb0, 0x1b, 0xae, 0xd5, 0xe0, 0x3b,
	0x83, 0xa7, 0x4d, 0x89, 0x2f, 0x75, 0x9f, 0xf7, 0x79, 0x4e, 0xb3, 0xb8, 0xea, 0xb2, 0xbd, 0xdc,
	0x0b, 0xd8, 0x5b, 0xa0, 0x12, 0x98, 0x47, 0x22, 0xe6, 0x97, 0x35, 0xc1, 0x84, 0x3e, 0xc6, 0x1c,
	0xbd, 0xb7, 0xe0, 0x1a, 0x6a, 0xdf, 0xe9, 0xb3, 0xe1, 0xde, 0xf1, 0xf3, 0x91, 0xb5, 0x37, 0x32,
	0x7d, 0x26, 0xd6, 0x4e, 0xb8, 0x7e, 0x3c, 0xc8, 0xe0, 0xf0, 0x0a, 0xcd, 0x04, 0x63, 0xcb, 0xf3,
	0x12, 0x7a, 0x56, 0x59, 0xaa, 0x7b, 0x36, 0x78, 0x92, 0x78, 0xaf, 0xe0, 0xfe, 0x5a, 0x93, 0xda,
	0xc4, 0xeb, 0x35, 0xd1, 0x8a, 0x70, 0xf0, 0x83, 0xc1, 0xe3, 0x2b, 0xbd, 0x3e, 0x95, 0x5a, 0x96,
	0xfa, 0x16, 0xae, 0x1f, 0x41, 0x3b, 0x13, 0x09, 0x7e, 0xab, 0x4d, 0xb7, 0x43, 0x23, 0x36, 0x66,
	0xd1, 0xfa, 0x97, 0x59, 0xfc, 0x62, 0xb0, 0xdf, 0x00, 0x86, 0xf1, 0x72, 0x97, 0xbd, 0x3c, 0x83,
	0x2e, 0x49, 0x54, 0xf5, 0xab, 0x35, 0x9f, 0x0d, 0x78, 0x87, 0xd0, 0x15, 0x25, 0x8f, 0xa6, 0x2b,
	0x8d, 0x45, 0x8d, 0xe8, 0x84, 0x1d, 0x51, 0xf2, 0x71, 0xa5, 0xbd, 0x87, 0xe0, 0x48, 0x6d, 0x10,
	0x5b, 0x61, 0xf5, 0x79, 0x7d, 0xda, 0xed, 0x9d, 0xa6, 0xed, 0xde, 0x34, 0xed, 0x9f, 0x0c, 0xfc,
	0x4d, 0x33, 0x7f, 0xad, 0xf5, 0x3f, 0xed, 0x5c, 0xe3, 0x73, 0x76, 0xe2, 0x6b, 0xdd, 0xc0, 0xe7,
	0x1d, 0x40, 0x27, 0xc1, 0x39, 0x2a, 0x85, 0xc6, 0x66, 0x27, 0x6c, 0xf4, 0xf8, 0x02, 0x7c, 0x52,
	0xe9, 0xe6, 0xd2, 0x9a, 0xd3, 0x1c, 0x3f, 0x38, 0xa5, 0xd4, 0xf8, 0x3a, 0xab, 0x2e, 0xb2, 0x38,
	0x63, 0x5f, 0x27, 0x69, 0xa6, 0xcf, 0xcb, 0xe9, 0x68, 0x46, 0x3c, 0x48, 0x14, 0xae, 0x2e, 0xef,
	0xb7, 0xca, 0x3c, 0x2a, 0x50, 0x5d, 0x64, 0x22, 0x3d, 0x4a, 0x29, 0x90, 0x8b, 0x34, 0xd8, 0x7a,
	0xf3, 0xbf, 0x19, 0x9b, 0xba, 0xf5, 0x91, 0xbf, 0xf9, 0x13, 0x00, 0x00, 0xff, 0xff, 0x00, 0xa9,
	0x77, 0x91, 0x52, 0x04, 0x00, 0x00,
}
