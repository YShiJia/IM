package svc

import (
	"github.com/YShiJia/IM/apps/user/rpc/internal/config"
	"github.com/YShiJia/IM/database/dbuser"
	email2 "github.com/YShiJia/IM/pkg/email"
	"github.com/jordan-wright/email"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config    config.Config
	Redis     *redis.Redis
	UserDb    dbuser.UserDB
	EmailPool *email.Pool
}

func NewServiceContext(c config.Config) *ServiceContext {

	Redis, err := redis.NewRedis(c.Redisx)
	if err != nil {
		panic(err)
	}

	gormDB, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	UserDb := dbuser.NewUserDbByGorm(gormDB)

	EmailPool, err := email2.GetEmailPool(&c.EmailConfig)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:    c,
		Redis:     Redis,
		UserDb:    UserDb,
		EmailPool: EmailPool,
	}
}
