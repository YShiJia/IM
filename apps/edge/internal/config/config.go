/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-02 15:08:35
 */

package config

import (
	"github.com/YShiJia/IM/apps/edge/internal/model"
	imModel "github.com/YShiJia/IM/model"
	"time"
)

const (
	AuthIMUserUID = "AuthIMUserUID"
	Authorization = "Authorization"
)

type Config struct {
	Name       string
	Number     int // 本服务编号
	NamePrefix string
	IPPrefix   string // 本集群IP网段

	HttpPort int
	GrpcPort int

	Env string

	WebSocketConnPoolMaxSize int           // ws连接池能容纳的最大连接数
	WebSocketConnectIdleTime time.Duration // ws连接最大空闲时间，超过该时间说明连接不可用
	DialTimeOut              time.Duration // 连接超时
	ReqTimeOut               time.Duration // 请求超时
	ServeInfoExpireTime      time.Duration // 存储在etcd的服务信息的超时时间
	ClientMaxSilenceTime     time.Duration // 客户端最大静默时间，超过该时间服务端认定为该连接不可用

	RedisConf imModel.RedisConfig

	EtcdConf imModel.EtcdConfig

	KafkaConf model.EdgeKafkaConfig

	AuthConf imModel.AuthConfig
}

var Conf = Config{
	WebSocketConnectIdleTime: time.Second * 10,
	DialTimeOut:              time.Second * 3,
	ReqTimeOut:               time.Second * 3,
	ServeInfoExpireTime:      time.Second * 10,
	ClientMaxSilenceTime:     time.Second * 10,
}
