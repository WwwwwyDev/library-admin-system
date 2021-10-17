package book

import (
	"context"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/common/errorx"
	"go-zero-admin-server/service/book/rpc/book"
	"go-zero-admin-server/service/book/rpc/bookclient"

	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateBookLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateBookLogic {
	return UpdateBookLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateBookLogic) UpdateBook(req types.UpdateBookReq) (*types.Reply, error) {
	isExist, err := l.svcCtx.BookRpc.IsExistBookById(l.ctx, &bookclient.IdReq{Id: uint64(req.Id)})
	if err != nil {
		return nil, err
	}
	if !isExist.IsExist{
		return nil,errorx.NewCodeError(code.NoFoundError,"图书不存在")
	}
	isExist, err = l.svcCtx.BookRpc.IsExistTypeById(l.ctx, &bookclient.IdReq{Id: uint64(req.TypeId)})
	if err != nil {
		return nil, err
	}
	if !isExist.IsExist{
		return nil, errorx.NewCodeError(code.NoFoundError,"不存在此图书类型")
	}
	oldBook, err := l.svcCtx.BookRpc.GetBookById(l.ctx, &bookclient.IdReq{Id: uint64(req.Id)})
	if oldBook.Name != req.Name{
		isExist, err = l.svcCtx.BookRpc.IsExistBookByName(l.ctx, &bookclient.NameReq{Name: req.Name})
		if err != nil {
			return nil, err
		}
		if isExist.IsExist{
			return nil, errorx.NewCodeError(code.RepeatError,"图书名重复")
		}
	}
	isSuccess, err := l.svcCtx.BookRpc.UpdateBook(l.ctx, &book.BookUpdateReq{Id: uint64(req.Id), Name: req.Name, Image: req.Image, Author: req.Author, Info: req.Info, TypeId: uint64(req.TypeId)})
	if err != nil {
		return nil, err
	}
	if !isSuccess.IsSuccess{
		return nil,errorx.NewCodeError(code.ChangeError,"修改图书失败")
	}
	return &types.Reply{Code:  code.Success, Msg: "修改图书成功"}, nil
}
