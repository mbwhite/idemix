// Code generated by protoc-gen-go. DO NOT EDIT.
// source: idemixmsp/msp_config.proto

package idemixmsp

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type IdemixMSPConfig struct {
	// Name holds the identifier of the MSP
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// ipk represents the (serialized) issuer public key
	Ipk []byte `protobuf:"bytes,2,opt,name=ipk,proto3" json:"ipk,omitempty"`
	// signer may contain crypto material to configure a default signer
	Signer *IdemixMSPSignerConfig `protobuf:"bytes,3,opt,name=signer,proto3" json:"signer,omitempty"`
	// revocation_pk is the public key used for revocation of credentials
	RevocationPk []byte `protobuf:"bytes,4,opt,name=revocation_pk,json=revocationPk,proto3" json:"revocation_pk,omitempty"`
	// epoch represents the current epoch (time interval) used for revocation
	Epoch int64 `protobuf:"varint,5,opt,name=epoch,proto3" json:"epoch,omitempty"`
	// curve_id indicates which Elliptic Curve should be used
	CurveId              string   `protobuf:"bytes,6,opt,name=curve_id,json=curveId,proto3" json:"curve_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdemixMSPConfig) Reset()         { *m = IdemixMSPConfig{} }
func (m *IdemixMSPConfig) String() string { return proto.CompactTextString(m) }
func (*IdemixMSPConfig) ProtoMessage()    {}
func (*IdemixMSPConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_783ba658a7862c73, []int{0}
}

func (m *IdemixMSPConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdemixMSPConfig.Unmarshal(m, b)
}
func (m *IdemixMSPConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdemixMSPConfig.Marshal(b, m, deterministic)
}
func (m *IdemixMSPConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdemixMSPConfig.Merge(m, src)
}
func (m *IdemixMSPConfig) XXX_Size() int {
	return xxx_messageInfo_IdemixMSPConfig.Size(m)
}
func (m *IdemixMSPConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_IdemixMSPConfig.DiscardUnknown(m)
}

var xxx_messageInfo_IdemixMSPConfig proto.InternalMessageInfo

func (m *IdemixMSPConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *IdemixMSPConfig) GetIpk() []byte {
	if m != nil {
		return m.Ipk
	}
	return nil
}

func (m *IdemixMSPConfig) GetSigner() *IdemixMSPSignerConfig {
	if m != nil {
		return m.Signer
	}
	return nil
}

func (m *IdemixMSPConfig) GetRevocationPk() []byte {
	if m != nil {
		return m.RevocationPk
	}
	return nil
}

func (m *IdemixMSPConfig) GetEpoch() int64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *IdemixMSPConfig) GetCurveId() string {
	if m != nil {
		return m.CurveId
	}
	return ""
}

// IdemixMSPSIgnerConfig contains the crypto material to set up an idemix signing identity
type IdemixMSPSignerConfig struct {
	// cred represents the serialized idemix credential of the default signer
	Cred []byte `protobuf:"bytes,1,opt,name=cred,proto3" json:"cred,omitempty"`
	// sk is the secret key of the default signer, corresponding to credential Cred
	Sk []byte `protobuf:"bytes,2,opt,name=sk,proto3" json:"sk,omitempty"`
	// organizational_unit_identifier defines the organizational unit the default signer is in
	OrganizationalUnitIdentifier string `protobuf:"bytes,3,opt,name=organizational_unit_identifier,json=organizationalUnitIdentifier,proto3" json:"organizational_unit_identifier,omitempty"`
	// role defines whether the default signer is admin, peer, member or client
	Role int32 `protobuf:"varint,4,opt,name=role,proto3" json:"role,omitempty"`
	// enrollment_id contains the enrollment id of this signer
	EnrollmentId string `protobuf:"bytes,5,opt,name=enrollment_id,json=enrollmentId,proto3" json:"enrollment_id,omitempty"`
	// credential_revocation_information contains a serialized CredentialRevocationInformation
	CredentialRevocationInformation []byte   `protobuf:"bytes,6,opt,name=credential_revocation_information,json=credentialRevocationInformation,proto3" json:"credential_revocation_information,omitempty"`
	XXX_NoUnkeyedLiteral            struct{} `json:"-"`
	XXX_unrecognized                []byte   `json:"-"`
	XXX_sizecache                   int32    `json:"-"`
}

func (m *IdemixMSPSignerConfig) Reset()         { *m = IdemixMSPSignerConfig{} }
func (m *IdemixMSPSignerConfig) String() string { return proto.CompactTextString(m) }
func (*IdemixMSPSignerConfig) ProtoMessage()    {}
func (*IdemixMSPSignerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_783ba658a7862c73, []int{1}
}

func (m *IdemixMSPSignerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdemixMSPSignerConfig.Unmarshal(m, b)
}
func (m *IdemixMSPSignerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdemixMSPSignerConfig.Marshal(b, m, deterministic)
}
func (m *IdemixMSPSignerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdemixMSPSignerConfig.Merge(m, src)
}
func (m *IdemixMSPSignerConfig) XXX_Size() int {
	return xxx_messageInfo_IdemixMSPSignerConfig.Size(m)
}
func (m *IdemixMSPSignerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_IdemixMSPSignerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_IdemixMSPSignerConfig proto.InternalMessageInfo

func (m *IdemixMSPSignerConfig) GetCred() []byte {
	if m != nil {
		return m.Cred
	}
	return nil
}

func (m *IdemixMSPSignerConfig) GetSk() []byte {
	if m != nil {
		return m.Sk
	}
	return nil
}

func (m *IdemixMSPSignerConfig) GetOrganizationalUnitIdentifier() string {
	if m != nil {
		return m.OrganizationalUnitIdentifier
	}
	return ""
}

func (m *IdemixMSPSignerConfig) GetRole() int32 {
	if m != nil {
		return m.Role
	}
	return 0
}

func (m *IdemixMSPSignerConfig) GetEnrollmentId() string {
	if m != nil {
		return m.EnrollmentId
	}
	return ""
}

func (m *IdemixMSPSignerConfig) GetCredentialRevocationInformation() []byte {
	if m != nil {
		return m.CredentialRevocationInformation
	}
	return nil
}

func init() {
	proto.RegisterType((*IdemixMSPConfig)(nil), "idemixmsp.IdemixMSPConfig")
	proto.RegisterType((*IdemixMSPSignerConfig)(nil), "idemixmsp.IdemixMSPSignerConfig")
}

func init() { proto.RegisterFile("idemixmsp/msp_config.proto", fileDescriptor_783ba658a7862c73) }

var fileDescriptor_783ba658a7862c73 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x41, 0x8b, 0xd4, 0x30,
	0x14, 0xc7, 0x49, 0x67, 0x67, 0xb4, 0xb1, 0xbb, 0x4a, 0x70, 0xa1, 0x8a, 0xb8, 0x75, 0xbd, 0xcc,
	0x69, 0x0a, 0x7a, 0xf1, 0x5c, 0x05, 0x89, 0x50, 0x28, 0x59, 0x84, 0x45, 0x84, 0x92, 0x6d, 0x33,
	0xdd, 0xd0, 0x26, 0x29, 0x69, 0x66, 0x11, 0xcf, 0xe2, 0x07, 0xf1, 0xe8, 0x17, 0x11, 0xfc, 0x18,
	0x1e, 0xfd, 0x14, 0x92, 0xd7, 0x9d, 0x76, 0x84, 0xbd, 0xfd, 0xdf, 0xeb, 0xaf, 0xff, 0xbc, 0x7f,
	0x5e, 0xf0, 0x53, 0x59, 0x0b, 0x25, 0xbf, 0xa8, 0xa1, 0x4f, 0xd5, 0xd0, 0x97, 0x95, 0xd1, 0x5b,
	0xd9, 0x6c, 0x7a, 0x6b, 0x9c, 0x21, 0xe1, 0xf4, 0xed, 0xfc, 0x17, 0xc2, 0x0f, 0x29, 0x54, 0xf9,
	0x45, 0xf1, 0x16, 0x20, 0x42, 0xf0, 0x91, 0xe6, 0x4a, 0xc4, 0x28, 0x41, 0xeb, 0x90, 0x81, 0x26,
	0x8f, 0xf0, 0x42, 0xf6, 0x6d, 0x1c, 0x24, 0x68, 0x1d, 0x31, 0x2f, 0xc9, 0x1b, 0xbc, 0x1a, 0x64,
	0xa3, 0x85, 0x8d, 0x17, 0x09, 0x5a, 0x3f, 0x78, 0x95, 0x6c, 0x26, 0xd7, 0xcd, 0xe4, 0x78, 0x01,
	0xc4, 0xe8, 0xcb, 0x6e, 0x79, 0xf2, 0x12, 0x1f, 0x5b, 0x71, 0x63, 0x2a, 0xee, 0xa4, 0xd1, 0x65,
	0xdf, 0xc6, 0x47, 0xe0, 0x1a, 0xcd, 0xcd, 0xa2, 0x25, 0x8f, 0xf1, 0x52, 0xf4, 0xa6, 0xba, 0x8e,
	0x97, 0x09, 0x5a, 0x2f, 0xd8, 0x58, 0x90, 0x27, 0xf8, 0x7e, 0xb5, 0xb3, 0x37, 0xa2, 0x94, 0x75,
	0xbc, 0x82, 0xf1, 0xee, 0x41, 0x4d, 0xeb, 0xf3, 0xef, 0x01, 0x3e, 0xbd, 0xf3, 0x5c, 0x9f, 0xa7,
	0xb2, 0xa2, 0x86, 0x3c, 0x11, 0x03, 0x4d, 0x4e, 0x70, 0x30, 0xec, 0xe3, 0x04, 0x43, 0x4b, 0xde,
	0xe1, 0xe7, 0xc6, 0x36, 0x5c, 0xcb, 0xaf, 0x30, 0x00, 0xef, 0xca, 0x9d, 0x96, 0xae, 0x94, 0xb5,
	0xd0, 0x4e, 0x6e, 0xe5, 0x6d, 0xca, 0x90, 0x3d, 0xfb, 0x9f, 0xfa, 0xa8, 0xa5, 0xa3, 0x13, 0xe3,
	0x4f, 0xb2, 0xa6, 0x13, 0x10, 0x68, 0xc9, 0x40, 0xfb, 0xb4, 0x42, 0x5b, 0xd3, 0x75, 0x4a, 0x68,
	0x6f, 0x08, 0x81, 0x42, 0x16, 0xcd, 0x4d, 0x5a, 0x93, 0x0f, 0xf8, 0x85, 0x1f, 0xcb, 0x1b, 0xf1,
	0xae, 0x3c, 0xb8, 0x1d, 0xa9, 0xb7, 0xc6, 0x2a, 0xd0, 0x10, 0x38, 0x62, 0x67, 0x33, 0xc8, 0x26,
	0x8e, 0xce, 0x58, 0xf6, 0x0d, 0xe1, 0xe3, 0xca, 0xa8, 0x79, 0x1d, 0xd9, 0x49, 0x3e, 0xf4, 0xe3,
	0x5d, 0x14, 0x7e, 0xff, 0x05, 0xfa, 0x74, 0xd6, 0x48, 0x77, 0xbd, 0xbb, 0xda, 0x54, 0x46, 0xa5,
	0x34, 0xcb, 0xd3, 0x91, 0x4d, 0xa7, 0x5f, 0x7e, 0x04, 0x0b, 0x7a, 0x79, 0xf9, 0x33, 0x08, 0xe9,
	0xbe, 0xf3, 0xfb, 0x40, 0xff, 0x09, 0x4e, 0x27, 0xfd, 0xf9, 0x7d, 0x91, 0xe5, 0xc2, 0xf1, 0x9a,
	0x3b, 0xfe, 0xf7, 0x80, 0xb9, 0x5a, 0xc1, 0x5b, 0x7b, 0xfd, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x5d,
	0xd6, 0x35, 0x81, 0x89, 0x02, 0x00, 0x00,
}