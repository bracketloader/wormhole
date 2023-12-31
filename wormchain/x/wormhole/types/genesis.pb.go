// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: wormhole/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// GenesisState defines the wormhole module's genesis state.
type GenesisState struct {
	GuardianSetList           []GuardianSet              `protobuf:"bytes,1,rep,name=guardianSetList,proto3" json:"guardianSetList"`
	Config                    *Config                    `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	ReplayProtectionList      []ReplayProtection         `protobuf:"bytes,3,rep,name=replayProtectionList,proto3" json:"replayProtectionList"`
	SequenceCounterList       []SequenceCounter          `protobuf:"bytes,4,rep,name=sequenceCounterList,proto3" json:"sequenceCounterList"`
	ConsensusGuardianSetIndex *ConsensusGuardianSetIndex `protobuf:"bytes,5,opt,name=consensusGuardianSetIndex,proto3" json:"consensusGuardianSetIndex,omitempty"`
	GuardianValidatorList     []GuardianValidator        `protobuf:"bytes,6,rep,name=guardianValidatorList,proto3" json:"guardianValidatorList"`
	AllowedAddresses          []ValidatorAllowedAddress  `protobuf:"bytes,7,rep,name=allowedAddresses,proto3" json:"allowedAddresses"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a7ced3fe0304831, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetGuardianSetList() []GuardianSet {
	if m != nil {
		return m.GuardianSetList
	}
	return nil
}

func (m *GenesisState) GetConfig() *Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *GenesisState) GetReplayProtectionList() []ReplayProtection {
	if m != nil {
		return m.ReplayProtectionList
	}
	return nil
}

func (m *GenesisState) GetSequenceCounterList() []SequenceCounter {
	if m != nil {
		return m.SequenceCounterList
	}
	return nil
}

func (m *GenesisState) GetConsensusGuardianSetIndex() *ConsensusGuardianSetIndex {
	if m != nil {
		return m.ConsensusGuardianSetIndex
	}
	return nil
}

func (m *GenesisState) GetGuardianValidatorList() []GuardianValidator {
	if m != nil {
		return m.GuardianValidatorList
	}
	return nil
}

func (m *GenesisState) GetAllowedAddresses() []ValidatorAllowedAddress {
	if m != nil {
		return m.AllowedAddresses
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "wormhole_foundation.wormchain.wormhole.GenesisState")
}

func init() { proto.RegisterFile("wormhole/genesis.proto", fileDescriptor_9a7ced3fe0304831) }

