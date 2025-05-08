/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2025-05-05 21:26:56
 */

package db

import (
	"fmt"
	"github.com/YShiJia/IM/model"
	"gorm.io/gorm"
)

type FileSliceDAO struct {
	*baseDAO
}

func NewFileSliceDAO(tx *gorm.DB) *FileSliceDAO {
	return &FileSliceDAO{baseDAO: NewBaseDAO(tx)}
}

var FileSlice = &FileSliceDAO{}

func (d *FileSliceDAO) OrderByOrder() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order("`order`")
	}
}

func (d *FileSliceDAO) ListByFileID(fileID uint, conds ...func(*gorm.DB) *gorm.DB) ([]*model.FileSlice, error) {
	var fileSlices []*model.FileSlice
	if err := d.getTx().
		Model(&model.FileSlice{}).
		Scopes(conds...).
		Find(&fileSlices).Error; err != nil {
		return nil, err
	}
	return fileSlices, nil
}

func (d *FileSliceDAO) Create(fileSlices []*model.FileSlice) ([]*model.FileSlice, error) {
	if err := d.getTx().CreateInBatches(fileSlices, 100).Error; err != nil {
		return nil, err
	}
	return fileSlices, nil
}

func (d *FileSliceDAO) FirstOrCreate(fileSlice *model.FileSlice) (*model.FileSlice, error) {
	if fileSlice == nil {
		return nil, fmt.Errorf("fileSlice is nil")
	}
	if err := d.getTx().
		Where(&model.FileSlice{Hash: fileSlice.Hash}).
		FirstOrCreate(fileSlice).Error; err != nil {
		return nil, err
	}
	return fileSlice, nil
}

func (d *FileSliceDAO) GetByHash(hash string) (*model.FileSlice, error) {
	fileSlice := &model.FileSlice{}
	if err := d.getTx().
		Where(&model.FileSlice{Hash: hash}).
		First(fileSlice).Error; err != nil {
		return nil, err
	}
	return fileSlice, nil
}
