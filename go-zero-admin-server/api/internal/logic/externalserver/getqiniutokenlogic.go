package externalserver

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go-zero-admin-server/common/code"

	"github.com/tal-tech/go-zero/core/logx"
	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"
)

type GetQiniuTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetQiniuTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetQiniuTokenLogic {
	return GetQiniuTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetQiniuTokenLogic) GetQiniuToken() (*types.Reply, error) {
	bucket:=l.svcCtx.Config.Qiniu.Bucket
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = l.svcCtx.Config.Qiniu.Expires
	accessKey := l.svcCtx.Config.Qiniu.AccessKey
	secretKey := l.svcCtx.Config.Qiniu.SecretKey
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	return &types.Reply{Code: code.Success, Data: map[string]interface{}{"token": upToken}, Msg: "获取成功"}, nil
}
