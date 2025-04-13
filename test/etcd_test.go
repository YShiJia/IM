package test

import (
	"context"
	log "github.com/sirupsen/logrus"
	etcdv3 "go.etcd.io/etcd/client/v3"
	"testing"
	"time"
)

func Test_ETCD(t *testing.T) {
	cli, err := etcdv3.New(etcdv3.Config{
		Endpoints:   []string{"10.120.0.40:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Infof("connect to etcd failed, err:%v", err)
	}
	put, err := cli.Put(context.Background(), "kt", "test")
	if err != nil {
		log.Infof("put to etcd failed, err:%v", err)
	}
	log.Infof("put to etcd success, putresp: %v", put)

	get, err := cli.Get(context.Background(), "kt", etcdv3.WithPrefix())
	if err != nil {
		log.Infof("get kt failed, err:%v", err)
	}
	log.Infof("get from etcd success, getresp: %v", get)
}
