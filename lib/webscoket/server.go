/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-01 10:22:50
 */

package webscoket

import (
	"fmt"
	"github.com/YShiJia/IM/lib/lock"
	"github.com/YShiJia/IM/lib/webscoket/conn"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
)

var defaultMaxConnectionIdleTime = time.Minute

var WsConnectionIdentity = "WsIdentity"

type WebServer interface {
	Register(method, path string) error
	Start() error
	Stop() error
}

type ConnHandler func(conn.Conn)

// TODO 现在只支持局部后续可以优化一下middleware的灵活度
type WsServer struct {
	//监听ip+端口
	listenOn string

	//路由树
	routes Route

	//WebSocket升级器
	upgrader websocket.Upgrader

	//连接存储模块
	connLocker sync.Locker
	conns      map[string]conn.Conn

	connHandler ConnHandler

	// 现在暂时使用切片实现middleware，使用gin的middleware执行流程方案
	// TODO 后续将其改造成一个接口，使用不同的middleware执行流程，grpc拦截器闭包方案也可以试一下
	Middlewares []Middleware
}

func NewWsServer(ListenOn string) *WsServer {
	return &WsServer{
		listenOn: ListenOn,
		routes:   NewRouteByMap(),
		upgrader: websocket.Upgrader{
			//允许跨域,TODO 做一个option回调
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		conns:       make(map[string]conn.Conn),
		connLocker:  lock.NewSpinLock(),
		Middlewares: make([]Middleware, 0),
	}
}

func (w *WsServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//判断对应业务逻辑是否存在
	method := request.Method
	path := request.URL.Path
	handleFunc, ok := w.routes.Get(method, path)
	if !ok {
		writer.WriteHeader(http.StatusNotFound)
	}

	webContext := NewWebContext(writer, request, w, append(w.Middlewares, handleFunc))

	//执行AOP
	// TODO 这里是全局粒度的AOP，后续会优化到节点粒度的AOP方案
	webContext.Next()
}

// 这个注册的handleFunc是如何连接，连接之后返回conn
func (w *WsServer) Register(method, path string) error {
	if ok := w.routes.Register(method, path, handleConnectionRequest); !ok {
		return ErrPathNodeExist
	}
	return nil
}

func (w *WsServer) WithConnHandler(handler ConnHandler) {
	w.connHandler = handler
}

func (w *WsServer) MiddleWare(middleware ...Middleware) error {
	w.Middlewares = append(w.Middlewares, middleware...)
	return nil
}

func (w *WsServer) Start() error {
	return http.ListenAndServe(w.listenOn, w)
}

func (w *WsServer) Stop() error {
	//关闭所有连接
	w.connLocker.Lock()
	defer w.connLocker.Unlock()
	//获取连接
	for _, c := range w.conns {
		c.Close()
	}
	return nil
}

func (w *WsServer) AddConn(key string, conn conn.Conn) {
	w.connLocker.Lock()
	defer w.connLocker.Unlock()
	//判断连接是否存在
	if conn, ok := w.conns[key]; ok {
		conn.Close()
	}
	w.conns[key] = conn

	//使用自定义业务逻辑
	w.connHandler(conn)
}

func (w *WsServer) GetConn(key string) conn.Conn {
	w.connLocker.Lock()
	defer w.connLocker.Unlock()
	return w.conns[key]
}

func (w *WsServer) RemoveConn(key string) {
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

func handleConnectionRequest(ctx *WebContext) {
	//没有认证
	id, ok := ctx.Value(WsConnectionIdentity).(string)
	if !ok {
		ctx.W.WriteHeader(http.StatusUnauthorized)
		ctx.W.Write([]byte(fmt.Sprintf("[err]: 认证标识 %s 异常", WsConnectionIdentity)))
		return
	}

	wsConn, err := conn.NewWsConn(&ctx.s.upgrader, ctx.W, ctx.R)
	if err != nil {
		return
	}

	if err = wsConn.SetValue(WsConnectionIdentity, id); err != nil {
		ctx.W.WriteHeader(http.StatusInternalServerError)
		ctx.W.Write([]byte(fmt.Sprintf("[err]: %s", err.Error())))
		return
	}

	//添加到当前服务器中
	ctx.s.AddConn(id, wsConn)
}
