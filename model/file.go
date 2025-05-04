/**
 * @author ysj
 * @email 2239831438@qq.com
 * @createTime: 2025-04-19 18:02:16
 */

package model

import (
	"gorm.io/gorm"
)

// File 文件
type File struct {
	ID        uint           `gorm:"comment:主键ID;"`
	Hash      string         `gorm:"comment:唯一标识;size:500;uniqueIndex;not null;"`
	Bucket    string         `gorm:"comment:存储桶;size:20;not null;"`
	Name      string         `gorm:"comment:文件名;size:30;index;not null;"`
	DeletedAt gorm.DeletedAt `gorm:"comment:删除时间;null"`
}
