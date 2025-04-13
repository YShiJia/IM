/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-08 10:18:56
 */

package bitmap

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

// TODO 以后有需要用到的功能再完善
type BitMap interface {
	CountArray(ctx context.Context, arr []int) ([]int, error)
}

type bitMapByRedis struct {
	redis *redis.Redis
	key   string
}

func NewBitMapByRedis(r *redis.Redis, key string) BitMap {
	return &bitMapByRedis{
		redis: r,
		key:   key,
	}
}

func (b *bitMapByRedis) CountArray(ctx context.Context, arr []int) ([]int, error) {
	res := make([]int, len(arr)/2)
	err := b.redis.PipelinedCtx(ctx, func(pp redis.Pipeliner) error {
		for i := range arr {
			result, err := pp.GetBit(ctx, b.key, int64(arr[i])).Result()
			if err != nil {
				return err
			}
			if result == 1 {
				res = append(res, arr[i])
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
