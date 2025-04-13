package test

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"testing"
	"time"
)

func Test_Kafka(t *testing.T) {
	address := "10.120.0.50:9092"
	topic := "my-topic"
	partition := 0

	// 连接至Kafka集群的Leader节点
	conn, err := kafka.DialLeader(context.Background(), "tcp", address, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	// 设置发送消息的超时时间
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	// 发送消息
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)

	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	// 关闭连接
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

func Test_Kafka2(t *testing.T) {
	address := "10.120.0.50:9092"

	// 连接至Kafka集群的Leader节点
	leaderConn, err := kafka.Dial("tcp", address)
	defer leaderConn.Close()
	if err != nil {
		log.Infof("failed to dial leader:%v", err)
	}

	// 设置发送消息的超时时间
	err = leaderConn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Infof("failed to set write deadline:%v", err)
	}
	//write, err := leaderConn.Write([]byte("ping"))
	//log.Infof("write:%v, err:%v", write, err)
	batch := leaderConn.ReadBatch(1e3, 1e6)
	b := make([]byte, 10e3)
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b[:n]))
	}

	partitions, err := leaderConn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	m := map[string]struct{}{}
	// 遍历所有分区取topic
	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}
	for k := range m {
		if err := leaderConn.DeleteTopics(k); err != nil {
			log.Infof("failed to delete topic:%v", err)
		}
		log.Infof("deleted topic:%v", k)
	}
}
