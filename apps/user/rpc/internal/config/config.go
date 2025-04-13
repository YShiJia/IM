package config

import (
	"github.com/YShiJia/IM/pkg/email"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	Redisx redis.RedisConf

	Mysql struct {
		DataSource string
	}

	EmailConfig email.EmailConfig

	JwtAuth struct { // JWT 认证需要的密钥和过期时间配置
		AccessSecret string
		AccessExpire int64
	}
}
