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

type UpdateTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateTypeLogic {
	return UpdateTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTypeLogic) UpdateType(req types.UpdateTypeReq) (*types.Reply, error) {
	isExist, err := l.svcCtx.BookRpc.IsExistTypeById(l.ctx, &bookclient.IdReq{Id: uint64(req.Id)})
	if err != nil {
		return nil, err
	}
	if !isExist.IsExist{
		return nil,errorx.NewCodeError(code.NoFoundError,"种类不存在")
	}
	oldType, err := l.svcCtx.BookRpc.GetTypeById(l.ctx, &bookclient.IdReq{Id: uint64(req.Id)})
	if oldType.Name != req.Name{
		isExist, err = l.svcCtx.BookRpc.IsExistTypeByName(l.ctx, &bookclient.NameReq{Name: req.Name})
		if err != nil {
			return nil, err
		}
		if isExist.IsExist{
			return nil, errorx.NewCodeError(code.RepeatError,"种类名重复")
		}
	}
	isSuccess, err := l.svcCtx.BookRpc.UpdateType(l.ctx, &bookclient.TypeUpdateReq{Id: uint64(req.Id), Name: req.Name, Intro: req.Intro})
	if err != nil {
		return nil, err
	}
	if !isSuccess.IsSuccess{
		return nil,errorx.NewCodeError(code.ChangeError,"修改种类失败")
	}
	return &types.Reply{Code:  code.Success, Msg: "修改种类成功"}, nil
}
