package logic

import (
	"context"
	"go-zero-admin-server/service/book/rpc/book"
	"go-zero-admin-server/service/book/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetBookByNameLikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBookByNameLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBookByNameLikeLogic {
	return &GetBookByNameLikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBookByNameLikeLogic) GetBookByNameLike(in *book.NameReq) (*book.BooksInfoReply, error) {
	_books, err := l.svcCtx.BookModel.GetBookByNameLike(in.Name)
	if err != nil {
		return nil,err
	}
	booksHandle := make([]*book.BookInfoReply, 0)
	for _, e := range _books {
		if e.Type == nil{
			booksHandle = append(booksHandle,&book.BookInfoReply{Id: uint64(e.ID), Name: e.Name,Author: e.Author,Image: e.Image,Info: e.Info,Type: ""})
		}else{
			booksHandle = append(booksHandle,&book.BookInfoReply{Id: uint64(e.ID), Name: e.Name,Author: e.Author,Image: e.Image,Info: e.Info,Type: e.Type.Name})
		}
	}
	return &book.BooksInfoReply{BooksInfo: booksHandle}, nil
}
