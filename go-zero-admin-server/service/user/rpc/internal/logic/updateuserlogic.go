package logic

import (
	"context"
	"go-zero-admin-server/service/user/model"
	"gorm.io/gorm"

	"go-zero-admin-server/service/user/rpc/internal/svc"
	"go-zero-admin-server/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *user.UserUpdateReq) (*user.IsSuccessReply, error) {
	isSuccess, err := l.svcCtx.UserModel.UpdateUser(&model.User{Model: gorm.Model{ID: uint(in.Id)}, Username: in.Username, Password: in.Password, Salt: in.Salt, Info: in.Info})
	if err != nil {
		return nil, err
	}
	return &user.IsSuccessReply{IsSuccess: isSuccess}, nil
}
