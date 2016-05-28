// Code generated by protoc-gen-go.
// source: identity.proto
// DO NOT EDIT!

/*
Package identity is a generated protocol buffer package.

It is generated from these files:
	identity.proto

It has these top-level messages:
	CallerIdentity
*/
package identity

import prpccommon "github.com/luci/luci-go/common/prpc"
import prpc "github.com/luci/luci-go/server/prpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/luci/luci-go/common/proto/google"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// CallerIdentity is returned by GetCallerIdentity.
//
// It contains identity of a caller as understood by auth system
// (e.g. "user:<email>"), as well as additional authentication related
// information about the requester.
type CallerIdentity struct {
	// Identity of a caller conveyed by their authentication (and possibly
	// delegation) tokens.
	//
	// It is the identity used in all authorization checks. Usually matches
	// peer_identity, but may be different if delegation is used.
	Identity string `protobuf:"bytes,1,opt,name=identity" json:"identity,omitempty"`
	// Identity of whoever is making the request ignoring delegation tokens.
	//
	// It's an identity directly extracted from user credentials.
	PeerIdentity string `protobuf:"bytes,2,opt,name=peer_identity,json=peerIdentity" json:"peer_identity,omitempty"`
	// IP address of the caller as seen by the server.
	PeerIp string `protobuf:"bytes,3,opt,name=peer_ip,json=peerIp" json:"peer_ip,omitempty"`
	// Client ID is set if the caller is using OAuth2 access token for
	// authentication.
	//
	// It is OAuth2 client ID used when making the token.
	Oauth2ClientId string `protobuf:"bytes,4,opt,name=oauth2_client_id,json=oauth2ClientId" json:"oauth2_client_id,omitempty"`
}

func (m *CallerIdentity) Reset()                    { *m = CallerIdentity{} }
func (m *CallerIdentity) String() string            { return proto.CompactTextString(m) }
func (*CallerIdentity) ProtoMessage()               {}
func (*CallerIdentity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*CallerIdentity)(nil), "tokenserver.identity.CallerIdentity")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for IdentityFetcher service

type IdentityFetcherClient interface {
	// GetCallerIdentity returns caller identity as understood by the auth layer.
	//
	// Is uses various authentication tokens supplied by the caller to
	// authenticate the request. It exercises exact same authentication paths as
	// regular services. Useful for debugging access tokens.
	GetCallerIdentity(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*CallerIdentity, error)
}
type identityFetcherPRPCClient struct {
	client *prpccommon.Client
}

func NewIdentityFetcherPRPCClient(client *prpccommon.Client) IdentityFetcherClient {
	return &identityFetcherPRPCClient{client}
}

func (c *identityFetcherPRPCClient) GetCallerIdentity(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*CallerIdentity, error) {
	out := new(CallerIdentity)
	err := c.client.Call(ctx, "tokenserver.identity.IdentityFetcher", "GetCallerIdentity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type identityFetcherClient struct {
	cc *grpc.ClientConn
}

func NewIdentityFetcherClient(cc *grpc.ClientConn) IdentityFetcherClient {
	return &identityFetcherClient{cc}
}

func (c *identityFetcherClient) GetCallerIdentity(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*CallerIdentity, error) {
	out := new(CallerIdentity)
	err := grpc.Invoke(ctx, "/tokenserver.identity.IdentityFetcher/GetCallerIdentity", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IdentityFetcher service

type IdentityFetcherServer interface {
	// GetCallerIdentity returns caller identity as understood by the auth layer.
	//
	// Is uses various authentication tokens supplied by the caller to
	// authenticate the request. It exercises exact same authentication paths as
	// regular services. Useful for debugging access tokens.
	GetCallerIdentity(context.Context, *google_protobuf.Empty) (*CallerIdentity, error)
}

func RegisterIdentityFetcherServer(s prpc.Registrar, srv IdentityFetcherServer) {
	s.RegisterService(&_IdentityFetcher_serviceDesc, srv)
}

func _IdentityFetcher_GetCallerIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityFetcherServer).GetCallerIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenserver.identity.IdentityFetcher/GetCallerIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityFetcherServer).GetCallerIdentity(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _IdentityFetcher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tokenserver.identity.IdentityFetcher",
	HandlerType: (*IdentityFetcherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCallerIdentity",
			Handler:    _IdentityFetcher_GetCallerIdentity_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x4c, 0x49, 0xcd,
	0x2b, 0xc9, 0x2c, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x29, 0xc9, 0xcf, 0x4e,
	0xcd, 0x2b, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x83, 0xc9, 0x49, 0x49, 0xa7, 0xe7, 0xe7, 0xa7,
	0xe7, 0xa4, 0xea, 0x83, 0xd5, 0x24, 0x95, 0xa6, 0xe9, 0xa7, 0xe6, 0x16, 0xc0, 0xb4, 0x28, 0x4d,
	0x61, 0xe4, 0xe2, 0x73, 0x4e, 0xcc, 0xc9, 0x49, 0x2d, 0xf2, 0x84, 0xaa, 0x17, 0x92, 0xe2, 0xe2,
	0x80, 0xe9, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf3, 0x85, 0x94, 0xb9, 0x78, 0x0b,
	0x52, 0x53, 0x8b, 0xe2, 0xe1, 0x0a, 0x98, 0xc0, 0x0a, 0x78, 0x40, 0x82, 0x70, 0x03, 0xc4, 0xb9,
	0xd8, 0x21, 0x8a, 0x0a, 0x24, 0x98, 0xc1, 0xd2, 0x6c, 0x60, 0xe9, 0x02, 0x21, 0x0d, 0x2e, 0x81,
	0xfc, 0xc4, 0xd2, 0x92, 0x0c, 0xa3, 0xf8, 0xe4, 0x9c, 0x4c, 0xa0, 0x62, 0xa0, 0x31, 0x12, 0x2c,
	0x60, 0x15, 0x7c, 0x10, 0x71, 0x67, 0xb0, 0xb0, 0x67, 0x8a, 0x51, 0x0a, 0x17, 0x3f, 0xcc, 0x38,
	0xb7, 0xd4, 0x92, 0xe4, 0x8c, 0xd4, 0x22, 0xa1, 0x40, 0x2e, 0x41, 0xf7, 0xd4, 0x12, 0x34, 0xb7,
	0x8a, 0xe9, 0x41, 0x3c, 0xa7, 0x07, 0xf3, 0x9c, 0x9e, 0x2b, 0xc8, 0x73, 0x52, 0x2a, 0x7a, 0xd8,
	0x82, 0x42, 0x0f, 0x55, 0x77, 0x12, 0x1b, 0x58, 0x97, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x98,
	0x8f, 0x4e, 0xf3, 0x48, 0x01, 0x00, 0x00,
}