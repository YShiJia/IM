/**
 * @author ysj
 * @email 2239831438@qq.com
 * @createTime: 2025-04-23 13:11:35
 */

package etcd

import (
	etcdv3 "go.etcd.io/etcd/client/v3"
)

var EtcdClient *etcdv3.Client

type etcdDao struct{}

var Etcd = &etcdDao{}
