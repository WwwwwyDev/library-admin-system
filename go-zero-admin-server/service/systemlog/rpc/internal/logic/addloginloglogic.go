package logic

import (
	"context"
	"go-zero-admin-server/service/systemlog/model"

	"go-zero-admin-server/service/systemlog/rpc/internal/svc"
	"go-zero-admin-server/service/systemlog/rpc/systemlog"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddLoginLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLoginLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLoginLogLogic {
	return &AddLoginLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLoginLogLogic) AddLoginLog(in *systemlog.LoginLogAddReq) (*systemlog.IsSuccessReply, error) {
	isSuccess, err := l.svcCtx.LoginLogModel.AddLoginLog(&model.LoginLog{Username: in.Username, Info: in.Info})
	if err != nil {
		return nil, err
	}
	return &systemlog.IsSuccessReply{IsSuccess: isSuccess}, nil
}
