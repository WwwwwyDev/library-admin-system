package logic

import (
	"context"
	"go-zero-admin-server/service/book/model"
	"gorm.io/gorm"

	"go-zero-admin-server/service/book/rpc/book"
	"go-zero-admin-server/service/book/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTypeLogic {
	return &UpdateTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTypeLogic) UpdateType(in *book.TypeUpdateReq) (*book.IsSuccessReply, error) {
	isSuccess,err := l.svcCtx.TypeModel.UpdateType(&model.Type{Model:gorm.Model{ID: uint(in.Id)},Name: in.Name, Intro: in.Intro})
	if err != nil {
		return nil, err
	}
	return &book.IsSuccessReply{IsSuccess: isSuccess}, nil
}
