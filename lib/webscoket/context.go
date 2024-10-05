/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-01 19:59:22
 */

package webscoket

import (
	"github.com/YShiJia/IM/lib/lock"
	"github.com/YShiJia/IM/lib/webscoket/conn"
	"math"
	"net/http"
	"sync"
)

var abortIndex int8 = math.MaxInt8 >> 1

// TODO 后续增加数据存储，路径参数存储，数据返回功能
type WebContext struct {
	store   map[string]any
	ctxLock sync.Locker

	W http.ResponseWriter
	R *http.Request

	index       int8
	Middlewares []Middleware
	s           *WsServer
}

func NewWebContext(w http.ResponseWriter, r *http.Request, s *WsServer, middlewares []Middleware) *WebContext {
	return &WebContext{
		store:       make(map[string]any),
		ctxLock:     lock.NewSpinLock(),
		W:           w,
		R:           r,
		Middlewares: middlewares,
		s:           s,
		index:       -1,
	}
}

// 继续执行下一个中间件
func (c *WebContext) Next() {
	c.index++
	for c.index < int8(len(c.Middlewares)) {
		c.Middlewares[c.index](c)
		c.index++
	}
}

// 直接跳过后续所有中间件
func (c *WebContext) Abort() {
	c.index = abortIndex
}

// 添加webscoket连接到服务器中
func (c *WebContext) AddWsConn(key string, conn conn.Conn) {
	c.s.AddConn(key, conn)
}

func (c *WebContext) Value(key string) any {
	c.ctxLock.Lock()
	defer c.ctxLock.Unlock()
	return c.store[key]
}

// 相同的key，覆盖
func (c *WebContext) Set(key string, value any) {
	c.ctxLock.Lock()
	defer c.ctxLock.Unlock()
	c.store[key] = value
}
