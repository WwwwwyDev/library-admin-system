package logic

import (
	"context"

	"go-zero-admin-server/service/book/rpc/book"
	"go-zero-admin-server/service/book/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type IsExistBookByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsExistBookByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsExistBookByNameLogic {
	return &IsExistBookByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsExistBookByNameLogic) IsExistBookByName(in *book.NameReq) (*book.IsExistReply, error) {
	isExist, err := l.svcCtx.BookModel.IsExistBookByName(in.Name)
	if err != nil {
		return nil, err
	}
	return &book.IsExistReply{IsExist: isExist}, nil
}
