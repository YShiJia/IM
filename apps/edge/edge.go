/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-02 15:00:28
 */

package main

import (
	"flag"
	"github.com/YShiJia/IM/apps/edge/internal/config"
	"github.com/YShiJia/IM/apps/edge/internal/server"
	"github.com/YShiJia/IM/apps/edge/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "etc/edge.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	srvCtx := svc.NewServiceContext(c)

	edgeServer := server.NewEdgeServer(srvCtx)

	defer edgeServer.Stop()

	if err := edgeServer.Start(); err != nil {
		logx.Infof("edge server exit : err %s", err.Error())
	}

}
