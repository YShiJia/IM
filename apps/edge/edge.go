/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-02 15:00:28
 */

package main

import (
	"flag"
	"fmt"
	"github.com/YShiJia/IM/apps/edge/internal"
	conf "github.com/YShiJia/IM/apps/edge/internal/config"
	initialize "github.com/YShiJia/IM/apps/edge/internal/init"
	log "github.com/sirupsen/logrus"
)

var configFile = flag.String("f", "./apps/edge/etc/edge.yaml", "the config file path")

// 初始化任务
var initTasks = []func() error{
	initialize.InitLog,
	initialize.InitRedis,
	initialize.InitEtcd,
	initialize.InitKafka,
	initialize.InitEdge,
}

func init() {
	initialize.InitConfig(*configFile)
	for _, initFunc := range initTasks {
		if err := initFunc(); err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	edgeServer := internal.NewEdgeServer(fmt.Sprintf(":%d", conf.Conf.HttpPort))
	defer edgeServer.Stop()
	log.Infof("Starting server at :%d...\n", conf.Conf.HttpPort)

	if err := edgeServer.Start(); err != nil {
		log.Infof("edge server start : err %v", err)
	}
}
