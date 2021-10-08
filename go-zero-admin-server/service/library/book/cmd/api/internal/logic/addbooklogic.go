package logic

import (
	"context"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/common/errorx"
	"go-zero-admin-server/service/library/book/cmd/api/internal/svc"
	"go-zero-admin-server/service/library/book/cmd/api/internal/types"
	"go-zero-admin-server/service/library/book/model"

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
	isExist, err := l.svcCtx.BookModel.IsExistBookByName(req.Name)
	if err != nil{
		return nil,err
	}
	if isExist{
		return nil, errorx.NewCodeError(code.RepeatError,"图书名重复")
	}
	err = l.svcCtx.BookModel.AddBook(&model.Book{Name: req.Name, Image: req.Image, Author: req.Author, Info: req.Info})
	if err != nil{
		return nil,err
	}
	return &types.Reply{Code: code.Success,Msg: "ok"}, nil
}
