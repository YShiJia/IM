/**
 * @author ysj
 * @email 2239831438@qq.com
 * @createTime: 2025-04-19 19:01:15
 */

package db

import (
	"fmt"
	"github.com/YShiJia/IM/apps/message/api/model"
	"gorm.io/gorm"
)

type fileDAO struct {
	*baseDAO
}

func NewFileDAO(tx *gorm.DB) *fileDAO {
	return &fileDAO{baseDAO: NewBaseDAO(tx)}
}

var File = &fileDAO{}

func (d *fileDAO) GetByHash(hash string) (*model.File, error) {
	if hash == "" {
		return nil, fmt.Errorf("invalid hash")
	}
	f := &model.File{}
	if err := d.getTx().Model(&model.File{}).Where(&model.File{Hash: hash}).First(f).Error; err != nil {
		return nil, err
	}
	return f, nil
}

func (d *fileDAO) Create(file *model.File) (*model.File, error) {
	if err := d.getTx().Model(&model.File{}).Create(file).Error; err != nil {
		return nil, err
	}
	return file, nil
}
