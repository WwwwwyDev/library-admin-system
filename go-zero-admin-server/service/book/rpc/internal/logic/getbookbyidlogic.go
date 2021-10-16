package logic

import (
	"context"

	"go-zero-admin-server/service/book/rpc/book"
	"go-zero-admin-server/service/book/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetBookByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBookByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBookByIdLogic {
	return &GetBookByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBookByIdLogic) GetBookById(in *book.IdReq) (*book.BookInfoReply, error) {
	_book, err := l.svcCtx.BookModel.GetBookByID(uint(in.Id))
	if err != nil {
		return nil,err
	}
	return &book.BookInfoReply{Id: uint64(_book.ID),Name: _book.Name,Author: _book.Author,Image: _book.Image,Type: _book.Type.Name}, nil
}
