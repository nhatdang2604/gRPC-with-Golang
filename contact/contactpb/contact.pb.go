// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.12.4
// source: contact/contactpb/contact.proto

package contactpb

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

type Contact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PhoneNumber string `protobuf:"bytes,2,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Address     string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Contact) Reset() {
	*x = Contact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_contactpb_contact_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contact) ProtoMessage() {}

func (x *Contact) ProtoReflect() protoreflect.Message {
	mi := &file_contact_contactpb_contact_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contact.ProtoReflect.Descriptor instead.
func (*Contact) Descriptor() ([]byte, []int) {
	return file_contact_contactpb_contact_proto_rawDescGZIP(), []int{0}
}

func (x *Contact) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Contact) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *Contact) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Contact) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type InsertContactRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contact *Contact `protobuf:"bytes,1,opt,name=contact,proto3" json:"contact,omitempty"`
}

func (x *InsertContactRequest) Reset() {
	*x = InsertContactRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_contactpb_contact_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertContactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertContactRequest) ProtoMessage() {}

func (x *InsertContactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contact_contactpb_contact_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertContactRequest.ProtoReflect.Descriptor instead.
func (*InsertContactRequest) Descriptor() ([]byte, []int) {
	return file_contact_contactpb_contact_proto_rawDescGZIP(), []int{1}
}

func (x *InsertContactRequest) GetContact() *Contact {
	if x != nil {
		return x.Contact
	}
	return nil
}

type InsertContactResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"` //0 == statusCode <=> success
	Message    string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *InsertContactResponse) Reset() {
	*x = InsertContactResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_contactpb_contact_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertContactResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertContactResponse) ProtoMessage() {}

func (x *InsertContactResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contact_contactpb_contact_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertContactResponse.ProtoReflect.Descriptor instead.
func (*InsertContactResponse) Descriptor() ([]byte, []int) {
	return file_contact_contactpb_contact_proto_rawDescGZIP(), []int{2}
}

func (x *InsertContactResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *InsertContactResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_contact_contactpb_contact_proto protoreflect.FileDescriptor

var file_contact_contactpb_contact_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x22, 0x69, 0x0a, 0x07, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x42, 0x0a, 0x14, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x22, 0x51, 0x0a, 0x15, 0x49, 0x6e, 0x73,
	0x65, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x5b, 0x0a, 0x0e,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49,
	0x0a, 0x06, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x1d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x15, 0x5a, 0x13, 0x2e, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contact_contactpb_contact_proto_rawDescOnce sync.Once
	file_contact_contactpb_contact_proto_rawDescData = file_contact_contactpb_contact_proto_rawDesc
)

func file_contact_contactpb_contact_proto_rawDescGZIP() []byte {
	file_contact_contactpb_contact_proto_rawDescOnce.Do(func() {
		file_contact_contactpb_contact_proto_rawDescData = protoimpl.X.CompressGZIP(file_contact_contactpb_contact_proto_rawDescData)
	})
	return file_contact_contactpb_contact_proto_rawDescData
}

var file_contact_contactpb_contact_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_contact_contactpb_contact_proto_goTypes = []interface{}{
	(*Contact)(nil),               // 0: contact.Contact
	(*InsertContactRequest)(nil),  // 1: contact.InsertContactRequest
	(*InsertContactResponse)(nil), // 2: contact.InsertContactResponse
}
var file_contact_contactpb_contact_proto_depIdxs = []int32{
	0, // 0: contact.InsertContactRequest.contact:type_name -> contact.Contact
	1, // 1: contact.ContactService.Insert:input_type -> contact.InsertContactRequest
	2, // 2: contact.ContactService.Insert:output_type -> contact.InsertContactResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_contact_contactpb_contact_proto_init() }
func file_contact_contactpb_contact_proto_init() {
	if File_contact_contactpb_contact_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contact_contactpb_contact_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contact); i {
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
		file_contact_contactpb_contact_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertContactRequest); i {
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
		file_contact_contactpb_contact_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertContactResponse); i {
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
			RawDescriptor: file_contact_contactpb_contact_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_contact_contactpb_contact_proto_goTypes,
		DependencyIndexes: file_contact_contactpb_contact_proto_depIdxs,
		MessageInfos:      file_contact_contactpb_contact_proto_msgTypes,
	}.Build()
	File_contact_contactpb_contact_proto = out.File
	file_contact_contactpb_contact_proto_rawDesc = nil
	file_contact_contactpb_contact_proto_goTypes = nil
	file_contact_contactpb_contact_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ContactServiceClient is the client API for ContactService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContactServiceClient interface {
	Insert(ctx context.Context, in *InsertContactRequest, opts ...grpc.CallOption) (*InsertContactResponse, error)
}

type contactServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContactServiceClient(cc grpc.ClientConnInterface) ContactServiceClient {
	return &contactServiceClient{cc}
}

func (c *contactServiceClient) Insert(ctx context.Context, in *InsertContactRequest, opts ...grpc.CallOption) (*InsertContactResponse, error) {
	out := new(InsertContactResponse)
	err := c.cc.Invoke(ctx, "/contact.ContactService/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContactServiceServer is the server API for ContactService service.
type ContactServiceServer interface {
	Insert(context.Context, *InsertContactRequest) (*InsertContactResponse, error)
}

// UnimplementedContactServiceServer can be embedded to have forward compatible implementations.
type UnimplementedContactServiceServer struct {
}

func (*UnimplementedContactServiceServer) Insert(context.Context, *InsertContactRequest) (*InsertContactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}

func RegisterContactServiceServer(s *grpc.Server, srv ContactServiceServer) {
	s.RegisterService(&_ContactService_serviceDesc, srv)
}

func _ContactService_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contact.ContactService/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).Insert(ctx, req.(*InsertContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContactService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "contact.ContactService",
	HandlerType: (*ContactServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Insert",
			Handler:    _ContactService_Insert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contact/contactpb/contact.proto",
}