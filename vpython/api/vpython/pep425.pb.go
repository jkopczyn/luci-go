// Code generated by protoc-gen-go.
// source: github.com/luci/luci-go/vpython/api/vpython/pep425.proto
// DO NOT EDIT!

package vpython

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Represnets a Python PEP425 tag.
type Pep425Tag struct {
	// Version is the PEP425 version tag (e.g., "cp27").
	Version string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	// ABI is the PEP425 ABI tag (e.g., "cp27mu", "none").
	Abi string `protobuf:"bytes,2,opt,name=abi" json:"abi,omitempty"`
	// Arch is the PEP425 architecture tag (e.g., "linux_x86_64", "armv7l",
	// "any").
	Arch string `protobuf:"bytes,3,opt,name=arch" json:"arch,omitempty"`
}

func (m *Pep425Tag) Reset()                    { *m = Pep425Tag{} }
func (m *Pep425Tag) String() string            { return proto.CompactTextString(m) }
func (*Pep425Tag) ProtoMessage()               {}
func (*Pep425Tag) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Pep425Tag) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Pep425Tag) GetAbi() string {
	if m != nil {
		return m.Abi
	}
	return ""
}

func (m *Pep425Tag) GetArch() string {
	if m != nil {
		return m.Arch
	}
	return ""
}

func init() {
	proto.RegisterType((*Pep425Tag)(nil), "vpython.Pep425Tag")
}

func init() {
	proto.RegisterFile("github.com/luci/luci-go/vpython/api/vpython/pep425.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x48, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x29, 0x4d, 0xce, 0x04, 0x13, 0xba, 0xe9, 0xf9,
	0xfa, 0x65, 0x05, 0x95, 0x25, 0x19, 0xf9, 0x79, 0xfa, 0x89, 0x05, 0x99, 0x70, 0x76, 0x41, 0x6a,
	0x81, 0x89, 0x91, 0xa9, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x10, 0x3b, 0x54, 0x54, 0xc9, 0x9b,
	0x8b, 0x33, 0x00, 0x2c, 0x11, 0x92, 0x98, 0x2e, 0x24, 0xc1, 0xc5, 0x5e, 0x96, 0x5a, 0x54, 0x9c,
	0x99, 0x9f, 0x27, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe3, 0x0a, 0x09, 0x70, 0x31, 0x27,
	0x26, 0x65, 0x4a, 0x30, 0x81, 0x45, 0x41, 0x4c, 0x21, 0x21, 0x2e, 0x96, 0xc4, 0xa2, 0xe4, 0x0c,
	0x09, 0x66, 0xb0, 0x10, 0x98, 0x9d, 0xc4, 0x06, 0x36, 0xdc, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0x17, 0x39, 0x2f, 0x5f, 0x98, 0x00, 0x00, 0x00,
}
