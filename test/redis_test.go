package test

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"testing"
)

func Test_Redis(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "10.120.0.20:6379",
		Password: "heathyang", // 密码
		DB:       0,           // 数据库
		PoolSize: 50,          // 连接池大小
	})
	if ping := rdb.Ping(context.Background()); ping.Err() != nil {
		fmt.Printf("create redis client failed: %v", ping.Err())
	}
	cmd := rdb.Get(context.TODO(), "k1")

	log.Infof("create redis success %+v", cmd)
}
