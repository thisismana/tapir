// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.3
// source: stroeer/web/article/v1/web_article_service.proto

package article

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

type GetWebArticlePageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetWebArticlePageRequest) Reset() {
	*x = GetWebArticlePageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stroeer_web_article_v1_web_article_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWebArticlePageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWebArticlePageRequest) ProtoMessage() {}

func (x *GetWebArticlePageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stroeer_web_article_v1_web_article_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWebArticlePageRequest.ProtoReflect.Descriptor instead.
func (*GetWebArticlePageRequest) Descriptor() ([]byte, []int) {
	return file_stroeer_web_article_v1_web_article_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetWebArticlePageRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetWebArticlePageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WebArticle         *WebArticle          `protobuf:"bytes,1,opt,name=web_article,json=webArticle,proto3" json:"web_article,omitempty"`
	RelatedWebArticles []*RelatedWebArticle `protobuf:"bytes,2,rep,name=related_web_articles,json=relatedWebArticles,proto3" json:"related_web_articles,omitempty"`
}

func (x *GetWebArticlePageResponse) Reset() {
	*x = GetWebArticlePageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stroeer_web_article_v1_web_article_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWebArticlePageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWebArticlePageResponse) ProtoMessage() {}

func (x *GetWebArticlePageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stroeer_web_article_v1_web_article_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWebArticlePageResponse.ProtoReflect.Descriptor instead.
func (*GetWebArticlePageResponse) Descriptor() ([]byte, []int) {
	return file_stroeer_web_article_v1_web_article_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetWebArticlePageResponse) GetWebArticle() *WebArticle {
	if x != nil {
		return x.WebArticle
	}
	return nil
}

func (x *GetWebArticlePageResponse) GetRelatedWebArticles() []*RelatedWebArticle {
	if x != nil {
		return x.RelatedWebArticles
	}
	return nil
}

var File_stroeer_web_article_v1_web_article_service_proto protoreflect.FileDescriptor

var file_stroeer_web_article_v1_web_article_service_proto_rawDesc = []byte{
	0x0a, 0x30, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x77, 0x65, 0x62, 0x2f, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x65, 0x62, 0x5f, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x16, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x77, 0x65, 0x62, 0x2e,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x28, 0x73, 0x74, 0x72, 0x6f,
	0x65, 0x65, 0x72, 0x2f, 0x77, 0x65, 0x62, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x77, 0x65, 0x62, 0x5f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x57, 0x65, 0x62, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0xbd, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x57, 0x65, 0x62, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43,
	0x0a, 0x0b, 0x77, 0x65, 0x62, 0x5f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x77, 0x65,
	0x62, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x62,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x0a, 0x77, 0x65, 0x62, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x12, 0x5b, 0x0a, 0x14, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x77,
	0x65, 0x62, 0x5f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x29, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x77, 0x65, 0x62, 0x2e,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x65, 0x64, 0x57, 0x65, 0x62, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x12, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x65, 0x64, 0x57, 0x65, 0x62, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73,
	0x32, 0x8f, 0x01, 0x0a, 0x11, 0x57, 0x65, 0x62, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7a, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x57, 0x65, 0x62,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x50, 0x61, 0x67, 0x65, 0x12, 0x30, 0x2e, 0x73, 0x74,
	0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x65, 0x62, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e,
	0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x65, 0x62, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x59, 0x0a, 0x19, 0x64, 0x65, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72,
	0x2e, 0x77, 0x65, 0x62, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x50,
	0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74,
	0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x70, 0x69, 0x72, 0x2f, 0x67, 0x6f, 0x2f, 0x73,
	0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x77, 0x65, 0x62, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stroeer_web_article_v1_web_article_service_proto_rawDescOnce sync.Once
	file_stroeer_web_article_v1_web_article_service_proto_rawDescData = file_stroeer_web_article_v1_web_article_service_proto_rawDesc
)

func file_stroeer_web_article_v1_web_article_service_proto_rawDescGZIP() []byte {
	file_stroeer_web_article_v1_web_article_service_proto_rawDescOnce.Do(func() {
		file_stroeer_web_article_v1_web_article_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_stroeer_web_article_v1_web_article_service_proto_rawDescData)
	})
	return file_stroeer_web_article_v1_web_article_service_proto_rawDescData
}

var file_stroeer_web_article_v1_web_article_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_stroeer_web_article_v1_web_article_service_proto_goTypes = []interface{}{
	(*GetWebArticlePageRequest)(nil),  // 0: stroeer.web.article.v1.GetWebArticlePageRequest
	(*GetWebArticlePageResponse)(nil), // 1: stroeer.web.article.v1.GetWebArticlePageResponse
	(*WebArticle)(nil),                // 2: stroeer.web.article.v1.WebArticle
	(*RelatedWebArticle)(nil),         // 3: stroeer.web.article.v1.RelatedWebArticle
}
var file_stroeer_web_article_v1_web_article_service_proto_depIdxs = []int32{
	2, // 0: stroeer.web.article.v1.GetWebArticlePageResponse.web_article:type_name -> stroeer.web.article.v1.WebArticle
	3, // 1: stroeer.web.article.v1.GetWebArticlePageResponse.related_web_articles:type_name -> stroeer.web.article.v1.RelatedWebArticle
	0, // 2: stroeer.web.article.v1.WebArticleService.GetWebArticlePage:input_type -> stroeer.web.article.v1.GetWebArticlePageRequest
	1, // 3: stroeer.web.article.v1.WebArticleService.GetWebArticlePage:output_type -> stroeer.web.article.v1.GetWebArticlePageResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_stroeer_web_article_v1_web_article_service_proto_init() }
func file_stroeer_web_article_v1_web_article_service_proto_init() {
	if File_stroeer_web_article_v1_web_article_service_proto != nil {
		return
	}
	file_stroeer_web_article_v1_web_article_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_stroeer_web_article_v1_web_article_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWebArticlePageRequest); i {
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
		file_stroeer_web_article_v1_web_article_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWebArticlePageResponse); i {
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
			RawDescriptor: file_stroeer_web_article_v1_web_article_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stroeer_web_article_v1_web_article_service_proto_goTypes,
		DependencyIndexes: file_stroeer_web_article_v1_web_article_service_proto_depIdxs,
		MessageInfos:      file_stroeer_web_article_v1_web_article_service_proto_msgTypes,
	}.Build()
	File_stroeer_web_article_v1_web_article_service_proto = out.File
	file_stroeer_web_article_v1_web_article_service_proto_rawDesc = nil
	file_stroeer_web_article_v1_web_article_service_proto_goTypes = nil
	file_stroeer_web_article_v1_web_article_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WebArticleServiceClient is the client API for WebArticleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WebArticleServiceClient interface {
	// Returns an Article page including all (SEO) relevant information to
	// render an Article with the given id as canonical web or AMP
	GetWebArticlePage(ctx context.Context, in *GetWebArticlePageRequest, opts ...grpc.CallOption) (*GetWebArticlePageResponse, error)
}

type webArticleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWebArticleServiceClient(cc grpc.ClientConnInterface) WebArticleServiceClient {
	return &webArticleServiceClient{cc}
}

func (c *webArticleServiceClient) GetWebArticlePage(ctx context.Context, in *GetWebArticlePageRequest, opts ...grpc.CallOption) (*GetWebArticlePageResponse, error) {
	out := new(GetWebArticlePageResponse)
	err := c.cc.Invoke(ctx, "/stroeer.web.article.v1.WebArticleService/GetWebArticlePage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebArticleServiceServer is the server API for WebArticleService service.
type WebArticleServiceServer interface {
	// Returns an Article page including all (SEO) relevant information to
	// render an Article with the given id as canonical web or AMP
	GetWebArticlePage(context.Context, *GetWebArticlePageRequest) (*GetWebArticlePageResponse, error)
}

// UnimplementedWebArticleServiceServer can be embedded to have forward compatible implementations.
type UnimplementedWebArticleServiceServer struct {
}

func (*UnimplementedWebArticleServiceServer) GetWebArticlePage(context.Context, *GetWebArticlePageRequest) (*GetWebArticlePageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWebArticlePage not implemented")
}

func RegisterWebArticleServiceServer(s *grpc.Server, srv WebArticleServiceServer) {
	s.RegisterService(&_WebArticleService_serviceDesc, srv)
}

func _WebArticleService_GetWebArticlePage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWebArticlePageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebArticleServiceServer).GetWebArticlePage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stroeer.web.article.v1.WebArticleService/GetWebArticlePage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebArticleServiceServer).GetWebArticlePage(ctx, req.(*GetWebArticlePageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WebArticleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stroeer.web.article.v1.WebArticleService",
	HandlerType: (*WebArticleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWebArticlePage",
			Handler:    _WebArticleService_GetWebArticlePage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stroeer/web/article/v1/web_article_service.proto",
}