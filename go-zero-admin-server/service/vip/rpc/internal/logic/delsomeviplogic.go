package logic

import (
	"context"
	"go-zero-admin-server/service/vip/rpc/internal/svc"
	"go-zero-admin-server/service/vip/rpc/vip"

	"github.com/tal-tech/go-zero/core/logx"
)

type DelSomeVipLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelSomeVipLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelSomeVipLogic {
	return &DelSomeVipLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelSomeVipLogic) DelSomeVip(in *vip.IdsReq) (*vip.IsSuccessReply, error) {
	ids := make([]uint, 0)
	for _,e := range in.Ids{
		ids = append(ids,uint(e.Id))
	}
	isSuccess, err := l.svcCtx.VipModel.DelSomeVip(ids)
	if err != nil {
		return nil, err
	}
	return &vip.IsSuccessReply{IsSuccess: isSuccess}, nil
}
