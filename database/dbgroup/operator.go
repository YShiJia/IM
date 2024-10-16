/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-16 21:30:05
 */

package dbgroup

import (
	"errors"
	"fmt"
	"github.com/YShiJia/IM/database/dbuser"
	"gorm.io/gorm"
	"time"
)

type GroupDB interface {
	FindGroupBySocialId(groupSocialId string) (*Group, error)
	FindGroupMembersBySocialId(groupSocialId string) ([]*dbuser.User, error)
	JoinGroup(userSocialId string, groupSocialId string) error
	QuitGroup(userSocialId string, groupSocialId string) error
}

var _ GroupDB = (*groupDbByGorm)(nil)

type groupDbByGorm struct {
	db *gorm.DB
}

/*
select id
from im_group
where social_id = 'xxx';
*/
func (g *groupDbByGorm) findGroupIdBySocialId(groupSocialId string) (int, error) {
	var groupId int
	err := g.db.Table(GroupTableName).Select("id").Where("social_id = ?", groupSocialId).Find(&groupId).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, fmt.Errorf("found groupid failed, err %w", err)
	}
	return groupId, nil
}

func (g *groupDbByGorm) findUserIdBySocialId(userSocialId string) (int, error) {
	var userId int
	err := g.db.Table(dbuser.UserTableName).Select("id").Where("social_id = ?", userSocialId).Find(&userId).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, fmt.Errorf("found userid failed, err %w", err)
	}
	return userId, nil
}

/*
select *
from im_group
where social_id = 'xxx';
*/

func (g *groupDbByGorm) FindGroupBySocialId(groupSocialId string) (*Group, error) {
	var group Group
	err := g.db.Table(GroupTableName).Where("social_id = ?", groupSocialId).Find(&group).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("group not found")
		}
		return nil, fmt.Errorf("found group failed, err %w", err)
	}
	return &group, nil
}

/*
select *
from im_user
where id in (select id

	from im_group_member
	where group_id = 'xxx');
*/
func (g *groupDbByGorm) FindGroupMembersBySocialId(groupSocialId string) ([]*dbuser.User, error) {
	// 1. 先查询group_id
	groupId, err := g.findGroupIdBySocialId(groupSocialId)
	if err != nil {
		return nil, err
	}
	// 2.再查询对应的users
	memberIds := g.db.Table(GroupMemberTableName).Where("group_id = ?", groupId)
	var users []*dbuser.User
	err = g.db.Table(dbuser.UserTableName).Where("id in (?)", memberIds).Find(&users).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 无成员数据，返回空数组
			return users, nil
		}
		return nil, fmt.Errorf("found group failed, err %w", err)
	}
	return users, nil
}

func (g *groupDbByGorm) JoinGroup(userSocialId string, groupSocialId string) error {
	// 1. 先查询group_id
	groupId, err := g.findGroupIdBySocialId(groupSocialId)
	if err != nil {
		return err
	}
	// 2. 查询对应的user_id
	userId, err := g.findUserIdBySocialId(userSocialId)
	if err != nil {
		return err
	}

	err = g.db.Table(GroupMemberTableName).Create(&GroupMember{
		GroupId:   groupId,
		MemberId:  userId,
		CreatedAt: time.Now().Unix(),
	}).Error
	if err != nil {
		return fmt.Errorf("join group failed, err %w", err)
	}
	return nil
}

func (g *groupDbByGorm) QuitGroup(userSocialId string, groupSocialId string) error {
	// 1. 先查询group_id
	groupId, err := g.findGroupIdBySocialId(groupSocialId)
	if err != nil {
		return err
	}
	// 2. 查询对应的user_id
	userId, err := g.findUserIdBySocialId(userSocialId)
	if err != nil {
		return err
	}
	err = g.db.Table(GroupMemberTableName).Where("group_id = ? and member_id = ?", groupId, userId).Delete(&GroupMember{}).Error
	if err != nil {
		return fmt.Errorf("quit from group failed, err %w", err)
	}
	return nil
}

func NewGroupDBDbByGorm(db *gorm.DB) *groupDbByGorm {
	return &groupDbByGorm{db: db}
}
