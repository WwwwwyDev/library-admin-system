package logic

import (
	"context"
	"go-zero-admin-server/service/vip/model"

	"go-zero-admin-server/service/vip/rpc/internal/svc"
	"go-zero-admin-server/service/vip/rpc/vip"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetVipsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVipsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVipsLogic {
	return &GetVipsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVipsLogic) GetVips(in *vip.VipsReq) (*vip.VipsInfoReply, error) {
	vips, total, err := l.svcCtx.VipModel.GetVips(int(in.Page), int(in.Limit), &model.Vip{CardNumber: in.CardNumber,Name: in.Name})
	if err != nil {
		return nil, err
	}
	vipsHandle := make([]*vip.VipInfoReply, 0)
	for _, e := range vips {
		vipsHandle = append(vipsHandle, &vip.VipInfoReply{Id: uint64(e.ID),CardNumber: e.CardNumber,Name: e.Name,Info: e.Info})
	}
	return &vip.VipsInfoReply{VipsInfo: vipsHandle,Total: total}, nil
}
