package logic

import (
	"context"

	"go-zero-admin-server/service/book/rpc/book"
	"go-zero-admin-server/service/book/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type IsExistBookByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsExistBookByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsExistBookByIdLogic {
	return &IsExistBookByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsExistBookByIdLogic) IsExistBookById(in *book.IdReq) (*book.IsExistReply, error) {
	isExist, err := l.svcCtx.BookModel.IsExistBookById(uint(in.Id))
	if err != nil {
		return nil, err
	}
	return &book.IsExistReply{IsExist: isExist}, nil
}
