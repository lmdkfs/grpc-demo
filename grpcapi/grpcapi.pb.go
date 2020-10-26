// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: grpcapi.proto

package grpcapi

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GrpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input string `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *GrpcRequest) Reset() {
	*x = GrpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcRequest) ProtoMessage() {}

func (x *GrpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpcapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcRequest.ProtoReflect.Descriptor instead.
func (*GrpcRequest) Descriptor() ([]byte, []int) {
	return file_grpcapi_proto_rawDescGZIP(), []int{0}
}

func (x *GrpcRequest) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

type GrpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *GrpcResponse) Reset() {
	*x = GrpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcResponse) ProtoMessage() {}

func (x *GrpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpcapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcResponse.ProtoReflect.Descriptor instead.
func (*GrpcResponse) Descriptor() ([]byte, []int) {
	return file_grpcapi_proto_rawDescGZIP(), []int{1}
}

func (x *GrpcResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_grpcapi_proto protoreflect.FileDescriptor

var file_grpcapi_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x67, 0x72, 0x65, 0x65, 0x74, 0x22, 0x23, 0x0a, 0x0b, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x2a, 0x0a, 0x0c, 0x47,
	0x72, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x47, 0x0a, 0x0b, 0x47, 0x72, 0x70, 0x63, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72,
	0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x67, 0x72, 0x65, 0x65,
	0x74, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x09, 0x5a, 0x07, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_grpcapi_proto_rawDescOnce sync.Once
	file_grpcapi_proto_rawDescData = file_grpcapi_proto_rawDesc
)

func file_grpcapi_proto_rawDescGZIP() []byte {
	file_grpcapi_proto_rawDescOnce.Do(func() {
		file_grpcapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpcapi_proto_rawDescData)
	})
	return file_grpcapi_proto_rawDescData
}

var file_grpcapi_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpcapi_proto_goTypes = []interface{}{
	(*GrpcRequest)(nil),  // 0: greet.GrpcRequest
	(*GrpcResponse)(nil), // 1: greet.GrpcResponse
}
var file_grpcapi_proto_depIdxs = []int32{
	0, // 0: greet.GrpcService.grpcService:input_type -> greet.GrpcRequest
	1, // 1: greet.GrpcService.grpcService:output_type -> greet.GrpcResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpcapi_proto_init() }
func file_grpcapi_proto_init() {
	if File_grpcapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpcapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpcapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpcapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpcapi_proto_goTypes,
		DependencyIndexes: file_grpcapi_proto_depIdxs,
		MessageInfos:      file_grpcapi_proto_msgTypes,
	}.Build()
	File_grpcapi_proto = out.File
	file_grpcapi_proto_rawDesc = nil
	file_grpcapi_proto_goTypes = nil
	file_grpcapi_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GrpcServiceClient is the client API for GrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GrpcServiceClient interface {
	GrpcService(ctx context.Context, in *GrpcRequest, opts ...grpc.CallOption) (*GrpcResponse, error)
}

type grpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGrpcServiceClient(cc grpc.ClientConnInterface) GrpcServiceClient {
	return &grpcServiceClient{cc}
}

func (c *grpcServiceClient) GrpcService(ctx context.Context, in *GrpcRequest, opts ...grpc.CallOption) (*GrpcResponse, error) {
	out := new(GrpcResponse)
	err := c.cc.Invoke(ctx, "/greet.GrpcService/grpcService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GrpcServiceServer is the server API for GrpcService service.
type GrpcServiceServer interface {
	GrpcService(context.Context, *GrpcRequest) (*GrpcResponse, error)
}

// UnimplementedGrpcServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGrpcServiceServer struct {
}

func (*UnimplementedGrpcServiceServer) GrpcService(context.Context, *GrpcRequest) (*GrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GrpcService not implemented")
}

func RegisterGrpcServiceServer(s *grpc.Server, srv GrpcServiceServer) {
	s.RegisterService(&_GrpcService_serviceDesc, srv)
}

func _GrpcService_GrpcService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServiceServer).GrpcService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.GrpcService/GrpcService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServiceServer).GrpcService(ctx, req.(*GrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GrpcService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GrpcService",
	HandlerType: (*GrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "grpcService",
			Handler:    _GrpcService_GrpcService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpcapi.proto",
}