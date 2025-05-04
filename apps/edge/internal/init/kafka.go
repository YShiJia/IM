package init

import (
	"fmt"
	conf "github.com/YShiJia/IM/apps/edge/internal/config"
	"github.com/YShiJia/IM/apps/edge/internal/dao"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

// InitKafka 获取kafka的Leader节点连接
// 1. 获取leader节点连接
func InitKafka() error {
	// 连接至Kafka集群的Leader节点
	leaderConn, err := kafka.Dial("tcp", conf.Conf.KafkaConf.RecvMessageQueue.Broker)
	if err != nil {
		return fmt.Errorf("create kafka client failed:%v", err)
	}
	dao.KafkaConn = leaderConn
	log.Info("connect to kafka success")
	return nil
}
