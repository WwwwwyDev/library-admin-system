package logic

import (
	"context"
	"go-zero-admin-server/service/user/model"

	"go-zero-admin-server/service/user/rpc/internal/svc"
	"go-zero-admin-server/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersLogic {
	return &GetUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUsersLogic) GetUsers(in *user.UsersReq) (*user.UsersInfoReply, error) {
	users, total, err := l.svcCtx.UserModel.GetUsers(int(in.Page), int(in.Limit), &model.User{Username: in.Username})
	if err != nil {
		return nil, err
	}
	usersHandle := make([]*user.UserInfoReply, 0)
	for _, e := range users {
		usersHandle = append(usersHandle, &user.UserInfoReply{Id: uint64(e.ID), Username: e.Username, Password: e.Password, Salt: e.Salt, Info: e.Info})
	}
	return &user.UsersInfoReply{UsersInfo: usersHandle,Total: total}, nil
}
