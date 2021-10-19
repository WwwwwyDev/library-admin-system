package vip

import (
	"context"
	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/common/errorx"
	"go-zero-admin-server/service/vip/rpc/vipclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteVipLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteVipLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteVipLogic {
	return DeleteVipLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteVipLogic) DeleteVip(req types.DeleteVipReq) (*types.Reply, error) {
	isExist, err := l.svcCtx.VipRpc.IsExistVipById(l.ctx, &vipclient.IdReq{Id: uint64(req.Id)})
	if !isExist.IsExist{
		return nil, errorx.NewCodeError(code.NoFoundError,"没有此会员")
	}
	isSuccess, err := l.svcCtx.VipRpc.DelVip(l.ctx, &vipclient.IdReq{Id: uint64(req.Id)})
	if err != nil {
		return nil, err
	}
	if !isSuccess.IsSuccess{
		return nil, errorx.NewCodeError(code.DelError,"删除失败")
	}
	return &types.Reply{Code: code.Success,Msg: "删除成功"}, nil
}
