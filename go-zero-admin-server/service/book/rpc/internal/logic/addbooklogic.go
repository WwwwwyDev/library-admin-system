package logic

import (
	"context"
	"go-zero-admin-server/service/book/model"

	"go-zero-admin-server/service/book/rpc/book"
	"go-zero-admin-server/service/book/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddBookLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddBookLogic {
	return &AddBookLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddBookLogic) AddBook(in *book.BookAddReq) (*book.IsSuccessReply, error) {
	isSuccess, err := l.svcCtx.BookModel.AddBook(&model.Book{Name: in.Name, Image: in.Image, Author: in.Author, Info: in.Info, TypeID: uint(in.TypeId)})
	if err != nil {
		return nil, err
	}
	return &book.IsSuccessReply{IsSuccess: isSuccess}, nil
}
