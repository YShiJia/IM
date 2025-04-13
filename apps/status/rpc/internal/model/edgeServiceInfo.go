/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-07 19:45:40
 */

package model

import "github.com/zeromicro/go-queue/kq"

type EdgeServiceInfo struct {
	RecvCmdMsgPusher    *kq.Pusher
	RecvCommonMsgPusher *kq.Pusher
	Address             string
}
