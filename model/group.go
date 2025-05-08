/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2025-05-03 00:22:08
 */

package model

import (
	"gorm.io/gorm"
	"time"
)

type Group struct {
	ID           uint           `gorm:"comment:主键ID;"`
	UID          string         `gorm:"comment:群UID;size:20;uniqueIndex;not null;"`
	Name         string         `gorm:"comment:群名称;size:20;index;not null;"`
	UserID       uint           `gorm:"comment:群主ID;not null;"`
	CreatorID    uint           `gorm:"comment:创建人ID;not null;"`
	CreatedAt    time.Time      `gorm:"comment:创建时间;not null;"`
	DeletedAt    gorm.DeletedAt `gorm:"comment:删除时间;null"`
	User         *User
	Creator      *User
	GroupMembers []GroupMember
}

type GroupMember struct {
	GroupID uint   `gorm:"comment:群ID;uniqueIndex:idx_gm_gu;not null;"`
	UserID  uint   `gorm:"comment:用户ID;uniqueIndex:idx_gm_gu;index;not null;"`
	Role    string `gorm:"comment:角色;size:20;not null;"`
	Group   *Group
	User    *User
}
