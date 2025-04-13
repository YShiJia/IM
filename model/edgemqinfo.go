/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-09 21:28:52
 */

package model

import "github.com/zeromicro/go-queue/kq"

type EdgeMQInfo struct {
	//消费命令消息
	RecvCmdMsgConsumerConf kq.KqConf
	//消费待发送的消息
	RecvCommonMsgConsumerConf kq.KqConf
	//edgeserver 地址
	Address string
}
