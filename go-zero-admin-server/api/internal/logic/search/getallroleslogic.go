package search

import (
	"context"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/service/user/rpc/userclient"

	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAllRolesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAllRolesLogic {
	return GetAllRolesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllRolesLogic) GetAllRoles() (*types.Reply, error) {
	roles, err := l.svcCtx.UserRpc.GetAllRole(l.ctx, &userclient.EmptyReq{})
	if err != nil {
		return nil, err
	}
	return &types.Reply{Code: code.Success, Data: map[string]interface{}{"roles": roles}, Msg: "查询成功"}, nil
}
