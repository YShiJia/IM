/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-16 18:01:58
 */

package dbfriend

import (
	"errors"
	"fmt"
	"github.com/YShiJia/IM/database/dbuser"
	"gorm.io/gorm"
	"time"
)

type FriendDB interface {
	GetFriendList(userSocialId string) ([]*dbuser.User, error)
	MakeFriend(FromSocialId string, ToSocialId string) error
	DeleteFriend(FromSocialId string, ToSocialId string) error
}

var _ FriendDB = (*friendDbByGorm)(nil)

type friendDbByGorm struct {
	db        *gorm.DB
	tableName string
}

/*
select *
from im_user
where id in (select friend_id
             from im_friend
             where user_id = (select id
                              from im_user
                              where social_id = 'xxx'))
*/

func (f *friendDbByGorm) GetFriendList(userSocialId string) ([]*dbuser.User, error) {
	var users []*dbuser.User
	//查询出user的id
	userid := f.db.Table(dbuser.UserTableName).Select("id").Where("social_id = ?", userSocialId)
	// 查询出user的所有friend_ids
	friendids := f.db.Table(FriendTableName).Select("friend_id").Where("user_id = ?", userid)
	// 根据friend_ids查询出所有friend信息
	err := f.db.Table(dbuser.UserTableName).Select("*").Where("id in (?)", friendids).Find(&users).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("found user failed, err %w", err)
	}
	return users, nil
}

func (f *friendDbByGorm) findUserIdsBySocialIds(socialIds []string) ([]int, error) {
	var userIds []int
	err := f.db.Table(dbuser.UserTableName).Select("id").Where("social_id in ?", socialIds).Find(&userIds).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("found user failed, err %w", err)
	}
	return userIds, nil
}

func (f *friendDbByGorm) MakeFriend(FromSocialId string, ToSocialId string) error {
	userIds, err := f.findUserIdsBySocialIds([]string{FromSocialId, ToSocialId})
	if err != nil {
		return err
	}
	if len(userIds) != 2 {
		return fmt.Errorf("user not found")
	}

	//插入两个记录
	friends := []*Friend{
		{
			UserId:    userIds[0],
			FriendId:  userIds[1],
			CreatedAt: time.Now().Unix(),
		},
		{
			UserId:    userIds[1],
			FriendId:  userIds[0],
			CreatedAt: time.Now().Unix(),
		},
	}
	if err = f.db.Table(FriendTableName).Create(friends).Error; err != nil {
		return fmt.Errorf("make friend failed")
	}
	return nil
}

/*
delete
from im_friend
where (user_id = 1xxx and friend_id = 2xxx)
   or (user_id = 2xxx and friend_id = 1xxx);
*/

func (f *friendDbByGorm) DeleteFriend(FromSocialId string, ToSocialId string) error {
	userIds, err := f.findUserIdsBySocialIds([]string{FromSocialId, ToSocialId})
	if err != nil {
		return err
	}
	if len(userIds) != 2 {
		return fmt.Errorf("user not found")
	}
	err = f.db.Table(FriendTableName).Delete(&Friend{}).Where("(user_id = ? and friend_id = ?) or (user_id = ? and friend_id = ?)",
		userIds[0], userIds[1], userIds[1], userIds[0]).Error
	if err != nil {
		return fmt.Errorf("delete friend failed, err %w", err)
	}
	return nil
}

func NewFriendDbByGorm(db *gorm.DB) *friendDbByGorm {
	return &friendDbByGorm{db: db, tableName: FriendTableName}
}
