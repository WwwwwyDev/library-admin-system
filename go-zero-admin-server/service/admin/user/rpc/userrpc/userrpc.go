// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package userrpc

import (
	"context"

	"go-zero-admin-server/service/admin/user/rpc/user"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	UserReply   = user.UserReply
	IDReq       = user.IDReq
	UsernameReq = user.UsernameReq
	Role        = user.Role

	UserRpc interface {
		GetUserById(ctx context.Context, in *IDReq) (*UserReply, error)
		GetUserByUsername(ctx context.Context, in *UsernameReq) (*UserReply, error)
	}

	defaultUserRpc struct {
		cli zrpc.Client
	}
)

func NewUserRpc(cli zrpc.Client) UserRpc {
	return &defaultUserRpc{
		cli: cli,
	}
}

func (m *defaultUserRpc) GetUserById(ctx context.Context, in *IDReq) (*UserReply, error) {
	client := user.NewUserRpcClient(m.cli.Conn())
	return client.GetUserById(ctx, in)
}

func (m *defaultUserRpc) GetUserByUsername(ctx context.Context, in *UsernameReq) (*UserReply, error) {
	client := user.NewUserRpcClient(m.cli.Conn())
	return client.GetUserByUsername(ctx, in)
}
