package main

import (
	"flag"
	"fmt"

	"github.com/YShiJia/IM/apps/status/rpc/internal/config"
	"github.com/YShiJia/IM/apps/status/rpc/internal/server"
	"github.com/YShiJia/IM/apps/status/rpc/internal/svc"
	"github.com/YShiJia/IM/apps/status/rpc/statusmodel"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/status.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	// 开启边缘服务动态监听
	ctx.ListenEdgeServer()
	// 开启边缘服务消息监听
	ctx.ConsumeEdgeMsg()

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		statusmodel.RegisterStatusServer(grpcServer, server.NewStatusServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
