// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: broker.proto

package broker

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

type Domcreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip      string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Domname string `protobuf:"bytes,2,opt,name=domname,proto3" json:"domname,omitempty"`
}

func (x *Domcreate) Reset() {
	*x = Domcreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Domcreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Domcreate) ProtoMessage() {}

func (x *Domcreate) ProtoReflect() protoreflect.Message {
	mi := &file_broker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Domcreate.ProtoReflect.Descriptor instead.
func (*Domcreate) Descriptor() ([]byte, []int) {
	return file_broker_proto_rawDescGZIP(), []int{0}
}

func (x *Domcreate) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Domcreate) GetDomname() string {
	if x != nil {
		return x.Domname
	}
	return ""
}

type DomUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OldDom string `protobuf:"bytes,1,opt,name=oldDom,proto3" json:"oldDom,omitempty"`
	NewDom string `protobuf:"bytes,2,opt,name=newDom,proto3" json:"newDom,omitempty"`
}

func (x *DomUpdate) Reset() {
	*x = DomUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomUpdate) ProtoMessage() {}

func (x *DomUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_broker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomUpdate.ProtoReflect.Descriptor instead.
func (*DomUpdate) Descriptor() ([]byte, []int) {
	return file_broker_proto_rawDescGZIP(), []int{1}
}

func (x *DomUpdate) GetOldDom() string {
	if x != nil {
		return x.OldDom
	}
	return ""
}

func (x *DomUpdate) GetNewDom() string {
	if x != nil {
		return x.NewDom
	}
	return ""
}

type Dom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domname string `protobuf:"bytes,1,opt,name=domname,proto3" json:"domname,omitempty"`
}

func (x *Dom) Reset() {
	*x = Dom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dom) ProtoMessage() {}

