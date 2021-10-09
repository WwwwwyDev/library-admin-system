package logic

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/gofrs/uuid"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/common/errorx"

	"go-zero-admin-server/service/admin/user/api/internal/svc"
	"go-zero-admin-server/service/admin/user/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateUserLogic {
	return UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req types.UpdateUserReq) (*types.Reply, error) {
	isExist, err := l.svcCtx.UserModel.IsExistUserByID(req.ID)
	if err != nil{
		return nil,err
	}
	if !isExist{
		return nil, errorx.NewCodeError(code.NoFoundError,"用户不存在")
	}
	userOld, err := l.svcCtx.UserModel.GetUserByID(req.ID)
	if userOld.Username!= req.Username{
		return nil, errorx.NewCodeError(code.NoChangeError,"拒绝修改用户名")
	}
	userNew := userOld
	saltb, err := uuid.NewV4()
	salts := saltb.String()
	userNew.Salt = salts
	passwordb := []byte(req.Password + salts)
	has := md5.Sum(passwordb)
	passwordmd5 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	userNew.Password = passwordmd5
	err = l.svcCtx.UserModel.UpdateUser(userNew)
	if err != nil {
		return nil, err
	}
	return &types.Reply{Code: code.Success, Msg: "ok"}, nil
}
