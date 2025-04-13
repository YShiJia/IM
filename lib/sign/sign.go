/**
 * @author ysj
 * @email 2239831438@qq.com
 * @createTime: 2025-04-13 10:57:16
 */

package sign

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func FileSignatureBySHA256(path string) (hash string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("open file %s error", path)
	}
	h_ob := sha256.New()
	if _, err = io.Copy(h_ob, file); err != nil {
		return "", fmt.Errorf("获取文件SHA256 hash错误 err: %v", err)
	}
	hashBytes := h_ob.Sum(nil)
	return hex.EncodeToString(hashBytes), nil
}
