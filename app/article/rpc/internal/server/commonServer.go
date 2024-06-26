// Code generated by goctl. DO NOT EDIT.
// Source: article.proto

package server

import (
	"context"

	"blog/app/article/rpc/internal/logic"
	"blog/app/article/rpc/internal/svc"
	"blog/app/article/rpc/pb/rpc"
)

type CommonServer struct {
	svcCtx *svc.ServiceContext
	rpc.UnimplementedCommonServer
}

func NewCommonServer(svcCtx *svc.ServiceContext) *CommonServer {
	return &CommonServer{
		svcCtx: svcCtx,
	}
}

// 公共文章
func (s *CommonServer) PageList(ctx context.Context, in *rpc.SearchReq) (*rpc.PageData, error) {
	l := logic.NewPageListLogic(ctx, s.svcCtx)
	return l.PageList(in)
}

func (s *CommonServer) CategoryList(ctx context.Context, in *rpc.EmptyReq) (*rpc.CategoryRes, error) {
	l := logic.NewCategoryListLogic(ctx, s.svcCtx)
	return l.CategoryList(in)
}

func (s *CommonServer) Info(ctx context.Context, in *rpc.InfoReq) (*rpc.Article, error) {
	l := logic.NewInfoLogic(ctx, s.svcCtx)
	return l.Info(in)
}
