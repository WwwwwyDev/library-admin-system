package logic

import (
	"context"
	"go-zero-admin-server/service/user/model"

	"go-zero-admin-server/service/user/rpc/internal/svc"
	"go-zero-admin-server/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserLogic) AddUser(in *user.UserAddReq) (*user.IsSuccessReply, error) {
	isSuccess, err := l.svcCtx.UserModel.AddUser(&model.User{Username: in.Username, Password: in.Password, Salt: in.Salt, Info: in.Info})
	if err != nil {
		return nil, err
	}
	return &user.IsSuccessReply{IsSuccess: isSuccess}, nil
}
