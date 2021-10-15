package user

import (
	"context"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/common/errorx"
	"go-zero-admin-server/service/user/rpc/userclient"

	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type EditUserRolesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditUserRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) EditUserRolesLogic {
	return EditUserRolesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditUserRolesLogic) EditUserRoles(req types.EditUserRolesReq) (*types.Reply, error) {
	if req.Id == 1 {
		return nil,errorx.NewCodeError(code.DelError,"拒绝修改超级管理员")
	}
	isExistResp, err := l.svcCtx.UserRpc.IsExistUserById(l.ctx,&userclient.IdReq{Id: uint64(req.Id)})
	if err != nil{
		return nil, err
	}
	if !isExistResp.IsExist{
		return nil, errorx.NewCodeError(code.NoFoundError,"用户不存在")
	}
	for _, e := range req.RoleIds {
		roleIsExistResp, err := l.svcCtx.UserRpc.IsExistRoleById(l.ctx, &userclient.IdReq{Id: uint64(e)})
		if err != nil {
			return nil, err
		}
		if !roleIsExistResp.IsExist {
			return nil, errorx.NewCodeError(code.NoFoundError, "不存在此权限")
		}
	}
	rolesHandle := make([]*userclient.RoleReq,0)
	for _,e := range req.RoleIds{
		rolesHandle = append(rolesHandle,&userclient.RoleReq{Id: uint64(e)})
	}
	isSuccessResp, err := l.svcCtx.UserRpc.EditUserRoles(l.ctx, &userclient.RolesReq{UserId: uint64(req.Id), Ids: rolesHandle})
	if err != nil {
		return nil, err
	}
	if !isSuccessResp.IsSuccess{
		return nil, errorx.NewCodeError(code.ChangeError, "编辑权限失败")
	}
	return &types.Reply{Code: code.Success,Msg: "编辑成功"}, nil
}
