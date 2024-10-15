/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-07 22:07:47
 */

package dbuser

import (
	"errors"
	"gorm.io/gorm"
)

type UserDB interface {
	GetIdBySocialId(socialIds []string) ([]int, error)
	GetSocialIdById(ids []int) ([]string, error)
	FindUserByEmail(email string) (*User, error)
	CreateUser(user *User) error
	GetUserBySocialId(socialId string) (*User, error)
}

var _ UserDB = (*userDbByGorm)(nil)

type userDbByGorm struct {
	db        *gorm.DB
	tableName string
}

func (u *userDbByGorm) GetUserBySocialId(socialId string) (*User, error) {
	var user User
	err := u.db.Model(&User{}).Select("*").Where("social_id = ?", socialId).Find(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &user, nil
}

func (u *userDbByGorm) CreateUser(user *User) error {
	return u.db.Model(&User{}).Create(user).Error
}

func (u *userDbByGorm) FindUserByEmail(email string) (*User, error) {
	var user User
	err := u.db.Model(&User{}).Select("*").Where("email = ?", email).Find(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &user, nil
}

func (u *userDbByGorm) GetIdBySocialId(socialIds []string) ([]int, error) {
	var ids []int
	err := u.db.Model(&User{}).Select("id").Where("social_id in (?)", socialIds).Find(&ids).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return ids, nil
}

func (u *userDbByGorm) GetSocialIdById(ids []int) ([]string, error) {
	var socialIds []string
	err := u.db.Model(&User{}).Select("social_id").Where("id in (?)", ids).Find(&socialIds).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return socialIds, nil
}

func NewUserDbByGorm(db *gorm.DB) *userDbByGorm {
	return &userDbByGorm{db: db, tableName: UserTableName}
}
