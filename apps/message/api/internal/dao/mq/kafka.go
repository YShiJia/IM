/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2025-05-03 00:57:24
 */

package mq

import (
	"github.com/segmentio/kafka-go"
)

var KafkaConn *kafka.Conn

type kafkaDao struct{}

var Kafka = &kafkaDao{}

func (*kafkaDao) DelTopic(topic string) error {
	return KafkaConn.DeleteTopics(topic)
}

func (*kafkaDao) CreateTopic(topic kafka.TopicConfig) error {
	return KafkaConn.CreateTopics(topic)
}
