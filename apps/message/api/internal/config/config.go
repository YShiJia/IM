package config

import (
	imModel "github.com/YShiJia/IM/model"
	"github.com/zeromicro/go-zero/rest"
	"time"
)

type Config struct {
	rest.RestConf

	MaxFileBytes      int64
	MaxTotalFileBytes int64
	IPPrefix          string // 本集群IP网段

	HttpPort int
	GrpcPort int

	Env string

	DialTimeOut time.Duration // 连接超时
	ReqTimeOut  time.Duration // 请求超时

	RedisConf imModel.RedisConfig

	EtcdConf imModel.EtcdConfig

	KafkaConf imModel.KafkaConfig

	MysqlConf imModel.MysqlConfig

	MinioConf imModel.MinioConfig

	AuthConf imModel.AuthConfig
}

var Conf = Config{
	DialTimeOut: time.Second * 3,
	ReqTimeOut:  time.Second * 3,
}
