package search

import (
	"context"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/service/book/rpc/bookclient"

	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetBookByNameLikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBookByNameLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetBookByNameLikeLogic {
	return GetBookByNameLikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBookByNameLikeLogic) GetBookByNameLike(req types.NameReq) (*types.Reply, error) {
	books, err := l.svcCtx.BookRpc.GetBookByNameLike(l.ctx, &bookclient.NameReq{Name: req.Name})
	if err != nil {
		return nil, err
	}
	return &types.Reply{Code: code.Success,Data: map[string]interface{}{"books":books.BooksInfo},Msg: "获取成功"}, nil
}
