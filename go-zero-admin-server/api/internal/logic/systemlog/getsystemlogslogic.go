package systemlog

import (
	"context"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/service/systemlog/rpc/systemlogclient"

	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetSystemlogsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSystemlogsLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetSystemlogsLogic {
	return GetSystemlogsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSystemlogsLogic) GetSystemlogs(req types.GetSystemlogsReq) (*types.Reply, error) {
	logs, err := l.svcCtx.SystemlogRpc.GetLoginLogs(l.ctx, &systemlogclient.LoginLogsReq{Page: int64(req.Page), Limit: int64(req.Limit), Username: req.Username})
	if err != nil {
		return nil, err
	}
	return &types.Reply{Code:code.Success,Data: map[string]interface{}{"logs":logs.LoginLogsInfo,"total":logs.Total},Msg:"获取成功"}, nil
}
