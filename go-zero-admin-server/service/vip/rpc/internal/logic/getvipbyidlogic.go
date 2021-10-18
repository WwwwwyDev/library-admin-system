package logic

import (
	"context"

	"go-zero-admin-server/service/vip/rpc/internal/svc"
	"go-zero-admin-server/service/vip/rpc/vip"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetVipByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVipByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVipByIdLogic {
	return &GetVipByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVipByIdLogic) GetVipById(in *vip.IdReq) (*vip.VipInfoReply, error) {
	_vip, err := l.svcCtx.VipModel.GetVipById(uint(in.Id))
	if err != nil {
		return nil, err
	}
	return &vip.VipInfoReply{Id: uint64(_vip.ID),CardNumber: _vip.CardNumber,Name: _vip.Name,Info: _vip.Info}, nil
}
