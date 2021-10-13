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
	if len(req.Password) < 6 {
		return nil,errorx.NewCodeError(code.ParameterError, "密码非法")
	}
	isExistResp, err := l.svcCtx.UserRpc.IsExistUserById(l.ctx, &userclient.IdReq{Id: uint64(req.Id)})
	if err != nil {
		return nil, err
	}
	if !isExistResp.IsExist {
		return nil, errorx.NewCodeError(code.NoFoundError, "用户不存在")
	}
	userOld, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &userclient.IdReq{Id: uint64(req.Id)})
	if err != nil {
		return nil, err
	}
	userNew := userOld
	saltb, err := uuid.NewV4()
	salts := saltb.String()
	userNew.Salt = salts
	userNew.Password = util.Str2Md5(req.Password + salts)
	isSuccessResp, err := l.svcCtx.UserRpc.UpdateUser(l.ctx, &userclient.UserUpdateReq{Id: userNew.Id, Username: userNew.Username, Password: userNew.Password, Salt: userNew.Salt,Avatar: req.Avatar,Info: req.Info})
	if err != nil {
		return nil, err
	}
	if !isSuccessResp.IsSuccess {
		return nil, errorx.NewCodeError(code.ChangeError, "修改用户失败")
	}
	return &types.Reply{Code: code.Success, Msg: "修改用户成功"}, nil
}
