// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"blog/app/user/rpc/internal/logic/userservice"
	"blog/app/user/rpc/internal/svc"
	"blog/app/user/rpc/pb/rpc"
)

type UserServiceServer struct {
	svcCtx *svc.ServiceContext
	rpc.UnimplementedUserServiceServer
}

func NewUserServiceServer(svcCtx *svc.ServiceContext) *UserServiceServer {
	return &UserServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServiceServer) Register(ctx context.Context, in *rpc.RegisterReq) (*rpc.RegisterRes, error) {
	l := userservicelogic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UserServiceServer) Login(ctx context.Context, in *rpc.LoginReq) (*rpc.LoginRes, error) {
	l := userservicelogic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServiceServer) Info(ctx context.Context, in *rpc.InfoReq) (*rpc.User, error) {
	l := userservicelogic.NewInfoLogic(ctx, s.svcCtx)
	return l.Info(in)
}
