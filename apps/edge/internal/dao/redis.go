package dao

import (
	"context"
	"fmt"
	conf "github.com/YShiJia/IM/apps/edge/internal/config"
	"github.com/redis/go-redis/v9"
	"time"
)

var RDB *redis.Client

type redisDao struct{}

var Redis = &redisDao{}

func (*redisDao) GetNextNumber(ctx context.Context, key string) (int, error) {
	// 获取key对应的值
	// 1. 如果存在且value > 0，value++，返回
	// 2. 不存在，创建value = 1，返回
	script := `
		local key = KEYS[1]
		local value = tonumber(redis.call("GET", key))

		if value and value > 0 then
		    redis.call("INCR", key)
		    return value + 1
		else
		    redis.call("SET", key, 1)
		    return 1
		end
	`

	// 执行 Lua 脚本
	number, err := RDB.Eval(ctx, script, []string{key}).Int()
	if err != nil {
		return 0, fmt.Errorf("executing Lua script failed, err: %v", err)
	}
	return number, nil
}

// RegisterUserUid 创建or更新user的过期时间
func (*redisDao) RegisterUserUid(ctx context.Context, userUid, edgeName string, expiration time.Duration) error {
	// 将用户的 UID 存入 Redis，设置过期时间
	err := RDB.SetEx(ctx, fmt.Sprintf("%s-%s", conf.Conf.RedisUserInfoPrefix, userUid), edgeName, expiration).Err()
	if err != nil {
		return fmt.Errorf("failed to set user UID with expiration, err %v", err)
	}
	return nil
}
