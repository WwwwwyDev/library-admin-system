package logic

import (
	"context"

	"go-zero-admin-server/service/lend/rpc/internal/svc"
	"go-zero-admin-server/service/lend/rpc/lend"

	"github.com/tal-tech/go-zero/core/logx"
)

type IsExistLendByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsExistLendByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsExistLendByIdLogic {
	return &IsExistLendByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsExistLendByIdLogic) IsExistLendById(in *lend.IdReq) (*lend.IsExistReply, error) {
	isExist, err := l.svcCtx.LendModel.IsExistLendById(in.Id)
	if err != nil {
		return nil, err
	}
	return &lend.IsExistReply{IsExist: isExist}, nil
}
