package search

import (
	"context"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/service/book/rpc/bookclient"

	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetTypeByNameLikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTypeByNameLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetTypeByNameLikeLogic {
	return GetTypeByNameLikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTypeByNameLikeLogic) GetTypeByNameLike(req types.NameReq) (*types.Reply, error) {
	_types, err := l.svcCtx.BookRpc.GetTypeByNameLike(l.ctx, &bookclient.NameReq{Name: req.Name})
	if err != nil {
		return nil, err
	}
	return &types.Reply{Code: code.Success,Data: map[string]interface{}{"types":_types},Msg: "获取成功"}, nil
}
