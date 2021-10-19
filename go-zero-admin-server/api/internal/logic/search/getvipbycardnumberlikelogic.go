package search

import (
	"context"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/service/vip/rpc/vipclient"

	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetVipByCardNumberLikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVipByCardNumberLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetVipByCardNumberLikeLogic {
	return GetVipByCardNumberLikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVipByCardNumberLikeLogic) GetVipByCardNumberLike(req types.CardNumberReq) (*types.Reply, error) {
	vips, err := l.svcCtx.VipRpc.GetVipByCardNumberLike(l.ctx, &vipclient.CardNumberReq{CardNumber: req.CardNumber})
	if err != nil {
		return nil, err
	}
	return &types.Reply{Code: code.Success,Data: map[string]interface{}{"vips":vips.VipsInfo},Msg: "搜索成功"}, nil
}