func (x *Dom) ProtoReflect() protoreflect.Message {
	mi := &file_broker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dom.ProtoReflect.Descriptor instead.
func (*Dom) Descriptor() ([]byte, []int) {
	return file_broker_proto_rawDescGZIP(), []int{2}
}

func (x *Dom) GetDomname() string {
	if x != nil {
		return x.Domname
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip          string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Verificador bool   `protobuf:"varint,2,opt,name=verificador,proto3" json:"verificador,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_broker_proto_msgTypes[3]
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
	return file_broker_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Response) GetVerificador() bool {
	if x != nil {
		return x.Verificador
	}
	return false
}

var File_broker_proto protoreflect.FileDescriptor

var file_broker_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x22, 0x35, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x6f, 0x6d, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x6f, 0x6d, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3b, 0x0a,
	0x09, 0x64, 0x6f, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x6c,
	0x64, 0x44, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x6c, 0x64, 0x44,
	0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x65, 0x77, 0x44, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6e, 0x65, 0x77, 0x44, 0x6f, 0x6d, 0x22, 0x1f, 0x0a, 0x03, 0x64, 0x6f,
	0x6d, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x6f, 0x6d, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x64, 0x6f, 0x6d, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3c, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x64, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x64, 0x6f, 0x72, 0x32, 0xf0, 0x01, 0x0a, 0x0d, 0x42, 0x72,
	0x6f, 0x6b, 0x65, 0x72, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x0d, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x0b, 0x2e, 0x62,
	0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x64, 0x6f, 0x6d, 0x1a, 0x10, 0x2e, 0x62, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30,
	0x01, 0x12, 0x39, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x12, 0x11, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x64, 0x6f, 0x6d, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x1a, 0x10, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x33, 0x0a, 0x0c,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x0b, 0x2e, 0x62,
	0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x64, 0x6f, 0x6d, 0x1a, 0x10, 0x2e, 0x62, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30,
	0x01, 0x12, 0x39, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x12, 0x11, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x64, 0x6f, 0x6d, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x1a, 0x10, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0f, 0x5a, 0x0d,
	0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_broker_proto_rawDescOnce sync.Once
	file_broker_proto_rawDescData = file_broker_proto_rawDesc
)

func file_broker_proto_rawDescGZIP() []byte {
	file_broker_proto_rawDescOnce.Do(func() {
		file_broker_proto_rawDescData = protoimpl.X.CompressGZIP(file_broker_proto_rawDescData)
	})
	return file_broker_proto_rawDescData
}

var file_broker_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_broker_proto_goTypes = []interface{}{
	(*Domcreate)(nil), // 0: broker.domcreate
	(*DomUpdate)(nil), // 1: broker.domUpdate
	(*Dom)(nil),       // 2: broker.dom
	(*Response)(nil),  // 3: broker.response
}
var file_broker_proto_depIdxs = []int32{
	2, // 0: broker.BrokerHandler.ConnectDomain:input_type -> broker.dom
	0, // 1: broker.BrokerHandler.CreateDomain:input_type -> broker.domcreate
	2, // 2: broker.BrokerHandler.DeleteDomain:input_type -> broker.dom
	1, // 3: broker.BrokerHandler.UpdateDomain:input_type -> broker.domUpdate
	3, // 4: broker.BrokerHandler.ConnectDomain:output_type -> broker.response
	3, // 5: broker.BrokerHandler.CreateDomain:output_type -> broker.response
	3, // 6: broker.BrokerHandler.DeleteDomain:output_type -> broker.response
	3, // 7: broker.BrokerHandler.UpdateDomain:output_type -> broker.response
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_broker_proto_init() }
func file_broker_proto_init() {
	if File_broker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_broker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Domcreate); i {
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
		file_broker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomUpdate); i {
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
		file_broker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dom); i {
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
		file_broker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_broker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_broker_proto_goTypes,
		DependencyIndexes: file_broker_proto_depIdxs,
		MessageInfos:      file_broker_proto_msgTypes,
	}.Build()
	File_broker_proto = out.File
	file_broker_proto_rawDesc = nil
	file_broker_proto_goTypes = nil
	file_broker_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BrokerHandlerClient is the client API for BrokerHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BrokerHandlerClient interface {
	ConnectDomain(ctx context.Context, opts ...grpc.CallOption) (BrokerHandler_ConnectDomainClient, error)
	CreateDomain(ctx context.Context, opts ...grpc.CallOption) (BrokerHandler_CreateDomainClient, error)
	DeleteDomain(ctx context.Context, opts ...grpc.CallOption) (BrokerHandler_DeleteDomainClient, error)
	UpdateDomain(ctx context.Context, opts ...grpc.CallOption) (BrokerHandler_UpdateDomainClient, error)
}

type brokerHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewBrokerHandlerClient(cc grpc.ClientConnInterface) BrokerHandlerClient {
	return &brokerHandlerClient{cc}
}

func (c *brokerHandlerClient) ConnectDomain(ctx context.Context, opts ...grpc.CallOption) (BrokerHandler_ConnectDomainClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BrokerHandler_serviceDesc.Streams[0], "/broker.BrokerHandler/ConnectDomain", opts...)
	if err != nil {
		return nil, err
	}
	x := &brokerHandlerConnectDomainClient{stream}
	return x, nil
}

type BrokerHandler_ConnectDomainClient interface {
	Send(*Dom) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type brokerHandlerConnectDomainClient struct {
	grpc.ClientStream
}

func (x *brokerHandlerConnectDomainClient) Send(m *Dom) error {
	return x.ClientStream.SendMsg(m)
}

func (x *brokerHandlerConnectDomainClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *brokerHandlerClient) CreateDomain(ctx context.Context, opts ...grpc.CallOption) (BrokerHandler_CreateDomainClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BrokerHandler_serviceDesc.Streams[1], "/broker.BrokerHandler/CreateDomain", opts...)
	if err != nil {
		return nil, err
	}
	x := &brokerHandlerCreateDomainClient{stream}
	return x, nil
}

type BrokerHandler_CreateDomainClient interface {
	Send(*Domcreate) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type brokerHandlerCreateDomainClient struct {
	grpc.ClientStream
}

func (x *brokerHandlerCreateDomainClient) Send(m *Domcreate) error {
	return x.ClientStream.SendMsg(m)
}

func (x *brokerHandlerCreateDomainClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *brokerHandlerClient) DeleteDomain(ctx context.Context, opts ...grpc.CallOption) (BrokerHandler_DeleteDomainClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BrokerHandler_serviceDesc.Streams[2], "/broker.BrokerHandler/DeleteDomain", opts...)
	if err != nil {
		return nil, err
	}
	x := &brokerHandlerDeleteDomainClient{stream}
	return x, nil
}

type BrokerHandler_DeleteDomainClient interface {
	Send(*Dom) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type brokerHandlerDeleteDomainClient struct {
	grpc.ClientStream
}

func (x *brokerHandlerDeleteDomainClient) Send(m *Dom) error {
	return x.ClientStream.SendMsg(m)
}

func (x *brokerHandlerDeleteDomainClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *brokerHandlerClient) UpdateDomain(ctx context.Context, opts ...grpc.CallOption) (BrokerHandler_UpdateDomainClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BrokerHandler_serviceDesc.Streams[3], "/broker.BrokerHandler/UpdateDomain", opts...)
	if err != nil {
		return nil, err
	}
	x := &brokerHandlerUpdateDomainClient{stream}
	return x, nil
}

type BrokerHandler_UpdateDomainClient interface {
	Send(*DomUpdate) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type brokerHandlerUpdateDomainClient struct {
	grpc.ClientStream
}

func (x *brokerHandlerUpdateDomainClient) Send(m *DomUpdate) error {
	return x.ClientStream.SendMsg(m)
}

func (x *brokerHandlerUpdateDomainClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BrokerHandlerServer is the server API for BrokerHandler service.
type BrokerHandlerServer interface {
	ConnectDomain(BrokerHandler_ConnectDomainServer) error
	CreateDomain(BrokerHandler_CreateDomainServer) error
	DeleteDomain(BrokerHandler_DeleteDomainServer) error
	UpdateDomain(BrokerHandler_UpdateDomainServer) error
}

// UnimplementedBrokerHandlerServer can be embedded to have forward compatible implementations.
type UnimplementedBrokerHandlerServer struct {
}

func (*UnimplementedBrokerHandlerServer) ConnectDomain(BrokerHandler_ConnectDomainServer) error {
	return status.Errorf(codes.Unimplemented, "method ConnectDomain not implemented")
}
func (*UnimplementedBrokerHandlerServer) CreateDomain(BrokerHandler_CreateDomainServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateDomain not implemented")
}
func (*UnimplementedBrokerHandlerServer) DeleteDomain(BrokerHandler_DeleteDomainServer) error {
	return status.Errorf(codes.Unimplemented, "method DeleteDomain not implemented")
}
func (*UnimplementedBrokerHandlerServer) UpdateDomain(BrokerHandler_UpdateDomainServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateDomain not implemented")
}

func RegisterBrokerHandlerServer(s *grpc.Server, srv BrokerHandlerServer) {
	s.RegisterService(&_BrokerHandler_serviceDesc, srv)
}

func _BrokerHandler_ConnectDomain_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BrokerHandlerServer).ConnectDomain(&brokerHandlerConnectDomainServer{stream})
}

type BrokerHandler_ConnectDomainServer interface {
	Send(*Response) error
	Recv() (*Dom, error)
	grpc.ServerStream
}

type brokerHandlerConnectDomainServer struct {
	grpc.ServerStream
}

func (x *brokerHandlerConnectDomainServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *brokerHandlerConnectDomainServer) Recv() (*Dom, error) {
	m := new(Dom)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _BrokerHandler_CreateDomain_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BrokerHandlerServer).CreateDomain(&brokerHandlerCreateDomainServer{stream})
}

type BrokerHandler_CreateDomainServer interface {
	Send(*Response) error
	Recv() (*Domcreate, error)
	grpc.ServerStream
}

type brokerHandlerCreateDomainServer struct {
	grpc.ServerStream
}

func (x *brokerHandlerCreateDomainServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *brokerHandlerCreateDomainServer) Recv() (*Domcreate, error) {
	m := new(Domcreate)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _BrokerHandler_DeleteDomain_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BrokerHandlerServer).DeleteDomain(&brokerHandlerDeleteDomainServer{stream})
}

type BrokerHandler_DeleteDomainServer interface {
	Send(*Response) error
	Recv() (*Dom, error)
	grpc.ServerStream
}

type brokerHandlerDeleteDomainServer struct {
	grpc.ServerStream
}

func (x *brokerHandlerDeleteDomainServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *brokerHandlerDeleteDomainServer) Recv() (*Dom, error) {
	m := new(Dom)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _BrokerHandler_UpdateDomain_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BrokerHandlerServer).UpdateDomain(&brokerHandlerUpdateDomainServer{stream})
}

type BrokerHandler_UpdateDomainServer interface {
	Send(*Response) error
	Recv() (*DomUpdate, error)
	grpc.ServerStream
}

type brokerHandlerUpdateDomainServer struct {
	grpc.ServerStream
}

func (x *brokerHandlerUpdateDomainServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *brokerHandlerUpdateDomainServer) Recv() (*DomUpdate, error) {
	m := new(DomUpdate)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _BrokerHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "broker.BrokerHandler",
	HandlerType: (*BrokerHandlerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConnectDomain",
			Handler:       _BrokerHandler_ConnectDomain_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "CreateDomain",
			Handler:       _BrokerHandler_CreateDomain_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeleteDomain",
			Handler:       _BrokerHandler_DeleteDomain_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "UpdateDomain",
			Handler:       _BrokerHandler_UpdateDomain_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "broker.proto",
}
