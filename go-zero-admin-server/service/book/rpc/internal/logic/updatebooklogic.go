package logic

import (
	"context"
	"go-zero-admin-server/service/book/model"
	"gorm.io/gorm"

	"go-zero-admin-server/service/book/rpc/book"
	"go-zero-admin-server/service/book/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateBookLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBookLogic {
	return &UpdateBookLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateBookLogic) UpdateBook(in *book.BookUpdateReq) (*book.IsSuccessReply, error) {
	isSuccess, err := l.svcCtx.BookModel.UpdateBook(&model.Book{Model: gorm.Model{ID: uint(in.Id)},Name: in.Name, Image: in.Image, Author: in.Author, Info: in.Info, TypeID: uint(in.TypeId)})
	if err != nil {
		return nil, err
	}
	return &book.IsSuccessReply{IsSuccess: isSuccess}, nil
}
