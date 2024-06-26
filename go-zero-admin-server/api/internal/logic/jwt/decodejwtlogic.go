package jwt

import (
	"context"
	"errors"
	"fmt"
	"go-zero-admin-server/common/code"
	"strconv"
	"strings"

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
	id := fmt.Sprintf("%v", l.ctx.Value("userId"))
	roles := fmt.Sprintf("%v", l.ctx.Value("roles"))
	exists,err:= l.svcCtx.Redis.Exists("loginUserId:"+id).Result()
	if err != nil {
		return nil, err
	}
	if exists!=1{
		return nil,errors.New("会话过期")
	}
	result, err := l.svcCtx.Redis.Get("loginUserId:"+id).Result()
	if err != nil {
		return nil, err
	}
	split := strings.Split(result, ";")
	idi,err := strconv.ParseUint(id,10,64)
	if err != nil {
		return nil, err
	}
	return &types.Reply{Code: code.Success, Data: map[string]interface{}{"id":idi,"username":split[0],"avatar":split[1],"info":split[2],"roles":roles},Msg: "解析成功"}, nil
}
