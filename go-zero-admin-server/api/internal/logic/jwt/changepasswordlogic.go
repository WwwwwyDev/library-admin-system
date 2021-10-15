package jwt

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/common/errorx"
	"go-zero-admin-server/common/util"
	"go-zero-admin-server/service/user/rpc/userclient"
	"strconv"

	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ChangePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) ChangePasswordLogic {
	return ChangePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePasswordLogic) ChangePassword(req types.ChangePasswordByJwtReq) (*types.Reply, error) {
	if req.NewPassword != req.NewPasswordAgain{
		return nil, errorx.NewCodeError(code.ParameterError,"两次输入密码不一致")
	}
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
	oldPasswordReq := util.Str2Md5(req.OldPassword + userOld.Salt)
	oldPassword := util.Str2Md5(userOld.Password+userOld.Salt)
	if oldPasswordReq != oldPassword{
		return nil, errorx.NewCodeError(code.ParameterError,"两次输入密码不一致")
	}
	saltb, err := uuid.NewV4()
	salts := saltb.String()
	newPassword := util.Str2Md5(req.NewPassword + salts)
	isSuccessResp, err := l.svcCtx.UserRpc.UpdateUser(l.ctx, &userclient.UserUpdateReq{Id: userOld.Id, Username: userOld.Username, Password: newPassword, Salt: salts,Avatar: userOld.Avatar,Info: userOld.Info})
	if err != nil {
		return nil, err
	}
	if !isSuccessResp.IsSuccess {
		return nil, errorx.NewCodeError(code.ChangeError, "修改用户失败")
	}
	return &types.Reply{}, nil
}
