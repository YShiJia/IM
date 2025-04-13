package dao

import (
	"github.com/segmentio/kafka-go"
)

var KafkaConn *kafka.Conn

var SendMsgQueueWriter *kafka.Writer
var RecvMsgQueueReader *kafka.Reader

type kafkaDao struct{}

var Kafka = &kafkaDao{}

func (*kafkaDao) DelTopic(topic string) error {
	return KafkaConn.DeleteTopics(topic)
}

func (*kafkaDao) CreateTopic(topic kafka.TopicConfig) error {
	return KafkaConn.CreateTopics(topic)
}
