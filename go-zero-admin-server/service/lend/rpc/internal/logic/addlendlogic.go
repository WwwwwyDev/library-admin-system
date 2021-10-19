package logic

import (
	"context"
	"go-zero-admin-server/service/lend/model"
	"time"

	"go-zero-admin-server/service/lend/rpc/internal/svc"
	"go-zero-admin-server/service/lend/rpc/lend"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddLendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLendLogic {
	return &AddLendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLendLogic) AddLend(in *lend.LendAddReq) (*lend.IsSuccessReply, error) {
	var _time time.Time
	if in.LendTime == nil{
		_time = time.Unix(0,0).UTC()
	}else {
		_time = time.Unix(in.LendTime.Seconds,int64(in.LendTime.Nanos)).UTC()
	}
	isSuccess, err := l.svcCtx.LendModel.AddLend(&model.Lend{BookName: in.BookName, VipCardNumber: in.VipCardNumber, VipId: uint(in.VipId), BookId: uint(in.BookId), LendTime: _time})
	if err != nil {
		return nil, err
	}
	return &lend.IsSuccessReply{IsSuccess: isSuccess}, nil
}
