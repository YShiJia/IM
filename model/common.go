package model

import "time"

// 节点类型
type ServeType uint

const (
	SERVE_TYPE_UNKNOW ServeType = iota
	SERVE_TYPE_EDGE
)

// Audit 审计信息
type Audit struct {
	Creator   string    `gorm:"comment:创建人;size:30;not null;"`
	CreatedAt time.Time `gorm:"comment:创建时间;not null;"`
	Updater   string    `gorm:"comment:更新人;size:30;not null;"`
	UpdatedAt time.Time `gorm:"comment:更新时间;not null;"`
}
