package logic

import (
	"context"

	"go-zero-admin-server/service/user/rpc/internal/svc"
	"go-zero-admin-server/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type IsExistRoleByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsExistRoleByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsExistRoleByIdLogic {
	return &IsExistRoleByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsExistRoleByIdLogic) IsExistRoleById(in *user.IdReq) (*user.IsExistReply, error) {
	isExist, err := l.svcCtx.RoleModel.IsExistRoleById(uint(in.Id))
	if err != nil {
		return nil, err
	}
	return &user.IsExistReply{IsExist: isExist}, nil
}
