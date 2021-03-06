// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/iam/credentials/v1/iamcredentials.proto

package credentials

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

func init() {
	proto.RegisterFile("google/iam/credentials/v1/iamcredentials.proto", fileDescriptor_ad032f4dbfa7437c)
}

var fileDescriptor_ad032f4dbfa7437c = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4f, 0x4b, 0xe3, 0x40,
	0x18, 0xc6, 0x99, 0x85, 0xfd, 0x43, 0x0e, 0xbb, 0x90, 0x9e, 0xb6, 0xec, 0x29, 0x0b, 0x0b, 0x9b,
	0x85, 0xcc, 0xb6, 0xbb, 0x55, 0x48, 0x15, 0x6c, 0x14, 0xa5, 0x05, 0xa1, 0xa8, 0x27, 0x6f, 0xd3,
	0xe9, 0x30, 0x8c, 0x26, 0xf3, 0xc6, 0xcc, 0xb4, 0x3d, 0x88, 0x17, 0xc1, 0x4f, 0xe0, 0xd5, 0x83,
	0x1f, 0xc4, 0x8f, 0xe0, 0xcd, 0x2f, 0xe0, 0xc1, 0x0f, 0xe1, 0x51, 0x26, 0x99, 0x60, 0xc5, 0x56,
	0x93, 0xe3, 0xbc, 0xf3, 0x3c, 0xef, 0xf3, 0x7b, 0x0e, 0xaf, 0x13, 0x70, 0x00, 0x1e, 0x33, 0x2c,
	0x48, 0x82, 0x69, 0xc6, 0xc6, 0x4c, 0x6a, 0x41, 0x62, 0x85, 0xa7, 0x2d, 0x33, 0x9a, 0x9b, 0x04,
	0x69, 0x06, 0x1a, 0xdc, 0xef, 0x85, 0x3e, 0x10, 0x24, 0x09, 0xe6, 0x7f, 0xa7, 0xad, 0xe6, 0x0f,
	0xbb, 0x8a, 0xa4, 0x02, 0x13, 0x29, 0x41, 0x13, 0x2d, 0x40, 0x5a, 0x63, 0xf3, 0xd7, 0xf2, 0x20,
	0x0a, 0x49, 0x02, 0xb2, 0xd0, 0xb5, 0xef, 0x3f, 0x3a, 0x5f, 0xfb, 0xbd, 0xdd, 0xcd, 0x67, 0x89,
	0x7b, 0x8b, 0x9c, 0xc6, 0x0e, 0x93, 0x2c, 0x23, 0x9a, 0xf5, 0x28, 0x65, 0x4a, 0x1d, 0xc0, 0x31,
	0x93, 0x6e, 0x27, 0x58, 0x0a, 0x13, 0x2c, 0xd0, 0xef, 0xb1, 0x93, 0x09, 0x53, 0xba, 0xb9, 0x52,
	0xd7, 0xa6, 0x52, 0x90, 0x8a, 0x79, 0xdb, 0xe7, 0x77, 0x0f, 0x97, 0x1f, 0x36, 0xbc, 0xae, 0x61,
	0x3e, 0x95, 0x24, 0x61, 0xeb, 0x69, 0x06, 0x47, 0x8c, 0x6a, 0x85, 0x7d, 0xac, 0x58, 0x36, 0x15,
	0xd4, 0x18, 0x61, 0x22, 0xcd, 0xe4, 0x2c, 0xe4, 0xaf, 0x97, 0x85, 0xc8, 0x77, 0x6f, 0x90, 0xf3,
	0xad, 0xcc, 0xe9, 0x8f, 0x8b, 0x2a, 0xad, 0x0a, 0x4c, 0x56, 0x5b, 0xd6, 0x68, 0xd7, 0xb1, 0xd8,
	0x0a, 0x51, 0x5e, 0x61, 0xcd, 0x5b, 0xad, 0x5b, 0xc1, 0x2e, 0x32, 0xf8, 0xd7, 0xc8, 0xf9, 0xb2,
	0x2f, 0xb8, 0x8c, 0x62, 0x18, 0xb9, 0xfe, 0x1b, 0x10, 0xa5, 0xa8, 0x04, 0xfe, 0x53, 0x49, 0x6b,
	0x49, 0xbb, 0x39, 0x69, 0xc7, 0xfb, 0x5b, 0x95, 0x54, 0xd9, 0x0d, 0x06, 0xf1, 0x0a, 0x39, 0x9f,
	0xcd, 0xc6, 0xc1, 0x4c, 0xbb, 0xbf, 0xdf, 0x49, 0x1d, 0xcc, 0x74, 0x09, 0xe8, 0x57, 0x91, 0x5a,
	0xbe, 0x30, 0xe7, 0xfb, 0xef, 0xe1, 0x3a, 0x7c, 0x83, 0x99, 0x0e, 0x91, 0x1f, 0x5d, 0x20, 0xe7,
	0x27, 0x85, 0xa4, 0x4c, 0xa3, 0x31, 0x4c, 0xc6, 0x0b, 0x32, 0xa3, 0xc6, 0xcb, 0x3b, 0x18, 0x9a,
	0xfb, 0x18, 0xa2, 0xc3, 0x2d, 0xeb, 0xe3, 0x10, 0x13, 0xc9, 0x03, 0xc8, 0x38, 0xe6, 0x4c, 0xe6,
	0xd7, 0x83, 0x8b, 0x2f, 0x92, 0x0a, 0xb5, 0xe0, 0xd0, 0xba, 0x73, 0xcf, 0x47, 0x84, 0x46, 0x9f,
	0x72, 0xcf, 0xbf, 0xa7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x83, 0x8b, 0xa8, 0x74, 0x04, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IAMCredentialsClient is the client API for IAMCredentials service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IAMCredentialsClient interface {
	// Generates an OAuth 2.0 access token for a service account.
	GenerateAccessToken(ctx context.Context, in *GenerateAccessTokenRequest, opts ...grpc.CallOption) (*GenerateAccessTokenResponse, error)
	// Generates an OpenID Connect ID token for a service account.
	GenerateIdToken(ctx context.Context, in *GenerateIdTokenRequest, opts ...grpc.CallOption) (*GenerateIdTokenResponse, error)
	// Signs a blob using a service account's system-managed private key.
	SignBlob(ctx context.Context, in *SignBlobRequest, opts ...grpc.CallOption) (*SignBlobResponse, error)
	// Signs a JWT using a service account's system-managed private key.
	SignJwt(ctx context.Context, in *SignJwtRequest, opts ...grpc.CallOption) (*SignJwtResponse, error)
}

