// Code generated by goctl. DO NOT EDIT.
// Source: verify.proto

package server

import (
	"context"

	"blog/app/verify/rpc/internal/logic/verifyservice"
	"blog/app/verify/rpc/internal/svc"
	"blog/app/verify/rpc/pb/rpc"
)

type VerifyServiceServer struct {
	svcCtx *svc.ServiceContext
	rpc.UnimplementedVerifyServiceServer
}

func NewVerifyServiceServer(svcCtx *svc.ServiceContext) *VerifyServiceServer {
	return &VerifyServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *VerifyServiceServer) GenToken(ctx context.Context, in *rpc.GenTokenReq) (*rpc.GenTokenRes, error) {
	l := verifyservicelogic.NewGenTokenLogic(ctx, s.svcCtx)
	return l.GenToken(in)
}

func (s *VerifyServiceServer) RemoveToken(ctx context.Context, in *rpc.RemoveTokenReq) (*rpc.EmptyRes, error) {
	l := verifyservicelogic.NewRemoveTokenLogic(ctx, s.svcCtx)
	return l.RemoveToken(in)
}

func (s *VerifyServiceServer) Auth(ctx context.Context, in *rpc.AuthReq) (*rpc.AuthRes, error) {
	l := verifyservicelogic.NewAuthLogic(ctx, s.svcCtx)
	return l.Auth(in)
}