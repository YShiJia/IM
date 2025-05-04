package main

import (
	"flag"
	conf "github.com/YShiJia/IM/apps/message/api/internal/config"
	"github.com/YShiJia/IM/apps/message/api/internal/handler"
	initialize "github.com/YShiJia/IM/apps/message/api/internal/init"
	"github.com/YShiJia/IM/apps/message/api/internal/svc"
	log "github.com/sirupsen/logrus"

	zc "github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "./apps/message/api/etc/message-api.yaml", "the config file")

// 初始化任务
var initTasks = []func() error{
	initialize.InitLog,
	initialize.InitRedis,
	initialize.InitEtcd,
	initialize.InitKafka,
	initialize.InitIMDB,
	initialize.InitMinio,
	initialize.InitMessage,
}

func init() {
	// 加载配置项
	flag.Parse()
	zc.MustLoad(*configFile, &conf.Conf)
	for _, initFunc := range initTasks {
		if err := initFunc(); err != nil {
			log.Fatal(err)
		}
	}
}

func main() {

	server := rest.MustNewServer(conf.Conf.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(conf.Conf)
	handler.RegisterHandlers(server, ctx)

	log.Infof("Starting server at %s:%d...\n", conf.Conf.Host, conf.Conf.Port)
	server.Start()
}
