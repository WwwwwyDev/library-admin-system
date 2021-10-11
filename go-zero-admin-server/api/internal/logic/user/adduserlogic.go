package user

import (
	"context"
	"github.com/gofrs/uuid"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/common/errorx"
	"go-zero-admin-server/common/util"
	"go-zero-admin-server/service/user/rpc/userclient"

	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"

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
	isExistResp, err := l.svcCtx.UserRpc.IsExistUserByUsername(l.ctx,&userclient.UsernameReq{Username: req.Username})
	if err != nil{
		return nil,errorx.NewCodeError(code.Error,err.Error())
	}
	if isExistResp.IsExist{
		return nil, errorx.NewCodeError(code.RepeatError,"用户名重复")
	}
	saltb, err := uuid.NewV4()
	salts := saltb.String()
	passwordmd5 := util.Str2Md5(req.Password + salts)
	isSuccess, err := l.svcCtx.UserRpc.AddUser(l.ctx,&userclient.UserAddReq{Username: req.Username,Password: passwordmd5,Salt: salts,Info: req.Info})
	if err != nil{
		return nil,errorx.NewCodeError(code.Error,err.Error())
	}
	if !isSuccess.IsSuccess{
		return nil, errorx.NewCodeError(code.AddError,"添加用户失败")
	}
	return &types.Reply{Code: code.Success,Msg: "添加成功"}, nil
}
