package user

import (
	"context"
	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/common/errorx"
	"go-zero-admin-server/service/user/rpc/userclient"

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
	if req.Id == 1 {
		return nil,errorx.NewCodeError(code.DelError,"拒绝删除超级管理员")
	}
	isExistResp, err := l.svcCtx.UserRpc.IsExistUserById(l.ctx,&userclient.IdReq{Id: uint64(req.Id)})
	if err != nil{
		return nil, err
	}
	if !isExistResp.IsExist{
		return nil, errorx.NewCodeError(code.NoFoundError,"用户不存在")
	}
	isSuccessResp, err := l.svcCtx.UserRpc.DelUser(l.ctx, &userclient.IdReq{Id: uint64(req.Id)})
	if err != nil {
		return nil, err
	}
	if !isSuccessResp.IsSuccess{
		return nil,errorx.NewCodeError(code.DelError,"删除用户失败")
	}
	return &types.Reply{Code: code.Success, Msg: "删除用户成功"}, nil
}
