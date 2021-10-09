package logic

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/gofrs/uuid"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/common/errorx"
	"go-zero-admin-server/service/admin/user/model"

	"go-zero-admin-server/service/admin/user/api/internal/svc"
	"go-zero-admin-server/service/admin/user/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddUserLogic {
	return AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req types.AddUserReq) (*types.Reply, error) {
	isExist, err := l.svcCtx.UserModel.IsExistUserByUsername(req.Username)
	if err != nil{
		return nil,err
	}
	if isExist{
		return nil, errorx.NewCodeError(code.RepeatError,"用户名重复")
	}
	saltb, err := uuid.NewV4()
	salts := saltb.String()
	passwordb := []byte(req.Password + salts)
	has := md5.Sum(passwordb)
	passwordmd5 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	err = l.svcCtx.UserModel.AddUser(&model.User{Username: req.Username, Password: passwordmd5,Salt: salts,Info: req.Info})
	if err != nil{
		return nil,err
	}
	return &types.Reply{Code: code.Success,Msg: "ok"}, nil
}
