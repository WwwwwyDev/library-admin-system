package logic

import (
	"context"

	"go-zero-admin-server/service/user/rpc/internal/svc"
	"go-zero-admin-server/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type IsExistUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsExistUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsExistUserByIdLogic {
	return &IsExistUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsExistUserByIdLogic) IsExistUserById(in *user.IdReq) (*user.IsExistReply, error) {
	isExist, err := l.svcCtx.UserModel.IsExistUserById(uint(in.Id))
	if err != nil {
		return nil, err
	}
	return &user.IsExistReply{IsExist: isExist}, nil
}
