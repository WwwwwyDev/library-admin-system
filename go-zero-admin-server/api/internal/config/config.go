package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserRpc zrpc.RpcClientConf
	CacheRedis cache.CacheConf
	Auth  struct {
		AccessSecret string
		AccessExpire int64
	}
	Redis struct{
		Host string
		Password string
		Db int
	}
}
