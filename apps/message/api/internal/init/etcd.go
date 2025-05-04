/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2025-04-29 12:00:53
 */

package init

import (
	"context"
	"fmt"
	conf "github.com/YShiJia/IM/apps/message/api/internal/config"
	"github.com/YShiJia/IM/apps/message/api/internal/dao/etcd"
	log "github.com/sirupsen/logrus"
	etcdv3 "go.etcd.io/etcd/client/v3"
)

func InitEtcd() error {
	cli, err := etcdv3.New(etcdv3.Config{
		Endpoints:   []string{fmt.Sprintf("%s:%d", conf.Conf.EtcdConf.Host, conf.Conf.EtcdConf.Port)},
		Username:    conf.Conf.EtcdConf.Username,
		Password:    conf.Conf.EtcdConf.Password,
		DialTimeout: conf.DialTimeOut,
	})
	if err != nil {
		return fmt.Errorf("create etcd client failed, err:%v", err)
	}
	ctx, _ := context.WithTimeout(context.Background(), conf.ReqTimeOut)
	if _, err = cli.Get(ctx, "ping"); err != nil {
		return fmt.Errorf("get ping from etcd failed, err:%v", err)
	}

	etcd.EtcdClient = cli

	log.Info("connect to etcd success")
	return nil
}
