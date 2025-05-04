/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2025-04-29 12:16:48
 */

package init

import (
	"context"
	"encoding/json"
	"fmt"
	conf "github.com/YShiJia/IM/apps/message/api/internal/config"
	"github.com/YShiJia/IM/apps/message/api/internal/dao/etcd"
	"github.com/YShiJia/IM/apps/message/api/internal/dao/mq"
	"github.com/YShiJia/IM/model"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// InitMessage 初始化message服务
// 1. 创建 sendMsgQueue 队列，及其reader，用于接收所有edge服务推送的消息
// 2. 开启监听worker，监听etcd中edge服务的上下线事件，并实时更新edge的recvMessageQueue队列的writer
// 3. 开启任务队列worker，消费sendMsgQueue队列的数据
func InitMessage() error {
	if err := createTopic(); err != nil {
		return err
	}
	if err := WatchEdgeService(); err != nil {
		return err
	}
	return nil
}

// 创建 sendMsgQueue 队列, 并创建其reader
func createTopic() error {
	tc := kafka.TopicConfig{
		Topic:             conf.Conf.SendMessageQueue.Topic,
		NumPartitions:     conf.Conf.SendMessageQueue.Partition,
		ReplicationFactor: conf.Conf.SendMessageQueue.Replication,
	}
	if err := mq.Kafka.CreateTopic(tc); err != nil {
		return fmt.Errorf("create topic failed, err: %v", err)
	}
	// 创建sendMsgQueue的Reader
	mq.SendMsgQueueReader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{conf.Conf.SendMessageQueue.Broker},
		GroupID:  conf.Conf.Name, // 使用本服务名作为消费者组id
		Topic:    conf.Conf.SendMessageQueue.Topic,
		MaxBytes: 10e6, // 10MB
	})
	return nil
}

func WatchEdgeService() error {
	// 1.监听所有edge服务
	watchChan := etcd.EtcdClient.Watch(context.Background(), conf.Conf.EdgeNamePrefix, clientv3.WithPrefix())

	go watchEdgeChange(&watchChan)

	// 2. 创建所有的writer
	ctx, cancel := context.WithTimeout(context.Background(), conf.ReqTimeOut)
	resp, err := etcd.EtcdClient.Get(ctx, conf.Conf.EdgeNamePrefix, clientv3.WithPrefix())
	cancel()
	if err != nil {
		log.Fatalf("Failed to get keys: %v", err)
	}
	for _, kv := range resp.Kvs {
		si := model.ServiceInfo{}
		if err := json.Unmarshal(kv.Value, &si); err != nil {
			log.Errorf("fail to unmarshal service info data: %s", string(kv.Value))
		}
		writer := &kafka.Writer{
			Addr:         kafka.TCP(si.RecvMessageQueue.Broker),
			Topic:        si.RecvMessageQueue.Topic,
			RequiredAcks: kafka.RequireAll, // ack模式
		}
		mq.SetRecvMsgQueueWriter(si.Name, writer)
	}
	return nil
}

func watchEdgeChange(watchChan *clientv3.WatchChan) {
	for watchResp := range *watchChan {
		for _, ev := range watchResp.Events {
			log.Infof("edge 服务集群变更: 事件类型 %s Key: %s Value: %s", ev.Type, ev.Kv.Key, ev.Kv.Value)
			if ev.Type == clientv3.EventTypePut {
				si := model.ServiceInfo{}
				if err := json.Unmarshal(ev.Kv.Value, &si); err != nil {
					log.Errorf("fail to unmarshal service info data: %s", string(ev.Kv.Value))
					continue
				}
				writer := &kafka.Writer{
					Addr:         kafka.TCP(si.RecvMessageQueue.Broker),
					Topic:        si.RecvMessageQueue.Topic,
					RequiredAcks: kafka.RequireAll, // ack模式
				}
				mq.SetRecvMsgQueueWriter(si.Name, writer)
			} else if ev.Type == clientv3.EventTypeDelete {
				mq.DelRecvMsgQueueWriter(string(ev.Kv.Key))
			}
		}
	}
}
