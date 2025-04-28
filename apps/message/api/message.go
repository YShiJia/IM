package main

import (
	"flag"
	"fmt"
	conf "github.com/YShiJia/IM/apps/message/api/internal/config"
	"github.com/YShiJia/IM/apps/message/api/internal/handler"
	"github.com/YShiJia/IM/apps/message/api/internal/svc"

	zc "github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "./apps/message/api/etc/message-api.yaml", "the config file")

func main() {
	flag.Parse()

	zc.MustLoad(*configFile, &conf.Conf)

	server := rest.MustNewServer(conf.Conf.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(conf.Conf)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", conf.Conf.Host, conf.Conf.Port)
	server.Start()
}
