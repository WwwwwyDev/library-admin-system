package logic

import (
	"context"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/common/errorx"
	"go-zero-admin-server/service/admin/user/model"

	"go-zero-admin-server/service/admin/user/api/internal/svc"
	"go-zero-admin-server/service/admin/user/api/internal/types"

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
	users, total, err := l.svcCtx.UserModel.GetUsers(page, limit,&model.User{Username: req.Username})
	if err != nil {
		return nil,err
	}
	for i,_ := range users{
		users[i].Password = "禁止访问该数据"
		users[i].Salt = "禁止访问该数据"
	}
	return &types.Reply{Code: code.Success, Data: map[string]interface{}{"users": users, "total": total}, Msg: "ok"}, nil
}
