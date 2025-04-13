package init

import (
	"fmt"
	conf "github.com/YShiJia/IM/apps/edge/internal/config"
	"github.com/YShiJia/IM/apps/edge/internal/dao"
	log "github.com/sirupsen/logrus"
	etcdv3 "go.etcd.io/etcd/client/v3"
)

func InitEtcd() error {
	cli, err := etcdv3.New(etcdv3.Config{
		Endpoints:   []string{fmt.Sprintf("%s:%d", conf.Conf.EtcdConf.Host, conf.Conf.EtcdConf.Port)},
		DialTimeout: conf.Conf.DialTimeOut,
	})
	if err != nil {
		return fmt.Errorf("create etcd client failed, err:%v", err)
	}
	dao.EtcdClient = cli

	log.Info("connect to etcd success")
	return nil
}
