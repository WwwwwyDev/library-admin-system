package search

import (
	"context"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/service/book/rpc/bookclient"

	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAllBookTypesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllBookTypesLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAllBookTypesLogic {
	return GetAllBookTypesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllBookTypesLogic) GetAllBookTypes() (*types.Reply, error) {
	allTypes, err := l.svcCtx.BookRpc.GetAllTypes(l.ctx, &bookclient.EmptyReq{})
	if err != nil {
		return nil, err
	}
	return &types.Reply{Code: code.Success,Data: map[string]interface{}{"types":allTypes.TypesInfo},Msg: "查询成功"}, nil
}
