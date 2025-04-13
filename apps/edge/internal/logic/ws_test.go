/**
 * @author ysj
 * @email 2239831438@qq.com
 * @createTime: 2025-04-13 12:52:05
 */

package logic

import (
	initialize "github.com/YShiJia/IM/apps/edge/internal/init"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCommunicateWithClient_HeartBeat(t *testing.T) {
	assert.Nil(t, initialize.InitConfig())
	err := initialize.InitRedis()
	assert.Nil(t, err)
	cwc := &CommunicateWithClient{
		ClientMaxSilenceTime: time.Second * 10,
	}
	cwc.HeartBeat()
}
