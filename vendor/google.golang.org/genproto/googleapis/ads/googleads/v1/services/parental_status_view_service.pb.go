// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/parental_status_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
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

// Request message for
// [ParentalStatusViewService.GetParentalStatusView][google.ads.googleads.v1.services.ParentalStatusViewService.GetParentalStatusView].
type GetParentalStatusViewRequest struct {
	// The resource name of the parental status view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetParentalStatusViewRequest) Reset()         { *m = GetParentalStatusViewRequest{} }
func (m *GetParentalStatusViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetParentalStatusViewRequest) ProtoMessage()    {}
func (*GetParentalStatusViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3875092f69544412, []int{0}
}

func (m *GetParentalStatusViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetParentalStatusViewRequest.Unmarshal(m, b)
}
func (m *GetParentalStatusViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetParentalStatusViewRequest.Marshal(b, m, deterministic)
}
func (m *GetParentalStatusViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetParentalStatusViewRequest.Merge(m, src)
}
func (m *GetParentalStatusViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetParentalStatusViewRequest.Size(m)
}
func (m *GetParentalStatusViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetParentalStatusViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetParentalStatusViewRequest proto.InternalMessageInfo

func (m *GetParentalStatusViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetParentalStatusViewRequest)(nil), "google.ads.googleads.v1.services.GetParentalStatusViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/parental_status_view_service.proto", fileDescriptor_3875092f69544412)
}

var fileDescriptor_3875092f69544412 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcf, 0x4a, 0xc3, 0x30,
	0x18, 0xa7, 0x15, 0x04, 0x8b, 0x5e, 0x0a, 0x82, 0x8e, 0x21, 0x63, 0xee, 0x20, 0x3b, 0x24, 0x54,
	0x19, 0x42, 0x74, 0x42, 0xb7, 0xc3, 0x3c, 0xc9, 0xd8, 0xa0, 0x07, 0x29, 0x94, 0xb8, 0x86, 0x52,
	0x58, 0x9b, 0xda, 0x2f, 0xed, 0x0e, 0xe2, 0x45, 0xf0, 0x09, 0x7c, 0x03, 0x8f, 0x3e, 0x8a, 0xe0,
	0xc9, 0x57, 0xf0, 0xa4, 0x2f, 0x21, 0x5d, 0x9a, 0x8a, 0xcc, 0xba, 0xdb, 0x8f, 0xe4, 0xf7, 0xe7,
	0xcb, 0xef, 0x8b, 0x31, 0x0c, 0x38, 0x0f, 0xe6, 0x0c, 0x53, 0x1f, 0xb0, 0x84, 0x05, 0xca, 0x2d,
	0x0c, 0x2c, 0xcd, 0xc3, 0x19, 0x03, 0x9c, 0xd0, 0x94, 0xc5, 0x82, 0xce, 0x3d, 0x10, 0x54, 0x64,
	0xe0, 0xe5, 0x21, 0x5b, 0x78, 0xe5, 0x2d, 0x4a, 0x52, 0x2e, 0xb8, 0xd9, 0x92, 0x4a, 0x44, 0x7d,
	0x40, 0x95, 0x09, 0xca, 0x2d, 0xa4, 0x4c, 0x1a, 0xe7, 0x75, 0x31, 0x29, 0x03, 0x9e, 0xa5, 0x75,
	0x39, 0xd2, 0xbf, 0xd1, 0x54, 0xea, 0x24, 0xc4, 0x34, 0x8e, 0xb9, 0xa0, 0x22, 0xe4, 0x31, 0xc8,
	0xdb, 0xf6, 0xd0, 0x68, 0x8e, 0x98, 0x18, 0x97, 0xf2, 0xe9, 0x52, 0xed, 0x84, 0x6c, 0x31, 0x61,
	0xb7, 0x19, 0x03, 0x61, 0x1e, 0x1a, 0x3b, 0x2a, 0xc5, 0x8b, 0x69, 0xc4, 0xf6, 0xb4, 0x96, 0x76,
	0xb4, 0x35, 0xd9, 0x56, 0x87, 0x57, 0x34, 0x62, 0xc7, 0x5f, 0x9a, 0xb1, 0xbf, 0x6a, 0x31, 0x95,
	0xf3, 0x9b, 0x6f, 0x9a, 0xb1, 0xfb, 0x67, 0x86, 0x79, 0x81, 0xd6, 0xbd, 0x1d, 0xfd, 0x37, 0x5c,
	0xa3, 0x57, 0xab, 0xaf, 0x9a, 0x41, 0xab, 0xea, 0x76, 0xff, 0xe1, 0xfd, 0xe3, 0x49, 0x3f, 0x35,
	0x7b, 0x45, 0x87, 0x77, 0xbf, 0x9e, 0xd7, 0x9f, 0x65, 0x20, 0x78, 0xc4, 0x52, 0xc0, 0xdd, 0xaa,
	0xd4, 0x1f, 0x29, 0xe0, 0xee, 0xfd, 0xe0, 0x51, 0x37, 0x3a, 0x33, 0x1e, 0xad, 0x9d, 0x7d, 0x70,
	0x50, 0xdb, 0xc9, 0xb8, 0xe8, 0x7e, 0xac, 0x5d, 0x5f, 0x96, 0x1e, 0x01, 0x9f, 0xd3, 0x38, 0x40,
	0x3c, 0x0d, 0x70, 0xc0, 0xe2, 0xe5, 0x66, 0xd4, 0xa6, 0x93, 0x10, 0xea, 0xff, 0xd7, 0x99, 0x02,
	0xcf, 0xfa, 0xc6, 0xc8, 0xb6, 0x5f, 0xf4, 0xd6, 0x48, 0x1a, 0xda, 0x3e, 0x20, 0x09, 0x0b, 0xe4,
	0x58, 0xa8, 0x0c, 0x86, 0x57, 0x45, 0x71, 0x6d, 0x1f, 0xdc, 0x8a, 0xe2, 0x3a, 0x96, 0xab, 0x28,
	0x9f, 0x7a, 0x47, 0x9e, 0x13, 0x62, 0xfb, 0x40, 0x48, 0x45, 0x22, 0xc4, 0xb1, 0x08, 0x51, 0xb4,
	0x9b, 0xcd, 0xe5, 0x9c, 0x27, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x08, 0x89, 0x70, 0x7e, 0x06,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ParentalStatusViewServiceClient is the client API for ParentalStatusViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ParentalStatusViewServiceClient interface {
	// Returns the requested parental status view in full detail.
	GetParentalStatusView(ctx context.Context, in *GetParentalStatusViewRequest, opts ...grpc.CallOption) (*resources.ParentalStatusView, error)
}

type parentalStatusViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewParentalStatusViewServiceClient(cc *grpc.ClientConn) ParentalStatusViewServiceClient {
	return &parentalStatusViewServiceClient{cc}
}

func (c *parentalStatusViewServiceClient) GetParentalStatusView(ctx context.Context, in *GetParentalStatusViewRequest, opts ...grpc.CallOption) (*resources.ParentalStatusView, error) {
	out := new(resources.ParentalStatusView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.ParentalStatusViewService/GetParentalStatusView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ParentalStatusViewServiceServer is the server API for ParentalStatusViewService service.
type ParentalStatusViewServiceServer interface {
	// Returns the requested parental status view in full detail.
	GetParentalStatusView(context.Context, *GetParentalStatusViewRequest) (*resources.ParentalStatusView, error)
}

// UnimplementedParentalStatusViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedParentalStatusViewServiceServer struct {
}

func (*UnimplementedParentalStatusViewServiceServer) GetParentalStatusView(ctx context.Context, req *GetParentalStatusViewRequest) (*resources.ParentalStatusView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParentalStatusView not implemented")
}

func RegisterParentalStatusViewServiceServer(s *grpc.Server, srv ParentalStatusViewServiceServer) {
	s.RegisterService(&_ParentalStatusViewService_serviceDesc, srv)
}

func _ParentalStatusViewService_GetParentalStatusView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParentalStatusViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParentalStatusViewServiceServer).GetParentalStatusView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.ParentalStatusViewService/GetParentalStatusView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParentalStatusViewServiceServer).GetParentalStatusView(ctx, req.(*GetParentalStatusViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ParentalStatusViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.ParentalStatusViewService",
	HandlerType: (*ParentalStatusViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetParentalStatusView",
			Handler:    _ParentalStatusViewService_GetParentalStatusView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/parental_status_view_service.proto",
}
