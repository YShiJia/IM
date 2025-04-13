/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-08 21:29:48
 */

package svc

import (
	"context"
	"encoding/json"
	"github.com/YShiJia/IM/apps/status/rpc/internal/model"
	model2 "github.com/YShiJia/IM/model"
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/logx"
	clientv3 "go.etcd.io/etcd/client/v3"
	"sync"
	"time"
)

type EdgeServices struct {
	esList map[string]*model.EdgeServiceInfo
	l      sync.RWMutex
}

func NewEdgeServices(conf discov.EtcdConf, prefixKey string) *EdgeServices {
	edgeServices := &EdgeServices{
		esList: make(map[string]*model.EdgeServiceInfo),
	}
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   conf.Hosts,
		DialTimeout: time.Second * 3,
	})
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := cli.Get(ctx, prefixKey, clientv3.WithPrefix())
	if err != nil {
		panic(err)
	}
	for _, kv := range res.Kvs {
		var edgeMqInfo model2.EdgeMQInfo
		if err := json.Unmarshal(kv.Value, &edgeMqInfo); err != nil {
			logx.Errorf("invalid data key is: %s value is: %s", string(kv.Key), string(kv.Value))
			continue
		}
		// 获取边缘服务信息
		esInfo := getEdgeServiceInfo(&edgeMqInfo)

		edgeServices.l.Lock()
		edgeServices.esList[string(kv.Key)] = esInfo
		edgeServices.l.Unlock()
	}
	return edgeServices
}

func (e *EdgeServices) Get(key string) (*model.EdgeServiceInfo, bool) {
	e.l.RLock()
	defer e.l.RUnlock()
	esInfo, ok := e.esList[key]
	return esInfo, ok
}

func (e *EdgeServices) Update(key string, data []byte) {
	var edgeMqInfo model2.EdgeMQInfo
	if err := json.Unmarshal(data, &edgeMqInfo); err != nil {
		logx.Errorf("invalid data key is: %s value is: %s", key, string(data))
	}
	// 获取边缘服务信息
	esInfo := getEdgeServiceInfo(&edgeMqInfo)
	e.l.Lock()
	e.esList[key] = esInfo
	e.l.Unlock()
}

func (e *EdgeServices) Delete(key string) {
	e.l.Lock()
	delete(e.esList, key)
	e.l.Unlock()
}

func getEdgeServiceInfo(edgeMqInfo *model2.EdgeMQInfo) *model.EdgeServiceInfo {
	// TODO gozero 没有RequireAckAll，后期舍弃kq包，使用原生kafka-go封装一个组件来做
	RecvCmdMsgPusher := kq.NewPusher(
		edgeMqInfo.RecvCmdMsgConsumerConf.Brokers,
		edgeMqInfo.RecvCmdMsgConsumerConf.Topic,
		kq.WithBalancer(&kafka.Hash{}),
	)
	RecvCommonMsgPusher := kq.NewPusher(
		edgeMqInfo.RecvCommonMsgConsumerConf.Brokers,
		edgeMqInfo.RecvCommonMsgConsumerConf.Topic,
		kq.WithBalancer(&kafka.Hash{}),
	)
	return &model.EdgeServiceInfo{
		RecvCmdMsgPusher:    RecvCmdMsgPusher,
		RecvCommonMsgPusher: RecvCommonMsgPusher,
		Address:             edgeMqInfo.Address,
	}
}
