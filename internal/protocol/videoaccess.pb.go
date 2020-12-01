// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: videoaccess.proto

package videoaccess

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

// 用戶传入的vid 列表
type UserVidListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid        string  `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	DeviceInfo string  `protobuf:"bytes,2,opt,name=device_info,json=deviceInfo,proto3" json:"device_info,omitempty"`
	VideoId    []int32 `protobuf:"varint,3,rep,packed,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
}

func (x *UserVidListRequest) Reset() {
	*x = UserVidListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_videoaccess_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserVidListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserVidListRequest) ProtoMessage() {}

func (x *UserVidListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_videoaccess_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserVidListRequest.ProtoReflect.Descriptor instead.
func (*UserVidListRequest) Descriptor() ([]byte, []int) {
	return file_videoaccess_proto_rawDescGZIP(), []int{0}
}

func (x *UserVidListRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *UserVidListRequest) GetDeviceInfo() string {
	if x != nil {
		return x.DeviceInfo
	}
	return ""
}

func (x *UserVidListRequest) GetVideoId() []int32 {
	if x != nil {
		return x.VideoId
	}
	return nil
}

// 返回有效的vid 列表-
type UserVideoListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid     string  `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	VideoId []int32 `protobuf:"varint,2,rep,packed,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
}

func (x *UserVideoListResponse) Reset() {
	*x = UserVideoListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_videoaccess_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserVideoListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserVideoListResponse) ProtoMessage() {}

func (x *UserVideoListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_videoaccess_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserVideoListResponse.ProtoReflect.Descriptor instead.
func (*UserVideoListResponse) Descriptor() ([]byte, []int) {
	return file_videoaccess_proto_rawDescGZIP(), []int{1}
}

func (x *UserVideoListResponse) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *UserVideoListResponse) GetVideoId() []int32 {
	if x != nil {
		return x.VideoId
	}
	return nil
}

// 返回操作结果
type OperResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetCode int32  `protobuf:"varint,1,opt,name=ret_code,json=retCode,proto3" json:"ret_code,omitempty"`
	RetMsg  string `protobuf:"bytes,2,opt,name=ret_msg,json=retMsg,proto3" json:"ret_msg,omitempty"`
}

func (x *OperResponse) Reset() {
	*x = OperResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_videoaccess_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperResponse) ProtoMessage() {}

