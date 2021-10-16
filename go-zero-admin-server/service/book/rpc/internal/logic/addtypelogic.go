package logic

import (
	"context"
	"go-zero-admin-server/service/book/model"

	"go-zero-admin-server/service/book/rpc/book"
	"go-zero-admin-server/service/book/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTypeLogic {
	return &AddTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddTypeLogic) AddType(in *book.TypeAddReq) (*book.IsSuccessReply, error) {
	isSuccess,err := l.svcCtx.TypeModel.AddType(&model.Type{Name: in.Name, Intro: in.Intro})
	if err != nil {
		return nil, err
	}
	return &book.IsSuccessReply{IsSuccess: isSuccess}, nil
}
