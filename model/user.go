/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2025-04-29 18:52:42
 */

package model

type User struct {
	ID     uint   `gorm:"comment:主键ID;"`
	UID    string `gorm:"comment:用户UID;size:20;uniqueIndex;not null;"`
	Age    uint   `gorm:"comment:年龄;size:3;not null;"`
	Name   string `gorm:"comment:用户名;size:20;not null;"`
	Avatar string `gorm:"comment:头像;size:100;not null;"`
	Gender bool   `gorm:"comment:性别;size:1;not null;"`
	Email  string `gorm:"comment:邮箱;size:30;null;"`
	Sign   string `gorm:"comment:个性签名;size:100;null;"`
	Source string `gorm:"comment:来源;size:20;not null;"`
}

type FriendGroup struct {
	ID     uint   `gorm:"comment:主键ID;"`
	UserID uint   `gorm:"comment:创建者ID;index;not null"`
	Name   string `gorm:"comment:分组名;size:20;not null;"`
	User   *User
}

type Friend struct {
	ID            uint `gorm:"comment:主键ID;"`
	UserID        uint `gorm:"comment:主用户id;uniqueIndex:idx_fd_ugf;not null;"`
	FriendGroupID uint `gorm:"comment:分组id;uniqueIndex:idx_fd_ugf;not null;"`
	FriendID      uint `gorm:"comment:友用户id;uniqueIndex:idx_fd_ugf;not null;"`
	User          *User
	FriendGroup   *FriendGroup
	Friend        *User
}