var fileDescriptor_9a7ced3fe0304831 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4f, 0x6b, 0xd4, 0x40,
	0x18, 0xc6, 0x13, 0x5b, 0x23, 0x4c, 0x05, 0x65, 0x6c, 0x35, 0xf6, 0x90, 0x2e, 0x1e, 0xa4, 0x20,
	0x26, 0xd0, 0x1e, 0xb4, 0x27, 0xd9, 0x16, 0x2c, 0x82, 0x07, 0x49, 0xc0, 0x83, 0x97, 0x30, 0x4d,
	0xde, 0x66, 0x07, 0xd2, 0x99, 0x6d, 0x66, 0x42, 0xbb, 0x27, 0x3f, 0x80, 0x17, 0x3f, 0xd6, 0x1e,
	0xf7, 0xe8, 0x49, 0x64, 0xf7, 0x8b, 0x48, 0x66, 0x26, 0xb3, 0x7f, 0x8c, 0x90, 0xbd, 0x0d, 0xcf,
	0xbc, 0xef, 0xef, 0x99, 0xf7, 0x79, 0x13, 0xf4, 0xfc, 0x8e, 0x57, 0x37, 0x23, 0x5e, 0x42, 0x54,
	0x00, 0x03, 0x41, 0x45, 0x38, 0xae, 0xb8, 0xe4, 0xf8, 0x75, 0xab, 0xa7, 0xd7, 0xbc, 0x66, 0x39,
	0x91, 0x94, 0xb3, 0xb0, 0xd1, 0xb2, 0x11, 0xa1, 0xfa, 0xd4, 0xdc, 0x1e, 0xbe, 0x58, 0xf6, 0xd7,
	0xa4, 0xca, 0x29, 0x61, 0x1a, 0x70, 0x78, 0x60, 0x2f, 0x32, 0xce, 0xae, 0x69, 0x61, 0xe4, 0x81,
	0x95, 0x2b, 0x18, 0x97, 0x64, 0x92, 0x36, 0x32, 0x64, 0x0a, 0xaf, 0x2b, 0x8e, 0x6c, 0x85, 0x80,
	0xdb, 0x1a, 0x58, 0x06, 0x69, 0xc6, 0x6b, 0x26, 0xa1, 0x32, 0x05, 0x6f, 0x56, 0xc9, 0x02, 0x98,
	0xa8, 0x45, 0xda, 0x9a, 0xa7, 0x02, 0x64, 0x4a, 0x59, 0x0e, 0xf7, 0xa6, 0x78, 0xbf, 0xe0, 0x05,
	0x57, 0xc7, 0xa8, 0x39, 0x69, 0xf5, 0xd5, 0x0f, 0x0f, 0x3d, 0xbe, 0xd4, 0xf3, 0x26, 0x92, 0x48,
	0xc0, 0x19, 0x7a, 0xd2, 0x22, 0x12, 0x90, 0x9f, 0xa9, 0x90, 0xbe, 0x3b, 0xd8, 0x39, 0xde, 0x3b,
	0x39, 0x0d, 0xfb, 0x05, 0x11, 0x5e, 0x2e, 0xdb, 0xcf, 0x77, 0xa7, 0xbf, 0x8f, 0x9c, 0x78, 0x93,
	0x88, 0x3f, 0x22, 0x4f, 0x67, 0xe1, 0x3f, 0x18, 0xb8, 0xc7, 0x7b, 0x27, 0x61, 0x5f, 0xf6, 0x85,
	0xea, 0x8a, 0x4d, 0x37, 0xae, 0xd0, 0xbe, 0x0e, 0xef, 0x8b, 0xcd, 0x4e, 0xbd, 0x78, 0x47, 0xbd,
	0xf8, 0x7d, 0x5f, 0x6a, 0xbc, 0xc1, 0x30, 0xcf, 0xee, 0x64, 0x63, 0x8e, 0x9e, 0xb5, 0xeb, 0xb8,
	0xd0, 0xdb, 0x50, 0x96, 0xbb, 0xca, 0xf2, 0x5d, 0x5f, 0xcb, 0x64, 0x1d, 0x61, 0x1c, 0xbb, 0xc8,
	0xf8, 0x3b, 0x7a, 0x69, 0xd7, 0xbb, 0x92, 0xed, 0xa7, 0x66, 0xb7, 0xfe, 0x43, 0x95, 0xdf, 0x70,
	0x8b, 0xfc, 0xba, 0x41, 0xf1, 0xff, 0x3d, 0x70, 0x8d, 0x0e, 0xda, 0x05, 0x7e, 0x25, 0x25, 0xcd,
	0x89, 0xe4, 0x7a, 0x66, 0x4f, 0xcd, 0x7c, 0xb6, 0xed, 0x87, 0x61, 0x21, 0x66, 0xea, 0x6e, 0x3a,
	0xbe, 0x45, 0x4f, 0x49, 0x59, 0xf2, 0x3b, 0xc8, 0x87, 0x79, 0x5e, 0x81, 0x10, 0x20, 0xfc, 0x47,
	0xca, 0xf1, 0x43, 0x5f, 0x47, 0x0b, 0x1c, 0xae, 0x81, 0x8c, 0xef, 0x3f, 0xf8, 0xf3, 0x64, 0x3a,
	0x0f, 0xdc, 0xd9, 0x3c, 0x70, 0xff, 0xcc, 0x03, 0xf7, 0xe7, 0x22, 0x70, 0x66, 0x8b, 0xc0, 0xf9,
	0xb5, 0x08, 0x9c, 0x6f, 0x67, 0x05, 0x95, 0xa3, 0xfa, 0x2a, 0xcc, 0xf8, 0x4d, 0xd4, 0xe2, 0xdf,
	0x2e, 0xcd, 0x23, 0x6b, 0x1e, 0xdd, 0xdb, 0xfb, 0x48, 0x4e, 0xc6, 0x20, 0xae, 0x3c, 0xf5, 0xa7,
	0x9d, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x77, 0xb5, 0xdb, 0x2c, 0x61, 0x04, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AllowedAddresses) > 0 {
		for iNdEx := len(m.AllowedAddresses) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AllowedAddresses[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.GuardianValidatorList) > 0 {
		for iNdEx := len(m.GuardianValidatorList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GuardianValidatorList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.ConsensusGuardianSetIndex != nil {
		{
			size, err := m.ConsensusGuardianSetIndex.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SequenceCounterList) > 0 {
		for iNdEx := len(m.SequenceCounterList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SequenceCounterList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ReplayProtectionList) > 0 {
		for iNdEx := len(m.ReplayProtectionList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ReplayProtectionList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Config != nil {
		{
			size, err := m.Config.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.GuardianSetList) > 0 {
		for iNdEx := len(m.GuardianSetList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GuardianSetList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.GuardianSetList) > 0 {
		for _, e := range m.GuardianSetList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.Config != nil {
		l = m.Config.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.ReplayProtectionList) > 0 {
		for _, e := range m.ReplayProtectionList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.SequenceCounterList) > 0 {
		for _, e := range m.SequenceCounterList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.ConsensusGuardianSetIndex != nil {
		l = m.ConsensusGuardianSetIndex.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.GuardianValidatorList) > 0 {
		for _, e := range m.GuardianValidatorList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.AllowedAddresses) > 0 {
		for _, e := range m.AllowedAddresses {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GuardianSetList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GuardianSetList = append(m.GuardianSetList, GuardianSet{})
			if err := m.GuardianSetList[len(m.GuardianSetList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Config == nil {
				m.Config = &Config{}
			}
			if err := m.Config.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplayProtectionList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReplayProtectionList = append(m.ReplayProtectionList, ReplayProtection{})
			if err := m.ReplayProtectionList[len(m.ReplayProtectionList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SequenceCounterList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SequenceCounterList = append(m.SequenceCounterList, SequenceCounter{})
			if err := m.SequenceCounterList[len(m.SequenceCounterList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusGuardianSetIndex", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ConsensusGuardianSetIndex == nil {
				m.ConsensusGuardianSetIndex = &ConsensusGuardianSetIndex{}
			}
			if err := m.ConsensusGuardianSetIndex.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GuardianValidatorList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GuardianValidatorList = append(m.GuardianValidatorList, GuardianValidator{})
			if err := m.GuardianValidatorList[len(m.GuardianValidatorList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedAddresses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllowedAddresses = append(m.AllowedAddresses, ValidatorAllowedAddress{})
			if err := m.AllowedAddresses[len(m.AllowedAddresses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
