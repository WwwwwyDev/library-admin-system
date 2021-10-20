package logic

import (
	"context"
	"go-zero-admin-server/service/systemlog/model"
	"go-zero-admin-server/service/systemlog/rpc/internal/svc"
	"go-zero-admin-server/service/systemlog/rpc/systemlog"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetLoginLogsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLoginLogsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLoginLogsLogic {
	return &GetLoginLogsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLoginLogsLogic) GetLoginLogs(in *systemlog.LoginLogsReq) (*systemlog.LoginLogsInfoReply, error) {
	loginLogs, total, err := l.svcCtx.LoginLogModel.GetLoginLogs(int(in.Page), int(in.Limit), &model.LoginLog{Username: in.Username})
	if err != nil {
		return nil, err
	}
	loginLogsHandle := make([]*systemlog.LoginLogInfoReply, 0)
	for _, e := range loginLogs {
		loginLogsHandle = append(loginLogsHandle, &systemlog.LoginLogInfoReply{Id:uint64(e.ID),Username: e.Username,Info: e.Info})
	}
	return &systemlog.LoginLogsInfoReply{LoginLogsInfo: loginLogsHandle,Total: total}, nil
}
