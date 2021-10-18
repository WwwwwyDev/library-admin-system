package logic

import (
	"context"
	"go-zero-admin-server/service/vip/model"
	"gorm.io/gorm"

	"go-zero-admin-server/service/vip/rpc/internal/svc"
	"go-zero-admin-server/service/vip/rpc/vip"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateVipLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateVipLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateVipLogic {
	return &UpdateVipLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateVipLogic) UpdateVip(in *vip.VipUpdateReq) (*vip.IsSuccessReply, error) {
	isSuccess, err := l.svcCtx.VipModel.UpdateVip(&model.Vip{Model: gorm.Model{ID: uint(in.Id)}, CardNumber: in.CardNumber, Name: in.Name, Info: in.Info})
	if err != nil {
		return nil, err
	}
	return &vip.IsSuccessReply{IsSuccess: isSuccess}, nil
}
