package logic

import (
	"context"
	"go-zero-admin-server/service/user/rpc/internal/svc"
	"go-zero-admin-server/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type DelSomeUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelSomeUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelSomeUserLogic {
	return &DelSomeUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelSomeUserLogic) DelSomeUser(in *user.IdsReq) (*user.IsSuccessReply, error) {
	ids := make([]uint, 0)
	for _,e := range in.Ids{
		ids = append(ids,uint(e.Id))
	}
	isSuccess, err := l.svcCtx.UserModel.DelSomeUser(ids)
	if err != nil {
		return nil, err
	}
	return &user.IsSuccessReply{IsSuccess: isSuccess}, nil
}
