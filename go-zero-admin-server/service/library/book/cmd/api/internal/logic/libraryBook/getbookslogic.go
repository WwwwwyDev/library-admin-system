package libraryBook

import (
	"context"

	"go-zero-admin-server/service/library/book/cmd/api/internal/svc"
	"go-zero-admin-server/service/library/book/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetBooksLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBooksLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetBooksLogic {
	return GetBooksLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBooksLogic) GetBooks(req types.GetBooksReq) (*types.Reply, error) {
	// todo: add your logic here and delete this line

	return &types.Reply{}, nil
}
