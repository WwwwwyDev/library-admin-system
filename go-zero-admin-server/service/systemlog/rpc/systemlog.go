package main

import (
	"flag"
	"fmt"

	"go-zero-admin-server/service/systemlog/rpc/internal/config"
	"go-zero-admin-server/service/systemlog/rpc/internal/server"
	"go-zero-admin-server/service/systemlog/rpc/internal/svc"
	"go-zero-admin-server/service/systemlog/rpc/systemlog"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/systemlog.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewSystemlogServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		systemlog.RegisterSystemlogServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
