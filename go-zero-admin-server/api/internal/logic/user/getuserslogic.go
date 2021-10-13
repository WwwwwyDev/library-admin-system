package user

import (
	"context"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/common/errorx"
	"go-zero-admin-server/service/user/rpc/userclient"

	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUsersLogic {
	return GetUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUsersLogic) GetUsers(req types.GetUsersReq) (*types.Reply, error) {
	page := req.Page
	limit := req.Limit
	if page <=0 || limit <=0{
		return nil, errorx.NewCodeError(code.ParameterError,"参数非法")
	}
	usersResp, err := l.svcCtx.UserRpc.GetUsers(l.ctx, &userclient.UsersReq{Page: int64(req.Page), Limit: int64(req.Limit), Username: req.Username})
	if err != nil {
		return nil, err
	}
	for i,_ := range usersResp.UsersInfo{
		usersResp.UsersInfo[i].Password = "禁止访问该数据"
		usersResp.UsersInfo[i].Salt = "禁止访问该数据"
	}
	return &types.Reply{Code: code.Success, Data: map[string]interface{}{"users": usersResp.UsersInfo, "total": usersResp.Total}, Msg: "查询用户成功"}, nil
}
