package lend

import (
	"context"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/common/errorx"
	"go-zero-admin-server/service/lend/rpc/lendclient"

	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteLendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLendLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteLendLogic {
	return DeleteLendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLendLogic) DeleteLend(req types.DeleteLendReq) (*types.Reply, error) {
	isExist, err := l.svcCtx.LendRpc.IsExistLendById(l.ctx, &lendclient.IdReq{Id: uint64(req.Id)})
	if err != nil {
		return nil, err
	}
	if !isExist.IsExist{
		return nil, errorx.NewCodeError(code.NoFoundError,"没有此记录")
	}
	isSuccess, err := l.svcCtx.LendRpc.DelLend(l.ctx, &lendclient.IdReq{Id: uint64(req.Id)})
	if err != nil {
		return nil, err
	}
	if !isSuccess.IsSuccess{
		return nil, errorx.NewCodeError(code.DelError,"删除失败")
	}
	return &types.Reply{Code: code.Success,Msg: "删除成功"}, nil
}
