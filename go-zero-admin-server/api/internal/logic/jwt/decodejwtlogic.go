package jwt

import (
	"context"
	"fmt"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/common/errorx"
	"go-zero-admin-server/service/user/rpc/userclient"
	"strconv"

	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DecodeJwtLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDecodeJwtLogic(ctx context.Context, svcCtx *svc.ServiceContext) DecodeJwtLogic {
	return DecodeJwtLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *DecodeJwtLogic) DecodeJwt() (*types.Reply, error) {
	sprintf := fmt.Sprintf("%v", l.ctx.Value("userId"))
	userId, err := strconv.ParseUint(sprintf,10,64)
	if err != nil{
		return nil, errorx.NewCodeError(code.Error, err.Error())
	}
	user, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &userclient.IdReq{Id: userId})
	if err != nil{
		return nil, errorx.NewCodeError(code.Error, err.Error())
	}
	user.Password = "禁止访问该数据"
	user.Salt = "禁止访问该数据"
	return &types.Reply{Code: code.Success, Data: map[string]interface{}{"user":user},Msg: "解析成功"}, nil
}
