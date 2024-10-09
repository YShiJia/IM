/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-03 11:08:02
 */

package discovery

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type ServerMaster struct {
	cli            *clientv3.Client
	prefixKey      string
	serverObserver ServerObserver
}

func NewServerMaster(prefixKey string, address []string) (*ServerMaster, error) {
	cfg := clientv3.Config{
		Endpoints:   address,
		DialTimeout: time.Second * 3,
	}
	cli, err := clientv3.New(cfg)
	if err != nil {
		return nil, err
	}
	return &ServerMaster{
		cli:       cli,
		prefixKey: prefixKey,
	}, nil
}

func (s *ServerMaster) Register(so ServerObserver) {
	s.serverObserver = so
}

func (s *ServerMaster) notifyUpdate(key string, data []byte) {
	s.serverObserver.Update(key, data)
}

func (s *ServerMaster) notifyDelete(key string) {
	s.serverObserver.Delete(key)
}

func (s *ServerMaster) updateQueueWorker(key string, data []byte) {
	s.notifyUpdate(key, data)
}

func (s *ServerMaster) deleteQueueWorker(key string) {
	s.notifyDelete(key)
}

func (s *ServerMaster) WatchQueueWorkers() {
	rch := s.cli.Watch(context.Background(), s.prefixKey, clientv3.WithPrefix())

	for wresp := range rch {
		if wresp.Err() != nil {
			logx.Severe(wresp.Err())
		}
		if wresp.Canceled {
			logx.Severe("watch is canceled")
		}
		for _, ev := range wresp.Events {
			switch ev.Type {
			case clientv3.EventTypePut:
				s.updateQueueWorker(string(ev.Kv.Key), ev.Kv.Value)
			case clientv3.EventTypeDelete:
				s.deleteQueueWorker(string(ev.Kv.Key))
			}
		}
	}
}
