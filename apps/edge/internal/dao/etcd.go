package dao

import (
	"context"
	"fmt"
	etcdv3 "go.etcd.io/etcd/client/v3"
)

var EtcdClient *etcdv3.Client

type etcdDao struct{}

var Etcd = &etcdDao{}

func (*etcdDao) PutWithExpireTime(ctx context.Context, key, value string, expire int, opts ...etcdv3.OpOption) (etcdv3.LeaseID, error) {
	leaseResp, err := EtcdClient.Grant(ctx, int64(expire))
	if err != nil {
		return 0, fmt.Errorf("create lease failed, err%v", err)
	}
	opts = append(opts, etcdv3.WithLease(leaseResp.ID))

	_, err = EtcdClient.Put(ctx, key, value, opts...)
	return leaseResp.ID, err
}

func (*etcdDao) KeepAlive(ctx context.Context, id etcdv3.LeaseID) (<-chan *etcdv3.LeaseKeepAliveResponse, error) {
	return EtcdClient.KeepAlive(ctx, id)
}
