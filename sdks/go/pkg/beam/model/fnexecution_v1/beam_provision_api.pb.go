// Code generated by protoc-gen-go. DO NOT EDIT.
// source: beam_provision_api.proto

package fnexecution_v1

import (
	context "context"
	fmt "fmt"
	pipeline_v1 "github.com/apache/beam/sdks/go/pkg/beam/model/pipeline_v1"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// A request to get the provision info of a SDK harness worker instance.
type GetProvisionInfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProvisionInfoRequest) Reset()         { *m = GetProvisionInfoRequest{} }
func (m *GetProvisionInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetProvisionInfoRequest) ProtoMessage()    {}
func (*GetProvisionInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_92e393e5933c7d6f, []int{0}
}

func (m *GetProvisionInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProvisionInfoRequest.Unmarshal(m, b)
}
func (m *GetProvisionInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProvisionInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetProvisionInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProvisionInfoRequest.Merge(m, src)
}
func (m *GetProvisionInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetProvisionInfoRequest.Size(m)
}
func (m *GetProvisionInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProvisionInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProvisionInfoRequest proto.InternalMessageInfo

// A response containing the provision info of a SDK harness worker instance.
type GetProvisionInfoResponse struct {
	Info                 *ProvisionInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetProvisionInfoResponse) Reset()         { *m = GetProvisionInfoResponse{} }
func (m *GetProvisionInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetProvisionInfoResponse) ProtoMessage()    {}
func (*GetProvisionInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_92e393e5933c7d6f, []int{1}
}

func (m *GetProvisionInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProvisionInfoResponse.Unmarshal(m, b)
}
func (m *GetProvisionInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProvisionInfoResponse.Marshal(b, m, deterministic)
}
func (m *GetProvisionInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProvisionInfoResponse.Merge(m, src)
}
func (m *GetProvisionInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetProvisionInfoResponse.Size(m)
}
func (m *GetProvisionInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProvisionInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetProvisionInfoResponse proto.InternalMessageInfo

func (m *GetProvisionInfoResponse) GetInfo() *ProvisionInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

// Runtime provisioning information for a SDK harness worker instance,
// such as pipeline options, resource constraints and other job metadata
type ProvisionInfo struct {
	// (required) Pipeline options. For non-template jobs, the options are
	// identical to what is passed to job submission.
	PipelineOptions *_struct.Struct `protobuf:"bytes,3,opt,name=pipeline_options,json=pipelineOptions,proto3" json:"pipeline_options,omitempty"`
	// (required) The artifact retrieval token produced by
	// LegacyArtifactStagingService.CommitManifestResponse.
	RetrievalToken string `protobuf:"bytes,6,opt,name=retrieval_token,json=retrievalToken,proto3" json:"retrieval_token,omitempty"`
	// (optional) The endpoint that the runner is hosting for the SDK to submit
	// status reports to during pipeline execution. This field will only be
	// populated if the runner supports SDK status reports. For more details see
	// https://s.apache.org/beam-fn-api-harness-status
	StatusEndpoint *pipeline_v1.ApiServiceDescriptor `protobuf:"bytes,7,opt,name=status_endpoint,json=statusEndpoint,proto3" json:"status_endpoint,omitempty"`
	// (optional) The logging endpoint this SDK should use.
	LoggingEndpoint *pipeline_v1.ApiServiceDescriptor `protobuf:"bytes,8,opt,name=logging_endpoint,json=loggingEndpoint,proto3" json:"logging_endpoint,omitempty"`
	// (optional) The artifact retrieval endpoint this SDK should use.
	ArtifactEndpoint *pipeline_v1.ApiServiceDescriptor `protobuf:"bytes,9,opt,name=artifact_endpoint,json=artifactEndpoint,proto3" json:"artifact_endpoint,omitempty"`
	// (optional) The control endpoint this SDK should use.
	ControlEndpoint *pipeline_v1.ApiServiceDescriptor `protobuf:"bytes,10,opt,name=control_endpoint,json=controlEndpoint,proto3" json:"control_endpoint,omitempty"`
	// The set of dependencies that should be staged into this environment.
	Dependencies         []*pipeline_v1.ArtifactInformation `protobuf:"bytes,11,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *ProvisionInfo) Reset()         { *m = ProvisionInfo{} }
func (m *ProvisionInfo) String() string { return proto.CompactTextString(m) }
func (*ProvisionInfo) ProtoMessage()    {}
func (*ProvisionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_92e393e5933c7d6f, []int{2}
}

func (m *ProvisionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProvisionInfo.Unmarshal(m, b)
}
func (m *ProvisionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProvisionInfo.Marshal(b, m, deterministic)
}
func (m *ProvisionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProvisionInfo.Merge(m, src)
}
func (m *ProvisionInfo) XXX_Size() int {
	return xxx_messageInfo_ProvisionInfo.Size(m)
}
func (m *ProvisionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProvisionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProvisionInfo proto.InternalMessageInfo

func (m *ProvisionInfo) GetPipelineOptions() *_struct.Struct {
	if m != nil {
		return m.PipelineOptions
	}
	return nil
}

func (m *ProvisionInfo) GetRetrievalToken() string {
	if m != nil {
		return m.RetrievalToken
	}
	return ""
}

func (m *ProvisionInfo) GetStatusEndpoint() *pipeline_v1.ApiServiceDescriptor {
	if m != nil {
		return m.StatusEndpoint
	}
	return nil
}

func (m *ProvisionInfo) GetLoggingEndpoint() *pipeline_v1.ApiServiceDescriptor {
	if m != nil {
		return m.LoggingEndpoint
	}
	return nil
}

func (m *ProvisionInfo) GetArtifactEndpoint() *pipeline_v1.ApiServiceDescriptor {
	if m != nil {
		return m.ArtifactEndpoint
	}
	return nil
}

func (m *ProvisionInfo) GetControlEndpoint() *pipeline_v1.ApiServiceDescriptor {
	if m != nil {
		return m.ControlEndpoint
	}
	return nil
}

func (m *ProvisionInfo) GetDependencies() []*pipeline_v1.ArtifactInformation {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

func init() {
	proto.RegisterType((*GetProvisionInfoRequest)(nil), "org.apache.beam.model.fn_execution.v1.GetProvisionInfoRequest")
	proto.RegisterType((*GetProvisionInfoResponse)(nil), "org.apache.beam.model.fn_execution.v1.GetProvisionInfoResponse")
	proto.RegisterType((*ProvisionInfo)(nil), "org.apache.beam.model.fn_execution.v1.ProvisionInfo")
}

func init() { proto.RegisterFile("beam_provision_api.proto", fileDescriptor_92e393e5933c7d6f) }

var fileDescriptor_92e393e5933c7d6f = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x65, 0xb5, 0x2a, 0x74, 0x53, 0x62, 0x63, 0x09, 0xd5, 0x44, 0x1c, 0xa2, 0x08, 0x44,
	0x4e, 0x5b, 0x35, 0x20, 0xb8, 0x81, 0x1a, 0x15, 0x01, 0x27, 0x90, 0xcb, 0xa9, 0x17, 0xb3, 0xb1,
	0xc7, 0x66, 0x85, 0x33, 0xb3, 0xec, 0xae, 0x2d, 0xde, 0x83, 0x97, 0xe0, 0xd1, 0x78, 0x0c, 0xe4,
	0xff, 0x09, 0x28, 0x52, 0xd4, 0x1c, 0x3d, 0xb3, 0xf3, 0xfb, 0xe6, 0xb3, 0xbe, 0x61, 0xc1, 0x0a,
	0xc4, 0x3a, 0x52, 0x9a, 0x4a, 0x69, 0x24, 0x61, 0x24, 0x94, 0xe4, 0x4a, 0x93, 0x25, 0xff, 0x19,
	0xe9, 0x8c, 0x0b, 0x25, 0xe2, 0x6f, 0xc0, 0xab, 0x47, 0x7c, 0x4d, 0x09, 0xe4, 0x3c, 0xc5, 0x08,
	0x7e, 0x42, 0x5c, 0x58, 0x49, 0xc8, 0xcb, 0xcb, 0xc9, 0xa3, 0x1a, 0xa0, 0x0b, 0x44, 0xd0, 0xc3,
	0xf4, 0xc4, 0x05, 0x4c, 0x14, 0x49, 0xb4, 0xa6, 0x2d, 0x3c, 0xc9, 0x88, 0xb2, 0x1c, 0x2e, 0xea,
	0xaf, 0x55, 0x91, 0x5e, 0x18, 0xab, 0x8b, 0xd8, 0x36, 0xdd, 0xd9, 0x63, 0x76, 0xfe, 0x1e, 0xec,
	0xe7, 0x6e, 0x8d, 0x8f, 0x98, 0x52, 0x08, 0x3f, 0x0a, 0x30, 0x76, 0x96, 0xb0, 0xe0, 0xff, 0x96,
	0x51, 0x84, 0x06, 0xfc, 0x0f, 0xec, 0x58, 0x62, 0x4a, 0x81, 0x33, 0x75, 0xe6, 0xa3, 0xc5, 0x4b,
	0xbe, 0xd7, 0xca, 0x7c, 0x9b, 0x55, 0x13, 0x66, 0x7f, 0x8e, 0xd9, 0x83, 0xad, 0xba, 0xbf, 0x64,
	0x9e, 0x92, 0x0a, 0x72, 0x89, 0x10, 0x91, 0xaa, 0x66, 0x4d, 0x70, 0x54, 0xeb, 0x9c, 0xf3, 0xc6,
	0x0b, 0xef, 0xbc, 0xf0, 0x9b, 0xda, 0x4b, 0xe8, 0x76, 0x03, 0x9f, 0x9a, 0xf7, 0xfe, 0x73, 0xe6,
	0x6a, 0xb0, 0x5a, 0x42, 0x29, 0xf2, 0xc8, 0xd2, 0x77, 0xc0, 0xe0, 0x64, 0xea, 0xcc, 0x4f, 0xc3,
	0x71, 0x5f, 0xfe, 0x52, 0x55, 0xfd, 0xaf, 0xcc, 0x35, 0x56, 0xd8, 0xc2, 0x44, 0xdd, 0x7f, 0x0b,
	0xee, 0xd5, 0x5a, 0xaf, 0x77, 0x78, 0xea, 0x94, 0x2a, 0x3f, 0x57, 0x4a, 0xde, 0x80, 0x2e, 0x65,
	0x0c, 0xd7, 0x60, 0x62, 0x2d, 0x95, 0x25, 0x1d, 0x8e, 0x1b, 0xde, 0xbb, 0x16, 0xe7, 0xaf, 0x98,
	0x97, 0x53, 0x96, 0x49, 0xcc, 0x06, 0x89, 0xfb, 0x87, 0x49, 0xb8, 0x2d, 0xb0, 0xd7, 0x48, 0xd8,
	0x43, 0xa1, 0xad, 0x4c, 0x45, 0x6c, 0x07, 0x91, 0xd3, 0xc3, 0x44, 0xbc, 0x8e, 0xb8, 0xe9, 0x24,
	0x26, 0xb4, 0x9a, 0xf2, 0x41, 0x84, 0x1d, 0xe8, 0xa4, 0x05, 0xf6, 0x1a, 0xb7, 0xec, 0x2c, 0x01,
	0x05, 0x98, 0x00, 0xc6, 0x12, 0x4c, 0x30, 0x9a, 0x1e, 0xcd, 0x47, 0x8b, 0x57, 0xfb, 0xf0, 0xdb,
	0x75, 0xab, 0x0c, 0xe9, 0xb5, 0xa8, 0x72, 0x10, 0x6e, 0xb1, 0x16, 0xbf, 0x1d, 0xe6, 0xf5, 0x51,
	0x6b, 0x77, 0xf1, 0x7f, 0x39, 0xcc, 0xfb, 0x37, 0xe6, 0xfe, 0x9b, 0x3d, 0x03, 0xbd, 0xe3, 0x74,
	0x26, 0x6f, 0xef, 0x3c, 0xdf, 0xdc, 0xd7, 0xf2, 0x9a, 0x3d, 0xdd, 0x45, 0xd8, 0x04, 0x2c, 0xcf,
	0xfa, 0xf1, 0x2b, 0x25, 0x6f, 0xc7, 0x1b, 0xdd, 0xa8, 0xbc, 0x5c, 0x9d, 0xd4, 0x77, 0xf2, 0xe2,
	0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xca, 0x2c, 0x91, 0x7e, 0x6c, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProvisionServiceClient is the client API for ProvisionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProvisionServiceClient interface {
	// Get provision information for the SDK harness worker instance.
	GetProvisionInfo(ctx context.Context, in *GetProvisionInfoRequest, opts ...grpc.CallOption) (*GetProvisionInfoResponse, error)
}

type provisionServiceClient struct {
	cc *grpc.ClientConn
}

func NewProvisionServiceClient(cc *grpc.ClientConn) ProvisionServiceClient {
	return &provisionServiceClient{cc}
}

func (c *provisionServiceClient) GetProvisionInfo(ctx context.Context, in *GetProvisionInfoRequest, opts ...grpc.CallOption) (*GetProvisionInfoResponse, error) {
	out := new(GetProvisionInfoResponse)
	err := c.cc.Invoke(ctx, "/org.apache.beam.model.fn_execution.v1.ProvisionService/GetProvisionInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProvisionServiceServer is the server API for ProvisionService service.
type ProvisionServiceServer interface {
	// Get provision information for the SDK harness worker instance.
	GetProvisionInfo(context.Context, *GetProvisionInfoRequest) (*GetProvisionInfoResponse, error)
}

// UnimplementedProvisionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProvisionServiceServer struct {
}

func (*UnimplementedProvisionServiceServer) GetProvisionInfo(ctx context.Context, req *GetProvisionInfoRequest) (*GetProvisionInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProvisionInfo not implemented")
}

func RegisterProvisionServiceServer(s *grpc.Server, srv ProvisionServiceServer) {
	s.RegisterService(&_ProvisionService_serviceDesc, srv)
}

func _ProvisionService_GetProvisionInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProvisionInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvisionServiceServer).GetProvisionInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.apache.beam.model.fn_execution.v1.ProvisionService/GetProvisionInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvisionServiceServer).GetProvisionInfo(ctx, req.(*GetProvisionInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProvisionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "org.apache.beam.model.fn_execution.v1.ProvisionService",
	HandlerType: (*ProvisionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProvisionInfo",
			Handler:    _ProvisionService_GetProvisionInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "beam_provision_api.proto",
}
