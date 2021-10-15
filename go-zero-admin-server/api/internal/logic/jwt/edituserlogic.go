package jwt

import (
	"context"
	"fmt"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/common/errorx"
	"go-zero-admin-server/service/user/rpc/userclient"
	"strconv"

	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type EditUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) EditUserLogic {
	return EditUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditUserLogic) EditUser(req types.UpdateUserByJwtReq) (*types.Reply, error) {
	sprintf := fmt.Sprintf("%v", l.ctx.Value("userId"))
	userId, err := strconv.ParseInt(sprintf,10,64)
	isExistResp, err := l.svcCtx.UserRpc.IsExistUserById(l.ctx, &userclient.IdReq{Id: uint64(userId)})
	if err != nil {
		return nil, err
	}
	if !isExistResp.IsExist {
		return nil, errorx.NewCodeError(code.NoFoundError, "用户不存在")
	}
	userOld, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &userclient.IdReq{Id: uint64(userId)})
	if err != nil {
		return nil, err
	}
	isSuccessResp, err := l.svcCtx.UserRpc.UpdateUser(l.ctx, &userclient.UserUpdateReq{Id: userOld.Id, Username: userOld.Username, Password: userOld.Password, Salt: userOld.Salt,Avatar: req.Avatar,Info: req.Info})
	if err != nil {
		return nil, err
	}
	if !isSuccessResp.IsSuccess {
		return nil, errorx.NewCodeError(code.ChangeError, "修改失败")
	}
	return &types.Reply{Code: code.Success, Msg: "修改成功"}, nil
}
