// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tendermint/p2p/pex.proto

package p2p

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

type PexAddress struct {
	ID   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	IP   string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Port uint32 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
}

func (m *PexAddress) Reset()         { *m = PexAddress{} }
func (m *PexAddress) String() string { return proto.CompactTextString(m) }
func (*PexAddress) ProtoMessage()    {}
func (*PexAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_81c2f011fd13be57, []int{0}
}
func (m *PexAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PexAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PexAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PexAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PexAddress.Merge(m, src)
}
func (m *PexAddress) XXX_Size() int {
	return m.Size()
}
func (m *PexAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_PexAddress.DiscardUnknown(m)
}

var xxx_messageInfo_PexAddress proto.InternalMessageInfo

func (m *PexAddress) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *PexAddress) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *PexAddress) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type PexRequest struct {
}

func (m *PexRequest) Reset()         { *m = PexRequest{} }
func (m *PexRequest) String() string { return proto.CompactTextString(m) }
func (*PexRequest) ProtoMessage()    {}
func (*PexRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_81c2f011fd13be57, []int{1}
}
func (m *PexRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PexRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PexRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PexRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PexRequest.Merge(m, src)
}
func (m *PexRequest) XXX_Size() int {
	return m.Size()
}
func (m *PexRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PexRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PexRequest proto.InternalMessageInfo

type PexResponse struct {
	Addresses []PexAddress `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses"`
}

func (m *PexResponse) Reset()         { *m = PexResponse{} }
func (m *PexResponse) String() string { return proto.CompactTextString(m) }
func (*PexResponse) ProtoMessage()    {}
func (*PexResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_81c2f011fd13be57, []int{2}
}
func (m *PexResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PexResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PexResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PexResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PexResponse.Merge(m, src)
}
func (m *PexResponse) XXX_Size() int {
	return m.Size()
}
func (m *PexResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PexResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PexResponse proto.InternalMessageInfo

func (m *PexResponse) GetAddresses() []PexAddress {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type PexAddressV2 struct {
	URL string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (m *PexAddressV2) Reset()         { *m = PexAddressV2{} }
func (m *PexAddressV2) String() string { return proto.CompactTextString(m) }
func (*PexAddressV2) ProtoMessage()    {}
func (*PexAddressV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_81c2f011fd13be57, []int{3}
}
func (m *PexAddressV2) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PexAddressV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PexAddressV2.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PexAddressV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PexAddressV2.Merge(m, src)
}
func (m *PexAddressV2) XXX_Size() int {
	return m.Size()
}
func (m *PexAddressV2) XXX_DiscardUnknown() {
	xxx_messageInfo_PexAddressV2.DiscardUnknown(m)
}

var xxx_messageInfo_PexAddressV2 proto.InternalMessageInfo

func (m *PexAddressV2) GetURL() string {
	if m != nil {
		return m.URL
	}
	return ""
}

type PexRequestV2 struct {
}

func (m *PexRequestV2) Reset()         { *m = PexRequestV2{} }
func (m *PexRequestV2) String() string { return proto.CompactTextString(m) }
func (*PexRequestV2) ProtoMessage()    {}
func (*PexRequestV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_81c2f011fd13be57, []int{4}
}
func (m *PexRequestV2) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PexRequestV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PexRequestV2.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PexRequestV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PexRequestV2.Merge(m, src)
}
func (m *PexRequestV2) XXX_Size() int {
	return m.Size()
}
func (m *PexRequestV2) XXX_DiscardUnknown() {
	xxx_messageInfo_PexRequestV2.DiscardUnknown(m)
}

var xxx_messageInfo_PexRequestV2 proto.InternalMessageInfo

type PexResponseV2 struct {
	Addresses []PexAddressV2 `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses"`
}

func (m *PexResponseV2) Reset()         { *m = PexResponseV2{} }
func (m *PexResponseV2) String() string { return proto.CompactTextString(m) }
func (*PexResponseV2) ProtoMessage()    {}
func (*PexResponseV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_81c2f011fd13be57, []int{5}
}
func (m *PexResponseV2) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PexResponseV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PexResponseV2.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PexResponseV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PexResponseV2.Merge(m, src)
}
func (m *PexResponseV2) XXX_Size() int {
	return m.Size()
}
func (m *PexResponseV2) XXX_DiscardUnknown() {
	xxx_messageInfo_PexResponseV2.DiscardUnknown(m)
}

var xxx_messageInfo_PexResponseV2 proto.InternalMessageInfo

func (m *PexResponseV2) GetAddresses() []PexAddressV2 {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type PexMessage struct {
	// Types that are valid to be assigned to Sum:
	//	*PexMessage_PexRequest
	//	*PexMessage_PexResponse
	//	*PexMessage_PexRequestV2
	//	*PexMessage_PexResponseV2
	Sum isPexMessage_Sum `protobuf_oneof:"sum"`
}

func (m *PexMessage) Reset()         { *m = PexMessage{} }
func (m *PexMessage) String() string { return proto.CompactTextString(m) }
func (*PexMessage) ProtoMessage()    {}
func (*PexMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_81c2f011fd13be57, []int{6}
}
func (m *PexMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PexMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PexMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PexMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PexMessage.Merge(m, src)
}
func (m *PexMessage) XXX_Size() int {
	return m.Size()
}
func (m *PexMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PexMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PexMessage proto.InternalMessageInfo

type isPexMessage_Sum interface {
	isPexMessage_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type PexMessage_PexRequest struct {
	PexRequest *PexRequest `protobuf:"bytes,1,opt,name=pex_request,json=pexRequest,proto3,oneof" json:"pex_request,omitempty"`
}
type PexMessage_PexResponse struct {
	PexResponse *PexResponse `protobuf:"bytes,2,opt,name=pex_response,json=pexResponse,proto3,oneof" json:"pex_response,omitempty"`
}
type PexMessage_PexRequestV2 struct {
	PexRequestV2 *PexRequestV2 `protobuf:"bytes,3,opt,name=pex_request_v2,json=pexRequestV2,proto3,oneof" json:"pex_request_v2,omitempty"`
}
type PexMessage_PexResponseV2 struct {
	PexResponseV2 *PexResponseV2 `protobuf:"bytes,4,opt,name=pex_response_v2,json=pexResponseV2,proto3,oneof" json:"pex_response_v2,omitempty"`
}

func (*PexMessage_PexRequest) isPexMessage_Sum()    {}
func (*PexMessage_PexResponse) isPexMessage_Sum()   {}
func (*PexMessage_PexRequestV2) isPexMessage_Sum()  {}
func (*PexMessage_PexResponseV2) isPexMessage_Sum() {}

func (m *PexMessage) GetSum() isPexMessage_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *PexMessage) GetPexRequest() *PexRequest {
	if x, ok := m.GetSum().(*PexMessage_PexRequest); ok {
		return x.PexRequest
	}
	return nil
}

func (m *PexMessage) GetPexResponse() *PexResponse {
	if x, ok := m.GetSum().(*PexMessage_PexResponse); ok {
		return x.PexResponse
	}
	return nil
}

func (m *PexMessage) GetPexRequestV2() *PexRequestV2 {
	if x, ok := m.GetSum().(*PexMessage_PexRequestV2); ok {
		return x.PexRequestV2
	}
	return nil
}

func (m *PexMessage) GetPexResponseV2() *PexResponseV2 {
	if x, ok := m.GetSum().(*PexMessage_PexResponseV2); ok {
		return x.PexResponseV2
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PexMessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PexMessage_PexRequest)(nil),
		(*PexMessage_PexResponse)(nil),
		(*PexMessage_PexRequestV2)(nil),
		(*PexMessage_PexResponseV2)(nil),
	}
}

func init() {
	proto.RegisterType((*PexAddress)(nil), "tendermint.p2p.PexAddress")
	proto.RegisterType((*PexRequest)(nil), "tendermint.p2p.PexRequest")
	proto.RegisterType((*PexResponse)(nil), "tendermint.p2p.PexResponse")
	proto.RegisterType((*PexAddressV2)(nil), "tendermint.p2p.PexAddressV2")
	proto.RegisterType((*PexRequestV2)(nil), "tendermint.p2p.PexRequestV2")
	proto.RegisterType((*PexResponseV2)(nil), "tendermint.p2p.PexResponseV2")
	proto.RegisterType((*PexMessage)(nil), "tendermint.p2p.PexMessage")
}

func init() { proto.RegisterFile("tendermint/p2p/pex.proto", fileDescriptor_81c2f011fd13be57) }

var fileDescriptor_81c2f011fd13be57 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x49, 0xcd, 0x4b,
	0x49, 0x2d, 0xca, 0xcd, 0xcc, 0x2b, 0xd1, 0x2f, 0x30, 0x2a, 0xd0, 0x2f, 0x48, 0xad, 0xd0, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x43, 0xc8, 0xe8, 0x15, 0x18, 0x15, 0x48, 0x49, 0xa1, 0xa9,
	0x2c, 0xa9, 0x2c, 0x48, 0x2d, 0x86, 0xa8, 0x95, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x33, 0xf5,
	0x41, 0x2c, 0x88, 0xa8, 0x12, 0x0f, 0x17, 0x57, 0x40, 0x6a, 0x45, 0x50, 0x6a, 0x61, 0x69, 0x6a,
	0x71, 0x89, 0x92, 0x13, 0x17, 0x47, 0x40, 0x6a, 0x85, 0x63, 0x4a, 0x4a, 0x51, 0xb1, 0x90, 0x19,
	0x17, 0x6b, 0x22, 0x88, 0x21, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x6d, 0x24, 0xa5, 0x87, 0x6a, 0x97,
	0x9e, 0x5f, 0x6a, 0x09, 0x48, 0x61, 0x6a, 0x71, 0xb1, 0x13, 0xcb, 0x89, 0x7b, 0xf2, 0x0c, 0x41,
	0x10, 0xe5, 0x4a, 0x1d, 0x8c, 0x5c, 0xec, 0xbe, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x42, 0xb6,
	0x5c, 0xdc, 0x05, 0xa9, 0x15, 0xf1, 0x45, 0x10, 0xe3, 0x25, 0x18, 0x15, 0x18, 0xb1, 0x99, 0x84,
	0x70, 0x80, 0x07, 0x43, 0x10, 0x57, 0x01, 0x9c, 0x27, 0x64, 0xce, 0xc5, 0x09, 0xd2, 0x0e, 0x71,
	0x06, 0x13, 0x58, 0xb3, 0x04, 0x16, 0xcd, 0x60, 0xf7, 0x7a, 0x30, 0x04, 0x71, 0x14, 0x40, 0xd9,
	0x4e, 0xac, 0x5c, 0xcc, 0xc5, 0xa5, 0xb9, 0x4e, 0xc1, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24,
	0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78,
	0x2c, 0xc7, 0x10, 0x65, 0x99, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f,
	0x9c, 0x9a, 0x93, 0x5a, 0x5c, 0x92, 0x99, 0x98, 0x5f, 0x94, 0x0e, 0x67, 0xeb, 0x26, 0xe7, 0x17,
	0xa5, 0xea, 0x43, 0x42, 0x0b, 0x35, 0x4c, 0x93, 0xd8, 0xc0, 0xa2, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x08, 0x3a, 0x5d, 0x95, 0x96, 0x01, 0x00, 0x00,
}

func (m *PexRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PexRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PexRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *PexResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PexResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PexResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Addresses) > 0 {
		for iNdEx := len(m.Addresses) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Addresses[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPex(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *PexAddressV2) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PexAddressV2) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PexAddressV2) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.URL) > 0 {
		i -= len(m.URL)
		copy(dAtA[i:], m.URL)
		i = encodeVarintPex(dAtA, i, uint64(len(m.URL)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PexRequestV2) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PexRequestV2) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PexRequestV2) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *PexResponseV2) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PexResponseV2) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PexResponseV2) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Addresses) > 0 {
		for iNdEx := len(m.Addresses) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Addresses[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPex(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *PexMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PexMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PexMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *PexMessage_PexRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PexMessage_PexRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PexRequest != nil {
		{
			size, err := m.PexRequest.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPex(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *PexMessage_PexResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PexMessage_PexResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PexResponse != nil {
		{
			size, err := m.PexResponse.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPex(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *PexMessage_PexRequestV2) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PexMessage_PexRequestV2) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PexRequestV2 != nil {
		{
			size, err := m.PexRequestV2.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPex(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *PexMessage_PexResponseV2) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PexMessage_PexResponseV2) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PexResponseV2 != nil {
		{
			size, err := m.PexResponseV2.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPex(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func encodeVarintPex(dAtA []byte, offset int, v uint64) int {
	offset -= sovPex(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PexAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovPex(uint64(l))
	}
	l = len(m.IP)
	if l > 0 {
		n += 1 + l + sovPex(uint64(l))
	}
	if m.Port != 0 {
		n += 1 + sovPex(uint64(m.Port))
	}
	return n
}

func (m *PexRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *PexResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Addresses) > 0 {
		for _, e := range m.Addresses {
			l = e.Size()
			n += 1 + l + sovPex(uint64(l))
		}
	}
	return n
}

func (m *PexAddressV2) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.URL)
	if l > 0 {
		n += 1 + l + sovPex(uint64(l))
	}
	return n
}

func (m *PexRequestV2) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *PexResponseV2) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Addresses) > 0 {
		for _, e := range m.Addresses {
			l = e.Size()
			n += 1 + l + sovPex(uint64(l))
		}
	}
	return n
}

func (m *PexMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	return n
}

func (m *PexMessage_PexRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PexRequest != nil {
		l = m.PexRequest.Size()
		n += 1 + l + sovPex(uint64(l))
	}
	return n
}
func (m *PexMessage_PexResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PexResponse != nil {
		l = m.PexResponse.Size()
		n += 1 + l + sovPex(uint64(l))
	}
	return n
}
func (m *PexMessage_PexRequestV2) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PexRequestV2 != nil {
		l = m.PexRequestV2.Size()
		n += 1 + l + sovPex(uint64(l))
	}
	return n
}
func (m *PexMessage_PexResponseV2) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PexResponseV2 != nil {
		l = m.PexResponseV2.Size()
		n += 1 + l + sovPex(uint64(l))
	}
	return n
}

func sovPex(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPex(x uint64) (n int) {
	return sovPex(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PexAddress) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPex
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
			return fmt.Errorf("proto: PexAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PexAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPex
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
				return ErrInvalidLengthPex
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IP", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPex
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
				return ErrInvalidLengthPex
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IP = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			m.Port = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Port |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPex(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPex
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
func (m *PexRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPex
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
			return fmt.Errorf("proto: PexRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PexRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPex(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPex
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
func (m *PexResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPex
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
			return fmt.Errorf("proto: PexResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PexResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addresses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPex
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
				return ErrInvalidLengthPex
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addresses = append(m.Addresses, PexAddress{})
			if err := m.Addresses[len(m.Addresses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPex(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPex
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
func (m *PexAddressV2) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPex
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
			return fmt.Errorf("proto: PexAddressV2: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PexAddressV2: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field URL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPex
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
				return ErrInvalidLengthPex
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.URL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPex(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPex
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
func (m *PexRequestV2) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPex
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
			return fmt.Errorf("proto: PexRequestV2: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PexRequestV2: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPex(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPex
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
func (m *PexResponseV2) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPex
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
			return fmt.Errorf("proto: PexResponseV2: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PexResponseV2: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addresses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPex
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
				return ErrInvalidLengthPex
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addresses = append(m.Addresses, PexAddressV2{})
			if err := m.Addresses[len(m.Addresses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPex(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPex
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
func (m *PexMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPex
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
			return fmt.Errorf("proto: PexMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PexMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PexRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPex
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
				return ErrInvalidLengthPex
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &PexRequest{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &PexMessage_PexRequest{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PexResponse", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPex
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
				return ErrInvalidLengthPex
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &PexResponse{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &PexMessage_PexResponse{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PexRequestV2", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPex
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
				return ErrInvalidLengthPex
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &PexRequestV2{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &PexMessage_PexRequestV2{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PexResponseV2", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPex
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
				return ErrInvalidLengthPex
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &PexResponseV2{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &PexMessage_PexResponseV2{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPex(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPex
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
func skipPex(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPex
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
					return 0, ErrIntOverflowPex
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
					return 0, ErrIntOverflowPex
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
				return 0, ErrInvalidLengthPex
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPex
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPex
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPex        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPex          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPex = fmt.Errorf("proto: unexpected end of group")
)
