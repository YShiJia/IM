/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-03 11:07:31
 */

package discovery

import (
	"github.com/zeromicro/go-zero/core/discov"
)

type ServerObserver interface {
	Update(key string, data []byte)
	Delete(key string)
}

func QueueDiscoveryProc(conf discov.EtcdConf, prefixKey string, so ServerObserver) {
	master, err := NewServerMaster(prefixKey, conf.Hosts)
	if err != nil {
		panic(err)
	}
	master.Register(so)
	master.WatchQueueWorkers()
}
