package logic

import (
	"context"

	"go-zero-admin-server/service/vip/rpc/internal/svc"
	"go-zero-admin-server/service/vip/rpc/vip"

	"github.com/tal-tech/go-zero/core/logx"
)

type IsExistVipByCardNumberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsExistVipByCardNumberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsExistVipByCardNumberLogic {
	return &IsExistVipByCardNumberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsExistVipByCardNumberLogic) IsExistVipByCardNumber(in *vip.CardNumberReq) (*vip.IsExistReply, error) {
	isExist, err := l.svcCtx.VipModel.IsExistVipByCardNumber(in.CardNumber)
	if err != nil {
		return nil, err
	}
	return &vip.IsExistReply{IsExist: isExist}, nil
}
