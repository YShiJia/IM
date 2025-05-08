/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2025-05-05 18:34:16
 */

package db

import (
	"github.com/YShiJia/IM/model"
	"gorm.io/gorm"
)

type MessageGroupDAO struct {
	*baseDAO
}

func NewMessageGroupDAO(tx *gorm.DB) *MessageGroupDAO {
	return &MessageGroupDAO{baseDAO: NewBaseDAO(tx)}
}

var MessageGroup = &MessageGroupDAO{}

func (m *MessageGroupDAO) Create(groupMessage *model.GroupMessage) error {
	return m.tx.Create(groupMessage).Error
}
