/**
 * @author ysj
 * @email 2239831438@qq.com
 * @createTime: 2025-04-19 18:13:41
 */

package db

import (
	"gorm.io/gorm"
)

var (
	IMDB *gorm.DB
)

type baseDAO struct {
	tx *gorm.DB
}

// NewBaseDAO 获取新的baseDAO
func NewBaseDAO(tx *gorm.DB) *baseDAO {
	return &baseDAO{tx: tx}
}

func (b *baseDAO) getTx() *gorm.DB {
	if (b == nil || b.tx == nil) && IMDB != nil {
		return IMDB
	}
	return b.tx
}
