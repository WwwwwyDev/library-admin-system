package logic

import (
	"context"
	"go-zero-admin-server/service/user/model"
	"gorm.io/gorm"

	"go-zero-admin-server/service/user/rpc/internal/svc"
	"go-zero-admin-server/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type DelUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserLogic {
	return &DelUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserLogic) DelUser(in *user.IdReq) (*user.IsSuccessReply, error) {
	isSuccess, err := l.svcCtx.UserModel.DelUser(&model.User{Model: gorm.Model{ID: uint(in.Id)}})
	if err != nil {
		return nil, err
	}
	return &user.IsSuccessReply{IsSuccess: isSuccess}, nil
}
