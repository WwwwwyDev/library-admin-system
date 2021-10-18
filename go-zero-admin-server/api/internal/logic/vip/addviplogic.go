package vip

import (
	"context"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/common/errorx"
	"go-zero-admin-server/service/vip/rpc/vipclient"

	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddVipLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddVipLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddVipLogic {
	return AddVipLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddVipLogic) AddVip(req types.AddVipReq) (*types.Reply, error) {
	isSuccess, err := l.svcCtx.VipRpc.AddVip(l.ctx, &vipclient.VipAddReq{CardNumber: req.CardNumber, Name: req.Name, Info: req.Info})
	if err != nil {
		return nil, err
	}
	if !isSuccess.IsSuccess{
		return nil,errorx.NewCodeError(code.AddError,"添加失败")
	}
	return &types.Reply{Code: code.Success,Msg: "添加成功"}, nil
}
