package logic

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	"go-zero-admin-server/service/admin/user/rpc/internal/svc"
	"go-zero-admin-server/service/admin/user/rpc/user"
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

func (l *GetUserByUsernameLogic) GetUserByUsername(in *user.UsernameReq) (*user.UserReply, error) {
	_user, err := l.svcCtx.UserModel.GetUserByUsername(in.Username)
	if err != nil {
		return nil, err
	}
	rolesHandle := make([]*user.Role,0)
	for _,e := range _user.Role{
		rolesHandle =append(rolesHandle, &user.Role{ID: uint64(e.ID),Name:e.Name,Info: e.Info})
	}
	return &user.UserReply{ID: uint64(_user.ID),Username: _user.Username,Salt: _user.Salt,Info: _user.Info,Roles: rolesHandle}, nil
}
