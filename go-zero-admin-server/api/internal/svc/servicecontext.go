package svc

import (
	"github.com/go-redis/redis"
	"github.com/tal-tech/go-zero/zrpc"
	"go-zero-admin-server/api/internal/config"
	"go-zero-admin-server/common/core"
	"go-zero-admin-server/service/book/rpc/bookclient"
	"go-zero-admin-server/service/lend/rpc/lendclient"
	"go-zero-admin-server/service/systemlog/rpc/systemlogclient"
	"go-zero-admin-server/service/user/rpc/userclient"
	"go-zero-admin-server/service/vip/rpc/vipclient"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
	BookRpc bookclient.Book
	VipRpc  vipclient.Vip
	LendRpc  lendclient.Lend
	SystemlogRpc systemlogclient.Systemlog
	Redis   *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	redis := core.NewRedis(c.Redis.Host, c.Redis.Password, c.Redis.Db)
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		BookRpc: bookclient.NewBook(zrpc.MustNewClient(c.BookRpc)),
		VipRpc:  vipclient.NewVip(zrpc.MustNewClient(c.VipRpc)),
		LendRpc: lendclient.NewLend(zrpc.MustNewClient(c.LendRpc)),
		SystemlogRpc: systemlogclient.NewSystemlog(zrpc.MustNewClient(c.SystemlogRpc)),
		Redis:   redis,
	}
}
