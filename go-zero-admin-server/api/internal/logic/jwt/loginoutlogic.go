package jwt

import (
	"context"
	"fmt"
	"go-zero-admin-server/common/code"

	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginOutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginOutLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginOutLogic {
	return LoginOutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginOutLogic) LoginOut() (*types.Reply, error) {
	id := fmt.Sprintf("%v", l.ctx.Value("userId"))
	l.svcCtx.Redis.Del("loginUserId:"+id)
	l.svcCtx.Redis.Do("srem","loginStatus",id)
	return &types.Reply{Code: code.Success, Msg: "登出成功"}, nil
}
