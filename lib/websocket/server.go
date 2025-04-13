/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-01 10:22:50
 */

package websocket

import (
	"fmt"
	"github.com/YShiJia/IM/lib/lock"
	libWsconn "github.com/YShiJia/IM/lib/websocket/conn"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
	"sync"
)

const (
	RequestUUID            = "RequestUUID"
	defaultConnPoolMaxSize = 500
)

type WebServer interface {
	http.Handler
	Register(method, path string, handler HandleFunc)
	AddMiddleWare(middleware ...Middleware)
	Start() error
	Stop() error
}

var _ WebServer = (*wsServer)(nil)

// websocket server
type wsServer struct {
	// 监听地址
	listenOn string

	// 路由树
	routes Route

	// 记录当前server有多少连接
	connLocker      sync.Locker
	conns           map[string]libWsconn.Conn
	connPoolMaxSize int

	// server 范围的中间件
	Middlewares []Middleware

	// http 升级为 WebSocket 的配置
	upgrader *websocket.Upgrader
}

func NewWsServer(ListenOn string, opts ...Option) WebServer {
	wss := &wsServer{
		listenOn:        ListenOn,
		routes:          NewRouteByMap(),
		connLocker:      lock.NewSpinLock(),
		conns:           make(map[string]libWsconn.Conn),
		connPoolMaxSize: defaultConnPoolMaxSize,
		Middlewares:     make([]Middleware, 0),
		upgrader: &websocket.Upgrader{
			//允许跨域,TODO 做一个option回调
			CheckOrigin: func(r *http.Request) bool {
				return true
			}},
	}
	for _, opt := range opts {
		opt(wss)
	}
	return wss
}

func (w *wsServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//判断对应业务逻辑是否存在
	method := request.Method
	path := request.URL.Path
	handleFunc, ok := w.routes.Get(method, path)
	if !ok {
		writer.WriteHeader(http.StatusNotFound)
	}

	// 创建 webContext
	middlewares := []Middleware{entryLogic, w.createWsConn} // 服务默认middleware
	middlewares = append(middlewares, w.Middlewares...)     // 用户自定义middlewares
	middlewares = append(middlewares, handleFunc)           // 用户在该路径上定义的业务逻辑
	webContext := NewWebContext(writer, request, w, middlewares)

	//执行具体逻辑
	webContext.Next()
}

// 服务入口逻辑
func entryLogic(ctx *WebContext) {
	uid, err := uuid.NewUUID()
	if err != nil {
		ctx.W.WriteHeader(http.StatusInternalServerError)
		ctx.W.Write([]byte(fmt.Sprintf("general request uuid error:%v", err)))
		return
	}
	ctx.Set(RequestUUID, uid.String())
	ctx.Next()
	log.Infof("request[%s] end", uid.String())
}

// 创建连接
func (w *wsServer) createWsConn(ctx *WebContext) {
	var respHeader http.Header
	if protocol := ctx.R.Header.Get("Sec-Websocket-Protocol"); protocol != "" {
		respHeader = http.Header{
			"Sec-Websocket-Protocol": []string{protocol},
		}
	}

	// 升级协议后，请求方就会建立好连接
	conn, err := w.upgrader.Upgrade(ctx.W, ctx.R, respHeader)
	if err != nil {
		log.Infof("create ws conn err: %v", err)
		ctx.SetError(fmt.Errorf("create ws conn err: %v", err))
		ctx.Abort()
	}

	wsConn, err := libWsconn.NewWsConn(conn)
	if err != nil {
		log.Infof("create webScoket connect err: %v", err)
		ctx.SetError(fmt.Errorf("create webScoket connect err: %v", err))
		return
	}

	reqUUID := ctx.Value(RequestUUID)
	w.addConn(reqUUID, wsConn)
	ctx.wsConn = wsConn

	// 执行用户自定义逻辑
	ctx.Next()

	// 逻辑执行完成之后，关闭该conn
	wsConn.Close()
	w.RemoveConn(reqUUID)
}

// 这个注册的handleFunc是如何连接，连接之后返回conn
func (w *wsServer) Register(method, path string, handler HandleFunc) {
	w.routes.Register(method, path, handler)
}

func (w *wsServer) AddMiddleWare(middleware ...Middleware) {
	w.Middlewares = append(w.Middlewares, middleware...)
}

func (w *wsServer) Start() error {
	return http.ListenAndServe(w.listenOn, w)
}

func (w *wsServer) Stop() error {
	//关闭所有连接
	w.connLocker.Lock()
	defer w.connLocker.Unlock()

	for _, c := range w.conns {
		if err := c.Close(); err != nil {
			return err
		}
	}
	return nil
}

func (w *wsServer) addConn(key string, conn libWsconn.Conn) {
	w.connLocker.Lock()
	defer w.connLocker.Unlock()

	if conn, ok := w.conns[key]; ok {
		// 关闭原来的连接
		conn.Close()
	}
	w.conns[key] = conn
}

func (w *wsServer) getConn(key string) (libWsconn.Conn, bool) {
	w.connLocker.Lock()
	defer w.connLocker.Unlock()
	conn, ok := w.conns[key]
	return conn, ok
}

func (w *wsServer) RemoveConn(key string) {
	w.connLocker.Lock()
	defer w.connLocker.Unlock()
	//获取连接
	conn, ok := w.conns[key]
	if !ok {
		return
	}
	//关闭连接，并删除
	conn.Close()
	delete(w.conns, key)
}
