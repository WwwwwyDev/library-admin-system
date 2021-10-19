package vip

import (
	"context"
	"github.com/gofrs/uuid"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/common/errorx"
	"go-zero-admin-server/service/vip/rpc/vipclient"

	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateVipLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateVipLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateVipLogic {
	return UpdateVipLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateVipLogic) UpdateVip(req types.UpdateVipReq) (*types.Reply, error) {
	cardb, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	cards := cardb.String()
	isSuccess, err := l.svcCtx.VipRpc.UpdateVip(l.ctx, &vipclient.VipUpdateReq{Id: uint64(req.Id), CardNumber: cards, Name: req.Name, Info: req.Info})
	if err != nil {
		return nil, err
	}
	if !isSuccess.IsSuccess{
		return nil, errorx.NewCodeError(code.ChangeError,"更新失败")
	}
	return &types.Reply{Code: code.Success,Msg: "更新成功"}, nil
}
