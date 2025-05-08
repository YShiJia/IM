package main

import (
	"flag"
	"fmt"
	initialize "github.com/YShiJia/IM/apps/file/internal/init"
	log "github.com/sirupsen/logrus"

	"github.com/YShiJia/IM/apps/file/internal/config"
	"github.com/YShiJia/IM/apps/file/internal/handler"
	"github.com/YShiJia/IM/apps/file/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "./apps/file/etc/file.yaml", "the config file")

// 初始化任务
var initTasks = []func() error{
	initialize.InitLog,
	initialize.InitIMDB,
	initialize.InitMinio,
}

func init() {
	// 加载配置项
	flag.Parse()
	conf.MustLoad(*configFile, &config.Conf)
	for _, initFunc := range initTasks {
		if err := initFunc(); err != nil {
			log.Fatal(err)
		}
	}
}

func main() {

	server := rest.MustNewServer(config.Conf.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(config.Conf)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", config.Conf.Host, config.Conf.Port)
	server.Start()
}
