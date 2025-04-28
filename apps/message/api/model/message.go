/**
 * @author ysj
 * @email 2239831438@qq.com
 * @createTime: 2025-04-20 21:10:33
 */

package model

import (
	im "github.com/YShiJia/IM/model"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type PrivateMessage struct {
	im.Audit
	ID        uint              `gorm:"comment:主键ID;"`
	UID       string            `gorm:"comment:唯一标识;size:20;uniqueIndex;not null;"`
	From      string            `gorm:"comment:发送者uid;size:30;uniqueIndex:idx_pm_ftt;not null;"`
	To        string            `gorm:"comment:接收者uid;size:30;uniqueIndex:idx_pm_ftt;;not null;"`
	timestamp int64             `gorm:"comment:消息发送时间戳;uniqueIndex:idx_pm_ftt;;not null;"`
	Type      string            `gorm:"comment:消息类型;size:10;not null;"`
	Content   datatypes.JSONMap `gorm:"comment:消息内容;null"`
	DeletedAt gorm.DeletedAt    `gorm:"comment:删除时间;null"`
}

type TextMessage struct {
	Text string `json:"text"`
}

type FileMessage struct {
	Name   string    `json:"name"`
	Size   int64     `json:"size"`
	Slices FileSlice `json:"slices"`
}

type FileSlice struct {
	Hash  string `json:"hash"`
	Order uint8  `json:"order"`
}
