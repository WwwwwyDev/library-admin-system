package logic

import (
	"context"
	"go-zero-admin-server/service/lend/model"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go-zero-admin-server/service/lend/rpc/internal/svc"
	"go-zero-admin-server/service/lend/rpc/lend"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetLendsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLendsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLendsLogic {
	return &GetLendsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLendsLogic) GetLends(in *lend.LendsReq) (*lend.LendsInfoReply, error) {
	lends, total, err := l.svcCtx.LendModel.GetLends(int(in.Page), int(in.Limit), &model.Lend{VipCardNumber: in.VipCardNumber,BookName: in.BookName})
	if err != nil {
		return nil, err
	}
	lendsHandle := make([]*lend.LendInfoReply, 0)
	for _, e := range lends {
		lendsHandle = append(lendsHandle, &lend.LendInfoReply{BookName: e.BookName,BookId: uint64(e.BookId),VipId: uint64(e.VipId),VipCardNumber: e.VipCardNumber,LendTime: &timestamppb.Timestamp{Seconds: e.LendTime.Unix(),Nanos: int32(e.LendTime.UnixNano())}})
	}
	return &lend.LendsInfoReply{LendsInfo: lendsHandle,Total: total}, nil
}
