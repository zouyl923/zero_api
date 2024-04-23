// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: article.proto

package rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CommonService_PageList_FullMethodName     = "/article.CommonService/PageList"
	CommonService_CategoryList_FullMethodName = "/article.CommonService/CategoryList"
	CommonService_Info_FullMethodName         = "/article.CommonService/Info"
)

// CommonServiceClient is the client API for CommonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommonServiceClient interface {
	// 公共文章
	PageList(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*PageData, error)
	CategoryList(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*CategoryRes, error)
	Info(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*Article, error)
}

type commonServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommonServiceClient(cc grpc.ClientConnInterface) CommonServiceClient {
	return &commonServiceClient{cc}
}

func (c *commonServiceClient) PageList(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*PageData, error) {
	out := new(PageData)
	err := c.cc.Invoke(ctx, CommonService_PageList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonServiceClient) CategoryList(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*CategoryRes, error) {
	out := new(CategoryRes)
	err := c.cc.Invoke(ctx, CommonService_CategoryList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonServiceClient) Info(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*Article, error) {
	out := new(Article)
	err := c.cc.Invoke(ctx, CommonService_Info_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommonServiceServer is the server API for CommonService service.
// All implementations must embed UnimplementedCommonServiceServer
// for forward compatibility
type CommonServiceServer interface {
	// 公共文章
	PageList(context.Context, *SearchReq) (*PageData, error)
	CategoryList(context.Context, *EmptyReq) (*CategoryRes, error)
	Info(context.Context, *InfoReq) (*Article, error)
	mustEmbedUnimplementedCommonServiceServer()
}

// UnimplementedCommonServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommonServiceServer struct {
}

func (UnimplementedCommonServiceServer) PageList(context.Context, *SearchReq) (*PageData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PageList not implemented")
}
func (UnimplementedCommonServiceServer) CategoryList(context.Context, *EmptyReq) (*CategoryRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CategoryList not implemented")
}
func (UnimplementedCommonServiceServer) Info(context.Context, *InfoReq) (*Article, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedCommonServiceServer) mustEmbedUnimplementedCommonServiceServer() {}

// UnsafeCommonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommonServiceServer will
// result in compilation errors.
type UnsafeCommonServiceServer interface {
	mustEmbedUnimplementedCommonServiceServer()
}

func RegisterCommonServiceServer(s grpc.ServiceRegistrar, srv CommonServiceServer) {
	s.RegisterService(&CommonService_ServiceDesc, srv)
}

func _CommonService_PageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServiceServer).PageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommonService_PageList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServiceServer).PageList(ctx, req.(*SearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommonService_CategoryList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServiceServer).CategoryList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommonService_CategoryList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServiceServer).CategoryList(ctx, req.(*EmptyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommonService_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServiceServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommonService_Info_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServiceServer).Info(ctx, req.(*InfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// CommonService_ServiceDesc is the grpc.ServiceDesc for CommonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "article.CommonService",
	HandlerType: (*CommonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PageList",
			Handler:    _CommonService_PageList_Handler,
		},
		{
			MethodName: "CategoryList",
			Handler:    _CommonService_CategoryList_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _CommonService_Info_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "article.proto",
}

const (
	UserService_PageList_FullMethodName = "/article.UserService/PageList"
	UserService_Info_FullMethodName     = "/article.UserService/Info"
	UserService_Push_FullMethodName     = "/article.UserService/Push"
	UserService_Upload_FullMethodName   = "/article.UserService/Upload"
	UserService_Delete_FullMethodName   = "/article.UserService/Delete"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	// 个人文章
	PageList(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*PageData, error)
	Info(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*Article, error)
	Push(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*EmptyRes, error)
	Upload(ctx context.Context, in *UploadReq, opts ...grpc.CallOption) (*UploadRes, error)
	Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*EmptyRes, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) PageList(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*PageData, error) {
	out := new(PageData)
	err := c.cc.Invoke(ctx, UserService_PageList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Info(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*Article, error) {
	out := new(Article)
	err := c.cc.Invoke(ctx, UserService_Info_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Push(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*EmptyRes, error) {
	out := new(EmptyRes)
	err := c.cc.Invoke(ctx, UserService_Push_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Upload(ctx context.Context, in *UploadReq, opts ...grpc.CallOption) (*UploadRes, error) {
	out := new(UploadRes)
	err := c.cc.Invoke(ctx, UserService_Upload_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*EmptyRes, error) {
	out := new(EmptyRes)
	err := c.cc.Invoke(ctx, UserService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	// 个人文章
	PageList(context.Context, *SearchReq) (*PageData, error)
	Info(context.Context, *InfoReq) (*Article, error)
	Push(context.Context, *UpdateReq) (*EmptyRes, error)
	Upload(context.Context, *UploadReq) (*UploadRes, error)
	Delete(context.Context, *DeleteReq) (*EmptyRes, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) PageList(context.Context, *SearchReq) (*PageData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PageList not implemented")
}
func (UnimplementedUserServiceServer) Info(context.Context, *InfoReq) (*Article, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedUserServiceServer) Push(context.Context, *UpdateReq) (*EmptyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}
func (UnimplementedUserServiceServer) Upload(context.Context, *UploadReq) (*UploadRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedUserServiceServer) Delete(context.Context, *DeleteReq) (*EmptyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_PageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).PageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_PageList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).PageList(ctx, req.(*SearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Info_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Info(ctx, req.(*InfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Push_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Push(ctx, req.(*UpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Upload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Upload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Upload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Upload(ctx, req.(*UploadReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Delete(ctx, req.(*DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "article.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PageList",
			Handler:    _UserService_PageList_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _UserService_Info_Handler,
		},
		{
			MethodName: "Push",
			Handler:    _UserService_Push_Handler,
		},
		{
			MethodName: "Upload",
			Handler:    _UserService_Upload_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "article.proto",
}

const (
	AdminService_PageList_FullMethodName = "/article.AdminService/PageList"
	AdminService_Delete_FullMethodName   = "/article.AdminService/Delete"
	AdminService_Examine_FullMethodName  = "/article.AdminService/Examine"
)

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminServiceClient interface {
	// 管理后台
	PageList(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*PageData, error)
	Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*EmptyRes, error)
	Examine(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*EmptyRes, error)
}

type adminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminServiceClient(cc grpc.ClientConnInterface) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) PageList(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*PageData, error) {
	out := new(PageData)
	err := c.cc.Invoke(ctx, AdminService_PageList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*EmptyRes, error) {
	out := new(EmptyRes)
	err := c.cc.Invoke(ctx, AdminService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) Examine(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*EmptyRes, error) {
	out := new(EmptyRes)
	err := c.cc.Invoke(ctx, AdminService_Examine_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
// All implementations must embed UnimplementedAdminServiceServer
// for forward compatibility
type AdminServiceServer interface {
	// 管理后台
	PageList(context.Context, *SearchReq) (*PageData, error)
	Delete(context.Context, *DeleteReq) (*EmptyRes, error)
	Examine(context.Context, *DeleteReq) (*EmptyRes, error)
	mustEmbedUnimplementedAdminServiceServer()
}

// UnimplementedAdminServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServiceServer struct {
}

func (UnimplementedAdminServiceServer) PageList(context.Context, *SearchReq) (*PageData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PageList not implemented")
}
func (UnimplementedAdminServiceServer) Delete(context.Context, *DeleteReq) (*EmptyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAdminServiceServer) Examine(context.Context, *DeleteReq) (*EmptyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Examine not implemented")
}
func (UnimplementedAdminServiceServer) mustEmbedUnimplementedAdminServiceServer() {}

// UnsafeAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServiceServer will
// result in compilation errors.
type UnsafeAdminServiceServer interface {
	mustEmbedUnimplementedAdminServiceServer()
}

func RegisterAdminServiceServer(s grpc.ServiceRegistrar, srv AdminServiceServer) {
	s.RegisterService(&AdminService_ServiceDesc, srv)
}

func _AdminService_PageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).PageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_PageList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).PageList(ctx, req.(*SearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).Delete(ctx, req.(*DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_Examine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).Examine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_Examine_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).Examine(ctx, req.(*DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminService_ServiceDesc is the grpc.ServiceDesc for AdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "article.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PageList",
			Handler:    _AdminService_PageList_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AdminService_Delete_Handler,
		},
		{
			MethodName: "Examine",
			Handler:    _AdminService_Examine_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "article.proto",
}
