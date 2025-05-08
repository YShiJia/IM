/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2025-05-08 16:01:41
 */

package logic

import (
	"crypto/md5"
	"encoding/hex"
)

func EncryptPasswordByMD5(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}
