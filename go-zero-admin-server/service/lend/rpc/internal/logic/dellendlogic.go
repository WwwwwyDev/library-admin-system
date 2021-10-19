package logic

import (
	"context"
	"go-zero-admin-server/service/lend/model"
	"gorm.io/gorm"

	"go-zero-admin-server/service/lend/rpc/internal/svc"
	"go-zero-admin-server/service/lend/rpc/lend"

	"github.com/tal-tech/go-zero/core/logx"
)

type DelLendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelLendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelLendLogic {
	return &DelLendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelLendLogic) DelLend(in *lend.IdReq) (*lend.IsSuccessReply, error) {
	isSuccess, err := l.svcCtx.LendModel.DelLend(&model.Lend{Model: gorm.Model{ID: uint(in.Id)}})
	if err != nil {
		return nil, err
	}
	return &lend.IsSuccessReply{IsSuccess: isSuccess}, nil
}
