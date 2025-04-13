package svc

import (
	"github.com/YShiJia/IM/apps/status/api/internal/config"
	"github.com/YShiJia/IM/apps/status/rpc/statusclient"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	*redis.Redis
	statusclient.Status
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Redis:  redis.MustNewRedis(c.Redisx),
		Status: statusclient.NewStatus(zrpc.MustNewClient(c.StatusRpc)),
		Config: c,
	}
}
