package logic

import (
	"context"

	"go-zero-admin-server/service/lend/rpc/internal/svc"
	"go-zero-admin-server/service/lend/rpc/lend"

	"github.com/tal-tech/go-zero/core/logx"
)

type IsLendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsLendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsLendLogic {
	return &IsLendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsLendLogic) IsLend(in *lend.BookIdReq) (*lend.IsLendReply, error) {
	isLend, err := l.svcCtx.LendModel.IsLend(uint64(in.BookId))
	if err != nil {
		return nil, err
	}
	return &lend.IsLendReply{IsLend: isLend}, nil
}
