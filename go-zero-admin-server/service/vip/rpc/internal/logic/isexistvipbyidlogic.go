package logic

import (
	"context"

	"go-zero-admin-server/service/vip/rpc/internal/svc"
	"go-zero-admin-server/service/vip/rpc/vip"

	"github.com/tal-tech/go-zero/core/logx"
)

type IsExistVipByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsExistVipByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsExistVipByIdLogic {
	return &IsExistVipByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsExistVipByIdLogic) IsExistVipById(in *vip.IdReq) (*vip.IsExistReply, error) {
	isExist, err := l.svcCtx.VipModel.IsExistVipById(uint(in.Id))
	if err != nil {
		return nil, err
	}
	return &vip.IsExistReply{IsExist: isExist}, nil
}
