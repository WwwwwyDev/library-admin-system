package verify

import (
	"context"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/common/errorx"
	"go-zero-admin-server/common/util"
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
	isExistResp, err := l.svcCtx.UserRpc.IsExistUserByUsername(l.ctx, &userclient.UsernameReq{Username: req.Username})
	if err != nil {
		return nil, errorx.NewCodeError(code.Error,err.Error())
	}
	if !isExistResp.IsExist {
		return nil, errorx.NewCodeError(code.NoFoundError,"用户不存在")
	}
	userResp, err := l.svcCtx.UserRpc.GetUserByUsername(l.ctx, &userclient.UsernameReq{Username: req.Username})
	if err != nil {
		return nil, errorx.NewCodeError(code.Error,err.Error())
	}
	passwordmd5 := util.Str2Md5(req.Password + userResp.Salt)
	if passwordmd5 != userResp.Password {
		return nil, errorx.NewCodeError(code.PasswordError,"用户密码错误")
	}
	return &types.Reply{Code:code.Success,Data: map[string]interface{}{"auth_token":"dwadadwadad"}, Msg: "登录成功"}, nil
}
