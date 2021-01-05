// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: dns.proto

package dns

import (
	context "context"
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

type Inmessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domname string `protobuf:"bytes,1,opt,name=domname,proto3" json:"domname,omitempty"`
}

func (x *Inmessage) Reset() {
	*x = Inmessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dns_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Inmessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Inmessage) ProtoMessage() {}

func (x *Inmessage) ProtoReflect() protoreflect.Message {
	mi := &file_dns_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Inmessage.ProtoReflect.Descriptor instead.
func (*Inmessage) Descriptor() ([]byte, []int) {
	return file_dns_proto_rawDescGZIP(), []int{0}
}

func (x *Inmessage) GetDomname() string {
	if x != nil {
		return x.Domname
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dns_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_dns_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_dns_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

var File_dns_proto protoreflect.FileDescriptor

var file_dns_proto_rawDesc = []byte{
	0x0a, 0x09, 0x64, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x64, 0x6e, 0x73,
	0x22, 0x25, 0x0a, 0x09, 0x69, 0x6e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x64, 0x6f, 0x6d, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x64, 0x6f, 0x6d, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x70, 0x32, 0x40, 0x0a, 0x0a, 0x44, 0x6e, 0x73, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x12, 0x32, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x70,
	0x12, 0x0e, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x69, 0x6e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x1a, 0x0d, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x44, 0x6e, 0x73, 0x2f, 0x64, 0x6e, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dns_proto_rawDescOnce sync.Once
	file_dns_proto_rawDescData = file_dns_proto_rawDesc
)

func file_dns_proto_rawDescGZIP() []byte {
	file_dns_proto_rawDescOnce.Do(func() {
		file_dns_proto_rawDescData = protoimpl.X.CompressGZIP(file_dns_proto_rawDescData)
	})
	return file_dns_proto_rawDescData
}

var file_dns_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dns_proto_goTypes = []interface{}{
	(*Inmessage)(nil), // 0: dns.inmessage
	(*Response)(nil),  // 1: dns.response
}
var file_dns_proto_depIdxs = []int32{
	0, // 0: dns.DnsHandler.GetDomainIp:input_type -> dns.inmessage
	1, // 1: dns.DnsHandler.GetDomainIp:output_type -> dns.response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dns_proto_init() }
func file_dns_proto_init() {
	if File_dns_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dns_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Inmessage); i {
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
		file_dns_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_dns_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dns_proto_goTypes,
		DependencyIndexes: file_dns_proto_depIdxs,
		MessageInfos:      file_dns_proto_msgTypes,
	}.Build()
	File_dns_proto = out.File
	file_dns_proto_rawDesc = nil
	file_dns_proto_goTypes = nil
	file_dns_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DnsHandlerClient is the client API for DnsHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DnsHandlerClient interface {
	GetDomainIp(ctx context.Context, opts ...grpc.CallOption) (DnsHandler_GetDomainIpClient, error)
}

type dnsHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewDnsHandlerClient(cc grpc.ClientConnInterface) DnsHandlerClient {
	return &dnsHandlerClient{cc}
}

func (c *dnsHandlerClient) GetDomainIp(ctx context.Context, opts ...grpc.CallOption) (DnsHandler_GetDomainIpClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DnsHandler_serviceDesc.Streams[0], "/dns.DnsHandler/GetDomainIp", opts...)
	if err != nil {
		return nil, err
	}
	x := &dnsHandlerGetDomainIpClient{stream}
	return x, nil
}

type DnsHandler_GetDomainIpClient interface {
	Send(*Inmessage) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type dnsHandlerGetDomainIpClient struct {
	grpc.ClientStream
}

func (x *dnsHandlerGetDomainIpClient) Send(m *Inmessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dnsHandlerGetDomainIpClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DnsHandlerServer is the server API for DnsHandler service.
type DnsHandlerServer interface {
	GetDomainIp(DnsHandler_GetDomainIpServer) error
}

// UnimplementedDnsHandlerServer can be embedded to have forward compatible implementations.
type UnimplementedDnsHandlerServer struct {
}

func (*UnimplementedDnsHandlerServer) GetDomainIp(DnsHandler_GetDomainIpServer) error {
	return status.Errorf(codes.Unimplemented, "method GetDomainIp not implemented")
}

func RegisterDnsHandlerServer(s *grpc.Server, srv DnsHandlerServer) {
	s.RegisterService(&_DnsHandler_serviceDesc, srv)
}

func _DnsHandler_GetDomainIp_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DnsHandlerServer).GetDomainIp(&dnsHandlerGetDomainIpServer{stream})
}

type DnsHandler_GetDomainIpServer interface {
	Send(*Response) error
	Recv() (*Inmessage, error)
	grpc.ServerStream
}

type dnsHandlerGetDomainIpServer struct {
	grpc.ServerStream
}

func (x *dnsHandlerGetDomainIpServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dnsHandlerGetDomainIpServer) Recv() (*Inmessage, error) {
	m := new(Inmessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DnsHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dns.DnsHandler",
	HandlerType: (*DnsHandlerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetDomainIp",
			Handler:       _DnsHandler_GetDomainIp_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "dns.proto",
}