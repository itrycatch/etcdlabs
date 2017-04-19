// Code generated by protoc-gen-gogo.
// source: clusterpb/clusterpb.proto
// DO NOT EDIT!

/*
	Package clusterpb is a generated protocol buffer package.

	It is generated from these files:
		clusterpb/clusterpb.proto

	It has these top-level messages:
		MemberStatus
*/
package clusterpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// MemberStatus defines node status information.
// Keep the json tag to make it parsable by Typescript.
type MemberStatus struct {
	Name      string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	ID        string `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
	Endpoint  string `protobuf:"bytes,3,opt,name=Endpoint,proto3" json:"Endpoint,omitempty"`
	IsLeader  bool   `protobuf:"varint,4,opt,name=IsLeader,proto3" json:"IsLeader,omitempty"`
	State     string `protobuf:"bytes,5,opt,name=State,proto3" json:"State,omitempty"`
	StateTxt  string `protobuf:"bytes,6,opt,name=StateTxt,proto3" json:"StateTxt,omitempty"`
	DBSize    uint64 `protobuf:"varint,7,opt,name=DBSize,proto3" json:"DBSize,omitempty"`
	DBSizeTxt string `protobuf:"bytes,8,opt,name=DBSizeTxt,proto3" json:"DBSizeTxt,omitempty"`
	Hash      int32  `protobuf:"varint,9,opt,name=Hash,proto3" json:"Hash,omitempty"`
}

func (m *MemberStatus) Reset()                    { *m = MemberStatus{} }
func (m *MemberStatus) String() string            { return proto.CompactTextString(m) }
func (*MemberStatus) ProtoMessage()               {}
func (*MemberStatus) Descriptor() ([]byte, []int) { return fileDescriptorClusterpb, []int{0} }

func init() {
	proto.RegisterType((*MemberStatus)(nil), "clusterpb.MemberStatus")
}
func (m *MemberStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MemberStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintClusterpb(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.ID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintClusterpb(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	if len(m.Endpoint) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintClusterpb(dAtA, i, uint64(len(m.Endpoint)))
		i += copy(dAtA[i:], m.Endpoint)
	}
	if m.IsLeader {
		dAtA[i] = 0x20
		i++
		if m.IsLeader {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.State) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintClusterpb(dAtA, i, uint64(len(m.State)))
		i += copy(dAtA[i:], m.State)
	}
	if len(m.StateTxt) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintClusterpb(dAtA, i, uint64(len(m.StateTxt)))
		i += copy(dAtA[i:], m.StateTxt)
	}
	if m.DBSize != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintClusterpb(dAtA, i, uint64(m.DBSize))
	}
	if len(m.DBSizeTxt) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintClusterpb(dAtA, i, uint64(len(m.DBSizeTxt)))
		i += copy(dAtA[i:], m.DBSizeTxt)
	}
	if m.Hash != 0 {
		dAtA[i] = 0x48
		i++
		i = encodeVarintClusterpb(dAtA, i, uint64(m.Hash))
	}
	return i, nil
}

func encodeFixed64Clusterpb(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Clusterpb(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintClusterpb(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *MemberStatus) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovClusterpb(uint64(l))
	}
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovClusterpb(uint64(l))
	}
	l = len(m.Endpoint)
	if l > 0 {
		n += 1 + l + sovClusterpb(uint64(l))
	}
	if m.IsLeader {
		n += 2
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovClusterpb(uint64(l))
	}
	l = len(m.StateTxt)
	if l > 0 {
		n += 1 + l + sovClusterpb(uint64(l))
	}
	if m.DBSize != 0 {
		n += 1 + sovClusterpb(uint64(m.DBSize))
	}
	l = len(m.DBSizeTxt)
	if l > 0 {
		n += 1 + l + sovClusterpb(uint64(l))
	}
	if m.Hash != 0 {
		n += 1 + sovClusterpb(uint64(m.Hash))
	}
	return n
}

func sovClusterpb(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozClusterpb(x uint64) (n int) {
	return sovClusterpb(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MemberStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClusterpb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MemberStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MemberStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusterpb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClusterpb
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusterpb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClusterpb
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Endpoint", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusterpb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClusterpb
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Endpoint = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsLeader", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusterpb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsLeader = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusterpb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClusterpb
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.State = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateTxt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusterpb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClusterpb
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StateTxt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DBSize", wireType)
			}
			m.DBSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusterpb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DBSize |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DBSizeTxt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusterpb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClusterpb
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DBSizeTxt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			m.Hash = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusterpb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Hash |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClusterpb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClusterpb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipClusterpb(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClusterpb
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowClusterpb
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowClusterpb
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthClusterpb
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowClusterpb
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipClusterpb(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthClusterpb = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClusterpb   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("clusterpb/clusterpb.proto", fileDescriptorClusterpb) }

var fileDescriptorClusterpb = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0x41, 0x4a, 0xc4, 0x30,
	0x14, 0x86, 0xe7, 0xd5, 0xb6, 0xb6, 0x41, 0x5c, 0x84, 0x41, 0xe2, 0x20, 0xa1, 0xb8, 0xea, 0xc6,
	0xe9, 0xc2, 0x1b, 0x0c, 0x15, 0x2c, 0xa8, 0x8b, 0x8e, 0x17, 0x68, 0x66, 0x9e, 0x9d, 0x82, 0x9d,
	0x94, 0x36, 0x05, 0xf1, 0x24, 0x1e, 0x69, 0x96, 0x1e, 0x41, 0x2b, 0x78, 0x0e, 0xc9, 0xab, 0xd6,
	0xdd, 0xf7, 0xfd, 0xc9, 0x9f, 0xf0, 0x1e, 0x3b, 0xdf, 0x3c, 0xf7, 0x9d, 0xc1, 0xb6, 0x51, 0xc9,
	0x44, 0xcb, 0xa6, 0xd5, 0x46, 0xf3, 0x70, 0x0a, 0x16, 0x57, 0x65, 0x65, 0x76, 0xbd, 0x5a, 0x6e,
	0x74, 0x9d, 0x94, 0xba, 0xd4, 0x09, 0xdd, 0x50, 0xfd, 0x13, 0x19, 0x09, 0xd1, 0xd8, 0xbc, 0xfc,
	0x06, 0x76, 0x72, 0x8f, 0xb5, 0xc2, 0x76, 0x6d, 0x0a, 0xd3, 0x77, 0x9c, 0x33, 0xf7, 0xa1, 0xa8,
	0x51, 0x40, 0x04, 0x71, 0x98, 0x13, 0xf3, 0x53, 0xe6, 0x64, 0xa9, 0x70, 0x28, 0x71, 0xb2, 0x94,
	0x2f, 0x58, 0x70, 0xb3, 0xdf, 0x36, 0xba, 0xda, 0x1b, 0x71, 0x44, 0xe9, 0xe4, 0xf6, 0x2c, 0xeb,
	0xee, 0xb0, 0xd8, 0x62, 0x2b, 0xdc, 0x08, 0xe2, 0x20, 0x9f, 0x9c, 0xcf, 0x99, 0x67, 0x7f, 0x41,
	0xe1, 0x51, 0x69, 0x14, 0xdb, 0x20, 0x78, 0x7c, 0x31, 0xc2, 0x1f, 0x5f, 0xfb, 0x73, 0x7e, 0xc6,
	0xfc, 0x74, 0xb5, 0xae, 0x5e, 0x51, 0x1c, 0x47, 0x10, 0xbb, 0xf9, 0xaf, 0xf1, 0x0b, 0x16, 0x8e,
	0x64, 0x4b, 0x01, 0x95, 0xfe, 0x03, 0x3b, 0xc3, 0x6d, 0xd1, 0xed, 0x44, 0x18, 0x41, 0xec, 0xe5,
	0xc4, 0xab, 0xf9, 0xe1, 0x53, 0xce, 0x0e, 0x83, 0x84, 0xf7, 0x41, 0xc2, 0xc7, 0x20, 0xe1, 0xed,
	0x4b, 0xce, 0x94, 0x4f, 0x5b, 0xb8, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x02, 0xf9, 0xc1, 0x85,
	0x5c, 0x01, 0x00, 0x00,
}
