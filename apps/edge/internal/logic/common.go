/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2025-05-04 17:47:05
 */

package logic

import (
	libWsconn "github.com/YShiJia/IM/lib/websocket/conn"
	"sync"
)

var (
	// 存放当前服务所有的wsConn
	GlobalConns = globalWsConn{
		wsConns: make(map[string]libWsconn.Conn),
	}
)

type globalWsConn struct {
	// uid -> conn
	wsConns map[string]libWsconn.Conn
	rwLock  sync.RWMutex
}

func (gwc *globalWsConn) Add(uid string, conn libWsconn.Conn) {
	gwc.rwLock.Lock()
	defer gwc.rwLock.Unlock()
	// 直接覆盖，不去理会原来的conn
	gwc.wsConns[uid] = conn
}

func (gwc *globalWsConn) Get(uid string) (libWsconn.Conn, bool) {
	gwc.rwLock.RLock()
	defer gwc.rwLock.RUnlock()
	conn, ok := gwc.wsConns[uid]
	return conn, ok
}

func (gwc *globalWsConn) Del(uid string) {
	gwc.rwLock.Lock()
	defer gwc.rwLock.Unlock()
	delete(gwc.wsConns, uid)
}
