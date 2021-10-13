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

type GetUserByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserByIdLogic {
	return GetUserByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserByIdLogic) GetUserById(req types.IdReq) (*types.Reply, error) {
	if req.Id <= 0{
		return nil,errorx.NewCodeError(code.ParameterError,"参数不合法")
	}
	isExistResp, err := l.svcCtx.UserRpc.IsExistUserById(l.ctx,&userclient.IdReq{Id: uint64(req.Id)})
	if err != nil{
		return nil, err
	}
	if !isExistResp.IsExist{
		return nil, errorx.NewCodeError(code.NoFoundError,"用户不存在")
	}
	user, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &userclient.IdReq{Id: uint64(req.Id)})
	if err != nil{
		return nil, err
	}
	user.Password = "禁止访问该数据"
	user.Salt = "禁止访问该数据"
	return &types.Reply{Code: code.Success, Data: map[string]interface{}{"user": user}, Msg: "查询用户成功"}, nil
}
