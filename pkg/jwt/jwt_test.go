/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-02 20:33:57
 */

package jwt

import (
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_JwtToken(t *testing.T) {
	secretKey := "heathyang"
	iat := time.Now().Unix()
	seconds := int64(time.Hour * 86400)
	payload := "im-dev-uid"

	token, err := GetJwtToken(secretKey, iat, seconds, payload)
	assert.NoError(t, err)

	parsePayload, err := ParseJwtToken(secretKey, token)
	assert.NoError(t, err)

	log.Infof("token: %v", token)

	assert.Equal(t, payload, parsePayload)
}
