/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-01 10:34:44
 */

package conn

import (
	"context"
	"errors"
	"github.com/YShiJia/IM/lib/encoder"
	"github.com/YShiJia/IM/lib/lock"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
)

var ErrConnectionClosed = errors.New("connection closed")

// 默认超时时间一分钟
var defaultMaxConnectionIdleTime = time.Minute

// Conn 封装websocket连接
type Conn interface {
	Send(ctx context.Context, data any) error
	Receive(ctx context.Context, data any) (err error)
	GetValue(key string) (value any, exit bool)
	SetValue(key string, value any) error
	Close() error
}

var _ Conn = (*wsConn)(nil)

// 支持并发安全的websocket连接
// 客户端主动进行心跳检测，超时则断开
type wsConn struct {
	//连接对象的上下文
	ctx     map[string]any
	ctxLock sync.Locker

	//封装websocket连接
	ws *websocket.Conn
	// "github.com/gorilla/websocket"支持并发单线程读写
	connReadLock  sync.Locker
	connWriteLock sync.Locker

	//传输数据编码器
	encoder encoder.Encoder

	//最后一次使用该连接的时间
	latestUseTime time.Time
	//最大空闲时间，超过该时间连接未通讯则说明连接失效
	maxConnectionIdleTime time.Duration

	//连接是否失效
	closed bool
}

func NewWsConn(upgrader *websocket.Upgrader, w http.ResponseWriter, r *http.Request, opts ...wsConnOption) (*wsConn, error) {

	var respHeader http.Header
	if protocol := r.Header.Get("Sec-Websocket-Protocol"); protocol != "" {
		respHeader = http.Header{
			"Sec-Websocket-Protocol": []string{protocol},
		}
	}

	ws, err := upgrader.Upgrade(w, r, respHeader)

	if err != nil {
		return nil, err
	}

	conn := &wsConn{
		ctx:           make(map[string]any),
		ctxLock:       lock.NewSpinLock(),
		ws:            ws,
		connReadLock:  lock.NewSpinLock(),
		connWriteLock: lock.NewSpinLock(),
		//encoder:               encoder.NewProtobufEncoder(),
		encoder:               encoder.NewJsonEncoder(),
		latestUseTime:         time.Now(),
		maxConnectionIdleTime: defaultMaxConnectionIdleTime,
		closed:                false,
	}

	for _, opt := range opts {
		opt(conn)
	}

	return conn, nil
}

func (w *wsConn) GetValue(key string) (value any, exit bool) {
	w.ctxLock.Lock()
	val, ok := w.ctx[key]
	w.ctxLock.Unlock()

	return val, ok
}

func (w *wsConn) SetValue(key string, value any) error {
	w.ctxLock.Lock()
	w.ctx[key] = value
	w.ctxLock.Unlock()

	return nil
}

func (w *wsConn) Send(ctx context.Context, data any) error {
	//先编码，再发送数据
	var msg []byte
	var err error
	if msg, err = w.encoder.Encode(data); err != nil {
		return err
	}

	w.connWriteLock.Lock()
	defer w.connWriteLock.Unlock()
	//超时或者连接正常关闭
	if time.Since(w.latestUseTime) > w.maxConnectionIdleTime || w.closed {
		return ErrConnectionClosed
	}

	//统一使用二进制消息，实际处理逻辑使用自定义消息类型
	if err = w.ws.WriteMessage(websocket.BinaryMessage, msg); err != nil && websocket.IsCloseError(err) {
		//连接异常关闭
		return ErrConnectionClosed
	}

	w.latestUseTime = time.Now()
	return nil
}

func (w *wsConn) Receive(ctx context.Context, data any) error {

	w.connReadLock.Lock()
	// 超时或者连接正常关闭
	if time.Since(w.latestUseTime) > w.maxConnectionIdleTime || w.closed {
		w.connReadLock.Unlock()
		return ErrConnectionClosed
	}
	_, pbData, err := w.ws.ReadMessage()
	w.latestUseTime = time.Now()

	w.connReadLock.Unlock()

	// 连接异常关闭
	if err != nil {
		if websocket.IsCloseError(err) {
			return ErrConnectionClosed
		}
		return err
	}
	return w.encoder.Decode(pbData, data)
}

func (w *wsConn) Close() error {
	w.ctxLock.Lock()
	defer w.ctxLock.Unlock()
	//保证关闭幂等性
	if w.closed {
		return nil
	}
	w.closed = true
	return w.ws.Close()
}
