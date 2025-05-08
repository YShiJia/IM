package config

import (
	imModel "github.com/YShiJia/IM/model"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	IPPrefix string // 本集群IP网段

	HttpPort int
	GrpcPort int

	Env string

	MysqlConf imModel.MysqlConfig

	AuthConf imModel.AuthConfig
}

var Conf = Config{}
