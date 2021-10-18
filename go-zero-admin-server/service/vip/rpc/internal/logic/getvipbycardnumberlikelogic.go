package logic

import (
	"context"
	"go-zero-admin-server/service/vip/rpc/internal/svc"
	"go-zero-admin-server/service/vip/rpc/vip"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetVipByCardNumberLikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVipByCardNumberLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVipByCardNumberLikeLogic {
	return &GetVipByCardNumberLikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVipByCardNumberLikeLogic) GetVipByCardNumberLike(in *vip.CardNumberReq) (*vip.VipsInfoReply, error) {
	vips, err := l.svcCtx.VipModel.GetVipByCardNumberLike(in.CardNumber)
	if err != nil {
		return nil, err
	}
	vipsHandle := make([]*vip.VipInfoReply, 0)
	for _, e := range vips {
		vipsHandle = append(vipsHandle, &vip.VipInfoReply{Id: uint64(e.ID), CardNumber: e.CardNumber,Name: e.Name,Info: e.Info})
	}
	return &vip.VipsInfoReply{VipsInfo: vipsHandle}, nil
}
