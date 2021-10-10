package login

import (
	"context"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/common/errorx"
	"go-zero-admin-server/service/user/rpc/userclient"

	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq) (*types.Reply, error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserRpc.GetUserByUsername(l.ctx, &userclient.UsernameReq{Username: req.Username})
	if err != nil {
		return nil, errorx.NewCodeError(code.Error,err.Error())
	}
	return &types.Reply{Code:code.Success,Data: map[string]interface{}{"user":user}, Msg: "登录成功"}, nil
}
