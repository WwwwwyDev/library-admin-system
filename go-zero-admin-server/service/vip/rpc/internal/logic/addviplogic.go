package logic

import (
	"context"
	"go-zero-admin-server/service/vip/model"

	"go-zero-admin-server/service/vip/rpc/internal/svc"
	"go-zero-admin-server/service/vip/rpc/vip"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddVipLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddVipLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddVipLogic {
	return &AddVipLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddVipLogic) AddVip(in *vip.VipAddReq) (*vip.IsSuccessReply, error) {
	isSuccess, err := l.svcCtx.VipModel.AddVip(&model.Vip{CardNumber: in.CardNumber, Name: in.Name, Info: in.Info})
	if err != nil {
		return nil, err
	}
	return &vip.IsSuccessReply{IsSuccess: isSuccess}, nil
}
