// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: protoc/friends.proto

package friends

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

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_friends_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_friends_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_protoc_friends_proto_rawDescGZIP(), []int{0}
}

func (x *Status) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type UserID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserID) Reset() {
	*x = UserID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_friends_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserID) ProtoMessage() {}

func (x *UserID) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_friends_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserID.ProtoReflect.Descriptor instead.
func (*UserID) Descriptor() ([]byte, []int) {
	return file_protoc_friends_proto_rawDescGZIP(), []int{1}
}

func (x *UserID) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FriendList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []int64 `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`
}

func (x *FriendList) Reset() {
	*x = FriendList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_friends_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendList) ProtoMessage() {}

func (x *FriendList) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_friends_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendList.ProtoReflect.Descriptor instead.
func (*FriendList) Descriptor() ([]byte, []int) {
	return file_protoc_friends_proto_rawDescGZIP(), []int{2}
}

func (x *FriendList) GetId() []int64 {
	if x != nil {
		return x.Id
	}
	return nil
}

type Range struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Range int64 `protobuf:"varint,1,opt,name=range,proto3" json:"range,omitempty"`
}

func (x *Range) Reset() {
	*x = Range{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_friends_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Range) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Range) ProtoMessage() {}

func (x *Range) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_friends_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Range.ProtoReflect.Descriptor instead.
func (*Range) Descriptor() ([]byte, []int) {
	return file_protoc_friends_proto_rawDescGZIP(), []int{3}
}

func (x *Range) GetRange() int64 {
	if x != nil {
		return x.Range
	}
	return 0
}

var File_protoc_friends_proto protoreflect.FileDescriptor

var file_protoc_friends_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x18, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b,
	0x22, 0x18, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1c, 0x0a, 0x0a, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1d, 0x0a, 0x05, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x32, 0xb6, 0x01, 0x0a, 0x07, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x73, 0x12, 0x25, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x07, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x0b, 0x2e,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x41, 0x64,
	0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x12, 0x07, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x1a, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x0d, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x12, 0x07, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x1a, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0c,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x12, 0x07, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20,
	0x0a, 0x0d, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x6f, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x12,
	0x07, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x06, 0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protoc_friends_proto_rawDescOnce sync.Once
	file_protoc_friends_proto_rawDescData = file_protoc_friends_proto_rawDesc
)

func file_protoc_friends_proto_rawDescGZIP() []byte {
	file_protoc_friends_proto_rawDescOnce.Do(func() {
		file_protoc_friends_proto_rawDescData = protoimpl.X.CompressGZIP(file_protoc_friends_proto_rawDescData)
	})
	return file_protoc_friends_proto_rawDescData
}

var file_protoc_friends_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protoc_friends_proto_goTypes = []interface{}{
	(*Status)(nil),     // 0: Status
	(*UserID)(nil),     // 1: UserID
	(*FriendList)(nil), // 2: FriendList
	(*Range)(nil),      // 3: Range
}
var file_protoc_friends_proto_depIdxs = []int32{
	1, // 0: Friends.GetFriendList:input_type -> UserID
	1, // 1: Friends.AddFriend:input_type -> UserID
	1, // 2: Friends.ConfirmFriend:input_type -> UserID
	1, // 3: Friends.DeleteFriend:input_type -> UserID
	1, // 4: Friends.RangeToFriend:input_type -> UserID
	2, // 5: Friends.GetFriendList:output_type -> FriendList
	0, // 6: Friends.AddFriend:output_type -> Status
	0, // 7: Friends.ConfirmFriend:output_type -> Status
	0, // 8: Friends.DeleteFriend:output_type -> Status
	3, // 9: Friends.RangeToFriend:output_type -> Range
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protoc_friends_proto_init() }
func file_protoc_friends_proto_init() {
	if File_protoc_friends_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protoc_friends_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_protoc_friends_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserID); i {
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
		file_protoc_friends_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendList); i {
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
		file_protoc_friends_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Range); i {
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
			RawDescriptor: file_protoc_friends_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protoc_friends_proto_goTypes,
		DependencyIndexes: file_protoc_friends_proto_depIdxs,
		MessageInfos:      file_protoc_friends_proto_msgTypes,
	}.Build()
	File_protoc_friends_proto = out.File
	file_protoc_friends_proto_rawDesc = nil
	file_protoc_friends_proto_goTypes = nil
	file_protoc_friends_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FriendsClient is the client API for Friends service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FriendsClient interface {
	GetFriendList(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*FriendList, error)
	//Current user id in ctx
	AddFriend(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Status, error)
	ConfirmFriend(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Status, error)
	DeleteFriend(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Status, error)
	RangeToFriend(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Range, error)
}

