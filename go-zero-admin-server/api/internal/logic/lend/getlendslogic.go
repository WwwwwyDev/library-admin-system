package lend

import (
	"context"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/service/lend/rpc/lendclient"

	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetLendsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLendsLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetLendsLogic {
	return GetLendsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLendsLogic) GetLends(req types.GetLendsReq) (*types.Reply, error) {
	lends, err := l.svcCtx.LendRpc.GetLends(l.ctx, &lendclient.LendsReq{Page: int64(req.Page), Limit: int64(req.Limit), BookName: req.BookName, VipCardNumber: req.VipCardNumber})
	if err != nil {
		return nil, err
	}
	return &types.Reply{Code: code.Success,Data: map[string]interface{}{"lends":lends.LendsInfo,"total":lends.Total}}, nil
}
