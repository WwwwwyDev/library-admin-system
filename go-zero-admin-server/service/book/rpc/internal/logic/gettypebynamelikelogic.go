package logic

import (
	"context"

	"go-zero-admin-server/service/book/rpc/book"
	"go-zero-admin-server/service/book/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetTypeByNameLikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTypeByNameLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTypeByNameLikeLogic {
	return &GetTypeByNameLikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTypeByNameLikeLogic) GetTypeByNameLike(in *book.NameReq) (*book.TypesInfoReply, error) {
	_type, err := l.svcCtx.TypeModel.GetTypeByNameLike(in.Name)
	if err != nil {
		return nil, err
	}
	typesHandle := make([]*book.TypeInfoReply, 0)
	for _, e := range _type {
		//booksHandle := make([]*book.BookInfoReply,0)
		//for _, e1 := range e.Book {
		//	booksHandle = append(booksHandle,&book.BookInfoReply{Id: uint64(e1.ID),Name: e1.Name,Image: e1.Image,Author: e1.Author,Info: e1.Info})
		//}
		typesHandle = append(typesHandle, &book.TypeInfoReply{Id: uint64(e.ID),Name: e.Name,Intro: e.Intro})
	}
	return &book.TypesInfoReply{}, nil
}