type iAMCredentialsClient struct {
	cc *grpc.ClientConn
}

func NewIAMCredentialsClient(cc *grpc.ClientConn) IAMCredentialsClient {
	return &iAMCredentialsClient{cc}
}

func (c *iAMCredentialsClient) GenerateAccessToken(ctx context.Context, in *GenerateAccessTokenRequest, opts ...grpc.CallOption) (*GenerateAccessTokenResponse, error) {
	out := new(GenerateAccessTokenResponse)
	err := c.cc.Invoke(ctx, "/google.iam.credentials.v1.IAMCredentials/GenerateAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMCredentialsClient) GenerateIdToken(ctx context.Context, in *GenerateIdTokenRequest, opts ...grpc.CallOption) (*GenerateIdTokenResponse, error) {
	out := new(GenerateIdTokenResponse)
	err := c.cc.Invoke(ctx, "/google.iam.credentials.v1.IAMCredentials/GenerateIdToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMCredentialsClient) SignBlob(ctx context.Context, in *SignBlobRequest, opts ...grpc.CallOption) (*SignBlobResponse, error) {
	out := new(SignBlobResponse)
	err := c.cc.Invoke(ctx, "/google.iam.credentials.v1.IAMCredentials/SignBlob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMCredentialsClient) SignJwt(ctx context.Context, in *SignJwtRequest, opts ...grpc.CallOption) (*SignJwtResponse, error) {
	out := new(SignJwtResponse)
	err := c.cc.Invoke(ctx, "/google.iam.credentials.v1.IAMCredentials/SignJwt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IAMCredentialsServer is the server API for IAMCredentials service.
type IAMCredentialsServer interface {
	// Generates an OAuth 2.0 access token for a service account.
	GenerateAccessToken(context.Context, *GenerateAccessTokenRequest) (*GenerateAccessTokenResponse, error)
	// Generates an OpenID Connect ID token for a service account.
	GenerateIdToken(context.Context, *GenerateIdTokenRequest) (*GenerateIdTokenResponse, error)
	// Signs a blob using a service account's system-managed private key.
	SignBlob(context.Context, *SignBlobRequest) (*SignBlobResponse, error)
	// Signs a JWT using a service account's system-managed private key.
	SignJwt(context.Context, *SignJwtRequest) (*SignJwtResponse, error)
}

// UnimplementedIAMCredentialsServer can be embedded to have forward compatible implementations.
type UnimplementedIAMCredentialsServer struct {
}

func (*UnimplementedIAMCredentialsServer) GenerateAccessToken(ctx context.Context, req *GenerateAccessTokenRequest) (*GenerateAccessTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateAccessToken not implemented")
}
func (*UnimplementedIAMCredentialsServer) GenerateIdToken(ctx context.Context, req *GenerateIdTokenRequest) (*GenerateIdTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateIdToken not implemented")
}
func (*UnimplementedIAMCredentialsServer) SignBlob(ctx context.Context, req *SignBlobRequest) (*SignBlobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignBlob not implemented")
}
func (*UnimplementedIAMCredentialsServer) SignJwt(ctx context.Context, req *SignJwtRequest) (*SignJwtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignJwt not implemented")
}

func RegisterIAMCredentialsServer(s *grpc.Server, srv IAMCredentialsServer) {
	s.RegisterService(&_IAMCredentials_serviceDesc, srv)
}

func _IAMCredentials_GenerateAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMCredentialsServer).GenerateAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.iam.credentials.v1.IAMCredentials/GenerateAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMCredentialsServer).GenerateAccessToken(ctx, req.(*GenerateAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMCredentials_GenerateIdToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateIdTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMCredentialsServer).GenerateIdToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.iam.credentials.v1.IAMCredentials/GenerateIdToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMCredentialsServer).GenerateIdToken(ctx, req.(*GenerateIdTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMCredentials_SignBlob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignBlobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMCredentialsServer).SignBlob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.iam.credentials.v1.IAMCredentials/SignBlob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMCredentialsServer).SignBlob(ctx, req.(*SignBlobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMCredentials_SignJwt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignJwtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMCredentialsServer).SignJwt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.iam.credentials.v1.IAMCredentials/SignJwt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMCredentialsServer).SignJwt(ctx, req.(*SignJwtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IAMCredentials_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.iam.credentials.v1.IAMCredentials",
	HandlerType: (*IAMCredentialsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateAccessToken",
			Handler:    _IAMCredentials_GenerateAccessToken_Handler,
		},
		{
			MethodName: "GenerateIdToken",
			Handler:    _IAMCredentials_GenerateIdToken_Handler,
		},
		{
			MethodName: "SignBlob",
			Handler:    _IAMCredentials_SignBlob_Handler,
		},
		{
			MethodName: "SignJwt",
			Handler:    _IAMCredentials_SignJwt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/iam/credentials/v1/iamcredentials.proto",
}
