package logic

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	"go-zero-admin-server/service/admin/user/rpc/internal/svc"
	"go-zero-admin-server/service/admin/user/rpc/user"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *user.IDReq) (*user.UserReply, error) {
	_user, err := l.svcCtx.UserModel.GetUserByID(uint(in.ID))
	if err != nil {
		return nil, err
	}
	rolesHandle := make([]*user.Role,0)
	for _,e := range _user.Role{
		rolesHandle =append(rolesHandle, &user.Role{ID: uint64(e.ID),Name:e.Name,Info: e.Info})
	}
	return &user.UserReply{ID: uint64(_user.ID),Username: _user.Username,Salt: _user.Salt,Info: _user.Info,Roles: rolesHandle}, nil
}
