package svc

import (
	"github.com/YShiJia/IM/apps/social/rpc/internal/config"
	"github.com/YShiJia/IM/database/dbfriend"
	"github.com/YShiJia/IM/database/dbgroup"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config   config.Config
	Redis    *redis.Redis
	FriendDB dbfriend.FriendDB
	GroupDB  dbgroup.GroupDB
}

func NewServiceContext(c config.Config) *ServiceContext {
	Redis, err := redis.NewRedis(c.Redisx)

	gormDB, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{})
	// TODO 后续可以添加一个获取本机的配置参数的pkg工具，根据该工具配置连接并发数等参数
	if err != nil {
		panic(err)
	}
	FriendDB := dbfriend.NewFriendDbByGorm(gormDB)
	GroupDB := dbgroup.NewGroupDBDbByGorm(gormDB)

	return &ServiceContext{
		Config:   c,
		Redis:    Redis,
		FriendDB: FriendDB,
		GroupDB:  GroupDB,
	}
}