type friendsClient struct {
	cc grpc.ClientConnInterface
}

func NewFriendsClient(cc grpc.ClientConnInterface) FriendsClient {
	return &friendsClient{cc}
}

func (c *friendsClient) GetFriendList(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*FriendList, error) {
	out := new(FriendList)
	err := c.cc.Invoke(ctx, "/Friends/GetFriendList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendsClient) AddFriend(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/Friends/AddFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendsClient) ConfirmFriend(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/Friends/ConfirmFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendsClient) DeleteFriend(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/Friends/DeleteFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendsClient) RangeToFriend(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Range, error) {
	out := new(Range)
	err := c.cc.Invoke(ctx, "/Friends/RangeToFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FriendsServer is the server API for Friends service.
type FriendsServer interface {
	GetFriendList(context.Context, *UserID) (*FriendList, error)
	//Current user id in ctx
	AddFriend(context.Context, *UserID) (*Status, error)
	ConfirmFriend(context.Context, *UserID) (*Status, error)
	DeleteFriend(context.Context, *UserID) (*Status, error)
	RangeToFriend(context.Context, *UserID) (*Range, error)
}

// UnimplementedFriendsServer can be embedded to have forward compatible implementations.
type UnimplementedFriendsServer struct {
}

func (*UnimplementedFriendsServer) GetFriendList(context.Context, *UserID) (*FriendList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriendList not implemented")
}
func (*UnimplementedFriendsServer) AddFriend(context.Context, *UserID) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFriend not implemented")
}
func (*UnimplementedFriendsServer) ConfirmFriend(context.Context, *UserID) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmFriend not implemented")
}
func (*UnimplementedFriendsServer) DeleteFriend(context.Context, *UserID) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFriend not implemented")
}
func (*UnimplementedFriendsServer) RangeToFriend(context.Context, *UserID) (*Range, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RangeToFriend not implemented")
}

func RegisterFriendsServer(s *grpc.Server, srv FriendsServer) {
	s.RegisterService(&_Friends_serviceDesc, srv)
}

func _Friends_GetFriendList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendsServer).GetFriendList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Friends/GetFriendList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendsServer).GetFriendList(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friends_AddFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendsServer).AddFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Friends/AddFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendsServer).AddFriend(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friends_ConfirmFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendsServer).ConfirmFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Friends/ConfirmFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendsServer).ConfirmFriend(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friends_DeleteFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendsServer).DeleteFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Friends/DeleteFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendsServer).DeleteFriend(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friends_RangeToFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendsServer).RangeToFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Friends/RangeToFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendsServer).RangeToFriend(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

var _Friends_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Friends",
	HandlerType: (*FriendsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFriendList",
			Handler:    _Friends_GetFriendList_Handler,
		},
		{
			MethodName: "AddFriend",
			Handler:    _Friends_AddFriend_Handler,
		},
		{
			MethodName: "ConfirmFriend",
			Handler:    _Friends_ConfirmFriend_Handler,
		},
		{
			MethodName: "DeleteFriend",
			Handler:    _Friends_DeleteFriend_Handler,
		},
		{
			MethodName: "RangeToFriend",
			Handler:    _Friends_RangeToFriend_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protoc/friends.proto",
}
