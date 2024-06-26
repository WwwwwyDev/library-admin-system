package config

import "github.com/tal-tech/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	Postgresql struct {
		User     string
		Password string
		Host     string
		Port     int
		Name     string
	}
}
