package systemlog

import (
	"context"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/common/errorx"
	"go-zero-admin-server/service/systemlog/rpc/systemlogclient"

	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteAllSystemlogsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteAllSystemlogsLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteAllSystemlogsLogic {
	return DeleteAllSystemlogsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteAllSystemlogsLogic) DeleteAllSystemlogs() (*types.Reply, error) {
	isSuccess, err := l.svcCtx.SystemlogRpc.DelLoginLog(l.ctx, &systemlogclient.EmptyReq{})
	if err != nil {
		return nil, err
	}
	if !isSuccess.IsSuccess{
		return nil, errorx.NewCodeError(code.DelError,"清空失败")
	}
	return &types.Reply{Code: code.Success,Msg: "清空成功"}, nil
}
