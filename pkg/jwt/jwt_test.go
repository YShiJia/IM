/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-02 20:33:57
 */

package jwt

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_JwtToken(t *testing.T) {
	secretKey := "password"
	iat := time.Now().Unix()
	seconds := int64(time.Hour)
	payload := "123"

	token, err := GetJwtToken(secretKey, iat, seconds, payload)
	assert.NoError(t, err)

	parsePayload, err := ParseJwtToken(secretKey, token)
	assert.NoError(t, err)

	assert.Equal(t, payload, parsePayload)
}
