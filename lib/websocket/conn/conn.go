/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-01 10:34:44
 */

package conn

import (
	"errors"
	"github.com/YShiJia/IM/lib/encoder"
	"github.com/YShiJia/IM/lib/lock"
	"github.com/gorilla/websocket"
	"sync"
	"sync/atomic"
	"time"
)

var ErrConnectionClosed = errors.New("connection closed")

// 默认超时时间一分钟
const (
	defaultMaxConnectionIdleTime = time.Minute
)

// Conn 封装websocket连接
type Conn interface {
	Send(int, []byte) error
	Receive() (int, []byte, error)
	Close() error
}

var _ Conn = (*wsConn)(nil)

// 支持并发安全的websocket连接
// 客户端主动进行心跳检测，超时则断开
type wsConn struct {
	//封装websocket连接
	conn *websocket.Conn

	// "github.com/gorilla/websocket"支持同一时刻读和写,细化锁的粒度，使用两个锁
	connReadLock  sync.Locker
	connWriteLock sync.Locker

	//传输数据编码器
	encoder encoder.Encoder

	//连接是否关闭 or 失效
	closeCh chan struct{}
	closed  atomic.Bool
}

func NewWsConn(wsconn *websocket.Conn, opts ...WsConnOption) (Conn, error) {
	conn := &wsConn{
		conn:          wsconn,
		connReadLock:  lock.NewSpinLock(),
		connWriteLock: lock.NewSpinLock(),
		// 使用json编码，易于调试
		encoder: encoder.NewJsonEncoder(),
		closeCh: make(chan struct{}),
		closed:  atomic.Bool{},
	}

	for _, opt := range opts {
		opt(conn)
	}

	return conn, nil
}

// Send 懒删除
func (wsc *wsConn) Send(messageType int, data []byte) error {
	wsc.connWriteLock.Lock()
	defer wsc.connWriteLock.Unlock()

	if wsc.closed.Load() {
		wsc.Close()
		return ErrConnectionClosed
	}

	//统一使用二进制消息，实际处理逻辑使用自定义消息类型
	if err := wsc.conn.WriteMessage(messageType, data); err != nil {
		wsc.Close()
		return ErrConnectionClosed
	}

	return nil
}

func (wsc *wsConn) Receive() (int, []byte, error) {
	wsc.connReadLock.Lock()
	// 超时或者已经关闭
	if wsc.closed.Load() {
		wsc.connReadLock.Unlock()
		wsc.Close()
		return 0, nil, ErrConnectionClosed
	}
	messageType, data, err := wsc.conn.ReadMessage()

	wsc.connReadLock.Unlock()

	// 连接异常关闭
	if err != nil {
		return 0, nil, err
	}
	return messageType, data, nil
}

func (wsc *wsConn) Close() error {
	// 关闭操作幂等性
	if swapped := wsc.closed.CompareAndSwap(false, true); !swapped {
		// 说明连接已经被关闭了
		return nil
	}
	close(wsc.closeCh)

	// 发送正常关闭的消息
	return wsc.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
}

func (wsc *wsConn) CloseCh() <-chan struct{} {
	return wsc.closeCh
}
