package svc

import (
	"context"
	"errors"
	"fmt"
	"github.com/YShiJia/IM/apps/status/rpc/internal/config"
	"github.com/YShiJia/IM/database/dbuser"
	"github.com/YShiJia/IM/lib/bitmap"
	"github.com/YShiJia/IM/lib/discovery"
	"github.com/YShiJia/IM/lib/encoder"
	"github.com/YShiJia/IM/pbmodel/pbmessage"
	csHash "github.com/YShiJia/consistentHash"
	"github.com/YShiJia/consistentHash/redisHashRing"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/threading"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	BitMapPrefix = "redis:bitmap"
)

type ServiceContext struct {
	Config  config.Config
	Encoder encoder.Encoder
	CsHash  *csHash.ConsistentHash

	Redis            *redis.Redis
	UserDb           dbuser.UserDB
	UserOnlineBitMap bitmap.BitMap

	EdgeService *EdgeServices

	SendMsgConsumer service.Service
}

func NewServiceContext(c config.Config) *ServiceContext {
	svcCtx := &ServiceContext{}

	Encoder := encoder.NewProtobufEncoder()

	client := redisHashRing.NewClient("tcp", c.Redisx.Host, c.Redisx.Pass)
	redisHashRing := redisHashRing.NewRedisHashRing(c.Name, client)
	CsHash := csHash.NewConsistentHash(redisHashRing, csHash.NewMurmurHasher32(), svcCtx.Migrate, logx.WithContext(context.Background()))

	Redis, err := redis.NewRedis(c.Redisx)
	if err != nil {
		panic(err)
	}

	//这里以后可以加上 gorm
	gormDB, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	UserDb := dbuser.NewUserDbByGorm(gormDB)

	UserOnlineBitMap := bitmap.NewBitMapByRedis(Redis, fmt.Sprintf("%s:%s:%s", BitMapPrefix, c.Name, "userOnline"))

	EdgeService := NewEdgeServices(svcCtx.Config.Etcd, svcCtx.Config.IMEdgePrefix)

	SendMsgConsumer := kq.MustNewQueue(
		c.SendMsgConsumerConf,
		NewSendMsgConsumerLogic(svcCtx),
	)
	svcCtx.Config = c
	svcCtx.Encoder = Encoder
	svcCtx.CsHash = CsHash
	svcCtx.Redis = Redis
	svcCtx.UserDb = UserDb
	svcCtx.UserOnlineBitMap = UserOnlineBitMap
	svcCtx.EdgeService = EdgeService
	svcCtx.SendMsgConsumer = SendMsgConsumer
	return svcCtx
}

// TODO 当有节点变动的时候，触发迁移，添加或者删除节点
func (svcCtx *ServiceContext) Migrate(ctx context.Context, dataKeys map[string]struct{}, from, to string) error {
	/**
	1. 根据数据 key(就是 to ) 获取相关的 pusher
	2. 向 pusher 发送连接转移的消息
	*/
	edgeServerInfo, ok := svcCtx.EdgeService.Get(to)
	if !ok {
		//该节点不存在
		_ = svcCtx.CsHash.RemoveNode(ctx, to)
		return errors.New("目标迁移节点不存在")
	}

	keys := make([]string, len(dataKeys))
	for key := range dataKeys {
		keys = append(keys, key)
	}

	message := pbmessage.NewUpdateConnPbMessage(keys)
	msgData, _ := svcCtx.Encoder.Encode(message)
	return edgeServerInfo.RecvCmdMsgPusher.KPush(ctx, svcCtx.Config.Name, string(msgData))
}

func (svcCtx *ServiceContext) ListenEdgeServer() {
	threading.GoSafe(func() {
		discovery.QueueDiscoveryProc(svcCtx.Config.Etcd, svcCtx.Config.IMEdgePrefix, svcCtx.EdgeService)
	})
}

func (svcCtx *ServiceContext) ConsumeEdgeMsg() {
	serviceGroup := service.NewServiceGroup()
	serviceGroup.Add(svcCtx.SendMsgConsumer)
	threading.GoSafe(func() {
		defer serviceGroup.Stop()
		serviceGroup.Start()
	})
}
