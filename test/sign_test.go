/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2025-05-05 20:12:07
 */

package test

import (
	"github.com/YShiJia/IM/lib/sign"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSign(t *testing.T) {
	fileName := "./Readme.md"
	file, err := os.Open(fileName)
	assert.NoError(t, err)

	hash, err := sign.FileSignatureByMD5(file)
	assert.NoError(t, err)
	log.Infof("hash: %s", hash)
}
