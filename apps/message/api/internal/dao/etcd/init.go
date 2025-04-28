/**
 * @author ysj
 * @email 2239831438@qq.com
 * @createTime: 2025-04-23 13:11:35
 */

package etcd

import (
	"fmt"
	conf "github.com/YShiJia/IM/apps/message/api/internal/config"
	log "github.com/sirupsen/logrus"
	etcdv3 "go.etcd.io/etcd/client/v3"
)

var EtcdClient *etcdv3.Client

type etcdDao struct{}

var Etcd = &etcdDao{}

func InitEtcd() error {
	cli, err := etcdv3.New(etcdv3.Config{
		Endpoints:   []string{fmt.Sprintf("%s:%d", conf.Conf.EtcdConf.Host, conf.Conf.EtcdConf.Port)},
		DialTimeout: conf.Conf.DialTimeOut,
	})
	if err != nil {
		return fmt.Errorf("create etcd client failed, err:%v", err)
	}
	EtcdClient = cli

	log.Info("connect to etcd success")
	return nil
}
