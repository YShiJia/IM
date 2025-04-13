/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-02 18:48:18
 */

package internal

import (
	"github.com/YShiJia/IM/apps/edge/internal/middleware"
	"github.com/YShiJia/IM/apps/edge/internal/route"
	libWs "github.com/YShiJia/IM/lib/websocket"
	log "github.com/sirupsen/logrus"
)

type EdgeServer struct {
	// 用于处理ws连接
	WsServer libWs.WebServer
}

var edgeServerEntity *EdgeServer

func NewEdgeServer(ListenOn string) *EdgeServer {
	wsServer := libWs.NewWsServer(ListenOn)
	// 中间件
	wsServer.AddMiddleWare(middleware.JwtAuthorize)
	// 注册路由
	route.RegisterHandlers(wsServer)

	return &EdgeServer{
		WsServer: wsServer,
	}
}

func (e *EdgeServer) Start() error {
	if err := e.WsServer.Start(); err != nil {
		log.Infof("wsServer start failed: %v", err)
		return err
	}
	return nil
}

func (e *EdgeServer) Stop() error {
	if err := e.WsServer.Stop(); err != nil {
		log.Infof("wsServer stop failed: %v", err)
		return err
	}
	return nil
}
