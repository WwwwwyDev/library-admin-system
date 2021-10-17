package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserRpc zrpc.RpcClientConf
	BookRpc zrpc.RpcClientConf
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
	Qiniu struct{
		AccessKey string
		SecretKey string
		Expires uint64
		Bucket string
	}
}
