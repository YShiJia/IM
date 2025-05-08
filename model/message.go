/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2025-04-29 18:43:53
 */

package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type PrivateMessage struct {
	ID      uint   `gorm:"comment:主键ID;"`
	FromUID string `gorm:"comment:发送者uid;size:20;not null;"`
	ToUID   string `gorm:"comment:接收者uid;size:20;not null;"`
	// 联合uid用于查找聊天记录
	UnionUID  string            `gorm:"comment:发送接收方联合uid，字典序小前大后;size:40;index;not null;"`
	SendTime  int64             `gorm:"comment:用户发送消息时间戳,用于前端显示时间;not null;"`
	Type      uint8             `gorm:"comment:消息类型;not null;"`
	Content   datatypes.JSONMap `gorm:"comment:消息内容;not null"`
	DeletedAt gorm.DeletedAt    `gorm:"comment:删除时间;null"`
}

type GroupMessage struct {
	ID        uint              `gorm:"comment:主键ID;"`
	FromUID   string            `gorm:"comment:发送者uid;size:20;not null;"`
	GroupUID  string            `gorm:"comment:接收群uid;size:20;index;not null;"`
	SendTime  int64             `gorm:"comment:用户发送消息时间戳,用于前端显示时间;not null;"`
	Type      uint8             `gorm:"comment:消息类型;not null;"`
	Content   datatypes.JSONMap `gorm:"comment:消息内容;not null"`
	DeletedAt gorm.DeletedAt    `gorm:"comment:删除时间;null"`
}

type TextContent struct {
	Text string `json:"text"`
}

// ImageContent 图片文件
type ImageContent struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
	Url  string `json:"url"`
}

// FileContent 文件
type FileContent struct {
	Name   string    `json:"name"`
	Size   int64     `json:"size"`
	Slices FileSlice `json:"slices"`
}

type ContentFileSlice struct {
	Hash  string `json:"hash"`
	Order uint8  `json:"order"`
}
