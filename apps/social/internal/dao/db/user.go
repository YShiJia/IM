/**
 * @author ysj
 * @email 2239831438@qq.com
 * @createTime: 2025-04-19 19:01:15
 */

package db

import (
	"github.com/YShiJia/IM/model"
	"gorm.io/gorm"
)

type UserDAO struct {
	*baseDAO
}

func NewUserDAO(tx *gorm.DB) *UserDAO {
	return &UserDAO{baseDAO: NewBaseDAO(tx)}
}

var User = &UserDAO{}

func (d *UserDAO) Create(user *model.User) (*model.User, error) {
	if err := d.getTx().Model(&model.User{}).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (d *UserDAO) GetByUID(uid string) (*model.User, error) {
	var user model.User

	if err := d.getTx().Model(&model.User{}).Where(&model.User{
		UID: uid,
	}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
