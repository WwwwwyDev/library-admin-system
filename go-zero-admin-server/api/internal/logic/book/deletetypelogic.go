package book

import (
	"context"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/common/errorx"
	"go-zero-admin-server/service/book/rpc/bookclient"

	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteTypeLogic {
	return DeleteTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTypeLogic) DeleteType(req types.DeleteTypeReq) (*types.Reply, error) {
	isExist, err := l.svcCtx.BookRpc.IsExistTypeById(l.ctx, &bookclient.IdReq{Id: uint64(req.Id)})
	if err != nil {
		return nil, err
	}
	if !isExist.IsExist{
		return nil,errorx.NewCodeError(code.NoFoundError,"种类不存在")
	}
	isSuccess, err := l.svcCtx.BookRpc.DelType(l.ctx, &bookclient.IdReq{Id: uint64(req.Id)})
	if err != nil {
		return nil, err
	}
	if !isSuccess.IsSuccess{
		return nil,errorx.NewCodeError(code.DelError,"删除失败")
	}
	return &types.Reply{Code : code.Success, Msg: "删除成功"}, nil
}
