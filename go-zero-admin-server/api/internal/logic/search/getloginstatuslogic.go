package search

import (
	"context"
	"go-zero-admin-server/common/code"

	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetLoginStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLoginStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetLoginStatusLogic {
	return GetLoginStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLoginStatusLogic) GetLoginStatus() (*types.Reply, error) {
	result, err := l.svcCtx.Redis.Do("smembers", "loginStatus").Result()
	if err != nil {
		return nil, err
	}
	return &types.Reply{Code: code.Success, Data: map[string]interface{}{"loginStatus": result}, Msg: "查询登录状态"}, nil
}
