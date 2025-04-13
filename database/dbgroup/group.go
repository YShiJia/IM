/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-16 21:29:56
 */

package dbgroup

var GroupTableName = "im_group"

type Group struct {
	Id           int    `json:"id" gorm:"column:id"`
	SocialId     string `json:"social_id" gorm:"column:social_id"`
	Name         string `json:"name" gorm:"column:name"`
	CreateUserId int    `json:"create_user_id" gorm:"column:create_user_id"`
	CreatedAt    int64  `json:"created_at" gorm:"column:created_at"`
	DeletedAt    int64  `json:"deleted_at" gorm:"column:deleted_at"`
}

func (g *Group) TableName() string {
	return "im_group"
}

var GroupMemberTableName = "im_group_member"

type GroupMember struct {
	Id        int   `json:"id" gorm:"column:id"`
	GroupId   int   `json:"group_id" gorm:"column:group_id"`
	MemberId  int   `json:"member_id" gorm:"column:member_id"`
	CreatedAt int64 `json:"created_at" gorm:"column:created_at"`
	DeletedAt int64 `json:"deleted_at" gorm:"column:deleted_at"`
}

func (g *GroupMember) TableName() string {
	return "im_group_member"
}
