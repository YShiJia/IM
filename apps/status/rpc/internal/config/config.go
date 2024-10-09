package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	IMEdgePrefix string

	Redisx redis.RedisConf

	Mysql struct {
		DataSource string
	}

	JwtAuth struct { // JWT 认证需要的密钥和过期时间配置
		AccessSecret string
		AccessExpire int64
	}

	//消费来着边缘服务的消息
	SendMsgConsumerConf kq.KqConf

	RecvCmdMsgPusherConf struct {
		Brokers []string
		Suffix  string
	}
	RecvCommonMsgPusherConf struct {
		Brokers []string
		Suffix  string
	}
}
