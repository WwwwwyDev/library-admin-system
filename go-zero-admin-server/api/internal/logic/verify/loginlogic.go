package verify

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/common/errorx"
	"go-zero-admin-server/common/util"
	"go-zero-admin-server/service/user/rpc/userclient"
	"time"

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
func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func (l *LoginLogic) Login(req types.LoginReq) (*types.Reply, error) {
	if len(req.Username) < 6 {
		return nil,errorx.NewCodeError(code.ParameterError, "用户名非法")
	}
	if len(req.Password) < 6 {
		return nil,errorx.NewCodeError(code.ParameterError, "密码非法")
	}
	isExistResp, err := l.svcCtx.UserRpc.IsExistUserByUsername(l.ctx, &userclient.UsernameReq{Username: req.Username})
	if err != nil {
		return nil, err
	}
	if !isExistResp.IsExist {
		return nil, errorx.NewCodeError(code.NoFoundError, "用户不存在")
	}
	userResp, err := l.svcCtx.UserRpc.GetUserByUsername(l.ctx, &userclient.UsernameReq{Username: req.Username})
	if err != nil {
		return nil, err
	}
	passwordmd5 := util.Str2Md5(req.Password + userResp.Salt)
	if passwordmd5 != userResp.Password {
		return nil, errorx.NewCodeError(code.PasswordError, "用户密码错误")
	}
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, int64(userResp.Id))
	if err != nil {
		return nil, err
	}
	idstr := fmt.Sprintf("%d",userResp.Id)
	l.svcCtx.Redis.Set("loginUserId:"+idstr,userResp.Username+";"+userResp.Avatar+";"+userResp.Info,time.Duration(accessExpire)*time.Second)
	return &types.Reply{Code: code.Success, Data: map[string]interface{}{"accessToken": jwtToken,
		"accessExpire": now + accessExpire,
		"refreshAfter": now + accessExpire/2}, Msg: "登录成功"}, nil
}
