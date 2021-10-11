package svc

import (
	"github.com/go-redis/redis"
	"github.com/tal-tech/go-zero/zrpc"
	"go-zero-admin-server/api/internal/config"
	"go-zero-admin-server/common/core"
	"go-zero-admin-server/service/user/rpc/userclient"
)

type ServiceContext struct {
	Config config.Config
	UserRpc userclient.User
	Redis *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	redis := core.NewRedis(c.Redis.Host, c.Redis.Password, c.Redis.Db)
	return &ServiceContext{
		Config: c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		Redis: redis,
	}
}
