package book

import (
	"context"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/common/errorx"
	"go-zero-admin-server/service/book/rpc/bookclient"

	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetTypesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTypesLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetTypesLogic {
	return GetTypesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTypesLogic) GetTypes(req types.GetTypesReq) (*types.Reply, error) {
	page := req.Page
	limit := req.Limit
	if page <=0 || limit <=0{
		return nil, errorx.NewCodeError(code.ParameterError,"参数非法")
	}
	typesResp, err := l.svcCtx.BookRpc.GetTypes(l.ctx, &bookclient.TypesReq{Page: int64(req.Page), Limit: int64(req.Limit), Name: req.Name})
	if err != nil {
		return nil, err
	}
	return &types.Reply{Code: code.Success, Data: map[string]interface{}{"types": typesResp.TypesInfo, "total": typesResp.Total}, Msg: "查询用户成功"}, nil
}
