package logic

import (
	"context"
	"go-zero-admin-server/service/vip/model"
	"gorm.io/gorm"

	"go-zero-admin-server/service/vip/rpc/internal/svc"
	"go-zero-admin-server/service/vip/rpc/vip"

	"github.com/tal-tech/go-zero/core/logx"
)

type DelVipLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelVipLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelVipLogic {
	return &DelVipLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelVipLogic) DelVip(in *vip.IdReq) (*vip.IsSuccessReply, error) {
	isSuccess, err := l.svcCtx.VipModel.DelVip(&model.Vip{Model: gorm.Model{ID: uint(in.Id)}})
	if err != nil {
		return nil, err
	}
	return &vip.IsSuccessReply{IsSuccess: isSuccess}, nil
}
