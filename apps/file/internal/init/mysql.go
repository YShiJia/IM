/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2025-04-29 11:59:28
 */

package init

import (
	"fmt"
	conf "github.com/YShiJia/IM/apps/file/internal/config"
	"github.com/YShiJia/IM/apps/file/internal/dao/db"
	"github.com/YShiJia/IM/database"
	"github.com/YShiJia/IM/model"
)

// InitIMDB 初始化业务数据库
func InitIMDB() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		conf.Conf.MysqlConf.Username,
		conf.Conf.MysqlConf.Password,
		conf.Conf.MysqlConf.Host,
		conf.Conf.MysqlConf.Port,
		conf.Conf.MysqlConf.Database,
		conf.Conf.MysqlConf.Config)
	db.IMDB, err = database.InitDB(dsn, model.GetFileModels())
	if err != nil {
		return fmt.Errorf("init im db error: %w", err)
	}
	return nil
}

func InitTestIMDB() (err error) {
	dsn := fmt.Sprintf("file::memory:?name=%s", conf.Conf.MysqlConf.Database)
	db.IMDB, err = database.InitTestDB(dsn, model.GetMessageModels())
	return err
}
