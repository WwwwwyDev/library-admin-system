// Code generated by goctl. DO NOT EDIT!
// Source: systemlog.proto

package server

import (
	"context"

	"go-zero-admin-server/service/systemlog/rpc/internal/logic"
	"go-zero-admin-server/service/systemlog/rpc/internal/svc"
	"go-zero-admin-server/service/systemlog/rpc/systemlog"
)

type SystemlogServer struct {
	svcCtx *svc.ServiceContext
}

func NewSystemlogServer(svcCtx *svc.ServiceContext) *SystemlogServer {
	return &SystemlogServer{
		svcCtx: svcCtx,
	}
}

func (s *SystemlogServer) GetLoginLogs(ctx context.Context, in *systemlog.LoginLogsReq) (*systemlog.LoginLogsInfoReply, error) {
	l := logic.NewGetLoginLogsLogic(ctx, s.svcCtx)
	return l.GetLoginLogs(in)
}

func (s *SystemlogServer) AddLoginLog(ctx context.Context, in *systemlog.LoginLogAddReq) (*systemlog.IsSuccessReply, error) {
	l := logic.NewAddLoginLogLogic(ctx, s.svcCtx)
	return l.AddLoginLog(in)
}

func (s *SystemlogServer) DelLoginLog(ctx context.Context, in *systemlog.EmptyReq) (*systemlog.IsSuccessReply, error) {
	l := logic.NewDelLoginLogLogic(ctx, s.svcCtx)
	return l.DelLoginLog(in)
}
