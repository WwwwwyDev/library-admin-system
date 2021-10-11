package search

import (
	"context"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/common/errorx"
	"go-zero-admin-server/service/user/rpc/userclient"

	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserByUserNameLikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserByUserNameLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserByUserNameLikeLogic {
	return GetUserByUserNameLikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserByUserNameLikeLogic) GetUserByUserNameLike(req types.UsernameReq) (*types.Reply, error) {
	userResp, err := l.svcCtx.UserRpc.GetUserByUsernameLike(l.ctx, &userclient.UsernameReq{Username: req.Username})
	if err != nil {
		return nil, errorx.NewCodeError(code.Error, err.Error())
	}
	users := userResp.UsersInfo
	for i, _ := range users {
		users[i].Password = "禁止访问该数据"
		users[i].Salt = "禁止访问该数据"
	}
	return &types.Reply{Code: code.Success, Data: map[string]interface{}{"users": users}, Msg: "查询用户成功"}, nil
}
