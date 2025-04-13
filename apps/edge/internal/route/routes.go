/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-02 18:58:09
 */

package route

import (
	"github.com/YShiJia/IM/apps/edge/internal/api"
	"github.com/YShiJia/IM/lib/websocket"
	"net/http"
)

type router struct {
	method  string
	path    string
	handler websocket.HandleFunc
}

var routers = []router{
	{
		method:  http.MethodGet,
		path:    "/wsconn",
		handler: api.MessageSRHandler,
	},
}

func RegisterHandlers(webServer websocket.WebServer) {
	for _, router := range routers {
		webServer.Register(router.method, router.path, router.handler)
	}
}
