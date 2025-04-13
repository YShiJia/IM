/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-01 13:02:26
 */

package conn

import (
	"github.com/YShiJia/IM/lib/encoder"
)

type WsConnOption func(conn *wsConn)

// WithProtobufEncoder 使用protobuf对传输数据进行编码
func WithProtobufEncoder() WsConnOption {
	return func(conn *wsConn) {
		conn.encoder = encoder.NewProtobufEncoder()
	}
}

// WithJsonEncoder 使用json对传输数据进行编码
func WithJsonEncoder() WsConnOption {
	return func(conn *wsConn) {
		conn.encoder = encoder.NewJsonEncoder()
	}
}
