/**
 * @author ysj
 * @email 2239831438@qq.com
 * @createTime: 2025-04-19 18:13:41
 */

package db

import (
	"fmt"
	conf "github.com/YShiJia/IM/apps/message/api/internal/config"
	"github.com/YShiJia/IM/apps/message/api/model"
	"github.com/YShiJia/IM/database"
	"gorm.io/gorm"
)

var (
	IMDB *gorm.DB
)

type baseDAO struct {
	tx *gorm.DB
}

// NewBaseDAO 获取新的baseDAO
func NewBaseDAO(tx *gorm.DB) *baseDAO {
	return &baseDAO{tx: tx}
}

func (b *baseDAO) getTx() *gorm.DB {
	if (b == nil || b.tx == nil) && IMDB != nil {
		return IMDB
	}
	return b.tx
}

// InitIMDB 初始化业务数据库
func InitIMDB() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		conf.Conf.MysqlConf.Username,
		conf.Conf.MysqlConf.Password,
		conf.Conf.MysqlConf.Host,
		conf.Conf.MysqlConf.Port,
		conf.Conf.MysqlConf.Database,
		conf.Conf.MysqlConf.Config)
	IMDB, err = database.InitDB(dsn, model.GetModels())
	if err != nil {
		return fmt.Errorf("init im db error: %w", err)
	}
	return nil
}

func InitTestCoreDB() (err error) {
	dsn := fmt.Sprintf("file::memory:?name=%s", conf.Conf.MysqlConf.Database)
	IMDB, err = database.InitTestDB(dsn, model.GetModels())
	return err
}
