package logic

import (
	"context"

	"go-zero-admin-server/service/systemlog/rpc/internal/svc"
	"go-zero-admin-server/service/systemlog/rpc/systemlog"

	"github.com/tal-tech/go-zero/core/logx"
)

type DelLoginLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelLoginLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelLoginLogLogic {
	return &DelLoginLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelLoginLogLogic) DelLoginLog(in *systemlog.EmptyReq) (*systemlog.IsSuccessReply, error) {
	isSuccess, err := l.svcCtx.LoginLogModel.ClearLoginLogs()
	if err != nil {
		return nil, err
	}
	return &systemlog.IsSuccessReply{IsSuccess: isSuccess}, nil
}
