package logic

import (
	"context"

	"go-zero-admin-server/service/user/rpc/internal/svc"
	"go-zero-admin-server/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAllRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllRoleLogic {
	return &GetAllRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllRoleLogic) GetAllRole(in *user.EmptyReq) (*user.RolesReply, error) {
	roles, err := l.svcCtx.RoleModel.GetAllRoles()
	if err != nil {
		return nil, err
	}
	rolesHandle := make([]*user.RoleReply,0)
	for _,e := range roles{
		rolesHandle =append(rolesHandle, &user.RoleReply{Id: uint64(e.ID),Name:e.Name,Info: e.Info})
	}
	return &user.RolesReply{Roles: rolesHandle}, nil
}
