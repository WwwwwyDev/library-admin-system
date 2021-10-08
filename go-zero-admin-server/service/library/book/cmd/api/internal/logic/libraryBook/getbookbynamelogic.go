package libraryBook

import (
	"context"

	"go-zero-admin-server/service/library/book/cmd/api/internal/svc"
	"go-zero-admin-server/service/library/book/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetBookByNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBookByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetBookByNameLogic {
	return GetBookByNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBookByNameLogic) GetBookByName(req types.GetBookByNameReq) (*types.Reply, error) {
	// todo: add your logic here and delete this line

	return &types.Reply{}, nil
}
