/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-02 15:08:35
 */

package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/service"
)

type Config struct {
	service.ServiceConf
	ListenOn   string
	IPv4Prefix string

	Etcd discov.EtcdConf

	Auth struct { // JWT 认证需要的密钥和过期时间配置
		AccessSecret string
		AccessExpire int64
	}

	//生产消息
	SendMsgPusherConf struct {
		Brokers []string
		Topic   string
	}
	//消费命令消息
	RecvCmdMsgConsumerConf kq.KqConf
	//消费待发送的消息
	RecvCommonMsgConsumerConf kq.KqConf
}
