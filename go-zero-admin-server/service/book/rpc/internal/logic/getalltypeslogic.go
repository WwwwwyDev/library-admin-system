package logic

import (
	"context"

	"go-zero-admin-server/service/book/rpc/book"
	"go-zero-admin-server/service/book/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAllTypesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllTypesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllTypesLogic {
	return &GetAllTypesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllTypesLogic) GetAllTypes(in *book.EmptyReq) (*book.TypesInfoReply, error) {
	types,  err := l.svcCtx.TypeModel.GetAllTypes()
	if err != nil {
		return nil, err
	}
	typesHandle := make([]*book.TypeInfoReply, 0)
	for _, e := range types {
		typesHandle = append(typesHandle, &book.TypeInfoReply{Id: uint64(e.ID), Name: e.Name,Intro: e.Intro})
	}
	return &book.TypesInfoReply{TypesInfo: typesHandle}, nil
}
