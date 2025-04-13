package websocket

import "github.com/gorilla/websocket"

// HandleFunc 处理逻辑
type HandleFunc func(ctx *WebContext)

// Middleware 中间件
type Middleware = HandleFunc

// Option 创建serve的option
type Option func(server *wsServer)

// 正常关闭连接的状态码
var wsCloseNormalCodes = []int{websocket.CloseNormalClosure, websocket.CloseGoingAway}
