package init

import (
	"context"
	"fmt"
	conf "github.com/YShiJia/IM/apps/edge/internal/config"
	"github.com/YShiJia/IM/apps/edge/internal/dao"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

func InitRedis() error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Conf.RedisConf.Host, conf.Conf.RedisConf.Port),
		Password: conf.Conf.RedisConf.Password, // 密码
		DB:       conf.Conf.RedisConf.Index,    // 数据库
		PoolSize: conf.Conf.RedisConf.PoolSize, // 连接池大小
	})
	ctx, cancel := context.WithTimeout(context.Background(), conf.Conf.ReqTimeOut)
	defer cancel()
	if ping := rdb.Ping(ctx); ping.Err() != nil {
		return fmt.Errorf("create redis client failed: %v", ping.Err())
	}
	dao.RDB = rdb
	log.Info("connect to redis success")
	return nil
}
