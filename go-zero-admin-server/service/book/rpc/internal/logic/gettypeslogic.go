package logic

import (
	"context"
	"go-zero-admin-server/service/book/model"

	"go-zero-admin-server/service/book/rpc/book"
	"go-zero-admin-server/service/book/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetTypesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTypesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTypesLogic {
	return &GetTypesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTypesLogic) GetTypes(in *book.TypesReq) (*book.TypesInfoReply, error) {
	types, total, err := l.svcCtx.TypeModel.GetTypes(int(in.Page), int(in.Limit), &model.Type{Name: in.Name})
	if err != nil {
		return nil, err
	}
	typesHandle := make([]*book.TypeInfoReply, 0)
	for _, e := range types {
		typesHandle = append(typesHandle, &book.TypeInfoReply{Id: uint64(e.ID), Name: e.Name,Intro: e.Intro})
	}
	return &book.TypesInfoReply{TypesInfo: typesHandle,Total: total}, nil
}
