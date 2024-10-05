/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-01 13:02:26
 */

package conn

import (
	"github.com/YShiJia/IM/lib/encoder"
	"github.com/YShiJia/IM/lib/lock"
	"sync"
	"time"
)

type wsConnOption func(conn *wsConn)

// 使用protobuf对传输数据进行编码
func WithProtobufEncoder() wsConnOption {
	return func(conn *wsConn) {
		conn.encoder = encoder.NewProtobufEncoder()
	}
}

// 使用json对传输数据进行编码
func WithJsonEncoder() wsConnOption {
	return func(conn *wsConn) {
		conn.encoder = encoder.NewJsonEncoder()
	}
}

// wsConn 使用自定义锁
func WithLocker(locker sync.Locker) wsConnOption {
	return func(conn *wsConn) {
		conn.ctxLock = locker
		conn.ctxLock = locker
	}
}

// 使 wsConn 操作不加锁
func WithoutLocker() wsConnOption {
	return func(conn *wsConn) {
		conn.ctxLock = lock.NewEmptyLock()
	}
}

func WithMaxConnectionIdleTime(idle time.Duration) wsConnOption {
	return func(conn *wsConn) {
		conn.maxConnectionIdleTime = idle
	}
}
