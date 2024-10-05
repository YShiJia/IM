/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-02 18:58:09
 */

package route

import (
	"github.com/YShiJia/IM/lib/webscoket"
	"net/http"
)

func RegisterHandlers(wsServer *webscoket.WsServer) {
	wsServer.Register(http.MethodGet, "/wsConn")
}
