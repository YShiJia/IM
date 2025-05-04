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
	UID          string         `gorm:"comment:群UID;size:20;uniqueIndex:idx_gp_u;not null;"`
	Name         string         `gorm:"comment:群名称;size:20;not null;"`
	Owner        string         `gorm:"comment:群主UID;size:30;not null;"`
	Creator      string         `gorm:"comment:创建人UID;size:30;not null;"`
	CreatedAt    time.Time      `gorm:"comment:创建时间;not null;"`
	DeletedAt    gorm.DeletedAt `gorm:"comment:删除时间;null"`
	GroupMembers []GroupMember
}

type GroupMember struct {
	GroupID uint   `gorm:"comment:群ID;uniqueIndex:idx_gm_gu;not null;"`
	UserID  uint   `gorm:"comment:用户ID;uniqueIndex:idx_gm_gu;index;not null;"`
	Role    string `gorm:"comment:角色;not null;"`
	Group   *Group
	User    *User
}
