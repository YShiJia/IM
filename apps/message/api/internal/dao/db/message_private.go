/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2025-05-05 16:59:39
 */

package db

import (
	"github.com/YShiJia/IM/model"
	"gorm.io/gorm"
)

type MessagePrivateDAO struct {
	*baseDAO
}

func NewMessagePrivateDAO(tx *gorm.DB) *MessagePrivateDAO {
	return &MessagePrivateDAO{baseDAO: NewBaseDAO(tx)}
}

var MessagePrivate = &MessagePrivateDAO{}

func (m *MessagePrivateDAO) Create(messagePrivate *model.PrivateMessage) error {
	return m.tx.Create(messagePrivate).Error
}
