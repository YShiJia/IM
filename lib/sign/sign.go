/**
 * @author ysj
 * @email 2239831438@qq.com
 * @createTime: 2025-04-13 10:57:16
 */

package sign

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

func FileSignatureBySHA256(reader io.Reader) (hash string, err error) {
	h_ob := sha256.New()
	if _, err = io.Copy(h_ob, reader); err != nil {
		return "", fmt.Errorf("获取文件SHA256 hash错误 err: %v", err)
	}
	hashBytes := h_ob.Sum(nil)
	return hex.EncodeToString(hashBytes), nil
}

func FileSignatureByMD5(reader io.Reader) (hash string, err error) {
	h_ob := md5.New() // 创建一个新的 MD5 哈希对象
	if _, err = io.Copy(h_ob, reader); err != nil {
		return "", fmt.Errorf("获取文件MD5 hash错误 err: %v", err)
	}
	hashBytes := h_ob.Sum(nil) // 计算哈希值
	return hex.EncodeToString(hashBytes), nil
}
