/**
 * @author ysj
 * @email 2239831438@qq.com
 * @createTime: 2025-04-19 20:10:11
 */

package nanoid

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	DefaultUIDLength = 10
	charSet          = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var (
	User  = NewUUID("uu")
	Group = NewUUID("gu")
)

type UUID struct {
	prefix string
	length int
}

func NewUUID(prefix string, length ...int) *UUID {
	uu := &UUID{prefix: prefix, length: length[0]}
	if uu.length == 0 {
		uu.length = DefaultUIDLength
	}
	if uu.length > len(charSet) {
		uu.length = len(charSet)
	}
	return uu
}

func (uu *UUID) generateRandomString() string {
	// 使用当前时间作为随机数种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 生成随机字符串
	var result []byte
	for i := 0; i < uu.length; i++ {
		randomIndex := r.Intn(len(charSet))           // 获取随机索引
		result = append(result, charSet[randomIndex]) // 添加字符到结果中
	}

	return string(result)
}

func (uu *UUID) Stand() string {
	str := uu.generateRandomString()[:uu.length]
	return fmt.Sprintf("%s-%s", uu.prefix, str)
}
