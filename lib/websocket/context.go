/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-01 19:59:22
 */

package websocket

import (
	"github.com/YShiJia/IM/lib/lock"
	libWsconn "github.com/YShiJia/IM/lib/websocket/conn"
	"math"
	"net/http"
	"sync"
)

var abortIndex int8 = math.MaxInt8 >> 1

// TODO 后续增加数据存储，路径参数存储，数据返回功能
type WebContext struct {
	// 当前连接请求范围内的数据存储
	store   map[string]string
	ctxLock sync.Locker

	// 请求连接
	W      http.ResponseWriter
	R      *http.Request
	wsConn libWsconn.Conn

	// AOP
	index       int8
	middlewares []Middleware

	// server
	s WebServer
	// 连接请求中困难出现的错误
	err error
}

func NewWebContext(w http.ResponseWriter, r *http.Request, s WebServer, middlewares []Middleware) *WebContext {
	return &WebContext{
		store:       make(map[string]string),
		ctxLock:     lock.NewSpinLock(),
		W:           w,
		R:           r,
		middlewares: middlewares,
		s:           s,
		index:       -1,
	}
}

// Next 继续执行下一个中间件
func (c *WebContext) Next() {
	c.index++
	for c.index < int8(len(c.middlewares)) {
		c.middlewares[c.index](c)
		c.index++
	}
}

// 直接跳过后续所有中间件
func (c *WebContext) Abort() {
	c.index = abortIndex
}

func (c *WebContext) Conn() libWsconn.Conn {
	return c.wsConn
}

func (c *WebContext) Value(key string) string {
	c.ctxLock.Lock()
	defer c.ctxLock.Unlock()
	return c.store[key]
}

// Set 相同的key，覆盖
func (c *WebContext) Set(key string, value string) {
	c.ctxLock.Lock()
	defer c.ctxLock.Unlock()
	c.store[key] = value
}

func (c *WebContext) SetError(err error) {
	c.err = err
}

func (c *WebContext) Error() error {
	return c.err
}
