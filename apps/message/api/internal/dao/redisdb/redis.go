/**
 * @author ysj
 * @email 2239831438@qq.com
 * @createTime: 2025-04-23 13:06:20
 */

package redisdb

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func GetValue(ctx context.Context, key string) (string, error) {
	val, err := RDB.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
