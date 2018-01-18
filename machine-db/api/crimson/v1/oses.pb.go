// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/machine-db/api/crimson/v1/oses.proto

package crimson

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// ListOSesRequest is a request to retrieve operating systems.
type ListOSesRequest struct {
	// The names of operating systems to retrieve.
	Names []string `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
}

func (m *ListOSesRequest) Reset()                    { *m = ListOSesRequest{} }
func (m *ListOSesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListOSesRequest) ProtoMessage()               {}
func (*ListOSesRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *ListOSesRequest) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

// OS describes an operating system.
type OS struct {
	// The name of this operating system. Uniquely identifies this operating system.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// A description of this operating system.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *OS) Reset()                    { *m = OS{} }
func (m *OS) String() string            { return proto.CompactTextString(m) }
func (*OS) ProtoMessage()               {}
func (*OS) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *OS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OS) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// ListOSesResponse is a response to a request to retrieve operating systems.
type ListOSesResponse struct {
	// The operating systems matching the request.
	Oses []*OS `protobuf:"bytes,1,rep,name=oses" json:"oses,omitempty"`
}

func (m *ListOSesResponse) Reset()                    { *m = ListOSesResponse{} }
func (m *ListOSesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListOSesResponse) ProtoMessage()               {}
func (*ListOSesResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *ListOSesResponse) GetOses() []*OS {
	if m != nil {
		return m.Oses
	}
	return nil
}

func init() {
	proto.RegisterType((*ListOSesRequest)(nil), "crimson.ListOSesRequest")
	proto.RegisterType((*OS)(nil), "crimson.OS")
	proto.RegisterType((*ListOSesResponse)(nil), "crimson.ListOSesResponse")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/machine-db/api/crimson/v1/oses.proto", fileDescriptor4)
}

var fileDescriptor4 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0x87, 0xe9, 0xba, 0x2a, 0x3b, 0x3d, 0x28, 0xc1, 0x43, 0x6f, 0x96, 0x5e, 0xdc, 0x8b, 0x09,
	0xba, 0x27, 0x7d, 0x06, 0xa1, 0x90, 0x3e, 0x41, 0x37, 0x3b, 0xb4, 0x03, 0x26, 0x13, 0x33, 0xa9,
	0xcf, 0x2f, 0x8d, 0x05, 0xbd, 0xfd, 0xfe, 0x1c, 0xbe, 0x0f, 0xde, 0x26, 0xd6, 0x6e, 0x4e, 0xec,
	0x69, 0xf1, 0x9a, 0xd3, 0x64, 0x3e, 0x17, 0x47, 0xc6, 0x8f, 0x6e, 0xa6, 0x80, 0xcf, 0x97, 0xb3,
	0x19, 0x23, 0x19, 0x97, 0xc8, 0x0b, 0x07, 0xf3, 0xfd, 0x62, 0x58, 0x50, 0x74, 0x4c, 0x9c, 0x59,
	0xdd, 0x6e, 0x73, 0xf7, 0x04, 0x77, 0x1f, 0x24, 0xb9, 0x1f, 0x50, 0x2c, 0x7e, 0x2d, 0x28, 0x59,
	0x3d, 0xc0, 0x75, 0x18, 0x3d, 0x4a, 0x53, 0xb5, 0x57, 0xc7, 0x83, 0xfd, 0x2d, 0xdd, 0x3b, 0xec,
	0xfa, 0x41, 0x29, 0xd8, 0xaf, 0xb5, 0xa9, 0xda, 0xea, 0x78, 0xb0, 0x25, 0xab, 0x16, 0xea, 0x0b,
	0x8a, 0x4b, 0x14, 0x33, 0x71, 0x68, 0x76, 0xe5, 0xfa, 0x3f, 0x75, 0x27, 0xb8, 0xff, 0x83, 0x48,
	0xe4, 0x20, 0xa8, 0x1e, 0x61, 0xbf, 0xfa, 0x14, 0x48, 0xfd, 0x5a, 0xeb, 0x4d, 0x48, 0xf7, 0x83,
	0x2d, 0xc7, 0xf9, 0xa6, 0x98, 0x9e, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x8d, 0xf8, 0xbe,
	0xe6, 0x00, 0x00, 0x00,
}