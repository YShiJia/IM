/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2025-05-05 18:30:11
 */

package db

import (
	"github.com/YShiJia/IM/model"
	"gorm.io/gorm"
)

type GroupDAO struct {
	*baseDAO
}

func NewGroupDAO(tx *gorm.DB) *GroupDAO {
	return &GroupDAO{baseDAO: NewBaseDAO(tx)}
}

var Group = &GroupDAO{}

func (d *GroupDAO) PreloadGroupMembers() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload("GroupMembers").
			Preload("GroupMembers.User")
	}
}

func (d *GroupDAO) GetByUID(uid string, conds ...func(*gorm.DB) *gorm.DB) (*model.Group, error) {
	group := &model.Group{}
	if err := d.getTx().Model(&model.Group{}).Scopes(conds...).
		Where(model.User{
			UID: uid,
		}).First(group).Error; err != nil {
		return nil, err
	}
	return group, nil

}
