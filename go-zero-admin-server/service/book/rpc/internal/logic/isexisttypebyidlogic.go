package logic

import (
	"context"

	"go-zero-admin-server/service/book/rpc/book"
	"go-zero-admin-server/service/book/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type IsExistTypeByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsExistTypeByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsExistTypeByIdLogic {
	return &IsExistTypeByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsExistTypeByIdLogic) IsExistTypeById(in *book.IdReq) (*book.IsExistReply, error) {
	isExist, err := l.svcCtx.TypeModel.IsExistTypeById(uint(in.Id))
	if err != nil {
		return nil, err
	}
	return &book.IsExistReply{IsExist: isExist}, nil
}
