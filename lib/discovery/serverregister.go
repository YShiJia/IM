/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-03 11:08:14
 */

package discovery

import (
	"context"
	"encoding/json"
	"github.com/YShiJia/IM/lib/encoder"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var etcdDialTimeout = 3 * time.Second

type ServerRegisterOption func(sr *ServerRegister)

type ServerRegister struct {
	key        string
	value      any
	encoder    encoder.Encoder
	etcdClient *clientv3.Client
}

func WithProtobufEncoder() ServerRegisterOption {
	return func(sr *ServerRegister) {
		sr.encoder = encoder.NewProtobufEncoder()
	}
}
func WithJsonEncoder() ServerRegisterOption {
	return func(sr *ServerRegister) {
		sr.encoder = encoder.NewJsonEncoder()
	}
}

// 创建一个注册服务维持器
func NewServerRegister(key string, endpoints []string, value any, pts ...ServerRegisterOption) *ServerRegister {
	cfg := clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: etcdDialTimeout,
	}
	etcdClient, err := clientv3.New(cfg)
	if err != nil {
		panic(err)
	}
	return &ServerRegister{
		key:        key,
		etcdClient: etcdClient,
		value:      value,
		// 为了调试可读性，使用json编码
		encoder: encoder.NewJsonEncoder(),
	}
}

func (s *ServerRegister) HeartBeat() {
	data, err := json.Marshal(s.value)
	if err != nil {
		panic(err)
	}
	s.register(string(data))
}

func (s *ServerRegister) register(value string) {
	//申请一个45秒的租约
	leaseGrantResp, err := s.etcdClient.Grant(context.TODO(), 45)
	if err != nil {
		panic(err)
	}

	//拿到租约的id
	leaseId := leaseGrantResp.ID
	logx.Infof("查看leaseId:%x", leaseId)

	//获得kv api子集
	kv := clientv3.NewKV(s.etcdClient)
	//put一个kv，让它与租约关联起来，从而实现45秒后自动过期
	putResp, err := kv.Put(context.TODO(), s.key, value, clientv3.WithLease(leaseId))
	if err != nil {
		panic(err)
	}

	//(自动续租)当我们申请了租约之后，我们就可以启动一个续租
	// 心跳检测，防止租约失效
	keepRespChan, err := s.etcdClient.KeepAlive(context.TODO(), leaseId)
	if err != nil {
		panic(err)
	}

	//处理续租应答的协程
	go func() {
		for {
			select {
			case keepResp, ok := <-keepRespChan:
				if !ok {
					logx.Infof("租约已经失效:%x", leaseId)
					s.register(value)
					return
				} else { //每秒会续租一次，所以就会受到一次应答
					logx.Infof("收到自动续租应答:%x", keepResp.ID)
				}
			}
		}
	}()

	logx.Info("写入成功:", putResp.Header.Revision)
}
