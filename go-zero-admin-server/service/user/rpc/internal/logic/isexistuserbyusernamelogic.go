package logic

import (
	"context"

	"go-zero-admin-server/service/user/rpc/internal/svc"
	"go-zero-admin-server/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type IsExistUserByUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsExistUserByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsExistUserByUsernameLogic {
	return &IsExistUserByUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsExistUserByUsernameLogic) IsExistUserByUsername(in *user.UsernameReq) (*user.IsExistReply, error) {
	isExist, err := l.svcCtx.UserModel.IsExistUserByUsername(in.Username)
	if err != nil {
		return nil, err
	}
	return &user.IsExistReply{IsExist: isExist}, nil
}
