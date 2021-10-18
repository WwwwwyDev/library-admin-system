package search

import (
	"context"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/service/book/rpc/bookclient"

	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetTypeByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTypeByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetTypeByIdLogic {
	return GetTypeByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTypeByIdLogic) GetTypeById(req types.IdReq) (*types.Reply, error) {
	_type, err := l.svcCtx.BookRpc.GetTypeById(l.ctx, &bookclient.IdReq{Id: uint64(req.Id)})
	if err != nil {
		return nil, err
	}
	return &types.Reply{Code: code.Success,Data: map[string]interface{}{"type": _type},Msg: "获取成功"}, nil
}
