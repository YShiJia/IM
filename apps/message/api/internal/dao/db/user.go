/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2025-05-05 17:35:10
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

func (d *UserDAO) GetByUID(uid string) (*model.User, error) {
	user := &model.User{}
	if err := d.getTx().Model(&model.User{}).Where(model.User{
		UID: uid,
	}).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