func (x *OperResponse) ProtoReflect() protoreflect.Message {
	mi := &file_videoaccess_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperResponse.ProtoReflect.Descriptor instead.
func (*OperResponse) Descriptor() ([]byte, []int) {
	return file_videoaccess_proto_rawDescGZIP(), []int{2}
}

func (x *OperResponse) GetRetCode() int32 {
	if x != nil {
		return x.RetCode
	}
	return 0
}

func (x *OperResponse) GetRetMsg() string {
	if x != nil {
		return x.RetMsg
	}
	return ""
}

var File_videoaccess_proto protoreflect.FileDescriptor

var file_videoaccess_proto_rawDesc = []byte{
	0x0a, 0x11, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x22, 0x62, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x56, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x49, 0x64, 0x22, 0x44, 0x0a, 0x15, 0x55, 0x73, 0x65, 0x72, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x05, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x0c, 0x4f, 0x70,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x65,
	0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65,
	0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x5f, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x32, 0xb6,
	0x01, 0x0a, 0x0b, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x59,
	0x0a, 0x12, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x12, 0x1f, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x56, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0e, 0x73, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1f, 0x2e, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x56, 0x69,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x12, 0x5a, 0x0d, 0x2e, 0x3b, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x80, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_videoaccess_proto_rawDescOnce sync.Once
	file_videoaccess_proto_rawDescData = file_videoaccess_proto_rawDesc
)

func file_videoaccess_proto_rawDescGZIP() []byte {
	file_videoaccess_proto_rawDescOnce.Do(func() {
		file_videoaccess_proto_rawDescData = protoimpl.X.CompressGZIP(file_videoaccess_proto_rawDescData)
	})
	return file_videoaccess_proto_rawDescData
}

var file_videoaccess_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_videoaccess_proto_goTypes = []interface{}{
	(*UserVidListRequest)(nil),    // 0: videoaccess.UserVidListRequest
	(*UserVideoListResponse)(nil), // 1: videoaccess.UserVideoListResponse
	(*OperResponse)(nil),          // 2: videoaccess.OperResponse
}
var file_videoaccess_proto_depIdxs = []int32{
	0, // 0: videoaccess.VideoAccess.removalUserHistory:input_type -> videoaccess.UserVidListRequest
	0, // 1: videoaccess.VideoAccess.setUserHistory:input_type -> videoaccess.UserVidListRequest
	1, // 2: videoaccess.VideoAccess.removalUserHistory:output_type -> videoaccess.UserVideoListResponse
	2, // 3: videoaccess.VideoAccess.setUserHistory:output_type -> videoaccess.OperResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_videoaccess_proto_init() }
func file_videoaccess_proto_init() {
	if File_videoaccess_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_videoaccess_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserVidListRequest); i {
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
		file_videoaccess_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserVideoListResponse); i {
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
		file_videoaccess_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperResponse); i {
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
			RawDescriptor: file_videoaccess_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_videoaccess_proto_goTypes,
		DependencyIndexes: file_videoaccess_proto_depIdxs,
		MessageInfos:      file_videoaccess_proto_msgTypes,
	}.Build()
	File_videoaccess_proto = out.File
	file_videoaccess_proto_rawDesc = nil
	file_videoaccess_proto_goTypes = nil
	file_videoaccess_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// VideoAccessClient is the client API for VideoAccess service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VideoAccessClient interface {
	// 用户历史去重
	RemovalUserHistory(ctx context.Context, in *UserVidListRequest, opts ...grpc.CallOption) (*UserVideoListResponse, error)
	// 设置用户已曝光的视频vid
	SetUserHistory(ctx context.Context, in *UserVidListRequest, opts ...grpc.CallOption) (*OperResponse, error)
}

type videoAccessClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoAccessClient(cc grpc.ClientConnInterface) VideoAccessClient {
	return &videoAccessClient{cc}
}

func (c *videoAccessClient) RemovalUserHistory(ctx context.Context, in *UserVidListRequest, opts ...grpc.CallOption) (*UserVideoListResponse, error) {
	out := new(UserVideoListResponse)
	err := c.cc.Invoke(ctx, "/videoaccess.VideoAccess/removalUserHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoAccessClient) SetUserHistory(ctx context.Context, in *UserVidListRequest, opts ...grpc.CallOption) (*OperResponse, error) {
	out := new(OperResponse)
	err := c.cc.Invoke(ctx, "/videoaccess.VideoAccess/setUserHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoAccessServer is the server API for VideoAccess service.
type VideoAccessServer interface {
	// 用户历史去重
	RemovalUserHistory(context.Context, *UserVidListRequest) (*UserVideoListResponse, error)
	// 设置用户已曝光的视频vid
	SetUserHistory(context.Context, *UserVidListRequest) (*OperResponse, error)
}

// UnimplementedVideoAccessServer can be embedded to have forward compatible implementations.
type UnimplementedVideoAccessServer struct {
}

func (*UnimplementedVideoAccessServer) RemovalUserHistory(context.Context, *UserVidListRequest) (*UserVideoListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovalUserHistory not implemented")
}
func (*UnimplementedVideoAccessServer) SetUserHistory(context.Context, *UserVidListRequest) (*OperResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUserHistory not implemented")
}

func RegisterVideoAccessServer(s *grpc.Server, srv VideoAccessServer) {
	s.RegisterService(&_VideoAccess_serviceDesc, srv)
}

func _VideoAccess_RemovalUserHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserVidListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoAccessServer).RemovalUserHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/videoaccess.VideoAccess/RemovalUserHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoAccessServer).RemovalUserHistory(ctx, req.(*UserVidListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoAccess_SetUserHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserVidListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoAccessServer).SetUserHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/videoaccess.VideoAccess/SetUserHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoAccessServer).SetUserHistory(ctx, req.(*UserVidListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VideoAccess_serviceDesc = grpc.ServiceDesc{
	ServiceName: "videoaccess.VideoAccess",
	HandlerType: (*VideoAccessServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "removalUserHistory",
			Handler:    _VideoAccess_RemovalUserHistory_Handler,
		},
		{
			MethodName: "setUserHistory",
			Handler:    _VideoAccess_SetUserHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "videoaccess.proto",
}
