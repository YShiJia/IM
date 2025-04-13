/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-02 15:17:23
 */

package svc

import (
	"github.com/YShiJia/IM/apps/edge/internal/config"
)

type ServiceContext struct {
	config.Config
}

// 在当前服务中弃用
func NewServiceContext(c config.Config) *ServiceContext {
	return nil
}
