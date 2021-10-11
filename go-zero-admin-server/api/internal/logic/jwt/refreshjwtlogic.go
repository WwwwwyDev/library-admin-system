package jwt

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/common/errorx"
	"strconv"
	"time"

	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RefreshJwtLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshJwtLogic(ctx context.Context, svcCtx *svc.ServiceContext) RefreshJwtLogic {
	return RefreshJwtLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshJwtLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func (l *RefreshJwtLogic) RefreshJwt() (*types.Reply, error) {
	sprintf := fmt.Sprintf("%v", l.ctx.Value("userId"))
	userId, err := strconv.ParseInt(sprintf,10,64)
	if err != nil{
		return nil, errorx.NewCodeError(code.Error, err.Error())
	}
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, userId)
	if err != nil {
		return nil, errorx.NewCodeError(code.Error, err.Error())
	}
	return &types.Reply{Code: code.Success, Data: map[string]interface{}{"accessToken": jwtToken,
		"accessExpire": now + accessExpire,
		"refreshAfter": now + accessExpire/2}, Msg: "刷新成功"}, nil
}
