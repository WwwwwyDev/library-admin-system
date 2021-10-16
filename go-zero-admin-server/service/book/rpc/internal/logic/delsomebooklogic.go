package logic

import (
	"context"

	"go-zero-admin-server/service/book/rpc/book"
	"go-zero-admin-server/service/book/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type DelSomeBookLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelSomeBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelSomeBookLogic {
	return &DelSomeBookLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelSomeBookLogic) DelSomeBook(in *book.IdsReq) (*book.IsSuccessReply, error) {
	ids := make([]uint, 0)
	for _,e := range in.Ids{
		ids = append(ids,uint(e.Id))
	}
	isSuccess, err := l.svcCtx.BookModel.DelSomeBook(ids)
	if err != nil {
		return nil, err
	}
	return &book.IsSuccessReply{IsSuccess: isSuccess}, nil
}
