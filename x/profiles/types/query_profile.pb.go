// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mage/profiles/v3/query_profile.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryProfileRequest is the request type for the Query/Profile RPC method.
type QueryProfileRequest struct {
	// Address or DTag of the user to query the profile for
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (m *QueryProfileRequest) Reset()         { *m = QueryProfileRequest{} }
func (m *QueryProfileRequest) String() string { return proto.CompactTextString(m) }
func (*QueryProfileRequest) ProtoMessage()    {}
func (*QueryProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c0e029b5401ec84, []int{0}
}
func (m *QueryProfileRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryProfileRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryProfileRequest.Merge(m, src)
}
func (m *QueryProfileRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryProfileRequest proto.InternalMessageInfo

// QueryProfileResponse is the response type for the Query/Profile RPC method.
type QueryProfileResponse struct {
	Profile *types.Any `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (m *QueryProfileResponse) Reset()         { *m = QueryProfileResponse{} }
func (m *QueryProfileResponse) String() string { return proto.CompactTextString(m) }
func (*QueryProfileResponse) ProtoMessage()    {}
func (*QueryProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c0e029b5401ec84, []int{1}
}
func (m *QueryProfileResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryProfileResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryProfileResponse.Merge(m, src)
}
func (m *QueryProfileResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryProfileResponse proto.InternalMessageInfo

func (m *QueryProfileResponse) GetProfile() *types.Any {
	if m != nil {
		return m.Profile
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryProfileRequest)(nil), "mage.profiles.v3.QueryProfileRequest")
	proto.RegisterType((*QueryProfileResponse)(nil), "mage.profiles.v3.QueryProfileResponse")
}

func init() {
	proto.RegisterFile("mage/profiles/v3/query_profile.proto", fileDescriptor_7c0e029b5401ec84)
}

var fileDescriptor_7c0e029b5401ec84 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x51, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0x4c, 0x24, 0x04, 0x25, 0x30, 0x85, 0x0e, 0x50, 0xa1, 0x14, 0x75, 0x40, 0x08, 0xa9, 0x7e,
	0x2a, 0x61, 0x42, 0x2c, 0xed, 0x86, 0x58, 0xa0, 0x62, 0x62, 0xa9, 0x9c, 0xe0, 0x86, 0x48, 0xad,
	0x9f, 0x1b, 0xdb, 0x11, 0xf9, 0x03, 0x46, 0x3e, 0xa1, 0x1f, 0xc1, 0x47, 0x20, 0xa6, 0x8e, 0x8c,
	0xa8, 0x5d, 0xf8, 0x0c, 0x14, 0xdb, 0x11, 0x62, 0x7b, 0xe7, 0xbb, 0xf3, 0xbd, 0xb3, 0x83, 0xd3,
	0x27, 0x26, 0xe7, 0x28, 0x41, 0x14, 0x38, 0xcd, 0x67, 0x4c, 0x42, 0x19, 0xc3, 0x42, 0xb3, 0xa2,
	0x9a, 0xb8, 0x13, 0x22, 0x0a, 0x54, 0x18, 0x86, 0x56, 0x47, 0x1a, 0x1d, 0x29, 0xe3, 0x4e, 0x3b,
	0xc3, 0x0c, 0x0d, 0x0d, 0xf5, 0x64, 0x95, 0x9d, 0xe3, 0x0c, 0x31, 0x9b, 0x31, 0xa0, 0x22, 0x07,
	0xca, 0x39, 0x2a, 0xaa, 0x72, 0xe4, 0xd2, 0xb1, 0x47, 0x8e, 0x35, 0x28, 0xd1, 0x53, 0xa0, 0xbc,
	0x6a, 0xa8, 0x14, 0xeb, 0x88, 0x89, 0xbd, 0xd1, 0x02, 0x47, 0x9d, 0x5b, 0x04, 0x09, 0x95, 0xcc,
	0xae, 0x07, 0xe5, 0x20, 0x61, 0x8a, 0x0e, 0x40, 0xd0, 0x2c, 0xe7, 0x26, 0xc2, 0x6a, 0x7b, 0x71,
	0x70, 0x70, 0x5f, 0x2b, 0xee, 0xec, 0xa6, 0x63, 0xb6, 0xd0, 0x4c, 0xaa, 0x30, 0x0c, 0xb6, 0xb4,
	0x64, 0xc5, 0xa1, 0x7f, 0xe2, 0x9f, 0xed, 0x8e, 0xcd, 0x7c, 0xd5, 0x7a, 0x5d, 0x76, 0xbd, 0x9f,
	0x65, 0xd7, 0xeb, 0x3d, 0x04, 0xed, 0xff, 0x26, 0x29, 0x90, 0x4b, 0x16, 0x5e, 0x07, 0x3b, 0xae,
	0xb1, 0x31, 0xee, 0x5d, 0xb4, 0x89, 0x2d, 0x40, 0x9a, 0x02, 0x64, 0xc8, 0xab, 0xd1, 0xfe, 0xe7,
	0x7b, 0xbf, 0x35, 0x4c, 0x53, 0xd4, 0x5c, 0xdd, 0x8c, 0x1b, 0xcb, 0xe8, 0xf6, 0x63, 0x1d, 0xf9,
	0xab, 0x75, 0xe4, 0x7f, 0xaf, 0x23, 0xff, 0x6d, 0x13, 0x79, 0xab, 0x4d, 0xe4, 0x7d, 0x6d, 0x22,
	0xef, 0x71, 0x90, 0xe5, 0xea, 0x59, 0x27, 0x24, 0xc5, 0x39, 0xd8, 0x97, 0xed, 0xcf, 0x68, 0x22,
	0xdd, 0x0c, 0xe5, 0x25, 0xbc, 0xfc, 0x7d, 0x89, 0xaa, 0x04, 0x93, 0xc9, 0xb6, 0x49, 0x8c, 0x7f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x01, 0x0a, 0x39, 0x3f, 0xb2, 0x01, 0x00, 0x00,
}

func (m *QueryProfileRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryProfileRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryProfileRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintQueryProfile(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryProfileResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryProfileResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryProfileResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Profile != nil {
		{
			size, err := m.Profile.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQueryProfile(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQueryProfile(dAtA []byte, offset int, v uint64) int {
	offset -= sovQueryProfile(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryProfileRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovQueryProfile(uint64(l))
	}
	return n
}

func (m *QueryProfileResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Profile != nil {
		l = m.Profile.Size()
		n += 1 + l + sovQueryProfile(uint64(l))
	}
	return n
}

func sovQueryProfile(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQueryProfile(x uint64) (n int) {
	return sovQueryProfile(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryProfileRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryProfile
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryProfileRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryProfileRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryProfile
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQueryProfile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryProfile(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryProfile
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
func (m *QueryProfileResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryProfile
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryProfileResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryProfileResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Profile", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryProfile
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQueryProfile
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Profile == nil {
				m.Profile = &types.Any{}
			}
			if err := m.Profile.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryProfile(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryProfile
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
func skipQueryProfile(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueryProfile
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
					return 0, ErrIntOverflowQueryProfile
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQueryProfile
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
			if length < 0 {
				return 0, ErrInvalidLengthQueryProfile
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQueryProfile
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQueryProfile
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQueryProfile        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueryProfile          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQueryProfile = fmt.Errorf("proto: unexpected end of group")
)
