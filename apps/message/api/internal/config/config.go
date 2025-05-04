package config

import (
	imModel "github.com/YShiJia/IM/model"
	"github.com/zeromicro/go-zero/rest"
	"time"
)

var (
	DialTimeOut = time.Second * 3
	ReqTimeOut  = time.Second * 3
)

type Config struct {
	rest.RestConf

	MaxFileBytes      int64
	MaxTotalFileBytes int64
	IPPrefix          string // 本集群IP网段

	HttpPort int
	GrpcPort int

	Env string

	EdgeNamePrefix string

	RedisConf imModel.RedisConfig

	EtcdConf imModel.EtcdConfig

	SendMessageQueue imModel.KafkaConfig

	MysqlConf imModel.MysqlConfig

	MinioConf imModel.MinioConfig

	AuthConf imModel.AuthConfig
}

var Conf = Config{}
