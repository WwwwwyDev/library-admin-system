package logic

import (
	"context"

	"go-zero-admin-server/service/book/rpc/book"
	"go-zero-admin-server/service/book/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type IsExistTypeByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsExistTypeByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsExistTypeByNameLogic {
	return &IsExistTypeByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsExistTypeByNameLogic) IsExistTypeByName(in *book.NameReq) (*book.IsExistReply, error) {
	isExist, err := l.svcCtx.TypeModel.IsExistTypeByName(in.Name)
	if err != nil {
		return nil, err
	}
	return &book.IsExistReply{IsExist: isExist}, nil
}
