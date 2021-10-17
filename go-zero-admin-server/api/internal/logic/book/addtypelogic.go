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

type AddTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddTypeLogic {
	return AddTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTypeLogic) AddType(req types.AddTypeReq) (*types.Reply, error) {
	isExist, err := l.svcCtx.BookRpc.IsExistTypeByName(l.ctx, &bookclient.NameReq{Name: req.Name})
	if isExist.IsExist{
		return nil, errorx.NewCodeError(code.RepeatError,"种类名称重复")
	}
	isSuccess, err := l.svcCtx.BookRpc.AddType(l.ctx, &bookclient.TypeAddReq{Name: req.Name, Intro: req.Intro})
	if err != nil {
		return nil, err
	}
	if !isSuccess.IsSuccess{
		return nil, errorx.NewCodeError(code.AddError,"添加失败")
	}
	return &types.Reply{Code: code.Success,Msg: "添加成功"}, nil
}
