// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/record/recordpb/recordpb.proto

/*
Package recordpb is a generated protocol buffer package.

It is generated from these files:
	pkg/record/recordpb/recordpb.proto

It has these top-level messages:
	Data
	Record
*/
package recordpb

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

// Data represents test data.
type Data struct {
	Endpoint      string   `protobuf:"bytes,1,opt,name=endpoint" json:"endpoint,omitempty"`
	ClusterSize   uint32   `protobuf:"varint,2,opt,name=clusterSize" json:"clusterSize,omitempty"`
	Started       string   `protobuf:"bytes,3,opt,name=started" json:"started,omitempty"`
	Current       uint64   `protobuf:"varint,4,opt,name=current" json:"current,omitempty"`
	CurrentFailed uint64   `protobuf:"varint,5,opt,name=currentFailed" json:"currentFailed,omitempty"`
	FailureCases  []string `protobuf:"bytes,6,rep,name=failureCases" json:"failureCases,omitempty"`
}

func (m *Data) Reset()                    { *m = Data{} }
func (m *Data) String() string            { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()               {}
func (*Data) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Data) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *Data) GetClusterSize() uint32 {
	if m != nil {
		return m.ClusterSize
	}
	return 0
}

func (m *Data) GetStarted() string {
	if m != nil {
		return m.Started
	}
	return ""
}

func (m *Data) GetCurrent() uint64 {
	if m != nil {
		return m.Current
	}
	return 0
}

func (m *Data) GetCurrentFailed() uint64 {
	if m != nil {
		return m.CurrentFailed
	}
	return 0
}

func (m *Data) GetFailureCases() []string {
	if m != nil {
		return m.FailureCases
	}
	return nil
}

// Record represents record data.
type Record struct {
	Total    uint64  `protobuf:"varint,1,opt,name=total" json:"total,omitempty"`
	TestData []*Data `protobuf:"bytes,2,rep,name=testData" json:"testData,omitempty"`
}

func (m *Record) Reset()                    { *m = Record{} }
func (m *Record) String() string            { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()               {}
func (*Record) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Record) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Record) GetTestData() []*Data {
	if m != nil {
		return m.TestData
	}
	return nil
}

func init() {
	proto.RegisterType((*Data)(nil), "recordpb.Data")
	proto.RegisterType((*Record)(nil), "recordpb.Record")
}

func init() { proto.RegisterFile("pkg/record/recordpb/recordpb.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xe5, 0x26, 0x0d, 0xe9, 0x95, 0x32, 0x58, 0x0c, 0x27, 0x26, 0x2b, 0x62, 0xb0, 0x18,
	0x8a, 0x04, 0x8f, 0x00, 0x62, 0x60, 0x34, 0x4f, 0xe0, 0x26, 0x07, 0x8a, 0x88, 0xe2, 0xe8, 0x7c,
	0x59, 0x78, 0x3c, 0x9e, 0x0c, 0xc5, 0x34, 0x51, 0x33, 0xf9, 0xbe, 0xef, 0x97, 0x87, 0xff, 0x87,
	0x6a, 0xf8, 0xfe, 0x7a, 0x64, 0xaa, 0x03, 0x37, 0xe7, 0x67, 0x38, 0x2d, 0xc7, 0x71, 0xe0, 0x20,
	0x41, 0x97, 0x33, 0x57, 0xbf, 0x0a, 0xf2, 0x57, 0x2f, 0x5e, 0xdf, 0x41, 0x49, 0x7d, 0x33, 0x84,
	0xb6, 0x17, 0x54, 0x46, 0xd9, 0x9d, 0x5b, 0x58, 0x1b, 0xd8, 0xd7, 0xdd, 0x18, 0x85, 0xf8, 0xa3,
	0xfd, 0x21, 0xdc, 0x18, 0x65, 0x0f, 0xee, 0x52, 0x69, 0x84, 0xab, 0x28, 0x9e, 0x85, 0x1a, 0xcc,
	0xd2, 0xe7, 0x19, 0xa7, 0xa4, 0x1e, 0x99, 0xa9, 0x17, 0xcc, 0x8d, 0xb2, 0xb9, 0x9b, 0x51, 0xdf,
	0xc3, 0xe1, 0x7c, 0xbe, 0xf9, 0xb6, 0xa3, 0x06, 0xb7, 0x29, 0x5f, 0x4b, 0x5d, 0xc1, 0xf5, 0xa7,
	0x6f, 0xbb, 0x91, 0xe9, 0xc5, 0x47, 0x8a, 0x58, 0x98, 0xcc, 0xee, 0xdc, 0xca, 0x55, 0xef, 0x50,
	0xb8, 0x54, 0x48, 0xdf, 0xc2, 0x56, 0x82, 0xf8, 0x2e, 0x55, 0xc8, 0xdd, 0x3f, 0xe8, 0x07, 0x28,
	0x85, 0xa2, 0x4c, 0x3d, 0x71, 0x63, 0x32, 0xbb, 0x7f, 0xba, 0x39, 0x2e, 0x8b, 0x4c, 0xd6, 0x2d,
	0xf9, 0xa9, 0x48, 0x0b, 0x3d, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0xc6, 0x06, 0x54, 0x7e, 0x47,
	0x01, 0x00, 0x00,
}
