package init

import (
	"context"
	"encoding/json"
	"fmt"
	conf "github.com/YShiJia/IM/apps/edge/internal/config"
	"github.com/YShiJia/IM/apps/edge/internal/dao"
	"github.com/YShiJia/IM/apps/edge/internal/logic"
	"github.com/YShiJia/IM/lib/ip"
	imModel "github.com/YShiJia/IM/model"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	etcdv3 "go.etcd.io/etcd/client/v3"
	"time"
)

// 1. 从redis中获取到自身的服务名
// 2. 创建自身的接收消息的队列
// 3. 创建对sendMsgQueue的Writer
// 4. 创建对recvMsgQueue对Reader
// 5. 将本edge服务的配置信息put到etcd中
// 6. 消费recvMsgQueue的数据
func InitEdge() error {
	if err := getAndUseServeNumber(); err != nil {
		return err
	}
	if err := createTopic(); err != nil {
		return err
	}
	if err := createWriterAndReader(); err != nil {
		return err
	}
	if err := registerInfoToEtcd(); err != nil {
		return err
	}
	go logic.Send()
	log.Info("init edge success")
	return nil
}

// getAndUseServeNumber 从redis中获取到自身的服务名、队列名和监听端口
func getAndUseServeNumber() error {
	ctx, cancel := context.WithTimeout(context.Background(), conf.Conf.ReqTimeOut)
	defer cancel()
	number, err := dao.Redis.GetNextNumber(ctx, conf.Conf.NamePrefix)
	if err != nil {
		return fmt.Errorf("get next number failed, err: %v", err)
	}
	conf.Conf.Number = number
	conf.Conf.Name = fmt.Sprintf("%s-%d", conf.Conf.NamePrefix, number)
	conf.Conf.KafkaConf.RecvMessageQueue.Topic = fmt.Sprintf("%s-%d", conf.Conf.RecvMessageQueueTopicPrefix, number)
	//conf.Conf.HttpPort = conf.Conf.HttpPort + number
	return nil
}

// 创建本服务接收消息队列
func createTopic() error {
	tc := kafka.TopicConfig{
		Topic:             conf.Conf.KafkaConf.RecvMessageQueue.Topic,
		NumPartitions:     conf.Conf.KafkaConf.RecvMessageQueue.Partition,
		ReplicationFactor: conf.Conf.KafkaConf.RecvMessageQueue.Replication,
	}
	if err := dao.Kafka.CreateTopic(tc); err != nil {
		return fmt.Errorf("create topic failed, err: %v", err)
	}
	return nil
}

// 创建对sendMsgQueue的Writer & 创建对recvMsgQueue的Reader
func createWriterAndReader() error {
	dao.SendMsgQueueWriter = &kafka.Writer{
		Addr:         kafka.TCP(conf.Conf.KafkaConf.SendMessageQueue.Broker),
		Topic:        conf.Conf.KafkaConf.SendMessageQueue.Topic,
		RequiredAcks: kafka.RequireAll, // ack模式
	}

	// 创建recvMsgQueue的Reader
	dao.RecvMsgQueueReader = kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{conf.Conf.KafkaConf.RecvMessageQueue.Broker},
		// 这里有一个薄弱点：当前是每一个edge都拥有自己的一个topic，都是从头开始消费
		// 如果后续
		GroupID:  conf.Conf.NamePrefix, // 统一使用edge作为消费者组id
		Topic:    conf.Conf.KafkaConf.RecvMessageQueue.Topic,
		MaxBytes: 10e6, // 10MB
	})
	return nil
}

// 将本节点的数据push到etcd中
func registerInfoToEtcd() error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), conf.Conf.ReqTimeOut)
	defer cancelFunc()

	ips, err := ip.GetIPv4Addr(conf.Conf.IPPrefix)
	if err != nil {
		return fmt.Errorf("get local ip addr failed, err: %v", err)
	}
	if len(ips) == 0 {
		return fmt.Errorf("can not get local serve cluster ip addr")
	}

	si := imModel.ServiceInfo{
		Name:             conf.Conf.Name,
		IP:               ips[0],
		HttpPort:         conf.Conf.HttpPort,
		GrpcPort:         conf.Conf.GrpcPort,
		Type:             imModel.SERVE_TYPE_EDGE,
		RecvMessageQueue: conf.Conf.KafkaConf.RecvMessageQueue,
	}

	siData, err := json.Marshal(si)
	if err != nil {
		return fmt.Errorf("marshal serve info to json data err:%v", err)
	}

	// put 本服务信息到etcd中，并设置超时时间
	leaseId, err := dao.Etcd.PutWithExpireTime(ctx, conf.Conf.Name, string(siData), int(conf.Conf.ReqTimeOut/time.Second))
	if err != nil {
		return fmt.Errorf("put edge service info to etcd err: %v", err)
	}

	// 监听租约续约
	kaChan, err := dao.Etcd.KeepAlive(context.TODO(), leaseId)
	if err != nil {
		return fmt.Errorf("keepalive edge serve info in etcd failed err %v", err)
	}

	// 在后台持续接收租约续约信息
	go keepalive(kaChan)

	return nil
}

func keepalive(kaChan <-chan *etcdv3.LeaseKeepAliveResponse) {
	for {
		select {
		case kaResp := <-kaChan:
			if kaResp == nil {
				log.Info("Lease expired.")
				return
			}
			log.Tracef("Lease renewed, ttl: %d seconds\n", kaResp)
		}
	}
}
