package logic

import (
	"context"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/common/errorx"
	"go-zero-admin-server/service/admin/user/model"
	"gorm.io/gorm"

	"go-zero-admin-server/service/admin/user/api/internal/svc"
	"go-zero-admin-server/service/admin/user/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteUserLogic {
	return DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req types.DeleteUserReq) (*types.Reply, error) {
	isExist, err := l.svcCtx.UserModel.IsExistUserByID(req.ID)
	if err != nil{
		return nil,err
	}
	if !isExist{
		return nil, errorx.NewCodeError(code.NoFoundError,"用户不存在")
	}
	err = l.svcCtx.UserModel.DelUser(&model.User{Model: gorm.Model{ID: req.ID}})
	if err != nil {
		return nil, err
	}
	return &types.Reply{Code: code.Success, Msg: "ok"}, nil
}
