package logic

import (
	"context"

	"go-zero-admin-server/service/book/rpc/book"
	"go-zero-admin-server/service/book/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetTypeByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTypeByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTypeByIdLogic {
	return &GetTypeByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTypeByIdLogic) GetTypeById(in *book.IdReq) (*book.TypeInfoReply, error) {
	_type, err := l.svcCtx.TypeModel.GetTypeById(uint(in.Id))
	if err != nil {
		return nil,err
	}
	booksHandle := make([]*book.BookInfoReply, 0)
	for _, e := range _type.Book {
		booksHandle = append(booksHandle, &book.BookInfoReply{Id: uint64(e.ID), Name: e.Name,Author: e.Author,Image: e.Image,Info: e.Info})
	}
	return &book.TypeInfoReply{Id:uint64(_type.ID),Name: _type.Name,Intro: _type.Intro,BooksInfo: booksHandle}, nil
}
