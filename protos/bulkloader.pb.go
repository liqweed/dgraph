// Code generated by protoc-gen-gogo.
// source: bulkloader.proto
// DO NOT EDIT!

/*
	Package protos is a generated protocol buffer package.

	It is generated from these files:
		bulkloader.proto
		facets.proto
		graphresponse.proto
		payload.proto
		schema.proto
		task.proto
		types.proto

	It has these top-level messages:
		MapEntry
		Facet
		Param
		Facets
		FacetsList
		Function
		FilterTree
		Num
		AssignedIds
		NQuad
		Value
		Mutation
		Request
		Latency
		Property
		Node
		Response
		Check
		Version
		Payload
		MovePredicatePayload
		ExportPayload
		SchemaRequest
		SchemaResult
		SchemaNode
		SchemaUpdate
		List
		TaskValue
		SrcFunction
		Query
		ValueList
		Result
		Order
		SortMessage
		SortResult
		RaftContext
		Member
		Group
		ZeroProposal
		MembershipState
		Tablet
		DirectedEdge
		Mutations
		KeyValues
		Proposal
		KV
		KC
		GroupKeys
		Posting
		PostingList
*/
package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type MapEntry struct {
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Only one should be set.
	Uid     uint64   `protobuf:"fixed64,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Posting *Posting `protobuf:"bytes,3,opt,name=posting" json:"posting,omitempty"`
}

func (m *MapEntry) Reset()                    { *m = MapEntry{} }
func (m *MapEntry) String() string            { return proto.CompactTextString(m) }
func (*MapEntry) ProtoMessage()               {}
func (*MapEntry) Descriptor() ([]byte, []int) { return fileDescriptorBulkloader, []int{0} }

func (m *MapEntry) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *MapEntry) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *MapEntry) GetPosting() *Posting {
	if m != nil {
		return m.Posting
	}
	return nil
}

func init() {
	proto.RegisterType((*MapEntry)(nil), "protos.MapEntry")
}
func (m *MapEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MapEntry) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintBulkloader(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if m.Uid != 0 {
		dAtA[i] = 0x11
		i++
		i = encodeFixed64Bulkloader(dAtA, i, uint64(m.Uid))
	}
	if m.Posting != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintBulkloader(dAtA, i, uint64(m.Posting.Size()))
		n1, err := m.Posting.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func encodeFixed64Bulkloader(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Bulkloader(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintBulkloader(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *MapEntry) Size() (n int) {
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovBulkloader(uint64(l))
	}
	if m.Uid != 0 {
		n += 9
	}
	if m.Posting != nil {
		l = m.Posting.Size()
		n += 1 + l + sovBulkloader(uint64(l))
	}
	return n
}

func sovBulkloader(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozBulkloader(x uint64) (n int) {
	return sovBulkloader(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MapEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBulkloader
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
			return fmt.Errorf("proto: MapEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MapEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBulkloader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBulkloader
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			m.Uid = uint64(dAtA[iNdEx-8])
			m.Uid |= uint64(dAtA[iNdEx-7]) << 8
			m.Uid |= uint64(dAtA[iNdEx-6]) << 16
			m.Uid |= uint64(dAtA[iNdEx-5]) << 24
			m.Uid |= uint64(dAtA[iNdEx-4]) << 32
			m.Uid |= uint64(dAtA[iNdEx-3]) << 40
			m.Uid |= uint64(dAtA[iNdEx-2]) << 48
			m.Uid |= uint64(dAtA[iNdEx-1]) << 56
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Posting", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBulkloader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBulkloader
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Posting == nil {
				m.Posting = &Posting{}
			}
			if err := m.Posting.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBulkloader(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBulkloader
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
func skipBulkloader(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBulkloader
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
					return 0, ErrIntOverflowBulkloader
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
					return 0, ErrIntOverflowBulkloader
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
				return 0, ErrInvalidLengthBulkloader
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBulkloader
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
				next, err := skipBulkloader(dAtA[start:])
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
	ErrInvalidLengthBulkloader = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBulkloader   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("bulkloader.proto", fileDescriptorBulkloader) }

var fileDescriptorBulkloader = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x2a, 0xcd, 0xc9,
	0xce, 0xc9, 0x4f, 0x4c, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53,
	0xc5, 0x52, 0xdc, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x10, 0x41, 0xa5, 0x48, 0x2e, 0x0e, 0xdf, 0xc4,
	0x02, 0xd7, 0xbc, 0x92, 0xa2, 0x4a, 0x21, 0x01, 0x2e, 0xe6, 0xec, 0xd4, 0x4a, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0x9e, 0x20, 0x10, 0x13, 0x24, 0x52, 0x9a, 0x99, 0x22, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1,
	0x16, 0x04, 0x62, 0x0a, 0x69, 0x72, 0xb1, 0x17, 0xe4, 0x17, 0x97, 0x64, 0xe6, 0xa5, 0x4b, 0x30,
	0x2b, 0x30, 0x6a, 0x70, 0x1b, 0xf1, 0x43, 0x0c, 0x2a, 0xd6, 0x0b, 0x80, 0x08, 0x07, 0xc1, 0xe4,
	0x9d, 0x04, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x19,
	0x8f, 0xe5, 0x18, 0x92, 0x20, 0x2e, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb5, 0xda, 0xed,
	0xae, 0x9c, 0x00, 0x00, 0x00,
}
