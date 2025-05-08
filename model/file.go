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
	ID         uint           `gorm:"comment:主键ID;"`
	Bucket     string         `gorm:"comment:存储桶;size:20;not null;"`
	Hash       string         `gorm:"comment:唯一标识;size:500;uniqueIndex;not null;"`
	Size       int64          `gorm:"comment:文件大小;not null;"`
	Name       string         `gorm:"comment:文件名;size:30;not null;"`
	DeletedAt  gorm.DeletedAt `gorm:"comment:删除时间;null"`
	FileSlices []*FileSlice
}

type FileSlice struct {
	ID     uint   `gorm:"comment:主键ID;"`
	Bucket string `gorm:"comment:存储桶;size:20;not null;"`
	Hash   string `gorm:"comment:唯一标识;size:500;uniqueIndex;not null;"`
	Size   int64  `gorm:"comment:文件大小;not null;"`
	FileID uint   `gorm:"comment:文件ID;uniqueIndex:idx_fs_fo;not null;"`
	Order  uint8  `gorm:"comment:分片顺序;uniqueIndex:idx_fs_fo;not null;"`

	File *File
}
