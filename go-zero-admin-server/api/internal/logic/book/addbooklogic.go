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

type AddBookLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddBookLogic {
	return AddBookLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddBookLogic) AddBook(req types.AddBookReq) (*types.Reply, error) {
	isExist, err := l.svcCtx.BookRpc.IsExistTypeById(l.ctx, &bookclient.IdReq{Id: uint64(req.TypeId)})
	if err != nil {
		return nil, err
	}
	if !isExist.IsExist{
		return nil, errorx.NewCodeError(code.NoFoundError,"不存在此图书类型")
	}
	isExist, err = l.svcCtx.BookRpc.IsExistBookByName(l.ctx, &bookclient.NameReq{Name: req.Name})
	if err != nil {
		return nil, err
	}
	if isExist.IsExist{
		return nil, errorx.NewCodeError(code.RepeatError,"图书名重复")
	}
	isSuccessResp, err := l.svcCtx.BookRpc.AddBook(l.ctx, &bookclient.BookAddReq{Name: req.Name, Image: req.Image, Author:req.Author,Info: req.Info, TypeId: uint64(req.TypeId)})
	if err != nil {
		return nil,err
	}
	if !isSuccessResp.IsSuccess {
		return nil, errorx.NewCodeError(code.AddError,"添加图书失败")
	}
	return &types.Reply{Code: code.Success,Msg: "添加成功"}, nil
}
