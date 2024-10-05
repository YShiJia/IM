/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-02 15:17:23
 */

package svc

import (
	"context"
	"github.com/YShiJia/IM/apps/edge/internal/config"
	"github.com/YShiJia/IM/apps/edge/internal/route"
	"github.com/YShiJia/IM/lib/encoder"
	"github.com/YShiJia/IM/lib/webscoket"
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"net"
	"strconv"
)

// 服务器级别的ctx
type ServiceContext struct {
	config.Config
	Encoder encoder.Encoder
	logx.Logger
	WsServer *webscoket.WsServer
	//发送消息
	SendMsgPusher *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	// wsServer初始化是在NewServiceContext
	// 填充处理连接逻辑是在EdgeServer
	wsServer := webscoket.NewWsServer(c.ListenOn)
	//注册路由
	route.RegisterHandlers(wsServer)

	if err := autoCreateMQ(c); err != nil {
		panic(err)
	}

	// 保证消息从本服务到目标消息队列的有序性，需要使用分区平衡算法
	// TODO gozero 没有RequireAckAll，后期舍弃kq包，使用原生kafka-go封装一个组件来做
	SendMsgPusher := kq.NewPusher(
		c.SendMsgPusherConf.Brokers,
		c.SendMsgPusherConf.Topic,
		kq.WithBalancer(&kafka.Hash{}),
	)
	svc := &ServiceContext{
		Config:        c,
		Encoder:       encoder.NewProtobufEncoder(),
		Logger:        logx.WithContext(context.TODO()),
		WsServer:      wsServer,
		SendMsgPusher: SendMsgPusher,
	}

	return svc
}

// TODO 下面的代码写的不是很满意，后续优化一下
func autoCreateMQ(c config.Config) error {
	//创建三个消息队列

	var topic []string
	//1.判断三个消息队列是否存在
	m, err := getKafkaTopicList(c)
	if err != nil {
		return err
	}

	if _, ok := m[c.SendMsgPusherConf.Topic]; !ok {
		topic = append(topic, c.SendMsgPusherConf.Topic)
	}
	if _, ok := m[c.RecvCmdMsgConsumerConf.Topic]; !ok {
		topic = append(topic, c.RecvCmdMsgConsumerConf.Topic)
	}
	if _, ok := m[c.RecvCommonMsgConsumerConf.Topic]; !ok {
		topic = append(topic, c.RecvCommonMsgConsumerConf.Topic)
	}

	// 获取kafka控制节点连接
	controllerConn, err := getKafkaControllerConn(c)
	if err != nil {
		return err
	}

	topicConfigs := []kafka.TopicConfig{}
	for _, t := range topic {
		topicConfigs = append(topicConfigs, kafka.TopicConfig{
			Topic:             t,
			NumPartitions:     3,
			ReplicationFactor: 3,
		})
	}

	// 创建topic
	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		return err
	}
	return nil
}

func getKafkaControllerConn(c config.Config) (*kafka.Conn, error) {
	conn, err := kafka.Dial("tcp", c.SendMsgPusherConf.Brokers[0])
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// 获取当前控制节点信息
	controller, err := conn.Controller()
	if err != nil {
		return nil, err
	}
	// 连接至leader节点
	controllerConn, err := kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		return nil, err
	}
	return controllerConn, nil
}

func getKafkaTopicList(c config.Config) (map[string]struct{}, error) {
	conn, err := kafka.Dial("tcp", c.SendMsgPusherConf.Brokers[0])
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		return nil, err
	}

	m := map[string]struct{}{}
	// 遍历所有分区取topic
	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}
	return m, nil
}
