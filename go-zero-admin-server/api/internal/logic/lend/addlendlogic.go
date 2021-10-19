package lend

import (
	"context"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/common/errorx"
	"go-zero-admin-server/service/book/rpc/bookclient"
	"go-zero-admin-server/service/lend/rpc/lendclient"
	"go-zero-admin-server/service/vip/rpc/vipclient"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"

	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddLendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLendLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddLendLogic {
	return AddLendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLendLogic) AddLend(req types.AddLendReq) (*types.Reply, error) {
	isExistVip, err := l.svcCtx.VipRpc.IsExistVipById(l.ctx, &vipclient.IdReq{Id: uint64(req.VipId)})
	if err != nil {
		return nil, err
	}
	if !isExistVip.IsExist{
		return nil, errorx.NewCodeError(code.NoFoundError,"没有该会员")
	}
	isExistBook, err := l.svcCtx.BookRpc.IsExistBookById(l.ctx, &bookclient.IdReq{Id: uint64(req.BookId)})
	if err != nil {
		return nil, err
	}
	if !isExistBook.IsExist{
		return nil, errorx.NewCodeError(code.NoFoundError,"没有该图书")
	}
	isLend, err := l.svcCtx.LendRpc.IsLend(l.ctx, &lendclient.BookIdReq{BookId: uint64(req.BookId)})
	if err != nil {
		return nil, err
	}
	if isLend.IsLend{
		return nil, errorx.NewCodeError(code.RepeatError,"该书已被借阅")
	}
	_time := time.Now()
	isSuccess, err := l.svcCtx.LendRpc.AddLend(l.ctx, &lendclient.LendAddReq{VipId: uint64(req.VipId), VipCardNumber: req.VipCardNumber, BookName: req.BookName, BookId: uint64(req.BookId), LendTime: &timestamppb.Timestamp{Seconds: _time.Unix(), Nanos: int32(_time.UnixNano())}})
	if err != nil {
		return nil, err
	}
	if !isSuccess.IsSuccess{
		return nil, errorx.NewCodeError(code.AddError,"借阅失败")
	}
	return &types.Reply{Code: code.Success,Msg: "借阅成功"}, nil
}
