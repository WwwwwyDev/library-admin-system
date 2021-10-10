package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"go-zero-admin-server/api/internal/config"
	"go-zero-admin-server/service/user/rpc/userclient"
)

type ServiceContext struct {
	Config config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
