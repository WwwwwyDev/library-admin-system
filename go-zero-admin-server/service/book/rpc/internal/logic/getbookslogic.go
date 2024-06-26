package logic

import (
	"context"
	"go-zero-admin-server/service/book/model"
	"go-zero-admin-server/service/book/rpc/book"
	"go-zero-admin-server/service/book/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetBooksLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBooksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBooksLogic {
	return &GetBooksLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBooksLogic) GetBooks(in *book.BooksReq) (*book.BooksInfoReply, error) {
	books, total, err := l.svcCtx.BookModel.GetBooks(int(in.Page), int(in.Limit), &model.Book{Name: in.Name, Author: in.Author})
	if err != nil {
		return nil, err
	}
	booksHandle := make([]*book.BookInfoReply, 0)
	for _, e := range books {
		if e.Type == nil{
			booksHandle = append(booksHandle,&book.BookInfoReply{Id: uint64(e.ID), Name: e.Name,Author: e.Author,Image: e.Image,Info: e.Info,Type: ""})
		}else{
			booksHandle = append(booksHandle,&book.BookInfoReply{Id: uint64(e.ID), Name: e.Name,Author: e.Author,Image: e.Image,Info: e.Info,Type: e.Type.Name,TypeId: uint64(e.TypeID)})
		}
	}
	return &book.BooksInfoReply{BooksInfo: booksHandle,Total: total}, nil
}
