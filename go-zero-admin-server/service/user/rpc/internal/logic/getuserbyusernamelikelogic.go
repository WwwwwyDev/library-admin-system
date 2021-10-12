package logic

import (
	"context"
	"go-zero-admin-server/service/user/rpc/internal/svc"
	"go-zero-admin-server/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserByUsernameLikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByUsernameLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByUsernameLikeLogic {
	return &GetUserByUsernameLikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByUsernameLikeLogic) GetUserByUsernameLike(in *user.UsernameReq) (*user.UsersInfoReply, error) {
	_users, err := l.svcCtx.UserModel.GetUserByUsernameLike(in.Username)
	if err != nil {
		return nil, err
	}
	usersHandle := make([]*user.UserInfoReply, 0)
	for _, e := range _users {
		usersHandle = append(usersHandle, &user.UserInfoReply{Id: uint64(e.ID), Username: e.Username, Password: e.Password, Salt: e.Salt, Avatar:e.Avatar,Info: e.Info})
	}
	return &user.UsersInfoReply{UsersInfo: usersHandle}, nil
}
