/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2025-04-29 12:01:03
 */

package init

import (
	"fmt"
	conf "github.com/YShiJia/IM/apps/message/api/internal/config"
	"github.com/YShiJia/IM/apps/message/api/internal/dao/mq"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

// InitKafka 获取kafka的Leader节点连接
// 1. 获取连接
func InitKafka() error {
	// 连接至Kafka集群的Leader节点
	leaderConn, err := kafka.Dial("tcp", conf.Conf.SendMessageQueue.Broker)
	if err != nil {
		return fmt.Errorf("create kafka client failed:%v", err)
	}
	mq.KafkaConn = leaderConn
	log.Info("connect to kafka success")
	return nil
}
