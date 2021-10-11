package logic

import (
	"context"
	"go-zero-admin-server/service/user/rpc/internal/svc"
	"go-zero-admin-server/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserByUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByUsernameLogic {
	return &GetUserByUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByUsernameLogic) GetUserByUsername(in *user.UsernameReq) (*user.UserInfoReply, error) {
	_user, err := l.svcCtx.UserModel.GetUserByUsername(in.Username)
	if err != nil {
		return nil, err
	}
	rolesHandle := make([]*user.RoleReply, 0)
	for _, e := range _user.Role {
		rolesHandle = append(rolesHandle, &user.RoleReply{Id: uint64(e.ID), Name: e.Name, Info: e.Info})
	}
	return &user.UserInfoReply{Id: uint64(_user.ID),Username: _user.Username,Password:_user.Password,Salt: _user.Salt,Info: _user.Info,Roles: rolesHandle}, nil
}
