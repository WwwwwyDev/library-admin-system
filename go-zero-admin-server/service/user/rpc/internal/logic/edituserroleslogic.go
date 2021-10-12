package logic

import (
	"context"
	"go-zero-admin-server/service/user/model"
	"go-zero-admin-server/service/user/rpc/internal/svc"
	"go-zero-admin-server/service/user/rpc/user"
	"gorm.io/gorm"

	"github.com/tal-tech/go-zero/core/logx"
)

type EditUserRolesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditUserRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditUserRolesLogic {
	return &EditUserRolesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EditUserRolesLogic) EditUserRoles(in *user.RolesReq) (*user.IsSuccessReply, error) {
	userId := uint(in.UserId)
	_user, err := l.svcCtx.UserModel.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	roleReplyHandle := make([]*model.Role,0)
	for _, e := range in.Ids{
		roleReplyHandle = append(roleReplyHandle,&model.Role{Model:gorm.Model{ID: uint(e.Id)}})
	}
	_user.Role = roleReplyHandle
	isSuccess, err := l.svcCtx.RoleModel.EditUserRoles(_user)
	if err != nil {
		return nil, err
	}
	return &user.IsSuccessReply{IsSuccess: isSuccess}, nil
}
