/**
 * @author ysj
 * @email 2239831438@qq.com
 * @createTime: 2025-04-23 13:13:28
 */

package mq

import (
	"fmt"
	conf "github.com/YShiJia/IM/apps/message/api/internal/config"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

var KafkaConn *kafka.Conn

var SendMsgQueueWriter *kafka.Writer
var RecvMsgQueueReader *kafka.Reader

type kafkaDao struct{}

var Kafka = &kafkaDao{}

// InitKafka 获取kafka的Leader节点连接
// 1. 获取连接
func InitKafka() error {
	// 连接至Kafka集群的Leader节点
	leaderConn, err := kafka.Dial("tcp", conf.Conf.KafkaConf.Broker)
	if err != nil {
		return fmt.Errorf("create kafka client failed:%v", err)
	}
	KafkaConn = leaderConn
	log.Info("connect to kafka success")
	return nil
}
