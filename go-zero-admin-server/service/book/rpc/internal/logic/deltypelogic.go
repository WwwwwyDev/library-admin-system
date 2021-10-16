package logic

import (
	"context"
	"go-zero-admin-server/service/book/model"
	"gorm.io/gorm"

	"go-zero-admin-server/service/book/rpc/book"
	"go-zero-admin-server/service/book/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type DelTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelTypeLogic {
	return &DelTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelTypeLogic) DelType(in *book.IdReq) (*book.IsSuccessReply, error) {
	isSuccess, err := l.svcCtx.TypeModel.DelType(&model.Type{Model: gorm.Model{ID: uint(in.Id)}})
	if err != nil {
		return nil,err
	}
	return &book.IsSuccessReply{IsSuccess: isSuccess}, nil
}
