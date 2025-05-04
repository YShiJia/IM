/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-05 15:10:35
 */

package test

import (
	"fmt"
	"github.com/YShiJia/IM/pkg/jwt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestJwt(t *testing.T) {

	secretKey := "heathyang"
	iat := time.Now().Unix()
	seconds := int64(time.Hour * 24 * 365)
	payload := "testUid1"

	token, err := jwt.GetJwtToken(secretKey, iat, seconds, payload)
	fmt.Println(token)
	assert.NoError(t, err)

	parsePayload, err := jwt.ParseJwtToken(secretKey, token)
	assert.NoError(t, err)

	assert.Equal(t, payload, parsePayload)
}
